package indieauth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/owncast/owncast/auth"
	ia "github.com/owncast/owncast/auth/indieauth"
	"github.com/owncast/owncast/core/user"
)

// StartAuthFlow will begin the IndieAuth flow for the current user.
func (c *Controller) StartAuthFlow(u user.User, w http.ResponseWriter, r *http.Request) {
	type request struct {
		AuthHost string `json:"authHost"`
	}

	type response struct {
		Redirect string `json:"redirect"`
	}

	var authRequest request
	p, err := io.ReadAll(r.Body)
	if err != nil {
		c.WriteSimpleResponse(w, false, err.Error())
		return
	}

	if err := json.Unmarshal(p, &authRequest); err != nil {
		c.WriteSimpleResponse(w, false, err.Error())
		return
	}

	accessToken := r.URL.Query().Get("accessToken")

	redirectURL, err := ia.StartAuthFlow(authRequest.AuthHost, u.ID, accessToken, u.DisplayName, c.Data)
	if err != nil {
		c.WriteSimpleResponse(w, false, err.Error())
		return
	}

	redirectResponse := response{
		Redirect: redirectURL.String(),
	}
	c.WriteResponse(w, redirectResponse)
}

// HandleRedirect will handle the redirect from an IndieAuth server to
// continue the auth flow.
func (c *Controller) HandleRedirect(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query().Get("state")
	code := r.URL.Query().Get("code")
	request, response, err := ia.HandleCallbackCode(code, state)
	if err != nil {
		log.Debugln(err)
		msg := `Unable to complete authentication. <a href="/">Go back.</a><hr/>`
		_ = c.WriteString(w, msg, http.StatusBadRequest)
		return
	}

	// Check if a user with this auth already exists, if so, log them in.
	if u := auth.GetUserByAuth(response.Me, auth.IndieAuth); u != nil {
		// Handle existing auth.
		log.Debugln("user with provided indieauth already exists, logging them in")

		// Update the current user's access token to point to the existing user id.
		accessToken := request.CurrentAccessToken
		userID := u.ID
		if err := user.SetAccessTokenToOwner(accessToken, userID, c.Data.Store); err != nil {
			c.WriteSimpleResponse(w, false, err.Error())
			return
		}

		if request.DisplayName != u.DisplayName {
			loginMessage := fmt.Sprintf("**%s** is now authenticated as **%s**", request.DisplayName, u.DisplayName)
			if err := c.Core.Chat.SendSystemAction(loginMessage, true); err != nil {
				log.Errorln(err)
			}
		}

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

		return
	}

	// Otherwise, save this as new auth.
	log.Debug("indieauth token does not already exist, saving it as a new one for the current user")
	if err := auth.AddAuth(request.UserID, response.Me, auth.IndieAuth); err != nil {
		c.WriteSimpleResponse(w, false, err.Error())
		return
	}

	// Update the current user's authenticated flag so we can show it in
	// the chat UI.
	if err := user.SetUserAsAuthenticated(request.UserID, c.Data.Store); err != nil {
		log.Errorln(err)
	}

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

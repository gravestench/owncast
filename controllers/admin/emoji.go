package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/owncast/owncast/config"
	"github.com/owncast/owncast/utils"
)

// UploadCustomEmoji allows POSTing a new custom emoji to the server.
func (c *Controller) UploadCustomEmoji(w http.ResponseWriter, r *http.Request) {
	if !c.requirePOST(w, r) {
		return
	}

	type postEmoji struct {
		Name string `json:"name"`
		Data string `json:"data"`
	}

	emoji := new(postEmoji)

	if err := json.NewDecoder(r.Body).Decode(emoji); err != nil {
		c.WriteSimpleResponse(w, false, err.Error())
		return
	}

	bytes, _, err := utils.DecodeBase64Image(emoji.Data)
	if err != nil {
		c.WriteSimpleResponse(w, false, err.Error())
		return
	}

	// Prevent path traversal attacks
	emojiFileName := filepath.Base(emoji.Name)
	targetPath := filepath.Join(config.CustomEmojiPath, emojiFileName)

	err = os.MkdirAll(config.CustomEmojiPath, 0o700)
	if err != nil {
		c.WriteSimpleResponse(w, false, err.Error())
		return
	}

	if utils.DoesFileExists(targetPath) {
		c.WriteSimpleResponse(w, false, fmt.Sprintf("An emoji with the name %q already exists", emojiFileName))
		return
	}

	if err = os.WriteFile(targetPath, bytes, 0o600); err != nil {
		c.WriteSimpleResponse(w, false, err.Error())
		return
	}

	c.WriteSimpleResponse(w, true, fmt.Sprintf("Emoji %q has been uploaded", emojiFileName))
}

// DeleteCustomEmoji deletes a custom emoji.
func (c *Controller) DeleteCustomEmoji(w http.ResponseWriter, r *http.Request) {
	if !c.requirePOST(w, r) {
		return
	}

	type deleteEmoji struct {
		Name string `json:"name"`
	}

	emoji := new(deleteEmoji)

	if err := json.NewDecoder(r.Body).Decode(emoji); err != nil {
		c.WriteSimpleResponse(w, false, err.Error())
		return
	}

	// var emojiFileName = filepath.Base(emoji.Name)
	targetPath := filepath.Join(config.CustomEmojiPath, emoji.Name)

	if err := os.Remove(targetPath); err != nil {
		if os.IsNotExist(err) {
			c.WriteSimpleResponse(w, false, fmt.Sprintf("Emoji %q doesn't exist", emoji.Name))
		} else {
			c.WriteSimpleResponse(w, false, err.Error())
		}
		return
	}

	c.WriteSimpleResponse(w, true, fmt.Sprintf("Emoji %q has been deleted", emoji.Name))
}

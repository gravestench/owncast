package chat

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/owncast/owncast/config"
	"github.com/owncast/owncast/core/chat/events"
	"github.com/owncast/owncast/core/user"
	"github.com/owncast/owncast/utils"
)

func (s *Server) userNameChanged(eventData chatClientEvent) {
	var receivedEvent events.NameChangeEvent
	if err := json.Unmarshal(eventData.data, &receivedEvent); err != nil {
		log.Errorln("error unmarshalling to NameChangeEvent", err)
		return
	}

	proposedUsername := receivedEvent.NewName

	// Check if name is on the blocklist
	blocklist := s.data.GetForbiddenUsernameList()

	// Names have a max length
	proposedUsername = utils.MakeSafeStringOfLength(proposedUsername, config.MaxChatDisplayNameLength)

	for _, blockedName := range blocklist {
		normalizedName := strings.TrimSpace(blockedName)
		normalizedName = strings.ToLower(normalizedName)
		if strings.Contains(normalizedName, proposedUsername) {
			// Denied.
			log.Debugln(logSanitize(eventData.client.User.DisplayName), "blocked from changing name to", logSanitize(proposedUsername), "due to blocked name", normalizedName)
			message := fmt.Sprintf("You cannot change your name to **%s**.", proposedUsername)
			s.sendActionToClient(eventData.client, message)

			// Resend the client's user so their username is in sync.
			eventData.client.sendConnectedClientInfo()

			return
		}
	}

	// Check if the name is not already assigned to a registered user.
	if available, err := user.IsDisplayNameAvailable(proposedUsername, s.data.Store); err != nil {
		log.Errorln("error checking if name is available", err)
		return
	} else if !available {
		message := fmt.Sprintf("The name **%s** has been already registered. If this is your name, please authenticate.", proposedUsername)
		s.sendActionToClient(eventData.client, message)

		// Resend the client's user so their username is in sync.
		eventData.client.sendConnectedClientInfo()

		return
	}

	savedUser := user.GetUserByToken(eventData.client.accessToken, s.data.Store)
	oldName := savedUser.DisplayName

	// Save the new name
	if err := user.ChangeUsername(eventData.client.User.ID, proposedUsername, s.data.Store); err != nil {
		log.Errorln("error changing username", err)
	}

	// Update the connected clients associated user with the new name
	now := time.Now()
	eventData.client.User = savedUser
	eventData.client.User.NameChangedAt = &now

	// Send chat event letting everyone about about the name change
	savedUser.DisplayName = proposedUsername

	broadcastEvent := events.NameChangeBroadcast{
		Oldname: oldName,
	}
	broadcastEvent.User = savedUser
	broadcastEvent.SetDefaults()
	payload := broadcastEvent.GetBroadcastPayload()
	if err := s.Broadcast(payload); err != nil {
		log.Errorln("error broadcasting NameChangeEvent", err)
		return
	}

	// Send chat user name changed webhook
	receivedEvent.User = savedUser
	receivedEvent.ClientID = eventData.client.Id
	s.webhooks.SendChatEventUsernameChanged(receivedEvent)

	// Resend the client's user so their username is in sync.
	eventData.client.sendConnectedClientInfo()
}

func (s *Server) userColorChanged(eventData chatClientEvent) {
	var receivedEvent events.ColorChangeEvent
	if err := json.Unmarshal(eventData.data, &receivedEvent); err != nil {
		log.Errorln("error unmarshalling to ColorChangeEvent", err)
		return
	}

	// Verify this color is valid
	if receivedEvent.NewColor > config.MaxUserColor {
		log.Errorln("invalid color requested when changing user display color")
		return
	}

	// Save the new color
	if err := user.ChangeUserColor(eventData.client.User.ID, receivedEvent.NewColor, s.data.Store); err != nil {
		log.Errorln("error changing user display color", err)
	}

	// Resend client's user info with new color, otherwise the name change dialog would still show the old color
	eventData.client.User.DisplayColor = receivedEvent.NewColor
	eventData.client.sendConnectedClientInfo()
}

func (s *Service) userMessageSent(eventData chatClientEvent) {
	var event events.UserMessageEvent
	if err := json.Unmarshal(eventData.data, &event); err != nil {
		log.Errorln("error unmarshalling to UserMessageEvent", err)
		return
	}

	event.SetDefaults()
	event.ClientID = eventData.client.Id

	// Ignore empty messages
	if event.Empty() {
		return
	}

	// Ignore if the stream has been offline
	if !s.getStatus().Online && s.getStatus().LastDisconnectTime != nil {
		disconnectedTime := s.getStatus().LastDisconnectTime.Time
		if time.Since(disconnectedTime) > 5*time.Minute {
			return
		}
	}

	event.User = user.GetUserByToken(eventData.client.accessToken, s.data.Store)

	// Guard against nil users
	if event.User == nil {
		return
	}

	if err := s.Broadcast(&event); err != nil {
		log.Errorln("error broadcasting UserMessageEvent payload", err)
		return
	}

	// Send chat message sent webhook
	s.webhooks.SendChatEvent(&event)
	s.chatMessagesSentCounter.Inc()

	SaveUserMessage(event)
	eventData.client.MessageCount++
	s.server.lastSeenCache[event.User.ID] = time.Now()
}

func logSanitize(userValue string) string {
	// strip carriage return and newline from user-submitted values to prevent log injection
	sanitizedValue := strings.ReplaceAll(userValue, "\n", "")
	sanitizedValue = strings.ReplaceAll(sanitizedValue, "\r", "")

	return fmt.Sprintf("userSuppliedValue(%s)", sanitizedValue)
}

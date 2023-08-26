package models

import (
	"github.com/owncast/owncast/utils"
)

type YpDetailsResponse struct {
	LastConnectTime *utils.NullTime `json:"lastConnectTime"`
	Name            string          `json:"name"`
	Description     string          `json:"description"`
	StreamTitle     string          `json:"streamTitle,omitempty"`
	Logo            string          `json:"logo"`
	Tags            []string        `json:"tags"`
	Social          []SocialHandle  `json:"social"`

	ViewerCount           int  `json:"viewerCount"`
	OverallMaxViewerCount int  `json:"overallMaxViewerCount"`
	SessionMaxViewerCount int  `json:"sessionMaxViewerCount"`
	NSFW                  bool `json:"nsfw"`
	Online                bool `json:"online"`
}

type YpPingResponse struct {
	Key       string `json:"key"`
	Error     string `json:"error"`
	ErrorCode int    `json:"errorCode"`
	Success   bool   `json:"success"`
}

type YpPingRequest struct {
	Key string `json:"key"`
	URL string `json:"url"`
}

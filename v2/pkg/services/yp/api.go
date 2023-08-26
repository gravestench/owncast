package yp

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/owncast/owncast/models"
)

// GetYPResponse gets the status of the server for Service purposes.
func (s *Service) GetYPResponse(w http.ResponseWriter, r *http.Request) {
	if !s.persistence.GetDirectoryEnabled() {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	status := s.getStatus()

	streamTitle := s.persistence.GetStreamTitle()

	response := models.YpDetailsResponse{
		Name:                  s.persistence.GetServerName(),
		Description:           s.persistence.GetServerSummary(),
		StreamTitle:           streamTitle,
		Logo:                  "/logo",
		NSFW:                  s.persistence.GetNSFW(),
		Tags:                  s.persistence.GetServerMetadataTags(),
		Online:                status.Online,
		ViewerCount:           status.ViewerCount,
		OverallMaxViewerCount: status.OverallMaxViewerCount,
		SessionMaxViewerCount: status.SessionMaxViewerCount,
		LastConnectTime:       status.LastConnectTime,
		Social:                s.persistence.GetSocialHandles(),
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Errorln(err)
	}
}

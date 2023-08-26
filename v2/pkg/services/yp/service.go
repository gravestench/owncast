package yp

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/gravestench/runtime/pkg"
	"github.com/sirupsen/logrus"

	"github.com/owncast/owncast/config"
	"github.com/owncast/owncast/models"
	"github.com/owncast/owncast/v2/pkg/services/persistence"
)

// New creates a new instance of the Service service handler.
func New(getStatusFunc func() models.Status) *Service {
	return &Service{
		getStatus: getStatusFunc,
	}
}

// Service is a service for handling listing in the Owncast directory.
type Service struct {
	timer  *time.Ticker
	logger *logrus.Logger

	getStatus    func() models.Status
	inErrorState bool

	persistence persistence.Dependency
}

func (s *Service) BindLogger(logger *logrus.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *logrus.Logger {
	return s.logger
}

func (s *Service) Init(rt pkg.IsRuntime) {
	s.Start()
}

func (s *Service) Name() string {
	return "Yellow Pages"
}

const pingInterval = 4 * time.Minute

// Start is run when a live stream begins to start pinging Service.
func (s *Service) Start() {
	s.Logger().Infoln("starting ping interval")

	s.timer = time.NewTicker(pingInterval)
	for range s.timer.C {
		s.ping()
	}

	s.ping()
}

// Stop stops the pinging of Service.
func (s *Service) Stop() {
	s.timer.Stop()
}

func (s *Service) ping() {
	if !s.persistence.GetDirectoryEnabled() {
		return
	}

	// Hack: Don't allow ping'ing when offline.
	// It shouldn't even be trying to, but on some instances the ping timer isn't stopping.
	if !s.getStatus().Online {
		return
	}

	myInstanceURL := s.persistence.GetServerURL()
	if myInstanceURL == "" {
		s.Logger().Warnln("Server URL not set in the configuration. Directory access is disabled until this is set.")
		return
	}
	isValidInstanceURL := isURL(myInstanceURL)
	if myInstanceURL == "" || !isValidInstanceURL {
		if !s.inErrorState {
			s.Logger().Warnln("Service Error: unable to use", myInstanceURL, "as a public instance URL. Fix this value in your configuration.")
		}
		s.inErrorState = true
		return
	}

	key := s.persistence.GetDirectoryRegistrationKey()

	s.Logger().Traceln("Pinging Service as: ", s.persistence.GetServerName(), "with key", key)

	request := models.YpPingRequest{
		Key: key,
		URL: myInstanceURL,
	}

	req, err := json.Marshal(request)
	if err != nil {
		s.Logger().Errorln(err)
		return
	}

	pingURL := config.GetDefaults().YPServer + "/api/ping"
	resp, err := http.Post(pingURL, "application/json", bytes.NewBuffer(req)) //nolint
	if err != nil {
		s.Logger().Errorln(err)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		s.Logger().Errorln(err)
	}

	pingResponse := models.YpPingResponse{}
	if err := json.Unmarshal(body, &pingResponse); err != nil {
		s.Logger().Errorln(err)
	}

	if !pingResponse.Success {
		if !s.inErrorState {
			s.Logger().Warnln("Service Ping error returned from service:", pingResponse.Error)
		}

		s.inErrorState = true

		return
	}

	s.inErrorState = false

	if pingResponse.Key == key {
		return
	}

	err = s.persistence.SetDirectoryRegistrationKey(key)
	if err != nil {
		s.Logger().Errorln("unable to save directory key:", err)
	}
}

func isURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

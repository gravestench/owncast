package persistence

import (
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/owncast/owncast/models"
)

const (
	datastoreValuesVersion   = 3
	datastoreValueVersionKey = "DATA_STORE_VERSION"
)

func (s *Service) migrateDatastoreValues(datastore *Service) {
	currentVersion, _ := s.GetNumber(datastoreValueVersionKey)
	if currentVersion == 0 {
		currentVersion = datastoreValuesVersion
	}

	for v := currentVersion; v < datastoreValuesVersion; v++ {
		log.Infof("Migration datastore values from %d to %d\n", int(v), int(v+1))
		switch v {
		case 0:
			s.migrateToDatastoreValues1(datastore)
		case 1:
			s.migrateToDatastoreValues2(datastore)
		case 2:
			s.migrateToDatastoreValues3ServingEndpoint3(datastore)
		default:
			log.Fatalln("missing datastore values migration step")
		}
	}
	if err := s.SetNumber(datastoreValueVersionKey, datastoreValuesVersion); err != nil {
		log.Errorln("error setting datastore value version:", err)
	}
}

func (s *Service) migrateToDatastoreValues1(datastore *Service) {
	// Migrate the forbidden usernames to be a slice instead of a string.
	forbiddenUsernamesString, _ := datastore.GetString(blockedUsernamesKey)
	if forbiddenUsernamesString != "" {
		forbiddenUsernamesSlice := strings.Split(forbiddenUsernamesString, ",")
		if err := datastore.SetStringSlice(blockedUsernamesKey, forbiddenUsernamesSlice); err != nil {
			log.Errorln("error migrating blocked username list:", err)
		}
	}

	// Migrate the suggested usernames to be a slice instead of a string.
	suggestedUsernamesString, _ := datastore.GetString(suggestedUsernamesKey)
	if suggestedUsernamesString != "" {
		suggestedUsernamesSlice := strings.Split(suggestedUsernamesString, ",")
		if err := datastore.SetStringSlice(suggestedUsernamesKey, suggestedUsernamesSlice); err != nil {
			log.Errorln("error migrating suggested username list:", err)
		}
	}
}

func (s *Service) migrateToDatastoreValues2(datastore *Service) {
	oldAdminPassword, _ := datastore.GetString("stream_key")
	_ = s.SetAdminPassword(oldAdminPassword)
	_ = s.SetStreamKeys([]models.StreamKey{
		{Key: oldAdminPassword, Comment: "Default stream key"},
	})
}

func (s *Service) migrateToDatastoreValues3ServingEndpoint3(_ *Service) {
	s3Config := s.GetS3Config()

	if !s3Config.Enabled {
		return
	}

	_ = s.SetVideoServingEndpoint(s3Config.ServingEndpoint)
}

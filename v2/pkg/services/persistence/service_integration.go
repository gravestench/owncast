package persistence

import (
	"time"

	"github.com/gravestench/runtime"

	"github.com/owncast/owncast/models"
	"github.com/owncast/owncast/utils"
	"github.com/owncast/owncast/v2/pkg/services/common_logging"
)

var (
	_ runtime.Service          = &Service{}
	_ runtime.HasDependencies  = &Service{}
	_ common_logging.HasLogger = &Service{}
	_ IsDatastore              = &Service{}
)

type Dependency = IsDatastore

type IsDatastore interface {
	WarmCache()
	Ready() bool // after cache is warmed

	GetCachedValue(key string) ([]byte, error)
	SetCachedValue(key string, b []byte)

	GetStringSlice(key string) ([]string, error)
	SetStringSlice(key string, value []string) error
	GetString(key string) (string, error)
	SetString(key string, value string) error
	GetNumber(key string) (float64, error)
	SetNumber(key string, value float64) error
	GetBool(key string) (bool, error)
	SetBool(key string, value bool) error
	GetStringMap(key string) (map[string]string, error)
	SetStringMap(key string, value map[string]string) error

	IsOwncastConfigManager
}

type IsOwncastConfigManager interface {
	GetExtraPageBodyContent() string
	SetExtraPageBodyContent(content string) error
	GetStreamTitle() string
	SetStreamTitle(title string) error
	GetAdminPassword() string
	SetAdminPassword(key string) error
	GetLogoPath() string
	SetLogoPath(logo string) error
	SetLogoUniquenessString(uniqueness string) error
	GetLogoUniquenessString() string
	GetServerSummary() string
	SetServerSummary(summary string) error
	GetServerWelcomeMessage() string
	SetServerWelcomeMessage(welcomeMessage string) error
	GetServerName() string
	SetServerName(name string) error
	GetServerURL() string
	SetServerURL(url string) error
	GetHTTPPortNumber() int
	SetWebsocketOverrideHost(host string) error
	GetWebsocketOverrideHost() string
	SetHTTPPortNumber(port float64) error
	GetHTTPListenAddress() string
	SetHTTPListenAddress(address string) error
	GetRTMPPortNumber() int
	SetRTMPPortNumber(port float64) error
	GetServerMetadataTags() []string
	SetServerMetadataTags(tags []string) error
	GetDirectoryEnabled() bool
	SetDirectoryEnabled(enabled bool) error
	SetDirectoryRegistrationKey(key string) error
	GetDirectoryRegistrationKey() string
	GetSocialHandles() []models.SocialHandle
	SetSocialHandles(socialHandles []models.SocialHandle) error
	GetPeakSessionViewerCount() int
	SetPeakSessionViewerCount(count int) error
	GetPeakOverallViewerCount() int
	SetPeakOverallViewerCount(count int) error
	GetLastDisconnectTime() (*utils.NullTime, error)
	SetLastDisconnectTime(disconnectTime time.Time) error
	SetNSFW(isNSFW bool) error
	GetNSFW() bool
	SetFfmpegPath(path string) error
	GetFfMpegPath() string
	GetS3Config() models.S3
	SetS3Config(config models.S3) error
	GetStreamLatencyLevel() models.LatencyLevel
	SetStreamLatencyLevel(level float64) error
	GetStreamOutputVariants() []models.StreamOutputVariant
	SetStreamOutputVariants(variants []models.StreamOutputVariant) error
	SetChatDisabled(disabled bool) error
	GetChatDisabled() bool
	SetChatEstablishedUsersOnlyMode(enabled bool) error
	GetChatEstbalishedUsersOnlyMode() bool
	GetExternalActions() []models.ExternalAction
	SetExternalActions(actions []models.ExternalAction) error
	SetCustomStyles(styles string) error
	GetCustomStyles() string
	SetCustomJavascript(styles string) error
	GetCustomJavascript() string
	SetVideoCodec(codec string) error
	GetVideoCodec() string
	VerifySettings() error
	FindHighestVideoQualityIndex(qualities []models.StreamOutputVariant) int
	GetForbiddenUsernameList() []string
	SetForbiddenUsernameList(usernames []string) error
	GetSuggestedUsernamesList() []string
	SetSuggestedUsernamesList(usernames []string) error
	GetServerInitTime() (*utils.NullTime, error)
	SetServerInitTime(t time.Time) error
	SetFederationEnabled(enabled bool) error
	GetFederationEnabled() bool
	SetFederationUsername(username string) error
	GetFederationUsername() string
	SetFederationGoLiveMessage(message string) error
	GetFederationGoLiveMessage() string
	SetFederationIsPrivate(isPrivate bool) error
	GetFederationIsPrivate() bool
	SetFederationShowEngagement(showEngagement bool) error
	GetFederationShowEngagement() bool
	SetBlockedFederatedDomains(domains []string) error
	GetBlockedFederatedDomains() []string
	SetChatJoinMessagesEnabled(enabled bool) error
	GetChatJoinMessagesEnabled() bool
	SetNotificationsEnabled(enabled bool) error
	GetNotificationsEnabled() bool
	GetDiscordConfig() models.DiscordConfiguration
	SetDiscordConfig(config models.DiscordConfiguration) error
	GetBrowserPushConfig() models.BrowserNotificationConfiguration
	SetBrowserPushConfig(config models.BrowserNotificationConfiguration) error
	SetBrowserPushPublicKey(key string) error
	GetBrowserPushPublicKey() (string, error)
	SetBrowserPushPrivateKey(key string) error
	GetBrowserPushPrivateKey() (string, error)
	SetHasPerformedInitialNotificationsConfig(hasConfigured bool) error
	GetHasPerformedInitialNotificationsConfig() bool
	GetHideViewerCount() bool
	SetHideViewerCount(hide bool) error
	GetCustomOfflineMessage() string
	SetCustomOfflineMessage(message string) error
	SetCustomColorVariableValues(variables map[string]string) error
	GetCustomColorVariableValues() map[string]string
	GetStreamKeys() []models.StreamKey
	SetStreamKeys(actions []models.StreamKey) error
	SetDisableSearchIndexing(disableSearchIndexing bool) error
	GetDisableSearchIndexing() bool
	GetVideoServingEndpoint() string
	SetVideoServingEndpoint(message string) error

	GetFederatedInboxMap() map[string]string
	GetDefaultFederationUsername() string

	GetPublicKey() string
	SetPublicKey(key string) error
	GetPrivateKey() string
	SetPrivateKey(key string) error

	GetMessagesCount() int64
	BanIPAddress(address, note string) error
	IsIPAddressBanned(address string) (bool, error)
	GetIPAddressBans() ([]models.IPAddress, error)
	RemoveIPAddressBan(address string) error
}

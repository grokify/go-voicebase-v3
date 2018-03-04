package clientutil

import (
	"github.com/grokify/go-voicebase-v3/client"
	"github.com/grokify/oauth2more"
)

const (
	ApiUrlV3Media = "https://apis.voicebase.com/v3/media"
)

func NewApiClientToken(token string) *voicebase.APIClient {
	apiConfig := voicebase.NewConfiguration()
	apiConfig.HTTPClient = oauth2more.NewClientAccessToken(token)
	return voicebase.NewAPIClient(apiConfig)
}

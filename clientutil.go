package clientutil

import (
	"github.com/grokify/go-voicebase-v3/client"
	"github.com/grokify/oauth2more"
)

func NewApiClientToken(token string) *voicebase.APIClient {
	apiConfig := voicebase.NewConfiguration()
	apiConfig.HTTPClient = oauth2more.NewClientAccessToken(token)
	return voicebase.NewAPIClient(apiConfig)
}

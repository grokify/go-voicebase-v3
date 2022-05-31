package clientutil

import (
	voicebase "github.com/grokify/go-voicebase-v3/client"
	"github.com/grokify/goauth"
)

const (
	ApiUrlV3Media = "https://apis.voicebase.com/v3/media"
)

func NewApiClientToken(token string) *voicebase.APIClient {
	apiConfig := voicebase.NewConfiguration()
	apiConfig.HTTPClient = goauth.NewClientToken(goauth.TokenBearer, token, false)
	return voicebase.NewAPIClient(apiConfig)
}

package clientutil

import (
	"github.com/grokify/goauth/authutil"

	voicebase "github.com/grokify/go-voicebase-v3/v3/client"
)

const (
	ApiUrlV3Media = "https://apis.voicebase.com/v3/media"
)

func NewApiClientToken(token string) *voicebase.APIClient {
	apiConfig := voicebase.NewConfiguration()
	apiConfig.HTTPClient = authutil.NewClientToken(authutil.TokenBearer, token, false)
	return voicebase.NewAPIClient(apiConfig)
}

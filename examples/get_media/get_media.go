package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/grokify/go-voicebase-v3"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/oauth2more"
	"github.com/joho/godotenv"
)

func LoadEnv() error {
	if len(strings.TrimSpace(os.Getenv("ENV_PATH"))) > 0 {
		return godotenv.Load(os.Getenv("ENV_PATH"))
	}
	return godotenv.Load()
}

func main() {
	err := LoadEnv()
	if err != nil {
		panic(err)
	}

	apiConfig := voicebase.NewConfiguration()
	apiConfig.HTTPClient = oauth2more.NewClientAccessToken(
		os.Getenv("VOICEBASE_BEARER_TOKEN"),
	)
	apiClient := voicebase.NewAPIClient(apiConfig)

	mediaResponse, resp, err := apiClient.MediaApi.MediaQuery(
		context.Background(),
		map[string]interface{}{},
	)
	if err != nil {
		panic(err)
	} else if resp.StatusCode > 299 {
		panic(fmt.Errorf(
			"VoiceBase media query returned status %v",
			resp.StatusCode,
		))
	}

	fmtutil.PrintJSON(mediaResponse)

	fmt.Println("DONE")
}

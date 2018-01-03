package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/grokify/go-voicebase-v3"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/gotilla/net/httputilmore"
	"github.com/grokify/oauth2more"
	"github.com/joho/godotenv"
)

const (
	EnvAccessToken = "VOICEBASE_BEARER_TOKEN"
)

func LoadEnv() error {
	if len(strings.TrimSpace(os.Getenv("ENV_PATH"))) > 0 {
		return godotenv.Load(os.Getenv("ENV_PATH"))
	}
	return godotenv.Load()
}

func getMediaHttpClient(client *http.Client) (*voicebase.VbMediaQueryResponse, error) {
	resp, err := client.Get("https://apis.voicebase.com/v3/media")
	if err != nil {
		return nil, err
	}

	err = httputilmore.PrintResponse(resp, true)
	if err != nil {
		return nil, err
	}

	mediaResponse := &voicebase.VbMediaQueryResponse{}

	err = httputilmore.UnmarshalResponseJSON(resp, mediaResponse)

	return mediaResponse, err
}

func getMediaApiClient(client *http.Client) (*voicebase.VbMediaQueryResponse, error) {
	apiConfig := voicebase.NewConfiguration()
	apiConfig.HTTPClient = client

	apiClient := voicebase.NewAPIClient(apiConfig)

	mediaResponse, resp, err := apiClient.MediaApi.MediaQuery(context.Background(), map[string]interface{}{})
	if err != nil {
		return nil, err
	} else if resp.StatusCode > 299 {
		return nil, fmt.Errorf("VoiceBase media query returned status %v", resp.StatusCode)
	}
	return &mediaResponse, nil
}

func main() {
	err := LoadEnv()
	if err != nil {
		panic(err)
	}

	client := oauth2more.NewClientAccessToken(os.Getenv(EnvAccessToken))

	useApiClient := true

	if useApiClient {
		mediaResponse, err := getMediaApiClient(client)
		if err != nil {
			panic(err)
		}
		fmtutil.PrintJSON(mediaResponse)
	} else {
		mediaResponse, err := getMediaHttpClient(client)
		if err != nil {
			panic(err)
		}
		fmtutil.PrintJSON(mediaResponse)
	}

	fmt.Println("DONE")
}

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/grokify/go-voicebase-v3"
	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"
)

const (
	EnvAccessToken = "VOICEBASE_BEARER_TOKEN"
	EnvMediaId     = "VOICEBASE_TRANSCRIPT_MEDIA_ID"
)

func main() {
	err := config.LoadDotEnvSkipEmpty(os.Getenv("ENV_PATH"), "./.env")
	if err != nil {
		panic(err)
	}

	apiClient := clientutil.NewApiClientToken(os.Getenv("VOICEBASE_BEARER_TOKEN"))

	mediaId := os.Getenv(EnvMediaId)

	vbt, resp, err := apiClient.MediaApi.GetTranscript(context.Background(), mediaId, map[string]interface{}{})
	if err != nil {
		panic(err)
	} else if resp.StatusCode >= 300 {
		panic(fmt.Sprintf("API Error %v", resp.StatusCode))
	}

	fmtutil.PrintJSON(vbt)

	fmt.Println("DONE")
}

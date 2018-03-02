package main

import (
	"context"
	"fmt"
	"os"

	"github.com/grokify/go-voicebase-v3"
	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"
)

func main() {
	err := config.LoadDotEnvSkipEmpty(os.Getenv("ENV_PATH"), "./.env")
	if err != nil {
		panic(err)
	}

	apiClient := clientutil.NewApiClientToken(os.Getenv("VOICEBASE_BEARER_TOKEN"))

	mediaResponse, resp, err := apiClient.MediaApi.MediaQuery(
		context.Background(),
		map[string]interface{}{},
	)
	if err != nil {
		panic(err)
	} else if resp.StatusCode >= 300 {
		panic(fmt.Errorf(
			"VoiceBase media query returned status %v",
			resp.StatusCode,
		))
	}

	fmtutil.PrintJSON(mediaResponse)

	fmt.Println("DONE")
}

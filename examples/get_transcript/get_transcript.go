package main

import (
	"context"
	"fmt"
	"os"

	clientutil "github.com/grokify/go-voicebase-v3"
	voicebase "github.com/grokify/go-voicebase-v3/client"
	"github.com/grokify/mogo/config"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/type/stringsutil"
)

const (
	EnvAccessToken = "VOICEBASE_BEARER_TOKEN"
	EnvMediaId     = "VOICEBASE_TRANSCRIPT_MEDIA_ID"
)

func TranscriptToSentence(vbt voicebase.VbTranscript) string {
	parts := []string{}
	for _, word := range vbt.Words {
		parts = append(parts, word.W)
	}
	return stringsutil.JoinTrimSpace(parts)
}

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

	fmt.Println(TranscriptToSentence(vbt))

	fmt.Println("DONE")
}

package main

import (
	"fmt"
	"os"

	"github.com/grokify/go-voicebase-v3"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/joho/godotenv"
)

const (
	EnvAccessToken = "VOICEBASE_BEARER_TOKEN"
	EnvMediaId     = "VOICEBASE_TRANSCRIPT_MEDIA_ID"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	accessToken := os.Getenv(EnvAccessToken)
	mediaId := os.Getenv(EnvMediaId)

	mediaApi := voicebase.NewMediaApi()
	mediaApi.Configuration.SetAccessToken(accessToken)

	vbt, resp, err := mediaApi.GetTranscript(mediaId, []string{})
	if err != nil {
		panic(err)
	}

	fmt.Println(string(resp.Payload))

	fmtutil.PrintJSON(vbt)

	fmt.Println("DONE")
}

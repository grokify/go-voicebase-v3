package main

import (
	"fmt"
	"os"

	"github.com/grokify/goauth"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/joho/godotenv"

	clientutil "github.com/grokify/go-voicebase-v3/v3"
)

const (
	EnvAccessToken = "VOICEBASE_BEARER_TOKEN"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	filename := "mpthreetest.mp3" // From https://archive.org/details/testmp3testfile
	vocabStrings := []string{"mpthree"}

	vocabs := clientutil.BuildVocabulariesForStrings(vocabStrings...)
	fmtutil.PrintJSON(vocabs)

	client := goauth.NewClientToken(goauth.TokenBearer, os.Getenv(EnvAccessToken), false)

	cfg := clientutil.MediaConfiguration{
		Language:               "en-US",
		EnableNumberFormatting: false,
		Vocabularies:           vocabs,
	}

	verbose := true
	fmtutil.PrintJSON(cfg)
	info, resp, err := clientutil.UploadMedia(client, filename, cfg.ToConfiguration(), verbose)
	if err != nil {
		panic(err)
	} else if resp.StatusCode >= 300 {
		panic(fmt.Sprintf("API Status %v", resp.StatusCode))
	}
	fmtutil.PrintJSON(info)
	fmt.Println("DONE")
}

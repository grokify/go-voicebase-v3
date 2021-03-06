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

	v := clientutil.Vocabulary{
		VocabularyName: "mpthree",
		Term:           "MP3",
		SoundsLike:     []string{"em pee three", "empee three", "m p three", "em p three"},
	}

	info, resp, err := apiClient.DefinitionApi.CreateVocabulary(
		context.Background(), v.VocabularyName, v.ToVbVocabulary(),
	)
	if err != nil {
		panic(err)
	} else if resp.StatusCode >= 300 {
		panic(fmt.Errorf("API Status %v", resp.StatusCode))
	}
	fmtutil.PrintJSON(info)

	fmt.Println("DONE")
}

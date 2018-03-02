package main

import (
	"context"
	"fmt"
	"os"

	"github.com/grokify/go-voicebase-v3"
	"github.com/grokify/go-voicebase-v3/client"
	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"
)

func BuildVocabulary(id string) voicebase.VbVocabulary {
	vb := voicebase.VbVocabulary{
		VocabularyName: id,
		VocabularyType: voicebase.TERMS,
		Terms: []voicebase.VbVocabularyTerm{
			{
				Term: "Embbnux",
				SoundsLike: []string{
					"emnux", "emnucks", "emnuks",
				},
			},
		},
	}
	return vb
}

/*
{
    "vocabularies": [
        {
            "vocabularyName": "rc-vocab",
            "vocabularyType": "terms",
            "terms": [
                {
                    "term": "Embbnux",
                    "soundsLike": [
                        "emnux",
                        "emnuks"
                    ]
                }
            ]
        }
    ],
    "_links": {
        "self": {
            "href": "https://apis.voicebase.com/v3/definition/vocabularies"
        }
    }
}
*/

func main() {
	err := config.LoadDotEnvSkipEmpty(os.Getenv("ENV_PATH"), "./.env")
	if err != nil {
		panic(err)
	}

	apiClient := clientutil.NewApiClientToken(os.Getenv("VOICEBASE_BEARER_TOKEN"))

	id := "embbnux"

	info, resp, err := apiClient.DefinitionApi.CreateVocabulary(
		context.Background(),
		id, BuildVocabulary(id),
	)
	if err != nil {
		panic(err)
	} else if resp.StatusCode >= 300 {
		panic(fmt.Errorf("API Status %v", resp.StatusCode))
	}
	fmtutil.PrintJSON(info)

	fmt.Println("DONE")
}

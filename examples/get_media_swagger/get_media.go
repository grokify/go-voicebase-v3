package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/grokify/go-voicebase-v3"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/joho/godotenv"
)

const (
	EnvAccessToken = "VOICEBASE_BEARER_TOKEN"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	accessToken := os.Getenv(EnvAccessToken)

	cfg := voicebase.NewConfiguration()

	resp, err := cfg.APIClient.CallAPI(
		"https://apis.voicebase.com/v3/media",
		"GET",
		"",
		map[string]string{"Authorization": fmt.Sprintf("Bearer %v", accessToken)},
		url.Values{},
		map[string]string{},
		"",
		[]byte(""))
	if err != nil {
		panic(err)
	}

	mediaResponse := voicebase.VbMediaQueryResponse{}

	fmt.Println(string(resp.Body()))

	err = json.Unmarshal(resp.Body(), &mediaResponse)
	if err != nil {
		panic(err)
	}

	fmtutil.PrintJSON(mediaResponse)

	fmt.Println("DONE")
}

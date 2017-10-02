package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/grokify/go-voicebase-v3"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/gotilla/net/httputilmore"
	"github.com/grokify/gotilla/time/timeutil"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

const (
	EnvAccessToken = "VOICEBASE_BEARER_TOKEN"
)

func NewClient() *http.Client {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	t0, _ := time.Parse(time.RFC3339, timeutil.RFC3339Zero)

	token := &oauth2.Token{
		AccessToken: os.Getenv(EnvAccessToken),
		TokenType:   "Bearer",
		Expiry:      t0}

	oAuthConfig := &oauth2.Config{}

	return oAuthConfig.Client(oauth2.NoContext, token)
}

func main() {
	client := NewClient()
	resp, err := client.Get("https://apis.voicebase.com/v3/media")
	if err != nil {
		panic(err)
	}

	mediaResponse := &voicebase.VbMediaQueryResponse{}

	err = httputilmore.UnmarshalResponseJSON(resp, mediaResponse)
	if err != nil {
		panic(err)
	}

	fmtutil.PrintJSON(mediaResponse)

	fmt.Println("DONE")
}

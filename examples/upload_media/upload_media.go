package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/gotilla/mime/multipartutil"
	"github.com/grokify/gotilla/net/httputilmore"
	"github.com/grokify/gotilla/time/timeutil"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

const (
	EnvAccessToken = "VOICEBASE_BEARER_TOKEN"
	SampleFile     = "SampleAudio_0.4mb.mp3"
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

type VbMediaUploadResponse struct {
	Links            Links       `json:"_links,omitempty"`
	FormatVersion    string      `json:"formatVersion,omitempty"`
	MediaId          string      `json:"mediaId,omitempty"`
	Status           string      `json:"status,omitempty"`
	DateCreated      time.Time   `json:"dateCreated,omitempty"`
	Metadata         interface{} `json:"metadata,omitempty"`
	MediaContentType string      `json:"mediaContentType,omitempty"`
	Length           int         `json:"length,omitempty"`
}

type Links struct {
	Self     LinkInfo `json:"metadata,omitempty"`
	Progress LinkInfo `json:"metadata,omitempty"`
	Metadata LinkInfo `json:"metadata,omitempty"`
}

type LinkInfo struct {
	Href string `json:"href,omitempty"`
}

func UploadMedia(client *http.Client, filepath string, verbose bool) *http.Response {
	builder := multipartutil.NewMultipartBuilder()

	err := builder.WriteFile("media", filepath)
	if err != nil {
		panic(err)
	}

	err = builder.Writer.Close()
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST",
		"https://apis.voicebase.com/v3/media",
		builder.Buffer)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", builder.Writer.FormDataContentType())

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response Status: %v\n", resp.StatusCode)

	if verbose {
		err = httputilmore.PrintResponse(resp, true)
		if err != nil {
			panic(err)
		}

		mediaResponse := &VbMediaUploadResponse{}

		err = httputilmore.UnmarshalResponseJSON(resp, mediaResponse)
		if err != nil {
			panic(err)
		}

		fmtutil.PrintJSON(mediaResponse)
	}
	return resp
}

func main() {
	client := NewClient()
	verbose := true
	UploadMedia(client, SampleFile, verbose)

	fmt.Println("DONE")
}

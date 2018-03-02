package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/gotilla/mime/multipartutil"
	"github.com/grokify/gotilla/net/httputilmore"
	"github.com/grokify/oauth2more"
	"github.com/joho/godotenv"
)

const (
	EnvAccessToken = "VOICEBASE_BEARER_TOKEN"

	// SampleFile from https://archive.org/details/testmp3testfile
	SampleFile = "mpthreetest.mp3"
)

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
	Self     LinkInfo `json:"self,omitempty"`
	Progress LinkInfo `json:"progress,omitempty"`
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
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	client := oauth2more.NewClientAccessToken(os.Getenv(EnvAccessToken))
	verbose := true
	UploadMedia(client, SampleFile, verbose)

	fmt.Println("DONE")
}

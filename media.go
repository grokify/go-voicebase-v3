package clientutil

import (
	"net/http"
	"time"

	voicebase "github.com/grokify/go-voicebase-v3/client"
	"github.com/grokify/mogo/encoding/jsonutil"
	"github.com/grokify/mogo/mime/multipartutil"
)

func UploadMedia(client *http.Client, filepath string, cfg UploadMediaConfiguration, verbose bool) (*UploadMediaResponse, *http.Response, error) {
	builder := multipartutil.NewMultipartBuilder()

	err := builder.WriteFilePath("media", filepath)
	if err != nil {
		return nil, nil, err
	}

	err = builder.WriteFieldAsJSON("configuration", cfg, false)
	if err != nil {
		return nil, nil, err
	}

	err = builder.Writer.Close()
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		ApiUrlV3Media,
		builder.Buffer)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", builder.Writer.FormDataContentType())

	resp, err := client.Do(req)
	if err != nil {
		return nil, resp, err
	}

	mediaResponse := &UploadMediaResponse{}

	_, err = jsonutil.UnmarshalReader(resp.Body, mediaResponse)

	return mediaResponse, resp, err
}

type MediaConfiguration struct {
	Language               string
	EnableNumberFormatting bool
	Vocabularies           []voicebase.VbVocabulary
}

func (mc *MediaConfiguration) ToConfiguration() UploadMediaConfiguration {
	return BuildUploadMediaConfiguration(mc.Language, mc.Vocabularies, mc.EnableNumberFormatting)
}

func BuildUploadMediaConfiguration(lang string, vocabs []voicebase.VbVocabulary, numFormatting bool) UploadMediaConfiguration {
	cfg := UploadMediaConfiguration{
		SpeechModel: SpeechModel{Language: "en-US"},
		Transcript: Transcript{
			Formatting: Formatting{
				EnableNumberFormatting: false,
			},
		},
		Vocabularies: vocabs,
	}
	return cfg
}

type UploadMediaConfiguration struct {
	SpeechModel  SpeechModel              `json:"speechModel,omitempty"`
	Transcript   Transcript               `json:"transcript"`
	Vocabularies []voicebase.VbVocabulary `json:"vocabularies,omitempty"`
}

type SpeechModel struct {
	Language string `json:"language,omitempty"`
}

type Transcript struct {
	Formatting Formatting `json:"formatting"`
}

type Formatting struct {
	EnableNumberFormatting bool `json:"enableNumberFormatting"`
}

type UploadMediaResponse struct {
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

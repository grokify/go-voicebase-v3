# Go API client for VoiceBase

[![Build Status][build-status-svg]][build-status-link]
[![Go Report Card][goreport-svg]][goreport-link]
[![Docs][docs-godoc-svg]][docs-godoc-link]
[![License][license-svg]][license-link]

APIs for speech recognition and speech analytics, powering insights every business needs.

## Overview

This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 3.0.1
- Package version: 2.3.0
- Build package: io.swagger.codegen.languages.GoClientCodegen
For more information, please visit [https://www.voicebase.com/contact-us/](https://www.voicebase.com/contact-us/)

## Installation

Install the client:

```
$ go get github.com/grokify/go-voicebase-v3
```

Add the following to your Go source:

```
import(
    "github.com/grokify/go-voicebase-v3"
)
```

## Documentation for API Endpoints

All URIs are relative to *https://apis.voicebase.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefinitionApi* | [**CreateKeywordSpottingGroup**](docs/DefinitionApi.md#createkeywordspottinggroup) | **Put** /v3/definition/spotting/groups/{groupId} | Create or update keyword spotting group
*DefinitionApi* | [**CreateVocabulary**](docs/DefinitionApi.md#createvocabulary) | **Put** /v3/definition/vocabularies/{vocabularyId} | Create or update a custom vocabulary
*DefinitionApi* | [**DeleteKeywordSpottingGroupById**](docs/DefinitionApi.md#deletekeywordspottinggroupbyid) | **Delete** /v3/definition/spotting/groups/{groupId} | Delete keyword spotting group
*DefinitionApi* | [**DeleteVocabularyById**](docs/DefinitionApi.md#deletevocabularybyid) | **Delete** /v3/definition/vocabularies/{vocabularyId} | Delete a custom vocabulary
*DefinitionApi* | [**GetClassifier**](docs/DefinitionApi.md#getclassifier) | **Get** /v3/definition/prediction/classifiers/{classifierId} | Retrieve details about a predictive classification model
*DefinitionApi* | [**GetClassifiers**](docs/DefinitionApi.md#getclassifiers) | **Get** /v3/definition/prediction/classifiers | List predictive classification models
*DefinitionApi* | [**GetDetector**](docs/DefinitionApi.md#getdetector) | **Get** /v3/definition/prediction/detectors/{detectorId} | Retrieve details about a detection model
*DefinitionApi* | [**GetDetectors**](docs/DefinitionApi.md#getdetectors) | **Get** /v3/definition/prediction/detectors | List detection models
*DefinitionApi* | [**GetKeywordSpottingGroupById**](docs/DefinitionApi.md#getkeywordspottinggroupbyid) | **Get** /v3/definition/spotting/groups/{groupId} | Retrieve a keyword spotting Group
*DefinitionApi* | [**GetKeywordSpottingGroups**](docs/DefinitionApi.md#getkeywordspottinggroups) | **Get** /v3/definition/spotting/groups | List keyword spotting groups
*DefinitionApi* | [**GetSearchableFields**](docs/DefinitionApi.md#getsearchablefields) | **Get** /v3/definition/media/search | Retrieve the list of extended metadata searchable fields
*DefinitionApi* | [**GetVocabularies**](docs/DefinitionApi.md#getvocabularies) | **Get** /v3/definition/vocabularies | List custom vocabularies
*DefinitionApi* | [**GetVocabularyById**](docs/DefinitionApi.md#getvocabularybyid) | **Get** /v3/definition/vocabularies/{vocabularyId} | Retrieve a custom vocabulary.
*DefinitionApi* | [**SetSearchableFields**](docs/DefinitionApi.md#setsearchablefields) | **Put** /v3/definition/media/search | Define extended metadata searchable fields
*KeyManagementApi* | [**CreateKey**](docs/KeyManagementApi.md#createkey) | **Post** /v3/profile/keys | Create API Key
*KeyManagementApi* | [**DeleteKeyById**](docs/KeyManagementApi.md#deletekeybyid) | **Delete** /v3/profile/keys/{keyId} | Delete API key
*KeyManagementApi* | [**GetKeyById**](docs/KeyManagementApi.md#getkeybyid) | **Get** /v3/profile/keys/{keyId} | Retrieve details about an API key
*KeyManagementApi* | [**GetKeys**](docs/KeyManagementApi.md#getkeys) | **Get** /v3/profile/keys | List API Keys
*MediaApi* | [**DeleteMediaById**](docs/MediaApi.md#deletemediabyid) | **Delete** /v3/media/{mediaId} | Delete a media, transcripts and analytics results
*MediaApi* | [**GetMediaById**](docs/MediaApi.md#getmediabyid) | **Get** /v3/media/{mediaId} | Retrieve transcript and analytical results from a media record
*MediaApi* | [**GetMetadataById**](docs/MediaApi.md#getmetadatabyid) | **Get** /v3/media/{mediaId}/metadata | Retrieve metadata associated with a media record
*MediaApi* | [**GetProgressById**](docs/MediaApi.md#getprogressbyid) | **Get** /v3/media/{mediaId}/progress | Retrieve processing progress
*MediaApi* | [**GetStreamByMediaIdAndName**](docs/MediaApi.md#getstreambymediaidandname) | **Get** /v3/media/{mediaId}/streams/{streamName} | Downloads the media identified by the stream name
*MediaApi* | [**GetStreamsById**](docs/MediaApi.md#getstreamsbyid) | **Get** /v3/media/{mediaId}/streams | List streams
*MediaApi* | [**GetSubtitleDfxpById**](docs/MediaApi.md#getsubtitledfxpbyid) | **Get** /v3/media/{mediaId}/transcript/dfxp | Retrieve DFXP transcript
*MediaApi* | [**GetSubtitleWebVttById**](docs/MediaApi.md#getsubtitlewebvttbyid) | **Get** /v3/media/{mediaId}/transcript/webvtt | Retrieve WEBVTT transcript
*MediaApi* | [**GetSubtitlesById**](docs/MediaApi.md#getsubtitlesbyid) | **Get** /v3/media/{mediaId}/transcript/srt | Retrieve SRT transcript
*MediaApi* | [**GetTextById**](docs/MediaApi.md#gettextbyid) | **Get** /v3/media/{mediaId}/transcript/text | Retrieve text transcript
*MediaApi* | [**GetTranscript**](docs/MediaApi.md#gettranscript) | **Get** /v3/media/{mediaId}/transcript | Retrieve transcript
*MediaApi* | [**MediaQuery**](docs/MediaApi.md#mediaquery) | **Get** /v3/media | List media records
*MediaApi* | [**PostMedia**](docs/MediaApi.md#postmedia) | **Post** /v3/media | Upload a media file for transcription and analysis
*MediaApi* | [**PostMediaById**](docs/MediaApi.md#postmediabyid) | **Post** /v3/media/{mediaId} | Align a transcript and re-run the job
*MediaApi* | [**SetMetadataById**](docs/MediaApi.md#setmetadatabyid) | **Put** /v3/media/{mediaId}/metadata | Update media metadata


## Documentation For Models

 - [VbAudioRedactorConfiguration](docs/VbAudioRedactorConfiguration.md)
 - [VbCallbackConfiguration](docs/VbCallbackConfiguration.md)
 - [VbCallbackFormatEnum](docs/VbCallbackFormatEnum.md)
 - [VbCallbackStreamEnum](docs/VbCallbackStreamEnum.md)
 - [VbCallbackTypeEnum](docs/VbCallbackTypeEnum.md)
 - [VbChannelConfiguration](docs/VbChannelConfiguration.md)
 - [VbClass](docs/VbClass.md)
 - [VbClassifier](docs/VbClassifier.md)
 - [VbClassifierConfiguration](docs/VbClassifierConfiguration.md)
 - [VbClassifierModel](docs/VbClassifierModel.md)
 - [VbClassifierModelsResponse](docs/VbClassifierModelsResponse.md)
 - [VbConfiguration](docs/VbConfiguration.md)
 - [VbContentFilteringConfiguration](docs/VbContentFilteringConfiguration.md)
 - [VbDetectedSegment](docs/VbDetectedSegment.md)
 - [VbDetection](docs/VbDetection.md)
 - [VbDetector](docs/VbDetector.md)
 - [VbDetectorConfiguration](docs/VbDetectorConfiguration.md)
 - [VbDetectorModel](docs/VbDetectorModel.md)
 - [VbDetectorModelsResponse](docs/VbDetectorModelsResponse.md)
 - [VbDiarization](docs/VbDiarization.md)
 - [VbError](docs/VbError.md)
 - [VbErrorResponse](docs/VbErrorResponse.md)
 - [VbFormattingConfiguration](docs/VbFormattingConfiguration.md)
 - [VbFrequency](docs/VbFrequency.md)
 - [VbGroup](docs/VbGroup.md)
 - [VbHttpMethodEnum](docs/VbHttpMethodEnum.md)
 - [VbIncludeTypeEnum](docs/VbIncludeTypeEnum.md)
 - [VbIngestConfiguration](docs/VbIngestConfiguration.md)
 - [VbJob](docs/VbJob.md)
 - [VbKey](docs/VbKey.md)
 - [VbKeyConfiguration](docs/VbKeyConfiguration.md)
 - [VbKeysResponse](docs/VbKeysResponse.md)
 - [VbKeyword](docs/VbKeyword.md)
 - [VbKeywordGroup](docs/VbKeywordGroup.md)
 - [VbKeywordGroupsResponse](docs/VbKeywordGroupsResponse.md)
 - [VbKnowledge](docs/VbKnowledge.md)
 - [VbKnowledgeConfiguration](docs/VbKnowledgeConfiguration.md)
 - [VbMediaQueryResponse](docs/VbMediaQueryResponse.md)
 - [VbMediaSummary](docs/VbMediaSummary.md)
 - [VbMention](docs/VbMention.md)
 - [VbMessage](docs/VbMessage.md)
 - [VbMetadata](docs/VbMetadata.md)
 - [VbOccurrence](docs/VbOccurrence.md)
 - [VbParameter](docs/VbParameter.md)
 - [VbParameterDefinition](docs/VbParameterDefinition.md)
 - [VbPrediction](docs/VbPrediction.md)
 - [VbPredictionConfiguration](docs/VbPredictionConfiguration.md)
 - [VbPriorityEnum](docs/VbPriorityEnum.md)
 - [VbProgress](docs/VbProgress.md)
 - [VbProgressTask](docs/VbProgressTask.md)
 - [VbPublishConfiguration](docs/VbPublishConfiguration.md)
 - [VbRedactorConfiguration](docs/VbRedactorConfiguration.md)
 - [VbReference](docs/VbReference.md)
 - [VbSearchableFields](docs/VbSearchableFields.md)
 - [VbSpeechModelConfiguration](docs/VbSpeechModelConfiguration.md)
 - [VbSpotting](docs/VbSpotting.md)
 - [VbSpottingConfiguration](docs/VbSpottingConfiguration.md)
 - [VbSpottingGroupConfiguration](docs/VbSpottingGroupConfiguration.md)
 - [VbStatusEnum](docs/VbStatusEnum.md)
 - [VbStereoConfiguration](docs/VbStereoConfiguration.md)
 - [VbStream](docs/VbStream.md)
 - [VbStreams](docs/VbStreams.md)
 - [VbTaskStatusEnum](docs/VbTaskStatusEnum.md)
 - [VbTopic](docs/VbTopic.md)
 - [VbTrackConfiguration](docs/VbTrackConfiguration.md)
 - [VbTranscript](docs/VbTranscript.md)
 - [VbTranscriptConfiguration](docs/VbTranscriptConfiguration.md)
 - [VbTranscriptFormat](docs/VbTranscriptFormat.md)
 - [VbTranscriptRedactorConfiguration](docs/VbTranscriptRedactorConfiguration.md)
 - [VbVocabulariesResponse](docs/VbVocabulariesResponse.md)
 - [VbVocabulary](docs/VbVocabulary.md)
 - [VbVocabularyConfiguration](docs/VbVocabularyConfiguration.md)
 - [VbVocabularyScript](docs/VbVocabularyScript.md)
 - [VbVocabularyScriptConfiguration](docs/VbVocabularyScriptConfiguration.md)
 - [VbVocabularyTerm](docs/VbVocabularyTerm.md)
 - [VbVocabularyTermConfiguration](docs/VbVocabularyTermConfiguration.md)
 - [VbVocabularyType](docs/VbVocabularyType.md)
 - [VbWord](docs/VbWord.md)
 - [VbWordTypeEnum](docs/VbWordTypeEnum.md)
 - [VbMedia](docs/VbMedia.md)


## Documentation For Authorization

## Authorization
- **Type**: API key 

Example
```
	auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
		Key: "APIKEY",
		Prefix: "Bearer", // Omit if not necessary.
	})
    r, err := client.Service.Operation(auth, args)
```

## Author

https://github.com/grokify/

## Support

https://stackoverflow.com/

 [build-status-svg]: https://api.travis-ci.org/grokify/go-voicebase-v3.svg?branch=master
 [build-status-link]: https://travis-ci.org/grokify/go-voicebase-v3
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/go-voicebase-v3
 [goreport-link]: https://goreportcard.com/report/github.com/grokify/go-voicebase-v3
 [docs-godoc-svg]: https://img.shields.io/badge/docs-godoc-blue.svg
 [docs-godoc-link]: https://godoc.org/github.com/grokify/go-voicebase-v3
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-link]: https://github.com/grokify/go-voicebase-v3/blob/master/LICENSE
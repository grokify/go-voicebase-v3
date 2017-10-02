# \MediaApi

All URIs are relative to *https://apis.voicebase.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMediaById**](MediaApi.md#DeleteMediaById) | **Delete** /v3/media/{mediaId} | Delete a media, transcripts and analytics results
[**GetMediaById**](MediaApi.md#GetMediaById) | **Get** /v3/media/{mediaId} | Retrieve transcript and analytical results from a media record
[**GetMetadataById**](MediaApi.md#GetMetadataById) | **Get** /v3/media/{mediaId}/metadata | Retrieve metadata associated with a media record
[**GetProgressById**](MediaApi.md#GetProgressById) | **Get** /v3/media/{mediaId}/progress | Retrieve processing progress
[**GetStreamByMediaIdAndName**](MediaApi.md#GetStreamByMediaIdAndName) | **Get** /v3/media/{mediaId}/streams/{streamName} | Downloads the media identified by the stream name
[**GetStreamsById**](MediaApi.md#GetStreamsById) | **Get** /v3/media/{mediaId}/streams | List streams
[**GetSubtitleDfxpById**](MediaApi.md#GetSubtitleDfxpById) | **Get** /v3/media/{mediaId}/transcript/dfxp | Retrieve DFXP transcript
[**GetSubtitleWebVttById**](MediaApi.md#GetSubtitleWebVttById) | **Get** /v3/media/{mediaId}/transcript/webvtt | Retrieve WEBVTT transcript
[**GetSubtitlesById**](MediaApi.md#GetSubtitlesById) | **Get** /v3/media/{mediaId}/transcript/srt | Retrieve SRT transcript
[**GetTextById**](MediaApi.md#GetTextById) | **Get** /v3/media/{mediaId}/transcript/text | Retrieve text transcript
[**GetTranscript**](MediaApi.md#GetTranscript) | **Get** /v3/media/{mediaId}/transcript | Retrieve transcript
[**MediaQuery**](MediaApi.md#MediaQuery) | **Get** /v3/media | List media records
[**PostMedia**](MediaApi.md#PostMedia) | **Post** /v3/media | Upload a media file for transcription and analysis
[**PostMediaById**](MediaApi.md#PostMediaById) | **Post** /v3/media/{mediaId} | Align a transcript and re-run the job
[**SetMetadataById**](MediaApi.md#SetMetadataById) | **Put** /v3/media/{mediaId}/metadata | Update media metadata


# **DeleteMediaById**
> DeleteMediaById($mediaId)

Delete a media, transcripts and analytics results

Delete this media


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mediaId** | **string**| Media identifier, a UUID. | 

### Return type

void (empty response body)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMediaById**
> VbMedia GetMediaById($mediaId, $includeAlternateFormat)

Retrieve transcript and analytical results from a media record

Retrieve analytical results from a previously uploaded media


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mediaId** | **string**| Media identifier, a UUID. | 
 **includeAlternateFormat** | [**[]string**](string.md)| Set of alternate formats to include in the response | [optional] 

### Return type

[**VbMedia**](VbMedia.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMetadataById**
> VbMetadata GetMetadataById($mediaId)

Retrieve metadata associated with a media record

Retrieve the media metadata


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mediaId** | **string**| Media identifier, a UUID. | 

### Return type

[**VbMetadata**](VbMetadata.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProgressById**
> VbJob GetProgressById($mediaId)

Retrieve processing progress

Retrieve processing progress for a given media


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mediaId** | **string**| Media identifier, a UUID. | 

### Return type

[**VbJob**](VbJob.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStreamByMediaIdAndName**
> VbStream GetStreamByMediaIdAndName($mediaId, $streamName)

Downloads the media identified by the stream name

Returns a redirect to the named media stream


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mediaId** | **string**| Media identifier, a UUID. | 
 **streamName** | **string**| A stream name | 

### Return type

[**VbStream**](VbStream.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStreamsById**
> VbStreams GetStreamsById($mediaId)

List streams

Get list of available media URLs


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mediaId** | **string**| Media identifier, a UUID. | 

### Return type

[**VbStreams**](VbStreams.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubtitleDfxpById**
> string GetSubtitleDfxpById($mediaId)

Retrieve DFXP transcript

Retrieve the transcript from a given media in DFXP format


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mediaId** | **string**| Media identifier, a UUID. | 

### Return type

**string**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/ttml+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubtitleWebVttById**
> string GetSubtitleWebVttById($mediaId)

Retrieve WEBVTT transcript

Retrieve the transcript from a given media in WEBVTT format


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mediaId** | **string**| Media identifier, a UUID. | 

### Return type

**string**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/vtt

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubtitlesById**
> string GetSubtitlesById($mediaId)

Retrieve SRT transcript

Retrieve the transcript from a given media in SRT (subtitles) format


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mediaId** | **string**| Media identifier, a UUID. | 

### Return type

**string**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/srt

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTextById**
> string GetTextById($mediaId)

Retrieve text transcript

Retrieve the transcript from a given media in plain text format


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mediaId** | **string**| Media identifier, a UUID. | 

### Return type

**string**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTranscript**
> VbTranscript GetTranscript($mediaId, $includeAlternateFormat)

Retrieve transcript

Retrieve the transcript from a given media


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mediaId** | **string**| Get media by Id. | 
 **includeAlternateFormat** | [**[]string**](string.md)| Set of alternate formats to include in the response | [optional] 

### Return type

[**VbTranscript**](VbTranscript.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MediaQuery**
> VbMediaQueryResponse MediaQuery($query, $externalId, $before, $after, $extendedFilter, $onOrAfterDate, $onOrBeforeDate, $sortOrder, $limit)

List media records

Retrieve a list of media previously uploaded that match a criteria


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string**| The full text search | [optional] 
 **externalId** | **string**| Media tagged with externalId in the metadata. | [optional] 
 **before** | **string**| Find media created before this mediaId | [optional] 
 **after** | **string**| Find media created after this mediaId | [optional] 
 **extendedFilter** | [**[]string**](string.md)| A special filter which is of the form &#39;extendedFilter&#x3D;Name;Value&#39; which allows you to filter by extended metadata. | [optional] 
 **onOrAfterDate** | **time.Time**| Media created on or after date. | [optional] 
 **onOrBeforeDate** | **time.Time**| Media created on or before date. | [optional] 
 **sortOrder** | **string**| Sort order. | [optional] 
 **limit** | **int32**| Control the number of values returned. | [optional] 

### Return type

[**VbMediaQueryResponse**](VbMediaQueryResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMedia**
> VbMedia PostMedia($media, $mediaUrl, $configuration, $metadata, $transcript)

Upload a media file for transcription and analysis

Upload new new media to the service as an attachment or from a url


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **media** | ***os.File**| Media file attached to the request. | [optional] 
 **mediaUrl** | **string**| URL where media file can be downloaded. | [optional] 
 **configuration** | **string**| A JSON object with configuration options. | [optional] 
 **metadata** | **string**| Metadata about the file being posted. | [optional] 
 **transcript** | **string**| A transcript | [optional] 

### Return type

[**VbMedia**](VbMedia.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: multipart/form-data, multipart/mixed
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMediaById**
> VbMedia PostMediaById($mediaId, $configuration, $metadata, $transcript)

Align a transcript and re-run the job

Upload a transcript to the service as an attachment for alignment and re-running of the job


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mediaId** | **string**| A JSON object with configuration options. | 
 **configuration** | **string**| A JSON object with configuration options. | [optional] 
 **metadata** | **string**| Metadata about the file being posted. | [optional] 
 **transcript** | **string**| A transcript | [optional] 

### Return type

[**VbMedia**](VbMedia.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: multipart/form-data, multipart/mixed
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetMetadataById**
> VbMetadata SetMetadataById($mediaId, $metadata)

Update media metadata

Set or update the media metadata


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mediaId** | **string**| Get media by Id. | 
 **metadata** | [**VbMetadata**](VbMetadata.md)| metadata. | 

### Return type

[**VbMetadata**](VbMetadata.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


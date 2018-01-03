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
> DeleteMediaById(ctx, mediaId)
Delete a media, transcripts and analytics results

Delete this media

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **mediaId** | **string**| Media identifier, a UUID. | 

### Return type

 (empty response body)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMediaById**
> VbMedia GetMediaById(ctx, mediaId, optional)
Retrieve transcript and analytical results from a media record

Retrieve analytical results from a previously uploaded media

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **mediaId** | **string**| Media identifier, a UUID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mediaId** | **string**| Media identifier, a UUID. | 
 **includeAlternateFormat** | [**[]string**](string.md)| Set of alternate formats to include in the response | 

### Return type

[**VbMedia**](VbMedia.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMetadataById**
> VbMetadata GetMetadataById(ctx, mediaId)
Retrieve metadata associated with a media record

Retrieve the media metadata

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
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
> VbJob GetProgressById(ctx, mediaId)
Retrieve processing progress

Retrieve processing progress for a given media

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
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
> VbStream GetStreamByMediaIdAndName(ctx, mediaId, streamName)
Downloads the media identified by the stream name

Returns a redirect to the named media stream

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
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
> VbStreams GetStreamsById(ctx, mediaId)
List streams

Get list of available media URLs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
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
> string GetSubtitleDfxpById(ctx, mediaId)
Retrieve DFXP transcript

Retrieve the transcript from a given media in DFXP format

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
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
> string GetSubtitleWebVttById(ctx, mediaId)
Retrieve WEBVTT transcript

Retrieve the transcript from a given media in WEBVTT format

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
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
> string GetSubtitlesById(ctx, mediaId)
Retrieve SRT transcript

Retrieve the transcript from a given media in SRT (subtitles) format

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
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
> string GetTextById(ctx, mediaId)
Retrieve text transcript

Retrieve the transcript from a given media in plain text format

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
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
> VbTranscript GetTranscript(ctx, mediaId, optional)
Retrieve transcript

Retrieve the transcript from a given media

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **mediaId** | **string**| Get media by Id. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mediaId** | **string**| Get media by Id. | 
 **includeAlternateFormat** | [**[]string**](string.md)| Set of alternate formats to include in the response | 

### Return type

[**VbTranscript**](VbTranscript.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MediaQuery**
> VbMediaQueryResponse MediaQuery(ctx, optional)
List media records

Retrieve a list of media previously uploaded that match a criteria

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string**| The full text search | 
 **externalId** | **string**| Media tagged with externalId in the metadata. | 
 **before** | **string**| Find media created before this mediaId | 
 **after** | **string**| Find media created after this mediaId | 
 **extendedFilter** | [**[]string**](string.md)| A special filter which is of the form &#39;extendedFilter&#x3D;Name;Value&#39; which allows you to filter by extended metadata. | 
 **onOrAfterDate** | **time.Time**| Media created on or after date. | 
 **onOrBeforeDate** | **time.Time**| Media created on or before date. | 
 **sortOrder** | **string**| Sort order. | 
 **limit** | **int32**| Control the number of values returned. | 

### Return type

[**VbMediaQueryResponse**](VbMediaQueryResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMedia**
> VbMedia PostMedia(ctx, optional)
Upload a media file for transcription and analysis

Upload new new media to the service as an attachment or from a url

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **media** | ***os.File**| Media file attached to the request. | 
 **mediaUrl** | **string**| URL where media file can be downloaded. | 
 **configuration** | **string**| A JSON object with configuration options. | 
 **metadata** | **string**| Metadata about the file being posted. | 
 **transcript** | **string**| A transcript | 

### Return type

[**VbMedia**](VbMedia.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: multipart/form-data, multipart/mixed
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMediaById**
> VbMedia PostMediaById(ctx, mediaId, optional)
Align a transcript and re-run the job

Upload a transcript to the service as an attachment for alignment and re-running of the job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **mediaId** | **string**| A JSON object with configuration options. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mediaId** | **string**| A JSON object with configuration options. | 
 **configuration** | **string**| A JSON object with configuration options. | 
 **metadata** | **string**| Metadata about the file being posted. | 
 **transcript** | **string**| A transcript | 

### Return type

[**VbMedia**](VbMedia.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: multipart/form-data, multipart/mixed
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetMetadataById**
> VbMetadata SetMetadataById(ctx, mediaId, metadata)
Update media metadata

Set or update the media metadata

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
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


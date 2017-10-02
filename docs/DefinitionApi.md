# \DefinitionApi

All URIs are relative to *https://apis.voicebase.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKeywordSpottingGroup**](DefinitionApi.md#CreateKeywordSpottingGroup) | **Put** /v3/definition/spotting/groups/{groupId} | Create or update keyword spotting group
[**CreateVocabulary**](DefinitionApi.md#CreateVocabulary) | **Put** /v3/definition/vocabularies/{vocabularyId} | Create or update a custom vocabulary
[**DeleteKeywordSpottingGroupById**](DefinitionApi.md#DeleteKeywordSpottingGroupById) | **Delete** /v3/definition/spotting/groups/{groupId} | Delete keyword spotting group
[**DeleteVocabularyById**](DefinitionApi.md#DeleteVocabularyById) | **Delete** /v3/definition/vocabularies/{vocabularyId} | Delete a custom vocabulary
[**GetClassifier**](DefinitionApi.md#GetClassifier) | **Get** /v3/definition/prediction/classifiers/{classifierId} | Retrieve details about a predictive classification model
[**GetClassifiers**](DefinitionApi.md#GetClassifiers) | **Get** /v3/definition/prediction/classifiers | List predictive classification models
[**GetDetector**](DefinitionApi.md#GetDetector) | **Get** /v3/definition/prediction/detectors/{detectorId} | Retrieve details about a detection model
[**GetDetectors**](DefinitionApi.md#GetDetectors) | **Get** /v3/definition/prediction/detectors | List detection models
[**GetKeywordSpottingGroupById**](DefinitionApi.md#GetKeywordSpottingGroupById) | **Get** /v3/definition/spotting/groups/{groupId} | Retrieve a keyword spotting Group
[**GetKeywordSpottingGroups**](DefinitionApi.md#GetKeywordSpottingGroups) | **Get** /v3/definition/spotting/groups | List keyword spotting groups
[**GetSearchableFields**](DefinitionApi.md#GetSearchableFields) | **Get** /v3/definition/media/search | Retrieve the list of extended metadata searchable fields
[**GetVocabularies**](DefinitionApi.md#GetVocabularies) | **Get** /v3/definition/vocabularies | List custom vocabularies
[**GetVocabularyById**](DefinitionApi.md#GetVocabularyById) | **Get** /v3/definition/vocabularies/{vocabularyId} | Retrieve a custom vocabulary.
[**SetSearchableFields**](DefinitionApi.md#SetSearchableFields) | **Put** /v3/definition/media/search | Define extended metadata searchable fields


# **CreateKeywordSpottingGroup**
> VbKeywordGroup CreateKeywordSpottingGroup($groupId, $keywordSpottingGroup)

Create or update keyword spotting group

Create or update a keyword spotting group with a set of keywords


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string**| The keyword spotting group identifier. | 
 **keywordSpottingGroup** | [**VbKeywordGroup**](VbKeywordGroup.md)| Keyword Spotting Group definition | 

### Return type

[**VbKeywordGroup**](VbKeywordGroup.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVocabulary**
> VbVocabulary CreateVocabulary($vocabularyId, $vocabulary)

Create or update a custom vocabulary

Create or update a custom vocabulary. Custom vocabularies are used to improve accurary of transcription


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vocabularyId** | **string**| The vocabulary identifier (name). | 
 **vocabulary** | [**VbVocabulary**](VbVocabulary.md)| Vocabulary defintion. | 

### Return type

[**VbVocabulary**](VbVocabulary.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteKeywordSpottingGroupById**
> DeleteKeywordSpottingGroupById($groupId)

Delete keyword spotting group

Delete the keyword spotting group


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string**| The keyword spotting group identifier. | 

### Return type

void (empty response body)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVocabularyById**
> DeleteVocabularyById($vocabularyId)

Delete a custom vocabulary

Delete a custom vocabulary


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vocabularyId** | **string**| The vocabulary identifier (name) | 

### Return type

void (empty response body)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClassifier**
> VbClassifierModel GetClassifier($classifierId)

Retrieve details about a predictive classification model

Get the classifier model


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **classifierId** | **string**| The classifier model identifier. | 

### Return type

[**VbClassifierModel**](VbClassifierModel.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClassifiers**
> VbClassifierModelsResponse GetClassifiers()

List predictive classification models

List available predictive models for classification


### Parameters
This endpoint does not need any parameter.

### Return type

[**VbClassifierModelsResponse**](VbClassifierModelsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDetector**
> VbDetectorModel GetDetector($detectorId)

Retrieve details about a detection model

Retrieve details about a detection model


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **detectorId** | **string**| The detector model identifier. | 

### Return type

[**VbDetectorModel**](VbDetectorModel.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDetectors**
> VbDetectorModelsResponse GetDetectors()

List detection models

List available detection models


### Parameters
This endpoint does not need any parameter.

### Return type

[**VbDetectorModelsResponse**](VbDetectorModelsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKeywordSpottingGroupById**
> VbKeywordGroup GetKeywordSpottingGroupById($groupId)

Retrieve a keyword spotting Group

Retrieve a keyword spotting group with its keywords


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string**| Keyword spotting group identifier | 

### Return type

[**VbKeywordGroup**](VbKeywordGroup.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKeywordSpottingGroups**
> VbKeywordGroupsResponse GetKeywordSpottingGroups()

List keyword spotting groups

Get all defined keyword spotting groups


### Parameters
This endpoint does not need any parameter.

### Return type

[**VbKeywordGroupsResponse**](VbKeywordGroupsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSearchableFields**
> VbSearchableFields GetSearchableFields()

Retrieve the list of extended metadata searchable fields

Get searchable fields


### Parameters
This endpoint does not need any parameter.

### Return type

[**VbSearchableFields**](VbSearchableFields.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVocabularies**
> VbVocabulariesResponse GetVocabularies()

List custom vocabularies

List all defined custom vocabularies. Custom vocabularies are used to improve accurary of transcription


### Parameters
This endpoint does not need any parameter.

### Return type

[**VbVocabulariesResponse**](VbVocabulariesResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVocabularyById**
> VbVocabulary GetVocabularyById($vocabularyId)

Retrieve a custom vocabulary.

Retrieve a custom vocabulary


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vocabularyId** | **string**| The vocabulary identifier (name) | 

### Return type

[**VbVocabulary**](VbVocabulary.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetSearchableFields**
> VbSearchableFields SetSearchableFields($searchableFields)

Define extended metadata searchable fields

Create or update custom parameters of metadata for search


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchableFields** | [**VbSearchableFields**](VbSearchableFields.md)| Key config. | 

### Return type

[**VbSearchableFields**](VbSearchableFields.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


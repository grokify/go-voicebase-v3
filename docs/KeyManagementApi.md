# \KeyManagementApi

All URIs are relative to *https://apis.voicebase.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKey**](KeyManagementApi.md#CreateKey) | **Post** /v3/profile/keys | Create API Key
[**DeleteKeyById**](KeyManagementApi.md#DeleteKeyById) | **Delete** /v3/profile/keys/{keyId} | Delete API key
[**GetKeyById**](KeyManagementApi.md#GetKeyById) | **Get** /v3/profile/keys/{keyId} | Retrieve details about an API key
[**GetKeys**](KeyManagementApi.md#GetKeys) | **Get** /v3/profile/keys | List API Keys


# **CreateKey**
> VbKey CreateKey($key)

Create API Key

Create a new API key for the current user


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | [**VbKeyConfiguration**](VbKeyConfiguration.md)| Key config. | 

### Return type

[**VbKey**](VbKey.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteKeyById**
> DeleteKeyById($keyId)

Delete API key

Delete and revoke this API key


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyId** | **string**| The key Id. | 

### Return type

void (empty response body)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKeyById**
> VbKey GetKeyById($keyId)

Retrieve details about an API key

Get information about this API key


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyId** | **string**| The key Id. | 

### Return type

[**VbKey**](VbKey.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKeys**
> VbKeysResponse GetKeys()

List API Keys

Returns all current API keys for current user


### Parameters
This endpoint does not need any parameter.

### Return type

[**VbKeysResponse**](VbKeysResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


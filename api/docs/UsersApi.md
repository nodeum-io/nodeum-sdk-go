# \UsersApi

All URIs are relative to *https://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiKey**](UsersApi.md#CreateApiKey) | **Post** /users/me/api_keys | Creates a new API Key for current user.
[**DestroyApiKey**](UsersApi.md#DestroyApiKey) | **Delete** /users/me/api_keys/{api_key_id} | Destroys a specific API Key.
[**IndexApiKeys**](UsersApi.md#IndexApiKeys) | **Get** /users/me/api_keys | Lists all API keys of current user.
[**ShowApiKey**](UsersApi.md#ShowApiKey) | **Get** /users/me/api_keys/{api_key_id} | Displays a specific API Key with its scopes.
[**UpdateApiKey**](UsersApi.md#UpdateApiKey) | **Put** /users/me/api_keys/{api_key_id} | Updates a specific API Key.


# **CreateApiKey**
> ApiKeyFull CreateApiKey(ctx, apiKeyBody)
Creates a new API Key for current user.

**API Key Scope**: api_keys / create

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKeyBody** | [**ApiKeyFull**](ApiKeyFull.md)|  | 

### Return type

[**ApiKeyFull**](api_key_full.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DestroyApiKey**
> DestroyApiKey(ctx, apiKeyId)
Destroys a specific API Key.

**API Key Scope**: api_keys / destroy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKeyId** | **int32**| Numeric ID of API Key. | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IndexApiKeys**
> ApiKeyCollection IndexApiKeys(ctx, optional)
Lists all API keys of current user.

**API Key Scope**: api_keys / index

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IndexApiKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IndexApiKeysOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 

### Return type

[**ApiKeyCollection**](api_key_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowApiKey**
> ApiKeyFull ShowApiKey(ctx, apiKeyId)
Displays a specific API Key with its scopes.

**API Key Scope**: api_keys / show

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKeyId** | **int32**| Numeric ID of API Key. | 

### Return type

[**ApiKeyFull**](api_key_full.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateApiKey**
> ApiKeyFull UpdateApiKey(ctx, apiKeyId, apiKeyBody)
Updates a specific API Key.

**API Key Scope**: api_keys / update

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKeyId** | **int32**| Numeric ID of API Key. | 
  **apiKeyBody** | [**ApiKeyFull**](ApiKeyFull.md)|  | 

### Return type

[**ApiKeyFull**](api_key_full.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


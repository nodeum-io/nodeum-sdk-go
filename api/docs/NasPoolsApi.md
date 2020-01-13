# \NasPoolsApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNasPool**](NasPoolsApi.md#CreateNasPool) | **Post** /nas_pools | Creates a new NAS pool.
[**DestroyNasPool**](NasPoolsApi.md#DestroyNasPool) | **Delete** /nas_pools/{nas_pool_id} | Destroys a specific NAS pool.
[**IndexNasPools**](NasPoolsApi.md#IndexNasPools) | **Get** /nas_pools | Lists all NAS pools.
[**MountStatusNasPool**](NasPoolsApi.md#MountStatusNasPool) | **Get** /nas_pools/{nas_pool_id}/mount | Get mount status of NAS pool.
[**ShowNasPool**](NasPoolsApi.md#ShowNasPool) | **Get** /nas_pools/{nas_pool_id} | Displays a specific NAS pool.
[**UpdateNasPool**](NasPoolsApi.md#UpdateNasPool) | **Put** /nas_pools/{nas_pool_id} | Updates a specific NAS pool.



## CreateNasPool

> NasPool CreateNasPool(ctx, nasPoolBody)

Creates a new NAS pool.

**API Key Scope**: nas_pools / create

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasPoolBody** | [**NasPoolUp**](NasPoolUp.md)|  | 

### Return type

[**NasPool**](nas_pool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyNasPool

> DestroyNasPool(ctx, nasPoolId)

Destroys a specific NAS pool.

**API Key Scope**: nas_pools / destroy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasPoolId** | **string**| Numeric ID or name of NAS pool. | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexNasPools

> NasPoolCollection IndexNasPools(ctx, optional)

Lists all NAS pools.

**API Key Scope**: nas_pools / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IndexNasPoolsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexNasPoolsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **name** | **optional.String**| Filter on name | 
 **comment** | **optional.String**| Filter on comment | 
 **type_** | **optional.String**| Filter on type | 

### Return type

[**NasPoolCollection**](nas_pool_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MountStatusNasPool

> MountStatus MountStatusNasPool(ctx, nasPoolId)

Get mount status of NAS pool.

**API Key Scope**: nas_pools / mount_status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasPoolId** | **string**| Numeric ID or name of NAS pool. | 

### Return type

[**MountStatus**](mount_status.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowNasPool

> NasPool ShowNasPool(ctx, nasPoolId)

Displays a specific NAS pool.

**API Key Scope**: nas_pools / show

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasPoolId** | **string**| Numeric ID or name of NAS pool. | 

### Return type

[**NasPool**](nas_pool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNasPool

> NasPool UpdateNasPool(ctx, nasPoolId, nasPoolBody)

Updates a specific NAS pool.

**API Key Scope**: nas_pools / update

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasPoolId** | **string**| Numeric ID or name of NAS pool. | 
**nasPoolBody** | [**NasPoolUp**](NasPoolUp.md)|  | 

### Return type

[**NasPool**](nas_pool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


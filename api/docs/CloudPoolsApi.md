# \CloudPoolsApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCloudPool**](CloudPoolsApi.md#CreateCloudPool) | **Post** /cloud_pools | Creates a new cloud pool.
[**DestroyCloudPool**](CloudPoolsApi.md#DestroyCloudPool) | **Delete** /cloud_pools/{cloud_pool_id} | Destroys a specific cloud pool.
[**IndexCloudPools**](CloudPoolsApi.md#IndexCloudPools) | **Get** /cloud_pools | Lists all cloud pools.
[**MountStatusCloudPool**](CloudPoolsApi.md#MountStatusCloudPool) | **Get** /cloud_pools/{cloud_pool_id}/mount | Get mount status of Cloud pool.
[**ShowCloudPool**](CloudPoolsApi.md#ShowCloudPool) | **Get** /cloud_pools/{cloud_pool_id} | Displays a specific cloud pool.
[**UpdateCloudPool**](CloudPoolsApi.md#UpdateCloudPool) | **Put** /cloud_pools/{cloud_pool_id} | Updates a specific cloud pool.



## CreateCloudPool

> CloudPool CreateCloudPool(ctx, cloudPoolBody)

Creates a new cloud pool.

**API Key Scope**: cloud_pools / create

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudPoolBody** | [**CloudPoolUp**](CloudPoolUp.md)|  | 

### Return type

[**CloudPool**](cloud_pool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyCloudPool

> DestroyCloudPool(ctx, cloudPoolId)

Destroys a specific cloud pool.

**API Key Scope**: cloud_pools / destroy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudPoolId** | **string**| Numeric ID or name of cloud pool. | 

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


## IndexCloudPools

> CloudPoolCollection IndexCloudPools(ctx, optional)

Lists all cloud pools.

**API Key Scope**: cloud_pools / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IndexCloudPoolsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexCloudPoolsOpts struct


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

[**CloudPoolCollection**](cloud_pool_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MountStatusCloudPool

> MountStatus MountStatusCloudPool(ctx, cloudPoolId)

Get mount status of Cloud pool.

**API Key Scope**: cloud_pools / mount_status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudPoolId** | **string**| Numeric ID or name of cloud pool. | 

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


## ShowCloudPool

> CloudPool ShowCloudPool(ctx, cloudPoolId)

Displays a specific cloud pool.

**API Key Scope**: cloud_pools / show

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudPoolId** | **string**| Numeric ID or name of cloud pool. | 

### Return type

[**CloudPool**](cloud_pool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCloudPool

> CloudPool UpdateCloudPool(ctx, cloudPoolId, cloudPoolBody)

Updates a specific cloud pool.

**API Key Scope**: cloud_pools / update

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudPoolId** | **string**| Numeric ID or name of cloud pool. | 
**cloudPoolBody** | [**CloudPoolUp**](CloudPoolUp.md)|  | 

### Return type

[**CloudPool**](cloud_pool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


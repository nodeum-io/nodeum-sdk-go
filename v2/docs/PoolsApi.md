# \PoolsApi

All URIs are relative to *https://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePool**](PoolsApi.md#CreatePool) | **Post** /pools | Creates a new pool.
[**DestroyPool**](PoolsApi.md#DestroyPool) | **Delete** /pools/{pool_id} | Destroys a specific tape pool.
[**IndexPools**](PoolsApi.md#IndexPools) | **Get** /pools | Lists all pools.
[**ShowPool**](PoolsApi.md#ShowPool) | **Get** /pools/{pool_id} | Displays a specific pool.
[**UpdatePool**](PoolsApi.md#UpdatePool) | **Put** /pools/{pool_id} | Updates a specific pool.


# **CreatePool**
> Pool CreatePool(ctx, poolBody)
Creates a new pool.

**API Key Scope**: pools / create

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolBody** | [**PoolUp**](PoolUp.md)|  | 

### Return type

[**Pool**](pool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DestroyPool**
> DestroyPool(ctx, poolId)
Destroys a specific tape pool.

**API Key Scope**: pools / destroy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**| Numeric ID, or name of pool. | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IndexPools**
> PoolCollection IndexPools(ctx, optional)
Lists all pools.

**API Key Scope**: pools / index

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IndexPoolsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IndexPoolsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **name** | **optional.String**| Filter on name | 
 **comment** | **optional.String**| Filter on comment | 
 **type_** | **optional.String**| Filter on type | 
 **content** | **optional.String**| Filter on content | 
 **primaryId** | **optional.String**| Filter on primary id | 

### Return type

[**PoolCollection**](pool_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowPool**
> Pool ShowPool(ctx, poolId)
Displays a specific pool.

**API Key Scope**: pools / show

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**| Numeric ID, or name of pool. | 

### Return type

[**Pool**](pool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePool**
> Pool UpdatePool(ctx, poolId, poolBody)
Updates a specific pool.

**API Key Scope**: pools / update

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**| Numeric ID, or name of pool. | 
  **poolBody** | [**PoolUp**](PoolUp.md)|  | 

### Return type

[**Pool**](pool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

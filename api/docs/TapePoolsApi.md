# \TapePoolsApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTapePool**](TapePoolsApi.md#CreateTapePool) | **Post** /tape_pools | Creates a new tape pool.
[**DestroyTapePool**](TapePoolsApi.md#DestroyTapePool) | **Delete** /tape_pools/{tape_pool_id} | Destroys a specific tape pool.
[**IndexTapePools**](TapePoolsApi.md#IndexTapePools) | **Get** /tape_pools | Lists all tape pools.
[**ShowTapePool**](TapePoolsApi.md#ShowTapePool) | **Get** /tape_pools/{tape_pool_id} | Displays a specific tape pool.
[**UpdateTapePool**](TapePoolsApi.md#UpdateTapePool) | **Put** /tape_pools/{tape_pool_id} | Updates a specific tape pool.



## CreateTapePool

> TapePool CreateTapePool(ctx, tapePoolBody)

Creates a new tape pool.

**API Key Scope**: tape_pools / create

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tapePoolBody** | [**TapePoolUp**](TapePoolUp.md)|  | 

### Return type

[**TapePool**](tape_pool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyTapePool

> DestroyTapePool(ctx, tapePoolId)

Destroys a specific tape pool.

**API Key Scope**: tape_pools / destroy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tapePoolId** | **string**| Numeric ID, or name of tape pool. | 

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


## IndexTapePools

> TapePoolCollection IndexTapePools(ctx, optional)

Lists all tape pools.

**API Key Scope**: tape_pools / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IndexTapePoolsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexTapePoolsOpts struct


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

[**TapePoolCollection**](tape_pool_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowTapePool

> TapePool ShowTapePool(ctx, tapePoolId)

Displays a specific tape pool.

**API Key Scope**: tape_pools / show

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tapePoolId** | **string**| Numeric ID, or name of tape pool. | 

### Return type

[**TapePool**](tape_pool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTapePool

> TapePool UpdateTapePool(ctx, tapePoolId, tapePoolBody)

Updates a specific tape pool.

**API Key Scope**: tape_pools / update

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tapePoolId** | **string**| Numeric ID, or name of tape pool. | 
**tapePoolBody** | [**TapePoolUp**](TapePoolUp.md)|  | 

### Return type

[**TapePool**](tape_pool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


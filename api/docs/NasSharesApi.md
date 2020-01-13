# \NasSharesApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNasShareByNas**](NasSharesApi.md#CreateNasShareByNas) | **Post** /nas/{nas_id}/nas_shares | Creates a new NAS share.
[**DestroyNasShare**](NasSharesApi.md#DestroyNasShare) | **Delete** /nas_shares/{nas_share_id} | Destroys a specific NAS share.
[**DestroyNasShareByNas**](NasSharesApi.md#DestroyNasShareByNas) | **Delete** /nas/{nas_id}/nas_shares/{nas_share_id} | Destroys a specific NAS share.
[**DestroyNasShareByNasPool**](NasSharesApi.md#DestroyNasShareByNasPool) | **Delete** /nas_pools/{nas_pool_id}/nas_shares/{nas_share_id} | Destroys a specific NAS share.
[**IndexNasShares**](NasSharesApi.md#IndexNasShares) | **Get** /nas_shares | Lists all NAS shares.
[**IndexNasSharesByNas**](NasSharesApi.md#IndexNasSharesByNas) | **Get** /nas/{nas_id}/nas_shares | Lists all NAS shares.
[**IndexNasSharesByNasPool**](NasSharesApi.md#IndexNasSharesByNasPool) | **Get** /nas_pools/{nas_pool_id}/nas_shares | Lists all NAS shares.
[**MountStatusNasShare**](NasSharesApi.md#MountStatusNasShare) | **Get** /nas_shares/{nas_share_id}/mount | Get mount status of NAS Share.
[**MountStatusNasShareByNas**](NasSharesApi.md#MountStatusNasShareByNas) | **Get** /nas/{nas_id}/nas_shares/{nas_share_id}/mount | Get mount status of NAS Share.
[**MountStatusNasShareByNasPool**](NasSharesApi.md#MountStatusNasShareByNasPool) | **Get** /nas_pools/{nas_pool_id}/nas_shares/{nas_share_id}/mount | Get mount status of NAS Share.
[**ShowNasShareByNas**](NasSharesApi.md#ShowNasShareByNas) | **Get** /nas/{nas_id}/nas_shares/{nas_share_id} | Displays a specific NAS share.
[**ShowNasShares**](NasSharesApi.md#ShowNasShares) | **Get** /nas_shares/{nas_share_id} | Displays a specific NAS share.
[**ShowNasSharesByNasPool**](NasSharesApi.md#ShowNasSharesByNasPool) | **Get** /nas_pools/{nas_pool_id}/nas_shares/{nas_share_id} | Displays a specific NAS share.
[**TestNasShare**](NasSharesApi.md#TestNasShare) | **Put** /nas/{nas_id}/nas_shares/-/test | Test an unsaved NAS Share.
[**TestResultNasShare**](NasSharesApi.md#TestResultNasShare) | **Get** /nas/{nas_id}/nas_shares/-/test | Check result of a NAS Share test job.
[**UpdateNasShare**](NasSharesApi.md#UpdateNasShare) | **Put** /nas_shares/{nas_share_id} | Updates a specific NAS share.
[**UpdateNasShareByNas**](NasSharesApi.md#UpdateNasShareByNas) | **Put** /nas/{nas_id}/nas_shares/{nas_share_id} | Updates a specific NAS share.
[**UpdateNasShareByNasPool**](NasSharesApi.md#UpdateNasShareByNasPool) | **Put** /nas_pools/{nas_pool_id}/nas_shares/{nas_share_id} | Updates a specific NAS share.



## CreateNasShareByNas

> NasShare CreateNasShareByNas(ctx, nasId, nasShareBody)

Creates a new NAS share.

**API Key Scope**: nas_shares / create

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasId** | **string**| Numeric ID or name of NAS. | 
**nasShareBody** | [**NasShare**](NasShare.md)|  | 

### Return type

[**NasShare**](nas_share.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyNasShare

> DestroyNasShare(ctx, nasShareId)

Destroys a specific NAS share.

**API Key Scope**: nas_shares / destroy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasShareId** | **int32**| Numeric ID of NAS share. | 

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


## DestroyNasShareByNas

> DestroyNasShareByNas(ctx, nasId, nasShareId)

Destroys a specific NAS share.

**API Key Scope**: nas_shares / destroy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasId** | **string**| Numeric ID or name of NAS. | 
**nasShareId** | **int32**| Numeric ID of NAS share. | 

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


## DestroyNasShareByNasPool

> DestroyNasShareByNasPool(ctx, nasPoolId, nasShareId)

Destroys a specific NAS share.

**API Key Scope**: nas_shares / destroy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasPoolId** | **string**| Numeric ID or name of NAS pool. | 
**nasShareId** | **int32**| Numeric ID of NAS share. | 

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


## IndexNasShares

> NasShareCollection IndexNasShares(ctx, optional)

Lists all NAS shares.

**API Key Scope**: nas_shares / index   Optional API Key Explicit Scope: nas_shares / get_password

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IndexNasSharesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexNasSharesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **path** | **optional.String**| Filter on path | 
 **options** | **optional.String**| Filter on options | 
 **username** | **optional.String**| Filter on username | 
 **nasId** | **optional.String**| Filter on NAS id | 
 **nasPoolId** | **optional.String**| Filter on NAS pool id | 

### Return type

[**NasShareCollection**](nas_share_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexNasSharesByNas

> NasShareCollection IndexNasSharesByNas(ctx, nasId, optional)

Lists all NAS shares.

**API Key Scope**: nas_shares / index   Optional API Key Explicit Scope: nas_shares / get_password

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasId** | **string**| Numeric ID or name of NAS. | 
 **optional** | ***IndexNasSharesByNasOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexNasSharesByNasOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **path** | **optional.String**| Filter on path | 
 **options** | **optional.String**| Filter on options | 
 **username** | **optional.String**| Filter on username | 
 **nasPoolId** | **optional.String**| Filter on NAS pool id | 

### Return type

[**NasShareCollection**](nas_share_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexNasSharesByNasPool

> NasShareCollection IndexNasSharesByNasPool(ctx, nasPoolId, optional)

Lists all NAS shares.

**API Key Scope**: nas_shares / index   Optional API Key Explicit Scope: nas_shares / get_password

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasPoolId** | **string**| Numeric ID or name of NAS pool. | 
 **optional** | ***IndexNasSharesByNasPoolOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexNasSharesByNasPoolOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **path** | **optional.String**| Filter on path | 
 **options** | **optional.String**| Filter on options | 
 **username** | **optional.String**| Filter on username | 
 **nasId** | **optional.String**| Filter on NAS id | 

### Return type

[**NasShareCollection**](nas_share_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MountStatusNasShare

> MountStatus MountStatusNasShare(ctx, nasShareId)

Get mount status of NAS Share.

**API Key Scope**: nas_shares / mount_status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasShareId** | **int32**| Numeric ID of NAS share. | 

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


## MountStatusNasShareByNas

> MountStatus MountStatusNasShareByNas(ctx, nasId, nasShareId)

Get mount status of NAS Share.

**API Key Scope**: nas_shares / mount_status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasId** | **string**| Numeric ID or name of NAS. | 
**nasShareId** | **int32**| Numeric ID of NAS share. | 

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


## MountStatusNasShareByNasPool

> MountStatus MountStatusNasShareByNasPool(ctx, nasPoolId, nasShareId)

Get mount status of NAS Share.

**API Key Scope**: nas_shares / mount_status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasPoolId** | **string**| Numeric ID or name of NAS pool. | 
**nasShareId** | **int32**| Numeric ID of NAS share. | 

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


## ShowNasShareByNas

> NasShare ShowNasShareByNas(ctx, nasId, nasShareId)

Displays a specific NAS share.

**API Key Scope**: nas_shares / show   Optional API Key Explicit Scope: nas_shares / get_password

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasId** | **string**| Numeric ID or name of NAS. | 
**nasShareId** | **int32**| Numeric ID of NAS share. | 

### Return type

[**NasShare**](nas_share.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowNasShares

> NasShare ShowNasShares(ctx, nasShareId)

Displays a specific NAS share.

**API Key Scope**: nas_shares / show   Optional API Key Explicit Scope: nas_shares / get_password

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasShareId** | **int32**| Numeric ID of NAS share. | 

### Return type

[**NasShare**](nas_share.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowNasSharesByNasPool

> NasShare ShowNasSharesByNasPool(ctx, nasPoolId, nasShareId)

Displays a specific NAS share.

**API Key Scope**: nas_shares / show   Optional API Key Explicit Scope: nas_shares / get_password

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasPoolId** | **string**| Numeric ID or name of NAS pool. | 
**nasShareId** | **int32**| Numeric ID of NAS share. | 

### Return type

[**NasShare**](nas_share.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestNasShare

> ActiveJobStatus TestNasShare(ctx, nasId, nasShareBody)

Test an unsaved NAS Share.

**API Key Scope**: nas_shares / test

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasId** | **string**| Numeric ID or name of NAS. | 
**nasShareBody** | [**NasShare**](NasShare.md)|  | 

### Return type

[**ActiveJobStatus**](active_job_status.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, queued, working, failed

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestResultNasShare

> ActiveJobStatus TestResultNasShare(ctx, nasId, optional)

Check result of a NAS Share test job.

**API Key Scope**: nas_shares / test

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasId** | **string**| Numeric ID or name of NAS. | 
 **optional** | ***TestResultNasShareOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TestResultNasShareOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jobId** | **optional.String**| ID of active job | 

### Return type

[**ActiveJobStatus**](active_job_status.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, queued, working, failed

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNasShare

> NasShare UpdateNasShare(ctx, nasShareId, nasShareBody)

Updates a specific NAS share.

**API Key Scope**: nas_shares / update

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasShareId** | **int32**| Numeric ID of NAS share. | 
**nasShareBody** | [**NasShare**](NasShare.md)|  | 

### Return type

[**NasShare**](nas_share.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNasShareByNas

> NasShare UpdateNasShareByNas(ctx, nasId, nasShareId, nasShareBody)

Updates a specific NAS share.

**API Key Scope**: nas_shares / update

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasId** | **string**| Numeric ID or name of NAS. | 
**nasShareId** | **int32**| Numeric ID of NAS share. | 
**nasShareBody** | [**NasShare**](NasShare.md)|  | 

### Return type

[**NasShare**](nas_share.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNasShareByNasPool

> NasShare UpdateNasShareByNasPool(ctx, nasPoolId, nasShareId, nasShareBody)

Updates a specific NAS share.

**API Key Scope**: nas_shares / update

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasPoolId** | **string**| Numeric ID or name of NAS pool. | 
**nasShareId** | **int32**| Numeric ID of NAS share. | 
**nasShareBody** | [**NasShare**](NasShare.md)|  | 

### Return type

[**NasShare**](nas_share.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


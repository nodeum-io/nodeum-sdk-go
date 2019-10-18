# \CloudBucketsApi

All URIs are relative to *https://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IndexCloudBuckets**](CloudBucketsApi.md#IndexCloudBuckets) | **Get** /cloud_buckets | Lists all cloud buckets.
[**IndexCloudBucketsByCloudConnector**](CloudBucketsApi.md#IndexCloudBucketsByCloudConnector) | **Get** /cloud_connectors/{cloud_connector_id}/cloud_buckets | Lists all cloud buckets.
[**IndexCloudBucketsByCloudPool**](CloudBucketsApi.md#IndexCloudBucketsByCloudPool) | **Get** /cloud_pools/{cloud_pool_id}/cloud_buckets | Lists all cloud buckets.
[**ShowCloudBucket**](CloudBucketsApi.md#ShowCloudBucket) | **Get** /cloud_buckets/{cloud_bucket_id} | Displays a specific cloud bucket.
[**ShowCloudBucketByCloudConnector**](CloudBucketsApi.md#ShowCloudBucketByCloudConnector) | **Get** /cloud_connectors/{cloud_connector_id}/cloud_buckets/{cloud_bucket_id} | Displays a specific cloud bucket.
[**ShowCloudBucketByCloudPool**](CloudBucketsApi.md#ShowCloudBucketByCloudPool) | **Get** /cloud_pools/{cloud_pool_id}/cloud_buckets/{cloud_bucket_id} | Displays a specific cloud bucket.
[**SyncCloudBuckets**](CloudBucketsApi.md#SyncCloudBuckets) | **Put** /cloud_connectors/{cloud_connector_id}/cloud_buckets/-/sync | Synchronize internal cloud buckets with their remote equivalent.
[**SyncResultCloudBuckets**](CloudBucketsApi.md#SyncResultCloudBuckets) | **Get** /cloud_connectors/{cloud_connector_id}/cloud_buckets/-/sync | Check result of cloud connector sync job.


# **IndexCloudBuckets**
> CloudBucketCollection IndexCloudBuckets(ctx, optional)
Lists all cloud buckets.

**API Key Scope**: cloud_buckets / index

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IndexCloudBucketsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IndexCloudBucketsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **cloudConnectorId** | **optional.String**| Filter on cloud connector id | 
 **cloudPoolId** | **optional.String**| Filter on cloud pool id | 
 **name** | **optional.String**| Filter on name | 
 **location** | **optional.String**| Filter on location | 
 **price** | **optional.String**| Filter on price | 

### Return type

[**CloudBucketCollection**](cloud_bucket_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IndexCloudBucketsByCloudConnector**
> CloudBucketCollection IndexCloudBucketsByCloudConnector(ctx, cloudConnectorId, optional)
Lists all cloud buckets.

**API Key Scope**: cloud_buckets / index

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudConnectorId** | **string**| Numeric ID or name of cloud connector. | 
 **optional** | ***IndexCloudBucketsByCloudConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IndexCloudBucketsByCloudConnectorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **cloudPoolId** | **optional.String**| Filter on cloud pool id | 
 **name** | **optional.String**| Filter on name | 
 **location** | **optional.String**| Filter on location | 
 **price** | **optional.String**| Filter on price | 

### Return type

[**CloudBucketCollection**](cloud_bucket_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IndexCloudBucketsByCloudPool**
> CloudBucketCollection IndexCloudBucketsByCloudPool(ctx, cloudPoolId, optional)
Lists all cloud buckets.

**API Key Scope**: cloud_buckets / index

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudPoolId** | **string**| Numeric ID or name of cloud pool. | 
 **optional** | ***IndexCloudBucketsByCloudPoolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IndexCloudBucketsByCloudPoolOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **cloudConnectorId** | **optional.String**| Filter on cloud connector id | 
 **name** | **optional.String**| Filter on name | 
 **location** | **optional.String**| Filter on location | 
 **price** | **optional.String**| Filter on price | 

### Return type

[**CloudBucketCollection**](cloud_bucket_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowCloudBucket**
> CloudBucket ShowCloudBucket(ctx, cloudBucketId)
Displays a specific cloud bucket.

**API Key Scope**: cloud_buckets / show

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudBucketId** | **string**| Numeric ID or name of cloud bucket. | 

### Return type

[**CloudBucket**](cloud_bucket.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowCloudBucketByCloudConnector**
> CloudBucket ShowCloudBucketByCloudConnector(ctx, cloudConnectorId, cloudBucketId)
Displays a specific cloud bucket.

**API Key Scope**: cloud_buckets / show

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudConnectorId** | **string**| Numeric ID or name of cloud connector. | 
  **cloudBucketId** | **string**| Numeric ID or name of cloud bucket. | 

### Return type

[**CloudBucket**](cloud_bucket.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowCloudBucketByCloudPool**
> CloudBucket ShowCloudBucketByCloudPool(ctx, cloudPoolId, cloudBucketId)
Displays a specific cloud bucket.

**API Key Scope**: cloud_buckets / show

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudPoolId** | **string**| Numeric ID or name of cloud pool. | 
  **cloudBucketId** | **string**| Numeric ID or name of cloud bucket. | 

### Return type

[**CloudBucket**](cloud_bucket.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncCloudBuckets**
> ActiveJobStatus SyncCloudBuckets(ctx, cloudConnectorId)
Synchronize internal cloud buckets with their remote equivalent.

**API Key Scope**: cloud_buckets / sync

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudConnectorId** | **string**| Numeric ID or name of cloud connector. | 

### Return type

[**ActiveJobStatus**](active_job_status.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncResultCloudBuckets**
> CloudBucketSimpleCollection SyncResultCloudBuckets(ctx, cloudConnectorId, optional)
Check result of cloud connector sync job.

**API Key Scope**: cloud_buckets / sync

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudConnectorId** | **string**| Numeric ID or name of cloud connector. | 
 **optional** | ***SyncResultCloudBucketsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SyncResultCloudBucketsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jobId** | **optional.String**| ID of active job | 

### Return type

[**CloudBucketSimpleCollection**](cloud_bucket_simple_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


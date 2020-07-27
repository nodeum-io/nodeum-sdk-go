# \StatisticsApi

All URIs are relative to *https://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StatisticsByDate**](StatisticsApi.md#StatisticsByDate) | **Get** /statistics/by_date | Get statistics about files, grouped by date
[**StatisticsByFileExtension**](StatisticsApi.md#StatisticsByFileExtension) | **Get** /statistics/by_file_extension | Get statistics about files, grouped by file extension
[**StatisticsByGroupOwner**](StatisticsApi.md#StatisticsByGroupOwner) | **Get** /statistics/by_group_owner | Get statistics about files, grouped by owner (group)
[**StatisticsByMetadata**](StatisticsApi.md#StatisticsByMetadata) | **Get** /statistics/by_metadata | Get statistics about files, grouped by metadata
[**StatisticsByPrimaryCloud**](StatisticsApi.md#StatisticsByPrimaryCloud) | **Get** /statistics/by_primary_cloud | Get statistics about files, grouped by primary Cloud
[**StatisticsByPrimaryName**](StatisticsApi.md#StatisticsByPrimaryName) | **Get** /statistics/by_primary_name | Get statistics about files, grouped by primary storages
[**StatisticsByPrimaryNas**](StatisticsApi.md#StatisticsByPrimaryNas) | **Get** /statistics/by_primary_nas | Get statistics about files, grouped by primary NAS
[**StatisticsByPrimaryStorage**](StatisticsApi.md#StatisticsByPrimaryStorage) | **Get** /statistics/by_primary_storage | Get statistics about files, grouped by primary storage
[**StatisticsBySecondaryCloud**](StatisticsApi.md#StatisticsBySecondaryCloud) | **Get** /statistics/by_secondary_cloud | Get statistics about files, grouped by secondary Cloud
[**StatisticsBySecondaryNas**](StatisticsApi.md#StatisticsBySecondaryNas) | **Get** /statistics/by_secondary_nas | Get statistics about files, grouped by secondary NAS
[**StatisticsBySecondaryStorage**](StatisticsApi.md#StatisticsBySecondaryStorage) | **Get** /statistics/by_secondary_storage | Get statistics about files, grouped by secondary storage
[**StatisticsBySecondaryTape**](StatisticsApi.md#StatisticsBySecondaryTape) | **Get** /statistics/by_secondary_tape | Get statistics about files, grouped by secondary Tape
[**StatisticsBySize**](StatisticsApi.md#StatisticsBySize) | **Get** /statistics/by_size | Get statistics about files, grouped by size
[**StatisticsByUserOwner**](StatisticsApi.md#StatisticsByUserOwner) | **Get** /statistics/by_user_owner | Get statistics about files, grouped by owner (user)
[**StatisticsStorage**](StatisticsApi.md#StatisticsStorage) | **Get** /statistics/storage | Get statistics about storages, grouped by types
[**StatisticsTaskByMetadata**](StatisticsApi.md#StatisticsTaskByMetadata) | **Get** /statistics/task_by_metadata | Get statistics about tasks executions, grouped by metadata
[**StatisticsTaskByStatus**](StatisticsApi.md#StatisticsTaskByStatus) | **Get** /statistics/task_by_status | Get statistics about tasks executions, grouped by status
[**StatisticsTaskByStorage**](StatisticsApi.md#StatisticsTaskByStorage) | **Get** /statistics/task_by_storage | Get statistics about tasks executions, grouped by source and destination
[**StatisticsTaskByWorkflow**](StatisticsApi.md#StatisticsTaskByWorkflow) | **Get** /statistics/task_by_workflow | Get statistics about tasks executions, grouped by workflow


# **StatisticsByDate**
> ByDateFacet StatisticsByDate(ctx, optional)
Get statistics about files, grouped by date

**API Key Scope**: statistics / by_date

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatisticsByDateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatisticsByDateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Solr query | 
 **fq** | [**optional.Interface of []string**](string.md)| Solr filter query  Multiple query can be separated by &#x60;|&#x60;. | 
 **dateAttr** | **optional.String**| Type of date to facet on | 

### Return type

[**ByDateFacet**](by_date_facet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatisticsByFileExtension**
> ByFileExtensionFacet StatisticsByFileExtension(ctx, optional)
Get statistics about files, grouped by file extension

**API Key Scope**: statistics / by_file_extension

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatisticsByFileExtensionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatisticsByFileExtensionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Solr query | 
 **fq** | [**optional.Interface of []string**](string.md)| Solr filter query  Multiple query can be separated by &#x60;|&#x60;. | 
 **dateAttr** | **optional.String**| Type of date to facet on | 
 **sort** | **optional.String**| Sort results of facet | [default to count]
 **limit** | **optional.Int32**| Limit results of facet | [default to 10]

### Return type

[**ByFileExtensionFacet**](by_file_extension_facet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatisticsByGroupOwner**
> ByGroupOwnerFacet StatisticsByGroupOwner(ctx, optional)
Get statistics about files, grouped by owner (group)

**API Key Scope**: statistics / by_group_owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatisticsByGroupOwnerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatisticsByGroupOwnerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Solr query | 
 **fq** | [**optional.Interface of []string**](string.md)| Solr filter query  Multiple query can be separated by &#x60;|&#x60;. | 
 **dateAttr** | **optional.String**| Type of date to facet on | 
 **sort** | **optional.String**| Sort results of facet | [default to count]
 **limit** | **optional.Int32**| Limit results of facet | [default to 10]

### Return type

[**ByGroupOwnerFacet**](by_group_owner_facet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatisticsByMetadata**
> ByMetadataFacet StatisticsByMetadata(ctx, optional)
Get statistics about files, grouped by metadata

**API Key Scope**: statistics / by_metadata

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatisticsByMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatisticsByMetadataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Solr query | 
 **fq** | [**optional.Interface of []string**](string.md)| Solr filter query  Multiple query can be separated by &#x60;|&#x60;. | 
 **dateAttr** | **optional.String**| Type of date to facet on | 
 **sort** | **optional.String**| Sort results of facet | [default to count]
 **limit** | **optional.Int32**| Limit results of facet | [default to 10]

### Return type

[**ByMetadataFacet**](by_metadata_facet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatisticsByPrimaryCloud**
> ByPrimaryCloudFacet StatisticsByPrimaryCloud(ctx, optional)
Get statistics about files, grouped by primary Cloud

**API Key Scope**: statistics / by_primary_cloud

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatisticsByPrimaryCloudOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatisticsByPrimaryCloudOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Solr query | 
 **fq** | [**optional.Interface of []string**](string.md)| Solr filter query  Multiple query can be separated by &#x60;|&#x60;. | 
 **dateAttr** | **optional.String**| Type of date to facet on | 
 **sort** | **optional.String**| Sort results of facet | [default to count]
 **limit** | **optional.Int32**| Limit results of facet | [default to 10]

### Return type

[**ByPrimaryCloudFacet**](by_primary_cloud_facet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatisticsByPrimaryName**
> ByPrimaryFacet StatisticsByPrimaryName(ctx, optional)
Get statistics about files, grouped by primary storages

**API Key Scope**: statistics / by_primary_name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatisticsByPrimaryNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatisticsByPrimaryNameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Solr query | 
 **fq** | [**optional.Interface of []string**](string.md)| Solr filter query  Multiple query can be separated by &#x60;|&#x60;. | 
 **dateAttr** | **optional.String**| Type of date to facet on | 
 **sort** | **optional.String**| Sort results of facet | [default to count]
 **limit** | **optional.Int32**| Limit results of facet | [default to 10]

### Return type

[**ByPrimaryFacet**](by_primary_facet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatisticsByPrimaryNas**
> ByPrimaryNasFacet StatisticsByPrimaryNas(ctx, optional)
Get statistics about files, grouped by primary NAS

**API Key Scope**: statistics / by_primary_nas

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatisticsByPrimaryNasOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatisticsByPrimaryNasOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Solr query | 
 **fq** | [**optional.Interface of []string**](string.md)| Solr filter query  Multiple query can be separated by &#x60;|&#x60;. | 
 **dateAttr** | **optional.String**| Type of date to facet on | 
 **sort** | **optional.String**| Sort results of facet | [default to count]
 **limit** | **optional.Int32**| Limit results of facet | [default to 10]

### Return type

[**ByPrimaryNasFacet**](by_primary_nas_facet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatisticsByPrimaryStorage**
> ByPrimaryStorageFacet StatisticsByPrimaryStorage(ctx, optional)
Get statistics about files, grouped by primary storage

**API Key Scope**: statistics / by_primary_storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatisticsByPrimaryStorageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatisticsByPrimaryStorageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Solr query | 
 **fq** | [**optional.Interface of []string**](string.md)| Solr filter query  Multiple query can be separated by &#x60;|&#x60;. | 
 **dateAttr** | **optional.String**| Type of date to facet on | 
 **sort** | **optional.String**| Sort results of facet | [default to count]
 **limit** | **optional.Int32**| Limit results of facet | [default to 10]

### Return type

[**ByPrimaryStorageFacet**](by_primary_storage_facet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatisticsBySecondaryCloud**
> BySecondaryCloudFacet StatisticsBySecondaryCloud(ctx, optional)
Get statistics about files, grouped by secondary Cloud

**API Key Scope**: statistics / by_secondary_cloud

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatisticsBySecondaryCloudOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatisticsBySecondaryCloudOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Solr query | 
 **fq** | [**optional.Interface of []string**](string.md)| Solr filter query  Multiple query can be separated by &#x60;|&#x60;. | 
 **dateAttr** | **optional.String**| Type of date to facet on | 
 **sort** | **optional.String**| Sort results of facet | [default to count]
 **limit** | **optional.Int32**| Limit results of facet | [default to 10]

### Return type

[**BySecondaryCloudFacet**](by_secondary_cloud_facet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatisticsBySecondaryNas**
> BySecondaryNasFacet StatisticsBySecondaryNas(ctx, optional)
Get statistics about files, grouped by secondary NAS

**API Key Scope**: statistics / by_secondary_nas

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatisticsBySecondaryNasOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatisticsBySecondaryNasOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Solr query | 
 **fq** | [**optional.Interface of []string**](string.md)| Solr filter query  Multiple query can be separated by &#x60;|&#x60;. | 
 **dateAttr** | **optional.String**| Type of date to facet on | 
 **sort** | **optional.String**| Sort results of facet | [default to count]
 **limit** | **optional.Int32**| Limit results of facet | [default to 10]

### Return type

[**BySecondaryNasFacet**](by_secondary_nas_facet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatisticsBySecondaryStorage**
> BySecondaryStorageFacet StatisticsBySecondaryStorage(ctx, optional)
Get statistics about files, grouped by secondary storage

**API Key Scope**: statistics / by_secondary_storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatisticsBySecondaryStorageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatisticsBySecondaryStorageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Solr query | 
 **fq** | [**optional.Interface of []string**](string.md)| Solr filter query  Multiple query can be separated by &#x60;|&#x60;. | 
 **dateAttr** | **optional.String**| Type of date to facet on | 
 **sort** | **optional.String**| Sort results of facet | [default to count]
 **limit** | **optional.Int32**| Limit results of facet | [default to 10]

### Return type

[**BySecondaryStorageFacet**](by_secondary_storage_facet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatisticsBySecondaryTape**
> BySecondaryTapeFacet StatisticsBySecondaryTape(ctx, optional)
Get statistics about files, grouped by secondary Tape

**API Key Scope**: statistics / by_secondary_tape

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatisticsBySecondaryTapeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatisticsBySecondaryTapeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Solr query | 
 **fq** | [**optional.Interface of []string**](string.md)| Solr filter query  Multiple query can be separated by &#x60;|&#x60;. | 
 **dateAttr** | **optional.String**| Type of date to facet on | 
 **sort** | **optional.String**| Sort results of facet | [default to count]
 **limit** | **optional.Int32**| Limit results of facet | [default to 10]

### Return type

[**BySecondaryTapeFacet**](by_secondary_tape_facet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatisticsBySize**
> BySizeFacet StatisticsBySize(ctx, optional)
Get statistics about files, grouped by size

**API Key Scope**: statistics / by_size

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatisticsBySizeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatisticsBySizeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Solr query | 
 **fq** | [**optional.Interface of []string**](string.md)| Solr filter query  Multiple query can be separated by &#x60;|&#x60;. | 
 **dateAttr** | **optional.String**| Type of date to facet on | 

### Return type

[**BySizeFacet**](by_size_facet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatisticsByUserOwner**
> ByUserOwnerFacet StatisticsByUserOwner(ctx, optional)
Get statistics about files, grouped by owner (user)

**API Key Scope**: statistics / by_user_owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatisticsByUserOwnerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatisticsByUserOwnerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Solr query | 
 **fq** | [**optional.Interface of []string**](string.md)| Solr filter query  Multiple query can be separated by &#x60;|&#x60;. | 
 **dateAttr** | **optional.String**| Type of date to facet on | 
 **sort** | **optional.String**| Sort results of facet | [default to count]
 **limit** | **optional.Int32**| Limit results of facet | [default to 10]

### Return type

[**ByUserOwnerFacet**](by_user_owner_facet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatisticsStorage**
> StorageFacet StatisticsStorage(ctx, optional)
Get statistics about storages, grouped by types

**API Key Scope**: statistics / storages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatisticsStorageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatisticsStorageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Solr query | 
 **fq** | [**optional.Interface of []string**](string.md)| Solr filter query  Multiple query can be separated by &#x60;|&#x60;. | 

### Return type

[**StorageFacet**](storage_facet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatisticsTaskByMetadata**
> ByTaskMetadataFacet StatisticsTaskByMetadata(ctx, optional)
Get statistics about tasks executions, grouped by metadata

**API Key Scope**: statistics / task_by_metadata

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatisticsTaskByMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatisticsTaskByMetadataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Solr query | 
 **fq** | [**optional.Interface of []string**](string.md)| Solr filter query  Multiple query can be separated by &#x60;|&#x60;. | 
 **sort** | **optional.String**| Sort results of facet on task | [default to count]
 **limit** | **optional.Int32**| Limit results of facet | [default to 10]

### Return type

[**ByTaskMetadataFacet**](by_task_metadata_facet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatisticsTaskByStatus**
> ByTaskStatusFacet StatisticsTaskByStatus(ctx, optional)
Get statistics about tasks executions, grouped by status

**API Key Scope**: statistics / task_by_status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatisticsTaskByStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatisticsTaskByStatusOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Solr query | 
 **fq** | [**optional.Interface of []string**](string.md)| Solr filter query  Multiple query can be separated by &#x60;|&#x60;. | 

### Return type

[**ByTaskStatusFacet**](by_task_status_facet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatisticsTaskByStorage**
> ByTaskStorageFacet StatisticsTaskByStorage(ctx, optional)
Get statistics about tasks executions, grouped by source and destination

**API Key Scope**: statistics / task_by_storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatisticsTaskByStorageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatisticsTaskByStorageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Solr query | 
 **fq** | [**optional.Interface of []string**](string.md)| Solr filter query  Multiple query can be separated by &#x60;|&#x60;. | 

### Return type

[**ByTaskStorageFacet**](by_task_storage_facet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatisticsTaskByWorkflow**
> ByTaskWorkflowFacet StatisticsTaskByWorkflow(ctx, optional)
Get statistics about tasks executions, grouped by workflow

**API Key Scope**: statistics / task_by_workflow

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatisticsTaskByWorkflowOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatisticsTaskByWorkflowOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Solr query | 
 **fq** | [**optional.Interface of []string**](string.md)| Solr filter query  Multiple query can be separated by &#x60;|&#x60;. | 

### Return type

[**ByTaskWorkflowFacet**](by_task_workflow_facet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


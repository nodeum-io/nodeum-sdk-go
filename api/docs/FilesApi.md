# \FilesApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesChildren**](FilesApi.md#FilesChildren) | **Get** /files/{file_parent_id}/children | Lists files under a specific folder.
[**FilesChildrenByCloudPool**](FilesApi.md#FilesChildrenByCloudPool) | **Get** /cloud_pools/{cloud_pool_id}/files/{file_parent_id}/children | Lists files under a specific folder.
[**FilesChildrenByContainer**](FilesApi.md#FilesChildrenByContainer) | **Get** /containers/{container_id}/files/{file_parent_id}/children | Lists files under a specific folder.
[**FilesChildrenByNasPool**](FilesApi.md#FilesChildrenByNasPool) | **Get** /nas_pools/{nas_pool_id}/files/{file_parent_id}/children | Lists files under a specific folder.
[**FilesChildrenByTapePool**](FilesApi.md#FilesChildrenByTapePool) | **Get** /tape_pools/{tape_pool_id}/files/{file_parent_id}/children | Lists files under a specific folder.
[**FilesChildrenByTask**](FilesApi.md#FilesChildrenByTask) | **Get** /tasks/{task_id}/files/{file_parent_id}/children | Lists files under a specific folder.
[**FilesChildrenByTaskExecution**](FilesApi.md#FilesChildrenByTaskExecution) | **Get** /task_executions/{task_execution_id}/files/{file_parent_id}/children | Lists files under a specific folder.
[**FilesChildrenByTaskExecutionByTask**](FilesApi.md#FilesChildrenByTaskExecutionByTask) | **Get** /tasks/{task_id}/task_executions/{task_execution_id}/files/{file_parent_id}/children | Lists files under a specific folder.
[**IndexFiles**](FilesApi.md#IndexFiles) | **Get** /files | Lists files on root.
[**IndexFilesByCloudPool**](FilesApi.md#IndexFilesByCloudPool) | **Get** /cloud_pools/{cloud_pool_id}/files | Lists files on root.
[**IndexFilesByContainer**](FilesApi.md#IndexFilesByContainer) | **Get** /containers/{container_id}/files | Lists files on root.
[**IndexFilesByNasPool**](FilesApi.md#IndexFilesByNasPool) | **Get** /nas_pools/{nas_pool_id}/files | Lists files on root.
[**IndexFilesByTapePool**](FilesApi.md#IndexFilesByTapePool) | **Get** /tape_pools/{tape_pool_id}/files | Lists files on root.
[**IndexFilesByTask**](FilesApi.md#IndexFilesByTask) | **Get** /tasks/{task_id}/files | Lists files on root.
[**IndexFilesByTaskExecution**](FilesApi.md#IndexFilesByTaskExecution) | **Get** /task_executions/{task_execution_id}/files | Lists files on root.
[**IndexFilesByTaskExecutionByTask**](FilesApi.md#IndexFilesByTaskExecutionByTask) | **Get** /tasks/{task_id}/task_executions/{task_execution_id}/files | Lists files on root.
[**ShowFile**](FilesApi.md#ShowFile) | **Get** /files/{file_id} | Displays a specific file.
[**ShowFileByCloudPool**](FilesApi.md#ShowFileByCloudPool) | **Get** /cloud_pools/{cloud_pool_id}/files/{file_id} | Displays a specific file.
[**ShowFileByContainer**](FilesApi.md#ShowFileByContainer) | **Get** /containers/{container_id}/files/{file_id} | Displays a specific file.
[**ShowFileByNasPool**](FilesApi.md#ShowFileByNasPool) | **Get** /nas_pools/{nas_pool_id}/files/{file_id} | Displays a specific file.
[**ShowFileByTapePool**](FilesApi.md#ShowFileByTapePool) | **Get** /tape_pools/{tape_pool_id}/files/{file_id} | Displays a specific file.
[**ShowFileByTask**](FilesApi.md#ShowFileByTask) | **Get** /tasks/{task_id}/files/{file_id} | Displays a specific file.
[**ShowFileByTaskExecution**](FilesApi.md#ShowFileByTaskExecution) | **Get** /task_executions/{task_execution_id}/files/{file_id} | Displays a specific file.
[**ShowFileByTaskExecutionByTask**](FilesApi.md#ShowFileByTaskExecutionByTask) | **Get** /tasks/{task_id}/task_executions/{task_execution_id}/files/{file_id} | Displays a specific file.



## FilesChildren

> NodeumFileCollection FilesChildren(ctx, fileParentId, optional)

Lists files under a specific folder.

**API Key Scope**: files / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileParentId** | **int32**| Numeric ID of parent folder. | 
 **optional** | ***FilesChildrenOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FilesChildrenOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **fileId** | **optional.String**| Filter on file id | 
 **name** | **optional.String**| Filter on name | 
 **type_** | **optional.String**| Filter on type | 
 **permission** | **optional.String**| Filter on permission | 
 **size** | **optional.String**| Filter on size | 
 **changeDate** | **optional.String**| Filter on change date | 
 **modificationDate** | **optional.String**| Filter on modification date | 
 **accessDate** | **optional.String**| Filter on access date | 
 **gid** | **optional.String**| Filter on gid | 
 **uid** | **optional.String**| Filter on uid | 

### Return type

[**NodeumFileCollection**](nodeum_file_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesChildrenByCloudPool

> NodeumFileCollection FilesChildrenByCloudPool(ctx, cloudPoolId, fileParentId, optional)

Lists files under a specific folder.

**API Key Scope**: files / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudPoolId** | **string**| Numeric ID or name of cloud pool. | 
**fileParentId** | **int32**| Numeric ID of parent folder. | 
 **optional** | ***FilesChildrenByCloudPoolOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FilesChildrenByCloudPoolOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **fileId** | **optional.String**| Filter on file id | 
 **name** | **optional.String**| Filter on name | 
 **type_** | **optional.String**| Filter on type | 
 **permission** | **optional.String**| Filter on permission | 
 **size** | **optional.String**| Filter on size | 
 **changeDate** | **optional.String**| Filter on change date | 
 **modificationDate** | **optional.String**| Filter on modification date | 
 **accessDate** | **optional.String**| Filter on access date | 
 **gid** | **optional.String**| Filter on gid | 
 **uid** | **optional.String**| Filter on uid | 

### Return type

[**NodeumFileCollection**](nodeum_file_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesChildrenByContainer

> NodeumFileCollection FilesChildrenByContainer(ctx, containerId, fileParentId, optional)

Lists files under a specific folder.

**API Key Scope**: files / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string**| Numeric ID or name of container. | 
**fileParentId** | **int32**| Numeric ID of parent folder. | 
 **optional** | ***FilesChildrenByContainerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FilesChildrenByContainerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **fileId** | **optional.String**| Filter on file id | 
 **name** | **optional.String**| Filter on name | 
 **type_** | **optional.String**| Filter on type | 
 **permission** | **optional.String**| Filter on permission | 
 **size** | **optional.String**| Filter on size | 
 **changeDate** | **optional.String**| Filter on change date | 
 **modificationDate** | **optional.String**| Filter on modification date | 
 **accessDate** | **optional.String**| Filter on access date | 
 **gid** | **optional.String**| Filter on gid | 
 **uid** | **optional.String**| Filter on uid | 

### Return type

[**NodeumFileCollection**](nodeum_file_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesChildrenByNasPool

> NodeumFileCollection FilesChildrenByNasPool(ctx, nasPoolId, fileParentId, optional)

Lists files under a specific folder.

**API Key Scope**: files / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasPoolId** | **string**| Numeric ID or name of NAS pool. | 
**fileParentId** | **int32**| Numeric ID of parent folder. | 
 **optional** | ***FilesChildrenByNasPoolOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FilesChildrenByNasPoolOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **fileId** | **optional.String**| Filter on file id | 
 **name** | **optional.String**| Filter on name | 
 **type_** | **optional.String**| Filter on type | 
 **permission** | **optional.String**| Filter on permission | 
 **size** | **optional.String**| Filter on size | 
 **changeDate** | **optional.String**| Filter on change date | 
 **modificationDate** | **optional.String**| Filter on modification date | 
 **accessDate** | **optional.String**| Filter on access date | 
 **gid** | **optional.String**| Filter on gid | 
 **uid** | **optional.String**| Filter on uid | 

### Return type

[**NodeumFileCollection**](nodeum_file_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesChildrenByTapePool

> NodeumFileCollection FilesChildrenByTapePool(ctx, tapePoolId, fileParentId, optional)

Lists files under a specific folder.

**API Key Scope**: files / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tapePoolId** | **string**| Numeric ID, or name of tape pool. | 
**fileParentId** | **int32**| Numeric ID of parent folder. | 
 **optional** | ***FilesChildrenByTapePoolOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FilesChildrenByTapePoolOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **fileId** | **optional.String**| Filter on file id | 
 **name** | **optional.String**| Filter on name | 
 **type_** | **optional.String**| Filter on type | 
 **permission** | **optional.String**| Filter on permission | 
 **size** | **optional.String**| Filter on size | 
 **changeDate** | **optional.String**| Filter on change date | 
 **modificationDate** | **optional.String**| Filter on modification date | 
 **accessDate** | **optional.String**| Filter on access date | 
 **gid** | **optional.String**| Filter on gid | 
 **uid** | **optional.String**| Filter on uid | 

### Return type

[**NodeumFileCollection**](nodeum_file_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesChildrenByTask

> NodeumFileCollection FilesChildrenByTask(ctx, taskId, fileParentId, optional)

Lists files under a specific folder.

**API Key Scope**: files / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
**fileParentId** | **int32**| Numeric ID of parent folder. | 
 **optional** | ***FilesChildrenByTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FilesChildrenByTaskOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **fileId** | **optional.String**| Filter on file id | 
 **name** | **optional.String**| Filter on name | 
 **type_** | **optional.String**| Filter on type | 
 **permission** | **optional.String**| Filter on permission | 
 **size** | **optional.String**| Filter on size | 
 **changeDate** | **optional.String**| Filter on change date | 
 **modificationDate** | **optional.String**| Filter on modification date | 
 **accessDate** | **optional.String**| Filter on access date | 
 **gid** | **optional.String**| Filter on gid | 
 **uid** | **optional.String**| Filter on uid | 

### Return type

[**NodeumFileCollection**](nodeum_file_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesChildrenByTaskExecution

> NodeumFileCollection FilesChildrenByTaskExecution(ctx, taskExecutionId, fileParentId, optional)

Lists files under a specific folder.

**API Key Scope**: files / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskExecutionId** | **int64**| Numeric ID of task execution. | 
**fileParentId** | **int32**| Numeric ID of parent folder. | 
 **optional** | ***FilesChildrenByTaskExecutionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FilesChildrenByTaskExecutionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **fileId** | **optional.String**| Filter on file id | 
 **name** | **optional.String**| Filter on name | 
 **type_** | **optional.String**| Filter on type | 
 **permission** | **optional.String**| Filter on permission | 
 **size** | **optional.String**| Filter on size | 
 **changeDate** | **optional.String**| Filter on change date | 
 **modificationDate** | **optional.String**| Filter on modification date | 
 **accessDate** | **optional.String**| Filter on access date | 
 **gid** | **optional.String**| Filter on gid | 
 **uid** | **optional.String**| Filter on uid | 

### Return type

[**NodeumFileCollection**](nodeum_file_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesChildrenByTaskExecutionByTask

> NodeumFileCollection FilesChildrenByTaskExecutionByTask(ctx, taskId, taskExecutionId, fileParentId, optional)

Lists files under a specific folder.

**API Key Scope**: files / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
**taskExecutionId** | **int64**| Numeric ID of task execution. | 
**fileParentId** | **int32**| Numeric ID of parent folder. | 
 **optional** | ***FilesChildrenByTaskExecutionByTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FilesChildrenByTaskExecutionByTaskOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **fileId** | **optional.String**| Filter on file id | 
 **name** | **optional.String**| Filter on name | 
 **type_** | **optional.String**| Filter on type | 
 **permission** | **optional.String**| Filter on permission | 
 **size** | **optional.String**| Filter on size | 
 **changeDate** | **optional.String**| Filter on change date | 
 **modificationDate** | **optional.String**| Filter on modification date | 
 **accessDate** | **optional.String**| Filter on access date | 
 **gid** | **optional.String**| Filter on gid | 
 **uid** | **optional.String**| Filter on uid | 

### Return type

[**NodeumFileCollection**](nodeum_file_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexFiles

> NodeumFileCollection IndexFiles(ctx, optional)

Lists files on root.

**API Key Scope**: files / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IndexFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexFilesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **fileId** | **optional.String**| Filter on file id | 
 **name** | **optional.String**| Filter on name | 
 **type_** | **optional.String**| Filter on type | 
 **permission** | **optional.String**| Filter on permission | 
 **size** | **optional.String**| Filter on size | 
 **changeDate** | **optional.String**| Filter on change date | 
 **modificationDate** | **optional.String**| Filter on modification date | 
 **accessDate** | **optional.String**| Filter on access date | 
 **gid** | **optional.String**| Filter on gid | 
 **uid** | **optional.String**| Filter on uid | 

### Return type

[**NodeumFileCollection**](nodeum_file_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexFilesByCloudPool

> NodeumFileCollection IndexFilesByCloudPool(ctx, cloudPoolId, optional)

Lists files on root.

**API Key Scope**: files / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudPoolId** | **string**| Numeric ID or name of cloud pool. | 
 **optional** | ***IndexFilesByCloudPoolOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexFilesByCloudPoolOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **fileId** | **optional.String**| Filter on file id | 
 **name** | **optional.String**| Filter on name | 
 **type_** | **optional.String**| Filter on type | 
 **permission** | **optional.String**| Filter on permission | 
 **size** | **optional.String**| Filter on size | 
 **changeDate** | **optional.String**| Filter on change date | 
 **modificationDate** | **optional.String**| Filter on modification date | 
 **accessDate** | **optional.String**| Filter on access date | 
 **gid** | **optional.String**| Filter on gid | 
 **uid** | **optional.String**| Filter on uid | 

### Return type

[**NodeumFileCollection**](nodeum_file_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexFilesByContainer

> NodeumFileCollection IndexFilesByContainer(ctx, containerId, optional)

Lists files on root.

**API Key Scope**: files / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string**| Numeric ID or name of container. | 
 **optional** | ***IndexFilesByContainerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexFilesByContainerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **fileId** | **optional.String**| Filter on file id | 
 **name** | **optional.String**| Filter on name | 
 **type_** | **optional.String**| Filter on type | 
 **permission** | **optional.String**| Filter on permission | 
 **size** | **optional.String**| Filter on size | 
 **changeDate** | **optional.String**| Filter on change date | 
 **modificationDate** | **optional.String**| Filter on modification date | 
 **accessDate** | **optional.String**| Filter on access date | 
 **gid** | **optional.String**| Filter on gid | 
 **uid** | **optional.String**| Filter on uid | 

### Return type

[**NodeumFileCollection**](nodeum_file_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexFilesByNasPool

> NodeumFileCollection IndexFilesByNasPool(ctx, nasPoolId, optional)

Lists files on root.

**API Key Scope**: files / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasPoolId** | **string**| Numeric ID or name of NAS pool. | 
 **optional** | ***IndexFilesByNasPoolOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexFilesByNasPoolOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **fileId** | **optional.String**| Filter on file id | 
 **name** | **optional.String**| Filter on name | 
 **type_** | **optional.String**| Filter on type | 
 **permission** | **optional.String**| Filter on permission | 
 **size** | **optional.String**| Filter on size | 
 **changeDate** | **optional.String**| Filter on change date | 
 **modificationDate** | **optional.String**| Filter on modification date | 
 **accessDate** | **optional.String**| Filter on access date | 
 **gid** | **optional.String**| Filter on gid | 
 **uid** | **optional.String**| Filter on uid | 

### Return type

[**NodeumFileCollection**](nodeum_file_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexFilesByTapePool

> NodeumFileCollection IndexFilesByTapePool(ctx, tapePoolId, optional)

Lists files on root.

**API Key Scope**: files / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tapePoolId** | **string**| Numeric ID, or name of tape pool. | 
 **optional** | ***IndexFilesByTapePoolOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexFilesByTapePoolOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **fileId** | **optional.String**| Filter on file id | 
 **name** | **optional.String**| Filter on name | 
 **type_** | **optional.String**| Filter on type | 
 **permission** | **optional.String**| Filter on permission | 
 **size** | **optional.String**| Filter on size | 
 **changeDate** | **optional.String**| Filter on change date | 
 **modificationDate** | **optional.String**| Filter on modification date | 
 **accessDate** | **optional.String**| Filter on access date | 
 **gid** | **optional.String**| Filter on gid | 
 **uid** | **optional.String**| Filter on uid | 

### Return type

[**NodeumFileCollection**](nodeum_file_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexFilesByTask

> NodeumFileCollection IndexFilesByTask(ctx, taskId, optional)

Lists files on root.

**API Key Scope**: files / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
 **optional** | ***IndexFilesByTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexFilesByTaskOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **fileId** | **optional.String**| Filter on file id | 
 **name** | **optional.String**| Filter on name | 
 **type_** | **optional.String**| Filter on type | 
 **permission** | **optional.String**| Filter on permission | 
 **size** | **optional.String**| Filter on size | 
 **changeDate** | **optional.String**| Filter on change date | 
 **modificationDate** | **optional.String**| Filter on modification date | 
 **accessDate** | **optional.String**| Filter on access date | 
 **gid** | **optional.String**| Filter on gid | 
 **uid** | **optional.String**| Filter on uid | 

### Return type

[**NodeumFileCollection**](nodeum_file_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexFilesByTaskExecution

> NodeumFileCollection IndexFilesByTaskExecution(ctx, taskExecutionId, optional)

Lists files on root.

**API Key Scope**: files / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskExecutionId** | **int64**| Numeric ID of task execution. | 
 **optional** | ***IndexFilesByTaskExecutionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexFilesByTaskExecutionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **fileId** | **optional.String**| Filter on file id | 
 **name** | **optional.String**| Filter on name | 
 **type_** | **optional.String**| Filter on type | 
 **permission** | **optional.String**| Filter on permission | 
 **size** | **optional.String**| Filter on size | 
 **changeDate** | **optional.String**| Filter on change date | 
 **modificationDate** | **optional.String**| Filter on modification date | 
 **accessDate** | **optional.String**| Filter on access date | 
 **gid** | **optional.String**| Filter on gid | 
 **uid** | **optional.String**| Filter on uid | 

### Return type

[**NodeumFileCollection**](nodeum_file_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexFilesByTaskExecutionByTask

> NodeumFileCollection IndexFilesByTaskExecutionByTask(ctx, taskId, taskExecutionId, optional)

Lists files on root.

**API Key Scope**: files / index

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
**taskExecutionId** | **int64**| Numeric ID of task execution. | 
 **optional** | ***IndexFilesByTaskExecutionByTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IndexFilesByTaskExecutionByTaskOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **fileId** | **optional.String**| Filter on file id | 
 **name** | **optional.String**| Filter on name | 
 **type_** | **optional.String**| Filter on type | 
 **permission** | **optional.String**| Filter on permission | 
 **size** | **optional.String**| Filter on size | 
 **changeDate** | **optional.String**| Filter on change date | 
 **modificationDate** | **optional.String**| Filter on modification date | 
 **accessDate** | **optional.String**| Filter on access date | 
 **gid** | **optional.String**| Filter on gid | 
 **uid** | **optional.String**| Filter on uid | 

### Return type

[**NodeumFileCollection**](nodeum_file_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowFile

> NodeumFileWithPath ShowFile(ctx, fileId)

Displays a specific file.

**API Key Scope**: files / show

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32**| Numeric ID of file. | 

### Return type

[**NodeumFileWithPath**](nodeum_file_with_path.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowFileByCloudPool

> NodeumFileWithPath ShowFileByCloudPool(ctx, cloudPoolId, fileId)

Displays a specific file.

**API Key Scope**: files / show

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudPoolId** | **string**| Numeric ID or name of cloud pool. | 
**fileId** | **int32**| Numeric ID of file. | 

### Return type

[**NodeumFileWithPath**](nodeum_file_with_path.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowFileByContainer

> NodeumFileWithPath ShowFileByContainer(ctx, containerId, fileId)

Displays a specific file.

**API Key Scope**: files / show

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string**| Numeric ID or name of container. | 
**fileId** | **int32**| Numeric ID of file. | 

### Return type

[**NodeumFileWithPath**](nodeum_file_with_path.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowFileByNasPool

> NodeumFileWithPath ShowFileByNasPool(ctx, nasPoolId, fileId)

Displays a specific file.

**API Key Scope**: files / show

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nasPoolId** | **string**| Numeric ID or name of NAS pool. | 
**fileId** | **int32**| Numeric ID of file. | 

### Return type

[**NodeumFileWithPath**](nodeum_file_with_path.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowFileByTapePool

> NodeumFileWithPath ShowFileByTapePool(ctx, tapePoolId, fileId)

Displays a specific file.

**API Key Scope**: files / show

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tapePoolId** | **string**| Numeric ID, or name of tape pool. | 
**fileId** | **int32**| Numeric ID of file. | 

### Return type

[**NodeumFileWithPath**](nodeum_file_with_path.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowFileByTask

> NodeumFileWithPath ShowFileByTask(ctx, taskId, fileId)

Displays a specific file.

**API Key Scope**: files / show

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
**fileId** | **int32**| Numeric ID of file. | 

### Return type

[**NodeumFileWithPath**](nodeum_file_with_path.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowFileByTaskExecution

> NodeumFileWithPath ShowFileByTaskExecution(ctx, taskExecutionId, fileId)

Displays a specific file.

**API Key Scope**: files / show

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskExecutionId** | **int64**| Numeric ID of task execution. | 
**fileId** | **int32**| Numeric ID of file. | 

### Return type

[**NodeumFileWithPath**](nodeum_file_with_path.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowFileByTaskExecutionByTask

> NodeumFileWithPath ShowFileByTaskExecutionByTask(ctx, taskId, taskExecutionId, fileId)

Displays a specific file.

**API Key Scope**: files / show

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
**taskExecutionId** | **int64**| Numeric ID of task execution. | 
**fileId** | **int32**| Numeric ID of file. | 

### Return type

[**NodeumFileWithPath**](nodeum_file_with_path.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


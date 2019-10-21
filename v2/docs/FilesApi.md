# \FilesApi

All URIs are relative to *https://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesChildren**](FilesApi.md#FilesChildren) | **Get** /files/{file_parent_id}/children | Lists files under a specific folder.
[**FilesChildrenByContainer**](FilesApi.md#FilesChildrenByContainer) | **Get** /containers/{container_id}/files/{file_parent_id}/children | Lists files under a specific folder.
[**FilesChildrenByPool**](FilesApi.md#FilesChildrenByPool) | **Get** /pools/{pool_id}/files/{file_parent_id}/children | Lists files under a specific folder.
[**FilesChildrenByTask**](FilesApi.md#FilesChildrenByTask) | **Get** /tasks/{task_id}/files/{file_parent_id}/children | Lists files under a specific folder.
[**FilesChildrenByTaskExecution**](FilesApi.md#FilesChildrenByTaskExecution) | **Get** /task_executions/{task_execution_id}/files/{file_parent_id}/children | Lists files under a specific folder.
[**FilesChildrenByTaskExecutionByTask**](FilesApi.md#FilesChildrenByTaskExecutionByTask) | **Get** /tasks/{task_id}/task_executions/{task_execution_id}/files/{file_parent_id}/children | Lists files under a specific folder.
[**IndexFiles**](FilesApi.md#IndexFiles) | **Get** /files | Lists files on root.
[**IndexFilesByContainer**](FilesApi.md#IndexFilesByContainer) | **Get** /containers/{container_id}/files | Lists files on root.
[**IndexFilesByPool**](FilesApi.md#IndexFilesByPool) | **Get** /pools/{pool_id}/files | Lists files on root.
[**IndexFilesByTask**](FilesApi.md#IndexFilesByTask) | **Get** /tasks/{task_id}/files | Lists files on root.
[**IndexFilesByTaskExecution**](FilesApi.md#IndexFilesByTaskExecution) | **Get** /task_executions/{task_execution_id}/files | Lists files on root.
[**IndexFilesByTaskExecutionByTask**](FilesApi.md#IndexFilesByTaskExecutionByTask) | **Get** /tasks/{task_id}/task_executions/{task_execution_id}/files | Lists files on root.
[**ShowFile**](FilesApi.md#ShowFile) | **Get** /files/{file_id} | Displays a specific file.
[**ShowFileByContainer**](FilesApi.md#ShowFileByContainer) | **Get** /containers/{container_id}/files/{file_id} | Displays a specific file.
[**ShowFileByPool**](FilesApi.md#ShowFileByPool) | **Get** /pools/{pool_id}/files/{file_id} | Displays a specific file.
[**ShowFileByTask**](FilesApi.md#ShowFileByTask) | **Get** /tasks/{task_id}/files/{file_id} | Displays a specific file.
[**ShowFileByTaskExecution**](FilesApi.md#ShowFileByTaskExecution) | **Get** /task_executions/{task_execution_id}/files/{file_id} | Displays a specific file.
[**ShowFileByTaskExecutionByTask**](FilesApi.md#ShowFileByTaskExecutionByTask) | **Get** /tasks/{task_id}/task_executions/{task_execution_id}/files/{file_id} | Displays a specific file.


# **FilesChildren**
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

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilesChildrenByContainer**
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

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilesChildrenByPool**
> NodeumFileCollection FilesChildrenByPool(ctx, poolId, fileParentId, optional)
Lists files under a specific folder.

**API Key Scope**: files / index

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**| Numeric ID, or name of pool. | 
  **fileParentId** | **int32**| Numeric ID of parent folder. | 
 **optional** | ***FilesChildrenByPoolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilesChildrenByPoolOpts struct

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

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilesChildrenByTask**
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

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilesChildrenByTaskExecution**
> NodeumFileCollection FilesChildrenByTaskExecution(ctx, taskExecutionId, fileParentId, optional)
Lists files under a specific folder.

**API Key Scope**: files / index

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskExecutionId** | **string**| Numeric ID of task execution. | 
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

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilesChildrenByTaskExecutionByTask**
> NodeumFileCollection FilesChildrenByTaskExecutionByTask(ctx, taskId, taskExecutionId, fileParentId, optional)
Lists files under a specific folder.

**API Key Scope**: files / index

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
  **taskExecutionId** | **string**| Numeric ID of task execution. | 
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

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IndexFiles**
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

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IndexFilesByContainer**
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

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IndexFilesByPool**
> NodeumFileCollection IndexFilesByPool(ctx, poolId, optional)
Lists files on root.

**API Key Scope**: files / index

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**| Numeric ID, or name of pool. | 
 **optional** | ***IndexFilesByPoolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IndexFilesByPoolOpts struct

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

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IndexFilesByTask**
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

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IndexFilesByTaskExecution**
> NodeumFileCollection IndexFilesByTaskExecution(ctx, taskExecutionId, optional)
Lists files on root.

**API Key Scope**: files / index

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskExecutionId** | **string**| Numeric ID of task execution. | 
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

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IndexFilesByTaskExecutionByTask**
> NodeumFileCollection IndexFilesByTaskExecutionByTask(ctx, taskId, taskExecutionId, optional)
Lists files on root.

**API Key Scope**: files / index

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
  **taskExecutionId** | **string**| Numeric ID of task execution. | 
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

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowFile**
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

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowFileByContainer**
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

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowFileByPool**
> NodeumFileWithPath ShowFileByPool(ctx, poolId, fileId)
Displays a specific file.

**API Key Scope**: files / show

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**| Numeric ID, or name of pool. | 
  **fileId** | **int32**| Numeric ID of file. | 

### Return type

[**NodeumFileWithPath**](nodeum_file_with_path.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowFileByTask**
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

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowFileByTaskExecution**
> NodeumFileWithPath ShowFileByTaskExecution(ctx, taskExecutionId, fileId)
Displays a specific file.

**API Key Scope**: files / show

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskExecutionId** | **string**| Numeric ID of task execution. | 
  **fileId** | **int32**| Numeric ID of file. | 

### Return type

[**NodeumFileWithPath**](nodeum_file_with_path.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowFileByTaskExecutionByTask**
> NodeumFileWithPath ShowFileByTaskExecutionByTask(ctx, taskId, taskExecutionId, fileId)
Displays a specific file.

**API Key Scope**: files / show

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
  **taskExecutionId** | **string**| Numeric ID of task execution. | 
  **fileId** | **int32**| Numeric ID of file. | 

### Return type

[**NodeumFileWithPath**](nodeum_file_with_path.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

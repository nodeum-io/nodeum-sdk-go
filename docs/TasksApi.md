# \TasksApi

All URIs are relative to *https://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTask**](TasksApi.md#CreateTask) | **Post** /tasks | Creates a new task.
[**DestroyTask**](TasksApi.md#DestroyTask) | **Delete** /tasks/{task_id} | Destroys a specific task.
[**IndexTasks**](TasksApi.md#IndexTasks) | **Get** /tasks | Lists all tasks.
[**ShowTask**](TasksApi.md#ShowTask) | **Get** /tasks/{task_id} | Displays a specific task.
[**UpdateTask**](TasksApi.md#UpdateTask) | **Put** /tasks/{task_id} | Updates a specific task.


# **CreateTask**
> Task CreateTask(ctx, taskBody)
Creates a new task.

**API Key Scope**: tasks / create

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskBody** | [**Task**](Task.md)|  | 

### Return type

[**Task**](task.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DestroyTask**
> DestroyTask(ctx, taskId)
Destroys a specific task.

**API Key Scope**: tasks / destroy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IndexTasks**
> TaskCollection IndexTasks(ctx, optional)
Lists all tasks.

**API Key Scope**: tasks / index

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IndexTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IndexTasksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **name** | **optional.String**| Filter on name | 
 **comment** | **optional.String**| Filter on comment | 
 **type_** | **optional.String**| Filter on type | 
 **priority** | **optional.String**| Filter on priority | 
 **conflictResolution** | **optional.String**| Filter on conflict resolution | 
 **action** | **optional.String**| Filter on action | 
 **activate** | **optional.String**| Filter on activate | 
 **creationDate** | **optional.String**| Filter on creation date | 
 **creationUsername** | **optional.String**| Filter on creation username | 
 **modificationDate** | **optional.String**| Filter on modification date | 
 **modificationUsername** | **optional.String**| Filter on modification username | 
 **jobStarted** | **optional.String**| Filter on job started | 
 **jobFinished** | **optional.String**| Filter on job finished | 
 **status** | **optional.String**| Filter on status | 
 **processedSize** | **optional.String**| Filter on processed size | 
 **toProcessSize** | **optional.String**| Filter on to process size | 

### Return type

[**TaskCollection**](task_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowTask**
> Task ShowTask(ctx, taskId)
Displays a specific task.

**API Key Scope**: tasks / show

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 

### Return type

[**Task**](task.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTask**
> Task UpdateTask(ctx, taskId, taskBody)
Updates a specific task.

**API Key Scope**: tasks / update

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**| Numeric ID or name of task. Task names are not unique, it&#39;s recommanded to use numeric ID. | 
  **taskBody** | [**Task**](Task.md)|  | 

### Return type

[**Task**](task.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


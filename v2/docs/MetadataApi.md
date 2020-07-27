# \MetadataApi

All URIs are relative to *https://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IndexFileMetadataDefinitions**](MetadataApi.md#IndexFileMetadataDefinitions) | **Get** /file_metadata_definitions | List file metadata definitions
[**IndexTaskMetadataDefinitions**](MetadataApi.md#IndexTaskMetadataDefinitions) | **Get** /task_metadata_definitions | List task metadata definitions
[**ShowFileMetadataDefinition**](MetadataApi.md#ShowFileMetadataDefinition) | **Get** /file_metadata_definitions/{metadata_definition_id} | Displays a specific task metadata definition.
[**ShowTaskMetadataDefinition**](MetadataApi.md#ShowTaskMetadataDefinition) | **Get** /task_metadata_definitions/{metadata_definition_id} | Displays a specific task metadata definition.


# **IndexFileMetadataDefinitions**
> FileMetadataDefinitionCollection IndexFileMetadataDefinitions(ctx, optional)
List file metadata definitions

**API Key Scope**: file_metadata_definitions / index

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IndexFileMetadataDefinitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IndexFileMetadataDefinitionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 

### Return type

[**FileMetadataDefinitionCollection**](file_metadata_definition_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IndexTaskMetadataDefinitions**
> TaskMetadataDefinitionCollection IndexTaskMetadataDefinitions(ctx, optional)
List task metadata definitions

**API Key Scope**: task_metadata_definitions / index

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IndexTaskMetadataDefinitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IndexTaskMetadataDefinitionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 

### Return type

[**TaskMetadataDefinitionCollection**](task_metadata_definition_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowFileMetadataDefinition**
> FileMetadataDefinition ShowFileMetadataDefinition(ctx, metadataDefinitionId)
Displays a specific task metadata definition.

**API Key Scope**: file_metadata_definitions / show

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **metadataDefinitionId** | **string**| Numeric ID or key of a metadata definition | 

### Return type

[**FileMetadataDefinition**](file_metadata_definition.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowTaskMetadataDefinition**
> TaskMetadataDefinition ShowTaskMetadataDefinition(ctx, metadataDefinitionId)
Displays a specific task metadata definition.

**API Key Scope**: task_metadata_definitions / show

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **metadataDefinitionId** | **string**| Numeric ID or key of a metadata definition | 

### Return type

[**TaskMetadataDefinition**](task_metadata_definition.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


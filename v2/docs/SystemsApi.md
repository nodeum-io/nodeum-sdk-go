# \SystemsApi

All URIs are relative to *https://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResultDownloadTraces**](SystemsApi.md#ResultDownloadTraces) | **Get** /systems/download_traces | Check result of a download traces job.
[**TriggerDownloadTraces**](SystemsApi.md#TriggerDownloadTraces) | **Put** /systems/download_traces | Trigger a download traces request.


# **ResultDownloadTraces**
> DownloadTracesActiveJobStatus ResultDownloadTraces(ctx, optional)
Check result of a download traces job.

**API Key Scope**: systems / download_traces

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ResultDownloadTracesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ResultDownloadTracesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobId** | **optional.String**| ID of active job | 

### Return type

[**DownloadTracesActiveJobStatus**](download_traces_active_job_status.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TriggerDownloadTraces**
> ActiveJobStatus TriggerDownloadTraces(ctx, optional)
Trigger a download traces request.

**API Key Scope**: systems / download_traces

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TriggerDownloadTracesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TriggerDownloadTracesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **optional.String**| Type of traces to download | 

### Return type

[**ActiveJobStatus**](active_job_status.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


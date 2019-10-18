# \MountPointsApi

All URIs are relative to *https://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMountPoint**](MountPointsApi.md#CreateMountPoint) | **Post** /mount_points | Creates a new mount point.
[**DestroyMountPoint**](MountPointsApi.md#DestroyMountPoint) | **Delete** /mount_points/{mount_point_id} | Destroys a specific mount point.
[**IndexMountPoints**](MountPointsApi.md#IndexMountPoints) | **Get** /mount_points | Lists all mount points.
[**MountMountPoint**](MountPointsApi.md#MountMountPoint) | **Put** /mount_points/{mount_point_id}/mount | Mount Mount Point.
[**MountStatusMountPoint**](MountPointsApi.md#MountStatusMountPoint) | **Get** /mount_points/{mount_point_id}/mount | Get mount status of Mount Point.
[**ShowMountPoint**](MountPointsApi.md#ShowMountPoint) | **Get** /mount_points/{mount_point_id} | Displays a specific mount point&#x60;.
[**UnmountMountPoint**](MountPointsApi.md#UnmountMountPoint) | **Delete** /mount_points/{mount_point_id}/mount | Unmount Mount Point.
[**UpdateMountPoint**](MountPointsApi.md#UpdateMountPoint) | **Put** /mount_points/{mount_point_id} | Updates a specific mount point&#x60;.


# **CreateMountPoint**
> MountPoint CreateMountPoint(ctx, mountPointBody)
Creates a new mount point.

It **does not** create and mount the file structure. Use API v1 instead.  **API Key Scope**: mount_points / create

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mountPointBody** | [**MountPoint**](MountPoint.md)|  | 

### Return type

[**MountPoint**](mount_point.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DestroyMountPoint**
> DestroyMountPoint(ctx, mountPointId)
Destroys a specific mount point.

**API Key Scope**: mount_points / destroy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mountPointId** | **string**| Numeric ID or name of mount point. | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IndexMountPoints**
> MountPointCollection IndexMountPoints(ctx, optional)
Lists all mount points.

**API Key Scope**: mount_points / index   Optional API Key Explicit Scope: mount_points / get_password

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IndexMountPointsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IndexMountPointsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **name** | **optional.String**| Filter on name | 
 **target** | **optional.String**| Filter on mount point target | 
 **type_** | **optional.String**| Filter on type | 
 **options** | **optional.String**| Filter on options | 
 **username** | **optional.String**| Filter on username | 
 **comment** | **optional.String**| Filter on comment | 
 **scanInterval** | **optional.String**| Filter on mount point scan interval | 
 **price** | **optional.String**| Filter on price | 

### Return type

[**MountPointCollection**](mount_point_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MountMountPoint**
> MountStatus MountMountPoint(ctx, mountPointId)
Mount Mount Point.

**API Key Scope**: mount_points / mount

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mountPointId** | **string**| Numeric ID or name of mount point. | 

### Return type

[**MountStatus**](mount_status.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MountStatusMountPoint**
> MountStatus MountStatusMountPoint(ctx, mountPointId)
Get mount status of Mount Point.

**API Key Scope**: mount_points / mount_status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mountPointId** | **string**| Numeric ID or name of mount point. | 

### Return type

[**MountStatus**](mount_status.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowMountPoint**
> MountPoint ShowMountPoint(ctx, mountPointId)
Displays a specific mount point`.

**API Key Scope**: mount_points / show   Optional API Key Explicit Scope: mount_points / get_password

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mountPointId** | **string**| Numeric ID or name of mount point. | 

### Return type

[**MountPoint**](mount_point.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnmountMountPoint**
> MountStatus UnmountMountPoint(ctx, mountPointId)
Unmount Mount Point.

**API Key Scope**: mount_points / unmount

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mountPointId** | **string**| Numeric ID or name of mount point. | 

### Return type

[**MountStatus**](mount_status.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMountPoint**
> MountPoint UpdateMountPoint(ctx, mountPointId, mountPointBody)
Updates a specific mount point`.

It **does not** create and mount the file structure. Use API v1 instead.  **API Key Scope**: mount_points / update

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mountPointId** | **string**| Numeric ID or name of mount point. | 
  **mountPointBody** | [**MountPoint**](MountPoint.md)|  | 

### Return type

[**MountPoint**](mount_point.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


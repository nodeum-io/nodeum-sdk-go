# \TapesApi

All URIs are relative to *https://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IndexTapes**](TapesApi.md#IndexTapes) | **Get** /tapes | Lists all tapes.
[**IndexTapesByTapeLibrary**](TapesApi.md#IndexTapesByTapeLibrary) | **Get** /tape_libraries/{tape_library_id}/tapes | Lists all tapes.
[**IndexTapesByTapePool**](TapesApi.md#IndexTapesByTapePool) | **Get** /tape_pools/{tape_pool_id}/tapes | Lists all tapes.
[**MountStatusTape**](TapesApi.md#MountStatusTape) | **Get** /tapes/{tape_id}/mount | Get mount status of Tape.
[**MountStatusTapeByTapeLibrary**](TapesApi.md#MountStatusTapeByTapeLibrary) | **Get** /tape_libraries/{tape_library_id}/tapes/{tape_id}/mount | Get mount status of Tape.
[**MountStatusTapeByTapePool**](TapesApi.md#MountStatusTapeByTapePool) | **Get** /tape_pools/{tape_pool_id}/tapes/{tape_id}/mount | Get mount status of Tape.
[**ShowTape**](TapesApi.md#ShowTape) | **Get** /tapes/{tape_id} | Displays a specific tape.
[**ShowTapeByTapeLibrary**](TapesApi.md#ShowTapeByTapeLibrary) | **Get** /tape_libraries/{tape_library_id}/tapes/{tape_id} | Displays a specific tape.
[**ShowTapeByTapePool**](TapesApi.md#ShowTapeByTapePool) | **Get** /tape_pools/{tape_pool_id}/tapes/{tape_id} | Displays a specific tape.


# **IndexTapes**
> TapeCollection IndexTapes(ctx, optional)
Lists all tapes.

**API Key Scope**: tapes / index

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IndexTapesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IndexTapesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **tapeLibraryId** | **optional.String**| Filter on tape library id | 
 **tapePoolId** | **optional.String**| Filter on tape pool id | 
 **barcode** | **optional.String**| Filter on barcode | 
 **location** | **optional.String**| Filter on location | 
 **type_** | **optional.String**| Filter on type | 
 **locked** | **optional.String**| Filter on locked | 
 **scratch** | **optional.String**| Filter on scratch | 
 **cleaning** | **optional.String**| Filter on cleaning | 
 **writeProtect** | **optional.String**| Filter on write protect | 
 **mounted** | **optional.String**| Filter on mounted | 
 **ejected** | **optional.String**| Filter on ejected | 
 **known** | **optional.String**| Filter on known | 
 **mountCount** | **optional.String**| Filter on mount count | 
 **dateIn** | **optional.String**| Filter on date in | 
 **dateMove** | **optional.String**| Filter on date move | 
 **free** | **optional.String**| Filter on free | 
 **max** | **optional.String**| Filter on max | 
 **lastSizeUpdate** | **optional.String**| Filter on last size update | 
 **lastMaintenance** | **optional.String**| Filter on last maintenance | 
 **lastRepack** | **optional.String**| Filter on last repack | 
 **repackStatus** | **optional.String**| Filter on repack status | 
 **hash** | **optional.String**| Filter on hash | 
 **forceImportType** | **optional.String**| Filter on force import type | 
 **needToCheck** | **optional.String**| Filter on need to check | 

### Return type

[**TapeCollection**](tape_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IndexTapesByTapeLibrary**
> TapeCollection IndexTapesByTapeLibrary(ctx, tapeLibraryId, optional)
Lists all tapes.

**API Key Scope**: tapes / index

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tapeLibraryId** | **string**| Numeric ID, serial, or name of tape library. | 
 **optional** | ***IndexTapesByTapeLibraryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IndexTapesByTapeLibraryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **tapePoolId** | **optional.String**| Filter on tape pool id | 
 **barcode** | **optional.String**| Filter on barcode | 
 **location** | **optional.String**| Filter on location | 
 **type_** | **optional.String**| Filter on type | 
 **locked** | **optional.String**| Filter on locked | 
 **scratch** | **optional.String**| Filter on scratch | 
 **cleaning** | **optional.String**| Filter on cleaning | 
 **writeProtect** | **optional.String**| Filter on write protect | 
 **mounted** | **optional.String**| Filter on mounted | 
 **ejected** | **optional.String**| Filter on ejected | 
 **known** | **optional.String**| Filter on known | 
 **mountCount** | **optional.String**| Filter on mount count | 
 **dateIn** | **optional.String**| Filter on date in | 
 **dateMove** | **optional.String**| Filter on date move | 
 **free** | **optional.String**| Filter on free | 
 **max** | **optional.String**| Filter on max | 
 **lastSizeUpdate** | **optional.String**| Filter on last size update | 
 **lastMaintenance** | **optional.String**| Filter on last maintenance | 
 **lastRepack** | **optional.String**| Filter on last repack | 
 **repackStatus** | **optional.String**| Filter on repack status | 
 **hash** | **optional.String**| Filter on hash | 
 **forceImportType** | **optional.String**| Filter on force import type | 
 **needToCheck** | **optional.String**| Filter on need to check | 

### Return type

[**TapeCollection**](tape_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IndexTapesByTapePool**
> TapeCollection IndexTapesByTapePool(ctx, tapePoolId, optional)
Lists all tapes.

**API Key Scope**: tapes / index

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tapePoolId** | **string**| Numeric ID, or name of tape pool. | 
 **optional** | ***IndexTapesByTapePoolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IndexTapesByTapePoolOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of items to display for pagination. | 
 **offset** | **optional.Int32**| The number of items to skip for pagination. | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sort results by attribute.  Can sort on multiple attributes, separated by &#x60;|&#x60;. Order direction can be suffixing the attribute by either &#x60;:asc&#x60; (default) or &#x60;:desc&#x60;. | 
 **id** | **optional.String**| Filter on id | 
 **tapeLibraryId** | **optional.String**| Filter on tape library id | 
 **barcode** | **optional.String**| Filter on barcode | 
 **location** | **optional.String**| Filter on location | 
 **type_** | **optional.String**| Filter on type | 
 **locked** | **optional.String**| Filter on locked | 
 **scratch** | **optional.String**| Filter on scratch | 
 **cleaning** | **optional.String**| Filter on cleaning | 
 **writeProtect** | **optional.String**| Filter on write protect | 
 **mounted** | **optional.String**| Filter on mounted | 
 **ejected** | **optional.String**| Filter on ejected | 
 **known** | **optional.String**| Filter on known | 
 **mountCount** | **optional.String**| Filter on mount count | 
 **dateIn** | **optional.String**| Filter on date in | 
 **dateMove** | **optional.String**| Filter on date move | 
 **free** | **optional.String**| Filter on free | 
 **max** | **optional.String**| Filter on max | 
 **lastSizeUpdate** | **optional.String**| Filter on last size update | 
 **lastMaintenance** | **optional.String**| Filter on last maintenance | 
 **lastRepack** | **optional.String**| Filter on last repack | 
 **repackStatus** | **optional.String**| Filter on repack status | 
 **hash** | **optional.String**| Filter on hash | 
 **forceImportType** | **optional.String**| Filter on force import type | 
 **needToCheck** | **optional.String**| Filter on need to check | 

### Return type

[**TapeCollection**](tape_collection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MountStatusTape**
> MountStatus MountStatusTape(ctx, tapeId)
Get mount status of Tape.

**API Key Scope**: tapes / mount_status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tapeId** | **string**| Numeric ID, or barcode of tape. | 

### Return type

[**MountStatus**](mount_status.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MountStatusTapeByTapeLibrary**
> MountStatus MountStatusTapeByTapeLibrary(ctx, tapeLibraryId, tapeId)
Get mount status of Tape.

**API Key Scope**: tapes / mount_status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tapeLibraryId** | **string**| Numeric ID, serial, or name of tape library. | 
  **tapeId** | **string**| Numeric ID, or barcode of tape. | 

### Return type

[**MountStatus**](mount_status.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MountStatusTapeByTapePool**
> MountStatus MountStatusTapeByTapePool(ctx, tapePoolId, tapeId)
Get mount status of Tape.

**API Key Scope**: tapes / mount_status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tapePoolId** | **string**| Numeric ID, or name of tape pool. | 
  **tapeId** | **string**| Numeric ID, or barcode of tape. | 

### Return type

[**MountStatus**](mount_status.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowTape**
> Tape ShowTape(ctx, tapeId)
Displays a specific tape.

**API Key Scope**: tapes / show

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tapeId** | **string**| Numeric ID, or barcode of tape. | 

### Return type

[**Tape**](tape.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowTapeByTapeLibrary**
> Tape ShowTapeByTapeLibrary(ctx, tapeLibraryId, tapeId)
Displays a specific tape.

**API Key Scope**: tapes / show

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tapeLibraryId** | **string**| Numeric ID, serial, or name of tape library. | 
  **tapeId** | **string**| Numeric ID, or barcode of tape. | 

### Return type

[**Tape**](tape.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowTapeByTapePool**
> Tape ShowTapeByTapePool(ctx, tapePoolId, tapeId)
Displays a specific tape.

**API Key Scope**: tapes / show

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tapePoolId** | **string**| Numeric ID, or name of tape pool. | 
  **tapeId** | **string**| Numeric ID, or barcode of tape. | 

### Return type

[**Tape**](tape.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


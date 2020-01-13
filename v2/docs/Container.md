# Container

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**QuotaTotalSize** | Pointer to **int32** |  | [optional] 
**QuotaOnCache** | Pointer to **int32** |  | [optional] 
**StatTotalFiles** | Pointer to **int32** |  | [optional] [readonly] 
**StatTotalSize** | Pointer to **int32** |  | [optional] [readonly] 
**StatSizeOnCache** | Pointer to **int32** |  | [optional] [readonly] 
**GuestRight** | Pointer to **string** |  | [optional] 
**LastUpdate** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### GetId

`func (o *Container) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Container) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Container) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Container) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetName

`func (o *Container) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Container) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Container) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Container) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetComment

`func (o *Container) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Container) GetCommentOk() (string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComment

`func (o *Container) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetComment

`func (o *Container) SetComment(v string)`

SetComment gets a reference to the given string and assigns it to the Comment field.

### GetQuotaTotalSize

`func (o *Container) GetQuotaTotalSize() int32`

GetQuotaTotalSize returns the QuotaTotalSize field if non-nil, zero value otherwise.

### GetQuotaTotalSizeOk

`func (o *Container) GetQuotaTotalSizeOk() (int32, bool)`

GetQuotaTotalSizeOk returns a tuple with the QuotaTotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQuotaTotalSize

`func (o *Container) HasQuotaTotalSize() bool`

HasQuotaTotalSize returns a boolean if a field has been set.

### SetQuotaTotalSize

`func (o *Container) SetQuotaTotalSize(v int32)`

SetQuotaTotalSize gets a reference to the given int32 and assigns it to the QuotaTotalSize field.

### GetQuotaOnCache

`func (o *Container) GetQuotaOnCache() int32`

GetQuotaOnCache returns the QuotaOnCache field if non-nil, zero value otherwise.

### GetQuotaOnCacheOk

`func (o *Container) GetQuotaOnCacheOk() (int32, bool)`

GetQuotaOnCacheOk returns a tuple with the QuotaOnCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQuotaOnCache

`func (o *Container) HasQuotaOnCache() bool`

HasQuotaOnCache returns a boolean if a field has been set.

### SetQuotaOnCache

`func (o *Container) SetQuotaOnCache(v int32)`

SetQuotaOnCache gets a reference to the given int32 and assigns it to the QuotaOnCache field.

### GetStatTotalFiles

`func (o *Container) GetStatTotalFiles() int32`

GetStatTotalFiles returns the StatTotalFiles field if non-nil, zero value otherwise.

### GetStatTotalFilesOk

`func (o *Container) GetStatTotalFilesOk() (int32, bool)`

GetStatTotalFilesOk returns a tuple with the StatTotalFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatTotalFiles

`func (o *Container) HasStatTotalFiles() bool`

HasStatTotalFiles returns a boolean if a field has been set.

### SetStatTotalFiles

`func (o *Container) SetStatTotalFiles(v int32)`

SetStatTotalFiles gets a reference to the given int32 and assigns it to the StatTotalFiles field.

### GetStatTotalSize

`func (o *Container) GetStatTotalSize() int32`

GetStatTotalSize returns the StatTotalSize field if non-nil, zero value otherwise.

### GetStatTotalSizeOk

`func (o *Container) GetStatTotalSizeOk() (int32, bool)`

GetStatTotalSizeOk returns a tuple with the StatTotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatTotalSize

`func (o *Container) HasStatTotalSize() bool`

HasStatTotalSize returns a boolean if a field has been set.

### SetStatTotalSize

`func (o *Container) SetStatTotalSize(v int32)`

SetStatTotalSize gets a reference to the given int32 and assigns it to the StatTotalSize field.

### GetStatSizeOnCache

`func (o *Container) GetStatSizeOnCache() int32`

GetStatSizeOnCache returns the StatSizeOnCache field if non-nil, zero value otherwise.

### GetStatSizeOnCacheOk

`func (o *Container) GetStatSizeOnCacheOk() (int32, bool)`

GetStatSizeOnCacheOk returns a tuple with the StatSizeOnCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatSizeOnCache

`func (o *Container) HasStatSizeOnCache() bool`

HasStatSizeOnCache returns a boolean if a field has been set.

### SetStatSizeOnCache

`func (o *Container) SetStatSizeOnCache(v int32)`

SetStatSizeOnCache gets a reference to the given int32 and assigns it to the StatSizeOnCache field.

### GetGuestRight

`func (o *Container) GetGuestRight() string`

GetGuestRight returns the GuestRight field if non-nil, zero value otherwise.

### GetGuestRightOk

`func (o *Container) GetGuestRightOk() (string, bool)`

GetGuestRightOk returns a tuple with the GuestRight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGuestRight

`func (o *Container) HasGuestRight() bool`

HasGuestRight returns a boolean if a field has been set.

### SetGuestRight

`func (o *Container) SetGuestRight(v string)`

SetGuestRight gets a reference to the given string and assigns it to the GuestRight field.

### GetLastUpdate

`func (o *Container) GetLastUpdate() string`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *Container) GetLastUpdateOk() (string, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastUpdate

`func (o *Container) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### SetLastUpdate

`func (o *Container) SetLastUpdate(v string)`

SetLastUpdate gets a reference to the given string and assigns it to the LastUpdate field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



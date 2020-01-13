# MountStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mounted** | Pointer to **bool** |  | [optional] [readonly] 
**Message** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to [**MountInfo**](mount_info.md) |  | [optional] 

## Methods

### GetMounted

`func (o *MountStatus) GetMounted() bool`

GetMounted returns the Mounted field if non-nil, zero value otherwise.

### GetMountedOk

`func (o *MountStatus) GetMountedOk() (bool, bool)`

GetMountedOk returns a tuple with the Mounted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMounted

`func (o *MountStatus) HasMounted() bool`

HasMounted returns a boolean if a field has been set.

### SetMounted

`func (o *MountStatus) SetMounted(v bool)`

SetMounted gets a reference to the given bool and assigns it to the Mounted field.

### GetMessage

`func (o *MountStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MountStatus) GetMessageOk() (string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessage

`func (o *MountStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessage

`func (o *MountStatus) SetMessage(v string)`

SetMessage gets a reference to the given string and assigns it to the Message field.

### GetStatus

`func (o *MountStatus) GetStatus() MountInfo`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MountStatus) GetStatusOk() (MountInfo, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *MountStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *MountStatus) SetStatus(v MountInfo)`

SetStatus gets a reference to the given MountInfo and assigns it to the Status field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



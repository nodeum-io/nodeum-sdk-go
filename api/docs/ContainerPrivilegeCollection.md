# ContainerPrivilegeCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] [readonly] 
**ContainerPrivileges** | Pointer to [**[]ContainerPrivilege**](container_privilege.md) |  | [optional] [readonly] 

## Methods

### GetCount

`func (o *ContainerPrivilegeCollection) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ContainerPrivilegeCollection) GetCountOk() (int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCount

`func (o *ContainerPrivilegeCollection) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCount

`func (o *ContainerPrivilegeCollection) SetCount(v int32)`

SetCount gets a reference to the given int32 and assigns it to the Count field.

### GetContainerPrivileges

`func (o *ContainerPrivilegeCollection) GetContainerPrivileges() []ContainerPrivilege`

GetContainerPrivileges returns the ContainerPrivileges field if non-nil, zero value otherwise.

### GetContainerPrivilegesOk

`func (o *ContainerPrivilegeCollection) GetContainerPrivilegesOk() ([]ContainerPrivilege, bool)`

GetContainerPrivilegesOk returns a tuple with the ContainerPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContainerPrivileges

`func (o *ContainerPrivilegeCollection) HasContainerPrivileges() bool`

HasContainerPrivileges returns a boolean if a field has been set.

### SetContainerPrivileges

`func (o *ContainerPrivilegeCollection) SetContainerPrivileges(v []ContainerPrivilege)`

SetContainerPrivileges gets a reference to the given []ContainerPrivilege and assigns it to the ContainerPrivileges field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



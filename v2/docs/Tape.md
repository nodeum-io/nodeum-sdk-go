# Tape

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**TapeLibraryId** | Pointer to **int32** |  | [optional] [readonly] 
**PoolId** | Pointer to **int32** |  | [optional] [readonly] 
**Barcode** | Pointer to **string** |  | [optional] [readonly] 
**Location** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] [readonly] 
**Locked** | Pointer to **bool** |  | [optional] [readonly] 
**Scratch** | Pointer to **bool** |  | [optional] [readonly] 
**Cleaning** | Pointer to **bool** |  | [optional] [readonly] 
**WriteProtect** | Pointer to **bool** |  | [optional] [readonly] 
**Mounted** | Pointer to **bool** |  | [optional] [readonly] 
**Ejected** | Pointer to **bool** |  | [optional] [readonly] 
**Known** | Pointer to **bool** |  | [optional] [readonly] 
**MountCount** | Pointer to **int32** |  | [optional] [readonly] 
**DateIn** | Pointer to **string** |  | [optional] [readonly] 
**DateMove** | Pointer to **string** |  | [optional] [readonly] 
**Free** | Pointer to **int32** |  | [optional] [readonly] 
**Max** | Pointer to **int32** |  | [optional] [readonly] 
**LastSizeUpdate** | Pointer to **string** |  | [optional] [readonly] 
**LastMaintenance** | Pointer to **string** |  | [optional] [readonly] 
**LastRepack** | Pointer to **string** |  | [optional] [readonly] 
**RepackStatus** | Pointer to **bool** |  | [optional] [readonly] 
**Hash** | Pointer to **string** |  | [optional] [readonly] 
**ForceImportType** | Pointer to **bool** |  | [optional] [readonly] 
**NeedToCheck** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### GetId

`func (o *Tape) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tape) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Tape) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Tape) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetTapeLibraryId

`func (o *Tape) GetTapeLibraryId() int32`

GetTapeLibraryId returns the TapeLibraryId field if non-nil, zero value otherwise.

### GetTapeLibraryIdOk

`func (o *Tape) GetTapeLibraryIdOk() (int32, bool)`

GetTapeLibraryIdOk returns a tuple with the TapeLibraryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTapeLibraryId

`func (o *Tape) HasTapeLibraryId() bool`

HasTapeLibraryId returns a boolean if a field has been set.

### SetTapeLibraryId

`func (o *Tape) SetTapeLibraryId(v int32)`

SetTapeLibraryId gets a reference to the given int32 and assigns it to the TapeLibraryId field.

### GetPoolId

`func (o *Tape) GetPoolId() int32`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *Tape) GetPoolIdOk() (int32, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPoolId

`func (o *Tape) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### SetPoolId

`func (o *Tape) SetPoolId(v int32)`

SetPoolId gets a reference to the given int32 and assigns it to the PoolId field.

### GetBarcode

`func (o *Tape) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *Tape) GetBarcodeOk() (string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBarcode

`func (o *Tape) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### SetBarcode

`func (o *Tape) SetBarcode(v string)`

SetBarcode gets a reference to the given string and assigns it to the Barcode field.

### GetLocation

`func (o *Tape) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Tape) GetLocationOk() (string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocation

`func (o *Tape) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocation

`func (o *Tape) SetLocation(v string)`

SetLocation gets a reference to the given string and assigns it to the Location field.

### GetType

`func (o *Tape) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Tape) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *Tape) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *Tape) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetLocked

`func (o *Tape) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *Tape) GetLockedOk() (bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocked

`func (o *Tape) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### SetLocked

`func (o *Tape) SetLocked(v bool)`

SetLocked gets a reference to the given bool and assigns it to the Locked field.

### GetScratch

`func (o *Tape) GetScratch() bool`

GetScratch returns the Scratch field if non-nil, zero value otherwise.

### GetScratchOk

`func (o *Tape) GetScratchOk() (bool, bool)`

GetScratchOk returns a tuple with the Scratch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScratch

`func (o *Tape) HasScratch() bool`

HasScratch returns a boolean if a field has been set.

### SetScratch

`func (o *Tape) SetScratch(v bool)`

SetScratch gets a reference to the given bool and assigns it to the Scratch field.

### GetCleaning

`func (o *Tape) GetCleaning() bool`

GetCleaning returns the Cleaning field if non-nil, zero value otherwise.

### GetCleaningOk

`func (o *Tape) GetCleaningOk() (bool, bool)`

GetCleaningOk returns a tuple with the Cleaning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCleaning

`func (o *Tape) HasCleaning() bool`

HasCleaning returns a boolean if a field has been set.

### SetCleaning

`func (o *Tape) SetCleaning(v bool)`

SetCleaning gets a reference to the given bool and assigns it to the Cleaning field.

### GetWriteProtect

`func (o *Tape) GetWriteProtect() bool`

GetWriteProtect returns the WriteProtect field if non-nil, zero value otherwise.

### GetWriteProtectOk

`func (o *Tape) GetWriteProtectOk() (bool, bool)`

GetWriteProtectOk returns a tuple with the WriteProtect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWriteProtect

`func (o *Tape) HasWriteProtect() bool`

HasWriteProtect returns a boolean if a field has been set.

### SetWriteProtect

`func (o *Tape) SetWriteProtect(v bool)`

SetWriteProtect gets a reference to the given bool and assigns it to the WriteProtect field.

### GetMounted

`func (o *Tape) GetMounted() bool`

GetMounted returns the Mounted field if non-nil, zero value otherwise.

### GetMountedOk

`func (o *Tape) GetMountedOk() (bool, bool)`

GetMountedOk returns a tuple with the Mounted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMounted

`func (o *Tape) HasMounted() bool`

HasMounted returns a boolean if a field has been set.

### SetMounted

`func (o *Tape) SetMounted(v bool)`

SetMounted gets a reference to the given bool and assigns it to the Mounted field.

### GetEjected

`func (o *Tape) GetEjected() bool`

GetEjected returns the Ejected field if non-nil, zero value otherwise.

### GetEjectedOk

`func (o *Tape) GetEjectedOk() (bool, bool)`

GetEjectedOk returns a tuple with the Ejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEjected

`func (o *Tape) HasEjected() bool`

HasEjected returns a boolean if a field has been set.

### SetEjected

`func (o *Tape) SetEjected(v bool)`

SetEjected gets a reference to the given bool and assigns it to the Ejected field.

### GetKnown

`func (o *Tape) GetKnown() bool`

GetKnown returns the Known field if non-nil, zero value otherwise.

### GetKnownOk

`func (o *Tape) GetKnownOk() (bool, bool)`

GetKnownOk returns a tuple with the Known field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKnown

`func (o *Tape) HasKnown() bool`

HasKnown returns a boolean if a field has been set.

### SetKnown

`func (o *Tape) SetKnown(v bool)`

SetKnown gets a reference to the given bool and assigns it to the Known field.

### GetMountCount

`func (o *Tape) GetMountCount() int32`

GetMountCount returns the MountCount field if non-nil, zero value otherwise.

### GetMountCountOk

`func (o *Tape) GetMountCountOk() (int32, bool)`

GetMountCountOk returns a tuple with the MountCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMountCount

`func (o *Tape) HasMountCount() bool`

HasMountCount returns a boolean if a field has been set.

### SetMountCount

`func (o *Tape) SetMountCount(v int32)`

SetMountCount gets a reference to the given int32 and assigns it to the MountCount field.

### GetDateIn

`func (o *Tape) GetDateIn() string`

GetDateIn returns the DateIn field if non-nil, zero value otherwise.

### GetDateInOk

`func (o *Tape) GetDateInOk() (string, bool)`

GetDateInOk returns a tuple with the DateIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDateIn

`func (o *Tape) HasDateIn() bool`

HasDateIn returns a boolean if a field has been set.

### SetDateIn

`func (o *Tape) SetDateIn(v string)`

SetDateIn gets a reference to the given string and assigns it to the DateIn field.

### GetDateMove

`func (o *Tape) GetDateMove() string`

GetDateMove returns the DateMove field if non-nil, zero value otherwise.

### GetDateMoveOk

`func (o *Tape) GetDateMoveOk() (string, bool)`

GetDateMoveOk returns a tuple with the DateMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDateMove

`func (o *Tape) HasDateMove() bool`

HasDateMove returns a boolean if a field has been set.

### SetDateMove

`func (o *Tape) SetDateMove(v string)`

SetDateMove gets a reference to the given string and assigns it to the DateMove field.

### GetFree

`func (o *Tape) GetFree() int32`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *Tape) GetFreeOk() (int32, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFree

`func (o *Tape) HasFree() bool`

HasFree returns a boolean if a field has been set.

### SetFree

`func (o *Tape) SetFree(v int32)`

SetFree gets a reference to the given int32 and assigns it to the Free field.

### GetMax

`func (o *Tape) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *Tape) GetMaxOk() (int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMax

`func (o *Tape) HasMax() bool`

HasMax returns a boolean if a field has been set.

### SetMax

`func (o *Tape) SetMax(v int32)`

SetMax gets a reference to the given int32 and assigns it to the Max field.

### GetLastSizeUpdate

`func (o *Tape) GetLastSizeUpdate() string`

GetLastSizeUpdate returns the LastSizeUpdate field if non-nil, zero value otherwise.

### GetLastSizeUpdateOk

`func (o *Tape) GetLastSizeUpdateOk() (string, bool)`

GetLastSizeUpdateOk returns a tuple with the LastSizeUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastSizeUpdate

`func (o *Tape) HasLastSizeUpdate() bool`

HasLastSizeUpdate returns a boolean if a field has been set.

### SetLastSizeUpdate

`func (o *Tape) SetLastSizeUpdate(v string)`

SetLastSizeUpdate gets a reference to the given string and assigns it to the LastSizeUpdate field.

### GetLastMaintenance

`func (o *Tape) GetLastMaintenance() string`

GetLastMaintenance returns the LastMaintenance field if non-nil, zero value otherwise.

### GetLastMaintenanceOk

`func (o *Tape) GetLastMaintenanceOk() (string, bool)`

GetLastMaintenanceOk returns a tuple with the LastMaintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastMaintenance

`func (o *Tape) HasLastMaintenance() bool`

HasLastMaintenance returns a boolean if a field has been set.

### SetLastMaintenance

`func (o *Tape) SetLastMaintenance(v string)`

SetLastMaintenance gets a reference to the given string and assigns it to the LastMaintenance field.

### GetLastRepack

`func (o *Tape) GetLastRepack() string`

GetLastRepack returns the LastRepack field if non-nil, zero value otherwise.

### GetLastRepackOk

`func (o *Tape) GetLastRepackOk() (string, bool)`

GetLastRepackOk returns a tuple with the LastRepack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastRepack

`func (o *Tape) HasLastRepack() bool`

HasLastRepack returns a boolean if a field has been set.

### SetLastRepack

`func (o *Tape) SetLastRepack(v string)`

SetLastRepack gets a reference to the given string and assigns it to the LastRepack field.

### GetRepackStatus

`func (o *Tape) GetRepackStatus() bool`

GetRepackStatus returns the RepackStatus field if non-nil, zero value otherwise.

### GetRepackStatusOk

`func (o *Tape) GetRepackStatusOk() (bool, bool)`

GetRepackStatusOk returns a tuple with the RepackStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRepackStatus

`func (o *Tape) HasRepackStatus() bool`

HasRepackStatus returns a boolean if a field has been set.

### SetRepackStatus

`func (o *Tape) SetRepackStatus(v bool)`

SetRepackStatus gets a reference to the given bool and assigns it to the RepackStatus field.

### GetHash

`func (o *Tape) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Tape) GetHashOk() (string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHash

`func (o *Tape) HasHash() bool`

HasHash returns a boolean if a field has been set.

### SetHash

`func (o *Tape) SetHash(v string)`

SetHash gets a reference to the given string and assigns it to the Hash field.

### GetForceImportType

`func (o *Tape) GetForceImportType() bool`

GetForceImportType returns the ForceImportType field if non-nil, zero value otherwise.

### GetForceImportTypeOk

`func (o *Tape) GetForceImportTypeOk() (bool, bool)`

GetForceImportTypeOk returns a tuple with the ForceImportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasForceImportType

`func (o *Tape) HasForceImportType() bool`

HasForceImportType returns a boolean if a field has been set.

### SetForceImportType

`func (o *Tape) SetForceImportType(v bool)`

SetForceImportType gets a reference to the given bool and assigns it to the ForceImportType field.

### GetNeedToCheck

`func (o *Tape) GetNeedToCheck() bool`

GetNeedToCheck returns the NeedToCheck field if non-nil, zero value otherwise.

### GetNeedToCheckOk

`func (o *Tape) GetNeedToCheckOk() (bool, bool)`

GetNeedToCheckOk returns a tuple with the NeedToCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNeedToCheck

`func (o *Tape) HasNeedToCheck() bool`

HasNeedToCheck returns a boolean if a field has been set.

### SetNeedToCheck

`func (o *Tape) SetNeedToCheck(v bool)`

SetNeedToCheck gets a reference to the given bool and assigns it to the NeedToCheck field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



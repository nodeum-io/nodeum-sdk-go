# TapeDrive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** |  | [optional] [readonly] 
**ScsiAddress** | Pointer to **int32** |  | [optional] [readonly] 
**Vendor** | Pointer to **string** |  | [optional] [readonly] 
**Product** | Pointer to **string** |  | [optional] [readonly] 
**Firmware** | Pointer to **string** |  | [optional] [readonly] 
**Device** | Pointer to **string** | When saved, device may be prefixed by *n* (eg. &#x60;/dev/nst5&#x60;) | [optional] [readonly] 
**Sgdevice** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**TapeLibraryId** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Libso** | Pointer to **string** |  | [optional] 
**Acs** | Pointer to **int32** |  | [optional] 
**Lsm** | Pointer to **int32** |  | [optional] 
**Panel** | Pointer to **int32** |  | [optional] 
**Transport** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] [readonly] 
**Full** | Pointer to **int32** |  | [optional] [readonly] 
**MountCount** | Pointer to **int32** |  | [optional] [readonly] 
**UseTo** | Pointer to **string** |  | [optional] [readonly] 
**UseBy** | Pointer to **string** |  | [optional] [readonly] 
**Barcode** | Pointer to **string** |  | [optional] [readonly] 
**TaskId** | Pointer to **int32** |  | [optional] [readonly] 
**UseFileProcessedSize** | Pointer to **int32** |  | [optional] [readonly] 
**UseFileSizeToProcess** | Pointer to **int32** |  | [optional] [readonly] 
**UseFileNameProcessed** | Pointer to **string** |  | [optional] [readonly] 
**Bandwidth** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### GetSerial

`func (o *TapeDrive) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *TapeDrive) GetSerialOk() (string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSerial

`func (o *TapeDrive) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### SetSerial

`func (o *TapeDrive) SetSerial(v string)`

SetSerial gets a reference to the given string and assigns it to the Serial field.

### GetScsiAddress

`func (o *TapeDrive) GetScsiAddress() int32`

GetScsiAddress returns the ScsiAddress field if non-nil, zero value otherwise.

### GetScsiAddressOk

`func (o *TapeDrive) GetScsiAddressOk() (int32, bool)`

GetScsiAddressOk returns a tuple with the ScsiAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScsiAddress

`func (o *TapeDrive) HasScsiAddress() bool`

HasScsiAddress returns a boolean if a field has been set.

### SetScsiAddress

`func (o *TapeDrive) SetScsiAddress(v int32)`

SetScsiAddress gets a reference to the given int32 and assigns it to the ScsiAddress field.

### GetVendor

`func (o *TapeDrive) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *TapeDrive) GetVendorOk() (string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVendor

`func (o *TapeDrive) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### SetVendor

`func (o *TapeDrive) SetVendor(v string)`

SetVendor gets a reference to the given string and assigns it to the Vendor field.

### GetProduct

`func (o *TapeDrive) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *TapeDrive) GetProductOk() (string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProduct

`func (o *TapeDrive) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### SetProduct

`func (o *TapeDrive) SetProduct(v string)`

SetProduct gets a reference to the given string and assigns it to the Product field.

### GetFirmware

`func (o *TapeDrive) GetFirmware() string`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *TapeDrive) GetFirmwareOk() (string, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFirmware

`func (o *TapeDrive) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### SetFirmware

`func (o *TapeDrive) SetFirmware(v string)`

SetFirmware gets a reference to the given string and assigns it to the Firmware field.

### GetDevice

`func (o *TapeDrive) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *TapeDrive) GetDeviceOk() (string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDevice

`func (o *TapeDrive) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDevice

`func (o *TapeDrive) SetDevice(v string)`

SetDevice gets a reference to the given string and assigns it to the Device field.

### GetSgdevice

`func (o *TapeDrive) GetSgdevice() string`

GetSgdevice returns the Sgdevice field if non-nil, zero value otherwise.

### GetSgdeviceOk

`func (o *TapeDrive) GetSgdeviceOk() (string, bool)`

GetSgdeviceOk returns a tuple with the Sgdevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSgdevice

`func (o *TapeDrive) HasSgdevice() bool`

HasSgdevice returns a boolean if a field has been set.

### SetSgdevice

`func (o *TapeDrive) SetSgdevice(v string)`

SetSgdevice gets a reference to the given string and assigns it to the Sgdevice field.

### GetId

`func (o *TapeDrive) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TapeDrive) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *TapeDrive) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *TapeDrive) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetTapeLibraryId

`func (o *TapeDrive) GetTapeLibraryId() int32`

GetTapeLibraryId returns the TapeLibraryId field if non-nil, zero value otherwise.

### GetTapeLibraryIdOk

`func (o *TapeDrive) GetTapeLibraryIdOk() (int32, bool)`

GetTapeLibraryIdOk returns a tuple with the TapeLibraryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTapeLibraryId

`func (o *TapeDrive) HasTapeLibraryId() bool`

HasTapeLibraryId returns a boolean if a field has been set.

### SetTapeLibraryId

`func (o *TapeDrive) SetTapeLibraryId(v int32)`

SetTapeLibraryId gets a reference to the given int32 and assigns it to the TapeLibraryId field.

### GetName

`func (o *TapeDrive) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TapeDrive) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *TapeDrive) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *TapeDrive) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetComment

`func (o *TapeDrive) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TapeDrive) GetCommentOk() (string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComment

`func (o *TapeDrive) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetComment

`func (o *TapeDrive) SetComment(v string)`

SetComment gets a reference to the given string and assigns it to the Comment field.

### GetLibso

`func (o *TapeDrive) GetLibso() string`

GetLibso returns the Libso field if non-nil, zero value otherwise.

### GetLibsoOk

`func (o *TapeDrive) GetLibsoOk() (string, bool)`

GetLibsoOk returns a tuple with the Libso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLibso

`func (o *TapeDrive) HasLibso() bool`

HasLibso returns a boolean if a field has been set.

### SetLibso

`func (o *TapeDrive) SetLibso(v string)`

SetLibso gets a reference to the given string and assigns it to the Libso field.

### GetAcs

`func (o *TapeDrive) GetAcs() int32`

GetAcs returns the Acs field if non-nil, zero value otherwise.

### GetAcsOk

`func (o *TapeDrive) GetAcsOk() (int32, bool)`

GetAcsOk returns a tuple with the Acs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAcs

`func (o *TapeDrive) HasAcs() bool`

HasAcs returns a boolean if a field has been set.

### SetAcs

`func (o *TapeDrive) SetAcs(v int32)`

SetAcs gets a reference to the given int32 and assigns it to the Acs field.

### GetLsm

`func (o *TapeDrive) GetLsm() int32`

GetLsm returns the Lsm field if non-nil, zero value otherwise.

### GetLsmOk

`func (o *TapeDrive) GetLsmOk() (int32, bool)`

GetLsmOk returns a tuple with the Lsm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLsm

`func (o *TapeDrive) HasLsm() bool`

HasLsm returns a boolean if a field has been set.

### SetLsm

`func (o *TapeDrive) SetLsm(v int32)`

SetLsm gets a reference to the given int32 and assigns it to the Lsm field.

### GetPanel

`func (o *TapeDrive) GetPanel() int32`

GetPanel returns the Panel field if non-nil, zero value otherwise.

### GetPanelOk

`func (o *TapeDrive) GetPanelOk() (int32, bool)`

GetPanelOk returns a tuple with the Panel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPanel

`func (o *TapeDrive) HasPanel() bool`

HasPanel returns a boolean if a field has been set.

### SetPanel

`func (o *TapeDrive) SetPanel(v int32)`

SetPanel gets a reference to the given int32 and assigns it to the Panel field.

### GetTransport

`func (o *TapeDrive) GetTransport() int32`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *TapeDrive) GetTransportOk() (int32, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTransport

`func (o *TapeDrive) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### SetTransport

`func (o *TapeDrive) SetTransport(v int32)`

SetTransport gets a reference to the given int32 and assigns it to the Transport field.

### GetStatus

`func (o *TapeDrive) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TapeDrive) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *TapeDrive) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *TapeDrive) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.

### GetFull

`func (o *TapeDrive) GetFull() int32`

GetFull returns the Full field if non-nil, zero value otherwise.

### GetFullOk

`func (o *TapeDrive) GetFullOk() (int32, bool)`

GetFullOk returns a tuple with the Full field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFull

`func (o *TapeDrive) HasFull() bool`

HasFull returns a boolean if a field has been set.

### SetFull

`func (o *TapeDrive) SetFull(v int32)`

SetFull gets a reference to the given int32 and assigns it to the Full field.

### GetMountCount

`func (o *TapeDrive) GetMountCount() int32`

GetMountCount returns the MountCount field if non-nil, zero value otherwise.

### GetMountCountOk

`func (o *TapeDrive) GetMountCountOk() (int32, bool)`

GetMountCountOk returns a tuple with the MountCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMountCount

`func (o *TapeDrive) HasMountCount() bool`

HasMountCount returns a boolean if a field has been set.

### SetMountCount

`func (o *TapeDrive) SetMountCount(v int32)`

SetMountCount gets a reference to the given int32 and assigns it to the MountCount field.

### GetUseTo

`func (o *TapeDrive) GetUseTo() string`

GetUseTo returns the UseTo field if non-nil, zero value otherwise.

### GetUseToOk

`func (o *TapeDrive) GetUseToOk() (string, bool)`

GetUseToOk returns a tuple with the UseTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUseTo

`func (o *TapeDrive) HasUseTo() bool`

HasUseTo returns a boolean if a field has been set.

### SetUseTo

`func (o *TapeDrive) SetUseTo(v string)`

SetUseTo gets a reference to the given string and assigns it to the UseTo field.

### GetUseBy

`func (o *TapeDrive) GetUseBy() string`

GetUseBy returns the UseBy field if non-nil, zero value otherwise.

### GetUseByOk

`func (o *TapeDrive) GetUseByOk() (string, bool)`

GetUseByOk returns a tuple with the UseBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUseBy

`func (o *TapeDrive) HasUseBy() bool`

HasUseBy returns a boolean if a field has been set.

### SetUseBy

`func (o *TapeDrive) SetUseBy(v string)`

SetUseBy gets a reference to the given string and assigns it to the UseBy field.

### GetBarcode

`func (o *TapeDrive) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *TapeDrive) GetBarcodeOk() (string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBarcode

`func (o *TapeDrive) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### SetBarcode

`func (o *TapeDrive) SetBarcode(v string)`

SetBarcode gets a reference to the given string and assigns it to the Barcode field.

### GetTaskId

`func (o *TapeDrive) GetTaskId() int32`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *TapeDrive) GetTaskIdOk() (int32, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTaskId

`func (o *TapeDrive) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskId

`func (o *TapeDrive) SetTaskId(v int32)`

SetTaskId gets a reference to the given int32 and assigns it to the TaskId field.

### GetUseFileProcessedSize

`func (o *TapeDrive) GetUseFileProcessedSize() int32`

GetUseFileProcessedSize returns the UseFileProcessedSize field if non-nil, zero value otherwise.

### GetUseFileProcessedSizeOk

`func (o *TapeDrive) GetUseFileProcessedSizeOk() (int32, bool)`

GetUseFileProcessedSizeOk returns a tuple with the UseFileProcessedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUseFileProcessedSize

`func (o *TapeDrive) HasUseFileProcessedSize() bool`

HasUseFileProcessedSize returns a boolean if a field has been set.

### SetUseFileProcessedSize

`func (o *TapeDrive) SetUseFileProcessedSize(v int32)`

SetUseFileProcessedSize gets a reference to the given int32 and assigns it to the UseFileProcessedSize field.

### GetUseFileSizeToProcess

`func (o *TapeDrive) GetUseFileSizeToProcess() int32`

GetUseFileSizeToProcess returns the UseFileSizeToProcess field if non-nil, zero value otherwise.

### GetUseFileSizeToProcessOk

`func (o *TapeDrive) GetUseFileSizeToProcessOk() (int32, bool)`

GetUseFileSizeToProcessOk returns a tuple with the UseFileSizeToProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUseFileSizeToProcess

`func (o *TapeDrive) HasUseFileSizeToProcess() bool`

HasUseFileSizeToProcess returns a boolean if a field has been set.

### SetUseFileSizeToProcess

`func (o *TapeDrive) SetUseFileSizeToProcess(v int32)`

SetUseFileSizeToProcess gets a reference to the given int32 and assigns it to the UseFileSizeToProcess field.

### GetUseFileNameProcessed

`func (o *TapeDrive) GetUseFileNameProcessed() string`

GetUseFileNameProcessed returns the UseFileNameProcessed field if non-nil, zero value otherwise.

### GetUseFileNameProcessedOk

`func (o *TapeDrive) GetUseFileNameProcessedOk() (string, bool)`

GetUseFileNameProcessedOk returns a tuple with the UseFileNameProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUseFileNameProcessed

`func (o *TapeDrive) HasUseFileNameProcessed() bool`

HasUseFileNameProcessed returns a boolean if a field has been set.

### SetUseFileNameProcessed

`func (o *TapeDrive) SetUseFileNameProcessed(v string)`

SetUseFileNameProcessed gets a reference to the given string and assigns it to the UseFileNameProcessed field.

### GetBandwidth

`func (o *TapeDrive) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *TapeDrive) GetBandwidthOk() (int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBandwidth

`func (o *TapeDrive) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### SetBandwidth

`func (o *TapeDrive) SetBandwidth(v int32)`

SetBandwidth gets a reference to the given int32 and assigns it to the Bandwidth field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



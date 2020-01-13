/*
 * Nodeum API
 *
 * # About  This document describes the Nodeum API version 2:  If you are looking for any information about the product itself, reach the product website https://www.nodeum.io. You can also contact us at this email address : info@nodeum.io  # Filter parameters When browsing a list of items, multiple filter parameters may be applied. Some operators can be added to the value as a prefix:  - `=` value is equal. Default operator, may be omitted  - `!=` value is different  - `>` greater than  - `>=` greater than or equal  - `<` lower than  - `>=` lower than or equal  - `><` included in list, items should be separated by `|`  - `!><` not included in list, items should be separated by `|`  - `~` pattern matching, may include `%` (any characters) and `_` (one character)  - `!~` pattern not matching, may include `%` (any characters) and `_` (one character)  
 *
 * API version: 2.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nodeum

import (
	"bytes"
	"encoding/json"
)

// TapeLibrary struct for TapeLibrary
type TapeLibrary struct {
	Serial *string `json:"serial,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
	Vendor *string `json:"vendor,omitempty"`
	Product *string `json:"product,omitempty"`
	Firmware *string `json:"firmware,omitempty"`
	Device *string `json:"device,omitempty"`
	Acs *int32 `json:"acs,omitempty"`
	StorageSlots *int32 `json:"storage_slots,omitempty"`
	StorageSlotsAddress *int32 `json:"storage_slots_address,omitempty"`
	IoSlots *int32 `json:"io_slots,omitempty"`
	IoSlotsAddress *int32 `json:"io_slots_address,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Comment *string `json:"comment,omitempty"`
	Libso *string `json:"libso,omitempty"`
	Status *string `json:"status,omitempty"`
	Price *string `json:"price,omitempty"`
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *TapeLibrary) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TapeLibrary) GetSerialOk() (string, bool) {
	if o == nil || o.Serial == nil {
		var ret string
		return ret, false
	}
	return *o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *TapeLibrary) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *TapeLibrary) SetSerial(v string) {
	o.Serial = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *TapeLibrary) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TapeLibrary) GetProtocolOk() (string, bool) {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret, false
	}
	return *o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *TapeLibrary) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *TapeLibrary) SetProtocol(v string) {
	o.Protocol = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *TapeLibrary) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TapeLibrary) GetVendorOk() (string, bool) {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret, false
	}
	return *o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *TapeLibrary) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *TapeLibrary) SetVendor(v string) {
	o.Vendor = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *TapeLibrary) GetProduct() string {
	if o == nil || o.Product == nil {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TapeLibrary) GetProductOk() (string, bool) {
	if o == nil || o.Product == nil {
		var ret string
		return ret, false
	}
	return *o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *TapeLibrary) HasProduct() bool {
	if o != nil && o.Product != nil {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *TapeLibrary) SetProduct(v string) {
	o.Product = &v
}

// GetFirmware returns the Firmware field value if set, zero value otherwise.
func (o *TapeLibrary) GetFirmware() string {
	if o == nil || o.Firmware == nil {
		var ret string
		return ret
	}
	return *o.Firmware
}

// GetFirmwareOk returns a tuple with the Firmware field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TapeLibrary) GetFirmwareOk() (string, bool) {
	if o == nil || o.Firmware == nil {
		var ret string
		return ret, false
	}
	return *o.Firmware, true
}

// HasFirmware returns a boolean if a field has been set.
func (o *TapeLibrary) HasFirmware() bool {
	if o != nil && o.Firmware != nil {
		return true
	}

	return false
}

// SetFirmware gets a reference to the given string and assigns it to the Firmware field.
func (o *TapeLibrary) SetFirmware(v string) {
	o.Firmware = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *TapeLibrary) GetDevice() string {
	if o == nil || o.Device == nil {
		var ret string
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TapeLibrary) GetDeviceOk() (string, bool) {
	if o == nil || o.Device == nil {
		var ret string
		return ret, false
	}
	return *o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *TapeLibrary) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given string and assigns it to the Device field.
func (o *TapeLibrary) SetDevice(v string) {
	o.Device = &v
}

// GetAcs returns the Acs field value if set, zero value otherwise.
func (o *TapeLibrary) GetAcs() int32 {
	if o == nil || o.Acs == nil {
		var ret int32
		return ret
	}
	return *o.Acs
}

// GetAcsOk returns a tuple with the Acs field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TapeLibrary) GetAcsOk() (int32, bool) {
	if o == nil || o.Acs == nil {
		var ret int32
		return ret, false
	}
	return *o.Acs, true
}

// HasAcs returns a boolean if a field has been set.
func (o *TapeLibrary) HasAcs() bool {
	if o != nil && o.Acs != nil {
		return true
	}

	return false
}

// SetAcs gets a reference to the given int32 and assigns it to the Acs field.
func (o *TapeLibrary) SetAcs(v int32) {
	o.Acs = &v
}

// GetStorageSlots returns the StorageSlots field value if set, zero value otherwise.
func (o *TapeLibrary) GetStorageSlots() int32 {
	if o == nil || o.StorageSlots == nil {
		var ret int32
		return ret
	}
	return *o.StorageSlots
}

// GetStorageSlotsOk returns a tuple with the StorageSlots field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TapeLibrary) GetStorageSlotsOk() (int32, bool) {
	if o == nil || o.StorageSlots == nil {
		var ret int32
		return ret, false
	}
	return *o.StorageSlots, true
}

// HasStorageSlots returns a boolean if a field has been set.
func (o *TapeLibrary) HasStorageSlots() bool {
	if o != nil && o.StorageSlots != nil {
		return true
	}

	return false
}

// SetStorageSlots gets a reference to the given int32 and assigns it to the StorageSlots field.
func (o *TapeLibrary) SetStorageSlots(v int32) {
	o.StorageSlots = &v
}

// GetStorageSlotsAddress returns the StorageSlotsAddress field value if set, zero value otherwise.
func (o *TapeLibrary) GetStorageSlotsAddress() int32 {
	if o == nil || o.StorageSlotsAddress == nil {
		var ret int32
		return ret
	}
	return *o.StorageSlotsAddress
}

// GetStorageSlotsAddressOk returns a tuple with the StorageSlotsAddress field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TapeLibrary) GetStorageSlotsAddressOk() (int32, bool) {
	if o == nil || o.StorageSlotsAddress == nil {
		var ret int32
		return ret, false
	}
	return *o.StorageSlotsAddress, true
}

// HasStorageSlotsAddress returns a boolean if a field has been set.
func (o *TapeLibrary) HasStorageSlotsAddress() bool {
	if o != nil && o.StorageSlotsAddress != nil {
		return true
	}

	return false
}

// SetStorageSlotsAddress gets a reference to the given int32 and assigns it to the StorageSlotsAddress field.
func (o *TapeLibrary) SetStorageSlotsAddress(v int32) {
	o.StorageSlotsAddress = &v
}

// GetIoSlots returns the IoSlots field value if set, zero value otherwise.
func (o *TapeLibrary) GetIoSlots() int32 {
	if o == nil || o.IoSlots == nil {
		var ret int32
		return ret
	}
	return *o.IoSlots
}

// GetIoSlotsOk returns a tuple with the IoSlots field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TapeLibrary) GetIoSlotsOk() (int32, bool) {
	if o == nil || o.IoSlots == nil {
		var ret int32
		return ret, false
	}
	return *o.IoSlots, true
}

// HasIoSlots returns a boolean if a field has been set.
func (o *TapeLibrary) HasIoSlots() bool {
	if o != nil && o.IoSlots != nil {
		return true
	}

	return false
}

// SetIoSlots gets a reference to the given int32 and assigns it to the IoSlots field.
func (o *TapeLibrary) SetIoSlots(v int32) {
	o.IoSlots = &v
}

// GetIoSlotsAddress returns the IoSlotsAddress field value if set, zero value otherwise.
func (o *TapeLibrary) GetIoSlotsAddress() int32 {
	if o == nil || o.IoSlotsAddress == nil {
		var ret int32
		return ret
	}
	return *o.IoSlotsAddress
}

// GetIoSlotsAddressOk returns a tuple with the IoSlotsAddress field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TapeLibrary) GetIoSlotsAddressOk() (int32, bool) {
	if o == nil || o.IoSlotsAddress == nil {
		var ret int32
		return ret, false
	}
	return *o.IoSlotsAddress, true
}

// HasIoSlotsAddress returns a boolean if a field has been set.
func (o *TapeLibrary) HasIoSlotsAddress() bool {
	if o != nil && o.IoSlotsAddress != nil {
		return true
	}

	return false
}

// SetIoSlotsAddress gets a reference to the given int32 and assigns it to the IoSlotsAddress field.
func (o *TapeLibrary) SetIoSlotsAddress(v int32) {
	o.IoSlotsAddress = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TapeLibrary) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TapeLibrary) GetIdOk() (int32, bool) {
	if o == nil || o.Id == nil {
		var ret int32
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TapeLibrary) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TapeLibrary) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TapeLibrary) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TapeLibrary) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TapeLibrary) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TapeLibrary) SetName(v string) {
	o.Name = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *TapeLibrary) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TapeLibrary) GetCommentOk() (string, bool) {
	if o == nil || o.Comment == nil {
		var ret string
		return ret, false
	}
	return *o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *TapeLibrary) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *TapeLibrary) SetComment(v string) {
	o.Comment = &v
}

// GetLibso returns the Libso field value if set, zero value otherwise.
func (o *TapeLibrary) GetLibso() string {
	if o == nil || o.Libso == nil {
		var ret string
		return ret
	}
	return *o.Libso
}

// GetLibsoOk returns a tuple with the Libso field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TapeLibrary) GetLibsoOk() (string, bool) {
	if o == nil || o.Libso == nil {
		var ret string
		return ret, false
	}
	return *o.Libso, true
}

// HasLibso returns a boolean if a field has been set.
func (o *TapeLibrary) HasLibso() bool {
	if o != nil && o.Libso != nil {
		return true
	}

	return false
}

// SetLibso gets a reference to the given string and assigns it to the Libso field.
func (o *TapeLibrary) SetLibso(v string) {
	o.Libso = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TapeLibrary) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TapeLibrary) GetStatusOk() (string, bool) {
	if o == nil || o.Status == nil {
		var ret string
		return ret, false
	}
	return *o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TapeLibrary) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TapeLibrary) SetStatus(v string) {
	o.Status = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *TapeLibrary) GetPrice() string {
	if o == nil || o.Price == nil {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TapeLibrary) GetPriceOk() (string, bool) {
	if o == nil || o.Price == nil {
		var ret string
		return ret, false
	}
	return *o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *TapeLibrary) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *TapeLibrary) SetPrice(v string) {
	o.Price = &v
}

type NullableTapeLibrary struct {
	Value TapeLibrary
	ExplicitNull bool
}

func (v NullableTapeLibrary) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableTapeLibrary) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

/*
 * Nodeum API Reference
 *
 * The Nodeum API makes it easy to tap into the digital data mesh that runs across your organisation. Make requests to our API endpoints and we’ll give you everything you need to interconnect your business workflows with your storage.  All production API requests are made to:  http://nodeumhostname/api/  The current production version of the API is v1.   **REST** The Nodeum API is a RESTful API. This means that the API is designed to allow you to get, create, update, & delete objects with the HTTP verbs GET, POST, PUT, PATCH, & DELETE.  **JSON** The Nodeum API speaks exclusively in JSON. This means that you should always set the Content-Type header to application/json to ensure that your requests are properly accepted and processed by the API.  **Authentication** All API calls require user-password authentication.   **Cross-Origin Resource Sharing** The Nodeum API supports CORS for communicating from Javascript for these endpoints. You will need to specify an Origin URI when creating your application to allow for CORS to be whitelisted for your domain.   **Pagination** Some endpoints such as File Listing return a potentially lengthy array of objects. In order to keep the response sizes manageable the API will take advantage of pagination. Pagination is a mechanism for returning a subset of the results for a request and allowing for subsequent requests to “page” through the rest of the results until the end is reached. Paginated endpoints follow a standard interface that accepts two query parameters, limit and offset, and return a payload that follows a standard form. These parameters names and their behavior are borrowed from SQL LIMIT and OFFSET keywords.  **Versioning** The Nodeum API is constantly being worked on to add features, make improvements, and fix bugs. This means that you should expect changes to be introduced and documented.   However, there are some changes or additions that are considered backwards-compatible and your applications should be flexible enough to handle them. These include:  - Adding new endpoints to the API - Adding new attributes to the response of an existing endpoint - Changing the order of attributes of responses (JSON by definition is an object of unordered key/value pairs)   **Filter parameters** When browsing a list of items, multiple filter parameters may be applied. Some operators can be added to the value as a prefix:  - `=` value is equal. Default operator, may be omitted  - `!=` value is different  - `>` greater than  - `>=` greater than or equal  - `<` lower than  - `>=` lower than or equal  - `><` included in list, items should be separated by `|`  - `!><` not included in list, items should be separated by `|`  - `~` pattern matching, may include `%` (any characters) and `_` (one character)  - `!~` pattern not matching, may include `%` (any characters) and `_` (one character)
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nodeum

import (
	"bytes"
	"encoding/json"
)

// TapeLibraryDevice struct for TapeLibraryDevice
type TapeLibraryDevice struct {
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
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *TapeLibraryDevice) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TapeLibraryDevice) GetSerialOk() (string, bool) {
	if o == nil || o.Serial == nil {
		var ret string
		return ret, false
	}
	return *o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *TapeLibraryDevice) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *TapeLibraryDevice) SetSerial(v string) {
	o.Serial = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *TapeLibraryDevice) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TapeLibraryDevice) GetProtocolOk() (string, bool) {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret, false
	}
	return *o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *TapeLibraryDevice) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *TapeLibraryDevice) SetProtocol(v string) {
	o.Protocol = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *TapeLibraryDevice) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TapeLibraryDevice) GetVendorOk() (string, bool) {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret, false
	}
	return *o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *TapeLibraryDevice) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *TapeLibraryDevice) SetVendor(v string) {
	o.Vendor = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *TapeLibraryDevice) GetProduct() string {
	if o == nil || o.Product == nil {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TapeLibraryDevice) GetProductOk() (string, bool) {
	if o == nil || o.Product == nil {
		var ret string
		return ret, false
	}
	return *o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *TapeLibraryDevice) HasProduct() bool {
	if o != nil && o.Product != nil {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *TapeLibraryDevice) SetProduct(v string) {
	o.Product = &v
}

// GetFirmware returns the Firmware field value if set, zero value otherwise.
func (o *TapeLibraryDevice) GetFirmware() string {
	if o == nil || o.Firmware == nil {
		var ret string
		return ret
	}
	return *o.Firmware
}

// GetFirmwareOk returns a tuple with the Firmware field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TapeLibraryDevice) GetFirmwareOk() (string, bool) {
	if o == nil || o.Firmware == nil {
		var ret string
		return ret, false
	}
	return *o.Firmware, true
}

// HasFirmware returns a boolean if a field has been set.
func (o *TapeLibraryDevice) HasFirmware() bool {
	if o != nil && o.Firmware != nil {
		return true
	}

	return false
}

// SetFirmware gets a reference to the given string and assigns it to the Firmware field.
func (o *TapeLibraryDevice) SetFirmware(v string) {
	o.Firmware = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *TapeLibraryDevice) GetDevice() string {
	if o == nil || o.Device == nil {
		var ret string
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TapeLibraryDevice) GetDeviceOk() (string, bool) {
	if o == nil || o.Device == nil {
		var ret string
		return ret, false
	}
	return *o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *TapeLibraryDevice) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given string and assigns it to the Device field.
func (o *TapeLibraryDevice) SetDevice(v string) {
	o.Device = &v
}

// GetAcs returns the Acs field value if set, zero value otherwise.
func (o *TapeLibraryDevice) GetAcs() int32 {
	if o == nil || o.Acs == nil {
		var ret int32
		return ret
	}
	return *o.Acs
}

// GetAcsOk returns a tuple with the Acs field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TapeLibraryDevice) GetAcsOk() (int32, bool) {
	if o == nil || o.Acs == nil {
		var ret int32
		return ret, false
	}
	return *o.Acs, true
}

// HasAcs returns a boolean if a field has been set.
func (o *TapeLibraryDevice) HasAcs() bool {
	if o != nil && o.Acs != nil {
		return true
	}

	return false
}

// SetAcs gets a reference to the given int32 and assigns it to the Acs field.
func (o *TapeLibraryDevice) SetAcs(v int32) {
	o.Acs = &v
}

// GetStorageSlots returns the StorageSlots field value if set, zero value otherwise.
func (o *TapeLibraryDevice) GetStorageSlots() int32 {
	if o == nil || o.StorageSlots == nil {
		var ret int32
		return ret
	}
	return *o.StorageSlots
}

// GetStorageSlotsOk returns a tuple with the StorageSlots field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TapeLibraryDevice) GetStorageSlotsOk() (int32, bool) {
	if o == nil || o.StorageSlots == nil {
		var ret int32
		return ret, false
	}
	return *o.StorageSlots, true
}

// HasStorageSlots returns a boolean if a field has been set.
func (o *TapeLibraryDevice) HasStorageSlots() bool {
	if o != nil && o.StorageSlots != nil {
		return true
	}

	return false
}

// SetStorageSlots gets a reference to the given int32 and assigns it to the StorageSlots field.
func (o *TapeLibraryDevice) SetStorageSlots(v int32) {
	o.StorageSlots = &v
}

// GetStorageSlotsAddress returns the StorageSlotsAddress field value if set, zero value otherwise.
func (o *TapeLibraryDevice) GetStorageSlotsAddress() int32 {
	if o == nil || o.StorageSlotsAddress == nil {
		var ret int32
		return ret
	}
	return *o.StorageSlotsAddress
}

// GetStorageSlotsAddressOk returns a tuple with the StorageSlotsAddress field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TapeLibraryDevice) GetStorageSlotsAddressOk() (int32, bool) {
	if o == nil || o.StorageSlotsAddress == nil {
		var ret int32
		return ret, false
	}
	return *o.StorageSlotsAddress, true
}

// HasStorageSlotsAddress returns a boolean if a field has been set.
func (o *TapeLibraryDevice) HasStorageSlotsAddress() bool {
	if o != nil && o.StorageSlotsAddress != nil {
		return true
	}

	return false
}

// SetStorageSlotsAddress gets a reference to the given int32 and assigns it to the StorageSlotsAddress field.
func (o *TapeLibraryDevice) SetStorageSlotsAddress(v int32) {
	o.StorageSlotsAddress = &v
}

// GetIoSlots returns the IoSlots field value if set, zero value otherwise.
func (o *TapeLibraryDevice) GetIoSlots() int32 {
	if o == nil || o.IoSlots == nil {
		var ret int32
		return ret
	}
	return *o.IoSlots
}

// GetIoSlotsOk returns a tuple with the IoSlots field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TapeLibraryDevice) GetIoSlotsOk() (int32, bool) {
	if o == nil || o.IoSlots == nil {
		var ret int32
		return ret, false
	}
	return *o.IoSlots, true
}

// HasIoSlots returns a boolean if a field has been set.
func (o *TapeLibraryDevice) HasIoSlots() bool {
	if o != nil && o.IoSlots != nil {
		return true
	}

	return false
}

// SetIoSlots gets a reference to the given int32 and assigns it to the IoSlots field.
func (o *TapeLibraryDevice) SetIoSlots(v int32) {
	o.IoSlots = &v
}

// GetIoSlotsAddress returns the IoSlotsAddress field value if set, zero value otherwise.
func (o *TapeLibraryDevice) GetIoSlotsAddress() int32 {
	if o == nil || o.IoSlotsAddress == nil {
		var ret int32
		return ret
	}
	return *o.IoSlotsAddress
}

// GetIoSlotsAddressOk returns a tuple with the IoSlotsAddress field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TapeLibraryDevice) GetIoSlotsAddressOk() (int32, bool) {
	if o == nil || o.IoSlotsAddress == nil {
		var ret int32
		return ret, false
	}
	return *o.IoSlotsAddress, true
}

// HasIoSlotsAddress returns a boolean if a field has been set.
func (o *TapeLibraryDevice) HasIoSlotsAddress() bool {
	if o != nil && o.IoSlotsAddress != nil {
		return true
	}

	return false
}

// SetIoSlotsAddress gets a reference to the given int32 and assigns it to the IoSlotsAddress field.
func (o *TapeLibraryDevice) SetIoSlotsAddress(v int32) {
	o.IoSlotsAddress = &v
}

type NullableTapeLibraryDevice struct {
	Value TapeLibraryDevice
	ExplicitNull bool
}

func (v NullableTapeLibraryDevice) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableTapeLibraryDevice) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

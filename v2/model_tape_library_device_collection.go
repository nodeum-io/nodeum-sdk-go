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

// TapeLibraryDeviceCollection struct for TapeLibraryDeviceCollection
type TapeLibraryDeviceCollection struct {
	TapeLibraries *[]TapeLibraryDevice `json:"tape_libraries,omitempty"`
}

// GetTapeLibraries returns the TapeLibraries field value if set, zero value otherwise.
func (o *TapeLibraryDeviceCollection) GetTapeLibraries() []TapeLibraryDevice {
	if o == nil || o.TapeLibraries == nil {
		var ret []TapeLibraryDevice
		return ret
	}
	return *o.TapeLibraries
}

// GetTapeLibrariesOk returns a tuple with the TapeLibraries field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TapeLibraryDeviceCollection) GetTapeLibrariesOk() ([]TapeLibraryDevice, bool) {
	if o == nil || o.TapeLibraries == nil {
		var ret []TapeLibraryDevice
		return ret, false
	}
	return *o.TapeLibraries, true
}

// HasTapeLibraries returns a boolean if a field has been set.
func (o *TapeLibraryDeviceCollection) HasTapeLibraries() bool {
	if o != nil && o.TapeLibraries != nil {
		return true
	}

	return false
}

// SetTapeLibraries gets a reference to the given []TapeLibraryDevice and assigns it to the TapeLibraries field.
func (o *TapeLibraryDeviceCollection) SetTapeLibraries(v []TapeLibraryDevice) {
	o.TapeLibraries = &v
}

type NullableTapeLibraryDeviceCollection struct {
	Value TapeLibraryDeviceCollection
	ExplicitNull bool
}

func (v NullableTapeLibraryDeviceCollection) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableTapeLibraryDeviceCollection) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

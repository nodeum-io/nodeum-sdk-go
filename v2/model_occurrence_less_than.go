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

// OccurrenceLessThan Must have less than %{count} occurrences (currently have %{value})
type OccurrenceLessThan struct {
	Error string `json:"error"`
	// Expected maximum number of occurrences
	Count *int32 `json:"count,omitempty"`
	// Current number of occurrences
	Value *int32 `json:"value,omitempty"`
}

// GetError returns the Error field value
func (o *OccurrenceLessThan) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// SetError sets field value
func (o *OccurrenceLessThan) SetError(v string) {
	o.Error = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *OccurrenceLessThan) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OccurrenceLessThan) GetCountOk() (int32, bool) {
	if o == nil || o.Count == nil {
		var ret int32
		return ret, false
	}
	return *o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *OccurrenceLessThan) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *OccurrenceLessThan) SetCount(v int32) {
	o.Count = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *OccurrenceLessThan) GetValue() int32 {
	if o == nil || o.Value == nil {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OccurrenceLessThan) GetValueOk() (int32, bool) {
	if o == nil || o.Value == nil {
		var ret int32
		return ret, false
	}
	return *o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *OccurrenceLessThan) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *OccurrenceLessThan) SetValue(v int32) {
	o.Value = &v
}

type NullableOccurrenceLessThan struct {
	Value OccurrenceLessThan
	ExplicitNull bool
}

func (v NullableOccurrenceLessThan) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableOccurrenceLessThan) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

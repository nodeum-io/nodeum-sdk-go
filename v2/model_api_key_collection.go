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

// ApiKeyCollection struct for ApiKeyCollection
type ApiKeyCollection struct {
	ApiKeys *[]ApiKey `json:"api_keys,omitempty"`
	Count *int32 `json:"count,omitempty"`
}

// GetApiKeys returns the ApiKeys field value if set, zero value otherwise.
func (o *ApiKeyCollection) GetApiKeys() []ApiKey {
	if o == nil || o.ApiKeys == nil {
		var ret []ApiKey
		return ret
	}
	return *o.ApiKeys
}

// GetApiKeysOk returns a tuple with the ApiKeys field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyCollection) GetApiKeysOk() ([]ApiKey, bool) {
	if o == nil || o.ApiKeys == nil {
		var ret []ApiKey
		return ret, false
	}
	return *o.ApiKeys, true
}

// HasApiKeys returns a boolean if a field has been set.
func (o *ApiKeyCollection) HasApiKeys() bool {
	if o != nil && o.ApiKeys != nil {
		return true
	}

	return false
}

// SetApiKeys gets a reference to the given []ApiKey and assigns it to the ApiKeys field.
func (o *ApiKeyCollection) SetApiKeys(v []ApiKey) {
	o.ApiKeys = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ApiKeyCollection) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyCollection) GetCountOk() (int32, bool) {
	if o == nil || o.Count == nil {
		var ret int32
		return ret, false
	}
	return *o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ApiKeyCollection) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ApiKeyCollection) SetCount(v int32) {
	o.Count = &v
}

type NullableApiKeyCollection struct {
	Value ApiKeyCollection
	ExplicitNull bool
}

func (v NullableApiKeyCollection) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableApiKeyCollection) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

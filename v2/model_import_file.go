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

// ImportFile struct for ImportFile
type ImportFile struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Parent *int32 `json:"parent,omitempty"`
	PrimaryId *int32 `json:"primary_id,omitempty"`
	Type *string `json:"type,omitempty"`
	Permission *int32 `json:"permission,omitempty"`
	Size *int32 `json:"size,omitempty"`
	ChangeDate *string `json:"change_date,omitempty"`
	ModificationDate *string `json:"modification_date,omitempty"`
	AccessDate *string `json:"access_date,omitempty"`
	Uid *int32 `json:"uid,omitempty"`
	Gid *int32 `json:"gid,omitempty"`
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ImportFile) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ImportFile) GetIdOk() (int32, bool) {
	if o == nil || o.Id == nil {
		var ret int32
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ImportFile) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ImportFile) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ImportFile) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ImportFile) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ImportFile) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ImportFile) SetName(v string) {
	o.Name = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *ImportFile) GetParent() int32 {
	if o == nil || o.Parent == nil {
		var ret int32
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ImportFile) GetParentOk() (int32, bool) {
	if o == nil || o.Parent == nil {
		var ret int32
		return ret, false
	}
	return *o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *ImportFile) HasParent() bool {
	if o != nil && o.Parent != nil {
		return true
	}

	return false
}

// SetParent gets a reference to the given int32 and assigns it to the Parent field.
func (o *ImportFile) SetParent(v int32) {
	o.Parent = &v
}

// GetPrimaryId returns the PrimaryId field value if set, zero value otherwise.
func (o *ImportFile) GetPrimaryId() int32 {
	if o == nil || o.PrimaryId == nil {
		var ret int32
		return ret
	}
	return *o.PrimaryId
}

// GetPrimaryIdOk returns a tuple with the PrimaryId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ImportFile) GetPrimaryIdOk() (int32, bool) {
	if o == nil || o.PrimaryId == nil {
		var ret int32
		return ret, false
	}
	return *o.PrimaryId, true
}

// HasPrimaryId returns a boolean if a field has been set.
func (o *ImportFile) HasPrimaryId() bool {
	if o != nil && o.PrimaryId != nil {
		return true
	}

	return false
}

// SetPrimaryId gets a reference to the given int32 and assigns it to the PrimaryId field.
func (o *ImportFile) SetPrimaryId(v int32) {
	o.PrimaryId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ImportFile) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ImportFile) GetTypeOk() (string, bool) {
	if o == nil || o.Type == nil {
		var ret string
		return ret, false
	}
	return *o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ImportFile) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ImportFile) SetType(v string) {
	o.Type = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *ImportFile) GetPermission() int32 {
	if o == nil || o.Permission == nil {
		var ret int32
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ImportFile) GetPermissionOk() (int32, bool) {
	if o == nil || o.Permission == nil {
		var ret int32
		return ret, false
	}
	return *o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *ImportFile) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given int32 and assigns it to the Permission field.
func (o *ImportFile) SetPermission(v int32) {
	o.Permission = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ImportFile) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ImportFile) GetSizeOk() (int32, bool) {
	if o == nil || o.Size == nil {
		var ret int32
		return ret, false
	}
	return *o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ImportFile) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *ImportFile) SetSize(v int32) {
	o.Size = &v
}

// GetChangeDate returns the ChangeDate field value if set, zero value otherwise.
func (o *ImportFile) GetChangeDate() string {
	if o == nil || o.ChangeDate == nil {
		var ret string
		return ret
	}
	return *o.ChangeDate
}

// GetChangeDateOk returns a tuple with the ChangeDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ImportFile) GetChangeDateOk() (string, bool) {
	if o == nil || o.ChangeDate == nil {
		var ret string
		return ret, false
	}
	return *o.ChangeDate, true
}

// HasChangeDate returns a boolean if a field has been set.
func (o *ImportFile) HasChangeDate() bool {
	if o != nil && o.ChangeDate != nil {
		return true
	}

	return false
}

// SetChangeDate gets a reference to the given string and assigns it to the ChangeDate field.
func (o *ImportFile) SetChangeDate(v string) {
	o.ChangeDate = &v
}

// GetModificationDate returns the ModificationDate field value if set, zero value otherwise.
func (o *ImportFile) GetModificationDate() string {
	if o == nil || o.ModificationDate == nil {
		var ret string
		return ret
	}
	return *o.ModificationDate
}

// GetModificationDateOk returns a tuple with the ModificationDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ImportFile) GetModificationDateOk() (string, bool) {
	if o == nil || o.ModificationDate == nil {
		var ret string
		return ret, false
	}
	return *o.ModificationDate, true
}

// HasModificationDate returns a boolean if a field has been set.
func (o *ImportFile) HasModificationDate() bool {
	if o != nil && o.ModificationDate != nil {
		return true
	}

	return false
}

// SetModificationDate gets a reference to the given string and assigns it to the ModificationDate field.
func (o *ImportFile) SetModificationDate(v string) {
	o.ModificationDate = &v
}

// GetAccessDate returns the AccessDate field value if set, zero value otherwise.
func (o *ImportFile) GetAccessDate() string {
	if o == nil || o.AccessDate == nil {
		var ret string
		return ret
	}
	return *o.AccessDate
}

// GetAccessDateOk returns a tuple with the AccessDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ImportFile) GetAccessDateOk() (string, bool) {
	if o == nil || o.AccessDate == nil {
		var ret string
		return ret, false
	}
	return *o.AccessDate, true
}

// HasAccessDate returns a boolean if a field has been set.
func (o *ImportFile) HasAccessDate() bool {
	if o != nil && o.AccessDate != nil {
		return true
	}

	return false
}

// SetAccessDate gets a reference to the given string and assigns it to the AccessDate field.
func (o *ImportFile) SetAccessDate(v string) {
	o.AccessDate = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *ImportFile) GetUid() int32 {
	if o == nil || o.Uid == nil {
		var ret int32
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ImportFile) GetUidOk() (int32, bool) {
	if o == nil || o.Uid == nil {
		var ret int32
		return ret, false
	}
	return *o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *ImportFile) HasUid() bool {
	if o != nil && o.Uid != nil {
		return true
	}

	return false
}

// SetUid gets a reference to the given int32 and assigns it to the Uid field.
func (o *ImportFile) SetUid(v int32) {
	o.Uid = &v
}

// GetGid returns the Gid field value if set, zero value otherwise.
func (o *ImportFile) GetGid() int32 {
	if o == nil || o.Gid == nil {
		var ret int32
		return ret
	}
	return *o.Gid
}

// GetGidOk returns a tuple with the Gid field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ImportFile) GetGidOk() (int32, bool) {
	if o == nil || o.Gid == nil {
		var ret int32
		return ret, false
	}
	return *o.Gid, true
}

// HasGid returns a boolean if a field has been set.
func (o *ImportFile) HasGid() bool {
	if o != nil && o.Gid != nil {
		return true
	}

	return false
}

// SetGid gets a reference to the given int32 and assigns it to the Gid field.
func (o *ImportFile) SetGid(v int32) {
	o.Gid = &v
}

type NullableImportFile struct {
	Value ImportFile
	ExplicitNull bool
}

func (v NullableImportFile) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableImportFile) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

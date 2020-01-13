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

// TaskSourceUp struct for TaskSourceUp
type TaskSourceUp struct {
	FileId *int32 `json:"file_id,omitempty"`
	FilePath *string `json:"file_path,omitempty"`
	ImportFileId *int32 `json:"import_file_id,omitempty"`
	ImportFilePath *string `json:"import_file_path,omitempty"`
	TapeId *int32 `json:"tape_id,omitempty"`
	TapeBarcode *string `json:"tape_barcode,omitempty"`
	TapePoolId *int32 `json:"tape_pool_id,omitempty"`
	TapePoolName *string `json:"tape_pool_name,omitempty"`
	CloudPoolId *int32 `json:"cloud_pool_id,omitempty"`
	CloudPoolName *string `json:"cloud_pool_name,omitempty"`
	NasPoolId *int32 `json:"nas_pool_id,omitempty"`
	NasPoolName *string `json:"nas_pool_name,omitempty"`
}

// GetFileId returns the FileId field value if set, zero value otherwise.
func (o *TaskSourceUp) GetFileId() int32 {
	if o == nil || o.FileId == nil {
		var ret int32
		return ret
	}
	return *o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TaskSourceUp) GetFileIdOk() (int32, bool) {
	if o == nil || o.FileId == nil {
		var ret int32
		return ret, false
	}
	return *o.FileId, true
}

// HasFileId returns a boolean if a field has been set.
func (o *TaskSourceUp) HasFileId() bool {
	if o != nil && o.FileId != nil {
		return true
	}

	return false
}

// SetFileId gets a reference to the given int32 and assigns it to the FileId field.
func (o *TaskSourceUp) SetFileId(v int32) {
	o.FileId = &v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *TaskSourceUp) GetFilePath() string {
	if o == nil || o.FilePath == nil {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TaskSourceUp) GetFilePathOk() (string, bool) {
	if o == nil || o.FilePath == nil {
		var ret string
		return ret, false
	}
	return *o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *TaskSourceUp) HasFilePath() bool {
	if o != nil && o.FilePath != nil {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *TaskSourceUp) SetFilePath(v string) {
	o.FilePath = &v
}

// GetImportFileId returns the ImportFileId field value if set, zero value otherwise.
func (o *TaskSourceUp) GetImportFileId() int32 {
	if o == nil || o.ImportFileId == nil {
		var ret int32
		return ret
	}
	return *o.ImportFileId
}

// GetImportFileIdOk returns a tuple with the ImportFileId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TaskSourceUp) GetImportFileIdOk() (int32, bool) {
	if o == nil || o.ImportFileId == nil {
		var ret int32
		return ret, false
	}
	return *o.ImportFileId, true
}

// HasImportFileId returns a boolean if a field has been set.
func (o *TaskSourceUp) HasImportFileId() bool {
	if o != nil && o.ImportFileId != nil {
		return true
	}

	return false
}

// SetImportFileId gets a reference to the given int32 and assigns it to the ImportFileId field.
func (o *TaskSourceUp) SetImportFileId(v int32) {
	o.ImportFileId = &v
}

// GetImportFilePath returns the ImportFilePath field value if set, zero value otherwise.
func (o *TaskSourceUp) GetImportFilePath() string {
	if o == nil || o.ImportFilePath == nil {
		var ret string
		return ret
	}
	return *o.ImportFilePath
}

// GetImportFilePathOk returns a tuple with the ImportFilePath field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TaskSourceUp) GetImportFilePathOk() (string, bool) {
	if o == nil || o.ImportFilePath == nil {
		var ret string
		return ret, false
	}
	return *o.ImportFilePath, true
}

// HasImportFilePath returns a boolean if a field has been set.
func (o *TaskSourceUp) HasImportFilePath() bool {
	if o != nil && o.ImportFilePath != nil {
		return true
	}

	return false
}

// SetImportFilePath gets a reference to the given string and assigns it to the ImportFilePath field.
func (o *TaskSourceUp) SetImportFilePath(v string) {
	o.ImportFilePath = &v
}

// GetTapeId returns the TapeId field value if set, zero value otherwise.
func (o *TaskSourceUp) GetTapeId() int32 {
	if o == nil || o.TapeId == nil {
		var ret int32
		return ret
	}
	return *o.TapeId
}

// GetTapeIdOk returns a tuple with the TapeId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TaskSourceUp) GetTapeIdOk() (int32, bool) {
	if o == nil || o.TapeId == nil {
		var ret int32
		return ret, false
	}
	return *o.TapeId, true
}

// HasTapeId returns a boolean if a field has been set.
func (o *TaskSourceUp) HasTapeId() bool {
	if o != nil && o.TapeId != nil {
		return true
	}

	return false
}

// SetTapeId gets a reference to the given int32 and assigns it to the TapeId field.
func (o *TaskSourceUp) SetTapeId(v int32) {
	o.TapeId = &v
}

// GetTapeBarcode returns the TapeBarcode field value if set, zero value otherwise.
func (o *TaskSourceUp) GetTapeBarcode() string {
	if o == nil || o.TapeBarcode == nil {
		var ret string
		return ret
	}
	return *o.TapeBarcode
}

// GetTapeBarcodeOk returns a tuple with the TapeBarcode field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TaskSourceUp) GetTapeBarcodeOk() (string, bool) {
	if o == nil || o.TapeBarcode == nil {
		var ret string
		return ret, false
	}
	return *o.TapeBarcode, true
}

// HasTapeBarcode returns a boolean if a field has been set.
func (o *TaskSourceUp) HasTapeBarcode() bool {
	if o != nil && o.TapeBarcode != nil {
		return true
	}

	return false
}

// SetTapeBarcode gets a reference to the given string and assigns it to the TapeBarcode field.
func (o *TaskSourceUp) SetTapeBarcode(v string) {
	o.TapeBarcode = &v
}

// GetTapePoolId returns the TapePoolId field value if set, zero value otherwise.
func (o *TaskSourceUp) GetTapePoolId() int32 {
	if o == nil || o.TapePoolId == nil {
		var ret int32
		return ret
	}
	return *o.TapePoolId
}

// GetTapePoolIdOk returns a tuple with the TapePoolId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TaskSourceUp) GetTapePoolIdOk() (int32, bool) {
	if o == nil || o.TapePoolId == nil {
		var ret int32
		return ret, false
	}
	return *o.TapePoolId, true
}

// HasTapePoolId returns a boolean if a field has been set.
func (o *TaskSourceUp) HasTapePoolId() bool {
	if o != nil && o.TapePoolId != nil {
		return true
	}

	return false
}

// SetTapePoolId gets a reference to the given int32 and assigns it to the TapePoolId field.
func (o *TaskSourceUp) SetTapePoolId(v int32) {
	o.TapePoolId = &v
}

// GetTapePoolName returns the TapePoolName field value if set, zero value otherwise.
func (o *TaskSourceUp) GetTapePoolName() string {
	if o == nil || o.TapePoolName == nil {
		var ret string
		return ret
	}
	return *o.TapePoolName
}

// GetTapePoolNameOk returns a tuple with the TapePoolName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TaskSourceUp) GetTapePoolNameOk() (string, bool) {
	if o == nil || o.TapePoolName == nil {
		var ret string
		return ret, false
	}
	return *o.TapePoolName, true
}

// HasTapePoolName returns a boolean if a field has been set.
func (o *TaskSourceUp) HasTapePoolName() bool {
	if o != nil && o.TapePoolName != nil {
		return true
	}

	return false
}

// SetTapePoolName gets a reference to the given string and assigns it to the TapePoolName field.
func (o *TaskSourceUp) SetTapePoolName(v string) {
	o.TapePoolName = &v
}

// GetCloudPoolId returns the CloudPoolId field value if set, zero value otherwise.
func (o *TaskSourceUp) GetCloudPoolId() int32 {
	if o == nil || o.CloudPoolId == nil {
		var ret int32
		return ret
	}
	return *o.CloudPoolId
}

// GetCloudPoolIdOk returns a tuple with the CloudPoolId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TaskSourceUp) GetCloudPoolIdOk() (int32, bool) {
	if o == nil || o.CloudPoolId == nil {
		var ret int32
		return ret, false
	}
	return *o.CloudPoolId, true
}

// HasCloudPoolId returns a boolean if a field has been set.
func (o *TaskSourceUp) HasCloudPoolId() bool {
	if o != nil && o.CloudPoolId != nil {
		return true
	}

	return false
}

// SetCloudPoolId gets a reference to the given int32 and assigns it to the CloudPoolId field.
func (o *TaskSourceUp) SetCloudPoolId(v int32) {
	o.CloudPoolId = &v
}

// GetCloudPoolName returns the CloudPoolName field value if set, zero value otherwise.
func (o *TaskSourceUp) GetCloudPoolName() string {
	if o == nil || o.CloudPoolName == nil {
		var ret string
		return ret
	}
	return *o.CloudPoolName
}

// GetCloudPoolNameOk returns a tuple with the CloudPoolName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TaskSourceUp) GetCloudPoolNameOk() (string, bool) {
	if o == nil || o.CloudPoolName == nil {
		var ret string
		return ret, false
	}
	return *o.CloudPoolName, true
}

// HasCloudPoolName returns a boolean if a field has been set.
func (o *TaskSourceUp) HasCloudPoolName() bool {
	if o != nil && o.CloudPoolName != nil {
		return true
	}

	return false
}

// SetCloudPoolName gets a reference to the given string and assigns it to the CloudPoolName field.
func (o *TaskSourceUp) SetCloudPoolName(v string) {
	o.CloudPoolName = &v
}

// GetNasPoolId returns the NasPoolId field value if set, zero value otherwise.
func (o *TaskSourceUp) GetNasPoolId() int32 {
	if o == nil || o.NasPoolId == nil {
		var ret int32
		return ret
	}
	return *o.NasPoolId
}

// GetNasPoolIdOk returns a tuple with the NasPoolId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TaskSourceUp) GetNasPoolIdOk() (int32, bool) {
	if o == nil || o.NasPoolId == nil {
		var ret int32
		return ret, false
	}
	return *o.NasPoolId, true
}

// HasNasPoolId returns a boolean if a field has been set.
func (o *TaskSourceUp) HasNasPoolId() bool {
	if o != nil && o.NasPoolId != nil {
		return true
	}

	return false
}

// SetNasPoolId gets a reference to the given int32 and assigns it to the NasPoolId field.
func (o *TaskSourceUp) SetNasPoolId(v int32) {
	o.NasPoolId = &v
}

// GetNasPoolName returns the NasPoolName field value if set, zero value otherwise.
func (o *TaskSourceUp) GetNasPoolName() string {
	if o == nil || o.NasPoolName == nil {
		var ret string
		return ret
	}
	return *o.NasPoolName
}

// GetNasPoolNameOk returns a tuple with the NasPoolName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TaskSourceUp) GetNasPoolNameOk() (string, bool) {
	if o == nil || o.NasPoolName == nil {
		var ret string
		return ret, false
	}
	return *o.NasPoolName, true
}

// HasNasPoolName returns a boolean if a field has been set.
func (o *TaskSourceUp) HasNasPoolName() bool {
	if o != nil && o.NasPoolName != nil {
		return true
	}

	return false
}

// SetNasPoolName gets a reference to the given string and assigns it to the NasPoolName field.
func (o *TaskSourceUp) SetNasPoolName(v string) {
	o.NasPoolName = &v
}

type NullableTaskSourceUp struct {
	Value TaskSourceUp
	ExplicitNull bool
}

func (v NullableTaskSourceUp) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableTaskSourceUp) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

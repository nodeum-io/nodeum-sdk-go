/*
 * # Nodeum API Reference The Nodeum API makes it easy to tap into the digital data mesh that runs across your organisation. Make requests to our API endpoints and we’ll give you everything you need to interconnect your business workflows with your storage.  All production API requests are made to:  http://<nodeum hostname>/api/  The current version of the API is v1. Backwards incompatible changes will result in a version bump. Some of our API endpoints require OAuth 2.0 credentials. Please see the Authentication & Authorization guide to get started.  There is also a sandbox to use when developing and testing applications, with requests being made to:  https://sandbox-api.uber.com/<version>  ¶ REST The Uber API is a RESTful API. This means that the API is designed to allow you to get, create, update, & delete objects with the HTTP verbs GET, POST, PUT, PATCH, & DELETE.  ¶ JSON The Uber API speaks exclusively in JSON. This means that you should always set the Content-Type header to application/json to ensure that your requests are properly accepted and processed by the API.    # Filter parameters When browsing a list of items, multiple filter parameters may be applied. Some operators can be added to the value as a prefix:  - `=` value is equal. Default operator, may be omitted  - `!=` value is different  - `>` greater than  - `>=` greater than or equal  - `<` lower than  - `>=` lower than or equal  - `><` included in list, items should be separated by `|`  - `!><` not included in list, items should be separated by `|`  - `~` pattern matching, may include `%` (any characters) and `_` (one character)  - `!~` pattern not matching, may include `%` (any characters) and `_` (one character)
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nodeum

type Tape struct {
	Id int32 `json:"id,omitempty"`
	TapeLibraryId int32 `json:"tape_library_id,omitempty"`
	TapePoolId int32 `json:"tape_pool_id,omitempty"`
	Barcode string `json:"barcode,omitempty"`
	Location string `json:"location,omitempty"`
	Type_ string `json:"type,omitempty"`
	Locked bool `json:"locked,omitempty"`
	Scratch bool `json:"scratch,omitempty"`
	Cleaning bool `json:"cleaning,omitempty"`
	WriteProtect bool `json:"write_protect,omitempty"`
	Mounted bool `json:"mounted,omitempty"`
	Ejected bool `json:"ejected,omitempty"`
	Known bool `json:"known,omitempty"`
	MountCount int32 `json:"mount_count,omitempty"`
	DateIn string `json:"date_in,omitempty"`
	DateMove string `json:"date_move,omitempty"`
	Free int32 `json:"free,omitempty"`
	Max int32 `json:"max,omitempty"`
	LastSizeUpdate string `json:"last_size_update,omitempty"`
	LastMaintenance string `json:"last_maintenance,omitempty"`
	LastRepack string `json:"last_repack,omitempty"`
	RepackStatus bool `json:"repack_status,omitempty"`
	Hash string `json:"hash,omitempty"`
	ForceImportType bool `json:"force_import_type,omitempty"`
	NeedToCheck bool `json:"need_to_check,omitempty"`
}

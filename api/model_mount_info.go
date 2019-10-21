/*
 * # Nodeum API Reference The Nodeum API makes it easy to tap into the digital data mesh that runs across your organisation. Make requests to our API endpoints and we’ll give you everything you need to interconnect your business workflows with your storage.  All production API requests are made to:  http://<nodeum hostname>/api/  The current production version of the API is v1. Backwards incompatible changes will result in a version bump.  There is also a sandbox to use when developing and testing applications, with requests being made to:  https://sandbox-api.uber.com/<version>  ¶ REST The Uber API is a RESTful API. This means that the API is designed to allow you to get, create, update, & delete objects with the HTTP verbs GET, POST, PUT, PATCH, & DELETE.  ¶ JSON The Uber API speaks exclusively in JSON. This means that you should always set the Content-Type header to application/json to ensure that your requests are properly accepted and processed by the API.    # Filter parameters When browsing a list of items, multiple filter parameters may be applied. Some operators can be added to the value as a prefix:  - `=` value is equal. Default operator, may be omitted  - `!=` value is different  - `>` greater than  - `>=` greater than or equal  - `<` lower than  - `>=` lower than or equal  - `><` included in list, items should be separated by `|`  - `!><` not included in list, items should be separated by `|`  - `~` pattern matching, may include `%` (any characters) and `_` (one character)  - `!~` pattern not matching, may include `%` (any characters) and `_` (one character)
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nodeum

type MountInfo struct {
	Id int32 `json:"id,omitempty"`
	MajorMinor string `json:"major_minor,omitempty"`
	Target string `json:"target,omitempty"`
	Source string `json:"source,omitempty"`
	Options string `json:"options,omitempty"`
	Type_ string `json:"type,omitempty"`
	NasId int32 `json:"nas_id,omitempty"`
	NasName string `json:"nas_name,omitempty"`
	NasShareId int32 `json:"nas_share_id,omitempty"`
	NasPoolId int32 `json:"nas_pool_id,omitempty"`
	NasPoolName string `json:"nas_pool_name,omitempty"`
	CloudConnectorId int32 `json:"cloud_connector_id,omitempty"`
	CloudConnectorName string `json:"cloud_connector_name,omitempty"`
	CloudBucketId int32 `json:"cloud_bucket_id,omitempty"`
	CloudBucketName string `json:"cloud_bucket_name,omitempty"`
	CloudPoolId int32 `json:"cloud_pool_id,omitempty"`
	CloudPoolName string `json:"cloud_pool_name,omitempty"`
	MountPointId int32 `json:"mount_point_id,omitempty"`
	MountPointName string `json:"mount_point_name,omitempty"`
	TapeId int32 `json:"tape_id,omitempty"`
	TapeBarcode string `json:"tape_barcode,omitempty"`
	TapeLibraryId int32 `json:"tape_library_id,omitempty"`
	TapeLibraryName string `json:"tape_library_name,omitempty"`
	TapeLibrarySerial string `json:"tape_library_serial,omitempty"`
}

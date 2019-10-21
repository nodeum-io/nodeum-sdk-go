/*
 * # Nodeum API Reference The Nodeum API makes it easy to tap into the digital data mesh that runs across your organisation. Make requests to our API endpoints and we’ll give you everything you need to interconnect your business workflows with your storage.  All production API requests are made to:  http://<nodeum hostname>/api/  The current version of the API is v1.2. Backwards incompatible changes will result in a version bump. Some of our API endpoints require OAuth 2.0 credentials. Please see the Authentication & Authorization guide to get started.  There is also a sandbox to use when developing and testing applications, with requests being made to:  https://sandbox-api.uber.com/<version>  ¶ REST The Uber API is a RESTful API. This means that the API is designed to allow you to get, create, update, & delete objects with the HTTP verbs GET, POST, PUT, PATCH, & DELETE.  ¶ JSON The Uber API speaks exclusively in JSON. This means that you should always set the Content-Type header to application/json to ensure that your requests are properly accepted and processed by the API.    # Filter parameters When browsing a list of items, multiple filter parameters may be applied. Some operators can be added to the value as a prefix:  - `=` value is equal. Default operator, may be omitted  - `!=` value is different  - `>` greater than  - `>=` greater than or equal  - `<` lower than  - `>=` lower than or equal  - `><` included in list, items should be separated by `|`  - `!><` not included in list, items should be separated by `|`  - `~` pattern matching, may include `%` (any characters) and `_` (one character)  - `!~` pattern not matching, may include `%` (any characters) and `_` (one character)
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nodeum

type MountPointCollection struct {
	Count int32 `json:"count,omitempty"`
	MountPoints []MountPoint `json:"mount_points,omitempty"`
}

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

// CloudPoolCollection struct for CloudPoolCollection
type CloudPoolCollection struct {
	Count *int32 `json:"count,omitempty"`
	CloudPools *[]CloudPool `json:"cloud_pools,omitempty"`
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *CloudPoolCollection) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CloudPoolCollection) GetCountOk() (int32, bool) {
	if o == nil || o.Count == nil {
		var ret int32
		return ret, false
	}
	return *o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *CloudPoolCollection) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *CloudPoolCollection) SetCount(v int32) {
	o.Count = &v
}

// GetCloudPools returns the CloudPools field value if set, zero value otherwise.
func (o *CloudPoolCollection) GetCloudPools() []CloudPool {
	if o == nil || o.CloudPools == nil {
		var ret []CloudPool
		return ret
	}
	return *o.CloudPools
}

// GetCloudPoolsOk returns a tuple with the CloudPools field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CloudPoolCollection) GetCloudPoolsOk() ([]CloudPool, bool) {
	if o == nil || o.CloudPools == nil {
		var ret []CloudPool
		return ret, false
	}
	return *o.CloudPools, true
}

// HasCloudPools returns a boolean if a field has been set.
func (o *CloudPoolCollection) HasCloudPools() bool {
	if o != nil && o.CloudPools != nil {
		return true
	}

	return false
}

// SetCloudPools gets a reference to the given []CloudPool and assigns it to the CloudPools field.
func (o *CloudPoolCollection) SetCloudPools(v []CloudPool) {
	o.CloudPools = &v
}

type NullableCloudPoolCollection struct {
	Value CloudPoolCollection
	ExplicitNull bool
}

func (v NullableCloudPoolCollection) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCloudPoolCollection) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

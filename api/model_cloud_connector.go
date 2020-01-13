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

// CloudConnector struct for CloudConnector
type CloudConnector struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Url *string `json:"url,omitempty"`
	UrlProxy *string `json:"url_proxy,omitempty"`
	Provider *string `json:"provider,omitempty"`
	Region *string `json:"region,omitempty"`
	AccessKey *string `json:"access_key,omitempty"`
	SecretKey *string `json:"secret_key,omitempty"`
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CloudConnector) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CloudConnector) GetIdOk() (int32, bool) {
	if o == nil || o.Id == nil {
		var ret int32
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CloudConnector) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CloudConnector) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudConnector) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CloudConnector) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudConnector) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudConnector) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CloudConnector) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CloudConnector) GetUrlOk() (string, bool) {
	if o == nil || o.Url == nil {
		var ret string
		return ret, false
	}
	return *o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CloudConnector) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CloudConnector) SetUrl(v string) {
	o.Url = &v
}

// GetUrlProxy returns the UrlProxy field value if set, zero value otherwise.
func (o *CloudConnector) GetUrlProxy() string {
	if o == nil || o.UrlProxy == nil {
		var ret string
		return ret
	}
	return *o.UrlProxy
}

// GetUrlProxyOk returns a tuple with the UrlProxy field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CloudConnector) GetUrlProxyOk() (string, bool) {
	if o == nil || o.UrlProxy == nil {
		var ret string
		return ret, false
	}
	return *o.UrlProxy, true
}

// HasUrlProxy returns a boolean if a field has been set.
func (o *CloudConnector) HasUrlProxy() bool {
	if o != nil && o.UrlProxy != nil {
		return true
	}

	return false
}

// SetUrlProxy gets a reference to the given string and assigns it to the UrlProxy field.
func (o *CloudConnector) SetUrlProxy(v string) {
	o.UrlProxy = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *CloudConnector) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CloudConnector) GetProviderOk() (string, bool) {
	if o == nil || o.Provider == nil {
		var ret string
		return ret, false
	}
	return *o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *CloudConnector) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *CloudConnector) SetProvider(v string) {
	o.Provider = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *CloudConnector) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CloudConnector) GetRegionOk() (string, bool) {
	if o == nil || o.Region == nil {
		var ret string
		return ret, false
	}
	return *o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *CloudConnector) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *CloudConnector) SetRegion(v string) {
	o.Region = &v
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *CloudConnector) GetAccessKey() string {
	if o == nil || o.AccessKey == nil {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CloudConnector) GetAccessKeyOk() (string, bool) {
	if o == nil || o.AccessKey == nil {
		var ret string
		return ret, false
	}
	return *o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *CloudConnector) HasAccessKey() bool {
	if o != nil && o.AccessKey != nil {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *CloudConnector) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *CloudConnector) GetSecretKey() string {
	if o == nil || o.SecretKey == nil {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CloudConnector) GetSecretKeyOk() (string, bool) {
	if o == nil || o.SecretKey == nil {
		var ret string
		return ret, false
	}
	return *o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *CloudConnector) HasSecretKey() bool {
	if o != nil && o.SecretKey != nil {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *CloudConnector) SetSecretKey(v string) {
	o.SecretKey = &v
}

type NullableCloudConnector struct {
	Value CloudConnector
	ExplicitNull bool
}

func (v NullableCloudConnector) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCloudConnector) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

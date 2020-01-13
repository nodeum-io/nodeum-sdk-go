# ApiKeyFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Key** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**ApiKeyScopes** | Pointer to [**[]ApiKeyScope**](api_key_scope.md) |  | [optional] 

## Methods

### GetId

`func (o *ApiKeyFull) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiKeyFull) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *ApiKeyFull) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *ApiKeyFull) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetKey

`func (o *ApiKeyFull) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ApiKeyFull) GetKeyOk() (string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKey

`func (o *ApiKeyFull) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKey

`func (o *ApiKeyFull) SetKey(v string)`

SetKey gets a reference to the given string and assigns it to the Key field.

### GetName

`func (o *ApiKeyFull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiKeyFull) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *ApiKeyFull) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *ApiKeyFull) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetApiKeyScopes

`func (o *ApiKeyFull) GetApiKeyScopes() []ApiKeyScope`

GetApiKeyScopes returns the ApiKeyScopes field if non-nil, zero value otherwise.

### GetApiKeyScopesOk

`func (o *ApiKeyFull) GetApiKeyScopesOk() ([]ApiKeyScope, bool)`

GetApiKeyScopesOk returns a tuple with the ApiKeyScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApiKeyScopes

`func (o *ApiKeyFull) HasApiKeyScopes() bool`

HasApiKeyScopes returns a boolean if a field has been set.

### SetApiKeyScopes

`func (o *ApiKeyFull) SetApiKeyScopes(v []ApiKeyScope)`

SetApiKeyScopes gets a reference to the given []ApiKeyScope and assigns it to the ApiKeyScopes field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



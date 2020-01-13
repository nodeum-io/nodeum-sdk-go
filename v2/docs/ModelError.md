# ModelError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**map[string][]AttributeError**](array.md) | Parsable objects describing the errors. The key is the invalid attribute name. | [optional] 
**Messages** | Pointer to **[]string** | English description of the errors. | [optional] 

## Methods

### GetDetails

`func (o *ModelError) GetDetails() map[string][]AttributeError`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ModelError) GetDetailsOk() (map[string][]AttributeError, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDetails

`func (o *ModelError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetails

`func (o *ModelError) SetDetails(v map[string][]AttributeError)`

SetDetails gets a reference to the given map[string][]AttributeError and assigns it to the Details field.

### GetMessages

`func (o *ModelError) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ModelError) GetMessagesOk() ([]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessages

`func (o *ModelError) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### SetMessages

`func (o *ModelError) SetMessages(v []string)`

SetMessages gets a reference to the given []string and assigns it to the Messages field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



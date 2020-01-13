# TapeStatCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] [readonly] 
**TapeStats** | Pointer to [**[]TapeStat**](tape_stat.md) |  | [optional] [readonly] 

## Methods

### GetCount

`func (o *TapeStatCollection) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TapeStatCollection) GetCountOk() (int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCount

`func (o *TapeStatCollection) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCount

`func (o *TapeStatCollection) SetCount(v int32)`

SetCount gets a reference to the given int32 and assigns it to the Count field.

### GetTapeStats

`func (o *TapeStatCollection) GetTapeStats() []TapeStat`

GetTapeStats returns the TapeStats field if non-nil, zero value otherwise.

### GetTapeStatsOk

`func (o *TapeStatCollection) GetTapeStatsOk() ([]TapeStat, bool)`

GetTapeStatsOk returns a tuple with the TapeStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTapeStats

`func (o *TapeStatCollection) HasTapeStats() bool`

HasTapeStats returns a boolean if a field has been set.

### SetTapeStats

`func (o *TapeStatCollection) SetTapeStats(v []TapeStat)`

SetTapeStats gets a reference to the given []TapeStat and assigns it to the TapeStats field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# TaskSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**TaskId** | Pointer to **int32** |  | [optional] 
**Rrule** | Pointer to **string** |  | [optional] 
**Done** | Pointer to **bool** |  | [optional] 
**Next** | Pointer to **string** |  | [optional] [readonly] 
**MissedCount** | Pointer to **int32** |  | [optional] [readonly] 
**SkippedCount** | Pointer to **int32** |  | [optional] [readonly] 
**MissedLast** | Pointer to **string** |  | [optional] [readonly] 
**MissedFirst** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### GetId

`func (o *TaskSchedule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskSchedule) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *TaskSchedule) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *TaskSchedule) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetTaskId

`func (o *TaskSchedule) GetTaskId() int32`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *TaskSchedule) GetTaskIdOk() (int32, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTaskId

`func (o *TaskSchedule) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskId

`func (o *TaskSchedule) SetTaskId(v int32)`

SetTaskId gets a reference to the given int32 and assigns it to the TaskId field.

### GetRrule

`func (o *TaskSchedule) GetRrule() string`

GetRrule returns the Rrule field if non-nil, zero value otherwise.

### GetRruleOk

`func (o *TaskSchedule) GetRruleOk() (string, bool)`

GetRruleOk returns a tuple with the Rrule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRrule

`func (o *TaskSchedule) HasRrule() bool`

HasRrule returns a boolean if a field has been set.

### SetRrule

`func (o *TaskSchedule) SetRrule(v string)`

SetRrule gets a reference to the given string and assigns it to the Rrule field.

### GetDone

`func (o *TaskSchedule) GetDone() bool`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *TaskSchedule) GetDoneOk() (bool, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDone

`func (o *TaskSchedule) HasDone() bool`

HasDone returns a boolean if a field has been set.

### SetDone

`func (o *TaskSchedule) SetDone(v bool)`

SetDone gets a reference to the given bool and assigns it to the Done field.

### GetNext

`func (o *TaskSchedule) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *TaskSchedule) GetNextOk() (string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNext

`func (o *TaskSchedule) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNext

`func (o *TaskSchedule) SetNext(v string)`

SetNext gets a reference to the given string and assigns it to the Next field.

### GetMissedCount

`func (o *TaskSchedule) GetMissedCount() int32`

GetMissedCount returns the MissedCount field if non-nil, zero value otherwise.

### GetMissedCountOk

`func (o *TaskSchedule) GetMissedCountOk() (int32, bool)`

GetMissedCountOk returns a tuple with the MissedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMissedCount

`func (o *TaskSchedule) HasMissedCount() bool`

HasMissedCount returns a boolean if a field has been set.

### SetMissedCount

`func (o *TaskSchedule) SetMissedCount(v int32)`

SetMissedCount gets a reference to the given int32 and assigns it to the MissedCount field.

### GetSkippedCount

`func (o *TaskSchedule) GetSkippedCount() int32`

GetSkippedCount returns the SkippedCount field if non-nil, zero value otherwise.

### GetSkippedCountOk

`func (o *TaskSchedule) GetSkippedCountOk() (int32, bool)`

GetSkippedCountOk returns a tuple with the SkippedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSkippedCount

`func (o *TaskSchedule) HasSkippedCount() bool`

HasSkippedCount returns a boolean if a field has been set.

### SetSkippedCount

`func (o *TaskSchedule) SetSkippedCount(v int32)`

SetSkippedCount gets a reference to the given int32 and assigns it to the SkippedCount field.

### GetMissedLast

`func (o *TaskSchedule) GetMissedLast() string`

GetMissedLast returns the MissedLast field if non-nil, zero value otherwise.

### GetMissedLastOk

`func (o *TaskSchedule) GetMissedLastOk() (string, bool)`

GetMissedLastOk returns a tuple with the MissedLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMissedLast

`func (o *TaskSchedule) HasMissedLast() bool`

HasMissedLast returns a boolean if a field has been set.

### SetMissedLast

`func (o *TaskSchedule) SetMissedLast(v string)`

SetMissedLast gets a reference to the given string and assigns it to the MissedLast field.

### GetMissedFirst

`func (o *TaskSchedule) GetMissedFirst() string`

GetMissedFirst returns the MissedFirst field if non-nil, zero value otherwise.

### GetMissedFirstOk

`func (o *TaskSchedule) GetMissedFirstOk() (string, bool)`

GetMissedFirstOk returns a tuple with the MissedFirst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMissedFirst

`func (o *TaskSchedule) HasMissedFirst() bool`

HasMissedFirst returns a boolean if a field has been set.

### SetMissedFirst

`func (o *TaskSchedule) SetMissedFirst(v string)`

SetMissedFirst gets a reference to the given string and assigns it to the MissedFirst field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



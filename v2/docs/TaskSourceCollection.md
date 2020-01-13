# TaskSourceCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] [readonly] 
**TaskSources** | Pointer to [**[]TaskSourceDown**](task_source_down.md) |  | [optional] [readonly] 

## Methods

### GetCount

`func (o *TaskSourceCollection) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TaskSourceCollection) GetCountOk() (int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCount

`func (o *TaskSourceCollection) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCount

`func (o *TaskSourceCollection) SetCount(v int32)`

SetCount gets a reference to the given int32 and assigns it to the Count field.

### GetTaskSources

`func (o *TaskSourceCollection) GetTaskSources() []TaskSourceDown`

GetTaskSources returns the TaskSources field if non-nil, zero value otherwise.

### GetTaskSourcesOk

`func (o *TaskSourceCollection) GetTaskSourcesOk() ([]TaskSourceDown, bool)`

GetTaskSourcesOk returns a tuple with the TaskSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTaskSources

`func (o *TaskSourceCollection) HasTaskSources() bool`

HasTaskSources returns a boolean if a field has been set.

### SetTaskSources

`func (o *TaskSourceCollection) SetTaskSources(v []TaskSourceDown)`

SetTaskSources gets a reference to the given []TaskSourceDown and assigns it to the TaskSources field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



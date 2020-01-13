# TaskCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] [readonly] 
**Tasks** | Pointer to [**[]Task**](task.md) |  | [optional] [readonly] 

## Methods

### GetCount

`func (o *TaskCollection) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TaskCollection) GetCountOk() (int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCount

`func (o *TaskCollection) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCount

`func (o *TaskCollection) SetCount(v int32)`

SetCount gets a reference to the given int32 and assigns it to the Count field.

### GetTasks

`func (o *TaskCollection) GetTasks() []Task`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *TaskCollection) GetTasksOk() ([]Task, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTasks

`func (o *TaskCollection) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### SetTasks

`func (o *TaskCollection) SetTasks(v []Task)`

SetTasks gets a reference to the given []Task and assigns it to the Tasks field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



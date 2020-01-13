# TaskExecutionCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] [readonly] 
**TaskExecutions** | Pointer to [**[]TaskExecution**](task_execution.md) |  | [optional] [readonly] 

## Methods

### GetCount

`func (o *TaskExecutionCollection) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TaskExecutionCollection) GetCountOk() (int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCount

`func (o *TaskExecutionCollection) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCount

`func (o *TaskExecutionCollection) SetCount(v int32)`

SetCount gets a reference to the given int32 and assigns it to the Count field.

### GetTaskExecutions

`func (o *TaskExecutionCollection) GetTaskExecutions() []TaskExecution`

GetTaskExecutions returns the TaskExecutions field if non-nil, zero value otherwise.

### GetTaskExecutionsOk

`func (o *TaskExecutionCollection) GetTaskExecutionsOk() ([]TaskExecution, bool)`

GetTaskExecutionsOk returns a tuple with the TaskExecutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTaskExecutions

`func (o *TaskExecutionCollection) HasTaskExecutions() bool`

HasTaskExecutions returns a boolean if a field has been set.

### SetTaskExecutions

`func (o *TaskExecutionCollection) SetTaskExecutions(v []TaskExecution)`

SetTaskExecutions gets a reference to the given []TaskExecution and assigns it to the TaskExecutions field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



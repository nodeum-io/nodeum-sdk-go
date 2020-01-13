# TaskExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**TaskId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**LogTime** | Pointer to **string** |  | [optional] 
**JobStarted** | Pointer to **string** |  | [optional] 
**JobFinished** | Pointer to **string** |  | [optional] 
**ToProcessSize** | Pointer to **int32** |  | [optional] 
**ProcessedSize** | Pointer to **int32** |  | [optional] 
**ToProcessFiles** | Pointer to **int32** |  | [optional] 
**ProcessedFiles** | Pointer to **int32** |  | [optional] 
**FinalizedFiles** | Pointer to **int32** |  | [optional] 
**EstimationTime** | Pointer to **int32** |  | [optional] 
**Bandwidth** | Pointer to **int32** |  | [optional] 

## Methods

### GetId

`func (o *TaskExecution) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskExecution) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *TaskExecution) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *TaskExecution) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetTaskId

`func (o *TaskExecution) GetTaskId() int32`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *TaskExecution) GetTaskIdOk() (int32, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTaskId

`func (o *TaskExecution) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskId

`func (o *TaskExecution) SetTaskId(v int32)`

SetTaskId gets a reference to the given int32 and assigns it to the TaskId field.

### GetName

`func (o *TaskExecution) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskExecution) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *TaskExecution) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *TaskExecution) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetType

`func (o *TaskExecution) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaskExecution) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *TaskExecution) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *TaskExecution) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetStatus

`func (o *TaskExecution) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskExecution) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *TaskExecution) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *TaskExecution) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.

### GetLogTime

`func (o *TaskExecution) GetLogTime() string`

GetLogTime returns the LogTime field if non-nil, zero value otherwise.

### GetLogTimeOk

`func (o *TaskExecution) GetLogTimeOk() (string, bool)`

GetLogTimeOk returns a tuple with the LogTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLogTime

`func (o *TaskExecution) HasLogTime() bool`

HasLogTime returns a boolean if a field has been set.

### SetLogTime

`func (o *TaskExecution) SetLogTime(v string)`

SetLogTime gets a reference to the given string and assigns it to the LogTime field.

### GetJobStarted

`func (o *TaskExecution) GetJobStarted() string`

GetJobStarted returns the JobStarted field if non-nil, zero value otherwise.

### GetJobStartedOk

`func (o *TaskExecution) GetJobStartedOk() (string, bool)`

GetJobStartedOk returns a tuple with the JobStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasJobStarted

`func (o *TaskExecution) HasJobStarted() bool`

HasJobStarted returns a boolean if a field has been set.

### SetJobStarted

`func (o *TaskExecution) SetJobStarted(v string)`

SetJobStarted gets a reference to the given string and assigns it to the JobStarted field.

### GetJobFinished

`func (o *TaskExecution) GetJobFinished() string`

GetJobFinished returns the JobFinished field if non-nil, zero value otherwise.

### GetJobFinishedOk

`func (o *TaskExecution) GetJobFinishedOk() (string, bool)`

GetJobFinishedOk returns a tuple with the JobFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasJobFinished

`func (o *TaskExecution) HasJobFinished() bool`

HasJobFinished returns a boolean if a field has been set.

### SetJobFinished

`func (o *TaskExecution) SetJobFinished(v string)`

SetJobFinished gets a reference to the given string and assigns it to the JobFinished field.

### GetToProcessSize

`func (o *TaskExecution) GetToProcessSize() int32`

GetToProcessSize returns the ToProcessSize field if non-nil, zero value otherwise.

### GetToProcessSizeOk

`func (o *TaskExecution) GetToProcessSizeOk() (int32, bool)`

GetToProcessSizeOk returns a tuple with the ToProcessSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasToProcessSize

`func (o *TaskExecution) HasToProcessSize() bool`

HasToProcessSize returns a boolean if a field has been set.

### SetToProcessSize

`func (o *TaskExecution) SetToProcessSize(v int32)`

SetToProcessSize gets a reference to the given int32 and assigns it to the ToProcessSize field.

### GetProcessedSize

`func (o *TaskExecution) GetProcessedSize() int32`

GetProcessedSize returns the ProcessedSize field if non-nil, zero value otherwise.

### GetProcessedSizeOk

`func (o *TaskExecution) GetProcessedSizeOk() (int32, bool)`

GetProcessedSizeOk returns a tuple with the ProcessedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProcessedSize

`func (o *TaskExecution) HasProcessedSize() bool`

HasProcessedSize returns a boolean if a field has been set.

### SetProcessedSize

`func (o *TaskExecution) SetProcessedSize(v int32)`

SetProcessedSize gets a reference to the given int32 and assigns it to the ProcessedSize field.

### GetToProcessFiles

`func (o *TaskExecution) GetToProcessFiles() int32`

GetToProcessFiles returns the ToProcessFiles field if non-nil, zero value otherwise.

### GetToProcessFilesOk

`func (o *TaskExecution) GetToProcessFilesOk() (int32, bool)`

GetToProcessFilesOk returns a tuple with the ToProcessFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasToProcessFiles

`func (o *TaskExecution) HasToProcessFiles() bool`

HasToProcessFiles returns a boolean if a field has been set.

### SetToProcessFiles

`func (o *TaskExecution) SetToProcessFiles(v int32)`

SetToProcessFiles gets a reference to the given int32 and assigns it to the ToProcessFiles field.

### GetProcessedFiles

`func (o *TaskExecution) GetProcessedFiles() int32`

GetProcessedFiles returns the ProcessedFiles field if non-nil, zero value otherwise.

### GetProcessedFilesOk

`func (o *TaskExecution) GetProcessedFilesOk() (int32, bool)`

GetProcessedFilesOk returns a tuple with the ProcessedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProcessedFiles

`func (o *TaskExecution) HasProcessedFiles() bool`

HasProcessedFiles returns a boolean if a field has been set.

### SetProcessedFiles

`func (o *TaskExecution) SetProcessedFiles(v int32)`

SetProcessedFiles gets a reference to the given int32 and assigns it to the ProcessedFiles field.

### GetFinalizedFiles

`func (o *TaskExecution) GetFinalizedFiles() int32`

GetFinalizedFiles returns the FinalizedFiles field if non-nil, zero value otherwise.

### GetFinalizedFilesOk

`func (o *TaskExecution) GetFinalizedFilesOk() (int32, bool)`

GetFinalizedFilesOk returns a tuple with the FinalizedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFinalizedFiles

`func (o *TaskExecution) HasFinalizedFiles() bool`

HasFinalizedFiles returns a boolean if a field has been set.

### SetFinalizedFiles

`func (o *TaskExecution) SetFinalizedFiles(v int32)`

SetFinalizedFiles gets a reference to the given int32 and assigns it to the FinalizedFiles field.

### GetEstimationTime

`func (o *TaskExecution) GetEstimationTime() int32`

GetEstimationTime returns the EstimationTime field if non-nil, zero value otherwise.

### GetEstimationTimeOk

`func (o *TaskExecution) GetEstimationTimeOk() (int32, bool)`

GetEstimationTimeOk returns a tuple with the EstimationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEstimationTime

`func (o *TaskExecution) HasEstimationTime() bool`

HasEstimationTime returns a boolean if a field has been set.

### SetEstimationTime

`func (o *TaskExecution) SetEstimationTime(v int32)`

SetEstimationTime gets a reference to the given int32 and assigns it to the EstimationTime field.

### GetBandwidth

`func (o *TaskExecution) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *TaskExecution) GetBandwidthOk() (int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBandwidth

`func (o *TaskExecution) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### SetBandwidth

`func (o *TaskExecution) SetBandwidth(v int32)`

SetBandwidth gets a reference to the given int32 and assigns it to the Bandwidth field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



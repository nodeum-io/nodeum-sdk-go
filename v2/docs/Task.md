# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**WorkflowType** | Pointer to **string** |  | [optional] 
**WorkflowAction** | Pointer to **string** |  | [optional] 
**SourceType** | Pointer to **string** |  | [optional] 
**DestinationType** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**ConflictResolution** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] [readonly] 
**Activate** | Pointer to **bool** |  | [optional] 
**CreationDate** | Pointer to **string** |  | [optional] [readonly] 
**ModificationDate** | Pointer to **string** |  | [optional] [readonly] 
**CreationUsername** | Pointer to **string** |  | [optional] [readonly] 
**ModificationUsername** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to **string** |  | [optional] [readonly] 
**JobStarted** | Pointer to **string** |  | [optional] [readonly] 
**JobFinished** | Pointer to **string** |  | [optional] [readonly] 
**ProcessedSize** | Pointer to **int32** |  | [optional] [readonly] 
**ToProcessSize** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### GetId

`func (o *Task) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Task) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Task) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Task) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetName

`func (o *Task) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Task) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Task) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Task) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetComment

`func (o *Task) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Task) GetCommentOk() (string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComment

`func (o *Task) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetComment

`func (o *Task) SetComment(v string)`

SetComment gets a reference to the given string and assigns it to the Comment field.

### GetWorkflowType

`func (o *Task) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *Task) GetWorkflowTypeOk() (string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkflowType

`func (o *Task) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.

### SetWorkflowType

`func (o *Task) SetWorkflowType(v string)`

SetWorkflowType gets a reference to the given string and assigns it to the WorkflowType field.

### GetWorkflowAction

`func (o *Task) GetWorkflowAction() string`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *Task) GetWorkflowActionOk() (string, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkflowAction

`func (o *Task) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.

### SetWorkflowAction

`func (o *Task) SetWorkflowAction(v string)`

SetWorkflowAction gets a reference to the given string and assigns it to the WorkflowAction field.

### GetSourceType

`func (o *Task) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *Task) GetSourceTypeOk() (string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSourceType

`func (o *Task) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### SetSourceType

`func (o *Task) SetSourceType(v string)`

SetSourceType gets a reference to the given string and assigns it to the SourceType field.

### GetDestinationType

`func (o *Task) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *Task) GetDestinationTypeOk() (string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDestinationType

`func (o *Task) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### SetDestinationType

`func (o *Task) SetDestinationType(v string)`

SetDestinationType gets a reference to the given string and assigns it to the DestinationType field.

### GetPriority

`func (o *Task) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Task) GetPriorityOk() (int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPriority

`func (o *Task) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriority

`func (o *Task) SetPriority(v int32)`

SetPriority gets a reference to the given int32 and assigns it to the Priority field.

### GetConflictResolution

`func (o *Task) GetConflictResolution() string`

GetConflictResolution returns the ConflictResolution field if non-nil, zero value otherwise.

### GetConflictResolutionOk

`func (o *Task) GetConflictResolutionOk() (string, bool)`

GetConflictResolutionOk returns a tuple with the ConflictResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConflictResolution

`func (o *Task) HasConflictResolution() bool`

HasConflictResolution returns a boolean if a field has been set.

### SetConflictResolution

`func (o *Task) SetConflictResolution(v string)`

SetConflictResolution gets a reference to the given string and assigns it to the ConflictResolution field.

### GetAction

`func (o *Task) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Task) GetActionOk() (string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAction

`func (o *Task) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetAction

`func (o *Task) SetAction(v string)`

SetAction gets a reference to the given string and assigns it to the Action field.

### GetActivate

`func (o *Task) GetActivate() bool`

GetActivate returns the Activate field if non-nil, zero value otherwise.

### GetActivateOk

`func (o *Task) GetActivateOk() (bool, bool)`

GetActivateOk returns a tuple with the Activate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivate

`func (o *Task) HasActivate() bool`

HasActivate returns a boolean if a field has been set.

### SetActivate

`func (o *Task) SetActivate(v bool)`

SetActivate gets a reference to the given bool and assigns it to the Activate field.

### GetCreationDate

`func (o *Task) GetCreationDate() string`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Task) GetCreationDateOk() (string, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreationDate

`func (o *Task) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### SetCreationDate

`func (o *Task) SetCreationDate(v string)`

SetCreationDate gets a reference to the given string and assigns it to the CreationDate field.

### GetModificationDate

`func (o *Task) GetModificationDate() string`

GetModificationDate returns the ModificationDate field if non-nil, zero value otherwise.

### GetModificationDateOk

`func (o *Task) GetModificationDateOk() (string, bool)`

GetModificationDateOk returns a tuple with the ModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModificationDate

`func (o *Task) HasModificationDate() bool`

HasModificationDate returns a boolean if a field has been set.

### SetModificationDate

`func (o *Task) SetModificationDate(v string)`

SetModificationDate gets a reference to the given string and assigns it to the ModificationDate field.

### GetCreationUsername

`func (o *Task) GetCreationUsername() string`

GetCreationUsername returns the CreationUsername field if non-nil, zero value otherwise.

### GetCreationUsernameOk

`func (o *Task) GetCreationUsernameOk() (string, bool)`

GetCreationUsernameOk returns a tuple with the CreationUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreationUsername

`func (o *Task) HasCreationUsername() bool`

HasCreationUsername returns a boolean if a field has been set.

### SetCreationUsername

`func (o *Task) SetCreationUsername(v string)`

SetCreationUsername gets a reference to the given string and assigns it to the CreationUsername field.

### GetModificationUsername

`func (o *Task) GetModificationUsername() string`

GetModificationUsername returns the ModificationUsername field if non-nil, zero value otherwise.

### GetModificationUsernameOk

`func (o *Task) GetModificationUsernameOk() (string, bool)`

GetModificationUsernameOk returns a tuple with the ModificationUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModificationUsername

`func (o *Task) HasModificationUsername() bool`

HasModificationUsername returns a boolean if a field has been set.

### SetModificationUsername

`func (o *Task) SetModificationUsername(v string)`

SetModificationUsername gets a reference to the given string and assigns it to the ModificationUsername field.

### GetStatus

`func (o *Task) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Task) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *Task) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *Task) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.

### GetJobStarted

`func (o *Task) GetJobStarted() string`

GetJobStarted returns the JobStarted field if non-nil, zero value otherwise.

### GetJobStartedOk

`func (o *Task) GetJobStartedOk() (string, bool)`

GetJobStartedOk returns a tuple with the JobStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasJobStarted

`func (o *Task) HasJobStarted() bool`

HasJobStarted returns a boolean if a field has been set.

### SetJobStarted

`func (o *Task) SetJobStarted(v string)`

SetJobStarted gets a reference to the given string and assigns it to the JobStarted field.

### GetJobFinished

`func (o *Task) GetJobFinished() string`

GetJobFinished returns the JobFinished field if non-nil, zero value otherwise.

### GetJobFinishedOk

`func (o *Task) GetJobFinishedOk() (string, bool)`

GetJobFinishedOk returns a tuple with the JobFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasJobFinished

`func (o *Task) HasJobFinished() bool`

HasJobFinished returns a boolean if a field has been set.

### SetJobFinished

`func (o *Task) SetJobFinished(v string)`

SetJobFinished gets a reference to the given string and assigns it to the JobFinished field.

### GetProcessedSize

`func (o *Task) GetProcessedSize() int32`

GetProcessedSize returns the ProcessedSize field if non-nil, zero value otherwise.

### GetProcessedSizeOk

`func (o *Task) GetProcessedSizeOk() (int32, bool)`

GetProcessedSizeOk returns a tuple with the ProcessedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProcessedSize

`func (o *Task) HasProcessedSize() bool`

HasProcessedSize returns a boolean if a field has been set.

### SetProcessedSize

`func (o *Task) SetProcessedSize(v int32)`

SetProcessedSize gets a reference to the given int32 and assigns it to the ProcessedSize field.

### GetToProcessSize

`func (o *Task) GetToProcessSize() int32`

GetToProcessSize returns the ToProcessSize field if non-nil, zero value otherwise.

### GetToProcessSizeOk

`func (o *Task) GetToProcessSizeOk() (int32, bool)`

GetToProcessSizeOk returns a tuple with the ToProcessSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasToProcessSize

`func (o *Task) HasToProcessSize() bool`

HasToProcessSize returns a boolean if a field has been set.

### SetToProcessSize

`func (o *Task) SetToProcessSize(v int32)`

SetToProcessSize gets a reference to the given int32 and assigns it to the ToProcessSize field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



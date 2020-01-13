/*
 * Nodeum API
 *
 * # About  This document describes the Nodeum API version 2:  If you are looking for any information about the product itself, reach the product website https://www.nodeum.io. You can also contact us at this email address : info@nodeum.io  # Filter parameters When browsing a list of items, multiple filter parameters may be applied. Some operators can be added to the value as a prefix:  - `=` value is equal. Default operator, may be omitted  - `!=` value is different  - `>` greater than  - `>=` greater than or equal  - `<` lower than  - `>=` lower than or equal  - `><` included in list, items should be separated by `|`  - `!><` not included in list, items should be separated by `|`  - `~` pattern matching, may include `%` (any characters) and `_` (one character)  - `!~` pattern not matching, may include `%` (any characters) and `_` (one character)  
 *
 * API version: 2.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nodeum

import (
	"bytes"
	"encoding/json"
)

// TaskExecutionCollection struct for TaskExecutionCollection
type TaskExecutionCollection struct {
	Count *int32 `json:"count,omitempty"`
	TaskExecutions *[]TaskExecution `json:"task_executions,omitempty"`
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *TaskExecutionCollection) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TaskExecutionCollection) GetCountOk() (int32, bool) {
	if o == nil || o.Count == nil {
		var ret int32
		return ret, false
	}
	return *o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *TaskExecutionCollection) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *TaskExecutionCollection) SetCount(v int32) {
	o.Count = &v
}

// GetTaskExecutions returns the TaskExecutions field value if set, zero value otherwise.
func (o *TaskExecutionCollection) GetTaskExecutions() []TaskExecution {
	if o == nil || o.TaskExecutions == nil {
		var ret []TaskExecution
		return ret
	}
	return *o.TaskExecutions
}

// GetTaskExecutionsOk returns a tuple with the TaskExecutions field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TaskExecutionCollection) GetTaskExecutionsOk() ([]TaskExecution, bool) {
	if o == nil || o.TaskExecutions == nil {
		var ret []TaskExecution
		return ret, false
	}
	return *o.TaskExecutions, true
}

// HasTaskExecutions returns a boolean if a field has been set.
func (o *TaskExecutionCollection) HasTaskExecutions() bool {
	if o != nil && o.TaskExecutions != nil {
		return true
	}

	return false
}

// SetTaskExecutions gets a reference to the given []TaskExecution and assigns it to the TaskExecutions field.
func (o *TaskExecutionCollection) SetTaskExecutions(v []TaskExecution) {
	o.TaskExecutions = &v
}

type NullableTaskExecutionCollection struct {
	Value TaskExecutionCollection
	ExplicitNull bool
}

func (v NullableTaskExecutionCollection) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableTaskExecutionCollection) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

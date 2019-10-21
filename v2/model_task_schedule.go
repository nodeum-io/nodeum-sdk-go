/*
 * Nodeum API
 *
 * # About  This document describes the Nodeum API version 2:  If you are looking for any information about the product itself, reach the product website https://www.nodeum.io. You can also contact us at this email address : info@nodeum.io  # Filter parameters When browsing a list of items, multiple filter parameters may be applied. Some operators can be added to the value as a prefix:  - `=` value is equal. Default operator, may be omitted  - `!=` value is different  - `>` greater than  - `>=` greater than or equal  - `<` lower than  - `>=` lower than or equal  - `><` included in list, items should be separated by `|`  - `!><` not included in list, items should be separated by `|`  - `~` pattern matching, may include `%` (any characters) and `_` (one character)  - `!~` pattern not matching, may include `%` (any characters) and `_` (one character)  
 *
 * API version: 2.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nodeum

type TaskSchedule struct {
	Id int32 `json:"id,omitempty"`
	TaskId int32 `json:"task_id,omitempty"`
	Rrule string `json:"rrule,omitempty"`
	Done bool `json:"done,omitempty"`
	Next string `json:"next,omitempty"`
	MissedCount int32 `json:"missed_count,omitempty"`
	SkippedCount int32 `json:"skipped_count,omitempty"`
	MissedLast string `json:"missed_last,omitempty"`
	MissedFirst string `json:"missed_first,omitempty"`
}

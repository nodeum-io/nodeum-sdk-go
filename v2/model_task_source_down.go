/*
 * Nodeum API
 *
 * # About  This document describes the Nodeum API version 2:  If you are looking for any information about the product itself, reach the product website https://www.nodeum.io. You can also contact us at this email address : info@nodeum.io  # Filter parameters When browsing a list of items, multiple filter parameters may be applied. Some operators can be added to the value as a prefix:  - `=` value is equal. Default operator, may be omitted  - `!=` value is different  - `>` greater than  - `>=` greater than or equal  - `<` lower than  - `>=` lower than or equal  - `><` included in list, items should be separated by `|`  - `!><` not included in list, items should be separated by `|`  - `~` pattern matching, may include `%` (any characters) and `_` (one character)  - `!~` pattern not matching, may include `%` (any characters) and `_` (one character)  
 *
 * API version: 2.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nodeum

type TaskSourceDown struct {
	Id int32 `json:"id,omitempty"`
	File *NodeumFile `json:"file,omitempty"`
	Range_ []int32 `json:"range,omitempty"`
	Container *Container `json:"container,omitempty"`
	ImportFile *ImportFile `json:"import_file,omitempty"`
	Tape *Tape `json:"tape,omitempty"`
	Pool *Pool `json:"pool,omitempty"`
}

/*
 * Nodeum API
 *
 * # About  This document describes the Nodeum API version 2:  If you are looking for any information about the product itself, reach the product website https://www.nodeum.io. You can also contact us at this email address : info@nodeum.io  # Filter parameters When browsing a list of items, multiple filter parameters may be applied. Some operators can be added to the value as a prefix:  - `=` value is equal. Default operator, may be omitted  - `!=` value is different  - `>` greater than  - `>=` greater than or equal  - `<` lower than  - `>=` lower than or equal  - `><` included in list, items should be separated by `|`  - `!><` not included in list, items should be separated by `|`  - `~` pattern matching, may include `%` (any characters) and `_` (one character)  - `!~` pattern not matching, may include `%` (any characters) and `_` (one character)  
 *
 * API version: 2.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nodeum
// NodeumFileWithPath struct for NodeumFileWithPath
type NodeumFileWithPath struct {
	Id int32 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Parent int32 `json:"parent,omitempty"`
	PrimaryId int32 `json:"primary_id,omitempty"`
	Type string `json:"type,omitempty"`
	Permission int32 `json:"permission,omitempty"`
	Size int32 `json:"size,omitempty"`
	ChangeDate string `json:"change_date,omitempty"`
	ModificationDate string `json:"modification_date,omitempty"`
	AccessDate string `json:"access_date,omitempty"`
	Uid int32 `json:"uid,omitempty"`
	Gid int32 `json:"gid,omitempty"`
	FilePath string `json:"file_path,omitempty"`
}

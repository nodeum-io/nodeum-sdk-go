/*
 * Nodeum API
 *
 * Nodeum API  # Filter parameters When browsing a list of items, multiple filter parameters may be applied. Some operators can be added to the value as a prefix:  - `=` value is equal. Default operator, may be omitted  - `!=` value is different  - `>` greater than  - `>=` greater than or equal  - `<` lower than  - `>=` lower than or equal  - `><` included in list, items should be separated by `|`  - `!><` not included in list, items should be separated by `|`  - `~` pattern matching, may include `%` (any characters) and `_` (one character)  - `!~` pattern not matching, may include `%` (any characters) and `_` (one character)  
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nodeum

type TaskDestinationDown struct {
	Id int32 `json:"id,omitempty"`
	Folder *NodeumFile `json:"folder,omitempty"`
	TapeId int32 `json:"tape_id,omitempty"`
	TapePoolId int32 `json:"tape_pool_id,omitempty"`
	CloudPoolId int32 `json:"cloud_pool_id,omitempty"`
	NasPoolId int32 `json:"nas_pool_id,omitempty"`
}

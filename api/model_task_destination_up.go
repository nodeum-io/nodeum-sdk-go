/*
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nodeum

type TaskDestinationUp struct {
	FolderId int32 `json:"folder_id,omitempty"`
	FolderPath string `json:"folder_path,omitempty"`
	TapeId int32 `json:"tape_id,omitempty"`
	TapeBarcode string `json:"tape_barcode,omitempty"`
	TapePoolId int32 `json:"tape_pool_id,omitempty"`
	TapePoolName string `json:"tape_pool_name,omitempty"`
	CloudPoolId int32 `json:"cloud_pool_id,omitempty"`
	CloudPoolName string `json:"cloud_pool_name,omitempty"`
	NasPoolId int32 `json:"nas_pool_id,omitempty"`
	NasPoolName string `json:"nas_pool_name,omitempty"`
}

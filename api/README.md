# Go API client for nodeum

Nodeum API  # Filter parameters When browsing a list of items, multiple filter parameters may be applied. Some operators can be added to the value as a prefix:  - `=` value is equal. Default operator, may be omitted  - `!=` value is different  - `>` greater than  - `>=` greater than or equal  - `<` lower than  - `>=` lower than or equal  - `><` included in list, items should be separated by `|`  - `!><` not included in list, items should be separated by `|`  - `~` pattern matching, may include `%` (any characters) and `_` (one character)  - `!~` pattern not matching, may include `%` (any characters) and `_` (one character)  

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2.0.0
- Package version: 2.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./nodeum"
```

## Documentation for API Endpoints

All URIs are relative to *https://localhost/api/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CloudBucketsApi* | [**IndexCloudBuckets**](docs/CloudBucketsApi.md#indexcloudbuckets) | **Get** /cloud_buckets | Lists all cloud buckets.
*CloudBucketsApi* | [**IndexCloudBucketsByCloudConnector**](docs/CloudBucketsApi.md#indexcloudbucketsbycloudconnector) | **Get** /cloud_connectors/{cloud_connector_id}/cloud_buckets | Lists all cloud buckets.
*CloudBucketsApi* | [**IndexCloudBucketsByCloudPool**](docs/CloudBucketsApi.md#indexcloudbucketsbycloudpool) | **Get** /cloud_pools/{cloud_pool_id}/cloud_buckets | Lists all cloud buckets.
*CloudBucketsApi* | [**ShowCloudBucket**](docs/CloudBucketsApi.md#showcloudbucket) | **Get** /cloud_buckets/{cloud_bucket_id} | Displays a specific cloud bucket.
*CloudBucketsApi* | [**ShowCloudBucketByCloudConnector**](docs/CloudBucketsApi.md#showcloudbucketbycloudconnector) | **Get** /cloud_connectors/{cloud_connector_id}/cloud_buckets/{cloud_bucket_id} | Displays a specific cloud bucket.
*CloudBucketsApi* | [**ShowCloudBucketByCloudPool**](docs/CloudBucketsApi.md#showcloudbucketbycloudpool) | **Get** /cloud_pools/{cloud_pool_id}/cloud_buckets/{cloud_bucket_id} | Displays a specific cloud bucket.
*CloudBucketsApi* | [**SyncCloudBuckets**](docs/CloudBucketsApi.md#synccloudbuckets) | **Put** /cloud_connectors/{cloud_connector_id}/cloud_buckets/-/sync | Synchronize internal cloud buckets with their remote equivalent.
*CloudBucketsApi* | [**SyncResultCloudBuckets**](docs/CloudBucketsApi.md#syncresultcloudbuckets) | **Get** /cloud_connectors/{cloud_connector_id}/cloud_buckets/-/sync | Check result of cloud connector sync job.
*CloudConnectorsApi* | [**CreateCloudConnector**](docs/CloudConnectorsApi.md#createcloudconnector) | **Post** /cloud_connectors | Creates a new cloud connector.
*CloudConnectorsApi* | [**DestroyCloudConnector**](docs/CloudConnectorsApi.md#destroycloudconnector) | **Delete** /cloud_connectors/{cloud_connector_id} | Destroys a specific cloud connector.
*CloudConnectorsApi* | [**IndexCloudConnectors**](docs/CloudConnectorsApi.md#indexcloudconnectors) | **Get** /cloud_connectors | Lists all cloud connectors.
*CloudConnectorsApi* | [**ShowCloudConnector**](docs/CloudConnectorsApi.md#showcloudconnector) | **Get** /cloud_connectors/{cloud_connector_id} | Displays a specific cloud connector.
*CloudConnectorsApi* | [**TestCloudConnector**](docs/CloudConnectorsApi.md#testcloudconnector) | **Put** /cloud_connectors/-/test | Test an unsaved cloud connector.
*CloudConnectorsApi* | [**TestResultCloudConnector**](docs/CloudConnectorsApi.md#testresultcloudconnector) | **Get** /cloud_connectors/-/test | Check result of cloud connector test job.
*CloudConnectorsApi* | [**UpdateCloudConnector**](docs/CloudConnectorsApi.md#updatecloudconnector) | **Put** /cloud_connectors/{cloud_connector_id} | Updates a specific cloud connector.
*CloudPoolsApi* | [**CreateCloudPool**](docs/CloudPoolsApi.md#createcloudpool) | **Post** /cloud_pools | Creates a new cloud pool.
*CloudPoolsApi* | [**DestroyCloudPool**](docs/CloudPoolsApi.md#destroycloudpool) | **Delete** /cloud_pools/{cloud_pool_id} | Destroys a specific cloud pool.
*CloudPoolsApi* | [**IndexCloudPools**](docs/CloudPoolsApi.md#indexcloudpools) | **Get** /cloud_pools | Lists all cloud pools.
*CloudPoolsApi* | [**ShowCloudPool**](docs/CloudPoolsApi.md#showcloudpool) | **Get** /cloud_pools/{cloud_pool_id} | Displays a specific cloud pool.
*CloudPoolsApi* | [**UpdateCloudPool**](docs/CloudPoolsApi.md#updatecloudpool) | **Put** /cloud_pools/{cloud_pool_id} | Updates a specific cloud pool.
*ContainersApi* | [**CreateContainer**](docs/ContainersApi.md#createcontainer) | **Post** /containers | Creates a new container.
*ContainersApi* | [**CreateContainerPrivilege**](docs/ContainersApi.md#createcontainerprivilege) | **Post** /containers/{container_id}/container_privileges | Creates a new privilege on the container.
*ContainersApi* | [**DestroyContainer**](docs/ContainersApi.md#destroycontainer) | **Delete** /containers/{container_id} | Destroys a specific container.
*ContainersApi* | [**DestroyContainerPrivilege**](docs/ContainersApi.md#destroycontainerprivilege) | **Delete** /containers/{container_id}/container_privileges/{container_privilege_id} | Destroys a specific privilege.
*ContainersApi* | [**IndexContainerPrivileges**](docs/ContainersApi.md#indexcontainerprivileges) | **Get** /containers/{container_id}/container_privileges | Lists all privilege on the container.
*ContainersApi* | [**IndexContainers**](docs/ContainersApi.md#indexcontainers) | **Get** /containers | Lists all containers.
*ContainersApi* | [**ShowContainer**](docs/ContainersApi.md#showcontainer) | **Get** /containers/{container_id} | Displays a specific container.
*ContainersApi* | [**ShowContainerPrivilege**](docs/ContainersApi.md#showcontainerprivilege) | **Get** /containers/{container_id}/container_privileges/{container_privilege_id} | Displays a specific privilege.
*ContainersApi* | [**UpdateContainer**](docs/ContainersApi.md#updatecontainer) | **Put** /containers/{container_id} | Updates a specific container.
*ContainersApi* | [**UpdateContainerPrivilege**](docs/ContainersApi.md#updatecontainerprivilege) | **Put** /containers/{container_id}/container_privileges/{container_privilege_id} | Updates a specific privilege.
*FilesApi* | [**FilesChildren**](docs/FilesApi.md#fileschildren) | **Get** /files/{file_parent_id}/children | Lists files under a specific folder.
*FilesApi* | [**FilesChildrenByCloudPool**](docs/FilesApi.md#fileschildrenbycloudpool) | **Get** /cloud_pools/{cloud_pool_id}/files/{file_parent_id}/children | Lists files under a specific folder.
*FilesApi* | [**FilesChildrenByContainer**](docs/FilesApi.md#fileschildrenbycontainer) | **Get** /containers/{container_id}/files/{file_parent_id}/children | Lists files under a specific folder.
*FilesApi* | [**FilesChildrenByNasPool**](docs/FilesApi.md#fileschildrenbynaspool) | **Get** /nas_pools/{nas_pool_id}/files/{file_parent_id}/children | Lists files under a specific folder.
*FilesApi* | [**FilesChildrenByTapePool**](docs/FilesApi.md#fileschildrenbytapepool) | **Get** /tape_pools/{tape_pool_id}/files/{file_parent_id}/children | Lists files under a specific folder.
*FilesApi* | [**FilesChildrenByTask**](docs/FilesApi.md#fileschildrenbytask) | **Get** /tasks/{task_id}/files/{file_parent_id}/children | Lists files under a specific folder.
*FilesApi* | [**FilesChildrenByTaskExecution**](docs/FilesApi.md#fileschildrenbytaskexecution) | **Get** /task_executions/{task_execution_id}/files/{file_parent_id}/children | Lists files under a specific folder.
*FilesApi* | [**FilesChildrenByTaskExecutionByTask**](docs/FilesApi.md#fileschildrenbytaskexecutionbytask) | **Get** /tasks/{task_id}/task_executions/{task_execution_id}/files/{file_parent_id}/children | Lists files under a specific folder.
*FilesApi* | [**IndexFiles**](docs/FilesApi.md#indexfiles) | **Get** /files | Lists files on root.
*FilesApi* | [**IndexFilesByCloudPool**](docs/FilesApi.md#indexfilesbycloudpool) | **Get** /cloud_pools/{cloud_pool_id}/files | Lists files on root.
*FilesApi* | [**IndexFilesByContainer**](docs/FilesApi.md#indexfilesbycontainer) | **Get** /containers/{container_id}/files | Lists files on root.
*FilesApi* | [**IndexFilesByNasPool**](docs/FilesApi.md#indexfilesbynaspool) | **Get** /nas_pools/{nas_pool_id}/files | Lists files on root.
*FilesApi* | [**IndexFilesByTapePool**](docs/FilesApi.md#indexfilesbytapepool) | **Get** /tape_pools/{tape_pool_id}/files | Lists files on root.
*FilesApi* | [**IndexFilesByTask**](docs/FilesApi.md#indexfilesbytask) | **Get** /tasks/{task_id}/files | Lists files on root.
*FilesApi* | [**IndexFilesByTaskExecution**](docs/FilesApi.md#indexfilesbytaskexecution) | **Get** /task_executions/{task_execution_id}/files | Lists files on root.
*FilesApi* | [**IndexFilesByTaskExecutionByTask**](docs/FilesApi.md#indexfilesbytaskexecutionbytask) | **Get** /tasks/{task_id}/task_executions/{task_execution_id}/files | Lists files on root.
*FilesApi* | [**ShowFile**](docs/FilesApi.md#showfile) | **Get** /files/{file_id} | Displays a specific file.
*FilesApi* | [**ShowFileByCloudPool**](docs/FilesApi.md#showfilebycloudpool) | **Get** /cloud_pools/{cloud_pool_id}/files/{file_id} | Displays a specific file.
*FilesApi* | [**ShowFileByContainer**](docs/FilesApi.md#showfilebycontainer) | **Get** /containers/{container_id}/files/{file_id} | Displays a specific file.
*FilesApi* | [**ShowFileByNasPool**](docs/FilesApi.md#showfilebynaspool) | **Get** /nas_pools/{nas_pool_id}/files/{file_id} | Displays a specific file.
*FilesApi* | [**ShowFileByTapePool**](docs/FilesApi.md#showfilebytapepool) | **Get** /tape_pools/{tape_pool_id}/files/{file_id} | Displays a specific file.
*FilesApi* | [**ShowFileByTask**](docs/FilesApi.md#showfilebytask) | **Get** /tasks/{task_id}/files/{file_id} | Displays a specific file.
*FilesApi* | [**ShowFileByTaskExecution**](docs/FilesApi.md#showfilebytaskexecution) | **Get** /task_executions/{task_execution_id}/files/{file_id} | Displays a specific file.
*FilesApi* | [**ShowFileByTaskExecutionByTask**](docs/FilesApi.md#showfilebytaskexecutionbytask) | **Get** /tasks/{task_id}/task_executions/{task_execution_id}/files/{file_id} | Displays a specific file.
*MountPointsApi* | [**CreateMountPoint**](docs/MountPointsApi.md#createmountpoint) | **Post** /mount_points | Creates a new mount point.
*MountPointsApi* | [**DestroyMountPoint**](docs/MountPointsApi.md#destroymountpoint) | **Delete** /mount_points/{mount_point_id} | Destroys a specific mount point.
*MountPointsApi* | [**IndexMountPoints**](docs/MountPointsApi.md#indexmountpoints) | **Get** /mount_points | Lists all mount points.
*MountPointsApi* | [**ShowMountPoint**](docs/MountPointsApi.md#showmountpoint) | **Get** /mount_points/{mount_point_id} | Displays a specific mount point&#x60;.
*MountPointsApi* | [**UpdateMountPoint**](docs/MountPointsApi.md#updatemountpoint) | **Put** /mount_points/{mount_point_id} | Updates a specific mount point&#x60;.
*NasApi* | [**CreateNas**](docs/NasApi.md#createnas) | **Post** /nas | Creates a new NAS.
*NasApi* | [**DestroyNas**](docs/NasApi.md#destroynas) | **Delete** /nas/{nas_id} | Destroys a specific NAS.
*NasApi* | [**IndexNas**](docs/NasApi.md#indexnas) | **Get** /nas | Lists all NAS.
*NasApi* | [**ShowNas**](docs/NasApi.md#shownas) | **Get** /nas/{nas_id} | Displays a specific NAS.
*NasApi* | [**UpdateNas**](docs/NasApi.md#updatenas) | **Put** /nas/{nas_id} | Updates a specific NAS.
*NasPoolsApi* | [**CreateNasPool**](docs/NasPoolsApi.md#createnaspool) | **Post** /nas_pools | Creates a new NAS pool.
*NasPoolsApi* | [**DestroyNasPool**](docs/NasPoolsApi.md#destroynaspool) | **Delete** /nas_pools/{nas_pool_id} | Destroys a specific NAS pool.
*NasPoolsApi* | [**IndexNasPools**](docs/NasPoolsApi.md#indexnaspools) | **Get** /nas_pools | Lists all NAS pools.
*NasPoolsApi* | [**ShowNasPool**](docs/NasPoolsApi.md#shownaspool) | **Get** /nas_pools/{nas_pool_id} | Displays a specific NAS pool.
*NasPoolsApi* | [**UpdateNasPool**](docs/NasPoolsApi.md#updatenaspool) | **Put** /nas_pools/{nas_pool_id} | Updates a specific NAS pool.
*NasSharesApi* | [**CreateNasShareByNas**](docs/NasSharesApi.md#createnassharebynas) | **Post** /nas/{nas_id}/nas_shares | Creates a new NAS share.
*NasSharesApi* | [**DestroyNasShare**](docs/NasSharesApi.md#destroynasshare) | **Delete** /nas_shares/{nas_share_id} | Destroys a specific NAS share.
*NasSharesApi* | [**DestroyNasShareByNas**](docs/NasSharesApi.md#destroynassharebynas) | **Delete** /nas/{nas_id}/nas_shares/{nas_share_id} | Destroys a specific NAS share.
*NasSharesApi* | [**DestroyNasShareByNasPool**](docs/NasSharesApi.md#destroynassharebynaspool) | **Delete** /nas_pools/{nas_pool_id}/nas_shares/{nas_share_id} | Destroys a specific NAS share.
*NasSharesApi* | [**IndexNasShares**](docs/NasSharesApi.md#indexnasshares) | **Get** /nas_shares | Lists all NAS shares.
*NasSharesApi* | [**IndexNasSharesByNas**](docs/NasSharesApi.md#indexnassharesbynas) | **Get** /nas/{nas_id}/nas_shares | Lists all NAS shares.
*NasSharesApi* | [**IndexNasSharesByNasPool**](docs/NasSharesApi.md#indexnassharesbynaspool) | **Get** /nas_pools/{nas_pool_id}/nas_shares | Lists all NAS shares.
*NasSharesApi* | [**ShowNasShareByNas**](docs/NasSharesApi.md#shownassharebynas) | **Get** /nas/{nas_id}/nas_shares/{nas_share_id} | Displays a specific NAS share.
*NasSharesApi* | [**ShowNasShares**](docs/NasSharesApi.md#shownasshares) | **Get** /nas_shares/{nas_share_id} | Displays a specific NAS share.
*NasSharesApi* | [**ShowNasSharesByNasPool**](docs/NasSharesApi.md#shownassharesbynaspool) | **Get** /nas_pools/{nas_pool_id}/nas_shares/{nas_share_id} | Displays a specific NAS share.
*NasSharesApi* | [**TestNasShare**](docs/NasSharesApi.md#testnasshare) | **Put** /nas/{nas_id}/nas_shares/-/test | Test an unsaved NAS Share.
*NasSharesApi* | [**TestResultNasShare**](docs/NasSharesApi.md#testresultnasshare) | **Get** /nas/{nas_id}/nas_shares/-/test | Check result of a NAS Share test job.
*NasSharesApi* | [**UpdateNasShare**](docs/NasSharesApi.md#updatenasshare) | **Put** /nas_shares/{nas_share_id} | Updates a specific NAS share.
*NasSharesApi* | [**UpdateNasShareByNas**](docs/NasSharesApi.md#updatenassharebynas) | **Put** /nas/{nas_id}/nas_shares/{nas_share_id} | Updates a specific NAS share.
*NasSharesApi* | [**UpdateNasShareByNasPool**](docs/NasSharesApi.md#updatenassharebynaspool) | **Put** /nas_pools/{nas_pool_id}/nas_shares/{nas_share_id} | Updates a specific NAS share.
*TapeDrivesApi* | [**CreateTapeDriveByTapeLibrary**](docs/TapeDrivesApi.md#createtapedrivebytapelibrary) | **Post** /tape_libraries/{tape_library_id}/tape_drives | Creates a new tape drive.
*TapeDrivesApi* | [**DestroyTapeDrive**](docs/TapeDrivesApi.md#destroytapedrive) | **Delete** /tape_drives/{tape_drive_id} | Destroys a specific tape drive.
*TapeDrivesApi* | [**DestroyTapeDriveByTapeLibrary**](docs/TapeDrivesApi.md#destroytapedrivebytapelibrary) | **Delete** /tape_libraries/{tape_library_id}/tape_drives/{tape_drive_id} | Destroys a specific tape drive.
*TapeDrivesApi* | [**IndexTapeDriveDevices**](docs/TapeDrivesApi.md#indextapedrivedevices) | **Get** /tape_libraries/{tape_library_id}/tape_drives/-/devices | Lists tape drives devices.
*TapeDrivesApi* | [**IndexTapeDrives**](docs/TapeDrivesApi.md#indextapedrives) | **Get** /tape_drives | Lists all tape drives.
*TapeDrivesApi* | [**IndexTapeDrivesByTapeLibrary**](docs/TapeDrivesApi.md#indextapedrivesbytapelibrary) | **Get** /tape_libraries/{tape_library_id}/tape_drives | Lists all tape drives.
*TapeDrivesApi* | [**ShowTapeDrive**](docs/TapeDrivesApi.md#showtapedrive) | **Get** /tape_drives/{tape_drive_id} | Displays a specific tape drive.
*TapeDrivesApi* | [**ShowTapeDriveByTapeLibrary**](docs/TapeDrivesApi.md#showtapedrivebytapelibrary) | **Get** /tape_libraries/{tape_library_id}/tape_drives/{tape_drive_id} | Displays a specific tape drive.
*TapeDrivesApi* | [**UpdateTapeDrive**](docs/TapeDrivesApi.md#updatetapedrive) | **Put** /tape_drives/{tape_drive_id} | Updates a specific tape drive.
*TapeDrivesApi* | [**UpdateTapeDriveByTapeLibrary**](docs/TapeDrivesApi.md#updatetapedrivebytapelibrary) | **Put** /tape_libraries/{tape_library_id}/tape_drives/{tape_drive_id} | Updates a specific tape drive.
*TapeLibrariesApi* | [**CreateTapeLibrary**](docs/TapeLibrariesApi.md#createtapelibrary) | **Post** /tape_libraries | Creates a new tape library.
*TapeLibrariesApi* | [**DestroyTapeLibrary**](docs/TapeLibrariesApi.md#destroytapelibrary) | **Delete** /tape_libraries/{tape_library_id} | Destroys a specific tape library.
*TapeLibrariesApi* | [**IndexTapeLibraries**](docs/TapeLibrariesApi.md#indextapelibraries) | **Get** /tape_libraries | Lists all tape libraries.
*TapeLibrariesApi* | [**IndexTapeLibraryDevices**](docs/TapeLibrariesApi.md#indextapelibrarydevices) | **Get** /tape_libraries/-/devices | Lists tape libraries devices.
*TapeLibrariesApi* | [**ShowTapeLibrary**](docs/TapeLibrariesApi.md#showtapelibrary) | **Get** /tape_libraries/{tape_library_id} | Displays a specific tape library.
*TapeLibrariesApi* | [**UpdateTapeLibrary**](docs/TapeLibrariesApi.md#updatetapelibrary) | **Put** /tape_libraries/{tape_library_id} | Updates a specific tape library.
*TapePoolsApi* | [**CreateTapePool**](docs/TapePoolsApi.md#createtapepool) | **Post** /tape_pools | Creates a new tape pool.
*TapePoolsApi* | [**DestroyTapePool**](docs/TapePoolsApi.md#destroytapepool) | **Delete** /tape_pools/{tape_pool_id} | Destroys a specific tape pool.
*TapePoolsApi* | [**IndexTapePools**](docs/TapePoolsApi.md#indextapepools) | **Get** /tape_pools | Lists all tape pools.
*TapePoolsApi* | [**ShowTapePool**](docs/TapePoolsApi.md#showtapepool) | **Get** /tape_pools/{tape_pool_id} | Displays a specific tape pool.
*TapePoolsApi* | [**UpdateTapePool**](docs/TapePoolsApi.md#updatetapepool) | **Put** /tape_pools/{tape_pool_id} | Updates a specific tape pool.
*TapesApi* | [**IndexTapes**](docs/TapesApi.md#indextapes) | **Get** /tapes | Lists all tapes.
*TapesApi* | [**IndexTapesByTapeLibrary**](docs/TapesApi.md#indextapesbytapelibrary) | **Get** /tape_libraries/{tape_library_id}/tapes | Lists all tapes.
*TapesApi* | [**IndexTapesByTapePool**](docs/TapesApi.md#indextapesbytapepool) | **Get** /tape_pools/{tape_pool_id}/tapes | Lists all tapes.
*TapesApi* | [**ShowTape**](docs/TapesApi.md#showtape) | **Get** /tapes/{tape_id} | Displays a specific tape.
*TapesApi* | [**ShowTapeByTapeLibrary**](docs/TapesApi.md#showtapebytapelibrary) | **Get** /tape_libraries/{tape_library_id}/tapes/{tape_id} | Displays a specific tape.
*TapesApi* | [**ShowTapeByTapePool**](docs/TapesApi.md#showtapebytapepool) | **Get** /tape_pools/{tape_pool_id}/tapes/{tape_id} | Displays a specific tape.
*TaskDestinationsApi* | [**CreateTaskDestination**](docs/TaskDestinationsApi.md#createtaskdestination) | **Post** /tasks/{task_id}/task_destinations | Creates a new task destination.
*TaskDestinationsApi* | [**DestroyTaskDestination**](docs/TaskDestinationsApi.md#destroytaskdestination) | **Delete** /tasks/{task_id}/task_destinations/{task_destination_id} | Destroys a specific task destination.
*TaskDestinationsApi* | [**IndexTaskDestinations**](docs/TaskDestinationsApi.md#indextaskdestinations) | **Get** /tasks/{task_id}/task_destinations | Lists all task destinations.
*TaskDestinationsApi* | [**ShowTaskDestination**](docs/TaskDestinationsApi.md#showtaskdestination) | **Get** /tasks/{task_id}/task_destinations/{task_destination_id} | Displays a specific task destination.
*TaskDestinationsApi* | [**UpdateTaskDestination**](docs/TaskDestinationsApi.md#updatetaskdestination) | **Put** /tasks/{task_id}/task_destinations/{task_destination_id} | Updates a specific task destination.
*TaskExecutionsApi* | [**IndexTaskExecutions**](docs/TaskExecutionsApi.md#indextaskexecutions) | **Get** /task_executions | Lists all task executions.
*TaskExecutionsApi* | [**IndexTaskExecutionsByTask**](docs/TaskExecutionsApi.md#indextaskexecutionsbytask) | **Get** /tasks/{task_id}/task_executions | Lists all task executions.
*TaskExecutionsApi* | [**ShowTaskExecution**](docs/TaskExecutionsApi.md#showtaskexecution) | **Get** /task_executions/{task_execution_id} | Displays a specific task execution.
*TaskExecutionsApi* | [**ShowTaskExecutionByTask**](docs/TaskExecutionsApi.md#showtaskexecutionbytask) | **Get** /tasks/{task_id}/task_executions/{task_execution_id} | Displays a specific task execution.
*TaskMetadataApi* | [**CreateTaskMetadatum**](docs/TaskMetadataApi.md#createtaskmetadatum) | **Post** /tasks/{task_id}/task_metadata | Creates a new task metadatum.
*TaskMetadataApi* | [**DestroyTaskMetadatum**](docs/TaskMetadataApi.md#destroytaskmetadatum) | **Delete** /tasks/{task_id}/task_metadata/{task_metadatum_id} | Destroys a specific task metadatum.
*TaskMetadataApi* | [**IndexTaskMetadata**](docs/TaskMetadataApi.md#indextaskmetadata) | **Get** /tasks/{task_id}/task_metadata | Lists all task metadata.
*TaskMetadataApi* | [**ShowTaskMetadat**](docs/TaskMetadataApi.md#showtaskmetadat) | **Get** /tasks/{task_id}/task_metadata/{task_metadatum_id} | Displays a specific task metadatum.
*TaskMetadataApi* | [**UpdateTaskMetadatum**](docs/TaskMetadataApi.md#updatetaskmetadatum) | **Put** /tasks/{task_id}/task_metadata/{task_metadatum_id} | Updates a specific task metadatum.
*TaskOptionsApi* | [**CreateTaskOption**](docs/TaskOptionsApi.md#createtaskoption) | **Post** /tasks/{task_id}/task_options | Creates a new task option.
*TaskOptionsApi* | [**DestroyTaskOption**](docs/TaskOptionsApi.md#destroytaskoption) | **Delete** /tasks/{task_id}/task_options/{task_option_id} | Destroys a specific task option.
*TaskOptionsApi* | [**IndexTaskOptions**](docs/TaskOptionsApi.md#indextaskoptions) | **Get** /tasks/{task_id}/task_options | Lists all task options.
*TaskOptionsApi* | [**ShowTaskOption**](docs/TaskOptionsApi.md#showtaskoption) | **Get** /tasks/{task_id}/task_options/{task_option_id} | Displays a specific task option.
*TaskOptionsApi* | [**UpdateTaskOption**](docs/TaskOptionsApi.md#updatetaskoption) | **Put** /tasks/{task_id}/task_options/{task_option_id} | Updates a specific task option.
*TaskSourcesApi* | [**CreateTaskSource**](docs/TaskSourcesApi.md#createtasksource) | **Post** /tasks/{task_id}/task_sources | Creates a new task source.
*TaskSourcesApi* | [**DestroyTaskSource**](docs/TaskSourcesApi.md#destroytasksource) | **Delete** /tasks/{task_id}/task_sources/{task_source_id} | Destroys a specific task source.
*TaskSourcesApi* | [**IndexTaskSources**](docs/TaskSourcesApi.md#indextasksources) | **Get** /tasks/{task_id}/task_sources | Lists all task sources.
*TaskSourcesApi* | [**ShowTaskSource**](docs/TaskSourcesApi.md#showtasksource) | **Get** /tasks/{task_id}/task_sources/{task_source_id} | Displays a specific task source.
*TaskSourcesApi* | [**UpdateTaskSource**](docs/TaskSourcesApi.md#updatetasksource) | **Put** /tasks/{task_id}/task_sources/{task_source_id} | Updates a specific task source.
*TasksApi* | [**CreateTask**](docs/TasksApi.md#createtask) | **Post** /tasks | Creates a new task.
*TasksApi* | [**DestroyTask**](docs/TasksApi.md#destroytask) | **Delete** /tasks/{task_id} | Destroys a specific task.
*TasksApi* | [**IndexTasks**](docs/TasksApi.md#indextasks) | **Get** /tasks | Lists all tasks.
*TasksApi* | [**ShowTask**](docs/TasksApi.md#showtask) | **Get** /tasks/{task_id} | Displays a specific task.
*TasksApi* | [**UpdateTask**](docs/TasksApi.md#updatetask) | **Put** /tasks/{task_id} | Updates a specific task.
*UsersApi* | [**CreateApiKey**](docs/UsersApi.md#createapikey) | **Post** /users/me/api_keys | Creates a new API Key for current user.
*UsersApi* | [**DestroyApiKey**](docs/UsersApi.md#destroyapikey) | **Delete** /users/me/api_keys/{api_key_id} | Destroys a specific API Key.
*UsersApi* | [**IndexApiKeys**](docs/UsersApi.md#indexapikeys) | **Get** /users/me/api_keys | Lists all API keys of current user.
*UsersApi* | [**ShowApiKey**](docs/UsersApi.md#showapikey) | **Get** /users/me/api_keys/{api_key_id} | Displays a specific API Key with its scopes.
*UsersApi* | [**UpdateApiKey**](docs/UsersApi.md#updateapikey) | **Put** /users/me/api_keys/{api_key_id} | Updates a specific API Key.


## Documentation For Models

 - [ActiveJobStatus](docs/ActiveJobStatus.md)
 - [ApiKey](docs/ApiKey.md)
 - [ApiKeyCollection](docs/ApiKeyCollection.md)
 - [ApiKeyFull](docs/ApiKeyFull.md)
 - [ApiKeyScope](docs/ApiKeyScope.md)
 - [AttributeError](docs/AttributeError.md)
 - [Blank](docs/Blank.md)
 - [CloudBucket](docs/CloudBucket.md)
 - [CloudBucketCollection](docs/CloudBucketCollection.md)
 - [CloudBucketSimpleCollection](docs/CloudBucketSimpleCollection.md)
 - [CloudConnector](docs/CloudConnector.md)
 - [CloudConnectorCollection](docs/CloudConnectorCollection.md)
 - [CloudPool](docs/CloudPool.md)
 - [CloudPoolCollection](docs/CloudPoolCollection.md)
 - [CloudPoolUp](docs/CloudPoolUp.md)
 - [Container](docs/Container.md)
 - [ContainerCollection](docs/ContainerCollection.md)
 - [ContainerPrivilege](docs/ContainerPrivilege.md)
 - [ContainerPrivilegeCollection](docs/ContainerPrivilegeCollection.md)
 - [Frozen](docs/Frozen.md)
 - [GreaterThan](docs/GreaterThan.md)
 - [GreaterThanOrEqualTo](docs/GreaterThanOrEqualTo.md)
 - [ImportFile](docs/ImportFile.md)
 - [Invalid](docs/Invalid.md)
 - [LessThan](docs/LessThan.md)
 - [LessThanOrEqualTo](docs/LessThanOrEqualTo.md)
 - [ModelError](docs/ModelError.md)
 - [MountPoint](docs/MountPoint.md)
 - [MountPointCollection](docs/MountPointCollection.md)
 - [Nas](docs/Nas.md)
 - [NasCollection](docs/NasCollection.md)
 - [NasPool](docs/NasPool.md)
 - [NasPoolCollection](docs/NasPoolCollection.md)
 - [NasPoolUp](docs/NasPoolUp.md)
 - [NasShare](docs/NasShare.md)
 - [NasShareCollection](docs/NasShareCollection.md)
 - [NodeumFile](docs/NodeumFile.md)
 - [NodeumFileCollection](docs/NodeumFileCollection.md)
 - [NodeumFileWithPath](docs/NodeumFileWithPath.md)
 - [OccurrenceLessThan](docs/OccurrenceLessThan.md)
 - [OccurrenceLessThanOrEqualTo](docs/OccurrenceLessThanOrEqualTo.md)
 - [QuotaOnCache](docs/QuotaOnCache.md)
 - [Taken](docs/Taken.md)
 - [Tape](docs/Tape.md)
 - [TapeCollection](docs/TapeCollection.md)
 - [TapeDrive](docs/TapeDrive.md)
 - [TapeDriveCollection](docs/TapeDriveCollection.md)
 - [TapeDriveDevice](docs/TapeDriveDevice.md)
 - [TapeDriveDeviceCollection](docs/TapeDriveDeviceCollection.md)
 - [TapeLibrary](docs/TapeLibrary.md)
 - [TapeLibraryCollection](docs/TapeLibraryCollection.md)
 - [TapeLibraryDevice](docs/TapeLibraryDevice.md)
 - [TapeLibraryDeviceCollection](docs/TapeLibraryDeviceCollection.md)
 - [TapePool](docs/TapePool.md)
 - [TapePoolCollection](docs/TapePoolCollection.md)
 - [TapePoolUp](docs/TapePoolUp.md)
 - [Task](docs/Task.md)
 - [TaskCollection](docs/TaskCollection.md)
 - [TaskDestinationCollection](docs/TaskDestinationCollection.md)
 - [TaskDestinationDown](docs/TaskDestinationDown.md)
 - [TaskDestinationUp](docs/TaskDestinationUp.md)
 - [TaskExecution](docs/TaskExecution.md)
 - [TaskExecutionCollection](docs/TaskExecutionCollection.md)
 - [TaskMetadatum](docs/TaskMetadatum.md)
 - [TaskMetadatumCollection](docs/TaskMetadatumCollection.md)
 - [TaskOption](docs/TaskOption.md)
 - [TaskOptionCollection](docs/TaskOptionCollection.md)
 - [TaskSourceCollection](docs/TaskSourceCollection.md)
 - [TaskSourceDown](docs/TaskSourceDown.md)
 - [TaskSourceUp](docs/TaskSourceUp.md)
 - [TooLong](docs/TooLong.md)
 - [TooShort](docs/TooShort.md)


## Documentation For Authorization

## BasicAuth
- **Type**: HTTP basic authentication

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```
## BearerAuth
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author



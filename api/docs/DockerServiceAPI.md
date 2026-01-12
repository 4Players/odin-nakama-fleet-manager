# \DockerServiceAPI

All URIs are relative to *https://fleet.4players.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBackup**](DockerServiceAPI.md#CreateBackup) | **Post** /v1/services/{dockerService}/backup | Create service backup
[**DockerServicesMetadataDeleteAll**](DockerServiceAPI.md#DockerServicesMetadataDeleteAll) | **Delete** /v1/services/{dockerService}/metadata | Delete all service metadata
[**DockerServicesMetadataDeleteKeys**](DockerServiceAPI.md#DockerServicesMetadataDeleteKeys) | **Delete** /v1/services/{dockerService}/metadata/keys | Delete service metadata keys
[**DockerServicesMetadataSet**](DockerServiceAPI.md#DockerServicesMetadataSet) | **Put** /v1/services/{dockerService}/metadata | Set service metadata
[**DockerServicesMetadataUpdate**](DockerServiceAPI.md#DockerServicesMetadataUpdate) | **Patch** /v1/services/{dockerService}/metadata | Update service metadata
[**DownloadServerLogs**](DockerServiceAPI.md#DownloadServerLogs) | **Get** /v1/services/{dockerService}/logs/download | Download service logs
[**GetBackups**](DockerServiceAPI.md#GetBackups) | **Get** /v1/services/{dockerService}/backups | List service backups
[**GetLatestBackup**](DockerServiceAPI.md#GetLatestBackup) | **Get** /v1/services/{dockerService}/backup | Get latest service backup
[**GetServerBackupDownloadUrl**](DockerServiceAPI.md#GetServerBackupDownloadUrl) | **Get** /v1/services/{dockerService}/backup/download | Get service backup download URL
[**GetServerById**](DockerServiceAPI.md#GetServerById) | **Get** /v1/apps/{app}/services/{dockerService} | Display a specific service
[**GetServerLogs**](DockerServiceAPI.md#GetServerLogs) | **Get** /v1/services/{dockerService}/logs | Get service logs
[**GetServers**](DockerServiceAPI.md#GetServers) | **Get** /v1/apps/{app}/services | List services
[**RestartServer**](DockerServiceAPI.md#RestartServer) | **Post** /v1/services/{dockerService}/restart | Restart service
[**RestoreBackup**](DockerServiceAPI.md#RestoreBackup) | **Post** /v1/services/{dockerService}/restore | Restore latest service backup
[**StartServer**](DockerServiceAPI.md#StartServer) | **Post** /v1/services/{dockerService}/start | Start service
[**StopServer**](DockerServiceAPI.md#StopServer) | **Post** /v1/services/{dockerService}/stop | Stop service



## CreateBackup

> CreateBackup(ctx, dockerService).CreateBackupDockerServiceRequest(createBackupDockerServiceRequest).Execute()

Create service backup

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	dockerService := int32(56) // int32 | The docker service ID
	createBackupDockerServiceRequest := *openapiclient.NewCreateBackupDockerServiceRequest("Name_example") // CreateBackupDockerServiceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DockerServiceAPI.CreateBackup(context.Background(), dockerService).CreateBackupDockerServiceRequest(createBackupDockerServiceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerServiceAPI.CreateBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dockerService** | **int32** | The docker service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createBackupDockerServiceRequest** | [**CreateBackupDockerServiceRequest**](CreateBackupDockerServiceRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DockerServicesMetadataDeleteAll

> Server DockerServicesMetadataDeleteAll(ctx, dockerService).Execute()

Delete all service metadata

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	dockerService := int32(56) // int32 | The docker service ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DockerServiceAPI.DockerServicesMetadataDeleteAll(context.Background(), dockerService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerServiceAPI.DockerServicesMetadataDeleteAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DockerServicesMetadataDeleteAll`: Server
	fmt.Fprintf(os.Stdout, "Response from `DockerServiceAPI.DockerServicesMetadataDeleteAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dockerService** | **int32** | The docker service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDockerServicesMetadataDeleteAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Server**](Server.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DockerServicesMetadataDeleteKeys

> Server DockerServicesMetadataDeleteKeys(ctx, dockerService).Metadata(metadata).Execute()

Delete service metadata keys

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	dockerService := int32(56) // int32 | The docker service ID
	metadata := []string{"Inner_example"} // []string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DockerServiceAPI.DockerServicesMetadataDeleteKeys(context.Background(), dockerService).Metadata(metadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerServiceAPI.DockerServicesMetadataDeleteKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DockerServicesMetadataDeleteKeys`: Server
	fmt.Fprintf(os.Stdout, "Response from `DockerServiceAPI.DockerServicesMetadataDeleteKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dockerService** | **int32** | The docker service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDockerServicesMetadataDeleteKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metadata** | **[]string** |  | 

### Return type

[**Server**](Server.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DockerServicesMetadataSet

> Server DockerServicesMetadataSet(ctx, dockerService).SetMetadataRequest(setMetadataRequest).Execute()

Set service metadata



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	dockerService := int32(56) // int32 | The docker service ID
	setMetadataRequest := *openapiclient.NewSetMetadataRequest() // SetMetadataRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DockerServiceAPI.DockerServicesMetadataSet(context.Background(), dockerService).SetMetadataRequest(setMetadataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerServiceAPI.DockerServicesMetadataSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DockerServicesMetadataSet`: Server
	fmt.Fprintf(os.Stdout, "Response from `DockerServiceAPI.DockerServicesMetadataSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dockerService** | **int32** | The docker service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDockerServicesMetadataSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setMetadataRequest** | [**SetMetadataRequest**](SetMetadataRequest.md) |  | 

### Return type

[**Server**](Server.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DockerServicesMetadataUpdate

> Server DockerServicesMetadataUpdate(ctx, dockerService).PatchMetadataRequest(patchMetadataRequest).Execute()

Update service metadata



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	dockerService := int32(56) // int32 | The docker service ID
	patchMetadataRequest := *openapiclient.NewPatchMetadataRequest() // PatchMetadataRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DockerServiceAPI.DockerServicesMetadataUpdate(context.Background(), dockerService).PatchMetadataRequest(patchMetadataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerServiceAPI.DockerServicesMetadataUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DockerServicesMetadataUpdate`: Server
	fmt.Fprintf(os.Stdout, "Response from `DockerServiceAPI.DockerServicesMetadataUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dockerService** | **int32** | The docker service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDockerServicesMetadataUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchMetadataRequest** | [**PatchMetadataRequest**](PatchMetadataRequest.md) |  | 

### Return type

[**Server**](Server.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadServerLogs

> ServiceLogs DownloadServerLogs(ctx, dockerService).StreamSource(streamSource).Execute()

Download service logs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	dockerService := int32(56) // int32 | The docker service ID
	streamSource := "streamSource_example" // string | Only return logs filtered by stream source like stdout or stderr. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DockerServiceAPI.DownloadServerLogs(context.Background(), dockerService).StreamSource(streamSource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerServiceAPI.DownloadServerLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadServerLogs`: ServiceLogs
	fmt.Fprintf(os.Stdout, "Response from `DockerServiceAPI.DownloadServerLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dockerService** | **int32** | The docker service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadServerLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **streamSource** | **string** | Only return logs filtered by stream source like stdout or stderr. | 

### Return type

[**ServiceLogs**](ServiceLogs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackups

> GetBackups200Response GetBackups(ctx, dockerService).PerPage(perPage).Page(page).Sort(sort).FilterName(filterName).FilterArchiveName(filterArchiveName).Execute()

List service backups

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	dockerService := int32(56) // int32 | The docker service ID
	perPage := int32(56) // int32 | The number of items to be shown per page. (optional)
	page := int32(56) // int32 | Specifies the page number to retrieve in the paginated results. (optional)
	sort := []string{"Inner_example"} // []string | Allows sorting of results. By default, sorting is in ascending order. To reverse the order, prepend the sort key with a hyphen (-).  **Simple Sort:** For example, to sort by name in ascending order or by archiveName in descending order:  ``` sort[]=name sort[]=-archiveName ```  **Multiple Sorts:** Combine multiple sorts by separating them with commas: ``` sort[]=name&sort[]=-archiveName ``` (optional)
	filterName := "filterName_example" // string | Filter by name. (optional)
	filterArchiveName := "filterArchiveName_example" // string | Filter by archive name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DockerServiceAPI.GetBackups(context.Background(), dockerService).PerPage(perPage).Page(page).Sort(sort).FilterName(filterName).FilterArchiveName(filterArchiveName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerServiceAPI.GetBackups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackups`: GetBackups200Response
	fmt.Fprintf(os.Stdout, "Response from `DockerServiceAPI.GetBackups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dockerService** | **int32** | The docker service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of items to be shown per page. | 
 **page** | **int32** | Specifies the page number to retrieve in the paginated results. | 
 **sort** | **[]string** | Allows sorting of results. By default, sorting is in ascending order. To reverse the order, prepend the sort key with a hyphen (-).  **Simple Sort:** For example, to sort by name in ascending order or by archiveName in descending order:  &#x60;&#x60;&#x60; sort[]&#x3D;name sort[]&#x3D;-archiveName &#x60;&#x60;&#x60;  **Multiple Sorts:** Combine multiple sorts by separating them with commas: &#x60;&#x60;&#x60; sort[]&#x3D;name&amp;sort[]&#x3D;-archiveName &#x60;&#x60;&#x60; | 
 **filterName** | **string** | Filter by name. | 
 **filterArchiveName** | **string** | Filter by archive name. | 

### Return type

[**GetBackups200Response**](GetBackups200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestBackup

> Backup GetLatestBackup(ctx, dockerService).Execute()

Get latest service backup

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	dockerService := int32(56) // int32 | The docker service ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DockerServiceAPI.GetLatestBackup(context.Background(), dockerService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerServiceAPI.GetLatestBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLatestBackup`: Backup
	fmt.Fprintf(os.Stdout, "Response from `DockerServiceAPI.GetLatestBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dockerService** | **int32** | The docker service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Backup**](Backup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerBackupDownloadUrl

> BackupDownload GetServerBackupDownloadUrl(ctx, dockerService).Execute()

Get service backup download URL

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	dockerService := int32(56) // int32 | The docker service ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DockerServiceAPI.GetServerBackupDownloadUrl(context.Background(), dockerService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerServiceAPI.GetServerBackupDownloadUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerBackupDownloadUrl`: BackupDownload
	fmt.Fprintf(os.Stdout, "Response from `DockerServiceAPI.GetServerBackupDownloadUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dockerService** | **int32** | The docker service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerBackupDownloadUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BackupDownload**](BackupDownload.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerById

> Server GetServerById(ctx, app, dockerService).Execute()

Display a specific service

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	app := int32(56) // int32 | The app ID
	dockerService := int32(56) // int32 | The docker service ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DockerServiceAPI.GetServerById(context.Background(), app, dockerService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerServiceAPI.GetServerById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerById`: Server
	fmt.Fprintf(os.Stdout, "Response from `DockerServiceAPI.GetServerById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **int32** | The app ID | 
**dockerService** | **int32** | The docker service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Server**](Server.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerLogs

> ServiceLogs GetServerLogs(ctx, dockerService).Limit(limit).Direction(direction).StreamSource(streamSource).Execute()

Get service logs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	dockerService := int32(56) // int32 | The docker service ID
	limit := int32(56) // int32 | The max number of entries to return. Default: 100 (optional)
	direction := "direction_example" // string | Determines the sort order of logs. Supported values are forward or backward. Default: forward (optional)
	streamSource := "streamSource_example" // string | Only return logs filtered by stream source like stdout or stderr. Default: null (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DockerServiceAPI.GetServerLogs(context.Background(), dockerService).Limit(limit).Direction(direction).StreamSource(streamSource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerServiceAPI.GetServerLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerLogs`: ServiceLogs
	fmt.Fprintf(os.Stdout, "Response from `DockerServiceAPI.GetServerLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dockerService** | **int32** | The docker service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The max number of entries to return. Default: 100 | 
 **direction** | **string** | Determines the sort order of logs. Supported values are forward or backward. Default: forward | 
 **streamSource** | **string** | Only return logs filtered by stream source like stdout or stderr. Default: null | 

### Return type

[**ServiceLogs**](ServiceLogs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServers

> GetServers200Response GetServers(ctx, app).PerPage(perPage).Page(page).FilterStatus(filterStatus).FilterAppLocationSettingId(filterAppLocationSettingId).FilterServerConfigId(filterServerConfigId).FilterServerConfigName(filterServerConfigName).FilterServerConfigNamePartial(filterServerConfigNamePartial).FilterLocationCity(filterLocationCity).FilterLocationCityDisplay(filterLocationCityDisplay).FilterLocationContinent(filterLocationContinent).FilterLocationCountry(filterLocationCountry).FilterIsBackupable(filterIsBackupable).FilterIsRestorable(filterIsRestorable).FilterIsPending(filterIsPending).FilterIsNotFound(filterIsNotFound).FilterIsHealthy(filterIsHealthy).FilterBinaryId(filterBinaryId).FilterIsStopped(filterIsStopped).FilterMetadata(filterMetadata).Sort(sort).Execute()

List services

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	app := int32(56) // int32 | The app ID
	perPage := int32(56) // int32 | The number of items to be shown per page. (optional)
	page := int32(56) // int32 | Specifies the page number to retrieve in the paginated results. (optional)
	filterStatus := "filterStatus_example" // string | Filter by status. (optional)
	filterAppLocationSettingId := int32(56) // int32 | Filter by AppLocationSetting ID. (optional)
	filterServerConfigId := int32(56) // int32 | Filter by ServerConfig ID. (optional)
	filterServerConfigName := "filterServerConfigName_example" // string | Filter by ServerConfig name. (optional)
	filterServerConfigNamePartial := "filterServerConfigNamePartial_example" // string | Filter by ServerConfig name using partial matching. For example, \"ann\" matches \"Joanna\" or \"Annie\". (optional)
	filterLocationCity := "filterLocationCity_example" // string | Filter by location city. (optional)
	filterLocationCityDisplay := "filterLocationCityDisplay_example" // string | Filter by location city display name. (optional)
	filterLocationContinent := "filterLocationContinent_example" // string | Filter by location continent. (optional)
	filterLocationCountry := "filterLocationCountry_example" // string | Filter by location country. (optional)
	filterIsBackupable := true // bool | Filter by whether the service can be backed up. (optional)
	filterIsRestorable := true // bool | Filter by whether the service can be restored. (optional)
	filterIsPending := true // bool | Filter by whether the service is pending (not running) due to insufficient resources on the node. (optional)
	filterIsNotFound := true // bool | Filter by whether the service is not found/missing in the cluster. (optional)
	filterIsHealthy := true // bool | Filter by whether the service is currently in an overall healthy state. (optional)
	filterBinaryId := int32(56) // int32 | Filter by Binary ID. (optional)
	filterIsStopped := true // bool | Filter by whether the service is currenctly stopped. (optional)
	filterMetadata := "filterMetadata_example" // string | Filter by metadata. Allows filtering based on metadata key-value pairs, supporting both simple and nested metadata fields using dot notation.  **Simple Filters:** To filter where `idle` is false (boolean): ``` filter[metadata]=idle=false ```  To filter where `string` is exactly \"a\": ``` filter[metadata]=string=\"a\" ```  **Filtering for Null Values:** To filter for a native null value, use unquoted null. For example, to filter where `score` is null: ``` filter[metadata]=score=null ```  **Nested Filters:** For nested metadata fields use dot notation. For example, to filter where `difficulty` within `gameSettings.survival` is exactly \"hardcore\": ``` filter[metadata]=gameSettings.survival.difficulty=\"hardcore\" ```  To filter for a nested field with a native `null` value, leave the null unquoted: ``` filter[metadata]=gameSettings.stats.score=null ```  **Array Contains Filter:** To filter where an array contains a given value (string, number, boolean or null): ``` filter[metadata]=players=\"foobar\" filter[metadata]=player_ids=37 filter[metadata]=array=true filter[metadata]=array=null ```  Works for nested arrays as well: ``` filter[metadata]=gameData.players=\"foobar\" ```  **Multiple Filters:** Combine multiple filters by separating them with commas: ``` filter[metadata]=idle=false,max_players=20,gameSettings.survival.difficulty=\"hardcore\" ``` (optional)
	sort := []string{"Inner_example"} // []string | Allows sorting of results. By default, sorting is in ascending order. To reverse the order, prepend the sort key with a hyphen (-).  **Simple Sort:** To sort by id in ascending order or by instance in descending order:  ``` sort[]=id sort[]=-instance ```  **Multiple Sorts:** Combine multiple sorts by separating them with commas: ``` sort[]=id&sort[]=-instance ``` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DockerServiceAPI.GetServers(context.Background(), app).PerPage(perPage).Page(page).FilterStatus(filterStatus).FilterAppLocationSettingId(filterAppLocationSettingId).FilterServerConfigId(filterServerConfigId).FilterServerConfigName(filterServerConfigName).FilterServerConfigNamePartial(filterServerConfigNamePartial).FilterLocationCity(filterLocationCity).FilterLocationCityDisplay(filterLocationCityDisplay).FilterLocationContinent(filterLocationContinent).FilterLocationCountry(filterLocationCountry).FilterIsBackupable(filterIsBackupable).FilterIsRestorable(filterIsRestorable).FilterIsPending(filterIsPending).FilterIsNotFound(filterIsNotFound).FilterIsHealthy(filterIsHealthy).FilterBinaryId(filterBinaryId).FilterIsStopped(filterIsStopped).FilterMetadata(filterMetadata).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerServiceAPI.GetServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServers`: GetServers200Response
	fmt.Fprintf(os.Stdout, "Response from `DockerServiceAPI.GetServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **int32** | The app ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of items to be shown per page. | 
 **page** | **int32** | Specifies the page number to retrieve in the paginated results. | 
 **filterStatus** | **string** | Filter by status. | 
 **filterAppLocationSettingId** | **int32** | Filter by AppLocationSetting ID. | 
 **filterServerConfigId** | **int32** | Filter by ServerConfig ID. | 
 **filterServerConfigName** | **string** | Filter by ServerConfig name. | 
 **filterServerConfigNamePartial** | **string** | Filter by ServerConfig name using partial matching. For example, \&quot;ann\&quot; matches \&quot;Joanna\&quot; or \&quot;Annie\&quot;. | 
 **filterLocationCity** | **string** | Filter by location city. | 
 **filterLocationCityDisplay** | **string** | Filter by location city display name. | 
 **filterLocationContinent** | **string** | Filter by location continent. | 
 **filterLocationCountry** | **string** | Filter by location country. | 
 **filterIsBackupable** | **bool** | Filter by whether the service can be backed up. | 
 **filterIsRestorable** | **bool** | Filter by whether the service can be restored. | 
 **filterIsPending** | **bool** | Filter by whether the service is pending (not running) due to insufficient resources on the node. | 
 **filterIsNotFound** | **bool** | Filter by whether the service is not found/missing in the cluster. | 
 **filterIsHealthy** | **bool** | Filter by whether the service is currently in an overall healthy state. | 
 **filterBinaryId** | **int32** | Filter by Binary ID. | 
 **filterIsStopped** | **bool** | Filter by whether the service is currenctly stopped. | 
 **filterMetadata** | **string** | Filter by metadata. Allows filtering based on metadata key-value pairs, supporting both simple and nested metadata fields using dot notation.  **Simple Filters:** To filter where &#x60;idle&#x60; is false (boolean): &#x60;&#x60;&#x60; filter[metadata]&#x3D;idle&#x3D;false &#x60;&#x60;&#x60;  To filter where &#x60;string&#x60; is exactly \&quot;a\&quot;: &#x60;&#x60;&#x60; filter[metadata]&#x3D;string&#x3D;\&quot;a\&quot; &#x60;&#x60;&#x60;  **Filtering for Null Values:** To filter for a native null value, use unquoted null. For example, to filter where &#x60;score&#x60; is null: &#x60;&#x60;&#x60; filter[metadata]&#x3D;score&#x3D;null &#x60;&#x60;&#x60;  **Nested Filters:** For nested metadata fields use dot notation. For example, to filter where &#x60;difficulty&#x60; within &#x60;gameSettings.survival&#x60; is exactly \&quot;hardcore\&quot;: &#x60;&#x60;&#x60; filter[metadata]&#x3D;gameSettings.survival.difficulty&#x3D;\&quot;hardcore\&quot; &#x60;&#x60;&#x60;  To filter for a nested field with a native &#x60;null&#x60; value, leave the null unquoted: &#x60;&#x60;&#x60; filter[metadata]&#x3D;gameSettings.stats.score&#x3D;null &#x60;&#x60;&#x60;  **Array Contains Filter:** To filter where an array contains a given value (string, number, boolean or null): &#x60;&#x60;&#x60; filter[metadata]&#x3D;players&#x3D;\&quot;foobar\&quot; filter[metadata]&#x3D;player_ids&#x3D;37 filter[metadata]&#x3D;array&#x3D;true filter[metadata]&#x3D;array&#x3D;null &#x60;&#x60;&#x60;  Works for nested arrays as well: &#x60;&#x60;&#x60; filter[metadata]&#x3D;gameData.players&#x3D;\&quot;foobar\&quot; &#x60;&#x60;&#x60;  **Multiple Filters:** Combine multiple filters by separating them with commas: &#x60;&#x60;&#x60; filter[metadata]&#x3D;idle&#x3D;false,max_players&#x3D;20,gameSettings.survival.difficulty&#x3D;\&quot;hardcore\&quot; &#x60;&#x60;&#x60; | 
 **sort** | **[]string** | Allows sorting of results. By default, sorting is in ascending order. To reverse the order, prepend the sort key with a hyphen (-).  **Simple Sort:** To sort by id in ascending order or by instance in descending order:  &#x60;&#x60;&#x60; sort[]&#x3D;id sort[]&#x3D;-instance &#x60;&#x60;&#x60;  **Multiple Sorts:** Combine multiple sorts by separating them with commas: &#x60;&#x60;&#x60; sort[]&#x3D;id&amp;sort[]&#x3D;-instance &#x60;&#x60;&#x60; | 

### Return type

[**GetServers200Response**](GetServers200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartServer

> RestartServer(ctx, dockerService).Execute()

Restart service

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	dockerService := int32(56) // int32 | The docker service ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DockerServiceAPI.RestartServer(context.Background(), dockerService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerServiceAPI.RestartServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dockerService** | **int32** | The docker service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreBackup

> RestoreBackup(ctx, dockerService).Execute()

Restore latest service backup

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	dockerService := int32(56) // int32 | The docker service ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DockerServiceAPI.RestoreBackup(context.Background(), dockerService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerServiceAPI.RestoreBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dockerService** | **int32** | The docker service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartServer

> StartServer(ctx, dockerService).Execute()

Start service

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	dockerService := int32(56) // int32 | The docker service ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DockerServiceAPI.StartServer(context.Background(), dockerService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerServiceAPI.StartServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dockerService** | **int32** | The docker service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopServer

> StopServer(ctx, dockerService).Execute()

Stop service

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	dockerService := int32(56) // int32 | The docker service ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DockerServiceAPI.StopServer(context.Background(), dockerService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerServiceAPI.StopServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dockerService** | **int32** | The docker service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
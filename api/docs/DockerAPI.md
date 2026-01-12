# \DockerAPI

All URIs are relative to *https://fleet.4players.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBackup**](DockerAPI.md#CreateBackup) | **Post** /v1/services/{dockerService}/backup | Create service backup
[**CreateDockerRegistry**](DockerAPI.md#CreateDockerRegistry) | **Post** /v1/docker-registries | Create a new docker registry
[**DeleteDockerRegistry**](DockerAPI.md#DeleteDockerRegistry) | **Delete** /v1/docker-registries/{dockerRegistry} | Delete a specific docker registry
[**DockerServicesMetadataDeleteAll**](DockerAPI.md#DockerServicesMetadataDeleteAll) | **Delete** /v1/services/{dockerService}/metadata | Delete all service metadata
[**DockerServicesMetadataDeleteKeys**](DockerAPI.md#DockerServicesMetadataDeleteKeys) | **Delete** /v1/services/{dockerService}/metadata/keys | Delete service metadata keys
[**DockerServicesMetadataSet**](DockerAPI.md#DockerServicesMetadataSet) | **Put** /v1/services/{dockerService}/metadata | Set service metadata
[**DockerServicesMetadataUpdate**](DockerAPI.md#DockerServicesMetadataUpdate) | **Patch** /v1/services/{dockerService}/metadata | Update service metadata
[**DownloadServerLogs**](DockerAPI.md#DownloadServerLogs) | **Get** /v1/services/{dockerService}/logs/download | Download service logs
[**GetBackups**](DockerAPI.md#GetBackups) | **Get** /v1/services/{dockerService}/backups | List service backups
[**GetDockerRegistries**](DockerAPI.md#GetDockerRegistries) | **Get** /v1/docker-registries | Show all docker registries
[**GetDockerRegistryById**](DockerAPI.md#GetDockerRegistryById) | **Get** /v1/docker-registries/{dockerRegistry} | Display a specific docker registry
[**GetLatestBackup**](DockerAPI.md#GetLatestBackup) | **Get** /v1/services/{dockerService}/backup | Get latest service backup
[**GetServerBackupDownloadUrl**](DockerAPI.md#GetServerBackupDownloadUrl) | **Get** /v1/services/{dockerService}/backup/download | Get service backup download URL
[**GetServerById**](DockerAPI.md#GetServerById) | **Get** /v1/apps/{app}/services/{dockerService} | Display a specific service
[**GetServerLogs**](DockerAPI.md#GetServerLogs) | **Get** /v1/services/{dockerService}/logs | Get service logs
[**GetServers**](DockerAPI.md#GetServers) | **Get** /v1/apps/{app}/services | List services
[**GetTaggedImages**](DockerAPI.md#GetTaggedImages) | **Get** /v1/docker-registries/{dockerRegistry}/tagged-images | List all available tagged images
[**RefreshTaggedImages**](DockerAPI.md#RefreshTaggedImages) | **Get** /v1/docker-registries/{dockerRegistry}/tagged-images/refresh | Refresh the cache for all available tagged images
[**RestartServer**](DockerAPI.md#RestartServer) | **Post** /v1/services/{dockerService}/restart | Restart service
[**RestoreBackup**](DockerAPI.md#RestoreBackup) | **Post** /v1/services/{dockerService}/restore | Restore latest service backup
[**StartServer**](DockerAPI.md#StartServer) | **Post** /v1/services/{dockerService}/start | Start service
[**StopServer**](DockerAPI.md#StopServer) | **Post** /v1/services/{dockerService}/stop | Stop service
[**UpdateDockerRegistry**](DockerAPI.md#UpdateDockerRegistry) | **Put** /v1/docker-registries/{dockerRegistry} | Update a specific docker registry



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
	r, err := apiClient.DockerAPI.CreateBackup(context.Background(), dockerService).CreateBackupDockerServiceRequest(createBackupDockerServiceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerAPI.CreateBackup``: %v\n", err)
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


## CreateDockerRegistry

> DockerRegistry CreateDockerRegistry(ctx).StoreDockerRegistryRequest(storeDockerRegistryRequest).Execute()

Create a new docker registry

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
	storeDockerRegistryRequest := *openapiclient.NewStoreDockerRegistryRequest(openapiclient.DockerRegistryType("default"), "Name_example") // StoreDockerRegistryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DockerAPI.CreateDockerRegistry(context.Background()).StoreDockerRegistryRequest(storeDockerRegistryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerAPI.CreateDockerRegistry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDockerRegistry`: DockerRegistry
	fmt.Fprintf(os.Stdout, "Response from `DockerAPI.CreateDockerRegistry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDockerRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeDockerRegistryRequest** | [**StoreDockerRegistryRequest**](StoreDockerRegistryRequest.md) |  | 

### Return type

[**DockerRegistry**](DockerRegistry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDockerRegistry

> DeleteDockerRegistry(ctx, dockerRegistry).Execute()

Delete a specific docker registry

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
	dockerRegistry := int32(56) // int32 | The docker registry ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DockerAPI.DeleteDockerRegistry(context.Background(), dockerRegistry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerAPI.DeleteDockerRegistry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dockerRegistry** | **int32** | The docker registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDockerRegistryRequest struct via the builder pattern


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
	resp, r, err := apiClient.DockerAPI.DockerServicesMetadataDeleteAll(context.Background(), dockerService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerAPI.DockerServicesMetadataDeleteAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DockerServicesMetadataDeleteAll`: Server
	fmt.Fprintf(os.Stdout, "Response from `DockerAPI.DockerServicesMetadataDeleteAll`: %v\n", resp)
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
	resp, r, err := apiClient.DockerAPI.DockerServicesMetadataDeleteKeys(context.Background(), dockerService).Metadata(metadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerAPI.DockerServicesMetadataDeleteKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DockerServicesMetadataDeleteKeys`: Server
	fmt.Fprintf(os.Stdout, "Response from `DockerAPI.DockerServicesMetadataDeleteKeys`: %v\n", resp)
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
	resp, r, err := apiClient.DockerAPI.DockerServicesMetadataSet(context.Background(), dockerService).SetMetadataRequest(setMetadataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerAPI.DockerServicesMetadataSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DockerServicesMetadataSet`: Server
	fmt.Fprintf(os.Stdout, "Response from `DockerAPI.DockerServicesMetadataSet`: %v\n", resp)
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
	resp, r, err := apiClient.DockerAPI.DockerServicesMetadataUpdate(context.Background(), dockerService).PatchMetadataRequest(patchMetadataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerAPI.DockerServicesMetadataUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DockerServicesMetadataUpdate`: Server
	fmt.Fprintf(os.Stdout, "Response from `DockerAPI.DockerServicesMetadataUpdate`: %v\n", resp)
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
	resp, r, err := apiClient.DockerAPI.DownloadServerLogs(context.Background(), dockerService).StreamSource(streamSource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerAPI.DownloadServerLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadServerLogs`: ServiceLogs
	fmt.Fprintf(os.Stdout, "Response from `DockerAPI.DownloadServerLogs`: %v\n", resp)
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
	resp, r, err := apiClient.DockerAPI.GetBackups(context.Background(), dockerService).PerPage(perPage).Page(page).Sort(sort).FilterName(filterName).FilterArchiveName(filterArchiveName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerAPI.GetBackups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackups`: GetBackups200Response
	fmt.Fprintf(os.Stdout, "Response from `DockerAPI.GetBackups`: %v\n", resp)
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


## GetDockerRegistries

> GetDockerRegistries200Response GetDockerRegistries(ctx).PerPage(perPage).Page(page).Sort(sort).FilterId(filterId).FilterType(filterType).FilterName(filterName).FilterNamePartial(filterNamePartial).FilterUrl(filterUrl).FilterOrganization(filterOrganization).Execute()

Show all docker registries

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
	perPage := int32(56) // int32 | The number of items to be shown per page. (optional)
	page := int32(56) // int32 | Specifies the page number to retrieve in the paginated results. (optional)
	sort := []string{"Inner_example"} // []string | Allows sorting of results. By default, sorting is in ascending order. To reverse the order, prepend the sort key with a hyphen (-).  **Simple Sort:** To sort by id in ascending order or by name in descending order:  ``` sort[]=id sort[]=-name ```  **Multiple Sorts:** Combine multiple sorts by separating them with commas: ``` sort[]=id&sort[]=-name ``` (optional)
	filterId := int32(56) // int32 | Filter by id. (optional)
	filterType := "filterType_example" // string | Filter by type. (optional)
	filterName := "filterName_example" // string | Filter by name. (optional)
	filterNamePartial := "filterNamePartial_example" // string | Filter by name using partial matching. For example, \"ann\" matches \"Joanna\" or \"Annie\". (optional)
	filterUrl := "filterUrl_example" // string | Filter by url. (optional)
	filterOrganization := "filterOrganization_example" // string | Filter by organization. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DockerAPI.GetDockerRegistries(context.Background()).PerPage(perPage).Page(page).Sort(sort).FilterId(filterId).FilterType(filterType).FilterName(filterName).FilterNamePartial(filterNamePartial).FilterUrl(filterUrl).FilterOrganization(filterOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerAPI.GetDockerRegistries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDockerRegistries`: GetDockerRegistries200Response
	fmt.Fprintf(os.Stdout, "Response from `DockerAPI.GetDockerRegistries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDockerRegistriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | The number of items to be shown per page. | 
 **page** | **int32** | Specifies the page number to retrieve in the paginated results. | 
 **sort** | **[]string** | Allows sorting of results. By default, sorting is in ascending order. To reverse the order, prepend the sort key with a hyphen (-).  **Simple Sort:** To sort by id in ascending order or by name in descending order:  &#x60;&#x60;&#x60; sort[]&#x3D;id sort[]&#x3D;-name &#x60;&#x60;&#x60;  **Multiple Sorts:** Combine multiple sorts by separating them with commas: &#x60;&#x60;&#x60; sort[]&#x3D;id&amp;sort[]&#x3D;-name &#x60;&#x60;&#x60; | 
 **filterId** | **int32** | Filter by id. | 
 **filterType** | **string** | Filter by type. | 
 **filterName** | **string** | Filter by name. | 
 **filterNamePartial** | **string** | Filter by name using partial matching. For example, \&quot;ann\&quot; matches \&quot;Joanna\&quot; or \&quot;Annie\&quot;. | 
 **filterUrl** | **string** | Filter by url. | 
 **filterOrganization** | **string** | Filter by organization. | 

### Return type

[**GetDockerRegistries200Response**](GetDockerRegistries200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDockerRegistryById

> DockerRegistry GetDockerRegistryById(ctx, dockerRegistry).Execute()

Display a specific docker registry

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
	dockerRegistry := int32(56) // int32 | The docker registry ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DockerAPI.GetDockerRegistryById(context.Background(), dockerRegistry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerAPI.GetDockerRegistryById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDockerRegistryById`: DockerRegistry
	fmt.Fprintf(os.Stdout, "Response from `DockerAPI.GetDockerRegistryById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dockerRegistry** | **int32** | The docker registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDockerRegistryByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DockerRegistry**](DockerRegistry.md)

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
	resp, r, err := apiClient.DockerAPI.GetLatestBackup(context.Background(), dockerService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerAPI.GetLatestBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLatestBackup`: Backup
	fmt.Fprintf(os.Stdout, "Response from `DockerAPI.GetLatestBackup`: %v\n", resp)
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
	resp, r, err := apiClient.DockerAPI.GetServerBackupDownloadUrl(context.Background(), dockerService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerAPI.GetServerBackupDownloadUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerBackupDownloadUrl`: BackupDownload
	fmt.Fprintf(os.Stdout, "Response from `DockerAPI.GetServerBackupDownloadUrl`: %v\n", resp)
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
	resp, r, err := apiClient.DockerAPI.GetServerById(context.Background(), app, dockerService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerAPI.GetServerById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerById`: Server
	fmt.Fprintf(os.Stdout, "Response from `DockerAPI.GetServerById`: %v\n", resp)
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
	resp, r, err := apiClient.DockerAPI.GetServerLogs(context.Background(), dockerService).Limit(limit).Direction(direction).StreamSource(streamSource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerAPI.GetServerLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerLogs`: ServiceLogs
	fmt.Fprintf(os.Stdout, "Response from `DockerAPI.GetServerLogs`: %v\n", resp)
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
	resp, r, err := apiClient.DockerAPI.GetServers(context.Background(), app).PerPage(perPage).Page(page).FilterStatus(filterStatus).FilterAppLocationSettingId(filterAppLocationSettingId).FilterServerConfigId(filterServerConfigId).FilterServerConfigName(filterServerConfigName).FilterServerConfigNamePartial(filterServerConfigNamePartial).FilterLocationCity(filterLocationCity).FilterLocationCityDisplay(filterLocationCityDisplay).FilterLocationContinent(filterLocationContinent).FilterLocationCountry(filterLocationCountry).FilterIsBackupable(filterIsBackupable).FilterIsRestorable(filterIsRestorable).FilterIsPending(filterIsPending).FilterIsNotFound(filterIsNotFound).FilterIsHealthy(filterIsHealthy).FilterBinaryId(filterBinaryId).FilterIsStopped(filterIsStopped).FilterMetadata(filterMetadata).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerAPI.GetServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServers`: GetServers200Response
	fmt.Fprintf(os.Stdout, "Response from `DockerAPI.GetServers`: %v\n", resp)
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


## GetTaggedImages

> GetTaggedImages200Response GetTaggedImages(ctx, dockerRegistry).Execute()

List all available tagged images

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
	dockerRegistry := int32(56) // int32 | The docker registry ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DockerAPI.GetTaggedImages(context.Background(), dockerRegistry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerAPI.GetTaggedImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaggedImages`: GetTaggedImages200Response
	fmt.Fprintf(os.Stdout, "Response from `DockerAPI.GetTaggedImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dockerRegistry** | **int32** | The docker registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaggedImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetTaggedImages200Response**](GetTaggedImages200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshTaggedImages

> GetTaggedImages200Response RefreshTaggedImages(ctx, dockerRegistry).Execute()

Refresh the cache for all available tagged images

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
	dockerRegistry := int32(56) // int32 | The docker registry ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DockerAPI.RefreshTaggedImages(context.Background(), dockerRegistry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerAPI.RefreshTaggedImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshTaggedImages`: GetTaggedImages200Response
	fmt.Fprintf(os.Stdout, "Response from `DockerAPI.RefreshTaggedImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dockerRegistry** | **int32** | The docker registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshTaggedImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetTaggedImages200Response**](GetTaggedImages200Response.md)

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
	r, err := apiClient.DockerAPI.RestartServer(context.Background(), dockerService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerAPI.RestartServer``: %v\n", err)
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
	r, err := apiClient.DockerAPI.RestoreBackup(context.Background(), dockerService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerAPI.RestoreBackup``: %v\n", err)
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
	r, err := apiClient.DockerAPI.StartServer(context.Background(), dockerService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerAPI.StartServer``: %v\n", err)
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
	r, err := apiClient.DockerAPI.StopServer(context.Background(), dockerService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerAPI.StopServer``: %v\n", err)
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


## UpdateDockerRegistry

> DockerRegistry UpdateDockerRegistry(ctx, dockerRegistry).UpdateDockerRegistryRequest(updateDockerRegistryRequest).Execute()

Update a specific docker registry

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
	dockerRegistry := int32(56) // int32 | The docker registry ID
	updateDockerRegistryRequest := *openapiclient.NewUpdateDockerRegistryRequest(openapiclient.DockerRegistryType("default"), "Name_example") // UpdateDockerRegistryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DockerAPI.UpdateDockerRegistry(context.Background(), dockerRegistry).UpdateDockerRegistryRequest(updateDockerRegistryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerAPI.UpdateDockerRegistry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDockerRegistry`: DockerRegistry
	fmt.Fprintf(os.Stdout, "Response from `DockerAPI.UpdateDockerRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dockerRegistry** | **int32** | The docker registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDockerRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDockerRegistryRequest** | [**UpdateDockerRegistryRequest**](UpdateDockerRegistryRequest.md) |  | 

### Return type

[**DockerRegistry**](DockerRegistry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
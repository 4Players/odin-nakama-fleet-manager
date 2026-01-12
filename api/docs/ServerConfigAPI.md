# \ServerConfigAPI

All URIs are relative to *https://fleet.4players.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServerConfig**](ServerConfigAPI.md#CreateServerConfig) | **Post** /v1/apps/{app}/configs | Create a new server config
[**DeleteServerConfig**](ServerConfigAPI.md#DeleteServerConfig) | **Delete** /v1/server-configs/{serverConfig} | Delete a specific server config
[**DeleteUnusedServerConfigs**](ServerConfigAPI.md#DeleteUnusedServerConfigs) | **Delete** /v1/apps/{app}/configs/unused | Delete all unused server configs
[**GetServerConfigById**](ServerConfigAPI.md#GetServerConfigById) | **Get** /v1/server-configs/{serverConfig} | Show a specific server config
[**GetServerConfigs**](ServerConfigAPI.md#GetServerConfigs) | **Get** /v1/apps/{app}/configs | Show all server configs
[**StartServersForServerConfig**](ServerConfigAPI.md#StartServersForServerConfig) | **Post** /v1/server-configs/{serverConfig}/services/start | Start all services related to a specific server config
[**StopServersForServerConfig**](ServerConfigAPI.md#StopServersForServerConfig) | **Post** /v1/server-configs/{serverConfig}/services/stop | Stop all services related to a specific server config
[**UpdateServerConfig**](ServerConfigAPI.md#UpdateServerConfig) | **Put** /v1/server-configs/{serverConfig} | Update a server config



## CreateServerConfig

> ServerConfig CreateServerConfig(ctx, app).StoreServerConfigRequest(storeServerConfigRequest).Execute()

Create a new server config

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
	storeServerConfigRequest := *openapiclient.NewStoreServerConfigRequest("Name_example", int32(123), "ResourcePackageSlug_example") // StoreServerConfigRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerConfigAPI.CreateServerConfig(context.Background(), app).StoreServerConfigRequest(storeServerConfigRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerConfigAPI.CreateServerConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServerConfig`: ServerConfig
	fmt.Fprintf(os.Stdout, "Response from `ServerConfigAPI.CreateServerConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **int32** | The app ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storeServerConfigRequest** | [**StoreServerConfigRequest**](StoreServerConfigRequest.md) |  | 

### Return type

[**ServerConfig**](ServerConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServerConfig

> DeleteServerConfig(ctx, serverConfig).Execute()

Delete a specific server config

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
	serverConfig := int32(56) // int32 | The server config ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerConfigAPI.DeleteServerConfig(context.Background(), serverConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerConfigAPI.DeleteServerConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverConfig** | **int32** | The server config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerConfigRequest struct via the builder pattern


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


## DeleteUnusedServerConfigs

> DeleteUnusedServerConfigs(ctx, app).Execute()

Delete all unused server configs

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerConfigAPI.DeleteUnusedServerConfigs(context.Background(), app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerConfigAPI.DeleteUnusedServerConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **int32** | The app ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUnusedServerConfigsRequest struct via the builder pattern


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


## GetServerConfigById

> ServerConfig GetServerConfigById(ctx, serverConfig).Execute()

Show a specific server config

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
	serverConfig := int32(56) // int32 | The server config ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerConfigAPI.GetServerConfigById(context.Background(), serverConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerConfigAPI.GetServerConfigById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerConfigById`: ServerConfig
	fmt.Fprintf(os.Stdout, "Response from `ServerConfigAPI.GetServerConfigById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverConfig** | **int32** | The server config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerConfigByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerConfig**](ServerConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerConfigs

> GetServerConfigs200Response GetServerConfigs(ctx, app).PerPage(perPage).Page(page).Sort(sort).FilterId(filterId).FilterBinaryId(filterBinaryId).FilterName(filterName).FilterNamePartial(filterNamePartial).FilterCommand(filterCommand).FilterArgs(filterArgs).FilterNotes(filterNotes).FilterStatus(filterStatus).FilterMaintenance(filterMaintenance).FilterResourcePackageSlug(filterResourcePackageSlug).FilterInUse(filterInUse).FilterBinaryName(filterBinaryName).FilterBinaryVersion(filterBinaryVersion).FilterBinaryType(filterBinaryType).FilterBinaryOs(filterBinaryOs).Execute()

Show all server configs

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
	sort := []string{"Inner_example"} // []string | Allows sorting of results. By default, sorting is in ascending order. To reverse the order, prepend the sort key with a hyphen (-).  **Simple Sort:** To sort by id in ascending order or by name in descending order:  ``` sort[]=id sort[]=-name ```  **Multiple Sorts:** Combine multiple sorts by separating them with commas: ``` sort[]=id&sort[]=-name ``` (optional)
	filterId := int32(56) // int32 | Filter by id. (optional)
	filterBinaryId := int32(56) // int32 | Filter by binary id. (optional)
	filterName := "filterName_example" // string | Filter by name. (optional)
	filterNamePartial := "filterNamePartial_example" // string | Filter by name using partial matching. For example, \"ann\" matches \"Joanna\" or \"Annie\". (optional)
	filterCommand := "filterCommand_example" // string | Filter by command. (optional)
	filterArgs := "filterArgs_example" // string | Filter by arguments. (optional)
	filterNotes := "filterNotes_example" // string | Filter by notes. (optional)
	filterStatus := "filterStatus_example" // string | Filter by status. (optional)
	filterMaintenance := true // bool | Filter by maintenance status. (optional)
	filterResourcePackageSlug := "filterResourcePackageSlug_example" // string | Filter by resource package slug. (optional)
	filterInUse := true // bool | Filter by in use flag. (optional)
	filterBinaryName := "filterBinaryName_example" // string | Filter by binary name. (optional)
	filterBinaryVersion := "filterBinaryVersion_example" // string | Filter by binary version. (optional)
	filterBinaryType := "filterBinaryType_example" // string | Filter by binary type. (optional)
	filterBinaryOs := "filterBinaryOs_example" // string | Filter by binary operating system. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerConfigAPI.GetServerConfigs(context.Background(), app).PerPage(perPage).Page(page).Sort(sort).FilterId(filterId).FilterBinaryId(filterBinaryId).FilterName(filterName).FilterNamePartial(filterNamePartial).FilterCommand(filterCommand).FilterArgs(filterArgs).FilterNotes(filterNotes).FilterStatus(filterStatus).FilterMaintenance(filterMaintenance).FilterResourcePackageSlug(filterResourcePackageSlug).FilterInUse(filterInUse).FilterBinaryName(filterBinaryName).FilterBinaryVersion(filterBinaryVersion).FilterBinaryType(filterBinaryType).FilterBinaryOs(filterBinaryOs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerConfigAPI.GetServerConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerConfigs`: GetServerConfigs200Response
	fmt.Fprintf(os.Stdout, "Response from `ServerConfigAPI.GetServerConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **int32** | The app ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of items to be shown per page. | 
 **page** | **int32** | Specifies the page number to retrieve in the paginated results. | 
 **sort** | **[]string** | Allows sorting of results. By default, sorting is in ascending order. To reverse the order, prepend the sort key with a hyphen (-).  **Simple Sort:** To sort by id in ascending order or by name in descending order:  &#x60;&#x60;&#x60; sort[]&#x3D;id sort[]&#x3D;-name &#x60;&#x60;&#x60;  **Multiple Sorts:** Combine multiple sorts by separating them with commas: &#x60;&#x60;&#x60; sort[]&#x3D;id&amp;sort[]&#x3D;-name &#x60;&#x60;&#x60; | 
 **filterId** | **int32** | Filter by id. | 
 **filterBinaryId** | **int32** | Filter by binary id. | 
 **filterName** | **string** | Filter by name. | 
 **filterNamePartial** | **string** | Filter by name using partial matching. For example, \&quot;ann\&quot; matches \&quot;Joanna\&quot; or \&quot;Annie\&quot;. | 
 **filterCommand** | **string** | Filter by command. | 
 **filterArgs** | **string** | Filter by arguments. | 
 **filterNotes** | **string** | Filter by notes. | 
 **filterStatus** | **string** | Filter by status. | 
 **filterMaintenance** | **bool** | Filter by maintenance status. | 
 **filterResourcePackageSlug** | **string** | Filter by resource package slug. | 
 **filterInUse** | **bool** | Filter by in use flag. | 
 **filterBinaryName** | **string** | Filter by binary name. | 
 **filterBinaryVersion** | **string** | Filter by binary version. | 
 **filterBinaryType** | **string** | Filter by binary type. | 
 **filterBinaryOs** | **string** | Filter by binary operating system. | 

### Return type

[**GetServerConfigs200Response**](GetServerConfigs200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartServersForServerConfig

> StartServersForServerConfig(ctx, serverConfig).Execute()

Start all services related to a specific server config

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
	serverConfig := int32(56) // int32 | The server config ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerConfigAPI.StartServersForServerConfig(context.Background(), serverConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerConfigAPI.StartServersForServerConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverConfig** | **int32** | The server config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartServersForServerConfigRequest struct via the builder pattern


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


## StopServersForServerConfig

> StopServersForServerConfig(ctx, serverConfig).Execute()

Stop all services related to a specific server config

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
	serverConfig := int32(56) // int32 | The server config ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerConfigAPI.StopServersForServerConfig(context.Background(), serverConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerConfigAPI.StopServersForServerConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverConfig** | **int32** | The server config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopServersForServerConfigRequest struct via the builder pattern


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


## UpdateServerConfig

> ServerConfig UpdateServerConfig(ctx, serverConfig).UpdateServerConfigRequest(updateServerConfigRequest).Execute()

Update a server config

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
	serverConfig := int32(56) // int32 | The server config ID
	updateServerConfigRequest := *openapiclient.NewUpdateServerConfigRequest("Name_example", int32(123), "ResourcePackageSlug_example") // UpdateServerConfigRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerConfigAPI.UpdateServerConfig(context.Background(), serverConfig).UpdateServerConfigRequest(updateServerConfigRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerConfigAPI.UpdateServerConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServerConfig`: ServerConfig
	fmt.Fprintf(os.Stdout, "Response from `ServerConfigAPI.UpdateServerConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverConfig** | **int32** | The server config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateServerConfigRequest** | [**UpdateServerConfigRequest**](UpdateServerConfigRequest.md) |  | 

### Return type

[**ServerConfig**](ServerConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
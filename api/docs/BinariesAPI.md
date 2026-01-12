# \BinariesAPI

All URIs are relative to *https://fleet.4players.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBinary**](BinariesAPI.md#CreateBinary) | **Post** /v1/apps/{app}/binaries | Create a binary and the related entity
[**DeleteBinary**](BinariesAPI.md#DeleteBinary) | **Delete** /v1/binaries/{binary} | Delete a specified binary
[**DeleteUnusedBinaries**](BinariesAPI.md#DeleteUnusedBinaries) | **Delete** /v1/apps/{app}/binaries/unused | Delete all unused binaries
[**GetBinaries**](BinariesAPI.md#GetBinaries) | **Get** /v1/apps/{app}/binaries | Show all binaries
[**GetBinaryById**](BinariesAPI.md#GetBinaryById) | **Get** /v1/binaries/{binary} | Show a specific binary
[**RefreshBinary**](BinariesAPI.md#RefreshBinary) | **Put** /v1/binaries/{binary}/refresh | Refresh a binary and the related entity
[**StartServersForBinary**](BinariesAPI.md#StartServersForBinary) | **Post** /v1/binaries/{binary}/services/start | Start all services related to a specific binary
[**StopServersForBinary**](BinariesAPI.md#StopServersForBinary) | **Post** /v1/binaries/{binary}/services/stop | Stop all services related to a specific binary
[**UpdateBinary**](BinariesAPI.md#UpdateBinary) | **Put** /v1/binaries/{binary} | Update a binary and the related entity



## CreateBinary

> Binary CreateBinary(ctx, app).StoreBinaryRequest(storeBinaryRequest).Execute()

Create a binary and the related entity

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
	storeBinaryRequest := *openapiclient.NewStoreBinaryRequest("Name_example", "Version_example", openapiclient.BinaryType("dockerImage"), openapiclient.OperatingSystem("windows")) // StoreBinaryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinariesAPI.CreateBinary(context.Background(), app).StoreBinaryRequest(storeBinaryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinariesAPI.CreateBinary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBinary`: Binary
	fmt.Fprintf(os.Stdout, "Response from `BinariesAPI.CreateBinary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **int32** | The app ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBinaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storeBinaryRequest** | [**StoreBinaryRequest**](StoreBinaryRequest.md) |  | 

### Return type

[**Binary**](Binary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBinary

> DeleteBinary(ctx, binary).Execute()

Delete a specified binary

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
	binary := int32(56) // int32 | The binary ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BinariesAPI.DeleteBinary(context.Background(), binary).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinariesAPI.DeleteBinary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**binary** | **int32** | The binary ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBinaryRequest struct via the builder pattern


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


## DeleteUnusedBinaries

> DeleteUnusedBinaries(ctx, app).Execute()

Delete all unused binaries

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
	r, err := apiClient.BinariesAPI.DeleteUnusedBinaries(context.Background(), app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinariesAPI.DeleteUnusedBinaries``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteUnusedBinariesRequest struct via the builder pattern


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


## GetBinaries

> GetBinaries200Response GetBinaries(ctx, app).PerPage(perPage).Page(page).Sort(sort).FilterName(filterName).FilterNamePartial(filterNamePartial).FilterVersion(filterVersion).FilterType(filterType).FilterOs(filterOs).FilterMaintenance(filterMaintenance).FilterStatus(filterStatus).FilterInUse(filterInUse).Execute()

Show all binaries

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
	filterName := "filterName_example" // string | Filter by name. (optional)
	filterNamePartial := "filterNamePartial_example" // string | Filter by name using partial matching. For example, \"ann\" matches \"Joanna\" or \"Annie\". (optional)
	filterVersion := "filterVersion_example" // string | Filter by version. (optional)
	filterType := "filterType_example" // string | Filter by type. (optional)
	filterOs := "filterOs_example" // string | Filter by operating system. (optional)
	filterMaintenance := true // bool | Filter by maintenance status. (optional)
	filterStatus := "filterStatus_example" // string | Filter by status. (optional)
	filterInUse := true // bool | Filter by in use flag. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinariesAPI.GetBinaries(context.Background(), app).PerPage(perPage).Page(page).Sort(sort).FilterName(filterName).FilterNamePartial(filterNamePartial).FilterVersion(filterVersion).FilterType(filterType).FilterOs(filterOs).FilterMaintenance(filterMaintenance).FilterStatus(filterStatus).FilterInUse(filterInUse).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinariesAPI.GetBinaries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBinaries`: GetBinaries200Response
	fmt.Fprintf(os.Stdout, "Response from `BinariesAPI.GetBinaries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **int32** | The app ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBinariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of items to be shown per page. | 
 **page** | **int32** | Specifies the page number to retrieve in the paginated results. | 
 **sort** | **[]string** | Allows sorting of results. By default, sorting is in ascending order. To reverse the order, prepend the sort key with a hyphen (-).  **Simple Sort:** To sort by id in ascending order or by name in descending order:  &#x60;&#x60;&#x60; sort[]&#x3D;id sort[]&#x3D;-name &#x60;&#x60;&#x60;  **Multiple Sorts:** Combine multiple sorts by separating them with commas: &#x60;&#x60;&#x60; sort[]&#x3D;id&amp;sort[]&#x3D;-name &#x60;&#x60;&#x60; | 
 **filterName** | **string** | Filter by name. | 
 **filterNamePartial** | **string** | Filter by name using partial matching. For example, \&quot;ann\&quot; matches \&quot;Joanna\&quot; or \&quot;Annie\&quot;. | 
 **filterVersion** | **string** | Filter by version. | 
 **filterType** | **string** | Filter by type. | 
 **filterOs** | **string** | Filter by operating system. | 
 **filterMaintenance** | **bool** | Filter by maintenance status. | 
 **filterStatus** | **string** | Filter by status. | 
 **filterInUse** | **bool** | Filter by in use flag. | 

### Return type

[**GetBinaries200Response**](GetBinaries200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBinaryById

> Binary GetBinaryById(ctx, binary).Execute()

Show a specific binary

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
	binary := int32(56) // int32 | The binary ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinariesAPI.GetBinaryById(context.Background(), binary).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinariesAPI.GetBinaryById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBinaryById`: Binary
	fmt.Fprintf(os.Stdout, "Response from `BinariesAPI.GetBinaryById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**binary** | **int32** | The binary ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBinaryByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Binary**](Binary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshBinary

> Binary RefreshBinary(ctx, binary).Execute()

Refresh a binary and the related entity

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
	binary := int32(56) // int32 | The binary ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinariesAPI.RefreshBinary(context.Background(), binary).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinariesAPI.RefreshBinary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshBinary`: Binary
	fmt.Fprintf(os.Stdout, "Response from `BinariesAPI.RefreshBinary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**binary** | **int32** | The binary ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshBinaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Binary**](Binary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartServersForBinary

> StartServersForBinary(ctx, binary).Execute()

Start all services related to a specific binary

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
	binary := int32(56) // int32 | The binary ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BinariesAPI.StartServersForBinary(context.Background(), binary).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinariesAPI.StartServersForBinary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**binary** | **int32** | The binary ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartServersForBinaryRequest struct via the builder pattern


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


## StopServersForBinary

> StopServersForBinary(ctx, binary).Execute()

Stop all services related to a specific binary

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
	binary := int32(56) // int32 | The binary ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BinariesAPI.StopServersForBinary(context.Background(), binary).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinariesAPI.StopServersForBinary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**binary** | **int32** | The binary ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopServersForBinaryRequest struct via the builder pattern


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


## UpdateBinary

> Binary UpdateBinary(ctx, binary).UpdateBinaryRequest(updateBinaryRequest).Execute()

Update a binary and the related entity

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
	binary := int32(56) // int32 | The binary ID
	updateBinaryRequest := *openapiclient.NewUpdateBinaryRequest("Name_example", "Version_example", openapiclient.BinaryType("dockerImage"), openapiclient.OperatingSystem("windows")) // UpdateBinaryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinariesAPI.UpdateBinary(context.Background(), binary).UpdateBinaryRequest(updateBinaryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinariesAPI.UpdateBinary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBinary`: Binary
	fmt.Fprintf(os.Stdout, "Response from `BinariesAPI.UpdateBinary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**binary** | **int32** | The binary ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBinaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateBinaryRequest** | [**UpdateBinaryRequest**](UpdateBinaryRequest.md) |  | 

### Return type

[**Binary**](Binary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
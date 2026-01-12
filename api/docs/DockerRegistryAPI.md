# \DockerRegistryAPI

All URIs are relative to *https://fleet.4players.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDockerRegistry**](DockerRegistryAPI.md#CreateDockerRegistry) | **Post** /v1/docker-registries | Create a new docker registry
[**DeleteDockerRegistry**](DockerRegistryAPI.md#DeleteDockerRegistry) | **Delete** /v1/docker-registries/{dockerRegistry} | Delete a specific docker registry
[**GetDockerRegistries**](DockerRegistryAPI.md#GetDockerRegistries) | **Get** /v1/docker-registries | Show all docker registries
[**GetDockerRegistryById**](DockerRegistryAPI.md#GetDockerRegistryById) | **Get** /v1/docker-registries/{dockerRegistry} | Display a specific docker registry
[**GetTaggedImages**](DockerRegistryAPI.md#GetTaggedImages) | **Get** /v1/docker-registries/{dockerRegistry}/tagged-images | List all available tagged images
[**RefreshTaggedImages**](DockerRegistryAPI.md#RefreshTaggedImages) | **Get** /v1/docker-registries/{dockerRegistry}/tagged-images/refresh | Refresh the cache for all available tagged images
[**UpdateDockerRegistry**](DockerRegistryAPI.md#UpdateDockerRegistry) | **Put** /v1/docker-registries/{dockerRegistry} | Update a specific docker registry



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
	resp, r, err := apiClient.DockerRegistryAPI.CreateDockerRegistry(context.Background()).StoreDockerRegistryRequest(storeDockerRegistryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerRegistryAPI.CreateDockerRegistry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDockerRegistry`: DockerRegistry
	fmt.Fprintf(os.Stdout, "Response from `DockerRegistryAPI.CreateDockerRegistry`: %v\n", resp)
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
	r, err := apiClient.DockerRegistryAPI.DeleteDockerRegistry(context.Background(), dockerRegistry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerRegistryAPI.DeleteDockerRegistry``: %v\n", err)
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
	resp, r, err := apiClient.DockerRegistryAPI.GetDockerRegistries(context.Background()).PerPage(perPage).Page(page).Sort(sort).FilterId(filterId).FilterType(filterType).FilterName(filterName).FilterNamePartial(filterNamePartial).FilterUrl(filterUrl).FilterOrganization(filterOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerRegistryAPI.GetDockerRegistries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDockerRegistries`: GetDockerRegistries200Response
	fmt.Fprintf(os.Stdout, "Response from `DockerRegistryAPI.GetDockerRegistries`: %v\n", resp)
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
	resp, r, err := apiClient.DockerRegistryAPI.GetDockerRegistryById(context.Background(), dockerRegistry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerRegistryAPI.GetDockerRegistryById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDockerRegistryById`: DockerRegistry
	fmt.Fprintf(os.Stdout, "Response from `DockerRegistryAPI.GetDockerRegistryById`: %v\n", resp)
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
	resp, r, err := apiClient.DockerRegistryAPI.GetTaggedImages(context.Background(), dockerRegistry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerRegistryAPI.GetTaggedImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaggedImages`: GetTaggedImages200Response
	fmt.Fprintf(os.Stdout, "Response from `DockerRegistryAPI.GetTaggedImages`: %v\n", resp)
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
	resp, r, err := apiClient.DockerRegistryAPI.RefreshTaggedImages(context.Background(), dockerRegistry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerRegistryAPI.RefreshTaggedImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshTaggedImages`: GetTaggedImages200Response
	fmt.Fprintf(os.Stdout, "Response from `DockerRegistryAPI.RefreshTaggedImages`: %v\n", resp)
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
	resp, r, err := apiClient.DockerRegistryAPI.UpdateDockerRegistry(context.Background(), dockerRegistry).UpdateDockerRegistryRequest(updateDockerRegistryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockerRegistryAPI.UpdateDockerRegistry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDockerRegistry`: DockerRegistry
	fmt.Fprintf(os.Stdout, "Response from `DockerRegistryAPI.UpdateDockerRegistry`: %v\n", resp)
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
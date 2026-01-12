# \ResourcePackageAPI

All URIs are relative to *https://fleet.4players.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetResourcePackageById**](ResourcePackageAPI.md#GetResourcePackageById) | **Get** /v1/resource-packages/{resourcePackage} | Show a specified resource package
[**GetResourcePackages**](ResourcePackageAPI.md#GetResourcePackages) | **Get** /v1/resource-packages | Show all available resource packages



## GetResourcePackageById

> ResourcePackage GetResourcePackageById(ctx, resourcePackage).Execute()

Show a specified resource package

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
	resourcePackage := int32(56) // int32 | The resource package ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcePackageAPI.GetResourcePackageById(context.Background(), resourcePackage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcePackageAPI.GetResourcePackageById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResourcePackageById`: ResourcePackage
	fmt.Fprintf(os.Stdout, "Response from `ResourcePackageAPI.GetResourcePackageById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourcePackage** | **int32** | The resource package ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourcePackageByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResourcePackage**](ResourcePackage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourcePackages

> GetResourcePackages200Response GetResourcePackages(ctx).PerPage(perPage).Page(page).Sort(sort).FilterId(filterId).FilterName(filterName).FilterNamePartial(filterNamePartial).FilterSlug(filterSlug).FilterType(filterType).FilterCpuLimit(filterCpuLimit).FilterMemoryLimitMiB(filterMemoryLimitMiB).Execute()

Show all available resource packages

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
	filterName := "filterName_example" // string | Filter by name. (optional)
	filterNamePartial := "filterNamePartial_example" // string | Filter by name using partial matching. For example, \"ann\" matches \"Joanna\" or \"Annie\". (optional)
	filterSlug := "filterSlug_example" // string | Filter by slug. (optional)
	filterType := "filterType_example" // string | Filter by type. (optional)
	filterCpuLimit := int32(56) // int32 | Filter by CPU limit. Maps to the `cpu_limit` column. (optional)
	filterMemoryLimitMiB := int32(56) // int32 | Filter by memory limit in MiB. Maps to the `memory_limit_mebibytes` column. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcePackageAPI.GetResourcePackages(context.Background()).PerPage(perPage).Page(page).Sort(sort).FilterId(filterId).FilterName(filterName).FilterNamePartial(filterNamePartial).FilterSlug(filterSlug).FilterType(filterType).FilterCpuLimit(filterCpuLimit).FilterMemoryLimitMiB(filterMemoryLimitMiB).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcePackageAPI.GetResourcePackages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResourcePackages`: GetResourcePackages200Response
	fmt.Fprintf(os.Stdout, "Response from `ResourcePackageAPI.GetResourcePackages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetResourcePackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | The number of items to be shown per page. | 
 **page** | **int32** | Specifies the page number to retrieve in the paginated results. | 
 **sort** | **[]string** | Allows sorting of results. By default, sorting is in ascending order. To reverse the order, prepend the sort key with a hyphen (-).  **Simple Sort:** To sort by id in ascending order or by name in descending order:  &#x60;&#x60;&#x60; sort[]&#x3D;id sort[]&#x3D;-name &#x60;&#x60;&#x60;  **Multiple Sorts:** Combine multiple sorts by separating them with commas: &#x60;&#x60;&#x60; sort[]&#x3D;id&amp;sort[]&#x3D;-name &#x60;&#x60;&#x60; | 
 **filterId** | **int32** | Filter by id. | 
 **filterName** | **string** | Filter by name. | 
 **filterNamePartial** | **string** | Filter by name using partial matching. For example, \&quot;ann\&quot; matches \&quot;Joanna\&quot; or \&quot;Annie\&quot;. | 
 **filterSlug** | **string** | Filter by slug. | 
 **filterType** | **string** | Filter by type. | 
 **filterCpuLimit** | **int32** | Filter by CPU limit. Maps to the &#x60;cpu_limit&#x60; column. | 
 **filterMemoryLimitMiB** | **int32** | Filter by memory limit in MiB. Maps to the &#x60;memory_limit_mebibytes&#x60; column. | 

### Return type

[**GetResourcePackages200Response**](GetResourcePackages200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
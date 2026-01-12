# \AppLocationSettingAPI

All URIs are relative to *https://fleet.4players.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAppLocationSetting**](AppLocationSettingAPI.md#CreateAppLocationSetting) | **Post** /v1/apps/{app}/location-settings | Create a new location setting
[**DeleteAppLocationSetting**](AppLocationSettingAPI.md#DeleteAppLocationSetting) | **Delete** /v1/app-location-settings/{appLocationSetting} | Delete a location setting
[**GetAppLocationSettingById**](AppLocationSettingAPI.md#GetAppLocationSettingById) | **Get** /v1/app-location-settings/{appLocationSetting} | Show a specific app location setting
[**GetAppLocationSettings**](AppLocationSettingAPI.md#GetAppLocationSettings) | **Get** /v1/apps/{app}/location-settings | Show all location settings
[**ListServicesForAppLocationSetting**](AppLocationSettingAPI.md#ListServicesForAppLocationSetting) | **Get** /v1/apps/{app}/location-settings/{appLocationSetting}/services | Show all services for a specific app location setting within a given app
[**StartServersForAppLocationSetting**](AppLocationSettingAPI.md#StartServersForAppLocationSetting) | **Post** /v1/app-location-settings/{appLocationSetting}/services/start | Start all services related to a specific app location setting
[**StopServersForAppLocationSetting**](AppLocationSettingAPI.md#StopServersForAppLocationSetting) | **Post** /v1/app-location-settings/{appLocationSetting}/services/stop | Stop all services related to a specific app location setting
[**UpdateAppLocationSetting**](AppLocationSettingAPI.md#UpdateAppLocationSetting) | **Put** /v1/app-location-settings/{appLocationSetting} | Update a location setting



## CreateAppLocationSetting

> AppLocationSetting CreateAppLocationSetting(ctx, app).StoreAppLocationSettingRequest(storeAppLocationSettingRequest).Execute()

Create a new location setting

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
	storeAppLocationSettingRequest := *openapiclient.NewStoreAppLocationSettingRequest("Name_example", int32(123), int32(123)) // StoreAppLocationSettingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppLocationSettingAPI.CreateAppLocationSetting(context.Background(), app).StoreAppLocationSettingRequest(storeAppLocationSettingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppLocationSettingAPI.CreateAppLocationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppLocationSetting`: AppLocationSetting
	fmt.Fprintf(os.Stdout, "Response from `AppLocationSettingAPI.CreateAppLocationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **int32** | The app ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppLocationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storeAppLocationSettingRequest** | [**StoreAppLocationSettingRequest**](StoreAppLocationSettingRequest.md) |  | 

### Return type

[**AppLocationSetting**](AppLocationSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAppLocationSetting

> DeleteAppLocationSetting(ctx, appLocationSetting).Execute()

Delete a location setting

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
	appLocationSetting := int32(56) // int32 | The app location setting ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppLocationSettingAPI.DeleteAppLocationSetting(context.Background(), appLocationSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppLocationSettingAPI.DeleteAppLocationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appLocationSetting** | **int32** | The app location setting ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppLocationSettingRequest struct via the builder pattern


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


## GetAppLocationSettingById

> AppLocationSetting GetAppLocationSettingById(ctx, appLocationSetting).Execute()

Show a specific app location setting

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
	appLocationSetting := int32(56) // int32 | The app location setting ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppLocationSettingAPI.GetAppLocationSettingById(context.Background(), appLocationSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppLocationSettingAPI.GetAppLocationSettingById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppLocationSettingById`: AppLocationSetting
	fmt.Fprintf(os.Stdout, "Response from `AppLocationSettingAPI.GetAppLocationSettingById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appLocationSetting** | **int32** | The app location setting ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppLocationSettingByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppLocationSetting**](AppLocationSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppLocationSettings

> GetAppLocationSettings200Response GetAppLocationSettings(ctx, app).PerPage(perPage).Page(page).Sort(sort).FilterId(filterId).FilterName(filterName).FilterNamePartial(filterNamePartial).FilterServerConfigId(filterServerConfigId).FilterNumInstances(filterNumInstances).FilterStatus(filterStatus).FilterMaintenance(filterMaintenance).FilterLocationCity(filterLocationCity).FilterLocationCityDisplay(filterLocationCityDisplay).FilterLocationContinent(filterLocationContinent).FilterLocationCountry(filterLocationCountry).FilterServerConfigName(filterServerConfigName).FilterServerConfigCommand(filterServerConfigCommand).FilterServerConfigArgs(filterServerConfigArgs).FilterServerConfigNotes(filterServerConfigNotes).FilterServerConfigStatus(filterServerConfigStatus).FilterServerConfigMaintenance(filterServerConfigMaintenance).FilterServerConfigResourcePackageSlug(filterServerConfigResourcePackageSlug).Execute()

Show all location settings

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
	filterName := "filterName_example" // string | Filter by name. (optional)
	filterNamePartial := "filterNamePartial_example" // string | Filter by name using partial matching. For example, \"ann\" matches \"Joanna\" or \"Annie\". (optional)
	filterServerConfigId := int32(56) // int32 | Filter by ServerConfig ID. (optional)
	filterNumInstances := int32(56) // int32 | Filter by number of instances. (optional)
	filterStatus := "filterStatus_example" // string | Filter by status. (optional)
	filterMaintenance := true // bool | Filter by maintenance. (optional)
	filterLocationCity := "filterLocationCity_example" // string | Filter by location city. (optional)
	filterLocationCityDisplay := "filterLocationCityDisplay_example" // string | Filter by location city display name. (optional)
	filterLocationContinent := "filterLocationContinent_example" // string | Filter by location continent. (optional)
	filterLocationCountry := "filterLocationCountry_example" // string | Filter by location country. (optional)
	filterServerConfigName := "filterServerConfigName_example" // string | Filter by ServerConfig name. (optional)
	filterServerConfigCommand := "filterServerConfigCommand_example" // string | Filter by ServerConfig command. (optional)
	filterServerConfigArgs := "filterServerConfigArgs_example" // string | Filter by ServerConfig arguments. (optional)
	filterServerConfigNotes := "filterServerConfigNotes_example" // string | Filter by ServerConfig notes. (optional)
	filterServerConfigStatus := "filterServerConfigStatus_example" // string | Filter by ServerConfig status. (optional)
	filterServerConfigMaintenance := true // bool | Filter by ServerConfig maintenance. (optional)
	filterServerConfigResourcePackageSlug := "filterServerConfigResourcePackageSlug_example" // string | Filter by ServerConfig resource package slug. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppLocationSettingAPI.GetAppLocationSettings(context.Background(), app).PerPage(perPage).Page(page).Sort(sort).FilterId(filterId).FilterName(filterName).FilterNamePartial(filterNamePartial).FilterServerConfigId(filterServerConfigId).FilterNumInstances(filterNumInstances).FilterStatus(filterStatus).FilterMaintenance(filterMaintenance).FilterLocationCity(filterLocationCity).FilterLocationCityDisplay(filterLocationCityDisplay).FilterLocationContinent(filterLocationContinent).FilterLocationCountry(filterLocationCountry).FilterServerConfigName(filterServerConfigName).FilterServerConfigCommand(filterServerConfigCommand).FilterServerConfigArgs(filterServerConfigArgs).FilterServerConfigNotes(filterServerConfigNotes).FilterServerConfigStatus(filterServerConfigStatus).FilterServerConfigMaintenance(filterServerConfigMaintenance).FilterServerConfigResourcePackageSlug(filterServerConfigResourcePackageSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppLocationSettingAPI.GetAppLocationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppLocationSettings`: GetAppLocationSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `AppLocationSettingAPI.GetAppLocationSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **int32** | The app ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppLocationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of items to be shown per page. | 
 **page** | **int32** | Specifies the page number to retrieve in the paginated results. | 
 **sort** | **[]string** | Allows sorting of results. By default, sorting is in ascending order. To reverse the order, prepend the sort key with a hyphen (-).  **Simple Sort:** To sort by id in ascending order or by name in descending order:  &#x60;&#x60;&#x60; sort[]&#x3D;id sort[]&#x3D;-name &#x60;&#x60;&#x60;  **Multiple Sorts:** Combine multiple sorts by separating them with commas: &#x60;&#x60;&#x60; sort[]&#x3D;id&amp;sort[]&#x3D;-name &#x60;&#x60;&#x60; | 
 **filterId** | **int32** | Filter by id. | 
 **filterName** | **string** | Filter by name. | 
 **filterNamePartial** | **string** | Filter by name using partial matching. For example, \&quot;ann\&quot; matches \&quot;Joanna\&quot; or \&quot;Annie\&quot;. | 
 **filterServerConfigId** | **int32** | Filter by ServerConfig ID. | 
 **filterNumInstances** | **int32** | Filter by number of instances. | 
 **filterStatus** | **string** | Filter by status. | 
 **filterMaintenance** | **bool** | Filter by maintenance. | 
 **filterLocationCity** | **string** | Filter by location city. | 
 **filterLocationCityDisplay** | **string** | Filter by location city display name. | 
 **filterLocationContinent** | **string** | Filter by location continent. | 
 **filterLocationCountry** | **string** | Filter by location country. | 
 **filterServerConfigName** | **string** | Filter by ServerConfig name. | 
 **filterServerConfigCommand** | **string** | Filter by ServerConfig command. | 
 **filterServerConfigArgs** | **string** | Filter by ServerConfig arguments. | 
 **filterServerConfigNotes** | **string** | Filter by ServerConfig notes. | 
 **filterServerConfigStatus** | **string** | Filter by ServerConfig status. | 
 **filterServerConfigMaintenance** | **bool** | Filter by ServerConfig maintenance. | 
 **filterServerConfigResourcePackageSlug** | **string** | Filter by ServerConfig resource package slug. | 

### Return type

[**GetAppLocationSettings200Response**](GetAppLocationSettings200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServicesForAppLocationSetting

> []Server ListServicesForAppLocationSetting(ctx, app, appLocationSetting).Execute()

Show all services for a specific app location setting within a given app

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
	appLocationSetting := int32(56) // int32 | The app location setting ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppLocationSettingAPI.ListServicesForAppLocationSetting(context.Background(), app, appLocationSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppLocationSettingAPI.ListServicesForAppLocationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServicesForAppLocationSetting`: []Server
	fmt.Fprintf(os.Stdout, "Response from `AppLocationSettingAPI.ListServicesForAppLocationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **int32** | The app ID | 
**appLocationSetting** | **int32** | The app location setting ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServicesForAppLocationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]Server**](Server.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartServersForAppLocationSetting

> StartServersForAppLocationSetting(ctx, appLocationSetting).Execute()

Start all services related to a specific app location setting

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
	appLocationSetting := int32(56) // int32 | The app location setting ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppLocationSettingAPI.StartServersForAppLocationSetting(context.Background(), appLocationSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppLocationSettingAPI.StartServersForAppLocationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appLocationSetting** | **int32** | The app location setting ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartServersForAppLocationSettingRequest struct via the builder pattern


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


## StopServersForAppLocationSetting

> StopServersForAppLocationSetting(ctx, appLocationSetting).Execute()

Stop all services related to a specific app location setting

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
	appLocationSetting := int32(56) // int32 | The app location setting ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppLocationSettingAPI.StopServersForAppLocationSetting(context.Background(), appLocationSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppLocationSettingAPI.StopServersForAppLocationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appLocationSetting** | **int32** | The app location setting ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopServersForAppLocationSettingRequest struct via the builder pattern


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


## UpdateAppLocationSetting

> AppLocationSetting UpdateAppLocationSetting(ctx, appLocationSetting).UpdateAppLocationSettingRequest(updateAppLocationSettingRequest).Execute()

Update a location setting

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
	appLocationSetting := int32(56) // int32 | The app location setting ID
	updateAppLocationSettingRequest := *openapiclient.NewUpdateAppLocationSettingRequest("Name_example", int32(123)) // UpdateAppLocationSettingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppLocationSettingAPI.UpdateAppLocationSetting(context.Background(), appLocationSetting).UpdateAppLocationSettingRequest(updateAppLocationSettingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppLocationSettingAPI.UpdateAppLocationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppLocationSetting`: AppLocationSetting
	fmt.Fprintf(os.Stdout, "Response from `AppLocationSettingAPI.UpdateAppLocationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appLocationSetting** | **int32** | The app location setting ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppLocationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAppLocationSettingRequest** | [**UpdateAppLocationSettingRequest**](UpdateAppLocationSettingRequest.md) |  | 

### Return type

[**AppLocationSetting**](AppLocationSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
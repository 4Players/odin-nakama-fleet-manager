# \TemplateAppAPI

All URIs are relative to *https://fleet.4players.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TemplateAppMinecraftStore**](TemplateAppAPI.md#TemplateAppMinecraftStore) | **Post** /v1/templates/apps/minecraft | Create a Minecraft template app
[**TemplateAppPalworldStore**](TemplateAppAPI.md#TemplateAppPalworldStore) | **Post** /v1/templates/apps/palworld | Create a Palworld template app



## TemplateAppMinecraftStore

> App TemplateAppMinecraftStore(ctx).StoreMinecraftTemplateRequest(storeMinecraftTemplateRequest).Execute()

Create a Minecraft template app

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
	storeMinecraftTemplateRequest := *openapiclient.NewStoreMinecraftTemplateRequest() // StoreMinecraftTemplateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplateAppAPI.TemplateAppMinecraftStore(context.Background()).StoreMinecraftTemplateRequest(storeMinecraftTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateAppAPI.TemplateAppMinecraftStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TemplateAppMinecraftStore`: App
	fmt.Fprintf(os.Stdout, "Response from `TemplateAppAPI.TemplateAppMinecraftStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemplateAppMinecraftStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeMinecraftTemplateRequest** | [**StoreMinecraftTemplateRequest**](StoreMinecraftTemplateRequest.md) |  | 

### Return type

[**App**](App.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateAppPalworldStore

> App TemplateAppPalworldStore(ctx).StorePalworldTemplateRequest(storePalworldTemplateRequest).Execute()

Create a Palworld template app

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
	storePalworldTemplateRequest := *openapiclient.NewStorePalworldTemplateRequest() // StorePalworldTemplateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplateAppAPI.TemplateAppPalworldStore(context.Background()).StorePalworldTemplateRequest(storePalworldTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateAppAPI.TemplateAppPalworldStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TemplateAppPalworldStore`: App
	fmt.Fprintf(os.Stdout, "Response from `TemplateAppAPI.TemplateAppPalworldStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemplateAppPalworldStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storePalworldTemplateRequest** | [**StorePalworldTemplateRequest**](StorePalworldTemplateRequest.md) |  | 

### Return type

[**App**](App.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
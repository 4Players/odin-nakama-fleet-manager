# \SteamAPI

All URIs are relative to *https://fleet.4players.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SteamGetBranches**](SteamAPI.md#SteamGetBranches) | **Get** /v1/binaries/steam/branches | Get branches for a specific steamworks app ID
[**SteamGetLauncher**](SteamAPI.md#SteamGetLauncher) | **Get** /v1/binaries/steam/launchers | Get launchers for a specific steamworks app ID, optionally filtered by OS



## SteamGetBranches

> []SteamBranch SteamGetBranches(ctx).AppId(appId).Execute()

Get branches for a specific steamworks app ID

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
	appId := int32(56) // int32 | The steamworks app id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SteamAPI.SteamGetBranches(context.Background()).AppId(appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SteamAPI.SteamGetBranches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SteamGetBranches`: []SteamBranch
	fmt.Fprintf(os.Stdout, "Response from `SteamAPI.SteamGetBranches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSteamGetBranchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **int32** | The steamworks app id | 

### Return type

[**[]SteamBranch**](SteamBranch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SteamGetLauncher

> []SteamLauncher SteamGetLauncher(ctx).AppId(appId).Os(os).Execute()

Get launchers for a specific steamworks app ID, optionally filtered by OS

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
	appId := int32(56) // int32 | The steamworks app id
	os := openapiclient.OperatingSystem("windows") // OperatingSystem | The operating system of the binary (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SteamAPI.SteamGetLauncher(context.Background()).AppId(appId).Os(os).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SteamAPI.SteamGetLauncher``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SteamGetLauncher`: []SteamLauncher
	fmt.Fprintf(os.Stdout, "Response from `SteamAPI.SteamGetLauncher`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSteamGetLauncherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **int32** | The steamworks app id | 
 **os** | [**OperatingSystem**](OperatingSystem.md) | The operating system of the binary | 

### Return type

[**[]SteamLauncher**](SteamLauncher.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
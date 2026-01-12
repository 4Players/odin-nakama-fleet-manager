# \AuthAPI

All URIs are relative to *https://fleet.4players.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuthToken**](AuthAPI.md#GetAuthToken) | **Post** /v1/auth/token | Get token
[**RefreshAuthToken**](AuthAPI.md#RefreshAuthToken) | **Put** /v1/auth/token | Refresh token



## GetAuthToken

> Auth GetAuthToken(ctx).AuthRequest(authRequest).Execute()

Get token



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
	authRequest := *openapiclient.NewAuthRequest("Email_example", "Password_example", "1234567890ABCDEFG") // AuthRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.GetAuthToken(context.Background()).AuthRequest(authRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.GetAuthToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthToken`: Auth
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.GetAuthToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authRequest** | [**AuthRequest**](AuthRequest.md) |  | 

### Return type

[**Auth**](Auth.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshAuthToken

> Auth RefreshAuthToken(ctx).Execute()

Refresh token

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.RefreshAuthToken(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.RefreshAuthToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshAuthToken`: Auth
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.RefreshAuthToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshAuthTokenRequest struct via the builder pattern


### Return type

[**Auth**](Auth.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
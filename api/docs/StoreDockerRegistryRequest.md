# StoreDockerRegistryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**DockerRegistryType**](DockerRegistryType.md) | The type of the docker registry | 
**Name** | **string** | The name of the docker registry | 
**Url** | Pointer to **NullableString** | The URL of the docker registry | [optional] 
**InstanceUrl** | Pointer to **NullableString** | The specific instance URL of the docker registry | [optional] 
**Username** | Pointer to **NullableString** | The username to use for authentication | [optional] 
**AccessToken** | Pointer to **NullableString** | The access token to use for authentication | [optional] 
**Password** | Pointer to **NullableString** | The password to use for authentication | [optional] 
**Organization** | Pointer to **bool** | Whether or not the registry is user-owned or organization-owned | [optional] 
**AwsAccessKey** | Pointer to **NullableString** | The AWS access key to use for authentication | [optional] 
**AwsSecretAccessKey** | Pointer to **NullableString** | The AWS secret access key to use for authentication | [optional] 

## Methods

### NewStoreDockerRegistryRequest

`func NewStoreDockerRegistryRequest(type_ DockerRegistryType, name string, ) *StoreDockerRegistryRequest`

NewStoreDockerRegistryRequest instantiates a new StoreDockerRegistryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreDockerRegistryRequestWithDefaults

`func NewStoreDockerRegistryRequestWithDefaults() *StoreDockerRegistryRequest`

NewStoreDockerRegistryRequestWithDefaults instantiates a new StoreDockerRegistryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StoreDockerRegistryRequest) GetType() DockerRegistryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StoreDockerRegistryRequest) GetTypeOk() (*DockerRegistryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StoreDockerRegistryRequest) SetType(v DockerRegistryType)`

SetType sets Type field to given value.


### GetName

`func (o *StoreDockerRegistryRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoreDockerRegistryRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoreDockerRegistryRequest) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *StoreDockerRegistryRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StoreDockerRegistryRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StoreDockerRegistryRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *StoreDockerRegistryRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *StoreDockerRegistryRequest) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *StoreDockerRegistryRequest) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetInstanceUrl

`func (o *StoreDockerRegistryRequest) GetInstanceUrl() string`

GetInstanceUrl returns the InstanceUrl field if non-nil, zero value otherwise.

### GetInstanceUrlOk

`func (o *StoreDockerRegistryRequest) GetInstanceUrlOk() (*string, bool)`

GetInstanceUrlOk returns a tuple with the InstanceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceUrl

`func (o *StoreDockerRegistryRequest) SetInstanceUrl(v string)`

SetInstanceUrl sets InstanceUrl field to given value.

### HasInstanceUrl

`func (o *StoreDockerRegistryRequest) HasInstanceUrl() bool`

HasInstanceUrl returns a boolean if a field has been set.

### SetInstanceUrlNil

`func (o *StoreDockerRegistryRequest) SetInstanceUrlNil(b bool)`

 SetInstanceUrlNil sets the value for InstanceUrl to be an explicit nil

### UnsetInstanceUrl
`func (o *StoreDockerRegistryRequest) UnsetInstanceUrl()`

UnsetInstanceUrl ensures that no value is present for InstanceUrl, not even an explicit nil
### GetUsername

`func (o *StoreDockerRegistryRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *StoreDockerRegistryRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *StoreDockerRegistryRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *StoreDockerRegistryRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *StoreDockerRegistryRequest) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *StoreDockerRegistryRequest) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetAccessToken

`func (o *StoreDockerRegistryRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *StoreDockerRegistryRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *StoreDockerRegistryRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *StoreDockerRegistryRequest) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### SetAccessTokenNil

`func (o *StoreDockerRegistryRequest) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *StoreDockerRegistryRequest) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
### GetPassword

`func (o *StoreDockerRegistryRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *StoreDockerRegistryRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *StoreDockerRegistryRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *StoreDockerRegistryRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *StoreDockerRegistryRequest) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *StoreDockerRegistryRequest) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetOrganization

`func (o *StoreDockerRegistryRequest) GetOrganization() bool`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *StoreDockerRegistryRequest) GetOrganizationOk() (*bool, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *StoreDockerRegistryRequest) SetOrganization(v bool)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *StoreDockerRegistryRequest) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetAwsAccessKey

`func (o *StoreDockerRegistryRequest) GetAwsAccessKey() string`

GetAwsAccessKey returns the AwsAccessKey field if non-nil, zero value otherwise.

### GetAwsAccessKeyOk

`func (o *StoreDockerRegistryRequest) GetAwsAccessKeyOk() (*string, bool)`

GetAwsAccessKeyOk returns a tuple with the AwsAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKey

`func (o *StoreDockerRegistryRequest) SetAwsAccessKey(v string)`

SetAwsAccessKey sets AwsAccessKey field to given value.

### HasAwsAccessKey

`func (o *StoreDockerRegistryRequest) HasAwsAccessKey() bool`

HasAwsAccessKey returns a boolean if a field has been set.

### SetAwsAccessKeyNil

`func (o *StoreDockerRegistryRequest) SetAwsAccessKeyNil(b bool)`

 SetAwsAccessKeyNil sets the value for AwsAccessKey to be an explicit nil

### UnsetAwsAccessKey
`func (o *StoreDockerRegistryRequest) UnsetAwsAccessKey()`

UnsetAwsAccessKey ensures that no value is present for AwsAccessKey, not even an explicit nil
### GetAwsSecretAccessKey

`func (o *StoreDockerRegistryRequest) GetAwsSecretAccessKey() string`

GetAwsSecretAccessKey returns the AwsSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsSecretAccessKeyOk

`func (o *StoreDockerRegistryRequest) GetAwsSecretAccessKeyOk() (*string, bool)`

GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretAccessKey

`func (o *StoreDockerRegistryRequest) SetAwsSecretAccessKey(v string)`

SetAwsSecretAccessKey sets AwsSecretAccessKey field to given value.

### HasAwsSecretAccessKey

`func (o *StoreDockerRegistryRequest) HasAwsSecretAccessKey() bool`

HasAwsSecretAccessKey returns a boolean if a field has been set.

### SetAwsSecretAccessKeyNil

`func (o *StoreDockerRegistryRequest) SetAwsSecretAccessKeyNil(b bool)`

 SetAwsSecretAccessKeyNil sets the value for AwsSecretAccessKey to be an explicit nil

### UnsetAwsSecretAccessKey
`func (o *StoreDockerRegistryRequest) UnsetAwsSecretAccessKey()`

UnsetAwsSecretAccessKey ensures that no value is present for AwsSecretAccessKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
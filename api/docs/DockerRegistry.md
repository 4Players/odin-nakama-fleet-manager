# DockerRegistry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of the registry | 
**Type** | [**DockerRegistryType**](DockerRegistryType.md) | The type of the registry | 
**Name** | **string** | The name of the registry | 
**Username** | **NullableString** | The username to use for authentication | 
**Password** | **NullableString** | The password to use for authentication | 
**Url** | **string** | The URL of the registry | 
**AccessToken** | **NullableString** | The access token to use for authentication | 
**InstanceUrl** | **NullableString** | The URL of the instance | 
**Organization** | **bool** | Whether or not the registry is user-owned or organization-owned | 
**AwsAccessKey** | **NullableString** | AWS access key ID | 
**AwsSecretAccessKey** | **NullableString** | AWS secret access key | 

## Methods

### NewDockerRegistry

`func NewDockerRegistry(id int32, type_ DockerRegistryType, name string, username NullableString, password NullableString, url string, accessToken NullableString, instanceUrl NullableString, organization bool, awsAccessKey NullableString, awsSecretAccessKey NullableString, ) *DockerRegistry`

NewDockerRegistry instantiates a new DockerRegistry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerRegistryWithDefaults

`func NewDockerRegistryWithDefaults() *DockerRegistry`

NewDockerRegistryWithDefaults instantiates a new DockerRegistry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DockerRegistry) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DockerRegistry) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DockerRegistry) SetId(v int32)`

SetId sets Id field to given value.


### GetType

`func (o *DockerRegistry) GetType() DockerRegistryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DockerRegistry) GetTypeOk() (*DockerRegistryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DockerRegistry) SetType(v DockerRegistryType)`

SetType sets Type field to given value.


### GetName

`func (o *DockerRegistry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DockerRegistry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DockerRegistry) SetName(v string)`

SetName sets Name field to given value.


### GetUsername

`func (o *DockerRegistry) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DockerRegistry) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DockerRegistry) SetUsername(v string)`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *DockerRegistry) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *DockerRegistry) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *DockerRegistry) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DockerRegistry) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DockerRegistry) SetPassword(v string)`

SetPassword sets Password field to given value.


### SetPasswordNil

`func (o *DockerRegistry) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *DockerRegistry) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetUrl

`func (o *DockerRegistry) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DockerRegistry) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DockerRegistry) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAccessToken

`func (o *DockerRegistry) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *DockerRegistry) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *DockerRegistry) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### SetAccessTokenNil

`func (o *DockerRegistry) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *DockerRegistry) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
### GetInstanceUrl

`func (o *DockerRegistry) GetInstanceUrl() string`

GetInstanceUrl returns the InstanceUrl field if non-nil, zero value otherwise.

### GetInstanceUrlOk

`func (o *DockerRegistry) GetInstanceUrlOk() (*string, bool)`

GetInstanceUrlOk returns a tuple with the InstanceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceUrl

`func (o *DockerRegistry) SetInstanceUrl(v string)`

SetInstanceUrl sets InstanceUrl field to given value.


### SetInstanceUrlNil

`func (o *DockerRegistry) SetInstanceUrlNil(b bool)`

 SetInstanceUrlNil sets the value for InstanceUrl to be an explicit nil

### UnsetInstanceUrl
`func (o *DockerRegistry) UnsetInstanceUrl()`

UnsetInstanceUrl ensures that no value is present for InstanceUrl, not even an explicit nil
### GetOrganization

`func (o *DockerRegistry) GetOrganization() bool`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *DockerRegistry) GetOrganizationOk() (*bool, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *DockerRegistry) SetOrganization(v bool)`

SetOrganization sets Organization field to given value.


### GetAwsAccessKey

`func (o *DockerRegistry) GetAwsAccessKey() string`

GetAwsAccessKey returns the AwsAccessKey field if non-nil, zero value otherwise.

### GetAwsAccessKeyOk

`func (o *DockerRegistry) GetAwsAccessKeyOk() (*string, bool)`

GetAwsAccessKeyOk returns a tuple with the AwsAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKey

`func (o *DockerRegistry) SetAwsAccessKey(v string)`

SetAwsAccessKey sets AwsAccessKey field to given value.


### SetAwsAccessKeyNil

`func (o *DockerRegistry) SetAwsAccessKeyNil(b bool)`

 SetAwsAccessKeyNil sets the value for AwsAccessKey to be an explicit nil

### UnsetAwsAccessKey
`func (o *DockerRegistry) UnsetAwsAccessKey()`

UnsetAwsAccessKey ensures that no value is present for AwsAccessKey, not even an explicit nil
### GetAwsSecretAccessKey

`func (o *DockerRegistry) GetAwsSecretAccessKey() string`

GetAwsSecretAccessKey returns the AwsSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsSecretAccessKeyOk

`func (o *DockerRegistry) GetAwsSecretAccessKeyOk() (*string, bool)`

GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretAccessKey

`func (o *DockerRegistry) SetAwsSecretAccessKey(v string)`

SetAwsSecretAccessKey sets AwsSecretAccessKey field to given value.


### SetAwsSecretAccessKeyNil

`func (o *DockerRegistry) SetAwsSecretAccessKeyNil(b bool)`

 SetAwsSecretAccessKeyNil sets the value for AwsSecretAccessKey to be an explicit nil

### UnsetAwsSecretAccessKey
`func (o *DockerRegistry) UnsetAwsSecretAccessKey()`

UnsetAwsSecretAccessKey ensures that no value is present for AwsSecretAccessKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
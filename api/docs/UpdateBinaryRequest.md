# UpdateBinaryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the binary | 
**Version** | **string** | The version of the binary | 
**Type** | [**BinaryType**](BinaryType.md) | The type of the binary | 
**Os** | [**OperatingSystem**](OperatingSystem.md) | The operating system of the binary | 
**Steam** | Pointer to [**CreateUpdateSteam**](CreateUpdateSteam.md) | The steam settings | [optional] 
**DockerImage** | Pointer to [**CreateUpdateDockerImage**](CreateUpdateDockerImage.md) | The docker image settings | [optional] 

## Methods

### NewUpdateBinaryRequest

`func NewUpdateBinaryRequest(name string, version string, type_ BinaryType, os OperatingSystem, ) *UpdateBinaryRequest`

NewUpdateBinaryRequest instantiates a new UpdateBinaryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBinaryRequestWithDefaults

`func NewUpdateBinaryRequestWithDefaults() *UpdateBinaryRequest`

NewUpdateBinaryRequestWithDefaults instantiates a new UpdateBinaryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateBinaryRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateBinaryRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateBinaryRequest) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *UpdateBinaryRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdateBinaryRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdateBinaryRequest) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetType

`func (o *UpdateBinaryRequest) GetType() BinaryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateBinaryRequest) GetTypeOk() (*BinaryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateBinaryRequest) SetType(v BinaryType)`

SetType sets Type field to given value.


### GetOs

`func (o *UpdateBinaryRequest) GetOs() OperatingSystem`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *UpdateBinaryRequest) GetOsOk() (*OperatingSystem, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *UpdateBinaryRequest) SetOs(v OperatingSystem)`

SetOs sets Os field to given value.


### GetSteam

`func (o *UpdateBinaryRequest) GetSteam() CreateUpdateSteam`

GetSteam returns the Steam field if non-nil, zero value otherwise.

### GetSteamOk

`func (o *UpdateBinaryRequest) GetSteamOk() (*CreateUpdateSteam, bool)`

GetSteamOk returns a tuple with the Steam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteam

`func (o *UpdateBinaryRequest) SetSteam(v CreateUpdateSteam)`

SetSteam sets Steam field to given value.

### HasSteam

`func (o *UpdateBinaryRequest) HasSteam() bool`

HasSteam returns a boolean if a field has been set.

### GetDockerImage

`func (o *UpdateBinaryRequest) GetDockerImage() CreateUpdateDockerImage`

GetDockerImage returns the DockerImage field if non-nil, zero value otherwise.

### GetDockerImageOk

`func (o *UpdateBinaryRequest) GetDockerImageOk() (*CreateUpdateDockerImage, bool)`

GetDockerImageOk returns a tuple with the DockerImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImage

`func (o *UpdateBinaryRequest) SetDockerImage(v CreateUpdateDockerImage)`

SetDockerImage sets DockerImage field to given value.

### HasDockerImage

`func (o *UpdateBinaryRequest) HasDockerImage() bool`

HasDockerImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
# StoreBinaryRequest

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

### NewStoreBinaryRequest

`func NewStoreBinaryRequest(name string, version string, type_ BinaryType, os OperatingSystem, ) *StoreBinaryRequest`

NewStoreBinaryRequest instantiates a new StoreBinaryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreBinaryRequestWithDefaults

`func NewStoreBinaryRequestWithDefaults() *StoreBinaryRequest`

NewStoreBinaryRequestWithDefaults instantiates a new StoreBinaryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StoreBinaryRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoreBinaryRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoreBinaryRequest) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *StoreBinaryRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StoreBinaryRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StoreBinaryRequest) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetType

`func (o *StoreBinaryRequest) GetType() BinaryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StoreBinaryRequest) GetTypeOk() (*BinaryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StoreBinaryRequest) SetType(v BinaryType)`

SetType sets Type field to given value.


### GetOs

`func (o *StoreBinaryRequest) GetOs() OperatingSystem`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *StoreBinaryRequest) GetOsOk() (*OperatingSystem, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *StoreBinaryRequest) SetOs(v OperatingSystem)`

SetOs sets Os field to given value.


### GetSteam

`func (o *StoreBinaryRequest) GetSteam() CreateUpdateSteam`

GetSteam returns the Steam field if non-nil, zero value otherwise.

### GetSteamOk

`func (o *StoreBinaryRequest) GetSteamOk() (*CreateUpdateSteam, bool)`

GetSteamOk returns a tuple with the Steam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteam

`func (o *StoreBinaryRequest) SetSteam(v CreateUpdateSteam)`

SetSteam sets Steam field to given value.

### HasSteam

`func (o *StoreBinaryRequest) HasSteam() bool`

HasSteam returns a boolean if a field has been set.

### GetDockerImage

`func (o *StoreBinaryRequest) GetDockerImage() CreateUpdateDockerImage`

GetDockerImage returns the DockerImage field if non-nil, zero value otherwise.

### GetDockerImageOk

`func (o *StoreBinaryRequest) GetDockerImageOk() (*CreateUpdateDockerImage, bool)`

GetDockerImageOk returns a tuple with the DockerImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImage

`func (o *StoreBinaryRequest) SetDockerImage(v CreateUpdateDockerImage)`

SetDockerImage sets DockerImage field to given value.

### HasDockerImage

`func (o *StoreBinaryRequest) HasDockerImage() bool`

HasDockerImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
# DockerImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageName** | **string** | The name of the docker image | 
**RegistryId** | **int32** | The id of the registry to load the image | 
**InUse** | **bool** | Indicates whether the docker image is currently in use | 

## Methods

### NewDockerImage

`func NewDockerImage(imageName string, registryId int32, inUse bool, ) *DockerImage`

NewDockerImage instantiates a new DockerImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerImageWithDefaults

`func NewDockerImageWithDefaults() *DockerImage`

NewDockerImageWithDefaults instantiates a new DockerImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageName

`func (o *DockerImage) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *DockerImage) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *DockerImage) SetImageName(v string)`

SetImageName sets ImageName field to given value.


### GetRegistryId

`func (o *DockerImage) GetRegistryId() int32`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *DockerImage) GetRegistryIdOk() (*int32, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *DockerImage) SetRegistryId(v int32)`

SetRegistryId sets RegistryId field to given value.


### GetInUse

`func (o *DockerImage) GetInUse() bool`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *DockerImage) GetInUseOk() (*bool, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *DockerImage) SetInUse(v bool)`

SetInUse sets InUse field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
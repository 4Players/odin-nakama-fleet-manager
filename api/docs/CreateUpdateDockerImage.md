# CreateUpdateDockerImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageName** | **string** | The name of the docker image | 
**RegistryId** | **int32** | The id of the registry to load the image | 

## Methods

### NewCreateUpdateDockerImage

`func NewCreateUpdateDockerImage(imageName string, registryId int32, ) *CreateUpdateDockerImage`

NewCreateUpdateDockerImage instantiates a new CreateUpdateDockerImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpdateDockerImageWithDefaults

`func NewCreateUpdateDockerImageWithDefaults() *CreateUpdateDockerImage`

NewCreateUpdateDockerImageWithDefaults instantiates a new CreateUpdateDockerImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageName

`func (o *CreateUpdateDockerImage) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *CreateUpdateDockerImage) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *CreateUpdateDockerImage) SetImageName(v string)`

SetImageName sets ImageName field to given value.


### GetRegistryId

`func (o *CreateUpdateDockerImage) GetRegistryId() int32`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *CreateUpdateDockerImage) GetRegistryIdOk() (*int32, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *CreateUpdateDockerImage) SetRegistryId(v int32)`

SetRegistryId sets RegistryId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
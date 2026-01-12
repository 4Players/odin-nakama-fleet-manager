# TaggedImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the docker image including the tag. | 
**Tag** | **string** | The tag of the docker image. | 
**Repository** | **string** | The name of the repository. | 
**Location** | **string** | The location URI of the tagged docker image. | 

## Methods

### NewTaggedImage

`func NewTaggedImage(name string, tag string, repository string, location string, ) *TaggedImage`

NewTaggedImage instantiates a new TaggedImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaggedImageWithDefaults

`func NewTaggedImageWithDefaults() *TaggedImage`

NewTaggedImageWithDefaults instantiates a new TaggedImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TaggedImage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaggedImage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaggedImage) SetName(v string)`

SetName sets Name field to given value.


### GetTag

`func (o *TaggedImage) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *TaggedImage) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *TaggedImage) SetTag(v string)`

SetTag sets Tag field to given value.


### GetRepository

`func (o *TaggedImage) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *TaggedImage) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *TaggedImage) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetLocation

`func (o *TaggedImage) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *TaggedImage) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *TaggedImage) SetLocation(v string)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
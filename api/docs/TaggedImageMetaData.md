# TaggedImageMetaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cached** | **bool** | Whether the data was retrieved from the cache. | 
**CachedAt** | **NullableString** | Timestamp of when the data was last cached. | 

## Methods

### NewTaggedImageMetaData

`func NewTaggedImageMetaData(cached bool, cachedAt NullableString, ) *TaggedImageMetaData`

NewTaggedImageMetaData instantiates a new TaggedImageMetaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaggedImageMetaDataWithDefaults

`func NewTaggedImageMetaDataWithDefaults() *TaggedImageMetaData`

NewTaggedImageMetaDataWithDefaults instantiates a new TaggedImageMetaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCached

`func (o *TaggedImageMetaData) GetCached() bool`

GetCached returns the Cached field if non-nil, zero value otherwise.

### GetCachedOk

`func (o *TaggedImageMetaData) GetCachedOk() (*bool, bool)`

GetCachedOk returns a tuple with the Cached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCached

`func (o *TaggedImageMetaData) SetCached(v bool)`

SetCached sets Cached field to given value.


### GetCachedAt

`func (o *TaggedImageMetaData) GetCachedAt() string`

GetCachedAt returns the CachedAt field if non-nil, zero value otherwise.

### GetCachedAtOk

`func (o *TaggedImageMetaData) GetCachedAtOk() (*string, bool)`

GetCachedAtOk returns a tuple with the CachedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedAt

`func (o *TaggedImageMetaData) SetCachedAt(v string)`

SetCachedAt sets CachedAt field to given value.


### SetCachedAtNil

`func (o *TaggedImageMetaData) SetCachedAtNil(b bool)`

 SetCachedAtNil sets the value for CachedAt to be an explicit nil

### UnsetCachedAt
`func (o *TaggedImageMetaData) UnsetCachedAt()`

UnsetCachedAt ensures that no value is present for CachedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
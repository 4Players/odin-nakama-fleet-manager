# GetAppLocationSettings200ResponseMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | **int32** |  | 
**From** | **NullableInt32** |  | 
**LastPage** | **int32** |  | 
**Links** | [**[]GetAppLocationSettings200ResponseMetaLinksInner**](GetAppLocationSettings200ResponseMetaLinksInner.md) | Generated paginator links. | 
**Path** | **NullableString** | Base path for paginator generated URLs. | 
**PerPage** | **int32** | Number of items shown per page. | 
**To** | **NullableInt32** | Number of the last item in the slice. | 
**Total** | **int32** | Total number of items being paginated. | 

## Methods

### NewGetAppLocationSettings200ResponseMeta

`func NewGetAppLocationSettings200ResponseMeta(currentPage int32, from NullableInt32, lastPage int32, links []GetAppLocationSettings200ResponseMetaLinksInner, path NullableString, perPage int32, to NullableInt32, total int32, ) *GetAppLocationSettings200ResponseMeta`

NewGetAppLocationSettings200ResponseMeta instantiates a new GetAppLocationSettings200ResponseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAppLocationSettings200ResponseMetaWithDefaults

`func NewGetAppLocationSettings200ResponseMetaWithDefaults() *GetAppLocationSettings200ResponseMeta`

NewGetAppLocationSettings200ResponseMetaWithDefaults instantiates a new GetAppLocationSettings200ResponseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *GetAppLocationSettings200ResponseMeta) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *GetAppLocationSettings200ResponseMeta) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *GetAppLocationSettings200ResponseMeta) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.


### GetFrom

`func (o *GetAppLocationSettings200ResponseMeta) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GetAppLocationSettings200ResponseMeta) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GetAppLocationSettings200ResponseMeta) SetFrom(v int32)`

SetFrom sets From field to given value.


### SetFromNil

`func (o *GetAppLocationSettings200ResponseMeta) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *GetAppLocationSettings200ResponseMeta) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetLastPage

`func (o *GetAppLocationSettings200ResponseMeta) GetLastPage() int32`

GetLastPage returns the LastPage field if non-nil, zero value otherwise.

### GetLastPageOk

`func (o *GetAppLocationSettings200ResponseMeta) GetLastPageOk() (*int32, bool)`

GetLastPageOk returns a tuple with the LastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPage

`func (o *GetAppLocationSettings200ResponseMeta) SetLastPage(v int32)`

SetLastPage sets LastPage field to given value.


### GetLinks

`func (o *GetAppLocationSettings200ResponseMeta) GetLinks() []GetAppLocationSettings200ResponseMetaLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetAppLocationSettings200ResponseMeta) GetLinksOk() (*[]GetAppLocationSettings200ResponseMetaLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetAppLocationSettings200ResponseMeta) SetLinks(v []GetAppLocationSettings200ResponseMetaLinksInner)`

SetLinks sets Links field to given value.


### GetPath

`func (o *GetAppLocationSettings200ResponseMeta) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GetAppLocationSettings200ResponseMeta) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GetAppLocationSettings200ResponseMeta) SetPath(v string)`

SetPath sets Path field to given value.


### SetPathNil

`func (o *GetAppLocationSettings200ResponseMeta) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *GetAppLocationSettings200ResponseMeta) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPerPage

`func (o *GetAppLocationSettings200ResponseMeta) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *GetAppLocationSettings200ResponseMeta) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *GetAppLocationSettings200ResponseMeta) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.


### GetTo

`func (o *GetAppLocationSettings200ResponseMeta) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GetAppLocationSettings200ResponseMeta) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GetAppLocationSettings200ResponseMeta) SetTo(v int32)`

SetTo sets To field to given value.


### SetToNil

`func (o *GetAppLocationSettings200ResponseMeta) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *GetAppLocationSettings200ResponseMeta) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetTotal

`func (o *GetAppLocationSettings200ResponseMeta) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetAppLocationSettings200ResponseMeta) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetAppLocationSettings200ResponseMeta) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
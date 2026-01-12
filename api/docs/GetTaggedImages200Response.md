# GetTaggedImages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TaggedImage**](TaggedImage.md) |  | 
**Meta** | [**TaggedImageMetaData**](TaggedImageMetaData.md) |  | 

## Methods

### NewGetTaggedImages200Response

`func NewGetTaggedImages200Response(data []TaggedImage, meta TaggedImageMetaData, ) *GetTaggedImages200Response`

NewGetTaggedImages200Response instantiates a new GetTaggedImages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTaggedImages200ResponseWithDefaults

`func NewGetTaggedImages200ResponseWithDefaults() *GetTaggedImages200Response`

NewGetTaggedImages200ResponseWithDefaults instantiates a new GetTaggedImages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetTaggedImages200Response) GetData() []TaggedImage`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTaggedImages200Response) GetDataOk() (*[]TaggedImage, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTaggedImages200Response) SetData(v []TaggedImage)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetTaggedImages200Response) GetMeta() TaggedImageMetaData`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTaggedImages200Response) GetMetaOk() (*TaggedImageMetaData, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTaggedImages200Response) SetMeta(v TaggedImageMetaData)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
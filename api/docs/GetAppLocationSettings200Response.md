# GetAppLocationSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AppLocationSetting**](AppLocationSetting.md) |  | 
**Links** | [**GetAppLocationSettings200ResponseLinks**](GetAppLocationSettings200ResponseLinks.md) |  | 
**Meta** | [**GetAppLocationSettings200ResponseMeta**](GetAppLocationSettings200ResponseMeta.md) |  | 

## Methods

### NewGetAppLocationSettings200Response

`func NewGetAppLocationSettings200Response(data []AppLocationSetting, links GetAppLocationSettings200ResponseLinks, meta GetAppLocationSettings200ResponseMeta, ) *GetAppLocationSettings200Response`

NewGetAppLocationSettings200Response instantiates a new GetAppLocationSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAppLocationSettings200ResponseWithDefaults

`func NewGetAppLocationSettings200ResponseWithDefaults() *GetAppLocationSettings200Response`

NewGetAppLocationSettings200ResponseWithDefaults instantiates a new GetAppLocationSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetAppLocationSettings200Response) GetData() []AppLocationSetting`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAppLocationSettings200Response) GetDataOk() (*[]AppLocationSetting, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAppLocationSettings200Response) SetData(v []AppLocationSetting)`

SetData sets Data field to given value.


### GetLinks

`func (o *GetAppLocationSettings200Response) GetLinks() GetAppLocationSettings200ResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetAppLocationSettings200Response) GetLinksOk() (*GetAppLocationSettings200ResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetAppLocationSettings200Response) SetLinks(v GetAppLocationSettings200ResponseLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *GetAppLocationSettings200Response) GetMeta() GetAppLocationSettings200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAppLocationSettings200Response) GetMetaOk() (*GetAppLocationSettings200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAppLocationSettings200Response) SetMeta(v GetAppLocationSettings200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
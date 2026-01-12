# GetDockerRegistries200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DockerRegistry**](DockerRegistry.md) |  | 
**Links** | [**GetAppLocationSettings200ResponseLinks**](GetAppLocationSettings200ResponseLinks.md) |  | 
**Meta** | [**GetAppLocationSettings200ResponseMeta**](GetAppLocationSettings200ResponseMeta.md) |  | 

## Methods

### NewGetDockerRegistries200Response

`func NewGetDockerRegistries200Response(data []DockerRegistry, links GetAppLocationSettings200ResponseLinks, meta GetAppLocationSettings200ResponseMeta, ) *GetDockerRegistries200Response`

NewGetDockerRegistries200Response instantiates a new GetDockerRegistries200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDockerRegistries200ResponseWithDefaults

`func NewGetDockerRegistries200ResponseWithDefaults() *GetDockerRegistries200Response`

NewGetDockerRegistries200ResponseWithDefaults instantiates a new GetDockerRegistries200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetDockerRegistries200Response) GetData() []DockerRegistry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetDockerRegistries200Response) GetDataOk() (*[]DockerRegistry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetDockerRegistries200Response) SetData(v []DockerRegistry)`

SetData sets Data field to given value.


### GetLinks

`func (o *GetDockerRegistries200Response) GetLinks() GetAppLocationSettings200ResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetDockerRegistries200Response) GetLinksOk() (*GetAppLocationSettings200ResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetDockerRegistries200Response) SetLinks(v GetAppLocationSettings200ResponseLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *GetDockerRegistries200Response) GetMeta() GetAppLocationSettings200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDockerRegistries200Response) GetMetaOk() (*GetAppLocationSettings200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDockerRegistries200Response) SetMeta(v GetAppLocationSettings200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
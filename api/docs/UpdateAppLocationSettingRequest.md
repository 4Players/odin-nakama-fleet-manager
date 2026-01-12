# UpdateAppLocationSettingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the location setting | 
**NumInstances** | **int32** | The number of instances that should run at the specific location | 

## Methods

### NewUpdateAppLocationSettingRequest

`func NewUpdateAppLocationSettingRequest(name string, numInstances int32, ) *UpdateAppLocationSettingRequest`

NewUpdateAppLocationSettingRequest instantiates a new UpdateAppLocationSettingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAppLocationSettingRequestWithDefaults

`func NewUpdateAppLocationSettingRequestWithDefaults() *UpdateAppLocationSettingRequest`

NewUpdateAppLocationSettingRequestWithDefaults instantiates a new UpdateAppLocationSettingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateAppLocationSettingRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAppLocationSettingRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAppLocationSettingRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNumInstances

`func (o *UpdateAppLocationSettingRequest) GetNumInstances() int32`

GetNumInstances returns the NumInstances field if non-nil, zero value otherwise.

### GetNumInstancesOk

`func (o *UpdateAppLocationSettingRequest) GetNumInstancesOk() (*int32, bool)`

GetNumInstancesOk returns a tuple with the NumInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumInstances

`func (o *UpdateAppLocationSettingRequest) SetNumInstances(v int32)`

SetNumInstances sets NumInstances field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
# StoreAppLocationSettingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the location setting | 
**ServerConfigId** | **int32** | The id of the server configuration that should be deployed | 
**NumInstances** | **int32** | The number of instances that should run at the specific location | 
**Placement** | Pointer to [**CreateUpdatePlacement**](CreateUpdatePlacement.md) | The placement settings that define the location and other constraints | [optional] 
**Password** | Pointer to **NullableString** | Password required to deploy services to a protected node location. Required when the selected location is password-protected. | [optional] 

## Methods

### NewStoreAppLocationSettingRequest

`func NewStoreAppLocationSettingRequest(name string, serverConfigId int32, numInstances int32, ) *StoreAppLocationSettingRequest`

NewStoreAppLocationSettingRequest instantiates a new StoreAppLocationSettingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreAppLocationSettingRequestWithDefaults

`func NewStoreAppLocationSettingRequestWithDefaults() *StoreAppLocationSettingRequest`

NewStoreAppLocationSettingRequestWithDefaults instantiates a new StoreAppLocationSettingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StoreAppLocationSettingRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoreAppLocationSettingRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoreAppLocationSettingRequest) SetName(v string)`

SetName sets Name field to given value.


### GetServerConfigId

`func (o *StoreAppLocationSettingRequest) GetServerConfigId() int32`

GetServerConfigId returns the ServerConfigId field if non-nil, zero value otherwise.

### GetServerConfigIdOk

`func (o *StoreAppLocationSettingRequest) GetServerConfigIdOk() (*int32, bool)`

GetServerConfigIdOk returns a tuple with the ServerConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerConfigId

`func (o *StoreAppLocationSettingRequest) SetServerConfigId(v int32)`

SetServerConfigId sets ServerConfigId field to given value.


### GetNumInstances

`func (o *StoreAppLocationSettingRequest) GetNumInstances() int32`

GetNumInstances returns the NumInstances field if non-nil, zero value otherwise.

### GetNumInstancesOk

`func (o *StoreAppLocationSettingRequest) GetNumInstancesOk() (*int32, bool)`

GetNumInstancesOk returns a tuple with the NumInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumInstances

`func (o *StoreAppLocationSettingRequest) SetNumInstances(v int32)`

SetNumInstances sets NumInstances field to given value.


### GetPlacement

`func (o *StoreAppLocationSettingRequest) GetPlacement() CreateUpdatePlacement`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *StoreAppLocationSettingRequest) GetPlacementOk() (*CreateUpdatePlacement, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *StoreAppLocationSettingRequest) SetPlacement(v CreateUpdatePlacement)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *StoreAppLocationSettingRequest) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### GetPassword

`func (o *StoreAppLocationSettingRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *StoreAppLocationSettingRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *StoreAppLocationSettingRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *StoreAppLocationSettingRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *StoreAppLocationSettingRequest) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *StoreAppLocationSettingRequest) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
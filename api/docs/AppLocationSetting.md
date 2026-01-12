# AppLocationSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The id of the app location setting | 
**Name** | **string** | The name of the app location setting | 
**AppId** | **int32** | The app id of the app location setting | 
**ServerConfigId** | **int32** | The server config id of the app location setting | 
**NumInstances** | **int32** | The number of instances that should run at the specific location | 
**AutoScalerEnabled** | **bool** | Whether the auto scaler is enabled | 
**AutoScalerMin** | Pointer to **int32** | The minimum number of instances that should run at the specific location | [optional] 
**AutoScalerMax** | Pointer to **int32** | The maximum number of instances that should run at the specific location | [optional] 
**Status** | [**AppLocationSettingStatus**](AppLocationSettingStatus.md) | The current status | 
**StatusMessage** | **NullableString** | An optional message | 
**Maintenance** | **bool** | Indicates if the app location setting is under maintenance | 
**InUse** | **bool** | Indicates whether the app location setting is currently in use | 
**Placement** | [**Placement**](Placement.md) | The placement to use | 
**ServerConfig** | Pointer to [**ServerConfig**](ServerConfig.md) | The server config to use | [optional] 

## Methods

### NewAppLocationSetting

`func NewAppLocationSetting(id int32, name string, appId int32, serverConfigId int32, numInstances int32, autoScalerEnabled bool, status AppLocationSettingStatus, statusMessage NullableString, maintenance bool, inUse bool, placement Placement, ) *AppLocationSetting`

NewAppLocationSetting instantiates a new AppLocationSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppLocationSettingWithDefaults

`func NewAppLocationSettingWithDefaults() *AppLocationSetting`

NewAppLocationSettingWithDefaults instantiates a new AppLocationSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppLocationSetting) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppLocationSetting) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppLocationSetting) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *AppLocationSetting) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppLocationSetting) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppLocationSetting) SetName(v string)`

SetName sets Name field to given value.


### GetAppId

`func (o *AppLocationSetting) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AppLocationSetting) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AppLocationSetting) SetAppId(v int32)`

SetAppId sets AppId field to given value.


### GetServerConfigId

`func (o *AppLocationSetting) GetServerConfigId() int32`

GetServerConfigId returns the ServerConfigId field if non-nil, zero value otherwise.

### GetServerConfigIdOk

`func (o *AppLocationSetting) GetServerConfigIdOk() (*int32, bool)`

GetServerConfigIdOk returns a tuple with the ServerConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerConfigId

`func (o *AppLocationSetting) SetServerConfigId(v int32)`

SetServerConfigId sets ServerConfigId field to given value.


### GetNumInstances

`func (o *AppLocationSetting) GetNumInstances() int32`

GetNumInstances returns the NumInstances field if non-nil, zero value otherwise.

### GetNumInstancesOk

`func (o *AppLocationSetting) GetNumInstancesOk() (*int32, bool)`

GetNumInstancesOk returns a tuple with the NumInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumInstances

`func (o *AppLocationSetting) SetNumInstances(v int32)`

SetNumInstances sets NumInstances field to given value.


### GetAutoScalerEnabled

`func (o *AppLocationSetting) GetAutoScalerEnabled() bool`

GetAutoScalerEnabled returns the AutoScalerEnabled field if non-nil, zero value otherwise.

### GetAutoScalerEnabledOk

`func (o *AppLocationSetting) GetAutoScalerEnabledOk() (*bool, bool)`

GetAutoScalerEnabledOk returns a tuple with the AutoScalerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScalerEnabled

`func (o *AppLocationSetting) SetAutoScalerEnabled(v bool)`

SetAutoScalerEnabled sets AutoScalerEnabled field to given value.


### GetAutoScalerMin

`func (o *AppLocationSetting) GetAutoScalerMin() int32`

GetAutoScalerMin returns the AutoScalerMin field if non-nil, zero value otherwise.

### GetAutoScalerMinOk

`func (o *AppLocationSetting) GetAutoScalerMinOk() (*int32, bool)`

GetAutoScalerMinOk returns a tuple with the AutoScalerMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScalerMin

`func (o *AppLocationSetting) SetAutoScalerMin(v int32)`

SetAutoScalerMin sets AutoScalerMin field to given value.

### HasAutoScalerMin

`func (o *AppLocationSetting) HasAutoScalerMin() bool`

HasAutoScalerMin returns a boolean if a field has been set.

### GetAutoScalerMax

`func (o *AppLocationSetting) GetAutoScalerMax() int32`

GetAutoScalerMax returns the AutoScalerMax field if non-nil, zero value otherwise.

### GetAutoScalerMaxOk

`func (o *AppLocationSetting) GetAutoScalerMaxOk() (*int32, bool)`

GetAutoScalerMaxOk returns a tuple with the AutoScalerMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScalerMax

`func (o *AppLocationSetting) SetAutoScalerMax(v int32)`

SetAutoScalerMax sets AutoScalerMax field to given value.

### HasAutoScalerMax

`func (o *AppLocationSetting) HasAutoScalerMax() bool`

HasAutoScalerMax returns a boolean if a field has been set.

### GetStatus

`func (o *AppLocationSetting) GetStatus() AppLocationSettingStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppLocationSetting) GetStatusOk() (*AppLocationSettingStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppLocationSetting) SetStatus(v AppLocationSettingStatus)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *AppLocationSetting) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *AppLocationSetting) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *AppLocationSetting) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.


### SetStatusMessageNil

`func (o *AppLocationSetting) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *AppLocationSetting) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetMaintenance

`func (o *AppLocationSetting) GetMaintenance() bool`

GetMaintenance returns the Maintenance field if non-nil, zero value otherwise.

### GetMaintenanceOk

`func (o *AppLocationSetting) GetMaintenanceOk() (*bool, bool)`

GetMaintenanceOk returns a tuple with the Maintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenance

`func (o *AppLocationSetting) SetMaintenance(v bool)`

SetMaintenance sets Maintenance field to given value.


### GetInUse

`func (o *AppLocationSetting) GetInUse() bool`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *AppLocationSetting) GetInUseOk() (*bool, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *AppLocationSetting) SetInUse(v bool)`

SetInUse sets InUse field to given value.


### GetPlacement

`func (o *AppLocationSetting) GetPlacement() Placement`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *AppLocationSetting) GetPlacementOk() (*Placement, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *AppLocationSetting) SetPlacement(v Placement)`

SetPlacement sets Placement field to given value.


### GetServerConfig

`func (o *AppLocationSetting) GetServerConfig() ServerConfig`

GetServerConfig returns the ServerConfig field if non-nil, zero value otherwise.

### GetServerConfigOk

`func (o *AppLocationSetting) GetServerConfigOk() (*ServerConfig, bool)`

GetServerConfigOk returns a tuple with the ServerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerConfig

`func (o *AppLocationSetting) SetServerConfig(v ServerConfig)`

SetServerConfig sets ServerConfig field to given value.

### HasServerConfig

`func (o *AppLocationSetting) HasServerConfig() bool`

HasServerConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
# App

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The id of the app | 
**Name** | **string** | The name of the app | 
**InUse** | **bool** | Indicates whether the app is currently in use | 
**BillingState** | [**AppBillingState**](AppBillingState.md) | The billing state of the app | 
**Maintenance** | **bool** | Indicates if the app is under maintenance | 
**Status** | [**AppStatus**](AppStatus.md) | The current status of the app | 
**StatusMessage** | **NullableString** | An optional status message | 

## Methods

### NewApp

`func NewApp(id int32, name string, inUse bool, billingState AppBillingState, maintenance bool, status AppStatus, statusMessage NullableString, ) *App`

NewApp instantiates a new App object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppWithDefaults

`func NewAppWithDefaults() *App`

NewAppWithDefaults instantiates a new App object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *App) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *App) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *App) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *App) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *App) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *App) SetName(v string)`

SetName sets Name field to given value.


### GetInUse

`func (o *App) GetInUse() bool`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *App) GetInUseOk() (*bool, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *App) SetInUse(v bool)`

SetInUse sets InUse field to given value.


### GetBillingState

`func (o *App) GetBillingState() AppBillingState`

GetBillingState returns the BillingState field if non-nil, zero value otherwise.

### GetBillingStateOk

`func (o *App) GetBillingStateOk() (*AppBillingState, bool)`

GetBillingStateOk returns a tuple with the BillingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingState

`func (o *App) SetBillingState(v AppBillingState)`

SetBillingState sets BillingState field to given value.


### GetMaintenance

`func (o *App) GetMaintenance() bool`

GetMaintenance returns the Maintenance field if non-nil, zero value otherwise.

### GetMaintenanceOk

`func (o *App) GetMaintenanceOk() (*bool, bool)`

GetMaintenanceOk returns a tuple with the Maintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenance

`func (o *App) SetMaintenance(v bool)`

SetMaintenance sets Maintenance field to given value.


### GetStatus

`func (o *App) GetStatus() AppStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *App) GetStatusOk() (*AppStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *App) SetStatus(v AppStatus)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *App) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *App) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *App) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.


### SetStatusMessageNil

`func (o *App) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *App) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
# Binary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The id of the binary | 
**AppId** | **int32** | The app id of the binary | 
**Name** | **string** | The app id of the binary | 
**Version** | **string** | The version of the binary | 
**Type** | [**BinaryType**](BinaryType.md) | The type of the binary | 
**Os** | [**OperatingSystem**](OperatingSystem.md) | The operating system of the binary | 
**Maintenance** | **bool** | Indicates if the binary is under maintenance | 
**Status** | [**BinaryStatus**](BinaryStatus.md) | The current status of the binary | 
**StatusMessage** | **NullableString** | An optional message returned by the build process | 
**Progress** | Pointer to **NullableFloat32** | The current progress percentage of the image build (0-100). Only present for binaries of type steam. | [optional] 
**ProgressMessage** | Pointer to **NullableString** | A message describing the current build step. Only present for binaries of type steam. | [optional] 
**InUse** | **bool** | Indicates whether the binary is currently in use | 
**DockerImage** | Pointer to [**DockerImage**](DockerImage.md) | The docker image of the binary | [optional] 
**Steam** | Pointer to [**Steam**](Steam.md) | The steam of the binary | [optional] 

## Methods

### NewBinary

`func NewBinary(id int32, appId int32, name string, version string, type_ BinaryType, os OperatingSystem, maintenance bool, status BinaryStatus, statusMessage NullableString, inUse bool, ) *Binary`

NewBinary instantiates a new Binary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBinaryWithDefaults

`func NewBinaryWithDefaults() *Binary`

NewBinaryWithDefaults instantiates a new Binary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Binary) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Binary) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Binary) SetId(v int32)`

SetId sets Id field to given value.


### GetAppId

`func (o *Binary) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *Binary) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *Binary) SetAppId(v int32)`

SetAppId sets AppId field to given value.


### GetName

`func (o *Binary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Binary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Binary) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *Binary) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Binary) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Binary) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetType

`func (o *Binary) GetType() BinaryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Binary) GetTypeOk() (*BinaryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Binary) SetType(v BinaryType)`

SetType sets Type field to given value.


### GetOs

`func (o *Binary) GetOs() OperatingSystem`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *Binary) GetOsOk() (*OperatingSystem, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *Binary) SetOs(v OperatingSystem)`

SetOs sets Os field to given value.


### GetMaintenance

`func (o *Binary) GetMaintenance() bool`

GetMaintenance returns the Maintenance field if non-nil, zero value otherwise.

### GetMaintenanceOk

`func (o *Binary) GetMaintenanceOk() (*bool, bool)`

GetMaintenanceOk returns a tuple with the Maintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenance

`func (o *Binary) SetMaintenance(v bool)`

SetMaintenance sets Maintenance field to given value.


### GetStatus

`func (o *Binary) GetStatus() BinaryStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Binary) GetStatusOk() (*BinaryStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Binary) SetStatus(v BinaryStatus)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *Binary) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *Binary) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *Binary) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.


### SetStatusMessageNil

`func (o *Binary) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *Binary) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetProgress

`func (o *Binary) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *Binary) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *Binary) SetProgress(v float32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *Binary) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### SetProgressNil

`func (o *Binary) SetProgressNil(b bool)`

 SetProgressNil sets the value for Progress to be an explicit nil

### UnsetProgress
`func (o *Binary) UnsetProgress()`

UnsetProgress ensures that no value is present for Progress, not even an explicit nil
### GetProgressMessage

`func (o *Binary) GetProgressMessage() string`

GetProgressMessage returns the ProgressMessage field if non-nil, zero value otherwise.

### GetProgressMessageOk

`func (o *Binary) GetProgressMessageOk() (*string, bool)`

GetProgressMessageOk returns a tuple with the ProgressMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressMessage

`func (o *Binary) SetProgressMessage(v string)`

SetProgressMessage sets ProgressMessage field to given value.

### HasProgressMessage

`func (o *Binary) HasProgressMessage() bool`

HasProgressMessage returns a boolean if a field has been set.

### SetProgressMessageNil

`func (o *Binary) SetProgressMessageNil(b bool)`

 SetProgressMessageNil sets the value for ProgressMessage to be an explicit nil

### UnsetProgressMessage
`func (o *Binary) UnsetProgressMessage()`

UnsetProgressMessage ensures that no value is present for ProgressMessage, not even an explicit nil
### GetInUse

`func (o *Binary) GetInUse() bool`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *Binary) GetInUseOk() (*bool, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *Binary) SetInUse(v bool)`

SetInUse sets InUse field to given value.


### GetDockerImage

`func (o *Binary) GetDockerImage() DockerImage`

GetDockerImage returns the DockerImage field if non-nil, zero value otherwise.

### GetDockerImageOk

`func (o *Binary) GetDockerImageOk() (*DockerImage, bool)`

GetDockerImageOk returns a tuple with the DockerImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImage

`func (o *Binary) SetDockerImage(v DockerImage)`

SetDockerImage sets DockerImage field to given value.

### HasDockerImage

`func (o *Binary) HasDockerImage() bool`

HasDockerImage returns a boolean if a field has been set.

### GetSteam

`func (o *Binary) GetSteam() Steam`

GetSteam returns the Steam field if non-nil, zero value otherwise.

### GetSteamOk

`func (o *Binary) GetSteamOk() (*Steam, bool)`

GetSteamOk returns a tuple with the Steam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteam

`func (o *Binary) SetSteam(v Steam)`

SetSteam sets Steam field to given value.

### HasSteam

`func (o *Binary) HasSteam() bool`

HasSteam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
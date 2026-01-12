# SteamBranch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the branch | 
**Buildid** | **NullableInt32** | The unique identifier for the build | 
**Pwdrequired** | **NullableBool** | Indicates whether a password is required | 
**Description** | **NullableString** | A brief description of the build | 
**Timeupdated** | **NullableInt32** | Last updated timestamp (in Unix epoch format) | 

## Methods

### NewSteamBranch

`func NewSteamBranch(name string, buildid NullableInt32, pwdrequired NullableBool, description NullableString, timeupdated NullableInt32, ) *SteamBranch`

NewSteamBranch instantiates a new SteamBranch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSteamBranchWithDefaults

`func NewSteamBranchWithDefaults() *SteamBranch`

NewSteamBranchWithDefaults instantiates a new SteamBranch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SteamBranch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SteamBranch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SteamBranch) SetName(v string)`

SetName sets Name field to given value.


### GetBuildid

`func (o *SteamBranch) GetBuildid() int32`

GetBuildid returns the Buildid field if non-nil, zero value otherwise.

### GetBuildidOk

`func (o *SteamBranch) GetBuildidOk() (*int32, bool)`

GetBuildidOk returns a tuple with the Buildid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildid

`func (o *SteamBranch) SetBuildid(v int32)`

SetBuildid sets Buildid field to given value.


### SetBuildidNil

`func (o *SteamBranch) SetBuildidNil(b bool)`

 SetBuildidNil sets the value for Buildid to be an explicit nil

### UnsetBuildid
`func (o *SteamBranch) UnsetBuildid()`

UnsetBuildid ensures that no value is present for Buildid, not even an explicit nil
### GetPwdrequired

`func (o *SteamBranch) GetPwdrequired() bool`

GetPwdrequired returns the Pwdrequired field if non-nil, zero value otherwise.

### GetPwdrequiredOk

`func (o *SteamBranch) GetPwdrequiredOk() (*bool, bool)`

GetPwdrequiredOk returns a tuple with the Pwdrequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwdrequired

`func (o *SteamBranch) SetPwdrequired(v bool)`

SetPwdrequired sets Pwdrequired field to given value.


### SetPwdrequiredNil

`func (o *SteamBranch) SetPwdrequiredNil(b bool)`

 SetPwdrequiredNil sets the value for Pwdrequired to be an explicit nil

### UnsetPwdrequired
`func (o *SteamBranch) UnsetPwdrequired()`

UnsetPwdrequired ensures that no value is present for Pwdrequired, not even an explicit nil
### GetDescription

`func (o *SteamBranch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SteamBranch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SteamBranch) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *SteamBranch) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SteamBranch) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTimeupdated

`func (o *SteamBranch) GetTimeupdated() int32`

GetTimeupdated returns the Timeupdated field if non-nil, zero value otherwise.

### GetTimeupdatedOk

`func (o *SteamBranch) GetTimeupdatedOk() (*int32, bool)`

GetTimeupdatedOk returns a tuple with the Timeupdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeupdated

`func (o *SteamBranch) SetTimeupdated(v int32)`

SetTimeupdated sets Timeupdated field to given value.


### SetTimeupdatedNil

`func (o *SteamBranch) SetTimeupdatedNil(b bool)`

 SetTimeupdatedNil sets the value for Timeupdated to be an explicit nil

### UnsetTimeupdated
`func (o *SteamBranch) UnsetTimeupdated()`

UnsetTimeupdated ensures that no value is present for Timeupdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
# SteamLauncher

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Executable** | **string** | The executable path | 
**Arguments** | **NullableString** | The arguments for the executable | 
**Osarch** | **NullableString** | The OS architecture of the executable | 
**Oslist** | **NullableString** | The OS of the executable | 
**Realm** | **NullableString** | The realm of the executable | 
**Betakey** | **NullableString** | The beta key of the executable | 
**Description** | **NullableString** | The description of the executable | 

## Methods

### NewSteamLauncher

`func NewSteamLauncher(executable string, arguments NullableString, osarch NullableString, oslist NullableString, realm NullableString, betakey NullableString, description NullableString, ) *SteamLauncher`

NewSteamLauncher instantiates a new SteamLauncher object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSteamLauncherWithDefaults

`func NewSteamLauncherWithDefaults() *SteamLauncher`

NewSteamLauncherWithDefaults instantiates a new SteamLauncher object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutable

`func (o *SteamLauncher) GetExecutable() string`

GetExecutable returns the Executable field if non-nil, zero value otherwise.

### GetExecutableOk

`func (o *SteamLauncher) GetExecutableOk() (*string, bool)`

GetExecutableOk returns a tuple with the Executable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutable

`func (o *SteamLauncher) SetExecutable(v string)`

SetExecutable sets Executable field to given value.


### GetArguments

`func (o *SteamLauncher) GetArguments() string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *SteamLauncher) GetArgumentsOk() (*string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *SteamLauncher) SetArguments(v string)`

SetArguments sets Arguments field to given value.


### SetArgumentsNil

`func (o *SteamLauncher) SetArgumentsNil(b bool)`

 SetArgumentsNil sets the value for Arguments to be an explicit nil

### UnsetArguments
`func (o *SteamLauncher) UnsetArguments()`

UnsetArguments ensures that no value is present for Arguments, not even an explicit nil
### GetOsarch

`func (o *SteamLauncher) GetOsarch() string`

GetOsarch returns the Osarch field if non-nil, zero value otherwise.

### GetOsarchOk

`func (o *SteamLauncher) GetOsarchOk() (*string, bool)`

GetOsarchOk returns a tuple with the Osarch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsarch

`func (o *SteamLauncher) SetOsarch(v string)`

SetOsarch sets Osarch field to given value.


### SetOsarchNil

`func (o *SteamLauncher) SetOsarchNil(b bool)`

 SetOsarchNil sets the value for Osarch to be an explicit nil

### UnsetOsarch
`func (o *SteamLauncher) UnsetOsarch()`

UnsetOsarch ensures that no value is present for Osarch, not even an explicit nil
### GetOslist

`func (o *SteamLauncher) GetOslist() string`

GetOslist returns the Oslist field if non-nil, zero value otherwise.

### GetOslistOk

`func (o *SteamLauncher) GetOslistOk() (*string, bool)`

GetOslistOk returns a tuple with the Oslist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOslist

`func (o *SteamLauncher) SetOslist(v string)`

SetOslist sets Oslist field to given value.


### SetOslistNil

`func (o *SteamLauncher) SetOslistNil(b bool)`

 SetOslistNil sets the value for Oslist to be an explicit nil

### UnsetOslist
`func (o *SteamLauncher) UnsetOslist()`

UnsetOslist ensures that no value is present for Oslist, not even an explicit nil
### GetRealm

`func (o *SteamLauncher) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *SteamLauncher) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *SteamLauncher) SetRealm(v string)`

SetRealm sets Realm field to given value.


### SetRealmNil

`func (o *SteamLauncher) SetRealmNil(b bool)`

 SetRealmNil sets the value for Realm to be an explicit nil

### UnsetRealm
`func (o *SteamLauncher) UnsetRealm()`

UnsetRealm ensures that no value is present for Realm, not even an explicit nil
### GetBetakey

`func (o *SteamLauncher) GetBetakey() string`

GetBetakey returns the Betakey field if non-nil, zero value otherwise.

### GetBetakeyOk

`func (o *SteamLauncher) GetBetakeyOk() (*string, bool)`

GetBetakeyOk returns a tuple with the Betakey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetakey

`func (o *SteamLauncher) SetBetakey(v string)`

SetBetakey sets Betakey field to given value.


### SetBetakeyNil

`func (o *SteamLauncher) SetBetakeyNil(b bool)`

 SetBetakeyNil sets the value for Betakey to be an explicit nil

### UnsetBetakey
`func (o *SteamLauncher) UnsetBetakey()`

UnsetBetakey ensures that no value is present for Betakey, not even an explicit nil
### GetDescription

`func (o *SteamLauncher) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SteamLauncher) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SteamLauncher) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *SteamLauncher) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SteamLauncher) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
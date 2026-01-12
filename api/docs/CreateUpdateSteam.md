# CreateUpdateSteam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SteamAppId** | **int32** | The Steam App ID | 
**Branch** | **string** | The branch to use | 
**Password** | **NullableString** | The password to use | 
**Command** | **string** | The command to execute | 
**SteamcmdUsername** | **NullableString** | The steam account username | 
**SteamcmdPassword** | **NullableString** | The steam account password | 
**Headful** | **bool** | Whether or not the binary needs a graphical interface | 
**RequestLicense** | **bool** | Whether or not the the license agreement needs to be requested | 
**Runtime** | [**SteamRuntime**](SteamRuntime.md) | The steam runtime the server depends on | 
**AdditionalPackages** | **NullableString** | A space-separated list of additional packages to install | 
**Unpublished** | **bool** | Whether or not the steamworks app is unpublished or published | 

## Methods

### NewCreateUpdateSteam

`func NewCreateUpdateSteam(steamAppId int32, branch string, password NullableString, command string, steamcmdUsername NullableString, steamcmdPassword NullableString, headful bool, requestLicense bool, runtime SteamRuntime, additionalPackages NullableString, unpublished bool, ) *CreateUpdateSteam`

NewCreateUpdateSteam instantiates a new CreateUpdateSteam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpdateSteamWithDefaults

`func NewCreateUpdateSteamWithDefaults() *CreateUpdateSteam`

NewCreateUpdateSteamWithDefaults instantiates a new CreateUpdateSteam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSteamAppId

`func (o *CreateUpdateSteam) GetSteamAppId() int32`

GetSteamAppId returns the SteamAppId field if non-nil, zero value otherwise.

### GetSteamAppIdOk

`func (o *CreateUpdateSteam) GetSteamAppIdOk() (*int32, bool)`

GetSteamAppIdOk returns a tuple with the SteamAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamAppId

`func (o *CreateUpdateSteam) SetSteamAppId(v int32)`

SetSteamAppId sets SteamAppId field to given value.


### GetBranch

`func (o *CreateUpdateSteam) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *CreateUpdateSteam) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *CreateUpdateSteam) SetBranch(v string)`

SetBranch sets Branch field to given value.


### GetPassword

`func (o *CreateUpdateSteam) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateUpdateSteam) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateUpdateSteam) SetPassword(v string)`

SetPassword sets Password field to given value.


### SetPasswordNil

`func (o *CreateUpdateSteam) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *CreateUpdateSteam) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetCommand

`func (o *CreateUpdateSteam) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *CreateUpdateSteam) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *CreateUpdateSteam) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetSteamcmdUsername

`func (o *CreateUpdateSteam) GetSteamcmdUsername() string`

GetSteamcmdUsername returns the SteamcmdUsername field if non-nil, zero value otherwise.

### GetSteamcmdUsernameOk

`func (o *CreateUpdateSteam) GetSteamcmdUsernameOk() (*string, bool)`

GetSteamcmdUsernameOk returns a tuple with the SteamcmdUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamcmdUsername

`func (o *CreateUpdateSteam) SetSteamcmdUsername(v string)`

SetSteamcmdUsername sets SteamcmdUsername field to given value.


### SetSteamcmdUsernameNil

`func (o *CreateUpdateSteam) SetSteamcmdUsernameNil(b bool)`

 SetSteamcmdUsernameNil sets the value for SteamcmdUsername to be an explicit nil

### UnsetSteamcmdUsername
`func (o *CreateUpdateSteam) UnsetSteamcmdUsername()`

UnsetSteamcmdUsername ensures that no value is present for SteamcmdUsername, not even an explicit nil
### GetSteamcmdPassword

`func (o *CreateUpdateSteam) GetSteamcmdPassword() string`

GetSteamcmdPassword returns the SteamcmdPassword field if non-nil, zero value otherwise.

### GetSteamcmdPasswordOk

`func (o *CreateUpdateSteam) GetSteamcmdPasswordOk() (*string, bool)`

GetSteamcmdPasswordOk returns a tuple with the SteamcmdPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamcmdPassword

`func (o *CreateUpdateSteam) SetSteamcmdPassword(v string)`

SetSteamcmdPassword sets SteamcmdPassword field to given value.


### SetSteamcmdPasswordNil

`func (o *CreateUpdateSteam) SetSteamcmdPasswordNil(b bool)`

 SetSteamcmdPasswordNil sets the value for SteamcmdPassword to be an explicit nil

### UnsetSteamcmdPassword
`func (o *CreateUpdateSteam) UnsetSteamcmdPassword()`

UnsetSteamcmdPassword ensures that no value is present for SteamcmdPassword, not even an explicit nil
### GetHeadful

`func (o *CreateUpdateSteam) GetHeadful() bool`

GetHeadful returns the Headful field if non-nil, zero value otherwise.

### GetHeadfulOk

`func (o *CreateUpdateSteam) GetHeadfulOk() (*bool, bool)`

GetHeadfulOk returns a tuple with the Headful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadful

`func (o *CreateUpdateSteam) SetHeadful(v bool)`

SetHeadful sets Headful field to given value.


### GetRequestLicense

`func (o *CreateUpdateSteam) GetRequestLicense() bool`

GetRequestLicense returns the RequestLicense field if non-nil, zero value otherwise.

### GetRequestLicenseOk

`func (o *CreateUpdateSteam) GetRequestLicenseOk() (*bool, bool)`

GetRequestLicenseOk returns a tuple with the RequestLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestLicense

`func (o *CreateUpdateSteam) SetRequestLicense(v bool)`

SetRequestLicense sets RequestLicense field to given value.


### GetRuntime

`func (o *CreateUpdateSteam) GetRuntime() SteamRuntime`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *CreateUpdateSteam) GetRuntimeOk() (*SteamRuntime, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *CreateUpdateSteam) SetRuntime(v SteamRuntime)`

SetRuntime sets Runtime field to given value.


### GetAdditionalPackages

`func (o *CreateUpdateSteam) GetAdditionalPackages() string`

GetAdditionalPackages returns the AdditionalPackages field if non-nil, zero value otherwise.

### GetAdditionalPackagesOk

`func (o *CreateUpdateSteam) GetAdditionalPackagesOk() (*string, bool)`

GetAdditionalPackagesOk returns a tuple with the AdditionalPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPackages

`func (o *CreateUpdateSteam) SetAdditionalPackages(v string)`

SetAdditionalPackages sets AdditionalPackages field to given value.


### SetAdditionalPackagesNil

`func (o *CreateUpdateSteam) SetAdditionalPackagesNil(b bool)`

 SetAdditionalPackagesNil sets the value for AdditionalPackages to be an explicit nil

### UnsetAdditionalPackages
`func (o *CreateUpdateSteam) UnsetAdditionalPackages()`

UnsetAdditionalPackages ensures that no value is present for AdditionalPackages, not even an explicit nil
### GetUnpublished

`func (o *CreateUpdateSteam) GetUnpublished() bool`

GetUnpublished returns the Unpublished field if non-nil, zero value otherwise.

### GetUnpublishedOk

`func (o *CreateUpdateSteam) GetUnpublishedOk() (*bool, bool)`

GetUnpublishedOk returns a tuple with the Unpublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpublished

`func (o *CreateUpdateSteam) SetUnpublished(v bool)`

SetUnpublished sets Unpublished field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
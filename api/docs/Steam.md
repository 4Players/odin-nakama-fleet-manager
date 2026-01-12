# Steam

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
**InUse** | **bool** | Indicates whether the steam binary is currently in use | 

## Methods

### NewSteam

`func NewSteam(steamAppId int32, branch string, password NullableString, command string, steamcmdUsername NullableString, steamcmdPassword NullableString, headful bool, requestLicense bool, runtime SteamRuntime, additionalPackages NullableString, unpublished bool, inUse bool, ) *Steam`

NewSteam instantiates a new Steam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSteamWithDefaults

`func NewSteamWithDefaults() *Steam`

NewSteamWithDefaults instantiates a new Steam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSteamAppId

`func (o *Steam) GetSteamAppId() int32`

GetSteamAppId returns the SteamAppId field if non-nil, zero value otherwise.

### GetSteamAppIdOk

`func (o *Steam) GetSteamAppIdOk() (*int32, bool)`

GetSteamAppIdOk returns a tuple with the SteamAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamAppId

`func (o *Steam) SetSteamAppId(v int32)`

SetSteamAppId sets SteamAppId field to given value.


### GetBranch

`func (o *Steam) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *Steam) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *Steam) SetBranch(v string)`

SetBranch sets Branch field to given value.


### GetPassword

`func (o *Steam) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Steam) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Steam) SetPassword(v string)`

SetPassword sets Password field to given value.


### SetPasswordNil

`func (o *Steam) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *Steam) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetCommand

`func (o *Steam) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *Steam) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *Steam) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetSteamcmdUsername

`func (o *Steam) GetSteamcmdUsername() string`

GetSteamcmdUsername returns the SteamcmdUsername field if non-nil, zero value otherwise.

### GetSteamcmdUsernameOk

`func (o *Steam) GetSteamcmdUsernameOk() (*string, bool)`

GetSteamcmdUsernameOk returns a tuple with the SteamcmdUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamcmdUsername

`func (o *Steam) SetSteamcmdUsername(v string)`

SetSteamcmdUsername sets SteamcmdUsername field to given value.


### SetSteamcmdUsernameNil

`func (o *Steam) SetSteamcmdUsernameNil(b bool)`

 SetSteamcmdUsernameNil sets the value for SteamcmdUsername to be an explicit nil

### UnsetSteamcmdUsername
`func (o *Steam) UnsetSteamcmdUsername()`

UnsetSteamcmdUsername ensures that no value is present for SteamcmdUsername, not even an explicit nil
### GetSteamcmdPassword

`func (o *Steam) GetSteamcmdPassword() string`

GetSteamcmdPassword returns the SteamcmdPassword field if non-nil, zero value otherwise.

### GetSteamcmdPasswordOk

`func (o *Steam) GetSteamcmdPasswordOk() (*string, bool)`

GetSteamcmdPasswordOk returns a tuple with the SteamcmdPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamcmdPassword

`func (o *Steam) SetSteamcmdPassword(v string)`

SetSteamcmdPassword sets SteamcmdPassword field to given value.


### SetSteamcmdPasswordNil

`func (o *Steam) SetSteamcmdPasswordNil(b bool)`

 SetSteamcmdPasswordNil sets the value for SteamcmdPassword to be an explicit nil

### UnsetSteamcmdPassword
`func (o *Steam) UnsetSteamcmdPassword()`

UnsetSteamcmdPassword ensures that no value is present for SteamcmdPassword, not even an explicit nil
### GetHeadful

`func (o *Steam) GetHeadful() bool`

GetHeadful returns the Headful field if non-nil, zero value otherwise.

### GetHeadfulOk

`func (o *Steam) GetHeadfulOk() (*bool, bool)`

GetHeadfulOk returns a tuple with the Headful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadful

`func (o *Steam) SetHeadful(v bool)`

SetHeadful sets Headful field to given value.


### GetRequestLicense

`func (o *Steam) GetRequestLicense() bool`

GetRequestLicense returns the RequestLicense field if non-nil, zero value otherwise.

### GetRequestLicenseOk

`func (o *Steam) GetRequestLicenseOk() (*bool, bool)`

GetRequestLicenseOk returns a tuple with the RequestLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestLicense

`func (o *Steam) SetRequestLicense(v bool)`

SetRequestLicense sets RequestLicense field to given value.


### GetRuntime

`func (o *Steam) GetRuntime() SteamRuntime`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *Steam) GetRuntimeOk() (*SteamRuntime, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *Steam) SetRuntime(v SteamRuntime)`

SetRuntime sets Runtime field to given value.


### GetAdditionalPackages

`func (o *Steam) GetAdditionalPackages() string`

GetAdditionalPackages returns the AdditionalPackages field if non-nil, zero value otherwise.

### GetAdditionalPackagesOk

`func (o *Steam) GetAdditionalPackagesOk() (*string, bool)`

GetAdditionalPackagesOk returns a tuple with the AdditionalPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPackages

`func (o *Steam) SetAdditionalPackages(v string)`

SetAdditionalPackages sets AdditionalPackages field to given value.


### SetAdditionalPackagesNil

`func (o *Steam) SetAdditionalPackagesNil(b bool)`

 SetAdditionalPackagesNil sets the value for AdditionalPackages to be an explicit nil

### UnsetAdditionalPackages
`func (o *Steam) UnsetAdditionalPackages()`

UnsetAdditionalPackages ensures that no value is present for AdditionalPackages, not even an explicit nil
### GetUnpublished

`func (o *Steam) GetUnpublished() bool`

GetUnpublished returns the Unpublished field if non-nil, zero value otherwise.

### GetUnpublishedOk

`func (o *Steam) GetUnpublishedOk() (*bool, bool)`

GetUnpublishedOk returns a tuple with the Unpublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpublished

`func (o *Steam) SetUnpublished(v bool)`

SetUnpublished sets Unpublished field to given value.


### GetInUse

`func (o *Steam) GetInUse() bool`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *Steam) GetInUseOk() (*bool, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *Steam) SetInUse(v bool)`

SetInUse sets InUse field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
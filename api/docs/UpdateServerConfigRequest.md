# UpdateServerConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the server configuration | 
**Args** | Pointer to **NullableString** | The arguments to pass to the command (overrides CMD of the Dockerfile) | [optional] 
**Command** | Pointer to **NullableString** | The command to run in the container (overrides ENTRYPOINT of the Dockerfile) | [optional] 
**Notes** | Pointer to **NullableString** | The notes of the server config - to keep track of things and to inform colleagues | [optional] 
**BinaryId** | **int32** | The binary id of the server configuration | 
**ResourcePackageSlug** | **string** | The slug of the resource package | 
**ConfigFiles** | Pointer to [**[]ConfigFile**](ConfigFile.md) | The config files used in this server configuration | [optional] 
**SecretFiles** | Pointer to [**[]SecretFile**](SecretFile.md) | The secret files used in this server configuration | [optional] 
**RestartPolicy** | Pointer to [**RestartPolicy**](RestartPolicy.md) | The restart policy of the server configuration | [optional] 
**Env** | Pointer to [**[]EnvironmentVariableDefinition**](EnvironmentVariableDefinition.md) | The environment variables used in this server configuration | [optional] 
**Mounts** | Pointer to [**[]Mount**](Mount.md) | The mounts used in this server configuration | [optional] 
**Ports** | Pointer to [**[]PortDefinition**](PortDefinition.md) | The port definitions | [optional] 

## Methods

### NewUpdateServerConfigRequest

`func NewUpdateServerConfigRequest(name string, binaryId int32, resourcePackageSlug string, ) *UpdateServerConfigRequest`

NewUpdateServerConfigRequest instantiates a new UpdateServerConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServerConfigRequestWithDefaults

`func NewUpdateServerConfigRequestWithDefaults() *UpdateServerConfigRequest`

NewUpdateServerConfigRequestWithDefaults instantiates a new UpdateServerConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateServerConfigRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateServerConfigRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateServerConfigRequest) SetName(v string)`

SetName sets Name field to given value.


### GetArgs

`func (o *UpdateServerConfigRequest) GetArgs() string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *UpdateServerConfigRequest) GetArgsOk() (*string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *UpdateServerConfigRequest) SetArgs(v string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *UpdateServerConfigRequest) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### SetArgsNil

`func (o *UpdateServerConfigRequest) SetArgsNil(b bool)`

 SetArgsNil sets the value for Args to be an explicit nil

### UnsetArgs
`func (o *UpdateServerConfigRequest) UnsetArgs()`

UnsetArgs ensures that no value is present for Args, not even an explicit nil
### GetCommand

`func (o *UpdateServerConfigRequest) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *UpdateServerConfigRequest) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *UpdateServerConfigRequest) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *UpdateServerConfigRequest) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### SetCommandNil

`func (o *UpdateServerConfigRequest) SetCommandNil(b bool)`

 SetCommandNil sets the value for Command to be an explicit nil

### UnsetCommand
`func (o *UpdateServerConfigRequest) UnsetCommand()`

UnsetCommand ensures that no value is present for Command, not even an explicit nil
### GetNotes

`func (o *UpdateServerConfigRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *UpdateServerConfigRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *UpdateServerConfigRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *UpdateServerConfigRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *UpdateServerConfigRequest) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *UpdateServerConfigRequest) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetBinaryId

`func (o *UpdateServerConfigRequest) GetBinaryId() int32`

GetBinaryId returns the BinaryId field if non-nil, zero value otherwise.

### GetBinaryIdOk

`func (o *UpdateServerConfigRequest) GetBinaryIdOk() (*int32, bool)`

GetBinaryIdOk returns a tuple with the BinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryId

`func (o *UpdateServerConfigRequest) SetBinaryId(v int32)`

SetBinaryId sets BinaryId field to given value.


### GetResourcePackageSlug

`func (o *UpdateServerConfigRequest) GetResourcePackageSlug() string`

GetResourcePackageSlug returns the ResourcePackageSlug field if non-nil, zero value otherwise.

### GetResourcePackageSlugOk

`func (o *UpdateServerConfigRequest) GetResourcePackageSlugOk() (*string, bool)`

GetResourcePackageSlugOk returns a tuple with the ResourcePackageSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePackageSlug

`func (o *UpdateServerConfigRequest) SetResourcePackageSlug(v string)`

SetResourcePackageSlug sets ResourcePackageSlug field to given value.


### GetConfigFiles

`func (o *UpdateServerConfigRequest) GetConfigFiles() []ConfigFile`

GetConfigFiles returns the ConfigFiles field if non-nil, zero value otherwise.

### GetConfigFilesOk

`func (o *UpdateServerConfigRequest) GetConfigFilesOk() (*[]ConfigFile, bool)`

GetConfigFilesOk returns a tuple with the ConfigFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigFiles

`func (o *UpdateServerConfigRequest) SetConfigFiles(v []ConfigFile)`

SetConfigFiles sets ConfigFiles field to given value.

### HasConfigFiles

`func (o *UpdateServerConfigRequest) HasConfigFiles() bool`

HasConfigFiles returns a boolean if a field has been set.

### GetSecretFiles

`func (o *UpdateServerConfigRequest) GetSecretFiles() []SecretFile`

GetSecretFiles returns the SecretFiles field if non-nil, zero value otherwise.

### GetSecretFilesOk

`func (o *UpdateServerConfigRequest) GetSecretFilesOk() (*[]SecretFile, bool)`

GetSecretFilesOk returns a tuple with the SecretFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretFiles

`func (o *UpdateServerConfigRequest) SetSecretFiles(v []SecretFile)`

SetSecretFiles sets SecretFiles field to given value.

### HasSecretFiles

`func (o *UpdateServerConfigRequest) HasSecretFiles() bool`

HasSecretFiles returns a boolean if a field has been set.

### GetRestartPolicy

`func (o *UpdateServerConfigRequest) GetRestartPolicy() RestartPolicy`

GetRestartPolicy returns the RestartPolicy field if non-nil, zero value otherwise.

### GetRestartPolicyOk

`func (o *UpdateServerConfigRequest) GetRestartPolicyOk() (*RestartPolicy, bool)`

GetRestartPolicyOk returns a tuple with the RestartPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartPolicy

`func (o *UpdateServerConfigRequest) SetRestartPolicy(v RestartPolicy)`

SetRestartPolicy sets RestartPolicy field to given value.

### HasRestartPolicy

`func (o *UpdateServerConfigRequest) HasRestartPolicy() bool`

HasRestartPolicy returns a boolean if a field has been set.

### GetEnv

`func (o *UpdateServerConfigRequest) GetEnv() []EnvironmentVariableDefinition`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *UpdateServerConfigRequest) GetEnvOk() (*[]EnvironmentVariableDefinition, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *UpdateServerConfigRequest) SetEnv(v []EnvironmentVariableDefinition)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *UpdateServerConfigRequest) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetMounts

`func (o *UpdateServerConfigRequest) GetMounts() []Mount`

GetMounts returns the Mounts field if non-nil, zero value otherwise.

### GetMountsOk

`func (o *UpdateServerConfigRequest) GetMountsOk() (*[]Mount, bool)`

GetMountsOk returns a tuple with the Mounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounts

`func (o *UpdateServerConfigRequest) SetMounts(v []Mount)`

SetMounts sets Mounts field to given value.

### HasMounts

`func (o *UpdateServerConfigRequest) HasMounts() bool`

HasMounts returns a boolean if a field has been set.

### GetPorts

`func (o *UpdateServerConfigRequest) GetPorts() []PortDefinition`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *UpdateServerConfigRequest) GetPortsOk() (*[]PortDefinition, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *UpdateServerConfigRequest) SetPorts(v []PortDefinition)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *UpdateServerConfigRequest) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
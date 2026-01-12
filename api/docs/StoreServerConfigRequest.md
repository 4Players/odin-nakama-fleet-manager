# StoreServerConfigRequest

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

### NewStoreServerConfigRequest

`func NewStoreServerConfigRequest(name string, binaryId int32, resourcePackageSlug string, ) *StoreServerConfigRequest`

NewStoreServerConfigRequest instantiates a new StoreServerConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreServerConfigRequestWithDefaults

`func NewStoreServerConfigRequestWithDefaults() *StoreServerConfigRequest`

NewStoreServerConfigRequestWithDefaults instantiates a new StoreServerConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StoreServerConfigRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoreServerConfigRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoreServerConfigRequest) SetName(v string)`

SetName sets Name field to given value.


### GetArgs

`func (o *StoreServerConfigRequest) GetArgs() string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *StoreServerConfigRequest) GetArgsOk() (*string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *StoreServerConfigRequest) SetArgs(v string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *StoreServerConfigRequest) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### SetArgsNil

`func (o *StoreServerConfigRequest) SetArgsNil(b bool)`

 SetArgsNil sets the value for Args to be an explicit nil

### UnsetArgs
`func (o *StoreServerConfigRequest) UnsetArgs()`

UnsetArgs ensures that no value is present for Args, not even an explicit nil
### GetCommand

`func (o *StoreServerConfigRequest) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *StoreServerConfigRequest) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *StoreServerConfigRequest) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *StoreServerConfigRequest) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### SetCommandNil

`func (o *StoreServerConfigRequest) SetCommandNil(b bool)`

 SetCommandNil sets the value for Command to be an explicit nil

### UnsetCommand
`func (o *StoreServerConfigRequest) UnsetCommand()`

UnsetCommand ensures that no value is present for Command, not even an explicit nil
### GetNotes

`func (o *StoreServerConfigRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *StoreServerConfigRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *StoreServerConfigRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *StoreServerConfigRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *StoreServerConfigRequest) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *StoreServerConfigRequest) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetBinaryId

`func (o *StoreServerConfigRequest) GetBinaryId() int32`

GetBinaryId returns the BinaryId field if non-nil, zero value otherwise.

### GetBinaryIdOk

`func (o *StoreServerConfigRequest) GetBinaryIdOk() (*int32, bool)`

GetBinaryIdOk returns a tuple with the BinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryId

`func (o *StoreServerConfigRequest) SetBinaryId(v int32)`

SetBinaryId sets BinaryId field to given value.


### GetResourcePackageSlug

`func (o *StoreServerConfigRequest) GetResourcePackageSlug() string`

GetResourcePackageSlug returns the ResourcePackageSlug field if non-nil, zero value otherwise.

### GetResourcePackageSlugOk

`func (o *StoreServerConfigRequest) GetResourcePackageSlugOk() (*string, bool)`

GetResourcePackageSlugOk returns a tuple with the ResourcePackageSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePackageSlug

`func (o *StoreServerConfigRequest) SetResourcePackageSlug(v string)`

SetResourcePackageSlug sets ResourcePackageSlug field to given value.


### GetConfigFiles

`func (o *StoreServerConfigRequest) GetConfigFiles() []ConfigFile`

GetConfigFiles returns the ConfigFiles field if non-nil, zero value otherwise.

### GetConfigFilesOk

`func (o *StoreServerConfigRequest) GetConfigFilesOk() (*[]ConfigFile, bool)`

GetConfigFilesOk returns a tuple with the ConfigFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigFiles

`func (o *StoreServerConfigRequest) SetConfigFiles(v []ConfigFile)`

SetConfigFiles sets ConfigFiles field to given value.

### HasConfigFiles

`func (o *StoreServerConfigRequest) HasConfigFiles() bool`

HasConfigFiles returns a boolean if a field has been set.

### GetSecretFiles

`func (o *StoreServerConfigRequest) GetSecretFiles() []SecretFile`

GetSecretFiles returns the SecretFiles field if non-nil, zero value otherwise.

### GetSecretFilesOk

`func (o *StoreServerConfigRequest) GetSecretFilesOk() (*[]SecretFile, bool)`

GetSecretFilesOk returns a tuple with the SecretFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretFiles

`func (o *StoreServerConfigRequest) SetSecretFiles(v []SecretFile)`

SetSecretFiles sets SecretFiles field to given value.

### HasSecretFiles

`func (o *StoreServerConfigRequest) HasSecretFiles() bool`

HasSecretFiles returns a boolean if a field has been set.

### GetRestartPolicy

`func (o *StoreServerConfigRequest) GetRestartPolicy() RestartPolicy`

GetRestartPolicy returns the RestartPolicy field if non-nil, zero value otherwise.

### GetRestartPolicyOk

`func (o *StoreServerConfigRequest) GetRestartPolicyOk() (*RestartPolicy, bool)`

GetRestartPolicyOk returns a tuple with the RestartPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartPolicy

`func (o *StoreServerConfigRequest) SetRestartPolicy(v RestartPolicy)`

SetRestartPolicy sets RestartPolicy field to given value.

### HasRestartPolicy

`func (o *StoreServerConfigRequest) HasRestartPolicy() bool`

HasRestartPolicy returns a boolean if a field has been set.

### GetEnv

`func (o *StoreServerConfigRequest) GetEnv() []EnvironmentVariableDefinition`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *StoreServerConfigRequest) GetEnvOk() (*[]EnvironmentVariableDefinition, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *StoreServerConfigRequest) SetEnv(v []EnvironmentVariableDefinition)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *StoreServerConfigRequest) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetMounts

`func (o *StoreServerConfigRequest) GetMounts() []Mount`

GetMounts returns the Mounts field if non-nil, zero value otherwise.

### GetMountsOk

`func (o *StoreServerConfigRequest) GetMountsOk() (*[]Mount, bool)`

GetMountsOk returns a tuple with the Mounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounts

`func (o *StoreServerConfigRequest) SetMounts(v []Mount)`

SetMounts sets Mounts field to given value.

### HasMounts

`func (o *StoreServerConfigRequest) HasMounts() bool`

HasMounts returns a boolean if a field has been set.

### GetPorts

`func (o *StoreServerConfigRequest) GetPorts() []PortDefinition`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *StoreServerConfigRequest) GetPortsOk() (*[]PortDefinition, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *StoreServerConfigRequest) SetPorts(v []PortDefinition)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *StoreServerConfigRequest) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
# ServerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The id of the server configuration | 
**BinaryId** | **int32** | The binary id of the server configuration | 
**AppId** | **int32** | The app id of the server configuration | 
**Name** | **string** | The name of the server config | 
**Command** | **NullableString** | The command to run in the container (overrides ENTRYPOINT of the Dockerfile) | 
**Args** | **NullableString** | The arguments to pass to the command (overrides CMD of the Dockerfile) | 
**Notes** | **NullableString** | The notes of the server config - to keep track of things and to inform colleagues | 
**Status** | [**ServerConfigStatus**](ServerConfigStatus.md) | The current status | 
**StatusMessage** | **NullableString** | An optional message | 
**Maintenance** | **bool** | Indicates if the server config is under maintenance | 
**ResourcePackageSlug** | **NullableString** | The slug of the resource package used in this server config | 
**InUse** | **bool** | Indicates whether the server config is currently in use | 
**RestartPolicy** | [**RestartPolicy**](RestartPolicy.md) | The policy used to restart this server | 
**Env** | [**[]EnvironmentVariableDefinition**](EnvironmentVariableDefinition.md) | The environment variable definitions to be used in this config | 
**Mounts** | [**[]Mount**](Mount.md) | The mounts to use | 
**Ports** | [**[]PortDefinition**](PortDefinition.md) | The ports to expose | 
**ConfigFiles** | [**[]ConfigFile**](ConfigFile.md) | The config files to use | 
**SecretFiles** | [**[]SecretFile**](SecretFile.md) | The secret files to use | 
**Binary** | Pointer to [**Binary**](Binary.md) | The image that is used in this server config | [optional] 

## Methods

### NewServerConfig

`func NewServerConfig(id int32, binaryId int32, appId int32, name string, command NullableString, args NullableString, notes NullableString, status ServerConfigStatus, statusMessage NullableString, maintenance bool, resourcePackageSlug NullableString, inUse bool, restartPolicy RestartPolicy, env []EnvironmentVariableDefinition, mounts []Mount, ports []PortDefinition, configFiles []ConfigFile, secretFiles []SecretFile, ) *ServerConfig`

NewServerConfig instantiates a new ServerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConfigWithDefaults

`func NewServerConfigWithDefaults() *ServerConfig`

NewServerConfigWithDefaults instantiates a new ServerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerConfig) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerConfig) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerConfig) SetId(v int32)`

SetId sets Id field to given value.


### GetBinaryId

`func (o *ServerConfig) GetBinaryId() int32`

GetBinaryId returns the BinaryId field if non-nil, zero value otherwise.

### GetBinaryIdOk

`func (o *ServerConfig) GetBinaryIdOk() (*int32, bool)`

GetBinaryIdOk returns a tuple with the BinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryId

`func (o *ServerConfig) SetBinaryId(v int32)`

SetBinaryId sets BinaryId field to given value.


### GetAppId

`func (o *ServerConfig) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ServerConfig) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ServerConfig) SetAppId(v int32)`

SetAppId sets AppId field to given value.


### GetName

`func (o *ServerConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerConfig) SetName(v string)`

SetName sets Name field to given value.


### GetCommand

`func (o *ServerConfig) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *ServerConfig) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *ServerConfig) SetCommand(v string)`

SetCommand sets Command field to given value.


### SetCommandNil

`func (o *ServerConfig) SetCommandNil(b bool)`

 SetCommandNil sets the value for Command to be an explicit nil

### UnsetCommand
`func (o *ServerConfig) UnsetCommand()`

UnsetCommand ensures that no value is present for Command, not even an explicit nil
### GetArgs

`func (o *ServerConfig) GetArgs() string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *ServerConfig) GetArgsOk() (*string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *ServerConfig) SetArgs(v string)`

SetArgs sets Args field to given value.


### SetArgsNil

`func (o *ServerConfig) SetArgsNil(b bool)`

 SetArgsNil sets the value for Args to be an explicit nil

### UnsetArgs
`func (o *ServerConfig) UnsetArgs()`

UnsetArgs ensures that no value is present for Args, not even an explicit nil
### GetNotes

`func (o *ServerConfig) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ServerConfig) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ServerConfig) SetNotes(v string)`

SetNotes sets Notes field to given value.


### SetNotesNil

`func (o *ServerConfig) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *ServerConfig) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetStatus

`func (o *ServerConfig) GetStatus() ServerConfigStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServerConfig) GetStatusOk() (*ServerConfigStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServerConfig) SetStatus(v ServerConfigStatus)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *ServerConfig) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ServerConfig) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ServerConfig) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.


### SetStatusMessageNil

`func (o *ServerConfig) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *ServerConfig) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetMaintenance

`func (o *ServerConfig) GetMaintenance() bool`

GetMaintenance returns the Maintenance field if non-nil, zero value otherwise.

### GetMaintenanceOk

`func (o *ServerConfig) GetMaintenanceOk() (*bool, bool)`

GetMaintenanceOk returns a tuple with the Maintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenance

`func (o *ServerConfig) SetMaintenance(v bool)`

SetMaintenance sets Maintenance field to given value.


### GetResourcePackageSlug

`func (o *ServerConfig) GetResourcePackageSlug() string`

GetResourcePackageSlug returns the ResourcePackageSlug field if non-nil, zero value otherwise.

### GetResourcePackageSlugOk

`func (o *ServerConfig) GetResourcePackageSlugOk() (*string, bool)`

GetResourcePackageSlugOk returns a tuple with the ResourcePackageSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePackageSlug

`func (o *ServerConfig) SetResourcePackageSlug(v string)`

SetResourcePackageSlug sets ResourcePackageSlug field to given value.


### SetResourcePackageSlugNil

`func (o *ServerConfig) SetResourcePackageSlugNil(b bool)`

 SetResourcePackageSlugNil sets the value for ResourcePackageSlug to be an explicit nil

### UnsetResourcePackageSlug
`func (o *ServerConfig) UnsetResourcePackageSlug()`

UnsetResourcePackageSlug ensures that no value is present for ResourcePackageSlug, not even an explicit nil
### GetInUse

`func (o *ServerConfig) GetInUse() bool`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *ServerConfig) GetInUseOk() (*bool, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *ServerConfig) SetInUse(v bool)`

SetInUse sets InUse field to given value.


### GetRestartPolicy

`func (o *ServerConfig) GetRestartPolicy() RestartPolicy`

GetRestartPolicy returns the RestartPolicy field if non-nil, zero value otherwise.

### GetRestartPolicyOk

`func (o *ServerConfig) GetRestartPolicyOk() (*RestartPolicy, bool)`

GetRestartPolicyOk returns a tuple with the RestartPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartPolicy

`func (o *ServerConfig) SetRestartPolicy(v RestartPolicy)`

SetRestartPolicy sets RestartPolicy field to given value.


### GetEnv

`func (o *ServerConfig) GetEnv() []EnvironmentVariableDefinition`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *ServerConfig) GetEnvOk() (*[]EnvironmentVariableDefinition, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *ServerConfig) SetEnv(v []EnvironmentVariableDefinition)`

SetEnv sets Env field to given value.


### GetMounts

`func (o *ServerConfig) GetMounts() []Mount`

GetMounts returns the Mounts field if non-nil, zero value otherwise.

### GetMountsOk

`func (o *ServerConfig) GetMountsOk() (*[]Mount, bool)`

GetMountsOk returns a tuple with the Mounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounts

`func (o *ServerConfig) SetMounts(v []Mount)`

SetMounts sets Mounts field to given value.


### GetPorts

`func (o *ServerConfig) GetPorts() []PortDefinition`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ServerConfig) GetPortsOk() (*[]PortDefinition, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ServerConfig) SetPorts(v []PortDefinition)`

SetPorts sets Ports field to given value.


### GetConfigFiles

`func (o *ServerConfig) GetConfigFiles() []ConfigFile`

GetConfigFiles returns the ConfigFiles field if non-nil, zero value otherwise.

### GetConfigFilesOk

`func (o *ServerConfig) GetConfigFilesOk() (*[]ConfigFile, bool)`

GetConfigFilesOk returns a tuple with the ConfigFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigFiles

`func (o *ServerConfig) SetConfigFiles(v []ConfigFile)`

SetConfigFiles sets ConfigFiles field to given value.


### GetSecretFiles

`func (o *ServerConfig) GetSecretFiles() []SecretFile`

GetSecretFiles returns the SecretFiles field if non-nil, zero value otherwise.

### GetSecretFilesOk

`func (o *ServerConfig) GetSecretFilesOk() (*[]SecretFile, bool)`

GetSecretFilesOk returns a tuple with the SecretFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretFiles

`func (o *ServerConfig) SetSecretFiles(v []SecretFile)`

SetSecretFiles sets SecretFiles field to given value.


### GetBinary

`func (o *ServerConfig) GetBinary() Binary`

GetBinary returns the Binary field if non-nil, zero value otherwise.

### GetBinaryOk

`func (o *ServerConfig) GetBinaryOk() (*Binary, bool)`

GetBinaryOk returns a tuple with the Binary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinary

`func (o *ServerConfig) SetBinary(v Binary)`

SetBinary sets Binary field to given value.

### HasBinary

`func (o *ServerConfig) HasBinary() bool`

HasBinary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
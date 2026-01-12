# Server

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of the service | 
**AppLocationSettingId** | **int32** | The ID of the location setting | 
**Instance** | **int32** | The ID of the instance | 
**ServerConfigId** | **int32** | The ID of the server config | 
**ServerConfigName** | **string** | The name of the server config | 
**Name** | **string** | The name of the service | 
**Status** | **string** | The current status | 
**StatusMessage** | **NullableString** | An optional message | 
**IsBackupable** | **bool** | Indicates whether the service can be backed up | 
**IsRestorable** | **bool** | Indicates whether the service can be restored | 
**IsPending** | **bool** | Indicates whether the service is pending (not running) due to insufficient resources on the node. | 
**IsNotFound** | **bool** | Indicates whether the service is currently not found/missing in the cluster. | 
**IsHealthy** | **bool** | Indicates whether the service is currently in an overall healthy state. | 
**IsStopped** | **bool** | Indicates whether the service is currently stopped. | 
**Maintenance** | **bool** | Indicates if the service is under maintenance | 
**Ports** | [**[]Port**](Port.md) | The port definitions of the service | 
**Env** | [**map[string]EnvironmentVariable**](EnvironmentVariable.md) | The environment variables of the service | 
**RestartPolicy** | [**RestartPolicy**](RestartPolicy.md) | The restart policy of the service | 
**Node** | Pointer to [**Node**](Node.md) | The assigned node. If missing, it likely means the service is overloaded and has not yet been assigned to any node. | [optional] 
**Location** | [**Location**](Location.md) | The location of the node | 
**Resources** | [**ResourcePackage**](ResourcePackage.md) | The assigned resource package | 
**Backups** | [**[]Backup**](Backup.md) | The backups of the service | 
**Metadata** | **map[string]interface{}** | The metadata for the service | 
**CreatedAt** | **NullableTime** | When the service was created | 
**UpdatedAt** | **NullableTime** | When the service was last updated | 

## Methods

### NewServer

`func NewServer(id int32, appLocationSettingId int32, instance int32, serverConfigId int32, serverConfigName string, name string, status string, statusMessage NullableString, isBackupable bool, isRestorable bool, isPending bool, isNotFound bool, isHealthy bool, isStopped bool, maintenance bool, ports []Port, env map[string]EnvironmentVariable, restartPolicy RestartPolicy, location Location, resources ResourcePackage, backups []Backup, metadata map[string]interface{}, createdAt NullableTime, updatedAt NullableTime, ) *Server`

NewServer instantiates a new Server object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerWithDefaults

`func NewServerWithDefaults() *Server`

NewServerWithDefaults instantiates a new Server object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Server) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Server) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Server) SetId(v int32)`

SetId sets Id field to given value.


### GetAppLocationSettingId

`func (o *Server) GetAppLocationSettingId() int32`

GetAppLocationSettingId returns the AppLocationSettingId field if non-nil, zero value otherwise.

### GetAppLocationSettingIdOk

`func (o *Server) GetAppLocationSettingIdOk() (*int32, bool)`

GetAppLocationSettingIdOk returns a tuple with the AppLocationSettingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppLocationSettingId

`func (o *Server) SetAppLocationSettingId(v int32)`

SetAppLocationSettingId sets AppLocationSettingId field to given value.


### GetInstance

`func (o *Server) GetInstance() int32`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *Server) GetInstanceOk() (*int32, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *Server) SetInstance(v int32)`

SetInstance sets Instance field to given value.


### GetServerConfigId

`func (o *Server) GetServerConfigId() int32`

GetServerConfigId returns the ServerConfigId field if non-nil, zero value otherwise.

### GetServerConfigIdOk

`func (o *Server) GetServerConfigIdOk() (*int32, bool)`

GetServerConfigIdOk returns a tuple with the ServerConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerConfigId

`func (o *Server) SetServerConfigId(v int32)`

SetServerConfigId sets ServerConfigId field to given value.


### GetServerConfigName

`func (o *Server) GetServerConfigName() string`

GetServerConfigName returns the ServerConfigName field if non-nil, zero value otherwise.

### GetServerConfigNameOk

`func (o *Server) GetServerConfigNameOk() (*string, bool)`

GetServerConfigNameOk returns a tuple with the ServerConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerConfigName

`func (o *Server) SetServerConfigName(v string)`

SetServerConfigName sets ServerConfigName field to given value.


### GetName

`func (o *Server) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Server) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Server) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *Server) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Server) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Server) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *Server) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *Server) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *Server) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.


### SetStatusMessageNil

`func (o *Server) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *Server) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetIsBackupable

`func (o *Server) GetIsBackupable() bool`

GetIsBackupable returns the IsBackupable field if non-nil, zero value otherwise.

### GetIsBackupableOk

`func (o *Server) GetIsBackupableOk() (*bool, bool)`

GetIsBackupableOk returns a tuple with the IsBackupable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBackupable

`func (o *Server) SetIsBackupable(v bool)`

SetIsBackupable sets IsBackupable field to given value.


### GetIsRestorable

`func (o *Server) GetIsRestorable() bool`

GetIsRestorable returns the IsRestorable field if non-nil, zero value otherwise.

### GetIsRestorableOk

`func (o *Server) GetIsRestorableOk() (*bool, bool)`

GetIsRestorableOk returns a tuple with the IsRestorable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRestorable

`func (o *Server) SetIsRestorable(v bool)`

SetIsRestorable sets IsRestorable field to given value.


### GetIsPending

`func (o *Server) GetIsPending() bool`

GetIsPending returns the IsPending field if non-nil, zero value otherwise.

### GetIsPendingOk

`func (o *Server) GetIsPendingOk() (*bool, bool)`

GetIsPendingOk returns a tuple with the IsPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPending

`func (o *Server) SetIsPending(v bool)`

SetIsPending sets IsPending field to given value.


### GetIsNotFound

`func (o *Server) GetIsNotFound() bool`

GetIsNotFound returns the IsNotFound field if non-nil, zero value otherwise.

### GetIsNotFoundOk

`func (o *Server) GetIsNotFoundOk() (*bool, bool)`

GetIsNotFoundOk returns a tuple with the IsNotFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNotFound

`func (o *Server) SetIsNotFound(v bool)`

SetIsNotFound sets IsNotFound field to given value.


### GetIsHealthy

`func (o *Server) GetIsHealthy() bool`

GetIsHealthy returns the IsHealthy field if non-nil, zero value otherwise.

### GetIsHealthyOk

`func (o *Server) GetIsHealthyOk() (*bool, bool)`

GetIsHealthyOk returns a tuple with the IsHealthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHealthy

`func (o *Server) SetIsHealthy(v bool)`

SetIsHealthy sets IsHealthy field to given value.


### GetIsStopped

`func (o *Server) GetIsStopped() bool`

GetIsStopped returns the IsStopped field if non-nil, zero value otherwise.

### GetIsStoppedOk

`func (o *Server) GetIsStoppedOk() (*bool, bool)`

GetIsStoppedOk returns a tuple with the IsStopped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStopped

`func (o *Server) SetIsStopped(v bool)`

SetIsStopped sets IsStopped field to given value.


### GetMaintenance

`func (o *Server) GetMaintenance() bool`

GetMaintenance returns the Maintenance field if non-nil, zero value otherwise.

### GetMaintenanceOk

`func (o *Server) GetMaintenanceOk() (*bool, bool)`

GetMaintenanceOk returns a tuple with the Maintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenance

`func (o *Server) SetMaintenance(v bool)`

SetMaintenance sets Maintenance field to given value.


### GetPorts

`func (o *Server) GetPorts() []Port`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *Server) GetPortsOk() (*[]Port, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *Server) SetPorts(v []Port)`

SetPorts sets Ports field to given value.


### GetEnv

`func (o *Server) GetEnv() map[string]EnvironmentVariable`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *Server) GetEnvOk() (*map[string]EnvironmentVariable, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *Server) SetEnv(v map[string]EnvironmentVariable)`

SetEnv sets Env field to given value.


### GetRestartPolicy

`func (o *Server) GetRestartPolicy() RestartPolicy`

GetRestartPolicy returns the RestartPolicy field if non-nil, zero value otherwise.

### GetRestartPolicyOk

`func (o *Server) GetRestartPolicyOk() (*RestartPolicy, bool)`

GetRestartPolicyOk returns a tuple with the RestartPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartPolicy

`func (o *Server) SetRestartPolicy(v RestartPolicy)`

SetRestartPolicy sets RestartPolicy field to given value.


### GetNode

`func (o *Server) GetNode() Node`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *Server) GetNodeOk() (*Node, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *Server) SetNode(v Node)`

SetNode sets Node field to given value.

### HasNode

`func (o *Server) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetLocation

`func (o *Server) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Server) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Server) SetLocation(v Location)`

SetLocation sets Location field to given value.


### GetResources

`func (o *Server) GetResources() ResourcePackage`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *Server) GetResourcesOk() (*ResourcePackage, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *Server) SetResources(v ResourcePackage)`

SetResources sets Resources field to given value.


### GetBackups

`func (o *Server) GetBackups() []Backup`

GetBackups returns the Backups field if non-nil, zero value otherwise.

### GetBackupsOk

`func (o *Server) GetBackupsOk() (*[]Backup, bool)`

GetBackupsOk returns a tuple with the Backups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackups

`func (o *Server) SetBackups(v []Backup)`

SetBackups sets Backups field to given value.


### GetMetadata

`func (o *Server) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Server) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Server) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetCreatedAt

`func (o *Server) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Server) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Server) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *Server) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Server) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *Server) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Server) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Server) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *Server) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Server) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
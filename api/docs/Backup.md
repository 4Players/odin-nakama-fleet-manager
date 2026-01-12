# Backup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the backup | 
**ArchiveName** | **string** | The name of the stored archive | 
**CreatedAt** | **time.Time** | The timestamp of the backups creation | 
**RestoredAt** | **NullableTime** | The timestamp of when the backup was restored, null if not restored | 

## Methods

### NewBackup

`func NewBackup(name string, archiveName string, createdAt time.Time, restoredAt NullableTime, ) *Backup`

NewBackup instantiates a new Backup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupWithDefaults

`func NewBackupWithDefaults() *Backup`

NewBackupWithDefaults instantiates a new Backup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Backup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Backup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Backup) SetName(v string)`

SetName sets Name field to given value.


### GetArchiveName

`func (o *Backup) GetArchiveName() string`

GetArchiveName returns the ArchiveName field if non-nil, zero value otherwise.

### GetArchiveNameOk

`func (o *Backup) GetArchiveNameOk() (*string, bool)`

GetArchiveNameOk returns a tuple with the ArchiveName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveName

`func (o *Backup) SetArchiveName(v string)`

SetArchiveName sets ArchiveName field to given value.


### GetCreatedAt

`func (o *Backup) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Backup) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Backup) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetRestoredAt

`func (o *Backup) GetRestoredAt() time.Time`

GetRestoredAt returns the RestoredAt field if non-nil, zero value otherwise.

### GetRestoredAtOk

`func (o *Backup) GetRestoredAtOk() (*time.Time, bool)`

GetRestoredAtOk returns a tuple with the RestoredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoredAt

`func (o *Backup) SetRestoredAt(v time.Time)`

SetRestoredAt sets RestoredAt field to given value.


### SetRestoredAtNil

`func (o *Backup) SetRestoredAtNil(b bool)`

 SetRestoredAtNil sets the value for RestoredAt to be an explicit nil

### UnsetRestoredAt
`func (o *Backup) UnsetRestoredAt()`

UnsetRestoredAt ensures that no value is present for RestoredAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
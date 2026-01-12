# ResourcePackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The id of the resource package | 
**Name** | **string** | The name of the resource package | 
**Slug** | **string** | The slug of the resource package | 
**Type** | [**ResourcePackageType**](ResourcePackageType.md) | The type of the resource package | 
**CpuLimit** | **float32** | The maximum CPU of the resource package | 
**MemoryLimitBytes** | **int64** | The maximum memory in bytes of the resource package | 
**MemoryLimitMiB** | **float32** | The maximum memory in mebibytes (MiB) of the resource package | 
**MemoryLimitFileSizeString** | **string** | The maximum memory in a human-readable format | 

## Methods

### NewResourcePackage

`func NewResourcePackage(id int32, name string, slug string, type_ ResourcePackageType, cpuLimit float32, memoryLimitBytes int64, memoryLimitMiB float32, memoryLimitFileSizeString string, ) *ResourcePackage`

NewResourcePackage instantiates a new ResourcePackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcePackageWithDefaults

`func NewResourcePackageWithDefaults() *ResourcePackage`

NewResourcePackageWithDefaults instantiates a new ResourcePackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourcePackage) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourcePackage) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourcePackage) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ResourcePackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourcePackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourcePackage) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *ResourcePackage) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ResourcePackage) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ResourcePackage) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetType

`func (o *ResourcePackage) GetType() ResourcePackageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourcePackage) GetTypeOk() (*ResourcePackageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourcePackage) SetType(v ResourcePackageType)`

SetType sets Type field to given value.


### GetCpuLimit

`func (o *ResourcePackage) GetCpuLimit() float32`

GetCpuLimit returns the CpuLimit field if non-nil, zero value otherwise.

### GetCpuLimitOk

`func (o *ResourcePackage) GetCpuLimitOk() (*float32, bool)`

GetCpuLimitOk returns a tuple with the CpuLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLimit

`func (o *ResourcePackage) SetCpuLimit(v float32)`

SetCpuLimit sets CpuLimit field to given value.


### GetMemoryLimitBytes

`func (o *ResourcePackage) GetMemoryLimitBytes() int64`

GetMemoryLimitBytes returns the MemoryLimitBytes field if non-nil, zero value otherwise.

### GetMemoryLimitBytesOk

`func (o *ResourcePackage) GetMemoryLimitBytesOk() (*int64, bool)`

GetMemoryLimitBytesOk returns a tuple with the MemoryLimitBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimitBytes

`func (o *ResourcePackage) SetMemoryLimitBytes(v int64)`

SetMemoryLimitBytes sets MemoryLimitBytes field to given value.


### GetMemoryLimitMiB

`func (o *ResourcePackage) GetMemoryLimitMiB() float32`

GetMemoryLimitMiB returns the MemoryLimitMiB field if non-nil, zero value otherwise.

### GetMemoryLimitMiBOk

`func (o *ResourcePackage) GetMemoryLimitMiBOk() (*float32, bool)`

GetMemoryLimitMiBOk returns a tuple with the MemoryLimitMiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimitMiB

`func (o *ResourcePackage) SetMemoryLimitMiB(v float32)`

SetMemoryLimitMiB sets MemoryLimitMiB field to given value.


### GetMemoryLimitFileSizeString

`func (o *ResourcePackage) GetMemoryLimitFileSizeString() string`

GetMemoryLimitFileSizeString returns the MemoryLimitFileSizeString field if non-nil, zero value otherwise.

### GetMemoryLimitFileSizeStringOk

`func (o *ResourcePackage) GetMemoryLimitFileSizeStringOk() (*string, bool)`

GetMemoryLimitFileSizeStringOk returns a tuple with the MemoryLimitFileSizeString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimitFileSizeString

`func (o *ResourcePackage) SetMemoryLimitFileSizeString(v string)`

SetMemoryLimitFileSizeString sets MemoryLimitFileSizeString field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
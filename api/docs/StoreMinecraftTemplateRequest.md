# StoreMinecraftTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | Pointer to **NullableString** | The name of the template app. If omitted, a default name will be used. | [optional] 
**ResourcePackageSlug** | Pointer to **NullableString** | The slug of the resource package. If omitted, a default resource package will be used. | [optional] 
**Placement** | Pointer to [**NullableCreateUpdatePlacement**](CreateUpdatePlacement.md) | The placement settings that define the location and other constraints If omitted, no deployment will be created. | [optional] 
**Password** | Pointer to **NullableString** | Password required to deploy services to a protected node location. Required when the selected location is password-protected. | [optional] 

## Methods

### NewStoreMinecraftTemplateRequest

`func NewStoreMinecraftTemplateRequest() *StoreMinecraftTemplateRequest`

NewStoreMinecraftTemplateRequest instantiates a new StoreMinecraftTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreMinecraftTemplateRequestWithDefaults

`func NewStoreMinecraftTemplateRequestWithDefaults() *StoreMinecraftTemplateRequest`

NewStoreMinecraftTemplateRequestWithDefaults instantiates a new StoreMinecraftTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *StoreMinecraftTemplateRequest) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *StoreMinecraftTemplateRequest) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *StoreMinecraftTemplateRequest) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *StoreMinecraftTemplateRequest) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### SetAppNameNil

`func (o *StoreMinecraftTemplateRequest) SetAppNameNil(b bool)`

 SetAppNameNil sets the value for AppName to be an explicit nil

### UnsetAppName
`func (o *StoreMinecraftTemplateRequest) UnsetAppName()`

UnsetAppName ensures that no value is present for AppName, not even an explicit nil
### GetResourcePackageSlug

`func (o *StoreMinecraftTemplateRequest) GetResourcePackageSlug() string`

GetResourcePackageSlug returns the ResourcePackageSlug field if non-nil, zero value otherwise.

### GetResourcePackageSlugOk

`func (o *StoreMinecraftTemplateRequest) GetResourcePackageSlugOk() (*string, bool)`

GetResourcePackageSlugOk returns a tuple with the ResourcePackageSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePackageSlug

`func (o *StoreMinecraftTemplateRequest) SetResourcePackageSlug(v string)`

SetResourcePackageSlug sets ResourcePackageSlug field to given value.

### HasResourcePackageSlug

`func (o *StoreMinecraftTemplateRequest) HasResourcePackageSlug() bool`

HasResourcePackageSlug returns a boolean if a field has been set.

### SetResourcePackageSlugNil

`func (o *StoreMinecraftTemplateRequest) SetResourcePackageSlugNil(b bool)`

 SetResourcePackageSlugNil sets the value for ResourcePackageSlug to be an explicit nil

### UnsetResourcePackageSlug
`func (o *StoreMinecraftTemplateRequest) UnsetResourcePackageSlug()`

UnsetResourcePackageSlug ensures that no value is present for ResourcePackageSlug, not even an explicit nil
### GetPlacement

`func (o *StoreMinecraftTemplateRequest) GetPlacement() CreateUpdatePlacement`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *StoreMinecraftTemplateRequest) GetPlacementOk() (*CreateUpdatePlacement, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *StoreMinecraftTemplateRequest) SetPlacement(v CreateUpdatePlacement)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *StoreMinecraftTemplateRequest) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### SetPlacementNil

`func (o *StoreMinecraftTemplateRequest) SetPlacementNil(b bool)`

 SetPlacementNil sets the value for Placement to be an explicit nil

### UnsetPlacement
`func (o *StoreMinecraftTemplateRequest) UnsetPlacement()`

UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
### GetPassword

`func (o *StoreMinecraftTemplateRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *StoreMinecraftTemplateRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *StoreMinecraftTemplateRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *StoreMinecraftTemplateRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *StoreMinecraftTemplateRequest) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *StoreMinecraftTemplateRequest) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
# StorePalworldTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | Pointer to **NullableString** | The name of the template app. If omitted, a default name will be used. | [optional] 
**ResourcePackageSlug** | Pointer to **NullableString** | The slug of the resource package. If omitted, a default resource package will be used. | [optional] 
**Placement** | Pointer to [**NullableCreateUpdatePlacement**](CreateUpdatePlacement.md) | The placement constraints that define the location. If omitted, no deployment will be created. | [optional] 
**Password** | Pointer to **NullableString** | Password required to deploy services to a protected location. Only required when the selected location is password-protected. | [optional] 

## Methods

### NewStorePalworldTemplateRequest

`func NewStorePalworldTemplateRequest() *StorePalworldTemplateRequest`

NewStorePalworldTemplateRequest instantiates a new StorePalworldTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorePalworldTemplateRequestWithDefaults

`func NewStorePalworldTemplateRequestWithDefaults() *StorePalworldTemplateRequest`

NewStorePalworldTemplateRequestWithDefaults instantiates a new StorePalworldTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *StorePalworldTemplateRequest) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *StorePalworldTemplateRequest) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *StorePalworldTemplateRequest) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *StorePalworldTemplateRequest) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### SetAppNameNil

`func (o *StorePalworldTemplateRequest) SetAppNameNil(b bool)`

 SetAppNameNil sets the value for AppName to be an explicit nil

### UnsetAppName
`func (o *StorePalworldTemplateRequest) UnsetAppName()`

UnsetAppName ensures that no value is present for AppName, not even an explicit nil
### GetResourcePackageSlug

`func (o *StorePalworldTemplateRequest) GetResourcePackageSlug() string`

GetResourcePackageSlug returns the ResourcePackageSlug field if non-nil, zero value otherwise.

### GetResourcePackageSlugOk

`func (o *StorePalworldTemplateRequest) GetResourcePackageSlugOk() (*string, bool)`

GetResourcePackageSlugOk returns a tuple with the ResourcePackageSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePackageSlug

`func (o *StorePalworldTemplateRequest) SetResourcePackageSlug(v string)`

SetResourcePackageSlug sets ResourcePackageSlug field to given value.

### HasResourcePackageSlug

`func (o *StorePalworldTemplateRequest) HasResourcePackageSlug() bool`

HasResourcePackageSlug returns a boolean if a field has been set.

### SetResourcePackageSlugNil

`func (o *StorePalworldTemplateRequest) SetResourcePackageSlugNil(b bool)`

 SetResourcePackageSlugNil sets the value for ResourcePackageSlug to be an explicit nil

### UnsetResourcePackageSlug
`func (o *StorePalworldTemplateRequest) UnsetResourcePackageSlug()`

UnsetResourcePackageSlug ensures that no value is present for ResourcePackageSlug, not even an explicit nil
### GetPlacement

`func (o *StorePalworldTemplateRequest) GetPlacement() CreateUpdatePlacement`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *StorePalworldTemplateRequest) GetPlacementOk() (*CreateUpdatePlacement, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *StorePalworldTemplateRequest) SetPlacement(v CreateUpdatePlacement)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *StorePalworldTemplateRequest) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### SetPlacementNil

`func (o *StorePalworldTemplateRequest) SetPlacementNil(b bool)`

 SetPlacementNil sets the value for Placement to be an explicit nil

### UnsetPlacement
`func (o *StorePalworldTemplateRequest) UnsetPlacement()`

UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
### GetPassword

`func (o *StorePalworldTemplateRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *StorePalworldTemplateRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *StorePalworldTemplateRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *StorePalworldTemplateRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *StorePalworldTemplateRequest) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *StorePalworldTemplateRequest) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
# BackupDownload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The url where the file is available | 
**ValidUntil** | **time.Time** | After this date the link is no longer valid | 

## Methods

### NewBackupDownload

`func NewBackupDownload(url string, validUntil time.Time, ) *BackupDownload`

NewBackupDownload instantiates a new BackupDownload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupDownloadWithDefaults

`func NewBackupDownloadWithDefaults() *BackupDownload`

NewBackupDownloadWithDefaults instantiates a new BackupDownload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *BackupDownload) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BackupDownload) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BackupDownload) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetValidUntil

`func (o *BackupDownload) GetValidUntil() time.Time`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *BackupDownload) GetValidUntilOk() (*time.Time, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *BackupDownload) SetValidUntil(v time.Time)`

SetValidUntil sets ValidUntil field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
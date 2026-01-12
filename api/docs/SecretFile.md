# SecretFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Target** | **string** | The path in the container | 
**Content** | **string** | The content of the secret file | 

## Methods

### NewSecretFile

`func NewSecretFile(target string, content string, ) *SecretFile`

NewSecretFile instantiates a new SecretFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretFileWithDefaults

`func NewSecretFileWithDefaults() *SecretFile`

NewSecretFileWithDefaults instantiates a new SecretFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTarget

`func (o *SecretFile) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SecretFile) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SecretFile) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetContent

`func (o *SecretFile) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SecretFile) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SecretFile) SetContent(v string)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
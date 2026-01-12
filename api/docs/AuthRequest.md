# AuthRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email of the user | 
**Password** | **string** | The password of the user | 
**Sid** | **string** | The fusion session id of the user | 

## Methods

### NewAuthRequest

`func NewAuthRequest(email string, password string, sid string, ) *AuthRequest`

NewAuthRequest instantiates a new AuthRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthRequestWithDefaults

`func NewAuthRequestWithDefaults() *AuthRequest`

NewAuthRequestWithDefaults instantiates a new AuthRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *AuthRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AuthRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AuthRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *AuthRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AuthRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AuthRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetSid

`func (o *AuthRequest) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *AuthRequest) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *AuthRequest) SetSid(v string)`

SetSid sets Sid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
# WalletCurrency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Unique code, e.g. &#39;CPUH&#39; | 
**Name** | **string** | Human-readable currency name, e.g. &#39;CPU-hours&#39; | 

## Methods

### NewWalletCurrency

`func NewWalletCurrency(code string, name string, ) *WalletCurrency`

NewWalletCurrency instantiates a new WalletCurrency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletCurrencyWithDefaults

`func NewWalletCurrencyWithDefaults() *WalletCurrency`

NewWalletCurrencyWithDefaults instantiates a new WalletCurrency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *WalletCurrency) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *WalletCurrency) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *WalletCurrency) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *WalletCurrency) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WalletCurrency) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WalletCurrency) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
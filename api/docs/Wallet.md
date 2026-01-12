# Wallet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The id of the wallet. | 
**Balance** | **float32** | The balance of the wallet. | 
**Scope** | [**WalletScope**](WalletScope.md) | The scope of the wallet. | 
**InitialBalance** | **float32** | The initial balance of the wallet. | 
**RechargeAmount** | **NullableFloat32** | The recharge amount or null if none. | 
**RechargeAt** | **NullableTime** | Next recharge date or null if none. | 
**UpdatedAt** | **time.Time** | When the wallet was last updated. | 
**Currency** | [**WalletCurrency**](WalletCurrency.md) | The currency of the wallet | 

## Methods

### NewWallet

`func NewWallet(id int32, balance float32, scope WalletScope, initialBalance float32, rechargeAmount NullableFloat32, rechargeAt NullableTime, updatedAt time.Time, currency WalletCurrency, ) *Wallet`

NewWallet instantiates a new Wallet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletWithDefaults

`func NewWalletWithDefaults() *Wallet`

NewWalletWithDefaults instantiates a new Wallet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Wallet) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Wallet) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Wallet) SetId(v int32)`

SetId sets Id field to given value.


### GetBalance

`func (o *Wallet) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *Wallet) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *Wallet) SetBalance(v float32)`

SetBalance sets Balance field to given value.


### GetScope

`func (o *Wallet) GetScope() WalletScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Wallet) GetScopeOk() (*WalletScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Wallet) SetScope(v WalletScope)`

SetScope sets Scope field to given value.


### GetInitialBalance

`func (o *Wallet) GetInitialBalance() float32`

GetInitialBalance returns the InitialBalance field if non-nil, zero value otherwise.

### GetInitialBalanceOk

`func (o *Wallet) GetInitialBalanceOk() (*float32, bool)`

GetInitialBalanceOk returns a tuple with the InitialBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialBalance

`func (o *Wallet) SetInitialBalance(v float32)`

SetInitialBalance sets InitialBalance field to given value.


### GetRechargeAmount

`func (o *Wallet) GetRechargeAmount() float32`

GetRechargeAmount returns the RechargeAmount field if non-nil, zero value otherwise.

### GetRechargeAmountOk

`func (o *Wallet) GetRechargeAmountOk() (*float32, bool)`

GetRechargeAmountOk returns a tuple with the RechargeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRechargeAmount

`func (o *Wallet) SetRechargeAmount(v float32)`

SetRechargeAmount sets RechargeAmount field to given value.


### SetRechargeAmountNil

`func (o *Wallet) SetRechargeAmountNil(b bool)`

 SetRechargeAmountNil sets the value for RechargeAmount to be an explicit nil

### UnsetRechargeAmount
`func (o *Wallet) UnsetRechargeAmount()`

UnsetRechargeAmount ensures that no value is present for RechargeAmount, not even an explicit nil
### GetRechargeAt

`func (o *Wallet) GetRechargeAt() time.Time`

GetRechargeAt returns the RechargeAt field if non-nil, zero value otherwise.

### GetRechargeAtOk

`func (o *Wallet) GetRechargeAtOk() (*time.Time, bool)`

GetRechargeAtOk returns a tuple with the RechargeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRechargeAt

`func (o *Wallet) SetRechargeAt(v time.Time)`

SetRechargeAt sets RechargeAt field to given value.


### SetRechargeAtNil

`func (o *Wallet) SetRechargeAtNil(b bool)`

 SetRechargeAtNil sets the value for RechargeAt to be an explicit nil

### UnsetRechargeAt
`func (o *Wallet) UnsetRechargeAt()`

UnsetRechargeAt ensures that no value is present for RechargeAt, not even an explicit nil
### GetUpdatedAt

`func (o *Wallet) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Wallet) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Wallet) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCurrency

`func (o *Wallet) GetCurrency() WalletCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Wallet) GetCurrencyOk() (*WalletCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Wallet) SetCurrency(v WalletCurrency)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
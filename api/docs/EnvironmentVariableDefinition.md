# EnvironmentVariableDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnvironmentVariableType**](EnvironmentVariableType.md) | The type of the environment variable | 
**Key** | **string** | The key of the variable | 
**Value** | Pointer to **string** | The value of the variable | [optional] 
**Variable** | Pointer to **string** | The variable definition of the environment variable | [optional] 

## Methods

### NewEnvironmentVariableDefinition

`func NewEnvironmentVariableDefinition(type_ EnvironmentVariableType, key string, ) *EnvironmentVariableDefinition`

NewEnvironmentVariableDefinition instantiates a new EnvironmentVariableDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentVariableDefinitionWithDefaults

`func NewEnvironmentVariableDefinitionWithDefaults() *EnvironmentVariableDefinition`

NewEnvironmentVariableDefinitionWithDefaults instantiates a new EnvironmentVariableDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EnvironmentVariableDefinition) GetType() EnvironmentVariableType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnvironmentVariableDefinition) GetTypeOk() (*EnvironmentVariableType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnvironmentVariableDefinition) SetType(v EnvironmentVariableType)`

SetType sets Type field to given value.


### GetKey

`func (o *EnvironmentVariableDefinition) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EnvironmentVariableDefinition) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EnvironmentVariableDefinition) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *EnvironmentVariableDefinition) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EnvironmentVariableDefinition) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EnvironmentVariableDefinition) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *EnvironmentVariableDefinition) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVariable

`func (o *EnvironmentVariableDefinition) GetVariable() string`

GetVariable returns the Variable field if non-nil, zero value otherwise.

### GetVariableOk

`func (o *EnvironmentVariableDefinition) GetVariableOk() (*string, bool)`

GetVariableOk returns a tuple with the Variable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariable

`func (o *EnvironmentVariableDefinition) SetVariable(v string)`

SetVariable sets Variable field to given value.

### HasVariable

`func (o *EnvironmentVariableDefinition) HasVariable() bool`

HasVariable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
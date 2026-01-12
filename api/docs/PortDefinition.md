# PortDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the server config | 
**Protocols** | [**[]Protocol**](Protocol.md) | The protocols to expose | 
**TargetPort** | **int32** | The port to expose | 

## Methods

### NewPortDefinition

`func NewPortDefinition(name string, protocols []Protocol, targetPort int32, ) *PortDefinition`

NewPortDefinition instantiates a new PortDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortDefinitionWithDefaults

`func NewPortDefinitionWithDefaults() *PortDefinition`

NewPortDefinitionWithDefaults instantiates a new PortDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PortDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PortDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PortDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetProtocols

`func (o *PortDefinition) GetProtocols() []Protocol`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *PortDefinition) GetProtocolsOk() (*[]Protocol, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *PortDefinition) SetProtocols(v []Protocol)`

SetProtocols sets Protocols field to given value.


### GetTargetPort

`func (o *PortDefinition) GetTargetPort() int32`

GetTargetPort returns the TargetPort field if non-nil, zero value otherwise.

### GetTargetPortOk

`func (o *PortDefinition) GetTargetPortOk() (*int32, bool)`

GetTargetPortOk returns a tuple with the TargetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPort

`func (o *PortDefinition) SetTargetPort(v int32)`

SetTargetPort sets TargetPort field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
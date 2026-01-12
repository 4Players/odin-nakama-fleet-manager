# Port

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name for the port. | 
**Protocol** | [**Protocol**](Protocol.md) | The protocol used by the port. | 
**TargetPort** | **int32** | The port inside the container. Incoming traffic is routed to this port within the container. | 
**PublishedPort** | **NullableInt32** | The externally published port, which allows external access. If &#x60;null&#x60;, the port has not yet been assigned. | 

## Methods

### NewPort

`func NewPort(name string, protocol Protocol, targetPort int32, publishedPort NullableInt32, ) *Port`

NewPort instantiates a new Port object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortWithDefaults

`func NewPortWithDefaults() *Port`

NewPortWithDefaults instantiates a new Port object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Port) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Port) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Port) SetName(v string)`

SetName sets Name field to given value.


### GetProtocol

`func (o *Port) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *Port) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *Port) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.


### GetTargetPort

`func (o *Port) GetTargetPort() int32`

GetTargetPort returns the TargetPort field if non-nil, zero value otherwise.

### GetTargetPortOk

`func (o *Port) GetTargetPortOk() (*int32, bool)`

GetTargetPortOk returns a tuple with the TargetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPort

`func (o *Port) SetTargetPort(v int32)`

SetTargetPort sets TargetPort field to given value.


### GetPublishedPort

`func (o *Port) GetPublishedPort() int32`

GetPublishedPort returns the PublishedPort field if non-nil, zero value otherwise.

### GetPublishedPortOk

`func (o *Port) GetPublishedPortOk() (*int32, bool)`

GetPublishedPortOk returns a tuple with the PublishedPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedPort

`func (o *Port) SetPublishedPort(v int32)`

SetPublishedPort sets PublishedPort field to given value.


### SetPublishedPortNil

`func (o *Port) SetPublishedPortNil(b bool)`

 SetPublishedPortNil sets the value for PublishedPort to be an explicit nil

### UnsetPublishedPort
`func (o *Port) UnsetPublishedPort()`

UnsetPublishedPort ensures that no value is present for PublishedPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
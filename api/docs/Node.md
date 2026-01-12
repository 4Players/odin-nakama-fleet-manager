# Node

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | [**Architecture**](Architecture.md) | The architecture of the node (e.g., &#39;x86_64&#39;) | 
**Os** | [**OperatingSystem**](OperatingSystem.md) | The operating system of the node | 
**Address** | **string** | The IP address of the node | 

## Methods

### NewNode

`func NewNode(architecture Architecture, os OperatingSystem, address string, ) *Node`

NewNode instantiates a new Node object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeWithDefaults

`func NewNodeWithDefaults() *Node`

NewNodeWithDefaults instantiates a new Node object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitecture

`func (o *Node) GetArchitecture() Architecture`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *Node) GetArchitectureOk() (*Architecture, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *Node) SetArchitecture(v Architecture)`

SetArchitecture sets Architecture field to given value.


### GetOs

`func (o *Node) GetOs() OperatingSystem`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *Node) GetOsOk() (*OperatingSystem, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *Node) SetOs(v OperatingSystem)`

SetOs sets Os field to given value.


### GetAddress

`func (o *Node) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Node) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Node) SetAddress(v string)`

SetAddress sets Address field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
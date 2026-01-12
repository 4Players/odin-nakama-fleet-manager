# Placement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Constraints** | [**Location**](Location.md) | The constraints to use | 

## Methods

### NewPlacement

`func NewPlacement(constraints Location, ) *Placement`

NewPlacement instantiates a new Placement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlacementWithDefaults

`func NewPlacementWithDefaults() *Placement`

NewPlacementWithDefaults instantiates a new Placement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstraints

`func (o *Placement) GetConstraints() Location`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *Placement) GetConstraintsOk() (*Location, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *Placement) SetConstraints(v Location)`

SetConstraints sets Constraints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
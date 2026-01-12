# RestartPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | [**RestartPolicyCondition**](RestartPolicyCondition.md) | The condition for a restart | 

## Methods

### NewRestartPolicy

`func NewRestartPolicy(condition RestartPolicyCondition, ) *RestartPolicy`

NewRestartPolicy instantiates a new RestartPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestartPolicyWithDefaults

`func NewRestartPolicyWithDefaults() *RestartPolicy`

NewRestartPolicyWithDefaults instantiates a new RestartPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *RestartPolicy) GetCondition() RestartPolicyCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *RestartPolicy) GetConditionOk() (*RestartPolicyCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *RestartPolicy) SetCondition(v RestartPolicyCondition)`

SetCondition sets Condition field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
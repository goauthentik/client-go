# DeviceUserBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Policy** | Pointer to **NullableString** |  | [optional] 
**Group** | Pointer to **NullableString** |  | [optional] 
**User** | Pointer to **NullableInt32** |  | [optional] 
**PolicyObj** | [**Policy**](Policy.md) |  | [readonly] 
**GroupObj** | [**PartialGroup**](PartialGroup.md) |  | [readonly] 
**UserObj** | [**PartialUser**](PartialUser.md) |  | [readonly] 
**Target** | **string** |  | 
**Negate** | Pointer to **bool** | Negates the outcome of the policy. Messages are unaffected. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Order** | **int32** |  | 
**Timeout** | Pointer to **int32** | Timeout after which Policy execution is terminated. | [optional] 
**FailureResult** | Pointer to **bool** | Result if the Policy execution fails. | [optional] 
**IsPrimary** | Pointer to **bool** |  | [optional] 
**Connector** | **NullableString** |  | [readonly] 
**ConnectorObj** | [**Connector**](Connector.md) |  | [readonly] 

## Methods

### NewDeviceUserBinding

`func NewDeviceUserBinding(pk string, policyObj Policy, groupObj PartialGroup, userObj PartialUser, target string, order int32, connector NullableString, connectorObj Connector, ) *DeviceUserBinding`

NewDeviceUserBinding instantiates a new DeviceUserBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceUserBindingWithDefaults

`func NewDeviceUserBindingWithDefaults() *DeviceUserBinding`

NewDeviceUserBindingWithDefaults instantiates a new DeviceUserBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *DeviceUserBinding) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *DeviceUserBinding) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *DeviceUserBinding) SetPk(v string)`

SetPk sets Pk field to given value.


### GetPolicy

`func (o *DeviceUserBinding) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *DeviceUserBinding) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *DeviceUserBinding) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *DeviceUserBinding) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicyNil

`func (o *DeviceUserBinding) SetPolicyNil(b bool)`

 SetPolicyNil sets the value for Policy to be an explicit nil

### UnsetPolicy
`func (o *DeviceUserBinding) UnsetPolicy()`

UnsetPolicy ensures that no value is present for Policy, not even an explicit nil
### GetGroup

`func (o *DeviceUserBinding) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *DeviceUserBinding) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *DeviceUserBinding) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *DeviceUserBinding) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *DeviceUserBinding) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *DeviceUserBinding) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetUser

`func (o *DeviceUserBinding) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DeviceUserBinding) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DeviceUserBinding) SetUser(v int32)`

SetUser sets User field to given value.

### HasUser

`func (o *DeviceUserBinding) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *DeviceUserBinding) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *DeviceUserBinding) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetPolicyObj

`func (o *DeviceUserBinding) GetPolicyObj() Policy`

GetPolicyObj returns the PolicyObj field if non-nil, zero value otherwise.

### GetPolicyObjOk

`func (o *DeviceUserBinding) GetPolicyObjOk() (*Policy, bool)`

GetPolicyObjOk returns a tuple with the PolicyObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObj

`func (o *DeviceUserBinding) SetPolicyObj(v Policy)`

SetPolicyObj sets PolicyObj field to given value.


### GetGroupObj

`func (o *DeviceUserBinding) GetGroupObj() PartialGroup`

GetGroupObj returns the GroupObj field if non-nil, zero value otherwise.

### GetGroupObjOk

`func (o *DeviceUserBinding) GetGroupObjOk() (*PartialGroup, bool)`

GetGroupObjOk returns a tuple with the GroupObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupObj

`func (o *DeviceUserBinding) SetGroupObj(v PartialGroup)`

SetGroupObj sets GroupObj field to given value.


### GetUserObj

`func (o *DeviceUserBinding) GetUserObj() PartialUser`

GetUserObj returns the UserObj field if non-nil, zero value otherwise.

### GetUserObjOk

`func (o *DeviceUserBinding) GetUserObjOk() (*PartialUser, bool)`

GetUserObjOk returns a tuple with the UserObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObj

`func (o *DeviceUserBinding) SetUserObj(v PartialUser)`

SetUserObj sets UserObj field to given value.


### GetTarget

`func (o *DeviceUserBinding) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *DeviceUserBinding) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *DeviceUserBinding) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetNegate

`func (o *DeviceUserBinding) GetNegate() bool`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *DeviceUserBinding) GetNegateOk() (*bool, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *DeviceUserBinding) SetNegate(v bool)`

SetNegate sets Negate field to given value.

### HasNegate

`func (o *DeviceUserBinding) HasNegate() bool`

HasNegate returns a boolean if a field has been set.

### GetEnabled

`func (o *DeviceUserBinding) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeviceUserBinding) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeviceUserBinding) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DeviceUserBinding) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOrder

`func (o *DeviceUserBinding) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *DeviceUserBinding) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *DeviceUserBinding) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetTimeout

`func (o *DeviceUserBinding) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *DeviceUserBinding) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *DeviceUserBinding) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *DeviceUserBinding) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetFailureResult

`func (o *DeviceUserBinding) GetFailureResult() bool`

GetFailureResult returns the FailureResult field if non-nil, zero value otherwise.

### GetFailureResultOk

`func (o *DeviceUserBinding) GetFailureResultOk() (*bool, bool)`

GetFailureResultOk returns a tuple with the FailureResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureResult

`func (o *DeviceUserBinding) SetFailureResult(v bool)`

SetFailureResult sets FailureResult field to given value.

### HasFailureResult

`func (o *DeviceUserBinding) HasFailureResult() bool`

HasFailureResult returns a boolean if a field has been set.

### GetIsPrimary

`func (o *DeviceUserBinding) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *DeviceUserBinding) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *DeviceUserBinding) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *DeviceUserBinding) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetConnector

`func (o *DeviceUserBinding) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *DeviceUserBinding) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *DeviceUserBinding) SetConnector(v string)`

SetConnector sets Connector field to given value.


### SetConnectorNil

`func (o *DeviceUserBinding) SetConnectorNil(b bool)`

 SetConnectorNil sets the value for Connector to be an explicit nil

### UnsetConnector
`func (o *DeviceUserBinding) UnsetConnector()`

UnsetConnector ensures that no value is present for Connector, not even an explicit nil
### GetConnectorObj

`func (o *DeviceUserBinding) GetConnectorObj() Connector`

GetConnectorObj returns the ConnectorObj field if non-nil, zero value otherwise.

### GetConnectorObjOk

`func (o *DeviceUserBinding) GetConnectorObjOk() (*Connector, bool)`

GetConnectorObjOk returns a tuple with the ConnectorObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorObj

`func (o *DeviceUserBinding) SetConnectorObj(v Connector)`

SetConnectorObj sets ConnectorObj field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



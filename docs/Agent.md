# Agent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**Username** | **string** | Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. | 
**Name** | **string** | User&#39;s display name. | 
**IsActive** | Pointer to **bool** | Designates whether this user should be treated as active. Unselect this instead of deleting accounts. | [optional] 
**LastLogin** | Pointer to **NullableTime** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**Uid** | **string** |  | [readonly] 
**Uuid** | **string** |  | [readonly] 
**Expiring** | Pointer to **bool** |  | [optional] 
**Expires** | Pointer to **NullableTime** |  | [optional] 
**Parent** | [**PartialUser**](PartialUser.md) |  | [readonly] 
**PolicyBehavior** | [**PolicyBehaviorEnum**](PolicyBehaviorEnum.md) |  | [readonly] 
**TokenIdentifier** | **NullableString** | Identifier of the agent&#39;s API token, so its key can be retrieved/copied later. | [readonly] 

## Methods

### NewAgent

`func NewAgent(pk int32, username string, name string, uid string, uuid string, parent PartialUser, policyBehavior PolicyBehaviorEnum, tokenIdentifier NullableString, ) *Agent`

NewAgent instantiates a new Agent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentWithDefaults

`func NewAgentWithDefaults() *Agent`

NewAgentWithDefaults instantiates a new Agent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *Agent) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *Agent) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *Agent) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetUsername

`func (o *Agent) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Agent) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Agent) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetName

`func (o *Agent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Agent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Agent) SetName(v string)`

SetName sets Name field to given value.


### GetIsActive

`func (o *Agent) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Agent) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Agent) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *Agent) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetLastLogin

`func (o *Agent) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *Agent) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *Agent) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *Agent) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### SetLastLoginNil

`func (o *Agent) SetLastLoginNil(b bool)`

 SetLastLoginNil sets the value for LastLogin to be an explicit nil

### UnsetLastLogin
`func (o *Agent) UnsetLastLogin()`

UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
### GetEmail

`func (o *Agent) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Agent) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Agent) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Agent) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAttributes

`func (o *Agent) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Agent) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Agent) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Agent) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetUid

`func (o *Agent) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Agent) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Agent) SetUid(v string)`

SetUid sets Uid field to given value.


### GetUuid

`func (o *Agent) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Agent) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Agent) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetExpiring

`func (o *Agent) GetExpiring() bool`

GetExpiring returns the Expiring field if non-nil, zero value otherwise.

### GetExpiringOk

`func (o *Agent) GetExpiringOk() (*bool, bool)`

GetExpiringOk returns a tuple with the Expiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiring

`func (o *Agent) SetExpiring(v bool)`

SetExpiring sets Expiring field to given value.

### HasExpiring

`func (o *Agent) HasExpiring() bool`

HasExpiring returns a boolean if a field has been set.

### GetExpires

`func (o *Agent) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *Agent) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *Agent) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *Agent) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *Agent) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *Agent) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetParent

`func (o *Agent) GetParent() PartialUser`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *Agent) GetParentOk() (*PartialUser, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *Agent) SetParent(v PartialUser)`

SetParent sets Parent field to given value.


### GetPolicyBehavior

`func (o *Agent) GetPolicyBehavior() PolicyBehaviorEnum`

GetPolicyBehavior returns the PolicyBehavior field if non-nil, zero value otherwise.

### GetPolicyBehaviorOk

`func (o *Agent) GetPolicyBehaviorOk() (*PolicyBehaviorEnum, bool)`

GetPolicyBehaviorOk returns a tuple with the PolicyBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyBehavior

`func (o *Agent) SetPolicyBehavior(v PolicyBehaviorEnum)`

SetPolicyBehavior sets PolicyBehavior field to given value.


### GetTokenIdentifier

`func (o *Agent) GetTokenIdentifier() string`

GetTokenIdentifier returns the TokenIdentifier field if non-nil, zero value otherwise.

### GetTokenIdentifierOk

`func (o *Agent) GetTokenIdentifierOk() (*string, bool)`

GetTokenIdentifierOk returns a tuple with the TokenIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIdentifier

`func (o *Agent) SetTokenIdentifier(v string)`

SetTokenIdentifier sets TokenIdentifier field to given value.


### SetTokenIdentifierNil

`func (o *Agent) SetTokenIdentifierNil(b bool)`

 SetTokenIdentifierNil sets the value for TokenIdentifier to be an explicit nil

### UnsetTokenIdentifier
`func (o *Agent) UnsetTokenIdentifier()`

UnsetTokenIdentifier ensures that no value is present for TokenIdentifier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



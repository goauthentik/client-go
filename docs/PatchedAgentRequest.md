# PatchedAgentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. | [optional] 
**Name** | Pointer to **string** | User&#39;s display name. | [optional] 
**IsActive** | Pointer to **bool** | Designates whether this user should be treated as active. Unselect this instead of deleting accounts. | [optional] 
**LastLogin** | Pointer to **NullableTime** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**Expiring** | Pointer to **bool** |  | [optional] 
**Expires** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewPatchedAgentRequest

`func NewPatchedAgentRequest() *PatchedAgentRequest`

NewPatchedAgentRequest instantiates a new PatchedAgentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedAgentRequestWithDefaults

`func NewPatchedAgentRequestWithDefaults() *PatchedAgentRequest`

NewPatchedAgentRequestWithDefaults instantiates a new PatchedAgentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *PatchedAgentRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PatchedAgentRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PatchedAgentRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PatchedAgentRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetName

`func (o *PatchedAgentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedAgentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedAgentRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedAgentRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsActive

`func (o *PatchedAgentRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PatchedAgentRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PatchedAgentRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *PatchedAgentRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetLastLogin

`func (o *PatchedAgentRequest) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *PatchedAgentRequest) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *PatchedAgentRequest) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *PatchedAgentRequest) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### SetLastLoginNil

`func (o *PatchedAgentRequest) SetLastLoginNil(b bool)`

 SetLastLoginNil sets the value for LastLogin to be an explicit nil

### UnsetLastLogin
`func (o *PatchedAgentRequest) UnsetLastLogin()`

UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
### GetEmail

`func (o *PatchedAgentRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PatchedAgentRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PatchedAgentRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PatchedAgentRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAttributes

`func (o *PatchedAgentRequest) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PatchedAgentRequest) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PatchedAgentRequest) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PatchedAgentRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetExpiring

`func (o *PatchedAgentRequest) GetExpiring() bool`

GetExpiring returns the Expiring field if non-nil, zero value otherwise.

### GetExpiringOk

`func (o *PatchedAgentRequest) GetExpiringOk() (*bool, bool)`

GetExpiringOk returns a tuple with the Expiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiring

`func (o *PatchedAgentRequest) SetExpiring(v bool)`

SetExpiring sets Expiring field to given value.

### HasExpiring

`func (o *PatchedAgentRequest) HasExpiring() bool`

HasExpiring returns a boolean if a field has been set.

### GetExpires

`func (o *PatchedAgentRequest) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *PatchedAgentRequest) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *PatchedAgentRequest) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *PatchedAgentRequest) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *PatchedAgentRequest) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *PatchedAgentRequest) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



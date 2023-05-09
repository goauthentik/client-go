# PatchedTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Managed** | Pointer to **NullableString** | Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update. | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**Intent** | Pointer to [**IntentEnum**](IntentEnum.md) |  | [optional] 
**User** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Expires** | Pointer to **time.Time** |  | [optional] 
**Expiring** | Pointer to **bool** |  | [optional] 

## Methods

### NewPatchedTokenRequest

`func NewPatchedTokenRequest() *PatchedTokenRequest`

NewPatchedTokenRequest instantiates a new PatchedTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedTokenRequestWithDefaults

`func NewPatchedTokenRequestWithDefaults() *PatchedTokenRequest`

NewPatchedTokenRequestWithDefaults instantiates a new PatchedTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManaged

`func (o *PatchedTokenRequest) GetManaged() string`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *PatchedTokenRequest) GetManagedOk() (*string, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *PatchedTokenRequest) SetManaged(v string)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *PatchedTokenRequest) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### SetManagedNil

`func (o *PatchedTokenRequest) SetManagedNil(b bool)`

 SetManagedNil sets the value for Managed to be an explicit nil

### UnsetManaged
`func (o *PatchedTokenRequest) UnsetManaged()`

UnsetManaged ensures that no value is present for Managed, not even an explicit nil
### GetIdentifier

`func (o *PatchedTokenRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PatchedTokenRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PatchedTokenRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *PatchedTokenRequest) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetIntent

`func (o *PatchedTokenRequest) GetIntent() IntentEnum`

GetIntent returns the Intent field if non-nil, zero value otherwise.

### GetIntentOk

`func (o *PatchedTokenRequest) GetIntentOk() (*IntentEnum, bool)`

GetIntentOk returns a tuple with the Intent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntent

`func (o *PatchedTokenRequest) SetIntent(v IntentEnum)`

SetIntent sets Intent field to given value.

### HasIntent

`func (o *PatchedTokenRequest) HasIntent() bool`

HasIntent returns a boolean if a field has been set.

### GetUser

`func (o *PatchedTokenRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchedTokenRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchedTokenRequest) SetUser(v int32)`

SetUser sets User field to given value.

### HasUser

`func (o *PatchedTokenRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedTokenRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedTokenRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedTokenRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedTokenRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpires

`func (o *PatchedTokenRequest) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *PatchedTokenRequest) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *PatchedTokenRequest) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *PatchedTokenRequest) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetExpiring

`func (o *PatchedTokenRequest) GetExpiring() bool`

GetExpiring returns the Expiring field if non-nil, zero value otherwise.

### GetExpiringOk

`func (o *PatchedTokenRequest) GetExpiringOk() (*bool, bool)`

GetExpiringOk returns a tuple with the Expiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiring

`func (o *PatchedTokenRequest) SetExpiring(v bool)`

SetExpiring sets Expiring field to given value.

### HasExpiring

`func (o *PatchedTokenRequest) HasExpiring() bool`

HasExpiring returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



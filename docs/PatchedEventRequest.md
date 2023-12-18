# PatchedEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to **interface{}** |  | [optional] 
**Action** | Pointer to [**EventActions**](EventActions.md) |  | [optional] 
**App** | Pointer to **string** |  | [optional] 
**Context** | Pointer to **interface{}** |  | [optional] 
**ClientIp** | Pointer to **NullableString** |  | [optional] 
**Expires** | Pointer to **time.Time** |  | [optional] 
**Tenant** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewPatchedEventRequest

`func NewPatchedEventRequest() *PatchedEventRequest`

NewPatchedEventRequest instantiates a new PatchedEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedEventRequestWithDefaults

`func NewPatchedEventRequestWithDefaults() *PatchedEventRequest`

NewPatchedEventRequestWithDefaults instantiates a new PatchedEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *PatchedEventRequest) GetUser() interface{}`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchedEventRequest) GetUserOk() (*interface{}, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchedEventRequest) SetUser(v interface{})`

SetUser sets User field to given value.

### HasUser

`func (o *PatchedEventRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *PatchedEventRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *PatchedEventRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetAction

`func (o *PatchedEventRequest) GetAction() EventActions`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PatchedEventRequest) GetActionOk() (*EventActions, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PatchedEventRequest) SetAction(v EventActions)`

SetAction sets Action field to given value.

### HasAction

`func (o *PatchedEventRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetApp

`func (o *PatchedEventRequest) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *PatchedEventRequest) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *PatchedEventRequest) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *PatchedEventRequest) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetContext

`func (o *PatchedEventRequest) GetContext() interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PatchedEventRequest) GetContextOk() (*interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PatchedEventRequest) SetContext(v interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *PatchedEventRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### SetContextNil

`func (o *PatchedEventRequest) SetContextNil(b bool)`

 SetContextNil sets the value for Context to be an explicit nil

### UnsetContext
`func (o *PatchedEventRequest) UnsetContext()`

UnsetContext ensures that no value is present for Context, not even an explicit nil
### GetClientIp

`func (o *PatchedEventRequest) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *PatchedEventRequest) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *PatchedEventRequest) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *PatchedEventRequest) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### SetClientIpNil

`func (o *PatchedEventRequest) SetClientIpNil(b bool)`

 SetClientIpNil sets the value for ClientIp to be an explicit nil

### UnsetClientIp
`func (o *PatchedEventRequest) UnsetClientIp()`

UnsetClientIp ensures that no value is present for ClientIp, not even an explicit nil
### GetExpires

`func (o *PatchedEventRequest) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *PatchedEventRequest) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *PatchedEventRequest) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *PatchedEventRequest) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedEventRequest) GetTenant() interface{}`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedEventRequest) GetTenantOk() (*interface{}, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedEventRequest) SetTenant(v interface{})`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedEventRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedEventRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedEventRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



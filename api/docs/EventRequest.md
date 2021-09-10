# EventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to **map[string]interface{}** |  | [optional] 
**Action** | [**EventActions**](EventActions.md) |  | 
**App** | **string** |  | 
**Context** | Pointer to **map[string]interface{}** |  | [optional] 
**ClientIp** | Pointer to **NullableString** |  | [optional] 
**Expires** | Pointer to **time.Time** |  | [optional] 
**Tenant** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewEventRequest

`func NewEventRequest(action EventActions, app string, ) *EventRequest`

NewEventRequest instantiates a new EventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventRequestWithDefaults

`func NewEventRequestWithDefaults() *EventRequest`

NewEventRequestWithDefaults instantiates a new EventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *EventRequest) GetUser() map[string]interface{}`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *EventRequest) GetUserOk() (*map[string]interface{}, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *EventRequest) SetUser(v map[string]interface{})`

SetUser sets User field to given value.

### HasUser

`func (o *EventRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetAction

`func (o *EventRequest) GetAction() EventActions`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *EventRequest) GetActionOk() (*EventActions, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *EventRequest) SetAction(v EventActions)`

SetAction sets Action field to given value.


### GetApp

`func (o *EventRequest) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *EventRequest) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *EventRequest) SetApp(v string)`

SetApp sets App field to given value.


### GetContext

`func (o *EventRequest) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *EventRequest) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *EventRequest) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *EventRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetClientIp

`func (o *EventRequest) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *EventRequest) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *EventRequest) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *EventRequest) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### SetClientIpNil

`func (o *EventRequest) SetClientIpNil(b bool)`

 SetClientIpNil sets the value for ClientIp to be an explicit nil

### UnsetClientIp
`func (o *EventRequest) UnsetClientIp()`

UnsetClientIp ensures that no value is present for ClientIp, not even an explicit nil
### GetExpires

`func (o *EventRequest) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *EventRequest) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *EventRequest) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *EventRequest) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetTenant

`func (o *EventRequest) GetTenant() map[string]interface{}`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *EventRequest) GetTenantOk() (*map[string]interface{}, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *EventRequest) SetTenant(v map[string]interface{})`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *EventRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# EventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to **interface{}** |  | [optional] 
**Action** | [**EventActions**](EventActions.md) |  | 
**App** | **string** |  | 
**Context** | Pointer to **interface{}** |  | [optional] 
**ClientIp** | Pointer to **NullableString** |  | [optional] 
**Expires** | Pointer to **time.Time** |  | [optional] 
**Brand** | Pointer to **interface{}** |  | [optional] 

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

`func (o *EventRequest) GetUser() interface{}`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *EventRequest) GetUserOk() (*interface{}, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *EventRequest) SetUser(v interface{})`

SetUser sets User field to given value.

### HasUser

`func (o *EventRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *EventRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *EventRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
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

`func (o *EventRequest) GetContext() interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *EventRequest) GetContextOk() (*interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *EventRequest) SetContext(v interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *EventRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### SetContextNil

`func (o *EventRequest) SetContextNil(b bool)`

 SetContextNil sets the value for Context to be an explicit nil

### UnsetContext
`func (o *EventRequest) UnsetContext()`

UnsetContext ensures that no value is present for Context, not even an explicit nil
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

### GetBrand

`func (o *EventRequest) GetBrand() interface{}`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *EventRequest) GetBrandOk() (*interface{}, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *EventRequest) SetBrand(v interface{})`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *EventRequest) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *EventRequest) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *EventRequest) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



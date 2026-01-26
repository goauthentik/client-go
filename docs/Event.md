# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**User** | Pointer to **interface{}** |  | [optional] 
**Action** | [**EventActions**](EventActions.md) |  | 
**App** | **string** |  | 
**Context** | Pointer to **interface{}** |  | [optional] 
**ClientIp** | Pointer to **NullableString** |  | [optional] 
**Created** | **time.Time** |  | [readonly] 
**Expires** | Pointer to **time.Time** |  | [optional] 
**Brand** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewEvent

`func NewEvent(pk string, action EventActions, app string, created time.Time, ) *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *Event) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *Event) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *Event) SetPk(v string)`

SetPk sets Pk field to given value.


### GetUser

`func (o *Event) GetUser() interface{}`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Event) GetUserOk() (*interface{}, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Event) SetUser(v interface{})`

SetUser sets User field to given value.

### HasUser

`func (o *Event) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *Event) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *Event) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetAction

`func (o *Event) GetAction() EventActions`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Event) GetActionOk() (*EventActions, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Event) SetAction(v EventActions)`

SetAction sets Action field to given value.


### GetApp

`func (o *Event) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *Event) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *Event) SetApp(v string)`

SetApp sets App field to given value.


### GetContext

`func (o *Event) GetContext() interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Event) GetContextOk() (*interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Event) SetContext(v interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *Event) HasContext() bool`

HasContext returns a boolean if a field has been set.

### SetContextNil

`func (o *Event) SetContextNil(b bool)`

 SetContextNil sets the value for Context to be an explicit nil

### UnsetContext
`func (o *Event) UnsetContext()`

UnsetContext ensures that no value is present for Context, not even an explicit nil
### GetClientIp

`func (o *Event) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *Event) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *Event) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *Event) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### SetClientIpNil

`func (o *Event) SetClientIpNil(b bool)`

 SetClientIpNil sets the value for ClientIp to be an explicit nil

### UnsetClientIp
`func (o *Event) UnsetClientIp()`

UnsetClientIp ensures that no value is present for ClientIp, not even an explicit nil
### GetCreated

`func (o *Event) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Event) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Event) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetExpires

`func (o *Event) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *Event) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *Event) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *Event) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetBrand

`func (o *Event) GetBrand() interface{}`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *Event) GetBrandOk() (*interface{}, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *Event) SetBrand(v interface{})`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *Event) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *Event) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *Event) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SCIMProviderUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ScimId** | **string** |  | 
**User** | **int32** |  | 
**UserObj** | [**PartialUser**](PartialUser.md) |  | [readonly] 
**Provider** | **int32** |  | 
**Attributes** | **map[string]interface{}** |  | [readonly] 

## Methods

### NewSCIMProviderUser

`func NewSCIMProviderUser(id string, scimId string, user int32, userObj PartialUser, provider int32, attributes map[string]interface{}, ) *SCIMProviderUser`

NewSCIMProviderUser instantiates a new SCIMProviderUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSCIMProviderUserWithDefaults

`func NewSCIMProviderUserWithDefaults() *SCIMProviderUser`

NewSCIMProviderUserWithDefaults instantiates a new SCIMProviderUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SCIMProviderUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SCIMProviderUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SCIMProviderUser) SetId(v string)`

SetId sets Id field to given value.


### GetScimId

`func (o *SCIMProviderUser) GetScimId() string`

GetScimId returns the ScimId field if non-nil, zero value otherwise.

### GetScimIdOk

`func (o *SCIMProviderUser) GetScimIdOk() (*string, bool)`

GetScimIdOk returns a tuple with the ScimId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScimId

`func (o *SCIMProviderUser) SetScimId(v string)`

SetScimId sets ScimId field to given value.


### GetUser

`func (o *SCIMProviderUser) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SCIMProviderUser) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SCIMProviderUser) SetUser(v int32)`

SetUser sets User field to given value.


### GetUserObj

`func (o *SCIMProviderUser) GetUserObj() PartialUser`

GetUserObj returns the UserObj field if non-nil, zero value otherwise.

### GetUserObjOk

`func (o *SCIMProviderUser) GetUserObjOk() (*PartialUser, bool)`

GetUserObjOk returns a tuple with the UserObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObj

`func (o *SCIMProviderUser) SetUserObj(v PartialUser)`

SetUserObj sets UserObj field to given value.


### GetProvider

`func (o *SCIMProviderUser) GetProvider() int32`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SCIMProviderUser) GetProviderOk() (*int32, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SCIMProviderUser) SetProvider(v int32)`

SetProvider sets Provider field to given value.


### GetAttributes

`func (o *SCIMProviderUser) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SCIMProviderUser) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SCIMProviderUser) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# MicrosoftEntraProviderUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**MicrosoftId** | **string** |  | 
**User** | **int32** |  | 
**UserObj** | [**GroupMember**](GroupMember.md) |  | [readonly] 
**Provider** | **int32** |  | 
**Attributes** | **map[string]interface{}** |  | [readonly] 

## Methods

### NewMicrosoftEntraProviderUser

`func NewMicrosoftEntraProviderUser(id string, microsoftId string, user int32, userObj GroupMember, provider int32, attributes map[string]interface{}, ) *MicrosoftEntraProviderUser`

NewMicrosoftEntraProviderUser instantiates a new MicrosoftEntraProviderUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftEntraProviderUserWithDefaults

`func NewMicrosoftEntraProviderUserWithDefaults() *MicrosoftEntraProviderUser`

NewMicrosoftEntraProviderUserWithDefaults instantiates a new MicrosoftEntraProviderUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftEntraProviderUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftEntraProviderUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftEntraProviderUser) SetId(v string)`

SetId sets Id field to given value.


### GetMicrosoftId

`func (o *MicrosoftEntraProviderUser) GetMicrosoftId() string`

GetMicrosoftId returns the MicrosoftId field if non-nil, zero value otherwise.

### GetMicrosoftIdOk

`func (o *MicrosoftEntraProviderUser) GetMicrosoftIdOk() (*string, bool)`

GetMicrosoftIdOk returns a tuple with the MicrosoftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftId

`func (o *MicrosoftEntraProviderUser) SetMicrosoftId(v string)`

SetMicrosoftId sets MicrosoftId field to given value.


### GetUser

`func (o *MicrosoftEntraProviderUser) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MicrosoftEntraProviderUser) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MicrosoftEntraProviderUser) SetUser(v int32)`

SetUser sets User field to given value.


### GetUserObj

`func (o *MicrosoftEntraProviderUser) GetUserObj() GroupMember`

GetUserObj returns the UserObj field if non-nil, zero value otherwise.

### GetUserObjOk

`func (o *MicrosoftEntraProviderUser) GetUserObjOk() (*GroupMember, bool)`

GetUserObjOk returns a tuple with the UserObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObj

`func (o *MicrosoftEntraProviderUser) SetUserObj(v GroupMember)`

SetUserObj sets UserObj field to given value.


### GetProvider

`func (o *MicrosoftEntraProviderUser) GetProvider() int32`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *MicrosoftEntraProviderUser) GetProviderOk() (*int32, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *MicrosoftEntraProviderUser) SetProvider(v int32)`

SetProvider sets Provider field to given value.


### GetAttributes

`func (o *MicrosoftEntraProviderUser) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MicrosoftEntraProviderUser) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MicrosoftEntraProviderUser) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SCIMSourceUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**User** | **int32** |  | 
**UserObj** | [**GroupMember**](GroupMember.md) |  | [readonly] 
**Source** | **string** |  | 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSCIMSourceUser

`func NewSCIMSourceUser(id string, user int32, userObj GroupMember, source string, ) *SCIMSourceUser`

NewSCIMSourceUser instantiates a new SCIMSourceUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSCIMSourceUserWithDefaults

`func NewSCIMSourceUserWithDefaults() *SCIMSourceUser`

NewSCIMSourceUserWithDefaults instantiates a new SCIMSourceUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SCIMSourceUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SCIMSourceUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SCIMSourceUser) SetId(v string)`

SetId sets Id field to given value.


### GetUser

`func (o *SCIMSourceUser) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SCIMSourceUser) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SCIMSourceUser) SetUser(v int32)`

SetUser sets User field to given value.


### GetUserObj

`func (o *SCIMSourceUser) GetUserObj() GroupMember`

GetUserObj returns the UserObj field if non-nil, zero value otherwise.

### GetUserObjOk

`func (o *SCIMSourceUser) GetUserObjOk() (*GroupMember, bool)`

GetUserObjOk returns a tuple with the UserObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObj

`func (o *SCIMSourceUser) SetUserObj(v GroupMember)`

SetUserObj sets UserObj field to given value.


### GetSource

`func (o *SCIMSourceUser) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SCIMSourceUser) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SCIMSourceUser) SetSource(v string)`

SetSource sets Source field to given value.


### GetAttributes

`func (o *SCIMSourceUser) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SCIMSourceUser) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SCIMSourceUser) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SCIMSourceUser) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



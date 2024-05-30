# GoogleWorkspaceProviderUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**User** | **int32** |  | 
**UserObj** | [**GroupMember**](GroupMember.md) |  | [readonly] 
**Provider** | **int32** |  | 
**Attributes** | **interface{}** |  | [readonly] 

## Methods

### NewGoogleWorkspaceProviderUser

`func NewGoogleWorkspaceProviderUser(id string, user int32, userObj GroupMember, provider int32, attributes interface{}, ) *GoogleWorkspaceProviderUser`

NewGoogleWorkspaceProviderUser instantiates a new GoogleWorkspaceProviderUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleWorkspaceProviderUserWithDefaults

`func NewGoogleWorkspaceProviderUserWithDefaults() *GoogleWorkspaceProviderUser`

NewGoogleWorkspaceProviderUserWithDefaults instantiates a new GoogleWorkspaceProviderUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GoogleWorkspaceProviderUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GoogleWorkspaceProviderUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GoogleWorkspaceProviderUser) SetId(v string)`

SetId sets Id field to given value.


### GetUser

`func (o *GoogleWorkspaceProviderUser) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GoogleWorkspaceProviderUser) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GoogleWorkspaceProviderUser) SetUser(v int32)`

SetUser sets User field to given value.


### GetUserObj

`func (o *GoogleWorkspaceProviderUser) GetUserObj() GroupMember`

GetUserObj returns the UserObj field if non-nil, zero value otherwise.

### GetUserObjOk

`func (o *GoogleWorkspaceProviderUser) GetUserObjOk() (*GroupMember, bool)`

GetUserObjOk returns a tuple with the UserObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObj

`func (o *GoogleWorkspaceProviderUser) SetUserObj(v GroupMember)`

SetUserObj sets UserObj field to given value.


### GetProvider

`func (o *GoogleWorkspaceProviderUser) GetProvider() int32`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *GoogleWorkspaceProviderUser) GetProviderOk() (*int32, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *GoogleWorkspaceProviderUser) SetProvider(v int32)`

SetProvider sets Provider field to given value.


### GetAttributes

`func (o *GoogleWorkspaceProviderUser) GetAttributes() interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GoogleWorkspaceProviderUser) GetAttributesOk() (*interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GoogleWorkspaceProviderUser) SetAttributes(v interface{})`

SetAttributes sets Attributes field to given value.


### SetAttributesNil

`func (o *GoogleWorkspaceProviderUser) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *GoogleWorkspaceProviderUser) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



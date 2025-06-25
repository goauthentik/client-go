# GoogleWorkspaceProviderGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**GoogleId** | **string** |  | 
**Group** | **string** |  | 
**GroupObj** | [**UserGroup**](UserGroup.md) |  | [readonly] 
**Provider** | **int32** |  | 
**Attributes** | **map[string]interface{}** |  | [readonly] 

## Methods

### NewGoogleWorkspaceProviderGroup

`func NewGoogleWorkspaceProviderGroup(id string, googleId string, group string, groupObj UserGroup, provider int32, attributes map[string]interface{}, ) *GoogleWorkspaceProviderGroup`

NewGoogleWorkspaceProviderGroup instantiates a new GoogleWorkspaceProviderGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleWorkspaceProviderGroupWithDefaults

`func NewGoogleWorkspaceProviderGroupWithDefaults() *GoogleWorkspaceProviderGroup`

NewGoogleWorkspaceProviderGroupWithDefaults instantiates a new GoogleWorkspaceProviderGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GoogleWorkspaceProviderGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GoogleWorkspaceProviderGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GoogleWorkspaceProviderGroup) SetId(v string)`

SetId sets Id field to given value.


### GetGoogleId

`func (o *GoogleWorkspaceProviderGroup) GetGoogleId() string`

GetGoogleId returns the GoogleId field if non-nil, zero value otherwise.

### GetGoogleIdOk

`func (o *GoogleWorkspaceProviderGroup) GetGoogleIdOk() (*string, bool)`

GetGoogleIdOk returns a tuple with the GoogleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleId

`func (o *GoogleWorkspaceProviderGroup) SetGoogleId(v string)`

SetGoogleId sets GoogleId field to given value.


### GetGroup

`func (o *GoogleWorkspaceProviderGroup) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GoogleWorkspaceProviderGroup) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GoogleWorkspaceProviderGroup) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetGroupObj

`func (o *GoogleWorkspaceProviderGroup) GetGroupObj() UserGroup`

GetGroupObj returns the GroupObj field if non-nil, zero value otherwise.

### GetGroupObjOk

`func (o *GoogleWorkspaceProviderGroup) GetGroupObjOk() (*UserGroup, bool)`

GetGroupObjOk returns a tuple with the GroupObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupObj

`func (o *GoogleWorkspaceProviderGroup) SetGroupObj(v UserGroup)`

SetGroupObj sets GroupObj field to given value.


### GetProvider

`func (o *GoogleWorkspaceProviderGroup) GetProvider() int32`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *GoogleWorkspaceProviderGroup) GetProviderOk() (*int32, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *GoogleWorkspaceProviderGroup) SetProvider(v int32)`

SetProvider sets Provider field to given value.


### GetAttributes

`func (o *GoogleWorkspaceProviderGroup) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GoogleWorkspaceProviderGroup) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GoogleWorkspaceProviderGroup) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# MicrosoftEntraProviderGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**MicrosoftId** | **string** |  | 
**Group** | **string** |  | 
**GroupObj** | [**UserGroup**](UserGroup.md) |  | [readonly] 
**Provider** | **int32** |  | 
**Attributes** | **map[string]interface{}** |  | [readonly] 

## Methods

### NewMicrosoftEntraProviderGroup

`func NewMicrosoftEntraProviderGroup(id string, microsoftId string, group string, groupObj UserGroup, provider int32, attributes map[string]interface{}, ) *MicrosoftEntraProviderGroup`

NewMicrosoftEntraProviderGroup instantiates a new MicrosoftEntraProviderGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftEntraProviderGroupWithDefaults

`func NewMicrosoftEntraProviderGroupWithDefaults() *MicrosoftEntraProviderGroup`

NewMicrosoftEntraProviderGroupWithDefaults instantiates a new MicrosoftEntraProviderGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftEntraProviderGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftEntraProviderGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftEntraProviderGroup) SetId(v string)`

SetId sets Id field to given value.


### GetMicrosoftId

`func (o *MicrosoftEntraProviderGroup) GetMicrosoftId() string`

GetMicrosoftId returns the MicrosoftId field if non-nil, zero value otherwise.

### GetMicrosoftIdOk

`func (o *MicrosoftEntraProviderGroup) GetMicrosoftIdOk() (*string, bool)`

GetMicrosoftIdOk returns a tuple with the MicrosoftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftId

`func (o *MicrosoftEntraProviderGroup) SetMicrosoftId(v string)`

SetMicrosoftId sets MicrosoftId field to given value.


### GetGroup

`func (o *MicrosoftEntraProviderGroup) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *MicrosoftEntraProviderGroup) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *MicrosoftEntraProviderGroup) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetGroupObj

`func (o *MicrosoftEntraProviderGroup) GetGroupObj() UserGroup`

GetGroupObj returns the GroupObj field if non-nil, zero value otherwise.

### GetGroupObjOk

`func (o *MicrosoftEntraProviderGroup) GetGroupObjOk() (*UserGroup, bool)`

GetGroupObjOk returns a tuple with the GroupObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupObj

`func (o *MicrosoftEntraProviderGroup) SetGroupObj(v UserGroup)`

SetGroupObj sets GroupObj field to given value.


### GetProvider

`func (o *MicrosoftEntraProviderGroup) GetProvider() int32`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *MicrosoftEntraProviderGroup) GetProviderOk() (*int32, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *MicrosoftEntraProviderGroup) SetProvider(v int32)`

SetProvider sets Provider field to given value.


### GetAttributes

`func (o *MicrosoftEntraProviderGroup) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MicrosoftEntraProviderGroup) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MicrosoftEntraProviderGroup) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



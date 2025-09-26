# SCIMProviderGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ScimId** | **string** |  | 
**Group** | **string** |  | 
**GroupObj** | [**PartialGroup**](PartialGroup.md) |  | [readonly] 
**Provider** | **int32** |  | 
**Attributes** | **map[string]interface{}** |  | [readonly] 

## Methods

### NewSCIMProviderGroup

`func NewSCIMProviderGroup(id string, scimId string, group string, groupObj PartialGroup, provider int32, attributes map[string]interface{}, ) *SCIMProviderGroup`

NewSCIMProviderGroup instantiates a new SCIMProviderGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSCIMProviderGroupWithDefaults

`func NewSCIMProviderGroupWithDefaults() *SCIMProviderGroup`

NewSCIMProviderGroupWithDefaults instantiates a new SCIMProviderGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SCIMProviderGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SCIMProviderGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SCIMProviderGroup) SetId(v string)`

SetId sets Id field to given value.


### GetScimId

`func (o *SCIMProviderGroup) GetScimId() string`

GetScimId returns the ScimId field if non-nil, zero value otherwise.

### GetScimIdOk

`func (o *SCIMProviderGroup) GetScimIdOk() (*string, bool)`

GetScimIdOk returns a tuple with the ScimId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScimId

`func (o *SCIMProviderGroup) SetScimId(v string)`

SetScimId sets ScimId field to given value.


### GetGroup

`func (o *SCIMProviderGroup) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SCIMProviderGroup) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SCIMProviderGroup) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetGroupObj

`func (o *SCIMProviderGroup) GetGroupObj() PartialGroup`

GetGroupObj returns the GroupObj field if non-nil, zero value otherwise.

### GetGroupObjOk

`func (o *SCIMProviderGroup) GetGroupObjOk() (*PartialGroup, bool)`

GetGroupObjOk returns a tuple with the GroupObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupObj

`func (o *SCIMProviderGroup) SetGroupObj(v PartialGroup)`

SetGroupObj sets GroupObj field to given value.


### GetProvider

`func (o *SCIMProviderGroup) GetProvider() int32`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SCIMProviderGroup) GetProviderOk() (*int32, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SCIMProviderGroup) SetProvider(v int32)`

SetProvider sets Provider field to given value.


### GetAttributes

`func (o *SCIMProviderGroup) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SCIMProviderGroup) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SCIMProviderGroup) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SCIMSourceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ExternalId** | **string** |  | 
**Group** | **string** |  | 
**GroupObj** | [**PartialGroup**](PartialGroup.md) |  | [readonly] 
**Source** | **string** |  | 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSCIMSourceGroup

`func NewSCIMSourceGroup(externalId string, group string, groupObj PartialGroup, source string, ) *SCIMSourceGroup`

NewSCIMSourceGroup instantiates a new SCIMSourceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSCIMSourceGroupWithDefaults

`func NewSCIMSourceGroupWithDefaults() *SCIMSourceGroup`

NewSCIMSourceGroupWithDefaults instantiates a new SCIMSourceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SCIMSourceGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SCIMSourceGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SCIMSourceGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SCIMSourceGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExternalId

`func (o *SCIMSourceGroup) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SCIMSourceGroup) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SCIMSourceGroup) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetGroup

`func (o *SCIMSourceGroup) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SCIMSourceGroup) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SCIMSourceGroup) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetGroupObj

`func (o *SCIMSourceGroup) GetGroupObj() PartialGroup`

GetGroupObj returns the GroupObj field if non-nil, zero value otherwise.

### GetGroupObjOk

`func (o *SCIMSourceGroup) GetGroupObjOk() (*PartialGroup, bool)`

GetGroupObjOk returns a tuple with the GroupObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupObj

`func (o *SCIMSourceGroup) SetGroupObj(v PartialGroup)`

SetGroupObj sets GroupObj field to given value.


### GetSource

`func (o *SCIMSourceGroup) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SCIMSourceGroup) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SCIMSourceGroup) SetSource(v string)`

SetSource sets Source field to given value.


### GetAttributes

`func (o *SCIMSourceGroup) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SCIMSourceGroup) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SCIMSourceGroup) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SCIMSourceGroup) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



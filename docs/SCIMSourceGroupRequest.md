# SCIMSourceGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ExternalId** | **string** |  | 
**Group** | **string** |  | 
**Source** | **string** |  | 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSCIMSourceGroupRequest

`func NewSCIMSourceGroupRequest(externalId string, group string, source string, ) *SCIMSourceGroupRequest`

NewSCIMSourceGroupRequest instantiates a new SCIMSourceGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSCIMSourceGroupRequestWithDefaults

`func NewSCIMSourceGroupRequestWithDefaults() *SCIMSourceGroupRequest`

NewSCIMSourceGroupRequestWithDefaults instantiates a new SCIMSourceGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SCIMSourceGroupRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SCIMSourceGroupRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SCIMSourceGroupRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SCIMSourceGroupRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExternalId

`func (o *SCIMSourceGroupRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SCIMSourceGroupRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SCIMSourceGroupRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetGroup

`func (o *SCIMSourceGroupRequest) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SCIMSourceGroupRequest) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SCIMSourceGroupRequest) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetSource

`func (o *SCIMSourceGroupRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SCIMSourceGroupRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SCIMSourceGroupRequest) SetSource(v string)`

SetSource sets Source field to given value.


### GetAttributes

`func (o *SCIMSourceGroupRequest) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SCIMSourceGroupRequest) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SCIMSourceGroupRequest) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SCIMSourceGroupRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



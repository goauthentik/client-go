# PatchedSCIMSourceUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**User** | Pointer to **int32** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewPatchedSCIMSourceUserRequest

`func NewPatchedSCIMSourceUserRequest() *PatchedSCIMSourceUserRequest`

NewPatchedSCIMSourceUserRequest instantiates a new PatchedSCIMSourceUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedSCIMSourceUserRequestWithDefaults

`func NewPatchedSCIMSourceUserRequestWithDefaults() *PatchedSCIMSourceUserRequest`

NewPatchedSCIMSourceUserRequestWithDefaults instantiates a new PatchedSCIMSourceUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedSCIMSourceUserRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedSCIMSourceUserRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedSCIMSourceUserRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedSCIMSourceUserRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUser

`func (o *PatchedSCIMSourceUserRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchedSCIMSourceUserRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchedSCIMSourceUserRequest) SetUser(v int32)`

SetUser sets User field to given value.

### HasUser

`func (o *PatchedSCIMSourceUserRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetSource

`func (o *PatchedSCIMSourceUserRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PatchedSCIMSourceUserRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PatchedSCIMSourceUserRequest) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PatchedSCIMSourceUserRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetAttributes

`func (o *PatchedSCIMSourceUserRequest) GetAttributes() interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PatchedSCIMSourceUserRequest) GetAttributesOk() (*interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PatchedSCIMSourceUserRequest) SetAttributes(v interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PatchedSCIMSourceUserRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *PatchedSCIMSourceUserRequest) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *PatchedSCIMSourceUserRequest) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



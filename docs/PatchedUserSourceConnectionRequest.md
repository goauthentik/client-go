# PatchedUserSourceConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to **int32** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedUserSourceConnectionRequest

`func NewPatchedUserSourceConnectionRequest() *PatchedUserSourceConnectionRequest`

NewPatchedUserSourceConnectionRequest instantiates a new PatchedUserSourceConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedUserSourceConnectionRequestWithDefaults

`func NewPatchedUserSourceConnectionRequestWithDefaults() *PatchedUserSourceConnectionRequest`

NewPatchedUserSourceConnectionRequestWithDefaults instantiates a new PatchedUserSourceConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *PatchedUserSourceConnectionRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchedUserSourceConnectionRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchedUserSourceConnectionRequest) SetUser(v int32)`

SetUser sets User field to given value.

### HasUser

`func (o *PatchedUserSourceConnectionRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetSource

`func (o *PatchedUserSourceConnectionRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PatchedUserSourceConnectionRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PatchedUserSourceConnectionRequest) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PatchedUserSourceConnectionRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetIdentifier

`func (o *PatchedUserSourceConnectionRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PatchedUserSourceConnectionRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PatchedUserSourceConnectionRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *PatchedUserSourceConnectionRequest) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



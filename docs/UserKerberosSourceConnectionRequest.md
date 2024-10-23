# UserKerberosSourceConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | **int32** |  | 
**Identifier** | **string** |  | 

## Methods

### NewUserKerberosSourceConnectionRequest

`func NewUserKerberosSourceConnectionRequest(user int32, identifier string, ) *UserKerberosSourceConnectionRequest`

NewUserKerberosSourceConnectionRequest instantiates a new UserKerberosSourceConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserKerberosSourceConnectionRequestWithDefaults

`func NewUserKerberosSourceConnectionRequestWithDefaults() *UserKerberosSourceConnectionRequest`

NewUserKerberosSourceConnectionRequestWithDefaults instantiates a new UserKerberosSourceConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *UserKerberosSourceConnectionRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserKerberosSourceConnectionRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserKerberosSourceConnectionRequest) SetUser(v int32)`

SetUser sets User field to given value.


### GetIdentifier

`func (o *UserKerberosSourceConnectionRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *UserKerberosSourceConnectionRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *UserKerberosSourceConnectionRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



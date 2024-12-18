# UserSourceConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | **int32** |  | 
**Source** | **string** |  | 

## Methods

### NewUserSourceConnectionRequest

`func NewUserSourceConnectionRequest(user int32, source string, ) *UserSourceConnectionRequest`

NewUserSourceConnectionRequest instantiates a new UserSourceConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSourceConnectionRequestWithDefaults

`func NewUserSourceConnectionRequestWithDefaults() *UserSourceConnectionRequest`

NewUserSourceConnectionRequestWithDefaults instantiates a new UserSourceConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *UserSourceConnectionRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserSourceConnectionRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserSourceConnectionRequest) SetUser(v int32)`

SetUser sets User field to given value.


### GetSource

`func (o *UserSourceConnectionRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UserSourceConnectionRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UserSourceConnectionRequest) SetSource(v string)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



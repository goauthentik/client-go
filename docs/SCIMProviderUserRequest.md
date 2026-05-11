# SCIMProviderUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScimId** | **string** |  | 
**User** | **int32** |  | 
**Provider** | **int32** |  | 

## Methods

### NewSCIMProviderUserRequest

`func NewSCIMProviderUserRequest(scimId string, user int32, provider int32, ) *SCIMProviderUserRequest`

NewSCIMProviderUserRequest instantiates a new SCIMProviderUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSCIMProviderUserRequestWithDefaults

`func NewSCIMProviderUserRequestWithDefaults() *SCIMProviderUserRequest`

NewSCIMProviderUserRequestWithDefaults instantiates a new SCIMProviderUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScimId

`func (o *SCIMProviderUserRequest) GetScimId() string`

GetScimId returns the ScimId field if non-nil, zero value otherwise.

### GetScimIdOk

`func (o *SCIMProviderUserRequest) GetScimIdOk() (*string, bool)`

GetScimIdOk returns a tuple with the ScimId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScimId

`func (o *SCIMProviderUserRequest) SetScimId(v string)`

SetScimId sets ScimId field to given value.


### GetUser

`func (o *SCIMProviderUserRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SCIMProviderUserRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SCIMProviderUserRequest) SetUser(v int32)`

SetUser sets User field to given value.


### GetProvider

`func (o *SCIMProviderUserRequest) GetProvider() int32`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SCIMProviderUserRequest) GetProviderOk() (*int32, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SCIMProviderUserRequest) SetProvider(v int32)`

SetProvider sets Provider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



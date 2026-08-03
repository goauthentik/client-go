# GoogleWorkspaceProviderUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GoogleId** | **string** |  | 
**User** | **int32** |  | 
**Provider** | **int32** |  | 

## Methods

### NewGoogleWorkspaceProviderUserRequest

`func NewGoogleWorkspaceProviderUserRequest(googleId string, user int32, provider int32, ) *GoogleWorkspaceProviderUserRequest`

NewGoogleWorkspaceProviderUserRequest instantiates a new GoogleWorkspaceProviderUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleWorkspaceProviderUserRequestWithDefaults

`func NewGoogleWorkspaceProviderUserRequestWithDefaults() *GoogleWorkspaceProviderUserRequest`

NewGoogleWorkspaceProviderUserRequestWithDefaults instantiates a new GoogleWorkspaceProviderUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGoogleId

`func (o *GoogleWorkspaceProviderUserRequest) GetGoogleId() string`

GetGoogleId returns the GoogleId field if non-nil, zero value otherwise.

### GetGoogleIdOk

`func (o *GoogleWorkspaceProviderUserRequest) GetGoogleIdOk() (*string, bool)`

GetGoogleIdOk returns a tuple with the GoogleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleId

`func (o *GoogleWorkspaceProviderUserRequest) SetGoogleId(v string)`

SetGoogleId sets GoogleId field to given value.


### GetUser

`func (o *GoogleWorkspaceProviderUserRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GoogleWorkspaceProviderUserRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GoogleWorkspaceProviderUserRequest) SetUser(v int32)`

SetUser sets User field to given value.


### GetProvider

`func (o *GoogleWorkspaceProviderUserRequest) GetProvider() int32`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *GoogleWorkspaceProviderUserRequest) GetProviderOk() (*int32, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *GoogleWorkspaceProviderUserRequest) SetProvider(v int32)`

SetProvider sets Provider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



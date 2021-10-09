# UserSelfRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. | 
**Name** | **string** | User&#39;s display name. | 
**Email** | Pointer to **string** |  | [optional] 
**Settings** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUserSelfRequest

`func NewUserSelfRequest(username string, name string, ) *UserSelfRequest`

NewUserSelfRequest instantiates a new UserSelfRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSelfRequestWithDefaults

`func NewUserSelfRequestWithDefaults() *UserSelfRequest`

NewUserSelfRequestWithDefaults instantiates a new UserSelfRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserSelfRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserSelfRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserSelfRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetName

`func (o *UserSelfRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserSelfRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserSelfRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *UserSelfRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserSelfRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserSelfRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserSelfRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSettings

`func (o *UserSelfRequest) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *UserSelfRequest) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *UserSelfRequest) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *UserSelfRequest) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



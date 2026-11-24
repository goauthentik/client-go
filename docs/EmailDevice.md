# EmailDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The human-readable name of this device. | 
**Pk** | **int32** |  | [readonly] 
**Email** | **string** |  | [readonly] 
**User** | [**PartialUser**](PartialUser.md) |  | [readonly] 

## Methods

### NewEmailDevice

`func NewEmailDevice(name string, pk int32, email string, user PartialUser, ) *EmailDevice`

NewEmailDevice instantiates a new EmailDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailDeviceWithDefaults

`func NewEmailDeviceWithDefaults() *EmailDevice`

NewEmailDeviceWithDefaults instantiates a new EmailDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EmailDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailDevice) SetName(v string)`

SetName sets Name field to given value.


### GetPk

`func (o *EmailDevice) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *EmailDevice) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *EmailDevice) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetEmail

`func (o *EmailDevice) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmailDevice) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmailDevice) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetUser

`func (o *EmailDevice) GetUser() PartialUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *EmailDevice) GetUserOk() (*PartialUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *EmailDevice) SetUser(v PartialUser)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



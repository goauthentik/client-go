# DeviceUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Username** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Home** | Pointer to **string** |  | [optional] 

## Methods

### NewDeviceUserRequest

`func NewDeviceUserRequest(id string, ) *DeviceUserRequest`

NewDeviceUserRequest instantiates a new DeviceUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceUserRequestWithDefaults

`func NewDeviceUserRequestWithDefaults() *DeviceUserRequest`

NewDeviceUserRequestWithDefaults instantiates a new DeviceUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceUserRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceUserRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceUserRequest) SetId(v string)`

SetId sets Id field to given value.


### GetUsername

`func (o *DeviceUserRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DeviceUserRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DeviceUserRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DeviceUserRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetName

`func (o *DeviceUserRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceUserRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceUserRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceUserRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHome

`func (o *DeviceUserRequest) GetHome() string`

GetHome returns the Home field if non-nil, zero value otherwise.

### GetHomeOk

`func (o *DeviceUserRequest) GetHomeOk() (*string, bool)`

GetHomeOk returns a tuple with the Home field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHome

`func (o *DeviceUserRequest) SetHome(v string)`

SetHome sets Home field to given value.

### HasHome

`func (o *DeviceUserRequest) HasHome() bool`

HasHome returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



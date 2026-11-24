# DeviceAccessGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDeviceAccessGroupRequest

`func NewDeviceAccessGroupRequest(name string, ) *DeviceAccessGroupRequest`

NewDeviceAccessGroupRequest instantiates a new DeviceAccessGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAccessGroupRequestWithDefaults

`func NewDeviceAccessGroupRequestWithDefaults() *DeviceAccessGroupRequest`

NewDeviceAccessGroupRequestWithDefaults instantiates a new DeviceAccessGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeviceAccessGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceAccessGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceAccessGroupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAttributes

`func (o *DeviceAccessGroupRequest) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *DeviceAccessGroupRequest) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *DeviceAccessGroupRequest) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *DeviceAccessGroupRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



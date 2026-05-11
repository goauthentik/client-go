# DeviceAccessGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PbmUuid** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDeviceAccessGroup

`func NewDeviceAccessGroup(pbmUuid string, name string, ) *DeviceAccessGroup`

NewDeviceAccessGroup instantiates a new DeviceAccessGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAccessGroupWithDefaults

`func NewDeviceAccessGroupWithDefaults() *DeviceAccessGroup`

NewDeviceAccessGroupWithDefaults instantiates a new DeviceAccessGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPbmUuid

`func (o *DeviceAccessGroup) GetPbmUuid() string`

GetPbmUuid returns the PbmUuid field if non-nil, zero value otherwise.

### GetPbmUuidOk

`func (o *DeviceAccessGroup) GetPbmUuidOk() (*string, bool)`

GetPbmUuidOk returns a tuple with the PbmUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPbmUuid

`func (o *DeviceAccessGroup) SetPbmUuid(v string)`

SetPbmUuid sets PbmUuid field to given value.


### GetName

`func (o *DeviceAccessGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceAccessGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceAccessGroup) SetName(v string)`

SetName sets Name field to given value.


### GetAttributes

`func (o *DeviceAccessGroup) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *DeviceAccessGroup) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *DeviceAccessGroup) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *DeviceAccessGroup) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



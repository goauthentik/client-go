# StaticDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The human-readable name of this device. | 
**TokenSet** | [**[]StaticDeviceToken**](StaticDeviceToken.md) |  | [readonly] 
**Pk** | **int32** |  | [readonly] 

## Methods

### NewStaticDevice

`func NewStaticDevice(name string, tokenSet []StaticDeviceToken, pk int32, ) *StaticDevice`

NewStaticDevice instantiates a new StaticDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticDeviceWithDefaults

`func NewStaticDeviceWithDefaults() *StaticDevice`

NewStaticDeviceWithDefaults instantiates a new StaticDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StaticDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StaticDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StaticDevice) SetName(v string)`

SetName sets Name field to given value.


### GetTokenSet

`func (o *StaticDevice) GetTokenSet() []StaticDeviceToken`

GetTokenSet returns the TokenSet field if non-nil, zero value otherwise.

### GetTokenSetOk

`func (o *StaticDevice) GetTokenSetOk() (*[]StaticDeviceToken, bool)`

GetTokenSetOk returns a tuple with the TokenSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSet

`func (o *StaticDevice) SetTokenSet(v []StaticDeviceToken)`

SetTokenSet sets TokenSet field to given value.


### GetPk

`func (o *StaticDevice) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *StaticDevice) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *StaticDevice) SetPk(v int32)`

SetPk sets Pk field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



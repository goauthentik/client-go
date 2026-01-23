# OperatingSystemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Family** | [**DeviceFactsOSFamily**](DeviceFactsOSFamily.md) |  | 
**Name** | Pointer to **string** | Operating System name, such as &#39;Server 2022&#39; or &#39;Ubuntu&#39; | [optional] 
**Version** | Pointer to **string** | Operating System version, must always be the version number but may contain build name | [optional] 
**Arch** | **string** |  | 

## Methods

### NewOperatingSystemRequest

`func NewOperatingSystemRequest(family DeviceFactsOSFamily, arch string, ) *OperatingSystemRequest`

NewOperatingSystemRequest instantiates a new OperatingSystemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatingSystemRequestWithDefaults

`func NewOperatingSystemRequestWithDefaults() *OperatingSystemRequest`

NewOperatingSystemRequestWithDefaults instantiates a new OperatingSystemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFamily

`func (o *OperatingSystemRequest) GetFamily() DeviceFactsOSFamily`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *OperatingSystemRequest) GetFamilyOk() (*DeviceFactsOSFamily, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *OperatingSystemRequest) SetFamily(v DeviceFactsOSFamily)`

SetFamily sets Family field to given value.


### GetName

`func (o *OperatingSystemRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OperatingSystemRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OperatingSystemRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OperatingSystemRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *OperatingSystemRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OperatingSystemRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OperatingSystemRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OperatingSystemRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetArch

`func (o *OperatingSystemRequest) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *OperatingSystemRequest) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *OperatingSystemRequest) SetArch(v string)`

SetArch sets Arch field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



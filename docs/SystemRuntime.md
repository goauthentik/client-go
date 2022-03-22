# SystemRuntime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PythonVersion** | **string** |  | 
**GunicornVersion** | **string** |  | 
**Environment** | **string** |  | 
**Architecture** | **string** |  | 
**Platform** | **string** |  | 
**Uname** | **string** |  | 

## Methods

### NewSystemRuntime

`func NewSystemRuntime(pythonVersion string, gunicornVersion string, environment string, architecture string, platform string, uname string, ) *SystemRuntime`

NewSystemRuntime instantiates a new SystemRuntime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemRuntimeWithDefaults

`func NewSystemRuntimeWithDefaults() *SystemRuntime`

NewSystemRuntimeWithDefaults instantiates a new SystemRuntime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPythonVersion

`func (o *SystemRuntime) GetPythonVersion() string`

GetPythonVersion returns the PythonVersion field if non-nil, zero value otherwise.

### GetPythonVersionOk

`func (o *SystemRuntime) GetPythonVersionOk() (*string, bool)`

GetPythonVersionOk returns a tuple with the PythonVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPythonVersion

`func (o *SystemRuntime) SetPythonVersion(v string)`

SetPythonVersion sets PythonVersion field to given value.


### GetGunicornVersion

`func (o *SystemRuntime) GetGunicornVersion() string`

GetGunicornVersion returns the GunicornVersion field if non-nil, zero value otherwise.

### GetGunicornVersionOk

`func (o *SystemRuntime) GetGunicornVersionOk() (*string, bool)`

GetGunicornVersionOk returns a tuple with the GunicornVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGunicornVersion

`func (o *SystemRuntime) SetGunicornVersion(v string)`

SetGunicornVersion sets GunicornVersion field to given value.


### GetEnvironment

`func (o *SystemRuntime) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SystemRuntime) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SystemRuntime) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetArchitecture

`func (o *SystemRuntime) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *SystemRuntime) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *SystemRuntime) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.


### GetPlatform

`func (o *SystemRuntime) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *SystemRuntime) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *SystemRuntime) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetUname

`func (o *SystemRuntime) GetUname() string`

GetUname returns the Uname field if non-nil, zero value otherwise.

### GetUnameOk

`func (o *SystemRuntime) GetUnameOk() (*string, bool)`

GetUnameOk returns a tuple with the Uname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUname

`func (o *SystemRuntime) SetUname(v string)`

SetUname sets Uname field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



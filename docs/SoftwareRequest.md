# SoftwareRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Version** | Pointer to **string** |  | [optional] 
**Source** | **string** |  | 
**Path** | Pointer to **string** |  | [optional] 

## Methods

### NewSoftwareRequest

`func NewSoftwareRequest(name string, source string, ) *SoftwareRequest`

NewSoftwareRequest instantiates a new SoftwareRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareRequestWithDefaults

`func NewSoftwareRequestWithDefaults() *SoftwareRequest`

NewSoftwareRequestWithDefaults instantiates a new SoftwareRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SoftwareRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SoftwareRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SoftwareRequest) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *SoftwareRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SoftwareRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SoftwareRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SoftwareRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSource

`func (o *SoftwareRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SoftwareRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SoftwareRequest) SetSource(v string)`

SetSource sets Source field to given value.


### GetPath

`func (o *SoftwareRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SoftwareRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SoftwareRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SoftwareRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



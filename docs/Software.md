# Software

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Version** | Pointer to **string** |  | [optional] 
**Source** | **string** |  | 
**Path** | Pointer to **string** |  | [optional] 

## Methods

### NewSoftware

`func NewSoftware(name string, source string, ) *Software`

NewSoftware instantiates a new Software object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareWithDefaults

`func NewSoftwareWithDefaults() *Software`

NewSoftwareWithDefaults instantiates a new Software object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Software) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Software) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Software) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *Software) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Software) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Software) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Software) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSource

`func (o *Software) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Software) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Software) SetSource(v string)`

SetSource sets Source field to given value.


### GetPath

`func (o *Software) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Software) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Software) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Software) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



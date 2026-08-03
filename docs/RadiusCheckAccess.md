# RadiusCheckAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to **string** |  | [optional] 
**Access** | [**PolicyTestResult**](PolicyTestResult.md) |  | 

## Methods

### NewRadiusCheckAccess

`func NewRadiusCheckAccess(access PolicyTestResult, ) *RadiusCheckAccess`

NewRadiusCheckAccess instantiates a new RadiusCheckAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadiusCheckAccessWithDefaults

`func NewRadiusCheckAccessWithDefaults() *RadiusCheckAccess`

NewRadiusCheckAccessWithDefaults instantiates a new RadiusCheckAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *RadiusCheckAccess) GetAttributes() string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *RadiusCheckAccess) GetAttributesOk() (*string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *RadiusCheckAccess) SetAttributes(v string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *RadiusCheckAccess) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetAccess

`func (o *RadiusCheckAccess) GetAccess() PolicyTestResult`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *RadiusCheckAccess) GetAccessOk() (*PolicyTestResult, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *RadiusCheckAccess) SetAccess(v PolicyTestResult)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



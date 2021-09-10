# ValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NonFieldErrors** | Pointer to **[]string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 

## Methods

### NewValidationError

`func NewValidationError() *ValidationError`

NewValidationError instantiates a new ValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationErrorWithDefaults

`func NewValidationErrorWithDefaults() *ValidationError`

NewValidationErrorWithDefaults instantiates a new ValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNonFieldErrors

`func (o *ValidationError) GetNonFieldErrors() []string`

GetNonFieldErrors returns the NonFieldErrors field if non-nil, zero value otherwise.

### GetNonFieldErrorsOk

`func (o *ValidationError) GetNonFieldErrorsOk() (*[]string, bool)`

GetNonFieldErrorsOk returns a tuple with the NonFieldErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonFieldErrors

`func (o *ValidationError) SetNonFieldErrors(v []string)`

SetNonFieldErrors sets NonFieldErrors field to given value.

### HasNonFieldErrors

`func (o *ValidationError) HasNonFieldErrors() bool`

HasNonFieldErrors returns a boolean if a field has been set.

### GetCode

`func (o *ValidationError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ValidationError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ValidationError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ValidationError) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



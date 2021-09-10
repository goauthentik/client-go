# IdentificationChallengeResponseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to **string** |  | [optional] [default to "ak-stage-identification"]
**UidField** | **string** |  | 
**Password** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIdentificationChallengeResponseRequest

`func NewIdentificationChallengeResponseRequest(uidField string, ) *IdentificationChallengeResponseRequest`

NewIdentificationChallengeResponseRequest instantiates a new IdentificationChallengeResponseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentificationChallengeResponseRequestWithDefaults

`func NewIdentificationChallengeResponseRequestWithDefaults() *IdentificationChallengeResponseRequest`

NewIdentificationChallengeResponseRequestWithDefaults instantiates a new IdentificationChallengeResponseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *IdentificationChallengeResponseRequest) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *IdentificationChallengeResponseRequest) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *IdentificationChallengeResponseRequest) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *IdentificationChallengeResponseRequest) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetUidField

`func (o *IdentificationChallengeResponseRequest) GetUidField() string`

GetUidField returns the UidField field if non-nil, zero value otherwise.

### GetUidFieldOk

`func (o *IdentificationChallengeResponseRequest) GetUidFieldOk() (*string, bool)`

GetUidFieldOk returns a tuple with the UidField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidField

`func (o *IdentificationChallengeResponseRequest) SetUidField(v string)`

SetUidField sets UidField field to given value.


### GetPassword

`func (o *IdentificationChallengeResponseRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *IdentificationChallengeResponseRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *IdentificationChallengeResponseRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *IdentificationChallengeResponseRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *IdentificationChallengeResponseRequest) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *IdentificationChallengeResponseRequest) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



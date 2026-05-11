# UserRecoveryEmailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenDuration** | Pointer to **string** |  | [optional] 
**EmailStage** | **string** |  | 

## Methods

### NewUserRecoveryEmailRequest

`func NewUserRecoveryEmailRequest(emailStage string, ) *UserRecoveryEmailRequest`

NewUserRecoveryEmailRequest instantiates a new UserRecoveryEmailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRecoveryEmailRequestWithDefaults

`func NewUserRecoveryEmailRequestWithDefaults() *UserRecoveryEmailRequest`

NewUserRecoveryEmailRequestWithDefaults instantiates a new UserRecoveryEmailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenDuration

`func (o *UserRecoveryEmailRequest) GetTokenDuration() string`

GetTokenDuration returns the TokenDuration field if non-nil, zero value otherwise.

### GetTokenDurationOk

`func (o *UserRecoveryEmailRequest) GetTokenDurationOk() (*string, bool)`

GetTokenDurationOk returns a tuple with the TokenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenDuration

`func (o *UserRecoveryEmailRequest) SetTokenDuration(v string)`

SetTokenDuration sets TokenDuration field to given value.

### HasTokenDuration

`func (o *UserRecoveryEmailRequest) HasTokenDuration() bool`

HasTokenDuration returns a boolean if a field has been set.

### GetEmailStage

`func (o *UserRecoveryEmailRequest) GetEmailStage() string`

GetEmailStage returns the EmailStage field if non-nil, zero value otherwise.

### GetEmailStageOk

`func (o *UserRecoveryEmailRequest) GetEmailStageOk() (*string, bool)`

GetEmailStageOk returns a tuple with the EmailStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailStage

`func (o *UserRecoveryEmailRequest) SetEmailStage(v string)`

SetEmailStage sets EmailStage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



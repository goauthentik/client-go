# PatchedCaptchaStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**PublicKey** | Pointer to **string** | Public key, acquired your captcha Provider. | [optional] 
**PrivateKey** | Pointer to **string** | Private key, acquired your captcha Provider. | [optional] 
**JsUrl** | Pointer to **string** |  | [optional] 
**ApiUrl** | Pointer to **string** |  | [optional] 
**Interactive** | Pointer to **bool** |  | [optional] 
**ScoreMinThreshold** | Pointer to **float64** |  | [optional] 
**ScoreMaxThreshold** | Pointer to **float64** |  | [optional] 
**ErrorOnInvalidScore** | Pointer to **bool** | When enabled and the received captcha score is outside of the given threshold, the stage will show an error message. When not enabled, the flow will continue, but the data from the captcha will be available in the context for policy decisions | [optional] 

## Methods

### NewPatchedCaptchaStageRequest

`func NewPatchedCaptchaStageRequest() *PatchedCaptchaStageRequest`

NewPatchedCaptchaStageRequest instantiates a new PatchedCaptchaStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedCaptchaStageRequestWithDefaults

`func NewPatchedCaptchaStageRequestWithDefaults() *PatchedCaptchaStageRequest`

NewPatchedCaptchaStageRequestWithDefaults instantiates a new PatchedCaptchaStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedCaptchaStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedCaptchaStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedCaptchaStageRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedCaptchaStageRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublicKey

`func (o *PatchedCaptchaStageRequest) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *PatchedCaptchaStageRequest) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *PatchedCaptchaStageRequest) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *PatchedCaptchaStageRequest) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetPrivateKey

`func (o *PatchedCaptchaStageRequest) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *PatchedCaptchaStageRequest) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *PatchedCaptchaStageRequest) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *PatchedCaptchaStageRequest) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetJsUrl

`func (o *PatchedCaptchaStageRequest) GetJsUrl() string`

GetJsUrl returns the JsUrl field if non-nil, zero value otherwise.

### GetJsUrlOk

`func (o *PatchedCaptchaStageRequest) GetJsUrlOk() (*string, bool)`

GetJsUrlOk returns a tuple with the JsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsUrl

`func (o *PatchedCaptchaStageRequest) SetJsUrl(v string)`

SetJsUrl sets JsUrl field to given value.

### HasJsUrl

`func (o *PatchedCaptchaStageRequest) HasJsUrl() bool`

HasJsUrl returns a boolean if a field has been set.

### GetApiUrl

`func (o *PatchedCaptchaStageRequest) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *PatchedCaptchaStageRequest) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *PatchedCaptchaStageRequest) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *PatchedCaptchaStageRequest) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### GetInteractive

`func (o *PatchedCaptchaStageRequest) GetInteractive() bool`

GetInteractive returns the Interactive field if non-nil, zero value otherwise.

### GetInteractiveOk

`func (o *PatchedCaptchaStageRequest) GetInteractiveOk() (*bool, bool)`

GetInteractiveOk returns a tuple with the Interactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractive

`func (o *PatchedCaptchaStageRequest) SetInteractive(v bool)`

SetInteractive sets Interactive field to given value.

### HasInteractive

`func (o *PatchedCaptchaStageRequest) HasInteractive() bool`

HasInteractive returns a boolean if a field has been set.

### GetScoreMinThreshold

`func (o *PatchedCaptchaStageRequest) GetScoreMinThreshold() float64`

GetScoreMinThreshold returns the ScoreMinThreshold field if non-nil, zero value otherwise.

### GetScoreMinThresholdOk

`func (o *PatchedCaptchaStageRequest) GetScoreMinThresholdOk() (*float64, bool)`

GetScoreMinThresholdOk returns a tuple with the ScoreMinThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreMinThreshold

`func (o *PatchedCaptchaStageRequest) SetScoreMinThreshold(v float64)`

SetScoreMinThreshold sets ScoreMinThreshold field to given value.

### HasScoreMinThreshold

`func (o *PatchedCaptchaStageRequest) HasScoreMinThreshold() bool`

HasScoreMinThreshold returns a boolean if a field has been set.

### GetScoreMaxThreshold

`func (o *PatchedCaptchaStageRequest) GetScoreMaxThreshold() float64`

GetScoreMaxThreshold returns the ScoreMaxThreshold field if non-nil, zero value otherwise.

### GetScoreMaxThresholdOk

`func (o *PatchedCaptchaStageRequest) GetScoreMaxThresholdOk() (*float64, bool)`

GetScoreMaxThresholdOk returns a tuple with the ScoreMaxThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreMaxThreshold

`func (o *PatchedCaptchaStageRequest) SetScoreMaxThreshold(v float64)`

SetScoreMaxThreshold sets ScoreMaxThreshold field to given value.

### HasScoreMaxThreshold

`func (o *PatchedCaptchaStageRequest) HasScoreMaxThreshold() bool`

HasScoreMaxThreshold returns a boolean if a field has been set.

### GetErrorOnInvalidScore

`func (o *PatchedCaptchaStageRequest) GetErrorOnInvalidScore() bool`

GetErrorOnInvalidScore returns the ErrorOnInvalidScore field if non-nil, zero value otherwise.

### GetErrorOnInvalidScoreOk

`func (o *PatchedCaptchaStageRequest) GetErrorOnInvalidScoreOk() (*bool, bool)`

GetErrorOnInvalidScoreOk returns a tuple with the ErrorOnInvalidScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorOnInvalidScore

`func (o *PatchedCaptchaStageRequest) SetErrorOnInvalidScore(v bool)`

SetErrorOnInvalidScore sets ErrorOnInvalidScore field to given value.

### HasErrorOnInvalidScore

`func (o *PatchedCaptchaStageRequest) HasErrorOnInvalidScore() bool`

HasErrorOnInvalidScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



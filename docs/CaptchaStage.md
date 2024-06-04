# CaptchaStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Component** | **string** | Get object type so that we know how to edit the object | [readonly] 
**VerboseName** | **string** | Return object&#39;s verbose_name | [readonly] 
**VerboseNamePlural** | **string** | Return object&#39;s plural verbose_name | [readonly] 
**MetaModelName** | **string** | Return internal model name | [readonly] 
**FlowSet** | Pointer to [**[]FlowSet**](FlowSet.md) |  | [optional] 
**PublicKey** | **string** | Public key, acquired your captcha Provider. | 
**JsUrl** | Pointer to **string** |  | [optional] 
**ApiUrl** | Pointer to **string** |  | [optional] 
**ScoreMinThreshold** | Pointer to **float64** |  | [optional] 
**ScoreMaxThreshold** | Pointer to **float64** |  | [optional] 
**ErrorOnInvalidScore** | Pointer to **bool** | When enabled and the received captcha score is outside of the given threshold, the stage will show an error message. When not enabled, the flow will continue, but the data from the captcha will be available in the context for policy decisions | [optional] 

## Methods

### NewCaptchaStage

`func NewCaptchaStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, publicKey string, ) *CaptchaStage`

NewCaptchaStage instantiates a new CaptchaStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptchaStageWithDefaults

`func NewCaptchaStageWithDefaults() *CaptchaStage`

NewCaptchaStageWithDefaults instantiates a new CaptchaStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *CaptchaStage) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *CaptchaStage) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *CaptchaStage) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *CaptchaStage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CaptchaStage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CaptchaStage) SetName(v string)`

SetName sets Name field to given value.


### GetComponent

`func (o *CaptchaStage) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *CaptchaStage) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *CaptchaStage) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *CaptchaStage) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *CaptchaStage) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *CaptchaStage) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *CaptchaStage) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *CaptchaStage) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *CaptchaStage) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *CaptchaStage) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *CaptchaStage) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *CaptchaStage) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetFlowSet

`func (o *CaptchaStage) GetFlowSet() []FlowSet`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *CaptchaStage) GetFlowSetOk() (*[]FlowSet, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *CaptchaStage) SetFlowSet(v []FlowSet)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *CaptchaStage) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetPublicKey

`func (o *CaptchaStage) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *CaptchaStage) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *CaptchaStage) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetJsUrl

`func (o *CaptchaStage) GetJsUrl() string`

GetJsUrl returns the JsUrl field if non-nil, zero value otherwise.

### GetJsUrlOk

`func (o *CaptchaStage) GetJsUrlOk() (*string, bool)`

GetJsUrlOk returns a tuple with the JsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsUrl

`func (o *CaptchaStage) SetJsUrl(v string)`

SetJsUrl sets JsUrl field to given value.

### HasJsUrl

`func (o *CaptchaStage) HasJsUrl() bool`

HasJsUrl returns a boolean if a field has been set.

### GetApiUrl

`func (o *CaptchaStage) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *CaptchaStage) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *CaptchaStage) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *CaptchaStage) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### GetScoreMinThreshold

`func (o *CaptchaStage) GetScoreMinThreshold() float64`

GetScoreMinThreshold returns the ScoreMinThreshold field if non-nil, zero value otherwise.

### GetScoreMinThresholdOk

`func (o *CaptchaStage) GetScoreMinThresholdOk() (*float64, bool)`

GetScoreMinThresholdOk returns a tuple with the ScoreMinThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreMinThreshold

`func (o *CaptchaStage) SetScoreMinThreshold(v float64)`

SetScoreMinThreshold sets ScoreMinThreshold field to given value.

### HasScoreMinThreshold

`func (o *CaptchaStage) HasScoreMinThreshold() bool`

HasScoreMinThreshold returns a boolean if a field has been set.

### GetScoreMaxThreshold

`func (o *CaptchaStage) GetScoreMaxThreshold() float64`

GetScoreMaxThreshold returns the ScoreMaxThreshold field if non-nil, zero value otherwise.

### GetScoreMaxThresholdOk

`func (o *CaptchaStage) GetScoreMaxThresholdOk() (*float64, bool)`

GetScoreMaxThresholdOk returns a tuple with the ScoreMaxThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreMaxThreshold

`func (o *CaptchaStage) SetScoreMaxThreshold(v float64)`

SetScoreMaxThreshold sets ScoreMaxThreshold field to given value.

### HasScoreMaxThreshold

`func (o *CaptchaStage) HasScoreMaxThreshold() bool`

HasScoreMaxThreshold returns a boolean if a field has been set.

### GetErrorOnInvalidScore

`func (o *CaptchaStage) GetErrorOnInvalidScore() bool`

GetErrorOnInvalidScore returns the ErrorOnInvalidScore field if non-nil, zero value otherwise.

### GetErrorOnInvalidScoreOk

`func (o *CaptchaStage) GetErrorOnInvalidScoreOk() (*bool, bool)`

GetErrorOnInvalidScoreOk returns a tuple with the ErrorOnInvalidScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorOnInvalidScore

`func (o *CaptchaStage) SetErrorOnInvalidScore(v bool)`

SetErrorOnInvalidScore sets ErrorOnInvalidScore field to given value.

### HasErrorOnInvalidScore

`func (o *CaptchaStage) HasErrorOnInvalidScore() bool`

HasErrorOnInvalidScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



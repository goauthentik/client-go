# CaptchaStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**FlowSet** | Pointer to [**[]FlowSetRequest**](FlowSetRequest.md) |  | [optional] 
**PublicKey** | **string** | Public key, acquired your captcha Provider. | 
**PrivateKey** | **string** | Private key, acquired your captcha Provider. | 
**JsUrl** | Pointer to **string** |  | [optional] 
**ApiUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewCaptchaStageRequest

`func NewCaptchaStageRequest(name string, publicKey string, privateKey string, ) *CaptchaStageRequest`

NewCaptchaStageRequest instantiates a new CaptchaStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptchaStageRequestWithDefaults

`func NewCaptchaStageRequestWithDefaults() *CaptchaStageRequest`

NewCaptchaStageRequestWithDefaults instantiates a new CaptchaStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CaptchaStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CaptchaStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CaptchaStageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFlowSet

`func (o *CaptchaStageRequest) GetFlowSet() []FlowSetRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *CaptchaStageRequest) GetFlowSetOk() (*[]FlowSetRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *CaptchaStageRequest) SetFlowSet(v []FlowSetRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *CaptchaStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetPublicKey

`func (o *CaptchaStageRequest) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *CaptchaStageRequest) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *CaptchaStageRequest) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetPrivateKey

`func (o *CaptchaStageRequest) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CaptchaStageRequest) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CaptchaStageRequest) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetJsUrl

`func (o *CaptchaStageRequest) GetJsUrl() string`

GetJsUrl returns the JsUrl field if non-nil, zero value otherwise.

### GetJsUrlOk

`func (o *CaptchaStageRequest) GetJsUrlOk() (*string, bool)`

GetJsUrlOk returns a tuple with the JsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsUrl

`func (o *CaptchaStageRequest) SetJsUrl(v string)`

SetJsUrl sets JsUrl field to given value.

### HasJsUrl

`func (o *CaptchaStageRequest) HasJsUrl() bool`

HasJsUrl returns a boolean if a field has been set.

### GetApiUrl

`func (o *CaptchaStageRequest) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *CaptchaStageRequest) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *CaptchaStageRequest) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *CaptchaStageRequest) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



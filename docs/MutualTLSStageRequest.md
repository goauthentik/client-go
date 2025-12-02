# MutualTLSStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**FlowSet** | Pointer to [**[]FlowSetRequest**](FlowSetRequest.md) |  | [optional] 
**Mode** | [**StageModeEnum**](StageModeEnum.md) |  | 
**CertificateAuthorities** | Pointer to **[]string** | Configure certificate authorities to validate the certificate against. This option has a higher priority than the &#x60;client_certificate&#x60; option on &#x60;Brand&#x60;. | [optional] 
**CertAttribute** | [**CertAttributeEnum**](CertAttributeEnum.md) |  | 
**UserAttribute** | [**UserAttributeEnum**](UserAttributeEnum.md) |  | 

## Methods

### NewMutualTLSStageRequest

`func NewMutualTLSStageRequest(name string, mode StageModeEnum, certAttribute CertAttributeEnum, userAttribute UserAttributeEnum, ) *MutualTLSStageRequest`

NewMutualTLSStageRequest instantiates a new MutualTLSStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutualTLSStageRequestWithDefaults

`func NewMutualTLSStageRequestWithDefaults() *MutualTLSStageRequest`

NewMutualTLSStageRequestWithDefaults instantiates a new MutualTLSStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MutualTLSStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MutualTLSStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MutualTLSStageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFlowSet

`func (o *MutualTLSStageRequest) GetFlowSet() []FlowSetRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *MutualTLSStageRequest) GetFlowSetOk() (*[]FlowSetRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *MutualTLSStageRequest) SetFlowSet(v []FlowSetRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *MutualTLSStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetMode

`func (o *MutualTLSStageRequest) GetMode() StageModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *MutualTLSStageRequest) GetModeOk() (*StageModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *MutualTLSStageRequest) SetMode(v StageModeEnum)`

SetMode sets Mode field to given value.


### GetCertificateAuthorities

`func (o *MutualTLSStageRequest) GetCertificateAuthorities() []string`

GetCertificateAuthorities returns the CertificateAuthorities field if non-nil, zero value otherwise.

### GetCertificateAuthoritiesOk

`func (o *MutualTLSStageRequest) GetCertificateAuthoritiesOk() (*[]string, bool)`

GetCertificateAuthoritiesOk returns a tuple with the CertificateAuthorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorities

`func (o *MutualTLSStageRequest) SetCertificateAuthorities(v []string)`

SetCertificateAuthorities sets CertificateAuthorities field to given value.

### HasCertificateAuthorities

`func (o *MutualTLSStageRequest) HasCertificateAuthorities() bool`

HasCertificateAuthorities returns a boolean if a field has been set.

### GetCertAttribute

`func (o *MutualTLSStageRequest) GetCertAttribute() CertAttributeEnum`

GetCertAttribute returns the CertAttribute field if non-nil, zero value otherwise.

### GetCertAttributeOk

`func (o *MutualTLSStageRequest) GetCertAttributeOk() (*CertAttributeEnum, bool)`

GetCertAttributeOk returns a tuple with the CertAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertAttribute

`func (o *MutualTLSStageRequest) SetCertAttribute(v CertAttributeEnum)`

SetCertAttribute sets CertAttribute field to given value.


### GetUserAttribute

`func (o *MutualTLSStageRequest) GetUserAttribute() UserAttributeEnum`

GetUserAttribute returns the UserAttribute field if non-nil, zero value otherwise.

### GetUserAttributeOk

`func (o *MutualTLSStageRequest) GetUserAttributeOk() (*UserAttributeEnum, bool)`

GetUserAttributeOk returns a tuple with the UserAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttribute

`func (o *MutualTLSStageRequest) SetUserAttribute(v UserAttributeEnum)`

SetUserAttribute sets UserAttribute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



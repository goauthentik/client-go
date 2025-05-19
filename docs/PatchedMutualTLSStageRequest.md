# PatchedMutualTLSStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**FlowSet** | Pointer to [**[]FlowSetRequest**](FlowSetRequest.md) |  | [optional] 
**Mode** | Pointer to [**MutualTLSStageModeEnum**](MutualTLSStageModeEnum.md) |  | [optional] 
**CertificateAuthorities** | Pointer to **[]string** | Configure certificate authorities to validate the certificate against. This option has a higher priority than the &#x60;client_certificate&#x60; option on &#x60;Brand&#x60;. | [optional] 
**CertAttribute** | Pointer to [**CertAttributeEnum**](CertAttributeEnum.md) |  | [optional] 
**UserAttribute** | Pointer to [**UserAttributeEnum**](UserAttributeEnum.md) |  | [optional] 

## Methods

### NewPatchedMutualTLSStageRequest

`func NewPatchedMutualTLSStageRequest() *PatchedMutualTLSStageRequest`

NewPatchedMutualTLSStageRequest instantiates a new PatchedMutualTLSStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedMutualTLSStageRequestWithDefaults

`func NewPatchedMutualTLSStageRequestWithDefaults() *PatchedMutualTLSStageRequest`

NewPatchedMutualTLSStageRequestWithDefaults instantiates a new PatchedMutualTLSStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedMutualTLSStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedMutualTLSStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedMutualTLSStageRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedMutualTLSStageRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFlowSet

`func (o *PatchedMutualTLSStageRequest) GetFlowSet() []FlowSetRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *PatchedMutualTLSStageRequest) GetFlowSetOk() (*[]FlowSetRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *PatchedMutualTLSStageRequest) SetFlowSet(v []FlowSetRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *PatchedMutualTLSStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetMode

`func (o *PatchedMutualTLSStageRequest) GetMode() MutualTLSStageModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PatchedMutualTLSStageRequest) GetModeOk() (*MutualTLSStageModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PatchedMutualTLSStageRequest) SetMode(v MutualTLSStageModeEnum)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PatchedMutualTLSStageRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetCertificateAuthorities

`func (o *PatchedMutualTLSStageRequest) GetCertificateAuthorities() []string`

GetCertificateAuthorities returns the CertificateAuthorities field if non-nil, zero value otherwise.

### GetCertificateAuthoritiesOk

`func (o *PatchedMutualTLSStageRequest) GetCertificateAuthoritiesOk() (*[]string, bool)`

GetCertificateAuthoritiesOk returns a tuple with the CertificateAuthorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorities

`func (o *PatchedMutualTLSStageRequest) SetCertificateAuthorities(v []string)`

SetCertificateAuthorities sets CertificateAuthorities field to given value.

### HasCertificateAuthorities

`func (o *PatchedMutualTLSStageRequest) HasCertificateAuthorities() bool`

HasCertificateAuthorities returns a boolean if a field has been set.

### GetCertAttribute

`func (o *PatchedMutualTLSStageRequest) GetCertAttribute() CertAttributeEnum`

GetCertAttribute returns the CertAttribute field if non-nil, zero value otherwise.

### GetCertAttributeOk

`func (o *PatchedMutualTLSStageRequest) GetCertAttributeOk() (*CertAttributeEnum, bool)`

GetCertAttributeOk returns a tuple with the CertAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertAttribute

`func (o *PatchedMutualTLSStageRequest) SetCertAttribute(v CertAttributeEnum)`

SetCertAttribute sets CertAttribute field to given value.

### HasCertAttribute

`func (o *PatchedMutualTLSStageRequest) HasCertAttribute() bool`

HasCertAttribute returns a boolean if a field has been set.

### GetUserAttribute

`func (o *PatchedMutualTLSStageRequest) GetUserAttribute() UserAttributeEnum`

GetUserAttribute returns the UserAttribute field if non-nil, zero value otherwise.

### GetUserAttributeOk

`func (o *PatchedMutualTLSStageRequest) GetUserAttributeOk() (*UserAttributeEnum, bool)`

GetUserAttributeOk returns a tuple with the UserAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttribute

`func (o *PatchedMutualTLSStageRequest) SetUserAttribute(v UserAttributeEnum)`

SetUserAttribute sets UserAttribute field to given value.

### HasUserAttribute

`func (o *PatchedMutualTLSStageRequest) HasUserAttribute() bool`

HasUserAttribute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



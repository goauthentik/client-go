# TenantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | Domain that activates this tenant. Can be a superset, i.e. &#x60;a.b&#x60; for &#x60;aa.b&#x60; and &#x60;ba.b&#x60; | 
**Default** | Pointer to **bool** |  | [optional] 
**BrandingTitle** | Pointer to **string** |  | [optional] 
**BrandingLogo** | Pointer to **string** |  | [optional] 
**BrandingFavicon** | Pointer to **string** |  | [optional] 
**FlowAuthentication** | Pointer to **NullableString** |  | [optional] 
**FlowInvalidation** | Pointer to **NullableString** |  | [optional] 
**FlowRecovery** | Pointer to **NullableString** |  | [optional] 
**FlowUnenrollment** | Pointer to **NullableString** |  | [optional] 
**FlowUserSettings** | Pointer to **NullableString** |  | [optional] 
**FlowDeviceCode** | Pointer to **NullableString** |  | [optional] 
**EventRetention** | Pointer to **string** | Events will be deleted after this duration.(Format: weeks&#x3D;3;days&#x3D;2;hours&#x3D;3,seconds&#x3D;2). | [optional] 
**WebCertificate** | Pointer to **NullableString** | Web Certificate used by the authentik Core webserver. | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewTenantRequest

`func NewTenantRequest(domain string, ) *TenantRequest`

NewTenantRequest instantiates a new TenantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantRequestWithDefaults

`func NewTenantRequestWithDefaults() *TenantRequest`

NewTenantRequestWithDefaults instantiates a new TenantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *TenantRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *TenantRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *TenantRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetDefault

`func (o *TenantRequest) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *TenantRequest) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *TenantRequest) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *TenantRequest) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetBrandingTitle

`func (o *TenantRequest) GetBrandingTitle() string`

GetBrandingTitle returns the BrandingTitle field if non-nil, zero value otherwise.

### GetBrandingTitleOk

`func (o *TenantRequest) GetBrandingTitleOk() (*string, bool)`

GetBrandingTitleOk returns a tuple with the BrandingTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingTitle

`func (o *TenantRequest) SetBrandingTitle(v string)`

SetBrandingTitle sets BrandingTitle field to given value.

### HasBrandingTitle

`func (o *TenantRequest) HasBrandingTitle() bool`

HasBrandingTitle returns a boolean if a field has been set.

### GetBrandingLogo

`func (o *TenantRequest) GetBrandingLogo() string`

GetBrandingLogo returns the BrandingLogo field if non-nil, zero value otherwise.

### GetBrandingLogoOk

`func (o *TenantRequest) GetBrandingLogoOk() (*string, bool)`

GetBrandingLogoOk returns a tuple with the BrandingLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingLogo

`func (o *TenantRequest) SetBrandingLogo(v string)`

SetBrandingLogo sets BrandingLogo field to given value.

### HasBrandingLogo

`func (o *TenantRequest) HasBrandingLogo() bool`

HasBrandingLogo returns a boolean if a field has been set.

### GetBrandingFavicon

`func (o *TenantRequest) GetBrandingFavicon() string`

GetBrandingFavicon returns the BrandingFavicon field if non-nil, zero value otherwise.

### GetBrandingFaviconOk

`func (o *TenantRequest) GetBrandingFaviconOk() (*string, bool)`

GetBrandingFaviconOk returns a tuple with the BrandingFavicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingFavicon

`func (o *TenantRequest) SetBrandingFavicon(v string)`

SetBrandingFavicon sets BrandingFavicon field to given value.

### HasBrandingFavicon

`func (o *TenantRequest) HasBrandingFavicon() bool`

HasBrandingFavicon returns a boolean if a field has been set.

### GetFlowAuthentication

`func (o *TenantRequest) GetFlowAuthentication() string`

GetFlowAuthentication returns the FlowAuthentication field if non-nil, zero value otherwise.

### GetFlowAuthenticationOk

`func (o *TenantRequest) GetFlowAuthenticationOk() (*string, bool)`

GetFlowAuthenticationOk returns a tuple with the FlowAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowAuthentication

`func (o *TenantRequest) SetFlowAuthentication(v string)`

SetFlowAuthentication sets FlowAuthentication field to given value.

### HasFlowAuthentication

`func (o *TenantRequest) HasFlowAuthentication() bool`

HasFlowAuthentication returns a boolean if a field has been set.

### SetFlowAuthenticationNil

`func (o *TenantRequest) SetFlowAuthenticationNil(b bool)`

 SetFlowAuthenticationNil sets the value for FlowAuthentication to be an explicit nil

### UnsetFlowAuthentication
`func (o *TenantRequest) UnsetFlowAuthentication()`

UnsetFlowAuthentication ensures that no value is present for FlowAuthentication, not even an explicit nil
### GetFlowInvalidation

`func (o *TenantRequest) GetFlowInvalidation() string`

GetFlowInvalidation returns the FlowInvalidation field if non-nil, zero value otherwise.

### GetFlowInvalidationOk

`func (o *TenantRequest) GetFlowInvalidationOk() (*string, bool)`

GetFlowInvalidationOk returns a tuple with the FlowInvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInvalidation

`func (o *TenantRequest) SetFlowInvalidation(v string)`

SetFlowInvalidation sets FlowInvalidation field to given value.

### HasFlowInvalidation

`func (o *TenantRequest) HasFlowInvalidation() bool`

HasFlowInvalidation returns a boolean if a field has been set.

### SetFlowInvalidationNil

`func (o *TenantRequest) SetFlowInvalidationNil(b bool)`

 SetFlowInvalidationNil sets the value for FlowInvalidation to be an explicit nil

### UnsetFlowInvalidation
`func (o *TenantRequest) UnsetFlowInvalidation()`

UnsetFlowInvalidation ensures that no value is present for FlowInvalidation, not even an explicit nil
### GetFlowRecovery

`func (o *TenantRequest) GetFlowRecovery() string`

GetFlowRecovery returns the FlowRecovery field if non-nil, zero value otherwise.

### GetFlowRecoveryOk

`func (o *TenantRequest) GetFlowRecoveryOk() (*string, bool)`

GetFlowRecoveryOk returns a tuple with the FlowRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowRecovery

`func (o *TenantRequest) SetFlowRecovery(v string)`

SetFlowRecovery sets FlowRecovery field to given value.

### HasFlowRecovery

`func (o *TenantRequest) HasFlowRecovery() bool`

HasFlowRecovery returns a boolean if a field has been set.

### SetFlowRecoveryNil

`func (o *TenantRequest) SetFlowRecoveryNil(b bool)`

 SetFlowRecoveryNil sets the value for FlowRecovery to be an explicit nil

### UnsetFlowRecovery
`func (o *TenantRequest) UnsetFlowRecovery()`

UnsetFlowRecovery ensures that no value is present for FlowRecovery, not even an explicit nil
### GetFlowUnenrollment

`func (o *TenantRequest) GetFlowUnenrollment() string`

GetFlowUnenrollment returns the FlowUnenrollment field if non-nil, zero value otherwise.

### GetFlowUnenrollmentOk

`func (o *TenantRequest) GetFlowUnenrollmentOk() (*string, bool)`

GetFlowUnenrollmentOk returns a tuple with the FlowUnenrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowUnenrollment

`func (o *TenantRequest) SetFlowUnenrollment(v string)`

SetFlowUnenrollment sets FlowUnenrollment field to given value.

### HasFlowUnenrollment

`func (o *TenantRequest) HasFlowUnenrollment() bool`

HasFlowUnenrollment returns a boolean if a field has been set.

### SetFlowUnenrollmentNil

`func (o *TenantRequest) SetFlowUnenrollmentNil(b bool)`

 SetFlowUnenrollmentNil sets the value for FlowUnenrollment to be an explicit nil

### UnsetFlowUnenrollment
`func (o *TenantRequest) UnsetFlowUnenrollment()`

UnsetFlowUnenrollment ensures that no value is present for FlowUnenrollment, not even an explicit nil
### GetFlowUserSettings

`func (o *TenantRequest) GetFlowUserSettings() string`

GetFlowUserSettings returns the FlowUserSettings field if non-nil, zero value otherwise.

### GetFlowUserSettingsOk

`func (o *TenantRequest) GetFlowUserSettingsOk() (*string, bool)`

GetFlowUserSettingsOk returns a tuple with the FlowUserSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowUserSettings

`func (o *TenantRequest) SetFlowUserSettings(v string)`

SetFlowUserSettings sets FlowUserSettings field to given value.

### HasFlowUserSettings

`func (o *TenantRequest) HasFlowUserSettings() bool`

HasFlowUserSettings returns a boolean if a field has been set.

### SetFlowUserSettingsNil

`func (o *TenantRequest) SetFlowUserSettingsNil(b bool)`

 SetFlowUserSettingsNil sets the value for FlowUserSettings to be an explicit nil

### UnsetFlowUserSettings
`func (o *TenantRequest) UnsetFlowUserSettings()`

UnsetFlowUserSettings ensures that no value is present for FlowUserSettings, not even an explicit nil
### GetFlowDeviceCode

`func (o *TenantRequest) GetFlowDeviceCode() string`

GetFlowDeviceCode returns the FlowDeviceCode field if non-nil, zero value otherwise.

### GetFlowDeviceCodeOk

`func (o *TenantRequest) GetFlowDeviceCodeOk() (*string, bool)`

GetFlowDeviceCodeOk returns a tuple with the FlowDeviceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDeviceCode

`func (o *TenantRequest) SetFlowDeviceCode(v string)`

SetFlowDeviceCode sets FlowDeviceCode field to given value.

### HasFlowDeviceCode

`func (o *TenantRequest) HasFlowDeviceCode() bool`

HasFlowDeviceCode returns a boolean if a field has been set.

### SetFlowDeviceCodeNil

`func (o *TenantRequest) SetFlowDeviceCodeNil(b bool)`

 SetFlowDeviceCodeNil sets the value for FlowDeviceCode to be an explicit nil

### UnsetFlowDeviceCode
`func (o *TenantRequest) UnsetFlowDeviceCode()`

UnsetFlowDeviceCode ensures that no value is present for FlowDeviceCode, not even an explicit nil
### GetEventRetention

`func (o *TenantRequest) GetEventRetention() string`

GetEventRetention returns the EventRetention field if non-nil, zero value otherwise.

### GetEventRetentionOk

`func (o *TenantRequest) GetEventRetentionOk() (*string, bool)`

GetEventRetentionOk returns a tuple with the EventRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRetention

`func (o *TenantRequest) SetEventRetention(v string)`

SetEventRetention sets EventRetention field to given value.

### HasEventRetention

`func (o *TenantRequest) HasEventRetention() bool`

HasEventRetention returns a boolean if a field has been set.

### GetWebCertificate

`func (o *TenantRequest) GetWebCertificate() string`

GetWebCertificate returns the WebCertificate field if non-nil, zero value otherwise.

### GetWebCertificateOk

`func (o *TenantRequest) GetWebCertificateOk() (*string, bool)`

GetWebCertificateOk returns a tuple with the WebCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebCertificate

`func (o *TenantRequest) SetWebCertificate(v string)`

SetWebCertificate sets WebCertificate field to given value.

### HasWebCertificate

`func (o *TenantRequest) HasWebCertificate() bool`

HasWebCertificate returns a boolean if a field has been set.

### SetWebCertificateNil

`func (o *TenantRequest) SetWebCertificateNil(b bool)`

 SetWebCertificateNil sets the value for WebCertificate to be an explicit nil

### UnsetWebCertificate
`func (o *TenantRequest) UnsetWebCertificate()`

UnsetWebCertificate ensures that no value is present for WebCertificate, not even an explicit nil
### GetAttributes

`func (o *TenantRequest) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TenantRequest) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TenantRequest) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TenantRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



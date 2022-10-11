# Tenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantUuid** | **string** |  | [readonly] 
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

### NewTenant

`func NewTenant(tenantUuid string, domain string, ) *Tenant`

NewTenant instantiates a new Tenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantWithDefaults

`func NewTenantWithDefaults() *Tenant`

NewTenantWithDefaults instantiates a new Tenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantUuid

`func (o *Tenant) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *Tenant) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *Tenant) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.


### GetDomain

`func (o *Tenant) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Tenant) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Tenant) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetDefault

`func (o *Tenant) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *Tenant) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *Tenant) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *Tenant) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetBrandingTitle

`func (o *Tenant) GetBrandingTitle() string`

GetBrandingTitle returns the BrandingTitle field if non-nil, zero value otherwise.

### GetBrandingTitleOk

`func (o *Tenant) GetBrandingTitleOk() (*string, bool)`

GetBrandingTitleOk returns a tuple with the BrandingTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingTitle

`func (o *Tenant) SetBrandingTitle(v string)`

SetBrandingTitle sets BrandingTitle field to given value.

### HasBrandingTitle

`func (o *Tenant) HasBrandingTitle() bool`

HasBrandingTitle returns a boolean if a field has been set.

### GetBrandingLogo

`func (o *Tenant) GetBrandingLogo() string`

GetBrandingLogo returns the BrandingLogo field if non-nil, zero value otherwise.

### GetBrandingLogoOk

`func (o *Tenant) GetBrandingLogoOk() (*string, bool)`

GetBrandingLogoOk returns a tuple with the BrandingLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingLogo

`func (o *Tenant) SetBrandingLogo(v string)`

SetBrandingLogo sets BrandingLogo field to given value.

### HasBrandingLogo

`func (o *Tenant) HasBrandingLogo() bool`

HasBrandingLogo returns a boolean if a field has been set.

### GetBrandingFavicon

`func (o *Tenant) GetBrandingFavicon() string`

GetBrandingFavicon returns the BrandingFavicon field if non-nil, zero value otherwise.

### GetBrandingFaviconOk

`func (o *Tenant) GetBrandingFaviconOk() (*string, bool)`

GetBrandingFaviconOk returns a tuple with the BrandingFavicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingFavicon

`func (o *Tenant) SetBrandingFavicon(v string)`

SetBrandingFavicon sets BrandingFavicon field to given value.

### HasBrandingFavicon

`func (o *Tenant) HasBrandingFavicon() bool`

HasBrandingFavicon returns a boolean if a field has been set.

### GetFlowAuthentication

`func (o *Tenant) GetFlowAuthentication() string`

GetFlowAuthentication returns the FlowAuthentication field if non-nil, zero value otherwise.

### GetFlowAuthenticationOk

`func (o *Tenant) GetFlowAuthenticationOk() (*string, bool)`

GetFlowAuthenticationOk returns a tuple with the FlowAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowAuthentication

`func (o *Tenant) SetFlowAuthentication(v string)`

SetFlowAuthentication sets FlowAuthentication field to given value.

### HasFlowAuthentication

`func (o *Tenant) HasFlowAuthentication() bool`

HasFlowAuthentication returns a boolean if a field has been set.

### SetFlowAuthenticationNil

`func (o *Tenant) SetFlowAuthenticationNil(b bool)`

 SetFlowAuthenticationNil sets the value for FlowAuthentication to be an explicit nil

### UnsetFlowAuthentication
`func (o *Tenant) UnsetFlowAuthentication()`

UnsetFlowAuthentication ensures that no value is present for FlowAuthentication, not even an explicit nil
### GetFlowInvalidation

`func (o *Tenant) GetFlowInvalidation() string`

GetFlowInvalidation returns the FlowInvalidation field if non-nil, zero value otherwise.

### GetFlowInvalidationOk

`func (o *Tenant) GetFlowInvalidationOk() (*string, bool)`

GetFlowInvalidationOk returns a tuple with the FlowInvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInvalidation

`func (o *Tenant) SetFlowInvalidation(v string)`

SetFlowInvalidation sets FlowInvalidation field to given value.

### HasFlowInvalidation

`func (o *Tenant) HasFlowInvalidation() bool`

HasFlowInvalidation returns a boolean if a field has been set.

### SetFlowInvalidationNil

`func (o *Tenant) SetFlowInvalidationNil(b bool)`

 SetFlowInvalidationNil sets the value for FlowInvalidation to be an explicit nil

### UnsetFlowInvalidation
`func (o *Tenant) UnsetFlowInvalidation()`

UnsetFlowInvalidation ensures that no value is present for FlowInvalidation, not even an explicit nil
### GetFlowRecovery

`func (o *Tenant) GetFlowRecovery() string`

GetFlowRecovery returns the FlowRecovery field if non-nil, zero value otherwise.

### GetFlowRecoveryOk

`func (o *Tenant) GetFlowRecoveryOk() (*string, bool)`

GetFlowRecoveryOk returns a tuple with the FlowRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowRecovery

`func (o *Tenant) SetFlowRecovery(v string)`

SetFlowRecovery sets FlowRecovery field to given value.

### HasFlowRecovery

`func (o *Tenant) HasFlowRecovery() bool`

HasFlowRecovery returns a boolean if a field has been set.

### SetFlowRecoveryNil

`func (o *Tenant) SetFlowRecoveryNil(b bool)`

 SetFlowRecoveryNil sets the value for FlowRecovery to be an explicit nil

### UnsetFlowRecovery
`func (o *Tenant) UnsetFlowRecovery()`

UnsetFlowRecovery ensures that no value is present for FlowRecovery, not even an explicit nil
### GetFlowUnenrollment

`func (o *Tenant) GetFlowUnenrollment() string`

GetFlowUnenrollment returns the FlowUnenrollment field if non-nil, zero value otherwise.

### GetFlowUnenrollmentOk

`func (o *Tenant) GetFlowUnenrollmentOk() (*string, bool)`

GetFlowUnenrollmentOk returns a tuple with the FlowUnenrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowUnenrollment

`func (o *Tenant) SetFlowUnenrollment(v string)`

SetFlowUnenrollment sets FlowUnenrollment field to given value.

### HasFlowUnenrollment

`func (o *Tenant) HasFlowUnenrollment() bool`

HasFlowUnenrollment returns a boolean if a field has been set.

### SetFlowUnenrollmentNil

`func (o *Tenant) SetFlowUnenrollmentNil(b bool)`

 SetFlowUnenrollmentNil sets the value for FlowUnenrollment to be an explicit nil

### UnsetFlowUnenrollment
`func (o *Tenant) UnsetFlowUnenrollment()`

UnsetFlowUnenrollment ensures that no value is present for FlowUnenrollment, not even an explicit nil
### GetFlowUserSettings

`func (o *Tenant) GetFlowUserSettings() string`

GetFlowUserSettings returns the FlowUserSettings field if non-nil, zero value otherwise.

### GetFlowUserSettingsOk

`func (o *Tenant) GetFlowUserSettingsOk() (*string, bool)`

GetFlowUserSettingsOk returns a tuple with the FlowUserSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowUserSettings

`func (o *Tenant) SetFlowUserSettings(v string)`

SetFlowUserSettings sets FlowUserSettings field to given value.

### HasFlowUserSettings

`func (o *Tenant) HasFlowUserSettings() bool`

HasFlowUserSettings returns a boolean if a field has been set.

### SetFlowUserSettingsNil

`func (o *Tenant) SetFlowUserSettingsNil(b bool)`

 SetFlowUserSettingsNil sets the value for FlowUserSettings to be an explicit nil

### UnsetFlowUserSettings
`func (o *Tenant) UnsetFlowUserSettings()`

UnsetFlowUserSettings ensures that no value is present for FlowUserSettings, not even an explicit nil
### GetFlowDeviceCode

`func (o *Tenant) GetFlowDeviceCode() string`

GetFlowDeviceCode returns the FlowDeviceCode field if non-nil, zero value otherwise.

### GetFlowDeviceCodeOk

`func (o *Tenant) GetFlowDeviceCodeOk() (*string, bool)`

GetFlowDeviceCodeOk returns a tuple with the FlowDeviceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDeviceCode

`func (o *Tenant) SetFlowDeviceCode(v string)`

SetFlowDeviceCode sets FlowDeviceCode field to given value.

### HasFlowDeviceCode

`func (o *Tenant) HasFlowDeviceCode() bool`

HasFlowDeviceCode returns a boolean if a field has been set.

### SetFlowDeviceCodeNil

`func (o *Tenant) SetFlowDeviceCodeNil(b bool)`

 SetFlowDeviceCodeNil sets the value for FlowDeviceCode to be an explicit nil

### UnsetFlowDeviceCode
`func (o *Tenant) UnsetFlowDeviceCode()`

UnsetFlowDeviceCode ensures that no value is present for FlowDeviceCode, not even an explicit nil
### GetEventRetention

`func (o *Tenant) GetEventRetention() string`

GetEventRetention returns the EventRetention field if non-nil, zero value otherwise.

### GetEventRetentionOk

`func (o *Tenant) GetEventRetentionOk() (*string, bool)`

GetEventRetentionOk returns a tuple with the EventRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRetention

`func (o *Tenant) SetEventRetention(v string)`

SetEventRetention sets EventRetention field to given value.

### HasEventRetention

`func (o *Tenant) HasEventRetention() bool`

HasEventRetention returns a boolean if a field has been set.

### GetWebCertificate

`func (o *Tenant) GetWebCertificate() string`

GetWebCertificate returns the WebCertificate field if non-nil, zero value otherwise.

### GetWebCertificateOk

`func (o *Tenant) GetWebCertificateOk() (*string, bool)`

GetWebCertificateOk returns a tuple with the WebCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebCertificate

`func (o *Tenant) SetWebCertificate(v string)`

SetWebCertificate sets WebCertificate field to given value.

### HasWebCertificate

`func (o *Tenant) HasWebCertificate() bool`

HasWebCertificate returns a boolean if a field has been set.

### SetWebCertificateNil

`func (o *Tenant) SetWebCertificateNil(b bool)`

 SetWebCertificateNil sets the value for WebCertificate to be an explicit nil

### UnsetWebCertificate
`func (o *Tenant) UnsetWebCertificate()`

UnsetWebCertificate ensures that no value is present for WebCertificate, not even an explicit nil
### GetAttributes

`func (o *Tenant) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Tenant) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Tenant) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Tenant) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



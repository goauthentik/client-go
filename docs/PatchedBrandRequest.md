# PatchedBrandRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | Domain that activates this brand. Can be a superset, i.e. &#x60;a.b&#x60; for &#x60;aa.b&#x60; and &#x60;ba.b&#x60; | [optional] 
**Default** | Pointer to **bool** |  | [optional] 
**BrandingTitle** | Pointer to **string** |  | [optional] 
**BrandingLogo** | Pointer to **string** |  | [optional] 
**BrandingFavicon** | Pointer to **string** |  | [optional] 
**BrandingCustomCss** | Pointer to **string** |  | [optional] 
**BrandingDefaultFlowBackground** | Pointer to **string** |  | [optional] 
**FlowAuthentication** | Pointer to **NullableString** |  | [optional] 
**FlowInvalidation** | Pointer to **NullableString** |  | [optional] 
**FlowRecovery** | Pointer to **NullableString** |  | [optional] 
**FlowUnenrollment** | Pointer to **NullableString** |  | [optional] 
**FlowUserSettings** | Pointer to **NullableString** |  | [optional] 
**FlowDeviceCode** | Pointer to **NullableString** |  | [optional] 
**DefaultApplication** | Pointer to **NullableString** | When set, external users will be redirected to this application after authenticating. | [optional] 
**WebCertificate** | Pointer to **NullableString** | Web Certificate used by the authentik Core webserver. | [optional] 
**ClientCertificates** | Pointer to **[]string** | Certificates used for client authentication. | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedBrandRequest

`func NewPatchedBrandRequest() *PatchedBrandRequest`

NewPatchedBrandRequest instantiates a new PatchedBrandRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBrandRequestWithDefaults

`func NewPatchedBrandRequestWithDefaults() *PatchedBrandRequest`

NewPatchedBrandRequestWithDefaults instantiates a new PatchedBrandRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *PatchedBrandRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PatchedBrandRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PatchedBrandRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *PatchedBrandRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDefault

`func (o *PatchedBrandRequest) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *PatchedBrandRequest) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *PatchedBrandRequest) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *PatchedBrandRequest) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetBrandingTitle

`func (o *PatchedBrandRequest) GetBrandingTitle() string`

GetBrandingTitle returns the BrandingTitle field if non-nil, zero value otherwise.

### GetBrandingTitleOk

`func (o *PatchedBrandRequest) GetBrandingTitleOk() (*string, bool)`

GetBrandingTitleOk returns a tuple with the BrandingTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingTitle

`func (o *PatchedBrandRequest) SetBrandingTitle(v string)`

SetBrandingTitle sets BrandingTitle field to given value.

### HasBrandingTitle

`func (o *PatchedBrandRequest) HasBrandingTitle() bool`

HasBrandingTitle returns a boolean if a field has been set.

### GetBrandingLogo

`func (o *PatchedBrandRequest) GetBrandingLogo() string`

GetBrandingLogo returns the BrandingLogo field if non-nil, zero value otherwise.

### GetBrandingLogoOk

`func (o *PatchedBrandRequest) GetBrandingLogoOk() (*string, bool)`

GetBrandingLogoOk returns a tuple with the BrandingLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingLogo

`func (o *PatchedBrandRequest) SetBrandingLogo(v string)`

SetBrandingLogo sets BrandingLogo field to given value.

### HasBrandingLogo

`func (o *PatchedBrandRequest) HasBrandingLogo() bool`

HasBrandingLogo returns a boolean if a field has been set.

### GetBrandingFavicon

`func (o *PatchedBrandRequest) GetBrandingFavicon() string`

GetBrandingFavicon returns the BrandingFavicon field if non-nil, zero value otherwise.

### GetBrandingFaviconOk

`func (o *PatchedBrandRequest) GetBrandingFaviconOk() (*string, bool)`

GetBrandingFaviconOk returns a tuple with the BrandingFavicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingFavicon

`func (o *PatchedBrandRequest) SetBrandingFavicon(v string)`

SetBrandingFavicon sets BrandingFavicon field to given value.

### HasBrandingFavicon

`func (o *PatchedBrandRequest) HasBrandingFavicon() bool`

HasBrandingFavicon returns a boolean if a field has been set.

### GetBrandingCustomCss

`func (o *PatchedBrandRequest) GetBrandingCustomCss() string`

GetBrandingCustomCss returns the BrandingCustomCss field if non-nil, zero value otherwise.

### GetBrandingCustomCssOk

`func (o *PatchedBrandRequest) GetBrandingCustomCssOk() (*string, bool)`

GetBrandingCustomCssOk returns a tuple with the BrandingCustomCss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingCustomCss

`func (o *PatchedBrandRequest) SetBrandingCustomCss(v string)`

SetBrandingCustomCss sets BrandingCustomCss field to given value.

### HasBrandingCustomCss

`func (o *PatchedBrandRequest) HasBrandingCustomCss() bool`

HasBrandingCustomCss returns a boolean if a field has been set.

### GetBrandingDefaultFlowBackground

`func (o *PatchedBrandRequest) GetBrandingDefaultFlowBackground() string`

GetBrandingDefaultFlowBackground returns the BrandingDefaultFlowBackground field if non-nil, zero value otherwise.

### GetBrandingDefaultFlowBackgroundOk

`func (o *PatchedBrandRequest) GetBrandingDefaultFlowBackgroundOk() (*string, bool)`

GetBrandingDefaultFlowBackgroundOk returns a tuple with the BrandingDefaultFlowBackground field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingDefaultFlowBackground

`func (o *PatchedBrandRequest) SetBrandingDefaultFlowBackground(v string)`

SetBrandingDefaultFlowBackground sets BrandingDefaultFlowBackground field to given value.

### HasBrandingDefaultFlowBackground

`func (o *PatchedBrandRequest) HasBrandingDefaultFlowBackground() bool`

HasBrandingDefaultFlowBackground returns a boolean if a field has been set.

### GetFlowAuthentication

`func (o *PatchedBrandRequest) GetFlowAuthentication() string`

GetFlowAuthentication returns the FlowAuthentication field if non-nil, zero value otherwise.

### GetFlowAuthenticationOk

`func (o *PatchedBrandRequest) GetFlowAuthenticationOk() (*string, bool)`

GetFlowAuthenticationOk returns a tuple with the FlowAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowAuthentication

`func (o *PatchedBrandRequest) SetFlowAuthentication(v string)`

SetFlowAuthentication sets FlowAuthentication field to given value.

### HasFlowAuthentication

`func (o *PatchedBrandRequest) HasFlowAuthentication() bool`

HasFlowAuthentication returns a boolean if a field has been set.

### SetFlowAuthenticationNil

`func (o *PatchedBrandRequest) SetFlowAuthenticationNil(b bool)`

 SetFlowAuthenticationNil sets the value for FlowAuthentication to be an explicit nil

### UnsetFlowAuthentication
`func (o *PatchedBrandRequest) UnsetFlowAuthentication()`

UnsetFlowAuthentication ensures that no value is present for FlowAuthentication, not even an explicit nil
### GetFlowInvalidation

`func (o *PatchedBrandRequest) GetFlowInvalidation() string`

GetFlowInvalidation returns the FlowInvalidation field if non-nil, zero value otherwise.

### GetFlowInvalidationOk

`func (o *PatchedBrandRequest) GetFlowInvalidationOk() (*string, bool)`

GetFlowInvalidationOk returns a tuple with the FlowInvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInvalidation

`func (o *PatchedBrandRequest) SetFlowInvalidation(v string)`

SetFlowInvalidation sets FlowInvalidation field to given value.

### HasFlowInvalidation

`func (o *PatchedBrandRequest) HasFlowInvalidation() bool`

HasFlowInvalidation returns a boolean if a field has been set.

### SetFlowInvalidationNil

`func (o *PatchedBrandRequest) SetFlowInvalidationNil(b bool)`

 SetFlowInvalidationNil sets the value for FlowInvalidation to be an explicit nil

### UnsetFlowInvalidation
`func (o *PatchedBrandRequest) UnsetFlowInvalidation()`

UnsetFlowInvalidation ensures that no value is present for FlowInvalidation, not even an explicit nil
### GetFlowRecovery

`func (o *PatchedBrandRequest) GetFlowRecovery() string`

GetFlowRecovery returns the FlowRecovery field if non-nil, zero value otherwise.

### GetFlowRecoveryOk

`func (o *PatchedBrandRequest) GetFlowRecoveryOk() (*string, bool)`

GetFlowRecoveryOk returns a tuple with the FlowRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowRecovery

`func (o *PatchedBrandRequest) SetFlowRecovery(v string)`

SetFlowRecovery sets FlowRecovery field to given value.

### HasFlowRecovery

`func (o *PatchedBrandRequest) HasFlowRecovery() bool`

HasFlowRecovery returns a boolean if a field has been set.

### SetFlowRecoveryNil

`func (o *PatchedBrandRequest) SetFlowRecoveryNil(b bool)`

 SetFlowRecoveryNil sets the value for FlowRecovery to be an explicit nil

### UnsetFlowRecovery
`func (o *PatchedBrandRequest) UnsetFlowRecovery()`

UnsetFlowRecovery ensures that no value is present for FlowRecovery, not even an explicit nil
### GetFlowUnenrollment

`func (o *PatchedBrandRequest) GetFlowUnenrollment() string`

GetFlowUnenrollment returns the FlowUnenrollment field if non-nil, zero value otherwise.

### GetFlowUnenrollmentOk

`func (o *PatchedBrandRequest) GetFlowUnenrollmentOk() (*string, bool)`

GetFlowUnenrollmentOk returns a tuple with the FlowUnenrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowUnenrollment

`func (o *PatchedBrandRequest) SetFlowUnenrollment(v string)`

SetFlowUnenrollment sets FlowUnenrollment field to given value.

### HasFlowUnenrollment

`func (o *PatchedBrandRequest) HasFlowUnenrollment() bool`

HasFlowUnenrollment returns a boolean if a field has been set.

### SetFlowUnenrollmentNil

`func (o *PatchedBrandRequest) SetFlowUnenrollmentNil(b bool)`

 SetFlowUnenrollmentNil sets the value for FlowUnenrollment to be an explicit nil

### UnsetFlowUnenrollment
`func (o *PatchedBrandRequest) UnsetFlowUnenrollment()`

UnsetFlowUnenrollment ensures that no value is present for FlowUnenrollment, not even an explicit nil
### GetFlowUserSettings

`func (o *PatchedBrandRequest) GetFlowUserSettings() string`

GetFlowUserSettings returns the FlowUserSettings field if non-nil, zero value otherwise.

### GetFlowUserSettingsOk

`func (o *PatchedBrandRequest) GetFlowUserSettingsOk() (*string, bool)`

GetFlowUserSettingsOk returns a tuple with the FlowUserSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowUserSettings

`func (o *PatchedBrandRequest) SetFlowUserSettings(v string)`

SetFlowUserSettings sets FlowUserSettings field to given value.

### HasFlowUserSettings

`func (o *PatchedBrandRequest) HasFlowUserSettings() bool`

HasFlowUserSettings returns a boolean if a field has been set.

### SetFlowUserSettingsNil

`func (o *PatchedBrandRequest) SetFlowUserSettingsNil(b bool)`

 SetFlowUserSettingsNil sets the value for FlowUserSettings to be an explicit nil

### UnsetFlowUserSettings
`func (o *PatchedBrandRequest) UnsetFlowUserSettings()`

UnsetFlowUserSettings ensures that no value is present for FlowUserSettings, not even an explicit nil
### GetFlowDeviceCode

`func (o *PatchedBrandRequest) GetFlowDeviceCode() string`

GetFlowDeviceCode returns the FlowDeviceCode field if non-nil, zero value otherwise.

### GetFlowDeviceCodeOk

`func (o *PatchedBrandRequest) GetFlowDeviceCodeOk() (*string, bool)`

GetFlowDeviceCodeOk returns a tuple with the FlowDeviceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDeviceCode

`func (o *PatchedBrandRequest) SetFlowDeviceCode(v string)`

SetFlowDeviceCode sets FlowDeviceCode field to given value.

### HasFlowDeviceCode

`func (o *PatchedBrandRequest) HasFlowDeviceCode() bool`

HasFlowDeviceCode returns a boolean if a field has been set.

### SetFlowDeviceCodeNil

`func (o *PatchedBrandRequest) SetFlowDeviceCodeNil(b bool)`

 SetFlowDeviceCodeNil sets the value for FlowDeviceCode to be an explicit nil

### UnsetFlowDeviceCode
`func (o *PatchedBrandRequest) UnsetFlowDeviceCode()`

UnsetFlowDeviceCode ensures that no value is present for FlowDeviceCode, not even an explicit nil
### GetDefaultApplication

`func (o *PatchedBrandRequest) GetDefaultApplication() string`

GetDefaultApplication returns the DefaultApplication field if non-nil, zero value otherwise.

### GetDefaultApplicationOk

`func (o *PatchedBrandRequest) GetDefaultApplicationOk() (*string, bool)`

GetDefaultApplicationOk returns a tuple with the DefaultApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultApplication

`func (o *PatchedBrandRequest) SetDefaultApplication(v string)`

SetDefaultApplication sets DefaultApplication field to given value.

### HasDefaultApplication

`func (o *PatchedBrandRequest) HasDefaultApplication() bool`

HasDefaultApplication returns a boolean if a field has been set.

### SetDefaultApplicationNil

`func (o *PatchedBrandRequest) SetDefaultApplicationNil(b bool)`

 SetDefaultApplicationNil sets the value for DefaultApplication to be an explicit nil

### UnsetDefaultApplication
`func (o *PatchedBrandRequest) UnsetDefaultApplication()`

UnsetDefaultApplication ensures that no value is present for DefaultApplication, not even an explicit nil
### GetWebCertificate

`func (o *PatchedBrandRequest) GetWebCertificate() string`

GetWebCertificate returns the WebCertificate field if non-nil, zero value otherwise.

### GetWebCertificateOk

`func (o *PatchedBrandRequest) GetWebCertificateOk() (*string, bool)`

GetWebCertificateOk returns a tuple with the WebCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebCertificate

`func (o *PatchedBrandRequest) SetWebCertificate(v string)`

SetWebCertificate sets WebCertificate field to given value.

### HasWebCertificate

`func (o *PatchedBrandRequest) HasWebCertificate() bool`

HasWebCertificate returns a boolean if a field has been set.

### SetWebCertificateNil

`func (o *PatchedBrandRequest) SetWebCertificateNil(b bool)`

 SetWebCertificateNil sets the value for WebCertificate to be an explicit nil

### UnsetWebCertificate
`func (o *PatchedBrandRequest) UnsetWebCertificate()`

UnsetWebCertificate ensures that no value is present for WebCertificate, not even an explicit nil
### GetClientCertificates

`func (o *PatchedBrandRequest) GetClientCertificates() []string`

GetClientCertificates returns the ClientCertificates field if non-nil, zero value otherwise.

### GetClientCertificatesOk

`func (o *PatchedBrandRequest) GetClientCertificatesOk() (*[]string, bool)`

GetClientCertificatesOk returns a tuple with the ClientCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificates

`func (o *PatchedBrandRequest) SetClientCertificates(v []string)`

SetClientCertificates sets ClientCertificates field to given value.

### HasClientCertificates

`func (o *PatchedBrandRequest) HasClientCertificates() bool`

HasClientCertificates returns a boolean if a field has been set.

### GetAttributes

`func (o *PatchedBrandRequest) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PatchedBrandRequest) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PatchedBrandRequest) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PatchedBrandRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



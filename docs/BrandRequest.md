# BrandRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | Domain that activates this brand. Can be a superset, i.e. &#x60;a.b&#x60; for &#x60;aa.b&#x60; and &#x60;ba.b&#x60; | 
**Default** | Pointer to **bool** |  | [optional] 
**BrandingTitle** | Pointer to **string** |  | [optional] 
**BrandingLogo** | Pointer to **string** |  | [optional] 
**BrandingFavicon** | Pointer to **string** |  | [optional] 
**BrandingCustomCss** | Pointer to **string** |  | [optional] 
**BrandingDefaultFlowBackground** | Pointer to **string** |  | [optional] 
**FlowAuthentication** | Pointer to **NullableString** |  | [optional] 
**FlowUserSwitch** | Pointer to **NullableString** |  | [optional] 
**FlowInvalidation** | Pointer to **NullableString** |  | [optional] 
**FlowRecovery** | Pointer to **NullableString** |  | [optional] 
**FlowUnenrollment** | Pointer to **NullableString** |  | [optional] 
**FlowUserSettings** | Pointer to **NullableString** |  | [optional] 
**FlowDeviceCode** | Pointer to **NullableString** |  | [optional] 
**FlowLockdown** | Pointer to **NullableString** |  | [optional] 
**FlowRequest** | Pointer to **NullableString** |  | [optional] 
**DefaultApplication** | Pointer to **NullableString** | When set, external users will be redirected to this application after authenticating. | [optional] 
**WebCertificate** | Pointer to **NullableString** | Web Certificate used by the authentik Core webserver. | [optional] 
**ClientCertificates** | Pointer to **[]string** | Certificates used for client authentication. | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBrandRequest

`func NewBrandRequest(domain string, ) *BrandRequest`

NewBrandRequest instantiates a new BrandRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandRequestWithDefaults

`func NewBrandRequestWithDefaults() *BrandRequest`

NewBrandRequestWithDefaults instantiates a new BrandRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *BrandRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *BrandRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *BrandRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetDefault

`func (o *BrandRequest) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *BrandRequest) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *BrandRequest) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *BrandRequest) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetBrandingTitle

`func (o *BrandRequest) GetBrandingTitle() string`

GetBrandingTitle returns the BrandingTitle field if non-nil, zero value otherwise.

### GetBrandingTitleOk

`func (o *BrandRequest) GetBrandingTitleOk() (*string, bool)`

GetBrandingTitleOk returns a tuple with the BrandingTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingTitle

`func (o *BrandRequest) SetBrandingTitle(v string)`

SetBrandingTitle sets BrandingTitle field to given value.

### HasBrandingTitle

`func (o *BrandRequest) HasBrandingTitle() bool`

HasBrandingTitle returns a boolean if a field has been set.

### GetBrandingLogo

`func (o *BrandRequest) GetBrandingLogo() string`

GetBrandingLogo returns the BrandingLogo field if non-nil, zero value otherwise.

### GetBrandingLogoOk

`func (o *BrandRequest) GetBrandingLogoOk() (*string, bool)`

GetBrandingLogoOk returns a tuple with the BrandingLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingLogo

`func (o *BrandRequest) SetBrandingLogo(v string)`

SetBrandingLogo sets BrandingLogo field to given value.

### HasBrandingLogo

`func (o *BrandRequest) HasBrandingLogo() bool`

HasBrandingLogo returns a boolean if a field has been set.

### GetBrandingFavicon

`func (o *BrandRequest) GetBrandingFavicon() string`

GetBrandingFavicon returns the BrandingFavicon field if non-nil, zero value otherwise.

### GetBrandingFaviconOk

`func (o *BrandRequest) GetBrandingFaviconOk() (*string, bool)`

GetBrandingFaviconOk returns a tuple with the BrandingFavicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingFavicon

`func (o *BrandRequest) SetBrandingFavicon(v string)`

SetBrandingFavicon sets BrandingFavicon field to given value.

### HasBrandingFavicon

`func (o *BrandRequest) HasBrandingFavicon() bool`

HasBrandingFavicon returns a boolean if a field has been set.

### GetBrandingCustomCss

`func (o *BrandRequest) GetBrandingCustomCss() string`

GetBrandingCustomCss returns the BrandingCustomCss field if non-nil, zero value otherwise.

### GetBrandingCustomCssOk

`func (o *BrandRequest) GetBrandingCustomCssOk() (*string, bool)`

GetBrandingCustomCssOk returns a tuple with the BrandingCustomCss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingCustomCss

`func (o *BrandRequest) SetBrandingCustomCss(v string)`

SetBrandingCustomCss sets BrandingCustomCss field to given value.

### HasBrandingCustomCss

`func (o *BrandRequest) HasBrandingCustomCss() bool`

HasBrandingCustomCss returns a boolean if a field has been set.

### GetBrandingDefaultFlowBackground

`func (o *BrandRequest) GetBrandingDefaultFlowBackground() string`

GetBrandingDefaultFlowBackground returns the BrandingDefaultFlowBackground field if non-nil, zero value otherwise.

### GetBrandingDefaultFlowBackgroundOk

`func (o *BrandRequest) GetBrandingDefaultFlowBackgroundOk() (*string, bool)`

GetBrandingDefaultFlowBackgroundOk returns a tuple with the BrandingDefaultFlowBackground field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingDefaultFlowBackground

`func (o *BrandRequest) SetBrandingDefaultFlowBackground(v string)`

SetBrandingDefaultFlowBackground sets BrandingDefaultFlowBackground field to given value.

### HasBrandingDefaultFlowBackground

`func (o *BrandRequest) HasBrandingDefaultFlowBackground() bool`

HasBrandingDefaultFlowBackground returns a boolean if a field has been set.

### GetFlowAuthentication

`func (o *BrandRequest) GetFlowAuthentication() string`

GetFlowAuthentication returns the FlowAuthentication field if non-nil, zero value otherwise.

### GetFlowAuthenticationOk

`func (o *BrandRequest) GetFlowAuthenticationOk() (*string, bool)`

GetFlowAuthenticationOk returns a tuple with the FlowAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowAuthentication

`func (o *BrandRequest) SetFlowAuthentication(v string)`

SetFlowAuthentication sets FlowAuthentication field to given value.

### HasFlowAuthentication

`func (o *BrandRequest) HasFlowAuthentication() bool`

HasFlowAuthentication returns a boolean if a field has been set.

### SetFlowAuthenticationNil

`func (o *BrandRequest) SetFlowAuthenticationNil(b bool)`

 SetFlowAuthenticationNil sets the value for FlowAuthentication to be an explicit nil

### UnsetFlowAuthentication
`func (o *BrandRequest) UnsetFlowAuthentication()`

UnsetFlowAuthentication ensures that no value is present for FlowAuthentication, not even an explicit nil
### GetFlowUserSwitch

`func (o *BrandRequest) GetFlowUserSwitch() string`

GetFlowUserSwitch returns the FlowUserSwitch field if non-nil, zero value otherwise.

### GetFlowUserSwitchOk

`func (o *BrandRequest) GetFlowUserSwitchOk() (*string, bool)`

GetFlowUserSwitchOk returns a tuple with the FlowUserSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowUserSwitch

`func (o *BrandRequest) SetFlowUserSwitch(v string)`

SetFlowUserSwitch sets FlowUserSwitch field to given value.

### HasFlowUserSwitch

`func (o *BrandRequest) HasFlowUserSwitch() bool`

HasFlowUserSwitch returns a boolean if a field has been set.

### SetFlowUserSwitchNil

`func (o *BrandRequest) SetFlowUserSwitchNil(b bool)`

 SetFlowUserSwitchNil sets the value for FlowUserSwitch to be an explicit nil

### UnsetFlowUserSwitch
`func (o *BrandRequest) UnsetFlowUserSwitch()`

UnsetFlowUserSwitch ensures that no value is present for FlowUserSwitch, not even an explicit nil
### GetFlowInvalidation

`func (o *BrandRequest) GetFlowInvalidation() string`

GetFlowInvalidation returns the FlowInvalidation field if non-nil, zero value otherwise.

### GetFlowInvalidationOk

`func (o *BrandRequest) GetFlowInvalidationOk() (*string, bool)`

GetFlowInvalidationOk returns a tuple with the FlowInvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInvalidation

`func (o *BrandRequest) SetFlowInvalidation(v string)`

SetFlowInvalidation sets FlowInvalidation field to given value.

### HasFlowInvalidation

`func (o *BrandRequest) HasFlowInvalidation() bool`

HasFlowInvalidation returns a boolean if a field has been set.

### SetFlowInvalidationNil

`func (o *BrandRequest) SetFlowInvalidationNil(b bool)`

 SetFlowInvalidationNil sets the value for FlowInvalidation to be an explicit nil

### UnsetFlowInvalidation
`func (o *BrandRequest) UnsetFlowInvalidation()`

UnsetFlowInvalidation ensures that no value is present for FlowInvalidation, not even an explicit nil
### GetFlowRecovery

`func (o *BrandRequest) GetFlowRecovery() string`

GetFlowRecovery returns the FlowRecovery field if non-nil, zero value otherwise.

### GetFlowRecoveryOk

`func (o *BrandRequest) GetFlowRecoveryOk() (*string, bool)`

GetFlowRecoveryOk returns a tuple with the FlowRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowRecovery

`func (o *BrandRequest) SetFlowRecovery(v string)`

SetFlowRecovery sets FlowRecovery field to given value.

### HasFlowRecovery

`func (o *BrandRequest) HasFlowRecovery() bool`

HasFlowRecovery returns a boolean if a field has been set.

### SetFlowRecoveryNil

`func (o *BrandRequest) SetFlowRecoveryNil(b bool)`

 SetFlowRecoveryNil sets the value for FlowRecovery to be an explicit nil

### UnsetFlowRecovery
`func (o *BrandRequest) UnsetFlowRecovery()`

UnsetFlowRecovery ensures that no value is present for FlowRecovery, not even an explicit nil
### GetFlowUnenrollment

`func (o *BrandRequest) GetFlowUnenrollment() string`

GetFlowUnenrollment returns the FlowUnenrollment field if non-nil, zero value otherwise.

### GetFlowUnenrollmentOk

`func (o *BrandRequest) GetFlowUnenrollmentOk() (*string, bool)`

GetFlowUnenrollmentOk returns a tuple with the FlowUnenrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowUnenrollment

`func (o *BrandRequest) SetFlowUnenrollment(v string)`

SetFlowUnenrollment sets FlowUnenrollment field to given value.

### HasFlowUnenrollment

`func (o *BrandRequest) HasFlowUnenrollment() bool`

HasFlowUnenrollment returns a boolean if a field has been set.

### SetFlowUnenrollmentNil

`func (o *BrandRequest) SetFlowUnenrollmentNil(b bool)`

 SetFlowUnenrollmentNil sets the value for FlowUnenrollment to be an explicit nil

### UnsetFlowUnenrollment
`func (o *BrandRequest) UnsetFlowUnenrollment()`

UnsetFlowUnenrollment ensures that no value is present for FlowUnenrollment, not even an explicit nil
### GetFlowUserSettings

`func (o *BrandRequest) GetFlowUserSettings() string`

GetFlowUserSettings returns the FlowUserSettings field if non-nil, zero value otherwise.

### GetFlowUserSettingsOk

`func (o *BrandRequest) GetFlowUserSettingsOk() (*string, bool)`

GetFlowUserSettingsOk returns a tuple with the FlowUserSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowUserSettings

`func (o *BrandRequest) SetFlowUserSettings(v string)`

SetFlowUserSettings sets FlowUserSettings field to given value.

### HasFlowUserSettings

`func (o *BrandRequest) HasFlowUserSettings() bool`

HasFlowUserSettings returns a boolean if a field has been set.

### SetFlowUserSettingsNil

`func (o *BrandRequest) SetFlowUserSettingsNil(b bool)`

 SetFlowUserSettingsNil sets the value for FlowUserSettings to be an explicit nil

### UnsetFlowUserSettings
`func (o *BrandRequest) UnsetFlowUserSettings()`

UnsetFlowUserSettings ensures that no value is present for FlowUserSettings, not even an explicit nil
### GetFlowDeviceCode

`func (o *BrandRequest) GetFlowDeviceCode() string`

GetFlowDeviceCode returns the FlowDeviceCode field if non-nil, zero value otherwise.

### GetFlowDeviceCodeOk

`func (o *BrandRequest) GetFlowDeviceCodeOk() (*string, bool)`

GetFlowDeviceCodeOk returns a tuple with the FlowDeviceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDeviceCode

`func (o *BrandRequest) SetFlowDeviceCode(v string)`

SetFlowDeviceCode sets FlowDeviceCode field to given value.

### HasFlowDeviceCode

`func (o *BrandRequest) HasFlowDeviceCode() bool`

HasFlowDeviceCode returns a boolean if a field has been set.

### SetFlowDeviceCodeNil

`func (o *BrandRequest) SetFlowDeviceCodeNil(b bool)`

 SetFlowDeviceCodeNil sets the value for FlowDeviceCode to be an explicit nil

### UnsetFlowDeviceCode
`func (o *BrandRequest) UnsetFlowDeviceCode()`

UnsetFlowDeviceCode ensures that no value is present for FlowDeviceCode, not even an explicit nil
### GetFlowLockdown

`func (o *BrandRequest) GetFlowLockdown() string`

GetFlowLockdown returns the FlowLockdown field if non-nil, zero value otherwise.

### GetFlowLockdownOk

`func (o *BrandRequest) GetFlowLockdownOk() (*string, bool)`

GetFlowLockdownOk returns a tuple with the FlowLockdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowLockdown

`func (o *BrandRequest) SetFlowLockdown(v string)`

SetFlowLockdown sets FlowLockdown field to given value.

### HasFlowLockdown

`func (o *BrandRequest) HasFlowLockdown() bool`

HasFlowLockdown returns a boolean if a field has been set.

### SetFlowLockdownNil

`func (o *BrandRequest) SetFlowLockdownNil(b bool)`

 SetFlowLockdownNil sets the value for FlowLockdown to be an explicit nil

### UnsetFlowLockdown
`func (o *BrandRequest) UnsetFlowLockdown()`

UnsetFlowLockdown ensures that no value is present for FlowLockdown, not even an explicit nil
### GetFlowRequest

`func (o *BrandRequest) GetFlowRequest() string`

GetFlowRequest returns the FlowRequest field if non-nil, zero value otherwise.

### GetFlowRequestOk

`func (o *BrandRequest) GetFlowRequestOk() (*string, bool)`

GetFlowRequestOk returns a tuple with the FlowRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowRequest

`func (o *BrandRequest) SetFlowRequest(v string)`

SetFlowRequest sets FlowRequest field to given value.

### HasFlowRequest

`func (o *BrandRequest) HasFlowRequest() bool`

HasFlowRequest returns a boolean if a field has been set.

### SetFlowRequestNil

`func (o *BrandRequest) SetFlowRequestNil(b bool)`

 SetFlowRequestNil sets the value for FlowRequest to be an explicit nil

### UnsetFlowRequest
`func (o *BrandRequest) UnsetFlowRequest()`

UnsetFlowRequest ensures that no value is present for FlowRequest, not even an explicit nil
### GetDefaultApplication

`func (o *BrandRequest) GetDefaultApplication() string`

GetDefaultApplication returns the DefaultApplication field if non-nil, zero value otherwise.

### GetDefaultApplicationOk

`func (o *BrandRequest) GetDefaultApplicationOk() (*string, bool)`

GetDefaultApplicationOk returns a tuple with the DefaultApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultApplication

`func (o *BrandRequest) SetDefaultApplication(v string)`

SetDefaultApplication sets DefaultApplication field to given value.

### HasDefaultApplication

`func (o *BrandRequest) HasDefaultApplication() bool`

HasDefaultApplication returns a boolean if a field has been set.

### SetDefaultApplicationNil

`func (o *BrandRequest) SetDefaultApplicationNil(b bool)`

 SetDefaultApplicationNil sets the value for DefaultApplication to be an explicit nil

### UnsetDefaultApplication
`func (o *BrandRequest) UnsetDefaultApplication()`

UnsetDefaultApplication ensures that no value is present for DefaultApplication, not even an explicit nil
### GetWebCertificate

`func (o *BrandRequest) GetWebCertificate() string`

GetWebCertificate returns the WebCertificate field if non-nil, zero value otherwise.

### GetWebCertificateOk

`func (o *BrandRequest) GetWebCertificateOk() (*string, bool)`

GetWebCertificateOk returns a tuple with the WebCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebCertificate

`func (o *BrandRequest) SetWebCertificate(v string)`

SetWebCertificate sets WebCertificate field to given value.

### HasWebCertificate

`func (o *BrandRequest) HasWebCertificate() bool`

HasWebCertificate returns a boolean if a field has been set.

### SetWebCertificateNil

`func (o *BrandRequest) SetWebCertificateNil(b bool)`

 SetWebCertificateNil sets the value for WebCertificate to be an explicit nil

### UnsetWebCertificate
`func (o *BrandRequest) UnsetWebCertificate()`

UnsetWebCertificate ensures that no value is present for WebCertificate, not even an explicit nil
### GetClientCertificates

`func (o *BrandRequest) GetClientCertificates() []string`

GetClientCertificates returns the ClientCertificates field if non-nil, zero value otherwise.

### GetClientCertificatesOk

`func (o *BrandRequest) GetClientCertificatesOk() (*[]string, bool)`

GetClientCertificatesOk returns a tuple with the ClientCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificates

`func (o *BrandRequest) SetClientCertificates(v []string)`

SetClientCertificates sets ClientCertificates field to given value.

### HasClientCertificates

`func (o *BrandRequest) HasClientCertificates() bool`

HasClientCertificates returns a boolean if a field has been set.

### GetAttributes

`func (o *BrandRequest) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BrandRequest) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BrandRequest) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *BrandRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



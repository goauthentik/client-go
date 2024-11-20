# CurrentBrand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchedDomain** | **string** |  | 
**BrandingTitle** | **string** |  | 
**BrandingLogo** | **string** |  | 
**BrandingFavicon** | **string** |  | 
**UiFooterLinks** | [**[]FooterLink**](FooterLink.md) |  | [readonly] 
**UiTheme** | [**UiThemeEnum**](UiThemeEnum.md) |  | [readonly] [default to UITHEMEENUM_AUTOMATIC]
**FlowAuthentication** | Pointer to **string** |  | [optional] 
**FlowInvalidation** | Pointer to **string** |  | [optional] 
**FlowRecovery** | Pointer to **string** |  | [optional] 
**FlowUnenrollment** | Pointer to **string** |  | [optional] 
**FlowUserSettings** | Pointer to **string** |  | [optional] 
**FlowDeviceCode** | Pointer to **string** |  | [optional] 
**DefaultLocale** | **string** |  | [readonly] 

## Methods

### NewCurrentBrand

`func NewCurrentBrand(matchedDomain string, brandingTitle string, brandingLogo string, brandingFavicon string, uiFooterLinks []FooterLink, uiTheme UiThemeEnum, defaultLocale string, ) *CurrentBrand`

NewCurrentBrand instantiates a new CurrentBrand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentBrandWithDefaults

`func NewCurrentBrandWithDefaults() *CurrentBrand`

NewCurrentBrandWithDefaults instantiates a new CurrentBrand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchedDomain

`func (o *CurrentBrand) GetMatchedDomain() string`

GetMatchedDomain returns the MatchedDomain field if non-nil, zero value otherwise.

### GetMatchedDomainOk

`func (o *CurrentBrand) GetMatchedDomainOk() (*string, bool)`

GetMatchedDomainOk returns a tuple with the MatchedDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedDomain

`func (o *CurrentBrand) SetMatchedDomain(v string)`

SetMatchedDomain sets MatchedDomain field to given value.


### GetBrandingTitle

`func (o *CurrentBrand) GetBrandingTitle() string`

GetBrandingTitle returns the BrandingTitle field if non-nil, zero value otherwise.

### GetBrandingTitleOk

`func (o *CurrentBrand) GetBrandingTitleOk() (*string, bool)`

GetBrandingTitleOk returns a tuple with the BrandingTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingTitle

`func (o *CurrentBrand) SetBrandingTitle(v string)`

SetBrandingTitle sets BrandingTitle field to given value.


### GetBrandingLogo

`func (o *CurrentBrand) GetBrandingLogo() string`

GetBrandingLogo returns the BrandingLogo field if non-nil, zero value otherwise.

### GetBrandingLogoOk

`func (o *CurrentBrand) GetBrandingLogoOk() (*string, bool)`

GetBrandingLogoOk returns a tuple with the BrandingLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingLogo

`func (o *CurrentBrand) SetBrandingLogo(v string)`

SetBrandingLogo sets BrandingLogo field to given value.


### GetBrandingFavicon

`func (o *CurrentBrand) GetBrandingFavicon() string`

GetBrandingFavicon returns the BrandingFavicon field if non-nil, zero value otherwise.

### GetBrandingFaviconOk

`func (o *CurrentBrand) GetBrandingFaviconOk() (*string, bool)`

GetBrandingFaviconOk returns a tuple with the BrandingFavicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingFavicon

`func (o *CurrentBrand) SetBrandingFavicon(v string)`

SetBrandingFavicon sets BrandingFavicon field to given value.


### GetUiFooterLinks

`func (o *CurrentBrand) GetUiFooterLinks() []FooterLink`

GetUiFooterLinks returns the UiFooterLinks field if non-nil, zero value otherwise.

### GetUiFooterLinksOk

`func (o *CurrentBrand) GetUiFooterLinksOk() (*[]FooterLink, bool)`

GetUiFooterLinksOk returns a tuple with the UiFooterLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiFooterLinks

`func (o *CurrentBrand) SetUiFooterLinks(v []FooterLink)`

SetUiFooterLinks sets UiFooterLinks field to given value.


### GetUiTheme

`func (o *CurrentBrand) GetUiTheme() UiThemeEnum`

GetUiTheme returns the UiTheme field if non-nil, zero value otherwise.

### GetUiThemeOk

`func (o *CurrentBrand) GetUiThemeOk() (*UiThemeEnum, bool)`

GetUiThemeOk returns a tuple with the UiTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiTheme

`func (o *CurrentBrand) SetUiTheme(v UiThemeEnum)`

SetUiTheme sets UiTheme field to given value.


### GetFlowAuthentication

`func (o *CurrentBrand) GetFlowAuthentication() string`

GetFlowAuthentication returns the FlowAuthentication field if non-nil, zero value otherwise.

### GetFlowAuthenticationOk

`func (o *CurrentBrand) GetFlowAuthenticationOk() (*string, bool)`

GetFlowAuthenticationOk returns a tuple with the FlowAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowAuthentication

`func (o *CurrentBrand) SetFlowAuthentication(v string)`

SetFlowAuthentication sets FlowAuthentication field to given value.

### HasFlowAuthentication

`func (o *CurrentBrand) HasFlowAuthentication() bool`

HasFlowAuthentication returns a boolean if a field has been set.

### GetFlowInvalidation

`func (o *CurrentBrand) GetFlowInvalidation() string`

GetFlowInvalidation returns the FlowInvalidation field if non-nil, zero value otherwise.

### GetFlowInvalidationOk

`func (o *CurrentBrand) GetFlowInvalidationOk() (*string, bool)`

GetFlowInvalidationOk returns a tuple with the FlowInvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInvalidation

`func (o *CurrentBrand) SetFlowInvalidation(v string)`

SetFlowInvalidation sets FlowInvalidation field to given value.

### HasFlowInvalidation

`func (o *CurrentBrand) HasFlowInvalidation() bool`

HasFlowInvalidation returns a boolean if a field has been set.

### GetFlowRecovery

`func (o *CurrentBrand) GetFlowRecovery() string`

GetFlowRecovery returns the FlowRecovery field if non-nil, zero value otherwise.

### GetFlowRecoveryOk

`func (o *CurrentBrand) GetFlowRecoveryOk() (*string, bool)`

GetFlowRecoveryOk returns a tuple with the FlowRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowRecovery

`func (o *CurrentBrand) SetFlowRecovery(v string)`

SetFlowRecovery sets FlowRecovery field to given value.

### HasFlowRecovery

`func (o *CurrentBrand) HasFlowRecovery() bool`

HasFlowRecovery returns a boolean if a field has been set.

### GetFlowUnenrollment

`func (o *CurrentBrand) GetFlowUnenrollment() string`

GetFlowUnenrollment returns the FlowUnenrollment field if non-nil, zero value otherwise.

### GetFlowUnenrollmentOk

`func (o *CurrentBrand) GetFlowUnenrollmentOk() (*string, bool)`

GetFlowUnenrollmentOk returns a tuple with the FlowUnenrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowUnenrollment

`func (o *CurrentBrand) SetFlowUnenrollment(v string)`

SetFlowUnenrollment sets FlowUnenrollment field to given value.

### HasFlowUnenrollment

`func (o *CurrentBrand) HasFlowUnenrollment() bool`

HasFlowUnenrollment returns a boolean if a field has been set.

### GetFlowUserSettings

`func (o *CurrentBrand) GetFlowUserSettings() string`

GetFlowUserSettings returns the FlowUserSettings field if non-nil, zero value otherwise.

### GetFlowUserSettingsOk

`func (o *CurrentBrand) GetFlowUserSettingsOk() (*string, bool)`

GetFlowUserSettingsOk returns a tuple with the FlowUserSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowUserSettings

`func (o *CurrentBrand) SetFlowUserSettings(v string)`

SetFlowUserSettings sets FlowUserSettings field to given value.

### HasFlowUserSettings

`func (o *CurrentBrand) HasFlowUserSettings() bool`

HasFlowUserSettings returns a boolean if a field has been set.

### GetFlowDeviceCode

`func (o *CurrentBrand) GetFlowDeviceCode() string`

GetFlowDeviceCode returns the FlowDeviceCode field if non-nil, zero value otherwise.

### GetFlowDeviceCodeOk

`func (o *CurrentBrand) GetFlowDeviceCodeOk() (*string, bool)`

GetFlowDeviceCodeOk returns a tuple with the FlowDeviceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDeviceCode

`func (o *CurrentBrand) SetFlowDeviceCode(v string)`

SetFlowDeviceCode sets FlowDeviceCode field to given value.

### HasFlowDeviceCode

`func (o *CurrentBrand) HasFlowDeviceCode() bool`

HasFlowDeviceCode returns a boolean if a field has been set.

### GetDefaultLocale

`func (o *CurrentBrand) GetDefaultLocale() string`

GetDefaultLocale returns the DefaultLocale field if non-nil, zero value otherwise.

### GetDefaultLocaleOk

`func (o *CurrentBrand) GetDefaultLocaleOk() (*string, bool)`

GetDefaultLocaleOk returns a tuple with the DefaultLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLocale

`func (o *CurrentBrand) SetDefaultLocale(v string)`

SetDefaultLocale sets DefaultLocale field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



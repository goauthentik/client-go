# PatchedTenantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | Domain that activates this tenant. Can be a superset, i.e. &#x60;a.b&#x60; for &#x60;aa.b&#x60; and &#x60;ba.b&#x60; | [optional] 
**Default** | Pointer to **bool** |  | [optional] 
**BrandingTitle** | Pointer to **string** |  | [optional] 
**BrandingLogo** | Pointer to **string** |  | [optional] 
**BrandingFavicon** | Pointer to **string** |  | [optional] 
**FlowAuthentication** | Pointer to **NullableString** |  | [optional] 
**FlowInvalidation** | Pointer to **NullableString** |  | [optional] 
**FlowRecovery** | Pointer to **NullableString** |  | [optional] 
**FlowUnenrollment** | Pointer to **NullableString** |  | [optional] 
**EventRetention** | Pointer to **string** | Events will be deleted after this duration.(Format: weeks&#x3D;3;days&#x3D;2;hours&#x3D;3,seconds&#x3D;2). | [optional] 

## Methods

### NewPatchedTenantRequest

`func NewPatchedTenantRequest() *PatchedTenantRequest`

NewPatchedTenantRequest instantiates a new PatchedTenantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedTenantRequestWithDefaults

`func NewPatchedTenantRequestWithDefaults() *PatchedTenantRequest`

NewPatchedTenantRequestWithDefaults instantiates a new PatchedTenantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *PatchedTenantRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PatchedTenantRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PatchedTenantRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *PatchedTenantRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDefault

`func (o *PatchedTenantRequest) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *PatchedTenantRequest) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *PatchedTenantRequest) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *PatchedTenantRequest) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetBrandingTitle

`func (o *PatchedTenantRequest) GetBrandingTitle() string`

GetBrandingTitle returns the BrandingTitle field if non-nil, zero value otherwise.

### GetBrandingTitleOk

`func (o *PatchedTenantRequest) GetBrandingTitleOk() (*string, bool)`

GetBrandingTitleOk returns a tuple with the BrandingTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingTitle

`func (o *PatchedTenantRequest) SetBrandingTitle(v string)`

SetBrandingTitle sets BrandingTitle field to given value.

### HasBrandingTitle

`func (o *PatchedTenantRequest) HasBrandingTitle() bool`

HasBrandingTitle returns a boolean if a field has been set.

### GetBrandingLogo

`func (o *PatchedTenantRequest) GetBrandingLogo() string`

GetBrandingLogo returns the BrandingLogo field if non-nil, zero value otherwise.

### GetBrandingLogoOk

`func (o *PatchedTenantRequest) GetBrandingLogoOk() (*string, bool)`

GetBrandingLogoOk returns a tuple with the BrandingLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingLogo

`func (o *PatchedTenantRequest) SetBrandingLogo(v string)`

SetBrandingLogo sets BrandingLogo field to given value.

### HasBrandingLogo

`func (o *PatchedTenantRequest) HasBrandingLogo() bool`

HasBrandingLogo returns a boolean if a field has been set.

### GetBrandingFavicon

`func (o *PatchedTenantRequest) GetBrandingFavicon() string`

GetBrandingFavicon returns the BrandingFavicon field if non-nil, zero value otherwise.

### GetBrandingFaviconOk

`func (o *PatchedTenantRequest) GetBrandingFaviconOk() (*string, bool)`

GetBrandingFaviconOk returns a tuple with the BrandingFavicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingFavicon

`func (o *PatchedTenantRequest) SetBrandingFavicon(v string)`

SetBrandingFavicon sets BrandingFavicon field to given value.

### HasBrandingFavicon

`func (o *PatchedTenantRequest) HasBrandingFavicon() bool`

HasBrandingFavicon returns a boolean if a field has been set.

### GetFlowAuthentication

`func (o *PatchedTenantRequest) GetFlowAuthentication() string`

GetFlowAuthentication returns the FlowAuthentication field if non-nil, zero value otherwise.

### GetFlowAuthenticationOk

`func (o *PatchedTenantRequest) GetFlowAuthenticationOk() (*string, bool)`

GetFlowAuthenticationOk returns a tuple with the FlowAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowAuthentication

`func (o *PatchedTenantRequest) SetFlowAuthentication(v string)`

SetFlowAuthentication sets FlowAuthentication field to given value.

### HasFlowAuthentication

`func (o *PatchedTenantRequest) HasFlowAuthentication() bool`

HasFlowAuthentication returns a boolean if a field has been set.

### SetFlowAuthenticationNil

`func (o *PatchedTenantRequest) SetFlowAuthenticationNil(b bool)`

 SetFlowAuthenticationNil sets the value for FlowAuthentication to be an explicit nil

### UnsetFlowAuthentication
`func (o *PatchedTenantRequest) UnsetFlowAuthentication()`

UnsetFlowAuthentication ensures that no value is present for FlowAuthentication, not even an explicit nil
### GetFlowInvalidation

`func (o *PatchedTenantRequest) GetFlowInvalidation() string`

GetFlowInvalidation returns the FlowInvalidation field if non-nil, zero value otherwise.

### GetFlowInvalidationOk

`func (o *PatchedTenantRequest) GetFlowInvalidationOk() (*string, bool)`

GetFlowInvalidationOk returns a tuple with the FlowInvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInvalidation

`func (o *PatchedTenantRequest) SetFlowInvalidation(v string)`

SetFlowInvalidation sets FlowInvalidation field to given value.

### HasFlowInvalidation

`func (o *PatchedTenantRequest) HasFlowInvalidation() bool`

HasFlowInvalidation returns a boolean if a field has been set.

### SetFlowInvalidationNil

`func (o *PatchedTenantRequest) SetFlowInvalidationNil(b bool)`

 SetFlowInvalidationNil sets the value for FlowInvalidation to be an explicit nil

### UnsetFlowInvalidation
`func (o *PatchedTenantRequest) UnsetFlowInvalidation()`

UnsetFlowInvalidation ensures that no value is present for FlowInvalidation, not even an explicit nil
### GetFlowRecovery

`func (o *PatchedTenantRequest) GetFlowRecovery() string`

GetFlowRecovery returns the FlowRecovery field if non-nil, zero value otherwise.

### GetFlowRecoveryOk

`func (o *PatchedTenantRequest) GetFlowRecoveryOk() (*string, bool)`

GetFlowRecoveryOk returns a tuple with the FlowRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowRecovery

`func (o *PatchedTenantRequest) SetFlowRecovery(v string)`

SetFlowRecovery sets FlowRecovery field to given value.

### HasFlowRecovery

`func (o *PatchedTenantRequest) HasFlowRecovery() bool`

HasFlowRecovery returns a boolean if a field has been set.

### SetFlowRecoveryNil

`func (o *PatchedTenantRequest) SetFlowRecoveryNil(b bool)`

 SetFlowRecoveryNil sets the value for FlowRecovery to be an explicit nil

### UnsetFlowRecovery
`func (o *PatchedTenantRequest) UnsetFlowRecovery()`

UnsetFlowRecovery ensures that no value is present for FlowRecovery, not even an explicit nil
### GetFlowUnenrollment

`func (o *PatchedTenantRequest) GetFlowUnenrollment() string`

GetFlowUnenrollment returns the FlowUnenrollment field if non-nil, zero value otherwise.

### GetFlowUnenrollmentOk

`func (o *PatchedTenantRequest) GetFlowUnenrollmentOk() (*string, bool)`

GetFlowUnenrollmentOk returns a tuple with the FlowUnenrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowUnenrollment

`func (o *PatchedTenantRequest) SetFlowUnenrollment(v string)`

SetFlowUnenrollment sets FlowUnenrollment field to given value.

### HasFlowUnenrollment

`func (o *PatchedTenantRequest) HasFlowUnenrollment() bool`

HasFlowUnenrollment returns a boolean if a field has been set.

### SetFlowUnenrollmentNil

`func (o *PatchedTenantRequest) SetFlowUnenrollmentNil(b bool)`

 SetFlowUnenrollmentNil sets the value for FlowUnenrollment to be an explicit nil

### UnsetFlowUnenrollment
`func (o *PatchedTenantRequest) UnsetFlowUnenrollment()`

UnsetFlowUnenrollment ensures that no value is present for FlowUnenrollment, not even an explicit nil
### GetEventRetention

`func (o *PatchedTenantRequest) GetEventRetention() string`

GetEventRetention returns the EventRetention field if non-nil, zero value otherwise.

### GetEventRetentionOk

`func (o *PatchedTenantRequest) GetEventRetentionOk() (*string, bool)`

GetEventRetentionOk returns a tuple with the EventRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRetention

`func (o *PatchedTenantRequest) SetEventRetention(v string)`

SetEventRetention sets EventRetention field to given value.

### HasEventRetention

`func (o *PatchedTenantRequest) HasEventRetention() bool`

HasEventRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



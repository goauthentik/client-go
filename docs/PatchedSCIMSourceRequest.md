# PatchedSCIMSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Source&#39;s display Name. | [optional] 
**Slug** | Pointer to **string** | Internal source name, used in URLs. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**UserPropertyMappings** | Pointer to **[]string** |  | [optional] 
**GroupPropertyMappings** | Pointer to **[]string** |  | [optional] 
**UserPathTemplate** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedSCIMSourceRequest

`func NewPatchedSCIMSourceRequest() *PatchedSCIMSourceRequest`

NewPatchedSCIMSourceRequest instantiates a new PatchedSCIMSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedSCIMSourceRequestWithDefaults

`func NewPatchedSCIMSourceRequestWithDefaults() *PatchedSCIMSourceRequest`

NewPatchedSCIMSourceRequestWithDefaults instantiates a new PatchedSCIMSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedSCIMSourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedSCIMSourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedSCIMSourceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedSCIMSourceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedSCIMSourceRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedSCIMSourceRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedSCIMSourceRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedSCIMSourceRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedSCIMSourceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedSCIMSourceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedSCIMSourceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedSCIMSourceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUserPropertyMappings

`func (o *PatchedSCIMSourceRequest) GetUserPropertyMappings() []string`

GetUserPropertyMappings returns the UserPropertyMappings field if non-nil, zero value otherwise.

### GetUserPropertyMappingsOk

`func (o *PatchedSCIMSourceRequest) GetUserPropertyMappingsOk() (*[]string, bool)`

GetUserPropertyMappingsOk returns a tuple with the UserPropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPropertyMappings

`func (o *PatchedSCIMSourceRequest) SetUserPropertyMappings(v []string)`

SetUserPropertyMappings sets UserPropertyMappings field to given value.

### HasUserPropertyMappings

`func (o *PatchedSCIMSourceRequest) HasUserPropertyMappings() bool`

HasUserPropertyMappings returns a boolean if a field has been set.

### GetGroupPropertyMappings

`func (o *PatchedSCIMSourceRequest) GetGroupPropertyMappings() []string`

GetGroupPropertyMappings returns the GroupPropertyMappings field if non-nil, zero value otherwise.

### GetGroupPropertyMappingsOk

`func (o *PatchedSCIMSourceRequest) GetGroupPropertyMappingsOk() (*[]string, bool)`

GetGroupPropertyMappingsOk returns a tuple with the GroupPropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPropertyMappings

`func (o *PatchedSCIMSourceRequest) SetGroupPropertyMappings(v []string)`

SetGroupPropertyMappings sets GroupPropertyMappings field to given value.

### HasGroupPropertyMappings

`func (o *PatchedSCIMSourceRequest) HasGroupPropertyMappings() bool`

HasGroupPropertyMappings returns a boolean if a field has been set.

### GetUserPathTemplate

`func (o *PatchedSCIMSourceRequest) GetUserPathTemplate() string`

GetUserPathTemplate returns the UserPathTemplate field if non-nil, zero value otherwise.

### GetUserPathTemplateOk

`func (o *PatchedSCIMSourceRequest) GetUserPathTemplateOk() (*string, bool)`

GetUserPathTemplateOk returns a tuple with the UserPathTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPathTemplate

`func (o *PatchedSCIMSourceRequest) SetUserPathTemplate(v string)`

SetUserPathTemplate sets UserPathTemplate field to given value.

### HasUserPathTemplate

`func (o *PatchedSCIMSourceRequest) HasUserPathTemplate() bool`

HasUserPathTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



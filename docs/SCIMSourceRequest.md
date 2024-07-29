# SCIMSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Source&#39;s display Name. | 
**Slug** | **string** | Internal source name, used in URLs. | 
**Enabled** | Pointer to **bool** |  | [optional] 
**UserPropertyMappings** | Pointer to **[]string** |  | [optional] 
**GroupPropertyMappings** | Pointer to **[]string** |  | [optional] 
**UserPathTemplate** | Pointer to **string** |  | [optional] 

## Methods

### NewSCIMSourceRequest

`func NewSCIMSourceRequest(name string, slug string, ) *SCIMSourceRequest`

NewSCIMSourceRequest instantiates a new SCIMSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSCIMSourceRequestWithDefaults

`func NewSCIMSourceRequestWithDefaults() *SCIMSourceRequest`

NewSCIMSourceRequestWithDefaults instantiates a new SCIMSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SCIMSourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SCIMSourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SCIMSourceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *SCIMSourceRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *SCIMSourceRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *SCIMSourceRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetEnabled

`func (o *SCIMSourceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SCIMSourceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SCIMSourceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SCIMSourceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUserPropertyMappings

`func (o *SCIMSourceRequest) GetUserPropertyMappings() []string`

GetUserPropertyMappings returns the UserPropertyMappings field if non-nil, zero value otherwise.

### GetUserPropertyMappingsOk

`func (o *SCIMSourceRequest) GetUserPropertyMappingsOk() (*[]string, bool)`

GetUserPropertyMappingsOk returns a tuple with the UserPropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPropertyMappings

`func (o *SCIMSourceRequest) SetUserPropertyMappings(v []string)`

SetUserPropertyMappings sets UserPropertyMappings field to given value.

### HasUserPropertyMappings

`func (o *SCIMSourceRequest) HasUserPropertyMappings() bool`

HasUserPropertyMappings returns a boolean if a field has been set.

### GetGroupPropertyMappings

`func (o *SCIMSourceRequest) GetGroupPropertyMappings() []string`

GetGroupPropertyMappings returns the GroupPropertyMappings field if non-nil, zero value otherwise.

### GetGroupPropertyMappingsOk

`func (o *SCIMSourceRequest) GetGroupPropertyMappingsOk() (*[]string, bool)`

GetGroupPropertyMappingsOk returns a tuple with the GroupPropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPropertyMappings

`func (o *SCIMSourceRequest) SetGroupPropertyMappings(v []string)`

SetGroupPropertyMappings sets GroupPropertyMappings field to given value.

### HasGroupPropertyMappings

`func (o *SCIMSourceRequest) HasGroupPropertyMappings() bool`

HasGroupPropertyMappings returns a boolean if a field has been set.

### GetUserPathTemplate

`func (o *SCIMSourceRequest) GetUserPathTemplate() string`

GetUserPathTemplate returns the UserPathTemplate field if non-nil, zero value otherwise.

### GetUserPathTemplateOk

`func (o *SCIMSourceRequest) GetUserPathTemplateOk() (*string, bool)`

GetUserPathTemplateOk returns a tuple with the UserPathTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPathTemplate

`func (o *SCIMSourceRequest) SetUserPathTemplate(v string)`

SetUserPathTemplate sets UserPathTemplate field to given value.

### HasUserPathTemplate

`func (o *SCIMSourceRequest) HasUserPathTemplate() bool`

HasUserPathTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



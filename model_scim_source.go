/*
authentik

Making authentication simple.

API version: 2025.6.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SCIMSource SCIMSource Serializer
type SCIMSource struct {
	Pk string `json:"pk"`
	// Source's display Name.
	Name string `json:"name"`
	// Internal source name, used in URLs.
	Slug                  string   `json:"slug"`
	Enabled               *bool    `json:"enabled,omitempty"`
	UserPropertyMappings  []string `json:"user_property_mappings,omitempty"`
	GroupPropertyMappings []string `json:"group_property_mappings,omitempty"`
	// Get object component so that we know how to edit the object
	Component string `json:"component"`
	// Return object's verbose_name
	VerboseName string `json:"verbose_name"`
	// Return object's plural verbose_name
	VerboseNamePlural string `json:"verbose_name_plural"`
	// Return internal model name
	MetaModelName string `json:"meta_model_name"`
	// Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update.
	Managed          NullableString `json:"managed"`
	UserPathTemplate *string        `json:"user_path_template,omitempty"`
	// Get Root URL
	RootUrl  string `json:"root_url"`
	TokenObj Token  `json:"token_obj"`
}

// NewSCIMSource instantiates a new SCIMSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSCIMSource(pk string, name string, slug string, component string, verboseName string, verboseNamePlural string, metaModelName string, managed NullableString, rootUrl string, tokenObj Token) *SCIMSource {
	this := SCIMSource{}
	this.Pk = pk
	this.Name = name
	this.Slug = slug
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.Managed = managed
	this.RootUrl = rootUrl
	this.TokenObj = tokenObj
	return &this
}

// NewSCIMSourceWithDefaults instantiates a new SCIMSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSCIMSourceWithDefaults() *SCIMSource {
	this := SCIMSource{}
	return &this
}

// GetPk returns the Pk field value
func (o *SCIMSource) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *SCIMSource) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *SCIMSource) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *SCIMSource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SCIMSource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SCIMSource) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *SCIMSource) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *SCIMSource) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *SCIMSource) SetSlug(v string) {
	o.Slug = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SCIMSource) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCIMSource) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SCIMSource) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SCIMSource) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetUserPropertyMappings returns the UserPropertyMappings field value if set, zero value otherwise.
func (o *SCIMSource) GetUserPropertyMappings() []string {
	if o == nil || o.UserPropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.UserPropertyMappings
}

// GetUserPropertyMappingsOk returns a tuple with the UserPropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCIMSource) GetUserPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.UserPropertyMappings == nil {
		return nil, false
	}
	return o.UserPropertyMappings, true
}

// HasUserPropertyMappings returns a boolean if a field has been set.
func (o *SCIMSource) HasUserPropertyMappings() bool {
	if o != nil && o.UserPropertyMappings != nil {
		return true
	}

	return false
}

// SetUserPropertyMappings gets a reference to the given []string and assigns it to the UserPropertyMappings field.
func (o *SCIMSource) SetUserPropertyMappings(v []string) {
	o.UserPropertyMappings = v
}

// GetGroupPropertyMappings returns the GroupPropertyMappings field value if set, zero value otherwise.
func (o *SCIMSource) GetGroupPropertyMappings() []string {
	if o == nil || o.GroupPropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.GroupPropertyMappings
}

// GetGroupPropertyMappingsOk returns a tuple with the GroupPropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCIMSource) GetGroupPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.GroupPropertyMappings == nil {
		return nil, false
	}
	return o.GroupPropertyMappings, true
}

// HasGroupPropertyMappings returns a boolean if a field has been set.
func (o *SCIMSource) HasGroupPropertyMappings() bool {
	if o != nil && o.GroupPropertyMappings != nil {
		return true
	}

	return false
}

// SetGroupPropertyMappings gets a reference to the given []string and assigns it to the GroupPropertyMappings field.
func (o *SCIMSource) SetGroupPropertyMappings(v []string) {
	o.GroupPropertyMappings = v
}

// GetComponent returns the Component field value
func (o *SCIMSource) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *SCIMSource) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *SCIMSource) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *SCIMSource) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *SCIMSource) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *SCIMSource) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *SCIMSource) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *SCIMSource) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *SCIMSource) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *SCIMSource) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *SCIMSource) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *SCIMSource) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetManaged returns the Managed field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SCIMSource) GetManaged() string {
	if o == nil || o.Managed.Get() == nil {
		var ret string
		return ret
	}

	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SCIMSource) GetManagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// SetManaged sets field value
func (o *SCIMSource) SetManaged(v string) {
	o.Managed.Set(&v)
}

// GetUserPathTemplate returns the UserPathTemplate field value if set, zero value otherwise.
func (o *SCIMSource) GetUserPathTemplate() string {
	if o == nil || o.UserPathTemplate == nil {
		var ret string
		return ret
	}
	return *o.UserPathTemplate
}

// GetUserPathTemplateOk returns a tuple with the UserPathTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCIMSource) GetUserPathTemplateOk() (*string, bool) {
	if o == nil || o.UserPathTemplate == nil {
		return nil, false
	}
	return o.UserPathTemplate, true
}

// HasUserPathTemplate returns a boolean if a field has been set.
func (o *SCIMSource) HasUserPathTemplate() bool {
	if o != nil && o.UserPathTemplate != nil {
		return true
	}

	return false
}

// SetUserPathTemplate gets a reference to the given string and assigns it to the UserPathTemplate field.
func (o *SCIMSource) SetUserPathTemplate(v string) {
	o.UserPathTemplate = &v
}

// GetRootUrl returns the RootUrl field value
func (o *SCIMSource) GetRootUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RootUrl
}

// GetRootUrlOk returns a tuple with the RootUrl field value
// and a boolean to check if the value has been set.
func (o *SCIMSource) GetRootUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootUrl, true
}

// SetRootUrl sets field value
func (o *SCIMSource) SetRootUrl(v string) {
	o.RootUrl = v
}

// GetTokenObj returns the TokenObj field value
func (o *SCIMSource) GetTokenObj() Token {
	if o == nil {
		var ret Token
		return ret
	}

	return o.TokenObj
}

// GetTokenObjOk returns a tuple with the TokenObj field value
// and a boolean to check if the value has been set.
func (o *SCIMSource) GetTokenObjOk() (*Token, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenObj, true
}

// SetTokenObj sets field value
func (o *SCIMSource) SetTokenObj(v Token) {
	o.TokenObj = v
}

func (o SCIMSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["slug"] = o.Slug
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.UserPropertyMappings != nil {
		toSerialize["user_property_mappings"] = o.UserPropertyMappings
	}
	if o.GroupPropertyMappings != nil {
		toSerialize["group_property_mappings"] = o.GroupPropertyMappings
	}
	if true {
		toSerialize["component"] = o.Component
	}
	if true {
		toSerialize["verbose_name"] = o.VerboseName
	}
	if true {
		toSerialize["verbose_name_plural"] = o.VerboseNamePlural
	}
	if true {
		toSerialize["meta_model_name"] = o.MetaModelName
	}
	if true {
		toSerialize["managed"] = o.Managed.Get()
	}
	if o.UserPathTemplate != nil {
		toSerialize["user_path_template"] = o.UserPathTemplate
	}
	if true {
		toSerialize["root_url"] = o.RootUrl
	}
	if true {
		toSerialize["token_obj"] = o.TokenObj
	}
	return json.Marshal(toSerialize)
}

type NullableSCIMSource struct {
	value *SCIMSource
	isSet bool
}

func (v NullableSCIMSource) Get() *SCIMSource {
	return v.value
}

func (v *NullableSCIMSource) Set(val *SCIMSource) {
	v.value = val
	v.isSet = true
}

func (v NullableSCIMSource) IsSet() bool {
	return v.isSet
}

func (v *NullableSCIMSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSCIMSource(val *SCIMSource) *NullableSCIMSource {
	return &NullableSCIMSource{value: val, isSet: true}
}

func (v NullableSCIMSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSCIMSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

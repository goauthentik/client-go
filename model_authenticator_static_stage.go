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

// AuthenticatorStaticStage AuthenticatorStaticStage Serializer
type AuthenticatorStaticStage struct {
	Pk   string `json:"pk"`
	Name string `json:"name"`
	// Get object type so that we know how to edit the object
	Component string `json:"component"`
	// Return object's verbose_name
	VerboseName string `json:"verbose_name"`
	// Return object's plural verbose_name
	VerboseNamePlural string `json:"verbose_name_plural"`
	// Return internal model name
	MetaModelName string    `json:"meta_model_name"`
	FlowSet       []FlowSet `json:"flow_set,omitempty"`
	// Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage.
	ConfigureFlow NullableString `json:"configure_flow,omitempty"`
	FriendlyName  NullableString `json:"friendly_name,omitempty"`
	TokenCount    *int32         `json:"token_count,omitempty"`
	TokenLength   *int32         `json:"token_length,omitempty"`
}

// NewAuthenticatorStaticStage instantiates a new AuthenticatorStaticStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorStaticStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string) *AuthenticatorStaticStage {
	this := AuthenticatorStaticStage{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	return &this
}

// NewAuthenticatorStaticStageWithDefaults instantiates a new AuthenticatorStaticStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorStaticStageWithDefaults() *AuthenticatorStaticStage {
	this := AuthenticatorStaticStage{}
	return &this
}

// GetPk returns the Pk field value
func (o *AuthenticatorStaticStage) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorStaticStage) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *AuthenticatorStaticStage) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *AuthenticatorStaticStage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorStaticStage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthenticatorStaticStage) SetName(v string) {
	o.Name = v
}

// GetComponent returns the Component field value
func (o *AuthenticatorStaticStage) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorStaticStage) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *AuthenticatorStaticStage) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *AuthenticatorStaticStage) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorStaticStage) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *AuthenticatorStaticStage) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *AuthenticatorStaticStage) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorStaticStage) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *AuthenticatorStaticStage) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *AuthenticatorStaticStage) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorStaticStage) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *AuthenticatorStaticStage) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *AuthenticatorStaticStage) GetFlowSet() []FlowSet {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSet
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorStaticStage) GetFlowSetOk() ([]FlowSet, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *AuthenticatorStaticStage) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSet and assigns it to the FlowSet field.
func (o *AuthenticatorStaticStage) SetFlowSet(v []FlowSet) {
	o.FlowSet = v
}

// GetConfigureFlow returns the ConfigureFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticatorStaticStage) GetConfigureFlow() string {
	if o == nil || o.ConfigureFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfigureFlow.Get()
}

// GetConfigureFlowOk returns a tuple with the ConfigureFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticatorStaticStage) GetConfigureFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigureFlow.Get(), o.ConfigureFlow.IsSet()
}

// HasConfigureFlow returns a boolean if a field has been set.
func (o *AuthenticatorStaticStage) HasConfigureFlow() bool {
	if o != nil && o.ConfigureFlow.IsSet() {
		return true
	}

	return false
}

// SetConfigureFlow gets a reference to the given NullableString and assigns it to the ConfigureFlow field.
func (o *AuthenticatorStaticStage) SetConfigureFlow(v string) {
	o.ConfigureFlow.Set(&v)
}

// SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil
func (o *AuthenticatorStaticStage) SetConfigureFlowNil() {
	o.ConfigureFlow.Set(nil)
}

// UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
func (o *AuthenticatorStaticStage) UnsetConfigureFlow() {
	o.ConfigureFlow.Unset()
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticatorStaticStage) GetFriendlyName() string {
	if o == nil || o.FriendlyName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FriendlyName.Get()
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticatorStaticStage) GetFriendlyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FriendlyName.Get(), o.FriendlyName.IsSet()
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *AuthenticatorStaticStage) HasFriendlyName() bool {
	if o != nil && o.FriendlyName.IsSet() {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given NullableString and assigns it to the FriendlyName field.
func (o *AuthenticatorStaticStage) SetFriendlyName(v string) {
	o.FriendlyName.Set(&v)
}

// SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil
func (o *AuthenticatorStaticStage) SetFriendlyNameNil() {
	o.FriendlyName.Set(nil)
}

// UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
func (o *AuthenticatorStaticStage) UnsetFriendlyName() {
	o.FriendlyName.Unset()
}

// GetTokenCount returns the TokenCount field value if set, zero value otherwise.
func (o *AuthenticatorStaticStage) GetTokenCount() int32 {
	if o == nil || o.TokenCount == nil {
		var ret int32
		return ret
	}
	return *o.TokenCount
}

// GetTokenCountOk returns a tuple with the TokenCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorStaticStage) GetTokenCountOk() (*int32, bool) {
	if o == nil || o.TokenCount == nil {
		return nil, false
	}
	return o.TokenCount, true
}

// HasTokenCount returns a boolean if a field has been set.
func (o *AuthenticatorStaticStage) HasTokenCount() bool {
	if o != nil && o.TokenCount != nil {
		return true
	}

	return false
}

// SetTokenCount gets a reference to the given int32 and assigns it to the TokenCount field.
func (o *AuthenticatorStaticStage) SetTokenCount(v int32) {
	o.TokenCount = &v
}

// GetTokenLength returns the TokenLength field value if set, zero value otherwise.
func (o *AuthenticatorStaticStage) GetTokenLength() int32 {
	if o == nil || o.TokenLength == nil {
		var ret int32
		return ret
	}
	return *o.TokenLength
}

// GetTokenLengthOk returns a tuple with the TokenLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorStaticStage) GetTokenLengthOk() (*int32, bool) {
	if o == nil || o.TokenLength == nil {
		return nil, false
	}
	return o.TokenLength, true
}

// HasTokenLength returns a boolean if a field has been set.
func (o *AuthenticatorStaticStage) HasTokenLength() bool {
	if o != nil && o.TokenLength != nil {
		return true
	}

	return false
}

// SetTokenLength gets a reference to the given int32 and assigns it to the TokenLength field.
func (o *AuthenticatorStaticStage) SetTokenLength(v int32) {
	o.TokenLength = &v
}

func (o AuthenticatorStaticStage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
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
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if o.ConfigureFlow.IsSet() {
		toSerialize["configure_flow"] = o.ConfigureFlow.Get()
	}
	if o.FriendlyName.IsSet() {
		toSerialize["friendly_name"] = o.FriendlyName.Get()
	}
	if o.TokenCount != nil {
		toSerialize["token_count"] = o.TokenCount
	}
	if o.TokenLength != nil {
		toSerialize["token_length"] = o.TokenLength
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticatorStaticStage struct {
	value *AuthenticatorStaticStage
	isSet bool
}

func (v NullableAuthenticatorStaticStage) Get() *AuthenticatorStaticStage {
	return v.value
}

func (v *NullableAuthenticatorStaticStage) Set(val *AuthenticatorStaticStage) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorStaticStage) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorStaticStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorStaticStage(val *AuthenticatorStaticStage) *NullableAuthenticatorStaticStage {
	return &NullableAuthenticatorStaticStage{value: val, isSet: true}
}

func (v NullableAuthenticatorStaticStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorStaticStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

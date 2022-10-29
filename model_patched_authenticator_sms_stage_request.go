/*
authentik

Making authentication simple.

API version: 2022.10.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedAuthenticatorSMSStageRequest AuthenticatorSMSStage Serializer
type PatchedAuthenticatorSMSStageRequest struct {
	Name    *string          `json:"name,omitempty"`
	FlowSet []FlowSetRequest `json:"flow_set,omitempty"`
	// Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage.
	ConfigureFlow NullableString `json:"configure_flow,omitempty"`
	Provider      *ProviderEnum  `json:"provider,omitempty"`
	FromNumber    *string        `json:"from_number,omitempty"`
	AccountSid    *string        `json:"account_sid,omitempty"`
	Auth          *string        `json:"auth,omitempty"`
	AuthPassword  *string        `json:"auth_password,omitempty"`
	AuthType      *AuthTypeEnum  `json:"auth_type,omitempty"`
	// When enabled, the Phone number is only used during enrollment to verify the users authenticity. Only a hash of the phone number is saved to ensure it is not re-used in the future.
	VerifyOnly *bool `json:"verify_only,omitempty"`
	// Optionally modify the payload being sent to custom providers.
	Mapping NullableString `json:"mapping,omitempty"`
}

// NewPatchedAuthenticatorSMSStageRequest instantiates a new PatchedAuthenticatorSMSStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedAuthenticatorSMSStageRequest() *PatchedAuthenticatorSMSStageRequest {
	this := PatchedAuthenticatorSMSStageRequest{}
	return &this
}

// NewPatchedAuthenticatorSMSStageRequestWithDefaults instantiates a new PatchedAuthenticatorSMSStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedAuthenticatorSMSStageRequestWithDefaults() *PatchedAuthenticatorSMSStageRequest {
	this := PatchedAuthenticatorSMSStageRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedAuthenticatorSMSStageRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorSMSStageRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedAuthenticatorSMSStageRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedAuthenticatorSMSStageRequest) SetName(v string) {
	o.Name = &v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PatchedAuthenticatorSMSStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorSMSStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PatchedAuthenticatorSMSStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *PatchedAuthenticatorSMSStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetConfigureFlow returns the ConfigureFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedAuthenticatorSMSStageRequest) GetConfigureFlow() string {
	if o == nil || o.ConfigureFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfigureFlow.Get()
}

// GetConfigureFlowOk returns a tuple with the ConfigureFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedAuthenticatorSMSStageRequest) GetConfigureFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigureFlow.Get(), o.ConfigureFlow.IsSet()
}

// HasConfigureFlow returns a boolean if a field has been set.
func (o *PatchedAuthenticatorSMSStageRequest) HasConfigureFlow() bool {
	if o != nil && o.ConfigureFlow.IsSet() {
		return true
	}

	return false
}

// SetConfigureFlow gets a reference to the given NullableString and assigns it to the ConfigureFlow field.
func (o *PatchedAuthenticatorSMSStageRequest) SetConfigureFlow(v string) {
	o.ConfigureFlow.Set(&v)
}

// SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil
func (o *PatchedAuthenticatorSMSStageRequest) SetConfigureFlowNil() {
	o.ConfigureFlow.Set(nil)
}

// UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
func (o *PatchedAuthenticatorSMSStageRequest) UnsetConfigureFlow() {
	o.ConfigureFlow.Unset()
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *PatchedAuthenticatorSMSStageRequest) GetProvider() ProviderEnum {
	if o == nil || o.Provider == nil {
		var ret ProviderEnum
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorSMSStageRequest) GetProviderOk() (*ProviderEnum, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *PatchedAuthenticatorSMSStageRequest) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given ProviderEnum and assigns it to the Provider field.
func (o *PatchedAuthenticatorSMSStageRequest) SetProvider(v ProviderEnum) {
	o.Provider = &v
}

// GetFromNumber returns the FromNumber field value if set, zero value otherwise.
func (o *PatchedAuthenticatorSMSStageRequest) GetFromNumber() string {
	if o == nil || o.FromNumber == nil {
		var ret string
		return ret
	}
	return *o.FromNumber
}

// GetFromNumberOk returns a tuple with the FromNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorSMSStageRequest) GetFromNumberOk() (*string, bool) {
	if o == nil || o.FromNumber == nil {
		return nil, false
	}
	return o.FromNumber, true
}

// HasFromNumber returns a boolean if a field has been set.
func (o *PatchedAuthenticatorSMSStageRequest) HasFromNumber() bool {
	if o != nil && o.FromNumber != nil {
		return true
	}

	return false
}

// SetFromNumber gets a reference to the given string and assigns it to the FromNumber field.
func (o *PatchedAuthenticatorSMSStageRequest) SetFromNumber(v string) {
	o.FromNumber = &v
}

// GetAccountSid returns the AccountSid field value if set, zero value otherwise.
func (o *PatchedAuthenticatorSMSStageRequest) GetAccountSid() string {
	if o == nil || o.AccountSid == nil {
		var ret string
		return ret
	}
	return *o.AccountSid
}

// GetAccountSidOk returns a tuple with the AccountSid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorSMSStageRequest) GetAccountSidOk() (*string, bool) {
	if o == nil || o.AccountSid == nil {
		return nil, false
	}
	return o.AccountSid, true
}

// HasAccountSid returns a boolean if a field has been set.
func (o *PatchedAuthenticatorSMSStageRequest) HasAccountSid() bool {
	if o != nil && o.AccountSid != nil {
		return true
	}

	return false
}

// SetAccountSid gets a reference to the given string and assigns it to the AccountSid field.
func (o *PatchedAuthenticatorSMSStageRequest) SetAccountSid(v string) {
	o.AccountSid = &v
}

// GetAuth returns the Auth field value if set, zero value otherwise.
func (o *PatchedAuthenticatorSMSStageRequest) GetAuth() string {
	if o == nil || o.Auth == nil {
		var ret string
		return ret
	}
	return *o.Auth
}

// GetAuthOk returns a tuple with the Auth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorSMSStageRequest) GetAuthOk() (*string, bool) {
	if o == nil || o.Auth == nil {
		return nil, false
	}
	return o.Auth, true
}

// HasAuth returns a boolean if a field has been set.
func (o *PatchedAuthenticatorSMSStageRequest) HasAuth() bool {
	if o != nil && o.Auth != nil {
		return true
	}

	return false
}

// SetAuth gets a reference to the given string and assigns it to the Auth field.
func (o *PatchedAuthenticatorSMSStageRequest) SetAuth(v string) {
	o.Auth = &v
}

// GetAuthPassword returns the AuthPassword field value if set, zero value otherwise.
func (o *PatchedAuthenticatorSMSStageRequest) GetAuthPassword() string {
	if o == nil || o.AuthPassword == nil {
		var ret string
		return ret
	}
	return *o.AuthPassword
}

// GetAuthPasswordOk returns a tuple with the AuthPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorSMSStageRequest) GetAuthPasswordOk() (*string, bool) {
	if o == nil || o.AuthPassword == nil {
		return nil, false
	}
	return o.AuthPassword, true
}

// HasAuthPassword returns a boolean if a field has been set.
func (o *PatchedAuthenticatorSMSStageRequest) HasAuthPassword() bool {
	if o != nil && o.AuthPassword != nil {
		return true
	}

	return false
}

// SetAuthPassword gets a reference to the given string and assigns it to the AuthPassword field.
func (o *PatchedAuthenticatorSMSStageRequest) SetAuthPassword(v string) {
	o.AuthPassword = &v
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *PatchedAuthenticatorSMSStageRequest) GetAuthType() AuthTypeEnum {
	if o == nil || o.AuthType == nil {
		var ret AuthTypeEnum
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorSMSStageRequest) GetAuthTypeOk() (*AuthTypeEnum, bool) {
	if o == nil || o.AuthType == nil {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *PatchedAuthenticatorSMSStageRequest) HasAuthType() bool {
	if o != nil && o.AuthType != nil {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given AuthTypeEnum and assigns it to the AuthType field.
func (o *PatchedAuthenticatorSMSStageRequest) SetAuthType(v AuthTypeEnum) {
	o.AuthType = &v
}

// GetVerifyOnly returns the VerifyOnly field value if set, zero value otherwise.
func (o *PatchedAuthenticatorSMSStageRequest) GetVerifyOnly() bool {
	if o == nil || o.VerifyOnly == nil {
		var ret bool
		return ret
	}
	return *o.VerifyOnly
}

// GetVerifyOnlyOk returns a tuple with the VerifyOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorSMSStageRequest) GetVerifyOnlyOk() (*bool, bool) {
	if o == nil || o.VerifyOnly == nil {
		return nil, false
	}
	return o.VerifyOnly, true
}

// HasVerifyOnly returns a boolean if a field has been set.
func (o *PatchedAuthenticatorSMSStageRequest) HasVerifyOnly() bool {
	if o != nil && o.VerifyOnly != nil {
		return true
	}

	return false
}

// SetVerifyOnly gets a reference to the given bool and assigns it to the VerifyOnly field.
func (o *PatchedAuthenticatorSMSStageRequest) SetVerifyOnly(v bool) {
	o.VerifyOnly = &v
}

// GetMapping returns the Mapping field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedAuthenticatorSMSStageRequest) GetMapping() string {
	if o == nil || o.Mapping.Get() == nil {
		var ret string
		return ret
	}
	return *o.Mapping.Get()
}

// GetMappingOk returns a tuple with the Mapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedAuthenticatorSMSStageRequest) GetMappingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mapping.Get(), o.Mapping.IsSet()
}

// HasMapping returns a boolean if a field has been set.
func (o *PatchedAuthenticatorSMSStageRequest) HasMapping() bool {
	if o != nil && o.Mapping.IsSet() {
		return true
	}

	return false
}

// SetMapping gets a reference to the given NullableString and assigns it to the Mapping field.
func (o *PatchedAuthenticatorSMSStageRequest) SetMapping(v string) {
	o.Mapping.Set(&v)
}

// SetMappingNil sets the value for Mapping to be an explicit nil
func (o *PatchedAuthenticatorSMSStageRequest) SetMappingNil() {
	o.Mapping.Set(nil)
}

// UnsetMapping ensures that no value is present for Mapping, not even an explicit nil
func (o *PatchedAuthenticatorSMSStageRequest) UnsetMapping() {
	o.Mapping.Unset()
}

func (o PatchedAuthenticatorSMSStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if o.ConfigureFlow.IsSet() {
		toSerialize["configure_flow"] = o.ConfigureFlow.Get()
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.FromNumber != nil {
		toSerialize["from_number"] = o.FromNumber
	}
	if o.AccountSid != nil {
		toSerialize["account_sid"] = o.AccountSid
	}
	if o.Auth != nil {
		toSerialize["auth"] = o.Auth
	}
	if o.AuthPassword != nil {
		toSerialize["auth_password"] = o.AuthPassword
	}
	if o.AuthType != nil {
		toSerialize["auth_type"] = o.AuthType
	}
	if o.VerifyOnly != nil {
		toSerialize["verify_only"] = o.VerifyOnly
	}
	if o.Mapping.IsSet() {
		toSerialize["mapping"] = o.Mapping.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedAuthenticatorSMSStageRequest struct {
	value *PatchedAuthenticatorSMSStageRequest
	isSet bool
}

func (v NullablePatchedAuthenticatorSMSStageRequest) Get() *PatchedAuthenticatorSMSStageRequest {
	return v.value
}

func (v *NullablePatchedAuthenticatorSMSStageRequest) Set(val *PatchedAuthenticatorSMSStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedAuthenticatorSMSStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedAuthenticatorSMSStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedAuthenticatorSMSStageRequest(val *PatchedAuthenticatorSMSStageRequest) *NullablePatchedAuthenticatorSMSStageRequest {
	return &NullablePatchedAuthenticatorSMSStageRequest{value: val, isSet: true}
}

func (v NullablePatchedAuthenticatorSMSStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedAuthenticatorSMSStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

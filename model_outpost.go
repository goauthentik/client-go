/*
authentik

Making authentication simple.

API version: 2025.2.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Outpost Outpost Serializer
type Outpost struct {
	Pk           string          `json:"pk"`
	Name         string          `json:"name"`
	Type         OutpostTypeEnum `json:"type"`
	Providers    []int32         `json:"providers"`
	ProvidersObj []Provider      `json:"providers_obj"`
	// Select Service-Connection authentik should use to manage this outpost. Leave empty if authentik should not handle the deployment.
	ServiceConnection    NullableString    `json:"service_connection,omitempty"`
	ServiceConnectionObj ServiceConnection `json:"service_connection_obj"`
	RefreshIntervalS     int32             `json:"refresh_interval_s"`
	// Get Token identifier
	TokenIdentifier string                 `json:"token_identifier"`
	Config          map[string]interface{} `json:"config"`
	// Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update.
	Managed NullableString `json:"managed,omitempty"`
}

// NewOutpost instantiates a new Outpost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutpost(pk string, name string, type_ OutpostTypeEnum, providers []int32, providersObj []Provider, serviceConnectionObj ServiceConnection, refreshIntervalS int32, tokenIdentifier string, config map[string]interface{}) *Outpost {
	this := Outpost{}
	this.Pk = pk
	this.Name = name
	this.Type = type_
	this.Providers = providers
	this.ProvidersObj = providersObj
	this.ServiceConnectionObj = serviceConnectionObj
	this.RefreshIntervalS = refreshIntervalS
	this.TokenIdentifier = tokenIdentifier
	this.Config = config
	return &this
}

// NewOutpostWithDefaults instantiates a new Outpost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutpostWithDefaults() *Outpost {
	this := Outpost{}
	return &this
}

// GetPk returns the Pk field value
func (o *Outpost) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *Outpost) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *Outpost) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *Outpost) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Outpost) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Outpost) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *Outpost) GetType() OutpostTypeEnum {
	if o == nil {
		var ret OutpostTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Outpost) GetTypeOk() (*OutpostTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Outpost) SetType(v OutpostTypeEnum) {
	o.Type = v
}

// GetProviders returns the Providers field value
func (o *Outpost) GetProviders() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value
// and a boolean to check if the value has been set.
func (o *Outpost) GetProvidersOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Providers, true
}

// SetProviders sets field value
func (o *Outpost) SetProviders(v []int32) {
	o.Providers = v
}

// GetProvidersObj returns the ProvidersObj field value
func (o *Outpost) GetProvidersObj() []Provider {
	if o == nil {
		var ret []Provider
		return ret
	}

	return o.ProvidersObj
}

// GetProvidersObjOk returns a tuple with the ProvidersObj field value
// and a boolean to check if the value has been set.
func (o *Outpost) GetProvidersObjOk() ([]Provider, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProvidersObj, true
}

// SetProvidersObj sets field value
func (o *Outpost) SetProvidersObj(v []Provider) {
	o.ProvidersObj = v
}

// GetServiceConnection returns the ServiceConnection field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Outpost) GetServiceConnection() string {
	if o == nil || o.ServiceConnection.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServiceConnection.Get()
}

// GetServiceConnectionOk returns a tuple with the ServiceConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Outpost) GetServiceConnectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceConnection.Get(), o.ServiceConnection.IsSet()
}

// HasServiceConnection returns a boolean if a field has been set.
func (o *Outpost) HasServiceConnection() bool {
	if o != nil && o.ServiceConnection.IsSet() {
		return true
	}

	return false
}

// SetServiceConnection gets a reference to the given NullableString and assigns it to the ServiceConnection field.
func (o *Outpost) SetServiceConnection(v string) {
	o.ServiceConnection.Set(&v)
}

// SetServiceConnectionNil sets the value for ServiceConnection to be an explicit nil
func (o *Outpost) SetServiceConnectionNil() {
	o.ServiceConnection.Set(nil)
}

// UnsetServiceConnection ensures that no value is present for ServiceConnection, not even an explicit nil
func (o *Outpost) UnsetServiceConnection() {
	o.ServiceConnection.Unset()
}

// GetServiceConnectionObj returns the ServiceConnectionObj field value
func (o *Outpost) GetServiceConnectionObj() ServiceConnection {
	if o == nil {
		var ret ServiceConnection
		return ret
	}

	return o.ServiceConnectionObj
}

// GetServiceConnectionObjOk returns a tuple with the ServiceConnectionObj field value
// and a boolean to check if the value has been set.
func (o *Outpost) GetServiceConnectionObjOk() (*ServiceConnection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceConnectionObj, true
}

// SetServiceConnectionObj sets field value
func (o *Outpost) SetServiceConnectionObj(v ServiceConnection) {
	o.ServiceConnectionObj = v
}

// GetRefreshIntervalS returns the RefreshIntervalS field value
func (o *Outpost) GetRefreshIntervalS() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RefreshIntervalS
}

// GetRefreshIntervalSOk returns a tuple with the RefreshIntervalS field value
// and a boolean to check if the value has been set.
func (o *Outpost) GetRefreshIntervalSOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefreshIntervalS, true
}

// SetRefreshIntervalS sets field value
func (o *Outpost) SetRefreshIntervalS(v int32) {
	o.RefreshIntervalS = v
}

// GetTokenIdentifier returns the TokenIdentifier field value
func (o *Outpost) GetTokenIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenIdentifier
}

// GetTokenIdentifierOk returns a tuple with the TokenIdentifier field value
// and a boolean to check if the value has been set.
func (o *Outpost) GetTokenIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenIdentifier, true
}

// SetTokenIdentifier sets field value
func (o *Outpost) SetTokenIdentifier(v string) {
	o.TokenIdentifier = v
}

// GetConfig returns the Config field value
func (o *Outpost) GetConfig() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *Outpost) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Config, true
}

// SetConfig sets field value
func (o *Outpost) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetManaged returns the Managed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Outpost) GetManaged() string {
	if o == nil || o.Managed.Get() == nil {
		var ret string
		return ret
	}
	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Outpost) GetManagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// HasManaged returns a boolean if a field has been set.
func (o *Outpost) HasManaged() bool {
	if o != nil && o.Managed.IsSet() {
		return true
	}

	return false
}

// SetManaged gets a reference to the given NullableString and assigns it to the Managed field.
func (o *Outpost) SetManaged(v string) {
	o.Managed.Set(&v)
}

// SetManagedNil sets the value for Managed to be an explicit nil
func (o *Outpost) SetManagedNil() {
	o.Managed.Set(nil)
}

// UnsetManaged ensures that no value is present for Managed, not even an explicit nil
func (o *Outpost) UnsetManaged() {
	o.Managed.Unset()
}

func (o Outpost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["providers"] = o.Providers
	}
	if true {
		toSerialize["providers_obj"] = o.ProvidersObj
	}
	if o.ServiceConnection.IsSet() {
		toSerialize["service_connection"] = o.ServiceConnection.Get()
	}
	if true {
		toSerialize["service_connection_obj"] = o.ServiceConnectionObj
	}
	if true {
		toSerialize["refresh_interval_s"] = o.RefreshIntervalS
	}
	if true {
		toSerialize["token_identifier"] = o.TokenIdentifier
	}
	if true {
		toSerialize["config"] = o.Config
	}
	if o.Managed.IsSet() {
		toSerialize["managed"] = o.Managed.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableOutpost struct {
	value *Outpost
	isSet bool
}

func (v NullableOutpost) Get() *Outpost {
	return v.value
}

func (v *NullableOutpost) Set(val *Outpost) {
	v.value = val
	v.isSet = true
}

func (v NullableOutpost) IsSet() bool {
	return v.isSet
}

func (v *NullableOutpost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutpost(val *Outpost) *NullableOutpost {
	return &NullableOutpost{value: val, isSet: true}
}

func (v NullableOutpost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutpost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

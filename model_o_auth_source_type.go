/*
authentik

Making authentication simple.

API version: 2023.3.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// OAuthSourceType struct for OAuthSourceType
type OAuthSourceType struct {
	Name             string         `json:"name"`
	Slug             string         `json:"slug"`
	UrlsCustomizable bool           `json:"urls_customizable"`
	RequestTokenUrl  NullableString `json:"request_token_url"`
	AuthorizationUrl NullableString `json:"authorization_url"`
	AccessTokenUrl   NullableString `json:"access_token_url"`
	ProfileUrl       NullableString `json:"profile_url"`
}

// NewOAuthSourceType instantiates a new OAuthSourceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthSourceType(name string, slug string, urlsCustomizable bool, requestTokenUrl NullableString, authorizationUrl NullableString, accessTokenUrl NullableString, profileUrl NullableString) *OAuthSourceType {
	this := OAuthSourceType{}
	this.Name = name
	this.Slug = slug
	this.UrlsCustomizable = urlsCustomizable
	this.RequestTokenUrl = requestTokenUrl
	this.AuthorizationUrl = authorizationUrl
	this.AccessTokenUrl = accessTokenUrl
	this.ProfileUrl = profileUrl
	return &this
}

// NewOAuthSourceTypeWithDefaults instantiates a new OAuthSourceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthSourceTypeWithDefaults() *OAuthSourceType {
	this := OAuthSourceType{}
	return &this
}

// GetName returns the Name field value
func (o *OAuthSourceType) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OAuthSourceType) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OAuthSourceType) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *OAuthSourceType) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *OAuthSourceType) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *OAuthSourceType) SetSlug(v string) {
	o.Slug = v
}

// GetUrlsCustomizable returns the UrlsCustomizable field value
func (o *OAuthSourceType) GetUrlsCustomizable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UrlsCustomizable
}

// GetUrlsCustomizableOk returns a tuple with the UrlsCustomizable field value
// and a boolean to check if the value has been set.
func (o *OAuthSourceType) GetUrlsCustomizableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UrlsCustomizable, true
}

// SetUrlsCustomizable sets field value
func (o *OAuthSourceType) SetUrlsCustomizable(v bool) {
	o.UrlsCustomizable = v
}

// GetRequestTokenUrl returns the RequestTokenUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OAuthSourceType) GetRequestTokenUrl() string {
	if o == nil || o.RequestTokenUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.RequestTokenUrl.Get()
}

// GetRequestTokenUrlOk returns a tuple with the RequestTokenUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuthSourceType) GetRequestTokenUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestTokenUrl.Get(), o.RequestTokenUrl.IsSet()
}

// SetRequestTokenUrl sets field value
func (o *OAuthSourceType) SetRequestTokenUrl(v string) {
	o.RequestTokenUrl.Set(&v)
}

// GetAuthorizationUrl returns the AuthorizationUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OAuthSourceType) GetAuthorizationUrl() string {
	if o == nil || o.AuthorizationUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.AuthorizationUrl.Get()
}

// GetAuthorizationUrlOk returns a tuple with the AuthorizationUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuthSourceType) GetAuthorizationUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorizationUrl.Get(), o.AuthorizationUrl.IsSet()
}

// SetAuthorizationUrl sets field value
func (o *OAuthSourceType) SetAuthorizationUrl(v string) {
	o.AuthorizationUrl.Set(&v)
}

// GetAccessTokenUrl returns the AccessTokenUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OAuthSourceType) GetAccessTokenUrl() string {
	if o == nil || o.AccessTokenUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.AccessTokenUrl.Get()
}

// GetAccessTokenUrlOk returns a tuple with the AccessTokenUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuthSourceType) GetAccessTokenUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessTokenUrl.Get(), o.AccessTokenUrl.IsSet()
}

// SetAccessTokenUrl sets field value
func (o *OAuthSourceType) SetAccessTokenUrl(v string) {
	o.AccessTokenUrl.Set(&v)
}

// GetProfileUrl returns the ProfileUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OAuthSourceType) GetProfileUrl() string {
	if o == nil || o.ProfileUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.ProfileUrl.Get()
}

// GetProfileUrlOk returns a tuple with the ProfileUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuthSourceType) GetProfileUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfileUrl.Get(), o.ProfileUrl.IsSet()
}

// SetProfileUrl sets field value
func (o *OAuthSourceType) SetProfileUrl(v string) {
	o.ProfileUrl.Set(&v)
}

func (o OAuthSourceType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["slug"] = o.Slug
	}
	if true {
		toSerialize["urls_customizable"] = o.UrlsCustomizable
	}
	if true {
		toSerialize["request_token_url"] = o.RequestTokenUrl.Get()
	}
	if true {
		toSerialize["authorization_url"] = o.AuthorizationUrl.Get()
	}
	if true {
		toSerialize["access_token_url"] = o.AccessTokenUrl.Get()
	}
	if true {
		toSerialize["profile_url"] = o.ProfileUrl.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableOAuthSourceType struct {
	value *OAuthSourceType
	isSet bool
}

func (v NullableOAuthSourceType) Get() *OAuthSourceType {
	return v.value
}

func (v *NullableOAuthSourceType) Set(val *OAuthSourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthSourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthSourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthSourceType(val *OAuthSourceType) *NullableOAuthSourceType {
	return &NullableOAuthSourceType{value: val, isSet: true}
}

func (v NullableOAuthSourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthSourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

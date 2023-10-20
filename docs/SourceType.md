# SourceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Slug** | **string** |  | 
**UrlsCustomizable** | **bool** |  | 
**RequestTokenUrl** | **NullableString** |  | [readonly] 
**AuthorizationUrl** | **NullableString** |  | [readonly] 
**AccessTokenUrl** | **NullableString** |  | [readonly] 
**ProfileUrl** | **NullableString** |  | [readonly] 
**OidcWellKnownUrl** | **NullableString** |  | [readonly] 
**OidcJwksUrl** | **NullableString** |  | [readonly] 

## Methods

### NewSourceType

`func NewSourceType(name string, slug string, urlsCustomizable bool, requestTokenUrl NullableString, authorizationUrl NullableString, accessTokenUrl NullableString, profileUrl NullableString, oidcWellKnownUrl NullableString, oidcJwksUrl NullableString, ) *SourceType`

NewSourceType instantiates a new SourceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceTypeWithDefaults

`func NewSourceTypeWithDefaults() *SourceType`

NewSourceTypeWithDefaults instantiates a new SourceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SourceType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourceType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourceType) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *SourceType) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *SourceType) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *SourceType) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetUrlsCustomizable

`func (o *SourceType) GetUrlsCustomizable() bool`

GetUrlsCustomizable returns the UrlsCustomizable field if non-nil, zero value otherwise.

### GetUrlsCustomizableOk

`func (o *SourceType) GetUrlsCustomizableOk() (*bool, bool)`

GetUrlsCustomizableOk returns a tuple with the UrlsCustomizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlsCustomizable

`func (o *SourceType) SetUrlsCustomizable(v bool)`

SetUrlsCustomizable sets UrlsCustomizable field to given value.


### GetRequestTokenUrl

`func (o *SourceType) GetRequestTokenUrl() string`

GetRequestTokenUrl returns the RequestTokenUrl field if non-nil, zero value otherwise.

### GetRequestTokenUrlOk

`func (o *SourceType) GetRequestTokenUrlOk() (*string, bool)`

GetRequestTokenUrlOk returns a tuple with the RequestTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTokenUrl

`func (o *SourceType) SetRequestTokenUrl(v string)`

SetRequestTokenUrl sets RequestTokenUrl field to given value.


### SetRequestTokenUrlNil

`func (o *SourceType) SetRequestTokenUrlNil(b bool)`

 SetRequestTokenUrlNil sets the value for RequestTokenUrl to be an explicit nil

### UnsetRequestTokenUrl
`func (o *SourceType) UnsetRequestTokenUrl()`

UnsetRequestTokenUrl ensures that no value is present for RequestTokenUrl, not even an explicit nil
### GetAuthorizationUrl

`func (o *SourceType) GetAuthorizationUrl() string`

GetAuthorizationUrl returns the AuthorizationUrl field if non-nil, zero value otherwise.

### GetAuthorizationUrlOk

`func (o *SourceType) GetAuthorizationUrlOk() (*string, bool)`

GetAuthorizationUrlOk returns a tuple with the AuthorizationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationUrl

`func (o *SourceType) SetAuthorizationUrl(v string)`

SetAuthorizationUrl sets AuthorizationUrl field to given value.


### SetAuthorizationUrlNil

`func (o *SourceType) SetAuthorizationUrlNil(b bool)`

 SetAuthorizationUrlNil sets the value for AuthorizationUrl to be an explicit nil

### UnsetAuthorizationUrl
`func (o *SourceType) UnsetAuthorizationUrl()`

UnsetAuthorizationUrl ensures that no value is present for AuthorizationUrl, not even an explicit nil
### GetAccessTokenUrl

`func (o *SourceType) GetAccessTokenUrl() string`

GetAccessTokenUrl returns the AccessTokenUrl field if non-nil, zero value otherwise.

### GetAccessTokenUrlOk

`func (o *SourceType) GetAccessTokenUrlOk() (*string, bool)`

GetAccessTokenUrlOk returns a tuple with the AccessTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenUrl

`func (o *SourceType) SetAccessTokenUrl(v string)`

SetAccessTokenUrl sets AccessTokenUrl field to given value.


### SetAccessTokenUrlNil

`func (o *SourceType) SetAccessTokenUrlNil(b bool)`

 SetAccessTokenUrlNil sets the value for AccessTokenUrl to be an explicit nil

### UnsetAccessTokenUrl
`func (o *SourceType) UnsetAccessTokenUrl()`

UnsetAccessTokenUrl ensures that no value is present for AccessTokenUrl, not even an explicit nil
### GetProfileUrl

`func (o *SourceType) GetProfileUrl() string`

GetProfileUrl returns the ProfileUrl field if non-nil, zero value otherwise.

### GetProfileUrlOk

`func (o *SourceType) GetProfileUrlOk() (*string, bool)`

GetProfileUrlOk returns a tuple with the ProfileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileUrl

`func (o *SourceType) SetProfileUrl(v string)`

SetProfileUrl sets ProfileUrl field to given value.


### SetProfileUrlNil

`func (o *SourceType) SetProfileUrlNil(b bool)`

 SetProfileUrlNil sets the value for ProfileUrl to be an explicit nil

### UnsetProfileUrl
`func (o *SourceType) UnsetProfileUrl()`

UnsetProfileUrl ensures that no value is present for ProfileUrl, not even an explicit nil
### GetOidcWellKnownUrl

`func (o *SourceType) GetOidcWellKnownUrl() string`

GetOidcWellKnownUrl returns the OidcWellKnownUrl field if non-nil, zero value otherwise.

### GetOidcWellKnownUrlOk

`func (o *SourceType) GetOidcWellKnownUrlOk() (*string, bool)`

GetOidcWellKnownUrlOk returns a tuple with the OidcWellKnownUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcWellKnownUrl

`func (o *SourceType) SetOidcWellKnownUrl(v string)`

SetOidcWellKnownUrl sets OidcWellKnownUrl field to given value.


### SetOidcWellKnownUrlNil

`func (o *SourceType) SetOidcWellKnownUrlNil(b bool)`

 SetOidcWellKnownUrlNil sets the value for OidcWellKnownUrl to be an explicit nil

### UnsetOidcWellKnownUrl
`func (o *SourceType) UnsetOidcWellKnownUrl()`

UnsetOidcWellKnownUrl ensures that no value is present for OidcWellKnownUrl, not even an explicit nil
### GetOidcJwksUrl

`func (o *SourceType) GetOidcJwksUrl() string`

GetOidcJwksUrl returns the OidcJwksUrl field if non-nil, zero value otherwise.

### GetOidcJwksUrlOk

`func (o *SourceType) GetOidcJwksUrlOk() (*string, bool)`

GetOidcJwksUrlOk returns a tuple with the OidcJwksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcJwksUrl

`func (o *SourceType) SetOidcJwksUrl(v string)`

SetOidcJwksUrl sets OidcJwksUrl field to given value.


### SetOidcJwksUrlNil

`func (o *SourceType) SetOidcJwksUrlNil(b bool)`

 SetOidcJwksUrlNil sets the value for OidcJwksUrl to be an explicit nil

### UnsetOidcJwksUrl
`func (o *SourceType) UnsetOidcJwksUrl()`

UnsetOidcJwksUrl ensures that no value is present for OidcJwksUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



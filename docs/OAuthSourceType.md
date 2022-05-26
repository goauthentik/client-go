# OAuthSourceType

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

## Methods

### NewOAuthSourceType

`func NewOAuthSourceType(name string, slug string, urlsCustomizable bool, requestTokenUrl NullableString, authorizationUrl NullableString, accessTokenUrl NullableString, profileUrl NullableString, ) *OAuthSourceType`

NewOAuthSourceType instantiates a new OAuthSourceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthSourceTypeWithDefaults

`func NewOAuthSourceTypeWithDefaults() *OAuthSourceType`

NewOAuthSourceTypeWithDefaults instantiates a new OAuthSourceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OAuthSourceType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OAuthSourceType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OAuthSourceType) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *OAuthSourceType) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *OAuthSourceType) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *OAuthSourceType) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetUrlsCustomizable

`func (o *OAuthSourceType) GetUrlsCustomizable() bool`

GetUrlsCustomizable returns the UrlsCustomizable field if non-nil, zero value otherwise.

### GetUrlsCustomizableOk

`func (o *OAuthSourceType) GetUrlsCustomizableOk() (*bool, bool)`

GetUrlsCustomizableOk returns a tuple with the UrlsCustomizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlsCustomizable

`func (o *OAuthSourceType) SetUrlsCustomizable(v bool)`

SetUrlsCustomizable sets UrlsCustomizable field to given value.


### GetRequestTokenUrl

`func (o *OAuthSourceType) GetRequestTokenUrl() string`

GetRequestTokenUrl returns the RequestTokenUrl field if non-nil, zero value otherwise.

### GetRequestTokenUrlOk

`func (o *OAuthSourceType) GetRequestTokenUrlOk() (*string, bool)`

GetRequestTokenUrlOk returns a tuple with the RequestTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTokenUrl

`func (o *OAuthSourceType) SetRequestTokenUrl(v string)`

SetRequestTokenUrl sets RequestTokenUrl field to given value.


### SetRequestTokenUrlNil

`func (o *OAuthSourceType) SetRequestTokenUrlNil(b bool)`

 SetRequestTokenUrlNil sets the value for RequestTokenUrl to be an explicit nil

### UnsetRequestTokenUrl
`func (o *OAuthSourceType) UnsetRequestTokenUrl()`

UnsetRequestTokenUrl ensures that no value is present for RequestTokenUrl, not even an explicit nil
### GetAuthorizationUrl

`func (o *OAuthSourceType) GetAuthorizationUrl() string`

GetAuthorizationUrl returns the AuthorizationUrl field if non-nil, zero value otherwise.

### GetAuthorizationUrlOk

`func (o *OAuthSourceType) GetAuthorizationUrlOk() (*string, bool)`

GetAuthorizationUrlOk returns a tuple with the AuthorizationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationUrl

`func (o *OAuthSourceType) SetAuthorizationUrl(v string)`

SetAuthorizationUrl sets AuthorizationUrl field to given value.


### SetAuthorizationUrlNil

`func (o *OAuthSourceType) SetAuthorizationUrlNil(b bool)`

 SetAuthorizationUrlNil sets the value for AuthorizationUrl to be an explicit nil

### UnsetAuthorizationUrl
`func (o *OAuthSourceType) UnsetAuthorizationUrl()`

UnsetAuthorizationUrl ensures that no value is present for AuthorizationUrl, not even an explicit nil
### GetAccessTokenUrl

`func (o *OAuthSourceType) GetAccessTokenUrl() string`

GetAccessTokenUrl returns the AccessTokenUrl field if non-nil, zero value otherwise.

### GetAccessTokenUrlOk

`func (o *OAuthSourceType) GetAccessTokenUrlOk() (*string, bool)`

GetAccessTokenUrlOk returns a tuple with the AccessTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenUrl

`func (o *OAuthSourceType) SetAccessTokenUrl(v string)`

SetAccessTokenUrl sets AccessTokenUrl field to given value.


### SetAccessTokenUrlNil

`func (o *OAuthSourceType) SetAccessTokenUrlNil(b bool)`

 SetAccessTokenUrlNil sets the value for AccessTokenUrl to be an explicit nil

### UnsetAccessTokenUrl
`func (o *OAuthSourceType) UnsetAccessTokenUrl()`

UnsetAccessTokenUrl ensures that no value is present for AccessTokenUrl, not even an explicit nil
### GetProfileUrl

`func (o *OAuthSourceType) GetProfileUrl() string`

GetProfileUrl returns the ProfileUrl field if non-nil, zero value otherwise.

### GetProfileUrlOk

`func (o *OAuthSourceType) GetProfileUrlOk() (*string, bool)`

GetProfileUrlOk returns a tuple with the ProfileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileUrl

`func (o *OAuthSourceType) SetProfileUrl(v string)`

SetProfileUrl sets ProfileUrl field to given value.


### SetProfileUrlNil

`func (o *OAuthSourceType) SetProfileUrlNil(b bool)`

 SetProfileUrlNil sets the value for ProfileUrl to be an explicit nil

### UnsetProfileUrl
`func (o *OAuthSourceType) UnsetProfileUrl()`

UnsetProfileUrl ensures that no value is present for ProfileUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



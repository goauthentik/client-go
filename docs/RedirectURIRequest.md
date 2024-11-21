# RedirectURIRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchingMode** | [**MatchingModeEnum**](MatchingModeEnum.md) |  | 
**Url** | **string** |  | 

## Methods

### NewRedirectURIRequest

`func NewRedirectURIRequest(matchingMode MatchingModeEnum, url string, ) *RedirectURIRequest`

NewRedirectURIRequest instantiates a new RedirectURIRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedirectURIRequestWithDefaults

`func NewRedirectURIRequestWithDefaults() *RedirectURIRequest`

NewRedirectURIRequestWithDefaults instantiates a new RedirectURIRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchingMode

`func (o *RedirectURIRequest) GetMatchingMode() MatchingModeEnum`

GetMatchingMode returns the MatchingMode field if non-nil, zero value otherwise.

### GetMatchingModeOk

`func (o *RedirectURIRequest) GetMatchingModeOk() (*MatchingModeEnum, bool)`

GetMatchingModeOk returns a tuple with the MatchingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingMode

`func (o *RedirectURIRequest) SetMatchingMode(v MatchingModeEnum)`

SetMatchingMode sets MatchingMode field to given value.


### GetUrl

`func (o *RedirectURIRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RedirectURIRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RedirectURIRequest) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



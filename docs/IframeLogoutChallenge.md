# IframeLogoutChallenge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowInfo** | Pointer to [**ContextualFlowInfo**](ContextualFlowInfo.md) |  | [optional] 
**Component** | Pointer to **string** |  | [optional] [default to "ak-provider-iframe-logout"]
**ResponseErrors** | Pointer to [**map[string][]ErrorDetail**](array.md) |  | [optional] 
**LogoutUrls** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewIframeLogoutChallenge

`func NewIframeLogoutChallenge() *IframeLogoutChallenge`

NewIframeLogoutChallenge instantiates a new IframeLogoutChallenge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIframeLogoutChallengeWithDefaults

`func NewIframeLogoutChallengeWithDefaults() *IframeLogoutChallenge`

NewIframeLogoutChallengeWithDefaults instantiates a new IframeLogoutChallenge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowInfo

`func (o *IframeLogoutChallenge) GetFlowInfo() ContextualFlowInfo`

GetFlowInfo returns the FlowInfo field if non-nil, zero value otherwise.

### GetFlowInfoOk

`func (o *IframeLogoutChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool)`

GetFlowInfoOk returns a tuple with the FlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInfo

`func (o *IframeLogoutChallenge) SetFlowInfo(v ContextualFlowInfo)`

SetFlowInfo sets FlowInfo field to given value.

### HasFlowInfo

`func (o *IframeLogoutChallenge) HasFlowInfo() bool`

HasFlowInfo returns a boolean if a field has been set.

### GetComponent

`func (o *IframeLogoutChallenge) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *IframeLogoutChallenge) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *IframeLogoutChallenge) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *IframeLogoutChallenge) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetResponseErrors

`func (o *IframeLogoutChallenge) GetResponseErrors() map[string][]ErrorDetail`

GetResponseErrors returns the ResponseErrors field if non-nil, zero value otherwise.

### GetResponseErrorsOk

`func (o *IframeLogoutChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool)`

GetResponseErrorsOk returns a tuple with the ResponseErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseErrors

`func (o *IframeLogoutChallenge) SetResponseErrors(v map[string][]ErrorDetail)`

SetResponseErrors sets ResponseErrors field to given value.

### HasResponseErrors

`func (o *IframeLogoutChallenge) HasResponseErrors() bool`

HasResponseErrors returns a boolean if a field has been set.

### GetLogoutUrls

`func (o *IframeLogoutChallenge) GetLogoutUrls() []map[string]interface{}`

GetLogoutUrls returns the LogoutUrls field if non-nil, zero value otherwise.

### GetLogoutUrlsOk

`func (o *IframeLogoutChallenge) GetLogoutUrlsOk() (*[]map[string]interface{}, bool)`

GetLogoutUrlsOk returns a tuple with the LogoutUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrls

`func (o *IframeLogoutChallenge) SetLogoutUrls(v []map[string]interface{})`

SetLogoutUrls sets LogoutUrls field to given value.

### HasLogoutUrls

`func (o *IframeLogoutChallenge) HasLogoutUrls() bool`

HasLogoutUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



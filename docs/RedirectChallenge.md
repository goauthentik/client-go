# RedirectChallenge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowInfo** | Pointer to [**ContextualFlowInfo**](ContextualFlowInfo.md) |  | [optional] 
**Component** | Pointer to **string** |  | [optional] [default to "xak-flow-redirect"]
**ResponseErrors** | Pointer to [**map[string][]ErrorDetail**](array.md) |  | [optional] 
**To** | **string** |  | 
**FinalRedirect** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewRedirectChallenge

`func NewRedirectChallenge(to string, ) *RedirectChallenge`

NewRedirectChallenge instantiates a new RedirectChallenge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedirectChallengeWithDefaults

`func NewRedirectChallengeWithDefaults() *RedirectChallenge`

NewRedirectChallengeWithDefaults instantiates a new RedirectChallenge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowInfo

`func (o *RedirectChallenge) GetFlowInfo() ContextualFlowInfo`

GetFlowInfo returns the FlowInfo field if non-nil, zero value otherwise.

### GetFlowInfoOk

`func (o *RedirectChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool)`

GetFlowInfoOk returns a tuple with the FlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInfo

`func (o *RedirectChallenge) SetFlowInfo(v ContextualFlowInfo)`

SetFlowInfo sets FlowInfo field to given value.

### HasFlowInfo

`func (o *RedirectChallenge) HasFlowInfo() bool`

HasFlowInfo returns a boolean if a field has been set.

### GetComponent

`func (o *RedirectChallenge) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *RedirectChallenge) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *RedirectChallenge) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *RedirectChallenge) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetResponseErrors

`func (o *RedirectChallenge) GetResponseErrors() map[string][]ErrorDetail`

GetResponseErrors returns the ResponseErrors field if non-nil, zero value otherwise.

### GetResponseErrorsOk

`func (o *RedirectChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool)`

GetResponseErrorsOk returns a tuple with the ResponseErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseErrors

`func (o *RedirectChallenge) SetResponseErrors(v map[string][]ErrorDetail)`

SetResponseErrors sets ResponseErrors field to given value.

### HasResponseErrors

`func (o *RedirectChallenge) HasResponseErrors() bool`

HasResponseErrors returns a boolean if a field has been set.

### GetTo

`func (o *RedirectChallenge) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *RedirectChallenge) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *RedirectChallenge) SetTo(v string)`

SetTo sets To field to given value.


### GetFinalRedirect

`func (o *RedirectChallenge) GetFinalRedirect() bool`

GetFinalRedirect returns the FinalRedirect field if non-nil, zero value otherwise.

### GetFinalRedirectOk

`func (o *RedirectChallenge) GetFinalRedirectOk() (*bool, bool)`

GetFinalRedirectOk returns a tuple with the FinalRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalRedirect

`func (o *RedirectChallenge) SetFinalRedirect(v bool)`

SetFinalRedirect sets FinalRedirect field to given value.

### HasFinalRedirect

`func (o *RedirectChallenge) HasFinalRedirect() bool`

HasFinalRedirect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



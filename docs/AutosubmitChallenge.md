# AutosubmitChallenge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ChallengeChoices**](ChallengeChoices.md) |  | 
**FlowInfo** | Pointer to [**ContextualFlowInfo**](ContextualFlowInfo.md) |  | [optional] 
**Component** | Pointer to **string** |  | [optional] [default to "ak-stage-autosubmit"]
**ResponseErrors** | Pointer to [**map[string][]ErrorDetail**](array.md) |  | [optional] 
**Url** | **string** |  | 
**Attrs** | **map[string]string** |  | 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewAutosubmitChallenge

`func NewAutosubmitChallenge(type_ ChallengeChoices, url string, attrs map[string]string, ) *AutosubmitChallenge`

NewAutosubmitChallenge instantiates a new AutosubmitChallenge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutosubmitChallengeWithDefaults

`func NewAutosubmitChallengeWithDefaults() *AutosubmitChallenge`

NewAutosubmitChallengeWithDefaults instantiates a new AutosubmitChallenge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AutosubmitChallenge) GetType() ChallengeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AutosubmitChallenge) GetTypeOk() (*ChallengeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AutosubmitChallenge) SetType(v ChallengeChoices)`

SetType sets Type field to given value.


### GetFlowInfo

`func (o *AutosubmitChallenge) GetFlowInfo() ContextualFlowInfo`

GetFlowInfo returns the FlowInfo field if non-nil, zero value otherwise.

### GetFlowInfoOk

`func (o *AutosubmitChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool)`

GetFlowInfoOk returns a tuple with the FlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInfo

`func (o *AutosubmitChallenge) SetFlowInfo(v ContextualFlowInfo)`

SetFlowInfo sets FlowInfo field to given value.

### HasFlowInfo

`func (o *AutosubmitChallenge) HasFlowInfo() bool`

HasFlowInfo returns a boolean if a field has been set.

### GetComponent

`func (o *AutosubmitChallenge) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *AutosubmitChallenge) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *AutosubmitChallenge) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *AutosubmitChallenge) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetResponseErrors

`func (o *AutosubmitChallenge) GetResponseErrors() map[string][]ErrorDetail`

GetResponseErrors returns the ResponseErrors field if non-nil, zero value otherwise.

### GetResponseErrorsOk

`func (o *AutosubmitChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool)`

GetResponseErrorsOk returns a tuple with the ResponseErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseErrors

`func (o *AutosubmitChallenge) SetResponseErrors(v map[string][]ErrorDetail)`

SetResponseErrors sets ResponseErrors field to given value.

### HasResponseErrors

`func (o *AutosubmitChallenge) HasResponseErrors() bool`

HasResponseErrors returns a boolean if a field has been set.

### GetUrl

`func (o *AutosubmitChallenge) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AutosubmitChallenge) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AutosubmitChallenge) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAttrs

`func (o *AutosubmitChallenge) GetAttrs() map[string]string`

GetAttrs returns the Attrs field if non-nil, zero value otherwise.

### GetAttrsOk

`func (o *AutosubmitChallenge) GetAttrsOk() (*map[string]string, bool)`

GetAttrsOk returns a tuple with the Attrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrs

`func (o *AutosubmitChallenge) SetAttrs(v map[string]string)`

SetAttrs sets Attrs field to given value.


### GetTitle

`func (o *AutosubmitChallenge) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AutosubmitChallenge) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AutosubmitChallenge) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AutosubmitChallenge) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



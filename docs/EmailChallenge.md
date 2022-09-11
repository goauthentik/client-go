# EmailChallenge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ChallengeChoices**](ChallengeChoices.md) |  | 
**FlowInfo** | Pointer to [**ContextualFlowInfo**](ContextualFlowInfo.md) |  | [optional] 
**Component** | Pointer to **string** |  | [optional] [default to "ak-stage-email"]
**ResponseErrors** | Pointer to [**map[string][]ErrorDetail**](array.md) |  | [optional] 

## Methods

### NewEmailChallenge

`func NewEmailChallenge(type_ ChallengeChoices, ) *EmailChallenge`

NewEmailChallenge instantiates a new EmailChallenge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailChallengeWithDefaults

`func NewEmailChallengeWithDefaults() *EmailChallenge`

NewEmailChallengeWithDefaults instantiates a new EmailChallenge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EmailChallenge) GetType() ChallengeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EmailChallenge) GetTypeOk() (*ChallengeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EmailChallenge) SetType(v ChallengeChoices)`

SetType sets Type field to given value.


### GetFlowInfo

`func (o *EmailChallenge) GetFlowInfo() ContextualFlowInfo`

GetFlowInfo returns the FlowInfo field if non-nil, zero value otherwise.

### GetFlowInfoOk

`func (o *EmailChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool)`

GetFlowInfoOk returns a tuple with the FlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInfo

`func (o *EmailChallenge) SetFlowInfo(v ContextualFlowInfo)`

SetFlowInfo sets FlowInfo field to given value.

### HasFlowInfo

`func (o *EmailChallenge) HasFlowInfo() bool`

HasFlowInfo returns a boolean if a field has been set.

### GetComponent

`func (o *EmailChallenge) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *EmailChallenge) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *EmailChallenge) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *EmailChallenge) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetResponseErrors

`func (o *EmailChallenge) GetResponseErrors() map[string][]ErrorDetail`

GetResponseErrors returns the ResponseErrors field if non-nil, zero value otherwise.

### GetResponseErrorsOk

`func (o *EmailChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool)`

GetResponseErrorsOk returns a tuple with the ResponseErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseErrors

`func (o *EmailChallenge) SetResponseErrors(v map[string][]ErrorDetail)`

SetResponseErrors sets ResponseErrors field to given value.

### HasResponseErrors

`func (o *EmailChallenge) HasResponseErrors() bool`

HasResponseErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



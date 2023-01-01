# FlowErrorChallenge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "native"]
**FlowInfo** | Pointer to [**ContextualFlowInfo**](ContextualFlowInfo.md) |  | [optional] 
**Component** | Pointer to **string** |  | [optional] [default to "ak-stage-flow-error"]
**ResponseErrors** | Pointer to [**map[string][]ErrorDetail**](array.md) |  | [optional] 
**RequestId** | **string** |  | 
**Error** | Pointer to **string** |  | [optional] 
**Traceback** | Pointer to **string** |  | [optional] 

## Methods

### NewFlowErrorChallenge

`func NewFlowErrorChallenge(requestId string, ) *FlowErrorChallenge`

NewFlowErrorChallenge instantiates a new FlowErrorChallenge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowErrorChallengeWithDefaults

`func NewFlowErrorChallengeWithDefaults() *FlowErrorChallenge`

NewFlowErrorChallengeWithDefaults instantiates a new FlowErrorChallenge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FlowErrorChallenge) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FlowErrorChallenge) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FlowErrorChallenge) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FlowErrorChallenge) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFlowInfo

`func (o *FlowErrorChallenge) GetFlowInfo() ContextualFlowInfo`

GetFlowInfo returns the FlowInfo field if non-nil, zero value otherwise.

### GetFlowInfoOk

`func (o *FlowErrorChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool)`

GetFlowInfoOk returns a tuple with the FlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInfo

`func (o *FlowErrorChallenge) SetFlowInfo(v ContextualFlowInfo)`

SetFlowInfo sets FlowInfo field to given value.

### HasFlowInfo

`func (o *FlowErrorChallenge) HasFlowInfo() bool`

HasFlowInfo returns a boolean if a field has been set.

### GetComponent

`func (o *FlowErrorChallenge) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *FlowErrorChallenge) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *FlowErrorChallenge) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *FlowErrorChallenge) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetResponseErrors

`func (o *FlowErrorChallenge) GetResponseErrors() map[string][]ErrorDetail`

GetResponseErrors returns the ResponseErrors field if non-nil, zero value otherwise.

### GetResponseErrorsOk

`func (o *FlowErrorChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool)`

GetResponseErrorsOk returns a tuple with the ResponseErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseErrors

`func (o *FlowErrorChallenge) SetResponseErrors(v map[string][]ErrorDetail)`

SetResponseErrors sets ResponseErrors field to given value.

### HasResponseErrors

`func (o *FlowErrorChallenge) HasResponseErrors() bool`

HasResponseErrors returns a boolean if a field has been set.

### GetRequestId

`func (o *FlowErrorChallenge) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *FlowErrorChallenge) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *FlowErrorChallenge) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetError

`func (o *FlowErrorChallenge) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *FlowErrorChallenge) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *FlowErrorChallenge) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *FlowErrorChallenge) HasError() bool`

HasError returns a boolean if a field has been set.

### GetTraceback

`func (o *FlowErrorChallenge) GetTraceback() string`

GetTraceback returns the Traceback field if non-nil, zero value otherwise.

### GetTracebackOk

`func (o *FlowErrorChallenge) GetTracebackOk() (*string, bool)`

GetTracebackOk returns a tuple with the Traceback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceback

`func (o *FlowErrorChallenge) SetTraceback(v string)`

SetTraceback sets Traceback field to given value.

### HasTraceback

`func (o *FlowErrorChallenge) HasTraceback() bool`

HasTraceback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



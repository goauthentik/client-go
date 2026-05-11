# FrameChallenge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowInfo** | Pointer to [**ContextualFlowInfo**](ContextualFlowInfo.md) |  | [optional] 
**Component** | Pointer to **string** |  | [optional] [default to "xak-flow-frame"]
**ResponseErrors** | Pointer to [**map[string][]ErrorDetail**](array.md) |  | [optional] 
**Url** | **string** |  | 
**LoadingOverlay** | Pointer to **bool** |  | [optional] [default to false]
**LoadingText** | **string** |  | 

## Methods

### NewFrameChallenge

`func NewFrameChallenge(url string, loadingText string, ) *FrameChallenge`

NewFrameChallenge instantiates a new FrameChallenge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameChallengeWithDefaults

`func NewFrameChallengeWithDefaults() *FrameChallenge`

NewFrameChallengeWithDefaults instantiates a new FrameChallenge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowInfo

`func (o *FrameChallenge) GetFlowInfo() ContextualFlowInfo`

GetFlowInfo returns the FlowInfo field if non-nil, zero value otherwise.

### GetFlowInfoOk

`func (o *FrameChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool)`

GetFlowInfoOk returns a tuple with the FlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInfo

`func (o *FrameChallenge) SetFlowInfo(v ContextualFlowInfo)`

SetFlowInfo sets FlowInfo field to given value.

### HasFlowInfo

`func (o *FrameChallenge) HasFlowInfo() bool`

HasFlowInfo returns a boolean if a field has been set.

### GetComponent

`func (o *FrameChallenge) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *FrameChallenge) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *FrameChallenge) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *FrameChallenge) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetResponseErrors

`func (o *FrameChallenge) GetResponseErrors() map[string][]ErrorDetail`

GetResponseErrors returns the ResponseErrors field if non-nil, zero value otherwise.

### GetResponseErrorsOk

`func (o *FrameChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool)`

GetResponseErrorsOk returns a tuple with the ResponseErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseErrors

`func (o *FrameChallenge) SetResponseErrors(v map[string][]ErrorDetail)`

SetResponseErrors sets ResponseErrors field to given value.

### HasResponseErrors

`func (o *FrameChallenge) HasResponseErrors() bool`

HasResponseErrors returns a boolean if a field has been set.

### GetUrl

`func (o *FrameChallenge) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FrameChallenge) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FrameChallenge) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetLoadingOverlay

`func (o *FrameChallenge) GetLoadingOverlay() bool`

GetLoadingOverlay returns the LoadingOverlay field if non-nil, zero value otherwise.

### GetLoadingOverlayOk

`func (o *FrameChallenge) GetLoadingOverlayOk() (*bool, bool)`

GetLoadingOverlayOk returns a tuple with the LoadingOverlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadingOverlay

`func (o *FrameChallenge) SetLoadingOverlay(v bool)`

SetLoadingOverlay sets LoadingOverlay field to given value.

### HasLoadingOverlay

`func (o *FrameChallenge) HasLoadingOverlay() bool`

HasLoadingOverlay returns a boolean if a field has been set.

### GetLoadingText

`func (o *FrameChallenge) GetLoadingText() string`

GetLoadingText returns the LoadingText field if non-nil, zero value otherwise.

### GetLoadingTextOk

`func (o *FrameChallenge) GetLoadingTextOk() (*string, bool)`

GetLoadingTextOk returns a tuple with the LoadingText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadingText

`func (o *FrameChallenge) SetLoadingText(v string)`

SetLoadingText sets LoadingText field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



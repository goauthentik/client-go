# InvitationStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**FlowSet** | Pointer to [**[]FlowSetRequest**](FlowSetRequest.md) |  | [optional] 
**ContinueFlowWithoutInvitation** | Pointer to **bool** | If this flag is set, this Stage will jump to the next Stage when no Invitation is given. By default this Stage will cancel the Flow when no invitation is given. | [optional] 

## Methods

### NewInvitationStageRequest

`func NewInvitationStageRequest(name string, ) *InvitationStageRequest`

NewInvitationStageRequest instantiates a new InvitationStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationStageRequestWithDefaults

`func NewInvitationStageRequestWithDefaults() *InvitationStageRequest`

NewInvitationStageRequestWithDefaults instantiates a new InvitationStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InvitationStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvitationStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvitationStageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFlowSet

`func (o *InvitationStageRequest) GetFlowSet() []FlowSetRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *InvitationStageRequest) GetFlowSetOk() (*[]FlowSetRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *InvitationStageRequest) SetFlowSet(v []FlowSetRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *InvitationStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetContinueFlowWithoutInvitation

`func (o *InvitationStageRequest) GetContinueFlowWithoutInvitation() bool`

GetContinueFlowWithoutInvitation returns the ContinueFlowWithoutInvitation field if non-nil, zero value otherwise.

### GetContinueFlowWithoutInvitationOk

`func (o *InvitationStageRequest) GetContinueFlowWithoutInvitationOk() (*bool, bool)`

GetContinueFlowWithoutInvitationOk returns a tuple with the ContinueFlowWithoutInvitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueFlowWithoutInvitation

`func (o *InvitationStageRequest) SetContinueFlowWithoutInvitation(v bool)`

SetContinueFlowWithoutInvitation sets ContinueFlowWithoutInvitation field to given value.

### HasContinueFlowWithoutInvitation

`func (o *InvitationStageRequest) HasContinueFlowWithoutInvitation() bool`

HasContinueFlowWithoutInvitation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



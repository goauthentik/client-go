# InvitationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Expires** | Pointer to **time.Time** |  | [optional] 
**FixedData** | Pointer to **map[string]interface{}** |  | [optional] 
**SingleUse** | Pointer to **bool** | When enabled, the invitation will be deleted after usage. | [optional] 
**Flow** | Pointer to **NullableString** | When set, only the configured flow can use this invitation. | [optional] 

## Methods

### NewInvitationRequest

`func NewInvitationRequest(name string, ) *InvitationRequest`

NewInvitationRequest instantiates a new InvitationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationRequestWithDefaults

`func NewInvitationRequestWithDefaults() *InvitationRequest`

NewInvitationRequestWithDefaults instantiates a new InvitationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InvitationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvitationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvitationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetExpires

`func (o *InvitationRequest) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *InvitationRequest) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *InvitationRequest) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *InvitationRequest) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetFixedData

`func (o *InvitationRequest) GetFixedData() map[string]interface{}`

GetFixedData returns the FixedData field if non-nil, zero value otherwise.

### GetFixedDataOk

`func (o *InvitationRequest) GetFixedDataOk() (*map[string]interface{}, bool)`

GetFixedDataOk returns a tuple with the FixedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedData

`func (o *InvitationRequest) SetFixedData(v map[string]interface{})`

SetFixedData sets FixedData field to given value.

### HasFixedData

`func (o *InvitationRequest) HasFixedData() bool`

HasFixedData returns a boolean if a field has been set.

### GetSingleUse

`func (o *InvitationRequest) GetSingleUse() bool`

GetSingleUse returns the SingleUse field if non-nil, zero value otherwise.

### GetSingleUseOk

`func (o *InvitationRequest) GetSingleUseOk() (*bool, bool)`

GetSingleUseOk returns a tuple with the SingleUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleUse

`func (o *InvitationRequest) SetSingleUse(v bool)`

SetSingleUse sets SingleUse field to given value.

### HasSingleUse

`func (o *InvitationRequest) HasSingleUse() bool`

HasSingleUse returns a boolean if a field has been set.

### GetFlow

`func (o *InvitationRequest) GetFlow() string`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *InvitationRequest) GetFlowOk() (*string, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *InvitationRequest) SetFlow(v string)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *InvitationRequest) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### SetFlowNil

`func (o *InvitationRequest) SetFlowNil(b bool)`

 SetFlowNil sets the value for Flow to be an explicit nil

### UnsetFlow
`func (o *InvitationRequest) UnsetFlow()`

UnsetFlow ensures that no value is present for Flow, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



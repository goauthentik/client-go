# Invitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Expires** | Pointer to **time.Time** |  | [optional] 
**FixedData** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedBy** | [**GroupMember**](GroupMember.md) |  | [readonly] 
**SingleUse** | Pointer to **bool** | When enabled, the invitation will be deleted after usage. | [optional] 
**Flow** | Pointer to **NullableString** | When set, only the configured flow can use this invitation. | [optional] 
**FlowObj** | [**Flow**](Flow.md) |  | [readonly] 

## Methods

### NewInvitation

`func NewInvitation(pk string, name string, createdBy GroupMember, flowObj Flow, ) *Invitation`

NewInvitation instantiates a new Invitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationWithDefaults

`func NewInvitationWithDefaults() *Invitation`

NewInvitationWithDefaults instantiates a new Invitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *Invitation) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *Invitation) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *Invitation) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *Invitation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Invitation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Invitation) SetName(v string)`

SetName sets Name field to given value.


### GetExpires

`func (o *Invitation) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *Invitation) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *Invitation) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *Invitation) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetFixedData

`func (o *Invitation) GetFixedData() map[string]interface{}`

GetFixedData returns the FixedData field if non-nil, zero value otherwise.

### GetFixedDataOk

`func (o *Invitation) GetFixedDataOk() (*map[string]interface{}, bool)`

GetFixedDataOk returns a tuple with the FixedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedData

`func (o *Invitation) SetFixedData(v map[string]interface{})`

SetFixedData sets FixedData field to given value.

### HasFixedData

`func (o *Invitation) HasFixedData() bool`

HasFixedData returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Invitation) GetCreatedBy() GroupMember`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Invitation) GetCreatedByOk() (*GroupMember, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Invitation) SetCreatedBy(v GroupMember)`

SetCreatedBy sets CreatedBy field to given value.


### GetSingleUse

`func (o *Invitation) GetSingleUse() bool`

GetSingleUse returns the SingleUse field if non-nil, zero value otherwise.

### GetSingleUseOk

`func (o *Invitation) GetSingleUseOk() (*bool, bool)`

GetSingleUseOk returns a tuple with the SingleUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleUse

`func (o *Invitation) SetSingleUse(v bool)`

SetSingleUse sets SingleUse field to given value.

### HasSingleUse

`func (o *Invitation) HasSingleUse() bool`

HasSingleUse returns a boolean if a field has been set.

### GetFlow

`func (o *Invitation) GetFlow() string`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *Invitation) GetFlowOk() (*string, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *Invitation) SetFlow(v string)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *Invitation) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### SetFlowNil

`func (o *Invitation) SetFlowNil(b bool)`

 SetFlowNil sets the value for Flow to be an explicit nil

### UnsetFlow
`func (o *Invitation) UnsetFlow()`

UnsetFlow ensures that no value is present for Flow, not even an explicit nil
### GetFlowObj

`func (o *Invitation) GetFlowObj() Flow`

GetFlowObj returns the FlowObj field if non-nil, zero value otherwise.

### GetFlowObjOk

`func (o *Invitation) GetFlowObjOk() (*Flow, bool)`

GetFlowObjOk returns a tuple with the FlowObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowObj

`func (o *Invitation) SetFlowObj(v Flow)`

SetFlowObj sets FlowObj field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



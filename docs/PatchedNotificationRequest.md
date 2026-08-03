# PatchedNotificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hyperlink** | Pointer to **NullableString** |  | [optional] 
**HyperlinkLabel** | Pointer to **NullableString** |  | [optional] 
**Event** | Pointer to [**EventRequest**](EventRequest.md) |  | [optional] 
**Seen** | Pointer to **bool** |  | [optional] 

## Methods

### NewPatchedNotificationRequest

`func NewPatchedNotificationRequest() *PatchedNotificationRequest`

NewPatchedNotificationRequest instantiates a new PatchedNotificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedNotificationRequestWithDefaults

`func NewPatchedNotificationRequestWithDefaults() *PatchedNotificationRequest`

NewPatchedNotificationRequestWithDefaults instantiates a new PatchedNotificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHyperlink

`func (o *PatchedNotificationRequest) GetHyperlink() string`

GetHyperlink returns the Hyperlink field if non-nil, zero value otherwise.

### GetHyperlinkOk

`func (o *PatchedNotificationRequest) GetHyperlinkOk() (*string, bool)`

GetHyperlinkOk returns a tuple with the Hyperlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperlink

`func (o *PatchedNotificationRequest) SetHyperlink(v string)`

SetHyperlink sets Hyperlink field to given value.

### HasHyperlink

`func (o *PatchedNotificationRequest) HasHyperlink() bool`

HasHyperlink returns a boolean if a field has been set.

### SetHyperlinkNil

`func (o *PatchedNotificationRequest) SetHyperlinkNil(b bool)`

 SetHyperlinkNil sets the value for Hyperlink to be an explicit nil

### UnsetHyperlink
`func (o *PatchedNotificationRequest) UnsetHyperlink()`

UnsetHyperlink ensures that no value is present for Hyperlink, not even an explicit nil
### GetHyperlinkLabel

`func (o *PatchedNotificationRequest) GetHyperlinkLabel() string`

GetHyperlinkLabel returns the HyperlinkLabel field if non-nil, zero value otherwise.

### GetHyperlinkLabelOk

`func (o *PatchedNotificationRequest) GetHyperlinkLabelOk() (*string, bool)`

GetHyperlinkLabelOk returns a tuple with the HyperlinkLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperlinkLabel

`func (o *PatchedNotificationRequest) SetHyperlinkLabel(v string)`

SetHyperlinkLabel sets HyperlinkLabel field to given value.

### HasHyperlinkLabel

`func (o *PatchedNotificationRequest) HasHyperlinkLabel() bool`

HasHyperlinkLabel returns a boolean if a field has been set.

### SetHyperlinkLabelNil

`func (o *PatchedNotificationRequest) SetHyperlinkLabelNil(b bool)`

 SetHyperlinkLabelNil sets the value for HyperlinkLabel to be an explicit nil

### UnsetHyperlinkLabel
`func (o *PatchedNotificationRequest) UnsetHyperlinkLabel()`

UnsetHyperlinkLabel ensures that no value is present for HyperlinkLabel, not even an explicit nil
### GetEvent

`func (o *PatchedNotificationRequest) GetEvent() EventRequest`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *PatchedNotificationRequest) GetEventOk() (*EventRequest, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *PatchedNotificationRequest) SetEvent(v EventRequest)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *PatchedNotificationRequest) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetSeen

`func (o *PatchedNotificationRequest) GetSeen() bool`

GetSeen returns the Seen field if non-nil, zero value otherwise.

### GetSeenOk

`func (o *PatchedNotificationRequest) GetSeenOk() (*bool, bool)`

GetSeenOk returns a tuple with the Seen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeen

`func (o *PatchedNotificationRequest) SetSeen(v bool)`

SetSeen sets Seen field to given value.

### HasSeen

`func (o *PatchedNotificationRequest) HasSeen() bool`

HasSeen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



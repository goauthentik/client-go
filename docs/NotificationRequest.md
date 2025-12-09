# NotificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hyperlink** | Pointer to **NullableString** |  | [optional] 
**HyperlinkLabel** | Pointer to **NullableString** |  | [optional] 
**Event** | Pointer to [**EventRequest**](EventRequest.md) |  | [optional] 
**Seen** | Pointer to **bool** |  | [optional] 

## Methods

### NewNotificationRequest

`func NewNotificationRequest() *NotificationRequest`

NewNotificationRequest instantiates a new NotificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRequestWithDefaults

`func NewNotificationRequestWithDefaults() *NotificationRequest`

NewNotificationRequestWithDefaults instantiates a new NotificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHyperlink

`func (o *NotificationRequest) GetHyperlink() string`

GetHyperlink returns the Hyperlink field if non-nil, zero value otherwise.

### GetHyperlinkOk

`func (o *NotificationRequest) GetHyperlinkOk() (*string, bool)`

GetHyperlinkOk returns a tuple with the Hyperlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperlink

`func (o *NotificationRequest) SetHyperlink(v string)`

SetHyperlink sets Hyperlink field to given value.

### HasHyperlink

`func (o *NotificationRequest) HasHyperlink() bool`

HasHyperlink returns a boolean if a field has been set.

### SetHyperlinkNil

`func (o *NotificationRequest) SetHyperlinkNil(b bool)`

 SetHyperlinkNil sets the value for Hyperlink to be an explicit nil

### UnsetHyperlink
`func (o *NotificationRequest) UnsetHyperlink()`

UnsetHyperlink ensures that no value is present for Hyperlink, not even an explicit nil
### GetHyperlinkLabel

`func (o *NotificationRequest) GetHyperlinkLabel() string`

GetHyperlinkLabel returns the HyperlinkLabel field if non-nil, zero value otherwise.

### GetHyperlinkLabelOk

`func (o *NotificationRequest) GetHyperlinkLabelOk() (*string, bool)`

GetHyperlinkLabelOk returns a tuple with the HyperlinkLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperlinkLabel

`func (o *NotificationRequest) SetHyperlinkLabel(v string)`

SetHyperlinkLabel sets HyperlinkLabel field to given value.

### HasHyperlinkLabel

`func (o *NotificationRequest) HasHyperlinkLabel() bool`

HasHyperlinkLabel returns a boolean if a field has been set.

### SetHyperlinkLabelNil

`func (o *NotificationRequest) SetHyperlinkLabelNil(b bool)`

 SetHyperlinkLabelNil sets the value for HyperlinkLabel to be an explicit nil

### UnsetHyperlinkLabel
`func (o *NotificationRequest) UnsetHyperlinkLabel()`

UnsetHyperlinkLabel ensures that no value is present for HyperlinkLabel, not even an explicit nil
### GetEvent

`func (o *NotificationRequest) GetEvent() EventRequest`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *NotificationRequest) GetEventOk() (*EventRequest, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *NotificationRequest) SetEvent(v EventRequest)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *NotificationRequest) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetSeen

`func (o *NotificationRequest) GetSeen() bool`

GetSeen returns the Seen field if non-nil, zero value otherwise.

### GetSeenOk

`func (o *NotificationRequest) GetSeenOk() (*bool, bool)`

GetSeenOk returns a tuple with the Seen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeen

`func (o *NotificationRequest) SetSeen(v bool)`

SetSeen sets Seen field to given value.

### HasSeen

`func (o *NotificationRequest) HasSeen() bool`

HasSeen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



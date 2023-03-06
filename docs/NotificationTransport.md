# NotificationTransport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Mode** | Pointer to [**NotificationTransportModeEnum**](NotificationTransportModeEnum.md) |  | [optional] 
**ModeVerbose** | **string** | Return selected mode with a UI Label | [readonly] 
**WebhookUrl** | Pointer to **string** |  | [optional] 
**WebhookMapping** | Pointer to **NullableString** |  | [optional] 
**SendOnce** | Pointer to **bool** | Only send notification once, for example when sending a webhook into a chat channel. | [optional] 

## Methods

### NewNotificationTransport

`func NewNotificationTransport(pk string, name string, modeVerbose string, ) *NotificationTransport`

NewNotificationTransport instantiates a new NotificationTransport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationTransportWithDefaults

`func NewNotificationTransportWithDefaults() *NotificationTransport`

NewNotificationTransportWithDefaults instantiates a new NotificationTransport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *NotificationTransport) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *NotificationTransport) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *NotificationTransport) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *NotificationTransport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationTransport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationTransport) SetName(v string)`

SetName sets Name field to given value.


### GetMode

`func (o *NotificationTransport) GetMode() NotificationTransportModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *NotificationTransport) GetModeOk() (*NotificationTransportModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *NotificationTransport) SetMode(v NotificationTransportModeEnum)`

SetMode sets Mode field to given value.

### HasMode

`func (o *NotificationTransport) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetModeVerbose

`func (o *NotificationTransport) GetModeVerbose() string`

GetModeVerbose returns the ModeVerbose field if non-nil, zero value otherwise.

### GetModeVerboseOk

`func (o *NotificationTransport) GetModeVerboseOk() (*string, bool)`

GetModeVerboseOk returns a tuple with the ModeVerbose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeVerbose

`func (o *NotificationTransport) SetModeVerbose(v string)`

SetModeVerbose sets ModeVerbose field to given value.


### GetWebhookUrl

`func (o *NotificationTransport) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *NotificationTransport) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *NotificationTransport) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *NotificationTransport) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetWebhookMapping

`func (o *NotificationTransport) GetWebhookMapping() string`

GetWebhookMapping returns the WebhookMapping field if non-nil, zero value otherwise.

### GetWebhookMappingOk

`func (o *NotificationTransport) GetWebhookMappingOk() (*string, bool)`

GetWebhookMappingOk returns a tuple with the WebhookMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookMapping

`func (o *NotificationTransport) SetWebhookMapping(v string)`

SetWebhookMapping sets WebhookMapping field to given value.

### HasWebhookMapping

`func (o *NotificationTransport) HasWebhookMapping() bool`

HasWebhookMapping returns a boolean if a field has been set.

### SetWebhookMappingNil

`func (o *NotificationTransport) SetWebhookMappingNil(b bool)`

 SetWebhookMappingNil sets the value for WebhookMapping to be an explicit nil

### UnsetWebhookMapping
`func (o *NotificationTransport) UnsetWebhookMapping()`

UnsetWebhookMapping ensures that no value is present for WebhookMapping, not even an explicit nil
### GetSendOnce

`func (o *NotificationTransport) GetSendOnce() bool`

GetSendOnce returns the SendOnce field if non-nil, zero value otherwise.

### GetSendOnceOk

`func (o *NotificationTransport) GetSendOnceOk() (*bool, bool)`

GetSendOnceOk returns a tuple with the SendOnce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendOnce

`func (o *NotificationTransport) SetSendOnce(v bool)`

SetSendOnce sets SendOnce field to given value.

### HasSendOnce

`func (o *NotificationTransport) HasSendOnce() bool`

HasSendOnce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



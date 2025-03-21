# NotificationTransportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Mode** | Pointer to [**NotificationTransportModeEnum**](NotificationTransportModeEnum.md) |  | [optional] 
**WebhookUrl** | Pointer to **string** |  | [optional] 
**WebhookMappingBody** | Pointer to **NullableString** | Customize the body of the request. Mapping should return data that is JSON-serializable. | [optional] 
**WebhookMappingHeaders** | Pointer to **NullableString** | Configure additional headers to be sent. Mapping should return a dictionary of key-value pairs | [optional] 
**SendOnce** | Pointer to **bool** | Only send notification once, for example when sending a webhook into a chat channel. | [optional] 

## Methods

### NewNotificationTransportRequest

`func NewNotificationTransportRequest(name string, ) *NotificationTransportRequest`

NewNotificationTransportRequest instantiates a new NotificationTransportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationTransportRequestWithDefaults

`func NewNotificationTransportRequestWithDefaults() *NotificationTransportRequest`

NewNotificationTransportRequestWithDefaults instantiates a new NotificationTransportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NotificationTransportRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationTransportRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationTransportRequest) SetName(v string)`

SetName sets Name field to given value.


### GetMode

`func (o *NotificationTransportRequest) GetMode() NotificationTransportModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *NotificationTransportRequest) GetModeOk() (*NotificationTransportModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *NotificationTransportRequest) SetMode(v NotificationTransportModeEnum)`

SetMode sets Mode field to given value.

### HasMode

`func (o *NotificationTransportRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *NotificationTransportRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *NotificationTransportRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *NotificationTransportRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *NotificationTransportRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetWebhookMappingBody

`func (o *NotificationTransportRequest) GetWebhookMappingBody() string`

GetWebhookMappingBody returns the WebhookMappingBody field if non-nil, zero value otherwise.

### GetWebhookMappingBodyOk

`func (o *NotificationTransportRequest) GetWebhookMappingBodyOk() (*string, bool)`

GetWebhookMappingBodyOk returns a tuple with the WebhookMappingBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookMappingBody

`func (o *NotificationTransportRequest) SetWebhookMappingBody(v string)`

SetWebhookMappingBody sets WebhookMappingBody field to given value.

### HasWebhookMappingBody

`func (o *NotificationTransportRequest) HasWebhookMappingBody() bool`

HasWebhookMappingBody returns a boolean if a field has been set.

### SetWebhookMappingBodyNil

`func (o *NotificationTransportRequest) SetWebhookMappingBodyNil(b bool)`

 SetWebhookMappingBodyNil sets the value for WebhookMappingBody to be an explicit nil

### UnsetWebhookMappingBody
`func (o *NotificationTransportRequest) UnsetWebhookMappingBody()`

UnsetWebhookMappingBody ensures that no value is present for WebhookMappingBody, not even an explicit nil
### GetWebhookMappingHeaders

`func (o *NotificationTransportRequest) GetWebhookMappingHeaders() string`

GetWebhookMappingHeaders returns the WebhookMappingHeaders field if non-nil, zero value otherwise.

### GetWebhookMappingHeadersOk

`func (o *NotificationTransportRequest) GetWebhookMappingHeadersOk() (*string, bool)`

GetWebhookMappingHeadersOk returns a tuple with the WebhookMappingHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookMappingHeaders

`func (o *NotificationTransportRequest) SetWebhookMappingHeaders(v string)`

SetWebhookMappingHeaders sets WebhookMappingHeaders field to given value.

### HasWebhookMappingHeaders

`func (o *NotificationTransportRequest) HasWebhookMappingHeaders() bool`

HasWebhookMappingHeaders returns a boolean if a field has been set.

### SetWebhookMappingHeadersNil

`func (o *NotificationTransportRequest) SetWebhookMappingHeadersNil(b bool)`

 SetWebhookMappingHeadersNil sets the value for WebhookMappingHeaders to be an explicit nil

### UnsetWebhookMappingHeaders
`func (o *NotificationTransportRequest) UnsetWebhookMappingHeaders()`

UnsetWebhookMappingHeaders ensures that no value is present for WebhookMappingHeaders, not even an explicit nil
### GetSendOnce

`func (o *NotificationTransportRequest) GetSendOnce() bool`

GetSendOnce returns the SendOnce field if non-nil, zero value otherwise.

### GetSendOnceOk

`func (o *NotificationTransportRequest) GetSendOnceOk() (*bool, bool)`

GetSendOnceOk returns a tuple with the SendOnce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendOnce

`func (o *NotificationTransportRequest) SetSendOnce(v bool)`

SetSendOnce sets SendOnce field to given value.

### HasSendOnce

`func (o *NotificationTransportRequest) HasSendOnce() bool`

HasSendOnce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



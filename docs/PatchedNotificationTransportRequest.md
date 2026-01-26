# PatchedNotificationTransportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to [**NotificationTransportModeEnum**](NotificationTransportModeEnum.md) |  | [optional] 
**WebhookUrl** | Pointer to **string** |  | [optional] 
**WebhookMappingBody** | Pointer to **NullableString** | Customize the body of the request. Mapping should return data that is JSON-serializable. | [optional] 
**WebhookMappingHeaders** | Pointer to **NullableString** | Configure additional headers to be sent. Mapping should return a dictionary of key-value pairs | [optional] 
**SendOnce** | Pointer to **bool** | Only send notification once, for example when sending a webhook into a chat channel. | [optional] 

## Methods

### NewPatchedNotificationTransportRequest

`func NewPatchedNotificationTransportRequest() *PatchedNotificationTransportRequest`

NewPatchedNotificationTransportRequest instantiates a new PatchedNotificationTransportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedNotificationTransportRequestWithDefaults

`func NewPatchedNotificationTransportRequestWithDefaults() *PatchedNotificationTransportRequest`

NewPatchedNotificationTransportRequestWithDefaults instantiates a new PatchedNotificationTransportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedNotificationTransportRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedNotificationTransportRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedNotificationTransportRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedNotificationTransportRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMode

`func (o *PatchedNotificationTransportRequest) GetMode() NotificationTransportModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PatchedNotificationTransportRequest) GetModeOk() (*NotificationTransportModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PatchedNotificationTransportRequest) SetMode(v NotificationTransportModeEnum)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PatchedNotificationTransportRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *PatchedNotificationTransportRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *PatchedNotificationTransportRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *PatchedNotificationTransportRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *PatchedNotificationTransportRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetWebhookMappingBody

`func (o *PatchedNotificationTransportRequest) GetWebhookMappingBody() string`

GetWebhookMappingBody returns the WebhookMappingBody field if non-nil, zero value otherwise.

### GetWebhookMappingBodyOk

`func (o *PatchedNotificationTransportRequest) GetWebhookMappingBodyOk() (*string, bool)`

GetWebhookMappingBodyOk returns a tuple with the WebhookMappingBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookMappingBody

`func (o *PatchedNotificationTransportRequest) SetWebhookMappingBody(v string)`

SetWebhookMappingBody sets WebhookMappingBody field to given value.

### HasWebhookMappingBody

`func (o *PatchedNotificationTransportRequest) HasWebhookMappingBody() bool`

HasWebhookMappingBody returns a boolean if a field has been set.

### SetWebhookMappingBodyNil

`func (o *PatchedNotificationTransportRequest) SetWebhookMappingBodyNil(b bool)`

 SetWebhookMappingBodyNil sets the value for WebhookMappingBody to be an explicit nil

### UnsetWebhookMappingBody
`func (o *PatchedNotificationTransportRequest) UnsetWebhookMappingBody()`

UnsetWebhookMappingBody ensures that no value is present for WebhookMappingBody, not even an explicit nil
### GetWebhookMappingHeaders

`func (o *PatchedNotificationTransportRequest) GetWebhookMappingHeaders() string`

GetWebhookMappingHeaders returns the WebhookMappingHeaders field if non-nil, zero value otherwise.

### GetWebhookMappingHeadersOk

`func (o *PatchedNotificationTransportRequest) GetWebhookMappingHeadersOk() (*string, bool)`

GetWebhookMappingHeadersOk returns a tuple with the WebhookMappingHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookMappingHeaders

`func (o *PatchedNotificationTransportRequest) SetWebhookMappingHeaders(v string)`

SetWebhookMappingHeaders sets WebhookMappingHeaders field to given value.

### HasWebhookMappingHeaders

`func (o *PatchedNotificationTransportRequest) HasWebhookMappingHeaders() bool`

HasWebhookMappingHeaders returns a boolean if a field has been set.

### SetWebhookMappingHeadersNil

`func (o *PatchedNotificationTransportRequest) SetWebhookMappingHeadersNil(b bool)`

 SetWebhookMappingHeadersNil sets the value for WebhookMappingHeaders to be an explicit nil

### UnsetWebhookMappingHeaders
`func (o *PatchedNotificationTransportRequest) UnsetWebhookMappingHeaders()`

UnsetWebhookMappingHeaders ensures that no value is present for WebhookMappingHeaders, not even an explicit nil
### GetSendOnce

`func (o *PatchedNotificationTransportRequest) GetSendOnce() bool`

GetSendOnce returns the SendOnce field if non-nil, zero value otherwise.

### GetSendOnceOk

`func (o *PatchedNotificationTransportRequest) GetSendOnceOk() (*bool, bool)`

GetSendOnceOk returns a tuple with the SendOnce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendOnce

`func (o *PatchedNotificationTransportRequest) SetSendOnce(v bool)`

SetSendOnce sets SendOnce field to given value.

### HasSendOnce

`func (o *PatchedNotificationTransportRequest) HasSendOnce() bool`

HasSendOnce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



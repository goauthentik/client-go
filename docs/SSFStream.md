# SSFStream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Provider** | **int32** |  | 
**ProviderObj** | [**SSFProvider**](SSFProvider.md) |  | [readonly] 
**DeliveryMethod** | [**DeliveryMethodEnum**](DeliveryMethodEnum.md) |  | 
**EndpointUrl** | Pointer to **NullableString** |  | [optional] 
**EventsRequested** | Pointer to [**[]EventsRequestedEnum**](EventsRequestedEnum.md) |  | [optional] 
**Format** | **string** |  | 
**Aud** | Pointer to **[]string** |  | [optional] 
**Iss** | **string** |  | 

## Methods

### NewSSFStream

`func NewSSFStream(pk string, provider int32, providerObj SSFProvider, deliveryMethod DeliveryMethodEnum, format string, iss string, ) *SSFStream`

NewSSFStream instantiates a new SSFStream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSFStreamWithDefaults

`func NewSSFStreamWithDefaults() *SSFStream`

NewSSFStreamWithDefaults instantiates a new SSFStream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *SSFStream) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *SSFStream) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *SSFStream) SetPk(v string)`

SetPk sets Pk field to given value.


### GetProvider

`func (o *SSFStream) GetProvider() int32`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SSFStream) GetProviderOk() (*int32, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SSFStream) SetProvider(v int32)`

SetProvider sets Provider field to given value.


### GetProviderObj

`func (o *SSFStream) GetProviderObj() SSFProvider`

GetProviderObj returns the ProviderObj field if non-nil, zero value otherwise.

### GetProviderObjOk

`func (o *SSFStream) GetProviderObjOk() (*SSFProvider, bool)`

GetProviderObjOk returns a tuple with the ProviderObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderObj

`func (o *SSFStream) SetProviderObj(v SSFProvider)`

SetProviderObj sets ProviderObj field to given value.


### GetDeliveryMethod

`func (o *SSFStream) GetDeliveryMethod() DeliveryMethodEnum`

GetDeliveryMethod returns the DeliveryMethod field if non-nil, zero value otherwise.

### GetDeliveryMethodOk

`func (o *SSFStream) GetDeliveryMethodOk() (*DeliveryMethodEnum, bool)`

GetDeliveryMethodOk returns a tuple with the DeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethod

`func (o *SSFStream) SetDeliveryMethod(v DeliveryMethodEnum)`

SetDeliveryMethod sets DeliveryMethod field to given value.


### GetEndpointUrl

`func (o *SSFStream) GetEndpointUrl() string`

GetEndpointUrl returns the EndpointUrl field if non-nil, zero value otherwise.

### GetEndpointUrlOk

`func (o *SSFStream) GetEndpointUrlOk() (*string, bool)`

GetEndpointUrlOk returns a tuple with the EndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointUrl

`func (o *SSFStream) SetEndpointUrl(v string)`

SetEndpointUrl sets EndpointUrl field to given value.

### HasEndpointUrl

`func (o *SSFStream) HasEndpointUrl() bool`

HasEndpointUrl returns a boolean if a field has been set.

### SetEndpointUrlNil

`func (o *SSFStream) SetEndpointUrlNil(b bool)`

 SetEndpointUrlNil sets the value for EndpointUrl to be an explicit nil

### UnsetEndpointUrl
`func (o *SSFStream) UnsetEndpointUrl()`

UnsetEndpointUrl ensures that no value is present for EndpointUrl, not even an explicit nil
### GetEventsRequested

`func (o *SSFStream) GetEventsRequested() []EventsRequestedEnum`

GetEventsRequested returns the EventsRequested field if non-nil, zero value otherwise.

### GetEventsRequestedOk

`func (o *SSFStream) GetEventsRequestedOk() (*[]EventsRequestedEnum, bool)`

GetEventsRequestedOk returns a tuple with the EventsRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsRequested

`func (o *SSFStream) SetEventsRequested(v []EventsRequestedEnum)`

SetEventsRequested sets EventsRequested field to given value.

### HasEventsRequested

`func (o *SSFStream) HasEventsRequested() bool`

HasEventsRequested returns a boolean if a field has been set.

### GetFormat

`func (o *SSFStream) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *SSFStream) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *SSFStream) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetAud

`func (o *SSFStream) GetAud() []string`

GetAud returns the Aud field if non-nil, zero value otherwise.

### GetAudOk

`func (o *SSFStream) GetAudOk() (*[]string, bool)`

GetAudOk returns a tuple with the Aud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAud

`func (o *SSFStream) SetAud(v []string)`

SetAud sets Aud field to given value.

### HasAud

`func (o *SSFStream) HasAud() bool`

HasAud returns a boolean if a field has been set.

### GetIss

`func (o *SSFStream) GetIss() string`

GetIss returns the Iss field if non-nil, zero value otherwise.

### GetIssOk

`func (o *SSFStream) GetIssOk() (*string, bool)`

GetIssOk returns a tuple with the Iss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIss

`func (o *SSFStream) SetIss(v string)`

SetIss sets Iss field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



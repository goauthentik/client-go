# TelegramLoginChallenge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowInfo** | Pointer to [**ContextualFlowInfo**](ContextualFlowInfo.md) |  | [optional] 
**Component** | Pointer to **string** |  | [optional] [default to "ak-source-telegram"]
**ResponseErrors** | Pointer to [**map[string][]ErrorDetail**](array.md) |  | [optional] 
**BotUsername** | **string** | Telegram bot username | 
**RequestMessageAccess** | **bool** |  | 

## Methods

### NewTelegramLoginChallenge

`func NewTelegramLoginChallenge(botUsername string, requestMessageAccess bool, ) *TelegramLoginChallenge`

NewTelegramLoginChallenge instantiates a new TelegramLoginChallenge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelegramLoginChallengeWithDefaults

`func NewTelegramLoginChallengeWithDefaults() *TelegramLoginChallenge`

NewTelegramLoginChallengeWithDefaults instantiates a new TelegramLoginChallenge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowInfo

`func (o *TelegramLoginChallenge) GetFlowInfo() ContextualFlowInfo`

GetFlowInfo returns the FlowInfo field if non-nil, zero value otherwise.

### GetFlowInfoOk

`func (o *TelegramLoginChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool)`

GetFlowInfoOk returns a tuple with the FlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInfo

`func (o *TelegramLoginChallenge) SetFlowInfo(v ContextualFlowInfo)`

SetFlowInfo sets FlowInfo field to given value.

### HasFlowInfo

`func (o *TelegramLoginChallenge) HasFlowInfo() bool`

HasFlowInfo returns a boolean if a field has been set.

### GetComponent

`func (o *TelegramLoginChallenge) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *TelegramLoginChallenge) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *TelegramLoginChallenge) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *TelegramLoginChallenge) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetResponseErrors

`func (o *TelegramLoginChallenge) GetResponseErrors() map[string][]ErrorDetail`

GetResponseErrors returns the ResponseErrors field if non-nil, zero value otherwise.

### GetResponseErrorsOk

`func (o *TelegramLoginChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool)`

GetResponseErrorsOk returns a tuple with the ResponseErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseErrors

`func (o *TelegramLoginChallenge) SetResponseErrors(v map[string][]ErrorDetail)`

SetResponseErrors sets ResponseErrors field to given value.

### HasResponseErrors

`func (o *TelegramLoginChallenge) HasResponseErrors() bool`

HasResponseErrors returns a boolean if a field has been set.

### GetBotUsername

`func (o *TelegramLoginChallenge) GetBotUsername() string`

GetBotUsername returns the BotUsername field if non-nil, zero value otherwise.

### GetBotUsernameOk

`func (o *TelegramLoginChallenge) GetBotUsernameOk() (*string, bool)`

GetBotUsernameOk returns a tuple with the BotUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotUsername

`func (o *TelegramLoginChallenge) SetBotUsername(v string)`

SetBotUsername sets BotUsername field to given value.


### GetRequestMessageAccess

`func (o *TelegramLoginChallenge) GetRequestMessageAccess() bool`

GetRequestMessageAccess returns the RequestMessageAccess field if non-nil, zero value otherwise.

### GetRequestMessageAccessOk

`func (o *TelegramLoginChallenge) GetRequestMessageAccessOk() (*bool, bool)`

GetRequestMessageAccessOk returns a tuple with the RequestMessageAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMessageAccess

`func (o *TelegramLoginChallenge) SetRequestMessageAccess(v bool)`

SetRequestMessageAccess sets RequestMessageAccess field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



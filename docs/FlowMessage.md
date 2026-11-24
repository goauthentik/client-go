# FlowMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | [**FlowMessageLevelEnum**](FlowMessageLevelEnum.md) |  | 
**Message** | **string** |  | 

## Methods

### NewFlowMessage

`func NewFlowMessage(level FlowMessageLevelEnum, message string, ) *FlowMessage`

NewFlowMessage instantiates a new FlowMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowMessageWithDefaults

`func NewFlowMessageWithDefaults() *FlowMessage`

NewFlowMessageWithDefaults instantiates a new FlowMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *FlowMessage) GetLevel() FlowMessageLevelEnum`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *FlowMessage) GetLevelOk() (*FlowMessageLevelEnum, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *FlowMessage) SetLevel(v FlowMessageLevelEnum)`

SetLevel sets Level field to given value.


### GetMessage

`func (o *FlowMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FlowMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FlowMessage) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



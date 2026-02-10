# LifecycleRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ContentType** | [**ContentTypeEnum**](ContentTypeEnum.md) |  | 
**ObjectId** | Pointer to **NullableString** |  | [optional] 
**Interval** | Pointer to **string** |  | [optional] 
**GracePeriod** | Pointer to **string** |  | [optional] 
**ReviewerGroups** | Pointer to **[]string** |  | [optional] 
**MinReviewers** | Pointer to **int32** |  | [optional] 
**MinReviewersIsPerGroup** | Pointer to **bool** |  | [optional] 
**Reviewers** | **[]string** |  | 
**NotificationTransports** | Pointer to **[]string** | Select which transports should be used to notify the reviewers. If none are selected, the notification will only be shown in the authentik UI. | [optional] 

## Methods

### NewLifecycleRuleRequest

`func NewLifecycleRuleRequest(name string, contentType ContentTypeEnum, reviewers []string, ) *LifecycleRuleRequest`

NewLifecycleRuleRequest instantiates a new LifecycleRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLifecycleRuleRequestWithDefaults

`func NewLifecycleRuleRequestWithDefaults() *LifecycleRuleRequest`

NewLifecycleRuleRequestWithDefaults instantiates a new LifecycleRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LifecycleRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LifecycleRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LifecycleRuleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetContentType

`func (o *LifecycleRuleRequest) GetContentType() ContentTypeEnum`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *LifecycleRuleRequest) GetContentTypeOk() (*ContentTypeEnum, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *LifecycleRuleRequest) SetContentType(v ContentTypeEnum)`

SetContentType sets ContentType field to given value.


### GetObjectId

`func (o *LifecycleRuleRequest) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *LifecycleRuleRequest) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *LifecycleRuleRequest) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *LifecycleRuleRequest) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### SetObjectIdNil

`func (o *LifecycleRuleRequest) SetObjectIdNil(b bool)`

 SetObjectIdNil sets the value for ObjectId to be an explicit nil

### UnsetObjectId
`func (o *LifecycleRuleRequest) UnsetObjectId()`

UnsetObjectId ensures that no value is present for ObjectId, not even an explicit nil
### GetInterval

`func (o *LifecycleRuleRequest) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *LifecycleRuleRequest) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *LifecycleRuleRequest) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *LifecycleRuleRequest) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetGracePeriod

`func (o *LifecycleRuleRequest) GetGracePeriod() string`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *LifecycleRuleRequest) GetGracePeriodOk() (*string, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *LifecycleRuleRequest) SetGracePeriod(v string)`

SetGracePeriod sets GracePeriod field to given value.

### HasGracePeriod

`func (o *LifecycleRuleRequest) HasGracePeriod() bool`

HasGracePeriod returns a boolean if a field has been set.

### GetReviewerGroups

`func (o *LifecycleRuleRequest) GetReviewerGroups() []string`

GetReviewerGroups returns the ReviewerGroups field if non-nil, zero value otherwise.

### GetReviewerGroupsOk

`func (o *LifecycleRuleRequest) GetReviewerGroupsOk() (*[]string, bool)`

GetReviewerGroupsOk returns a tuple with the ReviewerGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerGroups

`func (o *LifecycleRuleRequest) SetReviewerGroups(v []string)`

SetReviewerGroups sets ReviewerGroups field to given value.

### HasReviewerGroups

`func (o *LifecycleRuleRequest) HasReviewerGroups() bool`

HasReviewerGroups returns a boolean if a field has been set.

### GetMinReviewers

`func (o *LifecycleRuleRequest) GetMinReviewers() int32`

GetMinReviewers returns the MinReviewers field if non-nil, zero value otherwise.

### GetMinReviewersOk

`func (o *LifecycleRuleRequest) GetMinReviewersOk() (*int32, bool)`

GetMinReviewersOk returns a tuple with the MinReviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReviewers

`func (o *LifecycleRuleRequest) SetMinReviewers(v int32)`

SetMinReviewers sets MinReviewers field to given value.

### HasMinReviewers

`func (o *LifecycleRuleRequest) HasMinReviewers() bool`

HasMinReviewers returns a boolean if a field has been set.

### GetMinReviewersIsPerGroup

`func (o *LifecycleRuleRequest) GetMinReviewersIsPerGroup() bool`

GetMinReviewersIsPerGroup returns the MinReviewersIsPerGroup field if non-nil, zero value otherwise.

### GetMinReviewersIsPerGroupOk

`func (o *LifecycleRuleRequest) GetMinReviewersIsPerGroupOk() (*bool, bool)`

GetMinReviewersIsPerGroupOk returns a tuple with the MinReviewersIsPerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReviewersIsPerGroup

`func (o *LifecycleRuleRequest) SetMinReviewersIsPerGroup(v bool)`

SetMinReviewersIsPerGroup sets MinReviewersIsPerGroup field to given value.

### HasMinReviewersIsPerGroup

`func (o *LifecycleRuleRequest) HasMinReviewersIsPerGroup() bool`

HasMinReviewersIsPerGroup returns a boolean if a field has been set.

### GetReviewers

`func (o *LifecycleRuleRequest) GetReviewers() []string`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *LifecycleRuleRequest) GetReviewersOk() (*[]string, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *LifecycleRuleRequest) SetReviewers(v []string)`

SetReviewers sets Reviewers field to given value.


### GetNotificationTransports

`func (o *LifecycleRuleRequest) GetNotificationTransports() []string`

GetNotificationTransports returns the NotificationTransports field if non-nil, zero value otherwise.

### GetNotificationTransportsOk

`func (o *LifecycleRuleRequest) GetNotificationTransportsOk() (*[]string, bool)`

GetNotificationTransportsOk returns a tuple with the NotificationTransports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTransports

`func (o *LifecycleRuleRequest) SetNotificationTransports(v []string)`

SetNotificationTransports sets NotificationTransports field to given value.

### HasNotificationTransports

`func (o *LifecycleRuleRequest) HasNotificationTransports() bool`

HasNotificationTransports returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



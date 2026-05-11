# PatchedLifecycleRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to [**ContentTypeEnum**](ContentTypeEnum.md) |  | [optional] 
**ObjectId** | Pointer to **NullableString** |  | [optional] 
**Interval** | Pointer to **string** |  | [optional] 
**GracePeriod** | Pointer to **string** |  | [optional] 
**ReviewerGroups** | Pointer to **[]string** |  | [optional] 
**MinReviewers** | Pointer to **int32** |  | [optional] 
**MinReviewersIsPerGroup** | Pointer to **bool** |  | [optional] 
**Reviewers** | Pointer to **[]string** |  | [optional] 
**NotificationTransports** | Pointer to **[]string** | Select which transports should be used to notify the reviewers. If none are selected, the notification will only be shown in the authentik UI. | [optional] 

## Methods

### NewPatchedLifecycleRuleRequest

`func NewPatchedLifecycleRuleRequest() *PatchedLifecycleRuleRequest`

NewPatchedLifecycleRuleRequest instantiates a new PatchedLifecycleRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedLifecycleRuleRequestWithDefaults

`func NewPatchedLifecycleRuleRequestWithDefaults() *PatchedLifecycleRuleRequest`

NewPatchedLifecycleRuleRequestWithDefaults instantiates a new PatchedLifecycleRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedLifecycleRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedLifecycleRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedLifecycleRuleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedLifecycleRuleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetContentType

`func (o *PatchedLifecycleRuleRequest) GetContentType() ContentTypeEnum`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *PatchedLifecycleRuleRequest) GetContentTypeOk() (*ContentTypeEnum, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *PatchedLifecycleRuleRequest) SetContentType(v ContentTypeEnum)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *PatchedLifecycleRuleRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetObjectId

`func (o *PatchedLifecycleRuleRequest) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *PatchedLifecycleRuleRequest) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *PatchedLifecycleRuleRequest) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *PatchedLifecycleRuleRequest) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### SetObjectIdNil

`func (o *PatchedLifecycleRuleRequest) SetObjectIdNil(b bool)`

 SetObjectIdNil sets the value for ObjectId to be an explicit nil

### UnsetObjectId
`func (o *PatchedLifecycleRuleRequest) UnsetObjectId()`

UnsetObjectId ensures that no value is present for ObjectId, not even an explicit nil
### GetInterval

`func (o *PatchedLifecycleRuleRequest) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *PatchedLifecycleRuleRequest) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *PatchedLifecycleRuleRequest) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *PatchedLifecycleRuleRequest) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetGracePeriod

`func (o *PatchedLifecycleRuleRequest) GetGracePeriod() string`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *PatchedLifecycleRuleRequest) GetGracePeriodOk() (*string, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *PatchedLifecycleRuleRequest) SetGracePeriod(v string)`

SetGracePeriod sets GracePeriod field to given value.

### HasGracePeriod

`func (o *PatchedLifecycleRuleRequest) HasGracePeriod() bool`

HasGracePeriod returns a boolean if a field has been set.

### GetReviewerGroups

`func (o *PatchedLifecycleRuleRequest) GetReviewerGroups() []string`

GetReviewerGroups returns the ReviewerGroups field if non-nil, zero value otherwise.

### GetReviewerGroupsOk

`func (o *PatchedLifecycleRuleRequest) GetReviewerGroupsOk() (*[]string, bool)`

GetReviewerGroupsOk returns a tuple with the ReviewerGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerGroups

`func (o *PatchedLifecycleRuleRequest) SetReviewerGroups(v []string)`

SetReviewerGroups sets ReviewerGroups field to given value.

### HasReviewerGroups

`func (o *PatchedLifecycleRuleRequest) HasReviewerGroups() bool`

HasReviewerGroups returns a boolean if a field has been set.

### GetMinReviewers

`func (o *PatchedLifecycleRuleRequest) GetMinReviewers() int32`

GetMinReviewers returns the MinReviewers field if non-nil, zero value otherwise.

### GetMinReviewersOk

`func (o *PatchedLifecycleRuleRequest) GetMinReviewersOk() (*int32, bool)`

GetMinReviewersOk returns a tuple with the MinReviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReviewers

`func (o *PatchedLifecycleRuleRequest) SetMinReviewers(v int32)`

SetMinReviewers sets MinReviewers field to given value.

### HasMinReviewers

`func (o *PatchedLifecycleRuleRequest) HasMinReviewers() bool`

HasMinReviewers returns a boolean if a field has been set.

### GetMinReviewersIsPerGroup

`func (o *PatchedLifecycleRuleRequest) GetMinReviewersIsPerGroup() bool`

GetMinReviewersIsPerGroup returns the MinReviewersIsPerGroup field if non-nil, zero value otherwise.

### GetMinReviewersIsPerGroupOk

`func (o *PatchedLifecycleRuleRequest) GetMinReviewersIsPerGroupOk() (*bool, bool)`

GetMinReviewersIsPerGroupOk returns a tuple with the MinReviewersIsPerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReviewersIsPerGroup

`func (o *PatchedLifecycleRuleRequest) SetMinReviewersIsPerGroup(v bool)`

SetMinReviewersIsPerGroup sets MinReviewersIsPerGroup field to given value.

### HasMinReviewersIsPerGroup

`func (o *PatchedLifecycleRuleRequest) HasMinReviewersIsPerGroup() bool`

HasMinReviewersIsPerGroup returns a boolean if a field has been set.

### GetReviewers

`func (o *PatchedLifecycleRuleRequest) GetReviewers() []string`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *PatchedLifecycleRuleRequest) GetReviewersOk() (*[]string, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *PatchedLifecycleRuleRequest) SetReviewers(v []string)`

SetReviewers sets Reviewers field to given value.

### HasReviewers

`func (o *PatchedLifecycleRuleRequest) HasReviewers() bool`

HasReviewers returns a boolean if a field has been set.

### GetNotificationTransports

`func (o *PatchedLifecycleRuleRequest) GetNotificationTransports() []string`

GetNotificationTransports returns the NotificationTransports field if non-nil, zero value otherwise.

### GetNotificationTransportsOk

`func (o *PatchedLifecycleRuleRequest) GetNotificationTransportsOk() (*[]string, bool)`

GetNotificationTransportsOk returns a tuple with the NotificationTransports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTransports

`func (o *PatchedLifecycleRuleRequest) SetNotificationTransports(v []string)`

SetNotificationTransports sets NotificationTransports field to given value.

### HasNotificationTransports

`func (o *PatchedLifecycleRuleRequest) HasNotificationTransports() bool`

HasNotificationTransports returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



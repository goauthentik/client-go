# LifecycleRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Name** | **string** |  | 
**ContentType** | [**ContentTypeEnum**](ContentTypeEnum.md) |  | 
**ObjectId** | Pointer to **NullableString** |  | [optional] 
**Interval** | Pointer to **string** |  | [optional] 
**GracePeriod** | Pointer to **string** |  | [optional] 
**ReviewerGroups** | Pointer to **[]string** |  | [optional] 
**ReviewerGroupsObj** | [**[]ReviewerGroup**](ReviewerGroup.md) |  | [readonly] 
**MinReviewers** | Pointer to **int32** |  | [optional] 
**MinReviewersIsPerGroup** | Pointer to **bool** |  | [optional] 
**Reviewers** | **[]string** |  | 
**ReviewersObj** | [**[]ReviewerUser**](ReviewerUser.md) |  | [readonly] 
**NotificationTransports** | Pointer to **[]string** | Select which transports should be used to notify the reviewers. If none are selected, the notification will only be shown in the authentik UI. | [optional] 
**TargetVerbose** | **string** |  | [readonly] 

## Methods

### NewLifecycleRule

`func NewLifecycleRule(id string, name string, contentType ContentTypeEnum, reviewerGroupsObj []ReviewerGroup, reviewers []string, reviewersObj []ReviewerUser, targetVerbose string, ) *LifecycleRule`

NewLifecycleRule instantiates a new LifecycleRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLifecycleRuleWithDefaults

`func NewLifecycleRuleWithDefaults() *LifecycleRule`

NewLifecycleRuleWithDefaults instantiates a new LifecycleRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LifecycleRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LifecycleRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LifecycleRule) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *LifecycleRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LifecycleRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LifecycleRule) SetName(v string)`

SetName sets Name field to given value.


### GetContentType

`func (o *LifecycleRule) GetContentType() ContentTypeEnum`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *LifecycleRule) GetContentTypeOk() (*ContentTypeEnum, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *LifecycleRule) SetContentType(v ContentTypeEnum)`

SetContentType sets ContentType field to given value.


### GetObjectId

`func (o *LifecycleRule) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *LifecycleRule) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *LifecycleRule) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *LifecycleRule) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### SetObjectIdNil

`func (o *LifecycleRule) SetObjectIdNil(b bool)`

 SetObjectIdNil sets the value for ObjectId to be an explicit nil

### UnsetObjectId
`func (o *LifecycleRule) UnsetObjectId()`

UnsetObjectId ensures that no value is present for ObjectId, not even an explicit nil
### GetInterval

`func (o *LifecycleRule) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *LifecycleRule) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *LifecycleRule) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *LifecycleRule) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetGracePeriod

`func (o *LifecycleRule) GetGracePeriod() string`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *LifecycleRule) GetGracePeriodOk() (*string, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *LifecycleRule) SetGracePeriod(v string)`

SetGracePeriod sets GracePeriod field to given value.

### HasGracePeriod

`func (o *LifecycleRule) HasGracePeriod() bool`

HasGracePeriod returns a boolean if a field has been set.

### GetReviewerGroups

`func (o *LifecycleRule) GetReviewerGroups() []string`

GetReviewerGroups returns the ReviewerGroups field if non-nil, zero value otherwise.

### GetReviewerGroupsOk

`func (o *LifecycleRule) GetReviewerGroupsOk() (*[]string, bool)`

GetReviewerGroupsOk returns a tuple with the ReviewerGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerGroups

`func (o *LifecycleRule) SetReviewerGroups(v []string)`

SetReviewerGroups sets ReviewerGroups field to given value.

### HasReviewerGroups

`func (o *LifecycleRule) HasReviewerGroups() bool`

HasReviewerGroups returns a boolean if a field has been set.

### GetReviewerGroupsObj

`func (o *LifecycleRule) GetReviewerGroupsObj() []ReviewerGroup`

GetReviewerGroupsObj returns the ReviewerGroupsObj field if non-nil, zero value otherwise.

### GetReviewerGroupsObjOk

`func (o *LifecycleRule) GetReviewerGroupsObjOk() (*[]ReviewerGroup, bool)`

GetReviewerGroupsObjOk returns a tuple with the ReviewerGroupsObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerGroupsObj

`func (o *LifecycleRule) SetReviewerGroupsObj(v []ReviewerGroup)`

SetReviewerGroupsObj sets ReviewerGroupsObj field to given value.


### GetMinReviewers

`func (o *LifecycleRule) GetMinReviewers() int32`

GetMinReviewers returns the MinReviewers field if non-nil, zero value otherwise.

### GetMinReviewersOk

`func (o *LifecycleRule) GetMinReviewersOk() (*int32, bool)`

GetMinReviewersOk returns a tuple with the MinReviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReviewers

`func (o *LifecycleRule) SetMinReviewers(v int32)`

SetMinReviewers sets MinReviewers field to given value.

### HasMinReviewers

`func (o *LifecycleRule) HasMinReviewers() bool`

HasMinReviewers returns a boolean if a field has been set.

### GetMinReviewersIsPerGroup

`func (o *LifecycleRule) GetMinReviewersIsPerGroup() bool`

GetMinReviewersIsPerGroup returns the MinReviewersIsPerGroup field if non-nil, zero value otherwise.

### GetMinReviewersIsPerGroupOk

`func (o *LifecycleRule) GetMinReviewersIsPerGroupOk() (*bool, bool)`

GetMinReviewersIsPerGroupOk returns a tuple with the MinReviewersIsPerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReviewersIsPerGroup

`func (o *LifecycleRule) SetMinReviewersIsPerGroup(v bool)`

SetMinReviewersIsPerGroup sets MinReviewersIsPerGroup field to given value.

### HasMinReviewersIsPerGroup

`func (o *LifecycleRule) HasMinReviewersIsPerGroup() bool`

HasMinReviewersIsPerGroup returns a boolean if a field has been set.

### GetReviewers

`func (o *LifecycleRule) GetReviewers() []string`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *LifecycleRule) GetReviewersOk() (*[]string, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *LifecycleRule) SetReviewers(v []string)`

SetReviewers sets Reviewers field to given value.


### GetReviewersObj

`func (o *LifecycleRule) GetReviewersObj() []ReviewerUser`

GetReviewersObj returns the ReviewersObj field if non-nil, zero value otherwise.

### GetReviewersObjOk

`func (o *LifecycleRule) GetReviewersObjOk() (*[]ReviewerUser, bool)`

GetReviewersObjOk returns a tuple with the ReviewersObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewersObj

`func (o *LifecycleRule) SetReviewersObj(v []ReviewerUser)`

SetReviewersObj sets ReviewersObj field to given value.


### GetNotificationTransports

`func (o *LifecycleRule) GetNotificationTransports() []string`

GetNotificationTransports returns the NotificationTransports field if non-nil, zero value otherwise.

### GetNotificationTransportsOk

`func (o *LifecycleRule) GetNotificationTransportsOk() (*[]string, bool)`

GetNotificationTransportsOk returns a tuple with the NotificationTransports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTransports

`func (o *LifecycleRule) SetNotificationTransports(v []string)`

SetNotificationTransports sets NotificationTransports field to given value.

### HasNotificationTransports

`func (o *LifecycleRule) HasNotificationTransports() bool`

HasNotificationTransports returns a boolean if a field has been set.

### GetTargetVerbose

`func (o *LifecycleRule) GetTargetVerbose() string`

GetTargetVerbose returns the TargetVerbose field if non-nil, zero value otherwise.

### GetTargetVerboseOk

`func (o *LifecycleRule) GetTargetVerboseOk() (*string, bool)`

GetTargetVerboseOk returns a tuple with the TargetVerbose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVerbose

`func (o *LifecycleRule) SetTargetVerbose(v string)`

SetTargetVerbose sets TargetVerbose field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



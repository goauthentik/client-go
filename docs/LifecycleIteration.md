# LifecycleIteration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ContentType** | [**ContentTypeEnum**](ContentTypeEnum.md) |  | 
**ObjectId** | **string** |  | [readonly] 
**ObjectVerbose** | **string** |  | [readonly] 
**ObjectAdminUrl** | **string** |  | [readonly] 
**State** | [**LifecycleIterationStateEnum**](LifecycleIterationStateEnum.md) |  | [readonly] 
**OpenedOn** | **string** |  | [readonly] 
**GracePeriodEnd** | **string** |  | [readonly] 
**NextReviewDate** | **string** |  | [readonly] 
**Reviews** | [**[]Review**](Review.md) |  | [readonly] 
**UserCanReview** | **bool** |  | [readonly] 
**ReviewerGroups** | [**[]ReviewerGroup**](ReviewerGroup.md) |  | [readonly] 
**MinReviewers** | **int32** |  | [readonly] 
**Reviewers** | [**[]ReviewerUser**](ReviewerUser.md) |  | [readonly] 

## Methods

### NewLifecycleIteration

`func NewLifecycleIteration(id string, contentType ContentTypeEnum, objectId string, objectVerbose string, objectAdminUrl string, state LifecycleIterationStateEnum, openedOn string, gracePeriodEnd string, nextReviewDate string, reviews []Review, userCanReview bool, reviewerGroups []ReviewerGroup, minReviewers int32, reviewers []ReviewerUser, ) *LifecycleIteration`

NewLifecycleIteration instantiates a new LifecycleIteration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLifecycleIterationWithDefaults

`func NewLifecycleIterationWithDefaults() *LifecycleIteration`

NewLifecycleIterationWithDefaults instantiates a new LifecycleIteration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LifecycleIteration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LifecycleIteration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LifecycleIteration) SetId(v string)`

SetId sets Id field to given value.


### GetContentType

`func (o *LifecycleIteration) GetContentType() ContentTypeEnum`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *LifecycleIteration) GetContentTypeOk() (*ContentTypeEnum, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *LifecycleIteration) SetContentType(v ContentTypeEnum)`

SetContentType sets ContentType field to given value.


### GetObjectId

`func (o *LifecycleIteration) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *LifecycleIteration) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *LifecycleIteration) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetObjectVerbose

`func (o *LifecycleIteration) GetObjectVerbose() string`

GetObjectVerbose returns the ObjectVerbose field if non-nil, zero value otherwise.

### GetObjectVerboseOk

`func (o *LifecycleIteration) GetObjectVerboseOk() (*string, bool)`

GetObjectVerboseOk returns a tuple with the ObjectVerbose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectVerbose

`func (o *LifecycleIteration) SetObjectVerbose(v string)`

SetObjectVerbose sets ObjectVerbose field to given value.


### GetObjectAdminUrl

`func (o *LifecycleIteration) GetObjectAdminUrl() string`

GetObjectAdminUrl returns the ObjectAdminUrl field if non-nil, zero value otherwise.

### GetObjectAdminUrlOk

`func (o *LifecycleIteration) GetObjectAdminUrlOk() (*string, bool)`

GetObjectAdminUrlOk returns a tuple with the ObjectAdminUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectAdminUrl

`func (o *LifecycleIteration) SetObjectAdminUrl(v string)`

SetObjectAdminUrl sets ObjectAdminUrl field to given value.


### GetState

`func (o *LifecycleIteration) GetState() LifecycleIterationStateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LifecycleIteration) GetStateOk() (*LifecycleIterationStateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LifecycleIteration) SetState(v LifecycleIterationStateEnum)`

SetState sets State field to given value.


### GetOpenedOn

`func (o *LifecycleIteration) GetOpenedOn() string`

GetOpenedOn returns the OpenedOn field if non-nil, zero value otherwise.

### GetOpenedOnOk

`func (o *LifecycleIteration) GetOpenedOnOk() (*string, bool)`

GetOpenedOnOk returns a tuple with the OpenedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenedOn

`func (o *LifecycleIteration) SetOpenedOn(v string)`

SetOpenedOn sets OpenedOn field to given value.


### GetGracePeriodEnd

`func (o *LifecycleIteration) GetGracePeriodEnd() string`

GetGracePeriodEnd returns the GracePeriodEnd field if non-nil, zero value otherwise.

### GetGracePeriodEndOk

`func (o *LifecycleIteration) GetGracePeriodEndOk() (*string, bool)`

GetGracePeriodEndOk returns a tuple with the GracePeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodEnd

`func (o *LifecycleIteration) SetGracePeriodEnd(v string)`

SetGracePeriodEnd sets GracePeriodEnd field to given value.


### GetNextReviewDate

`func (o *LifecycleIteration) GetNextReviewDate() string`

GetNextReviewDate returns the NextReviewDate field if non-nil, zero value otherwise.

### GetNextReviewDateOk

`func (o *LifecycleIteration) GetNextReviewDateOk() (*string, bool)`

GetNextReviewDateOk returns a tuple with the NextReviewDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextReviewDate

`func (o *LifecycleIteration) SetNextReviewDate(v string)`

SetNextReviewDate sets NextReviewDate field to given value.


### GetReviews

`func (o *LifecycleIteration) GetReviews() []Review`

GetReviews returns the Reviews field if non-nil, zero value otherwise.

### GetReviewsOk

`func (o *LifecycleIteration) GetReviewsOk() (*[]Review, bool)`

GetReviewsOk returns a tuple with the Reviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviews

`func (o *LifecycleIteration) SetReviews(v []Review)`

SetReviews sets Reviews field to given value.


### GetUserCanReview

`func (o *LifecycleIteration) GetUserCanReview() bool`

GetUserCanReview returns the UserCanReview field if non-nil, zero value otherwise.

### GetUserCanReviewOk

`func (o *LifecycleIteration) GetUserCanReviewOk() (*bool, bool)`

GetUserCanReviewOk returns a tuple with the UserCanReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCanReview

`func (o *LifecycleIteration) SetUserCanReview(v bool)`

SetUserCanReview sets UserCanReview field to given value.


### GetReviewerGroups

`func (o *LifecycleIteration) GetReviewerGroups() []ReviewerGroup`

GetReviewerGroups returns the ReviewerGroups field if non-nil, zero value otherwise.

### GetReviewerGroupsOk

`func (o *LifecycleIteration) GetReviewerGroupsOk() (*[]ReviewerGroup, bool)`

GetReviewerGroupsOk returns a tuple with the ReviewerGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerGroups

`func (o *LifecycleIteration) SetReviewerGroups(v []ReviewerGroup)`

SetReviewerGroups sets ReviewerGroups field to given value.


### GetMinReviewers

`func (o *LifecycleIteration) GetMinReviewers() int32`

GetMinReviewers returns the MinReviewers field if non-nil, zero value otherwise.

### GetMinReviewersOk

`func (o *LifecycleIteration) GetMinReviewersOk() (*int32, bool)`

GetMinReviewersOk returns a tuple with the MinReviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReviewers

`func (o *LifecycleIteration) SetMinReviewers(v int32)`

SetMinReviewers sets MinReviewers field to given value.


### GetReviewers

`func (o *LifecycleIteration) GetReviewers() []ReviewerUser`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *LifecycleIteration) GetReviewersOk() (*[]ReviewerUser, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *LifecycleIteration) SetReviewers(v []ReviewerUser)`

SetReviewers sets Reviewers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# RelatedRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**ReviewerGroups** | [**[]ReviewerGroup**](ReviewerGroup.md) |  | [readonly] 
**MinReviewers** | **int32** |  | [readonly] 
**Reviewers** | [**[]ReviewerUser**](ReviewerUser.md) |  | [readonly] 

## Methods

### NewRelatedRule

`func NewRelatedRule(name string, reviewerGroups []ReviewerGroup, minReviewers int32, reviewers []ReviewerUser, ) *RelatedRule`

NewRelatedRule instantiates a new RelatedRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelatedRuleWithDefaults

`func NewRelatedRuleWithDefaults() *RelatedRule`

NewRelatedRuleWithDefaults instantiates a new RelatedRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RelatedRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelatedRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelatedRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RelatedRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RelatedRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RelatedRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RelatedRule) SetName(v string)`

SetName sets Name field to given value.


### GetReviewerGroups

`func (o *RelatedRule) GetReviewerGroups() []ReviewerGroup`

GetReviewerGroups returns the ReviewerGroups field if non-nil, zero value otherwise.

### GetReviewerGroupsOk

`func (o *RelatedRule) GetReviewerGroupsOk() (*[]ReviewerGroup, bool)`

GetReviewerGroupsOk returns a tuple with the ReviewerGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerGroups

`func (o *RelatedRule) SetReviewerGroups(v []ReviewerGroup)`

SetReviewerGroups sets ReviewerGroups field to given value.


### GetMinReviewers

`func (o *RelatedRule) GetMinReviewers() int32`

GetMinReviewers returns the MinReviewers field if non-nil, zero value otherwise.

### GetMinReviewersOk

`func (o *RelatedRule) GetMinReviewersOk() (*int32, bool)`

GetMinReviewersOk returns a tuple with the MinReviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReviewers

`func (o *RelatedRule) SetMinReviewers(v int32)`

SetMinReviewers sets MinReviewers field to given value.


### GetReviewers

`func (o *RelatedRule) GetReviewers() []ReviewerUser`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *RelatedRule) GetReviewersOk() (*[]ReviewerUser, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *RelatedRule) SetReviewers(v []ReviewerUser)`

SetReviewers sets Reviewers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



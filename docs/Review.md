# Review

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Iteration** | **string** |  | 
**Reviewer** | [**ReviewerUser**](ReviewerUser.md) |  | [readonly] 
**Timestamp** | **time.Time** |  | [readonly] 
**Note** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewReview

`func NewReview(id string, iteration string, reviewer ReviewerUser, timestamp time.Time, ) *Review`

NewReview instantiates a new Review object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewWithDefaults

`func NewReviewWithDefaults() *Review`

NewReviewWithDefaults instantiates a new Review object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Review) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Review) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Review) SetId(v string)`

SetId sets Id field to given value.


### GetIteration

`func (o *Review) GetIteration() string`

GetIteration returns the Iteration field if non-nil, zero value otherwise.

### GetIterationOk

`func (o *Review) GetIterationOk() (*string, bool)`

GetIterationOk returns a tuple with the Iteration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIteration

`func (o *Review) SetIteration(v string)`

SetIteration sets Iteration field to given value.


### GetReviewer

`func (o *Review) GetReviewer() ReviewerUser`

GetReviewer returns the Reviewer field if non-nil, zero value otherwise.

### GetReviewerOk

`func (o *Review) GetReviewerOk() (*ReviewerUser, bool)`

GetReviewerOk returns a tuple with the Reviewer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewer

`func (o *Review) SetReviewer(v ReviewerUser)`

SetReviewer sets Reviewer field to given value.


### GetTimestamp

`func (o *Review) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Review) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Review) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetNote

`func (o *Review) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *Review) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *Review) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *Review) HasNote() bool`

HasNote returns a boolean if a field has been set.

### SetNoteNil

`func (o *Review) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *Review) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



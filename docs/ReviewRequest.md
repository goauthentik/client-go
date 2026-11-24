# ReviewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iteration** | **string** |  | 
**Note** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewReviewRequest

`func NewReviewRequest(iteration string, ) *ReviewRequest`

NewReviewRequest instantiates a new ReviewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewRequestWithDefaults

`func NewReviewRequestWithDefaults() *ReviewRequest`

NewReviewRequestWithDefaults instantiates a new ReviewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIteration

`func (o *ReviewRequest) GetIteration() string`

GetIteration returns the Iteration field if non-nil, zero value otherwise.

### GetIterationOk

`func (o *ReviewRequest) GetIterationOk() (*string, bool)`

GetIterationOk returns a tuple with the Iteration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIteration

`func (o *ReviewRequest) SetIteration(v string)`

SetIteration sets Iteration field to given value.


### GetNote

`func (o *ReviewRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ReviewRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ReviewRequest) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *ReviewRequest) HasNote() bool`

HasNote returns a boolean if a field has been set.

### SetNoteNil

`func (o *ReviewRequest) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *ReviewRequest) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



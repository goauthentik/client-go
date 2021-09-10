# UserReputation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**Username** | **string** |  | 
**Score** | Pointer to **int32** |  | [optional] 
**Updated** | **time.Time** |  | [readonly] 

## Methods

### NewUserReputation

`func NewUserReputation(pk int32, username string, updated time.Time, ) *UserReputation`

NewUserReputation instantiates a new UserReputation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserReputationWithDefaults

`func NewUserReputationWithDefaults() *UserReputation`

NewUserReputationWithDefaults instantiates a new UserReputation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *UserReputation) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *UserReputation) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *UserReputation) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetUsername

`func (o *UserReputation) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserReputation) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserReputation) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetScore

`func (o *UserReputation) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *UserReputation) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *UserReputation) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *UserReputation) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetUpdated

`func (o *UserReputation) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *UserReputation) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *UserReputation) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



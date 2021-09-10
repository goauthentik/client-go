# IPReputation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**Ip** | **string** |  | 
**Score** | Pointer to **int32** |  | [optional] 
**Updated** | **time.Time** |  | [readonly] 

## Methods

### NewIPReputation

`func NewIPReputation(pk int32, ip string, updated time.Time, ) *IPReputation`

NewIPReputation instantiates a new IPReputation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPReputationWithDefaults

`func NewIPReputationWithDefaults() *IPReputation`

NewIPReputationWithDefaults instantiates a new IPReputation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *IPReputation) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *IPReputation) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *IPReputation) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetIp

`func (o *IPReputation) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *IPReputation) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *IPReputation) SetIp(v string)`

SetIp sets Ip field to given value.


### GetScore

`func (o *IPReputation) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *IPReputation) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *IPReputation) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *IPReputation) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetUpdated

`func (o *IPReputation) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *IPReputation) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *IPReputation) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



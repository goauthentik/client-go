# Reputation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | Pointer to **string** |  | [optional] 
**Identifier** | **string** |  | 
**Ip** | **string** |  | 
**IpGeoData** | Pointer to **map[string]interface{}** |  | [optional] 
**Score** | Pointer to **int64** |  | [optional] 
**Updated** | **time.Time** |  | [readonly] 

## Methods

### NewReputation

`func NewReputation(identifier string, ip string, updated time.Time, ) *Reputation`

NewReputation instantiates a new Reputation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReputationWithDefaults

`func NewReputationWithDefaults() *Reputation`

NewReputationWithDefaults instantiates a new Reputation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *Reputation) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *Reputation) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *Reputation) SetPk(v string)`

SetPk sets Pk field to given value.

### HasPk

`func (o *Reputation) HasPk() bool`

HasPk returns a boolean if a field has been set.

### GetIdentifier

`func (o *Reputation) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Reputation) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Reputation) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetIp

`func (o *Reputation) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *Reputation) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *Reputation) SetIp(v string)`

SetIp sets Ip field to given value.


### GetIpGeoData

`func (o *Reputation) GetIpGeoData() map[string]interface{}`

GetIpGeoData returns the IpGeoData field if non-nil, zero value otherwise.

### GetIpGeoDataOk

`func (o *Reputation) GetIpGeoDataOk() (*map[string]interface{}, bool)`

GetIpGeoDataOk returns a tuple with the IpGeoData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpGeoData

`func (o *Reputation) SetIpGeoData(v map[string]interface{})`

SetIpGeoData sets IpGeoData field to given value.

### HasIpGeoData

`func (o *Reputation) HasIpGeoData() bool`

HasIpGeoData returns a boolean if a field has been set.

### GetScore

`func (o *Reputation) GetScore() int64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *Reputation) GetScoreOk() (*int64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *Reputation) SetScore(v int64)`

SetScore sets Score field to given value.

### HasScore

`func (o *Reputation) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetUpdated

`func (o *Reputation) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Reputation) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Reputation) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



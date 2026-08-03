# GrantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** |  | [readonly] 
**CreatedBy** | [**PartialUser**](PartialUser.md) |  | [readonly] 
**RequesterData** | Pointer to **map[string]interface{}** |  | [optional] 
**FulfillerData** | Pointer to **map[string]interface{}** |  | [optional] 
**RevokedBy** | [**PartialUser**](PartialUser.md) |  | [readonly] 
**IsActive** | **bool** |  | [readonly] 
**Expires** | Pointer to **NullableTime** |  | [optional] 
**Status** | [**RequestStatus**](RequestStatus.md) |  | [readonly] 
**Targets** | **[]string** |  | [readonly] 
**TargetObjs** | [**[]RequestableTarget**](RequestableTarget.md) |  | [readonly] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewGrantRequest

`func NewGrantRequest(created time.Time, createdBy PartialUser, revokedBy PartialUser, isActive bool, status RequestStatus, targets []string, targetObjs []RequestableTarget, ) *GrantRequest`

NewGrantRequest instantiates a new GrantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantRequestWithDefaults

`func NewGrantRequestWithDefaults() *GrantRequest`

NewGrantRequestWithDefaults instantiates a new GrantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *GrantRequest) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GrantRequest) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GrantRequest) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCreatedBy

`func (o *GrantRequest) GetCreatedBy() PartialUser`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GrantRequest) GetCreatedByOk() (*PartialUser, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GrantRequest) SetCreatedBy(v PartialUser)`

SetCreatedBy sets CreatedBy field to given value.


### GetRequesterData

`func (o *GrantRequest) GetRequesterData() map[string]interface{}`

GetRequesterData returns the RequesterData field if non-nil, zero value otherwise.

### GetRequesterDataOk

`func (o *GrantRequest) GetRequesterDataOk() (*map[string]interface{}, bool)`

GetRequesterDataOk returns a tuple with the RequesterData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterData

`func (o *GrantRequest) SetRequesterData(v map[string]interface{})`

SetRequesterData sets RequesterData field to given value.

### HasRequesterData

`func (o *GrantRequest) HasRequesterData() bool`

HasRequesterData returns a boolean if a field has been set.

### GetFulfillerData

`func (o *GrantRequest) GetFulfillerData() map[string]interface{}`

GetFulfillerData returns the FulfillerData field if non-nil, zero value otherwise.

### GetFulfillerDataOk

`func (o *GrantRequest) GetFulfillerDataOk() (*map[string]interface{}, bool)`

GetFulfillerDataOk returns a tuple with the FulfillerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillerData

`func (o *GrantRequest) SetFulfillerData(v map[string]interface{})`

SetFulfillerData sets FulfillerData field to given value.

### HasFulfillerData

`func (o *GrantRequest) HasFulfillerData() bool`

HasFulfillerData returns a boolean if a field has been set.

### GetRevokedBy

`func (o *GrantRequest) GetRevokedBy() PartialUser`

GetRevokedBy returns the RevokedBy field if non-nil, zero value otherwise.

### GetRevokedByOk

`func (o *GrantRequest) GetRevokedByOk() (*PartialUser, bool)`

GetRevokedByOk returns a tuple with the RevokedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedBy

`func (o *GrantRequest) SetRevokedBy(v PartialUser)`

SetRevokedBy sets RevokedBy field to given value.


### GetIsActive

`func (o *GrantRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *GrantRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *GrantRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetExpires

`func (o *GrantRequest) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *GrantRequest) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *GrantRequest) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *GrantRequest) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *GrantRequest) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *GrantRequest) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetStatus

`func (o *GrantRequest) GetStatus() RequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GrantRequest) GetStatusOk() (*RequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GrantRequest) SetStatus(v RequestStatus)`

SetStatus sets Status field to given value.


### GetTargets

`func (o *GrantRequest) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *GrantRequest) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *GrantRequest) SetTargets(v []string)`

SetTargets sets Targets field to given value.


### GetTargetObjs

`func (o *GrantRequest) GetTargetObjs() []RequestableTarget`

GetTargetObjs returns the TargetObjs field if non-nil, zero value otherwise.

### GetTargetObjsOk

`func (o *GrantRequest) GetTargetObjsOk() (*[]RequestableTarget, bool)`

GetTargetObjsOk returns a tuple with the TargetObjs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetObjs

`func (o *GrantRequest) SetTargetObjs(v []RequestableTarget)`

SetTargetObjs sets TargetObjs field to given value.


### GetUuid

`func (o *GrantRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GrantRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GrantRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GrantRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UserExpirationRulePreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | [readonly] 
**Users** | [**[]UserExpirationRuleUser**](UserExpirationRuleUser.md) |  | [readonly] 
**Updated** | [**UserExpirationRulePendingPreviewGroup**](UserExpirationRulePendingPreviewGroup.md) |  | [readonly] 
**TakenOver** | [**UserExpirationRulePendingPreviewGroup**](UserExpirationRulePendingPreviewGroup.md) |  | [readonly] 
**Removed** | [**UserExpirationRulePendingPreviewGroup**](UserExpirationRulePendingPreviewGroup.md) |  | [readonly] 

## Methods

### NewUserExpirationRulePreview

`func NewUserExpirationRulePreview(count int32, users []UserExpirationRuleUser, updated UserExpirationRulePendingPreviewGroup, takenOver UserExpirationRulePendingPreviewGroup, removed UserExpirationRulePendingPreviewGroup, ) *UserExpirationRulePreview`

NewUserExpirationRulePreview instantiates a new UserExpirationRulePreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserExpirationRulePreviewWithDefaults

`func NewUserExpirationRulePreviewWithDefaults() *UserExpirationRulePreview`

NewUserExpirationRulePreviewWithDefaults instantiates a new UserExpirationRulePreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *UserExpirationRulePreview) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UserExpirationRulePreview) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UserExpirationRulePreview) SetCount(v int32)`

SetCount sets Count field to given value.


### GetUsers

`func (o *UserExpirationRulePreview) GetUsers() []UserExpirationRuleUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UserExpirationRulePreview) GetUsersOk() (*[]UserExpirationRuleUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UserExpirationRulePreview) SetUsers(v []UserExpirationRuleUser)`

SetUsers sets Users field to given value.


### GetUpdated

`func (o *UserExpirationRulePreview) GetUpdated() UserExpirationRulePendingPreviewGroup`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *UserExpirationRulePreview) GetUpdatedOk() (*UserExpirationRulePendingPreviewGroup, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *UserExpirationRulePreview) SetUpdated(v UserExpirationRulePendingPreviewGroup)`

SetUpdated sets Updated field to given value.


### GetTakenOver

`func (o *UserExpirationRulePreview) GetTakenOver() UserExpirationRulePendingPreviewGroup`

GetTakenOver returns the TakenOver field if non-nil, zero value otherwise.

### GetTakenOverOk

`func (o *UserExpirationRulePreview) GetTakenOverOk() (*UserExpirationRulePendingPreviewGroup, bool)`

GetTakenOverOk returns a tuple with the TakenOver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakenOver

`func (o *UserExpirationRulePreview) SetTakenOver(v UserExpirationRulePendingPreviewGroup)`

SetTakenOver sets TakenOver field to given value.


### GetRemoved

`func (o *UserExpirationRulePreview) GetRemoved() UserExpirationRulePendingPreviewGroup`

GetRemoved returns the Removed field if non-nil, zero value otherwise.

### GetRemovedOk

`func (o *UserExpirationRulePreview) GetRemovedOk() (*UserExpirationRulePendingPreviewGroup, bool)`

GetRemovedOk returns a tuple with the Removed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoved

`func (o *UserExpirationRulePreview) SetRemoved(v UserExpirationRulePendingPreviewGroup)`

SetRemoved sets Removed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



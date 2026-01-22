# BulkDeleteSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserIds** | Pointer to **[]int32** | List of user IDs to revoke all sessions for | [optional] 

## Methods

### NewBulkDeleteSession

`func NewBulkDeleteSession() *BulkDeleteSession`

NewBulkDeleteSession instantiates a new BulkDeleteSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkDeleteSessionWithDefaults

`func NewBulkDeleteSessionWithDefaults() *BulkDeleteSession`

NewBulkDeleteSessionWithDefaults instantiates a new BulkDeleteSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserIds

`func (o *BulkDeleteSession) GetUserIds() []int32`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *BulkDeleteSession) GetUserIdsOk() (*[]int32, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *BulkDeleteSession) SetUserIds(v []int32)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *BulkDeleteSession) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



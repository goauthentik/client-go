# LDAPDebug

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | **[]map[string]interface{}** |  | [readonly] 
**Group** | **[]map[string]interface{}** |  | [readonly] 
**Membership** | **[]map[string]interface{}** |  | [readonly] 

## Methods

### NewLDAPDebug

`func NewLDAPDebug(user []map[string]interface{}, group []map[string]interface{}, membership []map[string]interface{}, ) *LDAPDebug`

NewLDAPDebug instantiates a new LDAPDebug object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLDAPDebugWithDefaults

`func NewLDAPDebugWithDefaults() *LDAPDebug`

NewLDAPDebugWithDefaults instantiates a new LDAPDebug object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *LDAPDebug) GetUser() []map[string]interface{}`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LDAPDebug) GetUserOk() (*[]map[string]interface{}, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LDAPDebug) SetUser(v []map[string]interface{})`

SetUser sets User field to given value.


### GetGroup

`func (o *LDAPDebug) GetGroup() []map[string]interface{}`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *LDAPDebug) GetGroupOk() (*[]map[string]interface{}, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *LDAPDebug) SetGroup(v []map[string]interface{})`

SetGroup sets Group field to given value.


### GetMembership

`func (o *LDAPDebug) GetMembership() []map[string]interface{}`

GetMembership returns the Membership field if non-nil, zero value otherwise.

### GetMembershipOk

`func (o *LDAPDebug) GetMembershipOk() (*[]map[string]interface{}, bool)`

GetMembershipOk returns a tuple with the Membership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembership

`func (o *LDAPDebug) SetMembership(v []map[string]interface{})`

SetMembership sets Membership field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



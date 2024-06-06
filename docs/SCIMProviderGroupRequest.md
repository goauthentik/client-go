# SCIMProviderGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScimId** | **string** |  | 
**Group** | **string** |  | 
**Provider** | **int32** |  | 

## Methods

### NewSCIMProviderGroupRequest

`func NewSCIMProviderGroupRequest(scimId string, group string, provider int32, ) *SCIMProviderGroupRequest`

NewSCIMProviderGroupRequest instantiates a new SCIMProviderGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSCIMProviderGroupRequestWithDefaults

`func NewSCIMProviderGroupRequestWithDefaults() *SCIMProviderGroupRequest`

NewSCIMProviderGroupRequestWithDefaults instantiates a new SCIMProviderGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScimId

`func (o *SCIMProviderGroupRequest) GetScimId() string`

GetScimId returns the ScimId field if non-nil, zero value otherwise.

### GetScimIdOk

`func (o *SCIMProviderGroupRequest) GetScimIdOk() (*string, bool)`

GetScimIdOk returns a tuple with the ScimId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScimId

`func (o *SCIMProviderGroupRequest) SetScimId(v string)`

SetScimId sets ScimId field to given value.


### GetGroup

`func (o *SCIMProviderGroupRequest) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SCIMProviderGroupRequest) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SCIMProviderGroupRequest) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetProvider

`func (o *SCIMProviderGroupRequest) GetProvider() int32`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SCIMProviderGroupRequest) GetProviderOk() (*int32, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SCIMProviderGroupRequest) SetProvider(v int32)`

SetProvider sets Provider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



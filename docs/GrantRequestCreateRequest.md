# GrantRequestCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pbms** | **[]string** |  | 
**Expiry** | Pointer to **string** | Optional override for how long the grant should last once approved. Clamped to the granting rule binding(s)&#39; expiry_granted_max. | [optional] 

## Methods

### NewGrantRequestCreateRequest

`func NewGrantRequestCreateRequest(pbms []string, ) *GrantRequestCreateRequest`

NewGrantRequestCreateRequest instantiates a new GrantRequestCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantRequestCreateRequestWithDefaults

`func NewGrantRequestCreateRequestWithDefaults() *GrantRequestCreateRequest`

NewGrantRequestCreateRequestWithDefaults instantiates a new GrantRequestCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPbms

`func (o *GrantRequestCreateRequest) GetPbms() []string`

GetPbms returns the Pbms field if non-nil, zero value otherwise.

### GetPbmsOk

`func (o *GrantRequestCreateRequest) GetPbmsOk() (*[]string, bool)`

GetPbmsOk returns a tuple with the Pbms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPbms

`func (o *GrantRequestCreateRequest) SetPbms(v []string)`

SetPbms sets Pbms field to given value.


### GetExpiry

`func (o *GrantRequestCreateRequest) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *GrantRequestCreateRequest) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *GrantRequestCreateRequest) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *GrantRequestCreateRequest) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



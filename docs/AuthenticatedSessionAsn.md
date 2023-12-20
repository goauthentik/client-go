# AuthenticatedSessionAsn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asn** | **int32** |  | 
**AsOrg** | **NullableString** |  | 
**Network** | **NullableString** |  | 

## Methods

### NewAuthenticatedSessionAsn

`func NewAuthenticatedSessionAsn(asn int32, asOrg NullableString, network NullableString, ) *AuthenticatedSessionAsn`

NewAuthenticatedSessionAsn instantiates a new AuthenticatedSessionAsn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatedSessionAsnWithDefaults

`func NewAuthenticatedSessionAsnWithDefaults() *AuthenticatedSessionAsn`

NewAuthenticatedSessionAsnWithDefaults instantiates a new AuthenticatedSessionAsn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsn

`func (o *AuthenticatedSessionAsn) GetAsn() int32`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *AuthenticatedSessionAsn) GetAsnOk() (*int32, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *AuthenticatedSessionAsn) SetAsn(v int32)`

SetAsn sets Asn field to given value.


### GetAsOrg

`func (o *AuthenticatedSessionAsn) GetAsOrg() string`

GetAsOrg returns the AsOrg field if non-nil, zero value otherwise.

### GetAsOrgOk

`func (o *AuthenticatedSessionAsn) GetAsOrgOk() (*string, bool)`

GetAsOrgOk returns a tuple with the AsOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOrg

`func (o *AuthenticatedSessionAsn) SetAsOrg(v string)`

SetAsOrg sets AsOrg field to given value.


### SetAsOrgNil

`func (o *AuthenticatedSessionAsn) SetAsOrgNil(b bool)`

 SetAsOrgNil sets the value for AsOrg to be an explicit nil

### UnsetAsOrg
`func (o *AuthenticatedSessionAsn) UnsetAsOrg()`

UnsetAsOrg ensures that no value is present for AsOrg, not even an explicit nil
### GetNetwork

`func (o *AuthenticatedSessionAsn) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *AuthenticatedSessionAsn) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *AuthenticatedSessionAsn) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### SetNetworkNil

`func (o *AuthenticatedSessionAsn) SetNetworkNil(b bool)`

 SetNetworkNil sets the value for Network to be an explicit nil

### UnsetNetwork
`func (o *AuthenticatedSessionAsn) UnsetNetwork()`

UnsetNetwork ensures that no value is present for Network, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



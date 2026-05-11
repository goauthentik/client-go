# TenantRecoveryKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiry** | **time.Time** |  | 
**Url** | **string** |  | 

## Methods

### NewTenantRecoveryKeyResponse

`func NewTenantRecoveryKeyResponse(expiry time.Time, url string, ) *TenantRecoveryKeyResponse`

NewTenantRecoveryKeyResponse instantiates a new TenantRecoveryKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantRecoveryKeyResponseWithDefaults

`func NewTenantRecoveryKeyResponseWithDefaults() *TenantRecoveryKeyResponse`

NewTenantRecoveryKeyResponseWithDefaults instantiates a new TenantRecoveryKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiry

`func (o *TenantRecoveryKeyResponse) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *TenantRecoveryKeyResponse) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *TenantRecoveryKeyResponse) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.


### GetUrl

`func (o *TenantRecoveryKeyResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TenantRecoveryKeyResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TenantRecoveryKeyResponse) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



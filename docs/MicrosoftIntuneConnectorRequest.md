# MicrosoftIntuneConnectorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectorUuid** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**ClientId** | **string** |  | 
**ClientSecret** | **string** |  | 
**TenantId** | **string** |  | 
**CertificateAuthorities** | Pointer to **[]string** | Certificate authorities which issue device certificates via Intune (Cloud PKI, SCEP or PKCS profiles), used by the endpoint stage to validate certificates. This option has a higher priority than the &#x60;client_certificate&#x60; option on &#x60;Brand&#x60;. | [optional] 

## Methods

### NewMicrosoftIntuneConnectorRequest

`func NewMicrosoftIntuneConnectorRequest(name string, clientId string, clientSecret string, tenantId string, ) *MicrosoftIntuneConnectorRequest`

NewMicrosoftIntuneConnectorRequest instantiates a new MicrosoftIntuneConnectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftIntuneConnectorRequestWithDefaults

`func NewMicrosoftIntuneConnectorRequestWithDefaults() *MicrosoftIntuneConnectorRequest`

NewMicrosoftIntuneConnectorRequestWithDefaults instantiates a new MicrosoftIntuneConnectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectorUuid

`func (o *MicrosoftIntuneConnectorRequest) GetConnectorUuid() string`

GetConnectorUuid returns the ConnectorUuid field if non-nil, zero value otherwise.

### GetConnectorUuidOk

`func (o *MicrosoftIntuneConnectorRequest) GetConnectorUuidOk() (*string, bool)`

GetConnectorUuidOk returns a tuple with the ConnectorUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorUuid

`func (o *MicrosoftIntuneConnectorRequest) SetConnectorUuid(v string)`

SetConnectorUuid sets ConnectorUuid field to given value.

### HasConnectorUuid

`func (o *MicrosoftIntuneConnectorRequest) HasConnectorUuid() bool`

HasConnectorUuid returns a boolean if a field has been set.

### GetName

`func (o *MicrosoftIntuneConnectorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftIntuneConnectorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftIntuneConnectorRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *MicrosoftIntuneConnectorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MicrosoftIntuneConnectorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MicrosoftIntuneConnectorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MicrosoftIntuneConnectorRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetClientId

`func (o *MicrosoftIntuneConnectorRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *MicrosoftIntuneConnectorRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *MicrosoftIntuneConnectorRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *MicrosoftIntuneConnectorRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *MicrosoftIntuneConnectorRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *MicrosoftIntuneConnectorRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetTenantId

`func (o *MicrosoftIntuneConnectorRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *MicrosoftIntuneConnectorRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *MicrosoftIntuneConnectorRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCertificateAuthorities

`func (o *MicrosoftIntuneConnectorRequest) GetCertificateAuthorities() []string`

GetCertificateAuthorities returns the CertificateAuthorities field if non-nil, zero value otherwise.

### GetCertificateAuthoritiesOk

`func (o *MicrosoftIntuneConnectorRequest) GetCertificateAuthoritiesOk() (*[]string, bool)`

GetCertificateAuthoritiesOk returns a tuple with the CertificateAuthorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorities

`func (o *MicrosoftIntuneConnectorRequest) SetCertificateAuthorities(v []string)`

SetCertificateAuthorities sets CertificateAuthorities field to given value.

### HasCertificateAuthorities

`func (o *MicrosoftIntuneConnectorRequest) HasCertificateAuthorities() bool`

HasCertificateAuthorities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AgentPSSOUserRegistrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserAuth** | **string** |  | 
**UserSecureEnclaveKey** | **string** |  | 
**EnclaveKeyId** | **string** |  | 

## Methods

### NewAgentPSSOUserRegistrationRequest

`func NewAgentPSSOUserRegistrationRequest(userAuth string, userSecureEnclaveKey string, enclaveKeyId string, ) *AgentPSSOUserRegistrationRequest`

NewAgentPSSOUserRegistrationRequest instantiates a new AgentPSSOUserRegistrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentPSSOUserRegistrationRequestWithDefaults

`func NewAgentPSSOUserRegistrationRequestWithDefaults() *AgentPSSOUserRegistrationRequest`

NewAgentPSSOUserRegistrationRequestWithDefaults instantiates a new AgentPSSOUserRegistrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserAuth

`func (o *AgentPSSOUserRegistrationRequest) GetUserAuth() string`

GetUserAuth returns the UserAuth field if non-nil, zero value otherwise.

### GetUserAuthOk

`func (o *AgentPSSOUserRegistrationRequest) GetUserAuthOk() (*string, bool)`

GetUserAuthOk returns a tuple with the UserAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAuth

`func (o *AgentPSSOUserRegistrationRequest) SetUserAuth(v string)`

SetUserAuth sets UserAuth field to given value.


### GetUserSecureEnclaveKey

`func (o *AgentPSSOUserRegistrationRequest) GetUserSecureEnclaveKey() string`

GetUserSecureEnclaveKey returns the UserSecureEnclaveKey field if non-nil, zero value otherwise.

### GetUserSecureEnclaveKeyOk

`func (o *AgentPSSOUserRegistrationRequest) GetUserSecureEnclaveKeyOk() (*string, bool)`

GetUserSecureEnclaveKeyOk returns a tuple with the UserSecureEnclaveKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSecureEnclaveKey

`func (o *AgentPSSOUserRegistrationRequest) SetUserSecureEnclaveKey(v string)`

SetUserSecureEnclaveKey sets UserSecureEnclaveKey field to given value.


### GetEnclaveKeyId

`func (o *AgentPSSOUserRegistrationRequest) GetEnclaveKeyId() string`

GetEnclaveKeyId returns the EnclaveKeyId field if non-nil, zero value otherwise.

### GetEnclaveKeyIdOk

`func (o *AgentPSSOUserRegistrationRequest) GetEnclaveKeyIdOk() (*string, bool)`

GetEnclaveKeyIdOk returns a tuple with the EnclaveKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclaveKeyId

`func (o *AgentPSSOUserRegistrationRequest) SetEnclaveKeyId(v string)`

SetEnclaveKeyId sets EnclaveKeyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



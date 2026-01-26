# RACProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**AuthenticationFlow** | Pointer to **NullableString** | Flow used for authentication when the associated application is accessed by an un-authenticated user. | [optional] 
**AuthorizationFlow** | **string** | Flow used when authorizing this provider. | 
**PropertyMappings** | Pointer to **[]string** |  | [optional] 
**Settings** | Pointer to **interface{}** |  | [optional] 
**ConnectionExpiry** | Pointer to **string** | Determines how long a session lasts. Default of 0 means that the sessions lasts until the browser is closed. (Format: hours&#x3D;-1;minutes&#x3D;-2;seconds&#x3D;-3) | [optional] 
**DeleteTokenOnDisconnect** | Pointer to **bool** | When set to true, connection tokens will be deleted upon disconnect. | [optional] 

## Methods

### NewRACProviderRequest

`func NewRACProviderRequest(name string, authorizationFlow string, ) *RACProviderRequest`

NewRACProviderRequest instantiates a new RACProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRACProviderRequestWithDefaults

`func NewRACProviderRequestWithDefaults() *RACProviderRequest`

NewRACProviderRequestWithDefaults instantiates a new RACProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RACProviderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RACProviderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RACProviderRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAuthenticationFlow

`func (o *RACProviderRequest) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *RACProviderRequest) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *RACProviderRequest) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *RACProviderRequest) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *RACProviderRequest) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *RACProviderRequest) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetAuthorizationFlow

`func (o *RACProviderRequest) GetAuthorizationFlow() string`

GetAuthorizationFlow returns the AuthorizationFlow field if non-nil, zero value otherwise.

### GetAuthorizationFlowOk

`func (o *RACProviderRequest) GetAuthorizationFlowOk() (*string, bool)`

GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationFlow

`func (o *RACProviderRequest) SetAuthorizationFlow(v string)`

SetAuthorizationFlow sets AuthorizationFlow field to given value.


### GetPropertyMappings

`func (o *RACProviderRequest) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *RACProviderRequest) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *RACProviderRequest) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *RACProviderRequest) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetSettings

`func (o *RACProviderRequest) GetSettings() interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *RACProviderRequest) GetSettingsOk() (*interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *RACProviderRequest) SetSettings(v interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *RACProviderRequest) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *RACProviderRequest) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *RACProviderRequest) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil
### GetConnectionExpiry

`func (o *RACProviderRequest) GetConnectionExpiry() string`

GetConnectionExpiry returns the ConnectionExpiry field if non-nil, zero value otherwise.

### GetConnectionExpiryOk

`func (o *RACProviderRequest) GetConnectionExpiryOk() (*string, bool)`

GetConnectionExpiryOk returns a tuple with the ConnectionExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionExpiry

`func (o *RACProviderRequest) SetConnectionExpiry(v string)`

SetConnectionExpiry sets ConnectionExpiry field to given value.

### HasConnectionExpiry

`func (o *RACProviderRequest) HasConnectionExpiry() bool`

HasConnectionExpiry returns a boolean if a field has been set.

### GetDeleteTokenOnDisconnect

`func (o *RACProviderRequest) GetDeleteTokenOnDisconnect() bool`

GetDeleteTokenOnDisconnect returns the DeleteTokenOnDisconnect field if non-nil, zero value otherwise.

### GetDeleteTokenOnDisconnectOk

`func (o *RACProviderRequest) GetDeleteTokenOnDisconnectOk() (*bool, bool)`

GetDeleteTokenOnDisconnectOk returns a tuple with the DeleteTokenOnDisconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteTokenOnDisconnect

`func (o *RACProviderRequest) SetDeleteTokenOnDisconnect(v bool)`

SetDeleteTokenOnDisconnect sets DeleteTokenOnDisconnect field to given value.

### HasDeleteTokenOnDisconnect

`func (o *RACProviderRequest) HasDeleteTokenOnDisconnect() bool`

HasDeleteTokenOnDisconnect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



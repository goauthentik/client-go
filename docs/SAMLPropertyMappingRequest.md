# SAMLPropertyMappingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Managed** | Pointer to **NullableString** | Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update. | [optional] 
**Name** | **string** |  | 
**Expression** | **string** |  | 
**SamlName** | **string** |  | 
**FriendlyName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSAMLPropertyMappingRequest

`func NewSAMLPropertyMappingRequest(name string, expression string, samlName string, ) *SAMLPropertyMappingRequest`

NewSAMLPropertyMappingRequest instantiates a new SAMLPropertyMappingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAMLPropertyMappingRequestWithDefaults

`func NewSAMLPropertyMappingRequestWithDefaults() *SAMLPropertyMappingRequest`

NewSAMLPropertyMappingRequestWithDefaults instantiates a new SAMLPropertyMappingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManaged

`func (o *SAMLPropertyMappingRequest) GetManaged() string`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *SAMLPropertyMappingRequest) GetManagedOk() (*string, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *SAMLPropertyMappingRequest) SetManaged(v string)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *SAMLPropertyMappingRequest) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### SetManagedNil

`func (o *SAMLPropertyMappingRequest) SetManagedNil(b bool)`

 SetManagedNil sets the value for Managed to be an explicit nil

### UnsetManaged
`func (o *SAMLPropertyMappingRequest) UnsetManaged()`

UnsetManaged ensures that no value is present for Managed, not even an explicit nil
### GetName

`func (o *SAMLPropertyMappingRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SAMLPropertyMappingRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SAMLPropertyMappingRequest) SetName(v string)`

SetName sets Name field to given value.


### GetExpression

`func (o *SAMLPropertyMappingRequest) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *SAMLPropertyMappingRequest) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *SAMLPropertyMappingRequest) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetSamlName

`func (o *SAMLPropertyMappingRequest) GetSamlName() string`

GetSamlName returns the SamlName field if non-nil, zero value otherwise.

### GetSamlNameOk

`func (o *SAMLPropertyMappingRequest) GetSamlNameOk() (*string, bool)`

GetSamlNameOk returns a tuple with the SamlName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlName

`func (o *SAMLPropertyMappingRequest) SetSamlName(v string)`

SetSamlName sets SamlName field to given value.


### GetFriendlyName

`func (o *SAMLPropertyMappingRequest) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *SAMLPropertyMappingRequest) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *SAMLPropertyMappingRequest) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *SAMLPropertyMappingRequest) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *SAMLPropertyMappingRequest) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *SAMLPropertyMappingRequest) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



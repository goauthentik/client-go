# PatchedSAMLPropertyMappingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Managed** | Pointer to **NullableString** | Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Expression** | Pointer to **string** |  | [optional] 
**SamlName** | Pointer to **string** |  | [optional] 
**FriendlyName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPatchedSAMLPropertyMappingRequest

`func NewPatchedSAMLPropertyMappingRequest() *PatchedSAMLPropertyMappingRequest`

NewPatchedSAMLPropertyMappingRequest instantiates a new PatchedSAMLPropertyMappingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedSAMLPropertyMappingRequestWithDefaults

`func NewPatchedSAMLPropertyMappingRequestWithDefaults() *PatchedSAMLPropertyMappingRequest`

NewPatchedSAMLPropertyMappingRequestWithDefaults instantiates a new PatchedSAMLPropertyMappingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManaged

`func (o *PatchedSAMLPropertyMappingRequest) GetManaged() string`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *PatchedSAMLPropertyMappingRequest) GetManagedOk() (*string, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *PatchedSAMLPropertyMappingRequest) SetManaged(v string)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *PatchedSAMLPropertyMappingRequest) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### SetManagedNil

`func (o *PatchedSAMLPropertyMappingRequest) SetManagedNil(b bool)`

 SetManagedNil sets the value for Managed to be an explicit nil

### UnsetManaged
`func (o *PatchedSAMLPropertyMappingRequest) UnsetManaged()`

UnsetManaged ensures that no value is present for Managed, not even an explicit nil
### GetName

`func (o *PatchedSAMLPropertyMappingRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedSAMLPropertyMappingRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedSAMLPropertyMappingRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedSAMLPropertyMappingRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExpression

`func (o *PatchedSAMLPropertyMappingRequest) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *PatchedSAMLPropertyMappingRequest) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *PatchedSAMLPropertyMappingRequest) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *PatchedSAMLPropertyMappingRequest) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetSamlName

`func (o *PatchedSAMLPropertyMappingRequest) GetSamlName() string`

GetSamlName returns the SamlName field if non-nil, zero value otherwise.

### GetSamlNameOk

`func (o *PatchedSAMLPropertyMappingRequest) GetSamlNameOk() (*string, bool)`

GetSamlNameOk returns a tuple with the SamlName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlName

`func (o *PatchedSAMLPropertyMappingRequest) SetSamlName(v string)`

SetSamlName sets SamlName field to given value.

### HasSamlName

`func (o *PatchedSAMLPropertyMappingRequest) HasSamlName() bool`

HasSamlName returns a boolean if a field has been set.

### GetFriendlyName

`func (o *PatchedSAMLPropertyMappingRequest) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *PatchedSAMLPropertyMappingRequest) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *PatchedSAMLPropertyMappingRequest) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *PatchedSAMLPropertyMappingRequest) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *PatchedSAMLPropertyMappingRequest) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *PatchedSAMLPropertyMappingRequest) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



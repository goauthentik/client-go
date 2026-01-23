# PatchedFleetConnectorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectorUuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**HeadersMapping** | Pointer to **NullableString** | Configure additional headers to be sent. Mapping should return a dictionary of key-value pairs | [optional] 
**MapUsers** | Pointer to **bool** |  | [optional] 
**MapTeamsAccessGroup** | Pointer to **bool** |  | [optional] 

## Methods

### NewPatchedFleetConnectorRequest

`func NewPatchedFleetConnectorRequest() *PatchedFleetConnectorRequest`

NewPatchedFleetConnectorRequest instantiates a new PatchedFleetConnectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedFleetConnectorRequestWithDefaults

`func NewPatchedFleetConnectorRequestWithDefaults() *PatchedFleetConnectorRequest`

NewPatchedFleetConnectorRequestWithDefaults instantiates a new PatchedFleetConnectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectorUuid

`func (o *PatchedFleetConnectorRequest) GetConnectorUuid() string`

GetConnectorUuid returns the ConnectorUuid field if non-nil, zero value otherwise.

### GetConnectorUuidOk

`func (o *PatchedFleetConnectorRequest) GetConnectorUuidOk() (*string, bool)`

GetConnectorUuidOk returns a tuple with the ConnectorUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorUuid

`func (o *PatchedFleetConnectorRequest) SetConnectorUuid(v string)`

SetConnectorUuid sets ConnectorUuid field to given value.

### HasConnectorUuid

`func (o *PatchedFleetConnectorRequest) HasConnectorUuid() bool`

HasConnectorUuid returns a boolean if a field has been set.

### GetName

`func (o *PatchedFleetConnectorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedFleetConnectorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedFleetConnectorRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedFleetConnectorRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedFleetConnectorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedFleetConnectorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedFleetConnectorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedFleetConnectorRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUrl

`func (o *PatchedFleetConnectorRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedFleetConnectorRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedFleetConnectorRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedFleetConnectorRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetToken

`func (o *PatchedFleetConnectorRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PatchedFleetConnectorRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PatchedFleetConnectorRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PatchedFleetConnectorRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetHeadersMapping

`func (o *PatchedFleetConnectorRequest) GetHeadersMapping() string`

GetHeadersMapping returns the HeadersMapping field if non-nil, zero value otherwise.

### GetHeadersMappingOk

`func (o *PatchedFleetConnectorRequest) GetHeadersMappingOk() (*string, bool)`

GetHeadersMappingOk returns a tuple with the HeadersMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadersMapping

`func (o *PatchedFleetConnectorRequest) SetHeadersMapping(v string)`

SetHeadersMapping sets HeadersMapping field to given value.

### HasHeadersMapping

`func (o *PatchedFleetConnectorRequest) HasHeadersMapping() bool`

HasHeadersMapping returns a boolean if a field has been set.

### SetHeadersMappingNil

`func (o *PatchedFleetConnectorRequest) SetHeadersMappingNil(b bool)`

 SetHeadersMappingNil sets the value for HeadersMapping to be an explicit nil

### UnsetHeadersMapping
`func (o *PatchedFleetConnectorRequest) UnsetHeadersMapping()`

UnsetHeadersMapping ensures that no value is present for HeadersMapping, not even an explicit nil
### GetMapUsers

`func (o *PatchedFleetConnectorRequest) GetMapUsers() bool`

GetMapUsers returns the MapUsers field if non-nil, zero value otherwise.

### GetMapUsersOk

`func (o *PatchedFleetConnectorRequest) GetMapUsersOk() (*bool, bool)`

GetMapUsersOk returns a tuple with the MapUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapUsers

`func (o *PatchedFleetConnectorRequest) SetMapUsers(v bool)`

SetMapUsers sets MapUsers field to given value.

### HasMapUsers

`func (o *PatchedFleetConnectorRequest) HasMapUsers() bool`

HasMapUsers returns a boolean if a field has been set.

### GetMapTeamsAccessGroup

`func (o *PatchedFleetConnectorRequest) GetMapTeamsAccessGroup() bool`

GetMapTeamsAccessGroup returns the MapTeamsAccessGroup field if non-nil, zero value otherwise.

### GetMapTeamsAccessGroupOk

`func (o *PatchedFleetConnectorRequest) GetMapTeamsAccessGroupOk() (*bool, bool)`

GetMapTeamsAccessGroupOk returns a tuple with the MapTeamsAccessGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapTeamsAccessGroup

`func (o *PatchedFleetConnectorRequest) SetMapTeamsAccessGroup(v bool)`

SetMapTeamsAccessGroup sets MapTeamsAccessGroup field to given value.

### HasMapTeamsAccessGroup

`func (o *PatchedFleetConnectorRequest) HasMapTeamsAccessGroup() bool`

HasMapTeamsAccessGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



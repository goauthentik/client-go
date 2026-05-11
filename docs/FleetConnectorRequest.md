# FleetConnectorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectorUuid** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**Url** | **string** |  | 
**Token** | **string** |  | 
**HeadersMapping** | Pointer to **NullableString** | Configure additional headers to be sent. Mapping should return a dictionary of key-value pairs | [optional] 
**MapUsers** | Pointer to **bool** |  | [optional] 
**MapTeamsAccessGroup** | Pointer to **bool** |  | [optional] 

## Methods

### NewFleetConnectorRequest

`func NewFleetConnectorRequest(name string, url string, token string, ) *FleetConnectorRequest`

NewFleetConnectorRequest instantiates a new FleetConnectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetConnectorRequestWithDefaults

`func NewFleetConnectorRequestWithDefaults() *FleetConnectorRequest`

NewFleetConnectorRequestWithDefaults instantiates a new FleetConnectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectorUuid

`func (o *FleetConnectorRequest) GetConnectorUuid() string`

GetConnectorUuid returns the ConnectorUuid field if non-nil, zero value otherwise.

### GetConnectorUuidOk

`func (o *FleetConnectorRequest) GetConnectorUuidOk() (*string, bool)`

GetConnectorUuidOk returns a tuple with the ConnectorUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorUuid

`func (o *FleetConnectorRequest) SetConnectorUuid(v string)`

SetConnectorUuid sets ConnectorUuid field to given value.

### HasConnectorUuid

`func (o *FleetConnectorRequest) HasConnectorUuid() bool`

HasConnectorUuid returns a boolean if a field has been set.

### GetName

`func (o *FleetConnectorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FleetConnectorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FleetConnectorRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *FleetConnectorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FleetConnectorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FleetConnectorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FleetConnectorRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUrl

`func (o *FleetConnectorRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FleetConnectorRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FleetConnectorRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetToken

`func (o *FleetConnectorRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetConnectorRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetConnectorRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetHeadersMapping

`func (o *FleetConnectorRequest) GetHeadersMapping() string`

GetHeadersMapping returns the HeadersMapping field if non-nil, zero value otherwise.

### GetHeadersMappingOk

`func (o *FleetConnectorRequest) GetHeadersMappingOk() (*string, bool)`

GetHeadersMappingOk returns a tuple with the HeadersMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadersMapping

`func (o *FleetConnectorRequest) SetHeadersMapping(v string)`

SetHeadersMapping sets HeadersMapping field to given value.

### HasHeadersMapping

`func (o *FleetConnectorRequest) HasHeadersMapping() bool`

HasHeadersMapping returns a boolean if a field has been set.

### SetHeadersMappingNil

`func (o *FleetConnectorRequest) SetHeadersMappingNil(b bool)`

 SetHeadersMappingNil sets the value for HeadersMapping to be an explicit nil

### UnsetHeadersMapping
`func (o *FleetConnectorRequest) UnsetHeadersMapping()`

UnsetHeadersMapping ensures that no value is present for HeadersMapping, not even an explicit nil
### GetMapUsers

`func (o *FleetConnectorRequest) GetMapUsers() bool`

GetMapUsers returns the MapUsers field if non-nil, zero value otherwise.

### GetMapUsersOk

`func (o *FleetConnectorRequest) GetMapUsersOk() (*bool, bool)`

GetMapUsersOk returns a tuple with the MapUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapUsers

`func (o *FleetConnectorRequest) SetMapUsers(v bool)`

SetMapUsers sets MapUsers field to given value.

### HasMapUsers

`func (o *FleetConnectorRequest) HasMapUsers() bool`

HasMapUsers returns a boolean if a field has been set.

### GetMapTeamsAccessGroup

`func (o *FleetConnectorRequest) GetMapTeamsAccessGroup() bool`

GetMapTeamsAccessGroup returns the MapTeamsAccessGroup field if non-nil, zero value otherwise.

### GetMapTeamsAccessGroupOk

`func (o *FleetConnectorRequest) GetMapTeamsAccessGroupOk() (*bool, bool)`

GetMapTeamsAccessGroupOk returns a tuple with the MapTeamsAccessGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapTeamsAccessGroup

`func (o *FleetConnectorRequest) SetMapTeamsAccessGroup(v bool)`

SetMapTeamsAccessGroup sets MapTeamsAccessGroup field to given value.

### HasMapTeamsAccessGroup

`func (o *FleetConnectorRequest) HasMapTeamsAccessGroup() bool`

HasMapTeamsAccessGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



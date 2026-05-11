# FleetConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectorUuid** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**Component** | **string** | Get object component so that we know how to edit the object | [readonly] 
**VerboseName** | **string** | Return object&#39;s verbose_name | [readonly] 
**VerboseNamePlural** | **string** | Return object&#39;s plural verbose_name | [readonly] 
**MetaModelName** | **string** | Return internal model name | [readonly] 
**Url** | **string** |  | 
**HeadersMapping** | Pointer to **NullableString** | Configure additional headers to be sent. Mapping should return a dictionary of key-value pairs | [optional] 
**MapUsers** | Pointer to **bool** |  | [optional] 
**MapTeamsAccessGroup** | Pointer to **bool** |  | [optional] 

## Methods

### NewFleetConnector

`func NewFleetConnector(name string, component string, verboseName string, verboseNamePlural string, metaModelName string, url string, ) *FleetConnector`

NewFleetConnector instantiates a new FleetConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetConnectorWithDefaults

`func NewFleetConnectorWithDefaults() *FleetConnector`

NewFleetConnectorWithDefaults instantiates a new FleetConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectorUuid

`func (o *FleetConnector) GetConnectorUuid() string`

GetConnectorUuid returns the ConnectorUuid field if non-nil, zero value otherwise.

### GetConnectorUuidOk

`func (o *FleetConnector) GetConnectorUuidOk() (*string, bool)`

GetConnectorUuidOk returns a tuple with the ConnectorUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorUuid

`func (o *FleetConnector) SetConnectorUuid(v string)`

SetConnectorUuid sets ConnectorUuid field to given value.

### HasConnectorUuid

`func (o *FleetConnector) HasConnectorUuid() bool`

HasConnectorUuid returns a boolean if a field has been set.

### GetName

`func (o *FleetConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FleetConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FleetConnector) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *FleetConnector) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FleetConnector) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FleetConnector) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FleetConnector) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetComponent

`func (o *FleetConnector) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *FleetConnector) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *FleetConnector) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *FleetConnector) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *FleetConnector) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *FleetConnector) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *FleetConnector) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *FleetConnector) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *FleetConnector) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *FleetConnector) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *FleetConnector) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *FleetConnector) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetUrl

`func (o *FleetConnector) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FleetConnector) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FleetConnector) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeadersMapping

`func (o *FleetConnector) GetHeadersMapping() string`

GetHeadersMapping returns the HeadersMapping field if non-nil, zero value otherwise.

### GetHeadersMappingOk

`func (o *FleetConnector) GetHeadersMappingOk() (*string, bool)`

GetHeadersMappingOk returns a tuple with the HeadersMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadersMapping

`func (o *FleetConnector) SetHeadersMapping(v string)`

SetHeadersMapping sets HeadersMapping field to given value.

### HasHeadersMapping

`func (o *FleetConnector) HasHeadersMapping() bool`

HasHeadersMapping returns a boolean if a field has been set.

### SetHeadersMappingNil

`func (o *FleetConnector) SetHeadersMappingNil(b bool)`

 SetHeadersMappingNil sets the value for HeadersMapping to be an explicit nil

### UnsetHeadersMapping
`func (o *FleetConnector) UnsetHeadersMapping()`

UnsetHeadersMapping ensures that no value is present for HeadersMapping, not even an explicit nil
### GetMapUsers

`func (o *FleetConnector) GetMapUsers() bool`

GetMapUsers returns the MapUsers field if non-nil, zero value otherwise.

### GetMapUsersOk

`func (o *FleetConnector) GetMapUsersOk() (*bool, bool)`

GetMapUsersOk returns a tuple with the MapUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapUsers

`func (o *FleetConnector) SetMapUsers(v bool)`

SetMapUsers sets MapUsers field to given value.

### HasMapUsers

`func (o *FleetConnector) HasMapUsers() bool`

HasMapUsers returns a boolean if a field has been set.

### GetMapTeamsAccessGroup

`func (o *FleetConnector) GetMapTeamsAccessGroup() bool`

GetMapTeamsAccessGroup returns the MapTeamsAccessGroup field if non-nil, zero value otherwise.

### GetMapTeamsAccessGroupOk

`func (o *FleetConnector) GetMapTeamsAccessGroupOk() (*bool, bool)`

GetMapTeamsAccessGroupOk returns a tuple with the MapTeamsAccessGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapTeamsAccessGroup

`func (o *FleetConnector) SetMapTeamsAccessGroup(v bool)`

SetMapTeamsAccessGroup sets MapTeamsAccessGroup field to given value.

### HasMapTeamsAccessGroup

`func (o *FleetConnector) HasMapTeamsAccessGroup() bool`

HasMapTeamsAccessGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



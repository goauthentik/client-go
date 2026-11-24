# GoogleChromeConnector

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
**Credentials** | **map[string]interface{}** |  | 
**ChromeUrl** | **NullableString** | Full URL to be used in Google Workspace configuration | [readonly] 

## Methods

### NewGoogleChromeConnector

`func NewGoogleChromeConnector(name string, component string, verboseName string, verboseNamePlural string, metaModelName string, credentials map[string]interface{}, chromeUrl NullableString, ) *GoogleChromeConnector`

NewGoogleChromeConnector instantiates a new GoogleChromeConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleChromeConnectorWithDefaults

`func NewGoogleChromeConnectorWithDefaults() *GoogleChromeConnector`

NewGoogleChromeConnectorWithDefaults instantiates a new GoogleChromeConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectorUuid

`func (o *GoogleChromeConnector) GetConnectorUuid() string`

GetConnectorUuid returns the ConnectorUuid field if non-nil, zero value otherwise.

### GetConnectorUuidOk

`func (o *GoogleChromeConnector) GetConnectorUuidOk() (*string, bool)`

GetConnectorUuidOk returns a tuple with the ConnectorUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorUuid

`func (o *GoogleChromeConnector) SetConnectorUuid(v string)`

SetConnectorUuid sets ConnectorUuid field to given value.

### HasConnectorUuid

`func (o *GoogleChromeConnector) HasConnectorUuid() bool`

HasConnectorUuid returns a boolean if a field has been set.

### GetName

`func (o *GoogleChromeConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GoogleChromeConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GoogleChromeConnector) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *GoogleChromeConnector) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GoogleChromeConnector) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GoogleChromeConnector) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GoogleChromeConnector) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetComponent

`func (o *GoogleChromeConnector) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *GoogleChromeConnector) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *GoogleChromeConnector) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *GoogleChromeConnector) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *GoogleChromeConnector) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *GoogleChromeConnector) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *GoogleChromeConnector) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *GoogleChromeConnector) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *GoogleChromeConnector) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *GoogleChromeConnector) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *GoogleChromeConnector) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *GoogleChromeConnector) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetCredentials

`func (o *GoogleChromeConnector) GetCredentials() map[string]interface{}`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *GoogleChromeConnector) GetCredentialsOk() (*map[string]interface{}, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *GoogleChromeConnector) SetCredentials(v map[string]interface{})`

SetCredentials sets Credentials field to given value.


### GetChromeUrl

`func (o *GoogleChromeConnector) GetChromeUrl() string`

GetChromeUrl returns the ChromeUrl field if non-nil, zero value otherwise.

### GetChromeUrlOk

`func (o *GoogleChromeConnector) GetChromeUrlOk() (*string, bool)`

GetChromeUrlOk returns a tuple with the ChromeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromeUrl

`func (o *GoogleChromeConnector) SetChromeUrl(v string)`

SetChromeUrl sets ChromeUrl field to given value.


### SetChromeUrlNil

`func (o *GoogleChromeConnector) SetChromeUrlNil(b bool)`

 SetChromeUrlNil sets the value for ChromeUrl to be an explicit nil

### UnsetChromeUrl
`func (o *GoogleChromeConnector) UnsetChromeUrl()`

UnsetChromeUrl ensures that no value is present for ChromeUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



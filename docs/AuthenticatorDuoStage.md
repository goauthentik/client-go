# AuthenticatorDuoStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Component** | **string** |  | [readonly] 
**VerboseName** | **string** |  | [readonly] 
**VerboseNamePlural** | **string** |  | [readonly] 
**MetaModelName** | **string** |  | [readonly] 
**FlowSet** | Pointer to [**[]Flow**](Flow.md) |  | [optional] 
**ConfigureFlow** | Pointer to **NullableString** | Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage. | [optional] 
**ClientId** | **string** |  | 
**ApiHostname** | **string** |  | 

## Methods

### NewAuthenticatorDuoStage

`func NewAuthenticatorDuoStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, clientId string, apiHostname string, ) *AuthenticatorDuoStage`

NewAuthenticatorDuoStage instantiates a new AuthenticatorDuoStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorDuoStageWithDefaults

`func NewAuthenticatorDuoStageWithDefaults() *AuthenticatorDuoStage`

NewAuthenticatorDuoStageWithDefaults instantiates a new AuthenticatorDuoStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *AuthenticatorDuoStage) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *AuthenticatorDuoStage) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *AuthenticatorDuoStage) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *AuthenticatorDuoStage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthenticatorDuoStage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthenticatorDuoStage) SetName(v string)`

SetName sets Name field to given value.


### GetComponent

`func (o *AuthenticatorDuoStage) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *AuthenticatorDuoStage) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *AuthenticatorDuoStage) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *AuthenticatorDuoStage) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *AuthenticatorDuoStage) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *AuthenticatorDuoStage) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *AuthenticatorDuoStage) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *AuthenticatorDuoStage) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *AuthenticatorDuoStage) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *AuthenticatorDuoStage) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *AuthenticatorDuoStage) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *AuthenticatorDuoStage) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetFlowSet

`func (o *AuthenticatorDuoStage) GetFlowSet() []Flow`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *AuthenticatorDuoStage) GetFlowSetOk() (*[]Flow, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *AuthenticatorDuoStage) SetFlowSet(v []Flow)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *AuthenticatorDuoStage) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetConfigureFlow

`func (o *AuthenticatorDuoStage) GetConfigureFlow() string`

GetConfigureFlow returns the ConfigureFlow field if non-nil, zero value otherwise.

### GetConfigureFlowOk

`func (o *AuthenticatorDuoStage) GetConfigureFlowOk() (*string, bool)`

GetConfigureFlowOk returns a tuple with the ConfigureFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureFlow

`func (o *AuthenticatorDuoStage) SetConfigureFlow(v string)`

SetConfigureFlow sets ConfigureFlow field to given value.

### HasConfigureFlow

`func (o *AuthenticatorDuoStage) HasConfigureFlow() bool`

HasConfigureFlow returns a boolean if a field has been set.

### SetConfigureFlowNil

`func (o *AuthenticatorDuoStage) SetConfigureFlowNil(b bool)`

 SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil

### UnsetConfigureFlow
`func (o *AuthenticatorDuoStage) UnsetConfigureFlow()`

UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
### GetClientId

`func (o *AuthenticatorDuoStage) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AuthenticatorDuoStage) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AuthenticatorDuoStage) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetApiHostname

`func (o *AuthenticatorDuoStage) GetApiHostname() string`

GetApiHostname returns the ApiHostname field if non-nil, zero value otherwise.

### GetApiHostnameOk

`func (o *AuthenticatorDuoStage) GetApiHostnameOk() (*string, bool)`

GetApiHostnameOk returns a tuple with the ApiHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiHostname

`func (o *AuthenticatorDuoStage) SetApiHostname(v string)`

SetApiHostname sets ApiHostname field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



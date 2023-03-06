# ReputationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** |  | 
**ExecutionLogging** | Pointer to **bool** | When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged. | [optional] 
**Component** | **string** | Get object component so that we know how to edit the object | [readonly] 
**VerboseName** | **string** | Return object&#39;s verbose_name | [readonly] 
**VerboseNamePlural** | **string** | Return object&#39;s plural verbose_name | [readonly] 
**MetaModelName** | **string** | Return internal model name | [readonly] 
**BoundTo** | **int32** | Return objects policy is bound to | [readonly] 
**CheckIp** | Pointer to **bool** |  | [optional] 
**CheckUsername** | Pointer to **bool** |  | [optional] 
**Threshold** | Pointer to **int32** |  | [optional] 

## Methods

### NewReputationPolicy

`func NewReputationPolicy(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, boundTo int32, ) *ReputationPolicy`

NewReputationPolicy instantiates a new ReputationPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReputationPolicyWithDefaults

`func NewReputationPolicyWithDefaults() *ReputationPolicy`

NewReputationPolicyWithDefaults instantiates a new ReputationPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *ReputationPolicy) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *ReputationPolicy) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *ReputationPolicy) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *ReputationPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReputationPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReputationPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetExecutionLogging

`func (o *ReputationPolicy) GetExecutionLogging() bool`

GetExecutionLogging returns the ExecutionLogging field if non-nil, zero value otherwise.

### GetExecutionLoggingOk

`func (o *ReputationPolicy) GetExecutionLoggingOk() (*bool, bool)`

GetExecutionLoggingOk returns a tuple with the ExecutionLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionLogging

`func (o *ReputationPolicy) SetExecutionLogging(v bool)`

SetExecutionLogging sets ExecutionLogging field to given value.

### HasExecutionLogging

`func (o *ReputationPolicy) HasExecutionLogging() bool`

HasExecutionLogging returns a boolean if a field has been set.

### GetComponent

`func (o *ReputationPolicy) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ReputationPolicy) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ReputationPolicy) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *ReputationPolicy) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *ReputationPolicy) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *ReputationPolicy) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *ReputationPolicy) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *ReputationPolicy) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *ReputationPolicy) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *ReputationPolicy) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *ReputationPolicy) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *ReputationPolicy) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetBoundTo

`func (o *ReputationPolicy) GetBoundTo() int32`

GetBoundTo returns the BoundTo field if non-nil, zero value otherwise.

### GetBoundToOk

`func (o *ReputationPolicy) GetBoundToOk() (*int32, bool)`

GetBoundToOk returns a tuple with the BoundTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundTo

`func (o *ReputationPolicy) SetBoundTo(v int32)`

SetBoundTo sets BoundTo field to given value.


### GetCheckIp

`func (o *ReputationPolicy) GetCheckIp() bool`

GetCheckIp returns the CheckIp field if non-nil, zero value otherwise.

### GetCheckIpOk

`func (o *ReputationPolicy) GetCheckIpOk() (*bool, bool)`

GetCheckIpOk returns a tuple with the CheckIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckIp

`func (o *ReputationPolicy) SetCheckIp(v bool)`

SetCheckIp sets CheckIp field to given value.

### HasCheckIp

`func (o *ReputationPolicy) HasCheckIp() bool`

HasCheckIp returns a boolean if a field has been set.

### GetCheckUsername

`func (o *ReputationPolicy) GetCheckUsername() bool`

GetCheckUsername returns the CheckUsername field if non-nil, zero value otherwise.

### GetCheckUsernameOk

`func (o *ReputationPolicy) GetCheckUsernameOk() (*bool, bool)`

GetCheckUsernameOk returns a tuple with the CheckUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckUsername

`func (o *ReputationPolicy) SetCheckUsername(v bool)`

SetCheckUsername sets CheckUsername field to given value.

### HasCheckUsername

`func (o *ReputationPolicy) HasCheckUsername() bool`

HasCheckUsername returns a boolean if a field has been set.

### GetThreshold

`func (o *ReputationPolicy) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ReputationPolicy) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ReputationPolicy) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *ReputationPolicy) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DummyPolicy

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
**Result** | Pointer to **bool** |  | [optional] 
**WaitMin** | Pointer to **int32** |  | [optional] 
**WaitMax** | Pointer to **int32** |  | [optional] 

## Methods

### NewDummyPolicy

`func NewDummyPolicy(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, boundTo int32, ) *DummyPolicy`

NewDummyPolicy instantiates a new DummyPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDummyPolicyWithDefaults

`func NewDummyPolicyWithDefaults() *DummyPolicy`

NewDummyPolicyWithDefaults instantiates a new DummyPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *DummyPolicy) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *DummyPolicy) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *DummyPolicy) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *DummyPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DummyPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DummyPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetExecutionLogging

`func (o *DummyPolicy) GetExecutionLogging() bool`

GetExecutionLogging returns the ExecutionLogging field if non-nil, zero value otherwise.

### GetExecutionLoggingOk

`func (o *DummyPolicy) GetExecutionLoggingOk() (*bool, bool)`

GetExecutionLoggingOk returns a tuple with the ExecutionLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionLogging

`func (o *DummyPolicy) SetExecutionLogging(v bool)`

SetExecutionLogging sets ExecutionLogging field to given value.

### HasExecutionLogging

`func (o *DummyPolicy) HasExecutionLogging() bool`

HasExecutionLogging returns a boolean if a field has been set.

### GetComponent

`func (o *DummyPolicy) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *DummyPolicy) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *DummyPolicy) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *DummyPolicy) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *DummyPolicy) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *DummyPolicy) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *DummyPolicy) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *DummyPolicy) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *DummyPolicy) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *DummyPolicy) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *DummyPolicy) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *DummyPolicy) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetBoundTo

`func (o *DummyPolicy) GetBoundTo() int32`

GetBoundTo returns the BoundTo field if non-nil, zero value otherwise.

### GetBoundToOk

`func (o *DummyPolicy) GetBoundToOk() (*int32, bool)`

GetBoundToOk returns a tuple with the BoundTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundTo

`func (o *DummyPolicy) SetBoundTo(v int32)`

SetBoundTo sets BoundTo field to given value.


### GetResult

`func (o *DummyPolicy) GetResult() bool`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DummyPolicy) GetResultOk() (*bool, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DummyPolicy) SetResult(v bool)`

SetResult sets Result field to given value.

### HasResult

`func (o *DummyPolicy) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetWaitMin

`func (o *DummyPolicy) GetWaitMin() int32`

GetWaitMin returns the WaitMin field if non-nil, zero value otherwise.

### GetWaitMinOk

`func (o *DummyPolicy) GetWaitMinOk() (*int32, bool)`

GetWaitMinOk returns a tuple with the WaitMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitMin

`func (o *DummyPolicy) SetWaitMin(v int32)`

SetWaitMin sets WaitMin field to given value.

### HasWaitMin

`func (o *DummyPolicy) HasWaitMin() bool`

HasWaitMin returns a boolean if a field has been set.

### GetWaitMax

`func (o *DummyPolicy) GetWaitMax() int32`

GetWaitMax returns the WaitMax field if non-nil, zero value otherwise.

### GetWaitMaxOk

`func (o *DummyPolicy) GetWaitMaxOk() (*int32, bool)`

GetWaitMaxOk returns a tuple with the WaitMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitMax

`func (o *DummyPolicy) SetWaitMax(v int32)`

SetWaitMax sets WaitMax field to given value.

### HasWaitMax

`func (o *DummyPolicy) HasWaitMax() bool`

HasWaitMax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



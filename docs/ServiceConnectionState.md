# ServiceConnectionState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Healthy** | **bool** |  | [readonly] 
**Version** | **string** |  | [readonly] 

## Methods

### NewServiceConnectionState

`func NewServiceConnectionState(healthy bool, version string, ) *ServiceConnectionState`

NewServiceConnectionState instantiates a new ServiceConnectionState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceConnectionStateWithDefaults

`func NewServiceConnectionStateWithDefaults() *ServiceConnectionState`

NewServiceConnectionStateWithDefaults instantiates a new ServiceConnectionState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealthy

`func (o *ServiceConnectionState) GetHealthy() bool`

GetHealthy returns the Healthy field if non-nil, zero value otherwise.

### GetHealthyOk

`func (o *ServiceConnectionState) GetHealthyOk() (*bool, bool)`

GetHealthyOk returns a tuple with the Healthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthy

`func (o *ServiceConnectionState) SetHealthy(v bool)`

SetHealthy sets Healthy field to given value.


### GetVersion

`func (o *ServiceConnectionState) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServiceConnectionState) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServiceConnectionState) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



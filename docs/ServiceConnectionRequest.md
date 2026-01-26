# ServiceConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Local** | Pointer to **bool** | If enabled, use the local connection. Required Docker socket/Kubernetes Integration | [optional] 

## Methods

### NewServiceConnectionRequest

`func NewServiceConnectionRequest(name string, ) *ServiceConnectionRequest`

NewServiceConnectionRequest instantiates a new ServiceConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceConnectionRequestWithDefaults

`func NewServiceConnectionRequestWithDefaults() *ServiceConnectionRequest`

NewServiceConnectionRequestWithDefaults instantiates a new ServiceConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServiceConnectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceConnectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceConnectionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLocal

`func (o *ServiceConnectionRequest) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *ServiceConnectionRequest) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *ServiceConnectionRequest) SetLocal(v bool)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *ServiceConnectionRequest) HasLocal() bool`

HasLocal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



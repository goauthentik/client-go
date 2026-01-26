# RequestedBy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**Username** | **string** | Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. | 

## Methods

### NewRequestedBy

`func NewRequestedBy(pk int32, username string, ) *RequestedBy`

NewRequestedBy instantiates a new RequestedBy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestedByWithDefaults

`func NewRequestedByWithDefaults() *RequestedBy`

NewRequestedByWithDefaults instantiates a new RequestedBy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *RequestedBy) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *RequestedBy) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *RequestedBy) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetUsername

`func (o *RequestedBy) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RequestedBy) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RequestedBy) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# LoginMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoginsPer1h** | [**[]Coordinate**](Coordinate.md) |  | [readonly] 
**LoginsFailedPer1h** | [**[]Coordinate**](Coordinate.md) |  | [readonly] 

## Methods

### NewLoginMetrics

`func NewLoginMetrics(loginsPer1h []Coordinate, loginsFailedPer1h []Coordinate, ) *LoginMetrics`

NewLoginMetrics instantiates a new LoginMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginMetricsWithDefaults

`func NewLoginMetricsWithDefaults() *LoginMetrics`

NewLoginMetricsWithDefaults instantiates a new LoginMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoginsPer1h

`func (o *LoginMetrics) GetLoginsPer1h() []Coordinate`

GetLoginsPer1h returns the LoginsPer1h field if non-nil, zero value otherwise.

### GetLoginsPer1hOk

`func (o *LoginMetrics) GetLoginsPer1hOk() (*[]Coordinate, bool)`

GetLoginsPer1hOk returns a tuple with the LoginsPer1h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginsPer1h

`func (o *LoginMetrics) SetLoginsPer1h(v []Coordinate)`

SetLoginsPer1h sets LoginsPer1h field to given value.


### GetLoginsFailedPer1h

`func (o *LoginMetrics) GetLoginsFailedPer1h() []Coordinate`

GetLoginsFailedPer1h returns the LoginsFailedPer1h field if non-nil, zero value otherwise.

### GetLoginsFailedPer1hOk

`func (o *LoginMetrics) GetLoginsFailedPer1hOk() (*[]Coordinate, bool)`

GetLoginsFailedPer1hOk returns a tuple with the LoginsFailedPer1h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginsFailedPer1h

`func (o *LoginMetrics) SetLoginsFailedPer1h(v []Coordinate)`

SetLoginsFailedPer1h sets LoginsFailedPer1h field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



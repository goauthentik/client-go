# UserMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoginsPer1h** | [**[]Coordinate**](Coordinate.md) |  | [readonly] 
**LoginsFailedPer1h** | [**[]Coordinate**](Coordinate.md) |  | [readonly] 
**AuthorizationsPer1h** | [**[]Coordinate**](Coordinate.md) |  | [readonly] 

## Methods

### NewUserMetrics

`func NewUserMetrics(loginsPer1h []Coordinate, loginsFailedPer1h []Coordinate, authorizationsPer1h []Coordinate, ) *UserMetrics`

NewUserMetrics instantiates a new UserMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserMetricsWithDefaults

`func NewUserMetricsWithDefaults() *UserMetrics`

NewUserMetricsWithDefaults instantiates a new UserMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoginsPer1h

`func (o *UserMetrics) GetLoginsPer1h() []Coordinate`

GetLoginsPer1h returns the LoginsPer1h field if non-nil, zero value otherwise.

### GetLoginsPer1hOk

`func (o *UserMetrics) GetLoginsPer1hOk() (*[]Coordinate, bool)`

GetLoginsPer1hOk returns a tuple with the LoginsPer1h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginsPer1h

`func (o *UserMetrics) SetLoginsPer1h(v []Coordinate)`

SetLoginsPer1h sets LoginsPer1h field to given value.


### GetLoginsFailedPer1h

`func (o *UserMetrics) GetLoginsFailedPer1h() []Coordinate`

GetLoginsFailedPer1h returns the LoginsFailedPer1h field if non-nil, zero value otherwise.

### GetLoginsFailedPer1hOk

`func (o *UserMetrics) GetLoginsFailedPer1hOk() (*[]Coordinate, bool)`

GetLoginsFailedPer1hOk returns a tuple with the LoginsFailedPer1h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginsFailedPer1h

`func (o *UserMetrics) SetLoginsFailedPer1h(v []Coordinate)`

SetLoginsFailedPer1h sets LoginsFailedPer1h field to given value.


### GetAuthorizationsPer1h

`func (o *UserMetrics) GetAuthorizationsPer1h() []Coordinate`

GetAuthorizationsPer1h returns the AuthorizationsPer1h field if non-nil, zero value otherwise.

### GetAuthorizationsPer1hOk

`func (o *UserMetrics) GetAuthorizationsPer1hOk() (*[]Coordinate, bool)`

GetAuthorizationsPer1hOk returns a tuple with the AuthorizationsPer1h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationsPer1h

`func (o *UserMetrics) SetAuthorizationsPer1h(v []Coordinate)`

SetAuthorizationsPer1h sets AuthorizationsPer1h field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



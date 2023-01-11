# LoginMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logins** | [**[]Coordinate**](Coordinate.md) |  | [readonly] 
**LoginsFailed** | [**[]Coordinate**](Coordinate.md) |  | [readonly] 
**Authorizations** | [**[]Coordinate**](Coordinate.md) |  | [readonly] 

## Methods

### NewLoginMetrics

`func NewLoginMetrics(logins []Coordinate, loginsFailed []Coordinate, authorizations []Coordinate, ) *LoginMetrics`

NewLoginMetrics instantiates a new LoginMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginMetricsWithDefaults

`func NewLoginMetricsWithDefaults() *LoginMetrics`

NewLoginMetricsWithDefaults instantiates a new LoginMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogins

`func (o *LoginMetrics) GetLogins() []Coordinate`

GetLogins returns the Logins field if non-nil, zero value otherwise.

### GetLoginsOk

`func (o *LoginMetrics) GetLoginsOk() (*[]Coordinate, bool)`

GetLoginsOk returns a tuple with the Logins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogins

`func (o *LoginMetrics) SetLogins(v []Coordinate)`

SetLogins sets Logins field to given value.


### GetLoginsFailed

`func (o *LoginMetrics) GetLoginsFailed() []Coordinate`

GetLoginsFailed returns the LoginsFailed field if non-nil, zero value otherwise.

### GetLoginsFailedOk

`func (o *LoginMetrics) GetLoginsFailedOk() (*[]Coordinate, bool)`

GetLoginsFailedOk returns a tuple with the LoginsFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginsFailed

`func (o *LoginMetrics) SetLoginsFailed(v []Coordinate)`

SetLoginsFailed sets LoginsFailed field to given value.


### GetAuthorizations

`func (o *LoginMetrics) GetAuthorizations() []Coordinate`

GetAuthorizations returns the Authorizations field if non-nil, zero value otherwise.

### GetAuthorizationsOk

`func (o *LoginMetrics) GetAuthorizationsOk() (*[]Coordinate, bool)`

GetAuthorizationsOk returns a tuple with the Authorizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizations

`func (o *LoginMetrics) SetAuthorizations(v []Coordinate)`

SetAuthorizations sets Authorizations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



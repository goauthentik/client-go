# UserMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logins** | [**[]Coordinate**](Coordinate.md) |  | [readonly] 
**LoginsFailed** | [**[]Coordinate**](Coordinate.md) |  | [readonly] 
**Authorizations** | [**[]Coordinate**](Coordinate.md) |  | [readonly] 

## Methods

### NewUserMetrics

`func NewUserMetrics(logins []Coordinate, loginsFailed []Coordinate, authorizations []Coordinate, ) *UserMetrics`

NewUserMetrics instantiates a new UserMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserMetricsWithDefaults

`func NewUserMetricsWithDefaults() *UserMetrics`

NewUserMetricsWithDefaults instantiates a new UserMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogins

`func (o *UserMetrics) GetLogins() []Coordinate`

GetLogins returns the Logins field if non-nil, zero value otherwise.

### GetLoginsOk

`func (o *UserMetrics) GetLoginsOk() (*[]Coordinate, bool)`

GetLoginsOk returns a tuple with the Logins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogins

`func (o *UserMetrics) SetLogins(v []Coordinate)`

SetLogins sets Logins field to given value.


### GetLoginsFailed

`func (o *UserMetrics) GetLoginsFailed() []Coordinate`

GetLoginsFailed returns the LoginsFailed field if non-nil, zero value otherwise.

### GetLoginsFailedOk

`func (o *UserMetrics) GetLoginsFailedOk() (*[]Coordinate, bool)`

GetLoginsFailedOk returns a tuple with the LoginsFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginsFailed

`func (o *UserMetrics) SetLoginsFailed(v []Coordinate)`

SetLoginsFailed sets LoginsFailed field to given value.


### GetAuthorizations

`func (o *UserMetrics) GetAuthorizations() []Coordinate`

GetAuthorizations returns the Authorizations field if non-nil, zero value otherwise.

### GetAuthorizationsOk

`func (o *UserMetrics) GetAuthorizationsOk() (*[]Coordinate, bool)`

GetAuthorizationsOk returns a tuple with the Authorizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizations

`func (o *UserMetrics) SetAuthorizations(v []Coordinate)`

SetAuthorizations sets Authorizations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



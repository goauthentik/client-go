# SessionEndChallenge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowInfo** | Pointer to [**ContextualFlowInfo**](ContextualFlowInfo.md) |  | [optional] 
**Component** | Pointer to **string** |  | [optional] [default to "ak-stage-session-end"]
**ResponseErrors** | Pointer to [**map[string][]ErrorDetail**](array.md) |  | [optional] 
**PendingUser** | **string** |  | 
**PendingUserAvatar** | **string** |  | 
**ApplicationName** | Pointer to **string** |  | [optional] 
**ApplicationLaunchUrl** | Pointer to **string** |  | [optional] 
**InvalidationFlowUrl** | Pointer to **string** |  | [optional] 
**BrandName** | **string** |  | 

## Methods

### NewSessionEndChallenge

`func NewSessionEndChallenge(pendingUser string, pendingUserAvatar string, brandName string, ) *SessionEndChallenge`

NewSessionEndChallenge instantiates a new SessionEndChallenge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionEndChallengeWithDefaults

`func NewSessionEndChallengeWithDefaults() *SessionEndChallenge`

NewSessionEndChallengeWithDefaults instantiates a new SessionEndChallenge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowInfo

`func (o *SessionEndChallenge) GetFlowInfo() ContextualFlowInfo`

GetFlowInfo returns the FlowInfo field if non-nil, zero value otherwise.

### GetFlowInfoOk

`func (o *SessionEndChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool)`

GetFlowInfoOk returns a tuple with the FlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInfo

`func (o *SessionEndChallenge) SetFlowInfo(v ContextualFlowInfo)`

SetFlowInfo sets FlowInfo field to given value.

### HasFlowInfo

`func (o *SessionEndChallenge) HasFlowInfo() bool`

HasFlowInfo returns a boolean if a field has been set.

### GetComponent

`func (o *SessionEndChallenge) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *SessionEndChallenge) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *SessionEndChallenge) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *SessionEndChallenge) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetResponseErrors

`func (o *SessionEndChallenge) GetResponseErrors() map[string][]ErrorDetail`

GetResponseErrors returns the ResponseErrors field if non-nil, zero value otherwise.

### GetResponseErrorsOk

`func (o *SessionEndChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool)`

GetResponseErrorsOk returns a tuple with the ResponseErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseErrors

`func (o *SessionEndChallenge) SetResponseErrors(v map[string][]ErrorDetail)`

SetResponseErrors sets ResponseErrors field to given value.

### HasResponseErrors

`func (o *SessionEndChallenge) HasResponseErrors() bool`

HasResponseErrors returns a boolean if a field has been set.

### GetPendingUser

`func (o *SessionEndChallenge) GetPendingUser() string`

GetPendingUser returns the PendingUser field if non-nil, zero value otherwise.

### GetPendingUserOk

`func (o *SessionEndChallenge) GetPendingUserOk() (*string, bool)`

GetPendingUserOk returns a tuple with the PendingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUser

`func (o *SessionEndChallenge) SetPendingUser(v string)`

SetPendingUser sets PendingUser field to given value.


### GetPendingUserAvatar

`func (o *SessionEndChallenge) GetPendingUserAvatar() string`

GetPendingUserAvatar returns the PendingUserAvatar field if non-nil, zero value otherwise.

### GetPendingUserAvatarOk

`func (o *SessionEndChallenge) GetPendingUserAvatarOk() (*string, bool)`

GetPendingUserAvatarOk returns a tuple with the PendingUserAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUserAvatar

`func (o *SessionEndChallenge) SetPendingUserAvatar(v string)`

SetPendingUserAvatar sets PendingUserAvatar field to given value.


### GetApplicationName

`func (o *SessionEndChallenge) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *SessionEndChallenge) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *SessionEndChallenge) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *SessionEndChallenge) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetApplicationLaunchUrl

`func (o *SessionEndChallenge) GetApplicationLaunchUrl() string`

GetApplicationLaunchUrl returns the ApplicationLaunchUrl field if non-nil, zero value otherwise.

### GetApplicationLaunchUrlOk

`func (o *SessionEndChallenge) GetApplicationLaunchUrlOk() (*string, bool)`

GetApplicationLaunchUrlOk returns a tuple with the ApplicationLaunchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationLaunchUrl

`func (o *SessionEndChallenge) SetApplicationLaunchUrl(v string)`

SetApplicationLaunchUrl sets ApplicationLaunchUrl field to given value.

### HasApplicationLaunchUrl

`func (o *SessionEndChallenge) HasApplicationLaunchUrl() bool`

HasApplicationLaunchUrl returns a boolean if a field has been set.

### GetInvalidationFlowUrl

`func (o *SessionEndChallenge) GetInvalidationFlowUrl() string`

GetInvalidationFlowUrl returns the InvalidationFlowUrl field if non-nil, zero value otherwise.

### GetInvalidationFlowUrlOk

`func (o *SessionEndChallenge) GetInvalidationFlowUrlOk() (*string, bool)`

GetInvalidationFlowUrlOk returns a tuple with the InvalidationFlowUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidationFlowUrl

`func (o *SessionEndChallenge) SetInvalidationFlowUrl(v string)`

SetInvalidationFlowUrl sets InvalidationFlowUrl field to given value.

### HasInvalidationFlowUrl

`func (o *SessionEndChallenge) HasInvalidationFlowUrl() bool`

HasInvalidationFlowUrl returns a boolean if a field has been set.

### GetBrandName

`func (o *SessionEndChallenge) GetBrandName() string`

GetBrandName returns the BrandName field if non-nil, zero value otherwise.

### GetBrandNameOk

`func (o *SessionEndChallenge) GetBrandNameOk() (*string, bool)`

GetBrandNameOk returns a tuple with the BrandName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandName

`func (o *SessionEndChallenge) SetBrandName(v string)`

SetBrandName sets BrandName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# LoginChallengeTypes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ChallengeChoices**](ChallengeChoices.md) |  | 
**FlowInfo** | Pointer to [**ContextualFlowInfo**](ContextualFlowInfo.md) |  | [optional] 
**Component** | Pointer to **string** |  | [optional] [default to "ak-source-oauth-apple"]
**ResponseErrors** | Pointer to [**map[string][]ErrorDetail**](array.md) |  | [optional] 
**To** | **string** |  | 
**ClientId** | **string** |  | 
**Slug** | **string** |  | 
**Scope** | **string** |  | 
**RedirectUri** | **string** |  | 
**State** | **string** |  | 

## Methods

### NewLoginChallengeTypes

`func NewLoginChallengeTypes(type_ ChallengeChoices, to string, clientId string, slug string, scope string, redirectUri string, state string, ) *LoginChallengeTypes`

NewLoginChallengeTypes instantiates a new LoginChallengeTypes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginChallengeTypesWithDefaults

`func NewLoginChallengeTypesWithDefaults() *LoginChallengeTypes`

NewLoginChallengeTypesWithDefaults instantiates a new LoginChallengeTypes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LoginChallengeTypes) GetType() ChallengeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LoginChallengeTypes) GetTypeOk() (*ChallengeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LoginChallengeTypes) SetType(v ChallengeChoices)`

SetType sets Type field to given value.


### GetFlowInfo

`func (o *LoginChallengeTypes) GetFlowInfo() ContextualFlowInfo`

GetFlowInfo returns the FlowInfo field if non-nil, zero value otherwise.

### GetFlowInfoOk

`func (o *LoginChallengeTypes) GetFlowInfoOk() (*ContextualFlowInfo, bool)`

GetFlowInfoOk returns a tuple with the FlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInfo

`func (o *LoginChallengeTypes) SetFlowInfo(v ContextualFlowInfo)`

SetFlowInfo sets FlowInfo field to given value.

### HasFlowInfo

`func (o *LoginChallengeTypes) HasFlowInfo() bool`

HasFlowInfo returns a boolean if a field has been set.

### GetComponent

`func (o *LoginChallengeTypes) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *LoginChallengeTypes) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *LoginChallengeTypes) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *LoginChallengeTypes) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetResponseErrors

`func (o *LoginChallengeTypes) GetResponseErrors() map[string][]ErrorDetail`

GetResponseErrors returns the ResponseErrors field if non-nil, zero value otherwise.

### GetResponseErrorsOk

`func (o *LoginChallengeTypes) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool)`

GetResponseErrorsOk returns a tuple with the ResponseErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseErrors

`func (o *LoginChallengeTypes) SetResponseErrors(v map[string][]ErrorDetail)`

SetResponseErrors sets ResponseErrors field to given value.

### HasResponseErrors

`func (o *LoginChallengeTypes) HasResponseErrors() bool`

HasResponseErrors returns a boolean if a field has been set.

### GetTo

`func (o *LoginChallengeTypes) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *LoginChallengeTypes) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *LoginChallengeTypes) SetTo(v string)`

SetTo sets To field to given value.


### GetClientId

`func (o *LoginChallengeTypes) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *LoginChallengeTypes) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *LoginChallengeTypes) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetSlug

`func (o *LoginChallengeTypes) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *LoginChallengeTypes) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *LoginChallengeTypes) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetScope

`func (o *LoginChallengeTypes) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *LoginChallengeTypes) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *LoginChallengeTypes) SetScope(v string)`

SetScope sets Scope field to given value.


### GetRedirectUri

`func (o *LoginChallengeTypes) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *LoginChallengeTypes) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *LoginChallengeTypes) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.


### GetState

`func (o *LoginChallengeTypes) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LoginChallengeTypes) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LoginChallengeTypes) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



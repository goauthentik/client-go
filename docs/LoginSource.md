# LoginSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**IconUrl** | Pointer to **NullableString** |  | [optional] 
**Challenge** | [**LoginChallengeTypes**](LoginChallengeTypes.md) |  | 

## Methods

### NewLoginSource

`func NewLoginSource(name string, challenge LoginChallengeTypes, ) *LoginSource`

NewLoginSource instantiates a new LoginSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginSourceWithDefaults

`func NewLoginSourceWithDefaults() *LoginSource`

NewLoginSourceWithDefaults instantiates a new LoginSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LoginSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoginSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoginSource) SetName(v string)`

SetName sets Name field to given value.


### GetIconUrl

`func (o *LoginSource) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *LoginSource) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *LoginSource) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *LoginSource) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *LoginSource) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *LoginSource) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetChallenge

`func (o *LoginSource) GetChallenge() LoginChallengeTypes`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *LoginSource) GetChallengeOk() (*LoginChallengeTypes, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *LoginSource) SetChallenge(v LoginChallengeTypes)`

SetChallenge sets Challenge field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



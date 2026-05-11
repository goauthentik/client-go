# UserSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectUid** | **string** |  | 
**Component** | **string** |  | 
**Title** | **string** |  | 
**ConfigureUrl** | Pointer to **string** |  | [optional] 
**IconUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewUserSetting

`func NewUserSetting(objectUid string, component string, title string, ) *UserSetting`

NewUserSetting instantiates a new UserSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSettingWithDefaults

`func NewUserSettingWithDefaults() *UserSetting`

NewUserSettingWithDefaults instantiates a new UserSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectUid

`func (o *UserSetting) GetObjectUid() string`

GetObjectUid returns the ObjectUid field if non-nil, zero value otherwise.

### GetObjectUidOk

`func (o *UserSetting) GetObjectUidOk() (*string, bool)`

GetObjectUidOk returns a tuple with the ObjectUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectUid

`func (o *UserSetting) SetObjectUid(v string)`

SetObjectUid sets ObjectUid field to given value.


### GetComponent

`func (o *UserSetting) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *UserSetting) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *UserSetting) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetTitle

`func (o *UserSetting) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UserSetting) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UserSetting) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetConfigureUrl

`func (o *UserSetting) GetConfigureUrl() string`

GetConfigureUrl returns the ConfigureUrl field if non-nil, zero value otherwise.

### GetConfigureUrlOk

`func (o *UserSetting) GetConfigureUrlOk() (*string, bool)`

GetConfigureUrlOk returns a tuple with the ConfigureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureUrl

`func (o *UserSetting) SetConfigureUrl(v string)`

SetConfigureUrl sets ConfigureUrl field to given value.

### HasConfigureUrl

`func (o *UserSetting) HasConfigureUrl() bool`

HasConfigureUrl returns a boolean if a field has been set.

### GetIconUrl

`func (o *UserSetting) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *UserSetting) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *UserSetting) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *UserSetting) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



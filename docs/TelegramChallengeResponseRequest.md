# TelegramChallengeResponseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to **string** |  | [optional] [default to "ak-source-telegram"]
**Id** | **int32** |  | 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**PhotoUrl** | Pointer to **string** |  | [optional] 
**AuthDate** | **int32** |  | 
**Hash** | **string** |  | 

## Methods

### NewTelegramChallengeResponseRequest

`func NewTelegramChallengeResponseRequest(id int32, authDate int32, hash string, ) *TelegramChallengeResponseRequest`

NewTelegramChallengeResponseRequest instantiates a new TelegramChallengeResponseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelegramChallengeResponseRequestWithDefaults

`func NewTelegramChallengeResponseRequestWithDefaults() *TelegramChallengeResponseRequest`

NewTelegramChallengeResponseRequestWithDefaults instantiates a new TelegramChallengeResponseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *TelegramChallengeResponseRequest) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *TelegramChallengeResponseRequest) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *TelegramChallengeResponseRequest) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *TelegramChallengeResponseRequest) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetId

`func (o *TelegramChallengeResponseRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TelegramChallengeResponseRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TelegramChallengeResponseRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetFirstName

`func (o *TelegramChallengeResponseRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *TelegramChallengeResponseRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *TelegramChallengeResponseRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *TelegramChallengeResponseRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *TelegramChallengeResponseRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *TelegramChallengeResponseRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *TelegramChallengeResponseRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *TelegramChallengeResponseRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetUsername

`func (o *TelegramChallengeResponseRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TelegramChallengeResponseRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TelegramChallengeResponseRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TelegramChallengeResponseRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPhotoUrl

`func (o *TelegramChallengeResponseRequest) GetPhotoUrl() string`

GetPhotoUrl returns the PhotoUrl field if non-nil, zero value otherwise.

### GetPhotoUrlOk

`func (o *TelegramChallengeResponseRequest) GetPhotoUrlOk() (*string, bool)`

GetPhotoUrlOk returns a tuple with the PhotoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoUrl

`func (o *TelegramChallengeResponseRequest) SetPhotoUrl(v string)`

SetPhotoUrl sets PhotoUrl field to given value.

### HasPhotoUrl

`func (o *TelegramChallengeResponseRequest) HasPhotoUrl() bool`

HasPhotoUrl returns a boolean if a field has been set.

### GetAuthDate

`func (o *TelegramChallengeResponseRequest) GetAuthDate() int32`

GetAuthDate returns the AuthDate field if non-nil, zero value otherwise.

### GetAuthDateOk

`func (o *TelegramChallengeResponseRequest) GetAuthDateOk() (*int32, bool)`

GetAuthDateOk returns a tuple with the AuthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDate

`func (o *TelegramChallengeResponseRequest) SetAuthDate(v int32)`

SetAuthDate sets AuthDate field to given value.


### GetHash

`func (o *TelegramChallengeResponseRequest) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *TelegramChallengeResponseRequest) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *TelegramChallengeResponseRequest) SetHash(v string)`

SetHash sets Hash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



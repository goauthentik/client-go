# TelegramAuthRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**PhotoUrl** | Pointer to **string** |  | [optional] 
**AuthDate** | **int32** |  | 
**Hash** | **string** |  | 

## Methods

### NewTelegramAuthRequest

`func NewTelegramAuthRequest(id int32, authDate int32, hash string, ) *TelegramAuthRequest`

NewTelegramAuthRequest instantiates a new TelegramAuthRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelegramAuthRequestWithDefaults

`func NewTelegramAuthRequestWithDefaults() *TelegramAuthRequest`

NewTelegramAuthRequestWithDefaults instantiates a new TelegramAuthRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TelegramAuthRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TelegramAuthRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TelegramAuthRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetFirstName

`func (o *TelegramAuthRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *TelegramAuthRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *TelegramAuthRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *TelegramAuthRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *TelegramAuthRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *TelegramAuthRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *TelegramAuthRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *TelegramAuthRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetUsername

`func (o *TelegramAuthRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TelegramAuthRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TelegramAuthRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TelegramAuthRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPhotoUrl

`func (o *TelegramAuthRequest) GetPhotoUrl() string`

GetPhotoUrl returns the PhotoUrl field if non-nil, zero value otherwise.

### GetPhotoUrlOk

`func (o *TelegramAuthRequest) GetPhotoUrlOk() (*string, bool)`

GetPhotoUrlOk returns a tuple with the PhotoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoUrl

`func (o *TelegramAuthRequest) SetPhotoUrl(v string)`

SetPhotoUrl sets PhotoUrl field to given value.

### HasPhotoUrl

`func (o *TelegramAuthRequest) HasPhotoUrl() bool`

HasPhotoUrl returns a boolean if a field has been set.

### GetAuthDate

`func (o *TelegramAuthRequest) GetAuthDate() int32`

GetAuthDate returns the AuthDate field if non-nil, zero value otherwise.

### GetAuthDateOk

`func (o *TelegramAuthRequest) GetAuthDateOk() (*int32, bool)`

GetAuthDateOk returns a tuple with the AuthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDate

`func (o *TelegramAuthRequest) SetAuthDate(v int32)`

SetAuthDate sets AuthDate field to given value.


### GetHash

`func (o *TelegramAuthRequest) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *TelegramAuthRequest) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *TelegramAuthRequest) SetHash(v string)`

SetHash sets Hash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



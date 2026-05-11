# InvitationSendEmailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddresses** | **[]interface{}** |  | 
**CcAddresses** | Pointer to **[]interface{}** |  | [optional] 
**BccAddresses** | Pointer to **[]interface{}** |  | [optional] 
**Template** | Pointer to **string** |  | [optional] [default to "invitation"]

## Methods

### NewInvitationSendEmailRequest

`func NewInvitationSendEmailRequest(emailAddresses []interface{}, ) *InvitationSendEmailRequest`

NewInvitationSendEmailRequest instantiates a new InvitationSendEmailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationSendEmailRequestWithDefaults

`func NewInvitationSendEmailRequestWithDefaults() *InvitationSendEmailRequest`

NewInvitationSendEmailRequestWithDefaults instantiates a new InvitationSendEmailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddresses

`func (o *InvitationSendEmailRequest) GetEmailAddresses() []interface{}`

GetEmailAddresses returns the EmailAddresses field if non-nil, zero value otherwise.

### GetEmailAddressesOk

`func (o *InvitationSendEmailRequest) GetEmailAddressesOk() (*[]interface{}, bool)`

GetEmailAddressesOk returns a tuple with the EmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddresses

`func (o *InvitationSendEmailRequest) SetEmailAddresses(v []interface{})`

SetEmailAddresses sets EmailAddresses field to given value.


### GetCcAddresses

`func (o *InvitationSendEmailRequest) GetCcAddresses() []interface{}`

GetCcAddresses returns the CcAddresses field if non-nil, zero value otherwise.

### GetCcAddressesOk

`func (o *InvitationSendEmailRequest) GetCcAddressesOk() (*[]interface{}, bool)`

GetCcAddressesOk returns a tuple with the CcAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcAddresses

`func (o *InvitationSendEmailRequest) SetCcAddresses(v []interface{})`

SetCcAddresses sets CcAddresses field to given value.

### HasCcAddresses

`func (o *InvitationSendEmailRequest) HasCcAddresses() bool`

HasCcAddresses returns a boolean if a field has been set.

### GetBccAddresses

`func (o *InvitationSendEmailRequest) GetBccAddresses() []interface{}`

GetBccAddresses returns the BccAddresses field if non-nil, zero value otherwise.

### GetBccAddressesOk

`func (o *InvitationSendEmailRequest) GetBccAddressesOk() (*[]interface{}, bool)`

GetBccAddressesOk returns a tuple with the BccAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBccAddresses

`func (o *InvitationSendEmailRequest) SetBccAddresses(v []interface{})`

SetBccAddresses sets BccAddresses field to given value.

### HasBccAddresses

`func (o *InvitationSendEmailRequest) HasBccAddresses() bool`

HasBccAddresses returns a boolean if a field has been set.

### GetTemplate

`func (o *InvitationSendEmailRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *InvitationSendEmailRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *InvitationSendEmailRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *InvitationSendEmailRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



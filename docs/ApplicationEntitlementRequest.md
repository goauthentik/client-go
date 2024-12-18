# ApplicationEntitlementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**App** | **string** |  | 
**Attributes** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewApplicationEntitlementRequest

`func NewApplicationEntitlementRequest(name string, app string, ) *ApplicationEntitlementRequest`

NewApplicationEntitlementRequest instantiates a new ApplicationEntitlementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationEntitlementRequestWithDefaults

`func NewApplicationEntitlementRequestWithDefaults() *ApplicationEntitlementRequest`

NewApplicationEntitlementRequestWithDefaults instantiates a new ApplicationEntitlementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplicationEntitlementRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationEntitlementRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationEntitlementRequest) SetName(v string)`

SetName sets Name field to given value.


### GetApp

`func (o *ApplicationEntitlementRequest) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *ApplicationEntitlementRequest) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *ApplicationEntitlementRequest) SetApp(v string)`

SetApp sets App field to given value.


### GetAttributes

`func (o *ApplicationEntitlementRequest) GetAttributes() interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationEntitlementRequest) GetAttributesOk() (*interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApplicationEntitlementRequest) SetAttributes(v interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ApplicationEntitlementRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *ApplicationEntitlementRequest) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *ApplicationEntitlementRequest) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



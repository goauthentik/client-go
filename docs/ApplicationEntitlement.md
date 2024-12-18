# ApplicationEntitlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PbmUuid** | **string** |  | [readonly] 
**Name** | **string** |  | 
**App** | **string** |  | 
**Attributes** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewApplicationEntitlement

`func NewApplicationEntitlement(pbmUuid string, name string, app string, ) *ApplicationEntitlement`

NewApplicationEntitlement instantiates a new ApplicationEntitlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationEntitlementWithDefaults

`func NewApplicationEntitlementWithDefaults() *ApplicationEntitlement`

NewApplicationEntitlementWithDefaults instantiates a new ApplicationEntitlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPbmUuid

`func (o *ApplicationEntitlement) GetPbmUuid() string`

GetPbmUuid returns the PbmUuid field if non-nil, zero value otherwise.

### GetPbmUuidOk

`func (o *ApplicationEntitlement) GetPbmUuidOk() (*string, bool)`

GetPbmUuidOk returns a tuple with the PbmUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPbmUuid

`func (o *ApplicationEntitlement) SetPbmUuid(v string)`

SetPbmUuid sets PbmUuid field to given value.


### GetName

`func (o *ApplicationEntitlement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationEntitlement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationEntitlement) SetName(v string)`

SetName sets Name field to given value.


### GetApp

`func (o *ApplicationEntitlement) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *ApplicationEntitlement) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *ApplicationEntitlement) SetApp(v string)`

SetApp sets App field to given value.


### GetAttributes

`func (o *ApplicationEntitlement) GetAttributes() interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationEntitlement) GetAttributesOk() (*interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApplicationEntitlement) SetAttributes(v interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ApplicationEntitlement) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *ApplicationEntitlement) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *ApplicationEntitlement) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# BlueprintFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** |  | 
**LastM** | **time.Time** |  | 
**Hash** | **string** |  | 
**Meta** | [**Metadata**](Metadata.md) |  | [readonly] 

## Methods

### NewBlueprintFile

`func NewBlueprintFile(path string, lastM time.Time, hash string, meta Metadata, ) *BlueprintFile`

NewBlueprintFile instantiates a new BlueprintFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintFileWithDefaults

`func NewBlueprintFileWithDefaults() *BlueprintFile`

NewBlueprintFileWithDefaults instantiates a new BlueprintFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *BlueprintFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *BlueprintFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *BlueprintFile) SetPath(v string)`

SetPath sets Path field to given value.


### GetLastM

`func (o *BlueprintFile) GetLastM() time.Time`

GetLastM returns the LastM field if non-nil, zero value otherwise.

### GetLastMOk

`func (o *BlueprintFile) GetLastMOk() (*time.Time, bool)`

GetLastMOk returns a tuple with the LastM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastM

`func (o *BlueprintFile) SetLastM(v time.Time)`

SetLastM sets LastM field to given value.


### GetHash

`func (o *BlueprintFile) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *BlueprintFile) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *BlueprintFile) SetHash(v string)`

SetHash sets Hash field to given value.


### GetMeta

`func (o *BlueprintFile) GetMeta() Metadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BlueprintFile) GetMetaOk() (*Metadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BlueprintFile) SetMeta(v Metadata)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



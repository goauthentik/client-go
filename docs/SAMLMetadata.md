# SAMLMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | **string** |  | [readonly] 
**DownloadUrl** | **string** |  | [readonly] 

## Methods

### NewSAMLMetadata

`func NewSAMLMetadata(metadata string, downloadUrl string, ) *SAMLMetadata`

NewSAMLMetadata instantiates a new SAMLMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAMLMetadataWithDefaults

`func NewSAMLMetadataWithDefaults() *SAMLMetadata`

NewSAMLMetadataWithDefaults instantiates a new SAMLMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *SAMLMetadata) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SAMLMetadata) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SAMLMetadata) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.


### GetDownloadUrl

`func (o *SAMLMetadata) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *SAMLMetadata) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *SAMLMetadata) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



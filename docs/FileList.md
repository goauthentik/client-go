# FileList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**MimeType** | **string** |  | 
**Url** | **string** |  | 
**ThemedUrls** | Pointer to [**NullableThemedUrls**](ThemedUrls.md) |  | [optional] 

## Methods

### NewFileList

`func NewFileList(name string, mimeType string, url string, ) *FileList`

NewFileList instantiates a new FileList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileListWithDefaults

`func NewFileListWithDefaults() *FileList`

NewFileListWithDefaults instantiates a new FileList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FileList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileList) SetName(v string)`

SetName sets Name field to given value.


### GetMimeType

`func (o *FileList) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *FileList) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *FileList) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetUrl

`func (o *FileList) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FileList) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FileList) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetThemedUrls

`func (o *FileList) GetThemedUrls() ThemedUrls`

GetThemedUrls returns the ThemedUrls field if non-nil, zero value otherwise.

### GetThemedUrlsOk

`func (o *FileList) GetThemedUrlsOk() (*ThemedUrls, bool)`

GetThemedUrlsOk returns a tuple with the ThemedUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemedUrls

`func (o *FileList) SetThemedUrls(v ThemedUrls)`

SetThemedUrls sets ThemedUrls field to given value.

### HasThemedUrls

`func (o *FileList) HasThemedUrls() bool`

HasThemedUrls returns a boolean if a field has been set.

### SetThemedUrlsNil

`func (o *FileList) SetThemedUrlsNil(b bool)`

 SetThemedUrlsNil sets the value for ThemedUrls to be an explicit nil

### UnsetThemedUrls
`func (o *FileList) UnsetThemedUrls()`

UnsetThemedUrls ensures that no value is present for ThemedUrls, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



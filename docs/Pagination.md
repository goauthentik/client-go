# Pagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | **float32** |  | 
**Previous** | **float32** |  | 
**Count** | **float32** |  | 
**Current** | **float32** |  | 
**TotalPages** | **float32** |  | 
**StartIndex** | **float32** |  | 
**EndIndex** | **float32** |  | 

## Methods

### NewPagination

`func NewPagination(next float32, previous float32, count float32, current float32, totalPages float32, startIndex float32, endIndex float32, ) *Pagination`

NewPagination instantiates a new Pagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationWithDefaults

`func NewPaginationWithDefaults() *Pagination`

NewPaginationWithDefaults instantiates a new Pagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *Pagination) GetNext() float32`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *Pagination) GetNextOk() (*float32, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *Pagination) SetNext(v float32)`

SetNext sets Next field to given value.


### GetPrevious

`func (o *Pagination) GetPrevious() float32`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *Pagination) GetPreviousOk() (*float32, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *Pagination) SetPrevious(v float32)`

SetPrevious sets Previous field to given value.


### GetCount

`func (o *Pagination) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Pagination) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Pagination) SetCount(v float32)`

SetCount sets Count field to given value.


### GetCurrent

`func (o *Pagination) GetCurrent() float32`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *Pagination) GetCurrentOk() (*float32, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *Pagination) SetCurrent(v float32)`

SetCurrent sets Current field to given value.


### GetTotalPages

`func (o *Pagination) GetTotalPages() float32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *Pagination) GetTotalPagesOk() (*float32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *Pagination) SetTotalPages(v float32)`

SetTotalPages sets TotalPages field to given value.


### GetStartIndex

`func (o *Pagination) GetStartIndex() float32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *Pagination) GetStartIndexOk() (*float32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *Pagination) SetStartIndex(v float32)`

SetStartIndex sets StartIndex field to given value.


### GetEndIndex

`func (o *Pagination) GetEndIndex() float32`

GetEndIndex returns the EndIndex field if non-nil, zero value otherwise.

### GetEndIndexOk

`func (o *Pagination) GetEndIndexOk() (*float32, bool)`

GetEndIndexOk returns a tuple with the EndIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndIndex

`func (o *Pagination) SetEndIndex(v float32)`

SetEndIndex sets EndIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



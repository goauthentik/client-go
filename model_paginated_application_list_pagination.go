/*
authentik

Making authentication simple.

API version: 2022.5.1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedApplicationListPagination struct for PaginatedApplicationListPagination
type PaginatedApplicationListPagination struct {
	Next       float32 `json:"next"`
	Previous   float32 `json:"previous"`
	Count      float32 `json:"count"`
	Current    float32 `json:"current"`
	TotalPages float32 `json:"total_pages"`
	StartIndex float32 `json:"start_index"`
	EndIndex   float32 `json:"end_index"`
}

// NewPaginatedApplicationListPagination instantiates a new PaginatedApplicationListPagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedApplicationListPagination(next float32, previous float32, count float32, current float32, totalPages float32, startIndex float32, endIndex float32) *PaginatedApplicationListPagination {
	this := PaginatedApplicationListPagination{}
	this.Next = next
	this.Previous = previous
	this.Count = count
	this.Current = current
	this.TotalPages = totalPages
	this.StartIndex = startIndex
	this.EndIndex = endIndex
	return &this
}

// NewPaginatedApplicationListPaginationWithDefaults instantiates a new PaginatedApplicationListPagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedApplicationListPaginationWithDefaults() *PaginatedApplicationListPagination {
	this := PaginatedApplicationListPagination{}
	return &this
}

// GetNext returns the Next field value
func (o *PaginatedApplicationListPagination) GetNext() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Next
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
func (o *PaginatedApplicationListPagination) GetNextOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Next, true
}

// SetNext sets field value
func (o *PaginatedApplicationListPagination) SetNext(v float32) {
	o.Next = v
}

// GetPrevious returns the Previous field value
func (o *PaginatedApplicationListPagination) GetPrevious() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value
// and a boolean to check if the value has been set.
func (o *PaginatedApplicationListPagination) GetPreviousOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Previous, true
}

// SetPrevious sets field value
func (o *PaginatedApplicationListPagination) SetPrevious(v float32) {
	o.Previous = v
}

// GetCount returns the Count field value
func (o *PaginatedApplicationListPagination) GetCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *PaginatedApplicationListPagination) GetCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *PaginatedApplicationListPagination) SetCount(v float32) {
	o.Count = v
}

// GetCurrent returns the Current field value
func (o *PaginatedApplicationListPagination) GetCurrent() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Current
}

// GetCurrentOk returns a tuple with the Current field value
// and a boolean to check if the value has been set.
func (o *PaginatedApplicationListPagination) GetCurrentOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Current, true
}

// SetCurrent sets field value
func (o *PaginatedApplicationListPagination) SetCurrent(v float32) {
	o.Current = v
}

// GetTotalPages returns the TotalPages field value
func (o *PaginatedApplicationListPagination) GetTotalPages() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value
// and a boolean to check if the value has been set.
func (o *PaginatedApplicationListPagination) GetTotalPagesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPages, true
}

// SetTotalPages sets field value
func (o *PaginatedApplicationListPagination) SetTotalPages(v float32) {
	o.TotalPages = v
}

// GetStartIndex returns the StartIndex field value
func (o *PaginatedApplicationListPagination) GetStartIndex() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value
// and a boolean to check if the value has been set.
func (o *PaginatedApplicationListPagination) GetStartIndexOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartIndex, true
}

// SetStartIndex sets field value
func (o *PaginatedApplicationListPagination) SetStartIndex(v float32) {
	o.StartIndex = v
}

// GetEndIndex returns the EndIndex field value
func (o *PaginatedApplicationListPagination) GetEndIndex() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.EndIndex
}

// GetEndIndexOk returns a tuple with the EndIndex field value
// and a boolean to check if the value has been set.
func (o *PaginatedApplicationListPagination) GetEndIndexOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndIndex, true
}

// SetEndIndex sets field value
func (o *PaginatedApplicationListPagination) SetEndIndex(v float32) {
	o.EndIndex = v
}

func (o PaginatedApplicationListPagination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["next"] = o.Next
	}
	if true {
		toSerialize["previous"] = o.Previous
	}
	if true {
		toSerialize["count"] = o.Count
	}
	if true {
		toSerialize["current"] = o.Current
	}
	if true {
		toSerialize["total_pages"] = o.TotalPages
	}
	if true {
		toSerialize["start_index"] = o.StartIndex
	}
	if true {
		toSerialize["end_index"] = o.EndIndex
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedApplicationListPagination struct {
	value *PaginatedApplicationListPagination
	isSet bool
}

func (v NullablePaginatedApplicationListPagination) Get() *PaginatedApplicationListPagination {
	return v.value
}

func (v *NullablePaginatedApplicationListPagination) Set(val *PaginatedApplicationListPagination) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedApplicationListPagination) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedApplicationListPagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedApplicationListPagination(val *PaginatedApplicationListPagination) *NullablePaginatedApplicationListPagination {
	return &NullablePaginatedApplicationListPagination{value: val, isSet: true}
}

func (v NullablePaginatedApplicationListPagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedApplicationListPagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

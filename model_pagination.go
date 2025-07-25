/*
authentik

Making authentication simple.

API version: 2025.6.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Pagination struct for Pagination
type Pagination struct {
	Next       float32 `json:"next"`
	Previous   float32 `json:"previous"`
	Count      float32 `json:"count"`
	Current    float32 `json:"current"`
	TotalPages float32 `json:"total_pages"`
	StartIndex float32 `json:"start_index"`
	EndIndex   float32 `json:"end_index"`
}

// NewPagination instantiates a new Pagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagination(next float32, previous float32, count float32, current float32, totalPages float32, startIndex float32, endIndex float32) *Pagination {
	this := Pagination{}
	this.Next = next
	this.Previous = previous
	this.Count = count
	this.Current = current
	this.TotalPages = totalPages
	this.StartIndex = startIndex
	this.EndIndex = endIndex
	return &this
}

// NewPaginationWithDefaults instantiates a new Pagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationWithDefaults() *Pagination {
	this := Pagination{}
	return &this
}

// GetNext returns the Next field value
func (o *Pagination) GetNext() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Next
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
func (o *Pagination) GetNextOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Next, true
}

// SetNext sets field value
func (o *Pagination) SetNext(v float32) {
	o.Next = v
}

// GetPrevious returns the Previous field value
func (o *Pagination) GetPrevious() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value
// and a boolean to check if the value has been set.
func (o *Pagination) GetPreviousOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Previous, true
}

// SetPrevious sets field value
func (o *Pagination) SetPrevious(v float32) {
	o.Previous = v
}

// GetCount returns the Count field value
func (o *Pagination) GetCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *Pagination) GetCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *Pagination) SetCount(v float32) {
	o.Count = v
}

// GetCurrent returns the Current field value
func (o *Pagination) GetCurrent() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Current
}

// GetCurrentOk returns a tuple with the Current field value
// and a boolean to check if the value has been set.
func (o *Pagination) GetCurrentOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Current, true
}

// SetCurrent sets field value
func (o *Pagination) SetCurrent(v float32) {
	o.Current = v
}

// GetTotalPages returns the TotalPages field value
func (o *Pagination) GetTotalPages() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value
// and a boolean to check if the value has been set.
func (o *Pagination) GetTotalPagesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPages, true
}

// SetTotalPages sets field value
func (o *Pagination) SetTotalPages(v float32) {
	o.TotalPages = v
}

// GetStartIndex returns the StartIndex field value
func (o *Pagination) GetStartIndex() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value
// and a boolean to check if the value has been set.
func (o *Pagination) GetStartIndexOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartIndex, true
}

// SetStartIndex sets field value
func (o *Pagination) SetStartIndex(v float32) {
	o.StartIndex = v
}

// GetEndIndex returns the EndIndex field value
func (o *Pagination) GetEndIndex() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.EndIndex
}

// GetEndIndexOk returns a tuple with the EndIndex field value
// and a boolean to check if the value has been set.
func (o *Pagination) GetEndIndexOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndIndex, true
}

// SetEndIndex sets field value
func (o *Pagination) SetEndIndex(v float32) {
	o.EndIndex = v
}

func (o Pagination) MarshalJSON() ([]byte, error) {
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

type NullablePagination struct {
	value *Pagination
	isSet bool
}

func (v NullablePagination) Get() *Pagination {
	return v.value
}

func (v *NullablePagination) Set(val *Pagination) {
	v.value = val
	v.isSet = true
}

func (v NullablePagination) IsSet() bool {
	return v.isSet
}

func (v *NullablePagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagination(val *Pagination) *NullablePagination {
	return &NullablePagination{value: val, isSet: true}
}

func (v NullablePagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

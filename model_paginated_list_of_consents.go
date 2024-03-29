/*
My Data My Consent - Developer API

Unleashing the power of consent by establishing trust. The Platform Core Developer API defines a set of capabilities that can be used to request, issue, manage and update data, documents and credentials by organizations. The API can be used to request, manage and update Decentralised Identifiers, Financial Data, Health Data issue Documents, Credentials directly or using OpenID Connect flows, and verify Messages signed with DIDs and much more.

API version: 1.0
Contact: support@mydatamyconsent.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github.com/mydatamyconsent/sdk

import (
	"encoding/json"
)

// PaginatedListOfConsents struct for PaginatedListOfConsents
type PaginatedListOfConsents struct {
	PageNo int32 `json:"pageNo"`
	PageSize int32 `json:"pageSize"`
	TotalPage int32 `json:"totalPage"`
	Items []Consent `json:"items"`
}

// NewPaginatedListOfConsents instantiates a new PaginatedListOfConsents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedListOfConsents(pageNo int32, pageSize int32, totalPage int32, items []Consent) *PaginatedListOfConsents {
	this := PaginatedListOfConsents{}
	this.PageNo = pageNo
	this.PageSize = pageSize
	this.TotalPage = totalPage
	this.Items = items
	return &this
}

// NewPaginatedListOfConsentsWithDefaults instantiates a new PaginatedListOfConsents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedListOfConsentsWithDefaults() *PaginatedListOfConsents {
	this := PaginatedListOfConsents{}
	return &this
}

// GetPageNo returns the PageNo field value
func (o *PaginatedListOfConsents) GetPageNo() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageNo
}

// GetPageNoOk returns a tuple with the PageNo field value
// and a boolean to check if the value has been set.
func (o *PaginatedListOfConsents) GetPageNoOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageNo, true
}

// SetPageNo sets field value
func (o *PaginatedListOfConsents) SetPageNo(v int32) {
	o.PageNo = v
}

// GetPageSize returns the PageSize field value
func (o *PaginatedListOfConsents) GetPageSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *PaginatedListOfConsents) GetPageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *PaginatedListOfConsents) SetPageSize(v int32) {
	o.PageSize = v
}

// GetTotalPage returns the TotalPage field value
func (o *PaginatedListOfConsents) GetTotalPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalPage
}

// GetTotalPageOk returns a tuple with the TotalPage field value
// and a boolean to check if the value has been set.
func (o *PaginatedListOfConsents) GetTotalPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPage, true
}

// SetTotalPage sets field value
func (o *PaginatedListOfConsents) SetTotalPage(v int32) {
	o.TotalPage = v
}

// GetItems returns the Items field value
func (o *PaginatedListOfConsents) GetItems() []Consent {
	if o == nil {
		var ret []Consent
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *PaginatedListOfConsents) GetItemsOk() ([]Consent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *PaginatedListOfConsents) SetItems(v []Consent) {
	o.Items = v
}

func (o PaginatedListOfConsents) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pageNo"] = o.PageNo
	}
	if true {
		toSerialize["pageSize"] = o.PageSize
	}
	if true {
		toSerialize["totalPage"] = o.TotalPage
	}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedListOfConsents struct {
	value *PaginatedListOfConsents
	isSet bool
}

func (v NullablePaginatedListOfConsents) Get() *PaginatedListOfConsents {
	return v.value
}

func (v *NullablePaginatedListOfConsents) Set(val *PaginatedListOfConsents) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedListOfConsents) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedListOfConsents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedListOfConsents(val *PaginatedListOfConsents) *NullablePaginatedListOfConsents {
	return &NullablePaginatedListOfConsents{value: val, isSet: true}
}

func (v NullablePaginatedListOfConsents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedListOfConsents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



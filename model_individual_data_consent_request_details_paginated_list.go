/*
My Data My Consent - Developer API

Unleashing the power of data consent by establishing trust. The Platform Core Developer API defines a set of capabilities that can be used to request, issue, manage and update data, documents and credentials by organizations. The API can be used to request, manage and update Decentralised Identifiers, Financial Data, Health Data issue Documents, Credentials directly or using OpenID Connect flows, and verify Messages signed with DIDs and much more.

API version: v1
Contact: support@mydatamyconsent.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github.com/mydatamyconsent/sdk

import (
	"encoding/json"
)

// IndividualDataConsentRequestDetailsPaginatedList struct for IndividualDataConsentRequestDetailsPaginatedList
type IndividualDataConsentRequestDetailsPaginatedList struct {
	PageIndex *int32 `json:"pageIndex,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty"`
	TotalPages *int32 `json:"totalPages,omitempty"`
	TotalItems *int64 `json:"totalItems,omitempty"`
	Items []IndividualDataConsentRequestDetails `json:"items,omitempty"`
}

// NewIndividualDataConsentRequestDetailsPaginatedList instantiates a new IndividualDataConsentRequestDetailsPaginatedList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndividualDataConsentRequestDetailsPaginatedList() *IndividualDataConsentRequestDetailsPaginatedList {
	this := IndividualDataConsentRequestDetailsPaginatedList{}
	return &this
}

// NewIndividualDataConsentRequestDetailsPaginatedListWithDefaults instantiates a new IndividualDataConsentRequestDetailsPaginatedList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndividualDataConsentRequestDetailsPaginatedListWithDefaults() *IndividualDataConsentRequestDetailsPaginatedList {
	this := IndividualDataConsentRequestDetailsPaginatedList{}
	return &this
}

// GetPageIndex returns the PageIndex field value if set, zero value otherwise.
func (o *IndividualDataConsentRequestDetailsPaginatedList) GetPageIndex() int32 {
	if o == nil || o.PageIndex == nil {
		var ret int32
		return ret
	}
	return *o.PageIndex
}

// GetPageIndexOk returns a tuple with the PageIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualDataConsentRequestDetailsPaginatedList) GetPageIndexOk() (*int32, bool) {
	if o == nil || o.PageIndex == nil {
		return nil, false
	}
	return o.PageIndex, true
}

// HasPageIndex returns a boolean if a field has been set.
func (o *IndividualDataConsentRequestDetailsPaginatedList) HasPageIndex() bool {
	if o != nil && o.PageIndex != nil {
		return true
	}

	return false
}

// SetPageIndex gets a reference to the given int32 and assigns it to the PageIndex field.
func (o *IndividualDataConsentRequestDetailsPaginatedList) SetPageIndex(v int32) {
	o.PageIndex = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *IndividualDataConsentRequestDetailsPaginatedList) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualDataConsentRequestDetailsPaginatedList) GetPageSizeOk() (*int32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *IndividualDataConsentRequestDetailsPaginatedList) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *IndividualDataConsentRequestDetailsPaginatedList) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *IndividualDataConsentRequestDetailsPaginatedList) GetTotalPages() int32 {
	if o == nil || o.TotalPages == nil {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualDataConsentRequestDetailsPaginatedList) GetTotalPagesOk() (*int32, bool) {
	if o == nil || o.TotalPages == nil {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *IndividualDataConsentRequestDetailsPaginatedList) HasTotalPages() bool {
	if o != nil && o.TotalPages != nil {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *IndividualDataConsentRequestDetailsPaginatedList) SetTotalPages(v int32) {
	o.TotalPages = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *IndividualDataConsentRequestDetailsPaginatedList) GetTotalItems() int64 {
	if o == nil || o.TotalItems == nil {
		var ret int64
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualDataConsentRequestDetailsPaginatedList) GetTotalItemsOk() (*int64, bool) {
	if o == nil || o.TotalItems == nil {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *IndividualDataConsentRequestDetailsPaginatedList) HasTotalItems() bool {
	if o != nil && o.TotalItems != nil {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int64 and assigns it to the TotalItems field.
func (o *IndividualDataConsentRequestDetailsPaginatedList) SetTotalItems(v int64) {
	o.TotalItems = &v
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndividualDataConsentRequestDetailsPaginatedList) GetItems() []IndividualDataConsentRequestDetails {
	if o == nil  {
		var ret []IndividualDataConsentRequestDetails
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndividualDataConsentRequestDetailsPaginatedList) GetItemsOk() ([]IndividualDataConsentRequestDetails, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *IndividualDataConsentRequestDetailsPaginatedList) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []IndividualDataConsentRequestDetails and assigns it to the Items field.
func (o *IndividualDataConsentRequestDetailsPaginatedList) SetItems(v []IndividualDataConsentRequestDetails) {
	o.Items = v
}

func (o IndividualDataConsentRequestDetailsPaginatedList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PageIndex != nil {
		toSerialize["pageIndex"] = o.PageIndex
	}
	if o.PageSize != nil {
		toSerialize["pageSize"] = o.PageSize
	}
	if o.TotalPages != nil {
		toSerialize["totalPages"] = o.TotalPages
	}
	if o.TotalItems != nil {
		toSerialize["totalItems"] = o.TotalItems
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableIndividualDataConsentRequestDetailsPaginatedList struct {
	value *IndividualDataConsentRequestDetailsPaginatedList
	isSet bool
}

func (v NullableIndividualDataConsentRequestDetailsPaginatedList) Get() *IndividualDataConsentRequestDetailsPaginatedList {
	return v.value
}

func (v *NullableIndividualDataConsentRequestDetailsPaginatedList) Set(val *IndividualDataConsentRequestDetailsPaginatedList) {
	v.value = val
	v.isSet = true
}

func (v NullableIndividualDataConsentRequestDetailsPaginatedList) IsSet() bool {
	return v.isSet
}

func (v *NullableIndividualDataConsentRequestDetailsPaginatedList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndividualDataConsentRequestDetailsPaginatedList(val *IndividualDataConsentRequestDetailsPaginatedList) *NullableIndividualDataConsentRequestDetailsPaginatedList {
	return &NullableIndividualDataConsentRequestDetailsPaginatedList{value: val, isSet: true}
}

func (v NullableIndividualDataConsentRequestDetailsPaginatedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndividualDataConsentRequestDetailsPaginatedList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



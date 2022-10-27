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

// BondTransaction struct for BondTransaction
type BondTransaction struct {
	Id string `json:"id"`
}

// NewBondTransaction instantiates a new BondTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBondTransaction(id string) *BondTransaction {
	this := BondTransaction{}
	this.Id = id
	return &this
}

// NewBondTransactionWithDefaults instantiates a new BondTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBondTransactionWithDefaults() *BondTransaction {
	this := BondTransaction{}
	return &this
}

// GetId returns the Id field value
func (o *BondTransaction) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BondTransaction) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BondTransaction) SetId(v string) {
	o.Id = v
}

func (o BondTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableBondTransaction struct {
	value *BondTransaction
	isSet bool
}

func (v NullableBondTransaction) Get() *BondTransaction {
	return v.value
}

func (v *NullableBondTransaction) Set(val *BondTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableBondTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableBondTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBondTransaction(val *BondTransaction) *NullableBondTransaction {
	return &NullableBondTransaction{value: val, isSet: true}
}

func (v NullableBondTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBondTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


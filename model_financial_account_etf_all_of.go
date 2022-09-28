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

// FinancialAccountEtfAllOf struct for FinancialAccountEtfAllOf
type FinancialAccountEtfAllOf struct {
	Type string `json:"type"`
}

// NewFinancialAccountEtfAllOf instantiates a new FinancialAccountEtfAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialAccountEtfAllOf(type_ string) *FinancialAccountEtfAllOf {
	this := FinancialAccountEtfAllOf{}
	this.Type = type_
	return &this
}

// NewFinancialAccountEtfAllOfWithDefaults instantiates a new FinancialAccountEtfAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialAccountEtfAllOfWithDefaults() *FinancialAccountEtfAllOf {
	this := FinancialAccountEtfAllOf{}
	return &this
}

// GetType returns the Type field value
func (o *FinancialAccountEtfAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountEtfAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FinancialAccountEtfAllOf) SetType(v string) {
	o.Type = v
}

func (o FinancialAccountEtfAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableFinancialAccountEtfAllOf struct {
	value *FinancialAccountEtfAllOf
	isSet bool
}

func (v NullableFinancialAccountEtfAllOf) Get() *FinancialAccountEtfAllOf {
	return v.value
}

func (v *NullableFinancialAccountEtfAllOf) Set(val *FinancialAccountEtfAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialAccountEtfAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialAccountEtfAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialAccountEtfAllOf(val *FinancialAccountEtfAllOf) *NullableFinancialAccountEtfAllOf {
	return &NullableFinancialAccountEtfAllOf{value: val, isSet: true}
}

func (v NullableFinancialAccountEtfAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialAccountEtfAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



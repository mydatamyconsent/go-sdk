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

// FinancialAccountTermDepositAllOf struct for FinancialAccountTermDepositAllOf
type FinancialAccountTermDepositAllOf struct {
	Type string `json:"type"`
}

// NewFinancialAccountTermDepositAllOf instantiates a new FinancialAccountTermDepositAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialAccountTermDepositAllOf(type_ string) *FinancialAccountTermDepositAllOf {
	this := FinancialAccountTermDepositAllOf{}
	this.Type = type_
	return &this
}

// NewFinancialAccountTermDepositAllOfWithDefaults instantiates a new FinancialAccountTermDepositAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialAccountTermDepositAllOfWithDefaults() *FinancialAccountTermDepositAllOf {
	this := FinancialAccountTermDepositAllOf{}
	return &this
}

// GetType returns the Type field value
func (o *FinancialAccountTermDepositAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountTermDepositAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FinancialAccountTermDepositAllOf) SetType(v string) {
	o.Type = v
}

func (o FinancialAccountTermDepositAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableFinancialAccountTermDepositAllOf struct {
	value *FinancialAccountTermDepositAllOf
	isSet bool
}

func (v NullableFinancialAccountTermDepositAllOf) Get() *FinancialAccountTermDepositAllOf {
	return v.value
}

func (v *NullableFinancialAccountTermDepositAllOf) Set(val *FinancialAccountTermDepositAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialAccountTermDepositAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialAccountTermDepositAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialAccountTermDepositAllOf(val *FinancialAccountTermDepositAllOf) *NullableFinancialAccountTermDepositAllOf {
	return &NullableFinancialAccountTermDepositAllOf{value: val, isSet: true}
}

func (v NullableFinancialAccountTermDepositAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialAccountTermDepositAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



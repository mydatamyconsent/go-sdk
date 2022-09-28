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

// FinancialAccountBillPaymentAllOf struct for FinancialAccountBillPaymentAllOf
type FinancialAccountBillPaymentAllOf struct {
	Type string `json:"type"`
}

// NewFinancialAccountBillPaymentAllOf instantiates a new FinancialAccountBillPaymentAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialAccountBillPaymentAllOf(type_ string) *FinancialAccountBillPaymentAllOf {
	this := FinancialAccountBillPaymentAllOf{}
	this.Type = type_
	return &this
}

// NewFinancialAccountBillPaymentAllOfWithDefaults instantiates a new FinancialAccountBillPaymentAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialAccountBillPaymentAllOfWithDefaults() *FinancialAccountBillPaymentAllOf {
	this := FinancialAccountBillPaymentAllOf{}
	return &this
}

// GetType returns the Type field value
func (o *FinancialAccountBillPaymentAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountBillPaymentAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FinancialAccountBillPaymentAllOf) SetType(v string) {
	o.Type = v
}

func (o FinancialAccountBillPaymentAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableFinancialAccountBillPaymentAllOf struct {
	value *FinancialAccountBillPaymentAllOf
	isSet bool
}

func (v NullableFinancialAccountBillPaymentAllOf) Get() *FinancialAccountBillPaymentAllOf {
	return v.value
}

func (v *NullableFinancialAccountBillPaymentAllOf) Set(val *FinancialAccountBillPaymentAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialAccountBillPaymentAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialAccountBillPaymentAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialAccountBillPaymentAllOf(val *FinancialAccountBillPaymentAllOf) *NullableFinancialAccountBillPaymentAllOf {
	return &NullableFinancialAccountBillPaymentAllOf{value: val, isSet: true}
}

func (v NullableFinancialAccountBillPaymentAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialAccountBillPaymentAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// FinancialAccountTransactionMutualFundTransactionAllOf struct for FinancialAccountTransactionMutualFundTransactionAllOf
type FinancialAccountTransactionMutualFundTransactionAllOf struct {
	Type string `json:"type"`
}

// NewFinancialAccountTransactionMutualFundTransactionAllOf instantiates a new FinancialAccountTransactionMutualFundTransactionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialAccountTransactionMutualFundTransactionAllOf(type_ string) *FinancialAccountTransactionMutualFundTransactionAllOf {
	this := FinancialAccountTransactionMutualFundTransactionAllOf{}
	this.Type = type_
	return &this
}

// NewFinancialAccountTransactionMutualFundTransactionAllOfWithDefaults instantiates a new FinancialAccountTransactionMutualFundTransactionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialAccountTransactionMutualFundTransactionAllOfWithDefaults() *FinancialAccountTransactionMutualFundTransactionAllOf {
	this := FinancialAccountTransactionMutualFundTransactionAllOf{}
	return &this
}

// GetType returns the Type field value
func (o *FinancialAccountTransactionMutualFundTransactionAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountTransactionMutualFundTransactionAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FinancialAccountTransactionMutualFundTransactionAllOf) SetType(v string) {
	o.Type = v
}

func (o FinancialAccountTransactionMutualFundTransactionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableFinancialAccountTransactionMutualFundTransactionAllOf struct {
	value *FinancialAccountTransactionMutualFundTransactionAllOf
	isSet bool
}

func (v NullableFinancialAccountTransactionMutualFundTransactionAllOf) Get() *FinancialAccountTransactionMutualFundTransactionAllOf {
	return v.value
}

func (v *NullableFinancialAccountTransactionMutualFundTransactionAllOf) Set(val *FinancialAccountTransactionMutualFundTransactionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialAccountTransactionMutualFundTransactionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialAccountTransactionMutualFundTransactionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialAccountTransactionMutualFundTransactionAllOf(val *FinancialAccountTransactionMutualFundTransactionAllOf) *NullableFinancialAccountTransactionMutualFundTransactionAllOf {
	return &NullableFinancialAccountTransactionMutualFundTransactionAllOf{value: val, isSet: true}
}

func (v NullableFinancialAccountTransactionMutualFundTransactionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialAccountTransactionMutualFundTransactionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


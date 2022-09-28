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
	"time"
)

// FinancialAccountTransactionPeriod FinancialAccountTransactionPeriod : Financial account transaction period.
type FinancialAccountTransactionPeriod struct {
	From time.Time `json:"from"`
	To time.Time `json:"to"`
}

// NewFinancialAccountTransactionPeriod instantiates a new FinancialAccountTransactionPeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialAccountTransactionPeriod(from time.Time, to time.Time) *FinancialAccountTransactionPeriod {
	this := FinancialAccountTransactionPeriod{}
	this.From = from
	this.To = to
	return &this
}

// NewFinancialAccountTransactionPeriodWithDefaults instantiates a new FinancialAccountTransactionPeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialAccountTransactionPeriodWithDefaults() *FinancialAccountTransactionPeriod {
	this := FinancialAccountTransactionPeriod{}
	return &this
}

// GetFrom returns the From field value
func (o *FinancialAccountTransactionPeriod) GetFrom() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountTransactionPeriod) GetFromOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *FinancialAccountTransactionPeriod) SetFrom(v time.Time) {
	o.From = v
}

// GetTo returns the To field value
func (o *FinancialAccountTransactionPeriod) GetTo() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountTransactionPeriod) GetToOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *FinancialAccountTransactionPeriod) SetTo(v time.Time) {
	o.To = v
}

func (o FinancialAccountTransactionPeriod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["from"] = o.From
	}
	if true {
		toSerialize["to"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullableFinancialAccountTransactionPeriod struct {
	value *FinancialAccountTransactionPeriod
	isSet bool
}

func (v NullableFinancialAccountTransactionPeriod) Get() *FinancialAccountTransactionPeriod {
	return v.value
}

func (v *NullableFinancialAccountTransactionPeriod) Set(val *FinancialAccountTransactionPeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialAccountTransactionPeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialAccountTransactionPeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialAccountTransactionPeriod(val *FinancialAccountTransactionPeriod) *NullableFinancialAccountTransactionPeriod {
	return &NullableFinancialAccountTransactionPeriod{value: val, isSet: true}
}

func (v NullableFinancialAccountTransactionPeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialAccountTransactionPeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



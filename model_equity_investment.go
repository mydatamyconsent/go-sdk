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

// EquityInvestment struct for EquityInvestment
type EquityInvestment struct {
	Holdings EquityHoldings `json:"holdings"`
}

// NewEquityInvestment instantiates a new EquityInvestment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquityInvestment(holdings EquityHoldings) *EquityInvestment {
	this := EquityInvestment{}
	this.Holdings = holdings
	return &this
}

// NewEquityInvestmentWithDefaults instantiates a new EquityInvestment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquityInvestmentWithDefaults() *EquityInvestment {
	this := EquityInvestment{}
	return &this
}

// GetHoldings returns the Holdings field value
func (o *EquityInvestment) GetHoldings() EquityHoldings {
	if o == nil {
		var ret EquityHoldings
		return ret
	}

	return o.Holdings
}

// GetHoldingsOk returns a tuple with the Holdings field value
// and a boolean to check if the value has been set.
func (o *EquityInvestment) GetHoldingsOk() (*EquityHoldings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Holdings, true
}

// SetHoldings sets field value
func (o *EquityInvestment) SetHoldings(v EquityHoldings) {
	o.Holdings = v
}

func (o EquityInvestment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["holdings"] = o.Holdings
	}
	return json.Marshal(toSerialize)
}

type NullableEquityInvestment struct {
	value *EquityInvestment
	isSet bool
}

func (v NullableEquityInvestment) Get() *EquityInvestment {
	return v.value
}

func (v *NullableEquityInvestment) Set(val *EquityInvestment) {
	v.value = val
	v.isSet = true
}

func (v NullableEquityInvestment) IsSet() bool {
	return v.isSet
}

func (v *NullableEquityInvestment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquityInvestment(val *EquityInvestment) *NullableEquityInvestment {
	return &NullableEquityInvestment{value: val, isSet: true}
}

func (v NullableEquityInvestment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquityInvestment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// MutualFundHoldings struct for MutualFundHoldings
type MutualFundHoldings struct {
	Holding MutualFundHolding `json:"holding"`
}

// NewMutualFundHoldings instantiates a new MutualFundHoldings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMutualFundHoldings(holding MutualFundHolding) *MutualFundHoldings {
	this := MutualFundHoldings{}
	this.Holding = holding
	return &this
}

// NewMutualFundHoldingsWithDefaults instantiates a new MutualFundHoldings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMutualFundHoldingsWithDefaults() *MutualFundHoldings {
	this := MutualFundHoldings{}
	return &this
}

// GetHolding returns the Holding field value
func (o *MutualFundHoldings) GetHolding() MutualFundHolding {
	if o == nil {
		var ret MutualFundHolding
		return ret
	}

	return o.Holding
}

// GetHoldingOk returns a tuple with the Holding field value
// and a boolean to check if the value has been set.
func (o *MutualFundHoldings) GetHoldingOk() (*MutualFundHolding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Holding, true
}

// SetHolding sets field value
func (o *MutualFundHoldings) SetHolding(v MutualFundHolding) {
	o.Holding = v
}

func (o MutualFundHoldings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["holding"] = o.Holding
	}
	return json.Marshal(toSerialize)
}

type NullableMutualFundHoldings struct {
	value *MutualFundHoldings
	isSet bool
}

func (v NullableMutualFundHoldings) Get() *MutualFundHoldings {
	return v.value
}

func (v *NullableMutualFundHoldings) Set(val *MutualFundHoldings) {
	v.value = val
	v.isSet = true
}

func (v NullableMutualFundHoldings) IsSet() bool {
	return v.isSet
}

func (v *NullableMutualFundHoldings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMutualFundHoldings(val *MutualFundHoldings) *NullableMutualFundHoldings {
	return &NullableMutualFundHoldings{value: val, isSet: true}
}

func (v NullableMutualFundHoldings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMutualFundHoldings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



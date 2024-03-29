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
	"fmt"
)

// TermDepositInterestComputation the model 'TermDepositInterestComputation'
type TermDepositInterestComputation string

// List of TermDepositInterestComputation
const (
	SIMPLE TermDepositInterestComputation = "Simple"
	COMPOUND TermDepositInterestComputation = "Compound"
)

// All allowed values of TermDepositInterestComputation enum
var AllowedTermDepositInterestComputationEnumValues = []TermDepositInterestComputation{
	"Simple",
	"Compound",
}

func (v *TermDepositInterestComputation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TermDepositInterestComputation(value)
	for _, existing := range AllowedTermDepositInterestComputationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TermDepositInterestComputation", value)
}

// NewTermDepositInterestComputationFromValue returns a pointer to a valid TermDepositInterestComputation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTermDepositInterestComputationFromValue(v string) (*TermDepositInterestComputation, error) {
	ev := TermDepositInterestComputation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TermDepositInterestComputation: valid values are %v", v, AllowedTermDepositInterestComputationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TermDepositInterestComputation) IsValid() bool {
	for _, existing := range AllowedTermDepositInterestComputationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TermDepositInterestComputation value
func (v TermDepositInterestComputation) Ptr() *TermDepositInterestComputation {
	return &v
}

type NullableTermDepositInterestComputation struct {
	value *TermDepositInterestComputation
	isSet bool
}

func (v NullableTermDepositInterestComputation) Get() *TermDepositInterestComputation {
	return v.value
}

func (v *NullableTermDepositInterestComputation) Set(val *TermDepositInterestComputation) {
	v.value = val
	v.isSet = true
}

func (v NullableTermDepositInterestComputation) IsSet() bool {
	return v.isSet
}

func (v *NullableTermDepositInterestComputation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTermDepositInterestComputation(val *TermDepositInterestComputation) *NullableTermDepositInterestComputation {
	return &NullableTermDepositInterestComputation{value: val, isSet: true}
}

func (v NullableTermDepositInterestComputation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTermDepositInterestComputation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


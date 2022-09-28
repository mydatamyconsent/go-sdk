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

// MutualFundHoldingMode the model 'MutualFundHoldingMode'
type MutualFundHoldingMode string

// List of MutualFundHoldingMode
const (
	DEMAT MutualFundHoldingMode = "Demat"
	PHYSICAL MutualFundHoldingMode = "Physical"
)

// All allowed values of MutualFundHoldingMode enum
var AllowedMutualFundHoldingModeEnumValues = []MutualFundHoldingMode{
	"Demat",
	"Physical",
}

func (v *MutualFundHoldingMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MutualFundHoldingMode(value)
	for _, existing := range AllowedMutualFundHoldingModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MutualFundHoldingMode", value)
}

// NewMutualFundHoldingModeFromValue returns a pointer to a valid MutualFundHoldingMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMutualFundHoldingModeFromValue(v string) (*MutualFundHoldingMode, error) {
	ev := MutualFundHoldingMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MutualFundHoldingMode: valid values are %v", v, AllowedMutualFundHoldingModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MutualFundHoldingMode) IsValid() bool {
	for _, existing := range AllowedMutualFundHoldingModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MutualFundHoldingMode value
func (v MutualFundHoldingMode) Ptr() *MutualFundHoldingMode {
	return &v
}

type NullableMutualFundHoldingMode struct {
	value *MutualFundHoldingMode
	isSet bool
}

func (v NullableMutualFundHoldingMode) Get() *MutualFundHoldingMode {
	return v.value
}

func (v *NullableMutualFundHoldingMode) Set(val *MutualFundHoldingMode) {
	v.value = val
	v.isSet = true
}

func (v NullableMutualFundHoldingMode) IsSet() bool {
	return v.isSet
}

func (v *NullableMutualFundHoldingMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMutualFundHoldingMode(val *MutualFundHoldingMode) *NullableMutualFundHoldingMode {
	return &NullableMutualFundHoldingMode{value: val, isSet: true}
}

func (v NullableMutualFundHoldingMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMutualFundHoldingMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


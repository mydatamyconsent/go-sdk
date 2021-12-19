/*
My Data My Consent - Developer API

Unleashing the power of data consent by establishing trust. The Platform Core Developer API defines a set of capabilities that can be used to request, issue, manage and update data, documents and credentials by organizations. The API can be used to request, manage and update Decentralised Identifiers, Financial Data, Health Data issue Documents, Credentials directly or using OpenID Connect flows, and verify Messages signed with DIDs and much more.

API version: v1
Contact: support@mydatamyconsent.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mydatamyconsent

import (
	"encoding/json"
	"fmt"
)

// IdentificationStrategy the model 'IdentificationStrategy'
type IdentificationStrategy string

// List of IdentificationStrategy
const (
	MATCH_AT_LEAST_ONE_IDENTIFIER IdentificationStrategy = "MatchAtLeastOneIdentifier"
	MATCH_ANY_TWO_IDENTIFIERS IdentificationStrategy = "MatchAnyTwoIdentifiers"
	MATCH_ALL_IDENTIFIERS IdentificationStrategy = "MatchAllIdentifiers"
)

var allowedIdentificationStrategyEnumValues = []IdentificationStrategy{
	"MatchAtLeastOneIdentifier",
	"MatchAnyTwoIdentifiers",
	"MatchAllIdentifiers",
}

func (v *IdentificationStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IdentificationStrategy(value)
	for _, existing := range allowedIdentificationStrategyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IdentificationStrategy", value)
}

// NewIdentificationStrategyFromValue returns a pointer to a valid IdentificationStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIdentificationStrategyFromValue(v string) (*IdentificationStrategy, error) {
	ev := IdentificationStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IdentificationStrategy: valid values are %v", v, allowedIdentificationStrategyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IdentificationStrategy) IsValid() bool {
	for _, existing := range allowedIdentificationStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IdentificationStrategy value
func (v IdentificationStrategy) Ptr() *IdentificationStrategy {
	return &v
}

type NullableIdentificationStrategy struct {
	value *IdentificationStrategy
	isSet bool
}

func (v NullableIdentificationStrategy) Get() *IdentificationStrategy {
	return v.value
}

func (v *NullableIdentificationStrategy) Set(val *IdentificationStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentificationStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentificationStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentificationStrategy(val *IdentificationStrategy) *NullableIdentificationStrategy {
	return &NullableIdentificationStrategy{value: val, isSet: true}
}

func (v NullableIdentificationStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentificationStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


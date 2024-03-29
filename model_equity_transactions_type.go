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

// EquityTransactionsType the model 'EquityTransactionsType'
type EquityTransactionsType string

// List of EquityTransactionsType
const (
	BUY EquityTransactionsType = "Buy"
	SELL EquityTransactionsType = "Sell"
	BONUS EquityTransactionsType = "Bonus"
	SPLIT EquityTransactionsType = "Split"
	DIVIDEND EquityTransactionsType = "Dividend"
	RIGHTS EquityTransactionsType = "Rights"
	OTHERS EquityTransactionsType = "Others"
)

// All allowed values of EquityTransactionsType enum
var AllowedEquityTransactionsTypeEnumValues = []EquityTransactionsType{
	"Buy",
	"Sell",
	"Bonus",
	"Split",
	"Dividend",
	"Rights",
	"Others",
}

func (v *EquityTransactionsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EquityTransactionsType(value)
	for _, existing := range AllowedEquityTransactionsTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EquityTransactionsType", value)
}

// NewEquityTransactionsTypeFromValue returns a pointer to a valid EquityTransactionsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEquityTransactionsTypeFromValue(v string) (*EquityTransactionsType, error) {
	ev := EquityTransactionsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EquityTransactionsType: valid values are %v", v, AllowedEquityTransactionsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EquityTransactionsType) IsValid() bool {
	for _, existing := range AllowedEquityTransactionsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EquityTransactionsType value
func (v EquityTransactionsType) Ptr() *EquityTransactionsType {
	return &v
}

type NullableEquityTransactionsType struct {
	value *EquityTransactionsType
	isSet bool
}

func (v NullableEquityTransactionsType) Get() *EquityTransactionsType {
	return v.value
}

func (v *NullableEquityTransactionsType) Set(val *EquityTransactionsType) {
	v.value = val
	v.isSet = true
}

func (v NullableEquityTransactionsType) IsSet() bool {
	return v.isSet
}

func (v *NullableEquityTransactionsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquityTransactionsType(val *EquityTransactionsType) *NullableEquityTransactionsType {
	return &NullableEquityTransactionsType{value: val, isSet: true}
}

func (v NullableEquityTransactionsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquityTransactionsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


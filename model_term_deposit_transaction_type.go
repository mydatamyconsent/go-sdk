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

// TermDepositTransactionType the model 'TermDepositTransactionType'
type TermDepositTransactionType string

// List of TermDepositTransactionType
const (
	OPENING TermDepositTransactionType = "Opening"
	INTEREST TermDepositTransactionType = "Interest"
	TDS TermDepositTransactionType = "Tds"
	INSTALLMENT TermDepositTransactionType = "Installment"
	CLOSING TermDepositTransactionType = "Closing"
	OTHERS TermDepositTransactionType = "Others"
)

// All allowed values of TermDepositTransactionType enum
var AllowedTermDepositTransactionTypeEnumValues = []TermDepositTransactionType{
	"Opening",
	"Interest",
	"Tds",
	"Installment",
	"Closing",
	"Others",
}

func (v *TermDepositTransactionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TermDepositTransactionType(value)
	for _, existing := range AllowedTermDepositTransactionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TermDepositTransactionType", value)
}

// NewTermDepositTransactionTypeFromValue returns a pointer to a valid TermDepositTransactionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTermDepositTransactionTypeFromValue(v string) (*TermDepositTransactionType, error) {
	ev := TermDepositTransactionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TermDepositTransactionType: valid values are %v", v, AllowedTermDepositTransactionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TermDepositTransactionType) IsValid() bool {
	for _, existing := range AllowedTermDepositTransactionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TermDepositTransactionType value
func (v TermDepositTransactionType) Ptr() *TermDepositTransactionType {
	return &v
}

type NullableTermDepositTransactionType struct {
	value *TermDepositTransactionType
	isSet bool
}

func (v NullableTermDepositTransactionType) Get() *TermDepositTransactionType {
	return v.value
}

func (v *NullableTermDepositTransactionType) Set(val *TermDepositTransactionType) {
	v.value = val
	v.isSet = true
}

func (v NullableTermDepositTransactionType) IsSet() bool {
	return v.isSet
}

func (v *NullableTermDepositTransactionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTermDepositTransactionType(val *TermDepositTransactionType) *NullableTermDepositTransactionType {
	return &NullableTermDepositTransactionType{value: val, isSet: true}
}

func (v NullableTermDepositTransactionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTermDepositTransactionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


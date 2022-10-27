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

// MutualFundTransactionType the model 'MutualFundTransactionType'
type MutualFundTransactionType string

// List of MutualFundTransactionType
const (
	BUY MutualFundTransactionType = "Buy"
	SELL MutualFundTransactionType = "Sell"
	OTHERS MutualFundTransactionType = "Others"
)

// All allowed values of MutualFundTransactionType enum
var AllowedMutualFundTransactionTypeEnumValues = []MutualFundTransactionType{
	"Buy",
	"Sell",
	"Others",
}

func (v *MutualFundTransactionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MutualFundTransactionType(value)
	for _, existing := range AllowedMutualFundTransactionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MutualFundTransactionType", value)
}

// NewMutualFundTransactionTypeFromValue returns a pointer to a valid MutualFundTransactionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMutualFundTransactionTypeFromValue(v string) (*MutualFundTransactionType, error) {
	ev := MutualFundTransactionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MutualFundTransactionType: valid values are %v", v, AllowedMutualFundTransactionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MutualFundTransactionType) IsValid() bool {
	for _, existing := range AllowedMutualFundTransactionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MutualFundTransactionType value
func (v MutualFundTransactionType) Ptr() *MutualFundTransactionType {
	return &v
}

type NullableMutualFundTransactionType struct {
	value *MutualFundTransactionType
	isSet bool
}

func (v NullableMutualFundTransactionType) Get() *MutualFundTransactionType {
	return v.value
}

func (v *NullableMutualFundTransactionType) Set(val *MutualFundTransactionType) {
	v.value = val
	v.isSet = true
}

func (v NullableMutualFundTransactionType) IsSet() bool {
	return v.isSet
}

func (v *NullableMutualFundTransactionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMutualFundTransactionType(val *MutualFundTransactionType) *NullableMutualFundTransactionType {
	return &NullableMutualFundTransactionType{value: val, isSet: true}
}

func (v NullableMutualFundTransactionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMutualFundTransactionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

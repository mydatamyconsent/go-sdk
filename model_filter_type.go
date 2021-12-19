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

// FilterType the model 'FilterType'
type FilterType string

// List of FilterType
const (
	TRANSACTION_TYPE FilterType = "TransactionType"
	TRANSACTION_AMOUNT FilterType = "TransactionAmount"
	TRANSACTION_DESCRIPTION FilterType = "TransactionDescription"
)

var allowedFilterTypeEnumValues = []FilterType{
	"TransactionType",
	"TransactionAmount",
	"TransactionDescription",
}

func (v *FilterType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FilterType(value)
	for _, existing := range allowedFilterTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FilterType", value)
}

// NewFilterTypeFromValue returns a pointer to a valid FilterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFilterTypeFromValue(v string) (*FilterType, error) {
	ev := FilterType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FilterType: valid values are %v", v, allowedFilterTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FilterType) IsValid() bool {
	for _, existing := range allowedFilterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FilterType value
func (v FilterType) Ptr() *FilterType {
	return &v
}

type NullableFilterType struct {
	value *FilterType
	isSet bool
}

func (v NullableFilterType) Get() *FilterType {
	return v.value
}

func (v *NullableFilterType) Set(val *FilterType) {
	v.value = val
	v.isSet = true
}

func (v NullableFilterType) IsSet() bool {
	return v.isSet
}

func (v *NullableFilterType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilterType(val *FilterType) *NullableFilterType {
	return &NullableFilterType{value: val, isSet: true}
}

func (v NullableFilterType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilterType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


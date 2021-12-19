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

// DataFetchType the model 'DataFetchType'
type DataFetchType string

// List of DataFetchType
const (
	ONETIME DataFetchType = "Onetime"
	PERIODIC DataFetchType = "Periodic"
)

var allowedDataFetchTypeEnumValues = []DataFetchType{
	"Onetime",
	"Periodic",
}

func (v *DataFetchType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataFetchType(value)
	for _, existing := range allowedDataFetchTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataFetchType", value)
}

// NewDataFetchTypeFromValue returns a pointer to a valid DataFetchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataFetchTypeFromValue(v string) (*DataFetchType, error) {
	ev := DataFetchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataFetchType: valid values are %v", v, allowedDataFetchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataFetchType) IsValid() bool {
	for _, existing := range allowedDataFetchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataFetchType value
func (v DataFetchType) Ptr() *DataFetchType {
	return &v
}

type NullableDataFetchType struct {
	value *DataFetchType
	isSet bool
}

func (v NullableDataFetchType) Get() *DataFetchType {
	return v.value
}

func (v *NullableDataFetchType) Set(val *DataFetchType) {
	v.value = val
	v.isSet = true
}

func (v NullableDataFetchType) IsSet() bool {
	return v.isSet
}

func (v *NullableDataFetchType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataFetchType(val *DataFetchType) *NullableDataFetchType {
	return &NullableDataFetchType{value: val, isSet: true}
}

func (v NullableDataFetchType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataFetchType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


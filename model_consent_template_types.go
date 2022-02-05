/*
My Data My Consent - Developer API

Unleashing the power of data consent by establishing trust. The Platform Core Developer API defines a set of capabilities that can be used to request, issue, manage and update data, documents and credentials by organizations. The API can be used to request, manage and update Decentralised Identifiers, Financial Data, Health Data issue Documents, Credentials directly or using OpenID Connect flows, and verify Messages signed with DIDs and much more.

API version: v1
Contact: support@mydatamyconsent.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github.com/mydatamyconsent/sdk

import (
	"encoding/json"
	"fmt"
)

// ConsentTemplateTypes the model 'ConsentTemplateTypes'
type ConsentTemplateTypes string

// List of ConsentTemplateTypes
const (
	INDIVIDUAL ConsentTemplateTypes = "Individual"
	ORGANIZATION ConsentTemplateTypes = "Organization"
)

// All allowed values of ConsentTemplateTypes enum
var AllowedConsentTemplateTypesEnumValues = []ConsentTemplateTypes{
	"Individual",
	"Organization",
}

func (v *ConsentTemplateTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConsentTemplateTypes(value)
	for _, existing := range AllowedConsentTemplateTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConsentTemplateTypes", value)
}

// NewConsentTemplateTypesFromValue returns a pointer to a valid ConsentTemplateTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConsentTemplateTypesFromValue(v string) (*ConsentTemplateTypes, error) {
	ev := ConsentTemplateTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConsentTemplateTypes: valid values are %v", v, AllowedConsentTemplateTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConsentTemplateTypes) IsValid() bool {
	for _, existing := range AllowedConsentTemplateTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConsentTemplateTypes value
func (v ConsentTemplateTypes) Ptr() *ConsentTemplateTypes {
	return &v
}

type NullableConsentTemplateTypes struct {
	value *ConsentTemplateTypes
	isSet bool
}

func (v NullableConsentTemplateTypes) Get() *ConsentTemplateTypes {
	return v.value
}

func (v *NullableConsentTemplateTypes) Set(val *ConsentTemplateTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableConsentTemplateTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableConsentTemplateTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsentTemplateTypes(val *ConsentTemplateTypes) *NullableConsentTemplateTypes {
	return &NullableConsentTemplateTypes{value: val, isSet: true}
}

func (v NullableConsentTemplateTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsentTemplateTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


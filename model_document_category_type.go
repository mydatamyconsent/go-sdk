/*
My Data My Consent - Developer API

Unleashing the power of data consent by establishing trust. The Platform Core Developer API defines a set of capabilities that can be used to request, issue, manage and update data, documents and credentials by organizations. The API can be used to request, manage and update Decentralised Identifiers, Financial Data, Health Data issue Documents, Credentials directly or using OpenID Connect flows, and verify Messages signed with DIDs and much more.

API version: v1
Contact: support@mydatamyconsent.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github.com/mydatamyconsent/go-sdk

import (
	"encoding/json"
	"fmt"
)

// DocumentCategoryType the model 'DocumentCategoryType'
type DocumentCategoryType string

// List of DocumentCategoryType
const (
	HEALTH DocumentCategoryType = "Health"
	FINANCE DocumentCategoryType = "Finance"
	EDUCATION DocumentCategoryType = "Education"
	BILLS DocumentCategoryType = "Bills"
	TAX DocumentCategoryType = "Tax"
	CERTIFICATES DocumentCategoryType = "Certificates"
	INVOICES DocumentCategoryType = "Invoices"
)

// All allowed values of DocumentCategoryType enum
var AllowedDocumentCategoryTypeEnumValues = []DocumentCategoryType{
	"Health",
	"Finance",
	"Education",
	"Bills",
	"Tax",
	"Certificates",
	"Invoices",
}

func (v *DocumentCategoryType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DocumentCategoryType(value)
	for _, existing := range AllowedDocumentCategoryTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DocumentCategoryType", value)
}

// NewDocumentCategoryTypeFromValue returns a pointer to a valid DocumentCategoryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDocumentCategoryTypeFromValue(v string) (*DocumentCategoryType, error) {
	ev := DocumentCategoryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DocumentCategoryType: valid values are %v", v, AllowedDocumentCategoryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DocumentCategoryType) IsValid() bool {
	for _, existing := range AllowedDocumentCategoryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DocumentCategoryType value
func (v DocumentCategoryType) Ptr() *DocumentCategoryType {
	return &v
}

type NullableDocumentCategoryType struct {
	value *DocumentCategoryType
	isSet bool
}

func (v NullableDocumentCategoryType) Get() *DocumentCategoryType {
	return v.value
}

func (v *NullableDocumentCategoryType) Set(val *DocumentCategoryType) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentCategoryType) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentCategoryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentCategoryType(val *DocumentCategoryType) *NullableDocumentCategoryType {
	return &NullableDocumentCategoryType{value: val, isSet: true}
}

func (v NullableDocumentCategoryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentCategoryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


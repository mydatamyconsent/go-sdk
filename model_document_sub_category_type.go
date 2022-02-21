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

// DocumentSubCategoryType the model 'DocumentSubCategoryType'
type DocumentSubCategoryType string

// List of DocumentSubCategoryType
const (
	IDENTITY DocumentSubCategoryType = "Identity"
	HOUSE DocumentSubCategoryType = "House"
	VEHICLE DocumentSubCategoryType = "Vehicle"
	MARKS_MEMOS DocumentSubCategoryType = "MarksMemos"
	TRANSCRIPTS DocumentSubCategoryType = "Transcripts"
	CERTIFICATES DocumentSubCategoryType = "Certificates"
	PRESCRIPTIONS DocumentSubCategoryType = "Prescriptions"
	MEDICAL_REPORTS DocumentSubCategoryType = "MedicalReports"
	DISCHARGE_SUMMARY DocumentSubCategoryType = "DischargeSummary"
	BILLS DocumentSubCategoryType = "Bills"
	INVOICES DocumentSubCategoryType = "Invoices"
	TAXES DocumentSubCategoryType = "Taxes"
	CORPORATES DocumentSubCategoryType = "Corporates"
	COMPETITIONS DocumentSubCategoryType = "Competitions"
	INTELLECTUAL_PROPERTIES DocumentSubCategoryType = "IntellectualProperties"
)

// All allowed values of DocumentSubCategoryType enum
var AllowedDocumentSubCategoryTypeEnumValues = []DocumentSubCategoryType{
	"Identity",
	"House",
	"Vehicle",
	"MarksMemos",
	"Transcripts",
	"Certificates",
	"Prescriptions",
	"MedicalReports",
	"DischargeSummary",
	"Bills",
	"Invoices",
	"Taxes",
	"Corporates",
	"Competitions",
	"IntellectualProperties",
}

func (v *DocumentSubCategoryType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DocumentSubCategoryType(value)
	for _, existing := range AllowedDocumentSubCategoryTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DocumentSubCategoryType", value)
}

// NewDocumentSubCategoryTypeFromValue returns a pointer to a valid DocumentSubCategoryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDocumentSubCategoryTypeFromValue(v string) (*DocumentSubCategoryType, error) {
	ev := DocumentSubCategoryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DocumentSubCategoryType: valid values are %v", v, AllowedDocumentSubCategoryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DocumentSubCategoryType) IsValid() bool {
	for _, existing := range AllowedDocumentSubCategoryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DocumentSubCategoryType value
func (v DocumentSubCategoryType) Ptr() *DocumentSubCategoryType {
	return &v
}

type NullableDocumentSubCategoryType struct {
	value *DocumentSubCategoryType
	isSet bool
}

func (v NullableDocumentSubCategoryType) Get() *DocumentSubCategoryType {
	return v.value
}

func (v *NullableDocumentSubCategoryType) Set(val *DocumentSubCategoryType) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentSubCategoryType) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentSubCategoryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentSubCategoryType(val *DocumentSubCategoryType) *NullableDocumentSubCategoryType {
	return &NullableDocumentSubCategoryType{value: val, isSet: true}
}

func (v NullableDocumentSubCategoryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentSubCategoryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


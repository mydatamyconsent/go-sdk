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

// DocumentIssueRequestStatus Document issue request status.
type DocumentIssueRequestStatus string

// List of DocumentIssueRequestStatus
const (
	CREATED DocumentIssueRequestStatus = "Created"
	ISSUED DocumentIssueRequestStatus = "Issued"
	ACCEPTED DocumentIssueRequestStatus = "Accepted"
	REJECTED DocumentIssueRequestStatus = "Rejected"
)

// All allowed values of DocumentIssueRequestStatus enum
var AllowedDocumentIssueRequestStatusEnumValues = []DocumentIssueRequestStatus{
	"Created",
	"Issued",
	"Accepted",
	"Rejected",
}

func (v *DocumentIssueRequestStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DocumentIssueRequestStatus(value)
	for _, existing := range AllowedDocumentIssueRequestStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DocumentIssueRequestStatus", value)
}

// NewDocumentIssueRequestStatusFromValue returns a pointer to a valid DocumentIssueRequestStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDocumentIssueRequestStatusFromValue(v string) (*DocumentIssueRequestStatus, error) {
	ev := DocumentIssueRequestStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DocumentIssueRequestStatus: valid values are %v", v, AllowedDocumentIssueRequestStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DocumentIssueRequestStatus) IsValid() bool {
	for _, existing := range AllowedDocumentIssueRequestStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DocumentIssueRequestStatus value
func (v DocumentIssueRequestStatus) Ptr() *DocumentIssueRequestStatus {
	return &v
}

type NullableDocumentIssueRequestStatus struct {
	value *DocumentIssueRequestStatus
	isSet bool
}

func (v NullableDocumentIssueRequestStatus) Get() *DocumentIssueRequestStatus {
	return v.value
}

func (v *NullableDocumentIssueRequestStatus) Set(val *DocumentIssueRequestStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentIssueRequestStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentIssueRequestStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentIssueRequestStatus(val *DocumentIssueRequestStatus) *NullableDocumentIssueRequestStatus {
	return &NullableDocumentIssueRequestStatus{value: val, isSet: true}
}

func (v NullableDocumentIssueRequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentIssueRequestStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


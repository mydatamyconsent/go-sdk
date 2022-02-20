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

// DataConsentStatus the model 'DataConsentStatus'
type DataConsentStatus string

// List of DataConsentStatus
const (
	PENDING DataConsentStatus = "Pending"
	APPROVED DataConsentStatus = "Approved"
	REJECTED DataConsentStatus = "Rejected"
	REVOKED DataConsentStatus = "Revoked"
	EXPIRED DataConsentStatus = "Expired"
	CANCELED DataConsentStatus = "Canceled"
)

// All allowed values of DataConsentStatus enum
var AllowedDataConsentStatusEnumValues = []DataConsentStatus{
	"Pending",
	"Approved",
	"Rejected",
	"Revoked",
	"Expired",
	"Canceled",
}

func (v *DataConsentStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataConsentStatus(value)
	for _, existing := range AllowedDataConsentStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataConsentStatus", value)
}

// NewDataConsentStatusFromValue returns a pointer to a valid DataConsentStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataConsentStatusFromValue(v string) (*DataConsentStatus, error) {
	ev := DataConsentStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataConsentStatus: valid values are %v", v, AllowedDataConsentStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataConsentStatus) IsValid() bool {
	for _, existing := range AllowedDataConsentStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataConsentStatus value
func (v DataConsentStatus) Ptr() *DataConsentStatus {
	return &v
}

type NullableDataConsentStatus struct {
	value *DataConsentStatus
	isSet bool
}

func (v NullableDataConsentStatus) Get() *DataConsentStatus {
	return v.value
}

func (v *NullableDataConsentStatus) Set(val *DataConsentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDataConsentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDataConsentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataConsentStatus(val *DataConsentStatus) *NullableDataConsentStatus {
	return &NullableDataConsentStatus{value: val, isSet: true}
}

func (v NullableDataConsentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataConsentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


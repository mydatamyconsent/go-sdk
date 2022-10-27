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
)

// FinancialAccountGovtSecurityTransaction struct for FinancialAccountGovtSecurityTransaction
type FinancialAccountGovtSecurityTransaction struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewFinancialAccountGovtSecurityTransaction instantiates a new FinancialAccountGovtSecurityTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialAccountGovtSecurityTransaction(type_ string, id string) *FinancialAccountGovtSecurityTransaction {
	this := FinancialAccountGovtSecurityTransaction{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewFinancialAccountGovtSecurityTransactionWithDefaults instantiates a new FinancialAccountGovtSecurityTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialAccountGovtSecurityTransactionWithDefaults() *FinancialAccountGovtSecurityTransaction {
	this := FinancialAccountGovtSecurityTransaction{}
	return &this
}

// GetType returns the Type field value
func (o *FinancialAccountGovtSecurityTransaction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountGovtSecurityTransaction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FinancialAccountGovtSecurityTransaction) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *FinancialAccountGovtSecurityTransaction) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountGovtSecurityTransaction) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FinancialAccountGovtSecurityTransaction) SetId(v string) {
	o.Id = v
}

func (o FinancialAccountGovtSecurityTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableFinancialAccountGovtSecurityTransaction struct {
	value *FinancialAccountGovtSecurityTransaction
	isSet bool
}

func (v NullableFinancialAccountGovtSecurityTransaction) Get() *FinancialAccountGovtSecurityTransaction {
	return v.value
}

func (v *NullableFinancialAccountGovtSecurityTransaction) Set(val *FinancialAccountGovtSecurityTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialAccountGovtSecurityTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialAccountGovtSecurityTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialAccountGovtSecurityTransaction(val *FinancialAccountGovtSecurityTransaction) *NullableFinancialAccountGovtSecurityTransaction {
	return &NullableFinancialAccountGovtSecurityTransaction{value: val, isSet: true}
}

func (v NullableFinancialAccountGovtSecurityTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialAccountGovtSecurityTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


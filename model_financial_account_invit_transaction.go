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

// FinancialAccountInvitTransaction struct for FinancialAccountInvitTransaction
type FinancialAccountInvitTransaction struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewFinancialAccountInvitTransaction instantiates a new FinancialAccountInvitTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialAccountInvitTransaction(type_ string, id string) *FinancialAccountInvitTransaction {
	this := FinancialAccountInvitTransaction{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewFinancialAccountInvitTransactionWithDefaults instantiates a new FinancialAccountInvitTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialAccountInvitTransactionWithDefaults() *FinancialAccountInvitTransaction {
	this := FinancialAccountInvitTransaction{}
	return &this
}

// GetType returns the Type field value
func (o *FinancialAccountInvitTransaction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountInvitTransaction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FinancialAccountInvitTransaction) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *FinancialAccountInvitTransaction) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountInvitTransaction) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FinancialAccountInvitTransaction) SetId(v string) {
	o.Id = v
}

func (o FinancialAccountInvitTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableFinancialAccountInvitTransaction struct {
	value *FinancialAccountInvitTransaction
	isSet bool
}

func (v NullableFinancialAccountInvitTransaction) Get() *FinancialAccountInvitTransaction {
	return v.value
}

func (v *NullableFinancialAccountInvitTransaction) Set(val *FinancialAccountInvitTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialAccountInvitTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialAccountInvitTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialAccountInvitTransaction(val *FinancialAccountInvitTransaction) *NullableFinancialAccountInvitTransaction {
	return &NullableFinancialAccountInvitTransaction{value: val, isSet: true}
}

func (v NullableFinancialAccountInvitTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialAccountInvitTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



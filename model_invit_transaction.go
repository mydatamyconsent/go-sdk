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

// InvitTransaction struct for InvitTransaction
type InvitTransaction struct {
	Id string `json:"id"`
}

// NewInvitTransaction instantiates a new InvitTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitTransaction(id string) *InvitTransaction {
	this := InvitTransaction{}
	this.Id = id
	return &this
}

// NewInvitTransactionWithDefaults instantiates a new InvitTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitTransactionWithDefaults() *InvitTransaction {
	this := InvitTransaction{}
	return &this
}

// GetId returns the Id field value
func (o *InvitTransaction) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InvitTransaction) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InvitTransaction) SetId(v string) {
	o.Id = v
}

func (o InvitTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableInvitTransaction struct {
	value *InvitTransaction
	isSet bool
}

func (v NullableInvitTransaction) Get() *InvitTransaction {
	return v.value
}

func (v *NullableInvitTransaction) Set(val *InvitTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableInvitTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableInvitTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvitTransaction(val *InvitTransaction) *NullableInvitTransaction {
	return &NullableInvitTransaction{value: val, isSet: true}
}

func (v NullableInvitTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvitTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



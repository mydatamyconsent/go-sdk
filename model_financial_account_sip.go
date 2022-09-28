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

// FinancialAccountSip struct for FinancialAccountSip
type FinancialAccountSip struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Name string `json:"name"`
	Identifier string `json:"identifier"`
	Amount float64 `json:"amount"`
}

// NewFinancialAccountSip instantiates a new FinancialAccountSip object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialAccountSip(type_ string, id string, name string, identifier string, amount float64) *FinancialAccountSip {
	this := FinancialAccountSip{}
	this.Type = type_
	this.Id = id
	this.Name = name
	this.Identifier = identifier
	this.Amount = amount
	return &this
}

// NewFinancialAccountSipWithDefaults instantiates a new FinancialAccountSip object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialAccountSipWithDefaults() *FinancialAccountSip {
	this := FinancialAccountSip{}
	return &this
}

// GetType returns the Type field value
func (o *FinancialAccountSip) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountSip) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FinancialAccountSip) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *FinancialAccountSip) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountSip) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FinancialAccountSip) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *FinancialAccountSip) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountSip) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FinancialAccountSip) SetName(v string) {
	o.Name = v
}

// GetIdentifier returns the Identifier field value
func (o *FinancialAccountSip) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountSip) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *FinancialAccountSip) SetIdentifier(v string) {
	o.Identifier = v
}

// GetAmount returns the Amount field value
func (o *FinancialAccountSip) GetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountSip) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *FinancialAccountSip) SetAmount(v float64) {
	o.Amount = v
}

func (o FinancialAccountSip) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["identifier"] = o.Identifier
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullableFinancialAccountSip struct {
	value *FinancialAccountSip
	isSet bool
}

func (v NullableFinancialAccountSip) Get() *FinancialAccountSip {
	return v.value
}

func (v *NullableFinancialAccountSip) Set(val *FinancialAccountSip) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialAccountSip) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialAccountSip) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialAccountSip(val *FinancialAccountSip) *NullableFinancialAccountSip {
	return &NullableFinancialAccountSip{value: val, isSet: true}
}

func (v NullableFinancialAccountSip) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialAccountSip) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// TermDeposit struct for TermDeposit
type TermDeposit struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Identifier string `json:"identifier"`
	Amount float64 `json:"amount"`
}

// NewTermDeposit instantiates a new TermDeposit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTermDeposit(id string, name string, identifier string, amount float64) *TermDeposit {
	this := TermDeposit{}
	this.Id = id
	this.Name = name
	this.Identifier = identifier
	this.Amount = amount
	return &this
}

// NewTermDepositWithDefaults instantiates a new TermDeposit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTermDepositWithDefaults() *TermDeposit {
	this := TermDeposit{}
	return &this
}

// GetId returns the Id field value
func (o *TermDeposit) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TermDeposit) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TermDeposit) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TermDeposit) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TermDeposit) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TermDeposit) SetName(v string) {
	o.Name = v
}

// GetIdentifier returns the Identifier field value
func (o *TermDeposit) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *TermDeposit) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *TermDeposit) SetIdentifier(v string) {
	o.Identifier = v
}

// GetAmount returns the Amount field value
func (o *TermDeposit) GetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TermDeposit) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TermDeposit) SetAmount(v float64) {
	o.Amount = v
}

func (o TermDeposit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableTermDeposit struct {
	value *TermDeposit
	isSet bool
}

func (v NullableTermDeposit) Get() *TermDeposit {
	return v.value
}

func (v *NullableTermDeposit) Set(val *TermDeposit) {
	v.value = val
	v.isSet = true
}

func (v NullableTermDeposit) IsSet() bool {
	return v.isSet
}

func (v *NullableTermDeposit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTermDeposit(val *TermDeposit) *NullableTermDeposit {
	return &NullableTermDeposit{value: val, isSet: true}
}

func (v NullableTermDeposit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTermDeposit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



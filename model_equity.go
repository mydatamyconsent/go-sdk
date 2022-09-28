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

// Equity struct for Equity
type Equity struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Identifier string `json:"identifier"`
	Balance float64 `json:"balance"`
	Profile Profile `json:"profile"`
	Summary EquitySummary `json:"summary"`
	MaskedAccountNumber string `json:"masked_account_number"`
	LinkedAccountRef string `json:"linked_account_ref"`
	Version float32 `json:"version"`
}

// NewEquity instantiates a new Equity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquity(id string, name string, identifier string, balance float64, profile Profile, summary EquitySummary, maskedAccountNumber string, linkedAccountRef string, version float32) *Equity {
	this := Equity{}
	this.Id = id
	this.Name = name
	this.Identifier = identifier
	this.Balance = balance
	this.Profile = profile
	this.Summary = summary
	this.MaskedAccountNumber = maskedAccountNumber
	this.LinkedAccountRef = linkedAccountRef
	this.Version = version
	return &this
}

// NewEquityWithDefaults instantiates a new Equity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquityWithDefaults() *Equity {
	this := Equity{}
	return &this
}

// GetId returns the Id field value
func (o *Equity) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Equity) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Equity) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Equity) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Equity) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Equity) SetName(v string) {
	o.Name = v
}

// GetIdentifier returns the Identifier field value
func (o *Equity) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *Equity) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *Equity) SetIdentifier(v string) {
	o.Identifier = v
}

// GetBalance returns the Balance field value
func (o *Equity) GetBalance() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *Equity) GetBalanceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *Equity) SetBalance(v float64) {
	o.Balance = v
}

// GetProfile returns the Profile field value
func (o *Equity) GetProfile() Profile {
	if o == nil {
		var ret Profile
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *Equity) GetProfileOk() (*Profile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *Equity) SetProfile(v Profile) {
	o.Profile = v
}

// GetSummary returns the Summary field value
func (o *Equity) GetSummary() EquitySummary {
	if o == nil {
		var ret EquitySummary
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *Equity) GetSummaryOk() (*EquitySummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *Equity) SetSummary(v EquitySummary) {
	o.Summary = v
}

// GetMaskedAccountNumber returns the MaskedAccountNumber field value
func (o *Equity) GetMaskedAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaskedAccountNumber
}

// GetMaskedAccountNumberOk returns a tuple with the MaskedAccountNumber field value
// and a boolean to check if the value has been set.
func (o *Equity) GetMaskedAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaskedAccountNumber, true
}

// SetMaskedAccountNumber sets field value
func (o *Equity) SetMaskedAccountNumber(v string) {
	o.MaskedAccountNumber = v
}

// GetLinkedAccountRef returns the LinkedAccountRef field value
func (o *Equity) GetLinkedAccountRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkedAccountRef
}

// GetLinkedAccountRefOk returns a tuple with the LinkedAccountRef field value
// and a boolean to check if the value has been set.
func (o *Equity) GetLinkedAccountRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkedAccountRef, true
}

// SetLinkedAccountRef sets field value
func (o *Equity) SetLinkedAccountRef(v string) {
	o.LinkedAccountRef = v
}

// GetVersion returns the Version field value
func (o *Equity) GetVersion() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Equity) GetVersionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *Equity) SetVersion(v float32) {
	o.Version = v
}

func (o Equity) MarshalJSON() ([]byte, error) {
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
		toSerialize["balance"] = o.Balance
	}
	if true {
		toSerialize["profile"] = o.Profile
	}
	if true {
		toSerialize["summary"] = o.Summary
	}
	if true {
		toSerialize["masked_account_number"] = o.MaskedAccountNumber
	}
	if true {
		toSerialize["linked_account_ref"] = o.LinkedAccountRef
	}
	if true {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableEquity struct {
	value *Equity
	isSet bool
}

func (v NullableEquity) Get() *Equity {
	return v.value
}

func (v *NullableEquity) Set(val *Equity) {
	v.value = val
	v.isSet = true
}

func (v NullableEquity) IsSet() bool {
	return v.isSet
}

func (v *NullableEquity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquity(val *Equity) *NullableEquity {
	return &NullableEquity{value: val, isSet: true}
}

func (v NullableEquity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



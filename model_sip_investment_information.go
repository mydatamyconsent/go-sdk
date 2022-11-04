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
	"time"
)

// SipInvestmentInformation struct for SipInvestmentInformation
type SipInvestmentInformation struct {
	InstalmentAmount float64 `json:"instalment_amount"`
	Frequency string `json:"frequency"`
	CompletedInstalments float64 `json:"completed_instalments"`
	InvestmentValue float64 `json:"investment_value"`
	LastInstalmentDate *time.Time `json:"last_instalment_date,omitempty"`
	NextInstalmentDate *time.Time `json:"next_instalment_date,omitempty"`
}

// NewSipInvestmentInformation instantiates a new SipInvestmentInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSipInvestmentInformation(instalmentAmount float64, frequency string, completedInstalments float64, investmentValue float64) *SipInvestmentInformation {
	this := SipInvestmentInformation{}
	this.InstalmentAmount = instalmentAmount
	this.Frequency = frequency
	this.CompletedInstalments = completedInstalments
	this.InvestmentValue = investmentValue
	return &this
}

// NewSipInvestmentInformationWithDefaults instantiates a new SipInvestmentInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSipInvestmentInformationWithDefaults() *SipInvestmentInformation {
	this := SipInvestmentInformation{}
	return &this
}

// GetInstalmentAmount returns the InstalmentAmount field value
func (o *SipInvestmentInformation) GetInstalmentAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.InstalmentAmount
}

// GetInstalmentAmountOk returns a tuple with the InstalmentAmount field value
// and a boolean to check if the value has been set.
func (o *SipInvestmentInformation) GetInstalmentAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstalmentAmount, true
}

// SetInstalmentAmount sets field value
func (o *SipInvestmentInformation) SetInstalmentAmount(v float64) {
	o.InstalmentAmount = v
}

// GetFrequency returns the Frequency field value
func (o *SipInvestmentInformation) GetFrequency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
func (o *SipInvestmentInformation) GetFrequencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frequency, true
}

// SetFrequency sets field value
func (o *SipInvestmentInformation) SetFrequency(v string) {
	o.Frequency = v
}

// GetCompletedInstalments returns the CompletedInstalments field value
func (o *SipInvestmentInformation) GetCompletedInstalments() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.CompletedInstalments
}

// GetCompletedInstalmentsOk returns a tuple with the CompletedInstalments field value
// and a boolean to check if the value has been set.
func (o *SipInvestmentInformation) GetCompletedInstalmentsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletedInstalments, true
}

// SetCompletedInstalments sets field value
func (o *SipInvestmentInformation) SetCompletedInstalments(v float64) {
	o.CompletedInstalments = v
}

// GetInvestmentValue returns the InvestmentValue field value
func (o *SipInvestmentInformation) GetInvestmentValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.InvestmentValue
}

// GetInvestmentValueOk returns a tuple with the InvestmentValue field value
// and a boolean to check if the value has been set.
func (o *SipInvestmentInformation) GetInvestmentValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvestmentValue, true
}

// SetInvestmentValue sets field value
func (o *SipInvestmentInformation) SetInvestmentValue(v float64) {
	o.InvestmentValue = v
}

// GetLastInstalmentDate returns the LastInstalmentDate field value if set, zero value otherwise.
func (o *SipInvestmentInformation) GetLastInstalmentDate() time.Time {
	if o == nil || o.LastInstalmentDate == nil {
		var ret time.Time
		return ret
	}
	return *o.LastInstalmentDate
}

// GetLastInstalmentDateOk returns a tuple with the LastInstalmentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SipInvestmentInformation) GetLastInstalmentDateOk() (*time.Time, bool) {
	if o == nil || o.LastInstalmentDate == nil {
		return nil, false
	}
	return o.LastInstalmentDate, true
}

// HasLastInstalmentDate returns a boolean if a field has been set.
func (o *SipInvestmentInformation) HasLastInstalmentDate() bool {
	if o != nil && o.LastInstalmentDate != nil {
		return true
	}

	return false
}

// SetLastInstalmentDate gets a reference to the given time.Time and assigns it to the LastInstalmentDate field.
func (o *SipInvestmentInformation) SetLastInstalmentDate(v time.Time) {
	o.LastInstalmentDate = &v
}

// GetNextInstalmentDate returns the NextInstalmentDate field value if set, zero value otherwise.
func (o *SipInvestmentInformation) GetNextInstalmentDate() time.Time {
	if o == nil || o.NextInstalmentDate == nil {
		var ret time.Time
		return ret
	}
	return *o.NextInstalmentDate
}

// GetNextInstalmentDateOk returns a tuple with the NextInstalmentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SipInvestmentInformation) GetNextInstalmentDateOk() (*time.Time, bool) {
	if o == nil || o.NextInstalmentDate == nil {
		return nil, false
	}
	return o.NextInstalmentDate, true
}

// HasNextInstalmentDate returns a boolean if a field has been set.
func (o *SipInvestmentInformation) HasNextInstalmentDate() bool {
	if o != nil && o.NextInstalmentDate != nil {
		return true
	}

	return false
}

// SetNextInstalmentDate gets a reference to the given time.Time and assigns it to the NextInstalmentDate field.
func (o *SipInvestmentInformation) SetNextInstalmentDate(v time.Time) {
	o.NextInstalmentDate = &v
}

func (o SipInvestmentInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["instalment_amount"] = o.InstalmentAmount
	}
	if true {
		toSerialize["frequency"] = o.Frequency
	}
	if true {
		toSerialize["completed_instalments"] = o.CompletedInstalments
	}
	if true {
		toSerialize["investment_value"] = o.InvestmentValue
	}
	if o.LastInstalmentDate != nil {
		toSerialize["last_instalment_date"] = o.LastInstalmentDate
	}
	if o.NextInstalmentDate != nil {
		toSerialize["next_instalment_date"] = o.NextInstalmentDate
	}
	return json.Marshal(toSerialize)
}

type NullableSipInvestmentInformation struct {
	value *SipInvestmentInformation
	isSet bool
}

func (v NullableSipInvestmentInformation) Get() *SipInvestmentInformation {
	return v.value
}

func (v *NullableSipInvestmentInformation) Set(val *SipInvestmentInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableSipInvestmentInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableSipInvestmentInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSipInvestmentInformation(val *SipInvestmentInformation) *NullableSipInvestmentInformation {
	return &NullableSipInvestmentInformation{value: val, isSet: true}
}

func (v NullableSipInvestmentInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSipInvestmentInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


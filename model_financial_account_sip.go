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
	InvestmentValue float64 `json:"investment_value"`
	CurrentValue float64 `json:"current_value"`
	CurrencyCode string `json:"currency_code"`
	PlanInfo SipPlanInformation `json:"plan_info"`
	InvestmentInfo SipInvestmentInformation `json:"investment_info"`
	Holder Holder `json:"holder"`
	Transactions bool `json:"transactions"`
}

// NewFinancialAccountSip instantiates a new FinancialAccountSip object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialAccountSip(type_ string, id string, name string, investmentValue float64, currentValue float64, currencyCode string, planInfo SipPlanInformation, investmentInfo SipInvestmentInformation, holder Holder, transactions bool) *FinancialAccountSip {
	this := FinancialAccountSip{}
	this.Type = type_
	this.Id = id
	this.Name = name
	this.InvestmentValue = investmentValue
	this.CurrentValue = currentValue
	this.CurrencyCode = currencyCode
	this.PlanInfo = planInfo
	this.InvestmentInfo = investmentInfo
	this.Holder = holder
	this.Transactions = transactions
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

// GetInvestmentValue returns the InvestmentValue field value
func (o *FinancialAccountSip) GetInvestmentValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.InvestmentValue
}

// GetInvestmentValueOk returns a tuple with the InvestmentValue field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountSip) GetInvestmentValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvestmentValue, true
}

// SetInvestmentValue sets field value
func (o *FinancialAccountSip) SetInvestmentValue(v float64) {
	o.InvestmentValue = v
}

// GetCurrentValue returns the CurrentValue field value
func (o *FinancialAccountSip) GetCurrentValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.CurrentValue
}

// GetCurrentValueOk returns a tuple with the CurrentValue field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountSip) GetCurrentValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentValue, true
}

// SetCurrentValue sets field value
func (o *FinancialAccountSip) SetCurrentValue(v float64) {
	o.CurrentValue = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *FinancialAccountSip) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountSip) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *FinancialAccountSip) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

// GetPlanInfo returns the PlanInfo field value
func (o *FinancialAccountSip) GetPlanInfo() SipPlanInformation {
	if o == nil {
		var ret SipPlanInformation
		return ret
	}

	return o.PlanInfo
}

// GetPlanInfoOk returns a tuple with the PlanInfo field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountSip) GetPlanInfoOk() (*SipPlanInformation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanInfo, true
}

// SetPlanInfo sets field value
func (o *FinancialAccountSip) SetPlanInfo(v SipPlanInformation) {
	o.PlanInfo = v
}

// GetInvestmentInfo returns the InvestmentInfo field value
func (o *FinancialAccountSip) GetInvestmentInfo() SipInvestmentInformation {
	if o == nil {
		var ret SipInvestmentInformation
		return ret
	}

	return o.InvestmentInfo
}

// GetInvestmentInfoOk returns a tuple with the InvestmentInfo field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountSip) GetInvestmentInfoOk() (*SipInvestmentInformation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvestmentInfo, true
}

// SetInvestmentInfo sets field value
func (o *FinancialAccountSip) SetInvestmentInfo(v SipInvestmentInformation) {
	o.InvestmentInfo = v
}

// GetHolder returns the Holder field value
func (o *FinancialAccountSip) GetHolder() Holder {
	if o == nil {
		var ret Holder
		return ret
	}

	return o.Holder
}

// GetHolderOk returns a tuple with the Holder field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountSip) GetHolderOk() (*Holder, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Holder, true
}

// SetHolder sets field value
func (o *FinancialAccountSip) SetHolder(v Holder) {
	o.Holder = v
}

// GetTransactions returns the Transactions field value
func (o *FinancialAccountSip) GetTransactions() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountSip) GetTransactionsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transactions, true
}

// SetTransactions sets field value
func (o *FinancialAccountSip) SetTransactions(v bool) {
	o.Transactions = v
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
		toSerialize["investment_value"] = o.InvestmentValue
	}
	if true {
		toSerialize["current_value"] = o.CurrentValue
	}
	if true {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if true {
		toSerialize["plan_info"] = o.PlanInfo
	}
	if true {
		toSerialize["investment_info"] = o.InvestmentInfo
	}
	if true {
		toSerialize["holder"] = o.Holder
	}
	if true {
		toSerialize["transactions"] = o.Transactions
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



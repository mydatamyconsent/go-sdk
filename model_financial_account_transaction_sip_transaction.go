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

// FinancialAccountTransactionSipTransaction struct for FinancialAccountTransactionSipTransaction
type FinancialAccountTransactionSipTransaction struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Amount float64 `json:"amount"`
	CurrencyCode string `json:"currency_code"`
	TxnType SipTransactionType `json:"txn_type"`
	TransactedAtUtc time.Time `json:"transacted_at_utc"`
}

// NewFinancialAccountTransactionSipTransaction instantiates a new FinancialAccountTransactionSipTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialAccountTransactionSipTransaction(type_ string, id string, amount float64, currencyCode string, txnType SipTransactionType, transactedAtUtc time.Time) *FinancialAccountTransactionSipTransaction {
	this := FinancialAccountTransactionSipTransaction{}
	this.Type = type_
	this.Id = id
	this.Amount = amount
	this.CurrencyCode = currencyCode
	this.TxnType = txnType
	this.TransactedAtUtc = transactedAtUtc
	return &this
}

// NewFinancialAccountTransactionSipTransactionWithDefaults instantiates a new FinancialAccountTransactionSipTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialAccountTransactionSipTransactionWithDefaults() *FinancialAccountTransactionSipTransaction {
	this := FinancialAccountTransactionSipTransaction{}
	return &this
}

// GetType returns the Type field value
func (o *FinancialAccountTransactionSipTransaction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountTransactionSipTransaction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FinancialAccountTransactionSipTransaction) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *FinancialAccountTransactionSipTransaction) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountTransactionSipTransaction) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FinancialAccountTransactionSipTransaction) SetId(v string) {
	o.Id = v
}

// GetAmount returns the Amount field value
func (o *FinancialAccountTransactionSipTransaction) GetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountTransactionSipTransaction) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *FinancialAccountTransactionSipTransaction) SetAmount(v float64) {
	o.Amount = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *FinancialAccountTransactionSipTransaction) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountTransactionSipTransaction) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *FinancialAccountTransactionSipTransaction) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

// GetTxnType returns the TxnType field value
func (o *FinancialAccountTransactionSipTransaction) GetTxnType() SipTransactionType {
	if o == nil {
		var ret SipTransactionType
		return ret
	}

	return o.TxnType
}

// GetTxnTypeOk returns a tuple with the TxnType field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountTransactionSipTransaction) GetTxnTypeOk() (*SipTransactionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxnType, true
}

// SetTxnType sets field value
func (o *FinancialAccountTransactionSipTransaction) SetTxnType(v SipTransactionType) {
	o.TxnType = v
}

// GetTransactedAtUtc returns the TransactedAtUtc field value
func (o *FinancialAccountTransactionSipTransaction) GetTransactedAtUtc() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TransactedAtUtc
}

// GetTransactedAtUtcOk returns a tuple with the TransactedAtUtc field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountTransactionSipTransaction) GetTransactedAtUtcOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactedAtUtc, true
}

// SetTransactedAtUtc sets field value
func (o *FinancialAccountTransactionSipTransaction) SetTransactedAtUtc(v time.Time) {
	o.TransactedAtUtc = v
}

func (o FinancialAccountTransactionSipTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if true {
		toSerialize["txn_type"] = o.TxnType
	}
	if true {
		toSerialize["transacted_at_utc"] = o.TransactedAtUtc
	}
	return json.Marshal(toSerialize)
}

type NullableFinancialAccountTransactionSipTransaction struct {
	value *FinancialAccountTransactionSipTransaction
	isSet bool
}

func (v NullableFinancialAccountTransactionSipTransaction) Get() *FinancialAccountTransactionSipTransaction {
	return v.value
}

func (v *NullableFinancialAccountTransactionSipTransaction) Set(val *FinancialAccountTransactionSipTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialAccountTransactionSipTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialAccountTransactionSipTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialAccountTransactionSipTransaction(val *FinancialAccountTransactionSipTransaction) *NullableFinancialAccountTransactionSipTransaction {
	return &NullableFinancialAccountTransactionSipTransaction{value: val, isSet: true}
}

func (v NullableFinancialAccountTransactionSipTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialAccountTransactionSipTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



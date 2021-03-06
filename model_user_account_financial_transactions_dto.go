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
)

// UserAccountFinancialTransactionsDto struct for UserAccountFinancialTransactionsDto
type UserAccountFinancialTransactionsDto struct {
	Id *string `json:"id,omitempty"`
	AccountId *string `json:"accountId,omitempty"`
	Name NullableString `json:"name,omitempty"`
	InstrumentId *string `json:"instrumentId,omitempty"`
	TransactionType NullableString `json:"transactionType,omitempty"`
	Quantity *int32 `json:"quantity,omitempty"`
	AveragePrice *float64 `json:"averagePrice,omitempty"`
	Currency NullableString `json:"currency,omitempty"`
}

// NewUserAccountFinancialTransactionsDto instantiates a new UserAccountFinancialTransactionsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAccountFinancialTransactionsDto() *UserAccountFinancialTransactionsDto {
	this := UserAccountFinancialTransactionsDto{}
	return &this
}

// NewUserAccountFinancialTransactionsDtoWithDefaults instantiates a new UserAccountFinancialTransactionsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAccountFinancialTransactionsDtoWithDefaults() *UserAccountFinancialTransactionsDto {
	this := UserAccountFinancialTransactionsDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserAccountFinancialTransactionsDto) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccountFinancialTransactionsDto) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserAccountFinancialTransactionsDto) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserAccountFinancialTransactionsDto) SetId(v string) {
	o.Id = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *UserAccountFinancialTransactionsDto) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccountFinancialTransactionsDto) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *UserAccountFinancialTransactionsDto) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *UserAccountFinancialTransactionsDto) SetAccountId(v string) {
	o.AccountId = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserAccountFinancialTransactionsDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserAccountFinancialTransactionsDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UserAccountFinancialTransactionsDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UserAccountFinancialTransactionsDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *UserAccountFinancialTransactionsDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UserAccountFinancialTransactionsDto) UnsetName() {
	o.Name.Unset()
}

// GetInstrumentId returns the InstrumentId field value if set, zero value otherwise.
func (o *UserAccountFinancialTransactionsDto) GetInstrumentId() string {
	if o == nil || o.InstrumentId == nil {
		var ret string
		return ret
	}
	return *o.InstrumentId
}

// GetInstrumentIdOk returns a tuple with the InstrumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccountFinancialTransactionsDto) GetInstrumentIdOk() (*string, bool) {
	if o == nil || o.InstrumentId == nil {
		return nil, false
	}
	return o.InstrumentId, true
}

// HasInstrumentId returns a boolean if a field has been set.
func (o *UserAccountFinancialTransactionsDto) HasInstrumentId() bool {
	if o != nil && o.InstrumentId != nil {
		return true
	}

	return false
}

// SetInstrumentId gets a reference to the given string and assigns it to the InstrumentId field.
func (o *UserAccountFinancialTransactionsDto) SetInstrumentId(v string) {
	o.InstrumentId = &v
}

// GetTransactionType returns the TransactionType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserAccountFinancialTransactionsDto) GetTransactionType() string {
	if o == nil || o.TransactionType.Get() == nil {
		var ret string
		return ret
	}
	return *o.TransactionType.Get()
}

// GetTransactionTypeOk returns a tuple with the TransactionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserAccountFinancialTransactionsDto) GetTransactionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransactionType.Get(), o.TransactionType.IsSet()
}

// HasTransactionType returns a boolean if a field has been set.
func (o *UserAccountFinancialTransactionsDto) HasTransactionType() bool {
	if o != nil && o.TransactionType.IsSet() {
		return true
	}

	return false
}

// SetTransactionType gets a reference to the given NullableString and assigns it to the TransactionType field.
func (o *UserAccountFinancialTransactionsDto) SetTransactionType(v string) {
	o.TransactionType.Set(&v)
}
// SetTransactionTypeNil sets the value for TransactionType to be an explicit nil
func (o *UserAccountFinancialTransactionsDto) SetTransactionTypeNil() {
	o.TransactionType.Set(nil)
}

// UnsetTransactionType ensures that no value is present for TransactionType, not even an explicit nil
func (o *UserAccountFinancialTransactionsDto) UnsetTransactionType() {
	o.TransactionType.Unset()
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *UserAccountFinancialTransactionsDto) GetQuantity() int32 {
	if o == nil || o.Quantity == nil {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccountFinancialTransactionsDto) GetQuantityOk() (*int32, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *UserAccountFinancialTransactionsDto) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *UserAccountFinancialTransactionsDto) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetAveragePrice returns the AveragePrice field value if set, zero value otherwise.
func (o *UserAccountFinancialTransactionsDto) GetAveragePrice() float64 {
	if o == nil || o.AveragePrice == nil {
		var ret float64
		return ret
	}
	return *o.AveragePrice
}

// GetAveragePriceOk returns a tuple with the AveragePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccountFinancialTransactionsDto) GetAveragePriceOk() (*float64, bool) {
	if o == nil || o.AveragePrice == nil {
		return nil, false
	}
	return o.AveragePrice, true
}

// HasAveragePrice returns a boolean if a field has been set.
func (o *UserAccountFinancialTransactionsDto) HasAveragePrice() bool {
	if o != nil && o.AveragePrice != nil {
		return true
	}

	return false
}

// SetAveragePrice gets a reference to the given float64 and assigns it to the AveragePrice field.
func (o *UserAccountFinancialTransactionsDto) SetAveragePrice(v float64) {
	o.AveragePrice = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserAccountFinancialTransactionsDto) GetCurrency() string {
	if o == nil || o.Currency.Get() == nil {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserAccountFinancialTransactionsDto) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *UserAccountFinancialTransactionsDto) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *UserAccountFinancialTransactionsDto) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *UserAccountFinancialTransactionsDto) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *UserAccountFinancialTransactionsDto) UnsetCurrency() {
	o.Currency.Unset()
}

func (o UserAccountFinancialTransactionsDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.AccountId != nil {
		toSerialize["accountId"] = o.AccountId
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.InstrumentId != nil {
		toSerialize["instrumentId"] = o.InstrumentId
	}
	if o.TransactionType.IsSet() {
		toSerialize["transactionType"] = o.TransactionType.Get()
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.AveragePrice != nil {
		toSerialize["averagePrice"] = o.AveragePrice
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUserAccountFinancialTransactionsDto struct {
	value *UserAccountFinancialTransactionsDto
	isSet bool
}

func (v NullableUserAccountFinancialTransactionsDto) Get() *UserAccountFinancialTransactionsDto {
	return v.value
}

func (v *NullableUserAccountFinancialTransactionsDto) Set(val *UserAccountFinancialTransactionsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAccountFinancialTransactionsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAccountFinancialTransactionsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAccountFinancialTransactionsDto(val *UserAccountFinancialTransactionsDto) *NullableUserAccountFinancialTransactionsDto {
	return &NullableUserAccountFinancialTransactionsDto{value: val, isSet: true}
}

func (v NullableUserAccountFinancialTransactionsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAccountFinancialTransactionsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// Financial struct for Financial
type Financial struct {
	FieldName NullableString `json:"field_name,omitempty"`
	CustomKey NullableString `json:"custom_key,omitempty"`
	Accounts []FinancialAccounts `json:"accounts,omitempty"`
	Requirement *DocumentsRequired `json:"requirement,omitempty"`
}

// NewFinancial instantiates a new Financial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancial() *Financial {
	this := Financial{}
	return &this
}

// NewFinancialWithDefaults instantiates a new Financial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialWithDefaults() *Financial {
	this := Financial{}
	return &this
}

// GetFieldName returns the FieldName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Financial) GetFieldName() string {
	if o == nil || o.FieldName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FieldName.Get()
}

// GetFieldNameOk returns a tuple with the FieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Financial) GetFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FieldName.Get(), o.FieldName.IsSet()
}

// HasFieldName returns a boolean if a field has been set.
func (o *Financial) HasFieldName() bool {
	if o != nil && o.FieldName.IsSet() {
		return true
	}

	return false
}

// SetFieldName gets a reference to the given NullableString and assigns it to the FieldName field.
func (o *Financial) SetFieldName(v string) {
	o.FieldName.Set(&v)
}
// SetFieldNameNil sets the value for FieldName to be an explicit nil
func (o *Financial) SetFieldNameNil() {
	o.FieldName.Set(nil)
}

// UnsetFieldName ensures that no value is present for FieldName, not even an explicit nil
func (o *Financial) UnsetFieldName() {
	o.FieldName.Unset()
}

// GetCustomKey returns the CustomKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Financial) GetCustomKey() string {
	if o == nil || o.CustomKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.CustomKey.Get()
}

// GetCustomKeyOk returns a tuple with the CustomKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Financial) GetCustomKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomKey.Get(), o.CustomKey.IsSet()
}

// HasCustomKey returns a boolean if a field has been set.
func (o *Financial) HasCustomKey() bool {
	if o != nil && o.CustomKey.IsSet() {
		return true
	}

	return false
}

// SetCustomKey gets a reference to the given NullableString and assigns it to the CustomKey field.
func (o *Financial) SetCustomKey(v string) {
	o.CustomKey.Set(&v)
}
// SetCustomKeyNil sets the value for CustomKey to be an explicit nil
func (o *Financial) SetCustomKeyNil() {
	o.CustomKey.Set(nil)
}

// UnsetCustomKey ensures that no value is present for CustomKey, not even an explicit nil
func (o *Financial) UnsetCustomKey() {
	o.CustomKey.Unset()
}

// GetAccounts returns the Accounts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Financial) GetAccounts() []FinancialAccounts {
	if o == nil {
		var ret []FinancialAccounts
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Financial) GetAccountsOk() ([]FinancialAccounts, bool) {
	if o == nil || o.Accounts == nil {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *Financial) HasAccounts() bool {
	if o != nil && o.Accounts != nil {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []FinancialAccounts and assigns it to the Accounts field.
func (o *Financial) SetAccounts(v []FinancialAccounts) {
	o.Accounts = v
}

// GetRequirement returns the Requirement field value if set, zero value otherwise.
func (o *Financial) GetRequirement() DocumentsRequired {
	if o == nil || o.Requirement == nil {
		var ret DocumentsRequired
		return ret
	}
	return *o.Requirement
}

// GetRequirementOk returns a tuple with the Requirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Financial) GetRequirementOk() (*DocumentsRequired, bool) {
	if o == nil || o.Requirement == nil {
		return nil, false
	}
	return o.Requirement, true
}

// HasRequirement returns a boolean if a field has been set.
func (o *Financial) HasRequirement() bool {
	if o != nil && o.Requirement != nil {
		return true
	}

	return false
}

// SetRequirement gets a reference to the given DocumentsRequired and assigns it to the Requirement field.
func (o *Financial) SetRequirement(v DocumentsRequired) {
	o.Requirement = &v
}

func (o Financial) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FieldName.IsSet() {
		toSerialize["field_name"] = o.FieldName.Get()
	}
	if o.CustomKey.IsSet() {
		toSerialize["custom_key"] = o.CustomKey.Get()
	}
	if o.Accounts != nil {
		toSerialize["accounts"] = o.Accounts
	}
	if o.Requirement != nil {
		toSerialize["requirement"] = o.Requirement
	}
	return json.Marshal(toSerialize)
}

type NullableFinancial struct {
	value *Financial
	isSet bool
}

func (v NullableFinancial) Get() *Financial {
	return v.value
}

func (v *NullableFinancial) Set(val *Financial) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancial) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancial(val *Financial) *NullableFinancial {
	return &NullableFinancial{value: val, isSet: true}
}

func (v NullableFinancial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



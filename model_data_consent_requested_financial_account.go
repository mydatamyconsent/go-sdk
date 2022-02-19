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

// DataConsentRequestedFinancialAccount struct for DataConsentRequestedFinancialAccount
type DataConsentRequestedFinancialAccount struct {
	CustomKey NullableString `json:"custom_key,omitempty"`
	Drn NullableString `json:"drn,omitempty"`
	AccountTypeId NullableString `json:"accountTypeId,omitempty"`
	AccountIdentifier NullableString `json:"accountIdentifier,omitempty"`
}

// NewDataConsentRequestedFinancialAccount instantiates a new DataConsentRequestedFinancialAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataConsentRequestedFinancialAccount() *DataConsentRequestedFinancialAccount {
	this := DataConsentRequestedFinancialAccount{}
	return &this
}

// NewDataConsentRequestedFinancialAccountWithDefaults instantiates a new DataConsentRequestedFinancialAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataConsentRequestedFinancialAccountWithDefaults() *DataConsentRequestedFinancialAccount {
	this := DataConsentRequestedFinancialAccount{}
	return &this
}

// GetCustomKey returns the CustomKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsentRequestedFinancialAccount) GetCustomKey() string {
	if o == nil || o.CustomKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.CustomKey.Get()
}

// GetCustomKeyOk returns a tuple with the CustomKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsentRequestedFinancialAccount) GetCustomKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CustomKey.Get(), o.CustomKey.IsSet()
}

// HasCustomKey returns a boolean if a field has been set.
func (o *DataConsentRequestedFinancialAccount) HasCustomKey() bool {
	if o != nil && o.CustomKey.IsSet() {
		return true
	}

	return false
}

// SetCustomKey gets a reference to the given NullableString and assigns it to the CustomKey field.
func (o *DataConsentRequestedFinancialAccount) SetCustomKey(v string) {
	o.CustomKey.Set(&v)
}
// SetCustomKeyNil sets the value for CustomKey to be an explicit nil
func (o *DataConsentRequestedFinancialAccount) SetCustomKeyNil() {
	o.CustomKey.Set(nil)
}

// UnsetCustomKey ensures that no value is present for CustomKey, not even an explicit nil
func (o *DataConsentRequestedFinancialAccount) UnsetCustomKey() {
	o.CustomKey.Unset()
}

// GetDrn returns the Drn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsentRequestedFinancialAccount) GetDrn() string {
	if o == nil || o.Drn.Get() == nil {
		var ret string
		return ret
	}
	return *o.Drn.Get()
}

// GetDrnOk returns a tuple with the Drn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsentRequestedFinancialAccount) GetDrnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Drn.Get(), o.Drn.IsSet()
}

// HasDrn returns a boolean if a field has been set.
func (o *DataConsentRequestedFinancialAccount) HasDrn() bool {
	if o != nil && o.Drn.IsSet() {
		return true
	}

	return false
}

// SetDrn gets a reference to the given NullableString and assigns it to the Drn field.
func (o *DataConsentRequestedFinancialAccount) SetDrn(v string) {
	o.Drn.Set(&v)
}
// SetDrnNil sets the value for Drn to be an explicit nil
func (o *DataConsentRequestedFinancialAccount) SetDrnNil() {
	o.Drn.Set(nil)
}

// UnsetDrn ensures that no value is present for Drn, not even an explicit nil
func (o *DataConsentRequestedFinancialAccount) UnsetDrn() {
	o.Drn.Unset()
}

// GetAccountTypeId returns the AccountTypeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsentRequestedFinancialAccount) GetAccountTypeId() string {
	if o == nil || o.AccountTypeId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountTypeId.Get()
}

// GetAccountTypeIdOk returns a tuple with the AccountTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsentRequestedFinancialAccount) GetAccountTypeIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountTypeId.Get(), o.AccountTypeId.IsSet()
}

// HasAccountTypeId returns a boolean if a field has been set.
func (o *DataConsentRequestedFinancialAccount) HasAccountTypeId() bool {
	if o != nil && o.AccountTypeId.IsSet() {
		return true
	}

	return false
}

// SetAccountTypeId gets a reference to the given NullableString and assigns it to the AccountTypeId field.
func (o *DataConsentRequestedFinancialAccount) SetAccountTypeId(v string) {
	o.AccountTypeId.Set(&v)
}
// SetAccountTypeIdNil sets the value for AccountTypeId to be an explicit nil
func (o *DataConsentRequestedFinancialAccount) SetAccountTypeIdNil() {
	o.AccountTypeId.Set(nil)
}

// UnsetAccountTypeId ensures that no value is present for AccountTypeId, not even an explicit nil
func (o *DataConsentRequestedFinancialAccount) UnsetAccountTypeId() {
	o.AccountTypeId.Unset()
}

// GetAccountIdentifier returns the AccountIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsentRequestedFinancialAccount) GetAccountIdentifier() string {
	if o == nil || o.AccountIdentifier.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountIdentifier.Get()
}

// GetAccountIdentifierOk returns a tuple with the AccountIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsentRequestedFinancialAccount) GetAccountIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountIdentifier.Get(), o.AccountIdentifier.IsSet()
}

// HasAccountIdentifier returns a boolean if a field has been set.
func (o *DataConsentRequestedFinancialAccount) HasAccountIdentifier() bool {
	if o != nil && o.AccountIdentifier.IsSet() {
		return true
	}

	return false
}

// SetAccountIdentifier gets a reference to the given NullableString and assigns it to the AccountIdentifier field.
func (o *DataConsentRequestedFinancialAccount) SetAccountIdentifier(v string) {
	o.AccountIdentifier.Set(&v)
}
// SetAccountIdentifierNil sets the value for AccountIdentifier to be an explicit nil
func (o *DataConsentRequestedFinancialAccount) SetAccountIdentifierNil() {
	o.AccountIdentifier.Set(nil)
}

// UnsetAccountIdentifier ensures that no value is present for AccountIdentifier, not even an explicit nil
func (o *DataConsentRequestedFinancialAccount) UnsetAccountIdentifier() {
	o.AccountIdentifier.Unset()
}

func (o DataConsentRequestedFinancialAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomKey.IsSet() {
		toSerialize["custom_key"] = o.CustomKey.Get()
	}
	if o.Drn.IsSet() {
		toSerialize["drn"] = o.Drn.Get()
	}
	if o.AccountTypeId.IsSet() {
		toSerialize["accountTypeId"] = o.AccountTypeId.Get()
	}
	if o.AccountIdentifier.IsSet() {
		toSerialize["accountIdentifier"] = o.AccountIdentifier.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDataConsentRequestedFinancialAccount struct {
	value *DataConsentRequestedFinancialAccount
	isSet bool
}

func (v NullableDataConsentRequestedFinancialAccount) Get() *DataConsentRequestedFinancialAccount {
	return v.value
}

func (v *NullableDataConsentRequestedFinancialAccount) Set(val *DataConsentRequestedFinancialAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableDataConsentRequestedFinancialAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableDataConsentRequestedFinancialAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataConsentRequestedFinancialAccount(val *DataConsentRequestedFinancialAccount) *NullableDataConsentRequestedFinancialAccount {
	return &NullableDataConsentRequestedFinancialAccount{value: val, isSet: true}
}

func (v NullableDataConsentRequestedFinancialAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataConsentRequestedFinancialAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



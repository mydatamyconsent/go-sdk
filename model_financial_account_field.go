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

// FinancialAccountField FinancialAccountField : Financial account field of consent request template.
type FinancialAccountField struct {
	// Field title.
	FieldTitle string `json:"fieldTitle"`
	// Field slug.
	FieldSlug string `json:"fieldSlug"`
	// Selected financial account types.
	AccountTypes []SelectedFinancialAccountType `json:"accountTypes"`
}

// NewFinancialAccountField instantiates a new FinancialAccountField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialAccountField(fieldTitle string, fieldSlug string, accountTypes []SelectedFinancialAccountType) *FinancialAccountField {
	this := FinancialAccountField{}
	this.FieldTitle = fieldTitle
	this.FieldSlug = fieldSlug
	this.AccountTypes = accountTypes
	return &this
}

// NewFinancialAccountFieldWithDefaults instantiates a new FinancialAccountField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialAccountFieldWithDefaults() *FinancialAccountField {
	this := FinancialAccountField{}
	return &this
}

// GetFieldTitle returns the FieldTitle field value
func (o *FinancialAccountField) GetFieldTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldTitle
}

// GetFieldTitleOk returns a tuple with the FieldTitle field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountField) GetFieldTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldTitle, true
}

// SetFieldTitle sets field value
func (o *FinancialAccountField) SetFieldTitle(v string) {
	o.FieldTitle = v
}

// GetFieldSlug returns the FieldSlug field value
func (o *FinancialAccountField) GetFieldSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldSlug
}

// GetFieldSlugOk returns a tuple with the FieldSlug field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountField) GetFieldSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldSlug, true
}

// SetFieldSlug sets field value
func (o *FinancialAccountField) SetFieldSlug(v string) {
	o.FieldSlug = v
}

// GetAccountTypes returns the AccountTypes field value
func (o *FinancialAccountField) GetAccountTypes() []SelectedFinancialAccountType {
	if o == nil {
		var ret []SelectedFinancialAccountType
		return ret
	}

	return o.AccountTypes
}

// GetAccountTypesOk returns a tuple with the AccountTypes field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountField) GetAccountTypesOk() ([]SelectedFinancialAccountType, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountTypes, true
}

// SetAccountTypes sets field value
func (o *FinancialAccountField) SetAccountTypes(v []SelectedFinancialAccountType) {
	o.AccountTypes = v
}

func (o FinancialAccountField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fieldTitle"] = o.FieldTitle
	}
	if true {
		toSerialize["fieldSlug"] = o.FieldSlug
	}
	if true {
		toSerialize["accountTypes"] = o.AccountTypes
	}
	return json.Marshal(toSerialize)
}

type NullableFinancialAccountField struct {
	value *FinancialAccountField
	isSet bool
}

func (v NullableFinancialAccountField) Get() *FinancialAccountField {
	return v.value
}

func (v *NullableFinancialAccountField) Set(val *FinancialAccountField) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialAccountField) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialAccountField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialAccountField(val *FinancialAccountField) *NullableFinancialAccountField {
	return &NullableFinancialAccountField{value: val, isSet: true}
}

func (v NullableFinancialAccountField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialAccountField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



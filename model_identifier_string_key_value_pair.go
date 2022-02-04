/*
My Data My Consent - Developer API

Unleashing the power of data consent by establishing trust. The Platform Core Developer API defines a set of capabilities that can be used to request, issue, manage and update data, documents and credentials by organizations. The API can be used to request, manage and update Decentralised Identifiers, Financial Data, Health Data issue Documents, Credentials directly or using OpenID Connect flows, and verify Messages signed with DIDs and much more.

API version: v1
Contact: support@mydatamyconsent.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github.com/mydatamyconsent/go-sdk

import (
	"encoding/json"
)

// IdentifierStringKeyValuePair struct for IdentifierStringKeyValuePair
type IdentifierStringKeyValuePair struct {
	Key *Identifier `json:"key,omitempty"`
	Value NullableString `json:"value,omitempty"`
}

// NewIdentifierStringKeyValuePair instantiates a new IdentifierStringKeyValuePair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentifierStringKeyValuePair() *IdentifierStringKeyValuePair {
	this := IdentifierStringKeyValuePair{}
	return &this
}

// NewIdentifierStringKeyValuePairWithDefaults instantiates a new IdentifierStringKeyValuePair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentifierStringKeyValuePairWithDefaults() *IdentifierStringKeyValuePair {
	this := IdentifierStringKeyValuePair{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *IdentifierStringKeyValuePair) GetKey() Identifier {
	if o == nil || o.Key == nil {
		var ret Identifier
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentifierStringKeyValuePair) GetKeyOk() (*Identifier, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *IdentifierStringKeyValuePair) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given Identifier and assigns it to the Key field.
func (o *IdentifierStringKeyValuePair) SetKey(v Identifier) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentifierStringKeyValuePair) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentifierStringKeyValuePair) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *IdentifierStringKeyValuePair) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *IdentifierStringKeyValuePair) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *IdentifierStringKeyValuePair) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *IdentifierStringKeyValuePair) UnsetValue() {
	o.Value.Unset()
}

func (o IdentifierStringKeyValuePair) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIdentifierStringKeyValuePair struct {
	value *IdentifierStringKeyValuePair
	isSet bool
}

func (v NullableIdentifierStringKeyValuePair) Get() *IdentifierStringKeyValuePair {
	return v.value
}

func (v *NullableIdentifierStringKeyValuePair) Set(val *IdentifierStringKeyValuePair) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentifierStringKeyValuePair) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentifierStringKeyValuePair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentifierStringKeyValuePair(val *IdentifierStringKeyValuePair) *NullableIdentifierStringKeyValuePair {
	return &NullableIdentifierStringKeyValuePair{value: val, isSet: true}
}

func (v NullableIdentifierStringKeyValuePair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentifierStringKeyValuePair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// DataConsentIdentifier struct for DataConsentIdentifier
type DataConsentIdentifier struct {
	Key NullableString `json:"key,omitempty"`
	Value NullableString `json:"value,omitempty"`
}

// NewDataConsentIdentifier instantiates a new DataConsentIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataConsentIdentifier() *DataConsentIdentifier {
	this := DataConsentIdentifier{}
	return &this
}

// NewDataConsentIdentifierWithDefaults instantiates a new DataConsentIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataConsentIdentifierWithDefaults() *DataConsentIdentifier {
	this := DataConsentIdentifier{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsentIdentifier) GetKey() string {
	if o == nil || o.Key.Get() == nil {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsentIdentifier) GetKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *DataConsentIdentifier) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *DataConsentIdentifier) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *DataConsentIdentifier) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *DataConsentIdentifier) UnsetKey() {
	o.Key.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsentIdentifier) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsentIdentifier) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *DataConsentIdentifier) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *DataConsentIdentifier) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *DataConsentIdentifier) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *DataConsentIdentifier) UnsetValue() {
	o.Value.Unset()
}

func (o DataConsentIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDataConsentIdentifier struct {
	value *DataConsentIdentifier
	isSet bool
}

func (v NullableDataConsentIdentifier) Get() *DataConsentIdentifier {
	return v.value
}

func (v *NullableDataConsentIdentifier) Set(val *DataConsentIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableDataConsentIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableDataConsentIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataConsentIdentifier(val *DataConsentIdentifier) *NullableDataConsentIdentifier {
	return &NullableDataConsentIdentifier{value: val, isSet: true}
}

func (v NullableDataConsentIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataConsentIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



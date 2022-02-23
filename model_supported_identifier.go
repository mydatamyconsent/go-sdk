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

// SupportedIdentifier struct for SupportedIdentifier
type SupportedIdentifier struct {
	Key NullableString `json:"key,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
	ExampleValue NullableString `json:"exampleValue,omitempty"`
}

// NewSupportedIdentifier instantiates a new SupportedIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportedIdentifier() *SupportedIdentifier {
	this := SupportedIdentifier{}
	return &this
}

// NewSupportedIdentifierWithDefaults instantiates a new SupportedIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportedIdentifierWithDefaults() *SupportedIdentifier {
	this := SupportedIdentifier{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SupportedIdentifier) GetKey() string {
	if o == nil || o.Key.Get() == nil {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SupportedIdentifier) GetKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *SupportedIdentifier) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *SupportedIdentifier) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *SupportedIdentifier) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *SupportedIdentifier) UnsetKey() {
	o.Key.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SupportedIdentifier) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SupportedIdentifier) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SupportedIdentifier) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SupportedIdentifier) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SupportedIdentifier) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SupportedIdentifier) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SupportedIdentifier) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SupportedIdentifier) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *SupportedIdentifier) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *SupportedIdentifier) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *SupportedIdentifier) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *SupportedIdentifier) UnsetDescription() {
	o.Description.Unset()
}

// GetExampleValue returns the ExampleValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SupportedIdentifier) GetExampleValue() string {
	if o == nil || o.ExampleValue.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExampleValue.Get()
}

// GetExampleValueOk returns a tuple with the ExampleValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SupportedIdentifier) GetExampleValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExampleValue.Get(), o.ExampleValue.IsSet()
}

// HasExampleValue returns a boolean if a field has been set.
func (o *SupportedIdentifier) HasExampleValue() bool {
	if o != nil && o.ExampleValue.IsSet() {
		return true
	}

	return false
}

// SetExampleValue gets a reference to the given NullableString and assigns it to the ExampleValue field.
func (o *SupportedIdentifier) SetExampleValue(v string) {
	o.ExampleValue.Set(&v)
}
// SetExampleValueNil sets the value for ExampleValue to be an explicit nil
func (o *SupportedIdentifier) SetExampleValueNil() {
	o.ExampleValue.Set(nil)
}

// UnsetExampleValue ensures that no value is present for ExampleValue, not even an explicit nil
func (o *SupportedIdentifier) UnsetExampleValue() {
	o.ExampleValue.Unset()
}

func (o SupportedIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.ExampleValue.IsSet() {
		toSerialize["exampleValue"] = o.ExampleValue.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSupportedIdentifier struct {
	value *SupportedIdentifier
	isSet bool
}

func (v NullableSupportedIdentifier) Get() *SupportedIdentifier {
	return v.value
}

func (v *NullableSupportedIdentifier) Set(val *SupportedIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportedIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportedIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportedIdentifier(val *SupportedIdentifier) *NullableSupportedIdentifier {
	return &NullableSupportedIdentifier{value: val, isSet: true}
}

func (v NullableSupportedIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportedIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



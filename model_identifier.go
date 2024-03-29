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

// Identifier Identifier details.
type Identifier struct {
	// Identifier key. EMAIL, MOBILE_NUMBER, etc.
	Key string `json:"key"`
	// Identifier name. Email, Mobile Number, etc.
	Name string `json:"name"`
	// Identifier description. User's email, User's mobile number, etc.
	Description string `json:"description"`
	// Example value. example@email.com, +919090909090, etc.
	ExampleValue string `json:"example_value"`
}

// NewIdentifier instantiates a new Identifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentifier(key string, name string, description string, exampleValue string) *Identifier {
	this := Identifier{}
	this.Key = key
	this.Name = name
	this.Description = description
	this.ExampleValue = exampleValue
	return &this
}

// NewIdentifierWithDefaults instantiates a new Identifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentifierWithDefaults() *Identifier {
	this := Identifier{}
	return &this
}

// GetKey returns the Key field value
func (o *Identifier) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *Identifier) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *Identifier) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *Identifier) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Identifier) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Identifier) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *Identifier) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Identifier) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Identifier) SetDescription(v string) {
	o.Description = v
}

// GetExampleValue returns the ExampleValue field value
func (o *Identifier) GetExampleValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExampleValue
}

// GetExampleValueOk returns a tuple with the ExampleValue field value
// and a boolean to check if the value has been set.
func (o *Identifier) GetExampleValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExampleValue, true
}

// SetExampleValue sets field value
func (o *Identifier) SetExampleValue(v string) {
	o.ExampleValue = v
}

func (o Identifier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["example_value"] = o.ExampleValue
	}
	return json.Marshal(toSerialize)
}

type NullableIdentifier struct {
	value *Identifier
	isSet bool
}

func (v NullableIdentifier) Get() *Identifier {
	return v.value
}

func (v *NullableIdentifier) Set(val *Identifier) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentifier(val *Identifier) *NullableIdentifier {
	return &NullableIdentifier{value: val, isSet: true}
}

func (v NullableIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



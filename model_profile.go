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

// Profile struct for Profile
type Profile struct {
	Holders Holders `json:"holders"`
}

// NewProfile instantiates a new Profile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfile(holders Holders) *Profile {
	this := Profile{}
	this.Holders = holders
	return &this
}

// NewProfileWithDefaults instantiates a new Profile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileWithDefaults() *Profile {
	this := Profile{}
	return &this
}

// GetHolders returns the Holders field value
func (o *Profile) GetHolders() Holders {
	if o == nil {
		var ret Holders
		return ret
	}

	return o.Holders
}

// GetHoldersOk returns a tuple with the Holders field value
// and a boolean to check if the value has been set.
func (o *Profile) GetHoldersOk() (*Holders, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Holders, true
}

// SetHolders sets field value
func (o *Profile) SetHolders(v Holders) {
	o.Holders = v
}

func (o Profile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["holders"] = o.Holders
	}
	return json.Marshal(toSerialize)
}

type NullableProfile struct {
	value *Profile
	isSet bool
}

func (v NullableProfile) Get() *Profile {
	return v.value
}

func (v *NullableProfile) Set(val *Profile) {
	v.value = val
	v.isSet = true
}

func (v NullableProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfile(val *Profile) *NullableProfile {
	return &NullableProfile{value: val, isSet: true}
}

func (v NullableProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


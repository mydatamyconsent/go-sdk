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

// ConsentRequestReceiver Consent request receiver details
type ConsentRequestReceiver struct {
	// Consent request receiver country ISO 2 code
	CountryIso2Code string `json:"countryIso2Code"`
	// Consent request receiver identifiers
	Identifiers []StringStringKeyValuePair `json:"identifiers"`
	IdentificationStrategy IdentificationStrategy `json:"identificationStrategy"`
}

// NewConsentRequestReceiver instantiates a new ConsentRequestReceiver object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsentRequestReceiver(countryIso2Code string, identifiers []StringStringKeyValuePair, identificationStrategy IdentificationStrategy) *ConsentRequestReceiver {
	this := ConsentRequestReceiver{}
	this.CountryIso2Code = countryIso2Code
	this.Identifiers = identifiers
	this.IdentificationStrategy = identificationStrategy
	return &this
}

// NewConsentRequestReceiverWithDefaults instantiates a new ConsentRequestReceiver object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsentRequestReceiverWithDefaults() *ConsentRequestReceiver {
	this := ConsentRequestReceiver{}
	return &this
}

// GetCountryIso2Code returns the CountryIso2Code field value
func (o *ConsentRequestReceiver) GetCountryIso2Code() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryIso2Code
}

// GetCountryIso2CodeOk returns a tuple with the CountryIso2Code field value
// and a boolean to check if the value has been set.
func (o *ConsentRequestReceiver) GetCountryIso2CodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CountryIso2Code, true
}

// SetCountryIso2Code sets field value
func (o *ConsentRequestReceiver) SetCountryIso2Code(v string) {
	o.CountryIso2Code = v
}

// GetIdentifiers returns the Identifiers field value
func (o *ConsentRequestReceiver) GetIdentifiers() []StringStringKeyValuePair {
	if o == nil {
		var ret []StringStringKeyValuePair
		return ret
	}

	return o.Identifiers
}

// GetIdentifiersOk returns a tuple with the Identifiers field value
// and a boolean to check if the value has been set.
func (o *ConsentRequestReceiver) GetIdentifiersOk() ([]StringStringKeyValuePair, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Identifiers, true
}

// SetIdentifiers sets field value
func (o *ConsentRequestReceiver) SetIdentifiers(v []StringStringKeyValuePair) {
	o.Identifiers = v
}

// GetIdentificationStrategy returns the IdentificationStrategy field value
func (o *ConsentRequestReceiver) GetIdentificationStrategy() IdentificationStrategy {
	if o == nil {
		var ret IdentificationStrategy
		return ret
	}

	return o.IdentificationStrategy
}

// GetIdentificationStrategyOk returns a tuple with the IdentificationStrategy field value
// and a boolean to check if the value has been set.
func (o *ConsentRequestReceiver) GetIdentificationStrategyOk() (*IdentificationStrategy, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IdentificationStrategy, true
}

// SetIdentificationStrategy sets field value
func (o *ConsentRequestReceiver) SetIdentificationStrategy(v IdentificationStrategy) {
	o.IdentificationStrategy = v
}

func (o ConsentRequestReceiver) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["countryIso2Code"] = o.CountryIso2Code
	}
	if true {
		toSerialize["identifiers"] = o.Identifiers
	}
	if true {
		toSerialize["identificationStrategy"] = o.IdentificationStrategy
	}
	return json.Marshal(toSerialize)
}

type NullableConsentRequestReceiver struct {
	value *ConsentRequestReceiver
	isSet bool
}

func (v NullableConsentRequestReceiver) Get() *ConsentRequestReceiver {
	return v.value
}

func (v *NullableConsentRequestReceiver) Set(val *ConsentRequestReceiver) {
	v.value = val
	v.isSet = true
}

func (v NullableConsentRequestReceiver) IsSet() bool {
	return v.isSet
}

func (v *NullableConsentRequestReceiver) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsentRequestReceiver(val *ConsentRequestReceiver) *NullableConsentRequestReceiver {
	return &NullableConsentRequestReceiver{value: val, isSet: true}
}

func (v NullableConsentRequestReceiver) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsentRequestReceiver) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



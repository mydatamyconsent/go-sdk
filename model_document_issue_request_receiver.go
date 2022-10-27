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

// DocumentIssueRequestReceiver struct for DocumentIssueRequestReceiver
type DocumentIssueRequestReceiver struct {
	CountryIso2Code string `json:"countryIso2Code"`
	Identifiers []KeyValuePair `json:"identifiers"`
	IdentificationStrategy IdentificationStrategy `json:"identificationStrategy"`
}

// NewDocumentIssueRequestReceiver instantiates a new DocumentIssueRequestReceiver object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentIssueRequestReceiver(countryIso2Code string, identifiers []KeyValuePair, identificationStrategy IdentificationStrategy) *DocumentIssueRequestReceiver {
	this := DocumentIssueRequestReceiver{}
	this.CountryIso2Code = countryIso2Code
	this.Identifiers = identifiers
	this.IdentificationStrategy = identificationStrategy
	return &this
}

// NewDocumentIssueRequestReceiverWithDefaults instantiates a new DocumentIssueRequestReceiver object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentIssueRequestReceiverWithDefaults() *DocumentIssueRequestReceiver {
	this := DocumentIssueRequestReceiver{}
	return &this
}

// GetCountryIso2Code returns the CountryIso2Code field value
func (o *DocumentIssueRequestReceiver) GetCountryIso2Code() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryIso2Code
}

// GetCountryIso2CodeOk returns a tuple with the CountryIso2Code field value
// and a boolean to check if the value has been set.
func (o *DocumentIssueRequestReceiver) GetCountryIso2CodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryIso2Code, true
}

// SetCountryIso2Code sets field value
func (o *DocumentIssueRequestReceiver) SetCountryIso2Code(v string) {
	o.CountryIso2Code = v
}

// GetIdentifiers returns the Identifiers field value
func (o *DocumentIssueRequestReceiver) GetIdentifiers() []KeyValuePair {
	if o == nil {
		var ret []KeyValuePair
		return ret
	}

	return o.Identifiers
}

// GetIdentifiersOk returns a tuple with the Identifiers field value
// and a boolean to check if the value has been set.
func (o *DocumentIssueRequestReceiver) GetIdentifiersOk() ([]KeyValuePair, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identifiers, true
}

// SetIdentifiers sets field value
func (o *DocumentIssueRequestReceiver) SetIdentifiers(v []KeyValuePair) {
	o.Identifiers = v
}

// GetIdentificationStrategy returns the IdentificationStrategy field value
func (o *DocumentIssueRequestReceiver) GetIdentificationStrategy() IdentificationStrategy {
	if o == nil {
		var ret IdentificationStrategy
		return ret
	}

	return o.IdentificationStrategy
}

// GetIdentificationStrategyOk returns a tuple with the IdentificationStrategy field value
// and a boolean to check if the value has been set.
func (o *DocumentIssueRequestReceiver) GetIdentificationStrategyOk() (*IdentificationStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentificationStrategy, true
}

// SetIdentificationStrategy sets field value
func (o *DocumentIssueRequestReceiver) SetIdentificationStrategy(v IdentificationStrategy) {
	o.IdentificationStrategy = v
}

func (o DocumentIssueRequestReceiver) MarshalJSON() ([]byte, error) {
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

type NullableDocumentIssueRequestReceiver struct {
	value *DocumentIssueRequestReceiver
	isSet bool
}

func (v NullableDocumentIssueRequestReceiver) Get() *DocumentIssueRequestReceiver {
	return v.value
}

func (v *NullableDocumentIssueRequestReceiver) Set(val *DocumentIssueRequestReceiver) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentIssueRequestReceiver) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentIssueRequestReceiver) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentIssueRequestReceiver(val *DocumentIssueRequestReceiver) *NullableDocumentIssueRequestReceiver {
	return &NullableDocumentIssueRequestReceiver{value: val, isSet: true}
}

func (v NullableDocumentIssueRequestReceiver) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentIssueRequestReceiver) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


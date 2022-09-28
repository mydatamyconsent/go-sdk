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

// ConsentDocumentIssuer ConsentDocumentIssuer : Consent document issuer details.
type ConsentDocumentIssuer struct {
	// Document issuer id.
	Id string `json:"id"`
	// Document issuer name.
	Name string `json:"name"`
}

// NewConsentDocumentIssuer instantiates a new ConsentDocumentIssuer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsentDocumentIssuer(id string, name string) *ConsentDocumentIssuer {
	this := ConsentDocumentIssuer{}
	this.Id = id
	this.Name = name
	return &this
}

// NewConsentDocumentIssuerWithDefaults instantiates a new ConsentDocumentIssuer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsentDocumentIssuerWithDefaults() *ConsentDocumentIssuer {
	this := ConsentDocumentIssuer{}
	return &this
}

// GetId returns the Id field value
func (o *ConsentDocumentIssuer) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConsentDocumentIssuer) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConsentDocumentIssuer) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ConsentDocumentIssuer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConsentDocumentIssuer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConsentDocumentIssuer) SetName(v string) {
	o.Name = v
}

func (o ConsentDocumentIssuer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableConsentDocumentIssuer struct {
	value *ConsentDocumentIssuer
	isSet bool
}

func (v NullableConsentDocumentIssuer) Get() *ConsentDocumentIssuer {
	return v.value
}

func (v *NullableConsentDocumentIssuer) Set(val *ConsentDocumentIssuer) {
	v.value = val
	v.isSet = true
}

func (v NullableConsentDocumentIssuer) IsSet() bool {
	return v.isSet
}

func (v *NullableConsentDocumentIssuer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsentDocumentIssuer(val *ConsentDocumentIssuer) *NullableConsentDocumentIssuer {
	return &NullableConsentDocumentIssuer{value: val, isSet: true}
}

func (v NullableConsentDocumentIssuer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsentDocumentIssuer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



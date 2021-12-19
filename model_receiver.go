/*
My Data My Consent - Developer API

Unleashing the power of data consent by establishing trust. The Platform Core Developer API defines a set of capabilities that can be used to request, issue, manage and update data, documents and credentials by organizations. The API can be used to request, manage and update Decentralised Identifiers, Financial Data, Health Data issue Documents, Credentials directly or using OpenID Connect flows, and verify Messages signed with DIDs and much more.

API version: v1
Contact: support@mydatamyconsent.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mydatamyconsent

import (
	"encoding/json"
)

// Receiver struct for Receiver
type Receiver struct {
	Type *ReceiverType `json:"type,omitempty"`
	Identifiers []IdentifierStringKeyValuePair `json:"identifiers,omitempty"`
	IdentificationStrategy *IdentificationStrategy `json:"identificationStrategy,omitempty"`
}

// NewReceiver instantiates a new Receiver object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceiver() *Receiver {
	this := Receiver{}
	return &this
}

// NewReceiverWithDefaults instantiates a new Receiver object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceiverWithDefaults() *Receiver {
	this := Receiver{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Receiver) GetType() ReceiverType {
	if o == nil || o.Type == nil {
		var ret ReceiverType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receiver) GetTypeOk() (*ReceiverType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Receiver) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given ReceiverType and assigns it to the Type field.
func (o *Receiver) SetType(v ReceiverType) {
	o.Type = &v
}

// GetIdentifiers returns the Identifiers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Receiver) GetIdentifiers() []IdentifierStringKeyValuePair {
	if o == nil  {
		var ret []IdentifierStringKeyValuePair
		return ret
	}
	return o.Identifiers
}

// GetIdentifiersOk returns a tuple with the Identifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Receiver) GetIdentifiersOk() (*[]IdentifierStringKeyValuePair, bool) {
	if o == nil || o.Identifiers == nil {
		return nil, false
	}
	return &o.Identifiers, true
}

// HasIdentifiers returns a boolean if a field has been set.
func (o *Receiver) HasIdentifiers() bool {
	if o != nil && o.Identifiers != nil {
		return true
	}

	return false
}

// SetIdentifiers gets a reference to the given []IdentifierStringKeyValuePair and assigns it to the Identifiers field.
func (o *Receiver) SetIdentifiers(v []IdentifierStringKeyValuePair) {
	o.Identifiers = v
}

// GetIdentificationStrategy returns the IdentificationStrategy field value if set, zero value otherwise.
func (o *Receiver) GetIdentificationStrategy() IdentificationStrategy {
	if o == nil || o.IdentificationStrategy == nil {
		var ret IdentificationStrategy
		return ret
	}
	return *o.IdentificationStrategy
}

// GetIdentificationStrategyOk returns a tuple with the IdentificationStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receiver) GetIdentificationStrategyOk() (*IdentificationStrategy, bool) {
	if o == nil || o.IdentificationStrategy == nil {
		return nil, false
	}
	return o.IdentificationStrategy, true
}

// HasIdentificationStrategy returns a boolean if a field has been set.
func (o *Receiver) HasIdentificationStrategy() bool {
	if o != nil && o.IdentificationStrategy != nil {
		return true
	}

	return false
}

// SetIdentificationStrategy gets a reference to the given IdentificationStrategy and assigns it to the IdentificationStrategy field.
func (o *Receiver) SetIdentificationStrategy(v IdentificationStrategy) {
	o.IdentificationStrategy = &v
}

func (o Receiver) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Identifiers != nil {
		toSerialize["identifiers"] = o.Identifiers
	}
	if o.IdentificationStrategy != nil {
		toSerialize["identificationStrategy"] = o.IdentificationStrategy
	}
	return json.Marshal(toSerialize)
}

type NullableReceiver struct {
	value *Receiver
	isSet bool
}

func (v NullableReceiver) Get() *Receiver {
	return v.value
}

func (v *NullableReceiver) Set(val *Receiver) {
	v.value = val
	v.isSet = true
}

func (v NullableReceiver) IsSet() bool {
	return v.isSet
}

func (v *NullableReceiver) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceiver(val *Receiver) *NullableReceiver {
	return &NullableReceiver{value: val, isSet: true}
}

func (v NullableReceiver) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceiver) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



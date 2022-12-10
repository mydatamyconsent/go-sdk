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

// FhirBundleany struct for FhirBundleany
type FhirBundleany struct {
	ResourceType string `json:"resourceType"`
	Type string `json:"type"`
	Entry []interface{} `json:"entry"`
}

// NewFhirBundleany instantiates a new FhirBundleany object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFhirBundleany(resourceType string, type_ string, entry []interface{}) *FhirBundleany {
	this := FhirBundleany{}
	this.ResourceType = resourceType
	this.Type = type_
	this.Entry = entry
	return &this
}

// NewFhirBundleanyWithDefaults instantiates a new FhirBundleany object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFhirBundleanyWithDefaults() *FhirBundleany {
	this := FhirBundleany{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *FhirBundleany) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *FhirBundleany) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *FhirBundleany) SetResourceType(v string) {
	o.ResourceType = v
}

// GetType returns the Type field value
func (o *FhirBundleany) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FhirBundleany) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FhirBundleany) SetType(v string) {
	o.Type = v
}

// GetEntry returns the Entry field value
func (o *FhirBundleany) GetEntry() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.Entry
}

// GetEntryOk returns a tuple with the Entry field value
// and a boolean to check if the value has been set.
func (o *FhirBundleany) GetEntryOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Entry, true
}

// SetEntry sets field value
func (o *FhirBundleany) SetEntry(v []interface{}) {
	o.Entry = v
}

func (o FhirBundleany) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resourceType"] = o.ResourceType
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["entry"] = o.Entry
	}
	return json.Marshal(toSerialize)
}

type NullableFhirBundleany struct {
	value *FhirBundleany
	isSet bool
}

func (v NullableFhirBundleany) Get() *FhirBundleany {
	return v.value
}

func (v *NullableFhirBundleany) Set(val *FhirBundleany) {
	v.value = val
	v.isSet = true
}

func (v NullableFhirBundleany) IsSet() bool {
	return v.isSet
}

func (v *NullableFhirBundleany) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFhirBundleany(val *FhirBundleany) *NullableFhirBundleany {
	return &NullableFhirBundleany{value: val, isSet: true}
}

func (v NullableFhirBundleany) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFhirBundleany) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



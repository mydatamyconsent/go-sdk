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

// IdentityField IdentityField : Identity field of Consent request template.
type IdentityField struct {
	// Field title.
	Title string `json:"title"`
	// Field description.
	Description *string `json:"description,omitempty"`
	// Field key.
	Key string `json:"key"`
	// Field data type.
	DataType string `json:"dataType"`
}

// NewIdentityField instantiates a new IdentityField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityField(title string, key string, dataType string) *IdentityField {
	this := IdentityField{}
	this.Title = title
	this.Key = key
	this.DataType = dataType
	return &this
}

// NewIdentityFieldWithDefaults instantiates a new IdentityField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityFieldWithDefaults() *IdentityField {
	this := IdentityField{}
	return &this
}

// GetTitle returns the Title field value
func (o *IdentityField) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *IdentityField) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *IdentityField) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IdentityField) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityField) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IdentityField) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IdentityField) SetDescription(v string) {
	o.Description = &v
}

// GetKey returns the Key field value
func (o *IdentityField) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *IdentityField) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *IdentityField) SetKey(v string) {
	o.Key = v
}

// GetDataType returns the DataType field value
func (o *IdentityField) GetDataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value
// and a boolean to check if the value has been set.
func (o *IdentityField) GetDataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataType, true
}

// SetDataType sets field value
func (o *IdentityField) SetDataType(v string) {
	o.DataType = v
}

func (o IdentityField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["title"] = o.Title
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["dataType"] = o.DataType
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityField struct {
	value *IdentityField
	isSet bool
}

func (v NullableIdentityField) Get() *IdentityField {
	return v.value
}

func (v *NullableIdentityField) Set(val *IdentityField) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityField) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityField(val *IdentityField) *NullableIdentityField {
	return &NullableIdentityField{value: val, isSet: true}
}

func (v NullableIdentityField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



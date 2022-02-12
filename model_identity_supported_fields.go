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

// IdentitySupportedFields struct for IdentitySupportedFields
type IdentitySupportedFields struct {
	IconCodePoint int32 `json:"iconCodePoint"`
	Title string `json:"title"`
	Description NullableString `json:"description,omitempty"`
	Key string `json:"key"`
	DataType string `json:"dataType"`
}

// NewIdentitySupportedFields instantiates a new IdentitySupportedFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentitySupportedFields(iconCodePoint int32, title string, key string, dataType string) *IdentitySupportedFields {
	this := IdentitySupportedFields{}
	this.IconCodePoint = iconCodePoint
	this.Title = title
	this.Key = key
	this.DataType = dataType
	return &this
}

// NewIdentitySupportedFieldsWithDefaults instantiates a new IdentitySupportedFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentitySupportedFieldsWithDefaults() *IdentitySupportedFields {
	this := IdentitySupportedFields{}
	return &this
}

// GetIconCodePoint returns the IconCodePoint field value
func (o *IdentitySupportedFields) GetIconCodePoint() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IconCodePoint
}

// GetIconCodePointOk returns a tuple with the IconCodePoint field value
// and a boolean to check if the value has been set.
func (o *IdentitySupportedFields) GetIconCodePointOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IconCodePoint, true
}

// SetIconCodePoint sets field value
func (o *IdentitySupportedFields) SetIconCodePoint(v int32) {
	o.IconCodePoint = v
}

// GetTitle returns the Title field value
func (o *IdentitySupportedFields) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *IdentitySupportedFields) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *IdentitySupportedFields) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentitySupportedFields) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentitySupportedFields) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *IdentitySupportedFields) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *IdentitySupportedFields) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *IdentitySupportedFields) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *IdentitySupportedFields) UnsetDescription() {
	o.Description.Unset()
}

// GetKey returns the Key field value
func (o *IdentitySupportedFields) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *IdentitySupportedFields) GetKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *IdentitySupportedFields) SetKey(v string) {
	o.Key = v
}

// GetDataType returns the DataType field value
func (o *IdentitySupportedFields) GetDataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value
// and a boolean to check if the value has been set.
func (o *IdentitySupportedFields) GetDataTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DataType, true
}

// SetDataType sets field value
func (o *IdentitySupportedFields) SetDataType(v string) {
	o.DataType = v
}

func (o IdentitySupportedFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["iconCodePoint"] = o.IconCodePoint
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["dataType"] = o.DataType
	}
	return json.Marshal(toSerialize)
}

type NullableIdentitySupportedFields struct {
	value *IdentitySupportedFields
	isSet bool
}

func (v NullableIdentitySupportedFields) Get() *IdentitySupportedFields {
	return v.value
}

func (v *NullableIdentitySupportedFields) Set(val *IdentitySupportedFields) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentitySupportedFields) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentitySupportedFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentitySupportedFields(val *IdentitySupportedFields) *NullableIdentitySupportedFields {
	return &NullableIdentitySupportedFields{value: val, isSet: true}
}

func (v NullableIdentitySupportedFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentitySupportedFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



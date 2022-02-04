/*
My Data My Consent - Developer API

Unleashing the power of data consent by establishing trust. The Platform Core Developer API defines a set of capabilities that can be used to request, issue, manage and update data, documents and credentials by organizations. The API can be used to request, manage and update Decentralised Identifiers, Financial Data, Health Data issue Documents, Credentials directly or using OpenID Connect flows, and verify Messages signed with DIDs and much more.

API version: v1
Contact: support@mydatamyconsent.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github.com/mydatamyconsent/go-sdk

import (
	"encoding/json"
	"time"
)

// SharedWith struct for SharedWith
type SharedWith struct {
	Id *string `json:"id,omitempty"`
	ReceiverId *string `json:"receiverId,omitempty"`
	Name NullableString `json:"name,omitempty"`
	ContactNumber NullableString `json:"contactNumber,omitempty"`
	ProfileUrl NullableString `json:"profileUrl,omitempty"`
	ExpiresAtUtc NullableTime `json:"expiresAtUtc,omitempty"`
}

// NewSharedWith instantiates a new SharedWith object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharedWith() *SharedWith {
	this := SharedWith{}
	return &this
}

// NewSharedWithWithDefaults instantiates a new SharedWith object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharedWithWithDefaults() *SharedWith {
	this := SharedWith{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SharedWith) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedWith) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SharedWith) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SharedWith) SetId(v string) {
	o.Id = &v
}

// GetReceiverId returns the ReceiverId field value if set, zero value otherwise.
func (o *SharedWith) GetReceiverId() string {
	if o == nil || o.ReceiverId == nil {
		var ret string
		return ret
	}
	return *o.ReceiverId
}

// GetReceiverIdOk returns a tuple with the ReceiverId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedWith) GetReceiverIdOk() (*string, bool) {
	if o == nil || o.ReceiverId == nil {
		return nil, false
	}
	return o.ReceiverId, true
}

// HasReceiverId returns a boolean if a field has been set.
func (o *SharedWith) HasReceiverId() bool {
	if o != nil && o.ReceiverId != nil {
		return true
	}

	return false
}

// SetReceiverId gets a reference to the given string and assigns it to the ReceiverId field.
func (o *SharedWith) SetReceiverId(v string) {
	o.ReceiverId = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SharedWith) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SharedWith) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SharedWith) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SharedWith) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SharedWith) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SharedWith) UnsetName() {
	o.Name.Unset()
}

// GetContactNumber returns the ContactNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SharedWith) GetContactNumber() string {
	if o == nil || o.ContactNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContactNumber.Get()
}

// GetContactNumberOk returns a tuple with the ContactNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SharedWith) GetContactNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContactNumber.Get(), o.ContactNumber.IsSet()
}

// HasContactNumber returns a boolean if a field has been set.
func (o *SharedWith) HasContactNumber() bool {
	if o != nil && o.ContactNumber.IsSet() {
		return true
	}

	return false
}

// SetContactNumber gets a reference to the given NullableString and assigns it to the ContactNumber field.
func (o *SharedWith) SetContactNumber(v string) {
	o.ContactNumber.Set(&v)
}
// SetContactNumberNil sets the value for ContactNumber to be an explicit nil
func (o *SharedWith) SetContactNumberNil() {
	o.ContactNumber.Set(nil)
}

// UnsetContactNumber ensures that no value is present for ContactNumber, not even an explicit nil
func (o *SharedWith) UnsetContactNumber() {
	o.ContactNumber.Unset()
}

// GetProfileUrl returns the ProfileUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SharedWith) GetProfileUrl() string {
	if o == nil || o.ProfileUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProfileUrl.Get()
}

// GetProfileUrlOk returns a tuple with the ProfileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SharedWith) GetProfileUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProfileUrl.Get(), o.ProfileUrl.IsSet()
}

// HasProfileUrl returns a boolean if a field has been set.
func (o *SharedWith) HasProfileUrl() bool {
	if o != nil && o.ProfileUrl.IsSet() {
		return true
	}

	return false
}

// SetProfileUrl gets a reference to the given NullableString and assigns it to the ProfileUrl field.
func (o *SharedWith) SetProfileUrl(v string) {
	o.ProfileUrl.Set(&v)
}
// SetProfileUrlNil sets the value for ProfileUrl to be an explicit nil
func (o *SharedWith) SetProfileUrlNil() {
	o.ProfileUrl.Set(nil)
}

// UnsetProfileUrl ensures that no value is present for ProfileUrl, not even an explicit nil
func (o *SharedWith) UnsetProfileUrl() {
	o.ProfileUrl.Unset()
}

// GetExpiresAtUtc returns the ExpiresAtUtc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SharedWith) GetExpiresAtUtc() time.Time {
	if o == nil || o.ExpiresAtUtc.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAtUtc.Get()
}

// GetExpiresAtUtcOk returns a tuple with the ExpiresAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SharedWith) GetExpiresAtUtcOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpiresAtUtc.Get(), o.ExpiresAtUtc.IsSet()
}

// HasExpiresAtUtc returns a boolean if a field has been set.
func (o *SharedWith) HasExpiresAtUtc() bool {
	if o != nil && o.ExpiresAtUtc.IsSet() {
		return true
	}

	return false
}

// SetExpiresAtUtc gets a reference to the given NullableTime and assigns it to the ExpiresAtUtc field.
func (o *SharedWith) SetExpiresAtUtc(v time.Time) {
	o.ExpiresAtUtc.Set(&v)
}
// SetExpiresAtUtcNil sets the value for ExpiresAtUtc to be an explicit nil
func (o *SharedWith) SetExpiresAtUtcNil() {
	o.ExpiresAtUtc.Set(nil)
}

// UnsetExpiresAtUtc ensures that no value is present for ExpiresAtUtc, not even an explicit nil
func (o *SharedWith) UnsetExpiresAtUtc() {
	o.ExpiresAtUtc.Unset()
}

func (o SharedWith) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ReceiverId != nil {
		toSerialize["receiverId"] = o.ReceiverId
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ContactNumber.IsSet() {
		toSerialize["contactNumber"] = o.ContactNumber.Get()
	}
	if o.ProfileUrl.IsSet() {
		toSerialize["profileUrl"] = o.ProfileUrl.Get()
	}
	if o.ExpiresAtUtc.IsSet() {
		toSerialize["expiresAtUtc"] = o.ExpiresAtUtc.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSharedWith struct {
	value *SharedWith
	isSet bool
}

func (v NullableSharedWith) Get() *SharedWith {
	return v.value
}

func (v *NullableSharedWith) Set(val *SharedWith) {
	v.value = val
	v.isSet = true
}

func (v NullableSharedWith) IsSet() bool {
	return v.isSet
}

func (v *NullableSharedWith) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharedWith(val *SharedWith) *NullableSharedWith {
	return &NullableSharedWith{value: val, isSet: true}
}

func (v NullableSharedWith) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharedWith) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



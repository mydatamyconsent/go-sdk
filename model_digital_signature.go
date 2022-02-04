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

// DigitalSignature struct for DigitalSignature
type DigitalSignature struct {
	SignedBy NullableString `json:"signedBy,omitempty"`
	CertIssuedBy NullableString `json:"certIssuedBy,omitempty"`
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	ValidTill *time.Time `json:"validTill,omitempty"`
	Reason NullableString `json:"reason,omitempty"`
	Location NullableString `json:"location,omitempty"`
	Sha1Digest NullableString `json:"sha1Digest,omitempty"`
}

// NewDigitalSignature instantiates a new DigitalSignature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigitalSignature() *DigitalSignature {
	this := DigitalSignature{}
	return &this
}

// NewDigitalSignatureWithDefaults instantiates a new DigitalSignature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigitalSignatureWithDefaults() *DigitalSignature {
	this := DigitalSignature{}
	return &this
}

// GetSignedBy returns the SignedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DigitalSignature) GetSignedBy() string {
	if o == nil || o.SignedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.SignedBy.Get()
}

// GetSignedByOk returns a tuple with the SignedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DigitalSignature) GetSignedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SignedBy.Get(), o.SignedBy.IsSet()
}

// HasSignedBy returns a boolean if a field has been set.
func (o *DigitalSignature) HasSignedBy() bool {
	if o != nil && o.SignedBy.IsSet() {
		return true
	}

	return false
}

// SetSignedBy gets a reference to the given NullableString and assigns it to the SignedBy field.
func (o *DigitalSignature) SetSignedBy(v string) {
	o.SignedBy.Set(&v)
}
// SetSignedByNil sets the value for SignedBy to be an explicit nil
func (o *DigitalSignature) SetSignedByNil() {
	o.SignedBy.Set(nil)
}

// UnsetSignedBy ensures that no value is present for SignedBy, not even an explicit nil
func (o *DigitalSignature) UnsetSignedBy() {
	o.SignedBy.Unset()
}

// GetCertIssuedBy returns the CertIssuedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DigitalSignature) GetCertIssuedBy() string {
	if o == nil || o.CertIssuedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.CertIssuedBy.Get()
}

// GetCertIssuedByOk returns a tuple with the CertIssuedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DigitalSignature) GetCertIssuedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CertIssuedBy.Get(), o.CertIssuedBy.IsSet()
}

// HasCertIssuedBy returns a boolean if a field has been set.
func (o *DigitalSignature) HasCertIssuedBy() bool {
	if o != nil && o.CertIssuedBy.IsSet() {
		return true
	}

	return false
}

// SetCertIssuedBy gets a reference to the given NullableString and assigns it to the CertIssuedBy field.
func (o *DigitalSignature) SetCertIssuedBy(v string) {
	o.CertIssuedBy.Set(&v)
}
// SetCertIssuedByNil sets the value for CertIssuedBy to be an explicit nil
func (o *DigitalSignature) SetCertIssuedByNil() {
	o.CertIssuedBy.Set(nil)
}

// UnsetCertIssuedBy ensures that no value is present for CertIssuedBy, not even an explicit nil
func (o *DigitalSignature) UnsetCertIssuedBy() {
	o.CertIssuedBy.Unset()
}

// GetValidFrom returns the ValidFrom field value if set, zero value otherwise.
func (o *DigitalSignature) GetValidFrom() time.Time {
	if o == nil || o.ValidFrom == nil {
		var ret time.Time
		return ret
	}
	return *o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalSignature) GetValidFromOk() (*time.Time, bool) {
	if o == nil || o.ValidFrom == nil {
		return nil, false
	}
	return o.ValidFrom, true
}

// HasValidFrom returns a boolean if a field has been set.
func (o *DigitalSignature) HasValidFrom() bool {
	if o != nil && o.ValidFrom != nil {
		return true
	}

	return false
}

// SetValidFrom gets a reference to the given time.Time and assigns it to the ValidFrom field.
func (o *DigitalSignature) SetValidFrom(v time.Time) {
	o.ValidFrom = &v
}

// GetValidTill returns the ValidTill field value if set, zero value otherwise.
func (o *DigitalSignature) GetValidTill() time.Time {
	if o == nil || o.ValidTill == nil {
		var ret time.Time
		return ret
	}
	return *o.ValidTill
}

// GetValidTillOk returns a tuple with the ValidTill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalSignature) GetValidTillOk() (*time.Time, bool) {
	if o == nil || o.ValidTill == nil {
		return nil, false
	}
	return o.ValidTill, true
}

// HasValidTill returns a boolean if a field has been set.
func (o *DigitalSignature) HasValidTill() bool {
	if o != nil && o.ValidTill != nil {
		return true
	}

	return false
}

// SetValidTill gets a reference to the given time.Time and assigns it to the ValidTill field.
func (o *DigitalSignature) SetValidTill(v time.Time) {
	o.ValidTill = &v
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DigitalSignature) GetReason() string {
	if o == nil || o.Reason.Get() == nil {
		var ret string
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DigitalSignature) GetReasonOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a field has been set.
func (o *DigitalSignature) HasReason() bool {
	if o != nil && o.Reason.IsSet() {
		return true
	}

	return false
}

// SetReason gets a reference to the given NullableString and assigns it to the Reason field.
func (o *DigitalSignature) SetReason(v string) {
	o.Reason.Set(&v)
}
// SetReasonNil sets the value for Reason to be an explicit nil
func (o *DigitalSignature) SetReasonNil() {
	o.Reason.Set(nil)
}

// UnsetReason ensures that no value is present for Reason, not even an explicit nil
func (o *DigitalSignature) UnsetReason() {
	o.Reason.Unset()
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DigitalSignature) GetLocation() string {
	if o == nil || o.Location.Get() == nil {
		var ret string
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DigitalSignature) GetLocationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *DigitalSignature) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableString and assigns it to the Location field.
func (o *DigitalSignature) SetLocation(v string) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *DigitalSignature) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *DigitalSignature) UnsetLocation() {
	o.Location.Unset()
}

// GetSha1Digest returns the Sha1Digest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DigitalSignature) GetSha1Digest() string {
	if o == nil || o.Sha1Digest.Get() == nil {
		var ret string
		return ret
	}
	return *o.Sha1Digest.Get()
}

// GetSha1DigestOk returns a tuple with the Sha1Digest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DigitalSignature) GetSha1DigestOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Sha1Digest.Get(), o.Sha1Digest.IsSet()
}

// HasSha1Digest returns a boolean if a field has been set.
func (o *DigitalSignature) HasSha1Digest() bool {
	if o != nil && o.Sha1Digest.IsSet() {
		return true
	}

	return false
}

// SetSha1Digest gets a reference to the given NullableString and assigns it to the Sha1Digest field.
func (o *DigitalSignature) SetSha1Digest(v string) {
	o.Sha1Digest.Set(&v)
}
// SetSha1DigestNil sets the value for Sha1Digest to be an explicit nil
func (o *DigitalSignature) SetSha1DigestNil() {
	o.Sha1Digest.Set(nil)
}

// UnsetSha1Digest ensures that no value is present for Sha1Digest, not even an explicit nil
func (o *DigitalSignature) UnsetSha1Digest() {
	o.Sha1Digest.Unset()
}

func (o DigitalSignature) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SignedBy.IsSet() {
		toSerialize["signedBy"] = o.SignedBy.Get()
	}
	if o.CertIssuedBy.IsSet() {
		toSerialize["certIssuedBy"] = o.CertIssuedBy.Get()
	}
	if o.ValidFrom != nil {
		toSerialize["validFrom"] = o.ValidFrom
	}
	if o.ValidTill != nil {
		toSerialize["validTill"] = o.ValidTill
	}
	if o.Reason.IsSet() {
		toSerialize["reason"] = o.Reason.Get()
	}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if o.Sha1Digest.IsSet() {
		toSerialize["sha1Digest"] = o.Sha1Digest.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDigitalSignature struct {
	value *DigitalSignature
	isSet bool
}

func (v NullableDigitalSignature) Get() *DigitalSignature {
	return v.value
}

func (v *NullableDigitalSignature) Set(val *DigitalSignature) {
	v.value = val
	v.isSet = true
}

func (v NullableDigitalSignature) IsSet() bool {
	return v.isSet
}

func (v *NullableDigitalSignature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigitalSignature(val *DigitalSignature) *NullableDigitalSignature {
	return &NullableDigitalSignature{value: val, isSet: true}
}

func (v NullableDigitalSignature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigitalSignature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



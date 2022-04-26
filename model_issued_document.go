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
	"time"
)

// IssuedDocument Issued Document Identifier.
type IssuedDocument struct {
	// Document Id.
	Id string `json:"id"`
	// Document Identifier.
	Identifier string `json:"identifier"`
	// Document type name.
	DocumentType string `json:"documentType"`
	// User name.
	IssuedTo string `json:"issuedTo"`
	// Issued datetime in UTC timezone.
	IssuedAtUtc time.Time `json:"issuedAtUtc"`
	// Expires datetime in UTC timezone.
	ExpiresAtUtc NullableTime `json:"expiresAtUtc,omitempty"`
	// Accepted datetime in UTC timezone.
	AcceptedAtUtc NullableTime `json:"acceptedAtUtc,omitempty"`
}

// NewIssuedDocument instantiates a new IssuedDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssuedDocument(id string, identifier string, documentType string, issuedTo string, issuedAtUtc time.Time) *IssuedDocument {
	this := IssuedDocument{}
	this.Id = id
	this.Identifier = identifier
	this.DocumentType = documentType
	this.IssuedTo = issuedTo
	this.IssuedAtUtc = issuedAtUtc
	return &this
}

// NewIssuedDocumentWithDefaults instantiates a new IssuedDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssuedDocumentWithDefaults() *IssuedDocument {
	this := IssuedDocument{}
	return &this
}

// GetId returns the Id field value
func (o *IssuedDocument) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IssuedDocument) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IssuedDocument) SetId(v string) {
	o.Id = v
}

// GetIdentifier returns the Identifier field value
func (o *IssuedDocument) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *IssuedDocument) GetIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *IssuedDocument) SetIdentifier(v string) {
	o.Identifier = v
}

// GetDocumentType returns the DocumentType field value
func (o *IssuedDocument) GetDocumentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentType
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value
// and a boolean to check if the value has been set.
func (o *IssuedDocument) GetDocumentTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DocumentType, true
}

// SetDocumentType sets field value
func (o *IssuedDocument) SetDocumentType(v string) {
	o.DocumentType = v
}

// GetIssuedTo returns the IssuedTo field value
func (o *IssuedDocument) GetIssuedTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuedTo
}

// GetIssuedToOk returns a tuple with the IssuedTo field value
// and a boolean to check if the value has been set.
func (o *IssuedDocument) GetIssuedToOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IssuedTo, true
}

// SetIssuedTo sets field value
func (o *IssuedDocument) SetIssuedTo(v string) {
	o.IssuedTo = v
}

// GetIssuedAtUtc returns the IssuedAtUtc field value
func (o *IssuedDocument) GetIssuedAtUtc() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.IssuedAtUtc
}

// GetIssuedAtUtcOk returns a tuple with the IssuedAtUtc field value
// and a boolean to check if the value has been set.
func (o *IssuedDocument) GetIssuedAtUtcOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IssuedAtUtc, true
}

// SetIssuedAtUtc sets field value
func (o *IssuedDocument) SetIssuedAtUtc(v time.Time) {
	o.IssuedAtUtc = v
}

// GetExpiresAtUtc returns the ExpiresAtUtc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocument) GetExpiresAtUtc() time.Time {
	if o == nil || o.ExpiresAtUtc.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAtUtc.Get()
}

// GetExpiresAtUtcOk returns a tuple with the ExpiresAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocument) GetExpiresAtUtcOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpiresAtUtc.Get(), o.ExpiresAtUtc.IsSet()
}

// HasExpiresAtUtc returns a boolean if a field has been set.
func (o *IssuedDocument) HasExpiresAtUtc() bool {
	if o != nil && o.ExpiresAtUtc.IsSet() {
		return true
	}

	return false
}

// SetExpiresAtUtc gets a reference to the given NullableTime and assigns it to the ExpiresAtUtc field.
func (o *IssuedDocument) SetExpiresAtUtc(v time.Time) {
	o.ExpiresAtUtc.Set(&v)
}
// SetExpiresAtUtcNil sets the value for ExpiresAtUtc to be an explicit nil
func (o *IssuedDocument) SetExpiresAtUtcNil() {
	o.ExpiresAtUtc.Set(nil)
}

// UnsetExpiresAtUtc ensures that no value is present for ExpiresAtUtc, not even an explicit nil
func (o *IssuedDocument) UnsetExpiresAtUtc() {
	o.ExpiresAtUtc.Unset()
}

// GetAcceptedAtUtc returns the AcceptedAtUtc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuedDocument) GetAcceptedAtUtc() time.Time {
	if o == nil || o.AcceptedAtUtc.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.AcceptedAtUtc.Get()
}

// GetAcceptedAtUtcOk returns a tuple with the AcceptedAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuedDocument) GetAcceptedAtUtcOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AcceptedAtUtc.Get(), o.AcceptedAtUtc.IsSet()
}

// HasAcceptedAtUtc returns a boolean if a field has been set.
func (o *IssuedDocument) HasAcceptedAtUtc() bool {
	if o != nil && o.AcceptedAtUtc.IsSet() {
		return true
	}

	return false
}

// SetAcceptedAtUtc gets a reference to the given NullableTime and assigns it to the AcceptedAtUtc field.
func (o *IssuedDocument) SetAcceptedAtUtc(v time.Time) {
	o.AcceptedAtUtc.Set(&v)
}
// SetAcceptedAtUtcNil sets the value for AcceptedAtUtc to be an explicit nil
func (o *IssuedDocument) SetAcceptedAtUtcNil() {
	o.AcceptedAtUtc.Set(nil)
}

// UnsetAcceptedAtUtc ensures that no value is present for AcceptedAtUtc, not even an explicit nil
func (o *IssuedDocument) UnsetAcceptedAtUtc() {
	o.AcceptedAtUtc.Unset()
}

func (o IssuedDocument) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["identifier"] = o.Identifier
	}
	if true {
		toSerialize["documentType"] = o.DocumentType
	}
	if true {
		toSerialize["issuedTo"] = o.IssuedTo
	}
	if true {
		toSerialize["issuedAtUtc"] = o.IssuedAtUtc
	}
	if o.ExpiresAtUtc.IsSet() {
		toSerialize["expiresAtUtc"] = o.ExpiresAtUtc.Get()
	}
	if o.AcceptedAtUtc.IsSet() {
		toSerialize["acceptedAtUtc"] = o.AcceptedAtUtc.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIssuedDocument struct {
	value *IssuedDocument
	isSet bool
}

func (v NullableIssuedDocument) Get() *IssuedDocument {
	return v.value
}

func (v *NullableIssuedDocument) Set(val *IssuedDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuedDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuedDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuedDocument(val *IssuedDocument) *NullableIssuedDocument {
	return &NullableIssuedDocument{value: val, isSet: true}
}

func (v NullableIssuedDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuedDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



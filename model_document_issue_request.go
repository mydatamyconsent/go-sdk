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

// DocumentIssueRequest Document Issue Request.
type DocumentIssueRequest struct {
	DocumentTypeId string `json:"documentTypeId"`
	DocumentIdentifier string `json:"documentIdentifier"`
	Description string `json:"description"`
	Receiver DocumentReceiver `json:"receiver"`
	ExpiresAtUtc NullableTime `json:"expiresAtUtc,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

// NewDocumentIssueRequest instantiates a new DocumentIssueRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentIssueRequest(documentTypeId string, documentIdentifier string, description string, receiver DocumentReceiver) *DocumentIssueRequest {
	this := DocumentIssueRequest{}
	this.DocumentTypeId = documentTypeId
	this.DocumentIdentifier = documentIdentifier
	this.Description = description
	this.Receiver = receiver
	return &this
}

// NewDocumentIssueRequestWithDefaults instantiates a new DocumentIssueRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentIssueRequestWithDefaults() *DocumentIssueRequest {
	this := DocumentIssueRequest{}
	return &this
}

// GetDocumentTypeId returns the DocumentTypeId field value
func (o *DocumentIssueRequest) GetDocumentTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentTypeId
}

// GetDocumentTypeIdOk returns a tuple with the DocumentTypeId field value
// and a boolean to check if the value has been set.
func (o *DocumentIssueRequest) GetDocumentTypeIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DocumentTypeId, true
}

// SetDocumentTypeId sets field value
func (o *DocumentIssueRequest) SetDocumentTypeId(v string) {
	o.DocumentTypeId = v
}

// GetDocumentIdentifier returns the DocumentIdentifier field value
func (o *DocumentIssueRequest) GetDocumentIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentIdentifier
}

// GetDocumentIdentifierOk returns a tuple with the DocumentIdentifier field value
// and a boolean to check if the value has been set.
func (o *DocumentIssueRequest) GetDocumentIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DocumentIdentifier, true
}

// SetDocumentIdentifier sets field value
func (o *DocumentIssueRequest) SetDocumentIdentifier(v string) {
	o.DocumentIdentifier = v
}

// GetDescription returns the Description field value
func (o *DocumentIssueRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DocumentIssueRequest) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DocumentIssueRequest) SetDescription(v string) {
	o.Description = v
}

// GetReceiver returns the Receiver field value
func (o *DocumentIssueRequest) GetReceiver() DocumentReceiver {
	if o == nil {
		var ret DocumentReceiver
		return ret
	}

	return o.Receiver
}

// GetReceiverOk returns a tuple with the Receiver field value
// and a boolean to check if the value has been set.
func (o *DocumentIssueRequest) GetReceiverOk() (*DocumentReceiver, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Receiver, true
}

// SetReceiver sets field value
func (o *DocumentIssueRequest) SetReceiver(v DocumentReceiver) {
	o.Receiver = v
}

// GetExpiresAtUtc returns the ExpiresAtUtc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentIssueRequest) GetExpiresAtUtc() time.Time {
	if o == nil || o.ExpiresAtUtc.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAtUtc.Get()
}

// GetExpiresAtUtcOk returns a tuple with the ExpiresAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentIssueRequest) GetExpiresAtUtcOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpiresAtUtc.Get(), o.ExpiresAtUtc.IsSet()
}

// HasExpiresAtUtc returns a boolean if a field has been set.
func (o *DocumentIssueRequest) HasExpiresAtUtc() bool {
	if o != nil && o.ExpiresAtUtc.IsSet() {
		return true
	}

	return false
}

// SetExpiresAtUtc gets a reference to the given NullableTime and assigns it to the ExpiresAtUtc field.
func (o *DocumentIssueRequest) SetExpiresAtUtc(v time.Time) {
	o.ExpiresAtUtc.Set(&v)
}
// SetExpiresAtUtcNil sets the value for ExpiresAtUtc to be an explicit nil
func (o *DocumentIssueRequest) SetExpiresAtUtcNil() {
	o.ExpiresAtUtc.Set(nil)
}

// UnsetExpiresAtUtc ensures that no value is present for ExpiresAtUtc, not even an explicit nil
func (o *DocumentIssueRequest) UnsetExpiresAtUtc() {
	o.ExpiresAtUtc.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentIssueRequest) GetMetadata() map[string]string {
	if o == nil  {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentIssueRequest) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *DocumentIssueRequest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *DocumentIssueRequest) SetMetadata(v map[string]string) {
	o.Metadata = v
}

func (o DocumentIssueRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["documentTypeId"] = o.DocumentTypeId
	}
	if true {
		toSerialize["documentIdentifier"] = o.DocumentIdentifier
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["receiver"] = o.Receiver
	}
	if o.ExpiresAtUtc.IsSet() {
		toSerialize["expiresAtUtc"] = o.ExpiresAtUtc.Get()
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableDocumentIssueRequest struct {
	value *DocumentIssueRequest
	isSet bool
}

func (v NullableDocumentIssueRequest) Get() *DocumentIssueRequest {
	return v.value
}

func (v *NullableDocumentIssueRequest) Set(val *DocumentIssueRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentIssueRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentIssueRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentIssueRequest(val *DocumentIssueRequest) *NullableDocumentIssueRequest {
	return &NullableDocumentIssueRequest{value: val, isSet: true}
}

func (v NullableDocumentIssueRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentIssueRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



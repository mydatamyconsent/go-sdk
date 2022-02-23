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

// DocumentIssueRequestDetails Document issue request details.
type DocumentIssueRequestDetails struct {
	DocumentTypeId string `json:"documentTypeId"`
	DocumentTypeName string `json:"documentTypeName"`
	DocumentIdentifier string `json:"documentIdentifier"`
	Description string `json:"description"`
	Receiver interface{} `json:"receiver"`
	ExpiresAtUtc NullableTime `json:"expiresAtUtc,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
	CreatedAtUtc time.Time `json:"createdAtUtc"`
}

// NewDocumentIssueRequestDetails instantiates a new DocumentIssueRequestDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentIssueRequestDetails(documentTypeId string, documentTypeName string, documentIdentifier string, description string, receiver interface{}, createdAtUtc time.Time) *DocumentIssueRequestDetails {
	this := DocumentIssueRequestDetails{}
	this.DocumentTypeId = documentTypeId
	this.DocumentTypeName = documentTypeName
	this.DocumentIdentifier = documentIdentifier
	this.Description = description
	this.Receiver = receiver
	this.CreatedAtUtc = createdAtUtc
	return &this
}

// NewDocumentIssueRequestDetailsWithDefaults instantiates a new DocumentIssueRequestDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentIssueRequestDetailsWithDefaults() *DocumentIssueRequestDetails {
	this := DocumentIssueRequestDetails{}
	return &this
}

// GetDocumentTypeId returns the DocumentTypeId field value
func (o *DocumentIssueRequestDetails) GetDocumentTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentTypeId
}

// GetDocumentTypeIdOk returns a tuple with the DocumentTypeId field value
// and a boolean to check if the value has been set.
func (o *DocumentIssueRequestDetails) GetDocumentTypeIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DocumentTypeId, true
}

// SetDocumentTypeId sets field value
func (o *DocumentIssueRequestDetails) SetDocumentTypeId(v string) {
	o.DocumentTypeId = v
}

// GetDocumentTypeName returns the DocumentTypeName field value
func (o *DocumentIssueRequestDetails) GetDocumentTypeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentTypeName
}

// GetDocumentTypeNameOk returns a tuple with the DocumentTypeName field value
// and a boolean to check if the value has been set.
func (o *DocumentIssueRequestDetails) GetDocumentTypeNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DocumentTypeName, true
}

// SetDocumentTypeName sets field value
func (o *DocumentIssueRequestDetails) SetDocumentTypeName(v string) {
	o.DocumentTypeName = v
}

// GetDocumentIdentifier returns the DocumentIdentifier field value
func (o *DocumentIssueRequestDetails) GetDocumentIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentIdentifier
}

// GetDocumentIdentifierOk returns a tuple with the DocumentIdentifier field value
// and a boolean to check if the value has been set.
func (o *DocumentIssueRequestDetails) GetDocumentIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DocumentIdentifier, true
}

// SetDocumentIdentifier sets field value
func (o *DocumentIssueRequestDetails) SetDocumentIdentifier(v string) {
	o.DocumentIdentifier = v
}

// GetDescription returns the Description field value
func (o *DocumentIssueRequestDetails) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DocumentIssueRequestDetails) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DocumentIssueRequestDetails) SetDescription(v string) {
	o.Description = v
}

// GetReceiver returns the Receiver field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *DocumentIssueRequestDetails) GetReceiver() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Receiver
}

// GetReceiverOk returns a tuple with the Receiver field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentIssueRequestDetails) GetReceiverOk() (*interface{}, bool) {
	if o == nil || o.Receiver == nil {
		return nil, false
	}
	return &o.Receiver, true
}

// SetReceiver sets field value
func (o *DocumentIssueRequestDetails) SetReceiver(v interface{}) {
	o.Receiver = v
}

// GetExpiresAtUtc returns the ExpiresAtUtc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentIssueRequestDetails) GetExpiresAtUtc() time.Time {
	if o == nil || o.ExpiresAtUtc.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAtUtc.Get()
}

// GetExpiresAtUtcOk returns a tuple with the ExpiresAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentIssueRequestDetails) GetExpiresAtUtcOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpiresAtUtc.Get(), o.ExpiresAtUtc.IsSet()
}

// HasExpiresAtUtc returns a boolean if a field has been set.
func (o *DocumentIssueRequestDetails) HasExpiresAtUtc() bool {
	if o != nil && o.ExpiresAtUtc.IsSet() {
		return true
	}

	return false
}

// SetExpiresAtUtc gets a reference to the given NullableTime and assigns it to the ExpiresAtUtc field.
func (o *DocumentIssueRequestDetails) SetExpiresAtUtc(v time.Time) {
	o.ExpiresAtUtc.Set(&v)
}
// SetExpiresAtUtcNil sets the value for ExpiresAtUtc to be an explicit nil
func (o *DocumentIssueRequestDetails) SetExpiresAtUtcNil() {
	o.ExpiresAtUtc.Set(nil)
}

// UnsetExpiresAtUtc ensures that no value is present for ExpiresAtUtc, not even an explicit nil
func (o *DocumentIssueRequestDetails) UnsetExpiresAtUtc() {
	o.ExpiresAtUtc.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentIssueRequestDetails) GetMetadata() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentIssueRequestDetails) GetMetadataOk() (*interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *DocumentIssueRequestDetails) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *DocumentIssueRequestDetails) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetCreatedAtUtc returns the CreatedAtUtc field value
func (o *DocumentIssueRequestDetails) GetCreatedAtUtc() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAtUtc
}

// GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field value
// and a boolean to check if the value has been set.
func (o *DocumentIssueRequestDetails) GetCreatedAtUtcOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAtUtc, true
}

// SetCreatedAtUtc sets field value
func (o *DocumentIssueRequestDetails) SetCreatedAtUtc(v time.Time) {
	o.CreatedAtUtc = v
}

func (o DocumentIssueRequestDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["documentTypeId"] = o.DocumentTypeId
	}
	if true {
		toSerialize["documentTypeName"] = o.DocumentTypeName
	}
	if true {
		toSerialize["documentIdentifier"] = o.DocumentIdentifier
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.Receiver != nil {
		toSerialize["receiver"] = o.Receiver
	}
	if o.ExpiresAtUtc.IsSet() {
		toSerialize["expiresAtUtc"] = o.ExpiresAtUtc.Get()
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["createdAtUtc"] = o.CreatedAtUtc
	}
	return json.Marshal(toSerialize)
}

type NullableDocumentIssueRequestDetails struct {
	value *DocumentIssueRequestDetails
	isSet bool
}

func (v NullableDocumentIssueRequestDetails) Get() *DocumentIssueRequestDetails {
	return v.value
}

func (v *NullableDocumentIssueRequestDetails) Set(val *DocumentIssueRequestDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentIssueRequestDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentIssueRequestDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentIssueRequestDetails(val *DocumentIssueRequestDetails) *NullableDocumentIssueRequestDetails {
	return &NullableDocumentIssueRequestDetails{value: val, isSet: true}
}

func (v NullableDocumentIssueRequestDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentIssueRequestDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



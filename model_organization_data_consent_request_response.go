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

// OrganizationDataConsentRequestResponse Organization Data Consent Request Response.
type OrganizationDataConsentRequestResponse struct {
	Id *string `json:"id,omitempty"`
	TemplateId *string `json:"templateId,omitempty"`
	RequestedAtUtc *time.Time `json:"requestedAtUtc,omitempty"`
	RequestExpiresAtUtc *time.Time `json:"requestExpiresAtUtc,omitempty"`
	Status *DataConsentStatus `json:"status,omitempty"`
	TransactionId NullableString `json:"transactionId,omitempty"`
}

// NewOrganizationDataConsentRequestResponse instantiates a new OrganizationDataConsentRequestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationDataConsentRequestResponse() *OrganizationDataConsentRequestResponse {
	this := OrganizationDataConsentRequestResponse{}
	return &this
}

// NewOrganizationDataConsentRequestResponseWithDefaults instantiates a new OrganizationDataConsentRequestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationDataConsentRequestResponseWithDefaults() *OrganizationDataConsentRequestResponse {
	this := OrganizationDataConsentRequestResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationDataConsentRequestResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDataConsentRequestResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationDataConsentRequestResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationDataConsentRequestResponse) SetId(v string) {
	o.Id = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *OrganizationDataConsentRequestResponse) GetTemplateId() string {
	if o == nil || o.TemplateId == nil {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDataConsentRequestResponse) GetTemplateIdOk() (*string, bool) {
	if o == nil || o.TemplateId == nil {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *OrganizationDataConsentRequestResponse) HasTemplateId() bool {
	if o != nil && o.TemplateId != nil {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *OrganizationDataConsentRequestResponse) SetTemplateId(v string) {
	o.TemplateId = &v
}

// GetRequestedAtUtc returns the RequestedAtUtc field value if set, zero value otherwise.
func (o *OrganizationDataConsentRequestResponse) GetRequestedAtUtc() time.Time {
	if o == nil || o.RequestedAtUtc == nil {
		var ret time.Time
		return ret
	}
	return *o.RequestedAtUtc
}

// GetRequestedAtUtcOk returns a tuple with the RequestedAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDataConsentRequestResponse) GetRequestedAtUtcOk() (*time.Time, bool) {
	if o == nil || o.RequestedAtUtc == nil {
		return nil, false
	}
	return o.RequestedAtUtc, true
}

// HasRequestedAtUtc returns a boolean if a field has been set.
func (o *OrganizationDataConsentRequestResponse) HasRequestedAtUtc() bool {
	if o != nil && o.RequestedAtUtc != nil {
		return true
	}

	return false
}

// SetRequestedAtUtc gets a reference to the given time.Time and assigns it to the RequestedAtUtc field.
func (o *OrganizationDataConsentRequestResponse) SetRequestedAtUtc(v time.Time) {
	o.RequestedAtUtc = &v
}

// GetRequestExpiresAtUtc returns the RequestExpiresAtUtc field value if set, zero value otherwise.
func (o *OrganizationDataConsentRequestResponse) GetRequestExpiresAtUtc() time.Time {
	if o == nil || o.RequestExpiresAtUtc == nil {
		var ret time.Time
		return ret
	}
	return *o.RequestExpiresAtUtc
}

// GetRequestExpiresAtUtcOk returns a tuple with the RequestExpiresAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDataConsentRequestResponse) GetRequestExpiresAtUtcOk() (*time.Time, bool) {
	if o == nil || o.RequestExpiresAtUtc == nil {
		return nil, false
	}
	return o.RequestExpiresAtUtc, true
}

// HasRequestExpiresAtUtc returns a boolean if a field has been set.
func (o *OrganizationDataConsentRequestResponse) HasRequestExpiresAtUtc() bool {
	if o != nil && o.RequestExpiresAtUtc != nil {
		return true
	}

	return false
}

// SetRequestExpiresAtUtc gets a reference to the given time.Time and assigns it to the RequestExpiresAtUtc field.
func (o *OrganizationDataConsentRequestResponse) SetRequestExpiresAtUtc(v time.Time) {
	o.RequestExpiresAtUtc = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrganizationDataConsentRequestResponse) GetStatus() DataConsentStatus {
	if o == nil || o.Status == nil {
		var ret DataConsentStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDataConsentRequestResponse) GetStatusOk() (*DataConsentStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrganizationDataConsentRequestResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given DataConsentStatus and assigns it to the Status field.
func (o *OrganizationDataConsentRequestResponse) SetStatus(v DataConsentStatus) {
	o.Status = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationDataConsentRequestResponse) GetTransactionId() string {
	if o == nil || o.TransactionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TransactionId.Get()
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationDataConsentRequestResponse) GetTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TransactionId.Get(), o.TransactionId.IsSet()
}

// HasTransactionId returns a boolean if a field has been set.
func (o *OrganizationDataConsentRequestResponse) HasTransactionId() bool {
	if o != nil && o.TransactionId.IsSet() {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given NullableString and assigns it to the TransactionId field.
func (o *OrganizationDataConsentRequestResponse) SetTransactionId(v string) {
	o.TransactionId.Set(&v)
}
// SetTransactionIdNil sets the value for TransactionId to be an explicit nil
func (o *OrganizationDataConsentRequestResponse) SetTransactionIdNil() {
	o.TransactionId.Set(nil)
}

// UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
func (o *OrganizationDataConsentRequestResponse) UnsetTransactionId() {
	o.TransactionId.Unset()
}

func (o OrganizationDataConsentRequestResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.TemplateId != nil {
		toSerialize["templateId"] = o.TemplateId
	}
	if o.RequestedAtUtc != nil {
		toSerialize["requestedAtUtc"] = o.RequestedAtUtc
	}
	if o.RequestExpiresAtUtc != nil {
		toSerialize["requestExpiresAtUtc"] = o.RequestExpiresAtUtc
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.TransactionId.IsSet() {
		toSerialize["transactionId"] = o.TransactionId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationDataConsentRequestResponse struct {
	value *OrganizationDataConsentRequestResponse
	isSet bool
}

func (v NullableOrganizationDataConsentRequestResponse) Get() *OrganizationDataConsentRequestResponse {
	return v.value
}

func (v *NullableOrganizationDataConsentRequestResponse) Set(val *OrganizationDataConsentRequestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationDataConsentRequestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationDataConsentRequestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationDataConsentRequestResponse(val *OrganizationDataConsentRequestResponse) *NullableOrganizationDataConsentRequestResponse {
	return &NullableOrganizationDataConsentRequestResponse{value: val, isSet: true}
}

func (v NullableOrganizationDataConsentRequestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationDataConsentRequestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



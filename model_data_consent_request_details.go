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

// DataConsentRequestDetails DataConsentRequestResponse
type DataConsentRequestDetails struct {
	// Consent request id
	Id string `json:"id"`
	// Consent request template id
	TemplateId NullableString `json:"templateId,omitempty"`
	// Data Consent id
	ConsentId NullableString `json:"consentId,omitempty"`
	// Consent request title.
	Title string `json:"title"`
	// Consent request description.
	Description string `json:"description"`
	// Consent request purpose.
	Purpose NullableString `json:"purpose,omitempty"`
	Status DataConsentStatus `json:"status"`
	// Transaction id
	TransactionId NullableString `json:"transactionId,omitempty"`
	// Request creation datetime in UTC timezone
	CreatedAtUtc time.Time `json:"createdAtUtc"`
	// Request expiration datetime in UTC timezone
	ExpiresAtUtc time.Time `json:"expiresAtUtc"`
}

// NewDataConsentRequestDetails instantiates a new DataConsentRequestDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataConsentRequestDetails(id string, title string, description string, status DataConsentStatus, createdAtUtc time.Time, expiresAtUtc time.Time) *DataConsentRequestDetails {
	this := DataConsentRequestDetails{}
	this.Id = id
	this.Title = title
	this.Description = description
	this.Status = status
	this.CreatedAtUtc = createdAtUtc
	this.ExpiresAtUtc = expiresAtUtc
	return &this
}

// NewDataConsentRequestDetailsWithDefaults instantiates a new DataConsentRequestDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataConsentRequestDetailsWithDefaults() *DataConsentRequestDetails {
	this := DataConsentRequestDetails{}
	return &this
}

// GetId returns the Id field value
func (o *DataConsentRequestDetails) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DataConsentRequestDetails) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DataConsentRequestDetails) SetId(v string) {
	o.Id = v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsentRequestDetails) GetTemplateId() string {
	if o == nil || o.TemplateId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TemplateId.Get()
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsentRequestDetails) GetTemplateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TemplateId.Get(), o.TemplateId.IsSet()
}

// HasTemplateId returns a boolean if a field has been set.
func (o *DataConsentRequestDetails) HasTemplateId() bool {
	if o != nil && o.TemplateId.IsSet() {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given NullableString and assigns it to the TemplateId field.
func (o *DataConsentRequestDetails) SetTemplateId(v string) {
	o.TemplateId.Set(&v)
}
// SetTemplateIdNil sets the value for TemplateId to be an explicit nil
func (o *DataConsentRequestDetails) SetTemplateIdNil() {
	o.TemplateId.Set(nil)
}

// UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
func (o *DataConsentRequestDetails) UnsetTemplateId() {
	o.TemplateId.Unset()
}

// GetConsentId returns the ConsentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsentRequestDetails) GetConsentId() string {
	if o == nil || o.ConsentId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConsentId.Get()
}

// GetConsentIdOk returns a tuple with the ConsentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsentRequestDetails) GetConsentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConsentId.Get(), o.ConsentId.IsSet()
}

// HasConsentId returns a boolean if a field has been set.
func (o *DataConsentRequestDetails) HasConsentId() bool {
	if o != nil && o.ConsentId.IsSet() {
		return true
	}

	return false
}

// SetConsentId gets a reference to the given NullableString and assigns it to the ConsentId field.
func (o *DataConsentRequestDetails) SetConsentId(v string) {
	o.ConsentId.Set(&v)
}
// SetConsentIdNil sets the value for ConsentId to be an explicit nil
func (o *DataConsentRequestDetails) SetConsentIdNil() {
	o.ConsentId.Set(nil)
}

// UnsetConsentId ensures that no value is present for ConsentId, not even an explicit nil
func (o *DataConsentRequestDetails) UnsetConsentId() {
	o.ConsentId.Unset()
}

// GetTitle returns the Title field value
func (o *DataConsentRequestDetails) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *DataConsentRequestDetails) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *DataConsentRequestDetails) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value
func (o *DataConsentRequestDetails) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DataConsentRequestDetails) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DataConsentRequestDetails) SetDescription(v string) {
	o.Description = v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsentRequestDetails) GetPurpose() string {
	if o == nil || o.Purpose.Get() == nil {
		var ret string
		return ret
	}
	return *o.Purpose.Get()
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsentRequestDetails) GetPurposeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Purpose.Get(), o.Purpose.IsSet()
}

// HasPurpose returns a boolean if a field has been set.
func (o *DataConsentRequestDetails) HasPurpose() bool {
	if o != nil && o.Purpose.IsSet() {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given NullableString and assigns it to the Purpose field.
func (o *DataConsentRequestDetails) SetPurpose(v string) {
	o.Purpose.Set(&v)
}
// SetPurposeNil sets the value for Purpose to be an explicit nil
func (o *DataConsentRequestDetails) SetPurposeNil() {
	o.Purpose.Set(nil)
}

// UnsetPurpose ensures that no value is present for Purpose, not even an explicit nil
func (o *DataConsentRequestDetails) UnsetPurpose() {
	o.Purpose.Unset()
}

// GetStatus returns the Status field value
func (o *DataConsentRequestDetails) GetStatus() DataConsentStatus {
	if o == nil {
		var ret DataConsentStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DataConsentRequestDetails) GetStatusOk() (*DataConsentStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DataConsentRequestDetails) SetStatus(v DataConsentStatus) {
	o.Status = v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsentRequestDetails) GetTransactionId() string {
	if o == nil || o.TransactionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TransactionId.Get()
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsentRequestDetails) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransactionId.Get(), o.TransactionId.IsSet()
}

// HasTransactionId returns a boolean if a field has been set.
func (o *DataConsentRequestDetails) HasTransactionId() bool {
	if o != nil && o.TransactionId.IsSet() {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given NullableString and assigns it to the TransactionId field.
func (o *DataConsentRequestDetails) SetTransactionId(v string) {
	o.TransactionId.Set(&v)
}
// SetTransactionIdNil sets the value for TransactionId to be an explicit nil
func (o *DataConsentRequestDetails) SetTransactionIdNil() {
	o.TransactionId.Set(nil)
}

// UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
func (o *DataConsentRequestDetails) UnsetTransactionId() {
	o.TransactionId.Unset()
}

// GetCreatedAtUtc returns the CreatedAtUtc field value
func (o *DataConsentRequestDetails) GetCreatedAtUtc() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAtUtc
}

// GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field value
// and a boolean to check if the value has been set.
func (o *DataConsentRequestDetails) GetCreatedAtUtcOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAtUtc, true
}

// SetCreatedAtUtc sets field value
func (o *DataConsentRequestDetails) SetCreatedAtUtc(v time.Time) {
	o.CreatedAtUtc = v
}

// GetExpiresAtUtc returns the ExpiresAtUtc field value
func (o *DataConsentRequestDetails) GetExpiresAtUtc() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAtUtc
}

// GetExpiresAtUtcOk returns a tuple with the ExpiresAtUtc field value
// and a boolean to check if the value has been set.
func (o *DataConsentRequestDetails) GetExpiresAtUtcOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAtUtc, true
}

// SetExpiresAtUtc sets field value
func (o *DataConsentRequestDetails) SetExpiresAtUtc(v time.Time) {
	o.ExpiresAtUtc = v
}

func (o DataConsentRequestDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.TemplateId.IsSet() {
		toSerialize["templateId"] = o.TemplateId.Get()
	}
	if o.ConsentId.IsSet() {
		toSerialize["consentId"] = o.ConsentId.Get()
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.Purpose.IsSet() {
		toSerialize["purpose"] = o.Purpose.Get()
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.TransactionId.IsSet() {
		toSerialize["transactionId"] = o.TransactionId.Get()
	}
	if true {
		toSerialize["createdAtUtc"] = o.CreatedAtUtc
	}
	if true {
		toSerialize["expiresAtUtc"] = o.ExpiresAtUtc
	}
	return json.Marshal(toSerialize)
}

type NullableDataConsentRequestDetails struct {
	value *DataConsentRequestDetails
	isSet bool
}

func (v NullableDataConsentRequestDetails) Get() *DataConsentRequestDetails {
	return v.value
}

func (v *NullableDataConsentRequestDetails) Set(val *DataConsentRequestDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableDataConsentRequestDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableDataConsentRequestDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataConsentRequestDetails(val *DataConsentRequestDetails) *NullableDataConsentRequestDetails {
	return &NullableDataConsentRequestDetails{value: val, isSet: true}
}

func (v NullableDataConsentRequestDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataConsentRequestDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


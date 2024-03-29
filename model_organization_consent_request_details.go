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
	"time"
)

// OrganizationConsentRequestDetails OrganizationConsentRequestDetails : Organization consent request details.
type OrganizationConsentRequestDetails struct {
	// Name of request receiver organization.
	Receiver string `json:"receiver"`
	// Consent request id
	Id string `json:"id"`
	// Consent request template id
	TemplateId *string `json:"templateId,omitempty"`
	// Consent id
	ConsentId *string `json:"consentId,omitempty"`
	// Consent request title.
	Title string `json:"title"`
	// Consent request description.
	Description string `json:"description"`
	// Consent request purpose.
	Purpose *string `json:"purpose,omitempty"`
	Status DataConsentStatus `json:"status"`
	// Transaction id
	TransactionId *string `json:"transactionId,omitempty"`
	// Request creation datetime in UTC timezone
	CreatedAtUtc time.Time `json:"createdAtUtc"`
	// Request expiration datetime in UTC timezone
	ExpiresAtUtc time.Time `json:"expiresAtUtc"`
}

// NewOrganizationConsentRequestDetails instantiates a new OrganizationConsentRequestDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationConsentRequestDetails(receiver string, id string, title string, description string, status DataConsentStatus, createdAtUtc time.Time, expiresAtUtc time.Time) *OrganizationConsentRequestDetails {
	this := OrganizationConsentRequestDetails{}
	this.Receiver = receiver
	this.Id = id
	this.Title = title
	this.Description = description
	this.Status = status
	this.CreatedAtUtc = createdAtUtc
	this.ExpiresAtUtc = expiresAtUtc
	return &this
}

// NewOrganizationConsentRequestDetailsWithDefaults instantiates a new OrganizationConsentRequestDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationConsentRequestDetailsWithDefaults() *OrganizationConsentRequestDetails {
	this := OrganizationConsentRequestDetails{}
	return &this
}

// GetReceiver returns the Receiver field value
func (o *OrganizationConsentRequestDetails) GetReceiver() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Receiver
}

// GetReceiverOk returns a tuple with the Receiver field value
// and a boolean to check if the value has been set.
func (o *OrganizationConsentRequestDetails) GetReceiverOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Receiver, true
}

// SetReceiver sets field value
func (o *OrganizationConsentRequestDetails) SetReceiver(v string) {
	o.Receiver = v
}

// GetId returns the Id field value
func (o *OrganizationConsentRequestDetails) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrganizationConsentRequestDetails) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrganizationConsentRequestDetails) SetId(v string) {
	o.Id = v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *OrganizationConsentRequestDetails) GetTemplateId() string {
	if o == nil || o.TemplateId == nil {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationConsentRequestDetails) GetTemplateIdOk() (*string, bool) {
	if o == nil || o.TemplateId == nil {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *OrganizationConsentRequestDetails) HasTemplateId() bool {
	if o != nil && o.TemplateId != nil {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *OrganizationConsentRequestDetails) SetTemplateId(v string) {
	o.TemplateId = &v
}

// GetConsentId returns the ConsentId field value if set, zero value otherwise.
func (o *OrganizationConsentRequestDetails) GetConsentId() string {
	if o == nil || o.ConsentId == nil {
		var ret string
		return ret
	}
	return *o.ConsentId
}

// GetConsentIdOk returns a tuple with the ConsentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationConsentRequestDetails) GetConsentIdOk() (*string, bool) {
	if o == nil || o.ConsentId == nil {
		return nil, false
	}
	return o.ConsentId, true
}

// HasConsentId returns a boolean if a field has been set.
func (o *OrganizationConsentRequestDetails) HasConsentId() bool {
	if o != nil && o.ConsentId != nil {
		return true
	}

	return false
}

// SetConsentId gets a reference to the given string and assigns it to the ConsentId field.
func (o *OrganizationConsentRequestDetails) SetConsentId(v string) {
	o.ConsentId = &v
}

// GetTitle returns the Title field value
func (o *OrganizationConsentRequestDetails) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *OrganizationConsentRequestDetails) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *OrganizationConsentRequestDetails) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value
func (o *OrganizationConsentRequestDetails) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *OrganizationConsentRequestDetails) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *OrganizationConsentRequestDetails) SetDescription(v string) {
	o.Description = v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *OrganizationConsentRequestDetails) GetPurpose() string {
	if o == nil || o.Purpose == nil {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationConsentRequestDetails) GetPurposeOk() (*string, bool) {
	if o == nil || o.Purpose == nil {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *OrganizationConsentRequestDetails) HasPurpose() bool {
	if o != nil && o.Purpose != nil {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *OrganizationConsentRequestDetails) SetPurpose(v string) {
	o.Purpose = &v
}

// GetStatus returns the Status field value
func (o *OrganizationConsentRequestDetails) GetStatus() DataConsentStatus {
	if o == nil {
		var ret DataConsentStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OrganizationConsentRequestDetails) GetStatusOk() (*DataConsentStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *OrganizationConsentRequestDetails) SetStatus(v DataConsentStatus) {
	o.Status = v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *OrganizationConsentRequestDetails) GetTransactionId() string {
	if o == nil || o.TransactionId == nil {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationConsentRequestDetails) GetTransactionIdOk() (*string, bool) {
	if o == nil || o.TransactionId == nil {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *OrganizationConsentRequestDetails) HasTransactionId() bool {
	if o != nil && o.TransactionId != nil {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *OrganizationConsentRequestDetails) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetCreatedAtUtc returns the CreatedAtUtc field value
func (o *OrganizationConsentRequestDetails) GetCreatedAtUtc() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAtUtc
}

// GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field value
// and a boolean to check if the value has been set.
func (o *OrganizationConsentRequestDetails) GetCreatedAtUtcOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAtUtc, true
}

// SetCreatedAtUtc sets field value
func (o *OrganizationConsentRequestDetails) SetCreatedAtUtc(v time.Time) {
	o.CreatedAtUtc = v
}

// GetExpiresAtUtc returns the ExpiresAtUtc field value
func (o *OrganizationConsentRequestDetails) GetExpiresAtUtc() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAtUtc
}

// GetExpiresAtUtcOk returns a tuple with the ExpiresAtUtc field value
// and a boolean to check if the value has been set.
func (o *OrganizationConsentRequestDetails) GetExpiresAtUtcOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAtUtc, true
}

// SetExpiresAtUtc sets field value
func (o *OrganizationConsentRequestDetails) SetExpiresAtUtc(v time.Time) {
	o.ExpiresAtUtc = v
}

func (o OrganizationConsentRequestDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["receiver"] = o.Receiver
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.TemplateId != nil {
		toSerialize["templateId"] = o.TemplateId
	}
	if o.ConsentId != nil {
		toSerialize["consentId"] = o.ConsentId
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.Purpose != nil {
		toSerialize["purpose"] = o.Purpose
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.TransactionId != nil {
		toSerialize["transactionId"] = o.TransactionId
	}
	if true {
		toSerialize["createdAtUtc"] = o.CreatedAtUtc
	}
	if true {
		toSerialize["expiresAtUtc"] = o.ExpiresAtUtc
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationConsentRequestDetails struct {
	value *OrganizationConsentRequestDetails
	isSet bool
}

func (v NullableOrganizationConsentRequestDetails) Get() *OrganizationConsentRequestDetails {
	return v.value
}

func (v *NullableOrganizationConsentRequestDetails) Set(val *OrganizationConsentRequestDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationConsentRequestDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationConsentRequestDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationConsentRequestDetails(val *OrganizationConsentRequestDetails) *NullableOrganizationConsentRequestDetails {
	return &NullableOrganizationConsentRequestDetails{value: val, isSet: true}
}

func (v NullableOrganizationConsentRequestDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationConsentRequestDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



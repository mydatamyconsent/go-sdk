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

// ConsentRequest ConsentRequest :Consent request details.
type ConsentRequest struct {
	// Consent request id.
	Id string `json:"id"`
	// Consent request template id.
	TemplateId *string `json:"templateId,omitempty"`
	// Consent id.
	ConsentId *string `json:"consentId,omitempty"`
	// Consent title.
	Title string `json:"title"`
	// Consent description.
	Description string `json:"description"`
	// Consent purpose.
	Purpose *string `json:"purpose,omitempty"`
	DataLife *Life `json:"dataLife,omitempty"`
	// List of supported collectables.
	Collectables []CollectibleTypes `json:"collectables"`
	Receiver ConsentRequestReceiver `json:"receiver"`
	Status DataConsentStatus `json:"status"`
	// Request creation datetime in UTC timezone.
	CreatedAtUtc time.Time `json:"createdAtUtc"`
	// Request expiration datetime in UTC timezone.
	ExpiresAtUtc time.Time `json:"expiresAtUtc"`
	// Request approval datetime in UTC timezone.
	ApprovedAtUtc *time.Time `json:"approvedAtUtc,omitempty"`
	// Data access expiration datetime in UTC timezone.
	DataAccessExpiresAtUtc *time.Time `json:"dataAccessExpiresAtUtc,omitempty"`
	// Request rejection datetime in UTC timezone.
	RejectedAtUtc *time.Time `json:"rejectedAtUtc,omitempty"`
	// Request revocation datetime in UTC timezone.
	RevokedAtUtc *time.Time `json:"revokedAtUtc,omitempty"`
}

// NewConsentRequest instantiates a new ConsentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsentRequest(id string, title string, description string, collectables []CollectibleTypes, receiver ConsentRequestReceiver, status DataConsentStatus, createdAtUtc time.Time, expiresAtUtc time.Time) *ConsentRequest {
	this := ConsentRequest{}
	this.Id = id
	this.Title = title
	this.Description = description
	this.Collectables = collectables
	this.Receiver = receiver
	this.Status = status
	this.CreatedAtUtc = createdAtUtc
	this.ExpiresAtUtc = expiresAtUtc
	return &this
}

// NewConsentRequestWithDefaults instantiates a new ConsentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsentRequestWithDefaults() *ConsentRequest {
	this := ConsentRequest{}
	return &this
}

// GetId returns the Id field value
func (o *ConsentRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConsentRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConsentRequest) SetId(v string) {
	o.Id = v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *ConsentRequest) GetTemplateId() string {
	if o == nil || o.TemplateId == nil {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentRequest) GetTemplateIdOk() (*string, bool) {
	if o == nil || o.TemplateId == nil {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *ConsentRequest) HasTemplateId() bool {
	if o != nil && o.TemplateId != nil {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *ConsentRequest) SetTemplateId(v string) {
	o.TemplateId = &v
}

// GetConsentId returns the ConsentId field value if set, zero value otherwise.
func (o *ConsentRequest) GetConsentId() string {
	if o == nil || o.ConsentId == nil {
		var ret string
		return ret
	}
	return *o.ConsentId
}

// GetConsentIdOk returns a tuple with the ConsentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentRequest) GetConsentIdOk() (*string, bool) {
	if o == nil || o.ConsentId == nil {
		return nil, false
	}
	return o.ConsentId, true
}

// HasConsentId returns a boolean if a field has been set.
func (o *ConsentRequest) HasConsentId() bool {
	if o != nil && o.ConsentId != nil {
		return true
	}

	return false
}

// SetConsentId gets a reference to the given string and assigns it to the ConsentId field.
func (o *ConsentRequest) SetConsentId(v string) {
	o.ConsentId = &v
}

// GetTitle returns the Title field value
func (o *ConsentRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ConsentRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ConsentRequest) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value
func (o *ConsentRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ConsentRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ConsentRequest) SetDescription(v string) {
	o.Description = v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *ConsentRequest) GetPurpose() string {
	if o == nil || o.Purpose == nil {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentRequest) GetPurposeOk() (*string, bool) {
	if o == nil || o.Purpose == nil {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *ConsentRequest) HasPurpose() bool {
	if o != nil && o.Purpose != nil {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *ConsentRequest) SetPurpose(v string) {
	o.Purpose = &v
}

// GetDataLife returns the DataLife field value if set, zero value otherwise.
func (o *ConsentRequest) GetDataLife() Life {
	if o == nil || o.DataLife == nil {
		var ret Life
		return ret
	}
	return *o.DataLife
}

// GetDataLifeOk returns a tuple with the DataLife field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentRequest) GetDataLifeOk() (*Life, bool) {
	if o == nil || o.DataLife == nil {
		return nil, false
	}
	return o.DataLife, true
}

// HasDataLife returns a boolean if a field has been set.
func (o *ConsentRequest) HasDataLife() bool {
	if o != nil && o.DataLife != nil {
		return true
	}

	return false
}

// SetDataLife gets a reference to the given Life and assigns it to the DataLife field.
func (o *ConsentRequest) SetDataLife(v Life) {
	o.DataLife = &v
}

// GetCollectables returns the Collectables field value
func (o *ConsentRequest) GetCollectables() []CollectibleTypes {
	if o == nil {
		var ret []CollectibleTypes
		return ret
	}

	return o.Collectables
}

// GetCollectablesOk returns a tuple with the Collectables field value
// and a boolean to check if the value has been set.
func (o *ConsentRequest) GetCollectablesOk() ([]CollectibleTypes, bool) {
	if o == nil {
		return nil, false
	}
	return o.Collectables, true
}

// SetCollectables sets field value
func (o *ConsentRequest) SetCollectables(v []CollectibleTypes) {
	o.Collectables = v
}

// GetReceiver returns the Receiver field value
func (o *ConsentRequest) GetReceiver() ConsentRequestReceiver {
	if o == nil {
		var ret ConsentRequestReceiver
		return ret
	}

	return o.Receiver
}

// GetReceiverOk returns a tuple with the Receiver field value
// and a boolean to check if the value has been set.
func (o *ConsentRequest) GetReceiverOk() (*ConsentRequestReceiver, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Receiver, true
}

// SetReceiver sets field value
func (o *ConsentRequest) SetReceiver(v ConsentRequestReceiver) {
	o.Receiver = v
}

// GetStatus returns the Status field value
func (o *ConsentRequest) GetStatus() DataConsentStatus {
	if o == nil {
		var ret DataConsentStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ConsentRequest) GetStatusOk() (*DataConsentStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ConsentRequest) SetStatus(v DataConsentStatus) {
	o.Status = v
}

// GetCreatedAtUtc returns the CreatedAtUtc field value
func (o *ConsentRequest) GetCreatedAtUtc() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAtUtc
}

// GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field value
// and a boolean to check if the value has been set.
func (o *ConsentRequest) GetCreatedAtUtcOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAtUtc, true
}

// SetCreatedAtUtc sets field value
func (o *ConsentRequest) SetCreatedAtUtc(v time.Time) {
	o.CreatedAtUtc = v
}

// GetExpiresAtUtc returns the ExpiresAtUtc field value
func (o *ConsentRequest) GetExpiresAtUtc() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAtUtc
}

// GetExpiresAtUtcOk returns a tuple with the ExpiresAtUtc field value
// and a boolean to check if the value has been set.
func (o *ConsentRequest) GetExpiresAtUtcOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAtUtc, true
}

// SetExpiresAtUtc sets field value
func (o *ConsentRequest) SetExpiresAtUtc(v time.Time) {
	o.ExpiresAtUtc = v
}

// GetApprovedAtUtc returns the ApprovedAtUtc field value if set, zero value otherwise.
func (o *ConsentRequest) GetApprovedAtUtc() time.Time {
	if o == nil || o.ApprovedAtUtc == nil {
		var ret time.Time
		return ret
	}
	return *o.ApprovedAtUtc
}

// GetApprovedAtUtcOk returns a tuple with the ApprovedAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentRequest) GetApprovedAtUtcOk() (*time.Time, bool) {
	if o == nil || o.ApprovedAtUtc == nil {
		return nil, false
	}
	return o.ApprovedAtUtc, true
}

// HasApprovedAtUtc returns a boolean if a field has been set.
func (o *ConsentRequest) HasApprovedAtUtc() bool {
	if o != nil && o.ApprovedAtUtc != nil {
		return true
	}

	return false
}

// SetApprovedAtUtc gets a reference to the given time.Time and assigns it to the ApprovedAtUtc field.
func (o *ConsentRequest) SetApprovedAtUtc(v time.Time) {
	o.ApprovedAtUtc = &v
}

// GetDataAccessExpiresAtUtc returns the DataAccessExpiresAtUtc field value if set, zero value otherwise.
func (o *ConsentRequest) GetDataAccessExpiresAtUtc() time.Time {
	if o == nil || o.DataAccessExpiresAtUtc == nil {
		var ret time.Time
		return ret
	}
	return *o.DataAccessExpiresAtUtc
}

// GetDataAccessExpiresAtUtcOk returns a tuple with the DataAccessExpiresAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentRequest) GetDataAccessExpiresAtUtcOk() (*time.Time, bool) {
	if o == nil || o.DataAccessExpiresAtUtc == nil {
		return nil, false
	}
	return o.DataAccessExpiresAtUtc, true
}

// HasDataAccessExpiresAtUtc returns a boolean if a field has been set.
func (o *ConsentRequest) HasDataAccessExpiresAtUtc() bool {
	if o != nil && o.DataAccessExpiresAtUtc != nil {
		return true
	}

	return false
}

// SetDataAccessExpiresAtUtc gets a reference to the given time.Time and assigns it to the DataAccessExpiresAtUtc field.
func (o *ConsentRequest) SetDataAccessExpiresAtUtc(v time.Time) {
	o.DataAccessExpiresAtUtc = &v
}

// GetRejectedAtUtc returns the RejectedAtUtc field value if set, zero value otherwise.
func (o *ConsentRequest) GetRejectedAtUtc() time.Time {
	if o == nil || o.RejectedAtUtc == nil {
		var ret time.Time
		return ret
	}
	return *o.RejectedAtUtc
}

// GetRejectedAtUtcOk returns a tuple with the RejectedAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentRequest) GetRejectedAtUtcOk() (*time.Time, bool) {
	if o == nil || o.RejectedAtUtc == nil {
		return nil, false
	}
	return o.RejectedAtUtc, true
}

// HasRejectedAtUtc returns a boolean if a field has been set.
func (o *ConsentRequest) HasRejectedAtUtc() bool {
	if o != nil && o.RejectedAtUtc != nil {
		return true
	}

	return false
}

// SetRejectedAtUtc gets a reference to the given time.Time and assigns it to the RejectedAtUtc field.
func (o *ConsentRequest) SetRejectedAtUtc(v time.Time) {
	o.RejectedAtUtc = &v
}

// GetRevokedAtUtc returns the RevokedAtUtc field value if set, zero value otherwise.
func (o *ConsentRequest) GetRevokedAtUtc() time.Time {
	if o == nil || o.RevokedAtUtc == nil {
		var ret time.Time
		return ret
	}
	return *o.RevokedAtUtc
}

// GetRevokedAtUtcOk returns a tuple with the RevokedAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentRequest) GetRevokedAtUtcOk() (*time.Time, bool) {
	if o == nil || o.RevokedAtUtc == nil {
		return nil, false
	}
	return o.RevokedAtUtc, true
}

// HasRevokedAtUtc returns a boolean if a field has been set.
func (o *ConsentRequest) HasRevokedAtUtc() bool {
	if o != nil && o.RevokedAtUtc != nil {
		return true
	}

	return false
}

// SetRevokedAtUtc gets a reference to the given time.Time and assigns it to the RevokedAtUtc field.
func (o *ConsentRequest) SetRevokedAtUtc(v time.Time) {
	o.RevokedAtUtc = &v
}

func (o ConsentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if o.DataLife != nil {
		toSerialize["dataLife"] = o.DataLife
	}
	if true {
		toSerialize["collectables"] = o.Collectables
	}
	if true {
		toSerialize["receiver"] = o.Receiver
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["createdAtUtc"] = o.CreatedAtUtc
	}
	if true {
		toSerialize["expiresAtUtc"] = o.ExpiresAtUtc
	}
	if o.ApprovedAtUtc != nil {
		toSerialize["approvedAtUtc"] = o.ApprovedAtUtc
	}
	if o.DataAccessExpiresAtUtc != nil {
		toSerialize["dataAccessExpiresAtUtc"] = o.DataAccessExpiresAtUtc
	}
	if o.RejectedAtUtc != nil {
		toSerialize["rejectedAtUtc"] = o.RejectedAtUtc
	}
	if o.RevokedAtUtc != nil {
		toSerialize["revokedAtUtc"] = o.RevokedAtUtc
	}
	return json.Marshal(toSerialize)
}

type NullableConsentRequest struct {
	value *ConsentRequest
	isSet bool
}

func (v NullableConsentRequest) Get() *ConsentRequest {
	return v.value
}

func (v *NullableConsentRequest) Set(val *ConsentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConsentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConsentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsentRequest(val *ConsentRequest) *NullableConsentRequest {
	return &NullableConsentRequest{value: val, isSet: true}
}

func (v NullableConsentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



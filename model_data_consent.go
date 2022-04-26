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

// DataConsent Data Consent details.
type DataConsent struct {
	// Data consent id.
	Id string `json:"id"`
	// Data consent request id.
	RequestId string `json:"requestId"`
	// Consent template id.
	TemplateId NullableString `json:"templateId,omitempty"`
	// Consent title.
	Title string `json:"title"`
	// Consent description.
	Description string `json:"description"`
	// Consent purpose.
	Purpose NullableString `json:"purpose,omitempty"`
	Status DataConsentStatus `json:"status"`
	// Transaction id.
	TransactionId NullableString `json:"transactionId,omitempty"`
	// Consent approval datetime in UTC timezone.
	ApprovedAtUtc time.Time `json:"approvedAtUtc"`
	// Data access expiration datetime in UTC timezone.
	DataAccessExpiresAtUtc time.Time `json:"dataAccessExpiresAtUtc"`
	// Consent revocation datetime in UTC timezone.
	RevokedAtUtc NullableTime `json:"revokedAtUtc,omitempty"`
	// List of supported collectible types.
	Collectables []CollectibleTypes `json:"collectables"`
	// Consented identity details.
	Identifiers interface{} `json:"identifiers,omitempty"`
	// List of consented documents.
	Documents []DataConsentDocument `json:"documents,omitempty"`
}

// NewDataConsent instantiates a new DataConsent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataConsent(id string, requestId string, title string, description string, status DataConsentStatus, approvedAtUtc time.Time, dataAccessExpiresAtUtc time.Time, collectables []CollectibleTypes) *DataConsent {
	this := DataConsent{}
	this.Id = id
	this.RequestId = requestId
	this.Title = title
	this.Description = description
	this.Status = status
	this.ApprovedAtUtc = approvedAtUtc
	this.DataAccessExpiresAtUtc = dataAccessExpiresAtUtc
	this.Collectables = collectables
	return &this
}

// NewDataConsentWithDefaults instantiates a new DataConsent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataConsentWithDefaults() *DataConsent {
	this := DataConsent{}
	return &this
}

// GetId returns the Id field value
func (o *DataConsent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DataConsent) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DataConsent) SetId(v string) {
	o.Id = v
}

// GetRequestId returns the RequestId field value
func (o *DataConsent) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *DataConsent) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *DataConsent) SetRequestId(v string) {
	o.RequestId = v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsent) GetTemplateId() string {
	if o == nil || o.TemplateId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TemplateId.Get()
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsent) GetTemplateIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TemplateId.Get(), o.TemplateId.IsSet()
}

// HasTemplateId returns a boolean if a field has been set.
func (o *DataConsent) HasTemplateId() bool {
	if o != nil && o.TemplateId.IsSet() {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given NullableString and assigns it to the TemplateId field.
func (o *DataConsent) SetTemplateId(v string) {
	o.TemplateId.Set(&v)
}
// SetTemplateIdNil sets the value for TemplateId to be an explicit nil
func (o *DataConsent) SetTemplateIdNil() {
	o.TemplateId.Set(nil)
}

// UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
func (o *DataConsent) UnsetTemplateId() {
	o.TemplateId.Unset()
}

// GetTitle returns the Title field value
func (o *DataConsent) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *DataConsent) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *DataConsent) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value
func (o *DataConsent) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DataConsent) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DataConsent) SetDescription(v string) {
	o.Description = v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsent) GetPurpose() string {
	if o == nil || o.Purpose.Get() == nil {
		var ret string
		return ret
	}
	return *o.Purpose.Get()
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsent) GetPurposeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Purpose.Get(), o.Purpose.IsSet()
}

// HasPurpose returns a boolean if a field has been set.
func (o *DataConsent) HasPurpose() bool {
	if o != nil && o.Purpose.IsSet() {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given NullableString and assigns it to the Purpose field.
func (o *DataConsent) SetPurpose(v string) {
	o.Purpose.Set(&v)
}
// SetPurposeNil sets the value for Purpose to be an explicit nil
func (o *DataConsent) SetPurposeNil() {
	o.Purpose.Set(nil)
}

// UnsetPurpose ensures that no value is present for Purpose, not even an explicit nil
func (o *DataConsent) UnsetPurpose() {
	o.Purpose.Unset()
}

// GetStatus returns the Status field value
func (o *DataConsent) GetStatus() DataConsentStatus {
	if o == nil {
		var ret DataConsentStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DataConsent) GetStatusOk() (*DataConsentStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DataConsent) SetStatus(v DataConsentStatus) {
	o.Status = v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsent) GetTransactionId() string {
	if o == nil || o.TransactionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TransactionId.Get()
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsent) GetTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TransactionId.Get(), o.TransactionId.IsSet()
}

// HasTransactionId returns a boolean if a field has been set.
func (o *DataConsent) HasTransactionId() bool {
	if o != nil && o.TransactionId.IsSet() {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given NullableString and assigns it to the TransactionId field.
func (o *DataConsent) SetTransactionId(v string) {
	o.TransactionId.Set(&v)
}
// SetTransactionIdNil sets the value for TransactionId to be an explicit nil
func (o *DataConsent) SetTransactionIdNil() {
	o.TransactionId.Set(nil)
}

// UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
func (o *DataConsent) UnsetTransactionId() {
	o.TransactionId.Unset()
}

// GetApprovedAtUtc returns the ApprovedAtUtc field value
func (o *DataConsent) GetApprovedAtUtc() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ApprovedAtUtc
}

// GetApprovedAtUtcOk returns a tuple with the ApprovedAtUtc field value
// and a boolean to check if the value has been set.
func (o *DataConsent) GetApprovedAtUtcOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApprovedAtUtc, true
}

// SetApprovedAtUtc sets field value
func (o *DataConsent) SetApprovedAtUtc(v time.Time) {
	o.ApprovedAtUtc = v
}

// GetDataAccessExpiresAtUtc returns the DataAccessExpiresAtUtc field value
func (o *DataConsent) GetDataAccessExpiresAtUtc() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DataAccessExpiresAtUtc
}

// GetDataAccessExpiresAtUtcOk returns a tuple with the DataAccessExpiresAtUtc field value
// and a boolean to check if the value has been set.
func (o *DataConsent) GetDataAccessExpiresAtUtcOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DataAccessExpiresAtUtc, true
}

// SetDataAccessExpiresAtUtc sets field value
func (o *DataConsent) SetDataAccessExpiresAtUtc(v time.Time) {
	o.DataAccessExpiresAtUtc = v
}

// GetRevokedAtUtc returns the RevokedAtUtc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsent) GetRevokedAtUtc() time.Time {
	if o == nil || o.RevokedAtUtc.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.RevokedAtUtc.Get()
}

// GetRevokedAtUtcOk returns a tuple with the RevokedAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsent) GetRevokedAtUtcOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RevokedAtUtc.Get(), o.RevokedAtUtc.IsSet()
}

// HasRevokedAtUtc returns a boolean if a field has been set.
func (o *DataConsent) HasRevokedAtUtc() bool {
	if o != nil && o.RevokedAtUtc.IsSet() {
		return true
	}

	return false
}

// SetRevokedAtUtc gets a reference to the given NullableTime and assigns it to the RevokedAtUtc field.
func (o *DataConsent) SetRevokedAtUtc(v time.Time) {
	o.RevokedAtUtc.Set(&v)
}
// SetRevokedAtUtcNil sets the value for RevokedAtUtc to be an explicit nil
func (o *DataConsent) SetRevokedAtUtcNil() {
	o.RevokedAtUtc.Set(nil)
}

// UnsetRevokedAtUtc ensures that no value is present for RevokedAtUtc, not even an explicit nil
func (o *DataConsent) UnsetRevokedAtUtc() {
	o.RevokedAtUtc.Unset()
}

// GetCollectables returns the Collectables field value
func (o *DataConsent) GetCollectables() []CollectibleTypes {
	if o == nil {
		var ret []CollectibleTypes
		return ret
	}

	return o.Collectables
}

// GetCollectablesOk returns a tuple with the Collectables field value
// and a boolean to check if the value has been set.
func (o *DataConsent) GetCollectablesOk() ([]CollectibleTypes, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Collectables, true
}

// SetCollectables sets field value
func (o *DataConsent) SetCollectables(v []CollectibleTypes) {
	o.Collectables = v
}

// GetIdentifiers returns the Identifiers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsent) GetIdentifiers() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Identifiers
}

// GetIdentifiersOk returns a tuple with the Identifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsent) GetIdentifiersOk() (*interface{}, bool) {
	if o == nil || o.Identifiers == nil {
		return nil, false
	}
	return &o.Identifiers, true
}

// HasIdentifiers returns a boolean if a field has been set.
func (o *DataConsent) HasIdentifiers() bool {
	if o != nil && o.Identifiers != nil {
		return true
	}

	return false
}

// SetIdentifiers gets a reference to the given interface{} and assigns it to the Identifiers field.
func (o *DataConsent) SetIdentifiers(v interface{}) {
	o.Identifiers = v
}

// GetDocuments returns the Documents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsent) GetDocuments() []DataConsentDocument {
	if o == nil  {
		var ret []DataConsentDocument
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsent) GetDocumentsOk() ([]DataConsentDocument, bool) {
	if o == nil || o.Documents == nil {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *DataConsent) HasDocuments() bool {
	if o != nil && o.Documents != nil {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []DataConsentDocument and assigns it to the Documents field.
func (o *DataConsent) SetDocuments(v []DataConsentDocument) {
	o.Documents = v
}

func (o DataConsent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["requestId"] = o.RequestId
	}
	if o.TemplateId.IsSet() {
		toSerialize["templateId"] = o.TemplateId.Get()
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
		toSerialize["approvedAtUtc"] = o.ApprovedAtUtc
	}
	if true {
		toSerialize["dataAccessExpiresAtUtc"] = o.DataAccessExpiresAtUtc
	}
	if o.RevokedAtUtc.IsSet() {
		toSerialize["revokedAtUtc"] = o.RevokedAtUtc.Get()
	}
	if true {
		toSerialize["collectables"] = o.Collectables
	}
	if o.Identifiers != nil {
		toSerialize["identifiers"] = o.Identifiers
	}
	if o.Documents != nil {
		toSerialize["documents"] = o.Documents
	}
	return json.Marshal(toSerialize)
}

type NullableDataConsent struct {
	value *DataConsent
	isSet bool
}

func (v NullableDataConsent) Get() *DataConsent {
	return v.value
}

func (v *NullableDataConsent) Set(val *DataConsent) {
	v.value = val
	v.isSet = true
}

func (v NullableDataConsent) IsSet() bool {
	return v.isSet
}

func (v *NullableDataConsent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataConsent(val *DataConsent) *NullableDataConsent {
	return &NullableDataConsent{value: val, isSet: true}
}

func (v NullableDataConsent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataConsent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// DataConsentDetailsDto struct for DataConsentDetailsDto
type DataConsentDetailsDto struct {
	ConsentRequestId string `json:"consentRequestId"`
	Title NullableString `json:"title,omitempty"`
	Description NullableString `json:"description,omitempty"`
	DataLife *Life `json:"dataLife,omitempty"`
	RequestedByOrg *Requester `json:"requestedByOrg,omitempty"`
	Status *DataConsentStatus `json:"status,omitempty"`
	ApprovedAtUtc NullableTime `json:"approvedAtUtc,omitempty"`
	ApprovedExpiresAtUtc NullableTime `json:"approvedExpiresAtUtc,omitempty"`
	RejectedAtUtc NullableTime `json:"rejectedAtUtc,omitempty"`
	RevokedAtUtc NullableTime `json:"revokedAtUtc,omitempty"`
	RequestedExpiresAtUtc *time.Time `json:"requestedExpiresAtUtc,omitempty"`
	RequestedAtUtc *time.Time `json:"requestedAtUtc,omitempty"`
	Identifiers interface{} `json:"identifiers,omitempty"`
	Documents []DataConsentDocumentDetailsDto `json:"documents,omitempty"`
}

// NewDataConsentDetailsDto instantiates a new DataConsentDetailsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataConsentDetailsDto(consentRequestId string) *DataConsentDetailsDto {
	this := DataConsentDetailsDto{}
	this.ConsentRequestId = consentRequestId
	return &this
}

// NewDataConsentDetailsDtoWithDefaults instantiates a new DataConsentDetailsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataConsentDetailsDtoWithDefaults() *DataConsentDetailsDto {
	this := DataConsentDetailsDto{}
	return &this
}

// GetConsentRequestId returns the ConsentRequestId field value
func (o *DataConsentDetailsDto) GetConsentRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsentRequestId
}

// GetConsentRequestIdOk returns a tuple with the ConsentRequestId field value
// and a boolean to check if the value has been set.
func (o *DataConsentDetailsDto) GetConsentRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConsentRequestId, true
}

// SetConsentRequestId sets field value
func (o *DataConsentDetailsDto) SetConsentRequestId(v string) {
	o.ConsentRequestId = v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsentDetailsDto) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsentDetailsDto) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *DataConsentDetailsDto) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *DataConsentDetailsDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *DataConsentDetailsDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *DataConsentDetailsDto) UnsetTitle() {
	o.Title.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsentDetailsDto) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsentDetailsDto) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *DataConsentDetailsDto) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *DataConsentDetailsDto) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *DataConsentDetailsDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *DataConsentDetailsDto) UnsetDescription() {
	o.Description.Unset()
}

// GetDataLife returns the DataLife field value if set, zero value otherwise.
func (o *DataConsentDetailsDto) GetDataLife() Life {
	if o == nil || o.DataLife == nil {
		var ret Life
		return ret
	}
	return *o.DataLife
}

// GetDataLifeOk returns a tuple with the DataLife field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataConsentDetailsDto) GetDataLifeOk() (*Life, bool) {
	if o == nil || o.DataLife == nil {
		return nil, false
	}
	return o.DataLife, true
}

// HasDataLife returns a boolean if a field has been set.
func (o *DataConsentDetailsDto) HasDataLife() bool {
	if o != nil && o.DataLife != nil {
		return true
	}

	return false
}

// SetDataLife gets a reference to the given Life and assigns it to the DataLife field.
func (o *DataConsentDetailsDto) SetDataLife(v Life) {
	o.DataLife = &v
}

// GetRequestedByOrg returns the RequestedByOrg field value if set, zero value otherwise.
func (o *DataConsentDetailsDto) GetRequestedByOrg() Requester {
	if o == nil || o.RequestedByOrg == nil {
		var ret Requester
		return ret
	}
	return *o.RequestedByOrg
}

// GetRequestedByOrgOk returns a tuple with the RequestedByOrg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataConsentDetailsDto) GetRequestedByOrgOk() (*Requester, bool) {
	if o == nil || o.RequestedByOrg == nil {
		return nil, false
	}
	return o.RequestedByOrg, true
}

// HasRequestedByOrg returns a boolean if a field has been set.
func (o *DataConsentDetailsDto) HasRequestedByOrg() bool {
	if o != nil && o.RequestedByOrg != nil {
		return true
	}

	return false
}

// SetRequestedByOrg gets a reference to the given Requester and assigns it to the RequestedByOrg field.
func (o *DataConsentDetailsDto) SetRequestedByOrg(v Requester) {
	o.RequestedByOrg = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DataConsentDetailsDto) GetStatus() DataConsentStatus {
	if o == nil || o.Status == nil {
		var ret DataConsentStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataConsentDetailsDto) GetStatusOk() (*DataConsentStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DataConsentDetailsDto) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given DataConsentStatus and assigns it to the Status field.
func (o *DataConsentDetailsDto) SetStatus(v DataConsentStatus) {
	o.Status = &v
}

// GetApprovedAtUtc returns the ApprovedAtUtc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsentDetailsDto) GetApprovedAtUtc() time.Time {
	if o == nil || o.ApprovedAtUtc.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ApprovedAtUtc.Get()
}

// GetApprovedAtUtcOk returns a tuple with the ApprovedAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsentDetailsDto) GetApprovedAtUtcOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ApprovedAtUtc.Get(), o.ApprovedAtUtc.IsSet()
}

// HasApprovedAtUtc returns a boolean if a field has been set.
func (o *DataConsentDetailsDto) HasApprovedAtUtc() bool {
	if o != nil && o.ApprovedAtUtc.IsSet() {
		return true
	}

	return false
}

// SetApprovedAtUtc gets a reference to the given NullableTime and assigns it to the ApprovedAtUtc field.
func (o *DataConsentDetailsDto) SetApprovedAtUtc(v time.Time) {
	o.ApprovedAtUtc.Set(&v)
}
// SetApprovedAtUtcNil sets the value for ApprovedAtUtc to be an explicit nil
func (o *DataConsentDetailsDto) SetApprovedAtUtcNil() {
	o.ApprovedAtUtc.Set(nil)
}

// UnsetApprovedAtUtc ensures that no value is present for ApprovedAtUtc, not even an explicit nil
func (o *DataConsentDetailsDto) UnsetApprovedAtUtc() {
	o.ApprovedAtUtc.Unset()
}

// GetApprovedExpiresAtUtc returns the ApprovedExpiresAtUtc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsentDetailsDto) GetApprovedExpiresAtUtc() time.Time {
	if o == nil || o.ApprovedExpiresAtUtc.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ApprovedExpiresAtUtc.Get()
}

// GetApprovedExpiresAtUtcOk returns a tuple with the ApprovedExpiresAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsentDetailsDto) GetApprovedExpiresAtUtcOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ApprovedExpiresAtUtc.Get(), o.ApprovedExpiresAtUtc.IsSet()
}

// HasApprovedExpiresAtUtc returns a boolean if a field has been set.
func (o *DataConsentDetailsDto) HasApprovedExpiresAtUtc() bool {
	if o != nil && o.ApprovedExpiresAtUtc.IsSet() {
		return true
	}

	return false
}

// SetApprovedExpiresAtUtc gets a reference to the given NullableTime and assigns it to the ApprovedExpiresAtUtc field.
func (o *DataConsentDetailsDto) SetApprovedExpiresAtUtc(v time.Time) {
	o.ApprovedExpiresAtUtc.Set(&v)
}
// SetApprovedExpiresAtUtcNil sets the value for ApprovedExpiresAtUtc to be an explicit nil
func (o *DataConsentDetailsDto) SetApprovedExpiresAtUtcNil() {
	o.ApprovedExpiresAtUtc.Set(nil)
}

// UnsetApprovedExpiresAtUtc ensures that no value is present for ApprovedExpiresAtUtc, not even an explicit nil
func (o *DataConsentDetailsDto) UnsetApprovedExpiresAtUtc() {
	o.ApprovedExpiresAtUtc.Unset()
}

// GetRejectedAtUtc returns the RejectedAtUtc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsentDetailsDto) GetRejectedAtUtc() time.Time {
	if o == nil || o.RejectedAtUtc.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.RejectedAtUtc.Get()
}

// GetRejectedAtUtcOk returns a tuple with the RejectedAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsentDetailsDto) GetRejectedAtUtcOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RejectedAtUtc.Get(), o.RejectedAtUtc.IsSet()
}

// HasRejectedAtUtc returns a boolean if a field has been set.
func (o *DataConsentDetailsDto) HasRejectedAtUtc() bool {
	if o != nil && o.RejectedAtUtc.IsSet() {
		return true
	}

	return false
}

// SetRejectedAtUtc gets a reference to the given NullableTime and assigns it to the RejectedAtUtc field.
func (o *DataConsentDetailsDto) SetRejectedAtUtc(v time.Time) {
	o.RejectedAtUtc.Set(&v)
}
// SetRejectedAtUtcNil sets the value for RejectedAtUtc to be an explicit nil
func (o *DataConsentDetailsDto) SetRejectedAtUtcNil() {
	o.RejectedAtUtc.Set(nil)
}

// UnsetRejectedAtUtc ensures that no value is present for RejectedAtUtc, not even an explicit nil
func (o *DataConsentDetailsDto) UnsetRejectedAtUtc() {
	o.RejectedAtUtc.Unset()
}

// GetRevokedAtUtc returns the RevokedAtUtc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsentDetailsDto) GetRevokedAtUtc() time.Time {
	if o == nil || o.RevokedAtUtc.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.RevokedAtUtc.Get()
}

// GetRevokedAtUtcOk returns a tuple with the RevokedAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsentDetailsDto) GetRevokedAtUtcOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RevokedAtUtc.Get(), o.RevokedAtUtc.IsSet()
}

// HasRevokedAtUtc returns a boolean if a field has been set.
func (o *DataConsentDetailsDto) HasRevokedAtUtc() bool {
	if o != nil && o.RevokedAtUtc.IsSet() {
		return true
	}

	return false
}

// SetRevokedAtUtc gets a reference to the given NullableTime and assigns it to the RevokedAtUtc field.
func (o *DataConsentDetailsDto) SetRevokedAtUtc(v time.Time) {
	o.RevokedAtUtc.Set(&v)
}
// SetRevokedAtUtcNil sets the value for RevokedAtUtc to be an explicit nil
func (o *DataConsentDetailsDto) SetRevokedAtUtcNil() {
	o.RevokedAtUtc.Set(nil)
}

// UnsetRevokedAtUtc ensures that no value is present for RevokedAtUtc, not even an explicit nil
func (o *DataConsentDetailsDto) UnsetRevokedAtUtc() {
	o.RevokedAtUtc.Unset()
}

// GetRequestedExpiresAtUtc returns the RequestedExpiresAtUtc field value if set, zero value otherwise.
func (o *DataConsentDetailsDto) GetRequestedExpiresAtUtc() time.Time {
	if o == nil || o.RequestedExpiresAtUtc == nil {
		var ret time.Time
		return ret
	}
	return *o.RequestedExpiresAtUtc
}

// GetRequestedExpiresAtUtcOk returns a tuple with the RequestedExpiresAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataConsentDetailsDto) GetRequestedExpiresAtUtcOk() (*time.Time, bool) {
	if o == nil || o.RequestedExpiresAtUtc == nil {
		return nil, false
	}
	return o.RequestedExpiresAtUtc, true
}

// HasRequestedExpiresAtUtc returns a boolean if a field has been set.
func (o *DataConsentDetailsDto) HasRequestedExpiresAtUtc() bool {
	if o != nil && o.RequestedExpiresAtUtc != nil {
		return true
	}

	return false
}

// SetRequestedExpiresAtUtc gets a reference to the given time.Time and assigns it to the RequestedExpiresAtUtc field.
func (o *DataConsentDetailsDto) SetRequestedExpiresAtUtc(v time.Time) {
	o.RequestedExpiresAtUtc = &v
}

// GetRequestedAtUtc returns the RequestedAtUtc field value if set, zero value otherwise.
func (o *DataConsentDetailsDto) GetRequestedAtUtc() time.Time {
	if o == nil || o.RequestedAtUtc == nil {
		var ret time.Time
		return ret
	}
	return *o.RequestedAtUtc
}

// GetRequestedAtUtcOk returns a tuple with the RequestedAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataConsentDetailsDto) GetRequestedAtUtcOk() (*time.Time, bool) {
	if o == nil || o.RequestedAtUtc == nil {
		return nil, false
	}
	return o.RequestedAtUtc, true
}

// HasRequestedAtUtc returns a boolean if a field has been set.
func (o *DataConsentDetailsDto) HasRequestedAtUtc() bool {
	if o != nil && o.RequestedAtUtc != nil {
		return true
	}

	return false
}

// SetRequestedAtUtc gets a reference to the given time.Time and assigns it to the RequestedAtUtc field.
func (o *DataConsentDetailsDto) SetRequestedAtUtc(v time.Time) {
	o.RequestedAtUtc = &v
}

// GetIdentifiers returns the Identifiers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsentDetailsDto) GetIdentifiers() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Identifiers
}

// GetIdentifiersOk returns a tuple with the Identifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsentDetailsDto) GetIdentifiersOk() (*interface{}, bool) {
	if o == nil || o.Identifiers == nil {
		return nil, false
	}
	return &o.Identifiers, true
}

// HasIdentifiers returns a boolean if a field has been set.
func (o *DataConsentDetailsDto) HasIdentifiers() bool {
	if o != nil && o.Identifiers != nil {
		return true
	}

	return false
}

// SetIdentifiers gets a reference to the given interface{} and assigns it to the Identifiers field.
func (o *DataConsentDetailsDto) SetIdentifiers(v interface{}) {
	o.Identifiers = v
}

// GetDocuments returns the Documents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsentDetailsDto) GetDocuments() []DataConsentDocumentDetailsDto {
	if o == nil  {
		var ret []DataConsentDocumentDetailsDto
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsentDetailsDto) GetDocumentsOk() ([]DataConsentDocumentDetailsDto, bool) {
	if o == nil || o.Documents == nil {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *DataConsentDetailsDto) HasDocuments() bool {
	if o != nil && o.Documents != nil {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []DataConsentDocumentDetailsDto and assigns it to the Documents field.
func (o *DataConsentDetailsDto) SetDocuments(v []DataConsentDocumentDetailsDto) {
	o.Documents = v
}

func (o DataConsentDetailsDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["consentRequestId"] = o.ConsentRequestId
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DataLife != nil {
		toSerialize["dataLife"] = o.DataLife
	}
	if o.RequestedByOrg != nil {
		toSerialize["requestedByOrg"] = o.RequestedByOrg
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ApprovedAtUtc.IsSet() {
		toSerialize["approvedAtUtc"] = o.ApprovedAtUtc.Get()
	}
	if o.ApprovedExpiresAtUtc.IsSet() {
		toSerialize["approvedExpiresAtUtc"] = o.ApprovedExpiresAtUtc.Get()
	}
	if o.RejectedAtUtc.IsSet() {
		toSerialize["rejectedAtUtc"] = o.RejectedAtUtc.Get()
	}
	if o.RevokedAtUtc.IsSet() {
		toSerialize["revokedAtUtc"] = o.RevokedAtUtc.Get()
	}
	if o.RequestedExpiresAtUtc != nil {
		toSerialize["requestedExpiresAtUtc"] = o.RequestedExpiresAtUtc
	}
	if o.RequestedAtUtc != nil {
		toSerialize["requestedAtUtc"] = o.RequestedAtUtc
	}
	if o.Identifiers != nil {
		toSerialize["identifiers"] = o.Identifiers
	}
	if o.Documents != nil {
		toSerialize["documents"] = o.Documents
	}
	return json.Marshal(toSerialize)
}

type NullableDataConsentDetailsDto struct {
	value *DataConsentDetailsDto
	isSet bool
}

func (v NullableDataConsentDetailsDto) Get() *DataConsentDetailsDto {
	return v.value
}

func (v *NullableDataConsentDetailsDto) Set(val *DataConsentDetailsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDataConsentDetailsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDataConsentDetailsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataConsentDetailsDto(val *DataConsentDetailsDto) *NullableDataConsentDetailsDto {
	return &NullableDataConsentDetailsDto{value: val, isSet: true}
}

func (v NullableDataConsentDetailsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataConsentDetailsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



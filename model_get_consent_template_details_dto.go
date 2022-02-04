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

// GetConsentTemplateDetailsDto struct for GetConsentTemplateDetailsDto
type GetConsentTemplateDetailsDto struct {
	Id *string `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
	ConsentPurpose NullableString `json:"consentPurpose,omitempty"`
	Collectables []CollectibleTypes `json:"collectables,omitempty"`
	FetchType *FetchTypes `json:"fetchType,omitempty"`
	ShortId NullableString `json:"shortId,omitempty"`
	CreatedBy NullableString `json:"createdBy,omitempty"`
	CreatedAtUtc *time.Time `json:"createdAtUtc,omitempty"`
	Status NullableString `json:"status,omitempty"`
	TemplateType *ConsentTemplateTypes `json:"templateType,omitempty"`
	Frequency *Life `json:"frequency,omitempty"`
	Identity []IdentitySupportedFields `json:"identity,omitempty"`
	Documents []Document `json:"documents,omitempty"`
	Financials []Financial `json:"financials,omitempty"`
}

// NewGetConsentTemplateDetailsDto instantiates a new GetConsentTemplateDetailsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetConsentTemplateDetailsDto() *GetConsentTemplateDetailsDto {
	this := GetConsentTemplateDetailsDto{}
	return &this
}

// NewGetConsentTemplateDetailsDtoWithDefaults instantiates a new GetConsentTemplateDetailsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetConsentTemplateDetailsDtoWithDefaults() *GetConsentTemplateDetailsDto {
	this := GetConsentTemplateDetailsDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetConsentTemplateDetailsDto) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConsentTemplateDetailsDto) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetConsentTemplateDetailsDto) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetConsentTemplateDetailsDto) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetConsentTemplateDetailsDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetConsentTemplateDetailsDto) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *GetConsentTemplateDetailsDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *GetConsentTemplateDetailsDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *GetConsentTemplateDetailsDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *GetConsentTemplateDetailsDto) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetConsentTemplateDetailsDto) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetConsentTemplateDetailsDto) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *GetConsentTemplateDetailsDto) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *GetConsentTemplateDetailsDto) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *GetConsentTemplateDetailsDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *GetConsentTemplateDetailsDto) UnsetDescription() {
	o.Description.Unset()
}

// GetConsentPurpose returns the ConsentPurpose field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetConsentTemplateDetailsDto) GetConsentPurpose() string {
	if o == nil || o.ConsentPurpose.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConsentPurpose.Get()
}

// GetConsentPurposeOk returns a tuple with the ConsentPurpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetConsentTemplateDetailsDto) GetConsentPurposeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConsentPurpose.Get(), o.ConsentPurpose.IsSet()
}

// HasConsentPurpose returns a boolean if a field has been set.
func (o *GetConsentTemplateDetailsDto) HasConsentPurpose() bool {
	if o != nil && o.ConsentPurpose.IsSet() {
		return true
	}

	return false
}

// SetConsentPurpose gets a reference to the given NullableString and assigns it to the ConsentPurpose field.
func (o *GetConsentTemplateDetailsDto) SetConsentPurpose(v string) {
	o.ConsentPurpose.Set(&v)
}
// SetConsentPurposeNil sets the value for ConsentPurpose to be an explicit nil
func (o *GetConsentTemplateDetailsDto) SetConsentPurposeNil() {
	o.ConsentPurpose.Set(nil)
}

// UnsetConsentPurpose ensures that no value is present for ConsentPurpose, not even an explicit nil
func (o *GetConsentTemplateDetailsDto) UnsetConsentPurpose() {
	o.ConsentPurpose.Unset()
}

// GetCollectables returns the Collectables field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetConsentTemplateDetailsDto) GetCollectables() []CollectibleTypes {
	if o == nil  {
		var ret []CollectibleTypes
		return ret
	}
	return o.Collectables
}

// GetCollectablesOk returns a tuple with the Collectables field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetConsentTemplateDetailsDto) GetCollectablesOk() ([]CollectibleTypes, bool) {
	if o == nil || o.Collectables == nil {
		return nil, false
	}
	return o.Collectables, true
}

// HasCollectables returns a boolean if a field has been set.
func (o *GetConsentTemplateDetailsDto) HasCollectables() bool {
	if o != nil && o.Collectables != nil {
		return true
	}

	return false
}

// SetCollectables gets a reference to the given []CollectibleTypes and assigns it to the Collectables field.
func (o *GetConsentTemplateDetailsDto) SetCollectables(v []CollectibleTypes) {
	o.Collectables = v
}

// GetFetchType returns the FetchType field value if set, zero value otherwise.
func (o *GetConsentTemplateDetailsDto) GetFetchType() FetchTypes {
	if o == nil || o.FetchType == nil {
		var ret FetchTypes
		return ret
	}
	return *o.FetchType
}

// GetFetchTypeOk returns a tuple with the FetchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConsentTemplateDetailsDto) GetFetchTypeOk() (*FetchTypes, bool) {
	if o == nil || o.FetchType == nil {
		return nil, false
	}
	return o.FetchType, true
}

// HasFetchType returns a boolean if a field has been set.
func (o *GetConsentTemplateDetailsDto) HasFetchType() bool {
	if o != nil && o.FetchType != nil {
		return true
	}

	return false
}

// SetFetchType gets a reference to the given FetchTypes and assigns it to the FetchType field.
func (o *GetConsentTemplateDetailsDto) SetFetchType(v FetchTypes) {
	o.FetchType = &v
}

// GetShortId returns the ShortId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetConsentTemplateDetailsDto) GetShortId() string {
	if o == nil || o.ShortId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ShortId.Get()
}

// GetShortIdOk returns a tuple with the ShortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetConsentTemplateDetailsDto) GetShortIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ShortId.Get(), o.ShortId.IsSet()
}

// HasShortId returns a boolean if a field has been set.
func (o *GetConsentTemplateDetailsDto) HasShortId() bool {
	if o != nil && o.ShortId.IsSet() {
		return true
	}

	return false
}

// SetShortId gets a reference to the given NullableString and assigns it to the ShortId field.
func (o *GetConsentTemplateDetailsDto) SetShortId(v string) {
	o.ShortId.Set(&v)
}
// SetShortIdNil sets the value for ShortId to be an explicit nil
func (o *GetConsentTemplateDetailsDto) SetShortIdNil() {
	o.ShortId.Set(nil)
}

// UnsetShortId ensures that no value is present for ShortId, not even an explicit nil
func (o *GetConsentTemplateDetailsDto) UnsetShortId() {
	o.ShortId.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetConsentTemplateDetailsDto) GetCreatedBy() string {
	if o == nil || o.CreatedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetConsentTemplateDetailsDto) GetCreatedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *GetConsentTemplateDetailsDto) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *GetConsentTemplateDetailsDto) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}
// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *GetConsentTemplateDetailsDto) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *GetConsentTemplateDetailsDto) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetCreatedAtUtc returns the CreatedAtUtc field value if set, zero value otherwise.
func (o *GetConsentTemplateDetailsDto) GetCreatedAtUtc() time.Time {
	if o == nil || o.CreatedAtUtc == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAtUtc
}

// GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConsentTemplateDetailsDto) GetCreatedAtUtcOk() (*time.Time, bool) {
	if o == nil || o.CreatedAtUtc == nil {
		return nil, false
	}
	return o.CreatedAtUtc, true
}

// HasCreatedAtUtc returns a boolean if a field has been set.
func (o *GetConsentTemplateDetailsDto) HasCreatedAtUtc() bool {
	if o != nil && o.CreatedAtUtc != nil {
		return true
	}

	return false
}

// SetCreatedAtUtc gets a reference to the given time.Time and assigns it to the CreatedAtUtc field.
func (o *GetConsentTemplateDetailsDto) SetCreatedAtUtc(v time.Time) {
	o.CreatedAtUtc = &v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetConsentTemplateDetailsDto) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetConsentTemplateDetailsDto) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *GetConsentTemplateDetailsDto) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *GetConsentTemplateDetailsDto) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *GetConsentTemplateDetailsDto) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *GetConsentTemplateDetailsDto) UnsetStatus() {
	o.Status.Unset()
}

// GetTemplateType returns the TemplateType field value if set, zero value otherwise.
func (o *GetConsentTemplateDetailsDto) GetTemplateType() ConsentTemplateTypes {
	if o == nil || o.TemplateType == nil {
		var ret ConsentTemplateTypes
		return ret
	}
	return *o.TemplateType
}

// GetTemplateTypeOk returns a tuple with the TemplateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConsentTemplateDetailsDto) GetTemplateTypeOk() (*ConsentTemplateTypes, bool) {
	if o == nil || o.TemplateType == nil {
		return nil, false
	}
	return o.TemplateType, true
}

// HasTemplateType returns a boolean if a field has been set.
func (o *GetConsentTemplateDetailsDto) HasTemplateType() bool {
	if o != nil && o.TemplateType != nil {
		return true
	}

	return false
}

// SetTemplateType gets a reference to the given ConsentTemplateTypes and assigns it to the TemplateType field.
func (o *GetConsentTemplateDetailsDto) SetTemplateType(v ConsentTemplateTypes) {
	o.TemplateType = &v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *GetConsentTemplateDetailsDto) GetFrequency() Life {
	if o == nil || o.Frequency == nil {
		var ret Life
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetConsentTemplateDetailsDto) GetFrequencyOk() (*Life, bool) {
	if o == nil || o.Frequency == nil {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *GetConsentTemplateDetailsDto) HasFrequency() bool {
	if o != nil && o.Frequency != nil {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given Life and assigns it to the Frequency field.
func (o *GetConsentTemplateDetailsDto) SetFrequency(v Life) {
	o.Frequency = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetConsentTemplateDetailsDto) GetIdentity() []IdentitySupportedFields {
	if o == nil  {
		var ret []IdentitySupportedFields
		return ret
	}
	return o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetConsentTemplateDetailsDto) GetIdentityOk() ([]IdentitySupportedFields, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *GetConsentTemplateDetailsDto) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given []IdentitySupportedFields and assigns it to the Identity field.
func (o *GetConsentTemplateDetailsDto) SetIdentity(v []IdentitySupportedFields) {
	o.Identity = v
}

// GetDocuments returns the Documents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetConsentTemplateDetailsDto) GetDocuments() []Document {
	if o == nil  {
		var ret []Document
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetConsentTemplateDetailsDto) GetDocumentsOk() ([]Document, bool) {
	if o == nil || o.Documents == nil {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *GetConsentTemplateDetailsDto) HasDocuments() bool {
	if o != nil && o.Documents != nil {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []Document and assigns it to the Documents field.
func (o *GetConsentTemplateDetailsDto) SetDocuments(v []Document) {
	o.Documents = v
}

// GetFinancials returns the Financials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetConsentTemplateDetailsDto) GetFinancials() []Financial {
	if o == nil  {
		var ret []Financial
		return ret
	}
	return o.Financials
}

// GetFinancialsOk returns a tuple with the Financials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetConsentTemplateDetailsDto) GetFinancialsOk() ([]Financial, bool) {
	if o == nil || o.Financials == nil {
		return nil, false
	}
	return o.Financials, true
}

// HasFinancials returns a boolean if a field has been set.
func (o *GetConsentTemplateDetailsDto) HasFinancials() bool {
	if o != nil && o.Financials != nil {
		return true
	}

	return false
}

// SetFinancials gets a reference to the given []Financial and assigns it to the Financials field.
func (o *GetConsentTemplateDetailsDto) SetFinancials(v []Financial) {
	o.Financials = v
}

func (o GetConsentTemplateDetailsDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.ConsentPurpose.IsSet() {
		toSerialize["consentPurpose"] = o.ConsentPurpose.Get()
	}
	if o.Collectables != nil {
		toSerialize["collectables"] = o.Collectables
	}
	if o.FetchType != nil {
		toSerialize["fetchType"] = o.FetchType
	}
	if o.ShortId.IsSet() {
		toSerialize["shortId"] = o.ShortId.Get()
	}
	if o.CreatedBy.IsSet() {
		toSerialize["createdBy"] = o.CreatedBy.Get()
	}
	if o.CreatedAtUtc != nil {
		toSerialize["createdAtUtc"] = o.CreatedAtUtc
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.TemplateType != nil {
		toSerialize["templateType"] = o.TemplateType
	}
	if o.Frequency != nil {
		toSerialize["frequency"] = o.Frequency
	}
	if o.Identity != nil {
		toSerialize["identity"] = o.Identity
	}
	if o.Documents != nil {
		toSerialize["documents"] = o.Documents
	}
	if o.Financials != nil {
		toSerialize["financials"] = o.Financials
	}
	return json.Marshal(toSerialize)
}

type NullableGetConsentTemplateDetailsDto struct {
	value *GetConsentTemplateDetailsDto
	isSet bool
}

func (v NullableGetConsentTemplateDetailsDto) Get() *GetConsentTemplateDetailsDto {
	return v.value
}

func (v *NullableGetConsentTemplateDetailsDto) Set(val *GetConsentTemplateDetailsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableGetConsentTemplateDetailsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableGetConsentTemplateDetailsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetConsentTemplateDetailsDto(val *GetConsentTemplateDetailsDto) *NullableGetConsentTemplateDetailsDto {
	return &NullableGetConsentTemplateDetailsDto{value: val, isSet: true}
}

func (v NullableGetConsentTemplateDetailsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetConsentTemplateDetailsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


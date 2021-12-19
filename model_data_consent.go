/*
My Data My Consent - Developer API

Unleashing the power of data consent by establishing trust. The Platform Core Developer API defines a set of capabilities that can be used to request, issue, manage and update data, documents and credentials by organizations. The API can be used to request, manage and update Decentralised Identifiers, Financial Data, Health Data issue Documents, Credentials directly or using OpenID Connect flows, and verify Messages signed with DIDs and much more.

API version: v1
Contact: support@mydatamyconsent.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mydatamyconsent

import (
	"encoding/json"
	"time"
)

// DataConsent struct for DataConsent
type DataConsent struct {
	Id *string `json:"id,omitempty"`
	UserId NullableString `json:"userId,omitempty"`
	OrganizationId NullableString `json:"organizationId,omitempty"`
	RequestedByOrgId *string `json:"requestedByOrgId,omitempty"`
	TransactionId NullableString `json:"transactionId,omitempty"`
	StartDateTime NullableTime `json:"startDateTime,omitempty"`
	ExpiryDateTime *time.Time `json:"expiryDateTime,omitempty"`
	Description NullableString `json:"description,omitempty"`
	PurposeCode NullableString `json:"purposeCode,omitempty"`
	PurposeLink NullableString `json:"purposeLink,omitempty"`
	Location NullableString `json:"location,omitempty"`
	AgreementId *string `json:"agreementId,omitempty"`
	DataLifeUnit *DataLifeUnit `json:"dataLifeUnit,omitempty"`
	DataLifeValue *int32 `json:"dataLifeValue,omitempty"`
	DataFetchFrequencyUnit *DataFetchFrequencyUnit `json:"dataFetchFrequencyUnit,omitempty"`
	DataFetchFrequencyUnitValue *int32 `json:"dataFetchFrequencyUnitValue,omitempty"`
	DataFetchType *DataFetchType `json:"dataFetchType,omitempty"`
	Status *DataConsentStatus `json:"status,omitempty"`
	CreatedAtUtc *time.Time `json:"createdAtUtc,omitempty"`
	ApprovedAtUtc NullableTime `json:"approvedAtUtc,omitempty"`
	RejectedAtUtc NullableTime `json:"rejectedAtUtc,omitempty"`
	User *ApplicationUser `json:"user,omitempty"`
	Organization *Organization `json:"organization,omitempty"`
	RequestedByOrg *Organization `json:"requestedByOrg,omitempty"`
	Agreement *DataProcessingAgreement `json:"agreement,omitempty"`
	IdentityClaims []IdentityClaim `json:"identityClaims,omitempty"`
	Identifiers []DataConsentIdentifier `json:"identifiers,omitempty"`
	RequestedFinancialAccounts []DataConsentRequestedFinancialAccount `json:"requestedFinancialAccounts,omitempty"`
	RequestedDocuments []DataConsentRequestedDocument `json:"requestedDocuments,omitempty"`
}

// NewDataConsent instantiates a new DataConsent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataConsent() *DataConsent {
	this := DataConsent{}
	return &this
}

// NewDataConsentWithDefaults instantiates a new DataConsent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataConsentWithDefaults() *DataConsent {
	this := DataConsent{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DataConsent) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataConsent) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DataConsent) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DataConsent) SetId(v string) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsent) GetUserId() string {
	if o == nil || o.UserId.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsent) GetUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *DataConsent) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *DataConsent) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *DataConsent) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *DataConsent) UnsetUserId() {
	o.UserId.Unset()
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsent) GetOrganizationId() string {
	if o == nil || o.OrganizationId.Get() == nil {
		var ret string
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsent) GetOrganizationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *DataConsent) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableString and assigns it to the OrganizationId field.
func (o *DataConsent) SetOrganizationId(v string) {
	o.OrganizationId.Set(&v)
}
// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *DataConsent) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *DataConsent) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

// GetRequestedByOrgId returns the RequestedByOrgId field value if set, zero value otherwise.
func (o *DataConsent) GetRequestedByOrgId() string {
	if o == nil || o.RequestedByOrgId == nil {
		var ret string
		return ret
	}
	return *o.RequestedByOrgId
}

// GetRequestedByOrgIdOk returns a tuple with the RequestedByOrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataConsent) GetRequestedByOrgIdOk() (*string, bool) {
	if o == nil || o.RequestedByOrgId == nil {
		return nil, false
	}
	return o.RequestedByOrgId, true
}

// HasRequestedByOrgId returns a boolean if a field has been set.
func (o *DataConsent) HasRequestedByOrgId() bool {
	if o != nil && o.RequestedByOrgId != nil {
		return true
	}

	return false
}

// SetRequestedByOrgId gets a reference to the given string and assigns it to the RequestedByOrgId field.
func (o *DataConsent) SetRequestedByOrgId(v string) {
	o.RequestedByOrgId = &v
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

// GetStartDateTime returns the StartDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsent) GetStartDateTime() time.Time {
	if o == nil || o.StartDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDateTime.Get()
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsent) GetStartDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDateTime.Get(), o.StartDateTime.IsSet()
}

// HasStartDateTime returns a boolean if a field has been set.
func (o *DataConsent) HasStartDateTime() bool {
	if o != nil && o.StartDateTime.IsSet() {
		return true
	}

	return false
}

// SetStartDateTime gets a reference to the given NullableTime and assigns it to the StartDateTime field.
func (o *DataConsent) SetStartDateTime(v time.Time) {
	o.StartDateTime.Set(&v)
}
// SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil
func (o *DataConsent) SetStartDateTimeNil() {
	o.StartDateTime.Set(nil)
}

// UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil
func (o *DataConsent) UnsetStartDateTime() {
	o.StartDateTime.Unset()
}

// GetExpiryDateTime returns the ExpiryDateTime field value if set, zero value otherwise.
func (o *DataConsent) GetExpiryDateTime() time.Time {
	if o == nil || o.ExpiryDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDateTime
}

// GetExpiryDateTimeOk returns a tuple with the ExpiryDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataConsent) GetExpiryDateTimeOk() (*time.Time, bool) {
	if o == nil || o.ExpiryDateTime == nil {
		return nil, false
	}
	return o.ExpiryDateTime, true
}

// HasExpiryDateTime returns a boolean if a field has been set.
func (o *DataConsent) HasExpiryDateTime() bool {
	if o != nil && o.ExpiryDateTime != nil {
		return true
	}

	return false
}

// SetExpiryDateTime gets a reference to the given time.Time and assigns it to the ExpiryDateTime field.
func (o *DataConsent) SetExpiryDateTime(v time.Time) {
	o.ExpiryDateTime = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsent) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsent) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *DataConsent) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *DataConsent) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *DataConsent) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *DataConsent) UnsetDescription() {
	o.Description.Unset()
}

// GetPurposeCode returns the PurposeCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsent) GetPurposeCode() string {
	if o == nil || o.PurposeCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.PurposeCode.Get()
}

// GetPurposeCodeOk returns a tuple with the PurposeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsent) GetPurposeCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PurposeCode.Get(), o.PurposeCode.IsSet()
}

// HasPurposeCode returns a boolean if a field has been set.
func (o *DataConsent) HasPurposeCode() bool {
	if o != nil && o.PurposeCode.IsSet() {
		return true
	}

	return false
}

// SetPurposeCode gets a reference to the given NullableString and assigns it to the PurposeCode field.
func (o *DataConsent) SetPurposeCode(v string) {
	o.PurposeCode.Set(&v)
}
// SetPurposeCodeNil sets the value for PurposeCode to be an explicit nil
func (o *DataConsent) SetPurposeCodeNil() {
	o.PurposeCode.Set(nil)
}

// UnsetPurposeCode ensures that no value is present for PurposeCode, not even an explicit nil
func (o *DataConsent) UnsetPurposeCode() {
	o.PurposeCode.Unset()
}

// GetPurposeLink returns the PurposeLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsent) GetPurposeLink() string {
	if o == nil || o.PurposeLink.Get() == nil {
		var ret string
		return ret
	}
	return *o.PurposeLink.Get()
}

// GetPurposeLinkOk returns a tuple with the PurposeLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsent) GetPurposeLinkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PurposeLink.Get(), o.PurposeLink.IsSet()
}

// HasPurposeLink returns a boolean if a field has been set.
func (o *DataConsent) HasPurposeLink() bool {
	if o != nil && o.PurposeLink.IsSet() {
		return true
	}

	return false
}

// SetPurposeLink gets a reference to the given NullableString and assigns it to the PurposeLink field.
func (o *DataConsent) SetPurposeLink(v string) {
	o.PurposeLink.Set(&v)
}
// SetPurposeLinkNil sets the value for PurposeLink to be an explicit nil
func (o *DataConsent) SetPurposeLinkNil() {
	o.PurposeLink.Set(nil)
}

// UnsetPurposeLink ensures that no value is present for PurposeLink, not even an explicit nil
func (o *DataConsent) UnsetPurposeLink() {
	o.PurposeLink.Unset()
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsent) GetLocation() string {
	if o == nil || o.Location.Get() == nil {
		var ret string
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsent) GetLocationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *DataConsent) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableString and assigns it to the Location field.
func (o *DataConsent) SetLocation(v string) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *DataConsent) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *DataConsent) UnsetLocation() {
	o.Location.Unset()
}

// GetAgreementId returns the AgreementId field value if set, zero value otherwise.
func (o *DataConsent) GetAgreementId() string {
	if o == nil || o.AgreementId == nil {
		var ret string
		return ret
	}
	return *o.AgreementId
}

// GetAgreementIdOk returns a tuple with the AgreementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataConsent) GetAgreementIdOk() (*string, bool) {
	if o == nil || o.AgreementId == nil {
		return nil, false
	}
	return o.AgreementId, true
}

// HasAgreementId returns a boolean if a field has been set.
func (o *DataConsent) HasAgreementId() bool {
	if o != nil && o.AgreementId != nil {
		return true
	}

	return false
}

// SetAgreementId gets a reference to the given string and assigns it to the AgreementId field.
func (o *DataConsent) SetAgreementId(v string) {
	o.AgreementId = &v
}

// GetDataLifeUnit returns the DataLifeUnit field value if set, zero value otherwise.
func (o *DataConsent) GetDataLifeUnit() DataLifeUnit {
	if o == nil || o.DataLifeUnit == nil {
		var ret DataLifeUnit
		return ret
	}
	return *o.DataLifeUnit
}

// GetDataLifeUnitOk returns a tuple with the DataLifeUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataConsent) GetDataLifeUnitOk() (*DataLifeUnit, bool) {
	if o == nil || o.DataLifeUnit == nil {
		return nil, false
	}
	return o.DataLifeUnit, true
}

// HasDataLifeUnit returns a boolean if a field has been set.
func (o *DataConsent) HasDataLifeUnit() bool {
	if o != nil && o.DataLifeUnit != nil {
		return true
	}

	return false
}

// SetDataLifeUnit gets a reference to the given DataLifeUnit and assigns it to the DataLifeUnit field.
func (o *DataConsent) SetDataLifeUnit(v DataLifeUnit) {
	o.DataLifeUnit = &v
}

// GetDataLifeValue returns the DataLifeValue field value if set, zero value otherwise.
func (o *DataConsent) GetDataLifeValue() int32 {
	if o == nil || o.DataLifeValue == nil {
		var ret int32
		return ret
	}
	return *o.DataLifeValue
}

// GetDataLifeValueOk returns a tuple with the DataLifeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataConsent) GetDataLifeValueOk() (*int32, bool) {
	if o == nil || o.DataLifeValue == nil {
		return nil, false
	}
	return o.DataLifeValue, true
}

// HasDataLifeValue returns a boolean if a field has been set.
func (o *DataConsent) HasDataLifeValue() bool {
	if o != nil && o.DataLifeValue != nil {
		return true
	}

	return false
}

// SetDataLifeValue gets a reference to the given int32 and assigns it to the DataLifeValue field.
func (o *DataConsent) SetDataLifeValue(v int32) {
	o.DataLifeValue = &v
}

// GetDataFetchFrequencyUnit returns the DataFetchFrequencyUnit field value if set, zero value otherwise.
func (o *DataConsent) GetDataFetchFrequencyUnit() DataFetchFrequencyUnit {
	if o == nil || o.DataFetchFrequencyUnit == nil {
		var ret DataFetchFrequencyUnit
		return ret
	}
	return *o.DataFetchFrequencyUnit
}

// GetDataFetchFrequencyUnitOk returns a tuple with the DataFetchFrequencyUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataConsent) GetDataFetchFrequencyUnitOk() (*DataFetchFrequencyUnit, bool) {
	if o == nil || o.DataFetchFrequencyUnit == nil {
		return nil, false
	}
	return o.DataFetchFrequencyUnit, true
}

// HasDataFetchFrequencyUnit returns a boolean if a field has been set.
func (o *DataConsent) HasDataFetchFrequencyUnit() bool {
	if o != nil && o.DataFetchFrequencyUnit != nil {
		return true
	}

	return false
}

// SetDataFetchFrequencyUnit gets a reference to the given DataFetchFrequencyUnit and assigns it to the DataFetchFrequencyUnit field.
func (o *DataConsent) SetDataFetchFrequencyUnit(v DataFetchFrequencyUnit) {
	o.DataFetchFrequencyUnit = &v
}

// GetDataFetchFrequencyUnitValue returns the DataFetchFrequencyUnitValue field value if set, zero value otherwise.
func (o *DataConsent) GetDataFetchFrequencyUnitValue() int32 {
	if o == nil || o.DataFetchFrequencyUnitValue == nil {
		var ret int32
		return ret
	}
	return *o.DataFetchFrequencyUnitValue
}

// GetDataFetchFrequencyUnitValueOk returns a tuple with the DataFetchFrequencyUnitValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataConsent) GetDataFetchFrequencyUnitValueOk() (*int32, bool) {
	if o == nil || o.DataFetchFrequencyUnitValue == nil {
		return nil, false
	}
	return o.DataFetchFrequencyUnitValue, true
}

// HasDataFetchFrequencyUnitValue returns a boolean if a field has been set.
func (o *DataConsent) HasDataFetchFrequencyUnitValue() bool {
	if o != nil && o.DataFetchFrequencyUnitValue != nil {
		return true
	}

	return false
}

// SetDataFetchFrequencyUnitValue gets a reference to the given int32 and assigns it to the DataFetchFrequencyUnitValue field.
func (o *DataConsent) SetDataFetchFrequencyUnitValue(v int32) {
	o.DataFetchFrequencyUnitValue = &v
}

// GetDataFetchType returns the DataFetchType field value if set, zero value otherwise.
func (o *DataConsent) GetDataFetchType() DataFetchType {
	if o == nil || o.DataFetchType == nil {
		var ret DataFetchType
		return ret
	}
	return *o.DataFetchType
}

// GetDataFetchTypeOk returns a tuple with the DataFetchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataConsent) GetDataFetchTypeOk() (*DataFetchType, bool) {
	if o == nil || o.DataFetchType == nil {
		return nil, false
	}
	return o.DataFetchType, true
}

// HasDataFetchType returns a boolean if a field has been set.
func (o *DataConsent) HasDataFetchType() bool {
	if o != nil && o.DataFetchType != nil {
		return true
	}

	return false
}

// SetDataFetchType gets a reference to the given DataFetchType and assigns it to the DataFetchType field.
func (o *DataConsent) SetDataFetchType(v DataFetchType) {
	o.DataFetchType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DataConsent) GetStatus() DataConsentStatus {
	if o == nil || o.Status == nil {
		var ret DataConsentStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataConsent) GetStatusOk() (*DataConsentStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DataConsent) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given DataConsentStatus and assigns it to the Status field.
func (o *DataConsent) SetStatus(v DataConsentStatus) {
	o.Status = &v
}

// GetCreatedAtUtc returns the CreatedAtUtc field value if set, zero value otherwise.
func (o *DataConsent) GetCreatedAtUtc() time.Time {
	if o == nil || o.CreatedAtUtc == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAtUtc
}

// GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataConsent) GetCreatedAtUtcOk() (*time.Time, bool) {
	if o == nil || o.CreatedAtUtc == nil {
		return nil, false
	}
	return o.CreatedAtUtc, true
}

// HasCreatedAtUtc returns a boolean if a field has been set.
func (o *DataConsent) HasCreatedAtUtc() bool {
	if o != nil && o.CreatedAtUtc != nil {
		return true
	}

	return false
}

// SetCreatedAtUtc gets a reference to the given time.Time and assigns it to the CreatedAtUtc field.
func (o *DataConsent) SetCreatedAtUtc(v time.Time) {
	o.CreatedAtUtc = &v
}

// GetApprovedAtUtc returns the ApprovedAtUtc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsent) GetApprovedAtUtc() time.Time {
	if o == nil || o.ApprovedAtUtc.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ApprovedAtUtc.Get()
}

// GetApprovedAtUtcOk returns a tuple with the ApprovedAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsent) GetApprovedAtUtcOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ApprovedAtUtc.Get(), o.ApprovedAtUtc.IsSet()
}

// HasApprovedAtUtc returns a boolean if a field has been set.
func (o *DataConsent) HasApprovedAtUtc() bool {
	if o != nil && o.ApprovedAtUtc.IsSet() {
		return true
	}

	return false
}

// SetApprovedAtUtc gets a reference to the given NullableTime and assigns it to the ApprovedAtUtc field.
func (o *DataConsent) SetApprovedAtUtc(v time.Time) {
	o.ApprovedAtUtc.Set(&v)
}
// SetApprovedAtUtcNil sets the value for ApprovedAtUtc to be an explicit nil
func (o *DataConsent) SetApprovedAtUtcNil() {
	o.ApprovedAtUtc.Set(nil)
}

// UnsetApprovedAtUtc ensures that no value is present for ApprovedAtUtc, not even an explicit nil
func (o *DataConsent) UnsetApprovedAtUtc() {
	o.ApprovedAtUtc.Unset()
}

// GetRejectedAtUtc returns the RejectedAtUtc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsent) GetRejectedAtUtc() time.Time {
	if o == nil || o.RejectedAtUtc.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.RejectedAtUtc.Get()
}

// GetRejectedAtUtcOk returns a tuple with the RejectedAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsent) GetRejectedAtUtcOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RejectedAtUtc.Get(), o.RejectedAtUtc.IsSet()
}

// HasRejectedAtUtc returns a boolean if a field has been set.
func (o *DataConsent) HasRejectedAtUtc() bool {
	if o != nil && o.RejectedAtUtc.IsSet() {
		return true
	}

	return false
}

// SetRejectedAtUtc gets a reference to the given NullableTime and assigns it to the RejectedAtUtc field.
func (o *DataConsent) SetRejectedAtUtc(v time.Time) {
	o.RejectedAtUtc.Set(&v)
}
// SetRejectedAtUtcNil sets the value for RejectedAtUtc to be an explicit nil
func (o *DataConsent) SetRejectedAtUtcNil() {
	o.RejectedAtUtc.Set(nil)
}

// UnsetRejectedAtUtc ensures that no value is present for RejectedAtUtc, not even an explicit nil
func (o *DataConsent) UnsetRejectedAtUtc() {
	o.RejectedAtUtc.Unset()
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *DataConsent) GetUser() ApplicationUser {
	if o == nil || o.User == nil {
		var ret ApplicationUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataConsent) GetUserOk() (*ApplicationUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *DataConsent) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given ApplicationUser and assigns it to the User field.
func (o *DataConsent) SetUser(v ApplicationUser) {
	o.User = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *DataConsent) GetOrganization() Organization {
	if o == nil || o.Organization == nil {
		var ret Organization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataConsent) GetOrganizationOk() (*Organization, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *DataConsent) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Organization and assigns it to the Organization field.
func (o *DataConsent) SetOrganization(v Organization) {
	o.Organization = &v
}

// GetRequestedByOrg returns the RequestedByOrg field value if set, zero value otherwise.
func (o *DataConsent) GetRequestedByOrg() Organization {
	if o == nil || o.RequestedByOrg == nil {
		var ret Organization
		return ret
	}
	return *o.RequestedByOrg
}

// GetRequestedByOrgOk returns a tuple with the RequestedByOrg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataConsent) GetRequestedByOrgOk() (*Organization, bool) {
	if o == nil || o.RequestedByOrg == nil {
		return nil, false
	}
	return o.RequestedByOrg, true
}

// HasRequestedByOrg returns a boolean if a field has been set.
func (o *DataConsent) HasRequestedByOrg() bool {
	if o != nil && o.RequestedByOrg != nil {
		return true
	}

	return false
}

// SetRequestedByOrg gets a reference to the given Organization and assigns it to the RequestedByOrg field.
func (o *DataConsent) SetRequestedByOrg(v Organization) {
	o.RequestedByOrg = &v
}

// GetAgreement returns the Agreement field value if set, zero value otherwise.
func (o *DataConsent) GetAgreement() DataProcessingAgreement {
	if o == nil || o.Agreement == nil {
		var ret DataProcessingAgreement
		return ret
	}
	return *o.Agreement
}

// GetAgreementOk returns a tuple with the Agreement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataConsent) GetAgreementOk() (*DataProcessingAgreement, bool) {
	if o == nil || o.Agreement == nil {
		return nil, false
	}
	return o.Agreement, true
}

// HasAgreement returns a boolean if a field has been set.
func (o *DataConsent) HasAgreement() bool {
	if o != nil && o.Agreement != nil {
		return true
	}

	return false
}

// SetAgreement gets a reference to the given DataProcessingAgreement and assigns it to the Agreement field.
func (o *DataConsent) SetAgreement(v DataProcessingAgreement) {
	o.Agreement = &v
}

// GetIdentityClaims returns the IdentityClaims field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsent) GetIdentityClaims() []IdentityClaim {
	if o == nil  {
		var ret []IdentityClaim
		return ret
	}
	return o.IdentityClaims
}

// GetIdentityClaimsOk returns a tuple with the IdentityClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsent) GetIdentityClaimsOk() (*[]IdentityClaim, bool) {
	if o == nil || o.IdentityClaims == nil {
		return nil, false
	}
	return &o.IdentityClaims, true
}

// HasIdentityClaims returns a boolean if a field has been set.
func (o *DataConsent) HasIdentityClaims() bool {
	if o != nil && o.IdentityClaims != nil {
		return true
	}

	return false
}

// SetIdentityClaims gets a reference to the given []IdentityClaim and assigns it to the IdentityClaims field.
func (o *DataConsent) SetIdentityClaims(v []IdentityClaim) {
	o.IdentityClaims = v
}

// GetIdentifiers returns the Identifiers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsent) GetIdentifiers() []DataConsentIdentifier {
	if o == nil  {
		var ret []DataConsentIdentifier
		return ret
	}
	return o.Identifiers
}

// GetIdentifiersOk returns a tuple with the Identifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsent) GetIdentifiersOk() (*[]DataConsentIdentifier, bool) {
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

// SetIdentifiers gets a reference to the given []DataConsentIdentifier and assigns it to the Identifiers field.
func (o *DataConsent) SetIdentifiers(v []DataConsentIdentifier) {
	o.Identifiers = v
}

// GetRequestedFinancialAccounts returns the RequestedFinancialAccounts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsent) GetRequestedFinancialAccounts() []DataConsentRequestedFinancialAccount {
	if o == nil  {
		var ret []DataConsentRequestedFinancialAccount
		return ret
	}
	return o.RequestedFinancialAccounts
}

// GetRequestedFinancialAccountsOk returns a tuple with the RequestedFinancialAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsent) GetRequestedFinancialAccountsOk() (*[]DataConsentRequestedFinancialAccount, bool) {
	if o == nil || o.RequestedFinancialAccounts == nil {
		return nil, false
	}
	return &o.RequestedFinancialAccounts, true
}

// HasRequestedFinancialAccounts returns a boolean if a field has been set.
func (o *DataConsent) HasRequestedFinancialAccounts() bool {
	if o != nil && o.RequestedFinancialAccounts != nil {
		return true
	}

	return false
}

// SetRequestedFinancialAccounts gets a reference to the given []DataConsentRequestedFinancialAccount and assigns it to the RequestedFinancialAccounts field.
func (o *DataConsent) SetRequestedFinancialAccounts(v []DataConsentRequestedFinancialAccount) {
	o.RequestedFinancialAccounts = v
}

// GetRequestedDocuments returns the RequestedDocuments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsent) GetRequestedDocuments() []DataConsentRequestedDocument {
	if o == nil  {
		var ret []DataConsentRequestedDocument
		return ret
	}
	return o.RequestedDocuments
}

// GetRequestedDocumentsOk returns a tuple with the RequestedDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsent) GetRequestedDocumentsOk() (*[]DataConsentRequestedDocument, bool) {
	if o == nil || o.RequestedDocuments == nil {
		return nil, false
	}
	return &o.RequestedDocuments, true
}

// HasRequestedDocuments returns a boolean if a field has been set.
func (o *DataConsent) HasRequestedDocuments() bool {
	if o != nil && o.RequestedDocuments != nil {
		return true
	}

	return false
}

// SetRequestedDocuments gets a reference to the given []DataConsentRequestedDocument and assigns it to the RequestedDocuments field.
func (o *DataConsent) SetRequestedDocuments(v []DataConsentRequestedDocument) {
	o.RequestedDocuments = v
}

func (o DataConsent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.UserId.IsSet() {
		toSerialize["userId"] = o.UserId.Get()
	}
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
	}
	if o.RequestedByOrgId != nil {
		toSerialize["requestedByOrgId"] = o.RequestedByOrgId
	}
	if o.TransactionId.IsSet() {
		toSerialize["transactionId"] = o.TransactionId.Get()
	}
	if o.StartDateTime.IsSet() {
		toSerialize["startDateTime"] = o.StartDateTime.Get()
	}
	if o.ExpiryDateTime != nil {
		toSerialize["expiryDateTime"] = o.ExpiryDateTime
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.PurposeCode.IsSet() {
		toSerialize["purposeCode"] = o.PurposeCode.Get()
	}
	if o.PurposeLink.IsSet() {
		toSerialize["purposeLink"] = o.PurposeLink.Get()
	}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if o.AgreementId != nil {
		toSerialize["agreementId"] = o.AgreementId
	}
	if o.DataLifeUnit != nil {
		toSerialize["dataLifeUnit"] = o.DataLifeUnit
	}
	if o.DataLifeValue != nil {
		toSerialize["dataLifeValue"] = o.DataLifeValue
	}
	if o.DataFetchFrequencyUnit != nil {
		toSerialize["dataFetchFrequencyUnit"] = o.DataFetchFrequencyUnit
	}
	if o.DataFetchFrequencyUnitValue != nil {
		toSerialize["dataFetchFrequencyUnitValue"] = o.DataFetchFrequencyUnitValue
	}
	if o.DataFetchType != nil {
		toSerialize["dataFetchType"] = o.DataFetchType
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.CreatedAtUtc != nil {
		toSerialize["createdAtUtc"] = o.CreatedAtUtc
	}
	if o.ApprovedAtUtc.IsSet() {
		toSerialize["approvedAtUtc"] = o.ApprovedAtUtc.Get()
	}
	if o.RejectedAtUtc.IsSet() {
		toSerialize["rejectedAtUtc"] = o.RejectedAtUtc.Get()
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Organization != nil {
		toSerialize["organization"] = o.Organization
	}
	if o.RequestedByOrg != nil {
		toSerialize["requestedByOrg"] = o.RequestedByOrg
	}
	if o.Agreement != nil {
		toSerialize["agreement"] = o.Agreement
	}
	if o.IdentityClaims != nil {
		toSerialize["identityClaims"] = o.IdentityClaims
	}
	if o.Identifiers != nil {
		toSerialize["identifiers"] = o.Identifiers
	}
	if o.RequestedFinancialAccounts != nil {
		toSerialize["requestedFinancialAccounts"] = o.RequestedFinancialAccounts
	}
	if o.RequestedDocuments != nil {
		toSerialize["requestedDocuments"] = o.RequestedDocuments
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



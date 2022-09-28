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

// OrganizationConsentRequestTemplateDetails OrganizationConsentRequestTemplateDetails : Organization consent request template details.
type OrganizationConsentRequestTemplateDetails struct {
	// Consent request template id.
	Id string `json:"id"`
	// Consent request template title.
	Title string `json:"title"`
	// Consent request template description.
	Description string `json:"description"`
	// Consent request template purpose.
	Purpose *string `json:"purpose,omitempty"`
	// Consent request template short Id.
	ShortId string `json:"shortId"`
	Status ConsentRequestTemplateStatus `json:"status"`
	DataLife *IndividualConsentRequestTemplateDetailsDataLife `json:"dataLife,omitempty"`
	RequestLife *IndividualConsentRequestTemplateDetailsRequestLife `json:"requestLife,omitempty"`
	DataFrequency *IndividualConsentRequestTemplateDetailsDataFrequency `json:"dataFrequency,omitempty"`
	// Consent request template identity fields.
	Identifiers []IdentityField `json:"identifiers,omitempty"`
	// Consent request template document fields.
	Documents []DocumentField `json:"documents,omitempty"`
	// Consent request template financial account fields.
	FinancialAccounts []FinancialAccountField `json:"financialAccounts,omitempty"`
	// Consent request template created by user.
	CreatedBy string `json:"createdBy"`
	// Consent request template created datetime in UTC timezone.
	CreatedAtUtc time.Time `json:"createdAtUtc"`
	// Consent request template approval datetime in UTC timezone.
	ApprovedAtUtc *time.Time `json:"approvedAtUtc,omitempty"`
	// Consent request template published datetime in UTC timezone.
	PublishedAtUtc *time.Time `json:"publishedAtUtc,omitempty"`
	// Consent request template rejection datetime in UTC timezone.
	RejectedAtUtc *time.Time `json:"rejectedAtUtc,omitempty"`
	// Consent request template rejection reason.
	RejectionReason *string `json:"rejectionReason,omitempty"`
}

// NewOrganizationConsentRequestTemplateDetails instantiates a new OrganizationConsentRequestTemplateDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationConsentRequestTemplateDetails(id string, title string, description string, shortId string, status ConsentRequestTemplateStatus, createdBy string, createdAtUtc time.Time) *OrganizationConsentRequestTemplateDetails {
	this := OrganizationConsentRequestTemplateDetails{}
	this.Id = id
	this.Title = title
	this.Description = description
	this.ShortId = shortId
	this.Status = status
	this.CreatedBy = createdBy
	this.CreatedAtUtc = createdAtUtc
	return &this
}

// NewOrganizationConsentRequestTemplateDetailsWithDefaults instantiates a new OrganizationConsentRequestTemplateDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationConsentRequestTemplateDetailsWithDefaults() *OrganizationConsentRequestTemplateDetails {
	this := OrganizationConsentRequestTemplateDetails{}
	return &this
}

// GetId returns the Id field value
func (o *OrganizationConsentRequestTemplateDetails) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrganizationConsentRequestTemplateDetails) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrganizationConsentRequestTemplateDetails) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value
func (o *OrganizationConsentRequestTemplateDetails) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *OrganizationConsentRequestTemplateDetails) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *OrganizationConsentRequestTemplateDetails) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value
func (o *OrganizationConsentRequestTemplateDetails) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *OrganizationConsentRequestTemplateDetails) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *OrganizationConsentRequestTemplateDetails) SetDescription(v string) {
	o.Description = v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *OrganizationConsentRequestTemplateDetails) GetPurpose() string {
	if o == nil || o.Purpose == nil {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationConsentRequestTemplateDetails) GetPurposeOk() (*string, bool) {
	if o == nil || o.Purpose == nil {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *OrganizationConsentRequestTemplateDetails) HasPurpose() bool {
	if o != nil && o.Purpose != nil {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *OrganizationConsentRequestTemplateDetails) SetPurpose(v string) {
	o.Purpose = &v
}

// GetShortId returns the ShortId field value
func (o *OrganizationConsentRequestTemplateDetails) GetShortId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShortId
}

// GetShortIdOk returns a tuple with the ShortId field value
// and a boolean to check if the value has been set.
func (o *OrganizationConsentRequestTemplateDetails) GetShortIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShortId, true
}

// SetShortId sets field value
func (o *OrganizationConsentRequestTemplateDetails) SetShortId(v string) {
	o.ShortId = v
}

// GetStatus returns the Status field value
func (o *OrganizationConsentRequestTemplateDetails) GetStatus() ConsentRequestTemplateStatus {
	if o == nil {
		var ret ConsentRequestTemplateStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OrganizationConsentRequestTemplateDetails) GetStatusOk() (*ConsentRequestTemplateStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *OrganizationConsentRequestTemplateDetails) SetStatus(v ConsentRequestTemplateStatus) {
	o.Status = v
}

// GetDataLife returns the DataLife field value if set, zero value otherwise.
func (o *OrganizationConsentRequestTemplateDetails) GetDataLife() IndividualConsentRequestTemplateDetailsDataLife {
	if o == nil || o.DataLife == nil {
		var ret IndividualConsentRequestTemplateDetailsDataLife
		return ret
	}
	return *o.DataLife
}

// GetDataLifeOk returns a tuple with the DataLife field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationConsentRequestTemplateDetails) GetDataLifeOk() (*IndividualConsentRequestTemplateDetailsDataLife, bool) {
	if o == nil || o.DataLife == nil {
		return nil, false
	}
	return o.DataLife, true
}

// HasDataLife returns a boolean if a field has been set.
func (o *OrganizationConsentRequestTemplateDetails) HasDataLife() bool {
	if o != nil && o.DataLife != nil {
		return true
	}

	return false
}

// SetDataLife gets a reference to the given IndividualConsentRequestTemplateDetailsDataLife and assigns it to the DataLife field.
func (o *OrganizationConsentRequestTemplateDetails) SetDataLife(v IndividualConsentRequestTemplateDetailsDataLife) {
	o.DataLife = &v
}

// GetRequestLife returns the RequestLife field value if set, zero value otherwise.
func (o *OrganizationConsentRequestTemplateDetails) GetRequestLife() IndividualConsentRequestTemplateDetailsRequestLife {
	if o == nil || o.RequestLife == nil {
		var ret IndividualConsentRequestTemplateDetailsRequestLife
		return ret
	}
	return *o.RequestLife
}

// GetRequestLifeOk returns a tuple with the RequestLife field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationConsentRequestTemplateDetails) GetRequestLifeOk() (*IndividualConsentRequestTemplateDetailsRequestLife, bool) {
	if o == nil || o.RequestLife == nil {
		return nil, false
	}
	return o.RequestLife, true
}

// HasRequestLife returns a boolean if a field has been set.
func (o *OrganizationConsentRequestTemplateDetails) HasRequestLife() bool {
	if o != nil && o.RequestLife != nil {
		return true
	}

	return false
}

// SetRequestLife gets a reference to the given IndividualConsentRequestTemplateDetailsRequestLife and assigns it to the RequestLife field.
func (o *OrganizationConsentRequestTemplateDetails) SetRequestLife(v IndividualConsentRequestTemplateDetailsRequestLife) {
	o.RequestLife = &v
}

// GetDataFrequency returns the DataFrequency field value if set, zero value otherwise.
func (o *OrganizationConsentRequestTemplateDetails) GetDataFrequency() IndividualConsentRequestTemplateDetailsDataFrequency {
	if o == nil || o.DataFrequency == nil {
		var ret IndividualConsentRequestTemplateDetailsDataFrequency
		return ret
	}
	return *o.DataFrequency
}

// GetDataFrequencyOk returns a tuple with the DataFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationConsentRequestTemplateDetails) GetDataFrequencyOk() (*IndividualConsentRequestTemplateDetailsDataFrequency, bool) {
	if o == nil || o.DataFrequency == nil {
		return nil, false
	}
	return o.DataFrequency, true
}

// HasDataFrequency returns a boolean if a field has been set.
func (o *OrganizationConsentRequestTemplateDetails) HasDataFrequency() bool {
	if o != nil && o.DataFrequency != nil {
		return true
	}

	return false
}

// SetDataFrequency gets a reference to the given IndividualConsentRequestTemplateDetailsDataFrequency and assigns it to the DataFrequency field.
func (o *OrganizationConsentRequestTemplateDetails) SetDataFrequency(v IndividualConsentRequestTemplateDetailsDataFrequency) {
	o.DataFrequency = &v
}

// GetIdentifiers returns the Identifiers field value if set, zero value otherwise.
func (o *OrganizationConsentRequestTemplateDetails) GetIdentifiers() []IdentityField {
	if o == nil || o.Identifiers == nil {
		var ret []IdentityField
		return ret
	}
	return o.Identifiers
}

// GetIdentifiersOk returns a tuple with the Identifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationConsentRequestTemplateDetails) GetIdentifiersOk() ([]IdentityField, bool) {
	if o == nil || o.Identifiers == nil {
		return nil, false
	}
	return o.Identifiers, true
}

// HasIdentifiers returns a boolean if a field has been set.
func (o *OrganizationConsentRequestTemplateDetails) HasIdentifiers() bool {
	if o != nil && o.Identifiers != nil {
		return true
	}

	return false
}

// SetIdentifiers gets a reference to the given []IdentityField and assigns it to the Identifiers field.
func (o *OrganizationConsentRequestTemplateDetails) SetIdentifiers(v []IdentityField) {
	o.Identifiers = v
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *OrganizationConsentRequestTemplateDetails) GetDocuments() []DocumentField {
	if o == nil || o.Documents == nil {
		var ret []DocumentField
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationConsentRequestTemplateDetails) GetDocumentsOk() ([]DocumentField, bool) {
	if o == nil || o.Documents == nil {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *OrganizationConsentRequestTemplateDetails) HasDocuments() bool {
	if o != nil && o.Documents != nil {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []DocumentField and assigns it to the Documents field.
func (o *OrganizationConsentRequestTemplateDetails) SetDocuments(v []DocumentField) {
	o.Documents = v
}

// GetFinancialAccounts returns the FinancialAccounts field value if set, zero value otherwise.
func (o *OrganizationConsentRequestTemplateDetails) GetFinancialAccounts() []FinancialAccountField {
	if o == nil || o.FinancialAccounts == nil {
		var ret []FinancialAccountField
		return ret
	}
	return o.FinancialAccounts
}

// GetFinancialAccountsOk returns a tuple with the FinancialAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationConsentRequestTemplateDetails) GetFinancialAccountsOk() ([]FinancialAccountField, bool) {
	if o == nil || o.FinancialAccounts == nil {
		return nil, false
	}
	return o.FinancialAccounts, true
}

// HasFinancialAccounts returns a boolean if a field has been set.
func (o *OrganizationConsentRequestTemplateDetails) HasFinancialAccounts() bool {
	if o != nil && o.FinancialAccounts != nil {
		return true
	}

	return false
}

// SetFinancialAccounts gets a reference to the given []FinancialAccountField and assigns it to the FinancialAccounts field.
func (o *OrganizationConsentRequestTemplateDetails) SetFinancialAccounts(v []FinancialAccountField) {
	o.FinancialAccounts = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *OrganizationConsentRequestTemplateDetails) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *OrganizationConsentRequestTemplateDetails) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *OrganizationConsentRequestTemplateDetails) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreatedAtUtc returns the CreatedAtUtc field value
func (o *OrganizationConsentRequestTemplateDetails) GetCreatedAtUtc() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAtUtc
}

// GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field value
// and a boolean to check if the value has been set.
func (o *OrganizationConsentRequestTemplateDetails) GetCreatedAtUtcOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAtUtc, true
}

// SetCreatedAtUtc sets field value
func (o *OrganizationConsentRequestTemplateDetails) SetCreatedAtUtc(v time.Time) {
	o.CreatedAtUtc = v
}

// GetApprovedAtUtc returns the ApprovedAtUtc field value if set, zero value otherwise.
func (o *OrganizationConsentRequestTemplateDetails) GetApprovedAtUtc() time.Time {
	if o == nil || o.ApprovedAtUtc == nil {
		var ret time.Time
		return ret
	}
	return *o.ApprovedAtUtc
}

// GetApprovedAtUtcOk returns a tuple with the ApprovedAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationConsentRequestTemplateDetails) GetApprovedAtUtcOk() (*time.Time, bool) {
	if o == nil || o.ApprovedAtUtc == nil {
		return nil, false
	}
	return o.ApprovedAtUtc, true
}

// HasApprovedAtUtc returns a boolean if a field has been set.
func (o *OrganizationConsentRequestTemplateDetails) HasApprovedAtUtc() bool {
	if o != nil && o.ApprovedAtUtc != nil {
		return true
	}

	return false
}

// SetApprovedAtUtc gets a reference to the given time.Time and assigns it to the ApprovedAtUtc field.
func (o *OrganizationConsentRequestTemplateDetails) SetApprovedAtUtc(v time.Time) {
	o.ApprovedAtUtc = &v
}

// GetPublishedAtUtc returns the PublishedAtUtc field value if set, zero value otherwise.
func (o *OrganizationConsentRequestTemplateDetails) GetPublishedAtUtc() time.Time {
	if o == nil || o.PublishedAtUtc == nil {
		var ret time.Time
		return ret
	}
	return *o.PublishedAtUtc
}

// GetPublishedAtUtcOk returns a tuple with the PublishedAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationConsentRequestTemplateDetails) GetPublishedAtUtcOk() (*time.Time, bool) {
	if o == nil || o.PublishedAtUtc == nil {
		return nil, false
	}
	return o.PublishedAtUtc, true
}

// HasPublishedAtUtc returns a boolean if a field has been set.
func (o *OrganizationConsentRequestTemplateDetails) HasPublishedAtUtc() bool {
	if o != nil && o.PublishedAtUtc != nil {
		return true
	}

	return false
}

// SetPublishedAtUtc gets a reference to the given time.Time and assigns it to the PublishedAtUtc field.
func (o *OrganizationConsentRequestTemplateDetails) SetPublishedAtUtc(v time.Time) {
	o.PublishedAtUtc = &v
}

// GetRejectedAtUtc returns the RejectedAtUtc field value if set, zero value otherwise.
func (o *OrganizationConsentRequestTemplateDetails) GetRejectedAtUtc() time.Time {
	if o == nil || o.RejectedAtUtc == nil {
		var ret time.Time
		return ret
	}
	return *o.RejectedAtUtc
}

// GetRejectedAtUtcOk returns a tuple with the RejectedAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationConsentRequestTemplateDetails) GetRejectedAtUtcOk() (*time.Time, bool) {
	if o == nil || o.RejectedAtUtc == nil {
		return nil, false
	}
	return o.RejectedAtUtc, true
}

// HasRejectedAtUtc returns a boolean if a field has been set.
func (o *OrganizationConsentRequestTemplateDetails) HasRejectedAtUtc() bool {
	if o != nil && o.RejectedAtUtc != nil {
		return true
	}

	return false
}

// SetRejectedAtUtc gets a reference to the given time.Time and assigns it to the RejectedAtUtc field.
func (o *OrganizationConsentRequestTemplateDetails) SetRejectedAtUtc(v time.Time) {
	o.RejectedAtUtc = &v
}

// GetRejectionReason returns the RejectionReason field value if set, zero value otherwise.
func (o *OrganizationConsentRequestTemplateDetails) GetRejectionReason() string {
	if o == nil || o.RejectionReason == nil {
		var ret string
		return ret
	}
	return *o.RejectionReason
}

// GetRejectionReasonOk returns a tuple with the RejectionReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationConsentRequestTemplateDetails) GetRejectionReasonOk() (*string, bool) {
	if o == nil || o.RejectionReason == nil {
		return nil, false
	}
	return o.RejectionReason, true
}

// HasRejectionReason returns a boolean if a field has been set.
func (o *OrganizationConsentRequestTemplateDetails) HasRejectionReason() bool {
	if o != nil && o.RejectionReason != nil {
		return true
	}

	return false
}

// SetRejectionReason gets a reference to the given string and assigns it to the RejectionReason field.
func (o *OrganizationConsentRequestTemplateDetails) SetRejectionReason(v string) {
	o.RejectionReason = &v
}

func (o OrganizationConsentRequestTemplateDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
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
		toSerialize["shortId"] = o.ShortId
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.DataLife != nil {
		toSerialize["dataLife"] = o.DataLife
	}
	if o.RequestLife != nil {
		toSerialize["requestLife"] = o.RequestLife
	}
	if o.DataFrequency != nil {
		toSerialize["dataFrequency"] = o.DataFrequency
	}
	if o.Identifiers != nil {
		toSerialize["identifiers"] = o.Identifiers
	}
	if o.Documents != nil {
		toSerialize["documents"] = o.Documents
	}
	if o.FinancialAccounts != nil {
		toSerialize["financialAccounts"] = o.FinancialAccounts
	}
	if true {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if true {
		toSerialize["createdAtUtc"] = o.CreatedAtUtc
	}
	if o.ApprovedAtUtc != nil {
		toSerialize["approvedAtUtc"] = o.ApprovedAtUtc
	}
	if o.PublishedAtUtc != nil {
		toSerialize["publishedAtUtc"] = o.PublishedAtUtc
	}
	if o.RejectedAtUtc != nil {
		toSerialize["rejectedAtUtc"] = o.RejectedAtUtc
	}
	if o.RejectionReason != nil {
		toSerialize["rejectionReason"] = o.RejectionReason
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationConsentRequestTemplateDetails struct {
	value *OrganizationConsentRequestTemplateDetails
	isSet bool
}

func (v NullableOrganizationConsentRequestTemplateDetails) Get() *OrganizationConsentRequestTemplateDetails {
	return v.value
}

func (v *NullableOrganizationConsentRequestTemplateDetails) Set(val *OrganizationConsentRequestTemplateDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationConsentRequestTemplateDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationConsentRequestTemplateDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationConsentRequestTemplateDetails(val *OrganizationConsentRequestTemplateDetails) *NullableOrganizationConsentRequestTemplateDetails {
	return &NullableOrganizationConsentRequestTemplateDetails{value: val, isSet: true}
}

func (v NullableOrganizationConsentRequestTemplateDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationConsentRequestTemplateDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



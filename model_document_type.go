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

// DocumentType Issuable Document Type details.
type DocumentType struct {
	// Document Type Identifier.
	Id string `json:"id"`
	CategoryType DocumentCategoryType `json:"categoryType"`
	SubCategoryType DocumentSubCategoryType `json:"subCategoryType"`
	// Document Type Name. eg: Driving License.
	Name string `json:"name"`
	// Document Type Unique Slug. eg: \"in.gov.gj.transport.dl\".
	Slug string `json:"slug"`
	// Document Type description. eg: Gujarat State Driving License.
	Description NullableString `json:"description,omitempty"`
	// Logo URL of document type.
	LogoUrl string `json:"logoUrl"`
	// Document search repository service name.
	SearchServiceName NullableString `json:"searchServiceName,omitempty"`
	// Document repository service name.
	RepositoryServiceName NullableString `json:"repositoryServiceName,omitempty"`
	// Supported entity types. eg: Individual, Organization.
	SupportedEntityTypes []SupportedEntityType `json:"supportedEntityTypes"`
	// Name of the document type creator.
	AddedBy string `json:"addedBy"`
	// Payable amount if document is chargeable. eg: 10.25.
	PayableAmount NullableFloat64 `json:"payableAmount,omitempty"`
	// Payable amount currency. eg: INR, USD etc.,.
	PayableAmountCurrency NullableString `json:"payableAmountCurrency,omitempty"`
	// DateTime of approval in UTC timezone.
	ApprovedAtUtc NullableTime `json:"approvedAtUtc,omitempty"`
	// Document type approval status.
	Approved bool `json:"approved"`
}

// NewDocumentType instantiates a new DocumentType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentType(id string, categoryType DocumentCategoryType, subCategoryType DocumentSubCategoryType, name string, slug string, logoUrl string, supportedEntityTypes []SupportedEntityType, addedBy string, approved bool) *DocumentType {
	this := DocumentType{}
	this.Id = id
	this.CategoryType = categoryType
	this.SubCategoryType = subCategoryType
	this.Name = name
	this.Slug = slug
	this.LogoUrl = logoUrl
	this.SupportedEntityTypes = supportedEntityTypes
	this.AddedBy = addedBy
	this.Approved = approved
	return &this
}

// NewDocumentTypeWithDefaults instantiates a new DocumentType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentTypeWithDefaults() *DocumentType {
	this := DocumentType{}
	return &this
}

// GetId returns the Id field value
func (o *DocumentType) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DocumentType) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DocumentType) SetId(v string) {
	o.Id = v
}

// GetCategoryType returns the CategoryType field value
func (o *DocumentType) GetCategoryType() DocumentCategoryType {
	if o == nil {
		var ret DocumentCategoryType
		return ret
	}

	return o.CategoryType
}

// GetCategoryTypeOk returns a tuple with the CategoryType field value
// and a boolean to check if the value has been set.
func (o *DocumentType) GetCategoryTypeOk() (*DocumentCategoryType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CategoryType, true
}

// SetCategoryType sets field value
func (o *DocumentType) SetCategoryType(v DocumentCategoryType) {
	o.CategoryType = v
}

// GetSubCategoryType returns the SubCategoryType field value
func (o *DocumentType) GetSubCategoryType() DocumentSubCategoryType {
	if o == nil {
		var ret DocumentSubCategoryType
		return ret
	}

	return o.SubCategoryType
}

// GetSubCategoryTypeOk returns a tuple with the SubCategoryType field value
// and a boolean to check if the value has been set.
func (o *DocumentType) GetSubCategoryTypeOk() (*DocumentSubCategoryType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SubCategoryType, true
}

// SetSubCategoryType sets field value
func (o *DocumentType) SetSubCategoryType(v DocumentSubCategoryType) {
	o.SubCategoryType = v
}

// GetName returns the Name field value
func (o *DocumentType) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DocumentType) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DocumentType) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *DocumentType) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *DocumentType) GetSlugOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *DocumentType) SetSlug(v string) {
	o.Slug = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentType) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentType) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *DocumentType) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *DocumentType) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *DocumentType) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *DocumentType) UnsetDescription() {
	o.Description.Unset()
}

// GetLogoUrl returns the LogoUrl field value
func (o *DocumentType) GetLogoUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value
// and a boolean to check if the value has been set.
func (o *DocumentType) GetLogoUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LogoUrl, true
}

// SetLogoUrl sets field value
func (o *DocumentType) SetLogoUrl(v string) {
	o.LogoUrl = v
}

// GetSearchServiceName returns the SearchServiceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentType) GetSearchServiceName() string {
	if o == nil || o.SearchServiceName.Get() == nil {
		var ret string
		return ret
	}
	return *o.SearchServiceName.Get()
}

// GetSearchServiceNameOk returns a tuple with the SearchServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentType) GetSearchServiceNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SearchServiceName.Get(), o.SearchServiceName.IsSet()
}

// HasSearchServiceName returns a boolean if a field has been set.
func (o *DocumentType) HasSearchServiceName() bool {
	if o != nil && o.SearchServiceName.IsSet() {
		return true
	}

	return false
}

// SetSearchServiceName gets a reference to the given NullableString and assigns it to the SearchServiceName field.
func (o *DocumentType) SetSearchServiceName(v string) {
	o.SearchServiceName.Set(&v)
}
// SetSearchServiceNameNil sets the value for SearchServiceName to be an explicit nil
func (o *DocumentType) SetSearchServiceNameNil() {
	o.SearchServiceName.Set(nil)
}

// UnsetSearchServiceName ensures that no value is present for SearchServiceName, not even an explicit nil
func (o *DocumentType) UnsetSearchServiceName() {
	o.SearchServiceName.Unset()
}

// GetRepositoryServiceName returns the RepositoryServiceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentType) GetRepositoryServiceName() string {
	if o == nil || o.RepositoryServiceName.Get() == nil {
		var ret string
		return ret
	}
	return *o.RepositoryServiceName.Get()
}

// GetRepositoryServiceNameOk returns a tuple with the RepositoryServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentType) GetRepositoryServiceNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RepositoryServiceName.Get(), o.RepositoryServiceName.IsSet()
}

// HasRepositoryServiceName returns a boolean if a field has been set.
func (o *DocumentType) HasRepositoryServiceName() bool {
	if o != nil && o.RepositoryServiceName.IsSet() {
		return true
	}

	return false
}

// SetRepositoryServiceName gets a reference to the given NullableString and assigns it to the RepositoryServiceName field.
func (o *DocumentType) SetRepositoryServiceName(v string) {
	o.RepositoryServiceName.Set(&v)
}
// SetRepositoryServiceNameNil sets the value for RepositoryServiceName to be an explicit nil
func (o *DocumentType) SetRepositoryServiceNameNil() {
	o.RepositoryServiceName.Set(nil)
}

// UnsetRepositoryServiceName ensures that no value is present for RepositoryServiceName, not even an explicit nil
func (o *DocumentType) UnsetRepositoryServiceName() {
	o.RepositoryServiceName.Unset()
}

// GetSupportedEntityTypes returns the SupportedEntityTypes field value
func (o *DocumentType) GetSupportedEntityTypes() []SupportedEntityType {
	if o == nil {
		var ret []SupportedEntityType
		return ret
	}

	return o.SupportedEntityTypes
}

// GetSupportedEntityTypesOk returns a tuple with the SupportedEntityTypes field value
// and a boolean to check if the value has been set.
func (o *DocumentType) GetSupportedEntityTypesOk() ([]SupportedEntityType, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SupportedEntityTypes, true
}

// SetSupportedEntityTypes sets field value
func (o *DocumentType) SetSupportedEntityTypes(v []SupportedEntityType) {
	o.SupportedEntityTypes = v
}

// GetAddedBy returns the AddedBy field value
func (o *DocumentType) GetAddedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddedBy
}

// GetAddedByOk returns a tuple with the AddedBy field value
// and a boolean to check if the value has been set.
func (o *DocumentType) GetAddedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AddedBy, true
}

// SetAddedBy sets field value
func (o *DocumentType) SetAddedBy(v string) {
	o.AddedBy = v
}

// GetPayableAmount returns the PayableAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentType) GetPayableAmount() float64 {
	if o == nil || o.PayableAmount.Get() == nil {
		var ret float64
		return ret
	}
	return *o.PayableAmount.Get()
}

// GetPayableAmountOk returns a tuple with the PayableAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentType) GetPayableAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayableAmount.Get(), o.PayableAmount.IsSet()
}

// HasPayableAmount returns a boolean if a field has been set.
func (o *DocumentType) HasPayableAmount() bool {
	if o != nil && o.PayableAmount.IsSet() {
		return true
	}

	return false
}

// SetPayableAmount gets a reference to the given NullableFloat64 and assigns it to the PayableAmount field.
func (o *DocumentType) SetPayableAmount(v float64) {
	o.PayableAmount.Set(&v)
}
// SetPayableAmountNil sets the value for PayableAmount to be an explicit nil
func (o *DocumentType) SetPayableAmountNil() {
	o.PayableAmount.Set(nil)
}

// UnsetPayableAmount ensures that no value is present for PayableAmount, not even an explicit nil
func (o *DocumentType) UnsetPayableAmount() {
	o.PayableAmount.Unset()
}

// GetPayableAmountCurrency returns the PayableAmountCurrency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentType) GetPayableAmountCurrency() string {
	if o == nil || o.PayableAmountCurrency.Get() == nil {
		var ret string
		return ret
	}
	return *o.PayableAmountCurrency.Get()
}

// GetPayableAmountCurrencyOk returns a tuple with the PayableAmountCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentType) GetPayableAmountCurrencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayableAmountCurrency.Get(), o.PayableAmountCurrency.IsSet()
}

// HasPayableAmountCurrency returns a boolean if a field has been set.
func (o *DocumentType) HasPayableAmountCurrency() bool {
	if o != nil && o.PayableAmountCurrency.IsSet() {
		return true
	}

	return false
}

// SetPayableAmountCurrency gets a reference to the given NullableString and assigns it to the PayableAmountCurrency field.
func (o *DocumentType) SetPayableAmountCurrency(v string) {
	o.PayableAmountCurrency.Set(&v)
}
// SetPayableAmountCurrencyNil sets the value for PayableAmountCurrency to be an explicit nil
func (o *DocumentType) SetPayableAmountCurrencyNil() {
	o.PayableAmountCurrency.Set(nil)
}

// UnsetPayableAmountCurrency ensures that no value is present for PayableAmountCurrency, not even an explicit nil
func (o *DocumentType) UnsetPayableAmountCurrency() {
	o.PayableAmountCurrency.Unset()
}

// GetApprovedAtUtc returns the ApprovedAtUtc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentType) GetApprovedAtUtc() time.Time {
	if o == nil || o.ApprovedAtUtc.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ApprovedAtUtc.Get()
}

// GetApprovedAtUtcOk returns a tuple with the ApprovedAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentType) GetApprovedAtUtcOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ApprovedAtUtc.Get(), o.ApprovedAtUtc.IsSet()
}

// HasApprovedAtUtc returns a boolean if a field has been set.
func (o *DocumentType) HasApprovedAtUtc() bool {
	if o != nil && o.ApprovedAtUtc.IsSet() {
		return true
	}

	return false
}

// SetApprovedAtUtc gets a reference to the given NullableTime and assigns it to the ApprovedAtUtc field.
func (o *DocumentType) SetApprovedAtUtc(v time.Time) {
	o.ApprovedAtUtc.Set(&v)
}
// SetApprovedAtUtcNil sets the value for ApprovedAtUtc to be an explicit nil
func (o *DocumentType) SetApprovedAtUtcNil() {
	o.ApprovedAtUtc.Set(nil)
}

// UnsetApprovedAtUtc ensures that no value is present for ApprovedAtUtc, not even an explicit nil
func (o *DocumentType) UnsetApprovedAtUtc() {
	o.ApprovedAtUtc.Unset()
}

// GetApproved returns the Approved field value
func (o *DocumentType) GetApproved() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Approved
}

// GetApprovedOk returns a tuple with the Approved field value
// and a boolean to check if the value has been set.
func (o *DocumentType) GetApprovedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Approved, true
}

// SetApproved sets field value
func (o *DocumentType) SetApproved(v bool) {
	o.Approved = v
}

func (o DocumentType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["categoryType"] = o.CategoryType
	}
	if true {
		toSerialize["subCategoryType"] = o.SubCategoryType
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["slug"] = o.Slug
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if true {
		toSerialize["logoUrl"] = o.LogoUrl
	}
	if o.SearchServiceName.IsSet() {
		toSerialize["searchServiceName"] = o.SearchServiceName.Get()
	}
	if o.RepositoryServiceName.IsSet() {
		toSerialize["repositoryServiceName"] = o.RepositoryServiceName.Get()
	}
	if true {
		toSerialize["supportedEntityTypes"] = o.SupportedEntityTypes
	}
	if true {
		toSerialize["addedBy"] = o.AddedBy
	}
	if o.PayableAmount.IsSet() {
		toSerialize["payableAmount"] = o.PayableAmount.Get()
	}
	if o.PayableAmountCurrency.IsSet() {
		toSerialize["payableAmountCurrency"] = o.PayableAmountCurrency.Get()
	}
	if o.ApprovedAtUtc.IsSet() {
		toSerialize["approvedAtUtc"] = o.ApprovedAtUtc.Get()
	}
	if true {
		toSerialize["approved"] = o.Approved
	}
	return json.Marshal(toSerialize)
}

type NullableDocumentType struct {
	value *DocumentType
	isSet bool
}

func (v NullableDocumentType) Get() *DocumentType {
	return v.value
}

func (v *NullableDocumentType) Set(val *DocumentType) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentType) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentType(val *DocumentType) *NullableDocumentType {
	return &NullableDocumentType{value: val, isSet: true}
}

func (v NullableDocumentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// OrganizationKyoDocument struct for OrganizationKyoDocument
type OrganizationKyoDocument struct {
	Id *string `json:"id,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty"`
	DocumentName NullableString `json:"documentName,omitempty"`
	StorageUrl NullableString `json:"storageUrl,omitempty"`
	UploadedAtUtc *time.Time `json:"uploadedAtUtc,omitempty"`
	VerifiedBy NullableString `json:"verifiedBy,omitempty"`
	VerifiedAtUtc NullableTime `json:"verifiedAtUtc,omitempty"`
	DeletedBy NullableString `json:"deletedBy,omitempty"`
	DeletedAtUtc NullableTime `json:"deletedAtUtc,omitempty"`
	Organization *Organization `json:"organization,omitempty"`
	VerifiedByUser *ApplicationUser `json:"verifiedByUser,omitempty"`
	DeletedByUser *ApplicationUser `json:"deletedByUser,omitempty"`
	Rejection *Rejection `json:"rejection,omitempty"`
}

// NewOrganizationKyoDocument instantiates a new OrganizationKyoDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationKyoDocument() *OrganizationKyoDocument {
	this := OrganizationKyoDocument{}
	return &this
}

// NewOrganizationKyoDocumentWithDefaults instantiates a new OrganizationKyoDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationKyoDocumentWithDefaults() *OrganizationKyoDocument {
	this := OrganizationKyoDocument{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationKyoDocument) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationKyoDocument) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationKyoDocument) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationKyoDocument) SetId(v string) {
	o.Id = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *OrganizationKyoDocument) GetOrganizationId() string {
	if o == nil || o.OrganizationId == nil {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationKyoDocument) GetOrganizationIdOk() (*string, bool) {
	if o == nil || o.OrganizationId == nil {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *OrganizationKyoDocument) HasOrganizationId() bool {
	if o != nil && o.OrganizationId != nil {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *OrganizationKyoDocument) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetDocumentName returns the DocumentName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationKyoDocument) GetDocumentName() string {
	if o == nil || o.DocumentName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DocumentName.Get()
}

// GetDocumentNameOk returns a tuple with the DocumentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationKyoDocument) GetDocumentNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DocumentName.Get(), o.DocumentName.IsSet()
}

// HasDocumentName returns a boolean if a field has been set.
func (o *OrganizationKyoDocument) HasDocumentName() bool {
	if o != nil && o.DocumentName.IsSet() {
		return true
	}

	return false
}

// SetDocumentName gets a reference to the given NullableString and assigns it to the DocumentName field.
func (o *OrganizationKyoDocument) SetDocumentName(v string) {
	o.DocumentName.Set(&v)
}
// SetDocumentNameNil sets the value for DocumentName to be an explicit nil
func (o *OrganizationKyoDocument) SetDocumentNameNil() {
	o.DocumentName.Set(nil)
}

// UnsetDocumentName ensures that no value is present for DocumentName, not even an explicit nil
func (o *OrganizationKyoDocument) UnsetDocumentName() {
	o.DocumentName.Unset()
}

// GetStorageUrl returns the StorageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationKyoDocument) GetStorageUrl() string {
	if o == nil || o.StorageUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.StorageUrl.Get()
}

// GetStorageUrlOk returns a tuple with the StorageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationKyoDocument) GetStorageUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageUrl.Get(), o.StorageUrl.IsSet()
}

// HasStorageUrl returns a boolean if a field has been set.
func (o *OrganizationKyoDocument) HasStorageUrl() bool {
	if o != nil && o.StorageUrl.IsSet() {
		return true
	}

	return false
}

// SetStorageUrl gets a reference to the given NullableString and assigns it to the StorageUrl field.
func (o *OrganizationKyoDocument) SetStorageUrl(v string) {
	o.StorageUrl.Set(&v)
}
// SetStorageUrlNil sets the value for StorageUrl to be an explicit nil
func (o *OrganizationKyoDocument) SetStorageUrlNil() {
	o.StorageUrl.Set(nil)
}

// UnsetStorageUrl ensures that no value is present for StorageUrl, not even an explicit nil
func (o *OrganizationKyoDocument) UnsetStorageUrl() {
	o.StorageUrl.Unset()
}

// GetUploadedAtUtc returns the UploadedAtUtc field value if set, zero value otherwise.
func (o *OrganizationKyoDocument) GetUploadedAtUtc() time.Time {
	if o == nil || o.UploadedAtUtc == nil {
		var ret time.Time
		return ret
	}
	return *o.UploadedAtUtc
}

// GetUploadedAtUtcOk returns a tuple with the UploadedAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationKyoDocument) GetUploadedAtUtcOk() (*time.Time, bool) {
	if o == nil || o.UploadedAtUtc == nil {
		return nil, false
	}
	return o.UploadedAtUtc, true
}

// HasUploadedAtUtc returns a boolean if a field has been set.
func (o *OrganizationKyoDocument) HasUploadedAtUtc() bool {
	if o != nil && o.UploadedAtUtc != nil {
		return true
	}

	return false
}

// SetUploadedAtUtc gets a reference to the given time.Time and assigns it to the UploadedAtUtc field.
func (o *OrganizationKyoDocument) SetUploadedAtUtc(v time.Time) {
	o.UploadedAtUtc = &v
}

// GetVerifiedBy returns the VerifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationKyoDocument) GetVerifiedBy() string {
	if o == nil || o.VerifiedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.VerifiedBy.Get()
}

// GetVerifiedByOk returns a tuple with the VerifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationKyoDocument) GetVerifiedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VerifiedBy.Get(), o.VerifiedBy.IsSet()
}

// HasVerifiedBy returns a boolean if a field has been set.
func (o *OrganizationKyoDocument) HasVerifiedBy() bool {
	if o != nil && o.VerifiedBy.IsSet() {
		return true
	}

	return false
}

// SetVerifiedBy gets a reference to the given NullableString and assigns it to the VerifiedBy field.
func (o *OrganizationKyoDocument) SetVerifiedBy(v string) {
	o.VerifiedBy.Set(&v)
}
// SetVerifiedByNil sets the value for VerifiedBy to be an explicit nil
func (o *OrganizationKyoDocument) SetVerifiedByNil() {
	o.VerifiedBy.Set(nil)
}

// UnsetVerifiedBy ensures that no value is present for VerifiedBy, not even an explicit nil
func (o *OrganizationKyoDocument) UnsetVerifiedBy() {
	o.VerifiedBy.Unset()
}

// GetVerifiedAtUtc returns the VerifiedAtUtc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationKyoDocument) GetVerifiedAtUtc() time.Time {
	if o == nil || o.VerifiedAtUtc.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.VerifiedAtUtc.Get()
}

// GetVerifiedAtUtcOk returns a tuple with the VerifiedAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationKyoDocument) GetVerifiedAtUtcOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VerifiedAtUtc.Get(), o.VerifiedAtUtc.IsSet()
}

// HasVerifiedAtUtc returns a boolean if a field has been set.
func (o *OrganizationKyoDocument) HasVerifiedAtUtc() bool {
	if o != nil && o.VerifiedAtUtc.IsSet() {
		return true
	}

	return false
}

// SetVerifiedAtUtc gets a reference to the given NullableTime and assigns it to the VerifiedAtUtc field.
func (o *OrganizationKyoDocument) SetVerifiedAtUtc(v time.Time) {
	o.VerifiedAtUtc.Set(&v)
}
// SetVerifiedAtUtcNil sets the value for VerifiedAtUtc to be an explicit nil
func (o *OrganizationKyoDocument) SetVerifiedAtUtcNil() {
	o.VerifiedAtUtc.Set(nil)
}

// UnsetVerifiedAtUtc ensures that no value is present for VerifiedAtUtc, not even an explicit nil
func (o *OrganizationKyoDocument) UnsetVerifiedAtUtc() {
	o.VerifiedAtUtc.Unset()
}

// GetDeletedBy returns the DeletedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationKyoDocument) GetDeletedBy() string {
	if o == nil || o.DeletedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeletedBy.Get()
}

// GetDeletedByOk returns a tuple with the DeletedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationKyoDocument) GetDeletedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletedBy.Get(), o.DeletedBy.IsSet()
}

// HasDeletedBy returns a boolean if a field has been set.
func (o *OrganizationKyoDocument) HasDeletedBy() bool {
	if o != nil && o.DeletedBy.IsSet() {
		return true
	}

	return false
}

// SetDeletedBy gets a reference to the given NullableString and assigns it to the DeletedBy field.
func (o *OrganizationKyoDocument) SetDeletedBy(v string) {
	o.DeletedBy.Set(&v)
}
// SetDeletedByNil sets the value for DeletedBy to be an explicit nil
func (o *OrganizationKyoDocument) SetDeletedByNil() {
	o.DeletedBy.Set(nil)
}

// UnsetDeletedBy ensures that no value is present for DeletedBy, not even an explicit nil
func (o *OrganizationKyoDocument) UnsetDeletedBy() {
	o.DeletedBy.Unset()
}

// GetDeletedAtUtc returns the DeletedAtUtc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationKyoDocument) GetDeletedAtUtc() time.Time {
	if o == nil || o.DeletedAtUtc.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAtUtc.Get()
}

// GetDeletedAtUtcOk returns a tuple with the DeletedAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationKyoDocument) GetDeletedAtUtcOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletedAtUtc.Get(), o.DeletedAtUtc.IsSet()
}

// HasDeletedAtUtc returns a boolean if a field has been set.
func (o *OrganizationKyoDocument) HasDeletedAtUtc() bool {
	if o != nil && o.DeletedAtUtc.IsSet() {
		return true
	}

	return false
}

// SetDeletedAtUtc gets a reference to the given NullableTime and assigns it to the DeletedAtUtc field.
func (o *OrganizationKyoDocument) SetDeletedAtUtc(v time.Time) {
	o.DeletedAtUtc.Set(&v)
}
// SetDeletedAtUtcNil sets the value for DeletedAtUtc to be an explicit nil
func (o *OrganizationKyoDocument) SetDeletedAtUtcNil() {
	o.DeletedAtUtc.Set(nil)
}

// UnsetDeletedAtUtc ensures that no value is present for DeletedAtUtc, not even an explicit nil
func (o *OrganizationKyoDocument) UnsetDeletedAtUtc() {
	o.DeletedAtUtc.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *OrganizationKyoDocument) GetOrganization() Organization {
	if o == nil || o.Organization == nil {
		var ret Organization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationKyoDocument) GetOrganizationOk() (*Organization, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *OrganizationKyoDocument) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Organization and assigns it to the Organization field.
func (o *OrganizationKyoDocument) SetOrganization(v Organization) {
	o.Organization = &v
}

// GetVerifiedByUser returns the VerifiedByUser field value if set, zero value otherwise.
func (o *OrganizationKyoDocument) GetVerifiedByUser() ApplicationUser {
	if o == nil || o.VerifiedByUser == nil {
		var ret ApplicationUser
		return ret
	}
	return *o.VerifiedByUser
}

// GetVerifiedByUserOk returns a tuple with the VerifiedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationKyoDocument) GetVerifiedByUserOk() (*ApplicationUser, bool) {
	if o == nil || o.VerifiedByUser == nil {
		return nil, false
	}
	return o.VerifiedByUser, true
}

// HasVerifiedByUser returns a boolean if a field has been set.
func (o *OrganizationKyoDocument) HasVerifiedByUser() bool {
	if o != nil && o.VerifiedByUser != nil {
		return true
	}

	return false
}

// SetVerifiedByUser gets a reference to the given ApplicationUser and assigns it to the VerifiedByUser field.
func (o *OrganizationKyoDocument) SetVerifiedByUser(v ApplicationUser) {
	o.VerifiedByUser = &v
}

// GetDeletedByUser returns the DeletedByUser field value if set, zero value otherwise.
func (o *OrganizationKyoDocument) GetDeletedByUser() ApplicationUser {
	if o == nil || o.DeletedByUser == nil {
		var ret ApplicationUser
		return ret
	}
	return *o.DeletedByUser
}

// GetDeletedByUserOk returns a tuple with the DeletedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationKyoDocument) GetDeletedByUserOk() (*ApplicationUser, bool) {
	if o == nil || o.DeletedByUser == nil {
		return nil, false
	}
	return o.DeletedByUser, true
}

// HasDeletedByUser returns a boolean if a field has been set.
func (o *OrganizationKyoDocument) HasDeletedByUser() bool {
	if o != nil && o.DeletedByUser != nil {
		return true
	}

	return false
}

// SetDeletedByUser gets a reference to the given ApplicationUser and assigns it to the DeletedByUser field.
func (o *OrganizationKyoDocument) SetDeletedByUser(v ApplicationUser) {
	o.DeletedByUser = &v
}

// GetRejection returns the Rejection field value if set, zero value otherwise.
func (o *OrganizationKyoDocument) GetRejection() Rejection {
	if o == nil || o.Rejection == nil {
		var ret Rejection
		return ret
	}
	return *o.Rejection
}

// GetRejectionOk returns a tuple with the Rejection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationKyoDocument) GetRejectionOk() (*Rejection, bool) {
	if o == nil || o.Rejection == nil {
		return nil, false
	}
	return o.Rejection, true
}

// HasRejection returns a boolean if a field has been set.
func (o *OrganizationKyoDocument) HasRejection() bool {
	if o != nil && o.Rejection != nil {
		return true
	}

	return false
}

// SetRejection gets a reference to the given Rejection and assigns it to the Rejection field.
func (o *OrganizationKyoDocument) SetRejection(v Rejection) {
	o.Rejection = &v
}

func (o OrganizationKyoDocument) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.OrganizationId != nil {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if o.DocumentName.IsSet() {
		toSerialize["documentName"] = o.DocumentName.Get()
	}
	if o.StorageUrl.IsSet() {
		toSerialize["storageUrl"] = o.StorageUrl.Get()
	}
	if o.UploadedAtUtc != nil {
		toSerialize["uploadedAtUtc"] = o.UploadedAtUtc
	}
	if o.VerifiedBy.IsSet() {
		toSerialize["verifiedBy"] = o.VerifiedBy.Get()
	}
	if o.VerifiedAtUtc.IsSet() {
		toSerialize["verifiedAtUtc"] = o.VerifiedAtUtc.Get()
	}
	if o.DeletedBy.IsSet() {
		toSerialize["deletedBy"] = o.DeletedBy.Get()
	}
	if o.DeletedAtUtc.IsSet() {
		toSerialize["deletedAtUtc"] = o.DeletedAtUtc.Get()
	}
	if o.Organization != nil {
		toSerialize["organization"] = o.Organization
	}
	if o.VerifiedByUser != nil {
		toSerialize["verifiedByUser"] = o.VerifiedByUser
	}
	if o.DeletedByUser != nil {
		toSerialize["deletedByUser"] = o.DeletedByUser
	}
	if o.Rejection != nil {
		toSerialize["rejection"] = o.Rejection
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationKyoDocument struct {
	value *OrganizationKyoDocument
	isSet bool
}

func (v NullableOrganizationKyoDocument) Get() *OrganizationKyoDocument {
	return v.value
}

func (v *NullableOrganizationKyoDocument) Set(val *OrganizationKyoDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationKyoDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationKyoDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationKyoDocument(val *OrganizationKyoDocument) *NullableOrganizationKyoDocument {
	return &NullableOrganizationKyoDocument{value: val, isSet: true}
}

func (v NullableOrganizationKyoDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationKyoDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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
)

// OrganizationDocumentDownloadDto struct for OrganizationDocumentDownloadDto
type OrganizationDocumentDownloadDto struct {
	Id *string `json:"id,omitempty"`
	StorageUrl NullableString `json:"storageUrl,omitempty"`
	OwnerId *string `json:"ownerId,omitempty"`
}

// NewOrganizationDocumentDownloadDto instantiates a new OrganizationDocumentDownloadDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationDocumentDownloadDto() *OrganizationDocumentDownloadDto {
	this := OrganizationDocumentDownloadDto{}
	return &this
}

// NewOrganizationDocumentDownloadDtoWithDefaults instantiates a new OrganizationDocumentDownloadDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationDocumentDownloadDtoWithDefaults() *OrganizationDocumentDownloadDto {
	this := OrganizationDocumentDownloadDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationDocumentDownloadDto) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDocumentDownloadDto) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationDocumentDownloadDto) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationDocumentDownloadDto) SetId(v string) {
	o.Id = &v
}

// GetStorageUrl returns the StorageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationDocumentDownloadDto) GetStorageUrl() string {
	if o == nil || o.StorageUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.StorageUrl.Get()
}

// GetStorageUrlOk returns a tuple with the StorageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationDocumentDownloadDto) GetStorageUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageUrl.Get(), o.StorageUrl.IsSet()
}

// HasStorageUrl returns a boolean if a field has been set.
func (o *OrganizationDocumentDownloadDto) HasStorageUrl() bool {
	if o != nil && o.StorageUrl.IsSet() {
		return true
	}

	return false
}

// SetStorageUrl gets a reference to the given NullableString and assigns it to the StorageUrl field.
func (o *OrganizationDocumentDownloadDto) SetStorageUrl(v string) {
	o.StorageUrl.Set(&v)
}
// SetStorageUrlNil sets the value for StorageUrl to be an explicit nil
func (o *OrganizationDocumentDownloadDto) SetStorageUrlNil() {
	o.StorageUrl.Set(nil)
}

// UnsetStorageUrl ensures that no value is present for StorageUrl, not even an explicit nil
func (o *OrganizationDocumentDownloadDto) UnsetStorageUrl() {
	o.StorageUrl.Unset()
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *OrganizationDocumentDownloadDto) GetOwnerId() string {
	if o == nil || o.OwnerId == nil {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDocumentDownloadDto) GetOwnerIdOk() (*string, bool) {
	if o == nil || o.OwnerId == nil {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *OrganizationDocumentDownloadDto) HasOwnerId() bool {
	if o != nil && o.OwnerId != nil {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *OrganizationDocumentDownloadDto) SetOwnerId(v string) {
	o.OwnerId = &v
}

func (o OrganizationDocumentDownloadDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.StorageUrl.IsSet() {
		toSerialize["storageUrl"] = o.StorageUrl.Get()
	}
	if o.OwnerId != nil {
		toSerialize["ownerId"] = o.OwnerId
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationDocumentDownloadDto struct {
	value *OrganizationDocumentDownloadDto
	isSet bool
}

func (v NullableOrganizationDocumentDownloadDto) Get() *OrganizationDocumentDownloadDto {
	return v.value
}

func (v *NullableOrganizationDocumentDownloadDto) Set(val *OrganizationDocumentDownloadDto) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationDocumentDownloadDto) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationDocumentDownloadDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationDocumentDownloadDto(val *OrganizationDocumentDownloadDto) *NullableOrganizationDocumentDownloadDto {
	return &NullableOrganizationDocumentDownloadDto{value: val, isSet: true}
}

func (v NullableOrganizationDocumentDownloadDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationDocumentDownloadDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


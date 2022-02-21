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
)

// DataConsentRequesterDto struct for DataConsentRequesterDto
type DataConsentRequesterDto struct {
	Name NullableString `json:"name,omitempty"`
	LogoUrl NullableString `json:"logoUrl,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Location NullableString `json:"location,omitempty"`
	Website NullableString `json:"website,omitempty"`
	SupportEmail NullableString `json:"supportEmail,omitempty"`
	HelpLineNumber NullableString `json:"helpLineNumber,omitempty"`
}

// NewDataConsentRequesterDto instantiates a new DataConsentRequesterDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataConsentRequesterDto() *DataConsentRequesterDto {
	this := DataConsentRequesterDto{}
	return &this
}

// NewDataConsentRequesterDtoWithDefaults instantiates a new DataConsentRequesterDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataConsentRequesterDtoWithDefaults() *DataConsentRequesterDto {
	this := DataConsentRequesterDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsentRequesterDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsentRequesterDto) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DataConsentRequesterDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DataConsentRequesterDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *DataConsentRequesterDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DataConsentRequesterDto) UnsetName() {
	o.Name.Unset()
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsentRequesterDto) GetLogoUrl() string {
	if o == nil || o.LogoUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.LogoUrl.Get()
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsentRequesterDto) GetLogoUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LogoUrl.Get(), o.LogoUrl.IsSet()
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *DataConsentRequesterDto) HasLogoUrl() bool {
	if o != nil && o.LogoUrl.IsSet() {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given NullableString and assigns it to the LogoUrl field.
func (o *DataConsentRequesterDto) SetLogoUrl(v string) {
	o.LogoUrl.Set(&v)
}
// SetLogoUrlNil sets the value for LogoUrl to be an explicit nil
func (o *DataConsentRequesterDto) SetLogoUrlNil() {
	o.LogoUrl.Set(nil)
}

// UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
func (o *DataConsentRequesterDto) UnsetLogoUrl() {
	o.LogoUrl.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsentRequesterDto) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsentRequesterDto) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *DataConsentRequesterDto) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *DataConsentRequesterDto) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *DataConsentRequesterDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *DataConsentRequesterDto) UnsetDescription() {
	o.Description.Unset()
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsentRequesterDto) GetLocation() string {
	if o == nil || o.Location.Get() == nil {
		var ret string
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsentRequesterDto) GetLocationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *DataConsentRequesterDto) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableString and assigns it to the Location field.
func (o *DataConsentRequesterDto) SetLocation(v string) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *DataConsentRequesterDto) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *DataConsentRequesterDto) UnsetLocation() {
	o.Location.Unset()
}

// GetWebsite returns the Website field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsentRequesterDto) GetWebsite() string {
	if o == nil || o.Website.Get() == nil {
		var ret string
		return ret
	}
	return *o.Website.Get()
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsentRequesterDto) GetWebsiteOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Website.Get(), o.Website.IsSet()
}

// HasWebsite returns a boolean if a field has been set.
func (o *DataConsentRequesterDto) HasWebsite() bool {
	if o != nil && o.Website.IsSet() {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given NullableString and assigns it to the Website field.
func (o *DataConsentRequesterDto) SetWebsite(v string) {
	o.Website.Set(&v)
}
// SetWebsiteNil sets the value for Website to be an explicit nil
func (o *DataConsentRequesterDto) SetWebsiteNil() {
	o.Website.Set(nil)
}

// UnsetWebsite ensures that no value is present for Website, not even an explicit nil
func (o *DataConsentRequesterDto) UnsetWebsite() {
	o.Website.Unset()
}

// GetSupportEmail returns the SupportEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsentRequesterDto) GetSupportEmail() string {
	if o == nil || o.SupportEmail.Get() == nil {
		var ret string
		return ret
	}
	return *o.SupportEmail.Get()
}

// GetSupportEmailOk returns a tuple with the SupportEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsentRequesterDto) GetSupportEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SupportEmail.Get(), o.SupportEmail.IsSet()
}

// HasSupportEmail returns a boolean if a field has been set.
func (o *DataConsentRequesterDto) HasSupportEmail() bool {
	if o != nil && o.SupportEmail.IsSet() {
		return true
	}

	return false
}

// SetSupportEmail gets a reference to the given NullableString and assigns it to the SupportEmail field.
func (o *DataConsentRequesterDto) SetSupportEmail(v string) {
	o.SupportEmail.Set(&v)
}
// SetSupportEmailNil sets the value for SupportEmail to be an explicit nil
func (o *DataConsentRequesterDto) SetSupportEmailNil() {
	o.SupportEmail.Set(nil)
}

// UnsetSupportEmail ensures that no value is present for SupportEmail, not even an explicit nil
func (o *DataConsentRequesterDto) UnsetSupportEmail() {
	o.SupportEmail.Unset()
}

// GetHelpLineNumber returns the HelpLineNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsentRequesterDto) GetHelpLineNumber() string {
	if o == nil || o.HelpLineNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.HelpLineNumber.Get()
}

// GetHelpLineNumberOk returns a tuple with the HelpLineNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsentRequesterDto) GetHelpLineNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HelpLineNumber.Get(), o.HelpLineNumber.IsSet()
}

// HasHelpLineNumber returns a boolean if a field has been set.
func (o *DataConsentRequesterDto) HasHelpLineNumber() bool {
	if o != nil && o.HelpLineNumber.IsSet() {
		return true
	}

	return false
}

// SetHelpLineNumber gets a reference to the given NullableString and assigns it to the HelpLineNumber field.
func (o *DataConsentRequesterDto) SetHelpLineNumber(v string) {
	o.HelpLineNumber.Set(&v)
}
// SetHelpLineNumberNil sets the value for HelpLineNumber to be an explicit nil
func (o *DataConsentRequesterDto) SetHelpLineNumberNil() {
	o.HelpLineNumber.Set(nil)
}

// UnsetHelpLineNumber ensures that no value is present for HelpLineNumber, not even an explicit nil
func (o *DataConsentRequesterDto) UnsetHelpLineNumber() {
	o.HelpLineNumber.Unset()
}

func (o DataConsentRequesterDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.LogoUrl.IsSet() {
		toSerialize["logoUrl"] = o.LogoUrl.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if o.Website.IsSet() {
		toSerialize["website"] = o.Website.Get()
	}
	if o.SupportEmail.IsSet() {
		toSerialize["supportEmail"] = o.SupportEmail.Get()
	}
	if o.HelpLineNumber.IsSet() {
		toSerialize["helpLineNumber"] = o.HelpLineNumber.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDataConsentRequesterDto struct {
	value *DataConsentRequesterDto
	isSet bool
}

func (v NullableDataConsentRequesterDto) Get() *DataConsentRequesterDto {
	return v.value
}

func (v *NullableDataConsentRequesterDto) Set(val *DataConsentRequesterDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDataConsentRequesterDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDataConsentRequesterDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataConsentRequesterDto(val *DataConsentRequesterDto) *NullableDataConsentRequesterDto {
	return &NullableDataConsentRequesterDto{value: val, isSet: true}
}

func (v NullableDataConsentRequesterDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataConsentRequesterDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



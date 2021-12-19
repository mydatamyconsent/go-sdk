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
)

// DataConsentRequestedAccountDto struct for DataConsentRequestedAccountDto
type DataConsentRequestedAccountDto struct {
	Name NullableString `json:"name,omitempty"`
	AccountTypeId *string `json:"accountTypeId,omitempty"`
	SuggestedAccounts []SuggestedAccountDto `json:"suggestedAccounts,omitempty"`
	Issuer []string `json:"issuer,omitempty"`
	IssuerLogoUrls []string `json:"issuerLogoUrls,omitempty"`
	RequestedDataType NullableString `json:"requestedDataType,omitempty"`
	Optional *bool `json:"optional,omitempty"`
}

// NewDataConsentRequestedAccountDto instantiates a new DataConsentRequestedAccountDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataConsentRequestedAccountDto() *DataConsentRequestedAccountDto {
	this := DataConsentRequestedAccountDto{}
	return &this
}

// NewDataConsentRequestedAccountDtoWithDefaults instantiates a new DataConsentRequestedAccountDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataConsentRequestedAccountDtoWithDefaults() *DataConsentRequestedAccountDto {
	this := DataConsentRequestedAccountDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsentRequestedAccountDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsentRequestedAccountDto) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DataConsentRequestedAccountDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DataConsentRequestedAccountDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *DataConsentRequestedAccountDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DataConsentRequestedAccountDto) UnsetName() {
	o.Name.Unset()
}

// GetAccountTypeId returns the AccountTypeId field value if set, zero value otherwise.
func (o *DataConsentRequestedAccountDto) GetAccountTypeId() string {
	if o == nil || o.AccountTypeId == nil {
		var ret string
		return ret
	}
	return *o.AccountTypeId
}

// GetAccountTypeIdOk returns a tuple with the AccountTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataConsentRequestedAccountDto) GetAccountTypeIdOk() (*string, bool) {
	if o == nil || o.AccountTypeId == nil {
		return nil, false
	}
	return o.AccountTypeId, true
}

// HasAccountTypeId returns a boolean if a field has been set.
func (o *DataConsentRequestedAccountDto) HasAccountTypeId() bool {
	if o != nil && o.AccountTypeId != nil {
		return true
	}

	return false
}

// SetAccountTypeId gets a reference to the given string and assigns it to the AccountTypeId field.
func (o *DataConsentRequestedAccountDto) SetAccountTypeId(v string) {
	o.AccountTypeId = &v
}

// GetSuggestedAccounts returns the SuggestedAccounts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsentRequestedAccountDto) GetSuggestedAccounts() []SuggestedAccountDto {
	if o == nil  {
		var ret []SuggestedAccountDto
		return ret
	}
	return o.SuggestedAccounts
}

// GetSuggestedAccountsOk returns a tuple with the SuggestedAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsentRequestedAccountDto) GetSuggestedAccountsOk() (*[]SuggestedAccountDto, bool) {
	if o == nil || o.SuggestedAccounts == nil {
		return nil, false
	}
	return &o.SuggestedAccounts, true
}

// HasSuggestedAccounts returns a boolean if a field has been set.
func (o *DataConsentRequestedAccountDto) HasSuggestedAccounts() bool {
	if o != nil && o.SuggestedAccounts != nil {
		return true
	}

	return false
}

// SetSuggestedAccounts gets a reference to the given []SuggestedAccountDto and assigns it to the SuggestedAccounts field.
func (o *DataConsentRequestedAccountDto) SetSuggestedAccounts(v []SuggestedAccountDto) {
	o.SuggestedAccounts = v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsentRequestedAccountDto) GetIssuer() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsentRequestedAccountDto) GetIssuerOk() (*[]string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *DataConsentRequestedAccountDto) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given []string and assigns it to the Issuer field.
func (o *DataConsentRequestedAccountDto) SetIssuer(v []string) {
	o.Issuer = v
}

// GetIssuerLogoUrls returns the IssuerLogoUrls field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsentRequestedAccountDto) GetIssuerLogoUrls() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.IssuerLogoUrls
}

// GetIssuerLogoUrlsOk returns a tuple with the IssuerLogoUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsentRequestedAccountDto) GetIssuerLogoUrlsOk() (*[]string, bool) {
	if o == nil || o.IssuerLogoUrls == nil {
		return nil, false
	}
	return &o.IssuerLogoUrls, true
}

// HasIssuerLogoUrls returns a boolean if a field has been set.
func (o *DataConsentRequestedAccountDto) HasIssuerLogoUrls() bool {
	if o != nil && o.IssuerLogoUrls != nil {
		return true
	}

	return false
}

// SetIssuerLogoUrls gets a reference to the given []string and assigns it to the IssuerLogoUrls field.
func (o *DataConsentRequestedAccountDto) SetIssuerLogoUrls(v []string) {
	o.IssuerLogoUrls = v
}

// GetRequestedDataType returns the RequestedDataType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataConsentRequestedAccountDto) GetRequestedDataType() string {
	if o == nil || o.RequestedDataType.Get() == nil {
		var ret string
		return ret
	}
	return *o.RequestedDataType.Get()
}

// GetRequestedDataTypeOk returns a tuple with the RequestedDataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataConsentRequestedAccountDto) GetRequestedDataTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RequestedDataType.Get(), o.RequestedDataType.IsSet()
}

// HasRequestedDataType returns a boolean if a field has been set.
func (o *DataConsentRequestedAccountDto) HasRequestedDataType() bool {
	if o != nil && o.RequestedDataType.IsSet() {
		return true
	}

	return false
}

// SetRequestedDataType gets a reference to the given NullableString and assigns it to the RequestedDataType field.
func (o *DataConsentRequestedAccountDto) SetRequestedDataType(v string) {
	o.RequestedDataType.Set(&v)
}
// SetRequestedDataTypeNil sets the value for RequestedDataType to be an explicit nil
func (o *DataConsentRequestedAccountDto) SetRequestedDataTypeNil() {
	o.RequestedDataType.Set(nil)
}

// UnsetRequestedDataType ensures that no value is present for RequestedDataType, not even an explicit nil
func (o *DataConsentRequestedAccountDto) UnsetRequestedDataType() {
	o.RequestedDataType.Unset()
}

// GetOptional returns the Optional field value if set, zero value otherwise.
func (o *DataConsentRequestedAccountDto) GetOptional() bool {
	if o == nil || o.Optional == nil {
		var ret bool
		return ret
	}
	return *o.Optional
}

// GetOptionalOk returns a tuple with the Optional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataConsentRequestedAccountDto) GetOptionalOk() (*bool, bool) {
	if o == nil || o.Optional == nil {
		return nil, false
	}
	return o.Optional, true
}

// HasOptional returns a boolean if a field has been set.
func (o *DataConsentRequestedAccountDto) HasOptional() bool {
	if o != nil && o.Optional != nil {
		return true
	}

	return false
}

// SetOptional gets a reference to the given bool and assigns it to the Optional field.
func (o *DataConsentRequestedAccountDto) SetOptional(v bool) {
	o.Optional = &v
}

func (o DataConsentRequestedAccountDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.AccountTypeId != nil {
		toSerialize["accountTypeId"] = o.AccountTypeId
	}
	if o.SuggestedAccounts != nil {
		toSerialize["suggestedAccounts"] = o.SuggestedAccounts
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.IssuerLogoUrls != nil {
		toSerialize["issuerLogoUrls"] = o.IssuerLogoUrls
	}
	if o.RequestedDataType.IsSet() {
		toSerialize["requestedDataType"] = o.RequestedDataType.Get()
	}
	if o.Optional != nil {
		toSerialize["optional"] = o.Optional
	}
	return json.Marshal(toSerialize)
}

type NullableDataConsentRequestedAccountDto struct {
	value *DataConsentRequestedAccountDto
	isSet bool
}

func (v NullableDataConsentRequestedAccountDto) Get() *DataConsentRequestedAccountDto {
	return v.value
}

func (v *NullableDataConsentRequestedAccountDto) Set(val *DataConsentRequestedAccountDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDataConsentRequestedAccountDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDataConsentRequestedAccountDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataConsentRequestedAccountDto(val *DataConsentRequestedAccountDto) *NullableDataConsentRequestedAccountDto {
	return &NullableDataConsentRequestedAccountDto{value: val, isSet: true}
}

func (v NullableDataConsentRequestedAccountDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataConsentRequestedAccountDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



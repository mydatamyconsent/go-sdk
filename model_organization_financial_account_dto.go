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

// OrganizationFinancialAccountDto struct for OrganizationFinancialAccountDto
type OrganizationFinancialAccountDto struct {
	Id *string `json:"id,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty"`
	OrganizationName NullableString `json:"organizationName,omitempty"`
	BeneficiaryName NullableString `json:"beneficiaryName,omitempty"`
	AccountNumber NullableString `json:"accountNumber,omitempty"`
	RoutingNumber NullableString `json:"routingNumber,omitempty"`
	IsPrimary *bool `json:"isPrimary,omitempty"`
	IsVerified *bool `json:"isVerified,omitempty"`
	LogoUrl NullableString `json:"logoUrl,omitempty"`
	BankName NullableString `json:"bankName,omitempty"`
	BankAccountType *BankAccountType `json:"bankAccountType,omitempty"`
	BankAccountProofUrl NullableString `json:"bankAccountProofUrl,omitempty"`
}

// NewOrganizationFinancialAccountDto instantiates a new OrganizationFinancialAccountDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationFinancialAccountDto() *OrganizationFinancialAccountDto {
	this := OrganizationFinancialAccountDto{}
	return &this
}

// NewOrganizationFinancialAccountDtoWithDefaults instantiates a new OrganizationFinancialAccountDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationFinancialAccountDtoWithDefaults() *OrganizationFinancialAccountDto {
	this := OrganizationFinancialAccountDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationFinancialAccountDto) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationFinancialAccountDto) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationFinancialAccountDto) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationFinancialAccountDto) SetId(v string) {
	o.Id = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *OrganizationFinancialAccountDto) GetOrganizationId() string {
	if o == nil || o.OrganizationId == nil {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationFinancialAccountDto) GetOrganizationIdOk() (*string, bool) {
	if o == nil || o.OrganizationId == nil {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *OrganizationFinancialAccountDto) HasOrganizationId() bool {
	if o != nil && o.OrganizationId != nil {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *OrganizationFinancialAccountDto) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationFinancialAccountDto) GetOrganizationName() string {
	if o == nil || o.OrganizationName.Get() == nil {
		var ret string
		return ret
	}
	return *o.OrganizationName.Get()
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationFinancialAccountDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationName.Get(), o.OrganizationName.IsSet()
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *OrganizationFinancialAccountDto) HasOrganizationName() bool {
	if o != nil && o.OrganizationName.IsSet() {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given NullableString and assigns it to the OrganizationName field.
func (o *OrganizationFinancialAccountDto) SetOrganizationName(v string) {
	o.OrganizationName.Set(&v)
}
// SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil
func (o *OrganizationFinancialAccountDto) SetOrganizationNameNil() {
	o.OrganizationName.Set(nil)
}

// UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
func (o *OrganizationFinancialAccountDto) UnsetOrganizationName() {
	o.OrganizationName.Unset()
}

// GetBeneficiaryName returns the BeneficiaryName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationFinancialAccountDto) GetBeneficiaryName() string {
	if o == nil || o.BeneficiaryName.Get() == nil {
		var ret string
		return ret
	}
	return *o.BeneficiaryName.Get()
}

// GetBeneficiaryNameOk returns a tuple with the BeneficiaryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationFinancialAccountDto) GetBeneficiaryNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BeneficiaryName.Get(), o.BeneficiaryName.IsSet()
}

// HasBeneficiaryName returns a boolean if a field has been set.
func (o *OrganizationFinancialAccountDto) HasBeneficiaryName() bool {
	if o != nil && o.BeneficiaryName.IsSet() {
		return true
	}

	return false
}

// SetBeneficiaryName gets a reference to the given NullableString and assigns it to the BeneficiaryName field.
func (o *OrganizationFinancialAccountDto) SetBeneficiaryName(v string) {
	o.BeneficiaryName.Set(&v)
}
// SetBeneficiaryNameNil sets the value for BeneficiaryName to be an explicit nil
func (o *OrganizationFinancialAccountDto) SetBeneficiaryNameNil() {
	o.BeneficiaryName.Set(nil)
}

// UnsetBeneficiaryName ensures that no value is present for BeneficiaryName, not even an explicit nil
func (o *OrganizationFinancialAccountDto) UnsetBeneficiaryName() {
	o.BeneficiaryName.Unset()
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationFinancialAccountDto) GetAccountNumber() string {
	if o == nil || o.AccountNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountNumber.Get()
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationFinancialAccountDto) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountNumber.Get(), o.AccountNumber.IsSet()
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *OrganizationFinancialAccountDto) HasAccountNumber() bool {
	if o != nil && o.AccountNumber.IsSet() {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given NullableString and assigns it to the AccountNumber field.
func (o *OrganizationFinancialAccountDto) SetAccountNumber(v string) {
	o.AccountNumber.Set(&v)
}
// SetAccountNumberNil sets the value for AccountNumber to be an explicit nil
func (o *OrganizationFinancialAccountDto) SetAccountNumberNil() {
	o.AccountNumber.Set(nil)
}

// UnsetAccountNumber ensures that no value is present for AccountNumber, not even an explicit nil
func (o *OrganizationFinancialAccountDto) UnsetAccountNumber() {
	o.AccountNumber.Unset()
}

// GetRoutingNumber returns the RoutingNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationFinancialAccountDto) GetRoutingNumber() string {
	if o == nil || o.RoutingNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.RoutingNumber.Get()
}

// GetRoutingNumberOk returns a tuple with the RoutingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationFinancialAccountDto) GetRoutingNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoutingNumber.Get(), o.RoutingNumber.IsSet()
}

// HasRoutingNumber returns a boolean if a field has been set.
func (o *OrganizationFinancialAccountDto) HasRoutingNumber() bool {
	if o != nil && o.RoutingNumber.IsSet() {
		return true
	}

	return false
}

// SetRoutingNumber gets a reference to the given NullableString and assigns it to the RoutingNumber field.
func (o *OrganizationFinancialAccountDto) SetRoutingNumber(v string) {
	o.RoutingNumber.Set(&v)
}
// SetRoutingNumberNil sets the value for RoutingNumber to be an explicit nil
func (o *OrganizationFinancialAccountDto) SetRoutingNumberNil() {
	o.RoutingNumber.Set(nil)
}

// UnsetRoutingNumber ensures that no value is present for RoutingNumber, not even an explicit nil
func (o *OrganizationFinancialAccountDto) UnsetRoutingNumber() {
	o.RoutingNumber.Unset()
}

// GetIsPrimary returns the IsPrimary field value if set, zero value otherwise.
func (o *OrganizationFinancialAccountDto) GetIsPrimary() bool {
	if o == nil || o.IsPrimary == nil {
		var ret bool
		return ret
	}
	return *o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationFinancialAccountDto) GetIsPrimaryOk() (*bool, bool) {
	if o == nil || o.IsPrimary == nil {
		return nil, false
	}
	return o.IsPrimary, true
}

// HasIsPrimary returns a boolean if a field has been set.
func (o *OrganizationFinancialAccountDto) HasIsPrimary() bool {
	if o != nil && o.IsPrimary != nil {
		return true
	}

	return false
}

// SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.
func (o *OrganizationFinancialAccountDto) SetIsPrimary(v bool) {
	o.IsPrimary = &v
}

// GetIsVerified returns the IsVerified field value if set, zero value otherwise.
func (o *OrganizationFinancialAccountDto) GetIsVerified() bool {
	if o == nil || o.IsVerified == nil {
		var ret bool
		return ret
	}
	return *o.IsVerified
}

// GetIsVerifiedOk returns a tuple with the IsVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationFinancialAccountDto) GetIsVerifiedOk() (*bool, bool) {
	if o == nil || o.IsVerified == nil {
		return nil, false
	}
	return o.IsVerified, true
}

// HasIsVerified returns a boolean if a field has been set.
func (o *OrganizationFinancialAccountDto) HasIsVerified() bool {
	if o != nil && o.IsVerified != nil {
		return true
	}

	return false
}

// SetIsVerified gets a reference to the given bool and assigns it to the IsVerified field.
func (o *OrganizationFinancialAccountDto) SetIsVerified(v bool) {
	o.IsVerified = &v
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationFinancialAccountDto) GetLogoUrl() string {
	if o == nil || o.LogoUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.LogoUrl.Get()
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationFinancialAccountDto) GetLogoUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogoUrl.Get(), o.LogoUrl.IsSet()
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *OrganizationFinancialAccountDto) HasLogoUrl() bool {
	if o != nil && o.LogoUrl.IsSet() {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given NullableString and assigns it to the LogoUrl field.
func (o *OrganizationFinancialAccountDto) SetLogoUrl(v string) {
	o.LogoUrl.Set(&v)
}
// SetLogoUrlNil sets the value for LogoUrl to be an explicit nil
func (o *OrganizationFinancialAccountDto) SetLogoUrlNil() {
	o.LogoUrl.Set(nil)
}

// UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
func (o *OrganizationFinancialAccountDto) UnsetLogoUrl() {
	o.LogoUrl.Unset()
}

// GetBankName returns the BankName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationFinancialAccountDto) GetBankName() string {
	if o == nil || o.BankName.Get() == nil {
		var ret string
		return ret
	}
	return *o.BankName.Get()
}

// GetBankNameOk returns a tuple with the BankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationFinancialAccountDto) GetBankNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BankName.Get(), o.BankName.IsSet()
}

// HasBankName returns a boolean if a field has been set.
func (o *OrganizationFinancialAccountDto) HasBankName() bool {
	if o != nil && o.BankName.IsSet() {
		return true
	}

	return false
}

// SetBankName gets a reference to the given NullableString and assigns it to the BankName field.
func (o *OrganizationFinancialAccountDto) SetBankName(v string) {
	o.BankName.Set(&v)
}
// SetBankNameNil sets the value for BankName to be an explicit nil
func (o *OrganizationFinancialAccountDto) SetBankNameNil() {
	o.BankName.Set(nil)
}

// UnsetBankName ensures that no value is present for BankName, not even an explicit nil
func (o *OrganizationFinancialAccountDto) UnsetBankName() {
	o.BankName.Unset()
}

// GetBankAccountType returns the BankAccountType field value if set, zero value otherwise.
func (o *OrganizationFinancialAccountDto) GetBankAccountType() BankAccountType {
	if o == nil || o.BankAccountType == nil {
		var ret BankAccountType
		return ret
	}
	return *o.BankAccountType
}

// GetBankAccountTypeOk returns a tuple with the BankAccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationFinancialAccountDto) GetBankAccountTypeOk() (*BankAccountType, bool) {
	if o == nil || o.BankAccountType == nil {
		return nil, false
	}
	return o.BankAccountType, true
}

// HasBankAccountType returns a boolean if a field has been set.
func (o *OrganizationFinancialAccountDto) HasBankAccountType() bool {
	if o != nil && o.BankAccountType != nil {
		return true
	}

	return false
}

// SetBankAccountType gets a reference to the given BankAccountType and assigns it to the BankAccountType field.
func (o *OrganizationFinancialAccountDto) SetBankAccountType(v BankAccountType) {
	o.BankAccountType = &v
}

// GetBankAccountProofUrl returns the BankAccountProofUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationFinancialAccountDto) GetBankAccountProofUrl() string {
	if o == nil || o.BankAccountProofUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.BankAccountProofUrl.Get()
}

// GetBankAccountProofUrlOk returns a tuple with the BankAccountProofUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationFinancialAccountDto) GetBankAccountProofUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BankAccountProofUrl.Get(), o.BankAccountProofUrl.IsSet()
}

// HasBankAccountProofUrl returns a boolean if a field has been set.
func (o *OrganizationFinancialAccountDto) HasBankAccountProofUrl() bool {
	if o != nil && o.BankAccountProofUrl.IsSet() {
		return true
	}

	return false
}

// SetBankAccountProofUrl gets a reference to the given NullableString and assigns it to the BankAccountProofUrl field.
func (o *OrganizationFinancialAccountDto) SetBankAccountProofUrl(v string) {
	o.BankAccountProofUrl.Set(&v)
}
// SetBankAccountProofUrlNil sets the value for BankAccountProofUrl to be an explicit nil
func (o *OrganizationFinancialAccountDto) SetBankAccountProofUrlNil() {
	o.BankAccountProofUrl.Set(nil)
}

// UnsetBankAccountProofUrl ensures that no value is present for BankAccountProofUrl, not even an explicit nil
func (o *OrganizationFinancialAccountDto) UnsetBankAccountProofUrl() {
	o.BankAccountProofUrl.Unset()
}

func (o OrganizationFinancialAccountDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.OrganizationId != nil {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if o.OrganizationName.IsSet() {
		toSerialize["organizationName"] = o.OrganizationName.Get()
	}
	if o.BeneficiaryName.IsSet() {
		toSerialize["beneficiaryName"] = o.BeneficiaryName.Get()
	}
	if o.AccountNumber.IsSet() {
		toSerialize["accountNumber"] = o.AccountNumber.Get()
	}
	if o.RoutingNumber.IsSet() {
		toSerialize["routingNumber"] = o.RoutingNumber.Get()
	}
	if o.IsPrimary != nil {
		toSerialize["isPrimary"] = o.IsPrimary
	}
	if o.IsVerified != nil {
		toSerialize["isVerified"] = o.IsVerified
	}
	if o.LogoUrl.IsSet() {
		toSerialize["logoUrl"] = o.LogoUrl.Get()
	}
	if o.BankName.IsSet() {
		toSerialize["bankName"] = o.BankName.Get()
	}
	if o.BankAccountType != nil {
		toSerialize["bankAccountType"] = o.BankAccountType
	}
	if o.BankAccountProofUrl.IsSet() {
		toSerialize["bankAccountProofUrl"] = o.BankAccountProofUrl.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationFinancialAccountDto struct {
	value *OrganizationFinancialAccountDto
	isSet bool
}

func (v NullableOrganizationFinancialAccountDto) Get() *OrganizationFinancialAccountDto {
	return v.value
}

func (v *NullableOrganizationFinancialAccountDto) Set(val *OrganizationFinancialAccountDto) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationFinancialAccountDto) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationFinancialAccountDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationFinancialAccountDto(val *OrganizationFinancialAccountDto) *NullableOrganizationFinancialAccountDto {
	return &NullableOrganizationFinancialAccountDto{value: val, isSet: true}
}

func (v NullableOrganizationFinancialAccountDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationFinancialAccountDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



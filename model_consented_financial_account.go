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
)

// ConsentedFinancialAccount ConsentedFinancialAccount : Consented financial account details.
type ConsentedFinancialAccount struct {
	// Financial account id.
	Id string `json:"id"`
	// Financial account name.
	Name string `json:"name"`
	Category FinancialAccountCategoryType `json:"category"`
	SubCategory FinancialAccountSubCategoryType `json:"subCategory"`
	// Financial account identifier.
	Identifier string `json:"identifier"`
	// Financial account field title.
	FieldTitle string `json:"fieldTitle"`
	// Financial account field slug.
	FieldSlug string `json:"fieldSlug"`
	// Requested financial account details.
	RequestedDetails []FinancialAccountDetailsRequired `json:"requestedDetails"`
	TransactionPeriod *ConsentedFinancialAccountTransactionPeriod `json:"transactionPeriod,omitempty"`
	// Financial account issuer id.
	IssuerId string `json:"issuerId"`
	// Financial account issuer name.
	IssuerName string `json:"issuerName"`
}

// NewConsentedFinancialAccount instantiates a new ConsentedFinancialAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsentedFinancialAccount(id string, name string, category FinancialAccountCategoryType, subCategory FinancialAccountSubCategoryType, identifier string, fieldTitle string, fieldSlug string, requestedDetails []FinancialAccountDetailsRequired, issuerId string, issuerName string) *ConsentedFinancialAccount {
	this := ConsentedFinancialAccount{}
	this.Id = id
	this.Name = name
	this.Category = category
	this.SubCategory = subCategory
	this.Identifier = identifier
	this.FieldTitle = fieldTitle
	this.FieldSlug = fieldSlug
	this.RequestedDetails = requestedDetails
	this.IssuerId = issuerId
	this.IssuerName = issuerName
	return &this
}

// NewConsentedFinancialAccountWithDefaults instantiates a new ConsentedFinancialAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsentedFinancialAccountWithDefaults() *ConsentedFinancialAccount {
	this := ConsentedFinancialAccount{}
	return &this
}

// GetId returns the Id field value
func (o *ConsentedFinancialAccount) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConsentedFinancialAccount) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConsentedFinancialAccount) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ConsentedFinancialAccount) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConsentedFinancialAccount) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConsentedFinancialAccount) SetName(v string) {
	o.Name = v
}

// GetCategory returns the Category field value
func (o *ConsentedFinancialAccount) GetCategory() FinancialAccountCategoryType {
	if o == nil {
		var ret FinancialAccountCategoryType
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *ConsentedFinancialAccount) GetCategoryOk() (*FinancialAccountCategoryType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *ConsentedFinancialAccount) SetCategory(v FinancialAccountCategoryType) {
	o.Category = v
}

// GetSubCategory returns the SubCategory field value
func (o *ConsentedFinancialAccount) GetSubCategory() FinancialAccountSubCategoryType {
	if o == nil {
		var ret FinancialAccountSubCategoryType
		return ret
	}

	return o.SubCategory
}

// GetSubCategoryOk returns a tuple with the SubCategory field value
// and a boolean to check if the value has been set.
func (o *ConsentedFinancialAccount) GetSubCategoryOk() (*FinancialAccountSubCategoryType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubCategory, true
}

// SetSubCategory sets field value
func (o *ConsentedFinancialAccount) SetSubCategory(v FinancialAccountSubCategoryType) {
	o.SubCategory = v
}

// GetIdentifier returns the Identifier field value
func (o *ConsentedFinancialAccount) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *ConsentedFinancialAccount) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *ConsentedFinancialAccount) SetIdentifier(v string) {
	o.Identifier = v
}

// GetFieldTitle returns the FieldTitle field value
func (o *ConsentedFinancialAccount) GetFieldTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldTitle
}

// GetFieldTitleOk returns a tuple with the FieldTitle field value
// and a boolean to check if the value has been set.
func (o *ConsentedFinancialAccount) GetFieldTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldTitle, true
}

// SetFieldTitle sets field value
func (o *ConsentedFinancialAccount) SetFieldTitle(v string) {
	o.FieldTitle = v
}

// GetFieldSlug returns the FieldSlug field value
func (o *ConsentedFinancialAccount) GetFieldSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldSlug
}

// GetFieldSlugOk returns a tuple with the FieldSlug field value
// and a boolean to check if the value has been set.
func (o *ConsentedFinancialAccount) GetFieldSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldSlug, true
}

// SetFieldSlug sets field value
func (o *ConsentedFinancialAccount) SetFieldSlug(v string) {
	o.FieldSlug = v
}

// GetRequestedDetails returns the RequestedDetails field value
func (o *ConsentedFinancialAccount) GetRequestedDetails() []FinancialAccountDetailsRequired {
	if o == nil {
		var ret []FinancialAccountDetailsRequired
		return ret
	}

	return o.RequestedDetails
}

// GetRequestedDetailsOk returns a tuple with the RequestedDetails field value
// and a boolean to check if the value has been set.
func (o *ConsentedFinancialAccount) GetRequestedDetailsOk() ([]FinancialAccountDetailsRequired, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestedDetails, true
}

// SetRequestedDetails sets field value
func (o *ConsentedFinancialAccount) SetRequestedDetails(v []FinancialAccountDetailsRequired) {
	o.RequestedDetails = v
}

// GetTransactionPeriod returns the TransactionPeriod field value if set, zero value otherwise.
func (o *ConsentedFinancialAccount) GetTransactionPeriod() ConsentedFinancialAccountTransactionPeriod {
	if o == nil || o.TransactionPeriod == nil {
		var ret ConsentedFinancialAccountTransactionPeriod
		return ret
	}
	return *o.TransactionPeriod
}

// GetTransactionPeriodOk returns a tuple with the TransactionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentedFinancialAccount) GetTransactionPeriodOk() (*ConsentedFinancialAccountTransactionPeriod, bool) {
	if o == nil || o.TransactionPeriod == nil {
		return nil, false
	}
	return o.TransactionPeriod, true
}

// HasTransactionPeriod returns a boolean if a field has been set.
func (o *ConsentedFinancialAccount) HasTransactionPeriod() bool {
	if o != nil && o.TransactionPeriod != nil {
		return true
	}

	return false
}

// SetTransactionPeriod gets a reference to the given ConsentedFinancialAccountTransactionPeriod and assigns it to the TransactionPeriod field.
func (o *ConsentedFinancialAccount) SetTransactionPeriod(v ConsentedFinancialAccountTransactionPeriod) {
	o.TransactionPeriod = &v
}

// GetIssuerId returns the IssuerId field value
func (o *ConsentedFinancialAccount) GetIssuerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuerId
}

// GetIssuerIdOk returns a tuple with the IssuerId field value
// and a boolean to check if the value has been set.
func (o *ConsentedFinancialAccount) GetIssuerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuerId, true
}

// SetIssuerId sets field value
func (o *ConsentedFinancialAccount) SetIssuerId(v string) {
	o.IssuerId = v
}

// GetIssuerName returns the IssuerName field value
func (o *ConsentedFinancialAccount) GetIssuerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuerName
}

// GetIssuerNameOk returns a tuple with the IssuerName field value
// and a boolean to check if the value has been set.
func (o *ConsentedFinancialAccount) GetIssuerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuerName, true
}

// SetIssuerName sets field value
func (o *ConsentedFinancialAccount) SetIssuerName(v string) {
	o.IssuerName = v
}

func (o ConsentedFinancialAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["category"] = o.Category
	}
	if true {
		toSerialize["subCategory"] = o.SubCategory
	}
	if true {
		toSerialize["identifier"] = o.Identifier
	}
	if true {
		toSerialize["fieldTitle"] = o.FieldTitle
	}
	if true {
		toSerialize["fieldSlug"] = o.FieldSlug
	}
	if true {
		toSerialize["requestedDetails"] = o.RequestedDetails
	}
	if o.TransactionPeriod != nil {
		toSerialize["transactionPeriod"] = o.TransactionPeriod
	}
	if true {
		toSerialize["issuerId"] = o.IssuerId
	}
	if true {
		toSerialize["issuerName"] = o.IssuerName
	}
	return json.Marshal(toSerialize)
}

type NullableConsentedFinancialAccount struct {
	value *ConsentedFinancialAccount
	isSet bool
}

func (v NullableConsentedFinancialAccount) Get() *ConsentedFinancialAccount {
	return v.value
}

func (v *NullableConsentedFinancialAccount) Set(val *ConsentedFinancialAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableConsentedFinancialAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableConsentedFinancialAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsentedFinancialAccount(val *ConsentedFinancialAccount) *NullableConsentedFinancialAccount {
	return &NullableConsentedFinancialAccount{value: val, isSet: true}
}

func (v NullableConsentedFinancialAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsentedFinancialAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



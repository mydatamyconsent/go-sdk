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

// DocumentIssueRequestPaymentRequest struct for DocumentIssueRequestPaymentRequest
type DocumentIssueRequestPaymentRequest struct {
	Identifier string `json:"identifier"`
	Items []PaymentOrderItem `json:"items"`
	CurrencyCode string `json:"currencyCode"`
	PaymentUrl *string `json:"paymentUrl,omitempty"`
	Description string `json:"description"`
	DueByUtc time.Time `json:"dueByUtc"`
}

// NewDocumentIssueRequestPaymentRequest instantiates a new DocumentIssueRequestPaymentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentIssueRequestPaymentRequest(identifier string, items []PaymentOrderItem, currencyCode string, description string, dueByUtc time.Time) *DocumentIssueRequestPaymentRequest {
	this := DocumentIssueRequestPaymentRequest{}
	this.Identifier = identifier
	this.Items = items
	this.CurrencyCode = currencyCode
	this.Description = description
	this.DueByUtc = dueByUtc
	return &this
}

// NewDocumentIssueRequestPaymentRequestWithDefaults instantiates a new DocumentIssueRequestPaymentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentIssueRequestPaymentRequestWithDefaults() *DocumentIssueRequestPaymentRequest {
	this := DocumentIssueRequestPaymentRequest{}
	return &this
}

// GetIdentifier returns the Identifier field value
func (o *DocumentIssueRequestPaymentRequest) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *DocumentIssueRequestPaymentRequest) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *DocumentIssueRequestPaymentRequest) SetIdentifier(v string) {
	o.Identifier = v
}

// GetItems returns the Items field value
func (o *DocumentIssueRequestPaymentRequest) GetItems() []PaymentOrderItem {
	if o == nil {
		var ret []PaymentOrderItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *DocumentIssueRequestPaymentRequest) GetItemsOk() ([]PaymentOrderItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *DocumentIssueRequestPaymentRequest) SetItems(v []PaymentOrderItem) {
	o.Items = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *DocumentIssueRequestPaymentRequest) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *DocumentIssueRequestPaymentRequest) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *DocumentIssueRequestPaymentRequest) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

// GetPaymentUrl returns the PaymentUrl field value if set, zero value otherwise.
func (o *DocumentIssueRequestPaymentRequest) GetPaymentUrl() string {
	if o == nil || o.PaymentUrl == nil {
		var ret string
		return ret
	}
	return *o.PaymentUrl
}

// GetPaymentUrlOk returns a tuple with the PaymentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentIssueRequestPaymentRequest) GetPaymentUrlOk() (*string, bool) {
	if o == nil || o.PaymentUrl == nil {
		return nil, false
	}
	return o.PaymentUrl, true
}

// HasPaymentUrl returns a boolean if a field has been set.
func (o *DocumentIssueRequestPaymentRequest) HasPaymentUrl() bool {
	if o != nil && o.PaymentUrl != nil {
		return true
	}

	return false
}

// SetPaymentUrl gets a reference to the given string and assigns it to the PaymentUrl field.
func (o *DocumentIssueRequestPaymentRequest) SetPaymentUrl(v string) {
	o.PaymentUrl = &v
}

// GetDescription returns the Description field value
func (o *DocumentIssueRequestPaymentRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DocumentIssueRequestPaymentRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DocumentIssueRequestPaymentRequest) SetDescription(v string) {
	o.Description = v
}

// GetDueByUtc returns the DueByUtc field value
func (o *DocumentIssueRequestPaymentRequest) GetDueByUtc() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DueByUtc
}

// GetDueByUtcOk returns a tuple with the DueByUtc field value
// and a boolean to check if the value has been set.
func (o *DocumentIssueRequestPaymentRequest) GetDueByUtcOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DueByUtc, true
}

// SetDueByUtc sets field value
func (o *DocumentIssueRequestPaymentRequest) SetDueByUtc(v time.Time) {
	o.DueByUtc = v
}

func (o DocumentIssueRequestPaymentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["identifier"] = o.Identifier
	}
	if true {
		toSerialize["items"] = o.Items
	}
	if true {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if o.PaymentUrl != nil {
		toSerialize["paymentUrl"] = o.PaymentUrl
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["dueByUtc"] = o.DueByUtc
	}
	return json.Marshal(toSerialize)
}

type NullableDocumentIssueRequestPaymentRequest struct {
	value *DocumentIssueRequestPaymentRequest
	isSet bool
}

func (v NullableDocumentIssueRequestPaymentRequest) Get() *DocumentIssueRequestPaymentRequest {
	return v.value
}

func (v *NullableDocumentIssueRequestPaymentRequest) Set(val *DocumentIssueRequestPaymentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentIssueRequestPaymentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentIssueRequestPaymentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentIssueRequestPaymentRequest(val *DocumentIssueRequestPaymentRequest) *NullableDocumentIssueRequestPaymentRequest {
	return &NullableDocumentIssueRequestPaymentRequest{value: val, isSet: true}
}

func (v NullableDocumentIssueRequestPaymentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentIssueRequestPaymentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



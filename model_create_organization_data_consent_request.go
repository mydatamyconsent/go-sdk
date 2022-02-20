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

// CreateOrganizationDataConsentRequest Organization Data Consent Request.
type CreateOrganizationDataConsentRequest struct {
	ConsentTemplateId *string `json:"consentTemplateId,omitempty"`
	Receiver Receiver `json:"receiver"`
}

// NewCreateOrganizationDataConsentRequest instantiates a new CreateOrganizationDataConsentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationDataConsentRequest(receiver Receiver) *CreateOrganizationDataConsentRequest {
	this := CreateOrganizationDataConsentRequest{}
	this.Receiver = receiver
	return &this
}

// NewCreateOrganizationDataConsentRequestWithDefaults instantiates a new CreateOrganizationDataConsentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationDataConsentRequestWithDefaults() *CreateOrganizationDataConsentRequest {
	this := CreateOrganizationDataConsentRequest{}
	return &this
}

// GetConsentTemplateId returns the ConsentTemplateId field value if set, zero value otherwise.
func (o *CreateOrganizationDataConsentRequest) GetConsentTemplateId() string {
	if o == nil || o.ConsentTemplateId == nil {
		var ret string
		return ret
	}
	return *o.ConsentTemplateId
}

// GetConsentTemplateIdOk returns a tuple with the ConsentTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationDataConsentRequest) GetConsentTemplateIdOk() (*string, bool) {
	if o == nil || o.ConsentTemplateId == nil {
		return nil, false
	}
	return o.ConsentTemplateId, true
}

// HasConsentTemplateId returns a boolean if a field has been set.
func (o *CreateOrganizationDataConsentRequest) HasConsentTemplateId() bool {
	if o != nil && o.ConsentTemplateId != nil {
		return true
	}

	return false
}

// SetConsentTemplateId gets a reference to the given string and assigns it to the ConsentTemplateId field.
func (o *CreateOrganizationDataConsentRequest) SetConsentTemplateId(v string) {
	o.ConsentTemplateId = &v
}

// GetReceiver returns the Receiver field value
func (o *CreateOrganizationDataConsentRequest) GetReceiver() Receiver {
	if o == nil {
		var ret Receiver
		return ret
	}

	return o.Receiver
}

// GetReceiverOk returns a tuple with the Receiver field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationDataConsentRequest) GetReceiverOk() (*Receiver, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Receiver, true
}

// SetReceiver sets field value
func (o *CreateOrganizationDataConsentRequest) SetReceiver(v Receiver) {
	o.Receiver = v
}

func (o CreateOrganizationDataConsentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConsentTemplateId != nil {
		toSerialize["consentTemplateId"] = o.ConsentTemplateId
	}
	if true {
		toSerialize["receiver"] = o.Receiver
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrganizationDataConsentRequest struct {
	value *CreateOrganizationDataConsentRequest
	isSet bool
}

func (v NullableCreateOrganizationDataConsentRequest) Get() *CreateOrganizationDataConsentRequest {
	return v.value
}

func (v *NullableCreateOrganizationDataConsentRequest) Set(val *CreateOrganizationDataConsentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationDataConsentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationDataConsentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationDataConsentRequest(val *CreateOrganizationDataConsentRequest) *NullableCreateOrganizationDataConsentRequest {
	return &NullableCreateOrganizationDataConsentRequest{value: val, isSet: true}
}

func (v NullableCreateOrganizationDataConsentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationDataConsentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



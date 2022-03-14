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

// CreateDataConsentRequest Create individual data consent request
type CreateDataConsentRequest struct {
	// Consent template id
	ConsentTemplateId string `json:"consentTemplateId"`
	Receiver ConsentRequestReceiver `json:"receiver"`
}

// NewCreateDataConsentRequest instantiates a new CreateDataConsentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDataConsentRequest(consentTemplateId string, receiver ConsentRequestReceiver) *CreateDataConsentRequest {
	this := CreateDataConsentRequest{}
	this.ConsentTemplateId = consentTemplateId
	this.Receiver = receiver
	return &this
}

// NewCreateDataConsentRequestWithDefaults instantiates a new CreateDataConsentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDataConsentRequestWithDefaults() *CreateDataConsentRequest {
	this := CreateDataConsentRequest{}
	return &this
}

// GetConsentTemplateId returns the ConsentTemplateId field value
func (o *CreateDataConsentRequest) GetConsentTemplateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsentTemplateId
}

// GetConsentTemplateIdOk returns a tuple with the ConsentTemplateId field value
// and a boolean to check if the value has been set.
func (o *CreateDataConsentRequest) GetConsentTemplateIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConsentTemplateId, true
}

// SetConsentTemplateId sets field value
func (o *CreateDataConsentRequest) SetConsentTemplateId(v string) {
	o.ConsentTemplateId = v
}

// GetReceiver returns the Receiver field value
func (o *CreateDataConsentRequest) GetReceiver() ConsentRequestReceiver {
	if o == nil {
		var ret ConsentRequestReceiver
		return ret
	}

	return o.Receiver
}

// GetReceiverOk returns a tuple with the Receiver field value
// and a boolean to check if the value has been set.
func (o *CreateDataConsentRequest) GetReceiverOk() (*ConsentRequestReceiver, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Receiver, true
}

// SetReceiver sets field value
func (o *CreateDataConsentRequest) SetReceiver(v ConsentRequestReceiver) {
	o.Receiver = v
}

func (o CreateDataConsentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["consentTemplateId"] = o.ConsentTemplateId
	}
	if true {
		toSerialize["receiver"] = o.Receiver
	}
	return json.Marshal(toSerialize)
}

type NullableCreateDataConsentRequest struct {
	value *CreateDataConsentRequest
	isSet bool
}

func (v NullableCreateDataConsentRequest) Get() *CreateDataConsentRequest {
	return v.value
}

func (v *NullableCreateDataConsentRequest) Set(val *CreateDataConsentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDataConsentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDataConsentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDataConsentRequest(val *CreateDataConsentRequest) *NullableCreateDataConsentRequest {
	return &NullableCreateDataConsentRequest{value: val, isSet: true}
}

func (v NullableCreateDataConsentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDataConsentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



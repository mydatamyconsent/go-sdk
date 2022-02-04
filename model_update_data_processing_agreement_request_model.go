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

// UpdateDataProcessingAgreementRequestModel struct for UpdateDataProcessingAgreementRequestModel
type UpdateDataProcessingAgreementRequestModel struct {
	Version string `json:"version"`
	Body string `json:"body"`
	AttachmentUrl string `json:"attachmentUrl"`
}

// NewUpdateDataProcessingAgreementRequestModel instantiates a new UpdateDataProcessingAgreementRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDataProcessingAgreementRequestModel(version string, body string, attachmentUrl string) *UpdateDataProcessingAgreementRequestModel {
	this := UpdateDataProcessingAgreementRequestModel{}
	this.Version = version
	this.Body = body
	this.AttachmentUrl = attachmentUrl
	return &this
}

// NewUpdateDataProcessingAgreementRequestModelWithDefaults instantiates a new UpdateDataProcessingAgreementRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDataProcessingAgreementRequestModelWithDefaults() *UpdateDataProcessingAgreementRequestModel {
	this := UpdateDataProcessingAgreementRequestModel{}
	return &this
}

// GetVersion returns the Version field value
func (o *UpdateDataProcessingAgreementRequestModel) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *UpdateDataProcessingAgreementRequestModel) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *UpdateDataProcessingAgreementRequestModel) SetVersion(v string) {
	o.Version = v
}

// GetBody returns the Body field value
func (o *UpdateDataProcessingAgreementRequestModel) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *UpdateDataProcessingAgreementRequestModel) GetBodyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *UpdateDataProcessingAgreementRequestModel) SetBody(v string) {
	o.Body = v
}

// GetAttachmentUrl returns the AttachmentUrl field value
func (o *UpdateDataProcessingAgreementRequestModel) GetAttachmentUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttachmentUrl
}

// GetAttachmentUrlOk returns a tuple with the AttachmentUrl field value
// and a boolean to check if the value has been set.
func (o *UpdateDataProcessingAgreementRequestModel) GetAttachmentUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AttachmentUrl, true
}

// SetAttachmentUrl sets field value
func (o *UpdateDataProcessingAgreementRequestModel) SetAttachmentUrl(v string) {
	o.AttachmentUrl = v
}

func (o UpdateDataProcessingAgreementRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["body"] = o.Body
	}
	if true {
		toSerialize["attachmentUrl"] = o.AttachmentUrl
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateDataProcessingAgreementRequestModel struct {
	value *UpdateDataProcessingAgreementRequestModel
	isSet bool
}

func (v NullableUpdateDataProcessingAgreementRequestModel) Get() *UpdateDataProcessingAgreementRequestModel {
	return v.value
}

func (v *NullableUpdateDataProcessingAgreementRequestModel) Set(val *UpdateDataProcessingAgreementRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDataProcessingAgreementRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDataProcessingAgreementRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDataProcessingAgreementRequestModel(val *UpdateDataProcessingAgreementRequestModel) *NullableUpdateDataProcessingAgreementRequestModel {
	return &NullableUpdateDataProcessingAgreementRequestModel{value: val, isSet: true}
}

func (v NullableUpdateDataProcessingAgreementRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDataProcessingAgreementRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



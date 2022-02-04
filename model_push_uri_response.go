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

// PushUriResponse struct for PushUriResponse
type PushUriResponse struct {
	ResponseStatus NullableString `json:"responseStatus,omitempty"`
	ResponseMessage NullableString `json:"responseMessage,omitempty"`
	Ns2 NullableString `json:"ns2,omitempty"`
	Ver NullableString `json:"ver,omitempty"`
	Ts NullableString `json:"ts,omitempty"`
	Txn NullableString `json:"txn,omitempty"`
	OrgId NullableString `json:"orgId,omitempty"`
}

// NewPushUriResponse instantiates a new PushUriResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPushUriResponse() *PushUriResponse {
	this := PushUriResponse{}
	return &this
}

// NewPushUriResponseWithDefaults instantiates a new PushUriResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPushUriResponseWithDefaults() *PushUriResponse {
	this := PushUriResponse{}
	return &this
}

// GetResponseStatus returns the ResponseStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PushUriResponse) GetResponseStatus() string {
	if o == nil || o.ResponseStatus.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResponseStatus.Get()
}

// GetResponseStatusOk returns a tuple with the ResponseStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PushUriResponse) GetResponseStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResponseStatus.Get(), o.ResponseStatus.IsSet()
}

// HasResponseStatus returns a boolean if a field has been set.
func (o *PushUriResponse) HasResponseStatus() bool {
	if o != nil && o.ResponseStatus.IsSet() {
		return true
	}

	return false
}

// SetResponseStatus gets a reference to the given NullableString and assigns it to the ResponseStatus field.
func (o *PushUriResponse) SetResponseStatus(v string) {
	o.ResponseStatus.Set(&v)
}
// SetResponseStatusNil sets the value for ResponseStatus to be an explicit nil
func (o *PushUriResponse) SetResponseStatusNil() {
	o.ResponseStatus.Set(nil)
}

// UnsetResponseStatus ensures that no value is present for ResponseStatus, not even an explicit nil
func (o *PushUriResponse) UnsetResponseStatus() {
	o.ResponseStatus.Unset()
}

// GetResponseMessage returns the ResponseMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PushUriResponse) GetResponseMessage() string {
	if o == nil || o.ResponseMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResponseMessage.Get()
}

// GetResponseMessageOk returns a tuple with the ResponseMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PushUriResponse) GetResponseMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResponseMessage.Get(), o.ResponseMessage.IsSet()
}

// HasResponseMessage returns a boolean if a field has been set.
func (o *PushUriResponse) HasResponseMessage() bool {
	if o != nil && o.ResponseMessage.IsSet() {
		return true
	}

	return false
}

// SetResponseMessage gets a reference to the given NullableString and assigns it to the ResponseMessage field.
func (o *PushUriResponse) SetResponseMessage(v string) {
	o.ResponseMessage.Set(&v)
}
// SetResponseMessageNil sets the value for ResponseMessage to be an explicit nil
func (o *PushUriResponse) SetResponseMessageNil() {
	o.ResponseMessage.Set(nil)
}

// UnsetResponseMessage ensures that no value is present for ResponseMessage, not even an explicit nil
func (o *PushUriResponse) UnsetResponseMessage() {
	o.ResponseMessage.Unset()
}

// GetNs2 returns the Ns2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PushUriResponse) GetNs2() string {
	if o == nil || o.Ns2.Get() == nil {
		var ret string
		return ret
	}
	return *o.Ns2.Get()
}

// GetNs2Ok returns a tuple with the Ns2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PushUriResponse) GetNs2Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Ns2.Get(), o.Ns2.IsSet()
}

// HasNs2 returns a boolean if a field has been set.
func (o *PushUriResponse) HasNs2() bool {
	if o != nil && o.Ns2.IsSet() {
		return true
	}

	return false
}

// SetNs2 gets a reference to the given NullableString and assigns it to the Ns2 field.
func (o *PushUriResponse) SetNs2(v string) {
	o.Ns2.Set(&v)
}
// SetNs2Nil sets the value for Ns2 to be an explicit nil
func (o *PushUriResponse) SetNs2Nil() {
	o.Ns2.Set(nil)
}

// UnsetNs2 ensures that no value is present for Ns2, not even an explicit nil
func (o *PushUriResponse) UnsetNs2() {
	o.Ns2.Unset()
}

// GetVer returns the Ver field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PushUriResponse) GetVer() string {
	if o == nil || o.Ver.Get() == nil {
		var ret string
		return ret
	}
	return *o.Ver.Get()
}

// GetVerOk returns a tuple with the Ver field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PushUriResponse) GetVerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Ver.Get(), o.Ver.IsSet()
}

// HasVer returns a boolean if a field has been set.
func (o *PushUriResponse) HasVer() bool {
	if o != nil && o.Ver.IsSet() {
		return true
	}

	return false
}

// SetVer gets a reference to the given NullableString and assigns it to the Ver field.
func (o *PushUriResponse) SetVer(v string) {
	o.Ver.Set(&v)
}
// SetVerNil sets the value for Ver to be an explicit nil
func (o *PushUriResponse) SetVerNil() {
	o.Ver.Set(nil)
}

// UnsetVer ensures that no value is present for Ver, not even an explicit nil
func (o *PushUriResponse) UnsetVer() {
	o.Ver.Unset()
}

// GetTs returns the Ts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PushUriResponse) GetTs() string {
	if o == nil || o.Ts.Get() == nil {
		var ret string
		return ret
	}
	return *o.Ts.Get()
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PushUriResponse) GetTsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Ts.Get(), o.Ts.IsSet()
}

// HasTs returns a boolean if a field has been set.
func (o *PushUriResponse) HasTs() bool {
	if o != nil && o.Ts.IsSet() {
		return true
	}

	return false
}

// SetTs gets a reference to the given NullableString and assigns it to the Ts field.
func (o *PushUriResponse) SetTs(v string) {
	o.Ts.Set(&v)
}
// SetTsNil sets the value for Ts to be an explicit nil
func (o *PushUriResponse) SetTsNil() {
	o.Ts.Set(nil)
}

// UnsetTs ensures that no value is present for Ts, not even an explicit nil
func (o *PushUriResponse) UnsetTs() {
	o.Ts.Unset()
}

// GetTxn returns the Txn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PushUriResponse) GetTxn() string {
	if o == nil || o.Txn.Get() == nil {
		var ret string
		return ret
	}
	return *o.Txn.Get()
}

// GetTxnOk returns a tuple with the Txn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PushUriResponse) GetTxnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Txn.Get(), o.Txn.IsSet()
}

// HasTxn returns a boolean if a field has been set.
func (o *PushUriResponse) HasTxn() bool {
	if o != nil && o.Txn.IsSet() {
		return true
	}

	return false
}

// SetTxn gets a reference to the given NullableString and assigns it to the Txn field.
func (o *PushUriResponse) SetTxn(v string) {
	o.Txn.Set(&v)
}
// SetTxnNil sets the value for Txn to be an explicit nil
func (o *PushUriResponse) SetTxnNil() {
	o.Txn.Set(nil)
}

// UnsetTxn ensures that no value is present for Txn, not even an explicit nil
func (o *PushUriResponse) UnsetTxn() {
	o.Txn.Unset()
}

// GetOrgId returns the OrgId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PushUriResponse) GetOrgId() string {
	if o == nil || o.OrgId.Get() == nil {
		var ret string
		return ret
	}
	return *o.OrgId.Get()
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PushUriResponse) GetOrgIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OrgId.Get(), o.OrgId.IsSet()
}

// HasOrgId returns a boolean if a field has been set.
func (o *PushUriResponse) HasOrgId() bool {
	if o != nil && o.OrgId.IsSet() {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given NullableString and assigns it to the OrgId field.
func (o *PushUriResponse) SetOrgId(v string) {
	o.OrgId.Set(&v)
}
// SetOrgIdNil sets the value for OrgId to be an explicit nil
func (o *PushUriResponse) SetOrgIdNil() {
	o.OrgId.Set(nil)
}

// UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
func (o *PushUriResponse) UnsetOrgId() {
	o.OrgId.Unset()
}

func (o PushUriResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseStatus.IsSet() {
		toSerialize["responseStatus"] = o.ResponseStatus.Get()
	}
	if o.ResponseMessage.IsSet() {
		toSerialize["responseMessage"] = o.ResponseMessage.Get()
	}
	if o.Ns2.IsSet() {
		toSerialize["ns2"] = o.Ns2.Get()
	}
	if o.Ver.IsSet() {
		toSerialize["ver"] = o.Ver.Get()
	}
	if o.Ts.IsSet() {
		toSerialize["ts"] = o.Ts.Get()
	}
	if o.Txn.IsSet() {
		toSerialize["txn"] = o.Txn.Get()
	}
	if o.OrgId.IsSet() {
		toSerialize["orgId"] = o.OrgId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePushUriResponse struct {
	value *PushUriResponse
	isSet bool
}

func (v NullablePushUriResponse) Get() *PushUriResponse {
	return v.value
}

func (v *NullablePushUriResponse) Set(val *PushUriResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePushUriResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePushUriResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePushUriResponse(val *PushUriResponse) *NullablePushUriResponse {
	return &NullablePushUriResponse{value: val, isSet: true}
}

func (v NullablePushUriResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePushUriResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// SipPlanInformation struct for SipPlanInformation
type SipPlanInformation struct {
	Amc *string `json:"amc,omitempty"`
	Registrar *string `json:"registrar,omitempty"`
	Scheme string `json:"scheme"`
	Isin string `json:"isin"`
	FolioNumber *string `json:"folio_number,omitempty"`
	Nav *string `json:"nav,omitempty"`
	DividendType string `json:"dividend_type"`
	CreationDate *time.Time `json:"creation_date,omitempty"`
}

// NewSipPlanInformation instantiates a new SipPlanInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSipPlanInformation(scheme string, isin string, dividendType string) *SipPlanInformation {
	this := SipPlanInformation{}
	this.Scheme = scheme
	this.Isin = isin
	this.DividendType = dividendType
	return &this
}

// NewSipPlanInformationWithDefaults instantiates a new SipPlanInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSipPlanInformationWithDefaults() *SipPlanInformation {
	this := SipPlanInformation{}
	return &this
}

// GetAmc returns the Amc field value if set, zero value otherwise.
func (o *SipPlanInformation) GetAmc() string {
	if o == nil || o.Amc == nil {
		var ret string
		return ret
	}
	return *o.Amc
}

// GetAmcOk returns a tuple with the Amc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SipPlanInformation) GetAmcOk() (*string, bool) {
	if o == nil || o.Amc == nil {
		return nil, false
	}
	return o.Amc, true
}

// HasAmc returns a boolean if a field has been set.
func (o *SipPlanInformation) HasAmc() bool {
	if o != nil && o.Amc != nil {
		return true
	}

	return false
}

// SetAmc gets a reference to the given string and assigns it to the Amc field.
func (o *SipPlanInformation) SetAmc(v string) {
	o.Amc = &v
}

// GetRegistrar returns the Registrar field value if set, zero value otherwise.
func (o *SipPlanInformation) GetRegistrar() string {
	if o == nil || o.Registrar == nil {
		var ret string
		return ret
	}
	return *o.Registrar
}

// GetRegistrarOk returns a tuple with the Registrar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SipPlanInformation) GetRegistrarOk() (*string, bool) {
	if o == nil || o.Registrar == nil {
		return nil, false
	}
	return o.Registrar, true
}

// HasRegistrar returns a boolean if a field has been set.
func (o *SipPlanInformation) HasRegistrar() bool {
	if o != nil && o.Registrar != nil {
		return true
	}

	return false
}

// SetRegistrar gets a reference to the given string and assigns it to the Registrar field.
func (o *SipPlanInformation) SetRegistrar(v string) {
	o.Registrar = &v
}

// GetScheme returns the Scheme field value
func (o *SipPlanInformation) GetScheme() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value
// and a boolean to check if the value has been set.
func (o *SipPlanInformation) GetSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scheme, true
}

// SetScheme sets field value
func (o *SipPlanInformation) SetScheme(v string) {
	o.Scheme = v
}

// GetIsin returns the Isin field value
func (o *SipPlanInformation) GetIsin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Isin
}

// GetIsinOk returns a tuple with the Isin field value
// and a boolean to check if the value has been set.
func (o *SipPlanInformation) GetIsinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Isin, true
}

// SetIsin sets field value
func (o *SipPlanInformation) SetIsin(v string) {
	o.Isin = v
}

// GetFolioNumber returns the FolioNumber field value if set, zero value otherwise.
func (o *SipPlanInformation) GetFolioNumber() string {
	if o == nil || o.FolioNumber == nil {
		var ret string
		return ret
	}
	return *o.FolioNumber
}

// GetFolioNumberOk returns a tuple with the FolioNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SipPlanInformation) GetFolioNumberOk() (*string, bool) {
	if o == nil || o.FolioNumber == nil {
		return nil, false
	}
	return o.FolioNumber, true
}

// HasFolioNumber returns a boolean if a field has been set.
func (o *SipPlanInformation) HasFolioNumber() bool {
	if o != nil && o.FolioNumber != nil {
		return true
	}

	return false
}

// SetFolioNumber gets a reference to the given string and assigns it to the FolioNumber field.
func (o *SipPlanInformation) SetFolioNumber(v string) {
	o.FolioNumber = &v
}

// GetNav returns the Nav field value if set, zero value otherwise.
func (o *SipPlanInformation) GetNav() string {
	if o == nil || o.Nav == nil {
		var ret string
		return ret
	}
	return *o.Nav
}

// GetNavOk returns a tuple with the Nav field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SipPlanInformation) GetNavOk() (*string, bool) {
	if o == nil || o.Nav == nil {
		return nil, false
	}
	return o.Nav, true
}

// HasNav returns a boolean if a field has been set.
func (o *SipPlanInformation) HasNav() bool {
	if o != nil && o.Nav != nil {
		return true
	}

	return false
}

// SetNav gets a reference to the given string and assigns it to the Nav field.
func (o *SipPlanInformation) SetNav(v string) {
	o.Nav = &v
}

// GetDividendType returns the DividendType field value
func (o *SipPlanInformation) GetDividendType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DividendType
}

// GetDividendTypeOk returns a tuple with the DividendType field value
// and a boolean to check if the value has been set.
func (o *SipPlanInformation) GetDividendTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DividendType, true
}

// SetDividendType sets field value
func (o *SipPlanInformation) SetDividendType(v string) {
	o.DividendType = v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *SipPlanInformation) GetCreationDate() time.Time {
	if o == nil || o.CreationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SipPlanInformation) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *SipPlanInformation) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *SipPlanInformation) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

func (o SipPlanInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amc != nil {
		toSerialize["amc"] = o.Amc
	}
	if o.Registrar != nil {
		toSerialize["registrar"] = o.Registrar
	}
	if true {
		toSerialize["scheme"] = o.Scheme
	}
	if true {
		toSerialize["isin"] = o.Isin
	}
	if o.FolioNumber != nil {
		toSerialize["folio_number"] = o.FolioNumber
	}
	if o.Nav != nil {
		toSerialize["nav"] = o.Nav
	}
	if true {
		toSerialize["dividend_type"] = o.DividendType
	}
	if o.CreationDate != nil {
		toSerialize["creation_date"] = o.CreationDate
	}
	return json.Marshal(toSerialize)
}

type NullableSipPlanInformation struct {
	value *SipPlanInformation
	isSet bool
}

func (v NullableSipPlanInformation) Get() *SipPlanInformation {
	return v.value
}

func (v *NullableSipPlanInformation) Set(val *SipPlanInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableSipPlanInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableSipPlanInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSipPlanInformation(val *SipPlanInformation) *NullableSipPlanInformation {
	return &NullableSipPlanInformation{value: val, isSet: true}
}

func (v NullableSipPlanInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSipPlanInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



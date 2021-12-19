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
	"time"
)

// RefreshToken struct for RefreshToken
type RefreshToken struct {
	CreatedBy *string `json:"createdBy,omitempty"`
	CreatedAtUtc *time.Time `json:"createdAtUtc,omitempty"`
	UpdatedBy NullableString `json:"updatedBy,omitempty"`
	UpdatedAtUtc NullableTime `json:"updatedAtUtc,omitempty"`
	CreatedByUser *ApplicationUser `json:"createdByUser,omitempty"`
	UpdatedByUser *ApplicationUser `json:"updatedByUser,omitempty"`
	DeletedBy NullableString `json:"deletedBy,omitempty"`
	DeletedAtUtc NullableTime `json:"deletedAtUtc,omitempty"`
	DeletedByUser *ApplicationUser `json:"deletedByUser,omitempty"`
	Id *string `json:"id,omitempty"`
	InstallationId NullableString `json:"installationId,omitempty"`
	Token NullableString `json:"token,omitempty"`
	AccessToken NullableString `json:"accessToken,omitempty"`
	AccessTokenExpires *time.Time `json:"accessTokenExpires,omitempty"`
	Expires *time.Time `json:"expires,omitempty"`
	IsExpired *bool `json:"isExpired,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	Revoked NullableTime `json:"revoked,omitempty"`
	IsActive *bool `json:"isActive,omitempty"`
	UserId *string `json:"userId,omitempty"`
	User *ApplicationUser `json:"user,omitempty"`
}

// NewRefreshToken instantiates a new RefreshToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshToken() *RefreshToken {
	this := RefreshToken{}
	return &this
}

// NewRefreshTokenWithDefaults instantiates a new RefreshToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshTokenWithDefaults() *RefreshToken {
	this := RefreshToken{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *RefreshToken) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshToken) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *RefreshToken) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *RefreshToken) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedAtUtc returns the CreatedAtUtc field value if set, zero value otherwise.
func (o *RefreshToken) GetCreatedAtUtc() time.Time {
	if o == nil || o.CreatedAtUtc == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAtUtc
}

// GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshToken) GetCreatedAtUtcOk() (*time.Time, bool) {
	if o == nil || o.CreatedAtUtc == nil {
		return nil, false
	}
	return o.CreatedAtUtc, true
}

// HasCreatedAtUtc returns a boolean if a field has been set.
func (o *RefreshToken) HasCreatedAtUtc() bool {
	if o != nil && o.CreatedAtUtc != nil {
		return true
	}

	return false
}

// SetCreatedAtUtc gets a reference to the given time.Time and assigns it to the CreatedAtUtc field.
func (o *RefreshToken) SetCreatedAtUtc(v time.Time) {
	o.CreatedAtUtc = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RefreshToken) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy.Get()
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RefreshToken) GetUpdatedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedBy.Get(), o.UpdatedBy.IsSet()
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *RefreshToken) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy.IsSet() {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given NullableString and assigns it to the UpdatedBy field.
func (o *RefreshToken) SetUpdatedBy(v string) {
	o.UpdatedBy.Set(&v)
}
// SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil
func (o *RefreshToken) SetUpdatedByNil() {
	o.UpdatedBy.Set(nil)
}

// UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
func (o *RefreshToken) UnsetUpdatedBy() {
	o.UpdatedBy.Unset()
}

// GetUpdatedAtUtc returns the UpdatedAtUtc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RefreshToken) GetUpdatedAtUtc() time.Time {
	if o == nil || o.UpdatedAtUtc.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAtUtc.Get()
}

// GetUpdatedAtUtcOk returns a tuple with the UpdatedAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RefreshToken) GetUpdatedAtUtcOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedAtUtc.Get(), o.UpdatedAtUtc.IsSet()
}

// HasUpdatedAtUtc returns a boolean if a field has been set.
func (o *RefreshToken) HasUpdatedAtUtc() bool {
	if o != nil && o.UpdatedAtUtc.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAtUtc gets a reference to the given NullableTime and assigns it to the UpdatedAtUtc field.
func (o *RefreshToken) SetUpdatedAtUtc(v time.Time) {
	o.UpdatedAtUtc.Set(&v)
}
// SetUpdatedAtUtcNil sets the value for UpdatedAtUtc to be an explicit nil
func (o *RefreshToken) SetUpdatedAtUtcNil() {
	o.UpdatedAtUtc.Set(nil)
}

// UnsetUpdatedAtUtc ensures that no value is present for UpdatedAtUtc, not even an explicit nil
func (o *RefreshToken) UnsetUpdatedAtUtc() {
	o.UpdatedAtUtc.Unset()
}

// GetCreatedByUser returns the CreatedByUser field value if set, zero value otherwise.
func (o *RefreshToken) GetCreatedByUser() ApplicationUser {
	if o == nil || o.CreatedByUser == nil {
		var ret ApplicationUser
		return ret
	}
	return *o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshToken) GetCreatedByUserOk() (*ApplicationUser, bool) {
	if o == nil || o.CreatedByUser == nil {
		return nil, false
	}
	return o.CreatedByUser, true
}

// HasCreatedByUser returns a boolean if a field has been set.
func (o *RefreshToken) HasCreatedByUser() bool {
	if o != nil && o.CreatedByUser != nil {
		return true
	}

	return false
}

// SetCreatedByUser gets a reference to the given ApplicationUser and assigns it to the CreatedByUser field.
func (o *RefreshToken) SetCreatedByUser(v ApplicationUser) {
	o.CreatedByUser = &v
}

// GetUpdatedByUser returns the UpdatedByUser field value if set, zero value otherwise.
func (o *RefreshToken) GetUpdatedByUser() ApplicationUser {
	if o == nil || o.UpdatedByUser == nil {
		var ret ApplicationUser
		return ret
	}
	return *o.UpdatedByUser
}

// GetUpdatedByUserOk returns a tuple with the UpdatedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshToken) GetUpdatedByUserOk() (*ApplicationUser, bool) {
	if o == nil || o.UpdatedByUser == nil {
		return nil, false
	}
	return o.UpdatedByUser, true
}

// HasUpdatedByUser returns a boolean if a field has been set.
func (o *RefreshToken) HasUpdatedByUser() bool {
	if o != nil && o.UpdatedByUser != nil {
		return true
	}

	return false
}

// SetUpdatedByUser gets a reference to the given ApplicationUser and assigns it to the UpdatedByUser field.
func (o *RefreshToken) SetUpdatedByUser(v ApplicationUser) {
	o.UpdatedByUser = &v
}

// GetDeletedBy returns the DeletedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RefreshToken) GetDeletedBy() string {
	if o == nil || o.DeletedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeletedBy.Get()
}

// GetDeletedByOk returns a tuple with the DeletedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RefreshToken) GetDeletedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletedBy.Get(), o.DeletedBy.IsSet()
}

// HasDeletedBy returns a boolean if a field has been set.
func (o *RefreshToken) HasDeletedBy() bool {
	if o != nil && o.DeletedBy.IsSet() {
		return true
	}

	return false
}

// SetDeletedBy gets a reference to the given NullableString and assigns it to the DeletedBy field.
func (o *RefreshToken) SetDeletedBy(v string) {
	o.DeletedBy.Set(&v)
}
// SetDeletedByNil sets the value for DeletedBy to be an explicit nil
func (o *RefreshToken) SetDeletedByNil() {
	o.DeletedBy.Set(nil)
}

// UnsetDeletedBy ensures that no value is present for DeletedBy, not even an explicit nil
func (o *RefreshToken) UnsetDeletedBy() {
	o.DeletedBy.Unset()
}

// GetDeletedAtUtc returns the DeletedAtUtc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RefreshToken) GetDeletedAtUtc() time.Time {
	if o == nil || o.DeletedAtUtc.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAtUtc.Get()
}

// GetDeletedAtUtcOk returns a tuple with the DeletedAtUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RefreshToken) GetDeletedAtUtcOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletedAtUtc.Get(), o.DeletedAtUtc.IsSet()
}

// HasDeletedAtUtc returns a boolean if a field has been set.
func (o *RefreshToken) HasDeletedAtUtc() bool {
	if o != nil && o.DeletedAtUtc.IsSet() {
		return true
	}

	return false
}

// SetDeletedAtUtc gets a reference to the given NullableTime and assigns it to the DeletedAtUtc field.
func (o *RefreshToken) SetDeletedAtUtc(v time.Time) {
	o.DeletedAtUtc.Set(&v)
}
// SetDeletedAtUtcNil sets the value for DeletedAtUtc to be an explicit nil
func (o *RefreshToken) SetDeletedAtUtcNil() {
	o.DeletedAtUtc.Set(nil)
}

// UnsetDeletedAtUtc ensures that no value is present for DeletedAtUtc, not even an explicit nil
func (o *RefreshToken) UnsetDeletedAtUtc() {
	o.DeletedAtUtc.Unset()
}

// GetDeletedByUser returns the DeletedByUser field value if set, zero value otherwise.
func (o *RefreshToken) GetDeletedByUser() ApplicationUser {
	if o == nil || o.DeletedByUser == nil {
		var ret ApplicationUser
		return ret
	}
	return *o.DeletedByUser
}

// GetDeletedByUserOk returns a tuple with the DeletedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshToken) GetDeletedByUserOk() (*ApplicationUser, bool) {
	if o == nil || o.DeletedByUser == nil {
		return nil, false
	}
	return o.DeletedByUser, true
}

// HasDeletedByUser returns a boolean if a field has been set.
func (o *RefreshToken) HasDeletedByUser() bool {
	if o != nil && o.DeletedByUser != nil {
		return true
	}

	return false
}

// SetDeletedByUser gets a reference to the given ApplicationUser and assigns it to the DeletedByUser field.
func (o *RefreshToken) SetDeletedByUser(v ApplicationUser) {
	o.DeletedByUser = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RefreshToken) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshToken) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RefreshToken) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RefreshToken) SetId(v string) {
	o.Id = &v
}

// GetInstallationId returns the InstallationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RefreshToken) GetInstallationId() string {
	if o == nil || o.InstallationId.Get() == nil {
		var ret string
		return ret
	}
	return *o.InstallationId.Get()
}

// GetInstallationIdOk returns a tuple with the InstallationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RefreshToken) GetInstallationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InstallationId.Get(), o.InstallationId.IsSet()
}

// HasInstallationId returns a boolean if a field has been set.
func (o *RefreshToken) HasInstallationId() bool {
	if o != nil && o.InstallationId.IsSet() {
		return true
	}

	return false
}

// SetInstallationId gets a reference to the given NullableString and assigns it to the InstallationId field.
func (o *RefreshToken) SetInstallationId(v string) {
	o.InstallationId.Set(&v)
}
// SetInstallationIdNil sets the value for InstallationId to be an explicit nil
func (o *RefreshToken) SetInstallationIdNil() {
	o.InstallationId.Set(nil)
}

// UnsetInstallationId ensures that no value is present for InstallationId, not even an explicit nil
func (o *RefreshToken) UnsetInstallationId() {
	o.InstallationId.Unset()
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RefreshToken) GetToken() string {
	if o == nil || o.Token.Get() == nil {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RefreshToken) GetTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *RefreshToken) HasToken() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *RefreshToken) SetToken(v string) {
	o.Token.Set(&v)
}
// SetTokenNil sets the value for Token to be an explicit nil
func (o *RefreshToken) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *RefreshToken) UnsetToken() {
	o.Token.Unset()
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RefreshToken) GetAccessToken() string {
	if o == nil || o.AccessToken.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccessToken.Get()
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RefreshToken) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccessToken.Get(), o.AccessToken.IsSet()
}

// HasAccessToken returns a boolean if a field has been set.
func (o *RefreshToken) HasAccessToken() bool {
	if o != nil && o.AccessToken.IsSet() {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given NullableString and assigns it to the AccessToken field.
func (o *RefreshToken) SetAccessToken(v string) {
	o.AccessToken.Set(&v)
}
// SetAccessTokenNil sets the value for AccessToken to be an explicit nil
func (o *RefreshToken) SetAccessTokenNil() {
	o.AccessToken.Set(nil)
}

// UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
func (o *RefreshToken) UnsetAccessToken() {
	o.AccessToken.Unset()
}

// GetAccessTokenExpires returns the AccessTokenExpires field value if set, zero value otherwise.
func (o *RefreshToken) GetAccessTokenExpires() time.Time {
	if o == nil || o.AccessTokenExpires == nil {
		var ret time.Time
		return ret
	}
	return *o.AccessTokenExpires
}

// GetAccessTokenExpiresOk returns a tuple with the AccessTokenExpires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshToken) GetAccessTokenExpiresOk() (*time.Time, bool) {
	if o == nil || o.AccessTokenExpires == nil {
		return nil, false
	}
	return o.AccessTokenExpires, true
}

// HasAccessTokenExpires returns a boolean if a field has been set.
func (o *RefreshToken) HasAccessTokenExpires() bool {
	if o != nil && o.AccessTokenExpires != nil {
		return true
	}

	return false
}

// SetAccessTokenExpires gets a reference to the given time.Time and assigns it to the AccessTokenExpires field.
func (o *RefreshToken) SetAccessTokenExpires(v time.Time) {
	o.AccessTokenExpires = &v
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *RefreshToken) GetExpires() time.Time {
	if o == nil || o.Expires == nil {
		var ret time.Time
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshToken) GetExpiresOk() (*time.Time, bool) {
	if o == nil || o.Expires == nil {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *RefreshToken) HasExpires() bool {
	if o != nil && o.Expires != nil {
		return true
	}

	return false
}

// SetExpires gets a reference to the given time.Time and assigns it to the Expires field.
func (o *RefreshToken) SetExpires(v time.Time) {
	o.Expires = &v
}

// GetIsExpired returns the IsExpired field value if set, zero value otherwise.
func (o *RefreshToken) GetIsExpired() bool {
	if o == nil || o.IsExpired == nil {
		var ret bool
		return ret
	}
	return *o.IsExpired
}

// GetIsExpiredOk returns a tuple with the IsExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshToken) GetIsExpiredOk() (*bool, bool) {
	if o == nil || o.IsExpired == nil {
		return nil, false
	}
	return o.IsExpired, true
}

// HasIsExpired returns a boolean if a field has been set.
func (o *RefreshToken) HasIsExpired() bool {
	if o != nil && o.IsExpired != nil {
		return true
	}

	return false
}

// SetIsExpired gets a reference to the given bool and assigns it to the IsExpired field.
func (o *RefreshToken) SetIsExpired(v bool) {
	o.IsExpired = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *RefreshToken) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshToken) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *RefreshToken) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *RefreshToken) SetCreated(v time.Time) {
	o.Created = &v
}

// GetRevoked returns the Revoked field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RefreshToken) GetRevoked() time.Time {
	if o == nil || o.Revoked.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Revoked.Get()
}

// GetRevokedOk returns a tuple with the Revoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RefreshToken) GetRevokedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Revoked.Get(), o.Revoked.IsSet()
}

// HasRevoked returns a boolean if a field has been set.
func (o *RefreshToken) HasRevoked() bool {
	if o != nil && o.Revoked.IsSet() {
		return true
	}

	return false
}

// SetRevoked gets a reference to the given NullableTime and assigns it to the Revoked field.
func (o *RefreshToken) SetRevoked(v time.Time) {
	o.Revoked.Set(&v)
}
// SetRevokedNil sets the value for Revoked to be an explicit nil
func (o *RefreshToken) SetRevokedNil() {
	o.Revoked.Set(nil)
}

// UnsetRevoked ensures that no value is present for Revoked, not even an explicit nil
func (o *RefreshToken) UnsetRevoked() {
	o.Revoked.Unset()
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *RefreshToken) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshToken) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *RefreshToken) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *RefreshToken) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *RefreshToken) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshToken) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *RefreshToken) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *RefreshToken) SetUserId(v string) {
	o.UserId = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *RefreshToken) GetUser() ApplicationUser {
	if o == nil || o.User == nil {
		var ret ApplicationUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshToken) GetUserOk() (*ApplicationUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *RefreshToken) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given ApplicationUser and assigns it to the User field.
func (o *RefreshToken) SetUser(v ApplicationUser) {
	o.User = &v
}

func (o RefreshToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.CreatedAtUtc != nil {
		toSerialize["createdAtUtc"] = o.CreatedAtUtc
	}
	if o.UpdatedBy.IsSet() {
		toSerialize["updatedBy"] = o.UpdatedBy.Get()
	}
	if o.UpdatedAtUtc.IsSet() {
		toSerialize["updatedAtUtc"] = o.UpdatedAtUtc.Get()
	}
	if o.CreatedByUser != nil {
		toSerialize["createdByUser"] = o.CreatedByUser
	}
	if o.UpdatedByUser != nil {
		toSerialize["updatedByUser"] = o.UpdatedByUser
	}
	if o.DeletedBy.IsSet() {
		toSerialize["deletedBy"] = o.DeletedBy.Get()
	}
	if o.DeletedAtUtc.IsSet() {
		toSerialize["deletedAtUtc"] = o.DeletedAtUtc.Get()
	}
	if o.DeletedByUser != nil {
		toSerialize["deletedByUser"] = o.DeletedByUser
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.InstallationId.IsSet() {
		toSerialize["installationId"] = o.InstallationId.Get()
	}
	if o.Token.IsSet() {
		toSerialize["token"] = o.Token.Get()
	}
	if o.AccessToken.IsSet() {
		toSerialize["accessToken"] = o.AccessToken.Get()
	}
	if o.AccessTokenExpires != nil {
		toSerialize["accessTokenExpires"] = o.AccessTokenExpires
	}
	if o.Expires != nil {
		toSerialize["expires"] = o.Expires
	}
	if o.IsExpired != nil {
		toSerialize["isExpired"] = o.IsExpired
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Revoked.IsSet() {
		toSerialize["revoked"] = o.Revoked.Get()
	}
	if o.IsActive != nil {
		toSerialize["isActive"] = o.IsActive
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableRefreshToken struct {
	value *RefreshToken
	isSet bool
}

func (v NullableRefreshToken) Get() *RefreshToken {
	return v.value
}

func (v *NullableRefreshToken) Set(val *RefreshToken) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshToken) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshToken(val *RefreshToken) *NullableRefreshToken {
	return &NullableRefreshToken{value: val, isSet: true}
}

func (v NullableRefreshToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



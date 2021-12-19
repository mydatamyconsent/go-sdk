# RefreshToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedAtUtc** | Pointer to **time.Time** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**CreatedByUser** | Pointer to [**ApplicationUser**](ApplicationUser.md) |  | [optional] 
**UpdatedByUser** | Pointer to [**ApplicationUser**](ApplicationUser.md) |  | [optional] 
**DeletedBy** | Pointer to **NullableString** |  | [optional] 
**DeletedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**DeletedByUser** | Pointer to [**ApplicationUser**](ApplicationUser.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InstallationId** | Pointer to **NullableString** |  | [optional] 
**Token** | Pointer to **NullableString** |  | [optional] 
**AccessToken** | Pointer to **NullableString** |  | [optional] 
**AccessTokenExpires** | Pointer to **time.Time** |  | [optional] 
**Expires** | Pointer to **time.Time** |  | [optional] 
**IsExpired** | Pointer to **bool** |  | [optional] [readonly] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Revoked** | Pointer to **NullableTime** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] [readonly] 
**UserId** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**ApplicationUser**](ApplicationUser.md) |  | [optional] 

## Methods

### NewRefreshToken

`func NewRefreshToken() *RefreshToken`

NewRefreshToken instantiates a new RefreshToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshTokenWithDefaults

`func NewRefreshTokenWithDefaults() *RefreshToken`

NewRefreshTokenWithDefaults instantiates a new RefreshToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *RefreshToken) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RefreshToken) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RefreshToken) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *RefreshToken) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedAtUtc

`func (o *RefreshToken) GetCreatedAtUtc() time.Time`

GetCreatedAtUtc returns the CreatedAtUtc field if non-nil, zero value otherwise.

### GetCreatedAtUtcOk

`func (o *RefreshToken) GetCreatedAtUtcOk() (*time.Time, bool)`

GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtUtc

`func (o *RefreshToken) SetCreatedAtUtc(v time.Time)`

SetCreatedAtUtc sets CreatedAtUtc field to given value.

### HasCreatedAtUtc

`func (o *RefreshToken) HasCreatedAtUtc() bool`

HasCreatedAtUtc returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *RefreshToken) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *RefreshToken) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *RefreshToken) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *RefreshToken) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *RefreshToken) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *RefreshToken) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetUpdatedAtUtc

`func (o *RefreshToken) GetUpdatedAtUtc() time.Time`

GetUpdatedAtUtc returns the UpdatedAtUtc field if non-nil, zero value otherwise.

### GetUpdatedAtUtcOk

`func (o *RefreshToken) GetUpdatedAtUtcOk() (*time.Time, bool)`

GetUpdatedAtUtcOk returns a tuple with the UpdatedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtUtc

`func (o *RefreshToken) SetUpdatedAtUtc(v time.Time)`

SetUpdatedAtUtc sets UpdatedAtUtc field to given value.

### HasUpdatedAtUtc

`func (o *RefreshToken) HasUpdatedAtUtc() bool`

HasUpdatedAtUtc returns a boolean if a field has been set.

### SetUpdatedAtUtcNil

`func (o *RefreshToken) SetUpdatedAtUtcNil(b bool)`

 SetUpdatedAtUtcNil sets the value for UpdatedAtUtc to be an explicit nil

### UnsetUpdatedAtUtc
`func (o *RefreshToken) UnsetUpdatedAtUtc()`

UnsetUpdatedAtUtc ensures that no value is present for UpdatedAtUtc, not even an explicit nil
### GetCreatedByUser

`func (o *RefreshToken) GetCreatedByUser() ApplicationUser`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *RefreshToken) GetCreatedByUserOk() (*ApplicationUser, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *RefreshToken) SetCreatedByUser(v ApplicationUser)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *RefreshToken) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### GetUpdatedByUser

`func (o *RefreshToken) GetUpdatedByUser() ApplicationUser`

GetUpdatedByUser returns the UpdatedByUser field if non-nil, zero value otherwise.

### GetUpdatedByUserOk

`func (o *RefreshToken) GetUpdatedByUserOk() (*ApplicationUser, bool)`

GetUpdatedByUserOk returns a tuple with the UpdatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUser

`func (o *RefreshToken) SetUpdatedByUser(v ApplicationUser)`

SetUpdatedByUser sets UpdatedByUser field to given value.

### HasUpdatedByUser

`func (o *RefreshToken) HasUpdatedByUser() bool`

HasUpdatedByUser returns a boolean if a field has been set.

### GetDeletedBy

`func (o *RefreshToken) GetDeletedBy() string`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *RefreshToken) GetDeletedByOk() (*string, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *RefreshToken) SetDeletedBy(v string)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *RefreshToken) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### SetDeletedByNil

`func (o *RefreshToken) SetDeletedByNil(b bool)`

 SetDeletedByNil sets the value for DeletedBy to be an explicit nil

### UnsetDeletedBy
`func (o *RefreshToken) UnsetDeletedBy()`

UnsetDeletedBy ensures that no value is present for DeletedBy, not even an explicit nil
### GetDeletedAtUtc

`func (o *RefreshToken) GetDeletedAtUtc() time.Time`

GetDeletedAtUtc returns the DeletedAtUtc field if non-nil, zero value otherwise.

### GetDeletedAtUtcOk

`func (o *RefreshToken) GetDeletedAtUtcOk() (*time.Time, bool)`

GetDeletedAtUtcOk returns a tuple with the DeletedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAtUtc

`func (o *RefreshToken) SetDeletedAtUtc(v time.Time)`

SetDeletedAtUtc sets DeletedAtUtc field to given value.

### HasDeletedAtUtc

`func (o *RefreshToken) HasDeletedAtUtc() bool`

HasDeletedAtUtc returns a boolean if a field has been set.

### SetDeletedAtUtcNil

`func (o *RefreshToken) SetDeletedAtUtcNil(b bool)`

 SetDeletedAtUtcNil sets the value for DeletedAtUtc to be an explicit nil

### UnsetDeletedAtUtc
`func (o *RefreshToken) UnsetDeletedAtUtc()`

UnsetDeletedAtUtc ensures that no value is present for DeletedAtUtc, not even an explicit nil
### GetDeletedByUser

`func (o *RefreshToken) GetDeletedByUser() ApplicationUser`

GetDeletedByUser returns the DeletedByUser field if non-nil, zero value otherwise.

### GetDeletedByUserOk

`func (o *RefreshToken) GetDeletedByUserOk() (*ApplicationUser, bool)`

GetDeletedByUserOk returns a tuple with the DeletedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedByUser

`func (o *RefreshToken) SetDeletedByUser(v ApplicationUser)`

SetDeletedByUser sets DeletedByUser field to given value.

### HasDeletedByUser

`func (o *RefreshToken) HasDeletedByUser() bool`

HasDeletedByUser returns a boolean if a field has been set.

### GetId

`func (o *RefreshToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RefreshToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RefreshToken) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RefreshToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstallationId

`func (o *RefreshToken) GetInstallationId() string`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *RefreshToken) GetInstallationIdOk() (*string, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *RefreshToken) SetInstallationId(v string)`

SetInstallationId sets InstallationId field to given value.

### HasInstallationId

`func (o *RefreshToken) HasInstallationId() bool`

HasInstallationId returns a boolean if a field has been set.

### SetInstallationIdNil

`func (o *RefreshToken) SetInstallationIdNil(b bool)`

 SetInstallationIdNil sets the value for InstallationId to be an explicit nil

### UnsetInstallationId
`func (o *RefreshToken) UnsetInstallationId()`

UnsetInstallationId ensures that no value is present for InstallationId, not even an explicit nil
### GetToken

`func (o *RefreshToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RefreshToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RefreshToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RefreshToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *RefreshToken) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *RefreshToken) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetAccessToken

`func (o *RefreshToken) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *RefreshToken) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *RefreshToken) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *RefreshToken) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### SetAccessTokenNil

`func (o *RefreshToken) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *RefreshToken) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
### GetAccessTokenExpires

`func (o *RefreshToken) GetAccessTokenExpires() time.Time`

GetAccessTokenExpires returns the AccessTokenExpires field if non-nil, zero value otherwise.

### GetAccessTokenExpiresOk

`func (o *RefreshToken) GetAccessTokenExpiresOk() (*time.Time, bool)`

GetAccessTokenExpiresOk returns a tuple with the AccessTokenExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenExpires

`func (o *RefreshToken) SetAccessTokenExpires(v time.Time)`

SetAccessTokenExpires sets AccessTokenExpires field to given value.

### HasAccessTokenExpires

`func (o *RefreshToken) HasAccessTokenExpires() bool`

HasAccessTokenExpires returns a boolean if a field has been set.

### GetExpires

`func (o *RefreshToken) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *RefreshToken) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *RefreshToken) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *RefreshToken) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetIsExpired

`func (o *RefreshToken) GetIsExpired() bool`

GetIsExpired returns the IsExpired field if non-nil, zero value otherwise.

### GetIsExpiredOk

`func (o *RefreshToken) GetIsExpiredOk() (*bool, bool)`

GetIsExpiredOk returns a tuple with the IsExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpired

`func (o *RefreshToken) SetIsExpired(v bool)`

SetIsExpired sets IsExpired field to given value.

### HasIsExpired

`func (o *RefreshToken) HasIsExpired() bool`

HasIsExpired returns a boolean if a field has been set.

### GetCreated

`func (o *RefreshToken) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RefreshToken) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RefreshToken) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RefreshToken) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetRevoked

`func (o *RefreshToken) GetRevoked() time.Time`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *RefreshToken) GetRevokedOk() (*time.Time, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *RefreshToken) SetRevoked(v time.Time)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *RefreshToken) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.

### SetRevokedNil

`func (o *RefreshToken) SetRevokedNil(b bool)`

 SetRevokedNil sets the value for Revoked to be an explicit nil

### UnsetRevoked
`func (o *RefreshToken) UnsetRevoked()`

UnsetRevoked ensures that no value is present for Revoked, not even an explicit nil
### GetIsActive

`func (o *RefreshToken) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *RefreshToken) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *RefreshToken) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *RefreshToken) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetUserId

`func (o *RefreshToken) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RefreshToken) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RefreshToken) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *RefreshToken) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUser

`func (o *RefreshToken) GetUser() ApplicationUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RefreshToken) GetUserOk() (*ApplicationUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RefreshToken) SetUser(v ApplicationUser)`

SetUser sets User field to given value.

### HasUser

`func (o *RefreshToken) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ApplicationUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **NullableString** |  | [optional] 
**NormalizedUserName** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**NormalizedEmail** | Pointer to **NullableString** |  | [optional] 
**EmailConfirmed** | Pointer to **bool** |  | [optional] 
**PasswordHash** | Pointer to **NullableString** |  | [optional] 
**SecurityStamp** | Pointer to **NullableString** |  | [optional] 
**ConcurrencyStamp** | Pointer to **NullableString** |  | [optional] 
**PhoneNumber** | Pointer to **NullableString** |  | [optional] 
**PhoneNumberConfirmed** | Pointer to **bool** |  | [optional] 
**TwoFactorEnabled** | Pointer to **bool** |  | [optional] 
**LockoutEnd** | Pointer to **NullableTime** |  | [optional] 
**LockoutEnabled** | Pointer to **bool** |  | [optional] 
**AccessFailedCount** | Pointer to **int32** |  | [optional] 
**FirstName** | Pointer to **NullableString** |  | [optional] 
**MiddleName** | Pointer to **NullableString** |  | [optional] 
**LastName** | Pointer to **NullableString** |  | [optional] 
**FullName** | Pointer to **NullableString** |  | [optional] [readonly] 
**Gender** | Pointer to [**Gender**](Gender.md) |  | [optional] 
**DateOfBirth** | Pointer to **time.Time** |  | [optional] 
**CountryId** | Pointer to **string** |  | [optional] 
**PostCode** | Pointer to **NullableString** |  | [optional] 
**ReferredBy** | Pointer to **NullableString** |  | [optional] 
**Language** | Pointer to **NullableString** |  | [optional] 
**Theme** | Pointer to [**Theme**](Theme.md) |  | [optional] 
**Designation** | Pointer to **NullableString** |  | [optional] 
**IsMarkedForDeletion** | Pointer to **bool** |  | [optional] 
**HardDeleteDate** | Pointer to **NullableTime** |  | [optional] 
**SecurityPin** | Pointer to **NullableString** |  | [optional] 
**PhotoUrl** | Pointer to **NullableString** |  | [optional] 
**ReferralCode** | Pointer to **NullableString** |  | [optional] 
**RecoveryToken** | Pointer to **NullableString** |  | [optional] 
**DigiLockerLastSyncDate** | Pointer to **NullableTime** |  | [optional] 
**RefreshTokens** | Pointer to [**[]RefreshToken**](RefreshToken.md) |  | [optional] 
**Country** | Pointer to [**Country**](Country.md) |  | [optional] 
**ReferredByUser** | Pointer to [**ApplicationUser**](ApplicationUser.md) |  | [optional] 

## Methods

### NewApplicationUser

`func NewApplicationUser() *ApplicationUser`

NewApplicationUser instantiates a new ApplicationUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationUserWithDefaults

`func NewApplicationUserWithDefaults() *ApplicationUser`

NewApplicationUserWithDefaults instantiates a new ApplicationUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserName

`func (o *ApplicationUser) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ApplicationUser) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ApplicationUser) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ApplicationUser) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *ApplicationUser) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *ApplicationUser) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetNormalizedUserName

`func (o *ApplicationUser) GetNormalizedUserName() string`

GetNormalizedUserName returns the NormalizedUserName field if non-nil, zero value otherwise.

### GetNormalizedUserNameOk

`func (o *ApplicationUser) GetNormalizedUserNameOk() (*string, bool)`

GetNormalizedUserNameOk returns a tuple with the NormalizedUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizedUserName

`func (o *ApplicationUser) SetNormalizedUserName(v string)`

SetNormalizedUserName sets NormalizedUserName field to given value.

### HasNormalizedUserName

`func (o *ApplicationUser) HasNormalizedUserName() bool`

HasNormalizedUserName returns a boolean if a field has been set.

### SetNormalizedUserNameNil

`func (o *ApplicationUser) SetNormalizedUserNameNil(b bool)`

 SetNormalizedUserNameNil sets the value for NormalizedUserName to be an explicit nil

### UnsetNormalizedUserName
`func (o *ApplicationUser) UnsetNormalizedUserName()`

UnsetNormalizedUserName ensures that no value is present for NormalizedUserName, not even an explicit nil
### GetEmail

`func (o *ApplicationUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ApplicationUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ApplicationUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ApplicationUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ApplicationUser) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ApplicationUser) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetNormalizedEmail

`func (o *ApplicationUser) GetNormalizedEmail() string`

GetNormalizedEmail returns the NormalizedEmail field if non-nil, zero value otherwise.

### GetNormalizedEmailOk

`func (o *ApplicationUser) GetNormalizedEmailOk() (*string, bool)`

GetNormalizedEmailOk returns a tuple with the NormalizedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizedEmail

`func (o *ApplicationUser) SetNormalizedEmail(v string)`

SetNormalizedEmail sets NormalizedEmail field to given value.

### HasNormalizedEmail

`func (o *ApplicationUser) HasNormalizedEmail() bool`

HasNormalizedEmail returns a boolean if a field has been set.

### SetNormalizedEmailNil

`func (o *ApplicationUser) SetNormalizedEmailNil(b bool)`

 SetNormalizedEmailNil sets the value for NormalizedEmail to be an explicit nil

### UnsetNormalizedEmail
`func (o *ApplicationUser) UnsetNormalizedEmail()`

UnsetNormalizedEmail ensures that no value is present for NormalizedEmail, not even an explicit nil
### GetEmailConfirmed

`func (o *ApplicationUser) GetEmailConfirmed() bool`

GetEmailConfirmed returns the EmailConfirmed field if non-nil, zero value otherwise.

### GetEmailConfirmedOk

`func (o *ApplicationUser) GetEmailConfirmedOk() (*bool, bool)`

GetEmailConfirmedOk returns a tuple with the EmailConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConfirmed

`func (o *ApplicationUser) SetEmailConfirmed(v bool)`

SetEmailConfirmed sets EmailConfirmed field to given value.

### HasEmailConfirmed

`func (o *ApplicationUser) HasEmailConfirmed() bool`

HasEmailConfirmed returns a boolean if a field has been set.

### GetPasswordHash

`func (o *ApplicationUser) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *ApplicationUser) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *ApplicationUser) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *ApplicationUser) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### SetPasswordHashNil

`func (o *ApplicationUser) SetPasswordHashNil(b bool)`

 SetPasswordHashNil sets the value for PasswordHash to be an explicit nil

### UnsetPasswordHash
`func (o *ApplicationUser) UnsetPasswordHash()`

UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil
### GetSecurityStamp

`func (o *ApplicationUser) GetSecurityStamp() string`

GetSecurityStamp returns the SecurityStamp field if non-nil, zero value otherwise.

### GetSecurityStampOk

`func (o *ApplicationUser) GetSecurityStampOk() (*string, bool)`

GetSecurityStampOk returns a tuple with the SecurityStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityStamp

`func (o *ApplicationUser) SetSecurityStamp(v string)`

SetSecurityStamp sets SecurityStamp field to given value.

### HasSecurityStamp

`func (o *ApplicationUser) HasSecurityStamp() bool`

HasSecurityStamp returns a boolean if a field has been set.

### SetSecurityStampNil

`func (o *ApplicationUser) SetSecurityStampNil(b bool)`

 SetSecurityStampNil sets the value for SecurityStamp to be an explicit nil

### UnsetSecurityStamp
`func (o *ApplicationUser) UnsetSecurityStamp()`

UnsetSecurityStamp ensures that no value is present for SecurityStamp, not even an explicit nil
### GetConcurrencyStamp

`func (o *ApplicationUser) GetConcurrencyStamp() string`

GetConcurrencyStamp returns the ConcurrencyStamp field if non-nil, zero value otherwise.

### GetConcurrencyStampOk

`func (o *ApplicationUser) GetConcurrencyStampOk() (*string, bool)`

GetConcurrencyStampOk returns a tuple with the ConcurrencyStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyStamp

`func (o *ApplicationUser) SetConcurrencyStamp(v string)`

SetConcurrencyStamp sets ConcurrencyStamp field to given value.

### HasConcurrencyStamp

`func (o *ApplicationUser) HasConcurrencyStamp() bool`

HasConcurrencyStamp returns a boolean if a field has been set.

### SetConcurrencyStampNil

`func (o *ApplicationUser) SetConcurrencyStampNil(b bool)`

 SetConcurrencyStampNil sets the value for ConcurrencyStamp to be an explicit nil

### UnsetConcurrencyStamp
`func (o *ApplicationUser) UnsetConcurrencyStamp()`

UnsetConcurrencyStamp ensures that no value is present for ConcurrencyStamp, not even an explicit nil
### GetPhoneNumber

`func (o *ApplicationUser) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ApplicationUser) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ApplicationUser) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ApplicationUser) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *ApplicationUser) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *ApplicationUser) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetPhoneNumberConfirmed

`func (o *ApplicationUser) GetPhoneNumberConfirmed() bool`

GetPhoneNumberConfirmed returns the PhoneNumberConfirmed field if non-nil, zero value otherwise.

### GetPhoneNumberConfirmedOk

`func (o *ApplicationUser) GetPhoneNumberConfirmedOk() (*bool, bool)`

GetPhoneNumberConfirmedOk returns a tuple with the PhoneNumberConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberConfirmed

`func (o *ApplicationUser) SetPhoneNumberConfirmed(v bool)`

SetPhoneNumberConfirmed sets PhoneNumberConfirmed field to given value.

### HasPhoneNumberConfirmed

`func (o *ApplicationUser) HasPhoneNumberConfirmed() bool`

HasPhoneNumberConfirmed returns a boolean if a field has been set.

### GetTwoFactorEnabled

`func (o *ApplicationUser) GetTwoFactorEnabled() bool`

GetTwoFactorEnabled returns the TwoFactorEnabled field if non-nil, zero value otherwise.

### GetTwoFactorEnabledOk

`func (o *ApplicationUser) GetTwoFactorEnabledOk() (*bool, bool)`

GetTwoFactorEnabledOk returns a tuple with the TwoFactorEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorEnabled

`func (o *ApplicationUser) SetTwoFactorEnabled(v bool)`

SetTwoFactorEnabled sets TwoFactorEnabled field to given value.

### HasTwoFactorEnabled

`func (o *ApplicationUser) HasTwoFactorEnabled() bool`

HasTwoFactorEnabled returns a boolean if a field has been set.

### GetLockoutEnd

`func (o *ApplicationUser) GetLockoutEnd() time.Time`

GetLockoutEnd returns the LockoutEnd field if non-nil, zero value otherwise.

### GetLockoutEndOk

`func (o *ApplicationUser) GetLockoutEndOk() (*time.Time, bool)`

GetLockoutEndOk returns a tuple with the LockoutEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutEnd

`func (o *ApplicationUser) SetLockoutEnd(v time.Time)`

SetLockoutEnd sets LockoutEnd field to given value.

### HasLockoutEnd

`func (o *ApplicationUser) HasLockoutEnd() bool`

HasLockoutEnd returns a boolean if a field has been set.

### SetLockoutEndNil

`func (o *ApplicationUser) SetLockoutEndNil(b bool)`

 SetLockoutEndNil sets the value for LockoutEnd to be an explicit nil

### UnsetLockoutEnd
`func (o *ApplicationUser) UnsetLockoutEnd()`

UnsetLockoutEnd ensures that no value is present for LockoutEnd, not even an explicit nil
### GetLockoutEnabled

`func (o *ApplicationUser) GetLockoutEnabled() bool`

GetLockoutEnabled returns the LockoutEnabled field if non-nil, zero value otherwise.

### GetLockoutEnabledOk

`func (o *ApplicationUser) GetLockoutEnabledOk() (*bool, bool)`

GetLockoutEnabledOk returns a tuple with the LockoutEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutEnabled

`func (o *ApplicationUser) SetLockoutEnabled(v bool)`

SetLockoutEnabled sets LockoutEnabled field to given value.

### HasLockoutEnabled

`func (o *ApplicationUser) HasLockoutEnabled() bool`

HasLockoutEnabled returns a boolean if a field has been set.

### GetAccessFailedCount

`func (o *ApplicationUser) GetAccessFailedCount() int32`

GetAccessFailedCount returns the AccessFailedCount field if non-nil, zero value otherwise.

### GetAccessFailedCountOk

`func (o *ApplicationUser) GetAccessFailedCountOk() (*int32, bool)`

GetAccessFailedCountOk returns a tuple with the AccessFailedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessFailedCount

`func (o *ApplicationUser) SetAccessFailedCount(v int32)`

SetAccessFailedCount sets AccessFailedCount field to given value.

### HasAccessFailedCount

`func (o *ApplicationUser) HasAccessFailedCount() bool`

HasAccessFailedCount returns a boolean if a field has been set.

### GetFirstName

`func (o *ApplicationUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ApplicationUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ApplicationUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ApplicationUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *ApplicationUser) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *ApplicationUser) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetMiddleName

`func (o *ApplicationUser) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *ApplicationUser) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *ApplicationUser) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *ApplicationUser) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *ApplicationUser) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *ApplicationUser) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetLastName

`func (o *ApplicationUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ApplicationUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ApplicationUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ApplicationUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ApplicationUser) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ApplicationUser) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetFullName

`func (o *ApplicationUser) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ApplicationUser) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ApplicationUser) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *ApplicationUser) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *ApplicationUser) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *ApplicationUser) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetGender

`func (o *ApplicationUser) GetGender() Gender`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *ApplicationUser) GetGenderOk() (*Gender, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *ApplicationUser) SetGender(v Gender)`

SetGender sets Gender field to given value.

### HasGender

`func (o *ApplicationUser) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *ApplicationUser) GetDateOfBirth() time.Time`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *ApplicationUser) GetDateOfBirthOk() (*time.Time, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *ApplicationUser) SetDateOfBirth(v time.Time)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *ApplicationUser) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetCountryId

`func (o *ApplicationUser) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *ApplicationUser) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *ApplicationUser) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *ApplicationUser) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### GetPostCode

`func (o *ApplicationUser) GetPostCode() string`

GetPostCode returns the PostCode field if non-nil, zero value otherwise.

### GetPostCodeOk

`func (o *ApplicationUser) GetPostCodeOk() (*string, bool)`

GetPostCodeOk returns a tuple with the PostCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCode

`func (o *ApplicationUser) SetPostCode(v string)`

SetPostCode sets PostCode field to given value.

### HasPostCode

`func (o *ApplicationUser) HasPostCode() bool`

HasPostCode returns a boolean if a field has been set.

### SetPostCodeNil

`func (o *ApplicationUser) SetPostCodeNil(b bool)`

 SetPostCodeNil sets the value for PostCode to be an explicit nil

### UnsetPostCode
`func (o *ApplicationUser) UnsetPostCode()`

UnsetPostCode ensures that no value is present for PostCode, not even an explicit nil
### GetReferredBy

`func (o *ApplicationUser) GetReferredBy() string`

GetReferredBy returns the ReferredBy field if non-nil, zero value otherwise.

### GetReferredByOk

`func (o *ApplicationUser) GetReferredByOk() (*string, bool)`

GetReferredByOk returns a tuple with the ReferredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredBy

`func (o *ApplicationUser) SetReferredBy(v string)`

SetReferredBy sets ReferredBy field to given value.

### HasReferredBy

`func (o *ApplicationUser) HasReferredBy() bool`

HasReferredBy returns a boolean if a field has been set.

### SetReferredByNil

`func (o *ApplicationUser) SetReferredByNil(b bool)`

 SetReferredByNil sets the value for ReferredBy to be an explicit nil

### UnsetReferredBy
`func (o *ApplicationUser) UnsetReferredBy()`

UnsetReferredBy ensures that no value is present for ReferredBy, not even an explicit nil
### GetLanguage

`func (o *ApplicationUser) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ApplicationUser) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ApplicationUser) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ApplicationUser) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *ApplicationUser) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *ApplicationUser) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetTheme

`func (o *ApplicationUser) GetTheme() Theme`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *ApplicationUser) GetThemeOk() (*Theme, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *ApplicationUser) SetTheme(v Theme)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *ApplicationUser) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetDesignation

`func (o *ApplicationUser) GetDesignation() string`

GetDesignation returns the Designation field if non-nil, zero value otherwise.

### GetDesignationOk

`func (o *ApplicationUser) GetDesignationOk() (*string, bool)`

GetDesignationOk returns a tuple with the Designation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignation

`func (o *ApplicationUser) SetDesignation(v string)`

SetDesignation sets Designation field to given value.

### HasDesignation

`func (o *ApplicationUser) HasDesignation() bool`

HasDesignation returns a boolean if a field has been set.

### SetDesignationNil

`func (o *ApplicationUser) SetDesignationNil(b bool)`

 SetDesignationNil sets the value for Designation to be an explicit nil

### UnsetDesignation
`func (o *ApplicationUser) UnsetDesignation()`

UnsetDesignation ensures that no value is present for Designation, not even an explicit nil
### GetIsMarkedForDeletion

`func (o *ApplicationUser) GetIsMarkedForDeletion() bool`

GetIsMarkedForDeletion returns the IsMarkedForDeletion field if non-nil, zero value otherwise.

### GetIsMarkedForDeletionOk

`func (o *ApplicationUser) GetIsMarkedForDeletionOk() (*bool, bool)`

GetIsMarkedForDeletionOk returns a tuple with the IsMarkedForDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarkedForDeletion

`func (o *ApplicationUser) SetIsMarkedForDeletion(v bool)`

SetIsMarkedForDeletion sets IsMarkedForDeletion field to given value.

### HasIsMarkedForDeletion

`func (o *ApplicationUser) HasIsMarkedForDeletion() bool`

HasIsMarkedForDeletion returns a boolean if a field has been set.

### GetHardDeleteDate

`func (o *ApplicationUser) GetHardDeleteDate() time.Time`

GetHardDeleteDate returns the HardDeleteDate field if non-nil, zero value otherwise.

### GetHardDeleteDateOk

`func (o *ApplicationUser) GetHardDeleteDateOk() (*time.Time, bool)`

GetHardDeleteDateOk returns a tuple with the HardDeleteDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardDeleteDate

`func (o *ApplicationUser) SetHardDeleteDate(v time.Time)`

SetHardDeleteDate sets HardDeleteDate field to given value.

### HasHardDeleteDate

`func (o *ApplicationUser) HasHardDeleteDate() bool`

HasHardDeleteDate returns a boolean if a field has been set.

### SetHardDeleteDateNil

`func (o *ApplicationUser) SetHardDeleteDateNil(b bool)`

 SetHardDeleteDateNil sets the value for HardDeleteDate to be an explicit nil

### UnsetHardDeleteDate
`func (o *ApplicationUser) UnsetHardDeleteDate()`

UnsetHardDeleteDate ensures that no value is present for HardDeleteDate, not even an explicit nil
### GetSecurityPin

`func (o *ApplicationUser) GetSecurityPin() string`

GetSecurityPin returns the SecurityPin field if non-nil, zero value otherwise.

### GetSecurityPinOk

`func (o *ApplicationUser) GetSecurityPinOk() (*string, bool)`

GetSecurityPinOk returns a tuple with the SecurityPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPin

`func (o *ApplicationUser) SetSecurityPin(v string)`

SetSecurityPin sets SecurityPin field to given value.

### HasSecurityPin

`func (o *ApplicationUser) HasSecurityPin() bool`

HasSecurityPin returns a boolean if a field has been set.

### SetSecurityPinNil

`func (o *ApplicationUser) SetSecurityPinNil(b bool)`

 SetSecurityPinNil sets the value for SecurityPin to be an explicit nil

### UnsetSecurityPin
`func (o *ApplicationUser) UnsetSecurityPin()`

UnsetSecurityPin ensures that no value is present for SecurityPin, not even an explicit nil
### GetPhotoUrl

`func (o *ApplicationUser) GetPhotoUrl() string`

GetPhotoUrl returns the PhotoUrl field if non-nil, zero value otherwise.

### GetPhotoUrlOk

`func (o *ApplicationUser) GetPhotoUrlOk() (*string, bool)`

GetPhotoUrlOk returns a tuple with the PhotoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoUrl

`func (o *ApplicationUser) SetPhotoUrl(v string)`

SetPhotoUrl sets PhotoUrl field to given value.

### HasPhotoUrl

`func (o *ApplicationUser) HasPhotoUrl() bool`

HasPhotoUrl returns a boolean if a field has been set.

### SetPhotoUrlNil

`func (o *ApplicationUser) SetPhotoUrlNil(b bool)`

 SetPhotoUrlNil sets the value for PhotoUrl to be an explicit nil

### UnsetPhotoUrl
`func (o *ApplicationUser) UnsetPhotoUrl()`

UnsetPhotoUrl ensures that no value is present for PhotoUrl, not even an explicit nil
### GetReferralCode

`func (o *ApplicationUser) GetReferralCode() string`

GetReferralCode returns the ReferralCode field if non-nil, zero value otherwise.

### GetReferralCodeOk

`func (o *ApplicationUser) GetReferralCodeOk() (*string, bool)`

GetReferralCodeOk returns a tuple with the ReferralCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralCode

`func (o *ApplicationUser) SetReferralCode(v string)`

SetReferralCode sets ReferralCode field to given value.

### HasReferralCode

`func (o *ApplicationUser) HasReferralCode() bool`

HasReferralCode returns a boolean if a field has been set.

### SetReferralCodeNil

`func (o *ApplicationUser) SetReferralCodeNil(b bool)`

 SetReferralCodeNil sets the value for ReferralCode to be an explicit nil

### UnsetReferralCode
`func (o *ApplicationUser) UnsetReferralCode()`

UnsetReferralCode ensures that no value is present for ReferralCode, not even an explicit nil
### GetRecoveryToken

`func (o *ApplicationUser) GetRecoveryToken() string`

GetRecoveryToken returns the RecoveryToken field if non-nil, zero value otherwise.

### GetRecoveryTokenOk

`func (o *ApplicationUser) GetRecoveryTokenOk() (*string, bool)`

GetRecoveryTokenOk returns a tuple with the RecoveryToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryToken

`func (o *ApplicationUser) SetRecoveryToken(v string)`

SetRecoveryToken sets RecoveryToken field to given value.

### HasRecoveryToken

`func (o *ApplicationUser) HasRecoveryToken() bool`

HasRecoveryToken returns a boolean if a field has been set.

### SetRecoveryTokenNil

`func (o *ApplicationUser) SetRecoveryTokenNil(b bool)`

 SetRecoveryTokenNil sets the value for RecoveryToken to be an explicit nil

### UnsetRecoveryToken
`func (o *ApplicationUser) UnsetRecoveryToken()`

UnsetRecoveryToken ensures that no value is present for RecoveryToken, not even an explicit nil
### GetDigiLockerLastSyncDate

`func (o *ApplicationUser) GetDigiLockerLastSyncDate() time.Time`

GetDigiLockerLastSyncDate returns the DigiLockerLastSyncDate field if non-nil, zero value otherwise.

### GetDigiLockerLastSyncDateOk

`func (o *ApplicationUser) GetDigiLockerLastSyncDateOk() (*time.Time, bool)`

GetDigiLockerLastSyncDateOk returns a tuple with the DigiLockerLastSyncDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigiLockerLastSyncDate

`func (o *ApplicationUser) SetDigiLockerLastSyncDate(v time.Time)`

SetDigiLockerLastSyncDate sets DigiLockerLastSyncDate field to given value.

### HasDigiLockerLastSyncDate

`func (o *ApplicationUser) HasDigiLockerLastSyncDate() bool`

HasDigiLockerLastSyncDate returns a boolean if a field has been set.

### SetDigiLockerLastSyncDateNil

`func (o *ApplicationUser) SetDigiLockerLastSyncDateNil(b bool)`

 SetDigiLockerLastSyncDateNil sets the value for DigiLockerLastSyncDate to be an explicit nil

### UnsetDigiLockerLastSyncDate
`func (o *ApplicationUser) UnsetDigiLockerLastSyncDate()`

UnsetDigiLockerLastSyncDate ensures that no value is present for DigiLockerLastSyncDate, not even an explicit nil
### GetRefreshTokens

`func (o *ApplicationUser) GetRefreshTokens() []RefreshToken`

GetRefreshTokens returns the RefreshTokens field if non-nil, zero value otherwise.

### GetRefreshTokensOk

`func (o *ApplicationUser) GetRefreshTokensOk() (*[]RefreshToken, bool)`

GetRefreshTokensOk returns a tuple with the RefreshTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokens

`func (o *ApplicationUser) SetRefreshTokens(v []RefreshToken)`

SetRefreshTokens sets RefreshTokens field to given value.

### HasRefreshTokens

`func (o *ApplicationUser) HasRefreshTokens() bool`

HasRefreshTokens returns a boolean if a field has been set.

### SetRefreshTokensNil

`func (o *ApplicationUser) SetRefreshTokensNil(b bool)`

 SetRefreshTokensNil sets the value for RefreshTokens to be an explicit nil

### UnsetRefreshTokens
`func (o *ApplicationUser) UnsetRefreshTokens()`

UnsetRefreshTokens ensures that no value is present for RefreshTokens, not even an explicit nil
### GetCountry

`func (o *ApplicationUser) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ApplicationUser) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ApplicationUser) SetCountry(v Country)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ApplicationUser) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetReferredByUser

`func (o *ApplicationUser) GetReferredByUser() ApplicationUser`

GetReferredByUser returns the ReferredByUser field if non-nil, zero value otherwise.

### GetReferredByUserOk

`func (o *ApplicationUser) GetReferredByUserOk() (*ApplicationUser, bool)`

GetReferredByUserOk returns a tuple with the ReferredByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredByUser

`func (o *ApplicationUser) SetReferredByUser(v ApplicationUser)`

SetReferredByUser sets ReferredByUser field to given value.

### HasReferredByUser

`func (o *ApplicationUser) HasReferredByUser() bool`

HasReferredByUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



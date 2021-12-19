# DataConsent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **NullableString** |  | [optional] 
**RequestedByOrgId** | Pointer to **string** |  | [optional] 
**TransactionId** | Pointer to **NullableString** |  | [optional] 
**StartDateTime** | Pointer to **NullableTime** |  | [optional] 
**ExpiryDateTime** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**PurposeCode** | Pointer to **NullableString** |  | [optional] 
**PurposeLink** | Pointer to **NullableString** |  | [optional] 
**Location** | Pointer to **NullableString** |  | [optional] 
**AgreementId** | Pointer to **string** |  | [optional] 
**DataLifeUnit** | Pointer to [**DataLifeUnit**](DataLifeUnit.md) |  | [optional] 
**DataLifeValue** | Pointer to **int32** |  | [optional] 
**DataFetchFrequencyUnit** | Pointer to [**DataFetchFrequencyUnit**](DataFetchFrequencyUnit.md) |  | [optional] 
**DataFetchFrequencyUnitValue** | Pointer to **int32** |  | [optional] 
**DataFetchType** | Pointer to [**DataFetchType**](DataFetchType.md) |  | [optional] 
**Status** | Pointer to [**DataConsentStatus**](DataConsentStatus.md) |  | [optional] 
**CreatedAtUtc** | Pointer to **time.Time** |  | [optional] 
**ApprovedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**RejectedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**User** | Pointer to [**ApplicationUser**](ApplicationUser.md) |  | [optional] 
**Organization** | Pointer to [**Organization**](Organization.md) |  | [optional] 
**RequestedByOrg** | Pointer to [**Organization**](Organization.md) |  | [optional] 
**Agreement** | Pointer to [**DataProcessingAgreement**](DataProcessingAgreement.md) |  | [optional] 
**IdentityClaims** | Pointer to [**[]IdentityClaim**](IdentityClaim.md) |  | [optional] 
**Identifiers** | Pointer to [**[]DataConsentIdentifier**](DataConsentIdentifier.md) |  | [optional] 
**RequestedFinancialAccounts** | Pointer to [**[]DataConsentRequestedFinancialAccount**](DataConsentRequestedFinancialAccount.md) |  | [optional] 
**RequestedDocuments** | Pointer to [**[]DataConsentRequestedDocument**](DataConsentRequestedDocument.md) |  | [optional] 

## Methods

### NewDataConsent

`func NewDataConsent() *DataConsent`

NewDataConsent instantiates a new DataConsent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataConsentWithDefaults

`func NewDataConsentWithDefaults() *DataConsent`

NewDataConsentWithDefaults instantiates a new DataConsent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataConsent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataConsent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataConsent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataConsent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *DataConsent) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DataConsent) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DataConsent) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DataConsent) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *DataConsent) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *DataConsent) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetOrganizationId

`func (o *DataConsent) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *DataConsent) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *DataConsent) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *DataConsent) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *DataConsent) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *DataConsent) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetRequestedByOrgId

`func (o *DataConsent) GetRequestedByOrgId() string`

GetRequestedByOrgId returns the RequestedByOrgId field if non-nil, zero value otherwise.

### GetRequestedByOrgIdOk

`func (o *DataConsent) GetRequestedByOrgIdOk() (*string, bool)`

GetRequestedByOrgIdOk returns a tuple with the RequestedByOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedByOrgId

`func (o *DataConsent) SetRequestedByOrgId(v string)`

SetRequestedByOrgId sets RequestedByOrgId field to given value.

### HasRequestedByOrgId

`func (o *DataConsent) HasRequestedByOrgId() bool`

HasRequestedByOrgId returns a boolean if a field has been set.

### GetTransactionId

`func (o *DataConsent) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *DataConsent) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *DataConsent) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *DataConsent) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### SetTransactionIdNil

`func (o *DataConsent) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *DataConsent) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
### GetStartDateTime

`func (o *DataConsent) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *DataConsent) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *DataConsent) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *DataConsent) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### SetStartDateTimeNil

`func (o *DataConsent) SetStartDateTimeNil(b bool)`

 SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil

### UnsetStartDateTime
`func (o *DataConsent) UnsetStartDateTime()`

UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil
### GetExpiryDateTime

`func (o *DataConsent) GetExpiryDateTime() time.Time`

GetExpiryDateTime returns the ExpiryDateTime field if non-nil, zero value otherwise.

### GetExpiryDateTimeOk

`func (o *DataConsent) GetExpiryDateTimeOk() (*time.Time, bool)`

GetExpiryDateTimeOk returns a tuple with the ExpiryDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDateTime

`func (o *DataConsent) SetExpiryDateTime(v time.Time)`

SetExpiryDateTime sets ExpiryDateTime field to given value.

### HasExpiryDateTime

`func (o *DataConsent) HasExpiryDateTime() bool`

HasExpiryDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *DataConsent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DataConsent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DataConsent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DataConsent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DataConsent) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DataConsent) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPurposeCode

`func (o *DataConsent) GetPurposeCode() string`

GetPurposeCode returns the PurposeCode field if non-nil, zero value otherwise.

### GetPurposeCodeOk

`func (o *DataConsent) GetPurposeCodeOk() (*string, bool)`

GetPurposeCodeOk returns a tuple with the PurposeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurposeCode

`func (o *DataConsent) SetPurposeCode(v string)`

SetPurposeCode sets PurposeCode field to given value.

### HasPurposeCode

`func (o *DataConsent) HasPurposeCode() bool`

HasPurposeCode returns a boolean if a field has been set.

### SetPurposeCodeNil

`func (o *DataConsent) SetPurposeCodeNil(b bool)`

 SetPurposeCodeNil sets the value for PurposeCode to be an explicit nil

### UnsetPurposeCode
`func (o *DataConsent) UnsetPurposeCode()`

UnsetPurposeCode ensures that no value is present for PurposeCode, not even an explicit nil
### GetPurposeLink

`func (o *DataConsent) GetPurposeLink() string`

GetPurposeLink returns the PurposeLink field if non-nil, zero value otherwise.

### GetPurposeLinkOk

`func (o *DataConsent) GetPurposeLinkOk() (*string, bool)`

GetPurposeLinkOk returns a tuple with the PurposeLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurposeLink

`func (o *DataConsent) SetPurposeLink(v string)`

SetPurposeLink sets PurposeLink field to given value.

### HasPurposeLink

`func (o *DataConsent) HasPurposeLink() bool`

HasPurposeLink returns a boolean if a field has been set.

### SetPurposeLinkNil

`func (o *DataConsent) SetPurposeLinkNil(b bool)`

 SetPurposeLinkNil sets the value for PurposeLink to be an explicit nil

### UnsetPurposeLink
`func (o *DataConsent) UnsetPurposeLink()`

UnsetPurposeLink ensures that no value is present for PurposeLink, not even an explicit nil
### GetLocation

`func (o *DataConsent) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DataConsent) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DataConsent) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DataConsent) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *DataConsent) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *DataConsent) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetAgreementId

`func (o *DataConsent) GetAgreementId() string`

GetAgreementId returns the AgreementId field if non-nil, zero value otherwise.

### GetAgreementIdOk

`func (o *DataConsent) GetAgreementIdOk() (*string, bool)`

GetAgreementIdOk returns a tuple with the AgreementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementId

`func (o *DataConsent) SetAgreementId(v string)`

SetAgreementId sets AgreementId field to given value.

### HasAgreementId

`func (o *DataConsent) HasAgreementId() bool`

HasAgreementId returns a boolean if a field has been set.

### GetDataLifeUnit

`func (o *DataConsent) GetDataLifeUnit() DataLifeUnit`

GetDataLifeUnit returns the DataLifeUnit field if non-nil, zero value otherwise.

### GetDataLifeUnitOk

`func (o *DataConsent) GetDataLifeUnitOk() (*DataLifeUnit, bool)`

GetDataLifeUnitOk returns a tuple with the DataLifeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLifeUnit

`func (o *DataConsent) SetDataLifeUnit(v DataLifeUnit)`

SetDataLifeUnit sets DataLifeUnit field to given value.

### HasDataLifeUnit

`func (o *DataConsent) HasDataLifeUnit() bool`

HasDataLifeUnit returns a boolean if a field has been set.

### GetDataLifeValue

`func (o *DataConsent) GetDataLifeValue() int32`

GetDataLifeValue returns the DataLifeValue field if non-nil, zero value otherwise.

### GetDataLifeValueOk

`func (o *DataConsent) GetDataLifeValueOk() (*int32, bool)`

GetDataLifeValueOk returns a tuple with the DataLifeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLifeValue

`func (o *DataConsent) SetDataLifeValue(v int32)`

SetDataLifeValue sets DataLifeValue field to given value.

### HasDataLifeValue

`func (o *DataConsent) HasDataLifeValue() bool`

HasDataLifeValue returns a boolean if a field has been set.

### GetDataFetchFrequencyUnit

`func (o *DataConsent) GetDataFetchFrequencyUnit() DataFetchFrequencyUnit`

GetDataFetchFrequencyUnit returns the DataFetchFrequencyUnit field if non-nil, zero value otherwise.

### GetDataFetchFrequencyUnitOk

`func (o *DataConsent) GetDataFetchFrequencyUnitOk() (*DataFetchFrequencyUnit, bool)`

GetDataFetchFrequencyUnitOk returns a tuple with the DataFetchFrequencyUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFetchFrequencyUnit

`func (o *DataConsent) SetDataFetchFrequencyUnit(v DataFetchFrequencyUnit)`

SetDataFetchFrequencyUnit sets DataFetchFrequencyUnit field to given value.

### HasDataFetchFrequencyUnit

`func (o *DataConsent) HasDataFetchFrequencyUnit() bool`

HasDataFetchFrequencyUnit returns a boolean if a field has been set.

### GetDataFetchFrequencyUnitValue

`func (o *DataConsent) GetDataFetchFrequencyUnitValue() int32`

GetDataFetchFrequencyUnitValue returns the DataFetchFrequencyUnitValue field if non-nil, zero value otherwise.

### GetDataFetchFrequencyUnitValueOk

`func (o *DataConsent) GetDataFetchFrequencyUnitValueOk() (*int32, bool)`

GetDataFetchFrequencyUnitValueOk returns a tuple with the DataFetchFrequencyUnitValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFetchFrequencyUnitValue

`func (o *DataConsent) SetDataFetchFrequencyUnitValue(v int32)`

SetDataFetchFrequencyUnitValue sets DataFetchFrequencyUnitValue field to given value.

### HasDataFetchFrequencyUnitValue

`func (o *DataConsent) HasDataFetchFrequencyUnitValue() bool`

HasDataFetchFrequencyUnitValue returns a boolean if a field has been set.

### GetDataFetchType

`func (o *DataConsent) GetDataFetchType() DataFetchType`

GetDataFetchType returns the DataFetchType field if non-nil, zero value otherwise.

### GetDataFetchTypeOk

`func (o *DataConsent) GetDataFetchTypeOk() (*DataFetchType, bool)`

GetDataFetchTypeOk returns a tuple with the DataFetchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFetchType

`func (o *DataConsent) SetDataFetchType(v DataFetchType)`

SetDataFetchType sets DataFetchType field to given value.

### HasDataFetchType

`func (o *DataConsent) HasDataFetchType() bool`

HasDataFetchType returns a boolean if a field has been set.

### GetStatus

`func (o *DataConsent) GetStatus() DataConsentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DataConsent) GetStatusOk() (*DataConsentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DataConsent) SetStatus(v DataConsentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DataConsent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAtUtc

`func (o *DataConsent) GetCreatedAtUtc() time.Time`

GetCreatedAtUtc returns the CreatedAtUtc field if non-nil, zero value otherwise.

### GetCreatedAtUtcOk

`func (o *DataConsent) GetCreatedAtUtcOk() (*time.Time, bool)`

GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtUtc

`func (o *DataConsent) SetCreatedAtUtc(v time.Time)`

SetCreatedAtUtc sets CreatedAtUtc field to given value.

### HasCreatedAtUtc

`func (o *DataConsent) HasCreatedAtUtc() bool`

HasCreatedAtUtc returns a boolean if a field has been set.

### GetApprovedAtUtc

`func (o *DataConsent) GetApprovedAtUtc() time.Time`

GetApprovedAtUtc returns the ApprovedAtUtc field if non-nil, zero value otherwise.

### GetApprovedAtUtcOk

`func (o *DataConsent) GetApprovedAtUtcOk() (*time.Time, bool)`

GetApprovedAtUtcOk returns a tuple with the ApprovedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAtUtc

`func (o *DataConsent) SetApprovedAtUtc(v time.Time)`

SetApprovedAtUtc sets ApprovedAtUtc field to given value.

### HasApprovedAtUtc

`func (o *DataConsent) HasApprovedAtUtc() bool`

HasApprovedAtUtc returns a boolean if a field has been set.

### SetApprovedAtUtcNil

`func (o *DataConsent) SetApprovedAtUtcNil(b bool)`

 SetApprovedAtUtcNil sets the value for ApprovedAtUtc to be an explicit nil

### UnsetApprovedAtUtc
`func (o *DataConsent) UnsetApprovedAtUtc()`

UnsetApprovedAtUtc ensures that no value is present for ApprovedAtUtc, not even an explicit nil
### GetRejectedAtUtc

`func (o *DataConsent) GetRejectedAtUtc() time.Time`

GetRejectedAtUtc returns the RejectedAtUtc field if non-nil, zero value otherwise.

### GetRejectedAtUtcOk

`func (o *DataConsent) GetRejectedAtUtcOk() (*time.Time, bool)`

GetRejectedAtUtcOk returns a tuple with the RejectedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedAtUtc

`func (o *DataConsent) SetRejectedAtUtc(v time.Time)`

SetRejectedAtUtc sets RejectedAtUtc field to given value.

### HasRejectedAtUtc

`func (o *DataConsent) HasRejectedAtUtc() bool`

HasRejectedAtUtc returns a boolean if a field has been set.

### SetRejectedAtUtcNil

`func (o *DataConsent) SetRejectedAtUtcNil(b bool)`

 SetRejectedAtUtcNil sets the value for RejectedAtUtc to be an explicit nil

### UnsetRejectedAtUtc
`func (o *DataConsent) UnsetRejectedAtUtc()`

UnsetRejectedAtUtc ensures that no value is present for RejectedAtUtc, not even an explicit nil
### GetUser

`func (o *DataConsent) GetUser() ApplicationUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DataConsent) GetUserOk() (*ApplicationUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DataConsent) SetUser(v ApplicationUser)`

SetUser sets User field to given value.

### HasUser

`func (o *DataConsent) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetOrganization

`func (o *DataConsent) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *DataConsent) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *DataConsent) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *DataConsent) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetRequestedByOrg

`func (o *DataConsent) GetRequestedByOrg() Organization`

GetRequestedByOrg returns the RequestedByOrg field if non-nil, zero value otherwise.

### GetRequestedByOrgOk

`func (o *DataConsent) GetRequestedByOrgOk() (*Organization, bool)`

GetRequestedByOrgOk returns a tuple with the RequestedByOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedByOrg

`func (o *DataConsent) SetRequestedByOrg(v Organization)`

SetRequestedByOrg sets RequestedByOrg field to given value.

### HasRequestedByOrg

`func (o *DataConsent) HasRequestedByOrg() bool`

HasRequestedByOrg returns a boolean if a field has been set.

### GetAgreement

`func (o *DataConsent) GetAgreement() DataProcessingAgreement`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *DataConsent) GetAgreementOk() (*DataProcessingAgreement, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *DataConsent) SetAgreement(v DataProcessingAgreement)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *DataConsent) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetIdentityClaims

`func (o *DataConsent) GetIdentityClaims() []IdentityClaim`

GetIdentityClaims returns the IdentityClaims field if non-nil, zero value otherwise.

### GetIdentityClaimsOk

`func (o *DataConsent) GetIdentityClaimsOk() (*[]IdentityClaim, bool)`

GetIdentityClaimsOk returns a tuple with the IdentityClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityClaims

`func (o *DataConsent) SetIdentityClaims(v []IdentityClaim)`

SetIdentityClaims sets IdentityClaims field to given value.

### HasIdentityClaims

`func (o *DataConsent) HasIdentityClaims() bool`

HasIdentityClaims returns a boolean if a field has been set.

### SetIdentityClaimsNil

`func (o *DataConsent) SetIdentityClaimsNil(b bool)`

 SetIdentityClaimsNil sets the value for IdentityClaims to be an explicit nil

### UnsetIdentityClaims
`func (o *DataConsent) UnsetIdentityClaims()`

UnsetIdentityClaims ensures that no value is present for IdentityClaims, not even an explicit nil
### GetIdentifiers

`func (o *DataConsent) GetIdentifiers() []DataConsentIdentifier`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *DataConsent) GetIdentifiersOk() (*[]DataConsentIdentifier, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *DataConsent) SetIdentifiers(v []DataConsentIdentifier)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *DataConsent) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### SetIdentifiersNil

`func (o *DataConsent) SetIdentifiersNil(b bool)`

 SetIdentifiersNil sets the value for Identifiers to be an explicit nil

### UnsetIdentifiers
`func (o *DataConsent) UnsetIdentifiers()`

UnsetIdentifiers ensures that no value is present for Identifiers, not even an explicit nil
### GetRequestedFinancialAccounts

`func (o *DataConsent) GetRequestedFinancialAccounts() []DataConsentRequestedFinancialAccount`

GetRequestedFinancialAccounts returns the RequestedFinancialAccounts field if non-nil, zero value otherwise.

### GetRequestedFinancialAccountsOk

`func (o *DataConsent) GetRequestedFinancialAccountsOk() (*[]DataConsentRequestedFinancialAccount, bool)`

GetRequestedFinancialAccountsOk returns a tuple with the RequestedFinancialAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedFinancialAccounts

`func (o *DataConsent) SetRequestedFinancialAccounts(v []DataConsentRequestedFinancialAccount)`

SetRequestedFinancialAccounts sets RequestedFinancialAccounts field to given value.

### HasRequestedFinancialAccounts

`func (o *DataConsent) HasRequestedFinancialAccounts() bool`

HasRequestedFinancialAccounts returns a boolean if a field has been set.

### SetRequestedFinancialAccountsNil

`func (o *DataConsent) SetRequestedFinancialAccountsNil(b bool)`

 SetRequestedFinancialAccountsNil sets the value for RequestedFinancialAccounts to be an explicit nil

### UnsetRequestedFinancialAccounts
`func (o *DataConsent) UnsetRequestedFinancialAccounts()`

UnsetRequestedFinancialAccounts ensures that no value is present for RequestedFinancialAccounts, not even an explicit nil
### GetRequestedDocuments

`func (o *DataConsent) GetRequestedDocuments() []DataConsentRequestedDocument`

GetRequestedDocuments returns the RequestedDocuments field if non-nil, zero value otherwise.

### GetRequestedDocumentsOk

`func (o *DataConsent) GetRequestedDocumentsOk() (*[]DataConsentRequestedDocument, bool)`

GetRequestedDocumentsOk returns a tuple with the RequestedDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedDocuments

`func (o *DataConsent) SetRequestedDocuments(v []DataConsentRequestedDocument)`

SetRequestedDocuments sets RequestedDocuments field to given value.

### HasRequestedDocuments

`func (o *DataConsent) HasRequestedDocuments() bool`

HasRequestedDocuments returns a boolean if a field has been set.

### SetRequestedDocumentsNil

`func (o *DataConsent) SetRequestedDocumentsNil(b bool)`

 SetRequestedDocumentsNil sets the value for RequestedDocuments to be an explicit nil

### UnsetRequestedDocuments
`func (o *DataConsent) UnsetRequestedDocuments()`

UnsetRequestedDocuments ensures that no value is present for RequestedDocuments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



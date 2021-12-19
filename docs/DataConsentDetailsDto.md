# DataConsentDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Requester** | Pointer to [**DataConsentRequesterDto**](DataConsentRequesterDto.md) |  | [optional] 
**Location** | Pointer to **NullableString** |  | [optional] 
**PersonalInfoRequested** | Pointer to **bool** |  | [optional] 
**Documents** | Pointer to **int32** |  | [optional] 
**FinancialAccounts** | Pointer to **int32** |  | [optional] 
**TransactionId** | Pointer to **NullableString** |  | [optional] 
**IpAddress** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**PurposeCode** | Pointer to **NullableString** |  | [optional] 
**PurposeLink** | Pointer to **NullableString** |  | [optional] 
**AgreementId** | Pointer to **NullableString** |  | [optional] 
**DataLifeUnit** | Pointer to [**DataLifeUnit**](DataLifeUnit.md) |  | [optional] 
**DataLifeValue** | Pointer to **int32** |  | [optional] 
**DataFetchFrequencyUnit** | Pointer to [**DataFetchFrequencyUnit**](DataFetchFrequencyUnit.md) |  | [optional] 
**DataFetchFrequencyUnitValue** | Pointer to **int32** |  | [optional] 
**DataFetchType** | Pointer to [**DataFetchType**](DataFetchType.md) |  | [optional] 
**Status** | Pointer to [**DataConsentStatus**](DataConsentStatus.md) |  | [optional] 
**ApprovedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**RejectedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**ExpiresAtUtc** | Pointer to **time.Time** |  | [optional] 
**RequestedAtUtc** | Pointer to **time.Time** |  | [optional] 
**RequestedFinancialAccounts** | Pointer to [**[]DataConsentRequestedAccountDto**](DataConsentRequestedAccountDto.md) |  | [optional] 
**RequestedDocuments** | Pointer to [**[]DataConsentRequestedDocumentDto**](DataConsentRequestedDocumentDto.md) |  | [optional] 
**RequestedHealthData** | Pointer to [**[]DataConsentRequestedDocument**](DataConsentRequestedDocument.md) |  | [optional] 
**RequestedIdentityDetails** | Pointer to [**JsonSchema**](JsonSchema.md) |  | [optional] 

## Methods

### NewDataConsentDetailsDto

`func NewDataConsentDetailsDto() *DataConsentDetailsDto`

NewDataConsentDetailsDto instantiates a new DataConsentDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataConsentDetailsDtoWithDefaults

`func NewDataConsentDetailsDtoWithDefaults() *DataConsentDetailsDto`

NewDataConsentDetailsDtoWithDefaults instantiates a new DataConsentDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataConsentDetailsDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataConsentDetailsDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataConsentDetailsDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataConsentDetailsDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRequester

`func (o *DataConsentDetailsDto) GetRequester() DataConsentRequesterDto`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *DataConsentDetailsDto) GetRequesterOk() (*DataConsentRequesterDto, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *DataConsentDetailsDto) SetRequester(v DataConsentRequesterDto)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *DataConsentDetailsDto) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### GetLocation

`func (o *DataConsentDetailsDto) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DataConsentDetailsDto) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DataConsentDetailsDto) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DataConsentDetailsDto) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *DataConsentDetailsDto) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *DataConsentDetailsDto) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetPersonalInfoRequested

`func (o *DataConsentDetailsDto) GetPersonalInfoRequested() bool`

GetPersonalInfoRequested returns the PersonalInfoRequested field if non-nil, zero value otherwise.

### GetPersonalInfoRequestedOk

`func (o *DataConsentDetailsDto) GetPersonalInfoRequestedOk() (*bool, bool)`

GetPersonalInfoRequestedOk returns a tuple with the PersonalInfoRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalInfoRequested

`func (o *DataConsentDetailsDto) SetPersonalInfoRequested(v bool)`

SetPersonalInfoRequested sets PersonalInfoRequested field to given value.

### HasPersonalInfoRequested

`func (o *DataConsentDetailsDto) HasPersonalInfoRequested() bool`

HasPersonalInfoRequested returns a boolean if a field has been set.

### GetDocuments

`func (o *DataConsentDetailsDto) GetDocuments() int32`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *DataConsentDetailsDto) GetDocumentsOk() (*int32, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *DataConsentDetailsDto) SetDocuments(v int32)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *DataConsentDetailsDto) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetFinancialAccounts

`func (o *DataConsentDetailsDto) GetFinancialAccounts() int32`

GetFinancialAccounts returns the FinancialAccounts field if non-nil, zero value otherwise.

### GetFinancialAccountsOk

`func (o *DataConsentDetailsDto) GetFinancialAccountsOk() (*int32, bool)`

GetFinancialAccountsOk returns a tuple with the FinancialAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialAccounts

`func (o *DataConsentDetailsDto) SetFinancialAccounts(v int32)`

SetFinancialAccounts sets FinancialAccounts field to given value.

### HasFinancialAccounts

`func (o *DataConsentDetailsDto) HasFinancialAccounts() bool`

HasFinancialAccounts returns a boolean if a field has been set.

### GetTransactionId

`func (o *DataConsentDetailsDto) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *DataConsentDetailsDto) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *DataConsentDetailsDto) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *DataConsentDetailsDto) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### SetTransactionIdNil

`func (o *DataConsentDetailsDto) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *DataConsentDetailsDto) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
### GetIpAddress

`func (o *DataConsentDetailsDto) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *DataConsentDetailsDto) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *DataConsentDetailsDto) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *DataConsentDetailsDto) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *DataConsentDetailsDto) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *DataConsentDetailsDto) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetDescription

`func (o *DataConsentDetailsDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DataConsentDetailsDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DataConsentDetailsDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DataConsentDetailsDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DataConsentDetailsDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DataConsentDetailsDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPurposeCode

`func (o *DataConsentDetailsDto) GetPurposeCode() string`

GetPurposeCode returns the PurposeCode field if non-nil, zero value otherwise.

### GetPurposeCodeOk

`func (o *DataConsentDetailsDto) GetPurposeCodeOk() (*string, bool)`

GetPurposeCodeOk returns a tuple with the PurposeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurposeCode

`func (o *DataConsentDetailsDto) SetPurposeCode(v string)`

SetPurposeCode sets PurposeCode field to given value.

### HasPurposeCode

`func (o *DataConsentDetailsDto) HasPurposeCode() bool`

HasPurposeCode returns a boolean if a field has been set.

### SetPurposeCodeNil

`func (o *DataConsentDetailsDto) SetPurposeCodeNil(b bool)`

 SetPurposeCodeNil sets the value for PurposeCode to be an explicit nil

### UnsetPurposeCode
`func (o *DataConsentDetailsDto) UnsetPurposeCode()`

UnsetPurposeCode ensures that no value is present for PurposeCode, not even an explicit nil
### GetPurposeLink

`func (o *DataConsentDetailsDto) GetPurposeLink() string`

GetPurposeLink returns the PurposeLink field if non-nil, zero value otherwise.

### GetPurposeLinkOk

`func (o *DataConsentDetailsDto) GetPurposeLinkOk() (*string, bool)`

GetPurposeLinkOk returns a tuple with the PurposeLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurposeLink

`func (o *DataConsentDetailsDto) SetPurposeLink(v string)`

SetPurposeLink sets PurposeLink field to given value.

### HasPurposeLink

`func (o *DataConsentDetailsDto) HasPurposeLink() bool`

HasPurposeLink returns a boolean if a field has been set.

### SetPurposeLinkNil

`func (o *DataConsentDetailsDto) SetPurposeLinkNil(b bool)`

 SetPurposeLinkNil sets the value for PurposeLink to be an explicit nil

### UnsetPurposeLink
`func (o *DataConsentDetailsDto) UnsetPurposeLink()`

UnsetPurposeLink ensures that no value is present for PurposeLink, not even an explicit nil
### GetAgreementId

`func (o *DataConsentDetailsDto) GetAgreementId() string`

GetAgreementId returns the AgreementId field if non-nil, zero value otherwise.

### GetAgreementIdOk

`func (o *DataConsentDetailsDto) GetAgreementIdOk() (*string, bool)`

GetAgreementIdOk returns a tuple with the AgreementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementId

`func (o *DataConsentDetailsDto) SetAgreementId(v string)`

SetAgreementId sets AgreementId field to given value.

### HasAgreementId

`func (o *DataConsentDetailsDto) HasAgreementId() bool`

HasAgreementId returns a boolean if a field has been set.

### SetAgreementIdNil

`func (o *DataConsentDetailsDto) SetAgreementIdNil(b bool)`

 SetAgreementIdNil sets the value for AgreementId to be an explicit nil

### UnsetAgreementId
`func (o *DataConsentDetailsDto) UnsetAgreementId()`

UnsetAgreementId ensures that no value is present for AgreementId, not even an explicit nil
### GetDataLifeUnit

`func (o *DataConsentDetailsDto) GetDataLifeUnit() DataLifeUnit`

GetDataLifeUnit returns the DataLifeUnit field if non-nil, zero value otherwise.

### GetDataLifeUnitOk

`func (o *DataConsentDetailsDto) GetDataLifeUnitOk() (*DataLifeUnit, bool)`

GetDataLifeUnitOk returns a tuple with the DataLifeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLifeUnit

`func (o *DataConsentDetailsDto) SetDataLifeUnit(v DataLifeUnit)`

SetDataLifeUnit sets DataLifeUnit field to given value.

### HasDataLifeUnit

`func (o *DataConsentDetailsDto) HasDataLifeUnit() bool`

HasDataLifeUnit returns a boolean if a field has been set.

### GetDataLifeValue

`func (o *DataConsentDetailsDto) GetDataLifeValue() int32`

GetDataLifeValue returns the DataLifeValue field if non-nil, zero value otherwise.

### GetDataLifeValueOk

`func (o *DataConsentDetailsDto) GetDataLifeValueOk() (*int32, bool)`

GetDataLifeValueOk returns a tuple with the DataLifeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLifeValue

`func (o *DataConsentDetailsDto) SetDataLifeValue(v int32)`

SetDataLifeValue sets DataLifeValue field to given value.

### HasDataLifeValue

`func (o *DataConsentDetailsDto) HasDataLifeValue() bool`

HasDataLifeValue returns a boolean if a field has been set.

### GetDataFetchFrequencyUnit

`func (o *DataConsentDetailsDto) GetDataFetchFrequencyUnit() DataFetchFrequencyUnit`

GetDataFetchFrequencyUnit returns the DataFetchFrequencyUnit field if non-nil, zero value otherwise.

### GetDataFetchFrequencyUnitOk

`func (o *DataConsentDetailsDto) GetDataFetchFrequencyUnitOk() (*DataFetchFrequencyUnit, bool)`

GetDataFetchFrequencyUnitOk returns a tuple with the DataFetchFrequencyUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFetchFrequencyUnit

`func (o *DataConsentDetailsDto) SetDataFetchFrequencyUnit(v DataFetchFrequencyUnit)`

SetDataFetchFrequencyUnit sets DataFetchFrequencyUnit field to given value.

### HasDataFetchFrequencyUnit

`func (o *DataConsentDetailsDto) HasDataFetchFrequencyUnit() bool`

HasDataFetchFrequencyUnit returns a boolean if a field has been set.

### GetDataFetchFrequencyUnitValue

`func (o *DataConsentDetailsDto) GetDataFetchFrequencyUnitValue() int32`

GetDataFetchFrequencyUnitValue returns the DataFetchFrequencyUnitValue field if non-nil, zero value otherwise.

### GetDataFetchFrequencyUnitValueOk

`func (o *DataConsentDetailsDto) GetDataFetchFrequencyUnitValueOk() (*int32, bool)`

GetDataFetchFrequencyUnitValueOk returns a tuple with the DataFetchFrequencyUnitValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFetchFrequencyUnitValue

`func (o *DataConsentDetailsDto) SetDataFetchFrequencyUnitValue(v int32)`

SetDataFetchFrequencyUnitValue sets DataFetchFrequencyUnitValue field to given value.

### HasDataFetchFrequencyUnitValue

`func (o *DataConsentDetailsDto) HasDataFetchFrequencyUnitValue() bool`

HasDataFetchFrequencyUnitValue returns a boolean if a field has been set.

### GetDataFetchType

`func (o *DataConsentDetailsDto) GetDataFetchType() DataFetchType`

GetDataFetchType returns the DataFetchType field if non-nil, zero value otherwise.

### GetDataFetchTypeOk

`func (o *DataConsentDetailsDto) GetDataFetchTypeOk() (*DataFetchType, bool)`

GetDataFetchTypeOk returns a tuple with the DataFetchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFetchType

`func (o *DataConsentDetailsDto) SetDataFetchType(v DataFetchType)`

SetDataFetchType sets DataFetchType field to given value.

### HasDataFetchType

`func (o *DataConsentDetailsDto) HasDataFetchType() bool`

HasDataFetchType returns a boolean if a field has been set.

### GetStatus

`func (o *DataConsentDetailsDto) GetStatus() DataConsentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DataConsentDetailsDto) GetStatusOk() (*DataConsentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DataConsentDetailsDto) SetStatus(v DataConsentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DataConsentDetailsDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApprovedAtUtc

`func (o *DataConsentDetailsDto) GetApprovedAtUtc() time.Time`

GetApprovedAtUtc returns the ApprovedAtUtc field if non-nil, zero value otherwise.

### GetApprovedAtUtcOk

`func (o *DataConsentDetailsDto) GetApprovedAtUtcOk() (*time.Time, bool)`

GetApprovedAtUtcOk returns a tuple with the ApprovedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAtUtc

`func (o *DataConsentDetailsDto) SetApprovedAtUtc(v time.Time)`

SetApprovedAtUtc sets ApprovedAtUtc field to given value.

### HasApprovedAtUtc

`func (o *DataConsentDetailsDto) HasApprovedAtUtc() bool`

HasApprovedAtUtc returns a boolean if a field has been set.

### SetApprovedAtUtcNil

`func (o *DataConsentDetailsDto) SetApprovedAtUtcNil(b bool)`

 SetApprovedAtUtcNil sets the value for ApprovedAtUtc to be an explicit nil

### UnsetApprovedAtUtc
`func (o *DataConsentDetailsDto) UnsetApprovedAtUtc()`

UnsetApprovedAtUtc ensures that no value is present for ApprovedAtUtc, not even an explicit nil
### GetRejectedAtUtc

`func (o *DataConsentDetailsDto) GetRejectedAtUtc() time.Time`

GetRejectedAtUtc returns the RejectedAtUtc field if non-nil, zero value otherwise.

### GetRejectedAtUtcOk

`func (o *DataConsentDetailsDto) GetRejectedAtUtcOk() (*time.Time, bool)`

GetRejectedAtUtcOk returns a tuple with the RejectedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedAtUtc

`func (o *DataConsentDetailsDto) SetRejectedAtUtc(v time.Time)`

SetRejectedAtUtc sets RejectedAtUtc field to given value.

### HasRejectedAtUtc

`func (o *DataConsentDetailsDto) HasRejectedAtUtc() bool`

HasRejectedAtUtc returns a boolean if a field has been set.

### SetRejectedAtUtcNil

`func (o *DataConsentDetailsDto) SetRejectedAtUtcNil(b bool)`

 SetRejectedAtUtcNil sets the value for RejectedAtUtc to be an explicit nil

### UnsetRejectedAtUtc
`func (o *DataConsentDetailsDto) UnsetRejectedAtUtc()`

UnsetRejectedAtUtc ensures that no value is present for RejectedAtUtc, not even an explicit nil
### GetExpiresAtUtc

`func (o *DataConsentDetailsDto) GetExpiresAtUtc() time.Time`

GetExpiresAtUtc returns the ExpiresAtUtc field if non-nil, zero value otherwise.

### GetExpiresAtUtcOk

`func (o *DataConsentDetailsDto) GetExpiresAtUtcOk() (*time.Time, bool)`

GetExpiresAtUtcOk returns a tuple with the ExpiresAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAtUtc

`func (o *DataConsentDetailsDto) SetExpiresAtUtc(v time.Time)`

SetExpiresAtUtc sets ExpiresAtUtc field to given value.

### HasExpiresAtUtc

`func (o *DataConsentDetailsDto) HasExpiresAtUtc() bool`

HasExpiresAtUtc returns a boolean if a field has been set.

### GetRequestedAtUtc

`func (o *DataConsentDetailsDto) GetRequestedAtUtc() time.Time`

GetRequestedAtUtc returns the RequestedAtUtc field if non-nil, zero value otherwise.

### GetRequestedAtUtcOk

`func (o *DataConsentDetailsDto) GetRequestedAtUtcOk() (*time.Time, bool)`

GetRequestedAtUtcOk returns a tuple with the RequestedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAtUtc

`func (o *DataConsentDetailsDto) SetRequestedAtUtc(v time.Time)`

SetRequestedAtUtc sets RequestedAtUtc field to given value.

### HasRequestedAtUtc

`func (o *DataConsentDetailsDto) HasRequestedAtUtc() bool`

HasRequestedAtUtc returns a boolean if a field has been set.

### GetRequestedFinancialAccounts

`func (o *DataConsentDetailsDto) GetRequestedFinancialAccounts() []DataConsentRequestedAccountDto`

GetRequestedFinancialAccounts returns the RequestedFinancialAccounts field if non-nil, zero value otherwise.

### GetRequestedFinancialAccountsOk

`func (o *DataConsentDetailsDto) GetRequestedFinancialAccountsOk() (*[]DataConsentRequestedAccountDto, bool)`

GetRequestedFinancialAccountsOk returns a tuple with the RequestedFinancialAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedFinancialAccounts

`func (o *DataConsentDetailsDto) SetRequestedFinancialAccounts(v []DataConsentRequestedAccountDto)`

SetRequestedFinancialAccounts sets RequestedFinancialAccounts field to given value.

### HasRequestedFinancialAccounts

`func (o *DataConsentDetailsDto) HasRequestedFinancialAccounts() bool`

HasRequestedFinancialAccounts returns a boolean if a field has been set.

### SetRequestedFinancialAccountsNil

`func (o *DataConsentDetailsDto) SetRequestedFinancialAccountsNil(b bool)`

 SetRequestedFinancialAccountsNil sets the value for RequestedFinancialAccounts to be an explicit nil

### UnsetRequestedFinancialAccounts
`func (o *DataConsentDetailsDto) UnsetRequestedFinancialAccounts()`

UnsetRequestedFinancialAccounts ensures that no value is present for RequestedFinancialAccounts, not even an explicit nil
### GetRequestedDocuments

`func (o *DataConsentDetailsDto) GetRequestedDocuments() []DataConsentRequestedDocumentDto`

GetRequestedDocuments returns the RequestedDocuments field if non-nil, zero value otherwise.

### GetRequestedDocumentsOk

`func (o *DataConsentDetailsDto) GetRequestedDocumentsOk() (*[]DataConsentRequestedDocumentDto, bool)`

GetRequestedDocumentsOk returns a tuple with the RequestedDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedDocuments

`func (o *DataConsentDetailsDto) SetRequestedDocuments(v []DataConsentRequestedDocumentDto)`

SetRequestedDocuments sets RequestedDocuments field to given value.

### HasRequestedDocuments

`func (o *DataConsentDetailsDto) HasRequestedDocuments() bool`

HasRequestedDocuments returns a boolean if a field has been set.

### SetRequestedDocumentsNil

`func (o *DataConsentDetailsDto) SetRequestedDocumentsNil(b bool)`

 SetRequestedDocumentsNil sets the value for RequestedDocuments to be an explicit nil

### UnsetRequestedDocuments
`func (o *DataConsentDetailsDto) UnsetRequestedDocuments()`

UnsetRequestedDocuments ensures that no value is present for RequestedDocuments, not even an explicit nil
### GetRequestedHealthData

`func (o *DataConsentDetailsDto) GetRequestedHealthData() []DataConsentRequestedDocument`

GetRequestedHealthData returns the RequestedHealthData field if non-nil, zero value otherwise.

### GetRequestedHealthDataOk

`func (o *DataConsentDetailsDto) GetRequestedHealthDataOk() (*[]DataConsentRequestedDocument, bool)`

GetRequestedHealthDataOk returns a tuple with the RequestedHealthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedHealthData

`func (o *DataConsentDetailsDto) SetRequestedHealthData(v []DataConsentRequestedDocument)`

SetRequestedHealthData sets RequestedHealthData field to given value.

### HasRequestedHealthData

`func (o *DataConsentDetailsDto) HasRequestedHealthData() bool`

HasRequestedHealthData returns a boolean if a field has been set.

### SetRequestedHealthDataNil

`func (o *DataConsentDetailsDto) SetRequestedHealthDataNil(b bool)`

 SetRequestedHealthDataNil sets the value for RequestedHealthData to be an explicit nil

### UnsetRequestedHealthData
`func (o *DataConsentDetailsDto) UnsetRequestedHealthData()`

UnsetRequestedHealthData ensures that no value is present for RequestedHealthData, not even an explicit nil
### GetRequestedIdentityDetails

`func (o *DataConsentDetailsDto) GetRequestedIdentityDetails() JsonSchema`

GetRequestedIdentityDetails returns the RequestedIdentityDetails field if non-nil, zero value otherwise.

### GetRequestedIdentityDetailsOk

`func (o *DataConsentDetailsDto) GetRequestedIdentityDetailsOk() (*JsonSchema, bool)`

GetRequestedIdentityDetailsOk returns a tuple with the RequestedIdentityDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedIdentityDetails

`func (o *DataConsentDetailsDto) SetRequestedIdentityDetails(v JsonSchema)`

SetRequestedIdentityDetails sets RequestedIdentityDetails field to given value.

### HasRequestedIdentityDetails

`func (o *DataConsentDetailsDto) HasRequestedIdentityDetails() bool`

HasRequestedIdentityDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



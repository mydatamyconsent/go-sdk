# DataConsentRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | Pointer to **string** |  | [optional] 
**TransactionId** | Pointer to **NullableString** |  | [optional] 
**Identifiers** | Pointer to **map[string]string** |  | [optional] 
**StartDateTime** | Pointer to **NullableTime** |  | [optional] 
**ExpiryDateTime** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**PurposeCode** | Pointer to **NullableString** |  | [optional] 
**PurposeLink** | Pointer to **NullableString** |  | [optional] 
**DataLifeUnit** | Pointer to [**DataLifeUnit**](DataLifeUnit.md) |  | [optional] 
**DataLifeValue** | Pointer to **int32** |  | [optional] 
**DataFetchFrequencyUnit** | Pointer to [**DataFetchFrequencyUnit**](DataFetchFrequencyUnit.md) |  | [optional] 
**DataFetchFrequencyUnitValue** | Pointer to **int32** |  | [optional] 
**DataFetchType** | Pointer to [**DataFetchType**](DataFetchType.md) |  | [optional] 
**AgreementId** | Pointer to **string** |  | [optional] 
**IdentityClaims** | Pointer to [**[]IdentityClaim**](IdentityClaim.md) |  | [optional] 
**FinancialAccounts** | Pointer to [**[]DataConsentRequestedFaDto**](DataConsentRequestedFaDto.md) |  | [optional] 
**Documents** | Pointer to [**[]DataConsentRequestedDocumentDto**](DataConsentRequestedDocumentDto.md) |  | [optional] 

## Methods

### NewDataConsentRequestModel

`func NewDataConsentRequestModel() *DataConsentRequestModel`

NewDataConsentRequestModel instantiates a new DataConsentRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataConsentRequestModelWithDefaults

`func NewDataConsentRequestModelWithDefaults() *DataConsentRequestModel`

NewDataConsentRequestModelWithDefaults instantiates a new DataConsentRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *DataConsentRequestModel) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *DataConsentRequestModel) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *DataConsentRequestModel) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *DataConsentRequestModel) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetTransactionId

`func (o *DataConsentRequestModel) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *DataConsentRequestModel) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *DataConsentRequestModel) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *DataConsentRequestModel) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### SetTransactionIdNil

`func (o *DataConsentRequestModel) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *DataConsentRequestModel) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
### GetIdentifiers

`func (o *DataConsentRequestModel) GetIdentifiers() map[string]string`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *DataConsentRequestModel) GetIdentifiersOk() (*map[string]string, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *DataConsentRequestModel) SetIdentifiers(v map[string]string)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *DataConsentRequestModel) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### SetIdentifiersNil

`func (o *DataConsentRequestModel) SetIdentifiersNil(b bool)`

 SetIdentifiersNil sets the value for Identifiers to be an explicit nil

### UnsetIdentifiers
`func (o *DataConsentRequestModel) UnsetIdentifiers()`

UnsetIdentifiers ensures that no value is present for Identifiers, not even an explicit nil
### GetStartDateTime

`func (o *DataConsentRequestModel) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *DataConsentRequestModel) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *DataConsentRequestModel) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *DataConsentRequestModel) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### SetStartDateTimeNil

`func (o *DataConsentRequestModel) SetStartDateTimeNil(b bool)`

 SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil

### UnsetStartDateTime
`func (o *DataConsentRequestModel) UnsetStartDateTime()`

UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil
### GetExpiryDateTime

`func (o *DataConsentRequestModel) GetExpiryDateTime() time.Time`

GetExpiryDateTime returns the ExpiryDateTime field if non-nil, zero value otherwise.

### GetExpiryDateTimeOk

`func (o *DataConsentRequestModel) GetExpiryDateTimeOk() (*time.Time, bool)`

GetExpiryDateTimeOk returns a tuple with the ExpiryDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDateTime

`func (o *DataConsentRequestModel) SetExpiryDateTime(v time.Time)`

SetExpiryDateTime sets ExpiryDateTime field to given value.

### HasExpiryDateTime

`func (o *DataConsentRequestModel) HasExpiryDateTime() bool`

HasExpiryDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *DataConsentRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DataConsentRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DataConsentRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DataConsentRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DataConsentRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DataConsentRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPurposeCode

`func (o *DataConsentRequestModel) GetPurposeCode() string`

GetPurposeCode returns the PurposeCode field if non-nil, zero value otherwise.

### GetPurposeCodeOk

`func (o *DataConsentRequestModel) GetPurposeCodeOk() (*string, bool)`

GetPurposeCodeOk returns a tuple with the PurposeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurposeCode

`func (o *DataConsentRequestModel) SetPurposeCode(v string)`

SetPurposeCode sets PurposeCode field to given value.

### HasPurposeCode

`func (o *DataConsentRequestModel) HasPurposeCode() bool`

HasPurposeCode returns a boolean if a field has been set.

### SetPurposeCodeNil

`func (o *DataConsentRequestModel) SetPurposeCodeNil(b bool)`

 SetPurposeCodeNil sets the value for PurposeCode to be an explicit nil

### UnsetPurposeCode
`func (o *DataConsentRequestModel) UnsetPurposeCode()`

UnsetPurposeCode ensures that no value is present for PurposeCode, not even an explicit nil
### GetPurposeLink

`func (o *DataConsentRequestModel) GetPurposeLink() string`

GetPurposeLink returns the PurposeLink field if non-nil, zero value otherwise.

### GetPurposeLinkOk

`func (o *DataConsentRequestModel) GetPurposeLinkOk() (*string, bool)`

GetPurposeLinkOk returns a tuple with the PurposeLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurposeLink

`func (o *DataConsentRequestModel) SetPurposeLink(v string)`

SetPurposeLink sets PurposeLink field to given value.

### HasPurposeLink

`func (o *DataConsentRequestModel) HasPurposeLink() bool`

HasPurposeLink returns a boolean if a field has been set.

### SetPurposeLinkNil

`func (o *DataConsentRequestModel) SetPurposeLinkNil(b bool)`

 SetPurposeLinkNil sets the value for PurposeLink to be an explicit nil

### UnsetPurposeLink
`func (o *DataConsentRequestModel) UnsetPurposeLink()`

UnsetPurposeLink ensures that no value is present for PurposeLink, not even an explicit nil
### GetDataLifeUnit

`func (o *DataConsentRequestModel) GetDataLifeUnit() DataLifeUnit`

GetDataLifeUnit returns the DataLifeUnit field if non-nil, zero value otherwise.

### GetDataLifeUnitOk

`func (o *DataConsentRequestModel) GetDataLifeUnitOk() (*DataLifeUnit, bool)`

GetDataLifeUnitOk returns a tuple with the DataLifeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLifeUnit

`func (o *DataConsentRequestModel) SetDataLifeUnit(v DataLifeUnit)`

SetDataLifeUnit sets DataLifeUnit field to given value.

### HasDataLifeUnit

`func (o *DataConsentRequestModel) HasDataLifeUnit() bool`

HasDataLifeUnit returns a boolean if a field has been set.

### GetDataLifeValue

`func (o *DataConsentRequestModel) GetDataLifeValue() int32`

GetDataLifeValue returns the DataLifeValue field if non-nil, zero value otherwise.

### GetDataLifeValueOk

`func (o *DataConsentRequestModel) GetDataLifeValueOk() (*int32, bool)`

GetDataLifeValueOk returns a tuple with the DataLifeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLifeValue

`func (o *DataConsentRequestModel) SetDataLifeValue(v int32)`

SetDataLifeValue sets DataLifeValue field to given value.

### HasDataLifeValue

`func (o *DataConsentRequestModel) HasDataLifeValue() bool`

HasDataLifeValue returns a boolean if a field has been set.

### GetDataFetchFrequencyUnit

`func (o *DataConsentRequestModel) GetDataFetchFrequencyUnit() DataFetchFrequencyUnit`

GetDataFetchFrequencyUnit returns the DataFetchFrequencyUnit field if non-nil, zero value otherwise.

### GetDataFetchFrequencyUnitOk

`func (o *DataConsentRequestModel) GetDataFetchFrequencyUnitOk() (*DataFetchFrequencyUnit, bool)`

GetDataFetchFrequencyUnitOk returns a tuple with the DataFetchFrequencyUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFetchFrequencyUnit

`func (o *DataConsentRequestModel) SetDataFetchFrequencyUnit(v DataFetchFrequencyUnit)`

SetDataFetchFrequencyUnit sets DataFetchFrequencyUnit field to given value.

### HasDataFetchFrequencyUnit

`func (o *DataConsentRequestModel) HasDataFetchFrequencyUnit() bool`

HasDataFetchFrequencyUnit returns a boolean if a field has been set.

### GetDataFetchFrequencyUnitValue

`func (o *DataConsentRequestModel) GetDataFetchFrequencyUnitValue() int32`

GetDataFetchFrequencyUnitValue returns the DataFetchFrequencyUnitValue field if non-nil, zero value otherwise.

### GetDataFetchFrequencyUnitValueOk

`func (o *DataConsentRequestModel) GetDataFetchFrequencyUnitValueOk() (*int32, bool)`

GetDataFetchFrequencyUnitValueOk returns a tuple with the DataFetchFrequencyUnitValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFetchFrequencyUnitValue

`func (o *DataConsentRequestModel) SetDataFetchFrequencyUnitValue(v int32)`

SetDataFetchFrequencyUnitValue sets DataFetchFrequencyUnitValue field to given value.

### HasDataFetchFrequencyUnitValue

`func (o *DataConsentRequestModel) HasDataFetchFrequencyUnitValue() bool`

HasDataFetchFrequencyUnitValue returns a boolean if a field has been set.

### GetDataFetchType

`func (o *DataConsentRequestModel) GetDataFetchType() DataFetchType`

GetDataFetchType returns the DataFetchType field if non-nil, zero value otherwise.

### GetDataFetchTypeOk

`func (o *DataConsentRequestModel) GetDataFetchTypeOk() (*DataFetchType, bool)`

GetDataFetchTypeOk returns a tuple with the DataFetchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFetchType

`func (o *DataConsentRequestModel) SetDataFetchType(v DataFetchType)`

SetDataFetchType sets DataFetchType field to given value.

### HasDataFetchType

`func (o *DataConsentRequestModel) HasDataFetchType() bool`

HasDataFetchType returns a boolean if a field has been set.

### GetAgreementId

`func (o *DataConsentRequestModel) GetAgreementId() string`

GetAgreementId returns the AgreementId field if non-nil, zero value otherwise.

### GetAgreementIdOk

`func (o *DataConsentRequestModel) GetAgreementIdOk() (*string, bool)`

GetAgreementIdOk returns a tuple with the AgreementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementId

`func (o *DataConsentRequestModel) SetAgreementId(v string)`

SetAgreementId sets AgreementId field to given value.

### HasAgreementId

`func (o *DataConsentRequestModel) HasAgreementId() bool`

HasAgreementId returns a boolean if a field has been set.

### GetIdentityClaims

`func (o *DataConsentRequestModel) GetIdentityClaims() []IdentityClaim`

GetIdentityClaims returns the IdentityClaims field if non-nil, zero value otherwise.

### GetIdentityClaimsOk

`func (o *DataConsentRequestModel) GetIdentityClaimsOk() (*[]IdentityClaim, bool)`

GetIdentityClaimsOk returns a tuple with the IdentityClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityClaims

`func (o *DataConsentRequestModel) SetIdentityClaims(v []IdentityClaim)`

SetIdentityClaims sets IdentityClaims field to given value.

### HasIdentityClaims

`func (o *DataConsentRequestModel) HasIdentityClaims() bool`

HasIdentityClaims returns a boolean if a field has been set.

### SetIdentityClaimsNil

`func (o *DataConsentRequestModel) SetIdentityClaimsNil(b bool)`

 SetIdentityClaimsNil sets the value for IdentityClaims to be an explicit nil

### UnsetIdentityClaims
`func (o *DataConsentRequestModel) UnsetIdentityClaims()`

UnsetIdentityClaims ensures that no value is present for IdentityClaims, not even an explicit nil
### GetFinancialAccounts

`func (o *DataConsentRequestModel) GetFinancialAccounts() []DataConsentRequestedFaDto`

GetFinancialAccounts returns the FinancialAccounts field if non-nil, zero value otherwise.

### GetFinancialAccountsOk

`func (o *DataConsentRequestModel) GetFinancialAccountsOk() (*[]DataConsentRequestedFaDto, bool)`

GetFinancialAccountsOk returns a tuple with the FinancialAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialAccounts

`func (o *DataConsentRequestModel) SetFinancialAccounts(v []DataConsentRequestedFaDto)`

SetFinancialAccounts sets FinancialAccounts field to given value.

### HasFinancialAccounts

`func (o *DataConsentRequestModel) HasFinancialAccounts() bool`

HasFinancialAccounts returns a boolean if a field has been set.

### SetFinancialAccountsNil

`func (o *DataConsentRequestModel) SetFinancialAccountsNil(b bool)`

 SetFinancialAccountsNil sets the value for FinancialAccounts to be an explicit nil

### UnsetFinancialAccounts
`func (o *DataConsentRequestModel) UnsetFinancialAccounts()`

UnsetFinancialAccounts ensures that no value is present for FinancialAccounts, not even an explicit nil
### GetDocuments

`func (o *DataConsentRequestModel) GetDocuments() []DataConsentRequestedDocumentDto`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *DataConsentRequestModel) GetDocumentsOk() (*[]DataConsentRequestedDocumentDto, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *DataConsentRequestModel) SetDocuments(v []DataConsentRequestedDocumentDto)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *DataConsentRequestModel) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### SetDocumentsNil

`func (o *DataConsentRequestModel) SetDocumentsNil(b bool)`

 SetDocumentsNil sets the value for Documents to be an explicit nil

### UnsetDocuments
`func (o *DataConsentRequestModel) UnsetDocuments()`

UnsetDocuments ensures that no value is present for Documents, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



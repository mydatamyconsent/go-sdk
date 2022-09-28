# ConsentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Consent id. | 
**RequestId** | **string** | Consent request id. | 
**TemplateId** | Pointer to **string** | Consent request template id. | [optional] 
**Title** | **string** | Consent title. | 
**Description** | **string** | Consent description. | 
**Purpose** | Pointer to **string** | Consent purpose. | [optional] 
**Status** | [**DataConsentStatus**](DataConsentStatus.md) |  | 
**TransactionId** | Pointer to **string** | Transaction id. | [optional] 
**ApprovedAtUtc** | **time.Time** | Consent approval datetime in UTC timezone. | 
**DataAccessExpiresAtUtc** | **time.Time** | Data access expiration datetime in UTC timezone. | 
**RevokedAtUtc** | Pointer to **time.Time** | Consent revocation datetime in UTC timezone. | [optional] 
**Collectables** | [**[]CollectibleTypes**](CollectibleTypes.md) | List of supported collectible types. | 
**Identifiers** | Pointer to [**[]ConsentedIdentifier**](ConsentedIdentifier.md) | Consented identity details. | [optional] 
**Documents** | Pointer to [**[]ConsentedDocument**](ConsentedDocument.md) | List of consented documents. | [optional] 
**MedicalRecords** | Pointer to [**[]ConsentedMedicalRecord**](ConsentedMedicalRecord.md) | List of consented medical records. | [optional] 
**FinancialAccounts** | Pointer to [**[]ConsentedFinancialAccount**](ConsentedFinancialAccount.md) | List of consented financial accounts. | [optional] 

## Methods

### NewConsentDetails

`func NewConsentDetails(id string, requestId string, title string, description string, status DataConsentStatus, approvedAtUtc time.Time, dataAccessExpiresAtUtc time.Time, collectables []CollectibleTypes, ) *ConsentDetails`

NewConsentDetails instantiates a new ConsentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsentDetailsWithDefaults

`func NewConsentDetailsWithDefaults() *ConsentDetails`

NewConsentDetailsWithDefaults instantiates a new ConsentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConsentDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsentDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsentDetails) SetId(v string)`

SetId sets Id field to given value.


### GetRequestId

`func (o *ConsentDetails) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ConsentDetails) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ConsentDetails) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetTemplateId

`func (o *ConsentDetails) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *ConsentDetails) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *ConsentDetails) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *ConsentDetails) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetTitle

`func (o *ConsentDetails) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ConsentDetails) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ConsentDetails) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *ConsentDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConsentDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConsentDetails) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPurpose

`func (o *ConsentDetails) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *ConsentDetails) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *ConsentDetails) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *ConsentDetails) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetStatus

`func (o *ConsentDetails) GetStatus() DataConsentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConsentDetails) GetStatusOk() (*DataConsentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConsentDetails) SetStatus(v DataConsentStatus)`

SetStatus sets Status field to given value.


### GetTransactionId

`func (o *ConsentDetails) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ConsentDetails) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ConsentDetails) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *ConsentDetails) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetApprovedAtUtc

`func (o *ConsentDetails) GetApprovedAtUtc() time.Time`

GetApprovedAtUtc returns the ApprovedAtUtc field if non-nil, zero value otherwise.

### GetApprovedAtUtcOk

`func (o *ConsentDetails) GetApprovedAtUtcOk() (*time.Time, bool)`

GetApprovedAtUtcOk returns a tuple with the ApprovedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAtUtc

`func (o *ConsentDetails) SetApprovedAtUtc(v time.Time)`

SetApprovedAtUtc sets ApprovedAtUtc field to given value.


### GetDataAccessExpiresAtUtc

`func (o *ConsentDetails) GetDataAccessExpiresAtUtc() time.Time`

GetDataAccessExpiresAtUtc returns the DataAccessExpiresAtUtc field if non-nil, zero value otherwise.

### GetDataAccessExpiresAtUtcOk

`func (o *ConsentDetails) GetDataAccessExpiresAtUtcOk() (*time.Time, bool)`

GetDataAccessExpiresAtUtcOk returns a tuple with the DataAccessExpiresAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataAccessExpiresAtUtc

`func (o *ConsentDetails) SetDataAccessExpiresAtUtc(v time.Time)`

SetDataAccessExpiresAtUtc sets DataAccessExpiresAtUtc field to given value.


### GetRevokedAtUtc

`func (o *ConsentDetails) GetRevokedAtUtc() time.Time`

GetRevokedAtUtc returns the RevokedAtUtc field if non-nil, zero value otherwise.

### GetRevokedAtUtcOk

`func (o *ConsentDetails) GetRevokedAtUtcOk() (*time.Time, bool)`

GetRevokedAtUtcOk returns a tuple with the RevokedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedAtUtc

`func (o *ConsentDetails) SetRevokedAtUtc(v time.Time)`

SetRevokedAtUtc sets RevokedAtUtc field to given value.

### HasRevokedAtUtc

`func (o *ConsentDetails) HasRevokedAtUtc() bool`

HasRevokedAtUtc returns a boolean if a field has been set.

### GetCollectables

`func (o *ConsentDetails) GetCollectables() []CollectibleTypes`

GetCollectables returns the Collectables field if non-nil, zero value otherwise.

### GetCollectablesOk

`func (o *ConsentDetails) GetCollectablesOk() (*[]CollectibleTypes, bool)`

GetCollectablesOk returns a tuple with the Collectables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectables

`func (o *ConsentDetails) SetCollectables(v []CollectibleTypes)`

SetCollectables sets Collectables field to given value.


### GetIdentifiers

`func (o *ConsentDetails) GetIdentifiers() []ConsentedIdentifier`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *ConsentDetails) GetIdentifiersOk() (*[]ConsentedIdentifier, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *ConsentDetails) SetIdentifiers(v []ConsentedIdentifier)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *ConsentDetails) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### GetDocuments

`func (o *ConsentDetails) GetDocuments() []ConsentedDocument`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *ConsentDetails) GetDocumentsOk() (*[]ConsentedDocument, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *ConsentDetails) SetDocuments(v []ConsentedDocument)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *ConsentDetails) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetMedicalRecords

`func (o *ConsentDetails) GetMedicalRecords() []ConsentedMedicalRecord`

GetMedicalRecords returns the MedicalRecords field if non-nil, zero value otherwise.

### GetMedicalRecordsOk

`func (o *ConsentDetails) GetMedicalRecordsOk() (*[]ConsentedMedicalRecord, bool)`

GetMedicalRecordsOk returns a tuple with the MedicalRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedicalRecords

`func (o *ConsentDetails) SetMedicalRecords(v []ConsentedMedicalRecord)`

SetMedicalRecords sets MedicalRecords field to given value.

### HasMedicalRecords

`func (o *ConsentDetails) HasMedicalRecords() bool`

HasMedicalRecords returns a boolean if a field has been set.

### GetFinancialAccounts

`func (o *ConsentDetails) GetFinancialAccounts() []ConsentedFinancialAccount`

GetFinancialAccounts returns the FinancialAccounts field if non-nil, zero value otherwise.

### GetFinancialAccountsOk

`func (o *ConsentDetails) GetFinancialAccountsOk() (*[]ConsentedFinancialAccount, bool)`

GetFinancialAccountsOk returns a tuple with the FinancialAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialAccounts

`func (o *ConsentDetails) SetFinancialAccounts(v []ConsentedFinancialAccount)`

SetFinancialAccounts sets FinancialAccounts field to given value.

### HasFinancialAccounts

`func (o *ConsentDetails) HasFinancialAccounts() bool`

HasFinancialAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



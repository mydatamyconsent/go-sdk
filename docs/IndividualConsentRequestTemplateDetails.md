# IndividualConsentRequestTemplateDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Consent request template id. | 
**Title** | **string** | Consent request template title. | 
**Description** | **string** | Consent request template description. | 
**Purpose** | Pointer to **string** | Consent request template purpose. | [optional] 
**ShortId** | **string** | Consent request template short Id. | 
**Status** | [**ConsentRequestTemplateStatus**](ConsentRequestTemplateStatus.md) |  | 
**DataLife** | Pointer to [**IndividualConsentRequestTemplateDetailsDataLife**](IndividualConsentRequestTemplateDetailsDataLife.md) |  | [optional] 
**RequestLife** | Pointer to [**IndividualConsentRequestTemplateDetailsRequestLife**](IndividualConsentRequestTemplateDetailsRequestLife.md) |  | [optional] 
**DataFrequency** | Pointer to [**IndividualConsentRequestTemplateDetailsDataFrequency**](IndividualConsentRequestTemplateDetailsDataFrequency.md) |  | [optional] 
**Identifiers** | Pointer to [**[]IdentityField**](IdentityField.md) | Consent request template identity fields. | [optional] 
**Documents** | Pointer to [**[]DocumentField**](DocumentField.md) | Consent request template document fields. | [optional] 
**MedicalRecords** | Pointer to [**[]MedicalRecordField**](MedicalRecordField.md) | Consent request template medical record fields. | [optional] 
**FinancialAccounts** | Pointer to [**[]FinancialAccountField**](FinancialAccountField.md) | Consent request template financial account fields. | [optional] 
**CreatedBy** | **string** | Consent request template created by user. | 
**CreatedAtUtc** | **time.Time** | Consent request template created datetime in UTC timezone. | 
**ApprovedAtUtc** | Pointer to **time.Time** | Consent request template approval datetime in UTC timezone. | [optional] 
**PublishedAtUtc** | Pointer to **time.Time** | Consent request template published datetime in UTC timezone. | [optional] 
**RejectedAtUtc** | Pointer to **time.Time** | Consent request template rejection datetime in UTC timezone. | [optional] 
**RejectionReason** | Pointer to **string** | Consent request template rejection reason. | [optional] 

## Methods

### NewIndividualConsentRequestTemplateDetails

`func NewIndividualConsentRequestTemplateDetails(id string, title string, description string, shortId string, status ConsentRequestTemplateStatus, createdBy string, createdAtUtc time.Time, ) *IndividualConsentRequestTemplateDetails`

NewIndividualConsentRequestTemplateDetails instantiates a new IndividualConsentRequestTemplateDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndividualConsentRequestTemplateDetailsWithDefaults

`func NewIndividualConsentRequestTemplateDetailsWithDefaults() *IndividualConsentRequestTemplateDetails`

NewIndividualConsentRequestTemplateDetailsWithDefaults instantiates a new IndividualConsentRequestTemplateDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IndividualConsentRequestTemplateDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IndividualConsentRequestTemplateDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IndividualConsentRequestTemplateDetails) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *IndividualConsentRequestTemplateDetails) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IndividualConsentRequestTemplateDetails) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IndividualConsentRequestTemplateDetails) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *IndividualConsentRequestTemplateDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IndividualConsentRequestTemplateDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IndividualConsentRequestTemplateDetails) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPurpose

`func (o *IndividualConsentRequestTemplateDetails) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *IndividualConsentRequestTemplateDetails) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *IndividualConsentRequestTemplateDetails) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *IndividualConsentRequestTemplateDetails) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetShortId

`func (o *IndividualConsentRequestTemplateDetails) GetShortId() string`

GetShortId returns the ShortId field if non-nil, zero value otherwise.

### GetShortIdOk

`func (o *IndividualConsentRequestTemplateDetails) GetShortIdOk() (*string, bool)`

GetShortIdOk returns a tuple with the ShortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortId

`func (o *IndividualConsentRequestTemplateDetails) SetShortId(v string)`

SetShortId sets ShortId field to given value.


### GetStatus

`func (o *IndividualConsentRequestTemplateDetails) GetStatus() ConsentRequestTemplateStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IndividualConsentRequestTemplateDetails) GetStatusOk() (*ConsentRequestTemplateStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IndividualConsentRequestTemplateDetails) SetStatus(v ConsentRequestTemplateStatus)`

SetStatus sets Status field to given value.


### GetDataLife

`func (o *IndividualConsentRequestTemplateDetails) GetDataLife() IndividualConsentRequestTemplateDetailsDataLife`

GetDataLife returns the DataLife field if non-nil, zero value otherwise.

### GetDataLifeOk

`func (o *IndividualConsentRequestTemplateDetails) GetDataLifeOk() (*IndividualConsentRequestTemplateDetailsDataLife, bool)`

GetDataLifeOk returns a tuple with the DataLife field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLife

`func (o *IndividualConsentRequestTemplateDetails) SetDataLife(v IndividualConsentRequestTemplateDetailsDataLife)`

SetDataLife sets DataLife field to given value.

### HasDataLife

`func (o *IndividualConsentRequestTemplateDetails) HasDataLife() bool`

HasDataLife returns a boolean if a field has been set.

### GetRequestLife

`func (o *IndividualConsentRequestTemplateDetails) GetRequestLife() IndividualConsentRequestTemplateDetailsRequestLife`

GetRequestLife returns the RequestLife field if non-nil, zero value otherwise.

### GetRequestLifeOk

`func (o *IndividualConsentRequestTemplateDetails) GetRequestLifeOk() (*IndividualConsentRequestTemplateDetailsRequestLife, bool)`

GetRequestLifeOk returns a tuple with the RequestLife field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestLife

`func (o *IndividualConsentRequestTemplateDetails) SetRequestLife(v IndividualConsentRequestTemplateDetailsRequestLife)`

SetRequestLife sets RequestLife field to given value.

### HasRequestLife

`func (o *IndividualConsentRequestTemplateDetails) HasRequestLife() bool`

HasRequestLife returns a boolean if a field has been set.

### GetDataFrequency

`func (o *IndividualConsentRequestTemplateDetails) GetDataFrequency() IndividualConsentRequestTemplateDetailsDataFrequency`

GetDataFrequency returns the DataFrequency field if non-nil, zero value otherwise.

### GetDataFrequencyOk

`func (o *IndividualConsentRequestTemplateDetails) GetDataFrequencyOk() (*IndividualConsentRequestTemplateDetailsDataFrequency, bool)`

GetDataFrequencyOk returns a tuple with the DataFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFrequency

`func (o *IndividualConsentRequestTemplateDetails) SetDataFrequency(v IndividualConsentRequestTemplateDetailsDataFrequency)`

SetDataFrequency sets DataFrequency field to given value.

### HasDataFrequency

`func (o *IndividualConsentRequestTemplateDetails) HasDataFrequency() bool`

HasDataFrequency returns a boolean if a field has been set.

### GetIdentifiers

`func (o *IndividualConsentRequestTemplateDetails) GetIdentifiers() []IdentityField`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *IndividualConsentRequestTemplateDetails) GetIdentifiersOk() (*[]IdentityField, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *IndividualConsentRequestTemplateDetails) SetIdentifiers(v []IdentityField)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *IndividualConsentRequestTemplateDetails) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### GetDocuments

`func (o *IndividualConsentRequestTemplateDetails) GetDocuments() []DocumentField`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *IndividualConsentRequestTemplateDetails) GetDocumentsOk() (*[]DocumentField, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *IndividualConsentRequestTemplateDetails) SetDocuments(v []DocumentField)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *IndividualConsentRequestTemplateDetails) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetMedicalRecords

`func (o *IndividualConsentRequestTemplateDetails) GetMedicalRecords() []MedicalRecordField`

GetMedicalRecords returns the MedicalRecords field if non-nil, zero value otherwise.

### GetMedicalRecordsOk

`func (o *IndividualConsentRequestTemplateDetails) GetMedicalRecordsOk() (*[]MedicalRecordField, bool)`

GetMedicalRecordsOk returns a tuple with the MedicalRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedicalRecords

`func (o *IndividualConsentRequestTemplateDetails) SetMedicalRecords(v []MedicalRecordField)`

SetMedicalRecords sets MedicalRecords field to given value.

### HasMedicalRecords

`func (o *IndividualConsentRequestTemplateDetails) HasMedicalRecords() bool`

HasMedicalRecords returns a boolean if a field has been set.

### GetFinancialAccounts

`func (o *IndividualConsentRequestTemplateDetails) GetFinancialAccounts() []FinancialAccountField`

GetFinancialAccounts returns the FinancialAccounts field if non-nil, zero value otherwise.

### GetFinancialAccountsOk

`func (o *IndividualConsentRequestTemplateDetails) GetFinancialAccountsOk() (*[]FinancialAccountField, bool)`

GetFinancialAccountsOk returns a tuple with the FinancialAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialAccounts

`func (o *IndividualConsentRequestTemplateDetails) SetFinancialAccounts(v []FinancialAccountField)`

SetFinancialAccounts sets FinancialAccounts field to given value.

### HasFinancialAccounts

`func (o *IndividualConsentRequestTemplateDetails) HasFinancialAccounts() bool`

HasFinancialAccounts returns a boolean if a field has been set.

### GetCreatedBy

`func (o *IndividualConsentRequestTemplateDetails) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *IndividualConsentRequestTemplateDetails) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *IndividualConsentRequestTemplateDetails) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAtUtc

`func (o *IndividualConsentRequestTemplateDetails) GetCreatedAtUtc() time.Time`

GetCreatedAtUtc returns the CreatedAtUtc field if non-nil, zero value otherwise.

### GetCreatedAtUtcOk

`func (o *IndividualConsentRequestTemplateDetails) GetCreatedAtUtcOk() (*time.Time, bool)`

GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtUtc

`func (o *IndividualConsentRequestTemplateDetails) SetCreatedAtUtc(v time.Time)`

SetCreatedAtUtc sets CreatedAtUtc field to given value.


### GetApprovedAtUtc

`func (o *IndividualConsentRequestTemplateDetails) GetApprovedAtUtc() time.Time`

GetApprovedAtUtc returns the ApprovedAtUtc field if non-nil, zero value otherwise.

### GetApprovedAtUtcOk

`func (o *IndividualConsentRequestTemplateDetails) GetApprovedAtUtcOk() (*time.Time, bool)`

GetApprovedAtUtcOk returns a tuple with the ApprovedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAtUtc

`func (o *IndividualConsentRequestTemplateDetails) SetApprovedAtUtc(v time.Time)`

SetApprovedAtUtc sets ApprovedAtUtc field to given value.

### HasApprovedAtUtc

`func (o *IndividualConsentRequestTemplateDetails) HasApprovedAtUtc() bool`

HasApprovedAtUtc returns a boolean if a field has been set.

### GetPublishedAtUtc

`func (o *IndividualConsentRequestTemplateDetails) GetPublishedAtUtc() time.Time`

GetPublishedAtUtc returns the PublishedAtUtc field if non-nil, zero value otherwise.

### GetPublishedAtUtcOk

`func (o *IndividualConsentRequestTemplateDetails) GetPublishedAtUtcOk() (*time.Time, bool)`

GetPublishedAtUtcOk returns a tuple with the PublishedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAtUtc

`func (o *IndividualConsentRequestTemplateDetails) SetPublishedAtUtc(v time.Time)`

SetPublishedAtUtc sets PublishedAtUtc field to given value.

### HasPublishedAtUtc

`func (o *IndividualConsentRequestTemplateDetails) HasPublishedAtUtc() bool`

HasPublishedAtUtc returns a boolean if a field has been set.

### GetRejectedAtUtc

`func (o *IndividualConsentRequestTemplateDetails) GetRejectedAtUtc() time.Time`

GetRejectedAtUtc returns the RejectedAtUtc field if non-nil, zero value otherwise.

### GetRejectedAtUtcOk

`func (o *IndividualConsentRequestTemplateDetails) GetRejectedAtUtcOk() (*time.Time, bool)`

GetRejectedAtUtcOk returns a tuple with the RejectedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedAtUtc

`func (o *IndividualConsentRequestTemplateDetails) SetRejectedAtUtc(v time.Time)`

SetRejectedAtUtc sets RejectedAtUtc field to given value.

### HasRejectedAtUtc

`func (o *IndividualConsentRequestTemplateDetails) HasRejectedAtUtc() bool`

HasRejectedAtUtc returns a boolean if a field has been set.

### GetRejectionReason

`func (o *IndividualConsentRequestTemplateDetails) GetRejectionReason() string`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *IndividualConsentRequestTemplateDetails) GetRejectionReasonOk() (*string, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *IndividualConsentRequestTemplateDetails) SetRejectionReason(v string)`

SetRejectionReason sets RejectionReason field to given value.

### HasRejectionReason

`func (o *IndividualConsentRequestTemplateDetails) HasRejectionReason() bool`

HasRejectionReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



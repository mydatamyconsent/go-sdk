# OrganizationConsentRequestTemplateDetails

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
**FinancialAccounts** | Pointer to [**[]FinancialAccountField**](FinancialAccountField.md) | Consent request template financial account fields. | [optional] 
**CreatedBy** | **string** | Consent request template created by user. | 
**CreatedAtUtc** | **time.Time** | Consent request template created datetime in UTC timezone. | 
**ApprovedAtUtc** | Pointer to **time.Time** | Consent request template approval datetime in UTC timezone. | [optional] 
**PublishedAtUtc** | Pointer to **time.Time** | Consent request template published datetime in UTC timezone. | [optional] 
**RejectedAtUtc** | Pointer to **time.Time** | Consent request template rejection datetime in UTC timezone. | [optional] 
**RejectionReason** | Pointer to **string** | Consent request template rejection reason. | [optional] 

## Methods

### NewOrganizationConsentRequestTemplateDetails

`func NewOrganizationConsentRequestTemplateDetails(id string, title string, description string, shortId string, status ConsentRequestTemplateStatus, createdBy string, createdAtUtc time.Time, ) *OrganizationConsentRequestTemplateDetails`

NewOrganizationConsentRequestTemplateDetails instantiates a new OrganizationConsentRequestTemplateDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationConsentRequestTemplateDetailsWithDefaults

`func NewOrganizationConsentRequestTemplateDetailsWithDefaults() *OrganizationConsentRequestTemplateDetails`

NewOrganizationConsentRequestTemplateDetailsWithDefaults instantiates a new OrganizationConsentRequestTemplateDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationConsentRequestTemplateDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationConsentRequestTemplateDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationConsentRequestTemplateDetails) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *OrganizationConsentRequestTemplateDetails) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OrganizationConsentRequestTemplateDetails) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OrganizationConsentRequestTemplateDetails) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *OrganizationConsentRequestTemplateDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationConsentRequestTemplateDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationConsentRequestTemplateDetails) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPurpose

`func (o *OrganizationConsentRequestTemplateDetails) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *OrganizationConsentRequestTemplateDetails) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *OrganizationConsentRequestTemplateDetails) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *OrganizationConsentRequestTemplateDetails) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetShortId

`func (o *OrganizationConsentRequestTemplateDetails) GetShortId() string`

GetShortId returns the ShortId field if non-nil, zero value otherwise.

### GetShortIdOk

`func (o *OrganizationConsentRequestTemplateDetails) GetShortIdOk() (*string, bool)`

GetShortIdOk returns a tuple with the ShortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortId

`func (o *OrganizationConsentRequestTemplateDetails) SetShortId(v string)`

SetShortId sets ShortId field to given value.


### GetStatus

`func (o *OrganizationConsentRequestTemplateDetails) GetStatus() ConsentRequestTemplateStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrganizationConsentRequestTemplateDetails) GetStatusOk() (*ConsentRequestTemplateStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrganizationConsentRequestTemplateDetails) SetStatus(v ConsentRequestTemplateStatus)`

SetStatus sets Status field to given value.


### GetDataLife

`func (o *OrganizationConsentRequestTemplateDetails) GetDataLife() IndividualConsentRequestTemplateDetailsDataLife`

GetDataLife returns the DataLife field if non-nil, zero value otherwise.

### GetDataLifeOk

`func (o *OrganizationConsentRequestTemplateDetails) GetDataLifeOk() (*IndividualConsentRequestTemplateDetailsDataLife, bool)`

GetDataLifeOk returns a tuple with the DataLife field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLife

`func (o *OrganizationConsentRequestTemplateDetails) SetDataLife(v IndividualConsentRequestTemplateDetailsDataLife)`

SetDataLife sets DataLife field to given value.

### HasDataLife

`func (o *OrganizationConsentRequestTemplateDetails) HasDataLife() bool`

HasDataLife returns a boolean if a field has been set.

### GetRequestLife

`func (o *OrganizationConsentRequestTemplateDetails) GetRequestLife() IndividualConsentRequestTemplateDetailsRequestLife`

GetRequestLife returns the RequestLife field if non-nil, zero value otherwise.

### GetRequestLifeOk

`func (o *OrganizationConsentRequestTemplateDetails) GetRequestLifeOk() (*IndividualConsentRequestTemplateDetailsRequestLife, bool)`

GetRequestLifeOk returns a tuple with the RequestLife field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestLife

`func (o *OrganizationConsentRequestTemplateDetails) SetRequestLife(v IndividualConsentRequestTemplateDetailsRequestLife)`

SetRequestLife sets RequestLife field to given value.

### HasRequestLife

`func (o *OrganizationConsentRequestTemplateDetails) HasRequestLife() bool`

HasRequestLife returns a boolean if a field has been set.

### GetDataFrequency

`func (o *OrganizationConsentRequestTemplateDetails) GetDataFrequency() IndividualConsentRequestTemplateDetailsDataFrequency`

GetDataFrequency returns the DataFrequency field if non-nil, zero value otherwise.

### GetDataFrequencyOk

`func (o *OrganizationConsentRequestTemplateDetails) GetDataFrequencyOk() (*IndividualConsentRequestTemplateDetailsDataFrequency, bool)`

GetDataFrequencyOk returns a tuple with the DataFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFrequency

`func (o *OrganizationConsentRequestTemplateDetails) SetDataFrequency(v IndividualConsentRequestTemplateDetailsDataFrequency)`

SetDataFrequency sets DataFrequency field to given value.

### HasDataFrequency

`func (o *OrganizationConsentRequestTemplateDetails) HasDataFrequency() bool`

HasDataFrequency returns a boolean if a field has been set.

### GetIdentifiers

`func (o *OrganizationConsentRequestTemplateDetails) GetIdentifiers() []IdentityField`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *OrganizationConsentRequestTemplateDetails) GetIdentifiersOk() (*[]IdentityField, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *OrganizationConsentRequestTemplateDetails) SetIdentifiers(v []IdentityField)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *OrganizationConsentRequestTemplateDetails) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### GetDocuments

`func (o *OrganizationConsentRequestTemplateDetails) GetDocuments() []DocumentField`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *OrganizationConsentRequestTemplateDetails) GetDocumentsOk() (*[]DocumentField, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *OrganizationConsentRequestTemplateDetails) SetDocuments(v []DocumentField)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *OrganizationConsentRequestTemplateDetails) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetFinancialAccounts

`func (o *OrganizationConsentRequestTemplateDetails) GetFinancialAccounts() []FinancialAccountField`

GetFinancialAccounts returns the FinancialAccounts field if non-nil, zero value otherwise.

### GetFinancialAccountsOk

`func (o *OrganizationConsentRequestTemplateDetails) GetFinancialAccountsOk() (*[]FinancialAccountField, bool)`

GetFinancialAccountsOk returns a tuple with the FinancialAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialAccounts

`func (o *OrganizationConsentRequestTemplateDetails) SetFinancialAccounts(v []FinancialAccountField)`

SetFinancialAccounts sets FinancialAccounts field to given value.

### HasFinancialAccounts

`func (o *OrganizationConsentRequestTemplateDetails) HasFinancialAccounts() bool`

HasFinancialAccounts returns a boolean if a field has been set.

### GetCreatedBy

`func (o *OrganizationConsentRequestTemplateDetails) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *OrganizationConsentRequestTemplateDetails) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *OrganizationConsentRequestTemplateDetails) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAtUtc

`func (o *OrganizationConsentRequestTemplateDetails) GetCreatedAtUtc() time.Time`

GetCreatedAtUtc returns the CreatedAtUtc field if non-nil, zero value otherwise.

### GetCreatedAtUtcOk

`func (o *OrganizationConsentRequestTemplateDetails) GetCreatedAtUtcOk() (*time.Time, bool)`

GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtUtc

`func (o *OrganizationConsentRequestTemplateDetails) SetCreatedAtUtc(v time.Time)`

SetCreatedAtUtc sets CreatedAtUtc field to given value.


### GetApprovedAtUtc

`func (o *OrganizationConsentRequestTemplateDetails) GetApprovedAtUtc() time.Time`

GetApprovedAtUtc returns the ApprovedAtUtc field if non-nil, zero value otherwise.

### GetApprovedAtUtcOk

`func (o *OrganizationConsentRequestTemplateDetails) GetApprovedAtUtcOk() (*time.Time, bool)`

GetApprovedAtUtcOk returns a tuple with the ApprovedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAtUtc

`func (o *OrganizationConsentRequestTemplateDetails) SetApprovedAtUtc(v time.Time)`

SetApprovedAtUtc sets ApprovedAtUtc field to given value.

### HasApprovedAtUtc

`func (o *OrganizationConsentRequestTemplateDetails) HasApprovedAtUtc() bool`

HasApprovedAtUtc returns a boolean if a field has been set.

### GetPublishedAtUtc

`func (o *OrganizationConsentRequestTemplateDetails) GetPublishedAtUtc() time.Time`

GetPublishedAtUtc returns the PublishedAtUtc field if non-nil, zero value otherwise.

### GetPublishedAtUtcOk

`func (o *OrganizationConsentRequestTemplateDetails) GetPublishedAtUtcOk() (*time.Time, bool)`

GetPublishedAtUtcOk returns a tuple with the PublishedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAtUtc

`func (o *OrganizationConsentRequestTemplateDetails) SetPublishedAtUtc(v time.Time)`

SetPublishedAtUtc sets PublishedAtUtc field to given value.

### HasPublishedAtUtc

`func (o *OrganizationConsentRequestTemplateDetails) HasPublishedAtUtc() bool`

HasPublishedAtUtc returns a boolean if a field has been set.

### GetRejectedAtUtc

`func (o *OrganizationConsentRequestTemplateDetails) GetRejectedAtUtc() time.Time`

GetRejectedAtUtc returns the RejectedAtUtc field if non-nil, zero value otherwise.

### GetRejectedAtUtcOk

`func (o *OrganizationConsentRequestTemplateDetails) GetRejectedAtUtcOk() (*time.Time, bool)`

GetRejectedAtUtcOk returns a tuple with the RejectedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedAtUtc

`func (o *OrganizationConsentRequestTemplateDetails) SetRejectedAtUtc(v time.Time)`

SetRejectedAtUtc sets RejectedAtUtc field to given value.

### HasRejectedAtUtc

`func (o *OrganizationConsentRequestTemplateDetails) HasRejectedAtUtc() bool`

HasRejectedAtUtc returns a boolean if a field has been set.

### GetRejectionReason

`func (o *OrganizationConsentRequestTemplateDetails) GetRejectionReason() string`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *OrganizationConsentRequestTemplateDetails) GetRejectionReasonOk() (*string, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *OrganizationConsentRequestTemplateDetails) SetRejectionReason(v string)`

SetRejectionReason sets RejectionReason field to given value.

### HasRejectionReason

`func (o *OrganizationConsentRequestTemplateDetails) HasRejectionReason() bool`

HasRejectionReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



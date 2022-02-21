# GetConsentTemplateDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ConsentPurpose** | Pointer to **NullableString** |  | [optional] 
**Collectables** | Pointer to [**[]CollectibleTypes**](CollectibleTypes.md) |  | [optional] 
**FetchType** | Pointer to [**FetchTypes**](FetchTypes.md) |  | [optional] 
**ShortId** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAtUtc** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**TemplateType** | Pointer to [**ConsentTemplateTypes**](ConsentTemplateTypes.md) |  | [optional] 
**DataLife** | Pointer to [**Life**](Life.md) |  | [optional] 
**RequestLife** | Pointer to [**Life**](Life.md) |  | [optional] 
**Frequency** | Pointer to [**Life**](Life.md) |  | [optional] 
**Identity** | Pointer to [**[]IdentitySupportedFields**](IdentitySupportedFields.md) |  | [optional] 
**Documents** | Pointer to [**[]Document**](Document.md) |  | [optional] 
**Financials** | Pointer to [**[]Financial**](Financial.md) |  | [optional] 
**HealthRecords** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ApprovedBy** | Pointer to **NullableString** |  | [optional] 
**ApprovedAtUtc** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewGetConsentTemplateDetailsDto

`func NewGetConsentTemplateDetailsDto() *GetConsentTemplateDetailsDto`

NewGetConsentTemplateDetailsDto instantiates a new GetConsentTemplateDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConsentTemplateDetailsDtoWithDefaults

`func NewGetConsentTemplateDetailsDtoWithDefaults() *GetConsentTemplateDetailsDto`

NewGetConsentTemplateDetailsDtoWithDefaults instantiates a new GetConsentTemplateDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetConsentTemplateDetailsDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetConsentTemplateDetailsDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetConsentTemplateDetailsDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetConsentTemplateDetailsDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetConsentTemplateDetailsDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetConsentTemplateDetailsDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetConsentTemplateDetailsDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetConsentTemplateDetailsDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GetConsentTemplateDetailsDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GetConsentTemplateDetailsDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *GetConsentTemplateDetailsDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetConsentTemplateDetailsDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetConsentTemplateDetailsDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetConsentTemplateDetailsDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetConsentTemplateDetailsDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetConsentTemplateDetailsDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetConsentPurpose

`func (o *GetConsentTemplateDetailsDto) GetConsentPurpose() string`

GetConsentPurpose returns the ConsentPurpose field if non-nil, zero value otherwise.

### GetConsentPurposeOk

`func (o *GetConsentTemplateDetailsDto) GetConsentPurposeOk() (*string, bool)`

GetConsentPurposeOk returns a tuple with the ConsentPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentPurpose

`func (o *GetConsentTemplateDetailsDto) SetConsentPurpose(v string)`

SetConsentPurpose sets ConsentPurpose field to given value.

### HasConsentPurpose

`func (o *GetConsentTemplateDetailsDto) HasConsentPurpose() bool`

HasConsentPurpose returns a boolean if a field has been set.

### SetConsentPurposeNil

`func (o *GetConsentTemplateDetailsDto) SetConsentPurposeNil(b bool)`

 SetConsentPurposeNil sets the value for ConsentPurpose to be an explicit nil

### UnsetConsentPurpose
`func (o *GetConsentTemplateDetailsDto) UnsetConsentPurpose()`

UnsetConsentPurpose ensures that no value is present for ConsentPurpose, not even an explicit nil
### GetCollectables

`func (o *GetConsentTemplateDetailsDto) GetCollectables() []CollectibleTypes`

GetCollectables returns the Collectables field if non-nil, zero value otherwise.

### GetCollectablesOk

`func (o *GetConsentTemplateDetailsDto) GetCollectablesOk() (*[]CollectibleTypes, bool)`

GetCollectablesOk returns a tuple with the Collectables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectables

`func (o *GetConsentTemplateDetailsDto) SetCollectables(v []CollectibleTypes)`

SetCollectables sets Collectables field to given value.

### HasCollectables

`func (o *GetConsentTemplateDetailsDto) HasCollectables() bool`

HasCollectables returns a boolean if a field has been set.

### SetCollectablesNil

`func (o *GetConsentTemplateDetailsDto) SetCollectablesNil(b bool)`

 SetCollectablesNil sets the value for Collectables to be an explicit nil

### UnsetCollectables
`func (o *GetConsentTemplateDetailsDto) UnsetCollectables()`

UnsetCollectables ensures that no value is present for Collectables, not even an explicit nil
### GetFetchType

`func (o *GetConsentTemplateDetailsDto) GetFetchType() FetchTypes`

GetFetchType returns the FetchType field if non-nil, zero value otherwise.

### GetFetchTypeOk

`func (o *GetConsentTemplateDetailsDto) GetFetchTypeOk() (*FetchTypes, bool)`

GetFetchTypeOk returns a tuple with the FetchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchType

`func (o *GetConsentTemplateDetailsDto) SetFetchType(v FetchTypes)`

SetFetchType sets FetchType field to given value.

### HasFetchType

`func (o *GetConsentTemplateDetailsDto) HasFetchType() bool`

HasFetchType returns a boolean if a field has been set.

### GetShortId

`func (o *GetConsentTemplateDetailsDto) GetShortId() string`

GetShortId returns the ShortId field if non-nil, zero value otherwise.

### GetShortIdOk

`func (o *GetConsentTemplateDetailsDto) GetShortIdOk() (*string, bool)`

GetShortIdOk returns a tuple with the ShortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortId

`func (o *GetConsentTemplateDetailsDto) SetShortId(v string)`

SetShortId sets ShortId field to given value.

### HasShortId

`func (o *GetConsentTemplateDetailsDto) HasShortId() bool`

HasShortId returns a boolean if a field has been set.

### SetShortIdNil

`func (o *GetConsentTemplateDetailsDto) SetShortIdNil(b bool)`

 SetShortIdNil sets the value for ShortId to be an explicit nil

### UnsetShortId
`func (o *GetConsentTemplateDetailsDto) UnsetShortId()`

UnsetShortId ensures that no value is present for ShortId, not even an explicit nil
### GetCreatedBy

`func (o *GetConsentTemplateDetailsDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetConsentTemplateDetailsDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetConsentTemplateDetailsDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetConsentTemplateDetailsDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *GetConsentTemplateDetailsDto) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *GetConsentTemplateDetailsDto) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedAtUtc

`func (o *GetConsentTemplateDetailsDto) GetCreatedAtUtc() time.Time`

GetCreatedAtUtc returns the CreatedAtUtc field if non-nil, zero value otherwise.

### GetCreatedAtUtcOk

`func (o *GetConsentTemplateDetailsDto) GetCreatedAtUtcOk() (*time.Time, bool)`

GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtUtc

`func (o *GetConsentTemplateDetailsDto) SetCreatedAtUtc(v time.Time)`

SetCreatedAtUtc sets CreatedAtUtc field to given value.

### HasCreatedAtUtc

`func (o *GetConsentTemplateDetailsDto) HasCreatedAtUtc() bool`

HasCreatedAtUtc returns a boolean if a field has been set.

### GetStatus

`func (o *GetConsentTemplateDetailsDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetConsentTemplateDetailsDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetConsentTemplateDetailsDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetConsentTemplateDetailsDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GetConsentTemplateDetailsDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GetConsentTemplateDetailsDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTemplateType

`func (o *GetConsentTemplateDetailsDto) GetTemplateType() ConsentTemplateTypes`

GetTemplateType returns the TemplateType field if non-nil, zero value otherwise.

### GetTemplateTypeOk

`func (o *GetConsentTemplateDetailsDto) GetTemplateTypeOk() (*ConsentTemplateTypes, bool)`

GetTemplateTypeOk returns a tuple with the TemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateType

`func (o *GetConsentTemplateDetailsDto) SetTemplateType(v ConsentTemplateTypes)`

SetTemplateType sets TemplateType field to given value.

### HasTemplateType

`func (o *GetConsentTemplateDetailsDto) HasTemplateType() bool`

HasTemplateType returns a boolean if a field has been set.

### GetDataLife

`func (o *GetConsentTemplateDetailsDto) GetDataLife() Life`

GetDataLife returns the DataLife field if non-nil, zero value otherwise.

### GetDataLifeOk

`func (o *GetConsentTemplateDetailsDto) GetDataLifeOk() (*Life, bool)`

GetDataLifeOk returns a tuple with the DataLife field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLife

`func (o *GetConsentTemplateDetailsDto) SetDataLife(v Life)`

SetDataLife sets DataLife field to given value.

### HasDataLife

`func (o *GetConsentTemplateDetailsDto) HasDataLife() bool`

HasDataLife returns a boolean if a field has been set.

### GetRequestLife

`func (o *GetConsentTemplateDetailsDto) GetRequestLife() Life`

GetRequestLife returns the RequestLife field if non-nil, zero value otherwise.

### GetRequestLifeOk

`func (o *GetConsentTemplateDetailsDto) GetRequestLifeOk() (*Life, bool)`

GetRequestLifeOk returns a tuple with the RequestLife field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestLife

`func (o *GetConsentTemplateDetailsDto) SetRequestLife(v Life)`

SetRequestLife sets RequestLife field to given value.

### HasRequestLife

`func (o *GetConsentTemplateDetailsDto) HasRequestLife() bool`

HasRequestLife returns a boolean if a field has been set.

### GetFrequency

`func (o *GetConsentTemplateDetailsDto) GetFrequency() Life`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *GetConsentTemplateDetailsDto) GetFrequencyOk() (*Life, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *GetConsentTemplateDetailsDto) SetFrequency(v Life)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *GetConsentTemplateDetailsDto) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetIdentity

`func (o *GetConsentTemplateDetailsDto) GetIdentity() []IdentitySupportedFields`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *GetConsentTemplateDetailsDto) GetIdentityOk() (*[]IdentitySupportedFields, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *GetConsentTemplateDetailsDto) SetIdentity(v []IdentitySupportedFields)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *GetConsentTemplateDetailsDto) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### SetIdentityNil

`func (o *GetConsentTemplateDetailsDto) SetIdentityNil(b bool)`

 SetIdentityNil sets the value for Identity to be an explicit nil

### UnsetIdentity
`func (o *GetConsentTemplateDetailsDto) UnsetIdentity()`

UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
### GetDocuments

`func (o *GetConsentTemplateDetailsDto) GetDocuments() []Document`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *GetConsentTemplateDetailsDto) GetDocumentsOk() (*[]Document, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *GetConsentTemplateDetailsDto) SetDocuments(v []Document)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *GetConsentTemplateDetailsDto) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### SetDocumentsNil

`func (o *GetConsentTemplateDetailsDto) SetDocumentsNil(b bool)`

 SetDocumentsNil sets the value for Documents to be an explicit nil

### UnsetDocuments
`func (o *GetConsentTemplateDetailsDto) UnsetDocuments()`

UnsetDocuments ensures that no value is present for Documents, not even an explicit nil
### GetFinancials

`func (o *GetConsentTemplateDetailsDto) GetFinancials() []Financial`

GetFinancials returns the Financials field if non-nil, zero value otherwise.

### GetFinancialsOk

`func (o *GetConsentTemplateDetailsDto) GetFinancialsOk() (*[]Financial, bool)`

GetFinancialsOk returns a tuple with the Financials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancials

`func (o *GetConsentTemplateDetailsDto) SetFinancials(v []Financial)`

SetFinancials sets Financials field to given value.

### HasFinancials

`func (o *GetConsentTemplateDetailsDto) HasFinancials() bool`

HasFinancials returns a boolean if a field has been set.

### SetFinancialsNil

`func (o *GetConsentTemplateDetailsDto) SetFinancialsNil(b bool)`

 SetFinancialsNil sets the value for Financials to be an explicit nil

### UnsetFinancials
`func (o *GetConsentTemplateDetailsDto) UnsetFinancials()`

UnsetFinancials ensures that no value is present for Financials, not even an explicit nil
### GetHealthRecords

`func (o *GetConsentTemplateDetailsDto) GetHealthRecords() []map[string]interface{}`

GetHealthRecords returns the HealthRecords field if non-nil, zero value otherwise.

### GetHealthRecordsOk

`func (o *GetConsentTemplateDetailsDto) GetHealthRecordsOk() (*[]map[string]interface{}, bool)`

GetHealthRecordsOk returns a tuple with the HealthRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthRecords

`func (o *GetConsentTemplateDetailsDto) SetHealthRecords(v []map[string]interface{})`

SetHealthRecords sets HealthRecords field to given value.

### HasHealthRecords

`func (o *GetConsentTemplateDetailsDto) HasHealthRecords() bool`

HasHealthRecords returns a boolean if a field has been set.

### SetHealthRecordsNil

`func (o *GetConsentTemplateDetailsDto) SetHealthRecordsNil(b bool)`

 SetHealthRecordsNil sets the value for HealthRecords to be an explicit nil

### UnsetHealthRecords
`func (o *GetConsentTemplateDetailsDto) UnsetHealthRecords()`

UnsetHealthRecords ensures that no value is present for HealthRecords, not even an explicit nil
### GetApprovedBy

`func (o *GetConsentTemplateDetailsDto) GetApprovedBy() string`

GetApprovedBy returns the ApprovedBy field if non-nil, zero value otherwise.

### GetApprovedByOk

`func (o *GetConsentTemplateDetailsDto) GetApprovedByOk() (*string, bool)`

GetApprovedByOk returns a tuple with the ApprovedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedBy

`func (o *GetConsentTemplateDetailsDto) SetApprovedBy(v string)`

SetApprovedBy sets ApprovedBy field to given value.

### HasApprovedBy

`func (o *GetConsentTemplateDetailsDto) HasApprovedBy() bool`

HasApprovedBy returns a boolean if a field has been set.

### SetApprovedByNil

`func (o *GetConsentTemplateDetailsDto) SetApprovedByNil(b bool)`

 SetApprovedByNil sets the value for ApprovedBy to be an explicit nil

### UnsetApprovedBy
`func (o *GetConsentTemplateDetailsDto) UnsetApprovedBy()`

UnsetApprovedBy ensures that no value is present for ApprovedBy, not even an explicit nil
### GetApprovedAtUtc

`func (o *GetConsentTemplateDetailsDto) GetApprovedAtUtc() time.Time`

GetApprovedAtUtc returns the ApprovedAtUtc field if non-nil, zero value otherwise.

### GetApprovedAtUtcOk

`func (o *GetConsentTemplateDetailsDto) GetApprovedAtUtcOk() (*time.Time, bool)`

GetApprovedAtUtcOk returns a tuple with the ApprovedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAtUtc

`func (o *GetConsentTemplateDetailsDto) SetApprovedAtUtc(v time.Time)`

SetApprovedAtUtc sets ApprovedAtUtc field to given value.

### HasApprovedAtUtc

`func (o *GetConsentTemplateDetailsDto) HasApprovedAtUtc() bool`

HasApprovedAtUtc returns a boolean if a field has been set.

### SetApprovedAtUtcNil

`func (o *GetConsentTemplateDetailsDto) SetApprovedAtUtcNil(b bool)`

 SetApprovedAtUtcNil sets the value for ApprovedAtUtc to be an explicit nil

### UnsetApprovedAtUtc
`func (o *GetConsentTemplateDetailsDto) UnsetApprovedAtUtc()`

UnsetApprovedAtUtc ensures that no value is present for ApprovedAtUtc, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



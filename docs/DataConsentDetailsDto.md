# DataConsentDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**DataLife** | Pointer to [**Life**](Life.md) |  | [optional] 
**RequestedByOrg** | Pointer to [**Requester**](Requester.md) |  | [optional] 
**Status** | Pointer to [**DataConsentStatus**](DataConsentStatus.md) |  | [optional] 
**ApprovedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**RejectedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**ExpiresAtUtc** | Pointer to **time.Time** |  | [optional] 
**RequestedAtUtc** | Pointer to **time.Time** |  | [optional] 
**Identifiers** | Pointer to **interface{}** |  | [optional] 
**Documents** | Pointer to [**[]DataConsentDocumentDetailsDto**](DataConsentDocumentDetailsDto.md) |  | [optional] 
**Financials** | Pointer to **NullableString** |  | [optional] 
**HealthRecords** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDataConsentDetailsDto

`func NewDataConsentDetailsDto(id string, ) *DataConsentDetailsDto`

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


### GetTitle

`func (o *DataConsentDetailsDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DataConsentDetailsDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DataConsentDetailsDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DataConsentDetailsDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *DataConsentDetailsDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *DataConsentDetailsDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
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
### GetDataLife

`func (o *DataConsentDetailsDto) GetDataLife() Life`

GetDataLife returns the DataLife field if non-nil, zero value otherwise.

### GetDataLifeOk

`func (o *DataConsentDetailsDto) GetDataLifeOk() (*Life, bool)`

GetDataLifeOk returns a tuple with the DataLife field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLife

`func (o *DataConsentDetailsDto) SetDataLife(v Life)`

SetDataLife sets DataLife field to given value.

### HasDataLife

`func (o *DataConsentDetailsDto) HasDataLife() bool`

HasDataLife returns a boolean if a field has been set.

### GetRequestedByOrg

`func (o *DataConsentDetailsDto) GetRequestedByOrg() Requester`

GetRequestedByOrg returns the RequestedByOrg field if non-nil, zero value otherwise.

### GetRequestedByOrgOk

`func (o *DataConsentDetailsDto) GetRequestedByOrgOk() (*Requester, bool)`

GetRequestedByOrgOk returns a tuple with the RequestedByOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedByOrg

`func (o *DataConsentDetailsDto) SetRequestedByOrg(v Requester)`

SetRequestedByOrg sets RequestedByOrg field to given value.

### HasRequestedByOrg

`func (o *DataConsentDetailsDto) HasRequestedByOrg() bool`

HasRequestedByOrg returns a boolean if a field has been set.

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

### GetIdentifiers

`func (o *DataConsentDetailsDto) GetIdentifiers() interface{}`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *DataConsentDetailsDto) GetIdentifiersOk() (*interface{}, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *DataConsentDetailsDto) SetIdentifiers(v interface{})`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *DataConsentDetailsDto) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### SetIdentifiersNil

`func (o *DataConsentDetailsDto) SetIdentifiersNil(b bool)`

 SetIdentifiersNil sets the value for Identifiers to be an explicit nil

### UnsetIdentifiers
`func (o *DataConsentDetailsDto) UnsetIdentifiers()`

UnsetIdentifiers ensures that no value is present for Identifiers, not even an explicit nil
### GetDocuments

`func (o *DataConsentDetailsDto) GetDocuments() []DataConsentDocumentDetailsDto`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *DataConsentDetailsDto) GetDocumentsOk() (*[]DataConsentDocumentDetailsDto, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *DataConsentDetailsDto) SetDocuments(v []DataConsentDocumentDetailsDto)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *DataConsentDetailsDto) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### SetDocumentsNil

`func (o *DataConsentDetailsDto) SetDocumentsNil(b bool)`

 SetDocumentsNil sets the value for Documents to be an explicit nil

### UnsetDocuments
`func (o *DataConsentDetailsDto) UnsetDocuments()`

UnsetDocuments ensures that no value is present for Documents, not even an explicit nil
### GetFinancials

`func (o *DataConsentDetailsDto) GetFinancials() string`

GetFinancials returns the Financials field if non-nil, zero value otherwise.

### GetFinancialsOk

`func (o *DataConsentDetailsDto) GetFinancialsOk() (*string, bool)`

GetFinancialsOk returns a tuple with the Financials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancials

`func (o *DataConsentDetailsDto) SetFinancials(v string)`

SetFinancials sets Financials field to given value.

### HasFinancials

`func (o *DataConsentDetailsDto) HasFinancials() bool`

HasFinancials returns a boolean if a field has been set.

### SetFinancialsNil

`func (o *DataConsentDetailsDto) SetFinancialsNil(b bool)`

 SetFinancialsNil sets the value for Financials to be an explicit nil

### UnsetFinancials
`func (o *DataConsentDetailsDto) UnsetFinancials()`

UnsetFinancials ensures that no value is present for Financials, not even an explicit nil
### GetHealthRecords

`func (o *DataConsentDetailsDto) GetHealthRecords() string`

GetHealthRecords returns the HealthRecords field if non-nil, zero value otherwise.

### GetHealthRecordsOk

`func (o *DataConsentDetailsDto) GetHealthRecordsOk() (*string, bool)`

GetHealthRecordsOk returns a tuple with the HealthRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthRecords

`func (o *DataConsentDetailsDto) SetHealthRecords(v string)`

SetHealthRecords sets HealthRecords field to given value.

### HasHealthRecords

`func (o *DataConsentDetailsDto) HasHealthRecords() bool`

HasHealthRecords returns a boolean if a field has been set.

### SetHealthRecordsNil

`func (o *DataConsentDetailsDto) SetHealthRecordsNil(b bool)`

 SetHealthRecordsNil sets the value for HealthRecords to be an explicit nil

### UnsetHealthRecords
`func (o *DataConsentDetailsDto) UnsetHealthRecords()`

UnsetHealthRecords ensures that no value is present for HealthRecords, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DataConsent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Data consent id. | 
**RequestId** | **string** | Data consent request id. | 
**TemplateId** | Pointer to **NullableString** | Consent template id. | [optional] 
**Title** | **string** | Consent title. | 
**Description** | **string** | Consent description. | 
**Purpose** | Pointer to **NullableString** | Consent purpose. | [optional] 
**Status** | [**DataConsentStatus**](DataConsentStatus.md) |  | 
**TransactionId** | Pointer to **NullableString** | Transaction id. | [optional] 
**ApprovedAtUtc** | **time.Time** | Consent approval datetime in UTC timezone. | 
**DataAccessExpiresAtUtc** | **time.Time** | Data access expiration datetime in UTC timezone. | 
**RevokedAtUtc** | Pointer to **NullableTime** | Consent revocation datetime in UTC timezone. | [optional] 
**Collectables** | [**[]CollectibleTypes**](CollectibleTypes.md) | List of supported collectible types. | 
**Identifiers** | Pointer to **interface{}** | Consented identity details. | [optional] 
**Documents** | Pointer to [**[]DataConsentDocument**](DataConsentDocument.md) | List of consented documents. | [optional] 

## Methods

### NewDataConsent

`func NewDataConsent(id string, requestId string, title string, description string, status DataConsentStatus, approvedAtUtc time.Time, dataAccessExpiresAtUtc time.Time, collectables []CollectibleTypes, ) *DataConsent`

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


### GetRequestId

`func (o *DataConsent) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DataConsent) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DataConsent) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetTemplateId

`func (o *DataConsent) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *DataConsent) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *DataConsent) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *DataConsent) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *DataConsent) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *DataConsent) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetTitle

`func (o *DataConsent) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DataConsent) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DataConsent) SetTitle(v string)`

SetTitle sets Title field to given value.


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


### GetPurpose

`func (o *DataConsent) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *DataConsent) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *DataConsent) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *DataConsent) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### SetPurposeNil

`func (o *DataConsent) SetPurposeNil(b bool)`

 SetPurposeNil sets the value for Purpose to be an explicit nil

### UnsetPurpose
`func (o *DataConsent) UnsetPurpose()`

UnsetPurpose ensures that no value is present for Purpose, not even an explicit nil
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


### GetDataAccessExpiresAtUtc

`func (o *DataConsent) GetDataAccessExpiresAtUtc() time.Time`

GetDataAccessExpiresAtUtc returns the DataAccessExpiresAtUtc field if non-nil, zero value otherwise.

### GetDataAccessExpiresAtUtcOk

`func (o *DataConsent) GetDataAccessExpiresAtUtcOk() (*time.Time, bool)`

GetDataAccessExpiresAtUtcOk returns a tuple with the DataAccessExpiresAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataAccessExpiresAtUtc

`func (o *DataConsent) SetDataAccessExpiresAtUtc(v time.Time)`

SetDataAccessExpiresAtUtc sets DataAccessExpiresAtUtc field to given value.


### GetRevokedAtUtc

`func (o *DataConsent) GetRevokedAtUtc() time.Time`

GetRevokedAtUtc returns the RevokedAtUtc field if non-nil, zero value otherwise.

### GetRevokedAtUtcOk

`func (o *DataConsent) GetRevokedAtUtcOk() (*time.Time, bool)`

GetRevokedAtUtcOk returns a tuple with the RevokedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedAtUtc

`func (o *DataConsent) SetRevokedAtUtc(v time.Time)`

SetRevokedAtUtc sets RevokedAtUtc field to given value.

### HasRevokedAtUtc

`func (o *DataConsent) HasRevokedAtUtc() bool`

HasRevokedAtUtc returns a boolean if a field has been set.

### SetRevokedAtUtcNil

`func (o *DataConsent) SetRevokedAtUtcNil(b bool)`

 SetRevokedAtUtcNil sets the value for RevokedAtUtc to be an explicit nil

### UnsetRevokedAtUtc
`func (o *DataConsent) UnsetRevokedAtUtc()`

UnsetRevokedAtUtc ensures that no value is present for RevokedAtUtc, not even an explicit nil
### GetCollectables

`func (o *DataConsent) GetCollectables() []CollectibleTypes`

GetCollectables returns the Collectables field if non-nil, zero value otherwise.

### GetCollectablesOk

`func (o *DataConsent) GetCollectablesOk() (*[]CollectibleTypes, bool)`

GetCollectablesOk returns a tuple with the Collectables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectables

`func (o *DataConsent) SetCollectables(v []CollectibleTypes)`

SetCollectables sets Collectables field to given value.


### GetIdentifiers

`func (o *DataConsent) GetIdentifiers() interface{}`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *DataConsent) GetIdentifiersOk() (*interface{}, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *DataConsent) SetIdentifiers(v interface{})`

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
### GetDocuments

`func (o *DataConsent) GetDocuments() []DataConsentDocument`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *DataConsent) GetDocumentsOk() (*[]DataConsentDocument, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *DataConsent) SetDocuments(v []DataConsentDocument)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *DataConsent) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### SetDocumentsNil

`func (o *DataConsent) SetDocumentsNil(b bool)`

 SetDocumentsNil sets the value for Documents to be an explicit nil

### UnsetDocuments
`func (o *DataConsent) UnsetDocuments()`

UnsetDocuments ensures that no value is present for Documents, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



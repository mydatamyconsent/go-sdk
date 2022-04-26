# DataConsentRequestDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Consent request id | 
**TemplateId** | Pointer to **NullableString** | Consent request template id | [optional] 
**ConsentId** | Pointer to **NullableString** | Data Consent id | [optional] 
**Title** | **string** | Consent request title. | 
**Description** | **string** | Consent request description. | 
**Purpose** | Pointer to **NullableString** | Consent request purpose. | [optional] 
**Status** | [**DataConsentStatus**](DataConsentStatus.md) |  | 
**TransactionId** | Pointer to **NullableString** | Transaction id | [optional] 
**CreatedAtUtc** | **time.Time** | Request creation datetime in UTC timezone | 
**ExpiresAtUtc** | **time.Time** | Request expiration datetime in UTC timezone | 

## Methods

### NewDataConsentRequestDetails

`func NewDataConsentRequestDetails(id string, title string, description string, status DataConsentStatus, createdAtUtc time.Time, expiresAtUtc time.Time, ) *DataConsentRequestDetails`

NewDataConsentRequestDetails instantiates a new DataConsentRequestDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataConsentRequestDetailsWithDefaults

`func NewDataConsentRequestDetailsWithDefaults() *DataConsentRequestDetails`

NewDataConsentRequestDetailsWithDefaults instantiates a new DataConsentRequestDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataConsentRequestDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataConsentRequestDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataConsentRequestDetails) SetId(v string)`

SetId sets Id field to given value.


### GetTemplateId

`func (o *DataConsentRequestDetails) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *DataConsentRequestDetails) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *DataConsentRequestDetails) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *DataConsentRequestDetails) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *DataConsentRequestDetails) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *DataConsentRequestDetails) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetConsentId

`func (o *DataConsentRequestDetails) GetConsentId() string`

GetConsentId returns the ConsentId field if non-nil, zero value otherwise.

### GetConsentIdOk

`func (o *DataConsentRequestDetails) GetConsentIdOk() (*string, bool)`

GetConsentIdOk returns a tuple with the ConsentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentId

`func (o *DataConsentRequestDetails) SetConsentId(v string)`

SetConsentId sets ConsentId field to given value.

### HasConsentId

`func (o *DataConsentRequestDetails) HasConsentId() bool`

HasConsentId returns a boolean if a field has been set.

### SetConsentIdNil

`func (o *DataConsentRequestDetails) SetConsentIdNil(b bool)`

 SetConsentIdNil sets the value for ConsentId to be an explicit nil

### UnsetConsentId
`func (o *DataConsentRequestDetails) UnsetConsentId()`

UnsetConsentId ensures that no value is present for ConsentId, not even an explicit nil
### GetTitle

`func (o *DataConsentRequestDetails) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DataConsentRequestDetails) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DataConsentRequestDetails) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *DataConsentRequestDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DataConsentRequestDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DataConsentRequestDetails) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPurpose

`func (o *DataConsentRequestDetails) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *DataConsentRequestDetails) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *DataConsentRequestDetails) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *DataConsentRequestDetails) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### SetPurposeNil

`func (o *DataConsentRequestDetails) SetPurposeNil(b bool)`

 SetPurposeNil sets the value for Purpose to be an explicit nil

### UnsetPurpose
`func (o *DataConsentRequestDetails) UnsetPurpose()`

UnsetPurpose ensures that no value is present for Purpose, not even an explicit nil
### GetStatus

`func (o *DataConsentRequestDetails) GetStatus() DataConsentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DataConsentRequestDetails) GetStatusOk() (*DataConsentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DataConsentRequestDetails) SetStatus(v DataConsentStatus)`

SetStatus sets Status field to given value.


### GetTransactionId

`func (o *DataConsentRequestDetails) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *DataConsentRequestDetails) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *DataConsentRequestDetails) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *DataConsentRequestDetails) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### SetTransactionIdNil

`func (o *DataConsentRequestDetails) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *DataConsentRequestDetails) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
### GetCreatedAtUtc

`func (o *DataConsentRequestDetails) GetCreatedAtUtc() time.Time`

GetCreatedAtUtc returns the CreatedAtUtc field if non-nil, zero value otherwise.

### GetCreatedAtUtcOk

`func (o *DataConsentRequestDetails) GetCreatedAtUtcOk() (*time.Time, bool)`

GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtUtc

`func (o *DataConsentRequestDetails) SetCreatedAtUtc(v time.Time)`

SetCreatedAtUtc sets CreatedAtUtc field to given value.


### GetExpiresAtUtc

`func (o *DataConsentRequestDetails) GetExpiresAtUtc() time.Time`

GetExpiresAtUtc returns the ExpiresAtUtc field if non-nil, zero value otherwise.

### GetExpiresAtUtcOk

`func (o *DataConsentRequestDetails) GetExpiresAtUtcOk() (*time.Time, bool)`

GetExpiresAtUtcOk returns a tuple with the ExpiresAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAtUtc

`func (o *DataConsentRequestDetails) SetExpiresAtUtc(v time.Time)`

SetExpiresAtUtc sets ExpiresAtUtc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



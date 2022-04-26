# DataConsentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Data consent id. | 
**RequestId** | **string** | Consent request id. | 
**TemplateId** | Pointer to **NullableString** | Consent template id. | [optional] 
**Title** | **string** | Consent title. | 
**Description** | **string** | Consent description. | 
**Purpose** | Pointer to **NullableString** | Consent purpose. | [optional] 
**Status** | [**DataConsentStatus**](DataConsentStatus.md) |  | 
**TransactionId** | Pointer to **NullableString** | Transaction id. | [optional] 
**RequestedAtUtc** | **time.Time** | Consent requested datetime in UTC timezone. | 
**ApprovedAtUtc** | **time.Time** | Consent approval datetime in UTC timezone. | 
**DataAccessExpiresAtUtc** | **time.Time** | Data access expiration datetime in UTC timezone. | 
**RevokedAtUtc** | Pointer to **NullableTime** | Consent revocation datetime in UTC timezone. | [optional] 

## Methods

### NewDataConsentDetails

`func NewDataConsentDetails(id string, requestId string, title string, description string, status DataConsentStatus, requestedAtUtc time.Time, approvedAtUtc time.Time, dataAccessExpiresAtUtc time.Time, ) *DataConsentDetails`

NewDataConsentDetails instantiates a new DataConsentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataConsentDetailsWithDefaults

`func NewDataConsentDetailsWithDefaults() *DataConsentDetails`

NewDataConsentDetailsWithDefaults instantiates a new DataConsentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataConsentDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataConsentDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataConsentDetails) SetId(v string)`

SetId sets Id field to given value.


### GetRequestId

`func (o *DataConsentDetails) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DataConsentDetails) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DataConsentDetails) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetTemplateId

`func (o *DataConsentDetails) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *DataConsentDetails) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *DataConsentDetails) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *DataConsentDetails) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *DataConsentDetails) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *DataConsentDetails) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetTitle

`func (o *DataConsentDetails) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DataConsentDetails) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DataConsentDetails) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *DataConsentDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DataConsentDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DataConsentDetails) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPurpose

`func (o *DataConsentDetails) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *DataConsentDetails) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *DataConsentDetails) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *DataConsentDetails) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### SetPurposeNil

`func (o *DataConsentDetails) SetPurposeNil(b bool)`

 SetPurposeNil sets the value for Purpose to be an explicit nil

### UnsetPurpose
`func (o *DataConsentDetails) UnsetPurpose()`

UnsetPurpose ensures that no value is present for Purpose, not even an explicit nil
### GetStatus

`func (o *DataConsentDetails) GetStatus() DataConsentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DataConsentDetails) GetStatusOk() (*DataConsentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DataConsentDetails) SetStatus(v DataConsentStatus)`

SetStatus sets Status field to given value.


### GetTransactionId

`func (o *DataConsentDetails) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *DataConsentDetails) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *DataConsentDetails) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *DataConsentDetails) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### SetTransactionIdNil

`func (o *DataConsentDetails) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *DataConsentDetails) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
### GetRequestedAtUtc

`func (o *DataConsentDetails) GetRequestedAtUtc() time.Time`

GetRequestedAtUtc returns the RequestedAtUtc field if non-nil, zero value otherwise.

### GetRequestedAtUtcOk

`func (o *DataConsentDetails) GetRequestedAtUtcOk() (*time.Time, bool)`

GetRequestedAtUtcOk returns a tuple with the RequestedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAtUtc

`func (o *DataConsentDetails) SetRequestedAtUtc(v time.Time)`

SetRequestedAtUtc sets RequestedAtUtc field to given value.


### GetApprovedAtUtc

`func (o *DataConsentDetails) GetApprovedAtUtc() time.Time`

GetApprovedAtUtc returns the ApprovedAtUtc field if non-nil, zero value otherwise.

### GetApprovedAtUtcOk

`func (o *DataConsentDetails) GetApprovedAtUtcOk() (*time.Time, bool)`

GetApprovedAtUtcOk returns a tuple with the ApprovedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAtUtc

`func (o *DataConsentDetails) SetApprovedAtUtc(v time.Time)`

SetApprovedAtUtc sets ApprovedAtUtc field to given value.


### GetDataAccessExpiresAtUtc

`func (o *DataConsentDetails) GetDataAccessExpiresAtUtc() time.Time`

GetDataAccessExpiresAtUtc returns the DataAccessExpiresAtUtc field if non-nil, zero value otherwise.

### GetDataAccessExpiresAtUtcOk

`func (o *DataConsentDetails) GetDataAccessExpiresAtUtcOk() (*time.Time, bool)`

GetDataAccessExpiresAtUtcOk returns a tuple with the DataAccessExpiresAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataAccessExpiresAtUtc

`func (o *DataConsentDetails) SetDataAccessExpiresAtUtc(v time.Time)`

SetDataAccessExpiresAtUtc sets DataAccessExpiresAtUtc field to given value.


### GetRevokedAtUtc

`func (o *DataConsentDetails) GetRevokedAtUtc() time.Time`

GetRevokedAtUtc returns the RevokedAtUtc field if non-nil, zero value otherwise.

### GetRevokedAtUtcOk

`func (o *DataConsentDetails) GetRevokedAtUtcOk() (*time.Time, bool)`

GetRevokedAtUtcOk returns a tuple with the RevokedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedAtUtc

`func (o *DataConsentDetails) SetRevokedAtUtc(v time.Time)`

SetRevokedAtUtc sets RevokedAtUtc field to given value.

### HasRevokedAtUtc

`func (o *DataConsentDetails) HasRevokedAtUtc() bool`

HasRevokedAtUtc returns a boolean if a field has been set.

### SetRevokedAtUtcNil

`func (o *DataConsentDetails) SetRevokedAtUtcNil(b bool)`

 SetRevokedAtUtcNil sets the value for RevokedAtUtc to be an explicit nil

### UnsetRevokedAtUtc
`func (o *DataConsentDetails) UnsetRevokedAtUtc()`

UnsetRevokedAtUtc ensures that no value is present for RevokedAtUtc, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



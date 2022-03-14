# OrganizationDataConsent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approver** | **string** | Name of consent approver organization. | 
**Id** | **string** | Data consent id. | 
**TemplateId** | Pointer to **NullableString** | Consent template id. | [optional] 
**Title** | **string** | Consent title. | 
**Description** | **string** | Consent description. | 
**Purpose** | Pointer to **NullableString** | Consent purpose. | [optional] 
**Status** | [**DataConsentStatus**](DataConsentStatus.md) |  | 
**TransactionId** | Pointer to **NullableString** | Transaction id. | [optional] 
**ApprovedAtUtc** | **time.Time** | Consent approval datetime in UTC timezone. | 
**DataAccessExpiresAtUtc** | **time.Time** | Data access expiration datetime in UTC timezone. | 
**RevokedAtUtc** | Pointer to **NullableTime** | Consent revocation datetime in UTC timezone. | [optional] 

## Methods

### NewOrganizationDataConsent

`func NewOrganizationDataConsent(approver string, id string, title string, description string, status DataConsentStatus, approvedAtUtc time.Time, dataAccessExpiresAtUtc time.Time, ) *OrganizationDataConsent`

NewOrganizationDataConsent instantiates a new OrganizationDataConsent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationDataConsentWithDefaults

`func NewOrganizationDataConsentWithDefaults() *OrganizationDataConsent`

NewOrganizationDataConsentWithDefaults instantiates a new OrganizationDataConsent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprover

`func (o *OrganizationDataConsent) GetApprover() string`

GetApprover returns the Approver field if non-nil, zero value otherwise.

### GetApproverOk

`func (o *OrganizationDataConsent) GetApproverOk() (*string, bool)`

GetApproverOk returns a tuple with the Approver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprover

`func (o *OrganizationDataConsent) SetApprover(v string)`

SetApprover sets Approver field to given value.


### GetId

`func (o *OrganizationDataConsent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationDataConsent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationDataConsent) SetId(v string)`

SetId sets Id field to given value.


### GetTemplateId

`func (o *OrganizationDataConsent) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *OrganizationDataConsent) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *OrganizationDataConsent) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *OrganizationDataConsent) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *OrganizationDataConsent) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *OrganizationDataConsent) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetTitle

`func (o *OrganizationDataConsent) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OrganizationDataConsent) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OrganizationDataConsent) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *OrganizationDataConsent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationDataConsent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationDataConsent) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPurpose

`func (o *OrganizationDataConsent) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *OrganizationDataConsent) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *OrganizationDataConsent) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *OrganizationDataConsent) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### SetPurposeNil

`func (o *OrganizationDataConsent) SetPurposeNil(b bool)`

 SetPurposeNil sets the value for Purpose to be an explicit nil

### UnsetPurpose
`func (o *OrganizationDataConsent) UnsetPurpose()`

UnsetPurpose ensures that no value is present for Purpose, not even an explicit nil
### GetStatus

`func (o *OrganizationDataConsent) GetStatus() DataConsentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrganizationDataConsent) GetStatusOk() (*DataConsentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrganizationDataConsent) SetStatus(v DataConsentStatus)`

SetStatus sets Status field to given value.


### GetTransactionId

`func (o *OrganizationDataConsent) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *OrganizationDataConsent) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *OrganizationDataConsent) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *OrganizationDataConsent) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### SetTransactionIdNil

`func (o *OrganizationDataConsent) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *OrganizationDataConsent) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
### GetApprovedAtUtc

`func (o *OrganizationDataConsent) GetApprovedAtUtc() time.Time`

GetApprovedAtUtc returns the ApprovedAtUtc field if non-nil, zero value otherwise.

### GetApprovedAtUtcOk

`func (o *OrganizationDataConsent) GetApprovedAtUtcOk() (*time.Time, bool)`

GetApprovedAtUtcOk returns a tuple with the ApprovedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAtUtc

`func (o *OrganizationDataConsent) SetApprovedAtUtc(v time.Time)`

SetApprovedAtUtc sets ApprovedAtUtc field to given value.


### GetDataAccessExpiresAtUtc

`func (o *OrganizationDataConsent) GetDataAccessExpiresAtUtc() time.Time`

GetDataAccessExpiresAtUtc returns the DataAccessExpiresAtUtc field if non-nil, zero value otherwise.

### GetDataAccessExpiresAtUtcOk

`func (o *OrganizationDataConsent) GetDataAccessExpiresAtUtcOk() (*time.Time, bool)`

GetDataAccessExpiresAtUtcOk returns a tuple with the DataAccessExpiresAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataAccessExpiresAtUtc

`func (o *OrganizationDataConsent) SetDataAccessExpiresAtUtc(v time.Time)`

SetDataAccessExpiresAtUtc sets DataAccessExpiresAtUtc field to given value.


### GetRevokedAtUtc

`func (o *OrganizationDataConsent) GetRevokedAtUtc() time.Time`

GetRevokedAtUtc returns the RevokedAtUtc field if non-nil, zero value otherwise.

### GetRevokedAtUtcOk

`func (o *OrganizationDataConsent) GetRevokedAtUtcOk() (*time.Time, bool)`

GetRevokedAtUtcOk returns a tuple with the RevokedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedAtUtc

`func (o *OrganizationDataConsent) SetRevokedAtUtc(v time.Time)`

SetRevokedAtUtc sets RevokedAtUtc field to given value.

### HasRevokedAtUtc

`func (o *OrganizationDataConsent) HasRevokedAtUtc() bool`

HasRevokedAtUtc returns a boolean if a field has been set.

### SetRevokedAtUtcNil

`func (o *OrganizationDataConsent) SetRevokedAtUtcNil(b bool)`

 SetRevokedAtUtcNil sets the value for RevokedAtUtc to be an explicit nil

### UnsetRevokedAtUtc
`func (o *OrganizationDataConsent) UnsetRevokedAtUtc()`

UnsetRevokedAtUtc ensures that no value is present for RevokedAtUtc, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



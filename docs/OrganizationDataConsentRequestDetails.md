# OrganizationDataConsentRequestDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Receiver** | **string** | Name of request receiver organization. | 
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

### NewOrganizationDataConsentRequestDetails

`func NewOrganizationDataConsentRequestDetails(receiver string, id string, title string, description string, status DataConsentStatus, createdAtUtc time.Time, expiresAtUtc time.Time, ) *OrganizationDataConsentRequestDetails`

NewOrganizationDataConsentRequestDetails instantiates a new OrganizationDataConsentRequestDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationDataConsentRequestDetailsWithDefaults

`func NewOrganizationDataConsentRequestDetailsWithDefaults() *OrganizationDataConsentRequestDetails`

NewOrganizationDataConsentRequestDetailsWithDefaults instantiates a new OrganizationDataConsentRequestDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceiver

`func (o *OrganizationDataConsentRequestDetails) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *OrganizationDataConsentRequestDetails) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *OrganizationDataConsentRequestDetails) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.


### GetId

`func (o *OrganizationDataConsentRequestDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationDataConsentRequestDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationDataConsentRequestDetails) SetId(v string)`

SetId sets Id field to given value.


### GetTemplateId

`func (o *OrganizationDataConsentRequestDetails) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *OrganizationDataConsentRequestDetails) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *OrganizationDataConsentRequestDetails) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *OrganizationDataConsentRequestDetails) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *OrganizationDataConsentRequestDetails) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *OrganizationDataConsentRequestDetails) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetConsentId

`func (o *OrganizationDataConsentRequestDetails) GetConsentId() string`

GetConsentId returns the ConsentId field if non-nil, zero value otherwise.

### GetConsentIdOk

`func (o *OrganizationDataConsentRequestDetails) GetConsentIdOk() (*string, bool)`

GetConsentIdOk returns a tuple with the ConsentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentId

`func (o *OrganizationDataConsentRequestDetails) SetConsentId(v string)`

SetConsentId sets ConsentId field to given value.

### HasConsentId

`func (o *OrganizationDataConsentRequestDetails) HasConsentId() bool`

HasConsentId returns a boolean if a field has been set.

### SetConsentIdNil

`func (o *OrganizationDataConsentRequestDetails) SetConsentIdNil(b bool)`

 SetConsentIdNil sets the value for ConsentId to be an explicit nil

### UnsetConsentId
`func (o *OrganizationDataConsentRequestDetails) UnsetConsentId()`

UnsetConsentId ensures that no value is present for ConsentId, not even an explicit nil
### GetTitle

`func (o *OrganizationDataConsentRequestDetails) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OrganizationDataConsentRequestDetails) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OrganizationDataConsentRequestDetails) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *OrganizationDataConsentRequestDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationDataConsentRequestDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationDataConsentRequestDetails) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPurpose

`func (o *OrganizationDataConsentRequestDetails) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *OrganizationDataConsentRequestDetails) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *OrganizationDataConsentRequestDetails) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *OrganizationDataConsentRequestDetails) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### SetPurposeNil

`func (o *OrganizationDataConsentRequestDetails) SetPurposeNil(b bool)`

 SetPurposeNil sets the value for Purpose to be an explicit nil

### UnsetPurpose
`func (o *OrganizationDataConsentRequestDetails) UnsetPurpose()`

UnsetPurpose ensures that no value is present for Purpose, not even an explicit nil
### GetStatus

`func (o *OrganizationDataConsentRequestDetails) GetStatus() DataConsentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrganizationDataConsentRequestDetails) GetStatusOk() (*DataConsentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrganizationDataConsentRequestDetails) SetStatus(v DataConsentStatus)`

SetStatus sets Status field to given value.


### GetTransactionId

`func (o *OrganizationDataConsentRequestDetails) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *OrganizationDataConsentRequestDetails) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *OrganizationDataConsentRequestDetails) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *OrganizationDataConsentRequestDetails) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### SetTransactionIdNil

`func (o *OrganizationDataConsentRequestDetails) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *OrganizationDataConsentRequestDetails) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
### GetCreatedAtUtc

`func (o *OrganizationDataConsentRequestDetails) GetCreatedAtUtc() time.Time`

GetCreatedAtUtc returns the CreatedAtUtc field if non-nil, zero value otherwise.

### GetCreatedAtUtcOk

`func (o *OrganizationDataConsentRequestDetails) GetCreatedAtUtcOk() (*time.Time, bool)`

GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtUtc

`func (o *OrganizationDataConsentRequestDetails) SetCreatedAtUtc(v time.Time)`

SetCreatedAtUtc sets CreatedAtUtc field to given value.


### GetExpiresAtUtc

`func (o *OrganizationDataConsentRequestDetails) GetExpiresAtUtc() time.Time`

GetExpiresAtUtc returns the ExpiresAtUtc field if non-nil, zero value otherwise.

### GetExpiresAtUtcOk

`func (o *OrganizationDataConsentRequestDetails) GetExpiresAtUtcOk() (*time.Time, bool)`

GetExpiresAtUtcOk returns a tuple with the ExpiresAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAtUtc

`func (o *OrganizationDataConsentRequestDetails) SetExpiresAtUtc(v time.Time)`

SetExpiresAtUtc sets ExpiresAtUtc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# OrganizationConsentRequestDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Receiver** | **string** | Name of request receiver organization. | 
**Id** | **string** | Consent request id | 
**TemplateId** | Pointer to **string** | Consent request template id | [optional] 
**ConsentId** | Pointer to **string** | Consent id | [optional] 
**Title** | **string** | Consent request title. | 
**Description** | **string** | Consent request description. | 
**Purpose** | Pointer to **string** | Consent request purpose. | [optional] 
**Status** | [**DataConsentStatus**](DataConsentStatus.md) |  | 
**TransactionId** | Pointer to **string** | Transaction id | [optional] 
**CreatedAtUtc** | **time.Time** | Request creation datetime in UTC timezone | 
**ExpiresAtUtc** | **time.Time** | Request expiration datetime in UTC timezone | 

## Methods

### NewOrganizationConsentRequestDetails

`func NewOrganizationConsentRequestDetails(receiver string, id string, title string, description string, status DataConsentStatus, createdAtUtc time.Time, expiresAtUtc time.Time, ) *OrganizationConsentRequestDetails`

NewOrganizationConsentRequestDetails instantiates a new OrganizationConsentRequestDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationConsentRequestDetailsWithDefaults

`func NewOrganizationConsentRequestDetailsWithDefaults() *OrganizationConsentRequestDetails`

NewOrganizationConsentRequestDetailsWithDefaults instantiates a new OrganizationConsentRequestDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceiver

`func (o *OrganizationConsentRequestDetails) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *OrganizationConsentRequestDetails) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *OrganizationConsentRequestDetails) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.


### GetId

`func (o *OrganizationConsentRequestDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationConsentRequestDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationConsentRequestDetails) SetId(v string)`

SetId sets Id field to given value.


### GetTemplateId

`func (o *OrganizationConsentRequestDetails) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *OrganizationConsentRequestDetails) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *OrganizationConsentRequestDetails) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *OrganizationConsentRequestDetails) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetConsentId

`func (o *OrganizationConsentRequestDetails) GetConsentId() string`

GetConsentId returns the ConsentId field if non-nil, zero value otherwise.

### GetConsentIdOk

`func (o *OrganizationConsentRequestDetails) GetConsentIdOk() (*string, bool)`

GetConsentIdOk returns a tuple with the ConsentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentId

`func (o *OrganizationConsentRequestDetails) SetConsentId(v string)`

SetConsentId sets ConsentId field to given value.

### HasConsentId

`func (o *OrganizationConsentRequestDetails) HasConsentId() bool`

HasConsentId returns a boolean if a field has been set.

### GetTitle

`func (o *OrganizationConsentRequestDetails) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OrganizationConsentRequestDetails) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OrganizationConsentRequestDetails) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *OrganizationConsentRequestDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationConsentRequestDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationConsentRequestDetails) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPurpose

`func (o *OrganizationConsentRequestDetails) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *OrganizationConsentRequestDetails) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *OrganizationConsentRequestDetails) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *OrganizationConsentRequestDetails) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetStatus

`func (o *OrganizationConsentRequestDetails) GetStatus() DataConsentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrganizationConsentRequestDetails) GetStatusOk() (*DataConsentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrganizationConsentRequestDetails) SetStatus(v DataConsentStatus)`

SetStatus sets Status field to given value.


### GetTransactionId

`func (o *OrganizationConsentRequestDetails) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *OrganizationConsentRequestDetails) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *OrganizationConsentRequestDetails) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *OrganizationConsentRequestDetails) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetCreatedAtUtc

`func (o *OrganizationConsentRequestDetails) GetCreatedAtUtc() time.Time`

GetCreatedAtUtc returns the CreatedAtUtc field if non-nil, zero value otherwise.

### GetCreatedAtUtcOk

`func (o *OrganizationConsentRequestDetails) GetCreatedAtUtcOk() (*time.Time, bool)`

GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtUtc

`func (o *OrganizationConsentRequestDetails) SetCreatedAtUtc(v time.Time)`

SetCreatedAtUtc sets CreatedAtUtc field to given value.


### GetExpiresAtUtc

`func (o *OrganizationConsentRequestDetails) GetExpiresAtUtc() time.Time`

GetExpiresAtUtc returns the ExpiresAtUtc field if non-nil, zero value otherwise.

### GetExpiresAtUtcOk

`func (o *OrganizationConsentRequestDetails) GetExpiresAtUtcOk() (*time.Time, bool)`

GetExpiresAtUtcOk returns a tuple with the ExpiresAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAtUtc

`func (o *OrganizationConsentRequestDetails) SetExpiresAtUtc(v time.Time)`

SetExpiresAtUtc sets ExpiresAtUtc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



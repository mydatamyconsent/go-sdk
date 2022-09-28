# ConsentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Consent request id. | 
**TemplateId** | Pointer to **string** | Consent request template id. | [optional] 
**ConsentId** | Pointer to **string** | Consent id. | [optional] 
**Title** | **string** | Consent title. | 
**Description** | **string** | Consent description. | 
**Purpose** | Pointer to **string** | Consent purpose. | [optional] 
**DataLife** | Pointer to [**Life**](Life.md) |  | [optional] 
**Collectables** | [**[]CollectibleTypes**](CollectibleTypes.md) | List of supported collectables. | 
**Receiver** | [**ConsentRequestReceiver**](ConsentRequestReceiver.md) |  | 
**Status** | [**DataConsentStatus**](DataConsentStatus.md) |  | 
**CreatedAtUtc** | **time.Time** | Request creation datetime in UTC timezone. | 
**ExpiresAtUtc** | **time.Time** | Request expiration datetime in UTC timezone. | 
**ApprovedAtUtc** | Pointer to **time.Time** | Request approval datetime in UTC timezone. | [optional] 
**DataAccessExpiresAtUtc** | Pointer to **time.Time** | Data access expiration datetime in UTC timezone. | [optional] 
**RejectedAtUtc** | Pointer to **time.Time** | Request rejection datetime in UTC timezone. | [optional] 
**RevokedAtUtc** | Pointer to **time.Time** | Request revocation datetime in UTC timezone. | [optional] 

## Methods

### NewConsentRequest

`func NewConsentRequest(id string, title string, description string, collectables []CollectibleTypes, receiver ConsentRequestReceiver, status DataConsentStatus, createdAtUtc time.Time, expiresAtUtc time.Time, ) *ConsentRequest`

NewConsentRequest instantiates a new ConsentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsentRequestWithDefaults

`func NewConsentRequestWithDefaults() *ConsentRequest`

NewConsentRequestWithDefaults instantiates a new ConsentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConsentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetTemplateId

`func (o *ConsentRequest) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *ConsentRequest) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *ConsentRequest) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *ConsentRequest) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetConsentId

`func (o *ConsentRequest) GetConsentId() string`

GetConsentId returns the ConsentId field if non-nil, zero value otherwise.

### GetConsentIdOk

`func (o *ConsentRequest) GetConsentIdOk() (*string, bool)`

GetConsentIdOk returns a tuple with the ConsentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentId

`func (o *ConsentRequest) SetConsentId(v string)`

SetConsentId sets ConsentId field to given value.

### HasConsentId

`func (o *ConsentRequest) HasConsentId() bool`

HasConsentId returns a boolean if a field has been set.

### GetTitle

`func (o *ConsentRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ConsentRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ConsentRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *ConsentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConsentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConsentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPurpose

`func (o *ConsentRequest) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *ConsentRequest) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *ConsentRequest) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *ConsentRequest) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetDataLife

`func (o *ConsentRequest) GetDataLife() Life`

GetDataLife returns the DataLife field if non-nil, zero value otherwise.

### GetDataLifeOk

`func (o *ConsentRequest) GetDataLifeOk() (*Life, bool)`

GetDataLifeOk returns a tuple with the DataLife field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLife

`func (o *ConsentRequest) SetDataLife(v Life)`

SetDataLife sets DataLife field to given value.

### HasDataLife

`func (o *ConsentRequest) HasDataLife() bool`

HasDataLife returns a boolean if a field has been set.

### GetCollectables

`func (o *ConsentRequest) GetCollectables() []CollectibleTypes`

GetCollectables returns the Collectables field if non-nil, zero value otherwise.

### GetCollectablesOk

`func (o *ConsentRequest) GetCollectablesOk() (*[]CollectibleTypes, bool)`

GetCollectablesOk returns a tuple with the Collectables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectables

`func (o *ConsentRequest) SetCollectables(v []CollectibleTypes)`

SetCollectables sets Collectables field to given value.


### GetReceiver

`func (o *ConsentRequest) GetReceiver() ConsentRequestReceiver`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *ConsentRequest) GetReceiverOk() (*ConsentRequestReceiver, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *ConsentRequest) SetReceiver(v ConsentRequestReceiver)`

SetReceiver sets Receiver field to given value.


### GetStatus

`func (o *ConsentRequest) GetStatus() DataConsentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConsentRequest) GetStatusOk() (*DataConsentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConsentRequest) SetStatus(v DataConsentStatus)`

SetStatus sets Status field to given value.


### GetCreatedAtUtc

`func (o *ConsentRequest) GetCreatedAtUtc() time.Time`

GetCreatedAtUtc returns the CreatedAtUtc field if non-nil, zero value otherwise.

### GetCreatedAtUtcOk

`func (o *ConsentRequest) GetCreatedAtUtcOk() (*time.Time, bool)`

GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtUtc

`func (o *ConsentRequest) SetCreatedAtUtc(v time.Time)`

SetCreatedAtUtc sets CreatedAtUtc field to given value.


### GetExpiresAtUtc

`func (o *ConsentRequest) GetExpiresAtUtc() time.Time`

GetExpiresAtUtc returns the ExpiresAtUtc field if non-nil, zero value otherwise.

### GetExpiresAtUtcOk

`func (o *ConsentRequest) GetExpiresAtUtcOk() (*time.Time, bool)`

GetExpiresAtUtcOk returns a tuple with the ExpiresAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAtUtc

`func (o *ConsentRequest) SetExpiresAtUtc(v time.Time)`

SetExpiresAtUtc sets ExpiresAtUtc field to given value.


### GetApprovedAtUtc

`func (o *ConsentRequest) GetApprovedAtUtc() time.Time`

GetApprovedAtUtc returns the ApprovedAtUtc field if non-nil, zero value otherwise.

### GetApprovedAtUtcOk

`func (o *ConsentRequest) GetApprovedAtUtcOk() (*time.Time, bool)`

GetApprovedAtUtcOk returns a tuple with the ApprovedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAtUtc

`func (o *ConsentRequest) SetApprovedAtUtc(v time.Time)`

SetApprovedAtUtc sets ApprovedAtUtc field to given value.

### HasApprovedAtUtc

`func (o *ConsentRequest) HasApprovedAtUtc() bool`

HasApprovedAtUtc returns a boolean if a field has been set.

### GetDataAccessExpiresAtUtc

`func (o *ConsentRequest) GetDataAccessExpiresAtUtc() time.Time`

GetDataAccessExpiresAtUtc returns the DataAccessExpiresAtUtc field if non-nil, zero value otherwise.

### GetDataAccessExpiresAtUtcOk

`func (o *ConsentRequest) GetDataAccessExpiresAtUtcOk() (*time.Time, bool)`

GetDataAccessExpiresAtUtcOk returns a tuple with the DataAccessExpiresAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataAccessExpiresAtUtc

`func (o *ConsentRequest) SetDataAccessExpiresAtUtc(v time.Time)`

SetDataAccessExpiresAtUtc sets DataAccessExpiresAtUtc field to given value.

### HasDataAccessExpiresAtUtc

`func (o *ConsentRequest) HasDataAccessExpiresAtUtc() bool`

HasDataAccessExpiresAtUtc returns a boolean if a field has been set.

### GetRejectedAtUtc

`func (o *ConsentRequest) GetRejectedAtUtc() time.Time`

GetRejectedAtUtc returns the RejectedAtUtc field if non-nil, zero value otherwise.

### GetRejectedAtUtcOk

`func (o *ConsentRequest) GetRejectedAtUtcOk() (*time.Time, bool)`

GetRejectedAtUtcOk returns a tuple with the RejectedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedAtUtc

`func (o *ConsentRequest) SetRejectedAtUtc(v time.Time)`

SetRejectedAtUtc sets RejectedAtUtc field to given value.

### HasRejectedAtUtc

`func (o *ConsentRequest) HasRejectedAtUtc() bool`

HasRejectedAtUtc returns a boolean if a field has been set.

### GetRevokedAtUtc

`func (o *ConsentRequest) GetRevokedAtUtc() time.Time`

GetRevokedAtUtc returns the RevokedAtUtc field if non-nil, zero value otherwise.

### GetRevokedAtUtcOk

`func (o *ConsentRequest) GetRevokedAtUtcOk() (*time.Time, bool)`

GetRevokedAtUtcOk returns a tuple with the RevokedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedAtUtc

`func (o *ConsentRequest) SetRevokedAtUtc(v time.Time)`

SetRevokedAtUtc sets RevokedAtUtc field to given value.

### HasRevokedAtUtc

`func (o *ConsentRequest) HasRevokedAtUtc() bool`

HasRevokedAtUtc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



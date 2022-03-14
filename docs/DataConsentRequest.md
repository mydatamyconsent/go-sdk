# DataConsentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Data consent request id. | 
**TemplateId** | Pointer to **NullableString** | Data consent template id. | [optional] 
**Title** | **string** | Data consent title. | 
**Description** | **string** | Data consent description. | 
**Purpose** | Pointer to **NullableString** | Data consent purpose. | [optional] 
**DataLife** | Pointer to [**Life**](Life.md) |  | [optional] 
**Collectables** | [**[]CollectibleTypes**](CollectibleTypes.md) | List of supported collectables. | 
**Status** | [**DataConsentStatus**](DataConsentStatus.md) |  | 
**CreatedAtUtc** | **time.Time** | Request creation datetime in UTC timezone. | 
**ExpiresAtUtc** | **time.Time** | Request expiration datetime in UTC timezone. | 
**ApprovedAtUtc** | Pointer to **NullableTime** | Request approval datetime in UTC timezone. | [optional] 
**DataAccessExpiresAtUtc** | Pointer to **NullableTime** | Data access expiration datetime in UTC timezone. | [optional] 
**RejectedAtUtc** | Pointer to **NullableTime** | Request rejection datetime in UTC timezone. | [optional] 
**RevokedAtUtc** | Pointer to **NullableTime** | Request revocation datetime in UTC timezone. | [optional] 

## Methods

### NewDataConsentRequest

`func NewDataConsentRequest(id string, title string, description string, collectables []CollectibleTypes, status DataConsentStatus, createdAtUtc time.Time, expiresAtUtc time.Time, ) *DataConsentRequest`

NewDataConsentRequest instantiates a new DataConsentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataConsentRequestWithDefaults

`func NewDataConsentRequestWithDefaults() *DataConsentRequest`

NewDataConsentRequestWithDefaults instantiates a new DataConsentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataConsentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataConsentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataConsentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetTemplateId

`func (o *DataConsentRequest) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *DataConsentRequest) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *DataConsentRequest) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *DataConsentRequest) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *DataConsentRequest) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *DataConsentRequest) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetTitle

`func (o *DataConsentRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DataConsentRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DataConsentRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *DataConsentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DataConsentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DataConsentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPurpose

`func (o *DataConsentRequest) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *DataConsentRequest) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *DataConsentRequest) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *DataConsentRequest) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### SetPurposeNil

`func (o *DataConsentRequest) SetPurposeNil(b bool)`

 SetPurposeNil sets the value for Purpose to be an explicit nil

### UnsetPurpose
`func (o *DataConsentRequest) UnsetPurpose()`

UnsetPurpose ensures that no value is present for Purpose, not even an explicit nil
### GetDataLife

`func (o *DataConsentRequest) GetDataLife() Life`

GetDataLife returns the DataLife field if non-nil, zero value otherwise.

### GetDataLifeOk

`func (o *DataConsentRequest) GetDataLifeOk() (*Life, bool)`

GetDataLifeOk returns a tuple with the DataLife field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLife

`func (o *DataConsentRequest) SetDataLife(v Life)`

SetDataLife sets DataLife field to given value.

### HasDataLife

`func (o *DataConsentRequest) HasDataLife() bool`

HasDataLife returns a boolean if a field has been set.

### GetCollectables

`func (o *DataConsentRequest) GetCollectables() []CollectibleTypes`

GetCollectables returns the Collectables field if non-nil, zero value otherwise.

### GetCollectablesOk

`func (o *DataConsentRequest) GetCollectablesOk() (*[]CollectibleTypes, bool)`

GetCollectablesOk returns a tuple with the Collectables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectables

`func (o *DataConsentRequest) SetCollectables(v []CollectibleTypes)`

SetCollectables sets Collectables field to given value.


### GetStatus

`func (o *DataConsentRequest) GetStatus() DataConsentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DataConsentRequest) GetStatusOk() (*DataConsentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DataConsentRequest) SetStatus(v DataConsentStatus)`

SetStatus sets Status field to given value.


### GetCreatedAtUtc

`func (o *DataConsentRequest) GetCreatedAtUtc() time.Time`

GetCreatedAtUtc returns the CreatedAtUtc field if non-nil, zero value otherwise.

### GetCreatedAtUtcOk

`func (o *DataConsentRequest) GetCreatedAtUtcOk() (*time.Time, bool)`

GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtUtc

`func (o *DataConsentRequest) SetCreatedAtUtc(v time.Time)`

SetCreatedAtUtc sets CreatedAtUtc field to given value.


### GetExpiresAtUtc

`func (o *DataConsentRequest) GetExpiresAtUtc() time.Time`

GetExpiresAtUtc returns the ExpiresAtUtc field if non-nil, zero value otherwise.

### GetExpiresAtUtcOk

`func (o *DataConsentRequest) GetExpiresAtUtcOk() (*time.Time, bool)`

GetExpiresAtUtcOk returns a tuple with the ExpiresAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAtUtc

`func (o *DataConsentRequest) SetExpiresAtUtc(v time.Time)`

SetExpiresAtUtc sets ExpiresAtUtc field to given value.


### GetApprovedAtUtc

`func (o *DataConsentRequest) GetApprovedAtUtc() time.Time`

GetApprovedAtUtc returns the ApprovedAtUtc field if non-nil, zero value otherwise.

### GetApprovedAtUtcOk

`func (o *DataConsentRequest) GetApprovedAtUtcOk() (*time.Time, bool)`

GetApprovedAtUtcOk returns a tuple with the ApprovedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAtUtc

`func (o *DataConsentRequest) SetApprovedAtUtc(v time.Time)`

SetApprovedAtUtc sets ApprovedAtUtc field to given value.

### HasApprovedAtUtc

`func (o *DataConsentRequest) HasApprovedAtUtc() bool`

HasApprovedAtUtc returns a boolean if a field has been set.

### SetApprovedAtUtcNil

`func (o *DataConsentRequest) SetApprovedAtUtcNil(b bool)`

 SetApprovedAtUtcNil sets the value for ApprovedAtUtc to be an explicit nil

### UnsetApprovedAtUtc
`func (o *DataConsentRequest) UnsetApprovedAtUtc()`

UnsetApprovedAtUtc ensures that no value is present for ApprovedAtUtc, not even an explicit nil
### GetDataAccessExpiresAtUtc

`func (o *DataConsentRequest) GetDataAccessExpiresAtUtc() time.Time`

GetDataAccessExpiresAtUtc returns the DataAccessExpiresAtUtc field if non-nil, zero value otherwise.

### GetDataAccessExpiresAtUtcOk

`func (o *DataConsentRequest) GetDataAccessExpiresAtUtcOk() (*time.Time, bool)`

GetDataAccessExpiresAtUtcOk returns a tuple with the DataAccessExpiresAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataAccessExpiresAtUtc

`func (o *DataConsentRequest) SetDataAccessExpiresAtUtc(v time.Time)`

SetDataAccessExpiresAtUtc sets DataAccessExpiresAtUtc field to given value.

### HasDataAccessExpiresAtUtc

`func (o *DataConsentRequest) HasDataAccessExpiresAtUtc() bool`

HasDataAccessExpiresAtUtc returns a boolean if a field has been set.

### SetDataAccessExpiresAtUtcNil

`func (o *DataConsentRequest) SetDataAccessExpiresAtUtcNil(b bool)`

 SetDataAccessExpiresAtUtcNil sets the value for DataAccessExpiresAtUtc to be an explicit nil

### UnsetDataAccessExpiresAtUtc
`func (o *DataConsentRequest) UnsetDataAccessExpiresAtUtc()`

UnsetDataAccessExpiresAtUtc ensures that no value is present for DataAccessExpiresAtUtc, not even an explicit nil
### GetRejectedAtUtc

`func (o *DataConsentRequest) GetRejectedAtUtc() time.Time`

GetRejectedAtUtc returns the RejectedAtUtc field if non-nil, zero value otherwise.

### GetRejectedAtUtcOk

`func (o *DataConsentRequest) GetRejectedAtUtcOk() (*time.Time, bool)`

GetRejectedAtUtcOk returns a tuple with the RejectedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedAtUtc

`func (o *DataConsentRequest) SetRejectedAtUtc(v time.Time)`

SetRejectedAtUtc sets RejectedAtUtc field to given value.

### HasRejectedAtUtc

`func (o *DataConsentRequest) HasRejectedAtUtc() bool`

HasRejectedAtUtc returns a boolean if a field has been set.

### SetRejectedAtUtcNil

`func (o *DataConsentRequest) SetRejectedAtUtcNil(b bool)`

 SetRejectedAtUtcNil sets the value for RejectedAtUtc to be an explicit nil

### UnsetRejectedAtUtc
`func (o *DataConsentRequest) UnsetRejectedAtUtc()`

UnsetRejectedAtUtc ensures that no value is present for RejectedAtUtc, not even an explicit nil
### GetRevokedAtUtc

`func (o *DataConsentRequest) GetRevokedAtUtc() time.Time`

GetRevokedAtUtc returns the RevokedAtUtc field if non-nil, zero value otherwise.

### GetRevokedAtUtcOk

`func (o *DataConsentRequest) GetRevokedAtUtcOk() (*time.Time, bool)`

GetRevokedAtUtcOk returns a tuple with the RevokedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedAtUtc

`func (o *DataConsentRequest) SetRevokedAtUtc(v time.Time)`

SetRevokedAtUtc sets RevokedAtUtc field to given value.

### HasRevokedAtUtc

`func (o *DataConsentRequest) HasRevokedAtUtc() bool`

HasRevokedAtUtc returns a boolean if a field has been set.

### SetRevokedAtUtcNil

`func (o *DataConsentRequest) SetRevokedAtUtcNil(b bool)`

 SetRevokedAtUtcNil sets the value for RevokedAtUtc to be an explicit nil

### UnsetRevokedAtUtc
`func (o *DataConsentRequest) UnsetRevokedAtUtc()`

UnsetRevokedAtUtc ensures that no value is present for RevokedAtUtc, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



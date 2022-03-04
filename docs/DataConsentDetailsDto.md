# DataConsentDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsentRequestId** | **string** |  | 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**DataLife** | Pointer to [**Life**](Life.md) |  | [optional] 
**RequestedByOrg** | Pointer to [**Requester**](Requester.md) |  | [optional] 
**Status** | Pointer to [**DataConsentStatus**](DataConsentStatus.md) |  | [optional] 
**ApprovedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**ApprovedExpiresAtUtc** | Pointer to **NullableTime** |  | [optional] 
**RejectedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**RevokedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**RequestedExpiresAtUtc** | Pointer to **time.Time** |  | [optional] 
**RequestedAtUtc** | Pointer to **time.Time** |  | [optional] 
**Identifiers** | Pointer to **interface{}** |  | [optional] 
**Documents** | Pointer to [**[]DataConsentDocumentDetailsDto**](DataConsentDocumentDetailsDto.md) |  | [optional] 

## Methods

### NewDataConsentDetailsDto

`func NewDataConsentDetailsDto(consentRequestId string, ) *DataConsentDetailsDto`

NewDataConsentDetailsDto instantiates a new DataConsentDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataConsentDetailsDtoWithDefaults

`func NewDataConsentDetailsDtoWithDefaults() *DataConsentDetailsDto`

NewDataConsentDetailsDtoWithDefaults instantiates a new DataConsentDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsentRequestId

`func (o *DataConsentDetailsDto) GetConsentRequestId() string`

GetConsentRequestId returns the ConsentRequestId field if non-nil, zero value otherwise.

### GetConsentRequestIdOk

`func (o *DataConsentDetailsDto) GetConsentRequestIdOk() (*string, bool)`

GetConsentRequestIdOk returns a tuple with the ConsentRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentRequestId

`func (o *DataConsentDetailsDto) SetConsentRequestId(v string)`

SetConsentRequestId sets ConsentRequestId field to given value.


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
### GetApprovedExpiresAtUtc

`func (o *DataConsentDetailsDto) GetApprovedExpiresAtUtc() time.Time`

GetApprovedExpiresAtUtc returns the ApprovedExpiresAtUtc field if non-nil, zero value otherwise.

### GetApprovedExpiresAtUtcOk

`func (o *DataConsentDetailsDto) GetApprovedExpiresAtUtcOk() (*time.Time, bool)`

GetApprovedExpiresAtUtcOk returns a tuple with the ApprovedExpiresAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedExpiresAtUtc

`func (o *DataConsentDetailsDto) SetApprovedExpiresAtUtc(v time.Time)`

SetApprovedExpiresAtUtc sets ApprovedExpiresAtUtc field to given value.

### HasApprovedExpiresAtUtc

`func (o *DataConsentDetailsDto) HasApprovedExpiresAtUtc() bool`

HasApprovedExpiresAtUtc returns a boolean if a field has been set.

### SetApprovedExpiresAtUtcNil

`func (o *DataConsentDetailsDto) SetApprovedExpiresAtUtcNil(b bool)`

 SetApprovedExpiresAtUtcNil sets the value for ApprovedExpiresAtUtc to be an explicit nil

### UnsetApprovedExpiresAtUtc
`func (o *DataConsentDetailsDto) UnsetApprovedExpiresAtUtc()`

UnsetApprovedExpiresAtUtc ensures that no value is present for ApprovedExpiresAtUtc, not even an explicit nil
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
### GetRevokedAtUtc

`func (o *DataConsentDetailsDto) GetRevokedAtUtc() time.Time`

GetRevokedAtUtc returns the RevokedAtUtc field if non-nil, zero value otherwise.

### GetRevokedAtUtcOk

`func (o *DataConsentDetailsDto) GetRevokedAtUtcOk() (*time.Time, bool)`

GetRevokedAtUtcOk returns a tuple with the RevokedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedAtUtc

`func (o *DataConsentDetailsDto) SetRevokedAtUtc(v time.Time)`

SetRevokedAtUtc sets RevokedAtUtc field to given value.

### HasRevokedAtUtc

`func (o *DataConsentDetailsDto) HasRevokedAtUtc() bool`

HasRevokedAtUtc returns a boolean if a field has been set.

### SetRevokedAtUtcNil

`func (o *DataConsentDetailsDto) SetRevokedAtUtcNil(b bool)`

 SetRevokedAtUtcNil sets the value for RevokedAtUtc to be an explicit nil

### UnsetRevokedAtUtc
`func (o *DataConsentDetailsDto) UnsetRevokedAtUtc()`

UnsetRevokedAtUtc ensures that no value is present for RevokedAtUtc, not even an explicit nil
### GetRequestedExpiresAtUtc

`func (o *DataConsentDetailsDto) GetRequestedExpiresAtUtc() time.Time`

GetRequestedExpiresAtUtc returns the RequestedExpiresAtUtc field if non-nil, zero value otherwise.

### GetRequestedExpiresAtUtcOk

`func (o *DataConsentDetailsDto) GetRequestedExpiresAtUtcOk() (*time.Time, bool)`

GetRequestedExpiresAtUtcOk returns a tuple with the RequestedExpiresAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedExpiresAtUtc

`func (o *DataConsentDetailsDto) SetRequestedExpiresAtUtc(v time.Time)`

SetRequestedExpiresAtUtc sets RequestedExpiresAtUtc field to given value.

### HasRequestedExpiresAtUtc

`func (o *DataConsentDetailsDto) HasRequestedExpiresAtUtc() bool`

HasRequestedExpiresAtUtc returns a boolean if a field has been set.

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



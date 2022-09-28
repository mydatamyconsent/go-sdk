# ConsentRequestTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Consent request template id. | 
**Title** | **string** | Consent request template title. | 
**Description** | **string** | Consent request template description. | 
**Purpose** | Pointer to **string** | Consent request template purpose. | [optional] 
**ShortId** | **string** | Consent request template short Id. | 
**Status** | [**ConsentRequestTemplateStatus**](ConsentRequestTemplateStatus.md) |  | 
**CreatedBy** | **string** | Consent request template created by user. | 
**CreatedAtUtc** | **time.Time** | Consent request template created datetime in UTC timezone. | 
**ApprovedAtUtc** | Pointer to **time.Time** | Consent request template approval datetime in UTC timezone. | [optional] 
**RejectedAtUtc** | Pointer to **time.Time** | Consent request template rejection datetime in UTC timezone. | [optional] 
**RejectionReason** | Pointer to **string** | Consent request template rejection reason. | [optional] 

## Methods

### NewConsentRequestTemplate

`func NewConsentRequestTemplate(id string, title string, description string, shortId string, status ConsentRequestTemplateStatus, createdBy string, createdAtUtc time.Time, ) *ConsentRequestTemplate`

NewConsentRequestTemplate instantiates a new ConsentRequestTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsentRequestTemplateWithDefaults

`func NewConsentRequestTemplateWithDefaults() *ConsentRequestTemplate`

NewConsentRequestTemplateWithDefaults instantiates a new ConsentRequestTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConsentRequestTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsentRequestTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsentRequestTemplate) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *ConsentRequestTemplate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ConsentRequestTemplate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ConsentRequestTemplate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *ConsentRequestTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConsentRequestTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConsentRequestTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPurpose

`func (o *ConsentRequestTemplate) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *ConsentRequestTemplate) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *ConsentRequestTemplate) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *ConsentRequestTemplate) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetShortId

`func (o *ConsentRequestTemplate) GetShortId() string`

GetShortId returns the ShortId field if non-nil, zero value otherwise.

### GetShortIdOk

`func (o *ConsentRequestTemplate) GetShortIdOk() (*string, bool)`

GetShortIdOk returns a tuple with the ShortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortId

`func (o *ConsentRequestTemplate) SetShortId(v string)`

SetShortId sets ShortId field to given value.


### GetStatus

`func (o *ConsentRequestTemplate) GetStatus() ConsentRequestTemplateStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConsentRequestTemplate) GetStatusOk() (*ConsentRequestTemplateStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConsentRequestTemplate) SetStatus(v ConsentRequestTemplateStatus)`

SetStatus sets Status field to given value.


### GetCreatedBy

`func (o *ConsentRequestTemplate) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ConsentRequestTemplate) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ConsentRequestTemplate) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAtUtc

`func (o *ConsentRequestTemplate) GetCreatedAtUtc() time.Time`

GetCreatedAtUtc returns the CreatedAtUtc field if non-nil, zero value otherwise.

### GetCreatedAtUtcOk

`func (o *ConsentRequestTemplate) GetCreatedAtUtcOk() (*time.Time, bool)`

GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtUtc

`func (o *ConsentRequestTemplate) SetCreatedAtUtc(v time.Time)`

SetCreatedAtUtc sets CreatedAtUtc field to given value.


### GetApprovedAtUtc

`func (o *ConsentRequestTemplate) GetApprovedAtUtc() time.Time`

GetApprovedAtUtc returns the ApprovedAtUtc field if non-nil, zero value otherwise.

### GetApprovedAtUtcOk

`func (o *ConsentRequestTemplate) GetApprovedAtUtcOk() (*time.Time, bool)`

GetApprovedAtUtcOk returns a tuple with the ApprovedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAtUtc

`func (o *ConsentRequestTemplate) SetApprovedAtUtc(v time.Time)`

SetApprovedAtUtc sets ApprovedAtUtc field to given value.

### HasApprovedAtUtc

`func (o *ConsentRequestTemplate) HasApprovedAtUtc() bool`

HasApprovedAtUtc returns a boolean if a field has been set.

### GetRejectedAtUtc

`func (o *ConsentRequestTemplate) GetRejectedAtUtc() time.Time`

GetRejectedAtUtc returns the RejectedAtUtc field if non-nil, zero value otherwise.

### GetRejectedAtUtcOk

`func (o *ConsentRequestTemplate) GetRejectedAtUtcOk() (*time.Time, bool)`

GetRejectedAtUtcOk returns a tuple with the RejectedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedAtUtc

`func (o *ConsentRequestTemplate) SetRejectedAtUtc(v time.Time)`

SetRejectedAtUtc sets RejectedAtUtc field to given value.

### HasRejectedAtUtc

`func (o *ConsentRequestTemplate) HasRejectedAtUtc() bool`

HasRejectedAtUtc returns a boolean if a field has been set.

### GetRejectionReason

`func (o *ConsentRequestTemplate) GetRejectionReason() string`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *ConsentRequestTemplate) GetRejectionReasonOk() (*string, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *ConsentRequestTemplate) SetRejectionReason(v string)`

SetRejectionReason sets RejectionReason field to given value.

### HasRejectionReason

`func (o *ConsentRequestTemplate) HasRejectionReason() bool`

HasRejectionReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



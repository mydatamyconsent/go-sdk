# OrganizationDataConsentInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsentRequestId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**ConsentTemplateId** | Pointer to **NullableString** |  | [optional] 
**ConsentPurpose** | Pointer to **NullableString** |  | [optional] 
**ConsentDescription** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**DataConsentStatus**](DataConsentStatus.md) |  | [optional] 
**CreatedAtUtc** | Pointer to **time.Time** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**ConsentSentToOrganization** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewOrganizationDataConsentInfoDto

`func NewOrganizationDataConsentInfoDto() *OrganizationDataConsentInfoDto`

NewOrganizationDataConsentInfoDto instantiates a new OrganizationDataConsentInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationDataConsentInfoDtoWithDefaults

`func NewOrganizationDataConsentInfoDtoWithDefaults() *OrganizationDataConsentInfoDto`

NewOrganizationDataConsentInfoDtoWithDefaults instantiates a new OrganizationDataConsentInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsentRequestId

`func (o *OrganizationDataConsentInfoDto) GetConsentRequestId() string`

GetConsentRequestId returns the ConsentRequestId field if non-nil, zero value otherwise.

### GetConsentRequestIdOk

`func (o *OrganizationDataConsentInfoDto) GetConsentRequestIdOk() (*string, bool)`

GetConsentRequestIdOk returns a tuple with the ConsentRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentRequestId

`func (o *OrganizationDataConsentInfoDto) SetConsentRequestId(v string)`

SetConsentRequestId sets ConsentRequestId field to given value.

### HasConsentRequestId

`func (o *OrganizationDataConsentInfoDto) HasConsentRequestId() bool`

HasConsentRequestId returns a boolean if a field has been set.

### GetUserId

`func (o *OrganizationDataConsentInfoDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OrganizationDataConsentInfoDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OrganizationDataConsentInfoDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *OrganizationDataConsentInfoDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *OrganizationDataConsentInfoDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *OrganizationDataConsentInfoDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetConsentTemplateId

`func (o *OrganizationDataConsentInfoDto) GetConsentTemplateId() string`

GetConsentTemplateId returns the ConsentTemplateId field if non-nil, zero value otherwise.

### GetConsentTemplateIdOk

`func (o *OrganizationDataConsentInfoDto) GetConsentTemplateIdOk() (*string, bool)`

GetConsentTemplateIdOk returns a tuple with the ConsentTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentTemplateId

`func (o *OrganizationDataConsentInfoDto) SetConsentTemplateId(v string)`

SetConsentTemplateId sets ConsentTemplateId field to given value.

### HasConsentTemplateId

`func (o *OrganizationDataConsentInfoDto) HasConsentTemplateId() bool`

HasConsentTemplateId returns a boolean if a field has been set.

### SetConsentTemplateIdNil

`func (o *OrganizationDataConsentInfoDto) SetConsentTemplateIdNil(b bool)`

 SetConsentTemplateIdNil sets the value for ConsentTemplateId to be an explicit nil

### UnsetConsentTemplateId
`func (o *OrganizationDataConsentInfoDto) UnsetConsentTemplateId()`

UnsetConsentTemplateId ensures that no value is present for ConsentTemplateId, not even an explicit nil
### GetConsentPurpose

`func (o *OrganizationDataConsentInfoDto) GetConsentPurpose() string`

GetConsentPurpose returns the ConsentPurpose field if non-nil, zero value otherwise.

### GetConsentPurposeOk

`func (o *OrganizationDataConsentInfoDto) GetConsentPurposeOk() (*string, bool)`

GetConsentPurposeOk returns a tuple with the ConsentPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentPurpose

`func (o *OrganizationDataConsentInfoDto) SetConsentPurpose(v string)`

SetConsentPurpose sets ConsentPurpose field to given value.

### HasConsentPurpose

`func (o *OrganizationDataConsentInfoDto) HasConsentPurpose() bool`

HasConsentPurpose returns a boolean if a field has been set.

### SetConsentPurposeNil

`func (o *OrganizationDataConsentInfoDto) SetConsentPurposeNil(b bool)`

 SetConsentPurposeNil sets the value for ConsentPurpose to be an explicit nil

### UnsetConsentPurpose
`func (o *OrganizationDataConsentInfoDto) UnsetConsentPurpose()`

UnsetConsentPurpose ensures that no value is present for ConsentPurpose, not even an explicit nil
### GetConsentDescription

`func (o *OrganizationDataConsentInfoDto) GetConsentDescription() string`

GetConsentDescription returns the ConsentDescription field if non-nil, zero value otherwise.

### GetConsentDescriptionOk

`func (o *OrganizationDataConsentInfoDto) GetConsentDescriptionOk() (*string, bool)`

GetConsentDescriptionOk returns a tuple with the ConsentDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentDescription

`func (o *OrganizationDataConsentInfoDto) SetConsentDescription(v string)`

SetConsentDescription sets ConsentDescription field to given value.

### HasConsentDescription

`func (o *OrganizationDataConsentInfoDto) HasConsentDescription() bool`

HasConsentDescription returns a boolean if a field has been set.

### SetConsentDescriptionNil

`func (o *OrganizationDataConsentInfoDto) SetConsentDescriptionNil(b bool)`

 SetConsentDescriptionNil sets the value for ConsentDescription to be an explicit nil

### UnsetConsentDescription
`func (o *OrganizationDataConsentInfoDto) UnsetConsentDescription()`

UnsetConsentDescription ensures that no value is present for ConsentDescription, not even an explicit nil
### GetStatus

`func (o *OrganizationDataConsentInfoDto) GetStatus() DataConsentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrganizationDataConsentInfoDto) GetStatusOk() (*DataConsentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrganizationDataConsentInfoDto) SetStatus(v DataConsentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrganizationDataConsentInfoDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAtUtc

`func (o *OrganizationDataConsentInfoDto) GetCreatedAtUtc() time.Time`

GetCreatedAtUtc returns the CreatedAtUtc field if non-nil, zero value otherwise.

### GetCreatedAtUtcOk

`func (o *OrganizationDataConsentInfoDto) GetCreatedAtUtcOk() (*time.Time, bool)`

GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtUtc

`func (o *OrganizationDataConsentInfoDto) SetCreatedAtUtc(v time.Time)`

SetCreatedAtUtc sets CreatedAtUtc field to given value.

### HasCreatedAtUtc

`func (o *OrganizationDataConsentInfoDto) HasCreatedAtUtc() bool`

HasCreatedAtUtc returns a boolean if a field has been set.

### GetOrganizationId

`func (o *OrganizationDataConsentInfoDto) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OrganizationDataConsentInfoDto) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OrganizationDataConsentInfoDto) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *OrganizationDataConsentInfoDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetConsentSentToOrganization

`func (o *OrganizationDataConsentInfoDto) GetConsentSentToOrganization() string`

GetConsentSentToOrganization returns the ConsentSentToOrganization field if non-nil, zero value otherwise.

### GetConsentSentToOrganizationOk

`func (o *OrganizationDataConsentInfoDto) GetConsentSentToOrganizationOk() (*string, bool)`

GetConsentSentToOrganizationOk returns a tuple with the ConsentSentToOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentSentToOrganization

`func (o *OrganizationDataConsentInfoDto) SetConsentSentToOrganization(v string)`

SetConsentSentToOrganization sets ConsentSentToOrganization field to given value.

### HasConsentSentToOrganization

`func (o *OrganizationDataConsentInfoDto) HasConsentSentToOrganization() bool`

HasConsentSentToOrganization returns a boolean if a field has been set.

### SetConsentSentToOrganizationNil

`func (o *OrganizationDataConsentInfoDto) SetConsentSentToOrganizationNil(b bool)`

 SetConsentSentToOrganizationNil sets the value for ConsentSentToOrganization to be an explicit nil

### UnsetConsentSentToOrganization
`func (o *OrganizationDataConsentInfoDto) UnsetConsentSentToOrganization()`

UnsetConsentSentToOrganization ensures that no value is present for ConsentSentToOrganization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



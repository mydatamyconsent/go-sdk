# UserDataConsentInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsentRequestId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**ConsentTemplateId** | Pointer to **NullableString** |  | [optional] 
**ConsentPurpose** | Pointer to **NullableString** |  | [optional] 
**ConsentDescription** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**DataConsentStatus**](DataConsentStatus.md) |  | [optional] 
**CreatedAtUtc** | Pointer to **time.Time** |  | [optional] 
**ConsentSentToUser** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUserDataConsentInfoDto

`func NewUserDataConsentInfoDto() *UserDataConsentInfoDto`

NewUserDataConsentInfoDto instantiates a new UserDataConsentInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDataConsentInfoDtoWithDefaults

`func NewUserDataConsentInfoDtoWithDefaults() *UserDataConsentInfoDto`

NewUserDataConsentInfoDtoWithDefaults instantiates a new UserDataConsentInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsentRequestId

`func (o *UserDataConsentInfoDto) GetConsentRequestId() string`

GetConsentRequestId returns the ConsentRequestId field if non-nil, zero value otherwise.

### GetConsentRequestIdOk

`func (o *UserDataConsentInfoDto) GetConsentRequestIdOk() (*string, bool)`

GetConsentRequestIdOk returns a tuple with the ConsentRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentRequestId

`func (o *UserDataConsentInfoDto) SetConsentRequestId(v string)`

SetConsentRequestId sets ConsentRequestId field to given value.

### HasConsentRequestId

`func (o *UserDataConsentInfoDto) HasConsentRequestId() bool`

HasConsentRequestId returns a boolean if a field has been set.

### GetUserId

`func (o *UserDataConsentInfoDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserDataConsentInfoDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserDataConsentInfoDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserDataConsentInfoDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetConsentTemplateId

`func (o *UserDataConsentInfoDto) GetConsentTemplateId() string`

GetConsentTemplateId returns the ConsentTemplateId field if non-nil, zero value otherwise.

### GetConsentTemplateIdOk

`func (o *UserDataConsentInfoDto) GetConsentTemplateIdOk() (*string, bool)`

GetConsentTemplateIdOk returns a tuple with the ConsentTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentTemplateId

`func (o *UserDataConsentInfoDto) SetConsentTemplateId(v string)`

SetConsentTemplateId sets ConsentTemplateId field to given value.

### HasConsentTemplateId

`func (o *UserDataConsentInfoDto) HasConsentTemplateId() bool`

HasConsentTemplateId returns a boolean if a field has been set.

### SetConsentTemplateIdNil

`func (o *UserDataConsentInfoDto) SetConsentTemplateIdNil(b bool)`

 SetConsentTemplateIdNil sets the value for ConsentTemplateId to be an explicit nil

### UnsetConsentTemplateId
`func (o *UserDataConsentInfoDto) UnsetConsentTemplateId()`

UnsetConsentTemplateId ensures that no value is present for ConsentTemplateId, not even an explicit nil
### GetConsentPurpose

`func (o *UserDataConsentInfoDto) GetConsentPurpose() string`

GetConsentPurpose returns the ConsentPurpose field if non-nil, zero value otherwise.

### GetConsentPurposeOk

`func (o *UserDataConsentInfoDto) GetConsentPurposeOk() (*string, bool)`

GetConsentPurposeOk returns a tuple with the ConsentPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentPurpose

`func (o *UserDataConsentInfoDto) SetConsentPurpose(v string)`

SetConsentPurpose sets ConsentPurpose field to given value.

### HasConsentPurpose

`func (o *UserDataConsentInfoDto) HasConsentPurpose() bool`

HasConsentPurpose returns a boolean if a field has been set.

### SetConsentPurposeNil

`func (o *UserDataConsentInfoDto) SetConsentPurposeNil(b bool)`

 SetConsentPurposeNil sets the value for ConsentPurpose to be an explicit nil

### UnsetConsentPurpose
`func (o *UserDataConsentInfoDto) UnsetConsentPurpose()`

UnsetConsentPurpose ensures that no value is present for ConsentPurpose, not even an explicit nil
### GetConsentDescription

`func (o *UserDataConsentInfoDto) GetConsentDescription() string`

GetConsentDescription returns the ConsentDescription field if non-nil, zero value otherwise.

### GetConsentDescriptionOk

`func (o *UserDataConsentInfoDto) GetConsentDescriptionOk() (*string, bool)`

GetConsentDescriptionOk returns a tuple with the ConsentDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentDescription

`func (o *UserDataConsentInfoDto) SetConsentDescription(v string)`

SetConsentDescription sets ConsentDescription field to given value.

### HasConsentDescription

`func (o *UserDataConsentInfoDto) HasConsentDescription() bool`

HasConsentDescription returns a boolean if a field has been set.

### SetConsentDescriptionNil

`func (o *UserDataConsentInfoDto) SetConsentDescriptionNil(b bool)`

 SetConsentDescriptionNil sets the value for ConsentDescription to be an explicit nil

### UnsetConsentDescription
`func (o *UserDataConsentInfoDto) UnsetConsentDescription()`

UnsetConsentDescription ensures that no value is present for ConsentDescription, not even an explicit nil
### GetStatus

`func (o *UserDataConsentInfoDto) GetStatus() DataConsentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserDataConsentInfoDto) GetStatusOk() (*DataConsentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserDataConsentInfoDto) SetStatus(v DataConsentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserDataConsentInfoDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAtUtc

`func (o *UserDataConsentInfoDto) GetCreatedAtUtc() time.Time`

GetCreatedAtUtc returns the CreatedAtUtc field if non-nil, zero value otherwise.

### GetCreatedAtUtcOk

`func (o *UserDataConsentInfoDto) GetCreatedAtUtcOk() (*time.Time, bool)`

GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtUtc

`func (o *UserDataConsentInfoDto) SetCreatedAtUtc(v time.Time)`

SetCreatedAtUtc sets CreatedAtUtc field to given value.

### HasCreatedAtUtc

`func (o *UserDataConsentInfoDto) HasCreatedAtUtc() bool`

HasCreatedAtUtc returns a boolean if a field has been set.

### GetConsentSentToUser

`func (o *UserDataConsentInfoDto) GetConsentSentToUser() string`

GetConsentSentToUser returns the ConsentSentToUser field if non-nil, zero value otherwise.

### GetConsentSentToUserOk

`func (o *UserDataConsentInfoDto) GetConsentSentToUserOk() (*string, bool)`

GetConsentSentToUserOk returns a tuple with the ConsentSentToUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentSentToUser

`func (o *UserDataConsentInfoDto) SetConsentSentToUser(v string)`

SetConsentSentToUser sets ConsentSentToUser field to given value.

### HasConsentSentToUser

`func (o *UserDataConsentInfoDto) HasConsentSentToUser() bool`

HasConsentSentToUser returns a boolean if a field has been set.

### SetConsentSentToUserNil

`func (o *UserDataConsentInfoDto) SetConsentSentToUserNil(b bool)`

 SetConsentSentToUserNil sets the value for ConsentSentToUser to be an explicit nil

### UnsetConsentSentToUser
`func (o *UserDataConsentInfoDto) UnsetConsentSentToUser()`

UnsetConsentSentToUser ensures that no value is present for ConsentSentToUser, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



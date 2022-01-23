# DataConsentDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**DataConsentStatus**](DataConsentStatus.md) |  | [optional] 
**ApprovedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**RejectedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**ExpiresAtUtc** | Pointer to **time.Time** |  | [optional] 
**RequestedAtUtc** | Pointer to **time.Time** |  | [optional] 
**Requester** | Pointer to [**DataConsentRequesterDto**](DataConsentRequesterDto.md) |  | [optional] 
**ConsentDetails** | Pointer to [**GetConsentTemplateDetailsDto**](GetConsentTemplateDetailsDto.md) |  | [optional] 

## Methods

### NewDataConsentDetailsDto

`func NewDataConsentDetailsDto() *DataConsentDetailsDto`

NewDataConsentDetailsDto instantiates a new DataConsentDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataConsentDetailsDtoWithDefaults

`func NewDataConsentDetailsDtoWithDefaults() *DataConsentDetailsDto`

NewDataConsentDetailsDtoWithDefaults instantiates a new DataConsentDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataConsentDetailsDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataConsentDetailsDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataConsentDetailsDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataConsentDetailsDto) HasId() bool`

HasId returns a boolean if a field has been set.

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
### GetExpiresAtUtc

`func (o *DataConsentDetailsDto) GetExpiresAtUtc() time.Time`

GetExpiresAtUtc returns the ExpiresAtUtc field if non-nil, zero value otherwise.

### GetExpiresAtUtcOk

`func (o *DataConsentDetailsDto) GetExpiresAtUtcOk() (*time.Time, bool)`

GetExpiresAtUtcOk returns a tuple with the ExpiresAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAtUtc

`func (o *DataConsentDetailsDto) SetExpiresAtUtc(v time.Time)`

SetExpiresAtUtc sets ExpiresAtUtc field to given value.

### HasExpiresAtUtc

`func (o *DataConsentDetailsDto) HasExpiresAtUtc() bool`

HasExpiresAtUtc returns a boolean if a field has been set.

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

### GetRequester

`func (o *DataConsentDetailsDto) GetRequester() DataConsentRequesterDto`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *DataConsentDetailsDto) GetRequesterOk() (*DataConsentRequesterDto, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *DataConsentDetailsDto) SetRequester(v DataConsentRequesterDto)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *DataConsentDetailsDto) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### GetConsentDetails

`func (o *DataConsentDetailsDto) GetConsentDetails() GetConsentTemplateDetailsDto`

GetConsentDetails returns the ConsentDetails field if non-nil, zero value otherwise.

### GetConsentDetailsOk

`func (o *DataConsentDetailsDto) GetConsentDetailsOk() (*GetConsentTemplateDetailsDto, bool)`

GetConsentDetailsOk returns a tuple with the ConsentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentDetails

`func (o *DataConsentDetailsDto) SetConsentDetails(v GetConsentTemplateDetailsDto)`

SetConsentDetails sets ConsentDetails field to given value.

### HasConsentDetails

`func (o *DataConsentDetailsDto) HasConsentDetails() bool`

HasConsentDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



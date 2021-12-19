# Rejection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **NullableString** |  | [optional] 
**RejectedBy** | Pointer to **string** |  | [optional] 
**RejectedAtUtc** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRejection

`func NewRejection() *Rejection`

NewRejection instantiates a new Rejection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRejectionWithDefaults

`func NewRejectionWithDefaults() *Rejection`

NewRejectionWithDefaults instantiates a new Rejection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *Rejection) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Rejection) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Rejection) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Rejection) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *Rejection) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *Rejection) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetRejectedBy

`func (o *Rejection) GetRejectedBy() string`

GetRejectedBy returns the RejectedBy field if non-nil, zero value otherwise.

### GetRejectedByOk

`func (o *Rejection) GetRejectedByOk() (*string, bool)`

GetRejectedByOk returns a tuple with the RejectedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedBy

`func (o *Rejection) SetRejectedBy(v string)`

SetRejectedBy sets RejectedBy field to given value.

### HasRejectedBy

`func (o *Rejection) HasRejectedBy() bool`

HasRejectedBy returns a boolean if a field has been set.

### GetRejectedAtUtc

`func (o *Rejection) GetRejectedAtUtc() time.Time`

GetRejectedAtUtc returns the RejectedAtUtc field if non-nil, zero value otherwise.

### GetRejectedAtUtcOk

`func (o *Rejection) GetRejectedAtUtcOk() (*time.Time, bool)`

GetRejectedAtUtcOk returns a tuple with the RejectedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedAtUtc

`func (o *Rejection) SetRejectedAtUtc(v time.Time)`

SetRejectedAtUtc sets RejectedAtUtc field to given value.

### HasRejectedAtUtc

`func (o *Rejection) HasRejectedAtUtc() bool`

HasRejectedAtUtc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



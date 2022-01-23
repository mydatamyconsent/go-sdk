# FinancialAccounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Drn** | Pointer to **NullableString** |  | [optional] 
**FinancialAccountDetailsRequired** | Pointer to [**[]FinancialAccountDetailsRequired**](FinancialAccountDetailsRequired.md) |  | [optional] 
**StartDate** | Pointer to **NullableTime** |  | [optional] 
**EndDate** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewFinancialAccounts

`func NewFinancialAccounts() *FinancialAccounts`

NewFinancialAccounts instantiates a new FinancialAccounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialAccountsWithDefaults

`func NewFinancialAccountsWithDefaults() *FinancialAccounts`

NewFinancialAccountsWithDefaults instantiates a new FinancialAccounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrn

`func (o *FinancialAccounts) GetDrn() string`

GetDrn returns the Drn field if non-nil, zero value otherwise.

### GetDrnOk

`func (o *FinancialAccounts) GetDrnOk() (*string, bool)`

GetDrnOk returns a tuple with the Drn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrn

`func (o *FinancialAccounts) SetDrn(v string)`

SetDrn sets Drn field to given value.

### HasDrn

`func (o *FinancialAccounts) HasDrn() bool`

HasDrn returns a boolean if a field has been set.

### SetDrnNil

`func (o *FinancialAccounts) SetDrnNil(b bool)`

 SetDrnNil sets the value for Drn to be an explicit nil

### UnsetDrn
`func (o *FinancialAccounts) UnsetDrn()`

UnsetDrn ensures that no value is present for Drn, not even an explicit nil
### GetFinancialAccountDetailsRequired

`func (o *FinancialAccounts) GetFinancialAccountDetailsRequired() []FinancialAccountDetailsRequired`

GetFinancialAccountDetailsRequired returns the FinancialAccountDetailsRequired field if non-nil, zero value otherwise.

### GetFinancialAccountDetailsRequiredOk

`func (o *FinancialAccounts) GetFinancialAccountDetailsRequiredOk() (*[]FinancialAccountDetailsRequired, bool)`

GetFinancialAccountDetailsRequiredOk returns a tuple with the FinancialAccountDetailsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialAccountDetailsRequired

`func (o *FinancialAccounts) SetFinancialAccountDetailsRequired(v []FinancialAccountDetailsRequired)`

SetFinancialAccountDetailsRequired sets FinancialAccountDetailsRequired field to given value.

### HasFinancialAccountDetailsRequired

`func (o *FinancialAccounts) HasFinancialAccountDetailsRequired() bool`

HasFinancialAccountDetailsRequired returns a boolean if a field has been set.

### SetFinancialAccountDetailsRequiredNil

`func (o *FinancialAccounts) SetFinancialAccountDetailsRequiredNil(b bool)`

 SetFinancialAccountDetailsRequiredNil sets the value for FinancialAccountDetailsRequired to be an explicit nil

### UnsetFinancialAccountDetailsRequired
`func (o *FinancialAccounts) UnsetFinancialAccountDetailsRequired()`

UnsetFinancialAccountDetailsRequired ensures that no value is present for FinancialAccountDetailsRequired, not even an explicit nil
### GetStartDate

`func (o *FinancialAccounts) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *FinancialAccounts) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *FinancialAccounts) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *FinancialAccounts) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *FinancialAccounts) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *FinancialAccounts) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *FinancialAccounts) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *FinancialAccounts) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *FinancialAccounts) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *FinancialAccounts) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *FinancialAccounts) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *FinancialAccounts) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



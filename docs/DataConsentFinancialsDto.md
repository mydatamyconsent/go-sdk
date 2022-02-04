# DataConsentFinancialsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Financials** | Pointer to [**[]Financial**](Financial.md) |  | [optional] 
**ApprovedFinancials** | Pointer to [**[]DataConsentRequestedFinancialAccount**](DataConsentRequestedFinancialAccount.md) |  | [optional] 

## Methods

### NewDataConsentFinancialsDto

`func NewDataConsentFinancialsDto() *DataConsentFinancialsDto`

NewDataConsentFinancialsDto instantiates a new DataConsentFinancialsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataConsentFinancialsDtoWithDefaults

`func NewDataConsentFinancialsDtoWithDefaults() *DataConsentFinancialsDto`

NewDataConsentFinancialsDtoWithDefaults instantiates a new DataConsentFinancialsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataConsentFinancialsDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataConsentFinancialsDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataConsentFinancialsDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataConsentFinancialsDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFinancials

`func (o *DataConsentFinancialsDto) GetFinancials() []Financial`

GetFinancials returns the Financials field if non-nil, zero value otherwise.

### GetFinancialsOk

`func (o *DataConsentFinancialsDto) GetFinancialsOk() (*[]Financial, bool)`

GetFinancialsOk returns a tuple with the Financials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancials

`func (o *DataConsentFinancialsDto) SetFinancials(v []Financial)`

SetFinancials sets Financials field to given value.

### HasFinancials

`func (o *DataConsentFinancialsDto) HasFinancials() bool`

HasFinancials returns a boolean if a field has been set.

### SetFinancialsNil

`func (o *DataConsentFinancialsDto) SetFinancialsNil(b bool)`

 SetFinancialsNil sets the value for Financials to be an explicit nil

### UnsetFinancials
`func (o *DataConsentFinancialsDto) UnsetFinancials()`

UnsetFinancials ensures that no value is present for Financials, not even an explicit nil
### GetApprovedFinancials

`func (o *DataConsentFinancialsDto) GetApprovedFinancials() []DataConsentRequestedFinancialAccount`

GetApprovedFinancials returns the ApprovedFinancials field if non-nil, zero value otherwise.

### GetApprovedFinancialsOk

`func (o *DataConsentFinancialsDto) GetApprovedFinancialsOk() (*[]DataConsentRequestedFinancialAccount, bool)`

GetApprovedFinancialsOk returns a tuple with the ApprovedFinancials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedFinancials

`func (o *DataConsentFinancialsDto) SetApprovedFinancials(v []DataConsentRequestedFinancialAccount)`

SetApprovedFinancials sets ApprovedFinancials field to given value.

### HasApprovedFinancials

`func (o *DataConsentFinancialsDto) HasApprovedFinancials() bool`

HasApprovedFinancials returns a boolean if a field has been set.

### SetApprovedFinancialsNil

`func (o *DataConsentFinancialsDto) SetApprovedFinancialsNil(b bool)`

 SetApprovedFinancialsNil sets the value for ApprovedFinancials to be an explicit nil

### UnsetApprovedFinancials
`func (o *DataConsentFinancialsDto) UnsetApprovedFinancials()`

UnsetApprovedFinancials ensures that no value is present for ApprovedFinancials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



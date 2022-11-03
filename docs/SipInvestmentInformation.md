# SipInvestmentInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstalmentAmount** | **float64** |  | 
**Frequency** | **string** |  | 
**CompletedInstalments** | **float64** |  | 
**InvestmentValue** | **float64** |  | 
**LastInstalmentDate** | Pointer to **time.Time** |  | [optional] 
**NextInstalmentDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSipInvestmentInformation

`func NewSipInvestmentInformation(instalmentAmount float64, frequency string, completedInstalments float64, investmentValue float64, ) *SipInvestmentInformation`

NewSipInvestmentInformation instantiates a new SipInvestmentInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSipInvestmentInformationWithDefaults

`func NewSipInvestmentInformationWithDefaults() *SipInvestmentInformation`

NewSipInvestmentInformationWithDefaults instantiates a new SipInvestmentInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstalmentAmount

`func (o *SipInvestmentInformation) GetInstalmentAmount() float64`

GetInstalmentAmount returns the InstalmentAmount field if non-nil, zero value otherwise.

### GetInstalmentAmountOk

`func (o *SipInvestmentInformation) GetInstalmentAmountOk() (*float64, bool)`

GetInstalmentAmountOk returns a tuple with the InstalmentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalmentAmount

`func (o *SipInvestmentInformation) SetInstalmentAmount(v float64)`

SetInstalmentAmount sets InstalmentAmount field to given value.


### GetFrequency

`func (o *SipInvestmentInformation) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *SipInvestmentInformation) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *SipInvestmentInformation) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.


### GetCompletedInstalments

`func (o *SipInvestmentInformation) GetCompletedInstalments() float64`

GetCompletedInstalments returns the CompletedInstalments field if non-nil, zero value otherwise.

### GetCompletedInstalmentsOk

`func (o *SipInvestmentInformation) GetCompletedInstalmentsOk() (*float64, bool)`

GetCompletedInstalmentsOk returns a tuple with the CompletedInstalments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedInstalments

`func (o *SipInvestmentInformation) SetCompletedInstalments(v float64)`

SetCompletedInstalments sets CompletedInstalments field to given value.


### GetInvestmentValue

`func (o *SipInvestmentInformation) GetInvestmentValue() float64`

GetInvestmentValue returns the InvestmentValue field if non-nil, zero value otherwise.

### GetInvestmentValueOk

`func (o *SipInvestmentInformation) GetInvestmentValueOk() (*float64, bool)`

GetInvestmentValueOk returns a tuple with the InvestmentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentValue

`func (o *SipInvestmentInformation) SetInvestmentValue(v float64)`

SetInvestmentValue sets InvestmentValue field to given value.


### GetLastInstalmentDate

`func (o *SipInvestmentInformation) GetLastInstalmentDate() time.Time`

GetLastInstalmentDate returns the LastInstalmentDate field if non-nil, zero value otherwise.

### GetLastInstalmentDateOk

`func (o *SipInvestmentInformation) GetLastInstalmentDateOk() (*time.Time, bool)`

GetLastInstalmentDateOk returns a tuple with the LastInstalmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastInstalmentDate

`func (o *SipInvestmentInformation) SetLastInstalmentDate(v time.Time)`

SetLastInstalmentDate sets LastInstalmentDate field to given value.

### HasLastInstalmentDate

`func (o *SipInvestmentInformation) HasLastInstalmentDate() bool`

HasLastInstalmentDate returns a boolean if a field has been set.

### GetNextInstalmentDate

`func (o *SipInvestmentInformation) GetNextInstalmentDate() time.Time`

GetNextInstalmentDate returns the NextInstalmentDate field if non-nil, zero value otherwise.

### GetNextInstalmentDateOk

`func (o *SipInvestmentInformation) GetNextInstalmentDateOk() (*time.Time, bool)`

GetNextInstalmentDateOk returns a tuple with the NextInstalmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextInstalmentDate

`func (o *SipInvestmentInformation) SetNextInstalmentDate(v time.Time)`

SetNextInstalmentDate sets NextInstalmentDate field to given value.

### HasNextInstalmentDate

`func (o *SipInvestmentInformation) HasNextInstalmentDate() bool`

HasNextInstalmentDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



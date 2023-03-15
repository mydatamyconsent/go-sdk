# TermDepositInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepositType** | [**TermDepositAccountType**](TermDepositAccountType.md) |  | 
**Description** | **string** |  | 
**CurrencyCode** | **string** |  | 
**PrincipalAmount** | **float64** |  | 
**InterestRate** | **float32** |  | 
**InterestPayout** | [**TermDepositInterestPayoutType**](TermDepositInterestPayoutType.md) |  | 
**InterestComputation** | [**TermDepositInterestComputation**](TermDepositInterestComputation.md) |  | 
**CompoundingFrequency** | [**TermDepositCompoundingFrequency**](TermDepositCompoundingFrequency.md) |  | 
**InterestPeriodicPayoutAmount** | **float64** |  | 
**InterestOnMaturity** | **float64** |  | 
**DepositOpeningDate** | **time.Time** |  | 
**MaturityDate** | **time.Time** |  | 
**MaturityAmount** | **float64** |  | 
**TenureDays** | **int32** |  | 
**TenureMonths** | **int32** |  | 
**TenureYears** | **int32** |  | 

## Methods

### NewTermDepositInformation

`func NewTermDepositInformation(depositType TermDepositAccountType, description string, currencyCode string, principalAmount float64, interestRate float32, interestPayout TermDepositInterestPayoutType, interestComputation TermDepositInterestComputation, compoundingFrequency TermDepositCompoundingFrequency, interestPeriodicPayoutAmount float64, interestOnMaturity float64, depositOpeningDate time.Time, maturityDate time.Time, maturityAmount float64, tenureDays int32, tenureMonths int32, tenureYears int32, ) *TermDepositInformation`

NewTermDepositInformation instantiates a new TermDepositInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTermDepositInformationWithDefaults

`func NewTermDepositInformationWithDefaults() *TermDepositInformation`

NewTermDepositInformationWithDefaults instantiates a new TermDepositInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepositType

`func (o *TermDepositInformation) GetDepositType() TermDepositAccountType`

GetDepositType returns the DepositType field if non-nil, zero value otherwise.

### GetDepositTypeOk

`func (o *TermDepositInformation) GetDepositTypeOk() (*TermDepositAccountType, bool)`

GetDepositTypeOk returns a tuple with the DepositType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositType

`func (o *TermDepositInformation) SetDepositType(v TermDepositAccountType)`

SetDepositType sets DepositType field to given value.


### GetDescription

`func (o *TermDepositInformation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TermDepositInformation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TermDepositInformation) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCurrencyCode

`func (o *TermDepositInformation) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *TermDepositInformation) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *TermDepositInformation) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetPrincipalAmount

`func (o *TermDepositInformation) GetPrincipalAmount() float64`

GetPrincipalAmount returns the PrincipalAmount field if non-nil, zero value otherwise.

### GetPrincipalAmountOk

`func (o *TermDepositInformation) GetPrincipalAmountOk() (*float64, bool)`

GetPrincipalAmountOk returns a tuple with the PrincipalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalAmount

`func (o *TermDepositInformation) SetPrincipalAmount(v float64)`

SetPrincipalAmount sets PrincipalAmount field to given value.


### GetInterestRate

`func (o *TermDepositInformation) GetInterestRate() float32`

GetInterestRate returns the InterestRate field if non-nil, zero value otherwise.

### GetInterestRateOk

`func (o *TermDepositInformation) GetInterestRateOk() (*float32, bool)`

GetInterestRateOk returns a tuple with the InterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRate

`func (o *TermDepositInformation) SetInterestRate(v float32)`

SetInterestRate sets InterestRate field to given value.


### GetInterestPayout

`func (o *TermDepositInformation) GetInterestPayout() TermDepositInterestPayoutType`

GetInterestPayout returns the InterestPayout field if non-nil, zero value otherwise.

### GetInterestPayoutOk

`func (o *TermDepositInformation) GetInterestPayoutOk() (*TermDepositInterestPayoutType, bool)`

GetInterestPayoutOk returns a tuple with the InterestPayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestPayout

`func (o *TermDepositInformation) SetInterestPayout(v TermDepositInterestPayoutType)`

SetInterestPayout sets InterestPayout field to given value.


### GetInterestComputation

`func (o *TermDepositInformation) GetInterestComputation() TermDepositInterestComputation`

GetInterestComputation returns the InterestComputation field if non-nil, zero value otherwise.

### GetInterestComputationOk

`func (o *TermDepositInformation) GetInterestComputationOk() (*TermDepositInterestComputation, bool)`

GetInterestComputationOk returns a tuple with the InterestComputation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestComputation

`func (o *TermDepositInformation) SetInterestComputation(v TermDepositInterestComputation)`

SetInterestComputation sets InterestComputation field to given value.


### GetCompoundingFrequency

`func (o *TermDepositInformation) GetCompoundingFrequency() TermDepositCompoundingFrequency`

GetCompoundingFrequency returns the CompoundingFrequency field if non-nil, zero value otherwise.

### GetCompoundingFrequencyOk

`func (o *TermDepositInformation) GetCompoundingFrequencyOk() (*TermDepositCompoundingFrequency, bool)`

GetCompoundingFrequencyOk returns a tuple with the CompoundingFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompoundingFrequency

`func (o *TermDepositInformation) SetCompoundingFrequency(v TermDepositCompoundingFrequency)`

SetCompoundingFrequency sets CompoundingFrequency field to given value.


### GetInterestPeriodicPayoutAmount

`func (o *TermDepositInformation) GetInterestPeriodicPayoutAmount() float64`

GetInterestPeriodicPayoutAmount returns the InterestPeriodicPayoutAmount field if non-nil, zero value otherwise.

### GetInterestPeriodicPayoutAmountOk

`func (o *TermDepositInformation) GetInterestPeriodicPayoutAmountOk() (*float64, bool)`

GetInterestPeriodicPayoutAmountOk returns a tuple with the InterestPeriodicPayoutAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestPeriodicPayoutAmount

`func (o *TermDepositInformation) SetInterestPeriodicPayoutAmount(v float64)`

SetInterestPeriodicPayoutAmount sets InterestPeriodicPayoutAmount field to given value.


### GetInterestOnMaturity

`func (o *TermDepositInformation) GetInterestOnMaturity() float64`

GetInterestOnMaturity returns the InterestOnMaturity field if non-nil, zero value otherwise.

### GetInterestOnMaturityOk

`func (o *TermDepositInformation) GetInterestOnMaturityOk() (*float64, bool)`

GetInterestOnMaturityOk returns a tuple with the InterestOnMaturity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestOnMaturity

`func (o *TermDepositInformation) SetInterestOnMaturity(v float64)`

SetInterestOnMaturity sets InterestOnMaturity field to given value.


### GetDepositOpeningDate

`func (o *TermDepositInformation) GetDepositOpeningDate() time.Time`

GetDepositOpeningDate returns the DepositOpeningDate field if non-nil, zero value otherwise.

### GetDepositOpeningDateOk

`func (o *TermDepositInformation) GetDepositOpeningDateOk() (*time.Time, bool)`

GetDepositOpeningDateOk returns a tuple with the DepositOpeningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositOpeningDate

`func (o *TermDepositInformation) SetDepositOpeningDate(v time.Time)`

SetDepositOpeningDate sets DepositOpeningDate field to given value.


### GetMaturityDate

`func (o *TermDepositInformation) GetMaturityDate() time.Time`

GetMaturityDate returns the MaturityDate field if non-nil, zero value otherwise.

### GetMaturityDateOk

`func (o *TermDepositInformation) GetMaturityDateOk() (*time.Time, bool)`

GetMaturityDateOk returns a tuple with the MaturityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaturityDate

`func (o *TermDepositInformation) SetMaturityDate(v time.Time)`

SetMaturityDate sets MaturityDate field to given value.


### GetMaturityAmount

`func (o *TermDepositInformation) GetMaturityAmount() float64`

GetMaturityAmount returns the MaturityAmount field if non-nil, zero value otherwise.

### GetMaturityAmountOk

`func (o *TermDepositInformation) GetMaturityAmountOk() (*float64, bool)`

GetMaturityAmountOk returns a tuple with the MaturityAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaturityAmount

`func (o *TermDepositInformation) SetMaturityAmount(v float64)`

SetMaturityAmount sets MaturityAmount field to given value.


### GetTenureDays

`func (o *TermDepositInformation) GetTenureDays() int32`

GetTenureDays returns the TenureDays field if non-nil, zero value otherwise.

### GetTenureDaysOk

`func (o *TermDepositInformation) GetTenureDaysOk() (*int32, bool)`

GetTenureDaysOk returns a tuple with the TenureDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenureDays

`func (o *TermDepositInformation) SetTenureDays(v int32)`

SetTenureDays sets TenureDays field to given value.


### GetTenureMonths

`func (o *TermDepositInformation) GetTenureMonths() int32`

GetTenureMonths returns the TenureMonths field if non-nil, zero value otherwise.

### GetTenureMonthsOk

`func (o *TermDepositInformation) GetTenureMonthsOk() (*int32, bool)`

GetTenureMonthsOk returns a tuple with the TenureMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenureMonths

`func (o *TermDepositInformation) SetTenureMonths(v int32)`

SetTenureMonths sets TenureMonths field to given value.


### GetTenureYears

`func (o *TermDepositInformation) GetTenureYears() int32`

GetTenureYears returns the TenureYears field if non-nil, zero value otherwise.

### GetTenureYearsOk

`func (o *TermDepositInformation) GetTenureYearsOk() (*int32, bool)`

GetTenureYearsOk returns a tuple with the TenureYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenureYears

`func (o *TermDepositInformation) SetTenureYears(v int32)`

SetTenureYears sets TenureYears field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



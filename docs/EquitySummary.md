# EquitySummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Investment** | [**EquityInvestment**](EquityInvestment.md) |  | 
**InvestmentValue** | **float64** |  | 
**CurrentValue** | **float64** |  | 

## Methods

### NewEquitySummary

`func NewEquitySummary(investment EquityInvestment, investmentValue float64, currentValue float64, ) *EquitySummary`

NewEquitySummary instantiates a new EquitySummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquitySummaryWithDefaults

`func NewEquitySummaryWithDefaults() *EquitySummary`

NewEquitySummaryWithDefaults instantiates a new EquitySummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvestment

`func (o *EquitySummary) GetInvestment() EquityInvestment`

GetInvestment returns the Investment field if non-nil, zero value otherwise.

### GetInvestmentOk

`func (o *EquitySummary) GetInvestmentOk() (*EquityInvestment, bool)`

GetInvestmentOk returns a tuple with the Investment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestment

`func (o *EquitySummary) SetInvestment(v EquityInvestment)`

SetInvestment sets Investment field to given value.


### GetInvestmentValue

`func (o *EquitySummary) GetInvestmentValue() float64`

GetInvestmentValue returns the InvestmentValue field if non-nil, zero value otherwise.

### GetInvestmentValueOk

`func (o *EquitySummary) GetInvestmentValueOk() (*float64, bool)`

GetInvestmentValueOk returns a tuple with the InvestmentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentValue

`func (o *EquitySummary) SetInvestmentValue(v float64)`

SetInvestmentValue sets InvestmentValue field to given value.


### GetCurrentValue

`func (o *EquitySummary) GetCurrentValue() float64`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *EquitySummary) GetCurrentValueOk() (*float64, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *EquitySummary) SetCurrentValue(v float64)`

SetCurrentValue sets CurrentValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



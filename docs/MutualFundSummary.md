# MutualFundSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Investment** | [**MutualFundInvestment**](MutualFundInvestment.md) |  | 
**InvestmentValue** | **float64** |  | 
**CurrentValue** | **float64** |  | 

## Methods

### NewMutualFundSummary

`func NewMutualFundSummary(investment MutualFundInvestment, investmentValue float64, currentValue float64, ) *MutualFundSummary`

NewMutualFundSummary instantiates a new MutualFundSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutualFundSummaryWithDefaults

`func NewMutualFundSummaryWithDefaults() *MutualFundSummary`

NewMutualFundSummaryWithDefaults instantiates a new MutualFundSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvestment

`func (o *MutualFundSummary) GetInvestment() MutualFundInvestment`

GetInvestment returns the Investment field if non-nil, zero value otherwise.

### GetInvestmentOk

`func (o *MutualFundSummary) GetInvestmentOk() (*MutualFundInvestment, bool)`

GetInvestmentOk returns a tuple with the Investment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestment

`func (o *MutualFundSummary) SetInvestment(v MutualFundInvestment)`

SetInvestment sets Investment field to given value.


### GetInvestmentValue

`func (o *MutualFundSummary) GetInvestmentValue() float64`

GetInvestmentValue returns the InvestmentValue field if non-nil, zero value otherwise.

### GetInvestmentValueOk

`func (o *MutualFundSummary) GetInvestmentValueOk() (*float64, bool)`

GetInvestmentValueOk returns a tuple with the InvestmentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentValue

`func (o *MutualFundSummary) SetInvestmentValue(v float64)`

SetInvestmentValue sets InvestmentValue field to given value.


### GetCurrentValue

`func (o *MutualFundSummary) GetCurrentValue() float64`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *MutualFundSummary) GetCurrentValueOk() (*float64, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *MutualFundSummary) SetCurrentValue(v float64)`

SetCurrentValue sets CurrentValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



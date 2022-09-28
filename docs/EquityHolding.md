# EquityHolding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssuerName** | **string** |  | 
**Exchange** | **string** |  | 
**Isin** | **string** |  | 
**Units** | **int64** |  | 
**InvestmentDateTime** | Pointer to **time.Time** |  | [optional] 
**Rate** | Pointer to **string** |  | [optional] 
**LastTradedPrice** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewEquityHolding

`func NewEquityHolding(issuerName string, exchange string, isin string, units int64, ) *EquityHolding`

NewEquityHolding instantiates a new EquityHolding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquityHoldingWithDefaults

`func NewEquityHoldingWithDefaults() *EquityHolding`

NewEquityHoldingWithDefaults instantiates a new EquityHolding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuerName

`func (o *EquityHolding) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *EquityHolding) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *EquityHolding) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.


### GetExchange

`func (o *EquityHolding) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *EquityHolding) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *EquityHolding) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetIsin

`func (o *EquityHolding) GetIsin() string`

GetIsin returns the Isin field if non-nil, zero value otherwise.

### GetIsinOk

`func (o *EquityHolding) GetIsinOk() (*string, bool)`

GetIsinOk returns a tuple with the Isin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsin

`func (o *EquityHolding) SetIsin(v string)`

SetIsin sets Isin field to given value.


### GetUnits

`func (o *EquityHolding) GetUnits() int64`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *EquityHolding) GetUnitsOk() (*int64, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *EquityHolding) SetUnits(v int64)`

SetUnits sets Units field to given value.


### GetInvestmentDateTime

`func (o *EquityHolding) GetInvestmentDateTime() time.Time`

GetInvestmentDateTime returns the InvestmentDateTime field if non-nil, zero value otherwise.

### GetInvestmentDateTimeOk

`func (o *EquityHolding) GetInvestmentDateTimeOk() (*time.Time, bool)`

GetInvestmentDateTimeOk returns a tuple with the InvestmentDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentDateTime

`func (o *EquityHolding) SetInvestmentDateTime(v time.Time)`

SetInvestmentDateTime sets InvestmentDateTime field to given value.

### HasInvestmentDateTime

`func (o *EquityHolding) HasInvestmentDateTime() bool`

HasInvestmentDateTime returns a boolean if a field has been set.

### GetRate

`func (o *EquityHolding) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *EquityHolding) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *EquityHolding) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *EquityHolding) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetLastTradedPrice

`func (o *EquityHolding) GetLastTradedPrice() string`

GetLastTradedPrice returns the LastTradedPrice field if non-nil, zero value otherwise.

### GetLastTradedPriceOk

`func (o *EquityHolding) GetLastTradedPriceOk() (*string, bool)`

GetLastTradedPriceOk returns a tuple with the LastTradedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTradedPrice

`func (o *EquityHolding) SetLastTradedPrice(v string)`

SetLastTradedPrice sets LastTradedPrice field to given value.

### HasLastTradedPrice

`func (o *EquityHolding) HasLastTradedPrice() bool`

HasLastTradedPrice returns a boolean if a field has been set.

### GetDescription

`func (o *EquityHolding) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EquityHolding) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EquityHolding) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EquityHolding) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



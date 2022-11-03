# Equity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**IssuerName** | **string** |  | 
**Exchange** | **string** |  | 
**Isin** | **string** |  | 
**Units** | **int64** |  | 
**InvestmentValue** | **float64** |  | 
**CurrentValue** | **float64** |  | 
**CurrencyCode** | **string** |  | 
**Holder** | [**Holder**](Holder.md) |  | 
**Transactions** | **bool** |  | 

## Methods

### NewEquity

`func NewEquity(id string, name string, issuerName string, exchange string, isin string, units int64, investmentValue float64, currentValue float64, currencyCode string, holder Holder, transactions bool, ) *Equity`

NewEquity instantiates a new Equity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquityWithDefaults

`func NewEquityWithDefaults() *Equity`

NewEquityWithDefaults instantiates a new Equity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Equity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Equity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Equity) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Equity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Equity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Equity) SetName(v string)`

SetName sets Name field to given value.


### GetIssuerName

`func (o *Equity) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *Equity) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *Equity) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.


### GetExchange

`func (o *Equity) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *Equity) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *Equity) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetIsin

`func (o *Equity) GetIsin() string`

GetIsin returns the Isin field if non-nil, zero value otherwise.

### GetIsinOk

`func (o *Equity) GetIsinOk() (*string, bool)`

GetIsinOk returns a tuple with the Isin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsin

`func (o *Equity) SetIsin(v string)`

SetIsin sets Isin field to given value.


### GetUnits

`func (o *Equity) GetUnits() int64`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *Equity) GetUnitsOk() (*int64, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *Equity) SetUnits(v int64)`

SetUnits sets Units field to given value.


### GetInvestmentValue

`func (o *Equity) GetInvestmentValue() float64`

GetInvestmentValue returns the InvestmentValue field if non-nil, zero value otherwise.

### GetInvestmentValueOk

`func (o *Equity) GetInvestmentValueOk() (*float64, bool)`

GetInvestmentValueOk returns a tuple with the InvestmentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentValue

`func (o *Equity) SetInvestmentValue(v float64)`

SetInvestmentValue sets InvestmentValue field to given value.


### GetCurrentValue

`func (o *Equity) GetCurrentValue() float64`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *Equity) GetCurrentValueOk() (*float64, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *Equity) SetCurrentValue(v float64)`

SetCurrentValue sets CurrentValue field to given value.


### GetCurrencyCode

`func (o *Equity) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *Equity) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *Equity) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetHolder

`func (o *Equity) GetHolder() Holder`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *Equity) GetHolderOk() (*Holder, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *Equity) SetHolder(v Holder)`

SetHolder sets Holder field to given value.


### GetTransactions

`func (o *Equity) GetTransactions() bool`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *Equity) GetTransactionsOk() (*bool, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *Equity) SetTransactions(v bool)`

SetTransactions sets Transactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



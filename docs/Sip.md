# Sip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**InvestmentValue** | **float64** |  | 
**CurrentValue** | **float64** |  | 
**CurrencyCode** | **string** |  | 
**PlanInfo** | [**SipPlanInformation**](SipPlanInformation.md) |  | 
**InvestmentInfo** | [**SipInvestmentInformation**](SipInvestmentInformation.md) |  | 
**Holder** | [**Holder**](Holder.md) |  | 
**Transactions** | **bool** |  | 

## Methods

### NewSip

`func NewSip(id string, name string, investmentValue float64, currentValue float64, currencyCode string, planInfo SipPlanInformation, investmentInfo SipInvestmentInformation, holder Holder, transactions bool, ) *Sip`

NewSip instantiates a new Sip object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSipWithDefaults

`func NewSipWithDefaults() *Sip`

NewSipWithDefaults instantiates a new Sip object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Sip) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Sip) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Sip) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Sip) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Sip) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Sip) SetName(v string)`

SetName sets Name field to given value.


### GetInvestmentValue

`func (o *Sip) GetInvestmentValue() float64`

GetInvestmentValue returns the InvestmentValue field if non-nil, zero value otherwise.

### GetInvestmentValueOk

`func (o *Sip) GetInvestmentValueOk() (*float64, bool)`

GetInvestmentValueOk returns a tuple with the InvestmentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentValue

`func (o *Sip) SetInvestmentValue(v float64)`

SetInvestmentValue sets InvestmentValue field to given value.


### GetCurrentValue

`func (o *Sip) GetCurrentValue() float64`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *Sip) GetCurrentValueOk() (*float64, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *Sip) SetCurrentValue(v float64)`

SetCurrentValue sets CurrentValue field to given value.


### GetCurrencyCode

`func (o *Sip) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *Sip) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *Sip) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetPlanInfo

`func (o *Sip) GetPlanInfo() SipPlanInformation`

GetPlanInfo returns the PlanInfo field if non-nil, zero value otherwise.

### GetPlanInfoOk

`func (o *Sip) GetPlanInfoOk() (*SipPlanInformation, bool)`

GetPlanInfoOk returns a tuple with the PlanInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanInfo

`func (o *Sip) SetPlanInfo(v SipPlanInformation)`

SetPlanInfo sets PlanInfo field to given value.


### GetInvestmentInfo

`func (o *Sip) GetInvestmentInfo() SipInvestmentInformation`

GetInvestmentInfo returns the InvestmentInfo field if non-nil, zero value otherwise.

### GetInvestmentInfoOk

`func (o *Sip) GetInvestmentInfoOk() (*SipInvestmentInformation, bool)`

GetInvestmentInfoOk returns a tuple with the InvestmentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentInfo

`func (o *Sip) SetInvestmentInfo(v SipInvestmentInformation)`

SetInvestmentInfo sets InvestmentInfo field to given value.


### GetHolder

`func (o *Sip) GetHolder() Holder`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *Sip) GetHolderOk() (*Holder, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *Sip) SetHolder(v Holder)`

SetHolder sets Holder field to given value.


### GetTransactions

`func (o *Sip) GetTransactions() bool`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *Sip) GetTransactionsOk() (*bool, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *Sip) SetTransactions(v bool)`

SetTransactions sets Transactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



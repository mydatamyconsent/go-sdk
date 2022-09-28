# CreditCardSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditCardId** | **string** |  | 
**CurrentDue** | **float64** |  | 
**LastStatementDate** | **time.Time** |  | 
**DueDate** | **time.Time** |  | 
**PreviousDueAmount** | **float64** |  | 
**TotalDueAmount** | **float64** |  | 
**MinDueAmount** | **float64** |  | 
**CreditLimit** | **float64** |  | 
**CashLimit** | **float64** |  | 
**AvailableCredit** | **float64** |  | 
**LoyaltyPoints** | **string** |  | 
**FinanceCharges** | **float64** |  | 

## Methods

### NewCreditCardSummary

`func NewCreditCardSummary(creditCardId string, currentDue float64, lastStatementDate time.Time, dueDate time.Time, previousDueAmount float64, totalDueAmount float64, minDueAmount float64, creditLimit float64, cashLimit float64, availableCredit float64, loyaltyPoints string, financeCharges float64, ) *CreditCardSummary`

NewCreditCardSummary instantiates a new CreditCardSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditCardSummaryWithDefaults

`func NewCreditCardSummaryWithDefaults() *CreditCardSummary`

NewCreditCardSummaryWithDefaults instantiates a new CreditCardSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditCardId

`func (o *CreditCardSummary) GetCreditCardId() string`

GetCreditCardId returns the CreditCardId field if non-nil, zero value otherwise.

### GetCreditCardIdOk

`func (o *CreditCardSummary) GetCreditCardIdOk() (*string, bool)`

GetCreditCardIdOk returns a tuple with the CreditCardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCardId

`func (o *CreditCardSummary) SetCreditCardId(v string)`

SetCreditCardId sets CreditCardId field to given value.


### GetCurrentDue

`func (o *CreditCardSummary) GetCurrentDue() float64`

GetCurrentDue returns the CurrentDue field if non-nil, zero value otherwise.

### GetCurrentDueOk

`func (o *CreditCardSummary) GetCurrentDueOk() (*float64, bool)`

GetCurrentDueOk returns a tuple with the CurrentDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDue

`func (o *CreditCardSummary) SetCurrentDue(v float64)`

SetCurrentDue sets CurrentDue field to given value.


### GetLastStatementDate

`func (o *CreditCardSummary) GetLastStatementDate() time.Time`

GetLastStatementDate returns the LastStatementDate field if non-nil, zero value otherwise.

### GetLastStatementDateOk

`func (o *CreditCardSummary) GetLastStatementDateOk() (*time.Time, bool)`

GetLastStatementDateOk returns a tuple with the LastStatementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatementDate

`func (o *CreditCardSummary) SetLastStatementDate(v time.Time)`

SetLastStatementDate sets LastStatementDate field to given value.


### GetDueDate

`func (o *CreditCardSummary) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *CreditCardSummary) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *CreditCardSummary) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.


### GetPreviousDueAmount

`func (o *CreditCardSummary) GetPreviousDueAmount() float64`

GetPreviousDueAmount returns the PreviousDueAmount field if non-nil, zero value otherwise.

### GetPreviousDueAmountOk

`func (o *CreditCardSummary) GetPreviousDueAmountOk() (*float64, bool)`

GetPreviousDueAmountOk returns a tuple with the PreviousDueAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousDueAmount

`func (o *CreditCardSummary) SetPreviousDueAmount(v float64)`

SetPreviousDueAmount sets PreviousDueAmount field to given value.


### GetTotalDueAmount

`func (o *CreditCardSummary) GetTotalDueAmount() float64`

GetTotalDueAmount returns the TotalDueAmount field if non-nil, zero value otherwise.

### GetTotalDueAmountOk

`func (o *CreditCardSummary) GetTotalDueAmountOk() (*float64, bool)`

GetTotalDueAmountOk returns a tuple with the TotalDueAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDueAmount

`func (o *CreditCardSummary) SetTotalDueAmount(v float64)`

SetTotalDueAmount sets TotalDueAmount field to given value.


### GetMinDueAmount

`func (o *CreditCardSummary) GetMinDueAmount() float64`

GetMinDueAmount returns the MinDueAmount field if non-nil, zero value otherwise.

### GetMinDueAmountOk

`func (o *CreditCardSummary) GetMinDueAmountOk() (*float64, bool)`

GetMinDueAmountOk returns a tuple with the MinDueAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDueAmount

`func (o *CreditCardSummary) SetMinDueAmount(v float64)`

SetMinDueAmount sets MinDueAmount field to given value.


### GetCreditLimit

`func (o *CreditCardSummary) GetCreditLimit() float64`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *CreditCardSummary) GetCreditLimitOk() (*float64, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *CreditCardSummary) SetCreditLimit(v float64)`

SetCreditLimit sets CreditLimit field to given value.


### GetCashLimit

`func (o *CreditCardSummary) GetCashLimit() float64`

GetCashLimit returns the CashLimit field if non-nil, zero value otherwise.

### GetCashLimitOk

`func (o *CreditCardSummary) GetCashLimitOk() (*float64, bool)`

GetCashLimitOk returns a tuple with the CashLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashLimit

`func (o *CreditCardSummary) SetCashLimit(v float64)`

SetCashLimit sets CashLimit field to given value.


### GetAvailableCredit

`func (o *CreditCardSummary) GetAvailableCredit() float64`

GetAvailableCredit returns the AvailableCredit field if non-nil, zero value otherwise.

### GetAvailableCreditOk

`func (o *CreditCardSummary) GetAvailableCreditOk() (*float64, bool)`

GetAvailableCreditOk returns a tuple with the AvailableCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCredit

`func (o *CreditCardSummary) SetAvailableCredit(v float64)`

SetAvailableCredit sets AvailableCredit field to given value.


### GetLoyaltyPoints

`func (o *CreditCardSummary) GetLoyaltyPoints() string`

GetLoyaltyPoints returns the LoyaltyPoints field if non-nil, zero value otherwise.

### GetLoyaltyPointsOk

`func (o *CreditCardSummary) GetLoyaltyPointsOk() (*string, bool)`

GetLoyaltyPointsOk returns a tuple with the LoyaltyPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyPoints

`func (o *CreditCardSummary) SetLoyaltyPoints(v string)`

SetLoyaltyPoints sets LoyaltyPoints field to given value.


### GetFinanceCharges

`func (o *CreditCardSummary) GetFinanceCharges() float64`

GetFinanceCharges returns the FinanceCharges field if non-nil, zero value otherwise.

### GetFinanceChargesOk

`func (o *CreditCardSummary) GetFinanceChargesOk() (*float64, bool)`

GetFinanceChargesOk returns a tuple with the FinanceCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinanceCharges

`func (o *CreditCardSummary) SetFinanceCharges(v float64)`

SetFinanceCharges sets FinanceCharges field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



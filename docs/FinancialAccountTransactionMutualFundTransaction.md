# FinancialAccountTransactionMutualFundTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Amount** | **string** |  | 
**CurrencyCode** | **string** |  | 
**TxnType** | [**MutualFundTransactionType**](MutualFundTransactionType.md) |  | 
**Units** | **string** |  | 
**TransactedAtUtc** | **time.Time** |  | 

## Methods

### NewFinancialAccountTransactionMutualFundTransaction

`func NewFinancialAccountTransactionMutualFundTransaction(type_ string, id string, amount string, currencyCode string, txnType MutualFundTransactionType, units string, transactedAtUtc time.Time, ) *FinancialAccountTransactionMutualFundTransaction`

NewFinancialAccountTransactionMutualFundTransaction instantiates a new FinancialAccountTransactionMutualFundTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialAccountTransactionMutualFundTransactionWithDefaults

`func NewFinancialAccountTransactionMutualFundTransactionWithDefaults() *FinancialAccountTransactionMutualFundTransaction`

NewFinancialAccountTransactionMutualFundTransactionWithDefaults instantiates a new FinancialAccountTransactionMutualFundTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FinancialAccountTransactionMutualFundTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FinancialAccountTransactionMutualFundTransaction) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *FinancialAccountTransactionMutualFundTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FinancialAccountTransactionMutualFundTransaction) SetId(v string)`

SetId sets Id field to given value.


### GetAmount

`func (o *FinancialAccountTransactionMutualFundTransaction) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FinancialAccountTransactionMutualFundTransaction) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCurrencyCode

`func (o *FinancialAccountTransactionMutualFundTransaction) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *FinancialAccountTransactionMutualFundTransaction) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetTxnType

`func (o *FinancialAccountTransactionMutualFundTransaction) GetTxnType() MutualFundTransactionType`

GetTxnType returns the TxnType field if non-nil, zero value otherwise.

### GetTxnTypeOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetTxnTypeOk() (*MutualFundTransactionType, bool)`

GetTxnTypeOk returns a tuple with the TxnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnType

`func (o *FinancialAccountTransactionMutualFundTransaction) SetTxnType(v MutualFundTransactionType)`

SetTxnType sets TxnType field to given value.


### GetUnits

`func (o *FinancialAccountTransactionMutualFundTransaction) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *FinancialAccountTransactionMutualFundTransaction) SetUnits(v string)`

SetUnits sets Units field to given value.


### GetTransactedAtUtc

`func (o *FinancialAccountTransactionMutualFundTransaction) GetTransactedAtUtc() time.Time`

GetTransactedAtUtc returns the TransactedAtUtc field if non-nil, zero value otherwise.

### GetTransactedAtUtcOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetTransactedAtUtcOk() (*time.Time, bool)`

GetTransactedAtUtcOk returns a tuple with the TransactedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactedAtUtc

`func (o *FinancialAccountTransactionMutualFundTransaction) SetTransactedAtUtc(v time.Time)`

SetTransactedAtUtc sets TransactedAtUtc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



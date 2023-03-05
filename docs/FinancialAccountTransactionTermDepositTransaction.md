# FinancialAccountTransactionTermDepositTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Amount** | **float64** |  | 
**CurrencyCode** | **string** |  | 
**Narration** | **string** |  | 
**TxnType** | [**TermDepositTransactionType**](TermDepositTransactionType.md) |  | 
**Mode** | [**TermDepositTransactionMode**](TermDepositTransactionMode.md) |  | 
**TransactedAtUtc** | **time.Time** |  | 

## Methods

### NewFinancialAccountTransactionTermDepositTransaction

`func NewFinancialAccountTransactionTermDepositTransaction(type_ string, id string, amount float64, currencyCode string, narration string, txnType TermDepositTransactionType, mode TermDepositTransactionMode, transactedAtUtc time.Time, ) *FinancialAccountTransactionTermDepositTransaction`

NewFinancialAccountTransactionTermDepositTransaction instantiates a new FinancialAccountTransactionTermDepositTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialAccountTransactionTermDepositTransactionWithDefaults

`func NewFinancialAccountTransactionTermDepositTransactionWithDefaults() *FinancialAccountTransactionTermDepositTransaction`

NewFinancialAccountTransactionTermDepositTransactionWithDefaults instantiates a new FinancialAccountTransactionTermDepositTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FinancialAccountTransactionTermDepositTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FinancialAccountTransactionTermDepositTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FinancialAccountTransactionTermDepositTransaction) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *FinancialAccountTransactionTermDepositTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FinancialAccountTransactionTermDepositTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FinancialAccountTransactionTermDepositTransaction) SetId(v string)`

SetId sets Id field to given value.


### GetAmount

`func (o *FinancialAccountTransactionTermDepositTransaction) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FinancialAccountTransactionTermDepositTransaction) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FinancialAccountTransactionTermDepositTransaction) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetCurrencyCode

`func (o *FinancialAccountTransactionTermDepositTransaction) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *FinancialAccountTransactionTermDepositTransaction) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *FinancialAccountTransactionTermDepositTransaction) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetNarration

`func (o *FinancialAccountTransactionTermDepositTransaction) GetNarration() string`

GetNarration returns the Narration field if non-nil, zero value otherwise.

### GetNarrationOk

`func (o *FinancialAccountTransactionTermDepositTransaction) GetNarrationOk() (*string, bool)`

GetNarrationOk returns a tuple with the Narration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNarration

`func (o *FinancialAccountTransactionTermDepositTransaction) SetNarration(v string)`

SetNarration sets Narration field to given value.


### GetTxnType

`func (o *FinancialAccountTransactionTermDepositTransaction) GetTxnType() TermDepositTransactionType`

GetTxnType returns the TxnType field if non-nil, zero value otherwise.

### GetTxnTypeOk

`func (o *FinancialAccountTransactionTermDepositTransaction) GetTxnTypeOk() (*TermDepositTransactionType, bool)`

GetTxnTypeOk returns a tuple with the TxnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnType

`func (o *FinancialAccountTransactionTermDepositTransaction) SetTxnType(v TermDepositTransactionType)`

SetTxnType sets TxnType field to given value.


### GetMode

`func (o *FinancialAccountTransactionTermDepositTransaction) GetMode() TermDepositTransactionMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *FinancialAccountTransactionTermDepositTransaction) GetModeOk() (*TermDepositTransactionMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *FinancialAccountTransactionTermDepositTransaction) SetMode(v TermDepositTransactionMode)`

SetMode sets Mode field to given value.


### GetTransactedAtUtc

`func (o *FinancialAccountTransactionTermDepositTransaction) GetTransactedAtUtc() time.Time`

GetTransactedAtUtc returns the TransactedAtUtc field if non-nil, zero value otherwise.

### GetTransactedAtUtcOk

`func (o *FinancialAccountTransactionTermDepositTransaction) GetTransactedAtUtcOk() (*time.Time, bool)`

GetTransactedAtUtcOk returns a tuple with the TransactedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactedAtUtc

`func (o *FinancialAccountTransactionTermDepositTransaction) SetTransactedAtUtc(v time.Time)`

SetTransactedAtUtc sets TransactedAtUtc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



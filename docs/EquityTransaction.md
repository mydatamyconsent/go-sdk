# EquityTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Amount** | **string** |  | 
**CurrencyCode** | **string** |  | 
**TxnType** | [**EquityTransactionsType**](EquityTransactionsType.md) |  | 
**Units** | **string** |  | 
**TransactedAtUtc** | **time.Time** |  | 

## Methods

### NewEquityTransaction

`func NewEquityTransaction(id string, amount string, currencyCode string, txnType EquityTransactionsType, units string, transactedAtUtc time.Time, ) *EquityTransaction`

NewEquityTransaction instantiates a new EquityTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquityTransactionWithDefaults

`func NewEquityTransactionWithDefaults() *EquityTransaction`

NewEquityTransactionWithDefaults instantiates a new EquityTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EquityTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EquityTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EquityTransaction) SetId(v string)`

SetId sets Id field to given value.


### GetAmount

`func (o *EquityTransaction) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EquityTransaction) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EquityTransaction) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCurrencyCode

`func (o *EquityTransaction) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *EquityTransaction) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *EquityTransaction) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetTxnType

`func (o *EquityTransaction) GetTxnType() EquityTransactionsType`

GetTxnType returns the TxnType field if non-nil, zero value otherwise.

### GetTxnTypeOk

`func (o *EquityTransaction) GetTxnTypeOk() (*EquityTransactionsType, bool)`

GetTxnTypeOk returns a tuple with the TxnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnType

`func (o *EquityTransaction) SetTxnType(v EquityTransactionsType)`

SetTxnType sets TxnType field to given value.


### GetUnits

`func (o *EquityTransaction) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *EquityTransaction) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *EquityTransaction) SetUnits(v string)`

SetUnits sets Units field to given value.


### GetTransactedAtUtc

`func (o *EquityTransaction) GetTransactedAtUtc() time.Time`

GetTransactedAtUtc returns the TransactedAtUtc field if non-nil, zero value otherwise.

### GetTransactedAtUtcOk

`func (o *EquityTransaction) GetTransactedAtUtcOk() (*time.Time, bool)`

GetTransactedAtUtcOk returns a tuple with the TransactedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactedAtUtc

`func (o *EquityTransaction) SetTransactedAtUtc(v time.Time)`

SetTransactedAtUtc sets TransactedAtUtc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



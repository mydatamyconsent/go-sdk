# MutualFundTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Amount** | **string** |  | 
**CurrencyCode** | **string** |  | 
**TxnType** | [**MutualFundTransactionType**](MutualFundTransactionType.md) |  | 
**Units** | **string** |  | 
**TransactedAtUtc** | **time.Time** |  | 

## Methods

### NewMutualFundTransaction

`func NewMutualFundTransaction(id string, amount string, currencyCode string, txnType MutualFundTransactionType, units string, transactedAtUtc time.Time, ) *MutualFundTransaction`

NewMutualFundTransaction instantiates a new MutualFundTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutualFundTransactionWithDefaults

`func NewMutualFundTransactionWithDefaults() *MutualFundTransaction`

NewMutualFundTransactionWithDefaults instantiates a new MutualFundTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MutualFundTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MutualFundTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MutualFundTransaction) SetId(v string)`

SetId sets Id field to given value.


### GetAmount

`func (o *MutualFundTransaction) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *MutualFundTransaction) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *MutualFundTransaction) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCurrencyCode

`func (o *MutualFundTransaction) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *MutualFundTransaction) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *MutualFundTransaction) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetTxnType

`func (o *MutualFundTransaction) GetTxnType() MutualFundTransactionType`

GetTxnType returns the TxnType field if non-nil, zero value otherwise.

### GetTxnTypeOk

`func (o *MutualFundTransaction) GetTxnTypeOk() (*MutualFundTransactionType, bool)`

GetTxnTypeOk returns a tuple with the TxnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnType

`func (o *MutualFundTransaction) SetTxnType(v MutualFundTransactionType)`

SetTxnType sets TxnType field to given value.


### GetUnits

`func (o *MutualFundTransaction) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *MutualFundTransaction) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *MutualFundTransaction) SetUnits(v string)`

SetUnits sets Units field to given value.


### GetTransactedAtUtc

`func (o *MutualFundTransaction) GetTransactedAtUtc() time.Time`

GetTransactedAtUtc returns the TransactedAtUtc field if non-nil, zero value otherwise.

### GetTransactedAtUtcOk

`func (o *MutualFundTransaction) GetTransactedAtUtcOk() (*time.Time, bool)`

GetTransactedAtUtcOk returns a tuple with the TransactedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactedAtUtc

`func (o *MutualFundTransaction) SetTransactedAtUtc(v time.Time)`

SetTransactedAtUtc sets TransactedAtUtc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



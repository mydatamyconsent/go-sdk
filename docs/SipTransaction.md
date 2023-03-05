# SipTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Amount** | **float64** |  | 
**CurrencyCode** | **string** |  | 
**TxnType** | [**SipTransactionType**](SipTransactionType.md) |  | 
**TransactedAtUtc** | **time.Time** |  | 

## Methods

### NewSipTransaction

`func NewSipTransaction(id string, amount float64, currencyCode string, txnType SipTransactionType, transactedAtUtc time.Time, ) *SipTransaction`

NewSipTransaction instantiates a new SipTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSipTransactionWithDefaults

`func NewSipTransactionWithDefaults() *SipTransaction`

NewSipTransactionWithDefaults instantiates a new SipTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SipTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SipTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SipTransaction) SetId(v string)`

SetId sets Id field to given value.


### GetAmount

`func (o *SipTransaction) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SipTransaction) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SipTransaction) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetCurrencyCode

`func (o *SipTransaction) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *SipTransaction) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *SipTransaction) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetTxnType

`func (o *SipTransaction) GetTxnType() SipTransactionType`

GetTxnType returns the TxnType field if non-nil, zero value otherwise.

### GetTxnTypeOk

`func (o *SipTransaction) GetTxnTypeOk() (*SipTransactionType, bool)`

GetTxnTypeOk returns a tuple with the TxnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnType

`func (o *SipTransaction) SetTxnType(v SipTransactionType)`

SetTxnType sets TxnType field to given value.


### GetTransactedAtUtc

`func (o *SipTransaction) GetTransactedAtUtc() time.Time`

GetTransactedAtUtc returns the TransactedAtUtc field if non-nil, zero value otherwise.

### GetTransactedAtUtcOk

`func (o *SipTransaction) GetTransactedAtUtcOk() (*time.Time, bool)`

GetTransactedAtUtcOk returns a tuple with the TransactedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactedAtUtc

`func (o *SipTransaction) SetTransactedAtUtc(v time.Time)`

SetTransactedAtUtc sets TransactedAtUtc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



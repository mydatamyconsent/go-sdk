# CreditCardTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TxnType** | [**CreditCardTransactionType**](CreditCardTransactionType.md) |  | 
**TxnDate** | **time.Time** |  | 
**Amount** | **float32** |  | 
**ValueDate** | **time.Time** |  | 
**Narration** | **string** |  | 
**StatementDate** | **time.Time** |  | 
**Mcc** | **string** |  | 
**MaskedCardNumber** | **string** |  | 

## Methods

### NewCreditCardTransaction

`func NewCreditCardTransaction(id string, txnType CreditCardTransactionType, txnDate time.Time, amount float32, valueDate time.Time, narration string, statementDate time.Time, mcc string, maskedCardNumber string, ) *CreditCardTransaction`

NewCreditCardTransaction instantiates a new CreditCardTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditCardTransactionWithDefaults

`func NewCreditCardTransactionWithDefaults() *CreditCardTransaction`

NewCreditCardTransactionWithDefaults instantiates a new CreditCardTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreditCardTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreditCardTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreditCardTransaction) SetId(v string)`

SetId sets Id field to given value.


### GetTxnType

`func (o *CreditCardTransaction) GetTxnType() CreditCardTransactionType`

GetTxnType returns the TxnType field if non-nil, zero value otherwise.

### GetTxnTypeOk

`func (o *CreditCardTransaction) GetTxnTypeOk() (*CreditCardTransactionType, bool)`

GetTxnTypeOk returns a tuple with the TxnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnType

`func (o *CreditCardTransaction) SetTxnType(v CreditCardTransactionType)`

SetTxnType sets TxnType field to given value.


### GetTxnDate

`func (o *CreditCardTransaction) GetTxnDate() time.Time`

GetTxnDate returns the TxnDate field if non-nil, zero value otherwise.

### GetTxnDateOk

`func (o *CreditCardTransaction) GetTxnDateOk() (*time.Time, bool)`

GetTxnDateOk returns a tuple with the TxnDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnDate

`func (o *CreditCardTransaction) SetTxnDate(v time.Time)`

SetTxnDate sets TxnDate field to given value.


### GetAmount

`func (o *CreditCardTransaction) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreditCardTransaction) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreditCardTransaction) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetValueDate

`func (o *CreditCardTransaction) GetValueDate() time.Time`

GetValueDate returns the ValueDate field if non-nil, zero value otherwise.

### GetValueDateOk

`func (o *CreditCardTransaction) GetValueDateOk() (*time.Time, bool)`

GetValueDateOk returns a tuple with the ValueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueDate

`func (o *CreditCardTransaction) SetValueDate(v time.Time)`

SetValueDate sets ValueDate field to given value.


### GetNarration

`func (o *CreditCardTransaction) GetNarration() string`

GetNarration returns the Narration field if non-nil, zero value otherwise.

### GetNarrationOk

`func (o *CreditCardTransaction) GetNarrationOk() (*string, bool)`

GetNarrationOk returns a tuple with the Narration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNarration

`func (o *CreditCardTransaction) SetNarration(v string)`

SetNarration sets Narration field to given value.


### GetStatementDate

`func (o *CreditCardTransaction) GetStatementDate() time.Time`

GetStatementDate returns the StatementDate field if non-nil, zero value otherwise.

### GetStatementDateOk

`func (o *CreditCardTransaction) GetStatementDateOk() (*time.Time, bool)`

GetStatementDateOk returns a tuple with the StatementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementDate

`func (o *CreditCardTransaction) SetStatementDate(v time.Time)`

SetStatementDate sets StatementDate field to given value.


### GetMcc

`func (o *CreditCardTransaction) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *CreditCardTransaction) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *CreditCardTransaction) SetMcc(v string)`

SetMcc sets Mcc field to given value.


### GetMaskedCardNumber

`func (o *CreditCardTransaction) GetMaskedCardNumber() string`

GetMaskedCardNumber returns the MaskedCardNumber field if non-nil, zero value otherwise.

### GetMaskedCardNumberOk

`func (o *CreditCardTransaction) GetMaskedCardNumberOk() (*string, bool)`

GetMaskedCardNumberOk returns a tuple with the MaskedCardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedCardNumber

`func (o *CreditCardTransaction) SetMaskedCardNumber(v string)`

SetMaskedCardNumber sets MaskedCardNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



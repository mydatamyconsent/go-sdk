# FinancialAccountCreditCardTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
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

### NewFinancialAccountCreditCardTransaction

`func NewFinancialAccountCreditCardTransaction(type_ string, id string, txnType CreditCardTransactionType, txnDate time.Time, amount float32, valueDate time.Time, narration string, statementDate time.Time, mcc string, maskedCardNumber string, ) *FinancialAccountCreditCardTransaction`

NewFinancialAccountCreditCardTransaction instantiates a new FinancialAccountCreditCardTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialAccountCreditCardTransactionWithDefaults

`func NewFinancialAccountCreditCardTransactionWithDefaults() *FinancialAccountCreditCardTransaction`

NewFinancialAccountCreditCardTransactionWithDefaults instantiates a new FinancialAccountCreditCardTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FinancialAccountCreditCardTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FinancialAccountCreditCardTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FinancialAccountCreditCardTransaction) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *FinancialAccountCreditCardTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FinancialAccountCreditCardTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FinancialAccountCreditCardTransaction) SetId(v string)`

SetId sets Id field to given value.


### GetTxnType

`func (o *FinancialAccountCreditCardTransaction) GetTxnType() CreditCardTransactionType`

GetTxnType returns the TxnType field if non-nil, zero value otherwise.

### GetTxnTypeOk

`func (o *FinancialAccountCreditCardTransaction) GetTxnTypeOk() (*CreditCardTransactionType, bool)`

GetTxnTypeOk returns a tuple with the TxnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnType

`func (o *FinancialAccountCreditCardTransaction) SetTxnType(v CreditCardTransactionType)`

SetTxnType sets TxnType field to given value.


### GetTxnDate

`func (o *FinancialAccountCreditCardTransaction) GetTxnDate() time.Time`

GetTxnDate returns the TxnDate field if non-nil, zero value otherwise.

### GetTxnDateOk

`func (o *FinancialAccountCreditCardTransaction) GetTxnDateOk() (*time.Time, bool)`

GetTxnDateOk returns a tuple with the TxnDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnDate

`func (o *FinancialAccountCreditCardTransaction) SetTxnDate(v time.Time)`

SetTxnDate sets TxnDate field to given value.


### GetAmount

`func (o *FinancialAccountCreditCardTransaction) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FinancialAccountCreditCardTransaction) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FinancialAccountCreditCardTransaction) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetValueDate

`func (o *FinancialAccountCreditCardTransaction) GetValueDate() time.Time`

GetValueDate returns the ValueDate field if non-nil, zero value otherwise.

### GetValueDateOk

`func (o *FinancialAccountCreditCardTransaction) GetValueDateOk() (*time.Time, bool)`

GetValueDateOk returns a tuple with the ValueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueDate

`func (o *FinancialAccountCreditCardTransaction) SetValueDate(v time.Time)`

SetValueDate sets ValueDate field to given value.


### GetNarration

`func (o *FinancialAccountCreditCardTransaction) GetNarration() string`

GetNarration returns the Narration field if non-nil, zero value otherwise.

### GetNarrationOk

`func (o *FinancialAccountCreditCardTransaction) GetNarrationOk() (*string, bool)`

GetNarrationOk returns a tuple with the Narration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNarration

`func (o *FinancialAccountCreditCardTransaction) SetNarration(v string)`

SetNarration sets Narration field to given value.


### GetStatementDate

`func (o *FinancialAccountCreditCardTransaction) GetStatementDate() time.Time`

GetStatementDate returns the StatementDate field if non-nil, zero value otherwise.

### GetStatementDateOk

`func (o *FinancialAccountCreditCardTransaction) GetStatementDateOk() (*time.Time, bool)`

GetStatementDateOk returns a tuple with the StatementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementDate

`func (o *FinancialAccountCreditCardTransaction) SetStatementDate(v time.Time)`

SetStatementDate sets StatementDate field to given value.


### GetMcc

`func (o *FinancialAccountCreditCardTransaction) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *FinancialAccountCreditCardTransaction) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *FinancialAccountCreditCardTransaction) SetMcc(v string)`

SetMcc sets Mcc field to given value.


### GetMaskedCardNumber

`func (o *FinancialAccountCreditCardTransaction) GetMaskedCardNumber() string`

GetMaskedCardNumber returns the MaskedCardNumber field if non-nil, zero value otherwise.

### GetMaskedCardNumberOk

`func (o *FinancialAccountCreditCardTransaction) GetMaskedCardNumberOk() (*string, bool)`

GetMaskedCardNumberOk returns a tuple with the MaskedCardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedCardNumber

`func (o *FinancialAccountCreditCardTransaction) SetMaskedCardNumber(v string)`

SetMaskedCardNumber sets MaskedCardNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UserAccountFinancialTransactionsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**InstrumentId** | Pointer to **string** |  | [optional] 
**TransactionType** | Pointer to **NullableString** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**AveragePrice** | Pointer to **float64** |  | [optional] 
**Currency** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUserAccountFinancialTransactionsDto

`func NewUserAccountFinancialTransactionsDto() *UserAccountFinancialTransactionsDto`

NewUserAccountFinancialTransactionsDto instantiates a new UserAccountFinancialTransactionsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAccountFinancialTransactionsDtoWithDefaults

`func NewUserAccountFinancialTransactionsDtoWithDefaults() *UserAccountFinancialTransactionsDto`

NewUserAccountFinancialTransactionsDtoWithDefaults instantiates a new UserAccountFinancialTransactionsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserAccountFinancialTransactionsDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserAccountFinancialTransactionsDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserAccountFinancialTransactionsDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserAccountFinancialTransactionsDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *UserAccountFinancialTransactionsDto) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UserAccountFinancialTransactionsDto) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UserAccountFinancialTransactionsDto) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UserAccountFinancialTransactionsDto) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetName

`func (o *UserAccountFinancialTransactionsDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserAccountFinancialTransactionsDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserAccountFinancialTransactionsDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserAccountFinancialTransactionsDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UserAccountFinancialTransactionsDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UserAccountFinancialTransactionsDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetInstrumentId

`func (o *UserAccountFinancialTransactionsDto) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *UserAccountFinancialTransactionsDto) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *UserAccountFinancialTransactionsDto) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *UserAccountFinancialTransactionsDto) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetTransactionType

`func (o *UserAccountFinancialTransactionsDto) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *UserAccountFinancialTransactionsDto) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *UserAccountFinancialTransactionsDto) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *UserAccountFinancialTransactionsDto) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### SetTransactionTypeNil

`func (o *UserAccountFinancialTransactionsDto) SetTransactionTypeNil(b bool)`

 SetTransactionTypeNil sets the value for TransactionType to be an explicit nil

### UnsetTransactionType
`func (o *UserAccountFinancialTransactionsDto) UnsetTransactionType()`

UnsetTransactionType ensures that no value is present for TransactionType, not even an explicit nil
### GetQuantity

`func (o *UserAccountFinancialTransactionsDto) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UserAccountFinancialTransactionsDto) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UserAccountFinancialTransactionsDto) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UserAccountFinancialTransactionsDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetAveragePrice

`func (o *UserAccountFinancialTransactionsDto) GetAveragePrice() float64`

GetAveragePrice returns the AveragePrice field if non-nil, zero value otherwise.

### GetAveragePriceOk

`func (o *UserAccountFinancialTransactionsDto) GetAveragePriceOk() (*float64, bool)`

GetAveragePriceOk returns a tuple with the AveragePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAveragePrice

`func (o *UserAccountFinancialTransactionsDto) SetAveragePrice(v float64)`

SetAveragePrice sets AveragePrice field to given value.

### HasAveragePrice

`func (o *UserAccountFinancialTransactionsDto) HasAveragePrice() bool`

HasAveragePrice returns a boolean if a field has been set.

### GetCurrency

`func (o *UserAccountFinancialTransactionsDto) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UserAccountFinancialTransactionsDto) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UserAccountFinancialTransactionsDto) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UserAccountFinancialTransactionsDto) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *UserAccountFinancialTransactionsDto) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *UserAccountFinancialTransactionsDto) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



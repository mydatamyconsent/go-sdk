# OrganizationFinancialTransactionsDto

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

### NewOrganizationFinancialTransactionsDto

`func NewOrganizationFinancialTransactionsDto() *OrganizationFinancialTransactionsDto`

NewOrganizationFinancialTransactionsDto instantiates a new OrganizationFinancialTransactionsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationFinancialTransactionsDtoWithDefaults

`func NewOrganizationFinancialTransactionsDtoWithDefaults() *OrganizationFinancialTransactionsDto`

NewOrganizationFinancialTransactionsDtoWithDefaults instantiates a new OrganizationFinancialTransactionsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationFinancialTransactionsDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationFinancialTransactionsDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationFinancialTransactionsDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationFinancialTransactionsDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *OrganizationFinancialTransactionsDto) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *OrganizationFinancialTransactionsDto) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *OrganizationFinancialTransactionsDto) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *OrganizationFinancialTransactionsDto) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetName

`func (o *OrganizationFinancialTransactionsDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationFinancialTransactionsDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationFinancialTransactionsDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationFinancialTransactionsDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OrganizationFinancialTransactionsDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OrganizationFinancialTransactionsDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetInstrumentId

`func (o *OrganizationFinancialTransactionsDto) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *OrganizationFinancialTransactionsDto) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *OrganizationFinancialTransactionsDto) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *OrganizationFinancialTransactionsDto) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetTransactionType

`func (o *OrganizationFinancialTransactionsDto) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *OrganizationFinancialTransactionsDto) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *OrganizationFinancialTransactionsDto) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *OrganizationFinancialTransactionsDto) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### SetTransactionTypeNil

`func (o *OrganizationFinancialTransactionsDto) SetTransactionTypeNil(b bool)`

 SetTransactionTypeNil sets the value for TransactionType to be an explicit nil

### UnsetTransactionType
`func (o *OrganizationFinancialTransactionsDto) UnsetTransactionType()`

UnsetTransactionType ensures that no value is present for TransactionType, not even an explicit nil
### GetQuantity

`func (o *OrganizationFinancialTransactionsDto) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrganizationFinancialTransactionsDto) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrganizationFinancialTransactionsDto) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrganizationFinancialTransactionsDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetAveragePrice

`func (o *OrganizationFinancialTransactionsDto) GetAveragePrice() float64`

GetAveragePrice returns the AveragePrice field if non-nil, zero value otherwise.

### GetAveragePriceOk

`func (o *OrganizationFinancialTransactionsDto) GetAveragePriceOk() (*float64, bool)`

GetAveragePriceOk returns a tuple with the AveragePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAveragePrice

`func (o *OrganizationFinancialTransactionsDto) SetAveragePrice(v float64)`

SetAveragePrice sets AveragePrice field to given value.

### HasAveragePrice

`func (o *OrganizationFinancialTransactionsDto) HasAveragePrice() bool`

HasAveragePrice returns a boolean if a field has been set.

### GetCurrency

`func (o *OrganizationFinancialTransactionsDto) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *OrganizationFinancialTransactionsDto) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *OrganizationFinancialTransactionsDto) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *OrganizationFinancialTransactionsDto) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *OrganizationFinancialTransactionsDto) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *OrganizationFinancialTransactionsDto) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



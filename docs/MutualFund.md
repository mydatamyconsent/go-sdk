# MutualFund

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**InvestmentValue** | **float64** |  | 
**CurrentValue** | **float64** |  | 
**CurrencyCode** | **string** |  | 
**Amc** | Pointer to **string** |  | [optional] 
**Registrar** | Pointer to **string** |  | [optional] 
**FundName** | **string** |  | 
**Isin** | **string** |  | 
**FolioNumber** | **string** |  | 
**SchemeCode** | Pointer to **string** |  | [optional] 
**FundType** | Pointer to **string** |  | [optional] 
**FundCategory** | Pointer to **string** |  | [optional] 
**Units** | **float64** |  | 
**LienUnits** | Pointer to **string** |  | [optional] 
**CreationDate** | Pointer to **time.Time** |  | [optional] 
**Holder** | [**Holder**](Holder.md) |  | 
**Transactions** | **bool** |  | 

## Methods

### NewMutualFund

`func NewMutualFund(id string, name string, investmentValue float64, currentValue float64, currencyCode string, fundName string, isin string, folioNumber string, units float64, holder Holder, transactions bool, ) *MutualFund`

NewMutualFund instantiates a new MutualFund object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutualFundWithDefaults

`func NewMutualFundWithDefaults() *MutualFund`

NewMutualFundWithDefaults instantiates a new MutualFund object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MutualFund) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MutualFund) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MutualFund) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *MutualFund) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MutualFund) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MutualFund) SetName(v string)`

SetName sets Name field to given value.


### GetInvestmentValue

`func (o *MutualFund) GetInvestmentValue() float64`

GetInvestmentValue returns the InvestmentValue field if non-nil, zero value otherwise.

### GetInvestmentValueOk

`func (o *MutualFund) GetInvestmentValueOk() (*float64, bool)`

GetInvestmentValueOk returns a tuple with the InvestmentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentValue

`func (o *MutualFund) SetInvestmentValue(v float64)`

SetInvestmentValue sets InvestmentValue field to given value.


### GetCurrentValue

`func (o *MutualFund) GetCurrentValue() float64`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *MutualFund) GetCurrentValueOk() (*float64, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *MutualFund) SetCurrentValue(v float64)`

SetCurrentValue sets CurrentValue field to given value.


### GetCurrencyCode

`func (o *MutualFund) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *MutualFund) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *MutualFund) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetAmc

`func (o *MutualFund) GetAmc() string`

GetAmc returns the Amc field if non-nil, zero value otherwise.

### GetAmcOk

`func (o *MutualFund) GetAmcOk() (*string, bool)`

GetAmcOk returns a tuple with the Amc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmc

`func (o *MutualFund) SetAmc(v string)`

SetAmc sets Amc field to given value.

### HasAmc

`func (o *MutualFund) HasAmc() bool`

HasAmc returns a boolean if a field has been set.

### GetRegistrar

`func (o *MutualFund) GetRegistrar() string`

GetRegistrar returns the Registrar field if non-nil, zero value otherwise.

### GetRegistrarOk

`func (o *MutualFund) GetRegistrarOk() (*string, bool)`

GetRegistrarOk returns a tuple with the Registrar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrar

`func (o *MutualFund) SetRegistrar(v string)`

SetRegistrar sets Registrar field to given value.

### HasRegistrar

`func (o *MutualFund) HasRegistrar() bool`

HasRegistrar returns a boolean if a field has been set.

### GetFundName

`func (o *MutualFund) GetFundName() string`

GetFundName returns the FundName field if non-nil, zero value otherwise.

### GetFundNameOk

`func (o *MutualFund) GetFundNameOk() (*string, bool)`

GetFundNameOk returns a tuple with the FundName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundName

`func (o *MutualFund) SetFundName(v string)`

SetFundName sets FundName field to given value.


### GetIsin

`func (o *MutualFund) GetIsin() string`

GetIsin returns the Isin field if non-nil, zero value otherwise.

### GetIsinOk

`func (o *MutualFund) GetIsinOk() (*string, bool)`

GetIsinOk returns a tuple with the Isin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsin

`func (o *MutualFund) SetIsin(v string)`

SetIsin sets Isin field to given value.


### GetFolioNumber

`func (o *MutualFund) GetFolioNumber() string`

GetFolioNumber returns the FolioNumber field if non-nil, zero value otherwise.

### GetFolioNumberOk

`func (o *MutualFund) GetFolioNumberOk() (*string, bool)`

GetFolioNumberOk returns a tuple with the FolioNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolioNumber

`func (o *MutualFund) SetFolioNumber(v string)`

SetFolioNumber sets FolioNumber field to given value.


### GetSchemeCode

`func (o *MutualFund) GetSchemeCode() string`

GetSchemeCode returns the SchemeCode field if non-nil, zero value otherwise.

### GetSchemeCodeOk

`func (o *MutualFund) GetSchemeCodeOk() (*string, bool)`

GetSchemeCodeOk returns a tuple with the SchemeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeCode

`func (o *MutualFund) SetSchemeCode(v string)`

SetSchemeCode sets SchemeCode field to given value.

### HasSchemeCode

`func (o *MutualFund) HasSchemeCode() bool`

HasSchemeCode returns a boolean if a field has been set.

### GetFundType

`func (o *MutualFund) GetFundType() string`

GetFundType returns the FundType field if non-nil, zero value otherwise.

### GetFundTypeOk

`func (o *MutualFund) GetFundTypeOk() (*string, bool)`

GetFundTypeOk returns a tuple with the FundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundType

`func (o *MutualFund) SetFundType(v string)`

SetFundType sets FundType field to given value.

### HasFundType

`func (o *MutualFund) HasFundType() bool`

HasFundType returns a boolean if a field has been set.

### GetFundCategory

`func (o *MutualFund) GetFundCategory() string`

GetFundCategory returns the FundCategory field if non-nil, zero value otherwise.

### GetFundCategoryOk

`func (o *MutualFund) GetFundCategoryOk() (*string, bool)`

GetFundCategoryOk returns a tuple with the FundCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundCategory

`func (o *MutualFund) SetFundCategory(v string)`

SetFundCategory sets FundCategory field to given value.

### HasFundCategory

`func (o *MutualFund) HasFundCategory() bool`

HasFundCategory returns a boolean if a field has been set.

### GetUnits

`func (o *MutualFund) GetUnits() float64`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *MutualFund) GetUnitsOk() (*float64, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *MutualFund) SetUnits(v float64)`

SetUnits sets Units field to given value.


### GetLienUnits

`func (o *MutualFund) GetLienUnits() string`

GetLienUnits returns the LienUnits field if non-nil, zero value otherwise.

### GetLienUnitsOk

`func (o *MutualFund) GetLienUnitsOk() (*string, bool)`

GetLienUnitsOk returns a tuple with the LienUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLienUnits

`func (o *MutualFund) SetLienUnits(v string)`

SetLienUnits sets LienUnits field to given value.

### HasLienUnits

`func (o *MutualFund) HasLienUnits() bool`

HasLienUnits returns a boolean if a field has been set.

### GetCreationDate

`func (o *MutualFund) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *MutualFund) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *MutualFund) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *MutualFund) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetHolder

`func (o *MutualFund) GetHolder() Holder`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *MutualFund) GetHolderOk() (*Holder, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *MutualFund) SetHolder(v Holder)`

SetHolder sets Holder field to given value.


### GetTransactions

`func (o *MutualFund) GetTransactions() bool`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *MutualFund) GetTransactionsOk() (*bool, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *MutualFund) SetTransactions(v bool)`

SetTransactions sets Transactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



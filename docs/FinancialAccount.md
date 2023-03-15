# FinancialAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**IssuerName** | **string** |  | 
**Exchange** | **string** |  | 
**Isin** | **string** |  | 
**Units** | **float64** |  | 
**InvestmentValue** | **float64** |  | 
**CurrentValue** | **float64** |  | 
**CurrencyCode** | **string** |  | 
**Holder** | [**Holder**](Holder.md) |  | 
**Transactions** | **bool** |  | 
**Amc** | Pointer to **string** |  | [optional] 
**Registrar** | Pointer to **string** |  | [optional] 
**FundName** | **string** |  | 
**FolioNumber** | **string** |  | 
**SchemeCode** | Pointer to **string** |  | [optional] 
**FundType** | Pointer to **string** |  | [optional] 
**FundCategory** | Pointer to **string** |  | [optional] 
**LienUnits** | Pointer to **string** |  | [optional] 
**CreationDate** | Pointer to **time.Time** |  | [optional] 
**PlanInfo** | [**SipPlanInformation**](SipPlanInformation.md) |  | 
**InvestmentInfo** | [**SipInvestmentInformation**](SipInvestmentInformation.md) |  | 
**AccountNumber** | **string** |  | 
**AccountType** | [**TermDepositAccountType**](TermDepositAccountType.md) |  | 
**IssuerLogoUrl** | **string** |  | 
**AccountDetails** | [**TermDepositAccountDetails**](TermDepositAccountDetails.md) |  | 

## Methods

### NewFinancialAccount

`func NewFinancialAccount(type_ string, id string, name string, issuerName string, exchange string, isin string, units float64, investmentValue float64, currentValue float64, currencyCode string, holder Holder, transactions bool, fundName string, folioNumber string, planInfo SipPlanInformation, investmentInfo SipInvestmentInformation, accountNumber string, accountType TermDepositAccountType, issuerLogoUrl string, accountDetails TermDepositAccountDetails, ) *FinancialAccount`

NewFinancialAccount instantiates a new FinancialAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialAccountWithDefaults

`func NewFinancialAccountWithDefaults() *FinancialAccount`

NewFinancialAccountWithDefaults instantiates a new FinancialAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FinancialAccount) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FinancialAccount) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FinancialAccount) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *FinancialAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FinancialAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FinancialAccount) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *FinancialAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FinancialAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FinancialAccount) SetName(v string)`

SetName sets Name field to given value.


### GetIssuerName

`func (o *FinancialAccount) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *FinancialAccount) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *FinancialAccount) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.


### GetExchange

`func (o *FinancialAccount) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *FinancialAccount) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *FinancialAccount) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetIsin

`func (o *FinancialAccount) GetIsin() string`

GetIsin returns the Isin field if non-nil, zero value otherwise.

### GetIsinOk

`func (o *FinancialAccount) GetIsinOk() (*string, bool)`

GetIsinOk returns a tuple with the Isin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsin

`func (o *FinancialAccount) SetIsin(v string)`

SetIsin sets Isin field to given value.


### GetUnits

`func (o *FinancialAccount) GetUnits() float64`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *FinancialAccount) GetUnitsOk() (*float64, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *FinancialAccount) SetUnits(v float64)`

SetUnits sets Units field to given value.


### GetInvestmentValue

`func (o *FinancialAccount) GetInvestmentValue() float64`

GetInvestmentValue returns the InvestmentValue field if non-nil, zero value otherwise.

### GetInvestmentValueOk

`func (o *FinancialAccount) GetInvestmentValueOk() (*float64, bool)`

GetInvestmentValueOk returns a tuple with the InvestmentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentValue

`func (o *FinancialAccount) SetInvestmentValue(v float64)`

SetInvestmentValue sets InvestmentValue field to given value.


### GetCurrentValue

`func (o *FinancialAccount) GetCurrentValue() float64`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *FinancialAccount) GetCurrentValueOk() (*float64, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *FinancialAccount) SetCurrentValue(v float64)`

SetCurrentValue sets CurrentValue field to given value.


### GetCurrencyCode

`func (o *FinancialAccount) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *FinancialAccount) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *FinancialAccount) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetHolder

`func (o *FinancialAccount) GetHolder() Holder`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *FinancialAccount) GetHolderOk() (*Holder, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *FinancialAccount) SetHolder(v Holder)`

SetHolder sets Holder field to given value.


### GetTransactions

`func (o *FinancialAccount) GetTransactions() bool`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *FinancialAccount) GetTransactionsOk() (*bool, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *FinancialAccount) SetTransactions(v bool)`

SetTransactions sets Transactions field to given value.


### GetAmc

`func (o *FinancialAccount) GetAmc() string`

GetAmc returns the Amc field if non-nil, zero value otherwise.

### GetAmcOk

`func (o *FinancialAccount) GetAmcOk() (*string, bool)`

GetAmcOk returns a tuple with the Amc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmc

`func (o *FinancialAccount) SetAmc(v string)`

SetAmc sets Amc field to given value.

### HasAmc

`func (o *FinancialAccount) HasAmc() bool`

HasAmc returns a boolean if a field has been set.

### GetRegistrar

`func (o *FinancialAccount) GetRegistrar() string`

GetRegistrar returns the Registrar field if non-nil, zero value otherwise.

### GetRegistrarOk

`func (o *FinancialAccount) GetRegistrarOk() (*string, bool)`

GetRegistrarOk returns a tuple with the Registrar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrar

`func (o *FinancialAccount) SetRegistrar(v string)`

SetRegistrar sets Registrar field to given value.

### HasRegistrar

`func (o *FinancialAccount) HasRegistrar() bool`

HasRegistrar returns a boolean if a field has been set.

### GetFundName

`func (o *FinancialAccount) GetFundName() string`

GetFundName returns the FundName field if non-nil, zero value otherwise.

### GetFundNameOk

`func (o *FinancialAccount) GetFundNameOk() (*string, bool)`

GetFundNameOk returns a tuple with the FundName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundName

`func (o *FinancialAccount) SetFundName(v string)`

SetFundName sets FundName field to given value.


### GetFolioNumber

`func (o *FinancialAccount) GetFolioNumber() string`

GetFolioNumber returns the FolioNumber field if non-nil, zero value otherwise.

### GetFolioNumberOk

`func (o *FinancialAccount) GetFolioNumberOk() (*string, bool)`

GetFolioNumberOk returns a tuple with the FolioNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolioNumber

`func (o *FinancialAccount) SetFolioNumber(v string)`

SetFolioNumber sets FolioNumber field to given value.


### GetSchemeCode

`func (o *FinancialAccount) GetSchemeCode() string`

GetSchemeCode returns the SchemeCode field if non-nil, zero value otherwise.

### GetSchemeCodeOk

`func (o *FinancialAccount) GetSchemeCodeOk() (*string, bool)`

GetSchemeCodeOk returns a tuple with the SchemeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeCode

`func (o *FinancialAccount) SetSchemeCode(v string)`

SetSchemeCode sets SchemeCode field to given value.

### HasSchemeCode

`func (o *FinancialAccount) HasSchemeCode() bool`

HasSchemeCode returns a boolean if a field has been set.

### GetFundType

`func (o *FinancialAccount) GetFundType() string`

GetFundType returns the FundType field if non-nil, zero value otherwise.

### GetFundTypeOk

`func (o *FinancialAccount) GetFundTypeOk() (*string, bool)`

GetFundTypeOk returns a tuple with the FundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundType

`func (o *FinancialAccount) SetFundType(v string)`

SetFundType sets FundType field to given value.

### HasFundType

`func (o *FinancialAccount) HasFundType() bool`

HasFundType returns a boolean if a field has been set.

### GetFundCategory

`func (o *FinancialAccount) GetFundCategory() string`

GetFundCategory returns the FundCategory field if non-nil, zero value otherwise.

### GetFundCategoryOk

`func (o *FinancialAccount) GetFundCategoryOk() (*string, bool)`

GetFundCategoryOk returns a tuple with the FundCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundCategory

`func (o *FinancialAccount) SetFundCategory(v string)`

SetFundCategory sets FundCategory field to given value.

### HasFundCategory

`func (o *FinancialAccount) HasFundCategory() bool`

HasFundCategory returns a boolean if a field has been set.

### GetLienUnits

`func (o *FinancialAccount) GetLienUnits() string`

GetLienUnits returns the LienUnits field if non-nil, zero value otherwise.

### GetLienUnitsOk

`func (o *FinancialAccount) GetLienUnitsOk() (*string, bool)`

GetLienUnitsOk returns a tuple with the LienUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLienUnits

`func (o *FinancialAccount) SetLienUnits(v string)`

SetLienUnits sets LienUnits field to given value.

### HasLienUnits

`func (o *FinancialAccount) HasLienUnits() bool`

HasLienUnits returns a boolean if a field has been set.

### GetCreationDate

`func (o *FinancialAccount) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *FinancialAccount) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *FinancialAccount) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *FinancialAccount) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetPlanInfo

`func (o *FinancialAccount) GetPlanInfo() SipPlanInformation`

GetPlanInfo returns the PlanInfo field if non-nil, zero value otherwise.

### GetPlanInfoOk

`func (o *FinancialAccount) GetPlanInfoOk() (*SipPlanInformation, bool)`

GetPlanInfoOk returns a tuple with the PlanInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanInfo

`func (o *FinancialAccount) SetPlanInfo(v SipPlanInformation)`

SetPlanInfo sets PlanInfo field to given value.


### GetInvestmentInfo

`func (o *FinancialAccount) GetInvestmentInfo() SipInvestmentInformation`

GetInvestmentInfo returns the InvestmentInfo field if non-nil, zero value otherwise.

### GetInvestmentInfoOk

`func (o *FinancialAccount) GetInvestmentInfoOk() (*SipInvestmentInformation, bool)`

GetInvestmentInfoOk returns a tuple with the InvestmentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentInfo

`func (o *FinancialAccount) SetInvestmentInfo(v SipInvestmentInformation)`

SetInvestmentInfo sets InvestmentInfo field to given value.


### GetAccountNumber

`func (o *FinancialAccount) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *FinancialAccount) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *FinancialAccount) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetAccountType

`func (o *FinancialAccount) GetAccountType() TermDepositAccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *FinancialAccount) GetAccountTypeOk() (*TermDepositAccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *FinancialAccount) SetAccountType(v TermDepositAccountType)`

SetAccountType sets AccountType field to given value.


### GetIssuerLogoUrl

`func (o *FinancialAccount) GetIssuerLogoUrl() string`

GetIssuerLogoUrl returns the IssuerLogoUrl field if non-nil, zero value otherwise.

### GetIssuerLogoUrlOk

`func (o *FinancialAccount) GetIssuerLogoUrlOk() (*string, bool)`

GetIssuerLogoUrlOk returns a tuple with the IssuerLogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerLogoUrl

`func (o *FinancialAccount) SetIssuerLogoUrl(v string)`

SetIssuerLogoUrl sets IssuerLogoUrl field to given value.


### GetAccountDetails

`func (o *FinancialAccount) GetAccountDetails() TermDepositAccountDetails`

GetAccountDetails returns the AccountDetails field if non-nil, zero value otherwise.

### GetAccountDetailsOk

`func (o *FinancialAccount) GetAccountDetailsOk() (*TermDepositAccountDetails, bool)`

GetAccountDetailsOk returns a tuple with the AccountDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDetails

`func (o *FinancialAccount) SetAccountDetails(v TermDepositAccountDetails)`

SetAccountDetails sets AccountDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# FinancialAccountTermDeposit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**AccountNumber** | **string** |  | 
**AccountType** | [**TermDepositAccountType**](TermDepositAccountType.md) |  | 
**IssuerName** | **string** |  | 
**IssuerLogoUrl** | **string** |  | 
**CurrentValue** | **float64** |  | 
**CurrencyCode** | **string** |  | 
**AccountDetails** | [**TermDepositAccountDetails**](TermDepositAccountDetails.md) |  | 

## Methods

### NewFinancialAccountTermDeposit

`func NewFinancialAccountTermDeposit(type_ string, accountNumber string, accountType TermDepositAccountType, issuerName string, issuerLogoUrl string, currentValue float64, currencyCode string, accountDetails TermDepositAccountDetails, ) *FinancialAccountTermDeposit`

NewFinancialAccountTermDeposit instantiates a new FinancialAccountTermDeposit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialAccountTermDepositWithDefaults

`func NewFinancialAccountTermDepositWithDefaults() *FinancialAccountTermDeposit`

NewFinancialAccountTermDepositWithDefaults instantiates a new FinancialAccountTermDeposit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FinancialAccountTermDeposit) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FinancialAccountTermDeposit) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FinancialAccountTermDeposit) SetType(v string)`

SetType sets Type field to given value.


### GetAccountNumber

`func (o *FinancialAccountTermDeposit) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *FinancialAccountTermDeposit) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *FinancialAccountTermDeposit) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetAccountType

`func (o *FinancialAccountTermDeposit) GetAccountType() TermDepositAccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *FinancialAccountTermDeposit) GetAccountTypeOk() (*TermDepositAccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *FinancialAccountTermDeposit) SetAccountType(v TermDepositAccountType)`

SetAccountType sets AccountType field to given value.


### GetIssuerName

`func (o *FinancialAccountTermDeposit) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *FinancialAccountTermDeposit) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *FinancialAccountTermDeposit) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.


### GetIssuerLogoUrl

`func (o *FinancialAccountTermDeposit) GetIssuerLogoUrl() string`

GetIssuerLogoUrl returns the IssuerLogoUrl field if non-nil, zero value otherwise.

### GetIssuerLogoUrlOk

`func (o *FinancialAccountTermDeposit) GetIssuerLogoUrlOk() (*string, bool)`

GetIssuerLogoUrlOk returns a tuple with the IssuerLogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerLogoUrl

`func (o *FinancialAccountTermDeposit) SetIssuerLogoUrl(v string)`

SetIssuerLogoUrl sets IssuerLogoUrl field to given value.


### GetCurrentValue

`func (o *FinancialAccountTermDeposit) GetCurrentValue() float64`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *FinancialAccountTermDeposit) GetCurrentValueOk() (*float64, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *FinancialAccountTermDeposit) SetCurrentValue(v float64)`

SetCurrentValue sets CurrentValue field to given value.


### GetCurrencyCode

`func (o *FinancialAccountTermDeposit) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *FinancialAccountTermDeposit) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *FinancialAccountTermDeposit) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetAccountDetails

`func (o *FinancialAccountTermDeposit) GetAccountDetails() TermDepositAccountDetails`

GetAccountDetails returns the AccountDetails field if non-nil, zero value otherwise.

### GetAccountDetailsOk

`func (o *FinancialAccountTermDeposit) GetAccountDetailsOk() (*TermDepositAccountDetails, bool)`

GetAccountDetailsOk returns a tuple with the AccountDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDetails

`func (o *FinancialAccountTermDeposit) SetAccountDetails(v TermDepositAccountDetails)`

SetAccountDetails sets AccountDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



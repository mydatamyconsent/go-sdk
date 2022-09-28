# FinancialAccountTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**TxnType** | [**CreditCardTransactionType**](CreditCardTransactionType.md) |  | 
**TxnDate** | **time.Time** |  | 
**Amount** | **string** |  | 
**ValueDate** | **time.Time** |  | 
**Narration** | **string** |  | 
**StatementDate** | **time.Time** |  | 
**Mcc** | **string** |  | 
**MaskedCardNumber** | **string** |  | 
**Amc** | **string** |  | 
**Registrar** | **string** |  | 
**SchemeCode** | **string** |  | 
**SchemePlan** | [**MutualFundSchemePlan**](MutualFundSchemePlan.md) |  | 
**Isin** | **string** |  | 
**AmfiCode** | **string** |  | 
**FundType** | [**MutualFundFundType**](MutualFundFundType.md) |  | 
**SchemeOption** | [**MutualFundSchemeOption**](MutualFundSchemeOption.md) |  | 
**SchemeTypes** | [**MutualFundSchemeType**](MutualFundSchemeType.md) |  | 
**SchemeCategory** | [**MutualFundSchemeCategory**](MutualFundSchemeCategory.md) |  | 
**Ucc** | **string** |  | 
**ClosingUnits** | **string** |  | 
**LienUnits** | **string** |  | 
**Nav** | **string** |  | 
**NavDate** | **time.Time** |  | 
**OrderDate** | **time.Time** |  | 
**ExecutionDate** | **time.Time** |  | 
**LockinFlag** | **string** |  | 
**LockinDays** | **string** |  | 
**Mode** | [**MutualFundHoldingMode**](MutualFundHoldingMode.md) |  | 

## Methods

### NewFinancialAccountTransaction

`func NewFinancialAccountTransaction(type_ string, id string, txnType CreditCardTransactionType, txnDate time.Time, amount string, valueDate time.Time, narration string, statementDate time.Time, mcc string, maskedCardNumber string, amc string, registrar string, schemeCode string, schemePlan MutualFundSchemePlan, isin string, amfiCode string, fundType MutualFundFundType, schemeOption MutualFundSchemeOption, schemeTypes MutualFundSchemeType, schemeCategory MutualFundSchemeCategory, ucc string, closingUnits string, lienUnits string, nav string, navDate time.Time, orderDate time.Time, executionDate time.Time, lockinFlag string, lockinDays string, mode MutualFundHoldingMode, ) *FinancialAccountTransaction`

NewFinancialAccountTransaction instantiates a new FinancialAccountTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialAccountTransactionWithDefaults

`func NewFinancialAccountTransactionWithDefaults() *FinancialAccountTransaction`

NewFinancialAccountTransactionWithDefaults instantiates a new FinancialAccountTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FinancialAccountTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FinancialAccountTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FinancialAccountTransaction) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *FinancialAccountTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FinancialAccountTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FinancialAccountTransaction) SetId(v string)`

SetId sets Id field to given value.


### GetTxnType

`func (o *FinancialAccountTransaction) GetTxnType() CreditCardTransactionType`

GetTxnType returns the TxnType field if non-nil, zero value otherwise.

### GetTxnTypeOk

`func (o *FinancialAccountTransaction) GetTxnTypeOk() (*CreditCardTransactionType, bool)`

GetTxnTypeOk returns a tuple with the TxnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnType

`func (o *FinancialAccountTransaction) SetTxnType(v CreditCardTransactionType)`

SetTxnType sets TxnType field to given value.


### GetTxnDate

`func (o *FinancialAccountTransaction) GetTxnDate() time.Time`

GetTxnDate returns the TxnDate field if non-nil, zero value otherwise.

### GetTxnDateOk

`func (o *FinancialAccountTransaction) GetTxnDateOk() (*time.Time, bool)`

GetTxnDateOk returns a tuple with the TxnDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnDate

`func (o *FinancialAccountTransaction) SetTxnDate(v time.Time)`

SetTxnDate sets TxnDate field to given value.


### GetAmount

`func (o *FinancialAccountTransaction) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FinancialAccountTransaction) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FinancialAccountTransaction) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetValueDate

`func (o *FinancialAccountTransaction) GetValueDate() time.Time`

GetValueDate returns the ValueDate field if non-nil, zero value otherwise.

### GetValueDateOk

`func (o *FinancialAccountTransaction) GetValueDateOk() (*time.Time, bool)`

GetValueDateOk returns a tuple with the ValueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueDate

`func (o *FinancialAccountTransaction) SetValueDate(v time.Time)`

SetValueDate sets ValueDate field to given value.


### GetNarration

`func (o *FinancialAccountTransaction) GetNarration() string`

GetNarration returns the Narration field if non-nil, zero value otherwise.

### GetNarrationOk

`func (o *FinancialAccountTransaction) GetNarrationOk() (*string, bool)`

GetNarrationOk returns a tuple with the Narration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNarration

`func (o *FinancialAccountTransaction) SetNarration(v string)`

SetNarration sets Narration field to given value.


### GetStatementDate

`func (o *FinancialAccountTransaction) GetStatementDate() time.Time`

GetStatementDate returns the StatementDate field if non-nil, zero value otherwise.

### GetStatementDateOk

`func (o *FinancialAccountTransaction) GetStatementDateOk() (*time.Time, bool)`

GetStatementDateOk returns a tuple with the StatementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementDate

`func (o *FinancialAccountTransaction) SetStatementDate(v time.Time)`

SetStatementDate sets StatementDate field to given value.


### GetMcc

`func (o *FinancialAccountTransaction) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *FinancialAccountTransaction) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *FinancialAccountTransaction) SetMcc(v string)`

SetMcc sets Mcc field to given value.


### GetMaskedCardNumber

`func (o *FinancialAccountTransaction) GetMaskedCardNumber() string`

GetMaskedCardNumber returns the MaskedCardNumber field if non-nil, zero value otherwise.

### GetMaskedCardNumberOk

`func (o *FinancialAccountTransaction) GetMaskedCardNumberOk() (*string, bool)`

GetMaskedCardNumberOk returns a tuple with the MaskedCardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedCardNumber

`func (o *FinancialAccountTransaction) SetMaskedCardNumber(v string)`

SetMaskedCardNumber sets MaskedCardNumber field to given value.


### GetAmc

`func (o *FinancialAccountTransaction) GetAmc() string`

GetAmc returns the Amc field if non-nil, zero value otherwise.

### GetAmcOk

`func (o *FinancialAccountTransaction) GetAmcOk() (*string, bool)`

GetAmcOk returns a tuple with the Amc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmc

`func (o *FinancialAccountTransaction) SetAmc(v string)`

SetAmc sets Amc field to given value.


### GetRegistrar

`func (o *FinancialAccountTransaction) GetRegistrar() string`

GetRegistrar returns the Registrar field if non-nil, zero value otherwise.

### GetRegistrarOk

`func (o *FinancialAccountTransaction) GetRegistrarOk() (*string, bool)`

GetRegistrarOk returns a tuple with the Registrar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrar

`func (o *FinancialAccountTransaction) SetRegistrar(v string)`

SetRegistrar sets Registrar field to given value.


### GetSchemeCode

`func (o *FinancialAccountTransaction) GetSchemeCode() string`

GetSchemeCode returns the SchemeCode field if non-nil, zero value otherwise.

### GetSchemeCodeOk

`func (o *FinancialAccountTransaction) GetSchemeCodeOk() (*string, bool)`

GetSchemeCodeOk returns a tuple with the SchemeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeCode

`func (o *FinancialAccountTransaction) SetSchemeCode(v string)`

SetSchemeCode sets SchemeCode field to given value.


### GetSchemePlan

`func (o *FinancialAccountTransaction) GetSchemePlan() MutualFundSchemePlan`

GetSchemePlan returns the SchemePlan field if non-nil, zero value otherwise.

### GetSchemePlanOk

`func (o *FinancialAccountTransaction) GetSchemePlanOk() (*MutualFundSchemePlan, bool)`

GetSchemePlanOk returns a tuple with the SchemePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemePlan

`func (o *FinancialAccountTransaction) SetSchemePlan(v MutualFundSchemePlan)`

SetSchemePlan sets SchemePlan field to given value.


### GetIsin

`func (o *FinancialAccountTransaction) GetIsin() string`

GetIsin returns the Isin field if non-nil, zero value otherwise.

### GetIsinOk

`func (o *FinancialAccountTransaction) GetIsinOk() (*string, bool)`

GetIsinOk returns a tuple with the Isin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsin

`func (o *FinancialAccountTransaction) SetIsin(v string)`

SetIsin sets Isin field to given value.


### GetAmfiCode

`func (o *FinancialAccountTransaction) GetAmfiCode() string`

GetAmfiCode returns the AmfiCode field if non-nil, zero value otherwise.

### GetAmfiCodeOk

`func (o *FinancialAccountTransaction) GetAmfiCodeOk() (*string, bool)`

GetAmfiCodeOk returns a tuple with the AmfiCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfiCode

`func (o *FinancialAccountTransaction) SetAmfiCode(v string)`

SetAmfiCode sets AmfiCode field to given value.


### GetFundType

`func (o *FinancialAccountTransaction) GetFundType() MutualFundFundType`

GetFundType returns the FundType field if non-nil, zero value otherwise.

### GetFundTypeOk

`func (o *FinancialAccountTransaction) GetFundTypeOk() (*MutualFundFundType, bool)`

GetFundTypeOk returns a tuple with the FundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundType

`func (o *FinancialAccountTransaction) SetFundType(v MutualFundFundType)`

SetFundType sets FundType field to given value.


### GetSchemeOption

`func (o *FinancialAccountTransaction) GetSchemeOption() MutualFundSchemeOption`

GetSchemeOption returns the SchemeOption field if non-nil, zero value otherwise.

### GetSchemeOptionOk

`func (o *FinancialAccountTransaction) GetSchemeOptionOk() (*MutualFundSchemeOption, bool)`

GetSchemeOptionOk returns a tuple with the SchemeOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeOption

`func (o *FinancialAccountTransaction) SetSchemeOption(v MutualFundSchemeOption)`

SetSchemeOption sets SchemeOption field to given value.


### GetSchemeTypes

`func (o *FinancialAccountTransaction) GetSchemeTypes() MutualFundSchemeType`

GetSchemeTypes returns the SchemeTypes field if non-nil, zero value otherwise.

### GetSchemeTypesOk

`func (o *FinancialAccountTransaction) GetSchemeTypesOk() (*MutualFundSchemeType, bool)`

GetSchemeTypesOk returns a tuple with the SchemeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeTypes

`func (o *FinancialAccountTransaction) SetSchemeTypes(v MutualFundSchemeType)`

SetSchemeTypes sets SchemeTypes field to given value.


### GetSchemeCategory

`func (o *FinancialAccountTransaction) GetSchemeCategory() MutualFundSchemeCategory`

GetSchemeCategory returns the SchemeCategory field if non-nil, zero value otherwise.

### GetSchemeCategoryOk

`func (o *FinancialAccountTransaction) GetSchemeCategoryOk() (*MutualFundSchemeCategory, bool)`

GetSchemeCategoryOk returns a tuple with the SchemeCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeCategory

`func (o *FinancialAccountTransaction) SetSchemeCategory(v MutualFundSchemeCategory)`

SetSchemeCategory sets SchemeCategory field to given value.


### GetUcc

`func (o *FinancialAccountTransaction) GetUcc() string`

GetUcc returns the Ucc field if non-nil, zero value otherwise.

### GetUccOk

`func (o *FinancialAccountTransaction) GetUccOk() (*string, bool)`

GetUccOk returns a tuple with the Ucc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcc

`func (o *FinancialAccountTransaction) SetUcc(v string)`

SetUcc sets Ucc field to given value.


### GetClosingUnits

`func (o *FinancialAccountTransaction) GetClosingUnits() string`

GetClosingUnits returns the ClosingUnits field if non-nil, zero value otherwise.

### GetClosingUnitsOk

`func (o *FinancialAccountTransaction) GetClosingUnitsOk() (*string, bool)`

GetClosingUnitsOk returns a tuple with the ClosingUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingUnits

`func (o *FinancialAccountTransaction) SetClosingUnits(v string)`

SetClosingUnits sets ClosingUnits field to given value.


### GetLienUnits

`func (o *FinancialAccountTransaction) GetLienUnits() string`

GetLienUnits returns the LienUnits field if non-nil, zero value otherwise.

### GetLienUnitsOk

`func (o *FinancialAccountTransaction) GetLienUnitsOk() (*string, bool)`

GetLienUnitsOk returns a tuple with the LienUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLienUnits

`func (o *FinancialAccountTransaction) SetLienUnits(v string)`

SetLienUnits sets LienUnits field to given value.


### GetNav

`func (o *FinancialAccountTransaction) GetNav() string`

GetNav returns the Nav field if non-nil, zero value otherwise.

### GetNavOk

`func (o *FinancialAccountTransaction) GetNavOk() (*string, bool)`

GetNavOk returns a tuple with the Nav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNav

`func (o *FinancialAccountTransaction) SetNav(v string)`

SetNav sets Nav field to given value.


### GetNavDate

`func (o *FinancialAccountTransaction) GetNavDate() time.Time`

GetNavDate returns the NavDate field if non-nil, zero value otherwise.

### GetNavDateOk

`func (o *FinancialAccountTransaction) GetNavDateOk() (*time.Time, bool)`

GetNavDateOk returns a tuple with the NavDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavDate

`func (o *FinancialAccountTransaction) SetNavDate(v time.Time)`

SetNavDate sets NavDate field to given value.


### GetOrderDate

`func (o *FinancialAccountTransaction) GetOrderDate() time.Time`

GetOrderDate returns the OrderDate field if non-nil, zero value otherwise.

### GetOrderDateOk

`func (o *FinancialAccountTransaction) GetOrderDateOk() (*time.Time, bool)`

GetOrderDateOk returns a tuple with the OrderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDate

`func (o *FinancialAccountTransaction) SetOrderDate(v time.Time)`

SetOrderDate sets OrderDate field to given value.


### GetExecutionDate

`func (o *FinancialAccountTransaction) GetExecutionDate() time.Time`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *FinancialAccountTransaction) GetExecutionDateOk() (*time.Time, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *FinancialAccountTransaction) SetExecutionDate(v time.Time)`

SetExecutionDate sets ExecutionDate field to given value.


### GetLockinFlag

`func (o *FinancialAccountTransaction) GetLockinFlag() string`

GetLockinFlag returns the LockinFlag field if non-nil, zero value otherwise.

### GetLockinFlagOk

`func (o *FinancialAccountTransaction) GetLockinFlagOk() (*string, bool)`

GetLockinFlagOk returns a tuple with the LockinFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockinFlag

`func (o *FinancialAccountTransaction) SetLockinFlag(v string)`

SetLockinFlag sets LockinFlag field to given value.


### GetLockinDays

`func (o *FinancialAccountTransaction) GetLockinDays() string`

GetLockinDays returns the LockinDays field if non-nil, zero value otherwise.

### GetLockinDaysOk

`func (o *FinancialAccountTransaction) GetLockinDaysOk() (*string, bool)`

GetLockinDaysOk returns a tuple with the LockinDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockinDays

`func (o *FinancialAccountTransaction) SetLockinDays(v string)`

SetLockinDays sets LockinDays field to given value.


### GetMode

`func (o *FinancialAccountTransaction) GetMode() MutualFundHoldingMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *FinancialAccountTransaction) GetModeOk() (*MutualFundHoldingMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *FinancialAccountTransaction) SetMode(v MutualFundHoldingMode)`

SetMode sets Mode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



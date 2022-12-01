# FinancialAccountTransactionMutualFundTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MutualFundTransactionType**](MutualFundTransactionType.md) |  | 
**Id** | **string** |  | 
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
**Amount** | **string** |  | 
**ClosingUnits** | **string** |  | 
**LienUnits** | **string** |  | 
**Nav** | **string** |  | 
**NavDate** | **time.Time** |  | 
**OrderDate** | **time.Time** |  | 
**ExecutionDate** | **time.Time** |  | 
**LockinFlag** | **string** |  | 
**LockinDays** | **string** |  | 
**Mode** | [**MutualFundHoldingMode**](MutualFundHoldingMode.md) |  | 
**Narration** | **string** |  | 

## Methods

### NewFinancialAccountTransactionMutualFundTransaction

`func NewFinancialAccountTransactionMutualFundTransaction(type_ MutualFundTransactionType, id string, amc string, registrar string, schemeCode string, schemePlan MutualFundSchemePlan, isin string, amfiCode string, fundType MutualFundFundType, schemeOption MutualFundSchemeOption, schemeTypes MutualFundSchemeType, schemeCategory MutualFundSchemeCategory, ucc string, amount string, closingUnits string, lienUnits string, nav string, navDate time.Time, orderDate time.Time, executionDate time.Time, lockinFlag string, lockinDays string, mode MutualFundHoldingMode, narration string, ) *FinancialAccountTransactionMutualFundTransaction`

NewFinancialAccountTransactionMutualFundTransaction instantiates a new FinancialAccountTransactionMutualFundTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialAccountTransactionMutualFundTransactionWithDefaults

`func NewFinancialAccountTransactionMutualFundTransactionWithDefaults() *FinancialAccountTransactionMutualFundTransaction`

NewFinancialAccountTransactionMutualFundTransactionWithDefaults instantiates a new FinancialAccountTransactionMutualFundTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FinancialAccountTransactionMutualFundTransaction) GetType() MutualFundTransactionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetTypeOk() (*MutualFundTransactionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FinancialAccountTransactionMutualFundTransaction) SetType(v MutualFundTransactionType)`

SetType sets Type field to given value.


### GetId

`func (o *FinancialAccountTransactionMutualFundTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FinancialAccountTransactionMutualFundTransaction) SetId(v string)`

SetId sets Id field to given value.


### GetAmc

`func (o *FinancialAccountTransactionMutualFundTransaction) GetAmc() string`

GetAmc returns the Amc field if non-nil, zero value otherwise.

### GetAmcOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetAmcOk() (*string, bool)`

GetAmcOk returns a tuple with the Amc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmc

`func (o *FinancialAccountTransactionMutualFundTransaction) SetAmc(v string)`

SetAmc sets Amc field to given value.


### GetRegistrar

`func (o *FinancialAccountTransactionMutualFundTransaction) GetRegistrar() string`

GetRegistrar returns the Registrar field if non-nil, zero value otherwise.

### GetRegistrarOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetRegistrarOk() (*string, bool)`

GetRegistrarOk returns a tuple with the Registrar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrar

`func (o *FinancialAccountTransactionMutualFundTransaction) SetRegistrar(v string)`

SetRegistrar sets Registrar field to given value.


### GetSchemeCode

`func (o *FinancialAccountTransactionMutualFundTransaction) GetSchemeCode() string`

GetSchemeCode returns the SchemeCode field if non-nil, zero value otherwise.

### GetSchemeCodeOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetSchemeCodeOk() (*string, bool)`

GetSchemeCodeOk returns a tuple with the SchemeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeCode

`func (o *FinancialAccountTransactionMutualFundTransaction) SetSchemeCode(v string)`

SetSchemeCode sets SchemeCode field to given value.


### GetSchemePlan

`func (o *FinancialAccountTransactionMutualFundTransaction) GetSchemePlan() MutualFundSchemePlan`

GetSchemePlan returns the SchemePlan field if non-nil, zero value otherwise.

### GetSchemePlanOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetSchemePlanOk() (*MutualFundSchemePlan, bool)`

GetSchemePlanOk returns a tuple with the SchemePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemePlan

`func (o *FinancialAccountTransactionMutualFundTransaction) SetSchemePlan(v MutualFundSchemePlan)`

SetSchemePlan sets SchemePlan field to given value.


### GetIsin

`func (o *FinancialAccountTransactionMutualFundTransaction) GetIsin() string`

GetIsin returns the Isin field if non-nil, zero value otherwise.

### GetIsinOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetIsinOk() (*string, bool)`

GetIsinOk returns a tuple with the Isin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsin

`func (o *FinancialAccountTransactionMutualFundTransaction) SetIsin(v string)`

SetIsin sets Isin field to given value.


### GetAmfiCode

`func (o *FinancialAccountTransactionMutualFundTransaction) GetAmfiCode() string`

GetAmfiCode returns the AmfiCode field if non-nil, zero value otherwise.

### GetAmfiCodeOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetAmfiCodeOk() (*string, bool)`

GetAmfiCodeOk returns a tuple with the AmfiCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfiCode

`func (o *FinancialAccountTransactionMutualFundTransaction) SetAmfiCode(v string)`

SetAmfiCode sets AmfiCode field to given value.


### GetFundType

`func (o *FinancialAccountTransactionMutualFundTransaction) GetFundType() MutualFundFundType`

GetFundType returns the FundType field if non-nil, zero value otherwise.

### GetFundTypeOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetFundTypeOk() (*MutualFundFundType, bool)`

GetFundTypeOk returns a tuple with the FundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundType

`func (o *FinancialAccountTransactionMutualFundTransaction) SetFundType(v MutualFundFundType)`

SetFundType sets FundType field to given value.


### GetSchemeOption

`func (o *FinancialAccountTransactionMutualFundTransaction) GetSchemeOption() MutualFundSchemeOption`

GetSchemeOption returns the SchemeOption field if non-nil, zero value otherwise.

### GetSchemeOptionOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetSchemeOptionOk() (*MutualFundSchemeOption, bool)`

GetSchemeOptionOk returns a tuple with the SchemeOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeOption

`func (o *FinancialAccountTransactionMutualFundTransaction) SetSchemeOption(v MutualFundSchemeOption)`

SetSchemeOption sets SchemeOption field to given value.


### GetSchemeTypes

`func (o *FinancialAccountTransactionMutualFundTransaction) GetSchemeTypes() MutualFundSchemeType`

GetSchemeTypes returns the SchemeTypes field if non-nil, zero value otherwise.

### GetSchemeTypesOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetSchemeTypesOk() (*MutualFundSchemeType, bool)`

GetSchemeTypesOk returns a tuple with the SchemeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeTypes

`func (o *FinancialAccountTransactionMutualFundTransaction) SetSchemeTypes(v MutualFundSchemeType)`

SetSchemeTypes sets SchemeTypes field to given value.


### GetSchemeCategory

`func (o *FinancialAccountTransactionMutualFundTransaction) GetSchemeCategory() MutualFundSchemeCategory`

GetSchemeCategory returns the SchemeCategory field if non-nil, zero value otherwise.

### GetSchemeCategoryOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetSchemeCategoryOk() (*MutualFundSchemeCategory, bool)`

GetSchemeCategoryOk returns a tuple with the SchemeCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeCategory

`func (o *FinancialAccountTransactionMutualFundTransaction) SetSchemeCategory(v MutualFundSchemeCategory)`

SetSchemeCategory sets SchemeCategory field to given value.


### GetUcc

`func (o *FinancialAccountTransactionMutualFundTransaction) GetUcc() string`

GetUcc returns the Ucc field if non-nil, zero value otherwise.

### GetUccOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetUccOk() (*string, bool)`

GetUccOk returns a tuple with the Ucc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcc

`func (o *FinancialAccountTransactionMutualFundTransaction) SetUcc(v string)`

SetUcc sets Ucc field to given value.


### GetAmount

`func (o *FinancialAccountTransactionMutualFundTransaction) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FinancialAccountTransactionMutualFundTransaction) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetClosingUnits

`func (o *FinancialAccountTransactionMutualFundTransaction) GetClosingUnits() string`

GetClosingUnits returns the ClosingUnits field if non-nil, zero value otherwise.

### GetClosingUnitsOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetClosingUnitsOk() (*string, bool)`

GetClosingUnitsOk returns a tuple with the ClosingUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingUnits

`func (o *FinancialAccountTransactionMutualFundTransaction) SetClosingUnits(v string)`

SetClosingUnits sets ClosingUnits field to given value.


### GetLienUnits

`func (o *FinancialAccountTransactionMutualFundTransaction) GetLienUnits() string`

GetLienUnits returns the LienUnits field if non-nil, zero value otherwise.

### GetLienUnitsOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetLienUnitsOk() (*string, bool)`

GetLienUnitsOk returns a tuple with the LienUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLienUnits

`func (o *FinancialAccountTransactionMutualFundTransaction) SetLienUnits(v string)`

SetLienUnits sets LienUnits field to given value.


### GetNav

`func (o *FinancialAccountTransactionMutualFundTransaction) GetNav() string`

GetNav returns the Nav field if non-nil, zero value otherwise.

### GetNavOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetNavOk() (*string, bool)`

GetNavOk returns a tuple with the Nav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNav

`func (o *FinancialAccountTransactionMutualFundTransaction) SetNav(v string)`

SetNav sets Nav field to given value.


### GetNavDate

`func (o *FinancialAccountTransactionMutualFundTransaction) GetNavDate() time.Time`

GetNavDate returns the NavDate field if non-nil, zero value otherwise.

### GetNavDateOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetNavDateOk() (*time.Time, bool)`

GetNavDateOk returns a tuple with the NavDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavDate

`func (o *FinancialAccountTransactionMutualFundTransaction) SetNavDate(v time.Time)`

SetNavDate sets NavDate field to given value.


### GetOrderDate

`func (o *FinancialAccountTransactionMutualFundTransaction) GetOrderDate() time.Time`

GetOrderDate returns the OrderDate field if non-nil, zero value otherwise.

### GetOrderDateOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetOrderDateOk() (*time.Time, bool)`

GetOrderDateOk returns a tuple with the OrderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDate

`func (o *FinancialAccountTransactionMutualFundTransaction) SetOrderDate(v time.Time)`

SetOrderDate sets OrderDate field to given value.


### GetExecutionDate

`func (o *FinancialAccountTransactionMutualFundTransaction) GetExecutionDate() time.Time`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetExecutionDateOk() (*time.Time, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *FinancialAccountTransactionMutualFundTransaction) SetExecutionDate(v time.Time)`

SetExecutionDate sets ExecutionDate field to given value.


### GetLockinFlag

`func (o *FinancialAccountTransactionMutualFundTransaction) GetLockinFlag() string`

GetLockinFlag returns the LockinFlag field if non-nil, zero value otherwise.

### GetLockinFlagOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetLockinFlagOk() (*string, bool)`

GetLockinFlagOk returns a tuple with the LockinFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockinFlag

`func (o *FinancialAccountTransactionMutualFundTransaction) SetLockinFlag(v string)`

SetLockinFlag sets LockinFlag field to given value.


### GetLockinDays

`func (o *FinancialAccountTransactionMutualFundTransaction) GetLockinDays() string`

GetLockinDays returns the LockinDays field if non-nil, zero value otherwise.

### GetLockinDaysOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetLockinDaysOk() (*string, bool)`

GetLockinDaysOk returns a tuple with the LockinDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockinDays

`func (o *FinancialAccountTransactionMutualFundTransaction) SetLockinDays(v string)`

SetLockinDays sets LockinDays field to given value.


### GetMode

`func (o *FinancialAccountTransactionMutualFundTransaction) GetMode() MutualFundHoldingMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetModeOk() (*MutualFundHoldingMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *FinancialAccountTransactionMutualFundTransaction) SetMode(v MutualFundHoldingMode)`

SetMode sets Mode field to given value.


### GetNarration

`func (o *FinancialAccountTransactionMutualFundTransaction) GetNarration() string`

GetNarration returns the Narration field if non-nil, zero value otherwise.

### GetNarrationOk

`func (o *FinancialAccountTransactionMutualFundTransaction) GetNarrationOk() (*string, bool)`

GetNarrationOk returns a tuple with the Narration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNarration

`func (o *FinancialAccountTransactionMutualFundTransaction) SetNarration(v string)`

SetNarration sets Narration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



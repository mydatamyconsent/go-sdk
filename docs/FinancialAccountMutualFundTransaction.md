# FinancialAccountMutualFundTransaction

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

### NewFinancialAccountMutualFundTransaction

`func NewFinancialAccountMutualFundTransaction(type_ MutualFundTransactionType, id string, amc string, registrar string, schemeCode string, schemePlan MutualFundSchemePlan, isin string, amfiCode string, fundType MutualFundFundType, schemeOption MutualFundSchemeOption, schemeTypes MutualFundSchemeType, schemeCategory MutualFundSchemeCategory, ucc string, amount string, closingUnits string, lienUnits string, nav string, navDate time.Time, orderDate time.Time, executionDate time.Time, lockinFlag string, lockinDays string, mode MutualFundHoldingMode, narration string, ) *FinancialAccountMutualFundTransaction`

NewFinancialAccountMutualFundTransaction instantiates a new FinancialAccountMutualFundTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialAccountMutualFundTransactionWithDefaults

`func NewFinancialAccountMutualFundTransactionWithDefaults() *FinancialAccountMutualFundTransaction`

NewFinancialAccountMutualFundTransactionWithDefaults instantiates a new FinancialAccountMutualFundTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FinancialAccountMutualFundTransaction) GetType() MutualFundTransactionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FinancialAccountMutualFundTransaction) GetTypeOk() (*MutualFundTransactionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FinancialAccountMutualFundTransaction) SetType(v MutualFundTransactionType)`

SetType sets Type field to given value.


### GetId

`func (o *FinancialAccountMutualFundTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FinancialAccountMutualFundTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FinancialAccountMutualFundTransaction) SetId(v string)`

SetId sets Id field to given value.


### GetAmc

`func (o *FinancialAccountMutualFundTransaction) GetAmc() string`

GetAmc returns the Amc field if non-nil, zero value otherwise.

### GetAmcOk

`func (o *FinancialAccountMutualFundTransaction) GetAmcOk() (*string, bool)`

GetAmcOk returns a tuple with the Amc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmc

`func (o *FinancialAccountMutualFundTransaction) SetAmc(v string)`

SetAmc sets Amc field to given value.


### GetRegistrar

`func (o *FinancialAccountMutualFundTransaction) GetRegistrar() string`

GetRegistrar returns the Registrar field if non-nil, zero value otherwise.

### GetRegistrarOk

`func (o *FinancialAccountMutualFundTransaction) GetRegistrarOk() (*string, bool)`

GetRegistrarOk returns a tuple with the Registrar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrar

`func (o *FinancialAccountMutualFundTransaction) SetRegistrar(v string)`

SetRegistrar sets Registrar field to given value.


### GetSchemeCode

`func (o *FinancialAccountMutualFundTransaction) GetSchemeCode() string`

GetSchemeCode returns the SchemeCode field if non-nil, zero value otherwise.

### GetSchemeCodeOk

`func (o *FinancialAccountMutualFundTransaction) GetSchemeCodeOk() (*string, bool)`

GetSchemeCodeOk returns a tuple with the SchemeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeCode

`func (o *FinancialAccountMutualFundTransaction) SetSchemeCode(v string)`

SetSchemeCode sets SchemeCode field to given value.


### GetSchemePlan

`func (o *FinancialAccountMutualFundTransaction) GetSchemePlan() MutualFundSchemePlan`

GetSchemePlan returns the SchemePlan field if non-nil, zero value otherwise.

### GetSchemePlanOk

`func (o *FinancialAccountMutualFundTransaction) GetSchemePlanOk() (*MutualFundSchemePlan, bool)`

GetSchemePlanOk returns a tuple with the SchemePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemePlan

`func (o *FinancialAccountMutualFundTransaction) SetSchemePlan(v MutualFundSchemePlan)`

SetSchemePlan sets SchemePlan field to given value.


### GetIsin

`func (o *FinancialAccountMutualFundTransaction) GetIsin() string`

GetIsin returns the Isin field if non-nil, zero value otherwise.

### GetIsinOk

`func (o *FinancialAccountMutualFundTransaction) GetIsinOk() (*string, bool)`

GetIsinOk returns a tuple with the Isin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsin

`func (o *FinancialAccountMutualFundTransaction) SetIsin(v string)`

SetIsin sets Isin field to given value.


### GetAmfiCode

`func (o *FinancialAccountMutualFundTransaction) GetAmfiCode() string`

GetAmfiCode returns the AmfiCode field if non-nil, zero value otherwise.

### GetAmfiCodeOk

`func (o *FinancialAccountMutualFundTransaction) GetAmfiCodeOk() (*string, bool)`

GetAmfiCodeOk returns a tuple with the AmfiCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfiCode

`func (o *FinancialAccountMutualFundTransaction) SetAmfiCode(v string)`

SetAmfiCode sets AmfiCode field to given value.


### GetFundType

`func (o *FinancialAccountMutualFundTransaction) GetFundType() MutualFundFundType`

GetFundType returns the FundType field if non-nil, zero value otherwise.

### GetFundTypeOk

`func (o *FinancialAccountMutualFundTransaction) GetFundTypeOk() (*MutualFundFundType, bool)`

GetFundTypeOk returns a tuple with the FundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundType

`func (o *FinancialAccountMutualFundTransaction) SetFundType(v MutualFundFundType)`

SetFundType sets FundType field to given value.


### GetSchemeOption

`func (o *FinancialAccountMutualFundTransaction) GetSchemeOption() MutualFundSchemeOption`

GetSchemeOption returns the SchemeOption field if non-nil, zero value otherwise.

### GetSchemeOptionOk

`func (o *FinancialAccountMutualFundTransaction) GetSchemeOptionOk() (*MutualFundSchemeOption, bool)`

GetSchemeOptionOk returns a tuple with the SchemeOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeOption

`func (o *FinancialAccountMutualFundTransaction) SetSchemeOption(v MutualFundSchemeOption)`

SetSchemeOption sets SchemeOption field to given value.


### GetSchemeTypes

`func (o *FinancialAccountMutualFundTransaction) GetSchemeTypes() MutualFundSchemeType`

GetSchemeTypes returns the SchemeTypes field if non-nil, zero value otherwise.

### GetSchemeTypesOk

`func (o *FinancialAccountMutualFundTransaction) GetSchemeTypesOk() (*MutualFundSchemeType, bool)`

GetSchemeTypesOk returns a tuple with the SchemeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeTypes

`func (o *FinancialAccountMutualFundTransaction) SetSchemeTypes(v MutualFundSchemeType)`

SetSchemeTypes sets SchemeTypes field to given value.


### GetSchemeCategory

`func (o *FinancialAccountMutualFundTransaction) GetSchemeCategory() MutualFundSchemeCategory`

GetSchemeCategory returns the SchemeCategory field if non-nil, zero value otherwise.

### GetSchemeCategoryOk

`func (o *FinancialAccountMutualFundTransaction) GetSchemeCategoryOk() (*MutualFundSchemeCategory, bool)`

GetSchemeCategoryOk returns a tuple with the SchemeCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeCategory

`func (o *FinancialAccountMutualFundTransaction) SetSchemeCategory(v MutualFundSchemeCategory)`

SetSchemeCategory sets SchemeCategory field to given value.


### GetUcc

`func (o *FinancialAccountMutualFundTransaction) GetUcc() string`

GetUcc returns the Ucc field if non-nil, zero value otherwise.

### GetUccOk

`func (o *FinancialAccountMutualFundTransaction) GetUccOk() (*string, bool)`

GetUccOk returns a tuple with the Ucc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcc

`func (o *FinancialAccountMutualFundTransaction) SetUcc(v string)`

SetUcc sets Ucc field to given value.


### GetAmount

`func (o *FinancialAccountMutualFundTransaction) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FinancialAccountMutualFundTransaction) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FinancialAccountMutualFundTransaction) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetClosingUnits

`func (o *FinancialAccountMutualFundTransaction) GetClosingUnits() string`

GetClosingUnits returns the ClosingUnits field if non-nil, zero value otherwise.

### GetClosingUnitsOk

`func (o *FinancialAccountMutualFundTransaction) GetClosingUnitsOk() (*string, bool)`

GetClosingUnitsOk returns a tuple with the ClosingUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingUnits

`func (o *FinancialAccountMutualFundTransaction) SetClosingUnits(v string)`

SetClosingUnits sets ClosingUnits field to given value.


### GetLienUnits

`func (o *FinancialAccountMutualFundTransaction) GetLienUnits() string`

GetLienUnits returns the LienUnits field if non-nil, zero value otherwise.

### GetLienUnitsOk

`func (o *FinancialAccountMutualFundTransaction) GetLienUnitsOk() (*string, bool)`

GetLienUnitsOk returns a tuple with the LienUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLienUnits

`func (o *FinancialAccountMutualFundTransaction) SetLienUnits(v string)`

SetLienUnits sets LienUnits field to given value.


### GetNav

`func (o *FinancialAccountMutualFundTransaction) GetNav() string`

GetNav returns the Nav field if non-nil, zero value otherwise.

### GetNavOk

`func (o *FinancialAccountMutualFundTransaction) GetNavOk() (*string, bool)`

GetNavOk returns a tuple with the Nav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNav

`func (o *FinancialAccountMutualFundTransaction) SetNav(v string)`

SetNav sets Nav field to given value.


### GetNavDate

`func (o *FinancialAccountMutualFundTransaction) GetNavDate() time.Time`

GetNavDate returns the NavDate field if non-nil, zero value otherwise.

### GetNavDateOk

`func (o *FinancialAccountMutualFundTransaction) GetNavDateOk() (*time.Time, bool)`

GetNavDateOk returns a tuple with the NavDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavDate

`func (o *FinancialAccountMutualFundTransaction) SetNavDate(v time.Time)`

SetNavDate sets NavDate field to given value.


### GetOrderDate

`func (o *FinancialAccountMutualFundTransaction) GetOrderDate() time.Time`

GetOrderDate returns the OrderDate field if non-nil, zero value otherwise.

### GetOrderDateOk

`func (o *FinancialAccountMutualFundTransaction) GetOrderDateOk() (*time.Time, bool)`

GetOrderDateOk returns a tuple with the OrderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDate

`func (o *FinancialAccountMutualFundTransaction) SetOrderDate(v time.Time)`

SetOrderDate sets OrderDate field to given value.


### GetExecutionDate

`func (o *FinancialAccountMutualFundTransaction) GetExecutionDate() time.Time`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *FinancialAccountMutualFundTransaction) GetExecutionDateOk() (*time.Time, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *FinancialAccountMutualFundTransaction) SetExecutionDate(v time.Time)`

SetExecutionDate sets ExecutionDate field to given value.


### GetLockinFlag

`func (o *FinancialAccountMutualFundTransaction) GetLockinFlag() string`

GetLockinFlag returns the LockinFlag field if non-nil, zero value otherwise.

### GetLockinFlagOk

`func (o *FinancialAccountMutualFundTransaction) GetLockinFlagOk() (*string, bool)`

GetLockinFlagOk returns a tuple with the LockinFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockinFlag

`func (o *FinancialAccountMutualFundTransaction) SetLockinFlag(v string)`

SetLockinFlag sets LockinFlag field to given value.


### GetLockinDays

`func (o *FinancialAccountMutualFundTransaction) GetLockinDays() string`

GetLockinDays returns the LockinDays field if non-nil, zero value otherwise.

### GetLockinDaysOk

`func (o *FinancialAccountMutualFundTransaction) GetLockinDaysOk() (*string, bool)`

GetLockinDaysOk returns a tuple with the LockinDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockinDays

`func (o *FinancialAccountMutualFundTransaction) SetLockinDays(v string)`

SetLockinDays sets LockinDays field to given value.


### GetMode

`func (o *FinancialAccountMutualFundTransaction) GetMode() MutualFundHoldingMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *FinancialAccountMutualFundTransaction) GetModeOk() (*MutualFundHoldingMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *FinancialAccountMutualFundTransaction) SetMode(v MutualFundHoldingMode)`

SetMode sets Mode field to given value.


### GetNarration

`func (o *FinancialAccountMutualFundTransaction) GetNarration() string`

GetNarration returns the Narration field if non-nil, zero value otherwise.

### GetNarrationOk

`func (o *FinancialAccountMutualFundTransaction) GetNarrationOk() (*string, bool)`

GetNarrationOk returns a tuple with the Narration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNarration

`func (o *FinancialAccountMutualFundTransaction) SetNarration(v string)`

SetNarration sets Narration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



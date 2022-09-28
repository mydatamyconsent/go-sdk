# MutualFundTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**Type** | [**MutualFundTransactionType**](MutualFundTransactionType.md) |  | 
**OrderDate** | **time.Time** |  | 
**ExecutionDate** | **time.Time** |  | 
**LockinFlag** | **string** |  | 
**LockinDays** | **string** |  | 
**Mode** | [**MutualFundHoldingMode**](MutualFundHoldingMode.md) |  | 
**Narration** | **string** |  | 

## Methods

### NewMutualFundTransaction

`func NewMutualFundTransaction(id string, amc string, registrar string, schemeCode string, schemePlan MutualFundSchemePlan, isin string, amfiCode string, fundType MutualFundFundType, schemeOption MutualFundSchemeOption, schemeTypes MutualFundSchemeType, schemeCategory MutualFundSchemeCategory, ucc string, amount string, closingUnits string, lienUnits string, nav string, navDate time.Time, type_ MutualFundTransactionType, orderDate time.Time, executionDate time.Time, lockinFlag string, lockinDays string, mode MutualFundHoldingMode, narration string, ) *MutualFundTransaction`

NewMutualFundTransaction instantiates a new MutualFundTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutualFundTransactionWithDefaults

`func NewMutualFundTransactionWithDefaults() *MutualFundTransaction`

NewMutualFundTransactionWithDefaults instantiates a new MutualFundTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MutualFundTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MutualFundTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MutualFundTransaction) SetId(v string)`

SetId sets Id field to given value.


### GetAmc

`func (o *MutualFundTransaction) GetAmc() string`

GetAmc returns the Amc field if non-nil, zero value otherwise.

### GetAmcOk

`func (o *MutualFundTransaction) GetAmcOk() (*string, bool)`

GetAmcOk returns a tuple with the Amc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmc

`func (o *MutualFundTransaction) SetAmc(v string)`

SetAmc sets Amc field to given value.


### GetRegistrar

`func (o *MutualFundTransaction) GetRegistrar() string`

GetRegistrar returns the Registrar field if non-nil, zero value otherwise.

### GetRegistrarOk

`func (o *MutualFundTransaction) GetRegistrarOk() (*string, bool)`

GetRegistrarOk returns a tuple with the Registrar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrar

`func (o *MutualFundTransaction) SetRegistrar(v string)`

SetRegistrar sets Registrar field to given value.


### GetSchemeCode

`func (o *MutualFundTransaction) GetSchemeCode() string`

GetSchemeCode returns the SchemeCode field if non-nil, zero value otherwise.

### GetSchemeCodeOk

`func (o *MutualFundTransaction) GetSchemeCodeOk() (*string, bool)`

GetSchemeCodeOk returns a tuple with the SchemeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeCode

`func (o *MutualFundTransaction) SetSchemeCode(v string)`

SetSchemeCode sets SchemeCode field to given value.


### GetSchemePlan

`func (o *MutualFundTransaction) GetSchemePlan() MutualFundSchemePlan`

GetSchemePlan returns the SchemePlan field if non-nil, zero value otherwise.

### GetSchemePlanOk

`func (o *MutualFundTransaction) GetSchemePlanOk() (*MutualFundSchemePlan, bool)`

GetSchemePlanOk returns a tuple with the SchemePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemePlan

`func (o *MutualFundTransaction) SetSchemePlan(v MutualFundSchemePlan)`

SetSchemePlan sets SchemePlan field to given value.


### GetIsin

`func (o *MutualFundTransaction) GetIsin() string`

GetIsin returns the Isin field if non-nil, zero value otherwise.

### GetIsinOk

`func (o *MutualFundTransaction) GetIsinOk() (*string, bool)`

GetIsinOk returns a tuple with the Isin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsin

`func (o *MutualFundTransaction) SetIsin(v string)`

SetIsin sets Isin field to given value.


### GetAmfiCode

`func (o *MutualFundTransaction) GetAmfiCode() string`

GetAmfiCode returns the AmfiCode field if non-nil, zero value otherwise.

### GetAmfiCodeOk

`func (o *MutualFundTransaction) GetAmfiCodeOk() (*string, bool)`

GetAmfiCodeOk returns a tuple with the AmfiCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfiCode

`func (o *MutualFundTransaction) SetAmfiCode(v string)`

SetAmfiCode sets AmfiCode field to given value.


### GetFundType

`func (o *MutualFundTransaction) GetFundType() MutualFundFundType`

GetFundType returns the FundType field if non-nil, zero value otherwise.

### GetFundTypeOk

`func (o *MutualFundTransaction) GetFundTypeOk() (*MutualFundFundType, bool)`

GetFundTypeOk returns a tuple with the FundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundType

`func (o *MutualFundTransaction) SetFundType(v MutualFundFundType)`

SetFundType sets FundType field to given value.


### GetSchemeOption

`func (o *MutualFundTransaction) GetSchemeOption() MutualFundSchemeOption`

GetSchemeOption returns the SchemeOption field if non-nil, zero value otherwise.

### GetSchemeOptionOk

`func (o *MutualFundTransaction) GetSchemeOptionOk() (*MutualFundSchemeOption, bool)`

GetSchemeOptionOk returns a tuple with the SchemeOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeOption

`func (o *MutualFundTransaction) SetSchemeOption(v MutualFundSchemeOption)`

SetSchemeOption sets SchemeOption field to given value.


### GetSchemeTypes

`func (o *MutualFundTransaction) GetSchemeTypes() MutualFundSchemeType`

GetSchemeTypes returns the SchemeTypes field if non-nil, zero value otherwise.

### GetSchemeTypesOk

`func (o *MutualFundTransaction) GetSchemeTypesOk() (*MutualFundSchemeType, bool)`

GetSchemeTypesOk returns a tuple with the SchemeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeTypes

`func (o *MutualFundTransaction) SetSchemeTypes(v MutualFundSchemeType)`

SetSchemeTypes sets SchemeTypes field to given value.


### GetSchemeCategory

`func (o *MutualFundTransaction) GetSchemeCategory() MutualFundSchemeCategory`

GetSchemeCategory returns the SchemeCategory field if non-nil, zero value otherwise.

### GetSchemeCategoryOk

`func (o *MutualFundTransaction) GetSchemeCategoryOk() (*MutualFundSchemeCategory, bool)`

GetSchemeCategoryOk returns a tuple with the SchemeCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeCategory

`func (o *MutualFundTransaction) SetSchemeCategory(v MutualFundSchemeCategory)`

SetSchemeCategory sets SchemeCategory field to given value.


### GetUcc

`func (o *MutualFundTransaction) GetUcc() string`

GetUcc returns the Ucc field if non-nil, zero value otherwise.

### GetUccOk

`func (o *MutualFundTransaction) GetUccOk() (*string, bool)`

GetUccOk returns a tuple with the Ucc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcc

`func (o *MutualFundTransaction) SetUcc(v string)`

SetUcc sets Ucc field to given value.


### GetAmount

`func (o *MutualFundTransaction) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *MutualFundTransaction) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *MutualFundTransaction) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetClosingUnits

`func (o *MutualFundTransaction) GetClosingUnits() string`

GetClosingUnits returns the ClosingUnits field if non-nil, zero value otherwise.

### GetClosingUnitsOk

`func (o *MutualFundTransaction) GetClosingUnitsOk() (*string, bool)`

GetClosingUnitsOk returns a tuple with the ClosingUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingUnits

`func (o *MutualFundTransaction) SetClosingUnits(v string)`

SetClosingUnits sets ClosingUnits field to given value.


### GetLienUnits

`func (o *MutualFundTransaction) GetLienUnits() string`

GetLienUnits returns the LienUnits field if non-nil, zero value otherwise.

### GetLienUnitsOk

`func (o *MutualFundTransaction) GetLienUnitsOk() (*string, bool)`

GetLienUnitsOk returns a tuple with the LienUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLienUnits

`func (o *MutualFundTransaction) SetLienUnits(v string)`

SetLienUnits sets LienUnits field to given value.


### GetNav

`func (o *MutualFundTransaction) GetNav() string`

GetNav returns the Nav field if non-nil, zero value otherwise.

### GetNavOk

`func (o *MutualFundTransaction) GetNavOk() (*string, bool)`

GetNavOk returns a tuple with the Nav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNav

`func (o *MutualFundTransaction) SetNav(v string)`

SetNav sets Nav field to given value.


### GetNavDate

`func (o *MutualFundTransaction) GetNavDate() time.Time`

GetNavDate returns the NavDate field if non-nil, zero value otherwise.

### GetNavDateOk

`func (o *MutualFundTransaction) GetNavDateOk() (*time.Time, bool)`

GetNavDateOk returns a tuple with the NavDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavDate

`func (o *MutualFundTransaction) SetNavDate(v time.Time)`

SetNavDate sets NavDate field to given value.


### GetType

`func (o *MutualFundTransaction) GetType() MutualFundTransactionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MutualFundTransaction) GetTypeOk() (*MutualFundTransactionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MutualFundTransaction) SetType(v MutualFundTransactionType)`

SetType sets Type field to given value.


### GetOrderDate

`func (o *MutualFundTransaction) GetOrderDate() time.Time`

GetOrderDate returns the OrderDate field if non-nil, zero value otherwise.

### GetOrderDateOk

`func (o *MutualFundTransaction) GetOrderDateOk() (*time.Time, bool)`

GetOrderDateOk returns a tuple with the OrderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDate

`func (o *MutualFundTransaction) SetOrderDate(v time.Time)`

SetOrderDate sets OrderDate field to given value.


### GetExecutionDate

`func (o *MutualFundTransaction) GetExecutionDate() time.Time`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *MutualFundTransaction) GetExecutionDateOk() (*time.Time, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *MutualFundTransaction) SetExecutionDate(v time.Time)`

SetExecutionDate sets ExecutionDate field to given value.


### GetLockinFlag

`func (o *MutualFundTransaction) GetLockinFlag() string`

GetLockinFlag returns the LockinFlag field if non-nil, zero value otherwise.

### GetLockinFlagOk

`func (o *MutualFundTransaction) GetLockinFlagOk() (*string, bool)`

GetLockinFlagOk returns a tuple with the LockinFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockinFlag

`func (o *MutualFundTransaction) SetLockinFlag(v string)`

SetLockinFlag sets LockinFlag field to given value.


### GetLockinDays

`func (o *MutualFundTransaction) GetLockinDays() string`

GetLockinDays returns the LockinDays field if non-nil, zero value otherwise.

### GetLockinDaysOk

`func (o *MutualFundTransaction) GetLockinDaysOk() (*string, bool)`

GetLockinDaysOk returns a tuple with the LockinDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockinDays

`func (o *MutualFundTransaction) SetLockinDays(v string)`

SetLockinDays sets LockinDays field to given value.


### GetMode

`func (o *MutualFundTransaction) GetMode() MutualFundHoldingMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *MutualFundTransaction) GetModeOk() (*MutualFundHoldingMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *MutualFundTransaction) SetMode(v MutualFundHoldingMode)`

SetMode sets Mode field to given value.


### GetNarration

`func (o *MutualFundTransaction) GetNarration() string`

GetNarration returns the Narration field if non-nil, zero value otherwise.

### GetNarrationOk

`func (o *MutualFundTransaction) GetNarrationOk() (*string, bool)`

GetNarrationOk returns a tuple with the Narration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNarration

`func (o *MutualFundTransaction) SetNarration(v string)`

SetNarration sets Narration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# MutualFundHolding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Amc** | Pointer to **string** |  | [optional] 
**Registrar** | Pointer to **string** |  | [optional] 
**SchemeCode** | Pointer to **string** |  | [optional] 
**Isin** | **string** |  | 
**Ucc** | Pointer to **string** |  | [optional] 
**AmfiCode** | Pointer to **string** |  | [optional] 
**FolioNo** | **string** |  | 
**DividendType** | Pointer to **string** |  | [optional] 
**FatcaStatus** | Pointer to **string** |  | [optional] 
**Mode** | [**MutualFundHoldingMode**](MutualFundHoldingMode.md) |  | 
**Units** | **float64** |  | 
**ClosingUnits** | Pointer to **string** |  | [optional] 
**LienUnits** | Pointer to **string** |  | [optional] 
**Rate** | Pointer to **string** |  | [optional] 
**Nav** | Pointer to **string** |  | [optional] 
**LockingUnits** | Pointer to **string** |  | [optional] 

## Methods

### NewMutualFundHolding

`func NewMutualFundHolding(name string, isin string, folioNo string, mode MutualFundHoldingMode, units float64, ) *MutualFundHolding`

NewMutualFundHolding instantiates a new MutualFundHolding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutualFundHoldingWithDefaults

`func NewMutualFundHoldingWithDefaults() *MutualFundHolding`

NewMutualFundHoldingWithDefaults instantiates a new MutualFundHolding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MutualFundHolding) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MutualFundHolding) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MutualFundHolding) SetName(v string)`

SetName sets Name field to given value.


### GetAmc

`func (o *MutualFundHolding) GetAmc() string`

GetAmc returns the Amc field if non-nil, zero value otherwise.

### GetAmcOk

`func (o *MutualFundHolding) GetAmcOk() (*string, bool)`

GetAmcOk returns a tuple with the Amc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmc

`func (o *MutualFundHolding) SetAmc(v string)`

SetAmc sets Amc field to given value.

### HasAmc

`func (o *MutualFundHolding) HasAmc() bool`

HasAmc returns a boolean if a field has been set.

### GetRegistrar

`func (o *MutualFundHolding) GetRegistrar() string`

GetRegistrar returns the Registrar field if non-nil, zero value otherwise.

### GetRegistrarOk

`func (o *MutualFundHolding) GetRegistrarOk() (*string, bool)`

GetRegistrarOk returns a tuple with the Registrar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrar

`func (o *MutualFundHolding) SetRegistrar(v string)`

SetRegistrar sets Registrar field to given value.

### HasRegistrar

`func (o *MutualFundHolding) HasRegistrar() bool`

HasRegistrar returns a boolean if a field has been set.

### GetSchemeCode

`func (o *MutualFundHolding) GetSchemeCode() string`

GetSchemeCode returns the SchemeCode field if non-nil, zero value otherwise.

### GetSchemeCodeOk

`func (o *MutualFundHolding) GetSchemeCodeOk() (*string, bool)`

GetSchemeCodeOk returns a tuple with the SchemeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeCode

`func (o *MutualFundHolding) SetSchemeCode(v string)`

SetSchemeCode sets SchemeCode field to given value.

### HasSchemeCode

`func (o *MutualFundHolding) HasSchemeCode() bool`

HasSchemeCode returns a boolean if a field has been set.

### GetIsin

`func (o *MutualFundHolding) GetIsin() string`

GetIsin returns the Isin field if non-nil, zero value otherwise.

### GetIsinOk

`func (o *MutualFundHolding) GetIsinOk() (*string, bool)`

GetIsinOk returns a tuple with the Isin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsin

`func (o *MutualFundHolding) SetIsin(v string)`

SetIsin sets Isin field to given value.


### GetUcc

`func (o *MutualFundHolding) GetUcc() string`

GetUcc returns the Ucc field if non-nil, zero value otherwise.

### GetUccOk

`func (o *MutualFundHolding) GetUccOk() (*string, bool)`

GetUccOk returns a tuple with the Ucc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcc

`func (o *MutualFundHolding) SetUcc(v string)`

SetUcc sets Ucc field to given value.

### HasUcc

`func (o *MutualFundHolding) HasUcc() bool`

HasUcc returns a boolean if a field has been set.

### GetAmfiCode

`func (o *MutualFundHolding) GetAmfiCode() string`

GetAmfiCode returns the AmfiCode field if non-nil, zero value otherwise.

### GetAmfiCodeOk

`func (o *MutualFundHolding) GetAmfiCodeOk() (*string, bool)`

GetAmfiCodeOk returns a tuple with the AmfiCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfiCode

`func (o *MutualFundHolding) SetAmfiCode(v string)`

SetAmfiCode sets AmfiCode field to given value.

### HasAmfiCode

`func (o *MutualFundHolding) HasAmfiCode() bool`

HasAmfiCode returns a boolean if a field has been set.

### GetFolioNo

`func (o *MutualFundHolding) GetFolioNo() string`

GetFolioNo returns the FolioNo field if non-nil, zero value otherwise.

### GetFolioNoOk

`func (o *MutualFundHolding) GetFolioNoOk() (*string, bool)`

GetFolioNoOk returns a tuple with the FolioNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolioNo

`func (o *MutualFundHolding) SetFolioNo(v string)`

SetFolioNo sets FolioNo field to given value.


### GetDividendType

`func (o *MutualFundHolding) GetDividendType() string`

GetDividendType returns the DividendType field if non-nil, zero value otherwise.

### GetDividendTypeOk

`func (o *MutualFundHolding) GetDividendTypeOk() (*string, bool)`

GetDividendTypeOk returns a tuple with the DividendType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendType

`func (o *MutualFundHolding) SetDividendType(v string)`

SetDividendType sets DividendType field to given value.

### HasDividendType

`func (o *MutualFundHolding) HasDividendType() bool`

HasDividendType returns a boolean if a field has been set.

### GetFatcaStatus

`func (o *MutualFundHolding) GetFatcaStatus() string`

GetFatcaStatus returns the FatcaStatus field if non-nil, zero value otherwise.

### GetFatcaStatusOk

`func (o *MutualFundHolding) GetFatcaStatusOk() (*string, bool)`

GetFatcaStatusOk returns a tuple with the FatcaStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFatcaStatus

`func (o *MutualFundHolding) SetFatcaStatus(v string)`

SetFatcaStatus sets FatcaStatus field to given value.

### HasFatcaStatus

`func (o *MutualFundHolding) HasFatcaStatus() bool`

HasFatcaStatus returns a boolean if a field has been set.

### GetMode

`func (o *MutualFundHolding) GetMode() MutualFundHoldingMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *MutualFundHolding) GetModeOk() (*MutualFundHoldingMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *MutualFundHolding) SetMode(v MutualFundHoldingMode)`

SetMode sets Mode field to given value.


### GetUnits

`func (o *MutualFundHolding) GetUnits() float64`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *MutualFundHolding) GetUnitsOk() (*float64, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *MutualFundHolding) SetUnits(v float64)`

SetUnits sets Units field to given value.


### GetClosingUnits

`func (o *MutualFundHolding) GetClosingUnits() string`

GetClosingUnits returns the ClosingUnits field if non-nil, zero value otherwise.

### GetClosingUnitsOk

`func (o *MutualFundHolding) GetClosingUnitsOk() (*string, bool)`

GetClosingUnitsOk returns a tuple with the ClosingUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingUnits

`func (o *MutualFundHolding) SetClosingUnits(v string)`

SetClosingUnits sets ClosingUnits field to given value.

### HasClosingUnits

`func (o *MutualFundHolding) HasClosingUnits() bool`

HasClosingUnits returns a boolean if a field has been set.

### GetLienUnits

`func (o *MutualFundHolding) GetLienUnits() string`

GetLienUnits returns the LienUnits field if non-nil, zero value otherwise.

### GetLienUnitsOk

`func (o *MutualFundHolding) GetLienUnitsOk() (*string, bool)`

GetLienUnitsOk returns a tuple with the LienUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLienUnits

`func (o *MutualFundHolding) SetLienUnits(v string)`

SetLienUnits sets LienUnits field to given value.

### HasLienUnits

`func (o *MutualFundHolding) HasLienUnits() bool`

HasLienUnits returns a boolean if a field has been set.

### GetRate

`func (o *MutualFundHolding) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *MutualFundHolding) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *MutualFundHolding) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *MutualFundHolding) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetNav

`func (o *MutualFundHolding) GetNav() string`

GetNav returns the Nav field if non-nil, zero value otherwise.

### GetNavOk

`func (o *MutualFundHolding) GetNavOk() (*string, bool)`

GetNavOk returns a tuple with the Nav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNav

`func (o *MutualFundHolding) SetNav(v string)`

SetNav sets Nav field to given value.

### HasNav

`func (o *MutualFundHolding) HasNav() bool`

HasNav returns a boolean if a field has been set.

### GetLockingUnits

`func (o *MutualFundHolding) GetLockingUnits() string`

GetLockingUnits returns the LockingUnits field if non-nil, zero value otherwise.

### GetLockingUnitsOk

`func (o *MutualFundHolding) GetLockingUnitsOk() (*string, bool)`

GetLockingUnitsOk returns a tuple with the LockingUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockingUnits

`func (o *MutualFundHolding) SetLockingUnits(v string)`

SetLockingUnits sets LockingUnits field to given value.

### HasLockingUnits

`func (o *MutualFundHolding) HasLockingUnits() bool`

HasLockingUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



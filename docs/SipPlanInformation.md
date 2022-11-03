# SipPlanInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amc** | Pointer to **string** |  | [optional] 
**Registrar** | Pointer to **string** |  | [optional] 
**Scheme** | **string** |  | 
**Isin** | **string** |  | 
**FolioNumber** | Pointer to **string** |  | [optional] 
**Nav** | Pointer to **string** |  | [optional] 
**DividendType** | **string** |  | 
**CreationDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSipPlanInformation

`func NewSipPlanInformation(scheme string, isin string, dividendType string, ) *SipPlanInformation`

NewSipPlanInformation instantiates a new SipPlanInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSipPlanInformationWithDefaults

`func NewSipPlanInformationWithDefaults() *SipPlanInformation`

NewSipPlanInformationWithDefaults instantiates a new SipPlanInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmc

`func (o *SipPlanInformation) GetAmc() string`

GetAmc returns the Amc field if non-nil, zero value otherwise.

### GetAmcOk

`func (o *SipPlanInformation) GetAmcOk() (*string, bool)`

GetAmcOk returns a tuple with the Amc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmc

`func (o *SipPlanInformation) SetAmc(v string)`

SetAmc sets Amc field to given value.

### HasAmc

`func (o *SipPlanInformation) HasAmc() bool`

HasAmc returns a boolean if a field has been set.

### GetRegistrar

`func (o *SipPlanInformation) GetRegistrar() string`

GetRegistrar returns the Registrar field if non-nil, zero value otherwise.

### GetRegistrarOk

`func (o *SipPlanInformation) GetRegistrarOk() (*string, bool)`

GetRegistrarOk returns a tuple with the Registrar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrar

`func (o *SipPlanInformation) SetRegistrar(v string)`

SetRegistrar sets Registrar field to given value.

### HasRegistrar

`func (o *SipPlanInformation) HasRegistrar() bool`

HasRegistrar returns a boolean if a field has been set.

### GetScheme

`func (o *SipPlanInformation) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *SipPlanInformation) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *SipPlanInformation) SetScheme(v string)`

SetScheme sets Scheme field to given value.


### GetIsin

`func (o *SipPlanInformation) GetIsin() string`

GetIsin returns the Isin field if non-nil, zero value otherwise.

### GetIsinOk

`func (o *SipPlanInformation) GetIsinOk() (*string, bool)`

GetIsinOk returns a tuple with the Isin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsin

`func (o *SipPlanInformation) SetIsin(v string)`

SetIsin sets Isin field to given value.


### GetFolioNumber

`func (o *SipPlanInformation) GetFolioNumber() string`

GetFolioNumber returns the FolioNumber field if non-nil, zero value otherwise.

### GetFolioNumberOk

`func (o *SipPlanInformation) GetFolioNumberOk() (*string, bool)`

GetFolioNumberOk returns a tuple with the FolioNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolioNumber

`func (o *SipPlanInformation) SetFolioNumber(v string)`

SetFolioNumber sets FolioNumber field to given value.

### HasFolioNumber

`func (o *SipPlanInformation) HasFolioNumber() bool`

HasFolioNumber returns a boolean if a field has been set.

### GetNav

`func (o *SipPlanInformation) GetNav() string`

GetNav returns the Nav field if non-nil, zero value otherwise.

### GetNavOk

`func (o *SipPlanInformation) GetNavOk() (*string, bool)`

GetNavOk returns a tuple with the Nav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNav

`func (o *SipPlanInformation) SetNav(v string)`

SetNav sets Nav field to given value.

### HasNav

`func (o *SipPlanInformation) HasNav() bool`

HasNav returns a boolean if a field has been set.

### GetDividendType

`func (o *SipPlanInformation) GetDividendType() string`

GetDividendType returns the DividendType field if non-nil, zero value otherwise.

### GetDividendTypeOk

`func (o *SipPlanInformation) GetDividendTypeOk() (*string, bool)`

GetDividendTypeOk returns a tuple with the DividendType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendType

`func (o *SipPlanInformation) SetDividendType(v string)`

SetDividendType sets DividendType field to given value.


### GetCreationDate

`func (o *SipPlanInformation) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *SipPlanInformation) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *SipPlanInformation) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *SipPlanInformation) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



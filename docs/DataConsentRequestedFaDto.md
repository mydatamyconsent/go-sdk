# DataConsentRequestedFaDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Drn** | Pointer to **NullableString** |  | [optional] 
**FromDatetime** | Pointer to **time.Time** |  | [optional] 
**ToDatetime** | Pointer to **time.Time** |  | [optional] 
**ProviderId** | Pointer to **NullableString** |  | [optional] 
**AccountType** | Pointer to [**FinancialAccountTypes**](FinancialAccountTypes.md) |  | [optional] 
**AccountIdentifier** | Pointer to **NullableString** |  | [optional] 
**Filters** | Pointer to [**[]DataConsentRfaFilterDto**](DataConsentRfaFilterDto.md) |  | [optional] 

## Methods

### NewDataConsentRequestedFaDto

`func NewDataConsentRequestedFaDto() *DataConsentRequestedFaDto`

NewDataConsentRequestedFaDto instantiates a new DataConsentRequestedFaDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataConsentRequestedFaDtoWithDefaults

`func NewDataConsentRequestedFaDtoWithDefaults() *DataConsentRequestedFaDto`

NewDataConsentRequestedFaDtoWithDefaults instantiates a new DataConsentRequestedFaDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrn

`func (o *DataConsentRequestedFaDto) GetDrn() string`

GetDrn returns the Drn field if non-nil, zero value otherwise.

### GetDrnOk

`func (o *DataConsentRequestedFaDto) GetDrnOk() (*string, bool)`

GetDrnOk returns a tuple with the Drn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrn

`func (o *DataConsentRequestedFaDto) SetDrn(v string)`

SetDrn sets Drn field to given value.

### HasDrn

`func (o *DataConsentRequestedFaDto) HasDrn() bool`

HasDrn returns a boolean if a field has been set.

### SetDrnNil

`func (o *DataConsentRequestedFaDto) SetDrnNil(b bool)`

 SetDrnNil sets the value for Drn to be an explicit nil

### UnsetDrn
`func (o *DataConsentRequestedFaDto) UnsetDrn()`

UnsetDrn ensures that no value is present for Drn, not even an explicit nil
### GetFromDatetime

`func (o *DataConsentRequestedFaDto) GetFromDatetime() time.Time`

GetFromDatetime returns the FromDatetime field if non-nil, zero value otherwise.

### GetFromDatetimeOk

`func (o *DataConsentRequestedFaDto) GetFromDatetimeOk() (*time.Time, bool)`

GetFromDatetimeOk returns a tuple with the FromDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDatetime

`func (o *DataConsentRequestedFaDto) SetFromDatetime(v time.Time)`

SetFromDatetime sets FromDatetime field to given value.

### HasFromDatetime

`func (o *DataConsentRequestedFaDto) HasFromDatetime() bool`

HasFromDatetime returns a boolean if a field has been set.

### GetToDatetime

`func (o *DataConsentRequestedFaDto) GetToDatetime() time.Time`

GetToDatetime returns the ToDatetime field if non-nil, zero value otherwise.

### GetToDatetimeOk

`func (o *DataConsentRequestedFaDto) GetToDatetimeOk() (*time.Time, bool)`

GetToDatetimeOk returns a tuple with the ToDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDatetime

`func (o *DataConsentRequestedFaDto) SetToDatetime(v time.Time)`

SetToDatetime sets ToDatetime field to given value.

### HasToDatetime

`func (o *DataConsentRequestedFaDto) HasToDatetime() bool`

HasToDatetime returns a boolean if a field has been set.

### GetProviderId

`func (o *DataConsentRequestedFaDto) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *DataConsentRequestedFaDto) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *DataConsentRequestedFaDto) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *DataConsentRequestedFaDto) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *DataConsentRequestedFaDto) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *DataConsentRequestedFaDto) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
### GetAccountType

`func (o *DataConsentRequestedFaDto) GetAccountType() FinancialAccountTypes`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *DataConsentRequestedFaDto) GetAccountTypeOk() (*FinancialAccountTypes, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *DataConsentRequestedFaDto) SetAccountType(v FinancialAccountTypes)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *DataConsentRequestedFaDto) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetAccountIdentifier

`func (o *DataConsentRequestedFaDto) GetAccountIdentifier() string`

GetAccountIdentifier returns the AccountIdentifier field if non-nil, zero value otherwise.

### GetAccountIdentifierOk

`func (o *DataConsentRequestedFaDto) GetAccountIdentifierOk() (*string, bool)`

GetAccountIdentifierOk returns a tuple with the AccountIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIdentifier

`func (o *DataConsentRequestedFaDto) SetAccountIdentifier(v string)`

SetAccountIdentifier sets AccountIdentifier field to given value.

### HasAccountIdentifier

`func (o *DataConsentRequestedFaDto) HasAccountIdentifier() bool`

HasAccountIdentifier returns a boolean if a field has been set.

### SetAccountIdentifierNil

`func (o *DataConsentRequestedFaDto) SetAccountIdentifierNil(b bool)`

 SetAccountIdentifierNil sets the value for AccountIdentifier to be an explicit nil

### UnsetAccountIdentifier
`func (o *DataConsentRequestedFaDto) UnsetAccountIdentifier()`

UnsetAccountIdentifier ensures that no value is present for AccountIdentifier, not even an explicit nil
### GetFilters

`func (o *DataConsentRequestedFaDto) GetFilters() []DataConsentRfaFilterDto`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *DataConsentRequestedFaDto) GetFiltersOk() (*[]DataConsentRfaFilterDto, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *DataConsentRequestedFaDto) SetFilters(v []DataConsentRfaFilterDto)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *DataConsentRequestedFaDto) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *DataConsentRequestedFaDto) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *DataConsentRequestedFaDto) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



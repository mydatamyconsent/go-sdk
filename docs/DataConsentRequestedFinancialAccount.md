# DataConsentRequestedFinancialAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Drn** | Pointer to **NullableString** |  | [optional] 
**FromDatetime** | Pointer to **time.Time** |  | [optional] 
**ToDatetime** | Pointer to **time.Time** |  | [optional] 
**ProviderId** | Pointer to **NullableString** |  | [optional] 
**AccountType** | Pointer to [**FinancialAccountTypes**](FinancialAccountTypes.md) |  | [optional] 
**AccountIdentifier** | Pointer to **NullableString** |  | [optional] 
**Filters** | Pointer to [**[]DataConsentRfaFilter**](DataConsentRfaFilter.md) |  | [optional] 

## Methods

### NewDataConsentRequestedFinancialAccount

`func NewDataConsentRequestedFinancialAccount() *DataConsentRequestedFinancialAccount`

NewDataConsentRequestedFinancialAccount instantiates a new DataConsentRequestedFinancialAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataConsentRequestedFinancialAccountWithDefaults

`func NewDataConsentRequestedFinancialAccountWithDefaults() *DataConsentRequestedFinancialAccount`

NewDataConsentRequestedFinancialAccountWithDefaults instantiates a new DataConsentRequestedFinancialAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrn

`func (o *DataConsentRequestedFinancialAccount) GetDrn() string`

GetDrn returns the Drn field if non-nil, zero value otherwise.

### GetDrnOk

`func (o *DataConsentRequestedFinancialAccount) GetDrnOk() (*string, bool)`

GetDrnOk returns a tuple with the Drn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrn

`func (o *DataConsentRequestedFinancialAccount) SetDrn(v string)`

SetDrn sets Drn field to given value.

### HasDrn

`func (o *DataConsentRequestedFinancialAccount) HasDrn() bool`

HasDrn returns a boolean if a field has been set.

### SetDrnNil

`func (o *DataConsentRequestedFinancialAccount) SetDrnNil(b bool)`

 SetDrnNil sets the value for Drn to be an explicit nil

### UnsetDrn
`func (o *DataConsentRequestedFinancialAccount) UnsetDrn()`

UnsetDrn ensures that no value is present for Drn, not even an explicit nil
### GetFromDatetime

`func (o *DataConsentRequestedFinancialAccount) GetFromDatetime() time.Time`

GetFromDatetime returns the FromDatetime field if non-nil, zero value otherwise.

### GetFromDatetimeOk

`func (o *DataConsentRequestedFinancialAccount) GetFromDatetimeOk() (*time.Time, bool)`

GetFromDatetimeOk returns a tuple with the FromDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDatetime

`func (o *DataConsentRequestedFinancialAccount) SetFromDatetime(v time.Time)`

SetFromDatetime sets FromDatetime field to given value.

### HasFromDatetime

`func (o *DataConsentRequestedFinancialAccount) HasFromDatetime() bool`

HasFromDatetime returns a boolean if a field has been set.

### GetToDatetime

`func (o *DataConsentRequestedFinancialAccount) GetToDatetime() time.Time`

GetToDatetime returns the ToDatetime field if non-nil, zero value otherwise.

### GetToDatetimeOk

`func (o *DataConsentRequestedFinancialAccount) GetToDatetimeOk() (*time.Time, bool)`

GetToDatetimeOk returns a tuple with the ToDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDatetime

`func (o *DataConsentRequestedFinancialAccount) SetToDatetime(v time.Time)`

SetToDatetime sets ToDatetime field to given value.

### HasToDatetime

`func (o *DataConsentRequestedFinancialAccount) HasToDatetime() bool`

HasToDatetime returns a boolean if a field has been set.

### GetProviderId

`func (o *DataConsentRequestedFinancialAccount) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *DataConsentRequestedFinancialAccount) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *DataConsentRequestedFinancialAccount) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *DataConsentRequestedFinancialAccount) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *DataConsentRequestedFinancialAccount) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *DataConsentRequestedFinancialAccount) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
### GetAccountType

`func (o *DataConsentRequestedFinancialAccount) GetAccountType() FinancialAccountTypes`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *DataConsentRequestedFinancialAccount) GetAccountTypeOk() (*FinancialAccountTypes, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *DataConsentRequestedFinancialAccount) SetAccountType(v FinancialAccountTypes)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *DataConsentRequestedFinancialAccount) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetAccountIdentifier

`func (o *DataConsentRequestedFinancialAccount) GetAccountIdentifier() string`

GetAccountIdentifier returns the AccountIdentifier field if non-nil, zero value otherwise.

### GetAccountIdentifierOk

`func (o *DataConsentRequestedFinancialAccount) GetAccountIdentifierOk() (*string, bool)`

GetAccountIdentifierOk returns a tuple with the AccountIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIdentifier

`func (o *DataConsentRequestedFinancialAccount) SetAccountIdentifier(v string)`

SetAccountIdentifier sets AccountIdentifier field to given value.

### HasAccountIdentifier

`func (o *DataConsentRequestedFinancialAccount) HasAccountIdentifier() bool`

HasAccountIdentifier returns a boolean if a field has been set.

### SetAccountIdentifierNil

`func (o *DataConsentRequestedFinancialAccount) SetAccountIdentifierNil(b bool)`

 SetAccountIdentifierNil sets the value for AccountIdentifier to be an explicit nil

### UnsetAccountIdentifier
`func (o *DataConsentRequestedFinancialAccount) UnsetAccountIdentifier()`

UnsetAccountIdentifier ensures that no value is present for AccountIdentifier, not even an explicit nil
### GetFilters

`func (o *DataConsentRequestedFinancialAccount) GetFilters() []DataConsentRfaFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *DataConsentRequestedFinancialAccount) GetFiltersOk() (*[]DataConsentRfaFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *DataConsentRequestedFinancialAccount) SetFilters(v []DataConsentRfaFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *DataConsentRequestedFinancialAccount) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *DataConsentRequestedFinancialAccount) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *DataConsentRequestedFinancialAccount) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



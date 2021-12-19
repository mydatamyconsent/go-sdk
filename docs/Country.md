# Country

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Iso3** | Pointer to **NullableString** |  | [optional] 
**Iso2** | Pointer to **NullableString** |  | [optional] 
**PhoneCode** | Pointer to **NullableString** |  | [optional] 
**Capital** | Pointer to **NullableString** |  | [optional] 
**CurrencyCode** | Pointer to **NullableString** |  | [optional] 
**CurrencySymbol** | Pointer to **NullableString** |  | [optional] 
**FlagUrl** | Pointer to **NullableString** |  | [optional] 
**States** | Pointer to [**[]CountryState**](CountryState.md) |  | [optional] 

## Methods

### NewCountry

`func NewCountry() *Country`

NewCountry instantiates a new Country object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryWithDefaults

`func NewCountryWithDefaults() *Country`

NewCountryWithDefaults instantiates a new Country object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Country) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Country) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Country) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Country) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Country) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Country) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Country) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Country) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Country) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Country) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIso3

`func (o *Country) GetIso3() string`

GetIso3 returns the Iso3 field if non-nil, zero value otherwise.

### GetIso3Ok

`func (o *Country) GetIso3Ok() (*string, bool)`

GetIso3Ok returns a tuple with the Iso3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso3

`func (o *Country) SetIso3(v string)`

SetIso3 sets Iso3 field to given value.

### HasIso3

`func (o *Country) HasIso3() bool`

HasIso3 returns a boolean if a field has been set.

### SetIso3Nil

`func (o *Country) SetIso3Nil(b bool)`

 SetIso3Nil sets the value for Iso3 to be an explicit nil

### UnsetIso3
`func (o *Country) UnsetIso3()`

UnsetIso3 ensures that no value is present for Iso3, not even an explicit nil
### GetIso2

`func (o *Country) GetIso2() string`

GetIso2 returns the Iso2 field if non-nil, zero value otherwise.

### GetIso2Ok

`func (o *Country) GetIso2Ok() (*string, bool)`

GetIso2Ok returns a tuple with the Iso2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso2

`func (o *Country) SetIso2(v string)`

SetIso2 sets Iso2 field to given value.

### HasIso2

`func (o *Country) HasIso2() bool`

HasIso2 returns a boolean if a field has been set.

### SetIso2Nil

`func (o *Country) SetIso2Nil(b bool)`

 SetIso2Nil sets the value for Iso2 to be an explicit nil

### UnsetIso2
`func (o *Country) UnsetIso2()`

UnsetIso2 ensures that no value is present for Iso2, not even an explicit nil
### GetPhoneCode

`func (o *Country) GetPhoneCode() string`

GetPhoneCode returns the PhoneCode field if non-nil, zero value otherwise.

### GetPhoneCodeOk

`func (o *Country) GetPhoneCodeOk() (*string, bool)`

GetPhoneCodeOk returns a tuple with the PhoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneCode

`func (o *Country) SetPhoneCode(v string)`

SetPhoneCode sets PhoneCode field to given value.

### HasPhoneCode

`func (o *Country) HasPhoneCode() bool`

HasPhoneCode returns a boolean if a field has been set.

### SetPhoneCodeNil

`func (o *Country) SetPhoneCodeNil(b bool)`

 SetPhoneCodeNil sets the value for PhoneCode to be an explicit nil

### UnsetPhoneCode
`func (o *Country) UnsetPhoneCode()`

UnsetPhoneCode ensures that no value is present for PhoneCode, not even an explicit nil
### GetCapital

`func (o *Country) GetCapital() string`

GetCapital returns the Capital field if non-nil, zero value otherwise.

### GetCapitalOk

`func (o *Country) GetCapitalOk() (*string, bool)`

GetCapitalOk returns a tuple with the Capital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapital

`func (o *Country) SetCapital(v string)`

SetCapital sets Capital field to given value.

### HasCapital

`func (o *Country) HasCapital() bool`

HasCapital returns a boolean if a field has been set.

### SetCapitalNil

`func (o *Country) SetCapitalNil(b bool)`

 SetCapitalNil sets the value for Capital to be an explicit nil

### UnsetCapital
`func (o *Country) UnsetCapital()`

UnsetCapital ensures that no value is present for Capital, not even an explicit nil
### GetCurrencyCode

`func (o *Country) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *Country) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *Country) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *Country) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *Country) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *Country) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetCurrencySymbol

`func (o *Country) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *Country) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *Country) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.

### HasCurrencySymbol

`func (o *Country) HasCurrencySymbol() bool`

HasCurrencySymbol returns a boolean if a field has been set.

### SetCurrencySymbolNil

`func (o *Country) SetCurrencySymbolNil(b bool)`

 SetCurrencySymbolNil sets the value for CurrencySymbol to be an explicit nil

### UnsetCurrencySymbol
`func (o *Country) UnsetCurrencySymbol()`

UnsetCurrencySymbol ensures that no value is present for CurrencySymbol, not even an explicit nil
### GetFlagUrl

`func (o *Country) GetFlagUrl() string`

GetFlagUrl returns the FlagUrl field if non-nil, zero value otherwise.

### GetFlagUrlOk

`func (o *Country) GetFlagUrlOk() (*string, bool)`

GetFlagUrlOk returns a tuple with the FlagUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagUrl

`func (o *Country) SetFlagUrl(v string)`

SetFlagUrl sets FlagUrl field to given value.

### HasFlagUrl

`func (o *Country) HasFlagUrl() bool`

HasFlagUrl returns a boolean if a field has been set.

### SetFlagUrlNil

`func (o *Country) SetFlagUrlNil(b bool)`

 SetFlagUrlNil sets the value for FlagUrl to be an explicit nil

### UnsetFlagUrl
`func (o *Country) UnsetFlagUrl()`

UnsetFlagUrl ensures that no value is present for FlagUrl, not even an explicit nil
### GetStates

`func (o *Country) GetStates() []CountryState`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *Country) GetStatesOk() (*[]CountryState, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *Country) SetStates(v []CountryState)`

SetStates sets States field to given value.

### HasStates

`func (o *Country) HasStates() bool`

HasStates returns a boolean if a field has been set.

### SetStatesNil

`func (o *Country) SetStatesNil(b bool)`

 SetStatesNil sets the value for States to be an explicit nil

### UnsetStates
`func (o *Country) UnsetStates()`

UnsetStates ensures that no value is present for States, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



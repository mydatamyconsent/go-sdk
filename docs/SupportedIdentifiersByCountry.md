# SupportedIdentifiersByCountry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iso2** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Flag** | Pointer to **NullableString** |  | [optional] 
**IndividualIdentifiers** | Pointer to [**[]SupportedIdentifier**](SupportedIdentifier.md) |  | [optional] 
**OrganizationIdentifiers** | Pointer to [**[]SupportedIdentifier**](SupportedIdentifier.md) |  | [optional] 

## Methods

### NewSupportedIdentifiersByCountry

`func NewSupportedIdentifiersByCountry() *SupportedIdentifiersByCountry`

NewSupportedIdentifiersByCountry instantiates a new SupportedIdentifiersByCountry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedIdentifiersByCountryWithDefaults

`func NewSupportedIdentifiersByCountryWithDefaults() *SupportedIdentifiersByCountry`

NewSupportedIdentifiersByCountryWithDefaults instantiates a new SupportedIdentifiersByCountry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIso2

`func (o *SupportedIdentifiersByCountry) GetIso2() string`

GetIso2 returns the Iso2 field if non-nil, zero value otherwise.

### GetIso2Ok

`func (o *SupportedIdentifiersByCountry) GetIso2Ok() (*string, bool)`

GetIso2Ok returns a tuple with the Iso2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso2

`func (o *SupportedIdentifiersByCountry) SetIso2(v string)`

SetIso2 sets Iso2 field to given value.

### HasIso2

`func (o *SupportedIdentifiersByCountry) HasIso2() bool`

HasIso2 returns a boolean if a field has been set.

### SetIso2Nil

`func (o *SupportedIdentifiersByCountry) SetIso2Nil(b bool)`

 SetIso2Nil sets the value for Iso2 to be an explicit nil

### UnsetIso2
`func (o *SupportedIdentifiersByCountry) UnsetIso2()`

UnsetIso2 ensures that no value is present for Iso2, not even an explicit nil
### GetName

`func (o *SupportedIdentifiersByCountry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SupportedIdentifiersByCountry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SupportedIdentifiersByCountry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SupportedIdentifiersByCountry) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SupportedIdentifiersByCountry) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SupportedIdentifiersByCountry) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFlag

`func (o *SupportedIdentifiersByCountry) GetFlag() string`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *SupportedIdentifiersByCountry) GetFlagOk() (*string, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *SupportedIdentifiersByCountry) SetFlag(v string)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *SupportedIdentifiersByCountry) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### SetFlagNil

`func (o *SupportedIdentifiersByCountry) SetFlagNil(b bool)`

 SetFlagNil sets the value for Flag to be an explicit nil

### UnsetFlag
`func (o *SupportedIdentifiersByCountry) UnsetFlag()`

UnsetFlag ensures that no value is present for Flag, not even an explicit nil
### GetIndividualIdentifiers

`func (o *SupportedIdentifiersByCountry) GetIndividualIdentifiers() []SupportedIdentifier`

GetIndividualIdentifiers returns the IndividualIdentifiers field if non-nil, zero value otherwise.

### GetIndividualIdentifiersOk

`func (o *SupportedIdentifiersByCountry) GetIndividualIdentifiersOk() (*[]SupportedIdentifier, bool)`

GetIndividualIdentifiersOk returns a tuple with the IndividualIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualIdentifiers

`func (o *SupportedIdentifiersByCountry) SetIndividualIdentifiers(v []SupportedIdentifier)`

SetIndividualIdentifiers sets IndividualIdentifiers field to given value.

### HasIndividualIdentifiers

`func (o *SupportedIdentifiersByCountry) HasIndividualIdentifiers() bool`

HasIndividualIdentifiers returns a boolean if a field has been set.

### SetIndividualIdentifiersNil

`func (o *SupportedIdentifiersByCountry) SetIndividualIdentifiersNil(b bool)`

 SetIndividualIdentifiersNil sets the value for IndividualIdentifiers to be an explicit nil

### UnsetIndividualIdentifiers
`func (o *SupportedIdentifiersByCountry) UnsetIndividualIdentifiers()`

UnsetIndividualIdentifiers ensures that no value is present for IndividualIdentifiers, not even an explicit nil
### GetOrganizationIdentifiers

`func (o *SupportedIdentifiersByCountry) GetOrganizationIdentifiers() []SupportedIdentifier`

GetOrganizationIdentifiers returns the OrganizationIdentifiers field if non-nil, zero value otherwise.

### GetOrganizationIdentifiersOk

`func (o *SupportedIdentifiersByCountry) GetOrganizationIdentifiersOk() (*[]SupportedIdentifier, bool)`

GetOrganizationIdentifiersOk returns a tuple with the OrganizationIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationIdentifiers

`func (o *SupportedIdentifiersByCountry) SetOrganizationIdentifiers(v []SupportedIdentifier)`

SetOrganizationIdentifiers sets OrganizationIdentifiers field to given value.

### HasOrganizationIdentifiers

`func (o *SupportedIdentifiersByCountry) HasOrganizationIdentifiers() bool`

HasOrganizationIdentifiers returns a boolean if a field has been set.

### SetOrganizationIdentifiersNil

`func (o *SupportedIdentifiersByCountry) SetOrganizationIdentifiersNil(b bool)`

 SetOrganizationIdentifiersNil sets the value for OrganizationIdentifiers to be an explicit nil

### UnsetOrganizationIdentifiers
`func (o *SupportedIdentifiersByCountry) UnsetOrganizationIdentifiers()`

UnsetOrganizationIdentifiers ensures that no value is present for OrganizationIdentifiers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



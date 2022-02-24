# SupportedIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iso2** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**IndividualIdentifiers** | Pointer to [**[]Identifier**](Identifier.md) |  | [optional] 
**OrganizationIdentifiers** | Pointer to [**[]Identifier**](Identifier.md) |  | [optional] 

## Methods

### NewSupportedIdentifier

`func NewSupportedIdentifier() *SupportedIdentifier`

NewSupportedIdentifier instantiates a new SupportedIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedIdentifierWithDefaults

`func NewSupportedIdentifierWithDefaults() *SupportedIdentifier`

NewSupportedIdentifierWithDefaults instantiates a new SupportedIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIso2

`func (o *SupportedIdentifier) GetIso2() string`

GetIso2 returns the Iso2 field if non-nil, zero value otherwise.

### GetIso2Ok

`func (o *SupportedIdentifier) GetIso2Ok() (*string, bool)`

GetIso2Ok returns a tuple with the Iso2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso2

`func (o *SupportedIdentifier) SetIso2(v string)`

SetIso2 sets Iso2 field to given value.

### HasIso2

`func (o *SupportedIdentifier) HasIso2() bool`

HasIso2 returns a boolean if a field has been set.

### SetIso2Nil

`func (o *SupportedIdentifier) SetIso2Nil(b bool)`

 SetIso2Nil sets the value for Iso2 to be an explicit nil

### UnsetIso2
`func (o *SupportedIdentifier) UnsetIso2()`

UnsetIso2 ensures that no value is present for Iso2, not even an explicit nil
### GetName

`func (o *SupportedIdentifier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SupportedIdentifier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SupportedIdentifier) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SupportedIdentifier) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SupportedIdentifier) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SupportedIdentifier) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIndividualIdentifiers

`func (o *SupportedIdentifier) GetIndividualIdentifiers() []Identifier`

GetIndividualIdentifiers returns the IndividualIdentifiers field if non-nil, zero value otherwise.

### GetIndividualIdentifiersOk

`func (o *SupportedIdentifier) GetIndividualIdentifiersOk() (*[]Identifier, bool)`

GetIndividualIdentifiersOk returns a tuple with the IndividualIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualIdentifiers

`func (o *SupportedIdentifier) SetIndividualIdentifiers(v []Identifier)`

SetIndividualIdentifiers sets IndividualIdentifiers field to given value.

### HasIndividualIdentifiers

`func (o *SupportedIdentifier) HasIndividualIdentifiers() bool`

HasIndividualIdentifiers returns a boolean if a field has been set.

### SetIndividualIdentifiersNil

`func (o *SupportedIdentifier) SetIndividualIdentifiersNil(b bool)`

 SetIndividualIdentifiersNil sets the value for IndividualIdentifiers to be an explicit nil

### UnsetIndividualIdentifiers
`func (o *SupportedIdentifier) UnsetIndividualIdentifiers()`

UnsetIndividualIdentifiers ensures that no value is present for IndividualIdentifiers, not even an explicit nil
### GetOrganizationIdentifiers

`func (o *SupportedIdentifier) GetOrganizationIdentifiers() []Identifier`

GetOrganizationIdentifiers returns the OrganizationIdentifiers field if non-nil, zero value otherwise.

### GetOrganizationIdentifiersOk

`func (o *SupportedIdentifier) GetOrganizationIdentifiersOk() (*[]Identifier, bool)`

GetOrganizationIdentifiersOk returns a tuple with the OrganizationIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationIdentifiers

`func (o *SupportedIdentifier) SetOrganizationIdentifiers(v []Identifier)`

SetOrganizationIdentifiers sets OrganizationIdentifiers field to given value.

### HasOrganizationIdentifiers

`func (o *SupportedIdentifier) HasOrganizationIdentifiers() bool`

HasOrganizationIdentifiers returns a boolean if a field has been set.

### SetOrganizationIdentifiersNil

`func (o *SupportedIdentifier) SetOrganizationIdentifiersNil(b bool)`

 SetOrganizationIdentifiersNil sets the value for OrganizationIdentifiers to be an explicit nil

### UnsetOrganizationIdentifiers
`func (o *SupportedIdentifier) UnsetOrganizationIdentifiers()`

UnsetOrganizationIdentifiers ensures that no value is present for OrganizationIdentifiers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



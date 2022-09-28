# SupportedIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iso2** | **string** | Country ISO 2 code. Example: IN, US, etc. | 
**Name** | **string** | Country name. Example: India, United States of America, etc. | 
**IndividualIdentifiers** | [**[]Identifier**](Identifier.md) | List of supported identifiers for an individual. | 
**OrganizationIdentifiers** | [**[]Identifier**](Identifier.md) | List of supported identifiers for an organization. | 

## Methods

### NewSupportedIdentifier

`func NewSupportedIdentifier(iso2 string, name string, individualIdentifiers []Identifier, organizationIdentifiers []Identifier, ) *SupportedIdentifier`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



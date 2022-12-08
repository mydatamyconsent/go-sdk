# AllergyIntoleranceReaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Manifestation** | [**[]CodeableConcept**](CodeableConcept.md) |  | 

## Methods

### NewAllergyIntoleranceReaction

`func NewAllergyIntoleranceReaction(manifestation []CodeableConcept, ) *AllergyIntoleranceReaction`

NewAllergyIntoleranceReaction instantiates a new AllergyIntoleranceReaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllergyIntoleranceReactionWithDefaults

`func NewAllergyIntoleranceReactionWithDefaults() *AllergyIntoleranceReaction`

NewAllergyIntoleranceReactionWithDefaults instantiates a new AllergyIntoleranceReaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AllergyIntoleranceReaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AllergyIntoleranceReaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AllergyIntoleranceReaction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AllergyIntoleranceReaction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetManifestation

`func (o *AllergyIntoleranceReaction) GetManifestation() []CodeableConcept`

GetManifestation returns the Manifestation field if non-nil, zero value otherwise.

### GetManifestationOk

`func (o *AllergyIntoleranceReaction) GetManifestationOk() (*[]CodeableConcept, bool)`

GetManifestationOk returns a tuple with the Manifestation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestation

`func (o *AllergyIntoleranceReaction) SetManifestation(v []CodeableConcept)`

SetManifestation sets Manifestation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



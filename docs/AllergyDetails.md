# AllergyDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to [**CodeableConcept**](CodeableConcept.md) |  | [optional] 
**Category** | **[]string** |  | 
**LevelStatus** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**CodeableConcept**](CodeableConcept.md) |  | [optional] 
**ClinicalStatus** | Pointer to [**CodeableConcept**](CodeableConcept.md) |  | [optional] 
**Reactions** | [**[]AllergyIntoleranceReaction**](AllergyIntoleranceReaction.md) |  | 
**RecordedAtUtc** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAllergyDetails

`func NewAllergyDetails(id string, category []string, reactions []AllergyIntoleranceReaction, ) *AllergyDetails`

NewAllergyDetails instantiates a new AllergyDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllergyDetailsWithDefaults

`func NewAllergyDetailsWithDefaults() *AllergyDetails`

NewAllergyDetailsWithDefaults instantiates a new AllergyDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AllergyDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AllergyDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AllergyDetails) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AllergyDetails) GetName() CodeableConcept`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AllergyDetails) GetNameOk() (*CodeableConcept, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AllergyDetails) SetName(v CodeableConcept)`

SetName sets Name field to given value.

### HasName

`func (o *AllergyDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategory

`func (o *AllergyDetails) GetCategory() []string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AllergyDetails) GetCategoryOk() (*[]string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AllergyDetails) SetCategory(v []string)`

SetCategory sets Category field to given value.


### GetLevelStatus

`func (o *AllergyDetails) GetLevelStatus() string`

GetLevelStatus returns the LevelStatus field if non-nil, zero value otherwise.

### GetLevelStatusOk

`func (o *AllergyDetails) GetLevelStatusOk() (*string, bool)`

GetLevelStatusOk returns a tuple with the LevelStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelStatus

`func (o *AllergyDetails) SetLevelStatus(v string)`

SetLevelStatus sets LevelStatus field to given value.

### HasLevelStatus

`func (o *AllergyDetails) HasLevelStatus() bool`

HasLevelStatus returns a boolean if a field has been set.

### GetStatus

`func (o *AllergyDetails) GetStatus() CodeableConcept`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AllergyDetails) GetStatusOk() (*CodeableConcept, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AllergyDetails) SetStatus(v CodeableConcept)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AllergyDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetClinicalStatus

`func (o *AllergyDetails) GetClinicalStatus() CodeableConcept`

GetClinicalStatus returns the ClinicalStatus field if non-nil, zero value otherwise.

### GetClinicalStatusOk

`func (o *AllergyDetails) GetClinicalStatusOk() (*CodeableConcept, bool)`

GetClinicalStatusOk returns a tuple with the ClinicalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClinicalStatus

`func (o *AllergyDetails) SetClinicalStatus(v CodeableConcept)`

SetClinicalStatus sets ClinicalStatus field to given value.

### HasClinicalStatus

`func (o *AllergyDetails) HasClinicalStatus() bool`

HasClinicalStatus returns a boolean if a field has been set.

### GetReactions

`func (o *AllergyDetails) GetReactions() []AllergyIntoleranceReaction`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *AllergyDetails) GetReactionsOk() (*[]AllergyIntoleranceReaction, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *AllergyDetails) SetReactions(v []AllergyIntoleranceReaction)`

SetReactions sets Reactions field to given value.


### GetRecordedAtUtc

`func (o *AllergyDetails) GetRecordedAtUtc() time.Time`

GetRecordedAtUtc returns the RecordedAtUtc field if non-nil, zero value otherwise.

### GetRecordedAtUtcOk

`func (o *AllergyDetails) GetRecordedAtUtcOk() (*time.Time, bool)`

GetRecordedAtUtcOk returns a tuple with the RecordedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedAtUtc

`func (o *AllergyDetails) SetRecordedAtUtc(v time.Time)`

SetRecordedAtUtc sets RecordedAtUtc field to given value.

### HasRecordedAtUtc

`func (o *AllergyDetails) HasRecordedAtUtc() bool`

HasRecordedAtUtc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



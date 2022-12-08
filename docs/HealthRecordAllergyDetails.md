# HealthRecordAllergyDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Name** | Pointer to [**CodeableConcept**](CodeableConcept.md) |  | [optional] 
**Category** | **[]string** |  | 
**LevelStatus** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**CodeableConcept**](CodeableConcept.md) |  | [optional] 
**ClinicalStatus** | Pointer to [**CodeableConcept**](CodeableConcept.md) |  | [optional] 
**Reactions** | [**[]AllergyIntoleranceReaction**](AllergyIntoleranceReaction.md) |  | 
**RecordedAtUtc** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewHealthRecordAllergyDetails

`func NewHealthRecordAllergyDetails(type_ string, id string, category []string, reactions []AllergyIntoleranceReaction, ) *HealthRecordAllergyDetails`

NewHealthRecordAllergyDetails instantiates a new HealthRecordAllergyDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthRecordAllergyDetailsWithDefaults

`func NewHealthRecordAllergyDetailsWithDefaults() *HealthRecordAllergyDetails`

NewHealthRecordAllergyDetailsWithDefaults instantiates a new HealthRecordAllergyDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *HealthRecordAllergyDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HealthRecordAllergyDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HealthRecordAllergyDetails) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *HealthRecordAllergyDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HealthRecordAllergyDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HealthRecordAllergyDetails) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *HealthRecordAllergyDetails) GetName() CodeableConcept`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HealthRecordAllergyDetails) GetNameOk() (*CodeableConcept, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HealthRecordAllergyDetails) SetName(v CodeableConcept)`

SetName sets Name field to given value.

### HasName

`func (o *HealthRecordAllergyDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategory

`func (o *HealthRecordAllergyDetails) GetCategory() []string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *HealthRecordAllergyDetails) GetCategoryOk() (*[]string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *HealthRecordAllergyDetails) SetCategory(v []string)`

SetCategory sets Category field to given value.


### GetLevelStatus

`func (o *HealthRecordAllergyDetails) GetLevelStatus() string`

GetLevelStatus returns the LevelStatus field if non-nil, zero value otherwise.

### GetLevelStatusOk

`func (o *HealthRecordAllergyDetails) GetLevelStatusOk() (*string, bool)`

GetLevelStatusOk returns a tuple with the LevelStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelStatus

`func (o *HealthRecordAllergyDetails) SetLevelStatus(v string)`

SetLevelStatus sets LevelStatus field to given value.

### HasLevelStatus

`func (o *HealthRecordAllergyDetails) HasLevelStatus() bool`

HasLevelStatus returns a boolean if a field has been set.

### GetStatus

`func (o *HealthRecordAllergyDetails) GetStatus() CodeableConcept`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthRecordAllergyDetails) GetStatusOk() (*CodeableConcept, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthRecordAllergyDetails) SetStatus(v CodeableConcept)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HealthRecordAllergyDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetClinicalStatus

`func (o *HealthRecordAllergyDetails) GetClinicalStatus() CodeableConcept`

GetClinicalStatus returns the ClinicalStatus field if non-nil, zero value otherwise.

### GetClinicalStatusOk

`func (o *HealthRecordAllergyDetails) GetClinicalStatusOk() (*CodeableConcept, bool)`

GetClinicalStatusOk returns a tuple with the ClinicalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClinicalStatus

`func (o *HealthRecordAllergyDetails) SetClinicalStatus(v CodeableConcept)`

SetClinicalStatus sets ClinicalStatus field to given value.

### HasClinicalStatus

`func (o *HealthRecordAllergyDetails) HasClinicalStatus() bool`

HasClinicalStatus returns a boolean if a field has been set.

### GetReactions

`func (o *HealthRecordAllergyDetails) GetReactions() []AllergyIntoleranceReaction`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *HealthRecordAllergyDetails) GetReactionsOk() (*[]AllergyIntoleranceReaction, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *HealthRecordAllergyDetails) SetReactions(v []AllergyIntoleranceReaction)`

SetReactions sets Reactions field to given value.


### GetRecordedAtUtc

`func (o *HealthRecordAllergyDetails) GetRecordedAtUtc() time.Time`

GetRecordedAtUtc returns the RecordedAtUtc field if non-nil, zero value otherwise.

### GetRecordedAtUtcOk

`func (o *HealthRecordAllergyDetails) GetRecordedAtUtcOk() (*time.Time, bool)`

GetRecordedAtUtcOk returns a tuple with the RecordedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedAtUtc

`func (o *HealthRecordAllergyDetails) SetRecordedAtUtc(v time.Time)`

SetRecordedAtUtc sets RecordedAtUtc field to given value.

### HasRecordedAtUtc

`func (o *HealthRecordAllergyDetails) HasRecordedAtUtc() bool`

HasRecordedAtUtc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



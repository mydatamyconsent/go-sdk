# HealthRecord

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

### NewHealthRecord

`func NewHealthRecord(type_ string, id string, category []string, reactions []AllergyIntoleranceReaction, ) *HealthRecord`

NewHealthRecord instantiates a new HealthRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthRecordWithDefaults

`func NewHealthRecordWithDefaults() *HealthRecord`

NewHealthRecordWithDefaults instantiates a new HealthRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *HealthRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HealthRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HealthRecord) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *HealthRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HealthRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HealthRecord) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *HealthRecord) GetName() CodeableConcept`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HealthRecord) GetNameOk() (*CodeableConcept, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HealthRecord) SetName(v CodeableConcept)`

SetName sets Name field to given value.

### HasName

`func (o *HealthRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategory

`func (o *HealthRecord) GetCategory() []string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *HealthRecord) GetCategoryOk() (*[]string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *HealthRecord) SetCategory(v []string)`

SetCategory sets Category field to given value.


### GetLevelStatus

`func (o *HealthRecord) GetLevelStatus() string`

GetLevelStatus returns the LevelStatus field if non-nil, zero value otherwise.

### GetLevelStatusOk

`func (o *HealthRecord) GetLevelStatusOk() (*string, bool)`

GetLevelStatusOk returns a tuple with the LevelStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelStatus

`func (o *HealthRecord) SetLevelStatus(v string)`

SetLevelStatus sets LevelStatus field to given value.

### HasLevelStatus

`func (o *HealthRecord) HasLevelStatus() bool`

HasLevelStatus returns a boolean if a field has been set.

### GetStatus

`func (o *HealthRecord) GetStatus() CodeableConcept`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthRecord) GetStatusOk() (*CodeableConcept, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthRecord) SetStatus(v CodeableConcept)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HealthRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetClinicalStatus

`func (o *HealthRecord) GetClinicalStatus() CodeableConcept`

GetClinicalStatus returns the ClinicalStatus field if non-nil, zero value otherwise.

### GetClinicalStatusOk

`func (o *HealthRecord) GetClinicalStatusOk() (*CodeableConcept, bool)`

GetClinicalStatusOk returns a tuple with the ClinicalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClinicalStatus

`func (o *HealthRecord) SetClinicalStatus(v CodeableConcept)`

SetClinicalStatus sets ClinicalStatus field to given value.

### HasClinicalStatus

`func (o *HealthRecord) HasClinicalStatus() bool`

HasClinicalStatus returns a boolean if a field has been set.

### GetReactions

`func (o *HealthRecord) GetReactions() []AllergyIntoleranceReaction`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *HealthRecord) GetReactionsOk() (*[]AllergyIntoleranceReaction, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *HealthRecord) SetReactions(v []AllergyIntoleranceReaction)`

SetReactions sets Reactions field to given value.


### GetRecordedAtUtc

`func (o *HealthRecord) GetRecordedAtUtc() time.Time`

GetRecordedAtUtc returns the RecordedAtUtc field if non-nil, zero value otherwise.

### GetRecordedAtUtcOk

`func (o *HealthRecord) GetRecordedAtUtcOk() (*time.Time, bool)`

GetRecordedAtUtcOk returns a tuple with the RecordedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedAtUtc

`func (o *HealthRecord) SetRecordedAtUtc(v time.Time)`

SetRecordedAtUtc sets RecordedAtUtc field to given value.

### HasRecordedAtUtc

`func (o *HealthRecord) HasRecordedAtUtc() bool`

HasRecordedAtUtc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



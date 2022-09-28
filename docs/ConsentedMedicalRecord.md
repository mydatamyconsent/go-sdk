# ConsentedMedicalRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Health id. | [optional] 
**FieldTitle** | **string** | Health field title. | 
**FieldSlug** | **string** | Health field slug. | 
**Category** | **string** | health category type. | 
**ToDate** | Pointer to **time.Time** | To Date | [optional] 
**FromDate** | Pointer to **time.Time** | From Date | [optional] 

## Methods

### NewConsentedMedicalRecord

`func NewConsentedMedicalRecord(fieldTitle string, fieldSlug string, category string, ) *ConsentedMedicalRecord`

NewConsentedMedicalRecord instantiates a new ConsentedMedicalRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsentedMedicalRecordWithDefaults

`func NewConsentedMedicalRecordWithDefaults() *ConsentedMedicalRecord`

NewConsentedMedicalRecordWithDefaults instantiates a new ConsentedMedicalRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConsentedMedicalRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsentedMedicalRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsentedMedicalRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConsentedMedicalRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFieldTitle

`func (o *ConsentedMedicalRecord) GetFieldTitle() string`

GetFieldTitle returns the FieldTitle field if non-nil, zero value otherwise.

### GetFieldTitleOk

`func (o *ConsentedMedicalRecord) GetFieldTitleOk() (*string, bool)`

GetFieldTitleOk returns a tuple with the FieldTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldTitle

`func (o *ConsentedMedicalRecord) SetFieldTitle(v string)`

SetFieldTitle sets FieldTitle field to given value.


### GetFieldSlug

`func (o *ConsentedMedicalRecord) GetFieldSlug() string`

GetFieldSlug returns the FieldSlug field if non-nil, zero value otherwise.

### GetFieldSlugOk

`func (o *ConsentedMedicalRecord) GetFieldSlugOk() (*string, bool)`

GetFieldSlugOk returns a tuple with the FieldSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldSlug

`func (o *ConsentedMedicalRecord) SetFieldSlug(v string)`

SetFieldSlug sets FieldSlug field to given value.


### GetCategory

`func (o *ConsentedMedicalRecord) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ConsentedMedicalRecord) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ConsentedMedicalRecord) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetToDate

`func (o *ConsentedMedicalRecord) GetToDate() time.Time`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *ConsentedMedicalRecord) GetToDateOk() (*time.Time, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *ConsentedMedicalRecord) SetToDate(v time.Time)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *ConsentedMedicalRecord) HasToDate() bool`

HasToDate returns a boolean if a field has been set.

### GetFromDate

`func (o *ConsentedMedicalRecord) GetFromDate() time.Time`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *ConsentedMedicalRecord) GetFromDateOk() (*time.Time, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *ConsentedMedicalRecord) SetFromDate(v time.Time)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *ConsentedMedicalRecord) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



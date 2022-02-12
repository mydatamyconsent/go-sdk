# Document

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldTitle** | **string** |  | 
**FieldSlug** | **string** |  | 
**Drn** | **[]string** |  | 
**Requirement** | [**DocumentsRequired**](DocumentsRequired.md) |  | 

## Methods

### NewDocument

`func NewDocument(fieldTitle string, fieldSlug string, drn []string, requirement DocumentsRequired, ) *Document`

NewDocument instantiates a new Document object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentWithDefaults

`func NewDocumentWithDefaults() *Document`

NewDocumentWithDefaults instantiates a new Document object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldTitle

`func (o *Document) GetFieldTitle() string`

GetFieldTitle returns the FieldTitle field if non-nil, zero value otherwise.

### GetFieldTitleOk

`func (o *Document) GetFieldTitleOk() (*string, bool)`

GetFieldTitleOk returns a tuple with the FieldTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldTitle

`func (o *Document) SetFieldTitle(v string)`

SetFieldTitle sets FieldTitle field to given value.


### GetFieldSlug

`func (o *Document) GetFieldSlug() string`

GetFieldSlug returns the FieldSlug field if non-nil, zero value otherwise.

### GetFieldSlugOk

`func (o *Document) GetFieldSlugOk() (*string, bool)`

GetFieldSlugOk returns a tuple with the FieldSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldSlug

`func (o *Document) SetFieldSlug(v string)`

SetFieldSlug sets FieldSlug field to given value.


### GetDrn

`func (o *Document) GetDrn() []string`

GetDrn returns the Drn field if non-nil, zero value otherwise.

### GetDrnOk

`func (o *Document) GetDrnOk() (*[]string, bool)`

GetDrnOk returns a tuple with the Drn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrn

`func (o *Document) SetDrn(v []string)`

SetDrn sets Drn field to given value.


### GetRequirement

`func (o *Document) GetRequirement() DocumentsRequired`

GetRequirement returns the Requirement field if non-nil, zero value otherwise.

### GetRequirementOk

`func (o *Document) GetRequirementOk() (*DocumentsRequired, bool)`

GetRequirementOk returns a tuple with the Requirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirement

`func (o *Document) SetRequirement(v DocumentsRequired)`

SetRequirement sets Requirement field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



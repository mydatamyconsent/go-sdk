# Document

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentField** | Pointer to **NullableString** |  | [optional] 
**CustomKey** | Pointer to **NullableString** |  | [optional] 
**Drn** | Pointer to **[]string** |  | [optional] 
**Requirement** | Pointer to [**DocumentsRequired**](DocumentsRequired.md) |  | [optional] 

## Methods

### NewDocument

`func NewDocument() *Document`

NewDocument instantiates a new Document object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentWithDefaults

`func NewDocumentWithDefaults() *Document`

NewDocumentWithDefaults instantiates a new Document object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentField

`func (o *Document) GetDocumentField() string`

GetDocumentField returns the DocumentField field if non-nil, zero value otherwise.

### GetDocumentFieldOk

`func (o *Document) GetDocumentFieldOk() (*string, bool)`

GetDocumentFieldOk returns a tuple with the DocumentField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentField

`func (o *Document) SetDocumentField(v string)`

SetDocumentField sets DocumentField field to given value.

### HasDocumentField

`func (o *Document) HasDocumentField() bool`

HasDocumentField returns a boolean if a field has been set.

### SetDocumentFieldNil

`func (o *Document) SetDocumentFieldNil(b bool)`

 SetDocumentFieldNil sets the value for DocumentField to be an explicit nil

### UnsetDocumentField
`func (o *Document) UnsetDocumentField()`

UnsetDocumentField ensures that no value is present for DocumentField, not even an explicit nil
### GetCustomKey

`func (o *Document) GetCustomKey() string`

GetCustomKey returns the CustomKey field if non-nil, zero value otherwise.

### GetCustomKeyOk

`func (o *Document) GetCustomKeyOk() (*string, bool)`

GetCustomKeyOk returns a tuple with the CustomKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomKey

`func (o *Document) SetCustomKey(v string)`

SetCustomKey sets CustomKey field to given value.

### HasCustomKey

`func (o *Document) HasCustomKey() bool`

HasCustomKey returns a boolean if a field has been set.

### SetCustomKeyNil

`func (o *Document) SetCustomKeyNil(b bool)`

 SetCustomKeyNil sets the value for CustomKey to be an explicit nil

### UnsetCustomKey
`func (o *Document) UnsetCustomKey()`

UnsetCustomKey ensures that no value is present for CustomKey, not even an explicit nil
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

### HasDrn

`func (o *Document) HasDrn() bool`

HasDrn returns a boolean if a field has been set.

### SetDrnNil

`func (o *Document) SetDrnNil(b bool)`

 SetDrnNil sets the value for Drn to be an explicit nil

### UnsetDrn
`func (o *Document) UnsetDrn()`

UnsetDrn ensures that no value is present for Drn, not even an explicit nil
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

### HasRequirement

`func (o *Document) HasRequirement() bool`

HasRequirement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



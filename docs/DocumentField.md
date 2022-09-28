# DocumentField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldTitle** | **string** | Field title. | 
**FieldSlug** | **string** | Field slug. | 
**Drns** | **[]string** | Field DRNs. | 

## Methods

### NewDocumentField

`func NewDocumentField(fieldTitle string, fieldSlug string, drns []string, ) *DocumentField`

NewDocumentField instantiates a new DocumentField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentFieldWithDefaults

`func NewDocumentFieldWithDefaults() *DocumentField`

NewDocumentFieldWithDefaults instantiates a new DocumentField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldTitle

`func (o *DocumentField) GetFieldTitle() string`

GetFieldTitle returns the FieldTitle field if non-nil, zero value otherwise.

### GetFieldTitleOk

`func (o *DocumentField) GetFieldTitleOk() (*string, bool)`

GetFieldTitleOk returns a tuple with the FieldTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldTitle

`func (o *DocumentField) SetFieldTitle(v string)`

SetFieldTitle sets FieldTitle field to given value.


### GetFieldSlug

`func (o *DocumentField) GetFieldSlug() string`

GetFieldSlug returns the FieldSlug field if non-nil, zero value otherwise.

### GetFieldSlugOk

`func (o *DocumentField) GetFieldSlugOk() (*string, bool)`

GetFieldSlugOk returns a tuple with the FieldSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldSlug

`func (o *DocumentField) SetFieldSlug(v string)`

SetFieldSlug sets FieldSlug field to given value.


### GetDrns

`func (o *DocumentField) GetDrns() []string`

GetDrns returns the Drns field if non-nil, zero value otherwise.

### GetDrnsOk

`func (o *DocumentField) GetDrnsOk() (*[]string, bool)`

GetDrnsOk returns a tuple with the Drns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrns

`func (o *DocumentField) SetDrns(v []string)`

SetDrns sets Drns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DataConsentDocumentDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldTitle** | **string** |  | 
**FieldSlug** | **string** |  | 
**SupportedDocumentTypeCategoryDetails** | [**[]SupportedDocumentTypeCategoryDetailsDto**](SupportedDocumentTypeCategoryDetailsDto.md) |  | 
**Requirement** | [**DocumentsRequired**](DocumentsRequired.md) |  | 

## Methods

### NewDataConsentDocumentDetailsDto

`func NewDataConsentDocumentDetailsDto(fieldTitle string, fieldSlug string, supportedDocumentTypeCategoryDetails []SupportedDocumentTypeCategoryDetailsDto, requirement DocumentsRequired, ) *DataConsentDocumentDetailsDto`

NewDataConsentDocumentDetailsDto instantiates a new DataConsentDocumentDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataConsentDocumentDetailsDtoWithDefaults

`func NewDataConsentDocumentDetailsDtoWithDefaults() *DataConsentDocumentDetailsDto`

NewDataConsentDocumentDetailsDtoWithDefaults instantiates a new DataConsentDocumentDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldTitle

`func (o *DataConsentDocumentDetailsDto) GetFieldTitle() string`

GetFieldTitle returns the FieldTitle field if non-nil, zero value otherwise.

### GetFieldTitleOk

`func (o *DataConsentDocumentDetailsDto) GetFieldTitleOk() (*string, bool)`

GetFieldTitleOk returns a tuple with the FieldTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldTitle

`func (o *DataConsentDocumentDetailsDto) SetFieldTitle(v string)`

SetFieldTitle sets FieldTitle field to given value.


### GetFieldSlug

`func (o *DataConsentDocumentDetailsDto) GetFieldSlug() string`

GetFieldSlug returns the FieldSlug field if non-nil, zero value otherwise.

### GetFieldSlugOk

`func (o *DataConsentDocumentDetailsDto) GetFieldSlugOk() (*string, bool)`

GetFieldSlugOk returns a tuple with the FieldSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldSlug

`func (o *DataConsentDocumentDetailsDto) SetFieldSlug(v string)`

SetFieldSlug sets FieldSlug field to given value.


### GetSupportedDocumentTypeCategoryDetails

`func (o *DataConsentDocumentDetailsDto) GetSupportedDocumentTypeCategoryDetails() []SupportedDocumentTypeCategoryDetailsDto`

GetSupportedDocumentTypeCategoryDetails returns the SupportedDocumentTypeCategoryDetails field if non-nil, zero value otherwise.

### GetSupportedDocumentTypeCategoryDetailsOk

`func (o *DataConsentDocumentDetailsDto) GetSupportedDocumentTypeCategoryDetailsOk() (*[]SupportedDocumentTypeCategoryDetailsDto, bool)`

GetSupportedDocumentTypeCategoryDetailsOk returns a tuple with the SupportedDocumentTypeCategoryDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDocumentTypeCategoryDetails

`func (o *DataConsentDocumentDetailsDto) SetSupportedDocumentTypeCategoryDetails(v []SupportedDocumentTypeCategoryDetailsDto)`

SetSupportedDocumentTypeCategoryDetails sets SupportedDocumentTypeCategoryDetails field to given value.


### GetRequirement

`func (o *DataConsentDocumentDetailsDto) GetRequirement() DocumentsRequired`

GetRequirement returns the Requirement field if non-nil, zero value otherwise.

### GetRequirementOk

`func (o *DataConsentDocumentDetailsDto) GetRequirementOk() (*DocumentsRequired, bool)`

GetRequirementOk returns a tuple with the Requirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirement

`func (o *DataConsentDocumentDetailsDto) SetRequirement(v DocumentsRequired)`

SetRequirement sets Requirement field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



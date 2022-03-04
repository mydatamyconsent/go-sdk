# SupportedDocumentTypeCategoryDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentTypeCategoryId** | **string** |  | 
**DocumentTypeCategoryName** | **string** |  | 
**SupportedDocuments** | [**[]SupportedDocumentDetailsDto**](SupportedDocumentDetailsDto.md) |  | 
**SupportedDocumentProviderDetails** | Pointer to [**[]SupportedDocumentProviderDetailsDto**](SupportedDocumentProviderDetailsDto.md) |  | [optional] 

## Methods

### NewSupportedDocumentTypeCategoryDetailsDto

`func NewSupportedDocumentTypeCategoryDetailsDto(documentTypeCategoryId string, documentTypeCategoryName string, supportedDocuments []SupportedDocumentDetailsDto, ) *SupportedDocumentTypeCategoryDetailsDto`

NewSupportedDocumentTypeCategoryDetailsDto instantiates a new SupportedDocumentTypeCategoryDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedDocumentTypeCategoryDetailsDtoWithDefaults

`func NewSupportedDocumentTypeCategoryDetailsDtoWithDefaults() *SupportedDocumentTypeCategoryDetailsDto`

NewSupportedDocumentTypeCategoryDetailsDtoWithDefaults instantiates a new SupportedDocumentTypeCategoryDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentTypeCategoryId

`func (o *SupportedDocumentTypeCategoryDetailsDto) GetDocumentTypeCategoryId() string`

GetDocumentTypeCategoryId returns the DocumentTypeCategoryId field if non-nil, zero value otherwise.

### GetDocumentTypeCategoryIdOk

`func (o *SupportedDocumentTypeCategoryDetailsDto) GetDocumentTypeCategoryIdOk() (*string, bool)`

GetDocumentTypeCategoryIdOk returns a tuple with the DocumentTypeCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTypeCategoryId

`func (o *SupportedDocumentTypeCategoryDetailsDto) SetDocumentTypeCategoryId(v string)`

SetDocumentTypeCategoryId sets DocumentTypeCategoryId field to given value.


### GetDocumentTypeCategoryName

`func (o *SupportedDocumentTypeCategoryDetailsDto) GetDocumentTypeCategoryName() string`

GetDocumentTypeCategoryName returns the DocumentTypeCategoryName field if non-nil, zero value otherwise.

### GetDocumentTypeCategoryNameOk

`func (o *SupportedDocumentTypeCategoryDetailsDto) GetDocumentTypeCategoryNameOk() (*string, bool)`

GetDocumentTypeCategoryNameOk returns a tuple with the DocumentTypeCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTypeCategoryName

`func (o *SupportedDocumentTypeCategoryDetailsDto) SetDocumentTypeCategoryName(v string)`

SetDocumentTypeCategoryName sets DocumentTypeCategoryName field to given value.


### GetSupportedDocuments

`func (o *SupportedDocumentTypeCategoryDetailsDto) GetSupportedDocuments() []SupportedDocumentDetailsDto`

GetSupportedDocuments returns the SupportedDocuments field if non-nil, zero value otherwise.

### GetSupportedDocumentsOk

`func (o *SupportedDocumentTypeCategoryDetailsDto) GetSupportedDocumentsOk() (*[]SupportedDocumentDetailsDto, bool)`

GetSupportedDocumentsOk returns a tuple with the SupportedDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDocuments

`func (o *SupportedDocumentTypeCategoryDetailsDto) SetSupportedDocuments(v []SupportedDocumentDetailsDto)`

SetSupportedDocuments sets SupportedDocuments field to given value.


### GetSupportedDocumentProviderDetails

`func (o *SupportedDocumentTypeCategoryDetailsDto) GetSupportedDocumentProviderDetails() []SupportedDocumentProviderDetailsDto`

GetSupportedDocumentProviderDetails returns the SupportedDocumentProviderDetails field if non-nil, zero value otherwise.

### GetSupportedDocumentProviderDetailsOk

`func (o *SupportedDocumentTypeCategoryDetailsDto) GetSupportedDocumentProviderDetailsOk() (*[]SupportedDocumentProviderDetailsDto, bool)`

GetSupportedDocumentProviderDetailsOk returns a tuple with the SupportedDocumentProviderDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDocumentProviderDetails

`func (o *SupportedDocumentTypeCategoryDetailsDto) SetSupportedDocumentProviderDetails(v []SupportedDocumentProviderDetailsDto)`

SetSupportedDocumentProviderDetails sets SupportedDocumentProviderDetails field to given value.

### HasSupportedDocumentProviderDetails

`func (o *SupportedDocumentTypeCategoryDetailsDto) HasSupportedDocumentProviderDetails() bool`

HasSupportedDocumentProviderDetails returns a boolean if a field has been set.

### SetSupportedDocumentProviderDetailsNil

`func (o *SupportedDocumentTypeCategoryDetailsDto) SetSupportedDocumentProviderDetailsNil(b bool)`

 SetSupportedDocumentProviderDetailsNil sets the value for SupportedDocumentProviderDetails to be an explicit nil

### UnsetSupportedDocumentProviderDetails
`func (o *SupportedDocumentTypeCategoryDetailsDto) UnsetSupportedDocumentProviderDetails()`

UnsetSupportedDocumentProviderDetails ensures that no value is present for SupportedDocumentProviderDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DocumentTypeDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**OrganizationId** | **string** |  | 
**CategoryType** | [**DocumentCategoryType**](DocumentCategoryType.md) |  | 
**SubCategoryType** | [**DocumentSubCategoryType**](DocumentSubCategoryType.md) |  | 
**DocumentTypeCategoryId** | **string** |  | 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**LogoUrl** | **string** |  | 
**CountryIso2** | **string** |  | 
**CountryId** | **string** |  | 
**SearchServiceId** | Pointer to **NullableString** |  | [optional] 
**RepositoryServiceId** | Pointer to **NullableString** |  | [optional] 
**SupportedEntityType** | [**SupportedEntityType**](SupportedEntityType.md) |  | 
**AddedBy** | **string** |  | 
**PayableAmount** | Pointer to **NullableFloat64** |  | [optional] 
**ApprovedAtUtc** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewDocumentTypeDetailsDto

`func NewDocumentTypeDetailsDto(id string, organizationId string, categoryType DocumentCategoryType, subCategoryType DocumentSubCategoryType, documentTypeCategoryId string, name string, slug string, logoUrl string, countryIso2 string, countryId string, supportedEntityType SupportedEntityType, addedBy string, ) *DocumentTypeDetailsDto`

NewDocumentTypeDetailsDto instantiates a new DocumentTypeDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentTypeDetailsDtoWithDefaults

`func NewDocumentTypeDetailsDtoWithDefaults() *DocumentTypeDetailsDto`

NewDocumentTypeDetailsDtoWithDefaults instantiates a new DocumentTypeDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DocumentTypeDetailsDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentTypeDetailsDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentTypeDetailsDto) SetId(v string)`

SetId sets Id field to given value.


### GetOrganizationId

`func (o *DocumentTypeDetailsDto) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *DocumentTypeDetailsDto) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *DocumentTypeDetailsDto) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetCategoryType

`func (o *DocumentTypeDetailsDto) GetCategoryType() DocumentCategoryType`

GetCategoryType returns the CategoryType field if non-nil, zero value otherwise.

### GetCategoryTypeOk

`func (o *DocumentTypeDetailsDto) GetCategoryTypeOk() (*DocumentCategoryType, bool)`

GetCategoryTypeOk returns a tuple with the CategoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryType

`func (o *DocumentTypeDetailsDto) SetCategoryType(v DocumentCategoryType)`

SetCategoryType sets CategoryType field to given value.


### GetSubCategoryType

`func (o *DocumentTypeDetailsDto) GetSubCategoryType() DocumentSubCategoryType`

GetSubCategoryType returns the SubCategoryType field if non-nil, zero value otherwise.

### GetSubCategoryTypeOk

`func (o *DocumentTypeDetailsDto) GetSubCategoryTypeOk() (*DocumentSubCategoryType, bool)`

GetSubCategoryTypeOk returns a tuple with the SubCategoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCategoryType

`func (o *DocumentTypeDetailsDto) SetSubCategoryType(v DocumentSubCategoryType)`

SetSubCategoryType sets SubCategoryType field to given value.


### GetDocumentTypeCategoryId

`func (o *DocumentTypeDetailsDto) GetDocumentTypeCategoryId() string`

GetDocumentTypeCategoryId returns the DocumentTypeCategoryId field if non-nil, zero value otherwise.

### GetDocumentTypeCategoryIdOk

`func (o *DocumentTypeDetailsDto) GetDocumentTypeCategoryIdOk() (*string, bool)`

GetDocumentTypeCategoryIdOk returns a tuple with the DocumentTypeCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTypeCategoryId

`func (o *DocumentTypeDetailsDto) SetDocumentTypeCategoryId(v string)`

SetDocumentTypeCategoryId sets DocumentTypeCategoryId field to given value.


### GetName

`func (o *DocumentTypeDetailsDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DocumentTypeDetailsDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DocumentTypeDetailsDto) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *DocumentTypeDetailsDto) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *DocumentTypeDetailsDto) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *DocumentTypeDetailsDto) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *DocumentTypeDetailsDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DocumentTypeDetailsDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DocumentTypeDetailsDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DocumentTypeDetailsDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DocumentTypeDetailsDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DocumentTypeDetailsDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLogoUrl

`func (o *DocumentTypeDetailsDto) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *DocumentTypeDetailsDto) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *DocumentTypeDetailsDto) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.


### GetCountryIso2

`func (o *DocumentTypeDetailsDto) GetCountryIso2() string`

GetCountryIso2 returns the CountryIso2 field if non-nil, zero value otherwise.

### GetCountryIso2Ok

`func (o *DocumentTypeDetailsDto) GetCountryIso2Ok() (*string, bool)`

GetCountryIso2Ok returns a tuple with the CountryIso2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIso2

`func (o *DocumentTypeDetailsDto) SetCountryIso2(v string)`

SetCountryIso2 sets CountryIso2 field to given value.


### GetCountryId

`func (o *DocumentTypeDetailsDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *DocumentTypeDetailsDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *DocumentTypeDetailsDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.


### GetSearchServiceId

`func (o *DocumentTypeDetailsDto) GetSearchServiceId() string`

GetSearchServiceId returns the SearchServiceId field if non-nil, zero value otherwise.

### GetSearchServiceIdOk

`func (o *DocumentTypeDetailsDto) GetSearchServiceIdOk() (*string, bool)`

GetSearchServiceIdOk returns a tuple with the SearchServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchServiceId

`func (o *DocumentTypeDetailsDto) SetSearchServiceId(v string)`

SetSearchServiceId sets SearchServiceId field to given value.

### HasSearchServiceId

`func (o *DocumentTypeDetailsDto) HasSearchServiceId() bool`

HasSearchServiceId returns a boolean if a field has been set.

### SetSearchServiceIdNil

`func (o *DocumentTypeDetailsDto) SetSearchServiceIdNil(b bool)`

 SetSearchServiceIdNil sets the value for SearchServiceId to be an explicit nil

### UnsetSearchServiceId
`func (o *DocumentTypeDetailsDto) UnsetSearchServiceId()`

UnsetSearchServiceId ensures that no value is present for SearchServiceId, not even an explicit nil
### GetRepositoryServiceId

`func (o *DocumentTypeDetailsDto) GetRepositoryServiceId() string`

GetRepositoryServiceId returns the RepositoryServiceId field if non-nil, zero value otherwise.

### GetRepositoryServiceIdOk

`func (o *DocumentTypeDetailsDto) GetRepositoryServiceIdOk() (*string, bool)`

GetRepositoryServiceIdOk returns a tuple with the RepositoryServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryServiceId

`func (o *DocumentTypeDetailsDto) SetRepositoryServiceId(v string)`

SetRepositoryServiceId sets RepositoryServiceId field to given value.

### HasRepositoryServiceId

`func (o *DocumentTypeDetailsDto) HasRepositoryServiceId() bool`

HasRepositoryServiceId returns a boolean if a field has been set.

### SetRepositoryServiceIdNil

`func (o *DocumentTypeDetailsDto) SetRepositoryServiceIdNil(b bool)`

 SetRepositoryServiceIdNil sets the value for RepositoryServiceId to be an explicit nil

### UnsetRepositoryServiceId
`func (o *DocumentTypeDetailsDto) UnsetRepositoryServiceId()`

UnsetRepositoryServiceId ensures that no value is present for RepositoryServiceId, not even an explicit nil
### GetSupportedEntityType

`func (o *DocumentTypeDetailsDto) GetSupportedEntityType() SupportedEntityType`

GetSupportedEntityType returns the SupportedEntityType field if non-nil, zero value otherwise.

### GetSupportedEntityTypeOk

`func (o *DocumentTypeDetailsDto) GetSupportedEntityTypeOk() (*SupportedEntityType, bool)`

GetSupportedEntityTypeOk returns a tuple with the SupportedEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedEntityType

`func (o *DocumentTypeDetailsDto) SetSupportedEntityType(v SupportedEntityType)`

SetSupportedEntityType sets SupportedEntityType field to given value.


### GetAddedBy

`func (o *DocumentTypeDetailsDto) GetAddedBy() string`

GetAddedBy returns the AddedBy field if non-nil, zero value otherwise.

### GetAddedByOk

`func (o *DocumentTypeDetailsDto) GetAddedByOk() (*string, bool)`

GetAddedByOk returns a tuple with the AddedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedBy

`func (o *DocumentTypeDetailsDto) SetAddedBy(v string)`

SetAddedBy sets AddedBy field to given value.


### GetPayableAmount

`func (o *DocumentTypeDetailsDto) GetPayableAmount() float64`

GetPayableAmount returns the PayableAmount field if non-nil, zero value otherwise.

### GetPayableAmountOk

`func (o *DocumentTypeDetailsDto) GetPayableAmountOk() (*float64, bool)`

GetPayableAmountOk returns a tuple with the PayableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayableAmount

`func (o *DocumentTypeDetailsDto) SetPayableAmount(v float64)`

SetPayableAmount sets PayableAmount field to given value.

### HasPayableAmount

`func (o *DocumentTypeDetailsDto) HasPayableAmount() bool`

HasPayableAmount returns a boolean if a field has been set.

### SetPayableAmountNil

`func (o *DocumentTypeDetailsDto) SetPayableAmountNil(b bool)`

 SetPayableAmountNil sets the value for PayableAmount to be an explicit nil

### UnsetPayableAmount
`func (o *DocumentTypeDetailsDto) UnsetPayableAmount()`

UnsetPayableAmount ensures that no value is present for PayableAmount, not even an explicit nil
### GetApprovedAtUtc

`func (o *DocumentTypeDetailsDto) GetApprovedAtUtc() time.Time`

GetApprovedAtUtc returns the ApprovedAtUtc field if non-nil, zero value otherwise.

### GetApprovedAtUtcOk

`func (o *DocumentTypeDetailsDto) GetApprovedAtUtcOk() (*time.Time, bool)`

GetApprovedAtUtcOk returns a tuple with the ApprovedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAtUtc

`func (o *DocumentTypeDetailsDto) SetApprovedAtUtc(v time.Time)`

SetApprovedAtUtc sets ApprovedAtUtc field to given value.

### HasApprovedAtUtc

`func (o *DocumentTypeDetailsDto) HasApprovedAtUtc() bool`

HasApprovedAtUtc returns a boolean if a field has been set.

### SetApprovedAtUtcNil

`func (o *DocumentTypeDetailsDto) SetApprovedAtUtcNil(b bool)`

 SetApprovedAtUtcNil sets the value for ApprovedAtUtc to be an explicit nil

### UnsetApprovedAtUtc
`func (o *DocumentTypeDetailsDto) UnsetApprovedAtUtc()`

UnsetApprovedAtUtc ensures that no value is present for ApprovedAtUtc, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



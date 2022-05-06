# DocumentType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Document Type Identifier. | 
**CategoryType** | [**DocumentCategoryType**](DocumentCategoryType.md) |  | 
**SubCategoryType** | [**DocumentSubCategoryType**](DocumentSubCategoryType.md) |  | 
**Name** | **string** | Document Type Name. eg: Driving License. | 
**Slug** | **string** | Document Type Unique Slug. eg: \&quot;in.gov.gj.transport.dl\&quot;. | 
**Description** | Pointer to **NullableString** | Document Type description. eg: Gujarat State Driving License. | [optional] 
**LogoUrl** | **string** | Logo URL of document type. | 
**SearchServiceName** | Pointer to **NullableString** | Document search repository service name. | [optional] 
**RepositoryServiceName** | Pointer to **NullableString** | Document repository service name. | [optional] 
**SupportedEntityTypes** | [**[]SupportedEntityType**](SupportedEntityType.md) | Supported entity types. eg: Individual, Organization. | 
**AddedBy** | **string** | Name of the document type creator. | 
**PayableAmount** | **float64** | Payable amount if document is chargeable. eg: 10.25. | 
**PayableAmountCurrency** | Pointer to **NullableString** | Payable amount currency. eg: INR, USD etc.,. | [optional] 
**ApprovedAtUtc** | Pointer to **NullableTime** | DateTime of approval in UTC timezone. | [optional] 

## Methods

### NewDocumentType

`func NewDocumentType(id string, categoryType DocumentCategoryType, subCategoryType DocumentSubCategoryType, name string, slug string, logoUrl string, supportedEntityTypes []SupportedEntityType, addedBy string, payableAmount float64, ) *DocumentType`

NewDocumentType instantiates a new DocumentType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentTypeWithDefaults

`func NewDocumentTypeWithDefaults() *DocumentType`

NewDocumentTypeWithDefaults instantiates a new DocumentType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DocumentType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentType) SetId(v string)`

SetId sets Id field to given value.


### GetCategoryType

`func (o *DocumentType) GetCategoryType() DocumentCategoryType`

GetCategoryType returns the CategoryType field if non-nil, zero value otherwise.

### GetCategoryTypeOk

`func (o *DocumentType) GetCategoryTypeOk() (*DocumentCategoryType, bool)`

GetCategoryTypeOk returns a tuple with the CategoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryType

`func (o *DocumentType) SetCategoryType(v DocumentCategoryType)`

SetCategoryType sets CategoryType field to given value.


### GetSubCategoryType

`func (o *DocumentType) GetSubCategoryType() DocumentSubCategoryType`

GetSubCategoryType returns the SubCategoryType field if non-nil, zero value otherwise.

### GetSubCategoryTypeOk

`func (o *DocumentType) GetSubCategoryTypeOk() (*DocumentSubCategoryType, bool)`

GetSubCategoryTypeOk returns a tuple with the SubCategoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCategoryType

`func (o *DocumentType) SetSubCategoryType(v DocumentSubCategoryType)`

SetSubCategoryType sets SubCategoryType field to given value.


### GetName

`func (o *DocumentType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DocumentType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DocumentType) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *DocumentType) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *DocumentType) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *DocumentType) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *DocumentType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DocumentType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DocumentType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DocumentType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DocumentType) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DocumentType) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLogoUrl

`func (o *DocumentType) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *DocumentType) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *DocumentType) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.


### GetSearchServiceName

`func (o *DocumentType) GetSearchServiceName() string`

GetSearchServiceName returns the SearchServiceName field if non-nil, zero value otherwise.

### GetSearchServiceNameOk

`func (o *DocumentType) GetSearchServiceNameOk() (*string, bool)`

GetSearchServiceNameOk returns a tuple with the SearchServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchServiceName

`func (o *DocumentType) SetSearchServiceName(v string)`

SetSearchServiceName sets SearchServiceName field to given value.

### HasSearchServiceName

`func (o *DocumentType) HasSearchServiceName() bool`

HasSearchServiceName returns a boolean if a field has been set.

### SetSearchServiceNameNil

`func (o *DocumentType) SetSearchServiceNameNil(b bool)`

 SetSearchServiceNameNil sets the value for SearchServiceName to be an explicit nil

### UnsetSearchServiceName
`func (o *DocumentType) UnsetSearchServiceName()`

UnsetSearchServiceName ensures that no value is present for SearchServiceName, not even an explicit nil
### GetRepositoryServiceName

`func (o *DocumentType) GetRepositoryServiceName() string`

GetRepositoryServiceName returns the RepositoryServiceName field if non-nil, zero value otherwise.

### GetRepositoryServiceNameOk

`func (o *DocumentType) GetRepositoryServiceNameOk() (*string, bool)`

GetRepositoryServiceNameOk returns a tuple with the RepositoryServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryServiceName

`func (o *DocumentType) SetRepositoryServiceName(v string)`

SetRepositoryServiceName sets RepositoryServiceName field to given value.

### HasRepositoryServiceName

`func (o *DocumentType) HasRepositoryServiceName() bool`

HasRepositoryServiceName returns a boolean if a field has been set.

### SetRepositoryServiceNameNil

`func (o *DocumentType) SetRepositoryServiceNameNil(b bool)`

 SetRepositoryServiceNameNil sets the value for RepositoryServiceName to be an explicit nil

### UnsetRepositoryServiceName
`func (o *DocumentType) UnsetRepositoryServiceName()`

UnsetRepositoryServiceName ensures that no value is present for RepositoryServiceName, not even an explicit nil
### GetSupportedEntityTypes

`func (o *DocumentType) GetSupportedEntityTypes() []SupportedEntityType`

GetSupportedEntityTypes returns the SupportedEntityTypes field if non-nil, zero value otherwise.

### GetSupportedEntityTypesOk

`func (o *DocumentType) GetSupportedEntityTypesOk() (*[]SupportedEntityType, bool)`

GetSupportedEntityTypesOk returns a tuple with the SupportedEntityTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedEntityTypes

`func (o *DocumentType) SetSupportedEntityTypes(v []SupportedEntityType)`

SetSupportedEntityTypes sets SupportedEntityTypes field to given value.


### GetAddedBy

`func (o *DocumentType) GetAddedBy() string`

GetAddedBy returns the AddedBy field if non-nil, zero value otherwise.

### GetAddedByOk

`func (o *DocumentType) GetAddedByOk() (*string, bool)`

GetAddedByOk returns a tuple with the AddedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedBy

`func (o *DocumentType) SetAddedBy(v string)`

SetAddedBy sets AddedBy field to given value.


### GetPayableAmount

`func (o *DocumentType) GetPayableAmount() float64`

GetPayableAmount returns the PayableAmount field if non-nil, zero value otherwise.

### GetPayableAmountOk

`func (o *DocumentType) GetPayableAmountOk() (*float64, bool)`

GetPayableAmountOk returns a tuple with the PayableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayableAmount

`func (o *DocumentType) SetPayableAmount(v float64)`

SetPayableAmount sets PayableAmount field to given value.


### GetPayableAmountCurrency

`func (o *DocumentType) GetPayableAmountCurrency() string`

GetPayableAmountCurrency returns the PayableAmountCurrency field if non-nil, zero value otherwise.

### GetPayableAmountCurrencyOk

`func (o *DocumentType) GetPayableAmountCurrencyOk() (*string, bool)`

GetPayableAmountCurrencyOk returns a tuple with the PayableAmountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayableAmountCurrency

`func (o *DocumentType) SetPayableAmountCurrency(v string)`

SetPayableAmountCurrency sets PayableAmountCurrency field to given value.

### HasPayableAmountCurrency

`func (o *DocumentType) HasPayableAmountCurrency() bool`

HasPayableAmountCurrency returns a boolean if a field has been set.

### SetPayableAmountCurrencyNil

`func (o *DocumentType) SetPayableAmountCurrencyNil(b bool)`

 SetPayableAmountCurrencyNil sets the value for PayableAmountCurrency to be an explicit nil

### UnsetPayableAmountCurrency
`func (o *DocumentType) UnsetPayableAmountCurrency()`

UnsetPayableAmountCurrency ensures that no value is present for PayableAmountCurrency, not even an explicit nil
### GetApprovedAtUtc

`func (o *DocumentType) GetApprovedAtUtc() time.Time`

GetApprovedAtUtc returns the ApprovedAtUtc field if non-nil, zero value otherwise.

### GetApprovedAtUtcOk

`func (o *DocumentType) GetApprovedAtUtcOk() (*time.Time, bool)`

GetApprovedAtUtcOk returns a tuple with the ApprovedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAtUtc

`func (o *DocumentType) SetApprovedAtUtc(v time.Time)`

SetApprovedAtUtc sets ApprovedAtUtc field to given value.

### HasApprovedAtUtc

`func (o *DocumentType) HasApprovedAtUtc() bool`

HasApprovedAtUtc returns a boolean if a field has been set.

### SetApprovedAtUtcNil

`func (o *DocumentType) SetApprovedAtUtcNil(b bool)`

 SetApprovedAtUtcNil sets the value for ApprovedAtUtc to be an explicit nil

### UnsetApprovedAtUtc
`func (o *DocumentType) UnsetApprovedAtUtc()`

UnsetApprovedAtUtc ensures that no value is present for ApprovedAtUtc, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



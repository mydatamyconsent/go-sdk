# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedAtUtc** | Pointer to **time.Time** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**CreatedByUser** | Pointer to [**ApplicationUser**](ApplicationUser.md) |  | [optional] 
**UpdatedByUser** | Pointer to [**ApplicationUser**](ApplicationUser.md) |  | [optional] 
**DeletedBy** | Pointer to **NullableString** |  | [optional] 
**DeletedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**DeletedByUser** | Pointer to [**ApplicationUser**](ApplicationUser.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**DataPartnerId** | Pointer to **NullableString** |  | [optional] 
**TypeId** | Pointer to **string** |  | [optional] 
**CategoryId** | Pointer to **string** |  | [optional] 
**DocumentProviderCategoryId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**RegulatorName** | Pointer to **NullableString** |  | [optional] 
**BrandName** | Pointer to **NullableString** |  | [optional] 
**AuthorizedPersonnelTaxId** | Pointer to **NullableString** |  | [optional] 
**AuthorizedPersonnelName** | Pointer to **NullableString** |  | [optional] 
**RegistrationId** | Pointer to **NullableString** |  | [optional] 
**VatId** | Pointer to **NullableString** |  | [optional] 
**TaxId** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**LogoUrl** | Pointer to **NullableString** |  | [optional] 
**ContactEmail** | Pointer to **NullableString** |  | [optional] 
**PhoneNumber** | Pointer to **NullableString** |  | [optional] 
**CountryId** | Pointer to **string** |  | [optional] 
**StateId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**OrganizationStatus**](OrganizationStatus.md) |  | [optional] 
**CompanyCode** | Pointer to **NullableString** |  | [optional] 
**PrivacyPolicyUrl** | Pointer to **NullableString** |  | [optional] 
**TermsOfServiceUrl** | Pointer to **NullableString** |  | [optional] 
**WebsiteUrl** | Pointer to **NullableString** |  | [optional] 
**IsGovernmentOrganization** | Pointer to **bool** |  | [optional] 
**DlApiKey** | Pointer to **NullableString** |  | [optional] 
**IsKyoCompleted** | Pointer to **bool** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**IsDataProvider** | Pointer to **bool** |  | [optional] 
**IsDataConsumer** | Pointer to **bool** |  | [optional] 
**SubmittedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**ApprovedBy** | Pointer to **NullableString** |  | [optional] 
**ApprovedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**IsDigiLockerOrganization** | Pointer to **bool** |  | [optional] 
**IsMdmcMaintained** | Pointer to **bool** |  | [optional] 
**IsBbps** | Pointer to **bool** |  | [optional] 
**MetaData** | Pointer to [**OrganizationMetaData**](OrganizationMetaData.md) |  | [optional] 
**CountryCode** | Pointer to **NullableString** |  | [optional] 
**OrganizationType** | Pointer to [**OrganizationType**](OrganizationType.md) |  | [optional] 
**OrganizationCategory** | Pointer to [**OrganizationCategory**](OrganizationCategory.md) |  | [optional] 
**DocumentProviderCategory** | Pointer to [**DocumentProviderCategory**](DocumentProviderCategory.md) |  | [optional] 
**Addresses** | Pointer to [**[]OrganizationAddress**](OrganizationAddress.md) |  | [optional] 
**FinancialAccounts** | Pointer to [**[]OrganizationFinancialAccount**](OrganizationFinancialAccount.md) |  | [optional] 
**CountryState** | Pointer to [**Country**](Country.md) |  | [optional] 
**ApprovedByUser** | Pointer to [**ApplicationUser**](ApplicationUser.md) |  | [optional] 
**KyoDocuments** | Pointer to [**[]OrganizationKyoDocument**](OrganizationKyoDocument.md) |  | [optional] 
**IsDelete** | Pointer to **bool** |  | [optional] 
**RecoveryToken** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewOrganization

`func NewOrganization() *Organization`

NewOrganization instantiates a new Organization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *Organization) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Organization) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Organization) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Organization) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedAtUtc

`func (o *Organization) GetCreatedAtUtc() time.Time`

GetCreatedAtUtc returns the CreatedAtUtc field if non-nil, zero value otherwise.

### GetCreatedAtUtcOk

`func (o *Organization) GetCreatedAtUtcOk() (*time.Time, bool)`

GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtUtc

`func (o *Organization) SetCreatedAtUtc(v time.Time)`

SetCreatedAtUtc sets CreatedAtUtc field to given value.

### HasCreatedAtUtc

`func (o *Organization) HasCreatedAtUtc() bool`

HasCreatedAtUtc returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Organization) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Organization) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Organization) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Organization) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *Organization) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *Organization) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetUpdatedAtUtc

`func (o *Organization) GetUpdatedAtUtc() time.Time`

GetUpdatedAtUtc returns the UpdatedAtUtc field if non-nil, zero value otherwise.

### GetUpdatedAtUtcOk

`func (o *Organization) GetUpdatedAtUtcOk() (*time.Time, bool)`

GetUpdatedAtUtcOk returns a tuple with the UpdatedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtUtc

`func (o *Organization) SetUpdatedAtUtc(v time.Time)`

SetUpdatedAtUtc sets UpdatedAtUtc field to given value.

### HasUpdatedAtUtc

`func (o *Organization) HasUpdatedAtUtc() bool`

HasUpdatedAtUtc returns a boolean if a field has been set.

### SetUpdatedAtUtcNil

`func (o *Organization) SetUpdatedAtUtcNil(b bool)`

 SetUpdatedAtUtcNil sets the value for UpdatedAtUtc to be an explicit nil

### UnsetUpdatedAtUtc
`func (o *Organization) UnsetUpdatedAtUtc()`

UnsetUpdatedAtUtc ensures that no value is present for UpdatedAtUtc, not even an explicit nil
### GetCreatedByUser

`func (o *Organization) GetCreatedByUser() ApplicationUser`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *Organization) GetCreatedByUserOk() (*ApplicationUser, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *Organization) SetCreatedByUser(v ApplicationUser)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *Organization) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### GetUpdatedByUser

`func (o *Organization) GetUpdatedByUser() ApplicationUser`

GetUpdatedByUser returns the UpdatedByUser field if non-nil, zero value otherwise.

### GetUpdatedByUserOk

`func (o *Organization) GetUpdatedByUserOk() (*ApplicationUser, bool)`

GetUpdatedByUserOk returns a tuple with the UpdatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUser

`func (o *Organization) SetUpdatedByUser(v ApplicationUser)`

SetUpdatedByUser sets UpdatedByUser field to given value.

### HasUpdatedByUser

`func (o *Organization) HasUpdatedByUser() bool`

HasUpdatedByUser returns a boolean if a field has been set.

### GetDeletedBy

`func (o *Organization) GetDeletedBy() string`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *Organization) GetDeletedByOk() (*string, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *Organization) SetDeletedBy(v string)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *Organization) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### SetDeletedByNil

`func (o *Organization) SetDeletedByNil(b bool)`

 SetDeletedByNil sets the value for DeletedBy to be an explicit nil

### UnsetDeletedBy
`func (o *Organization) UnsetDeletedBy()`

UnsetDeletedBy ensures that no value is present for DeletedBy, not even an explicit nil
### GetDeletedAtUtc

`func (o *Organization) GetDeletedAtUtc() time.Time`

GetDeletedAtUtc returns the DeletedAtUtc field if non-nil, zero value otherwise.

### GetDeletedAtUtcOk

`func (o *Organization) GetDeletedAtUtcOk() (*time.Time, bool)`

GetDeletedAtUtcOk returns a tuple with the DeletedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAtUtc

`func (o *Organization) SetDeletedAtUtc(v time.Time)`

SetDeletedAtUtc sets DeletedAtUtc field to given value.

### HasDeletedAtUtc

`func (o *Organization) HasDeletedAtUtc() bool`

HasDeletedAtUtc returns a boolean if a field has been set.

### SetDeletedAtUtcNil

`func (o *Organization) SetDeletedAtUtcNil(b bool)`

 SetDeletedAtUtcNil sets the value for DeletedAtUtc to be an explicit nil

### UnsetDeletedAtUtc
`func (o *Organization) UnsetDeletedAtUtc()`

UnsetDeletedAtUtc ensures that no value is present for DeletedAtUtc, not even an explicit nil
### GetDeletedByUser

`func (o *Organization) GetDeletedByUser() ApplicationUser`

GetDeletedByUser returns the DeletedByUser field if non-nil, zero value otherwise.

### GetDeletedByUserOk

`func (o *Organization) GetDeletedByUserOk() (*ApplicationUser, bool)`

GetDeletedByUserOk returns a tuple with the DeletedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedByUser

`func (o *Organization) SetDeletedByUser(v ApplicationUser)`

SetDeletedByUser sets DeletedByUser field to given value.

### HasDeletedByUser

`func (o *Organization) HasDeletedByUser() bool`

HasDeletedByUser returns a boolean if a field has been set.

### GetId

`func (o *Organization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Organization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Organization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Organization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDataPartnerId

`func (o *Organization) GetDataPartnerId() string`

GetDataPartnerId returns the DataPartnerId field if non-nil, zero value otherwise.

### GetDataPartnerIdOk

`func (o *Organization) GetDataPartnerIdOk() (*string, bool)`

GetDataPartnerIdOk returns a tuple with the DataPartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPartnerId

`func (o *Organization) SetDataPartnerId(v string)`

SetDataPartnerId sets DataPartnerId field to given value.

### HasDataPartnerId

`func (o *Organization) HasDataPartnerId() bool`

HasDataPartnerId returns a boolean if a field has been set.

### SetDataPartnerIdNil

`func (o *Organization) SetDataPartnerIdNil(b bool)`

 SetDataPartnerIdNil sets the value for DataPartnerId to be an explicit nil

### UnsetDataPartnerId
`func (o *Organization) UnsetDataPartnerId()`

UnsetDataPartnerId ensures that no value is present for DataPartnerId, not even an explicit nil
### GetTypeId

`func (o *Organization) GetTypeId() string`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *Organization) GetTypeIdOk() (*string, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *Organization) SetTypeId(v string)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *Organization) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetCategoryId

`func (o *Organization) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *Organization) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *Organization) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *Organization) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetDocumentProviderCategoryId

`func (o *Organization) GetDocumentProviderCategoryId() string`

GetDocumentProviderCategoryId returns the DocumentProviderCategoryId field if non-nil, zero value otherwise.

### GetDocumentProviderCategoryIdOk

`func (o *Organization) GetDocumentProviderCategoryIdOk() (*string, bool)`

GetDocumentProviderCategoryIdOk returns a tuple with the DocumentProviderCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentProviderCategoryId

`func (o *Organization) SetDocumentProviderCategoryId(v string)`

SetDocumentProviderCategoryId sets DocumentProviderCategoryId field to given value.

### HasDocumentProviderCategoryId

`func (o *Organization) HasDocumentProviderCategoryId() bool`

HasDocumentProviderCategoryId returns a boolean if a field has been set.

### SetDocumentProviderCategoryIdNil

`func (o *Organization) SetDocumentProviderCategoryIdNil(b bool)`

 SetDocumentProviderCategoryIdNil sets the value for DocumentProviderCategoryId to be an explicit nil

### UnsetDocumentProviderCategoryId
`func (o *Organization) UnsetDocumentProviderCategoryId()`

UnsetDocumentProviderCategoryId ensures that no value is present for DocumentProviderCategoryId, not even an explicit nil
### GetName

`func (o *Organization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Organization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Organization) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Organization) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Organization) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Organization) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRegulatorName

`func (o *Organization) GetRegulatorName() string`

GetRegulatorName returns the RegulatorName field if non-nil, zero value otherwise.

### GetRegulatorNameOk

`func (o *Organization) GetRegulatorNameOk() (*string, bool)`

GetRegulatorNameOk returns a tuple with the RegulatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatorName

`func (o *Organization) SetRegulatorName(v string)`

SetRegulatorName sets RegulatorName field to given value.

### HasRegulatorName

`func (o *Organization) HasRegulatorName() bool`

HasRegulatorName returns a boolean if a field has been set.

### SetRegulatorNameNil

`func (o *Organization) SetRegulatorNameNil(b bool)`

 SetRegulatorNameNil sets the value for RegulatorName to be an explicit nil

### UnsetRegulatorName
`func (o *Organization) UnsetRegulatorName()`

UnsetRegulatorName ensures that no value is present for RegulatorName, not even an explicit nil
### GetBrandName

`func (o *Organization) GetBrandName() string`

GetBrandName returns the BrandName field if non-nil, zero value otherwise.

### GetBrandNameOk

`func (o *Organization) GetBrandNameOk() (*string, bool)`

GetBrandNameOk returns a tuple with the BrandName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandName

`func (o *Organization) SetBrandName(v string)`

SetBrandName sets BrandName field to given value.

### HasBrandName

`func (o *Organization) HasBrandName() bool`

HasBrandName returns a boolean if a field has been set.

### SetBrandNameNil

`func (o *Organization) SetBrandNameNil(b bool)`

 SetBrandNameNil sets the value for BrandName to be an explicit nil

### UnsetBrandName
`func (o *Organization) UnsetBrandName()`

UnsetBrandName ensures that no value is present for BrandName, not even an explicit nil
### GetAuthorizedPersonnelTaxId

`func (o *Organization) GetAuthorizedPersonnelTaxId() string`

GetAuthorizedPersonnelTaxId returns the AuthorizedPersonnelTaxId field if non-nil, zero value otherwise.

### GetAuthorizedPersonnelTaxIdOk

`func (o *Organization) GetAuthorizedPersonnelTaxIdOk() (*string, bool)`

GetAuthorizedPersonnelTaxIdOk returns a tuple with the AuthorizedPersonnelTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedPersonnelTaxId

`func (o *Organization) SetAuthorizedPersonnelTaxId(v string)`

SetAuthorizedPersonnelTaxId sets AuthorizedPersonnelTaxId field to given value.

### HasAuthorizedPersonnelTaxId

`func (o *Organization) HasAuthorizedPersonnelTaxId() bool`

HasAuthorizedPersonnelTaxId returns a boolean if a field has been set.

### SetAuthorizedPersonnelTaxIdNil

`func (o *Organization) SetAuthorizedPersonnelTaxIdNil(b bool)`

 SetAuthorizedPersonnelTaxIdNil sets the value for AuthorizedPersonnelTaxId to be an explicit nil

### UnsetAuthorizedPersonnelTaxId
`func (o *Organization) UnsetAuthorizedPersonnelTaxId()`

UnsetAuthorizedPersonnelTaxId ensures that no value is present for AuthorizedPersonnelTaxId, not even an explicit nil
### GetAuthorizedPersonnelName

`func (o *Organization) GetAuthorizedPersonnelName() string`

GetAuthorizedPersonnelName returns the AuthorizedPersonnelName field if non-nil, zero value otherwise.

### GetAuthorizedPersonnelNameOk

`func (o *Organization) GetAuthorizedPersonnelNameOk() (*string, bool)`

GetAuthorizedPersonnelNameOk returns a tuple with the AuthorizedPersonnelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedPersonnelName

`func (o *Organization) SetAuthorizedPersonnelName(v string)`

SetAuthorizedPersonnelName sets AuthorizedPersonnelName field to given value.

### HasAuthorizedPersonnelName

`func (o *Organization) HasAuthorizedPersonnelName() bool`

HasAuthorizedPersonnelName returns a boolean if a field has been set.

### SetAuthorizedPersonnelNameNil

`func (o *Organization) SetAuthorizedPersonnelNameNil(b bool)`

 SetAuthorizedPersonnelNameNil sets the value for AuthorizedPersonnelName to be an explicit nil

### UnsetAuthorizedPersonnelName
`func (o *Organization) UnsetAuthorizedPersonnelName()`

UnsetAuthorizedPersonnelName ensures that no value is present for AuthorizedPersonnelName, not even an explicit nil
### GetRegistrationId

`func (o *Organization) GetRegistrationId() string`

GetRegistrationId returns the RegistrationId field if non-nil, zero value otherwise.

### GetRegistrationIdOk

`func (o *Organization) GetRegistrationIdOk() (*string, bool)`

GetRegistrationIdOk returns a tuple with the RegistrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationId

`func (o *Organization) SetRegistrationId(v string)`

SetRegistrationId sets RegistrationId field to given value.

### HasRegistrationId

`func (o *Organization) HasRegistrationId() bool`

HasRegistrationId returns a boolean if a field has been set.

### SetRegistrationIdNil

`func (o *Organization) SetRegistrationIdNil(b bool)`

 SetRegistrationIdNil sets the value for RegistrationId to be an explicit nil

### UnsetRegistrationId
`func (o *Organization) UnsetRegistrationId()`

UnsetRegistrationId ensures that no value is present for RegistrationId, not even an explicit nil
### GetVatId

`func (o *Organization) GetVatId() string`

GetVatId returns the VatId field if non-nil, zero value otherwise.

### GetVatIdOk

`func (o *Organization) GetVatIdOk() (*string, bool)`

GetVatIdOk returns a tuple with the VatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatId

`func (o *Organization) SetVatId(v string)`

SetVatId sets VatId field to given value.

### HasVatId

`func (o *Organization) HasVatId() bool`

HasVatId returns a boolean if a field has been set.

### SetVatIdNil

`func (o *Organization) SetVatIdNil(b bool)`

 SetVatIdNil sets the value for VatId to be an explicit nil

### UnsetVatId
`func (o *Organization) UnsetVatId()`

UnsetVatId ensures that no value is present for VatId, not even an explicit nil
### GetTaxId

`func (o *Organization) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *Organization) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *Organization) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *Organization) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### SetTaxIdNil

`func (o *Organization) SetTaxIdNil(b bool)`

 SetTaxIdNil sets the value for TaxId to be an explicit nil

### UnsetTaxId
`func (o *Organization) UnsetTaxId()`

UnsetTaxId ensures that no value is present for TaxId, not even an explicit nil
### GetDescription

`func (o *Organization) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Organization) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Organization) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Organization) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Organization) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Organization) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLogoUrl

`func (o *Organization) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *Organization) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *Organization) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *Organization) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### SetLogoUrlNil

`func (o *Organization) SetLogoUrlNil(b bool)`

 SetLogoUrlNil sets the value for LogoUrl to be an explicit nil

### UnsetLogoUrl
`func (o *Organization) UnsetLogoUrl()`

UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
### GetContactEmail

`func (o *Organization) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *Organization) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *Organization) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *Organization) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### SetContactEmailNil

`func (o *Organization) SetContactEmailNil(b bool)`

 SetContactEmailNil sets the value for ContactEmail to be an explicit nil

### UnsetContactEmail
`func (o *Organization) UnsetContactEmail()`

UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
### GetPhoneNumber

`func (o *Organization) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *Organization) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *Organization) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *Organization) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *Organization) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *Organization) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetCountryId

`func (o *Organization) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *Organization) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *Organization) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *Organization) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### GetStateId

`func (o *Organization) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *Organization) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *Organization) SetStateId(v string)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *Organization) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### GetStatus

`func (o *Organization) GetStatus() OrganizationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Organization) GetStatusOk() (*OrganizationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Organization) SetStatus(v OrganizationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Organization) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCompanyCode

`func (o *Organization) GetCompanyCode() string`

GetCompanyCode returns the CompanyCode field if non-nil, zero value otherwise.

### GetCompanyCodeOk

`func (o *Organization) GetCompanyCodeOk() (*string, bool)`

GetCompanyCodeOk returns a tuple with the CompanyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyCode

`func (o *Organization) SetCompanyCode(v string)`

SetCompanyCode sets CompanyCode field to given value.

### HasCompanyCode

`func (o *Organization) HasCompanyCode() bool`

HasCompanyCode returns a boolean if a field has been set.

### SetCompanyCodeNil

`func (o *Organization) SetCompanyCodeNil(b bool)`

 SetCompanyCodeNil sets the value for CompanyCode to be an explicit nil

### UnsetCompanyCode
`func (o *Organization) UnsetCompanyCode()`

UnsetCompanyCode ensures that no value is present for CompanyCode, not even an explicit nil
### GetPrivacyPolicyUrl

`func (o *Organization) GetPrivacyPolicyUrl() string`

GetPrivacyPolicyUrl returns the PrivacyPolicyUrl field if non-nil, zero value otherwise.

### GetPrivacyPolicyUrlOk

`func (o *Organization) GetPrivacyPolicyUrlOk() (*string, bool)`

GetPrivacyPolicyUrlOk returns a tuple with the PrivacyPolicyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicyUrl

`func (o *Organization) SetPrivacyPolicyUrl(v string)`

SetPrivacyPolicyUrl sets PrivacyPolicyUrl field to given value.

### HasPrivacyPolicyUrl

`func (o *Organization) HasPrivacyPolicyUrl() bool`

HasPrivacyPolicyUrl returns a boolean if a field has been set.

### SetPrivacyPolicyUrlNil

`func (o *Organization) SetPrivacyPolicyUrlNil(b bool)`

 SetPrivacyPolicyUrlNil sets the value for PrivacyPolicyUrl to be an explicit nil

### UnsetPrivacyPolicyUrl
`func (o *Organization) UnsetPrivacyPolicyUrl()`

UnsetPrivacyPolicyUrl ensures that no value is present for PrivacyPolicyUrl, not even an explicit nil
### GetTermsOfServiceUrl

`func (o *Organization) GetTermsOfServiceUrl() string`

GetTermsOfServiceUrl returns the TermsOfServiceUrl field if non-nil, zero value otherwise.

### GetTermsOfServiceUrlOk

`func (o *Organization) GetTermsOfServiceUrlOk() (*string, bool)`

GetTermsOfServiceUrlOk returns a tuple with the TermsOfServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfServiceUrl

`func (o *Organization) SetTermsOfServiceUrl(v string)`

SetTermsOfServiceUrl sets TermsOfServiceUrl field to given value.

### HasTermsOfServiceUrl

`func (o *Organization) HasTermsOfServiceUrl() bool`

HasTermsOfServiceUrl returns a boolean if a field has been set.

### SetTermsOfServiceUrlNil

`func (o *Organization) SetTermsOfServiceUrlNil(b bool)`

 SetTermsOfServiceUrlNil sets the value for TermsOfServiceUrl to be an explicit nil

### UnsetTermsOfServiceUrl
`func (o *Organization) UnsetTermsOfServiceUrl()`

UnsetTermsOfServiceUrl ensures that no value is present for TermsOfServiceUrl, not even an explicit nil
### GetWebsiteUrl

`func (o *Organization) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *Organization) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *Organization) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *Organization) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### SetWebsiteUrlNil

`func (o *Organization) SetWebsiteUrlNil(b bool)`

 SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil

### UnsetWebsiteUrl
`func (o *Organization) UnsetWebsiteUrl()`

UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
### GetIsGovernmentOrganization

`func (o *Organization) GetIsGovernmentOrganization() bool`

GetIsGovernmentOrganization returns the IsGovernmentOrganization field if non-nil, zero value otherwise.

### GetIsGovernmentOrganizationOk

`func (o *Organization) GetIsGovernmentOrganizationOk() (*bool, bool)`

GetIsGovernmentOrganizationOk returns a tuple with the IsGovernmentOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGovernmentOrganization

`func (o *Organization) SetIsGovernmentOrganization(v bool)`

SetIsGovernmentOrganization sets IsGovernmentOrganization field to given value.

### HasIsGovernmentOrganization

`func (o *Organization) HasIsGovernmentOrganization() bool`

HasIsGovernmentOrganization returns a boolean if a field has been set.

### GetDlApiKey

`func (o *Organization) GetDlApiKey() string`

GetDlApiKey returns the DlApiKey field if non-nil, zero value otherwise.

### GetDlApiKeyOk

`func (o *Organization) GetDlApiKeyOk() (*string, bool)`

GetDlApiKeyOk returns a tuple with the DlApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlApiKey

`func (o *Organization) SetDlApiKey(v string)`

SetDlApiKey sets DlApiKey field to given value.

### HasDlApiKey

`func (o *Organization) HasDlApiKey() bool`

HasDlApiKey returns a boolean if a field has been set.

### SetDlApiKeyNil

`func (o *Organization) SetDlApiKeyNil(b bool)`

 SetDlApiKeyNil sets the value for DlApiKey to be an explicit nil

### UnsetDlApiKey
`func (o *Organization) UnsetDlApiKey()`

UnsetDlApiKey ensures that no value is present for DlApiKey, not even an explicit nil
### GetIsKyoCompleted

`func (o *Organization) GetIsKyoCompleted() bool`

GetIsKyoCompleted returns the IsKyoCompleted field if non-nil, zero value otherwise.

### GetIsKyoCompletedOk

`func (o *Organization) GetIsKyoCompletedOk() (*bool, bool)`

GetIsKyoCompletedOk returns a tuple with the IsKyoCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKyoCompleted

`func (o *Organization) SetIsKyoCompleted(v bool)`

SetIsKyoCompleted sets IsKyoCompleted field to given value.

### HasIsKyoCompleted

`func (o *Organization) HasIsKyoCompleted() bool`

HasIsKyoCompleted returns a boolean if a field has been set.

### GetIsEnabled

`func (o *Organization) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *Organization) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *Organization) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *Organization) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsDataProvider

`func (o *Organization) GetIsDataProvider() bool`

GetIsDataProvider returns the IsDataProvider field if non-nil, zero value otherwise.

### GetIsDataProviderOk

`func (o *Organization) GetIsDataProviderOk() (*bool, bool)`

GetIsDataProviderOk returns a tuple with the IsDataProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDataProvider

`func (o *Organization) SetIsDataProvider(v bool)`

SetIsDataProvider sets IsDataProvider field to given value.

### HasIsDataProvider

`func (o *Organization) HasIsDataProvider() bool`

HasIsDataProvider returns a boolean if a field has been set.

### GetIsDataConsumer

`func (o *Organization) GetIsDataConsumer() bool`

GetIsDataConsumer returns the IsDataConsumer field if non-nil, zero value otherwise.

### GetIsDataConsumerOk

`func (o *Organization) GetIsDataConsumerOk() (*bool, bool)`

GetIsDataConsumerOk returns a tuple with the IsDataConsumer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDataConsumer

`func (o *Organization) SetIsDataConsumer(v bool)`

SetIsDataConsumer sets IsDataConsumer field to given value.

### HasIsDataConsumer

`func (o *Organization) HasIsDataConsumer() bool`

HasIsDataConsumer returns a boolean if a field has been set.

### GetSubmittedAtUtc

`func (o *Organization) GetSubmittedAtUtc() time.Time`

GetSubmittedAtUtc returns the SubmittedAtUtc field if non-nil, zero value otherwise.

### GetSubmittedAtUtcOk

`func (o *Organization) GetSubmittedAtUtcOk() (*time.Time, bool)`

GetSubmittedAtUtcOk returns a tuple with the SubmittedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAtUtc

`func (o *Organization) SetSubmittedAtUtc(v time.Time)`

SetSubmittedAtUtc sets SubmittedAtUtc field to given value.

### HasSubmittedAtUtc

`func (o *Organization) HasSubmittedAtUtc() bool`

HasSubmittedAtUtc returns a boolean if a field has been set.

### SetSubmittedAtUtcNil

`func (o *Organization) SetSubmittedAtUtcNil(b bool)`

 SetSubmittedAtUtcNil sets the value for SubmittedAtUtc to be an explicit nil

### UnsetSubmittedAtUtc
`func (o *Organization) UnsetSubmittedAtUtc()`

UnsetSubmittedAtUtc ensures that no value is present for SubmittedAtUtc, not even an explicit nil
### GetApprovedBy

`func (o *Organization) GetApprovedBy() string`

GetApprovedBy returns the ApprovedBy field if non-nil, zero value otherwise.

### GetApprovedByOk

`func (o *Organization) GetApprovedByOk() (*string, bool)`

GetApprovedByOk returns a tuple with the ApprovedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedBy

`func (o *Organization) SetApprovedBy(v string)`

SetApprovedBy sets ApprovedBy field to given value.

### HasApprovedBy

`func (o *Organization) HasApprovedBy() bool`

HasApprovedBy returns a boolean if a field has been set.

### SetApprovedByNil

`func (o *Organization) SetApprovedByNil(b bool)`

 SetApprovedByNil sets the value for ApprovedBy to be an explicit nil

### UnsetApprovedBy
`func (o *Organization) UnsetApprovedBy()`

UnsetApprovedBy ensures that no value is present for ApprovedBy, not even an explicit nil
### GetApprovedAtUtc

`func (o *Organization) GetApprovedAtUtc() time.Time`

GetApprovedAtUtc returns the ApprovedAtUtc field if non-nil, zero value otherwise.

### GetApprovedAtUtcOk

`func (o *Organization) GetApprovedAtUtcOk() (*time.Time, bool)`

GetApprovedAtUtcOk returns a tuple with the ApprovedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAtUtc

`func (o *Organization) SetApprovedAtUtc(v time.Time)`

SetApprovedAtUtc sets ApprovedAtUtc field to given value.

### HasApprovedAtUtc

`func (o *Organization) HasApprovedAtUtc() bool`

HasApprovedAtUtc returns a boolean if a field has been set.

### SetApprovedAtUtcNil

`func (o *Organization) SetApprovedAtUtcNil(b bool)`

 SetApprovedAtUtcNil sets the value for ApprovedAtUtc to be an explicit nil

### UnsetApprovedAtUtc
`func (o *Organization) UnsetApprovedAtUtc()`

UnsetApprovedAtUtc ensures that no value is present for ApprovedAtUtc, not even an explicit nil
### GetIsDigiLockerOrganization

`func (o *Organization) GetIsDigiLockerOrganization() bool`

GetIsDigiLockerOrganization returns the IsDigiLockerOrganization field if non-nil, zero value otherwise.

### GetIsDigiLockerOrganizationOk

`func (o *Organization) GetIsDigiLockerOrganizationOk() (*bool, bool)`

GetIsDigiLockerOrganizationOk returns a tuple with the IsDigiLockerOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDigiLockerOrganization

`func (o *Organization) SetIsDigiLockerOrganization(v bool)`

SetIsDigiLockerOrganization sets IsDigiLockerOrganization field to given value.

### HasIsDigiLockerOrganization

`func (o *Organization) HasIsDigiLockerOrganization() bool`

HasIsDigiLockerOrganization returns a boolean if a field has been set.

### GetIsMdmcMaintained

`func (o *Organization) GetIsMdmcMaintained() bool`

GetIsMdmcMaintained returns the IsMdmcMaintained field if non-nil, zero value otherwise.

### GetIsMdmcMaintainedOk

`func (o *Organization) GetIsMdmcMaintainedOk() (*bool, bool)`

GetIsMdmcMaintainedOk returns a tuple with the IsMdmcMaintained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMdmcMaintained

`func (o *Organization) SetIsMdmcMaintained(v bool)`

SetIsMdmcMaintained sets IsMdmcMaintained field to given value.

### HasIsMdmcMaintained

`func (o *Organization) HasIsMdmcMaintained() bool`

HasIsMdmcMaintained returns a boolean if a field has been set.

### GetIsBbps

`func (o *Organization) GetIsBbps() bool`

GetIsBbps returns the IsBbps field if non-nil, zero value otherwise.

### GetIsBbpsOk

`func (o *Organization) GetIsBbpsOk() (*bool, bool)`

GetIsBbpsOk returns a tuple with the IsBbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBbps

`func (o *Organization) SetIsBbps(v bool)`

SetIsBbps sets IsBbps field to given value.

### HasIsBbps

`func (o *Organization) HasIsBbps() bool`

HasIsBbps returns a boolean if a field has been set.

### GetMetaData

`func (o *Organization) GetMetaData() OrganizationMetaData`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *Organization) GetMetaDataOk() (*OrganizationMetaData, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *Organization) SetMetaData(v OrganizationMetaData)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *Organization) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetCountryCode

`func (o *Organization) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Organization) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Organization) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *Organization) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *Organization) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *Organization) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetOrganizationType

`func (o *Organization) GetOrganizationType() OrganizationType`

GetOrganizationType returns the OrganizationType field if non-nil, zero value otherwise.

### GetOrganizationTypeOk

`func (o *Organization) GetOrganizationTypeOk() (*OrganizationType, bool)`

GetOrganizationTypeOk returns a tuple with the OrganizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationType

`func (o *Organization) SetOrganizationType(v OrganizationType)`

SetOrganizationType sets OrganizationType field to given value.

### HasOrganizationType

`func (o *Organization) HasOrganizationType() bool`

HasOrganizationType returns a boolean if a field has been set.

### GetOrganizationCategory

`func (o *Organization) GetOrganizationCategory() OrganizationCategory`

GetOrganizationCategory returns the OrganizationCategory field if non-nil, zero value otherwise.

### GetOrganizationCategoryOk

`func (o *Organization) GetOrganizationCategoryOk() (*OrganizationCategory, bool)`

GetOrganizationCategoryOk returns a tuple with the OrganizationCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationCategory

`func (o *Organization) SetOrganizationCategory(v OrganizationCategory)`

SetOrganizationCategory sets OrganizationCategory field to given value.

### HasOrganizationCategory

`func (o *Organization) HasOrganizationCategory() bool`

HasOrganizationCategory returns a boolean if a field has been set.

### GetDocumentProviderCategory

`func (o *Organization) GetDocumentProviderCategory() DocumentProviderCategory`

GetDocumentProviderCategory returns the DocumentProviderCategory field if non-nil, zero value otherwise.

### GetDocumentProviderCategoryOk

`func (o *Organization) GetDocumentProviderCategoryOk() (*DocumentProviderCategory, bool)`

GetDocumentProviderCategoryOk returns a tuple with the DocumentProviderCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentProviderCategory

`func (o *Organization) SetDocumentProviderCategory(v DocumentProviderCategory)`

SetDocumentProviderCategory sets DocumentProviderCategory field to given value.

### HasDocumentProviderCategory

`func (o *Organization) HasDocumentProviderCategory() bool`

HasDocumentProviderCategory returns a boolean if a field has been set.

### GetAddresses

`func (o *Organization) GetAddresses() []OrganizationAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *Organization) GetAddressesOk() (*[]OrganizationAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *Organization) SetAddresses(v []OrganizationAddress)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *Organization) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### SetAddressesNil

`func (o *Organization) SetAddressesNil(b bool)`

 SetAddressesNil sets the value for Addresses to be an explicit nil

### UnsetAddresses
`func (o *Organization) UnsetAddresses()`

UnsetAddresses ensures that no value is present for Addresses, not even an explicit nil
### GetFinancialAccounts

`func (o *Organization) GetFinancialAccounts() []OrganizationFinancialAccount`

GetFinancialAccounts returns the FinancialAccounts field if non-nil, zero value otherwise.

### GetFinancialAccountsOk

`func (o *Organization) GetFinancialAccountsOk() (*[]OrganizationFinancialAccount, bool)`

GetFinancialAccountsOk returns a tuple with the FinancialAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialAccounts

`func (o *Organization) SetFinancialAccounts(v []OrganizationFinancialAccount)`

SetFinancialAccounts sets FinancialAccounts field to given value.

### HasFinancialAccounts

`func (o *Organization) HasFinancialAccounts() bool`

HasFinancialAccounts returns a boolean if a field has been set.

### SetFinancialAccountsNil

`func (o *Organization) SetFinancialAccountsNil(b bool)`

 SetFinancialAccountsNil sets the value for FinancialAccounts to be an explicit nil

### UnsetFinancialAccounts
`func (o *Organization) UnsetFinancialAccounts()`

UnsetFinancialAccounts ensures that no value is present for FinancialAccounts, not even an explicit nil
### GetCountryState

`func (o *Organization) GetCountryState() Country`

GetCountryState returns the CountryState field if non-nil, zero value otherwise.

### GetCountryStateOk

`func (o *Organization) GetCountryStateOk() (*Country, bool)`

GetCountryStateOk returns a tuple with the CountryState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryState

`func (o *Organization) SetCountryState(v Country)`

SetCountryState sets CountryState field to given value.

### HasCountryState

`func (o *Organization) HasCountryState() bool`

HasCountryState returns a boolean if a field has been set.

### GetApprovedByUser

`func (o *Organization) GetApprovedByUser() ApplicationUser`

GetApprovedByUser returns the ApprovedByUser field if non-nil, zero value otherwise.

### GetApprovedByUserOk

`func (o *Organization) GetApprovedByUserOk() (*ApplicationUser, bool)`

GetApprovedByUserOk returns a tuple with the ApprovedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedByUser

`func (o *Organization) SetApprovedByUser(v ApplicationUser)`

SetApprovedByUser sets ApprovedByUser field to given value.

### HasApprovedByUser

`func (o *Organization) HasApprovedByUser() bool`

HasApprovedByUser returns a boolean if a field has been set.

### GetKyoDocuments

`func (o *Organization) GetKyoDocuments() []OrganizationKyoDocument`

GetKyoDocuments returns the KyoDocuments field if non-nil, zero value otherwise.

### GetKyoDocumentsOk

`func (o *Organization) GetKyoDocumentsOk() (*[]OrganizationKyoDocument, bool)`

GetKyoDocumentsOk returns a tuple with the KyoDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKyoDocuments

`func (o *Organization) SetKyoDocuments(v []OrganizationKyoDocument)`

SetKyoDocuments sets KyoDocuments field to given value.

### HasKyoDocuments

`func (o *Organization) HasKyoDocuments() bool`

HasKyoDocuments returns a boolean if a field has been set.

### SetKyoDocumentsNil

`func (o *Organization) SetKyoDocumentsNil(b bool)`

 SetKyoDocumentsNil sets the value for KyoDocuments to be an explicit nil

### UnsetKyoDocuments
`func (o *Organization) UnsetKyoDocuments()`

UnsetKyoDocuments ensures that no value is present for KyoDocuments, not even an explicit nil
### GetIsDelete

`func (o *Organization) GetIsDelete() bool`

GetIsDelete returns the IsDelete field if non-nil, zero value otherwise.

### GetIsDeleteOk

`func (o *Organization) GetIsDeleteOk() (*bool, bool)`

GetIsDeleteOk returns a tuple with the IsDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDelete

`func (o *Organization) SetIsDelete(v bool)`

SetIsDelete sets IsDelete field to given value.

### HasIsDelete

`func (o *Organization) HasIsDelete() bool`

HasIsDelete returns a boolean if a field has been set.

### GetRecoveryToken

`func (o *Organization) GetRecoveryToken() string`

GetRecoveryToken returns the RecoveryToken field if non-nil, zero value otherwise.

### GetRecoveryTokenOk

`func (o *Organization) GetRecoveryTokenOk() (*string, bool)`

GetRecoveryTokenOk returns a tuple with the RecoveryToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryToken

`func (o *Organization) SetRecoveryToken(v string)`

SetRecoveryToken sets RecoveryToken field to given value.

### HasRecoveryToken

`func (o *Organization) HasRecoveryToken() bool`

HasRecoveryToken returns a boolean if a field has been set.

### SetRecoveryTokenNil

`func (o *Organization) SetRecoveryTokenNil(b bool)`

 SetRecoveryTokenNil sets the value for RecoveryToken to be an explicit nil

### UnsetRecoveryToken
`func (o *Organization) UnsetRecoveryToken()`

UnsetRecoveryToken ensures that no value is present for RecoveryToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



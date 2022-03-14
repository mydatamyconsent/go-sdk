# DataProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Category** | **string** |  | 
**LogoUrl** | Pointer to **NullableString** |  | [optional] 
**Website** | Pointer to **NullableString** |  | [optional] 
**SupportEmail** | Pointer to **NullableString** |  | [optional] 
**HelpLineNumber** | Pointer to **NullableString** |  | [optional] 
**PrivacyPolicy** | Pointer to **NullableString** |  | [optional] 
**TermOfService** | Pointer to **NullableString** |  | [optional] 
**DataProtectionOfficer** | Pointer to [**DataProtectionOfficer**](DataProtectionOfficer.md) |  | [optional] 
**SupportedDocumentTypes** | **[]string** |  | 
**SupportedAccountTypes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDataProvider

`func NewDataProvider(id string, name string, category string, supportedDocumentTypes []string, ) *DataProvider`

NewDataProvider instantiates a new DataProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataProviderWithDefaults

`func NewDataProviderWithDefaults() *DataProvider`

NewDataProviderWithDefaults instantiates a new DataProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataProvider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataProvider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataProvider) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DataProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataProvider) SetName(v string)`

SetName sets Name field to given value.


### GetCategory

`func (o *DataProvider) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *DataProvider) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *DataProvider) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetLogoUrl

`func (o *DataProvider) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *DataProvider) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *DataProvider) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *DataProvider) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### SetLogoUrlNil

`func (o *DataProvider) SetLogoUrlNil(b bool)`

 SetLogoUrlNil sets the value for LogoUrl to be an explicit nil

### UnsetLogoUrl
`func (o *DataProvider) UnsetLogoUrl()`

UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
### GetWebsite

`func (o *DataProvider) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *DataProvider) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *DataProvider) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *DataProvider) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### SetWebsiteNil

`func (o *DataProvider) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *DataProvider) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil
### GetSupportEmail

`func (o *DataProvider) GetSupportEmail() string`

GetSupportEmail returns the SupportEmail field if non-nil, zero value otherwise.

### GetSupportEmailOk

`func (o *DataProvider) GetSupportEmailOk() (*string, bool)`

GetSupportEmailOk returns a tuple with the SupportEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmail

`func (o *DataProvider) SetSupportEmail(v string)`

SetSupportEmail sets SupportEmail field to given value.

### HasSupportEmail

`func (o *DataProvider) HasSupportEmail() bool`

HasSupportEmail returns a boolean if a field has been set.

### SetSupportEmailNil

`func (o *DataProvider) SetSupportEmailNil(b bool)`

 SetSupportEmailNil sets the value for SupportEmail to be an explicit nil

### UnsetSupportEmail
`func (o *DataProvider) UnsetSupportEmail()`

UnsetSupportEmail ensures that no value is present for SupportEmail, not even an explicit nil
### GetHelpLineNumber

`func (o *DataProvider) GetHelpLineNumber() string`

GetHelpLineNumber returns the HelpLineNumber field if non-nil, zero value otherwise.

### GetHelpLineNumberOk

`func (o *DataProvider) GetHelpLineNumberOk() (*string, bool)`

GetHelpLineNumberOk returns a tuple with the HelpLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpLineNumber

`func (o *DataProvider) SetHelpLineNumber(v string)`

SetHelpLineNumber sets HelpLineNumber field to given value.

### HasHelpLineNumber

`func (o *DataProvider) HasHelpLineNumber() bool`

HasHelpLineNumber returns a boolean if a field has been set.

### SetHelpLineNumberNil

`func (o *DataProvider) SetHelpLineNumberNil(b bool)`

 SetHelpLineNumberNil sets the value for HelpLineNumber to be an explicit nil

### UnsetHelpLineNumber
`func (o *DataProvider) UnsetHelpLineNumber()`

UnsetHelpLineNumber ensures that no value is present for HelpLineNumber, not even an explicit nil
### GetPrivacyPolicy

`func (o *DataProvider) GetPrivacyPolicy() string`

GetPrivacyPolicy returns the PrivacyPolicy field if non-nil, zero value otherwise.

### GetPrivacyPolicyOk

`func (o *DataProvider) GetPrivacyPolicyOk() (*string, bool)`

GetPrivacyPolicyOk returns a tuple with the PrivacyPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicy

`func (o *DataProvider) SetPrivacyPolicy(v string)`

SetPrivacyPolicy sets PrivacyPolicy field to given value.

### HasPrivacyPolicy

`func (o *DataProvider) HasPrivacyPolicy() bool`

HasPrivacyPolicy returns a boolean if a field has been set.

### SetPrivacyPolicyNil

`func (o *DataProvider) SetPrivacyPolicyNil(b bool)`

 SetPrivacyPolicyNil sets the value for PrivacyPolicy to be an explicit nil

### UnsetPrivacyPolicy
`func (o *DataProvider) UnsetPrivacyPolicy()`

UnsetPrivacyPolicy ensures that no value is present for PrivacyPolicy, not even an explicit nil
### GetTermOfService

`func (o *DataProvider) GetTermOfService() string`

GetTermOfService returns the TermOfService field if non-nil, zero value otherwise.

### GetTermOfServiceOk

`func (o *DataProvider) GetTermOfServiceOk() (*string, bool)`

GetTermOfServiceOk returns a tuple with the TermOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermOfService

`func (o *DataProvider) SetTermOfService(v string)`

SetTermOfService sets TermOfService field to given value.

### HasTermOfService

`func (o *DataProvider) HasTermOfService() bool`

HasTermOfService returns a boolean if a field has been set.

### SetTermOfServiceNil

`func (o *DataProvider) SetTermOfServiceNil(b bool)`

 SetTermOfServiceNil sets the value for TermOfService to be an explicit nil

### UnsetTermOfService
`func (o *DataProvider) UnsetTermOfService()`

UnsetTermOfService ensures that no value is present for TermOfService, not even an explicit nil
### GetDataProtectionOfficer

`func (o *DataProvider) GetDataProtectionOfficer() DataProtectionOfficer`

GetDataProtectionOfficer returns the DataProtectionOfficer field if non-nil, zero value otherwise.

### GetDataProtectionOfficerOk

`func (o *DataProvider) GetDataProtectionOfficerOk() (*DataProtectionOfficer, bool)`

GetDataProtectionOfficerOk returns a tuple with the DataProtectionOfficer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProtectionOfficer

`func (o *DataProvider) SetDataProtectionOfficer(v DataProtectionOfficer)`

SetDataProtectionOfficer sets DataProtectionOfficer field to given value.

### HasDataProtectionOfficer

`func (o *DataProvider) HasDataProtectionOfficer() bool`

HasDataProtectionOfficer returns a boolean if a field has been set.

### GetSupportedDocumentTypes

`func (o *DataProvider) GetSupportedDocumentTypes() []string`

GetSupportedDocumentTypes returns the SupportedDocumentTypes field if non-nil, zero value otherwise.

### GetSupportedDocumentTypesOk

`func (o *DataProvider) GetSupportedDocumentTypesOk() (*[]string, bool)`

GetSupportedDocumentTypesOk returns a tuple with the SupportedDocumentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDocumentTypes

`func (o *DataProvider) SetSupportedDocumentTypes(v []string)`

SetSupportedDocumentTypes sets SupportedDocumentTypes field to given value.


### GetSupportedAccountTypes

`func (o *DataProvider) GetSupportedAccountTypes() []string`

GetSupportedAccountTypes returns the SupportedAccountTypes field if non-nil, zero value otherwise.

### GetSupportedAccountTypesOk

`func (o *DataProvider) GetSupportedAccountTypesOk() (*[]string, bool)`

GetSupportedAccountTypesOk returns a tuple with the SupportedAccountTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedAccountTypes

`func (o *DataProvider) SetSupportedAccountTypes(v []string)`

SetSupportedAccountTypes sets SupportedAccountTypes field to given value.

### HasSupportedAccountTypes

`func (o *DataProvider) HasSupportedAccountTypes() bool`

HasSupportedAccountTypes returns a boolean if a field has been set.

### SetSupportedAccountTypesNil

`func (o *DataProvider) SetSupportedAccountTypesNil(b bool)`

 SetSupportedAccountTypesNil sets the value for SupportedAccountTypes to be an explicit nil

### UnsetSupportedAccountTypes
`func (o *DataProvider) UnsetSupportedAccountTypes()`

UnsetSupportedAccountTypes ensures that no value is present for SupportedAccountTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



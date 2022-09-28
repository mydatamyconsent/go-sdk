# DataProviderDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Category** | **string** |  | 
**LogoUrl** | Pointer to **string** |  | [optional] 
**Website** | Pointer to **string** |  | [optional] 
**SupportEmail** | Pointer to **string** |  | [optional] 
**HelpLineNumber** | Pointer to **string** |  | [optional] 
**PrivacyPolicy** | Pointer to **string** |  | [optional] 
**TermOfService** | Pointer to **string** |  | [optional] 
**DataProtectionOfficer** | [**DataProtectionOfficer**](DataProtectionOfficer.md) |  | 
**SupportedDocumentTypes** | [**[]SupportedDocumentType**](SupportedDocumentType.md) |  | 
**SupportedAccountTypes** | **[]string** |  | 

## Methods

### NewDataProviderDetails

`func NewDataProviderDetails(id string, name string, category string, dataProtectionOfficer DataProtectionOfficer, supportedDocumentTypes []SupportedDocumentType, supportedAccountTypes []string, ) *DataProviderDetails`

NewDataProviderDetails instantiates a new DataProviderDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataProviderDetailsWithDefaults

`func NewDataProviderDetailsWithDefaults() *DataProviderDetails`

NewDataProviderDetailsWithDefaults instantiates a new DataProviderDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataProviderDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataProviderDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataProviderDetails) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DataProviderDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataProviderDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataProviderDetails) SetName(v string)`

SetName sets Name field to given value.


### GetCategory

`func (o *DataProviderDetails) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *DataProviderDetails) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *DataProviderDetails) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetLogoUrl

`func (o *DataProviderDetails) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *DataProviderDetails) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *DataProviderDetails) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *DataProviderDetails) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetWebsite

`func (o *DataProviderDetails) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *DataProviderDetails) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *DataProviderDetails) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *DataProviderDetails) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetSupportEmail

`func (o *DataProviderDetails) GetSupportEmail() string`

GetSupportEmail returns the SupportEmail field if non-nil, zero value otherwise.

### GetSupportEmailOk

`func (o *DataProviderDetails) GetSupportEmailOk() (*string, bool)`

GetSupportEmailOk returns a tuple with the SupportEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmail

`func (o *DataProviderDetails) SetSupportEmail(v string)`

SetSupportEmail sets SupportEmail field to given value.

### HasSupportEmail

`func (o *DataProviderDetails) HasSupportEmail() bool`

HasSupportEmail returns a boolean if a field has been set.

### GetHelpLineNumber

`func (o *DataProviderDetails) GetHelpLineNumber() string`

GetHelpLineNumber returns the HelpLineNumber field if non-nil, zero value otherwise.

### GetHelpLineNumberOk

`func (o *DataProviderDetails) GetHelpLineNumberOk() (*string, bool)`

GetHelpLineNumberOk returns a tuple with the HelpLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpLineNumber

`func (o *DataProviderDetails) SetHelpLineNumber(v string)`

SetHelpLineNumber sets HelpLineNumber field to given value.

### HasHelpLineNumber

`func (o *DataProviderDetails) HasHelpLineNumber() bool`

HasHelpLineNumber returns a boolean if a field has been set.

### GetPrivacyPolicy

`func (o *DataProviderDetails) GetPrivacyPolicy() string`

GetPrivacyPolicy returns the PrivacyPolicy field if non-nil, zero value otherwise.

### GetPrivacyPolicyOk

`func (o *DataProviderDetails) GetPrivacyPolicyOk() (*string, bool)`

GetPrivacyPolicyOk returns a tuple with the PrivacyPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicy

`func (o *DataProviderDetails) SetPrivacyPolicy(v string)`

SetPrivacyPolicy sets PrivacyPolicy field to given value.

### HasPrivacyPolicy

`func (o *DataProviderDetails) HasPrivacyPolicy() bool`

HasPrivacyPolicy returns a boolean if a field has been set.

### GetTermOfService

`func (o *DataProviderDetails) GetTermOfService() string`

GetTermOfService returns the TermOfService field if non-nil, zero value otherwise.

### GetTermOfServiceOk

`func (o *DataProviderDetails) GetTermOfServiceOk() (*string, bool)`

GetTermOfServiceOk returns a tuple with the TermOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermOfService

`func (o *DataProviderDetails) SetTermOfService(v string)`

SetTermOfService sets TermOfService field to given value.

### HasTermOfService

`func (o *DataProviderDetails) HasTermOfService() bool`

HasTermOfService returns a boolean if a field has been set.

### GetDataProtectionOfficer

`func (o *DataProviderDetails) GetDataProtectionOfficer() DataProtectionOfficer`

GetDataProtectionOfficer returns the DataProtectionOfficer field if non-nil, zero value otherwise.

### GetDataProtectionOfficerOk

`func (o *DataProviderDetails) GetDataProtectionOfficerOk() (*DataProtectionOfficer, bool)`

GetDataProtectionOfficerOk returns a tuple with the DataProtectionOfficer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProtectionOfficer

`func (o *DataProviderDetails) SetDataProtectionOfficer(v DataProtectionOfficer)`

SetDataProtectionOfficer sets DataProtectionOfficer field to given value.


### GetSupportedDocumentTypes

`func (o *DataProviderDetails) GetSupportedDocumentTypes() []SupportedDocumentType`

GetSupportedDocumentTypes returns the SupportedDocumentTypes field if non-nil, zero value otherwise.

### GetSupportedDocumentTypesOk

`func (o *DataProviderDetails) GetSupportedDocumentTypesOk() (*[]SupportedDocumentType, bool)`

GetSupportedDocumentTypesOk returns a tuple with the SupportedDocumentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDocumentTypes

`func (o *DataProviderDetails) SetSupportedDocumentTypes(v []SupportedDocumentType)`

SetSupportedDocumentTypes sets SupportedDocumentTypes field to given value.


### GetSupportedAccountTypes

`func (o *DataProviderDetails) GetSupportedAccountTypes() []string`

GetSupportedAccountTypes returns the SupportedAccountTypes field if non-nil, zero value otherwise.

### GetSupportedAccountTypesOk

`func (o *DataProviderDetails) GetSupportedAccountTypesOk() (*[]string, bool)`

GetSupportedAccountTypesOk returns a tuple with the SupportedAccountTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedAccountTypes

`func (o *DataProviderDetails) SetSupportedAccountTypes(v []string)`

SetSupportedAccountTypes sets SupportedAccountTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



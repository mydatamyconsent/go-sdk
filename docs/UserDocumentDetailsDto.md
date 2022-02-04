# UserDocumentDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CategoryType** | Pointer to [**DocumentCategoryType**](DocumentCategoryType.md) |  | [optional] 
**TypeId** | Pointer to **string** |  | [optional] 
**TypeName** | Pointer to **NullableString** |  | [optional] 
**FullName** | Pointer to **NullableString** |  | [optional] 
**Identifier** | Pointer to **NullableString** |  | [optional] 
**AccountId** | Pointer to **NullableString** |  | [optional] 
**IssuerId** | Pointer to **string** |  | [optional] 
**IssuerName** | Pointer to **NullableString** |  | [optional] 
**StorageUrl** | Pointer to **NullableString** |  | [optional] 
**IsQuickAccessEnabled** | Pointer to **bool** |  | [optional] 
**IsOwner** | Pointer to **bool** |  | [optional] 
**DigitalSignatureDetails** | Pointer to [**[]DigitalSignature**](DigitalSignature.md) |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 

## Methods

### NewUserDocumentDetailsDto

`func NewUserDocumentDetailsDto() *UserDocumentDetailsDto`

NewUserDocumentDetailsDto instantiates a new UserDocumentDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDocumentDetailsDtoWithDefaults

`func NewUserDocumentDetailsDtoWithDefaults() *UserDocumentDetailsDto`

NewUserDocumentDetailsDtoWithDefaults instantiates a new UserDocumentDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserDocumentDetailsDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserDocumentDetailsDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserDocumentDetailsDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserDocumentDetailsDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCategoryType

`func (o *UserDocumentDetailsDto) GetCategoryType() DocumentCategoryType`

GetCategoryType returns the CategoryType field if non-nil, zero value otherwise.

### GetCategoryTypeOk

`func (o *UserDocumentDetailsDto) GetCategoryTypeOk() (*DocumentCategoryType, bool)`

GetCategoryTypeOk returns a tuple with the CategoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryType

`func (o *UserDocumentDetailsDto) SetCategoryType(v DocumentCategoryType)`

SetCategoryType sets CategoryType field to given value.

### HasCategoryType

`func (o *UserDocumentDetailsDto) HasCategoryType() bool`

HasCategoryType returns a boolean if a field has been set.

### GetTypeId

`func (o *UserDocumentDetailsDto) GetTypeId() string`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *UserDocumentDetailsDto) GetTypeIdOk() (*string, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *UserDocumentDetailsDto) SetTypeId(v string)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *UserDocumentDetailsDto) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetTypeName

`func (o *UserDocumentDetailsDto) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *UserDocumentDetailsDto) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *UserDocumentDetailsDto) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *UserDocumentDetailsDto) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### SetTypeNameNil

`func (o *UserDocumentDetailsDto) SetTypeNameNil(b bool)`

 SetTypeNameNil sets the value for TypeName to be an explicit nil

### UnsetTypeName
`func (o *UserDocumentDetailsDto) UnsetTypeName()`

UnsetTypeName ensures that no value is present for TypeName, not even an explicit nil
### GetFullName

`func (o *UserDocumentDetailsDto) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *UserDocumentDetailsDto) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *UserDocumentDetailsDto) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *UserDocumentDetailsDto) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *UserDocumentDetailsDto) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *UserDocumentDetailsDto) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetIdentifier

`func (o *UserDocumentDetailsDto) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *UserDocumentDetailsDto) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *UserDocumentDetailsDto) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *UserDocumentDetailsDto) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *UserDocumentDetailsDto) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *UserDocumentDetailsDto) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetAccountId

`func (o *UserDocumentDetailsDto) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UserDocumentDetailsDto) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UserDocumentDetailsDto) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UserDocumentDetailsDto) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *UserDocumentDetailsDto) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *UserDocumentDetailsDto) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetIssuerId

`func (o *UserDocumentDetailsDto) GetIssuerId() string`

GetIssuerId returns the IssuerId field if non-nil, zero value otherwise.

### GetIssuerIdOk

`func (o *UserDocumentDetailsDto) GetIssuerIdOk() (*string, bool)`

GetIssuerIdOk returns a tuple with the IssuerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerId

`func (o *UserDocumentDetailsDto) SetIssuerId(v string)`

SetIssuerId sets IssuerId field to given value.

### HasIssuerId

`func (o *UserDocumentDetailsDto) HasIssuerId() bool`

HasIssuerId returns a boolean if a field has been set.

### GetIssuerName

`func (o *UserDocumentDetailsDto) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *UserDocumentDetailsDto) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *UserDocumentDetailsDto) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.

### HasIssuerName

`func (o *UserDocumentDetailsDto) HasIssuerName() bool`

HasIssuerName returns a boolean if a field has been set.

### SetIssuerNameNil

`func (o *UserDocumentDetailsDto) SetIssuerNameNil(b bool)`

 SetIssuerNameNil sets the value for IssuerName to be an explicit nil

### UnsetIssuerName
`func (o *UserDocumentDetailsDto) UnsetIssuerName()`

UnsetIssuerName ensures that no value is present for IssuerName, not even an explicit nil
### GetStorageUrl

`func (o *UserDocumentDetailsDto) GetStorageUrl() string`

GetStorageUrl returns the StorageUrl field if non-nil, zero value otherwise.

### GetStorageUrlOk

`func (o *UserDocumentDetailsDto) GetStorageUrlOk() (*string, bool)`

GetStorageUrlOk returns a tuple with the StorageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUrl

`func (o *UserDocumentDetailsDto) SetStorageUrl(v string)`

SetStorageUrl sets StorageUrl field to given value.

### HasStorageUrl

`func (o *UserDocumentDetailsDto) HasStorageUrl() bool`

HasStorageUrl returns a boolean if a field has been set.

### SetStorageUrlNil

`func (o *UserDocumentDetailsDto) SetStorageUrlNil(b bool)`

 SetStorageUrlNil sets the value for StorageUrl to be an explicit nil

### UnsetStorageUrl
`func (o *UserDocumentDetailsDto) UnsetStorageUrl()`

UnsetStorageUrl ensures that no value is present for StorageUrl, not even an explicit nil
### GetIsQuickAccessEnabled

`func (o *UserDocumentDetailsDto) GetIsQuickAccessEnabled() bool`

GetIsQuickAccessEnabled returns the IsQuickAccessEnabled field if non-nil, zero value otherwise.

### GetIsQuickAccessEnabledOk

`func (o *UserDocumentDetailsDto) GetIsQuickAccessEnabledOk() (*bool, bool)`

GetIsQuickAccessEnabledOk returns a tuple with the IsQuickAccessEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQuickAccessEnabled

`func (o *UserDocumentDetailsDto) SetIsQuickAccessEnabled(v bool)`

SetIsQuickAccessEnabled sets IsQuickAccessEnabled field to given value.

### HasIsQuickAccessEnabled

`func (o *UserDocumentDetailsDto) HasIsQuickAccessEnabled() bool`

HasIsQuickAccessEnabled returns a boolean if a field has been set.

### GetIsOwner

`func (o *UserDocumentDetailsDto) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *UserDocumentDetailsDto) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *UserDocumentDetailsDto) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.

### HasIsOwner

`func (o *UserDocumentDetailsDto) HasIsOwner() bool`

HasIsOwner returns a boolean if a field has been set.

### GetDigitalSignatureDetails

`func (o *UserDocumentDetailsDto) GetDigitalSignatureDetails() []DigitalSignature`

GetDigitalSignatureDetails returns the DigitalSignatureDetails field if non-nil, zero value otherwise.

### GetDigitalSignatureDetailsOk

`func (o *UserDocumentDetailsDto) GetDigitalSignatureDetailsOk() (*[]DigitalSignature, bool)`

GetDigitalSignatureDetailsOk returns a tuple with the DigitalSignatureDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalSignatureDetails

`func (o *UserDocumentDetailsDto) SetDigitalSignatureDetails(v []DigitalSignature)`

SetDigitalSignatureDetails sets DigitalSignatureDetails field to given value.

### HasDigitalSignatureDetails

`func (o *UserDocumentDetailsDto) HasDigitalSignatureDetails() bool`

HasDigitalSignatureDetails returns a boolean if a field has been set.

### SetDigitalSignatureDetailsNil

`func (o *UserDocumentDetailsDto) SetDigitalSignatureDetailsNil(b bool)`

 SetDigitalSignatureDetailsNil sets the value for DigitalSignatureDetails to be an explicit nil

### UnsetDigitalSignatureDetails
`func (o *UserDocumentDetailsDto) UnsetDigitalSignatureDetails()`

UnsetDigitalSignatureDetails ensures that no value is present for DigitalSignatureDetails, not even an explicit nil
### GetOwnerId

`func (o *UserDocumentDetailsDto) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *UserDocumentDetailsDto) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *UserDocumentDetailsDto) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *UserDocumentDetailsDto) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



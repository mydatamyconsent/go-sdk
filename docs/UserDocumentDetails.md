# UserDocumentDetails

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

### NewUserDocumentDetails

`func NewUserDocumentDetails() *UserDocumentDetails`

NewUserDocumentDetails instantiates a new UserDocumentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDocumentDetailsWithDefaults

`func NewUserDocumentDetailsWithDefaults() *UserDocumentDetails`

NewUserDocumentDetailsWithDefaults instantiates a new UserDocumentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserDocumentDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserDocumentDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserDocumentDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserDocumentDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCategoryType

`func (o *UserDocumentDetails) GetCategoryType() DocumentCategoryType`

GetCategoryType returns the CategoryType field if non-nil, zero value otherwise.

### GetCategoryTypeOk

`func (o *UserDocumentDetails) GetCategoryTypeOk() (*DocumentCategoryType, bool)`

GetCategoryTypeOk returns a tuple with the CategoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryType

`func (o *UserDocumentDetails) SetCategoryType(v DocumentCategoryType)`

SetCategoryType sets CategoryType field to given value.

### HasCategoryType

`func (o *UserDocumentDetails) HasCategoryType() bool`

HasCategoryType returns a boolean if a field has been set.

### GetTypeId

`func (o *UserDocumentDetails) GetTypeId() string`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *UserDocumentDetails) GetTypeIdOk() (*string, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *UserDocumentDetails) SetTypeId(v string)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *UserDocumentDetails) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetTypeName

`func (o *UserDocumentDetails) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *UserDocumentDetails) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *UserDocumentDetails) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *UserDocumentDetails) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### SetTypeNameNil

`func (o *UserDocumentDetails) SetTypeNameNil(b bool)`

 SetTypeNameNil sets the value for TypeName to be an explicit nil

### UnsetTypeName
`func (o *UserDocumentDetails) UnsetTypeName()`

UnsetTypeName ensures that no value is present for TypeName, not even an explicit nil
### GetFullName

`func (o *UserDocumentDetails) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *UserDocumentDetails) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *UserDocumentDetails) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *UserDocumentDetails) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *UserDocumentDetails) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *UserDocumentDetails) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetIdentifier

`func (o *UserDocumentDetails) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *UserDocumentDetails) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *UserDocumentDetails) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *UserDocumentDetails) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *UserDocumentDetails) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *UserDocumentDetails) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetAccountId

`func (o *UserDocumentDetails) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UserDocumentDetails) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UserDocumentDetails) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UserDocumentDetails) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *UserDocumentDetails) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *UserDocumentDetails) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetIssuerId

`func (o *UserDocumentDetails) GetIssuerId() string`

GetIssuerId returns the IssuerId field if non-nil, zero value otherwise.

### GetIssuerIdOk

`func (o *UserDocumentDetails) GetIssuerIdOk() (*string, bool)`

GetIssuerIdOk returns a tuple with the IssuerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerId

`func (o *UserDocumentDetails) SetIssuerId(v string)`

SetIssuerId sets IssuerId field to given value.

### HasIssuerId

`func (o *UserDocumentDetails) HasIssuerId() bool`

HasIssuerId returns a boolean if a field has been set.

### GetIssuerName

`func (o *UserDocumentDetails) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *UserDocumentDetails) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *UserDocumentDetails) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.

### HasIssuerName

`func (o *UserDocumentDetails) HasIssuerName() bool`

HasIssuerName returns a boolean if a field has been set.

### SetIssuerNameNil

`func (o *UserDocumentDetails) SetIssuerNameNil(b bool)`

 SetIssuerNameNil sets the value for IssuerName to be an explicit nil

### UnsetIssuerName
`func (o *UserDocumentDetails) UnsetIssuerName()`

UnsetIssuerName ensures that no value is present for IssuerName, not even an explicit nil
### GetStorageUrl

`func (o *UserDocumentDetails) GetStorageUrl() string`

GetStorageUrl returns the StorageUrl field if non-nil, zero value otherwise.

### GetStorageUrlOk

`func (o *UserDocumentDetails) GetStorageUrlOk() (*string, bool)`

GetStorageUrlOk returns a tuple with the StorageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUrl

`func (o *UserDocumentDetails) SetStorageUrl(v string)`

SetStorageUrl sets StorageUrl field to given value.

### HasStorageUrl

`func (o *UserDocumentDetails) HasStorageUrl() bool`

HasStorageUrl returns a boolean if a field has been set.

### SetStorageUrlNil

`func (o *UserDocumentDetails) SetStorageUrlNil(b bool)`

 SetStorageUrlNil sets the value for StorageUrl to be an explicit nil

### UnsetStorageUrl
`func (o *UserDocumentDetails) UnsetStorageUrl()`

UnsetStorageUrl ensures that no value is present for StorageUrl, not even an explicit nil
### GetIsQuickAccessEnabled

`func (o *UserDocumentDetails) GetIsQuickAccessEnabled() bool`

GetIsQuickAccessEnabled returns the IsQuickAccessEnabled field if non-nil, zero value otherwise.

### GetIsQuickAccessEnabledOk

`func (o *UserDocumentDetails) GetIsQuickAccessEnabledOk() (*bool, bool)`

GetIsQuickAccessEnabledOk returns a tuple with the IsQuickAccessEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQuickAccessEnabled

`func (o *UserDocumentDetails) SetIsQuickAccessEnabled(v bool)`

SetIsQuickAccessEnabled sets IsQuickAccessEnabled field to given value.

### HasIsQuickAccessEnabled

`func (o *UserDocumentDetails) HasIsQuickAccessEnabled() bool`

HasIsQuickAccessEnabled returns a boolean if a field has been set.

### GetIsOwner

`func (o *UserDocumentDetails) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *UserDocumentDetails) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *UserDocumentDetails) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.

### HasIsOwner

`func (o *UserDocumentDetails) HasIsOwner() bool`

HasIsOwner returns a boolean if a field has been set.

### GetDigitalSignatureDetails

`func (o *UserDocumentDetails) GetDigitalSignatureDetails() []DigitalSignature`

GetDigitalSignatureDetails returns the DigitalSignatureDetails field if non-nil, zero value otherwise.

### GetDigitalSignatureDetailsOk

`func (o *UserDocumentDetails) GetDigitalSignatureDetailsOk() (*[]DigitalSignature, bool)`

GetDigitalSignatureDetailsOk returns a tuple with the DigitalSignatureDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalSignatureDetails

`func (o *UserDocumentDetails) SetDigitalSignatureDetails(v []DigitalSignature)`

SetDigitalSignatureDetails sets DigitalSignatureDetails field to given value.

### HasDigitalSignatureDetails

`func (o *UserDocumentDetails) HasDigitalSignatureDetails() bool`

HasDigitalSignatureDetails returns a boolean if a field has been set.

### SetDigitalSignatureDetailsNil

`func (o *UserDocumentDetails) SetDigitalSignatureDetailsNil(b bool)`

 SetDigitalSignatureDetailsNil sets the value for DigitalSignatureDetails to be an explicit nil

### UnsetDigitalSignatureDetails
`func (o *UserDocumentDetails) UnsetDigitalSignatureDetails()`

UnsetDigitalSignatureDetails ensures that no value is present for DigitalSignatureDetails, not even an explicit nil
### GetOwnerId

`func (o *UserDocumentDetails) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *UserDocumentDetails) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *UserDocumentDetails) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *UserDocumentDetails) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



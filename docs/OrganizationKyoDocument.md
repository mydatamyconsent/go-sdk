# OrganizationKyoDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**DocumentName** | Pointer to **NullableString** |  | [optional] 
**StorageUrl** | Pointer to **NullableString** |  | [optional] 
**UploadedAtUtc** | Pointer to **time.Time** |  | [optional] 
**VerifiedBy** | Pointer to **NullableString** |  | [optional] 
**VerifiedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**DeletedBy** | Pointer to **NullableString** |  | [optional] 
**DeletedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**Organization** | Pointer to [**Organization**](Organization.md) |  | [optional] 
**VerifiedByUser** | Pointer to [**ApplicationUser**](ApplicationUser.md) |  | [optional] 
**DeletedByUser** | Pointer to [**ApplicationUser**](ApplicationUser.md) |  | [optional] 
**Rejection** | Pointer to [**Rejection**](Rejection.md) |  | [optional] 

## Methods

### NewOrganizationKyoDocument

`func NewOrganizationKyoDocument() *OrganizationKyoDocument`

NewOrganizationKyoDocument instantiates a new OrganizationKyoDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationKyoDocumentWithDefaults

`func NewOrganizationKyoDocumentWithDefaults() *OrganizationKyoDocument`

NewOrganizationKyoDocumentWithDefaults instantiates a new OrganizationKyoDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationKyoDocument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationKyoDocument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationKyoDocument) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationKyoDocument) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *OrganizationKyoDocument) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OrganizationKyoDocument) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OrganizationKyoDocument) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *OrganizationKyoDocument) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetDocumentName

`func (o *OrganizationKyoDocument) GetDocumentName() string`

GetDocumentName returns the DocumentName field if non-nil, zero value otherwise.

### GetDocumentNameOk

`func (o *OrganizationKyoDocument) GetDocumentNameOk() (*string, bool)`

GetDocumentNameOk returns a tuple with the DocumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentName

`func (o *OrganizationKyoDocument) SetDocumentName(v string)`

SetDocumentName sets DocumentName field to given value.

### HasDocumentName

`func (o *OrganizationKyoDocument) HasDocumentName() bool`

HasDocumentName returns a boolean if a field has been set.

### SetDocumentNameNil

`func (o *OrganizationKyoDocument) SetDocumentNameNil(b bool)`

 SetDocumentNameNil sets the value for DocumentName to be an explicit nil

### UnsetDocumentName
`func (o *OrganizationKyoDocument) UnsetDocumentName()`

UnsetDocumentName ensures that no value is present for DocumentName, not even an explicit nil
### GetStorageUrl

`func (o *OrganizationKyoDocument) GetStorageUrl() string`

GetStorageUrl returns the StorageUrl field if non-nil, zero value otherwise.

### GetStorageUrlOk

`func (o *OrganizationKyoDocument) GetStorageUrlOk() (*string, bool)`

GetStorageUrlOk returns a tuple with the StorageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUrl

`func (o *OrganizationKyoDocument) SetStorageUrl(v string)`

SetStorageUrl sets StorageUrl field to given value.

### HasStorageUrl

`func (o *OrganizationKyoDocument) HasStorageUrl() bool`

HasStorageUrl returns a boolean if a field has been set.

### SetStorageUrlNil

`func (o *OrganizationKyoDocument) SetStorageUrlNil(b bool)`

 SetStorageUrlNil sets the value for StorageUrl to be an explicit nil

### UnsetStorageUrl
`func (o *OrganizationKyoDocument) UnsetStorageUrl()`

UnsetStorageUrl ensures that no value is present for StorageUrl, not even an explicit nil
### GetUploadedAtUtc

`func (o *OrganizationKyoDocument) GetUploadedAtUtc() time.Time`

GetUploadedAtUtc returns the UploadedAtUtc field if non-nil, zero value otherwise.

### GetUploadedAtUtcOk

`func (o *OrganizationKyoDocument) GetUploadedAtUtcOk() (*time.Time, bool)`

GetUploadedAtUtcOk returns a tuple with the UploadedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedAtUtc

`func (o *OrganizationKyoDocument) SetUploadedAtUtc(v time.Time)`

SetUploadedAtUtc sets UploadedAtUtc field to given value.

### HasUploadedAtUtc

`func (o *OrganizationKyoDocument) HasUploadedAtUtc() bool`

HasUploadedAtUtc returns a boolean if a field has been set.

### GetVerifiedBy

`func (o *OrganizationKyoDocument) GetVerifiedBy() string`

GetVerifiedBy returns the VerifiedBy field if non-nil, zero value otherwise.

### GetVerifiedByOk

`func (o *OrganizationKyoDocument) GetVerifiedByOk() (*string, bool)`

GetVerifiedByOk returns a tuple with the VerifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedBy

`func (o *OrganizationKyoDocument) SetVerifiedBy(v string)`

SetVerifiedBy sets VerifiedBy field to given value.

### HasVerifiedBy

`func (o *OrganizationKyoDocument) HasVerifiedBy() bool`

HasVerifiedBy returns a boolean if a field has been set.

### SetVerifiedByNil

`func (o *OrganizationKyoDocument) SetVerifiedByNil(b bool)`

 SetVerifiedByNil sets the value for VerifiedBy to be an explicit nil

### UnsetVerifiedBy
`func (o *OrganizationKyoDocument) UnsetVerifiedBy()`

UnsetVerifiedBy ensures that no value is present for VerifiedBy, not even an explicit nil
### GetVerifiedAtUtc

`func (o *OrganizationKyoDocument) GetVerifiedAtUtc() time.Time`

GetVerifiedAtUtc returns the VerifiedAtUtc field if non-nil, zero value otherwise.

### GetVerifiedAtUtcOk

`func (o *OrganizationKyoDocument) GetVerifiedAtUtcOk() (*time.Time, bool)`

GetVerifiedAtUtcOk returns a tuple with the VerifiedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAtUtc

`func (o *OrganizationKyoDocument) SetVerifiedAtUtc(v time.Time)`

SetVerifiedAtUtc sets VerifiedAtUtc field to given value.

### HasVerifiedAtUtc

`func (o *OrganizationKyoDocument) HasVerifiedAtUtc() bool`

HasVerifiedAtUtc returns a boolean if a field has been set.

### SetVerifiedAtUtcNil

`func (o *OrganizationKyoDocument) SetVerifiedAtUtcNil(b bool)`

 SetVerifiedAtUtcNil sets the value for VerifiedAtUtc to be an explicit nil

### UnsetVerifiedAtUtc
`func (o *OrganizationKyoDocument) UnsetVerifiedAtUtc()`

UnsetVerifiedAtUtc ensures that no value is present for VerifiedAtUtc, not even an explicit nil
### GetDeletedBy

`func (o *OrganizationKyoDocument) GetDeletedBy() string`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *OrganizationKyoDocument) GetDeletedByOk() (*string, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *OrganizationKyoDocument) SetDeletedBy(v string)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *OrganizationKyoDocument) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### SetDeletedByNil

`func (o *OrganizationKyoDocument) SetDeletedByNil(b bool)`

 SetDeletedByNil sets the value for DeletedBy to be an explicit nil

### UnsetDeletedBy
`func (o *OrganizationKyoDocument) UnsetDeletedBy()`

UnsetDeletedBy ensures that no value is present for DeletedBy, not even an explicit nil
### GetDeletedAtUtc

`func (o *OrganizationKyoDocument) GetDeletedAtUtc() time.Time`

GetDeletedAtUtc returns the DeletedAtUtc field if non-nil, zero value otherwise.

### GetDeletedAtUtcOk

`func (o *OrganizationKyoDocument) GetDeletedAtUtcOk() (*time.Time, bool)`

GetDeletedAtUtcOk returns a tuple with the DeletedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAtUtc

`func (o *OrganizationKyoDocument) SetDeletedAtUtc(v time.Time)`

SetDeletedAtUtc sets DeletedAtUtc field to given value.

### HasDeletedAtUtc

`func (o *OrganizationKyoDocument) HasDeletedAtUtc() bool`

HasDeletedAtUtc returns a boolean if a field has been set.

### SetDeletedAtUtcNil

`func (o *OrganizationKyoDocument) SetDeletedAtUtcNil(b bool)`

 SetDeletedAtUtcNil sets the value for DeletedAtUtc to be an explicit nil

### UnsetDeletedAtUtc
`func (o *OrganizationKyoDocument) UnsetDeletedAtUtc()`

UnsetDeletedAtUtc ensures that no value is present for DeletedAtUtc, not even an explicit nil
### GetOrganization

`func (o *OrganizationKyoDocument) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *OrganizationKyoDocument) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *OrganizationKyoDocument) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *OrganizationKyoDocument) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetVerifiedByUser

`func (o *OrganizationKyoDocument) GetVerifiedByUser() ApplicationUser`

GetVerifiedByUser returns the VerifiedByUser field if non-nil, zero value otherwise.

### GetVerifiedByUserOk

`func (o *OrganizationKyoDocument) GetVerifiedByUserOk() (*ApplicationUser, bool)`

GetVerifiedByUserOk returns a tuple with the VerifiedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedByUser

`func (o *OrganizationKyoDocument) SetVerifiedByUser(v ApplicationUser)`

SetVerifiedByUser sets VerifiedByUser field to given value.

### HasVerifiedByUser

`func (o *OrganizationKyoDocument) HasVerifiedByUser() bool`

HasVerifiedByUser returns a boolean if a field has been set.

### GetDeletedByUser

`func (o *OrganizationKyoDocument) GetDeletedByUser() ApplicationUser`

GetDeletedByUser returns the DeletedByUser field if non-nil, zero value otherwise.

### GetDeletedByUserOk

`func (o *OrganizationKyoDocument) GetDeletedByUserOk() (*ApplicationUser, bool)`

GetDeletedByUserOk returns a tuple with the DeletedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedByUser

`func (o *OrganizationKyoDocument) SetDeletedByUser(v ApplicationUser)`

SetDeletedByUser sets DeletedByUser field to given value.

### HasDeletedByUser

`func (o *OrganizationKyoDocument) HasDeletedByUser() bool`

HasDeletedByUser returns a boolean if a field has been set.

### GetRejection

`func (o *OrganizationKyoDocument) GetRejection() Rejection`

GetRejection returns the Rejection field if non-nil, zero value otherwise.

### GetRejectionOk

`func (o *OrganizationKyoDocument) GetRejectionOk() (*Rejection, bool)`

GetRejectionOk returns a tuple with the Rejection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejection

`func (o *OrganizationKyoDocument) SetRejection(v Rejection)`

SetRejection sets Rejection field to given value.

### HasRejection

`func (o *OrganizationKyoDocument) HasRejection() bool`

HasRejection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



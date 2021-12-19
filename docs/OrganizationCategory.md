# OrganizationCategory

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
**Name** | Pointer to **NullableString** |  | [optional] 
**LogoUrl** | Pointer to **NullableString** |  | [optional] 
**MetaData** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewOrganizationCategory

`func NewOrganizationCategory() *OrganizationCategory`

NewOrganizationCategory instantiates a new OrganizationCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCategoryWithDefaults

`func NewOrganizationCategoryWithDefaults() *OrganizationCategory`

NewOrganizationCategoryWithDefaults instantiates a new OrganizationCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *OrganizationCategory) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *OrganizationCategory) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *OrganizationCategory) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *OrganizationCategory) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedAtUtc

`func (o *OrganizationCategory) GetCreatedAtUtc() time.Time`

GetCreatedAtUtc returns the CreatedAtUtc field if non-nil, zero value otherwise.

### GetCreatedAtUtcOk

`func (o *OrganizationCategory) GetCreatedAtUtcOk() (*time.Time, bool)`

GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtUtc

`func (o *OrganizationCategory) SetCreatedAtUtc(v time.Time)`

SetCreatedAtUtc sets CreatedAtUtc field to given value.

### HasCreatedAtUtc

`func (o *OrganizationCategory) HasCreatedAtUtc() bool`

HasCreatedAtUtc returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *OrganizationCategory) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *OrganizationCategory) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *OrganizationCategory) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *OrganizationCategory) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *OrganizationCategory) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *OrganizationCategory) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetUpdatedAtUtc

`func (o *OrganizationCategory) GetUpdatedAtUtc() time.Time`

GetUpdatedAtUtc returns the UpdatedAtUtc field if non-nil, zero value otherwise.

### GetUpdatedAtUtcOk

`func (o *OrganizationCategory) GetUpdatedAtUtcOk() (*time.Time, bool)`

GetUpdatedAtUtcOk returns a tuple with the UpdatedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtUtc

`func (o *OrganizationCategory) SetUpdatedAtUtc(v time.Time)`

SetUpdatedAtUtc sets UpdatedAtUtc field to given value.

### HasUpdatedAtUtc

`func (o *OrganizationCategory) HasUpdatedAtUtc() bool`

HasUpdatedAtUtc returns a boolean if a field has been set.

### SetUpdatedAtUtcNil

`func (o *OrganizationCategory) SetUpdatedAtUtcNil(b bool)`

 SetUpdatedAtUtcNil sets the value for UpdatedAtUtc to be an explicit nil

### UnsetUpdatedAtUtc
`func (o *OrganizationCategory) UnsetUpdatedAtUtc()`

UnsetUpdatedAtUtc ensures that no value is present for UpdatedAtUtc, not even an explicit nil
### GetCreatedByUser

`func (o *OrganizationCategory) GetCreatedByUser() ApplicationUser`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *OrganizationCategory) GetCreatedByUserOk() (*ApplicationUser, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *OrganizationCategory) SetCreatedByUser(v ApplicationUser)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *OrganizationCategory) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### GetUpdatedByUser

`func (o *OrganizationCategory) GetUpdatedByUser() ApplicationUser`

GetUpdatedByUser returns the UpdatedByUser field if non-nil, zero value otherwise.

### GetUpdatedByUserOk

`func (o *OrganizationCategory) GetUpdatedByUserOk() (*ApplicationUser, bool)`

GetUpdatedByUserOk returns a tuple with the UpdatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUser

`func (o *OrganizationCategory) SetUpdatedByUser(v ApplicationUser)`

SetUpdatedByUser sets UpdatedByUser field to given value.

### HasUpdatedByUser

`func (o *OrganizationCategory) HasUpdatedByUser() bool`

HasUpdatedByUser returns a boolean if a field has been set.

### GetDeletedBy

`func (o *OrganizationCategory) GetDeletedBy() string`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *OrganizationCategory) GetDeletedByOk() (*string, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *OrganizationCategory) SetDeletedBy(v string)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *OrganizationCategory) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### SetDeletedByNil

`func (o *OrganizationCategory) SetDeletedByNil(b bool)`

 SetDeletedByNil sets the value for DeletedBy to be an explicit nil

### UnsetDeletedBy
`func (o *OrganizationCategory) UnsetDeletedBy()`

UnsetDeletedBy ensures that no value is present for DeletedBy, not even an explicit nil
### GetDeletedAtUtc

`func (o *OrganizationCategory) GetDeletedAtUtc() time.Time`

GetDeletedAtUtc returns the DeletedAtUtc field if non-nil, zero value otherwise.

### GetDeletedAtUtcOk

`func (o *OrganizationCategory) GetDeletedAtUtcOk() (*time.Time, bool)`

GetDeletedAtUtcOk returns a tuple with the DeletedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAtUtc

`func (o *OrganizationCategory) SetDeletedAtUtc(v time.Time)`

SetDeletedAtUtc sets DeletedAtUtc field to given value.

### HasDeletedAtUtc

`func (o *OrganizationCategory) HasDeletedAtUtc() bool`

HasDeletedAtUtc returns a boolean if a field has been set.

### SetDeletedAtUtcNil

`func (o *OrganizationCategory) SetDeletedAtUtcNil(b bool)`

 SetDeletedAtUtcNil sets the value for DeletedAtUtc to be an explicit nil

### UnsetDeletedAtUtc
`func (o *OrganizationCategory) UnsetDeletedAtUtc()`

UnsetDeletedAtUtc ensures that no value is present for DeletedAtUtc, not even an explicit nil
### GetDeletedByUser

`func (o *OrganizationCategory) GetDeletedByUser() ApplicationUser`

GetDeletedByUser returns the DeletedByUser field if non-nil, zero value otherwise.

### GetDeletedByUserOk

`func (o *OrganizationCategory) GetDeletedByUserOk() (*ApplicationUser, bool)`

GetDeletedByUserOk returns a tuple with the DeletedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedByUser

`func (o *OrganizationCategory) SetDeletedByUser(v ApplicationUser)`

SetDeletedByUser sets DeletedByUser field to given value.

### HasDeletedByUser

`func (o *OrganizationCategory) HasDeletedByUser() bool`

HasDeletedByUser returns a boolean if a field has been set.

### GetId

`func (o *OrganizationCategory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationCategory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationCategory) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationCategory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OrganizationCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationCategory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationCategory) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OrganizationCategory) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OrganizationCategory) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLogoUrl

`func (o *OrganizationCategory) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *OrganizationCategory) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *OrganizationCategory) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *OrganizationCategory) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### SetLogoUrlNil

`func (o *OrganizationCategory) SetLogoUrlNil(b bool)`

 SetLogoUrlNil sets the value for LogoUrl to be an explicit nil

### UnsetLogoUrl
`func (o *OrganizationCategory) UnsetLogoUrl()`

UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
### GetMetaData

`func (o *OrganizationCategory) GetMetaData() interface{}`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *OrganizationCategory) GetMetaDataOk() (*interface{}, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *OrganizationCategory) SetMetaData(v interface{})`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *OrganizationCategory) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### SetMetaDataNil

`func (o *OrganizationCategory) SetMetaDataNil(b bool)`

 SetMetaDataNil sets the value for MetaData to be an explicit nil

### UnsetMetaData
`func (o *OrganizationCategory) UnsetMetaData()`

UnsetMetaData ensures that no value is present for MetaData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DocumentProviderCategory

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
**IconCodePoint** | Pointer to **int32** |  | [optional] 
**IconFontFamily** | Pointer to [**CategoryIconFontFamily**](CategoryIconFontFamily.md) |  | [optional] 
**MetaData** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewDocumentProviderCategory

`func NewDocumentProviderCategory() *DocumentProviderCategory`

NewDocumentProviderCategory instantiates a new DocumentProviderCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentProviderCategoryWithDefaults

`func NewDocumentProviderCategoryWithDefaults() *DocumentProviderCategory`

NewDocumentProviderCategoryWithDefaults instantiates a new DocumentProviderCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *DocumentProviderCategory) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DocumentProviderCategory) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DocumentProviderCategory) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DocumentProviderCategory) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedAtUtc

`func (o *DocumentProviderCategory) GetCreatedAtUtc() time.Time`

GetCreatedAtUtc returns the CreatedAtUtc field if non-nil, zero value otherwise.

### GetCreatedAtUtcOk

`func (o *DocumentProviderCategory) GetCreatedAtUtcOk() (*time.Time, bool)`

GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtUtc

`func (o *DocumentProviderCategory) SetCreatedAtUtc(v time.Time)`

SetCreatedAtUtc sets CreatedAtUtc field to given value.

### HasCreatedAtUtc

`func (o *DocumentProviderCategory) HasCreatedAtUtc() bool`

HasCreatedAtUtc returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *DocumentProviderCategory) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DocumentProviderCategory) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DocumentProviderCategory) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DocumentProviderCategory) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *DocumentProviderCategory) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *DocumentProviderCategory) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetUpdatedAtUtc

`func (o *DocumentProviderCategory) GetUpdatedAtUtc() time.Time`

GetUpdatedAtUtc returns the UpdatedAtUtc field if non-nil, zero value otherwise.

### GetUpdatedAtUtcOk

`func (o *DocumentProviderCategory) GetUpdatedAtUtcOk() (*time.Time, bool)`

GetUpdatedAtUtcOk returns a tuple with the UpdatedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtUtc

`func (o *DocumentProviderCategory) SetUpdatedAtUtc(v time.Time)`

SetUpdatedAtUtc sets UpdatedAtUtc field to given value.

### HasUpdatedAtUtc

`func (o *DocumentProviderCategory) HasUpdatedAtUtc() bool`

HasUpdatedAtUtc returns a boolean if a field has been set.

### SetUpdatedAtUtcNil

`func (o *DocumentProviderCategory) SetUpdatedAtUtcNil(b bool)`

 SetUpdatedAtUtcNil sets the value for UpdatedAtUtc to be an explicit nil

### UnsetUpdatedAtUtc
`func (o *DocumentProviderCategory) UnsetUpdatedAtUtc()`

UnsetUpdatedAtUtc ensures that no value is present for UpdatedAtUtc, not even an explicit nil
### GetCreatedByUser

`func (o *DocumentProviderCategory) GetCreatedByUser() ApplicationUser`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *DocumentProviderCategory) GetCreatedByUserOk() (*ApplicationUser, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *DocumentProviderCategory) SetCreatedByUser(v ApplicationUser)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *DocumentProviderCategory) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### GetUpdatedByUser

`func (o *DocumentProviderCategory) GetUpdatedByUser() ApplicationUser`

GetUpdatedByUser returns the UpdatedByUser field if non-nil, zero value otherwise.

### GetUpdatedByUserOk

`func (o *DocumentProviderCategory) GetUpdatedByUserOk() (*ApplicationUser, bool)`

GetUpdatedByUserOk returns a tuple with the UpdatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUser

`func (o *DocumentProviderCategory) SetUpdatedByUser(v ApplicationUser)`

SetUpdatedByUser sets UpdatedByUser field to given value.

### HasUpdatedByUser

`func (o *DocumentProviderCategory) HasUpdatedByUser() bool`

HasUpdatedByUser returns a boolean if a field has been set.

### GetDeletedBy

`func (o *DocumentProviderCategory) GetDeletedBy() string`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *DocumentProviderCategory) GetDeletedByOk() (*string, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *DocumentProviderCategory) SetDeletedBy(v string)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *DocumentProviderCategory) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### SetDeletedByNil

`func (o *DocumentProviderCategory) SetDeletedByNil(b bool)`

 SetDeletedByNil sets the value for DeletedBy to be an explicit nil

### UnsetDeletedBy
`func (o *DocumentProviderCategory) UnsetDeletedBy()`

UnsetDeletedBy ensures that no value is present for DeletedBy, not even an explicit nil
### GetDeletedAtUtc

`func (o *DocumentProviderCategory) GetDeletedAtUtc() time.Time`

GetDeletedAtUtc returns the DeletedAtUtc field if non-nil, zero value otherwise.

### GetDeletedAtUtcOk

`func (o *DocumentProviderCategory) GetDeletedAtUtcOk() (*time.Time, bool)`

GetDeletedAtUtcOk returns a tuple with the DeletedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAtUtc

`func (o *DocumentProviderCategory) SetDeletedAtUtc(v time.Time)`

SetDeletedAtUtc sets DeletedAtUtc field to given value.

### HasDeletedAtUtc

`func (o *DocumentProviderCategory) HasDeletedAtUtc() bool`

HasDeletedAtUtc returns a boolean if a field has been set.

### SetDeletedAtUtcNil

`func (o *DocumentProviderCategory) SetDeletedAtUtcNil(b bool)`

 SetDeletedAtUtcNil sets the value for DeletedAtUtc to be an explicit nil

### UnsetDeletedAtUtc
`func (o *DocumentProviderCategory) UnsetDeletedAtUtc()`

UnsetDeletedAtUtc ensures that no value is present for DeletedAtUtc, not even an explicit nil
### GetDeletedByUser

`func (o *DocumentProviderCategory) GetDeletedByUser() ApplicationUser`

GetDeletedByUser returns the DeletedByUser field if non-nil, zero value otherwise.

### GetDeletedByUserOk

`func (o *DocumentProviderCategory) GetDeletedByUserOk() (*ApplicationUser, bool)`

GetDeletedByUserOk returns a tuple with the DeletedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedByUser

`func (o *DocumentProviderCategory) SetDeletedByUser(v ApplicationUser)`

SetDeletedByUser sets DeletedByUser field to given value.

### HasDeletedByUser

`func (o *DocumentProviderCategory) HasDeletedByUser() bool`

HasDeletedByUser returns a boolean if a field has been set.

### GetId

`func (o *DocumentProviderCategory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentProviderCategory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentProviderCategory) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DocumentProviderCategory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DocumentProviderCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DocumentProviderCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DocumentProviderCategory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DocumentProviderCategory) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DocumentProviderCategory) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DocumentProviderCategory) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIconCodePoint

`func (o *DocumentProviderCategory) GetIconCodePoint() int32`

GetIconCodePoint returns the IconCodePoint field if non-nil, zero value otherwise.

### GetIconCodePointOk

`func (o *DocumentProviderCategory) GetIconCodePointOk() (*int32, bool)`

GetIconCodePointOk returns a tuple with the IconCodePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconCodePoint

`func (o *DocumentProviderCategory) SetIconCodePoint(v int32)`

SetIconCodePoint sets IconCodePoint field to given value.

### HasIconCodePoint

`func (o *DocumentProviderCategory) HasIconCodePoint() bool`

HasIconCodePoint returns a boolean if a field has been set.

### GetIconFontFamily

`func (o *DocumentProviderCategory) GetIconFontFamily() CategoryIconFontFamily`

GetIconFontFamily returns the IconFontFamily field if non-nil, zero value otherwise.

### GetIconFontFamilyOk

`func (o *DocumentProviderCategory) GetIconFontFamilyOk() (*CategoryIconFontFamily, bool)`

GetIconFontFamilyOk returns a tuple with the IconFontFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconFontFamily

`func (o *DocumentProviderCategory) SetIconFontFamily(v CategoryIconFontFamily)`

SetIconFontFamily sets IconFontFamily field to given value.

### HasIconFontFamily

`func (o *DocumentProviderCategory) HasIconFontFamily() bool`

HasIconFontFamily returns a boolean if a field has been set.

### GetMetaData

`func (o *DocumentProviderCategory) GetMetaData() interface{}`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *DocumentProviderCategory) GetMetaDataOk() (*interface{}, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *DocumentProviderCategory) SetMetaData(v interface{})`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *DocumentProviderCategory) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### SetMetaDataNil

`func (o *DocumentProviderCategory) SetMetaDataNil(b bool)`

 SetMetaDataNil sets the value for MetaData to be an explicit nil

### UnsetMetaData
`func (o *DocumentProviderCategory) UnsetMetaData()`

UnsetMetaData ensures that no value is present for MetaData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



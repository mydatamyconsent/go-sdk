# DataProcessingAgreement

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
**OrganizationId** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**Body** | Pointer to **NullableString** |  | [optional] 
**AttachmentUrl** | Pointer to **NullableString** |  | [optional] 
**Organization** | Pointer to [**Organization**](Organization.md) |  | [optional] 

## Methods

### NewDataProcessingAgreement

`func NewDataProcessingAgreement() *DataProcessingAgreement`

NewDataProcessingAgreement instantiates a new DataProcessingAgreement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataProcessingAgreementWithDefaults

`func NewDataProcessingAgreementWithDefaults() *DataProcessingAgreement`

NewDataProcessingAgreementWithDefaults instantiates a new DataProcessingAgreement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *DataProcessingAgreement) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DataProcessingAgreement) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DataProcessingAgreement) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DataProcessingAgreement) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedAtUtc

`func (o *DataProcessingAgreement) GetCreatedAtUtc() time.Time`

GetCreatedAtUtc returns the CreatedAtUtc field if non-nil, zero value otherwise.

### GetCreatedAtUtcOk

`func (o *DataProcessingAgreement) GetCreatedAtUtcOk() (*time.Time, bool)`

GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtUtc

`func (o *DataProcessingAgreement) SetCreatedAtUtc(v time.Time)`

SetCreatedAtUtc sets CreatedAtUtc field to given value.

### HasCreatedAtUtc

`func (o *DataProcessingAgreement) HasCreatedAtUtc() bool`

HasCreatedAtUtc returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *DataProcessingAgreement) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DataProcessingAgreement) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DataProcessingAgreement) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DataProcessingAgreement) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *DataProcessingAgreement) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *DataProcessingAgreement) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetUpdatedAtUtc

`func (o *DataProcessingAgreement) GetUpdatedAtUtc() time.Time`

GetUpdatedAtUtc returns the UpdatedAtUtc field if non-nil, zero value otherwise.

### GetUpdatedAtUtcOk

`func (o *DataProcessingAgreement) GetUpdatedAtUtcOk() (*time.Time, bool)`

GetUpdatedAtUtcOk returns a tuple with the UpdatedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtUtc

`func (o *DataProcessingAgreement) SetUpdatedAtUtc(v time.Time)`

SetUpdatedAtUtc sets UpdatedAtUtc field to given value.

### HasUpdatedAtUtc

`func (o *DataProcessingAgreement) HasUpdatedAtUtc() bool`

HasUpdatedAtUtc returns a boolean if a field has been set.

### SetUpdatedAtUtcNil

`func (o *DataProcessingAgreement) SetUpdatedAtUtcNil(b bool)`

 SetUpdatedAtUtcNil sets the value for UpdatedAtUtc to be an explicit nil

### UnsetUpdatedAtUtc
`func (o *DataProcessingAgreement) UnsetUpdatedAtUtc()`

UnsetUpdatedAtUtc ensures that no value is present for UpdatedAtUtc, not even an explicit nil
### GetCreatedByUser

`func (o *DataProcessingAgreement) GetCreatedByUser() ApplicationUser`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *DataProcessingAgreement) GetCreatedByUserOk() (*ApplicationUser, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *DataProcessingAgreement) SetCreatedByUser(v ApplicationUser)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *DataProcessingAgreement) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### GetUpdatedByUser

`func (o *DataProcessingAgreement) GetUpdatedByUser() ApplicationUser`

GetUpdatedByUser returns the UpdatedByUser field if non-nil, zero value otherwise.

### GetUpdatedByUserOk

`func (o *DataProcessingAgreement) GetUpdatedByUserOk() (*ApplicationUser, bool)`

GetUpdatedByUserOk returns a tuple with the UpdatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUser

`func (o *DataProcessingAgreement) SetUpdatedByUser(v ApplicationUser)`

SetUpdatedByUser sets UpdatedByUser field to given value.

### HasUpdatedByUser

`func (o *DataProcessingAgreement) HasUpdatedByUser() bool`

HasUpdatedByUser returns a boolean if a field has been set.

### GetDeletedBy

`func (o *DataProcessingAgreement) GetDeletedBy() string`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *DataProcessingAgreement) GetDeletedByOk() (*string, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *DataProcessingAgreement) SetDeletedBy(v string)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *DataProcessingAgreement) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### SetDeletedByNil

`func (o *DataProcessingAgreement) SetDeletedByNil(b bool)`

 SetDeletedByNil sets the value for DeletedBy to be an explicit nil

### UnsetDeletedBy
`func (o *DataProcessingAgreement) UnsetDeletedBy()`

UnsetDeletedBy ensures that no value is present for DeletedBy, not even an explicit nil
### GetDeletedAtUtc

`func (o *DataProcessingAgreement) GetDeletedAtUtc() time.Time`

GetDeletedAtUtc returns the DeletedAtUtc field if non-nil, zero value otherwise.

### GetDeletedAtUtcOk

`func (o *DataProcessingAgreement) GetDeletedAtUtcOk() (*time.Time, bool)`

GetDeletedAtUtcOk returns a tuple with the DeletedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAtUtc

`func (o *DataProcessingAgreement) SetDeletedAtUtc(v time.Time)`

SetDeletedAtUtc sets DeletedAtUtc field to given value.

### HasDeletedAtUtc

`func (o *DataProcessingAgreement) HasDeletedAtUtc() bool`

HasDeletedAtUtc returns a boolean if a field has been set.

### SetDeletedAtUtcNil

`func (o *DataProcessingAgreement) SetDeletedAtUtcNil(b bool)`

 SetDeletedAtUtcNil sets the value for DeletedAtUtc to be an explicit nil

### UnsetDeletedAtUtc
`func (o *DataProcessingAgreement) UnsetDeletedAtUtc()`

UnsetDeletedAtUtc ensures that no value is present for DeletedAtUtc, not even an explicit nil
### GetDeletedByUser

`func (o *DataProcessingAgreement) GetDeletedByUser() ApplicationUser`

GetDeletedByUser returns the DeletedByUser field if non-nil, zero value otherwise.

### GetDeletedByUserOk

`func (o *DataProcessingAgreement) GetDeletedByUserOk() (*ApplicationUser, bool)`

GetDeletedByUserOk returns a tuple with the DeletedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedByUser

`func (o *DataProcessingAgreement) SetDeletedByUser(v ApplicationUser)`

SetDeletedByUser sets DeletedByUser field to given value.

### HasDeletedByUser

`func (o *DataProcessingAgreement) HasDeletedByUser() bool`

HasDeletedByUser returns a boolean if a field has been set.

### GetId

`func (o *DataProcessingAgreement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataProcessingAgreement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataProcessingAgreement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataProcessingAgreement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *DataProcessingAgreement) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *DataProcessingAgreement) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *DataProcessingAgreement) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *DataProcessingAgreement) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetVersion

`func (o *DataProcessingAgreement) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DataProcessingAgreement) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DataProcessingAgreement) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DataProcessingAgreement) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *DataProcessingAgreement) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *DataProcessingAgreement) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetBody

`func (o *DataProcessingAgreement) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *DataProcessingAgreement) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *DataProcessingAgreement) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *DataProcessingAgreement) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *DataProcessingAgreement) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *DataProcessingAgreement) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetAttachmentUrl

`func (o *DataProcessingAgreement) GetAttachmentUrl() string`

GetAttachmentUrl returns the AttachmentUrl field if non-nil, zero value otherwise.

### GetAttachmentUrlOk

`func (o *DataProcessingAgreement) GetAttachmentUrlOk() (*string, bool)`

GetAttachmentUrlOk returns a tuple with the AttachmentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentUrl

`func (o *DataProcessingAgreement) SetAttachmentUrl(v string)`

SetAttachmentUrl sets AttachmentUrl field to given value.

### HasAttachmentUrl

`func (o *DataProcessingAgreement) HasAttachmentUrl() bool`

HasAttachmentUrl returns a boolean if a field has been set.

### SetAttachmentUrlNil

`func (o *DataProcessingAgreement) SetAttachmentUrlNil(b bool)`

 SetAttachmentUrlNil sets the value for AttachmentUrl to be an explicit nil

### UnsetAttachmentUrl
`func (o *DataProcessingAgreement) UnsetAttachmentUrl()`

UnsetAttachmentUrl ensures that no value is present for AttachmentUrl, not even an explicit nil
### GetOrganization

`func (o *DataProcessingAgreement) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *DataProcessingAgreement) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *DataProcessingAgreement) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *DataProcessingAgreement) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



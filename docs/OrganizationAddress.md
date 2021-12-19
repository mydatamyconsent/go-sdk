# OrganizationAddress

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
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**OrganizationAddressType**](OrganizationAddressType.md) |  | [optional] 
**AddressLine1** | Pointer to **NullableString** |  | [optional] 
**AddressLine2** | Pointer to **NullableString** |  | [optional] 
**LandMark** | Pointer to **NullableString** |  | [optional] 
**CountryId** | Pointer to **string** |  | [optional] 
**StateId** | Pointer to **string** |  | [optional] 
**City** | Pointer to **NullableString** |  | [optional] 
**PostCode** | Pointer to **NullableString** |  | [optional] 
**ProofDocumentType** | Pointer to [**ProofDocumentType**](ProofDocumentType.md) |  | [optional] 
**ProofDocumentUrl** | Pointer to **NullableString** |  | [optional] 
**Country** | Pointer to [**Country**](Country.md) |  | [optional] 
**Organization** | Pointer to [**Organization**](Organization.md) |  | [optional] 
**CountryState** | Pointer to [**CountryState**](CountryState.md) |  | [optional] 
**IsPrimary** | Pointer to **bool** |  | [optional] 
**IsVerified** | Pointer to **bool** |  | [optional] 

## Methods

### NewOrganizationAddress

`func NewOrganizationAddress() *OrganizationAddress`

NewOrganizationAddress instantiates a new OrganizationAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationAddressWithDefaults

`func NewOrganizationAddressWithDefaults() *OrganizationAddress`

NewOrganizationAddressWithDefaults instantiates a new OrganizationAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *OrganizationAddress) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *OrganizationAddress) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *OrganizationAddress) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *OrganizationAddress) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedAtUtc

`func (o *OrganizationAddress) GetCreatedAtUtc() time.Time`

GetCreatedAtUtc returns the CreatedAtUtc field if non-nil, zero value otherwise.

### GetCreatedAtUtcOk

`func (o *OrganizationAddress) GetCreatedAtUtcOk() (*time.Time, bool)`

GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtUtc

`func (o *OrganizationAddress) SetCreatedAtUtc(v time.Time)`

SetCreatedAtUtc sets CreatedAtUtc field to given value.

### HasCreatedAtUtc

`func (o *OrganizationAddress) HasCreatedAtUtc() bool`

HasCreatedAtUtc returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *OrganizationAddress) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *OrganizationAddress) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *OrganizationAddress) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *OrganizationAddress) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *OrganizationAddress) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *OrganizationAddress) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetUpdatedAtUtc

`func (o *OrganizationAddress) GetUpdatedAtUtc() time.Time`

GetUpdatedAtUtc returns the UpdatedAtUtc field if non-nil, zero value otherwise.

### GetUpdatedAtUtcOk

`func (o *OrganizationAddress) GetUpdatedAtUtcOk() (*time.Time, bool)`

GetUpdatedAtUtcOk returns a tuple with the UpdatedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtUtc

`func (o *OrganizationAddress) SetUpdatedAtUtc(v time.Time)`

SetUpdatedAtUtc sets UpdatedAtUtc field to given value.

### HasUpdatedAtUtc

`func (o *OrganizationAddress) HasUpdatedAtUtc() bool`

HasUpdatedAtUtc returns a boolean if a field has been set.

### SetUpdatedAtUtcNil

`func (o *OrganizationAddress) SetUpdatedAtUtcNil(b bool)`

 SetUpdatedAtUtcNil sets the value for UpdatedAtUtc to be an explicit nil

### UnsetUpdatedAtUtc
`func (o *OrganizationAddress) UnsetUpdatedAtUtc()`

UnsetUpdatedAtUtc ensures that no value is present for UpdatedAtUtc, not even an explicit nil
### GetCreatedByUser

`func (o *OrganizationAddress) GetCreatedByUser() ApplicationUser`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *OrganizationAddress) GetCreatedByUserOk() (*ApplicationUser, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *OrganizationAddress) SetCreatedByUser(v ApplicationUser)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *OrganizationAddress) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### GetUpdatedByUser

`func (o *OrganizationAddress) GetUpdatedByUser() ApplicationUser`

GetUpdatedByUser returns the UpdatedByUser field if non-nil, zero value otherwise.

### GetUpdatedByUserOk

`func (o *OrganizationAddress) GetUpdatedByUserOk() (*ApplicationUser, bool)`

GetUpdatedByUserOk returns a tuple with the UpdatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUser

`func (o *OrganizationAddress) SetUpdatedByUser(v ApplicationUser)`

SetUpdatedByUser sets UpdatedByUser field to given value.

### HasUpdatedByUser

`func (o *OrganizationAddress) HasUpdatedByUser() bool`

HasUpdatedByUser returns a boolean if a field has been set.

### GetDeletedBy

`func (o *OrganizationAddress) GetDeletedBy() string`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *OrganizationAddress) GetDeletedByOk() (*string, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *OrganizationAddress) SetDeletedBy(v string)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *OrganizationAddress) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### SetDeletedByNil

`func (o *OrganizationAddress) SetDeletedByNil(b bool)`

 SetDeletedByNil sets the value for DeletedBy to be an explicit nil

### UnsetDeletedBy
`func (o *OrganizationAddress) UnsetDeletedBy()`

UnsetDeletedBy ensures that no value is present for DeletedBy, not even an explicit nil
### GetDeletedAtUtc

`func (o *OrganizationAddress) GetDeletedAtUtc() time.Time`

GetDeletedAtUtc returns the DeletedAtUtc field if non-nil, zero value otherwise.

### GetDeletedAtUtcOk

`func (o *OrganizationAddress) GetDeletedAtUtcOk() (*time.Time, bool)`

GetDeletedAtUtcOk returns a tuple with the DeletedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAtUtc

`func (o *OrganizationAddress) SetDeletedAtUtc(v time.Time)`

SetDeletedAtUtc sets DeletedAtUtc field to given value.

### HasDeletedAtUtc

`func (o *OrganizationAddress) HasDeletedAtUtc() bool`

HasDeletedAtUtc returns a boolean if a field has been set.

### SetDeletedAtUtcNil

`func (o *OrganizationAddress) SetDeletedAtUtcNil(b bool)`

 SetDeletedAtUtcNil sets the value for DeletedAtUtc to be an explicit nil

### UnsetDeletedAtUtc
`func (o *OrganizationAddress) UnsetDeletedAtUtc()`

UnsetDeletedAtUtc ensures that no value is present for DeletedAtUtc, not even an explicit nil
### GetDeletedByUser

`func (o *OrganizationAddress) GetDeletedByUser() ApplicationUser`

GetDeletedByUser returns the DeletedByUser field if non-nil, zero value otherwise.

### GetDeletedByUserOk

`func (o *OrganizationAddress) GetDeletedByUserOk() (*ApplicationUser, bool)`

GetDeletedByUserOk returns a tuple with the DeletedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedByUser

`func (o *OrganizationAddress) SetDeletedByUser(v ApplicationUser)`

SetDeletedByUser sets DeletedByUser field to given value.

### HasDeletedByUser

`func (o *OrganizationAddress) HasDeletedByUser() bool`

HasDeletedByUser returns a boolean if a field has been set.

### GetId

`func (o *OrganizationAddress) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationAddress) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationAddress) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationAddress) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *OrganizationAddress) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OrganizationAddress) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OrganizationAddress) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *OrganizationAddress) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOrganizationName

`func (o *OrganizationAddress) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *OrganizationAddress) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *OrganizationAddress) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *OrganizationAddress) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *OrganizationAddress) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *OrganizationAddress) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetType

`func (o *OrganizationAddress) GetType() OrganizationAddressType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrganizationAddress) GetTypeOk() (*OrganizationAddressType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrganizationAddress) SetType(v OrganizationAddressType)`

SetType sets Type field to given value.

### HasType

`func (o *OrganizationAddress) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAddressLine1

`func (o *OrganizationAddress) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *OrganizationAddress) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *OrganizationAddress) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *OrganizationAddress) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### SetAddressLine1Nil

`func (o *OrganizationAddress) SetAddressLine1Nil(b bool)`

 SetAddressLine1Nil sets the value for AddressLine1 to be an explicit nil

### UnsetAddressLine1
`func (o *OrganizationAddress) UnsetAddressLine1()`

UnsetAddressLine1 ensures that no value is present for AddressLine1, not even an explicit nil
### GetAddressLine2

`func (o *OrganizationAddress) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *OrganizationAddress) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *OrganizationAddress) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *OrganizationAddress) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### SetAddressLine2Nil

`func (o *OrganizationAddress) SetAddressLine2Nil(b bool)`

 SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil

### UnsetAddressLine2
`func (o *OrganizationAddress) UnsetAddressLine2()`

UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
### GetLandMark

`func (o *OrganizationAddress) GetLandMark() string`

GetLandMark returns the LandMark field if non-nil, zero value otherwise.

### GetLandMarkOk

`func (o *OrganizationAddress) GetLandMarkOk() (*string, bool)`

GetLandMarkOk returns a tuple with the LandMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandMark

`func (o *OrganizationAddress) SetLandMark(v string)`

SetLandMark sets LandMark field to given value.

### HasLandMark

`func (o *OrganizationAddress) HasLandMark() bool`

HasLandMark returns a boolean if a field has been set.

### SetLandMarkNil

`func (o *OrganizationAddress) SetLandMarkNil(b bool)`

 SetLandMarkNil sets the value for LandMark to be an explicit nil

### UnsetLandMark
`func (o *OrganizationAddress) UnsetLandMark()`

UnsetLandMark ensures that no value is present for LandMark, not even an explicit nil
### GetCountryId

`func (o *OrganizationAddress) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *OrganizationAddress) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *OrganizationAddress) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *OrganizationAddress) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### GetStateId

`func (o *OrganizationAddress) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *OrganizationAddress) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *OrganizationAddress) SetStateId(v string)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *OrganizationAddress) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### GetCity

`func (o *OrganizationAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *OrganizationAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *OrganizationAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *OrganizationAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *OrganizationAddress) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *OrganizationAddress) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetPostCode

`func (o *OrganizationAddress) GetPostCode() string`

GetPostCode returns the PostCode field if non-nil, zero value otherwise.

### GetPostCodeOk

`func (o *OrganizationAddress) GetPostCodeOk() (*string, bool)`

GetPostCodeOk returns a tuple with the PostCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCode

`func (o *OrganizationAddress) SetPostCode(v string)`

SetPostCode sets PostCode field to given value.

### HasPostCode

`func (o *OrganizationAddress) HasPostCode() bool`

HasPostCode returns a boolean if a field has been set.

### SetPostCodeNil

`func (o *OrganizationAddress) SetPostCodeNil(b bool)`

 SetPostCodeNil sets the value for PostCode to be an explicit nil

### UnsetPostCode
`func (o *OrganizationAddress) UnsetPostCode()`

UnsetPostCode ensures that no value is present for PostCode, not even an explicit nil
### GetProofDocumentType

`func (o *OrganizationAddress) GetProofDocumentType() ProofDocumentType`

GetProofDocumentType returns the ProofDocumentType field if non-nil, zero value otherwise.

### GetProofDocumentTypeOk

`func (o *OrganizationAddress) GetProofDocumentTypeOk() (*ProofDocumentType, bool)`

GetProofDocumentTypeOk returns a tuple with the ProofDocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProofDocumentType

`func (o *OrganizationAddress) SetProofDocumentType(v ProofDocumentType)`

SetProofDocumentType sets ProofDocumentType field to given value.

### HasProofDocumentType

`func (o *OrganizationAddress) HasProofDocumentType() bool`

HasProofDocumentType returns a boolean if a field has been set.

### GetProofDocumentUrl

`func (o *OrganizationAddress) GetProofDocumentUrl() string`

GetProofDocumentUrl returns the ProofDocumentUrl field if non-nil, zero value otherwise.

### GetProofDocumentUrlOk

`func (o *OrganizationAddress) GetProofDocumentUrlOk() (*string, bool)`

GetProofDocumentUrlOk returns a tuple with the ProofDocumentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProofDocumentUrl

`func (o *OrganizationAddress) SetProofDocumentUrl(v string)`

SetProofDocumentUrl sets ProofDocumentUrl field to given value.

### HasProofDocumentUrl

`func (o *OrganizationAddress) HasProofDocumentUrl() bool`

HasProofDocumentUrl returns a boolean if a field has been set.

### SetProofDocumentUrlNil

`func (o *OrganizationAddress) SetProofDocumentUrlNil(b bool)`

 SetProofDocumentUrlNil sets the value for ProofDocumentUrl to be an explicit nil

### UnsetProofDocumentUrl
`func (o *OrganizationAddress) UnsetProofDocumentUrl()`

UnsetProofDocumentUrl ensures that no value is present for ProofDocumentUrl, not even an explicit nil
### GetCountry

`func (o *OrganizationAddress) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *OrganizationAddress) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *OrganizationAddress) SetCountry(v Country)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *OrganizationAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetOrganization

`func (o *OrganizationAddress) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *OrganizationAddress) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *OrganizationAddress) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *OrganizationAddress) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetCountryState

`func (o *OrganizationAddress) GetCountryState() CountryState`

GetCountryState returns the CountryState field if non-nil, zero value otherwise.

### GetCountryStateOk

`func (o *OrganizationAddress) GetCountryStateOk() (*CountryState, bool)`

GetCountryStateOk returns a tuple with the CountryState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryState

`func (o *OrganizationAddress) SetCountryState(v CountryState)`

SetCountryState sets CountryState field to given value.

### HasCountryState

`func (o *OrganizationAddress) HasCountryState() bool`

HasCountryState returns a boolean if a field has been set.

### GetIsPrimary

`func (o *OrganizationAddress) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *OrganizationAddress) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *OrganizationAddress) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *OrganizationAddress) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetIsVerified

`func (o *OrganizationAddress) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *OrganizationAddress) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *OrganizationAddress) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *OrganizationAddress) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DigitalSignature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignedBy** | Pointer to **NullableString** |  | [optional] 
**CertIssuedBy** | Pointer to **NullableString** |  | [optional] 
**ValidFrom** | Pointer to **time.Time** |  | [optional] 
**ValidTill** | Pointer to **time.Time** |  | [optional] 
**Reason** | Pointer to **NullableString** |  | [optional] 
**Location** | Pointer to **NullableString** |  | [optional] 
**Sha1Digest** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDigitalSignature

`func NewDigitalSignature() *DigitalSignature`

NewDigitalSignature instantiates a new DigitalSignature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigitalSignatureWithDefaults

`func NewDigitalSignatureWithDefaults() *DigitalSignature`

NewDigitalSignatureWithDefaults instantiates a new DigitalSignature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignedBy

`func (o *DigitalSignature) GetSignedBy() string`

GetSignedBy returns the SignedBy field if non-nil, zero value otherwise.

### GetSignedByOk

`func (o *DigitalSignature) GetSignedByOk() (*string, bool)`

GetSignedByOk returns a tuple with the SignedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedBy

`func (o *DigitalSignature) SetSignedBy(v string)`

SetSignedBy sets SignedBy field to given value.

### HasSignedBy

`func (o *DigitalSignature) HasSignedBy() bool`

HasSignedBy returns a boolean if a field has been set.

### SetSignedByNil

`func (o *DigitalSignature) SetSignedByNil(b bool)`

 SetSignedByNil sets the value for SignedBy to be an explicit nil

### UnsetSignedBy
`func (o *DigitalSignature) UnsetSignedBy()`

UnsetSignedBy ensures that no value is present for SignedBy, not even an explicit nil
### GetCertIssuedBy

`func (o *DigitalSignature) GetCertIssuedBy() string`

GetCertIssuedBy returns the CertIssuedBy field if non-nil, zero value otherwise.

### GetCertIssuedByOk

`func (o *DigitalSignature) GetCertIssuedByOk() (*string, bool)`

GetCertIssuedByOk returns a tuple with the CertIssuedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertIssuedBy

`func (o *DigitalSignature) SetCertIssuedBy(v string)`

SetCertIssuedBy sets CertIssuedBy field to given value.

### HasCertIssuedBy

`func (o *DigitalSignature) HasCertIssuedBy() bool`

HasCertIssuedBy returns a boolean if a field has been set.

### SetCertIssuedByNil

`func (o *DigitalSignature) SetCertIssuedByNil(b bool)`

 SetCertIssuedByNil sets the value for CertIssuedBy to be an explicit nil

### UnsetCertIssuedBy
`func (o *DigitalSignature) UnsetCertIssuedBy()`

UnsetCertIssuedBy ensures that no value is present for CertIssuedBy, not even an explicit nil
### GetValidFrom

`func (o *DigitalSignature) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *DigitalSignature) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *DigitalSignature) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *DigitalSignature) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidTill

`func (o *DigitalSignature) GetValidTill() time.Time`

GetValidTill returns the ValidTill field if non-nil, zero value otherwise.

### GetValidTillOk

`func (o *DigitalSignature) GetValidTillOk() (*time.Time, bool)`

GetValidTillOk returns a tuple with the ValidTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTill

`func (o *DigitalSignature) SetValidTill(v time.Time)`

SetValidTill sets ValidTill field to given value.

### HasValidTill

`func (o *DigitalSignature) HasValidTill() bool`

HasValidTill returns a boolean if a field has been set.

### GetReason

`func (o *DigitalSignature) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DigitalSignature) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DigitalSignature) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *DigitalSignature) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *DigitalSignature) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *DigitalSignature) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetLocation

`func (o *DigitalSignature) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DigitalSignature) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DigitalSignature) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DigitalSignature) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *DigitalSignature) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *DigitalSignature) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetSha1Digest

`func (o *DigitalSignature) GetSha1Digest() string`

GetSha1Digest returns the Sha1Digest field if non-nil, zero value otherwise.

### GetSha1DigestOk

`func (o *DigitalSignature) GetSha1DigestOk() (*string, bool)`

GetSha1DigestOk returns a tuple with the Sha1Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1Digest

`func (o *DigitalSignature) SetSha1Digest(v string)`

SetSha1Digest sets Sha1Digest field to given value.

### HasSha1Digest

`func (o *DigitalSignature) HasSha1Digest() bool`

HasSha1Digest returns a boolean if a field has been set.

### SetSha1DigestNil

`func (o *DigitalSignature) SetSha1DigestNil(b bool)`

 SetSha1DigestNil sets the value for Sha1Digest to be an explicit nil

### UnsetSha1Digest
`func (o *DigitalSignature) UnsetSha1Digest()`

UnsetSha1Digest ensures that no value is present for Sha1Digest, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



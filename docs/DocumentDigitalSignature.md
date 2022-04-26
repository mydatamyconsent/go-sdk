# DocumentDigitalSignature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name. | 
**IssuedBy** | **string** | Signature issued by. | 
**IssuerName** | **string** | Signature issuer name. | 
**ValidFromUtc** | **time.Time** | Signature valid from datatime in UTC timezone. | 
**ValidToUtc** | **time.Time** | Signature valid to datatime in UTC timezone. | 

## Methods

### NewDocumentDigitalSignature

`func NewDocumentDigitalSignature(name string, issuedBy string, issuerName string, validFromUtc time.Time, validToUtc time.Time, ) *DocumentDigitalSignature`

NewDocumentDigitalSignature instantiates a new DocumentDigitalSignature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentDigitalSignatureWithDefaults

`func NewDocumentDigitalSignatureWithDefaults() *DocumentDigitalSignature`

NewDocumentDigitalSignatureWithDefaults instantiates a new DocumentDigitalSignature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DocumentDigitalSignature) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DocumentDigitalSignature) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DocumentDigitalSignature) SetName(v string)`

SetName sets Name field to given value.


### GetIssuedBy

`func (o *DocumentDigitalSignature) GetIssuedBy() string`

GetIssuedBy returns the IssuedBy field if non-nil, zero value otherwise.

### GetIssuedByOk

`func (o *DocumentDigitalSignature) GetIssuedByOk() (*string, bool)`

GetIssuedByOk returns a tuple with the IssuedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedBy

`func (o *DocumentDigitalSignature) SetIssuedBy(v string)`

SetIssuedBy sets IssuedBy field to given value.


### GetIssuerName

`func (o *DocumentDigitalSignature) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *DocumentDigitalSignature) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *DocumentDigitalSignature) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.


### GetValidFromUtc

`func (o *DocumentDigitalSignature) GetValidFromUtc() time.Time`

GetValidFromUtc returns the ValidFromUtc field if non-nil, zero value otherwise.

### GetValidFromUtcOk

`func (o *DocumentDigitalSignature) GetValidFromUtcOk() (*time.Time, bool)`

GetValidFromUtcOk returns a tuple with the ValidFromUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFromUtc

`func (o *DocumentDigitalSignature) SetValidFromUtc(v time.Time)`

SetValidFromUtc sets ValidFromUtc field to given value.


### GetValidToUtc

`func (o *DocumentDigitalSignature) GetValidToUtc() time.Time`

GetValidToUtc returns the ValidToUtc field if non-nil, zero value otherwise.

### GetValidToUtcOk

`func (o *DocumentDigitalSignature) GetValidToUtcOk() (*time.Time, bool)`

GetValidToUtcOk returns a tuple with the ValidToUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidToUtc

`func (o *DocumentDigitalSignature) SetValidToUtc(v time.Time)`

SetValidToUtc sets ValidToUtc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



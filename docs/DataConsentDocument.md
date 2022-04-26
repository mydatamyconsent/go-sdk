# DataConsentDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Document id. | 
**Name** | **string** | Document name. | 
**Category** | **string** | Document category. | 
**Identifier** | **string** | Document identifier. | 
**FieldTitle** | **string** | Document field title. | 
**FieldSlug** | **string** | Document field slug. | 
**IssuedAtUtc** | **time.Time** | Document issued at datetime in UTC timezone. | 
**ExpiresAtUtc** | Pointer to **NullableTime** | Document expires at datetime in UTC timezone. | [optional] 
**Issuer** | [**DataConsentDocumentIssuer**](DataConsentDocumentIssuer.md) |  | 
**DigitalSignatures** | [**[]DocumentDigitalSignature**](DocumentDigitalSignature.md) | Digital signatures. | 

## Methods

### NewDataConsentDocument

`func NewDataConsentDocument(id string, name string, category string, identifier string, fieldTitle string, fieldSlug string, issuedAtUtc time.Time, issuer DataConsentDocumentIssuer, digitalSignatures []DocumentDigitalSignature, ) *DataConsentDocument`

NewDataConsentDocument instantiates a new DataConsentDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataConsentDocumentWithDefaults

`func NewDataConsentDocumentWithDefaults() *DataConsentDocument`

NewDataConsentDocumentWithDefaults instantiates a new DataConsentDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataConsentDocument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataConsentDocument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataConsentDocument) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DataConsentDocument) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataConsentDocument) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataConsentDocument) SetName(v string)`

SetName sets Name field to given value.


### GetCategory

`func (o *DataConsentDocument) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *DataConsentDocument) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *DataConsentDocument) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetIdentifier

`func (o *DataConsentDocument) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *DataConsentDocument) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *DataConsentDocument) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetFieldTitle

`func (o *DataConsentDocument) GetFieldTitle() string`

GetFieldTitle returns the FieldTitle field if non-nil, zero value otherwise.

### GetFieldTitleOk

`func (o *DataConsentDocument) GetFieldTitleOk() (*string, bool)`

GetFieldTitleOk returns a tuple with the FieldTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldTitle

`func (o *DataConsentDocument) SetFieldTitle(v string)`

SetFieldTitle sets FieldTitle field to given value.


### GetFieldSlug

`func (o *DataConsentDocument) GetFieldSlug() string`

GetFieldSlug returns the FieldSlug field if non-nil, zero value otherwise.

### GetFieldSlugOk

`func (o *DataConsentDocument) GetFieldSlugOk() (*string, bool)`

GetFieldSlugOk returns a tuple with the FieldSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldSlug

`func (o *DataConsentDocument) SetFieldSlug(v string)`

SetFieldSlug sets FieldSlug field to given value.


### GetIssuedAtUtc

`func (o *DataConsentDocument) GetIssuedAtUtc() time.Time`

GetIssuedAtUtc returns the IssuedAtUtc field if non-nil, zero value otherwise.

### GetIssuedAtUtcOk

`func (o *DataConsentDocument) GetIssuedAtUtcOk() (*time.Time, bool)`

GetIssuedAtUtcOk returns a tuple with the IssuedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAtUtc

`func (o *DataConsentDocument) SetIssuedAtUtc(v time.Time)`

SetIssuedAtUtc sets IssuedAtUtc field to given value.


### GetExpiresAtUtc

`func (o *DataConsentDocument) GetExpiresAtUtc() time.Time`

GetExpiresAtUtc returns the ExpiresAtUtc field if non-nil, zero value otherwise.

### GetExpiresAtUtcOk

`func (o *DataConsentDocument) GetExpiresAtUtcOk() (*time.Time, bool)`

GetExpiresAtUtcOk returns a tuple with the ExpiresAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAtUtc

`func (o *DataConsentDocument) SetExpiresAtUtc(v time.Time)`

SetExpiresAtUtc sets ExpiresAtUtc field to given value.

### HasExpiresAtUtc

`func (o *DataConsentDocument) HasExpiresAtUtc() bool`

HasExpiresAtUtc returns a boolean if a field has been set.

### SetExpiresAtUtcNil

`func (o *DataConsentDocument) SetExpiresAtUtcNil(b bool)`

 SetExpiresAtUtcNil sets the value for ExpiresAtUtc to be an explicit nil

### UnsetExpiresAtUtc
`func (o *DataConsentDocument) UnsetExpiresAtUtc()`

UnsetExpiresAtUtc ensures that no value is present for ExpiresAtUtc, not even an explicit nil
### GetIssuer

`func (o *DataConsentDocument) GetIssuer() DataConsentDocumentIssuer`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *DataConsentDocument) GetIssuerOk() (*DataConsentDocumentIssuer, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *DataConsentDocument) SetIssuer(v DataConsentDocumentIssuer)`

SetIssuer sets Issuer field to given value.


### GetDigitalSignatures

`func (o *DataConsentDocument) GetDigitalSignatures() []DocumentDigitalSignature`

GetDigitalSignatures returns the DigitalSignatures field if non-nil, zero value otherwise.

### GetDigitalSignaturesOk

`func (o *DataConsentDocument) GetDigitalSignaturesOk() (*[]DocumentDigitalSignature, bool)`

GetDigitalSignaturesOk returns a tuple with the DigitalSignatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalSignatures

`func (o *DataConsentDocument) SetDigitalSignatures(v []DocumentDigitalSignature)`

SetDigitalSignatures sets DigitalSignatures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



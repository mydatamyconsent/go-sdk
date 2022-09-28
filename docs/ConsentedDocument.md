# ConsentedDocument

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
**ExpiresAtUtc** | Pointer to **time.Time** | Document expires at datetime in UTC timezone. | [optional] 
**Issuer** | [**ConsentDocumentIssuer**](ConsentDocumentIssuer.md) |  | 
**DigitalSignatures** | [**[]DocumentDigitalSignature**](DocumentDigitalSignature.md) | Digital signatures. | 

## Methods

### NewConsentedDocument

`func NewConsentedDocument(id string, name string, category string, identifier string, fieldTitle string, fieldSlug string, issuedAtUtc time.Time, issuer ConsentDocumentIssuer, digitalSignatures []DocumentDigitalSignature, ) *ConsentedDocument`

NewConsentedDocument instantiates a new ConsentedDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsentedDocumentWithDefaults

`func NewConsentedDocumentWithDefaults() *ConsentedDocument`

NewConsentedDocumentWithDefaults instantiates a new ConsentedDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConsentedDocument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsentedDocument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsentedDocument) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ConsentedDocument) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsentedDocument) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsentedDocument) SetName(v string)`

SetName sets Name field to given value.


### GetCategory

`func (o *ConsentedDocument) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ConsentedDocument) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ConsentedDocument) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetIdentifier

`func (o *ConsentedDocument) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ConsentedDocument) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ConsentedDocument) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetFieldTitle

`func (o *ConsentedDocument) GetFieldTitle() string`

GetFieldTitle returns the FieldTitle field if non-nil, zero value otherwise.

### GetFieldTitleOk

`func (o *ConsentedDocument) GetFieldTitleOk() (*string, bool)`

GetFieldTitleOk returns a tuple with the FieldTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldTitle

`func (o *ConsentedDocument) SetFieldTitle(v string)`

SetFieldTitle sets FieldTitle field to given value.


### GetFieldSlug

`func (o *ConsentedDocument) GetFieldSlug() string`

GetFieldSlug returns the FieldSlug field if non-nil, zero value otherwise.

### GetFieldSlugOk

`func (o *ConsentedDocument) GetFieldSlugOk() (*string, bool)`

GetFieldSlugOk returns a tuple with the FieldSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldSlug

`func (o *ConsentedDocument) SetFieldSlug(v string)`

SetFieldSlug sets FieldSlug field to given value.


### GetIssuedAtUtc

`func (o *ConsentedDocument) GetIssuedAtUtc() time.Time`

GetIssuedAtUtc returns the IssuedAtUtc field if non-nil, zero value otherwise.

### GetIssuedAtUtcOk

`func (o *ConsentedDocument) GetIssuedAtUtcOk() (*time.Time, bool)`

GetIssuedAtUtcOk returns a tuple with the IssuedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAtUtc

`func (o *ConsentedDocument) SetIssuedAtUtc(v time.Time)`

SetIssuedAtUtc sets IssuedAtUtc field to given value.


### GetExpiresAtUtc

`func (o *ConsentedDocument) GetExpiresAtUtc() time.Time`

GetExpiresAtUtc returns the ExpiresAtUtc field if non-nil, zero value otherwise.

### GetExpiresAtUtcOk

`func (o *ConsentedDocument) GetExpiresAtUtcOk() (*time.Time, bool)`

GetExpiresAtUtcOk returns a tuple with the ExpiresAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAtUtc

`func (o *ConsentedDocument) SetExpiresAtUtc(v time.Time)`

SetExpiresAtUtc sets ExpiresAtUtc field to given value.

### HasExpiresAtUtc

`func (o *ConsentedDocument) HasExpiresAtUtc() bool`

HasExpiresAtUtc returns a boolean if a field has been set.

### GetIssuer

`func (o *ConsentedDocument) GetIssuer() ConsentDocumentIssuer`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *ConsentedDocument) GetIssuerOk() (*ConsentDocumentIssuer, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *ConsentedDocument) SetIssuer(v ConsentDocumentIssuer)`

SetIssuer sets Issuer field to given value.


### GetDigitalSignatures

`func (o *ConsentedDocument) GetDigitalSignatures() []DocumentDigitalSignature`

GetDigitalSignatures returns the DigitalSignatures field if non-nil, zero value otherwise.

### GetDigitalSignaturesOk

`func (o *ConsentedDocument) GetDigitalSignaturesOk() (*[]DocumentDigitalSignature, bool)`

GetDigitalSignaturesOk returns a tuple with the DigitalSignatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalSignatures

`func (o *ConsentedDocument) SetDigitalSignatures(v []DocumentDigitalSignature)`

SetDigitalSignatures sets DigitalSignatures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



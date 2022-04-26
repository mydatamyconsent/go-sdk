# IssuedDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Document Id. | 
**Identifier** | **string** | Document Identifier. | 
**DocumentType** | **string** | Document type name. | 
**IssuedTo** | **string** | User name. | 
**IssuedAtUtc** | **time.Time** | Issued datetime in UTC timezone. | 
**ExpiresAtUtc** | Pointer to **NullableTime** | Expires datetime in UTC timezone. | [optional] 
**AcceptedAtUtc** | Pointer to **NullableTime** | Accepted datetime in UTC timezone. | [optional] 

## Methods

### NewIssuedDocument

`func NewIssuedDocument(id string, identifier string, documentType string, issuedTo string, issuedAtUtc time.Time, ) *IssuedDocument`

NewIssuedDocument instantiates a new IssuedDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuedDocumentWithDefaults

`func NewIssuedDocumentWithDefaults() *IssuedDocument`

NewIssuedDocumentWithDefaults instantiates a new IssuedDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IssuedDocument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssuedDocument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssuedDocument) SetId(v string)`

SetId sets Id field to given value.


### GetIdentifier

`func (o *IssuedDocument) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *IssuedDocument) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *IssuedDocument) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetDocumentType

`func (o *IssuedDocument) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *IssuedDocument) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *IssuedDocument) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.


### GetIssuedTo

`func (o *IssuedDocument) GetIssuedTo() string`

GetIssuedTo returns the IssuedTo field if non-nil, zero value otherwise.

### GetIssuedToOk

`func (o *IssuedDocument) GetIssuedToOk() (*string, bool)`

GetIssuedToOk returns a tuple with the IssuedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedTo

`func (o *IssuedDocument) SetIssuedTo(v string)`

SetIssuedTo sets IssuedTo field to given value.


### GetIssuedAtUtc

`func (o *IssuedDocument) GetIssuedAtUtc() time.Time`

GetIssuedAtUtc returns the IssuedAtUtc field if non-nil, zero value otherwise.

### GetIssuedAtUtcOk

`func (o *IssuedDocument) GetIssuedAtUtcOk() (*time.Time, bool)`

GetIssuedAtUtcOk returns a tuple with the IssuedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAtUtc

`func (o *IssuedDocument) SetIssuedAtUtc(v time.Time)`

SetIssuedAtUtc sets IssuedAtUtc field to given value.


### GetExpiresAtUtc

`func (o *IssuedDocument) GetExpiresAtUtc() time.Time`

GetExpiresAtUtc returns the ExpiresAtUtc field if non-nil, zero value otherwise.

### GetExpiresAtUtcOk

`func (o *IssuedDocument) GetExpiresAtUtcOk() (*time.Time, bool)`

GetExpiresAtUtcOk returns a tuple with the ExpiresAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAtUtc

`func (o *IssuedDocument) SetExpiresAtUtc(v time.Time)`

SetExpiresAtUtc sets ExpiresAtUtc field to given value.

### HasExpiresAtUtc

`func (o *IssuedDocument) HasExpiresAtUtc() bool`

HasExpiresAtUtc returns a boolean if a field has been set.

### SetExpiresAtUtcNil

`func (o *IssuedDocument) SetExpiresAtUtcNil(b bool)`

 SetExpiresAtUtcNil sets the value for ExpiresAtUtc to be an explicit nil

### UnsetExpiresAtUtc
`func (o *IssuedDocument) UnsetExpiresAtUtc()`

UnsetExpiresAtUtc ensures that no value is present for ExpiresAtUtc, not even an explicit nil
### GetAcceptedAtUtc

`func (o *IssuedDocument) GetAcceptedAtUtc() time.Time`

GetAcceptedAtUtc returns the AcceptedAtUtc field if non-nil, zero value otherwise.

### GetAcceptedAtUtcOk

`func (o *IssuedDocument) GetAcceptedAtUtcOk() (*time.Time, bool)`

GetAcceptedAtUtcOk returns a tuple with the AcceptedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedAtUtc

`func (o *IssuedDocument) SetAcceptedAtUtc(v time.Time)`

SetAcceptedAtUtc sets AcceptedAtUtc field to given value.

### HasAcceptedAtUtc

`func (o *IssuedDocument) HasAcceptedAtUtc() bool`

HasAcceptedAtUtc returns a boolean if a field has been set.

### SetAcceptedAtUtcNil

`func (o *IssuedDocument) SetAcceptedAtUtcNil(b bool)`

 SetAcceptedAtUtcNil sets the value for AcceptedAtUtc to be an explicit nil

### UnsetAcceptedAtUtc
`func (o *IssuedDocument) UnsetAcceptedAtUtc()`

UnsetAcceptedAtUtc ensures that no value is present for AcceptedAtUtc, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



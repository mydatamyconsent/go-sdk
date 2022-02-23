# IssuedDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Document Identifier. | 
**Identifier** | **string** | Document Identifier. eg: GJ05FG67866586. | 
**DocumentType** | **string** | Document type name. eg: Driving License. | 
**IssuedTo** | **string** |  | 
**IssuedAtUtc** | **time.Time** |  | 

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



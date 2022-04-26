# IssuedDocumentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Receiver** | [**DocumentReceiver**](DocumentReceiver.md) |  | 
**Metadata** | Pointer to **map[string]string** | Metadata. | [optional] 
**DigitalSignatures** | [**[]DocumentDigitalSignature**](DocumentDigitalSignature.md) | Digital signatures. | 
**Id** | **string** | Document Id. | 
**Identifier** | **string** | Document Identifier. | 
**DocumentType** | **string** | Document type name. | 
**IssuedTo** | **string** | User name. | 
**IssuedAtUtc** | **time.Time** | Issued datetime in UTC timezone. | 
**ExpiresAtUtc** | Pointer to **NullableTime** | Expires datetime in UTC timezone. | [optional] 
**AcceptedAtUtc** | Pointer to **NullableTime** | Accepted datetime in UTC timezone. | [optional] 

## Methods

### NewIssuedDocumentDetails

`func NewIssuedDocumentDetails(receiver DocumentReceiver, digitalSignatures []DocumentDigitalSignature, id string, identifier string, documentType string, issuedTo string, issuedAtUtc time.Time, ) *IssuedDocumentDetails`

NewIssuedDocumentDetails instantiates a new IssuedDocumentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuedDocumentDetailsWithDefaults

`func NewIssuedDocumentDetailsWithDefaults() *IssuedDocumentDetails`

NewIssuedDocumentDetailsWithDefaults instantiates a new IssuedDocumentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceiver

`func (o *IssuedDocumentDetails) GetReceiver() DocumentReceiver`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *IssuedDocumentDetails) GetReceiverOk() (*DocumentReceiver, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *IssuedDocumentDetails) SetReceiver(v DocumentReceiver)`

SetReceiver sets Receiver field to given value.


### GetMetadata

`func (o *IssuedDocumentDetails) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IssuedDocumentDetails) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IssuedDocumentDetails) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IssuedDocumentDetails) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *IssuedDocumentDetails) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *IssuedDocumentDetails) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetDigitalSignatures

`func (o *IssuedDocumentDetails) GetDigitalSignatures() []DocumentDigitalSignature`

GetDigitalSignatures returns the DigitalSignatures field if non-nil, zero value otherwise.

### GetDigitalSignaturesOk

`func (o *IssuedDocumentDetails) GetDigitalSignaturesOk() (*[]DocumentDigitalSignature, bool)`

GetDigitalSignaturesOk returns a tuple with the DigitalSignatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalSignatures

`func (o *IssuedDocumentDetails) SetDigitalSignatures(v []DocumentDigitalSignature)`

SetDigitalSignatures sets DigitalSignatures field to given value.


### GetId

`func (o *IssuedDocumentDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssuedDocumentDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssuedDocumentDetails) SetId(v string)`

SetId sets Id field to given value.


### GetIdentifier

`func (o *IssuedDocumentDetails) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *IssuedDocumentDetails) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *IssuedDocumentDetails) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetDocumentType

`func (o *IssuedDocumentDetails) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *IssuedDocumentDetails) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *IssuedDocumentDetails) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.


### GetIssuedTo

`func (o *IssuedDocumentDetails) GetIssuedTo() string`

GetIssuedTo returns the IssuedTo field if non-nil, zero value otherwise.

### GetIssuedToOk

`func (o *IssuedDocumentDetails) GetIssuedToOk() (*string, bool)`

GetIssuedToOk returns a tuple with the IssuedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedTo

`func (o *IssuedDocumentDetails) SetIssuedTo(v string)`

SetIssuedTo sets IssuedTo field to given value.


### GetIssuedAtUtc

`func (o *IssuedDocumentDetails) GetIssuedAtUtc() time.Time`

GetIssuedAtUtc returns the IssuedAtUtc field if non-nil, zero value otherwise.

### GetIssuedAtUtcOk

`func (o *IssuedDocumentDetails) GetIssuedAtUtcOk() (*time.Time, bool)`

GetIssuedAtUtcOk returns a tuple with the IssuedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAtUtc

`func (o *IssuedDocumentDetails) SetIssuedAtUtc(v time.Time)`

SetIssuedAtUtc sets IssuedAtUtc field to given value.


### GetExpiresAtUtc

`func (o *IssuedDocumentDetails) GetExpiresAtUtc() time.Time`

GetExpiresAtUtc returns the ExpiresAtUtc field if non-nil, zero value otherwise.

### GetExpiresAtUtcOk

`func (o *IssuedDocumentDetails) GetExpiresAtUtcOk() (*time.Time, bool)`

GetExpiresAtUtcOk returns a tuple with the ExpiresAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAtUtc

`func (o *IssuedDocumentDetails) SetExpiresAtUtc(v time.Time)`

SetExpiresAtUtc sets ExpiresAtUtc field to given value.

### HasExpiresAtUtc

`func (o *IssuedDocumentDetails) HasExpiresAtUtc() bool`

HasExpiresAtUtc returns a boolean if a field has been set.

### SetExpiresAtUtcNil

`func (o *IssuedDocumentDetails) SetExpiresAtUtcNil(b bool)`

 SetExpiresAtUtcNil sets the value for ExpiresAtUtc to be an explicit nil

### UnsetExpiresAtUtc
`func (o *IssuedDocumentDetails) UnsetExpiresAtUtc()`

UnsetExpiresAtUtc ensures that no value is present for ExpiresAtUtc, not even an explicit nil
### GetAcceptedAtUtc

`func (o *IssuedDocumentDetails) GetAcceptedAtUtc() time.Time`

GetAcceptedAtUtc returns the AcceptedAtUtc field if non-nil, zero value otherwise.

### GetAcceptedAtUtcOk

`func (o *IssuedDocumentDetails) GetAcceptedAtUtcOk() (*time.Time, bool)`

GetAcceptedAtUtcOk returns a tuple with the AcceptedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedAtUtc

`func (o *IssuedDocumentDetails) SetAcceptedAtUtc(v time.Time)`

SetAcceptedAtUtc sets AcceptedAtUtc field to given value.

### HasAcceptedAtUtc

`func (o *IssuedDocumentDetails) HasAcceptedAtUtc() bool`

HasAcceptedAtUtc returns a boolean if a field has been set.

### SetAcceptedAtUtcNil

`func (o *IssuedDocumentDetails) SetAcceptedAtUtcNil(b bool)`

 SetAcceptedAtUtcNil sets the value for AcceptedAtUtc to be an explicit nil

### UnsetAcceptedAtUtc
`func (o *IssuedDocumentDetails) UnsetAcceptedAtUtc()`

UnsetAcceptedAtUtc ensures that no value is present for AcceptedAtUtc, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



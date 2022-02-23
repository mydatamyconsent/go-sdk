# DocumentIssueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentTypeId** | **string** |  | 
**DocumentIdentifier** | **string** |  | 
**Description** | **string** |  | 
**Receiver** | [**DocumentReceiver**](DocumentReceiver.md) |  | 
**ExpiresAtUtc** | Pointer to **NullableTime** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewDocumentIssueRequest

`func NewDocumentIssueRequest(documentTypeId string, documentIdentifier string, description string, receiver DocumentReceiver, ) *DocumentIssueRequest`

NewDocumentIssueRequest instantiates a new DocumentIssueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentIssueRequestWithDefaults

`func NewDocumentIssueRequestWithDefaults() *DocumentIssueRequest`

NewDocumentIssueRequestWithDefaults instantiates a new DocumentIssueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentTypeId

`func (o *DocumentIssueRequest) GetDocumentTypeId() string`

GetDocumentTypeId returns the DocumentTypeId field if non-nil, zero value otherwise.

### GetDocumentTypeIdOk

`func (o *DocumentIssueRequest) GetDocumentTypeIdOk() (*string, bool)`

GetDocumentTypeIdOk returns a tuple with the DocumentTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTypeId

`func (o *DocumentIssueRequest) SetDocumentTypeId(v string)`

SetDocumentTypeId sets DocumentTypeId field to given value.


### GetDocumentIdentifier

`func (o *DocumentIssueRequest) GetDocumentIdentifier() string`

GetDocumentIdentifier returns the DocumentIdentifier field if non-nil, zero value otherwise.

### GetDocumentIdentifierOk

`func (o *DocumentIssueRequest) GetDocumentIdentifierOk() (*string, bool)`

GetDocumentIdentifierOk returns a tuple with the DocumentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIdentifier

`func (o *DocumentIssueRequest) SetDocumentIdentifier(v string)`

SetDocumentIdentifier sets DocumentIdentifier field to given value.


### GetDescription

`func (o *DocumentIssueRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DocumentIssueRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DocumentIssueRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetReceiver

`func (o *DocumentIssueRequest) GetReceiver() DocumentReceiver`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *DocumentIssueRequest) GetReceiverOk() (*DocumentReceiver, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *DocumentIssueRequest) SetReceiver(v DocumentReceiver)`

SetReceiver sets Receiver field to given value.


### GetExpiresAtUtc

`func (o *DocumentIssueRequest) GetExpiresAtUtc() time.Time`

GetExpiresAtUtc returns the ExpiresAtUtc field if non-nil, zero value otherwise.

### GetExpiresAtUtcOk

`func (o *DocumentIssueRequest) GetExpiresAtUtcOk() (*time.Time, bool)`

GetExpiresAtUtcOk returns a tuple with the ExpiresAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAtUtc

`func (o *DocumentIssueRequest) SetExpiresAtUtc(v time.Time)`

SetExpiresAtUtc sets ExpiresAtUtc field to given value.

### HasExpiresAtUtc

`func (o *DocumentIssueRequest) HasExpiresAtUtc() bool`

HasExpiresAtUtc returns a boolean if a field has been set.

### SetExpiresAtUtcNil

`func (o *DocumentIssueRequest) SetExpiresAtUtcNil(b bool)`

 SetExpiresAtUtcNil sets the value for ExpiresAtUtc to be an explicit nil

### UnsetExpiresAtUtc
`func (o *DocumentIssueRequest) UnsetExpiresAtUtc()`

UnsetExpiresAtUtc ensures that no value is present for ExpiresAtUtc, not even an explicit nil
### GetMetadata

`func (o *DocumentIssueRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DocumentIssueRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DocumentIssueRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DocumentIssueRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *DocumentIssueRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *DocumentIssueRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



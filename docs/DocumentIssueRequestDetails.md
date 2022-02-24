# DocumentIssueRequestDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Request Id. | 
**DocumentTypeId** | **string** |  | 
**DocumentTypeName** | **string** |  | 
**DocumentIdentifier** | **string** |  | 
**Status** | Pointer to [**DocumentIssueRequestStatus**](DocumentIssueRequestStatus.md) |  | [optional] 
**Description** | **string** |  | 
**Receiver** | **interface{}** |  | 
**ExpiresAtUtc** | Pointer to **NullableTime** |  | [optional] 
**Metadata** | Pointer to **interface{}** |  | [optional] 
**CreatedAtUtc** | **time.Time** |  | 

## Methods

### NewDocumentIssueRequestDetails

`func NewDocumentIssueRequestDetails(id string, documentTypeId string, documentTypeName string, documentIdentifier string, description string, receiver interface{}, createdAtUtc time.Time, ) *DocumentIssueRequestDetails`

NewDocumentIssueRequestDetails instantiates a new DocumentIssueRequestDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentIssueRequestDetailsWithDefaults

`func NewDocumentIssueRequestDetailsWithDefaults() *DocumentIssueRequestDetails`

NewDocumentIssueRequestDetailsWithDefaults instantiates a new DocumentIssueRequestDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DocumentIssueRequestDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentIssueRequestDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentIssueRequestDetails) SetId(v string)`

SetId sets Id field to given value.


### GetDocumentTypeId

`func (o *DocumentIssueRequestDetails) GetDocumentTypeId() string`

GetDocumentTypeId returns the DocumentTypeId field if non-nil, zero value otherwise.

### GetDocumentTypeIdOk

`func (o *DocumentIssueRequestDetails) GetDocumentTypeIdOk() (*string, bool)`

GetDocumentTypeIdOk returns a tuple with the DocumentTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTypeId

`func (o *DocumentIssueRequestDetails) SetDocumentTypeId(v string)`

SetDocumentTypeId sets DocumentTypeId field to given value.


### GetDocumentTypeName

`func (o *DocumentIssueRequestDetails) GetDocumentTypeName() string`

GetDocumentTypeName returns the DocumentTypeName field if non-nil, zero value otherwise.

### GetDocumentTypeNameOk

`func (o *DocumentIssueRequestDetails) GetDocumentTypeNameOk() (*string, bool)`

GetDocumentTypeNameOk returns a tuple with the DocumentTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTypeName

`func (o *DocumentIssueRequestDetails) SetDocumentTypeName(v string)`

SetDocumentTypeName sets DocumentTypeName field to given value.


### GetDocumentIdentifier

`func (o *DocumentIssueRequestDetails) GetDocumentIdentifier() string`

GetDocumentIdentifier returns the DocumentIdentifier field if non-nil, zero value otherwise.

### GetDocumentIdentifierOk

`func (o *DocumentIssueRequestDetails) GetDocumentIdentifierOk() (*string, bool)`

GetDocumentIdentifierOk returns a tuple with the DocumentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIdentifier

`func (o *DocumentIssueRequestDetails) SetDocumentIdentifier(v string)`

SetDocumentIdentifier sets DocumentIdentifier field to given value.


### GetStatus

`func (o *DocumentIssueRequestDetails) GetStatus() DocumentIssueRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DocumentIssueRequestDetails) GetStatusOk() (*DocumentIssueRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DocumentIssueRequestDetails) SetStatus(v DocumentIssueRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DocumentIssueRequestDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDescription

`func (o *DocumentIssueRequestDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DocumentIssueRequestDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DocumentIssueRequestDetails) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetReceiver

`func (o *DocumentIssueRequestDetails) GetReceiver() interface{}`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *DocumentIssueRequestDetails) GetReceiverOk() (*interface{}, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *DocumentIssueRequestDetails) SetReceiver(v interface{})`

SetReceiver sets Receiver field to given value.


### SetReceiverNil

`func (o *DocumentIssueRequestDetails) SetReceiverNil(b bool)`

 SetReceiverNil sets the value for Receiver to be an explicit nil

### UnsetReceiver
`func (o *DocumentIssueRequestDetails) UnsetReceiver()`

UnsetReceiver ensures that no value is present for Receiver, not even an explicit nil
### GetExpiresAtUtc

`func (o *DocumentIssueRequestDetails) GetExpiresAtUtc() time.Time`

GetExpiresAtUtc returns the ExpiresAtUtc field if non-nil, zero value otherwise.

### GetExpiresAtUtcOk

`func (o *DocumentIssueRequestDetails) GetExpiresAtUtcOk() (*time.Time, bool)`

GetExpiresAtUtcOk returns a tuple with the ExpiresAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAtUtc

`func (o *DocumentIssueRequestDetails) SetExpiresAtUtc(v time.Time)`

SetExpiresAtUtc sets ExpiresAtUtc field to given value.

### HasExpiresAtUtc

`func (o *DocumentIssueRequestDetails) HasExpiresAtUtc() bool`

HasExpiresAtUtc returns a boolean if a field has been set.

### SetExpiresAtUtcNil

`func (o *DocumentIssueRequestDetails) SetExpiresAtUtcNil(b bool)`

 SetExpiresAtUtcNil sets the value for ExpiresAtUtc to be an explicit nil

### UnsetExpiresAtUtc
`func (o *DocumentIssueRequestDetails) UnsetExpiresAtUtc()`

UnsetExpiresAtUtc ensures that no value is present for ExpiresAtUtc, not even an explicit nil
### GetMetadata

`func (o *DocumentIssueRequestDetails) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DocumentIssueRequestDetails) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DocumentIssueRequestDetails) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DocumentIssueRequestDetails) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *DocumentIssueRequestDetails) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *DocumentIssueRequestDetails) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCreatedAtUtc

`func (o *DocumentIssueRequestDetails) GetCreatedAtUtc() time.Time`

GetCreatedAtUtc returns the CreatedAtUtc field if non-nil, zero value otherwise.

### GetCreatedAtUtcOk

`func (o *DocumentIssueRequestDetails) GetCreatedAtUtcOk() (*time.Time, bool)`

GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtUtc

`func (o *DocumentIssueRequestDetails) SetCreatedAtUtc(v time.Time)`

SetCreatedAtUtc sets CreatedAtUtc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



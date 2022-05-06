# DocumentIssueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentTypeId** | **string** | Document type id. | 
**Identifier** | **string** | Document identifier. | 
**Description** | **string** | Document description. | 
**Receiver** | [**DocumentReceiver**](DocumentReceiver.md) |  | 
**IssuedAtUtc** | **time.Time** | Datetime of issue in UTC timezone. | 
**ValidFromUtc** | **time.Time** | Valid from datetime in UTC timezone. | 
**ExpiresAtUtc** | Pointer to **NullableTime** | Datetime of expiry in UTC timezone. | [optional] 
**PaymentRequest** | Pointer to [**PaymentRequest**](PaymentRequest.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** | Metadata. | [optional] 

## Methods

### NewDocumentIssueRequest

`func NewDocumentIssueRequest(documentTypeId string, identifier string, description string, receiver DocumentReceiver, issuedAtUtc time.Time, validFromUtc time.Time, ) *DocumentIssueRequest`

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


### GetIdentifier

`func (o *DocumentIssueRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *DocumentIssueRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *DocumentIssueRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


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


### GetIssuedAtUtc

`func (o *DocumentIssueRequest) GetIssuedAtUtc() time.Time`

GetIssuedAtUtc returns the IssuedAtUtc field if non-nil, zero value otherwise.

### GetIssuedAtUtcOk

`func (o *DocumentIssueRequest) GetIssuedAtUtcOk() (*time.Time, bool)`

GetIssuedAtUtcOk returns a tuple with the IssuedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAtUtc

`func (o *DocumentIssueRequest) SetIssuedAtUtc(v time.Time)`

SetIssuedAtUtc sets IssuedAtUtc field to given value.


### GetValidFromUtc

`func (o *DocumentIssueRequest) GetValidFromUtc() time.Time`

GetValidFromUtc returns the ValidFromUtc field if non-nil, zero value otherwise.

### GetValidFromUtcOk

`func (o *DocumentIssueRequest) GetValidFromUtcOk() (*time.Time, bool)`

GetValidFromUtcOk returns a tuple with the ValidFromUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFromUtc

`func (o *DocumentIssueRequest) SetValidFromUtc(v time.Time)`

SetValidFromUtc sets ValidFromUtc field to given value.


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
### GetPaymentRequest

`func (o *DocumentIssueRequest) GetPaymentRequest() PaymentRequest`

GetPaymentRequest returns the PaymentRequest field if non-nil, zero value otherwise.

### GetPaymentRequestOk

`func (o *DocumentIssueRequest) GetPaymentRequestOk() (*PaymentRequest, bool)`

GetPaymentRequestOk returns a tuple with the PaymentRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequest

`func (o *DocumentIssueRequest) SetPaymentRequest(v PaymentRequest)`

SetPaymentRequest sets PaymentRequest field to given value.

### HasPaymentRequest

`func (o *DocumentIssueRequest) HasPaymentRequest() bool`

HasPaymentRequest returns a boolean if a field has been set.

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



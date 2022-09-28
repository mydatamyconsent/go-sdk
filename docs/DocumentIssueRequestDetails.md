# DocumentIssueRequestDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Document issue request Id. | 
**DocumentTypeId** | **string** | Document type Id. | 
**TypeName** | **string** | Document type name. | 
**Identifier** | **string** | Document identifier. | 
**Status** | [**DocumentIssueRequestStatus**](DocumentIssueRequestStatus.md) |  | 
**Description** | **string** | Document description. | 
**Receiver** | [**DocumentIssueRequestDetailsReceiver**](DocumentIssueRequestDetailsReceiver.md) |  | 
**PaymentRequest** | Pointer to [**PaymentRequest**](PaymentRequest.md) |  | [optional] 
**IssuedAtUtc** | **time.Time** | Datetime of issue in UTC timezone. | 
**ValidFromUtc** | **time.Time** | Valid from datetime in UTC timezone. | 
**ExpiresAtUtc** | Pointer to **time.Time** | Datetime of expiry in UTC timezone. | [optional] 
**MetaData** | Pointer to **interface{}** | Metadata. | [optional] 
**CreatedAtUtc** | **time.Time** | Creation datetime of issue request in UTC timezone. | 

## Methods

### NewDocumentIssueRequestDetails

`func NewDocumentIssueRequestDetails(id string, documentTypeId string, typeName string, identifier string, status DocumentIssueRequestStatus, description string, receiver DocumentIssueRequestDetailsReceiver, issuedAtUtc time.Time, validFromUtc time.Time, createdAtUtc time.Time, ) *DocumentIssueRequestDetails`

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


### GetTypeName

`func (o *DocumentIssueRequestDetails) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *DocumentIssueRequestDetails) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *DocumentIssueRequestDetails) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.


### GetIdentifier

`func (o *DocumentIssueRequestDetails) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *DocumentIssueRequestDetails) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *DocumentIssueRequestDetails) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


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

`func (o *DocumentIssueRequestDetails) GetReceiver() DocumentIssueRequestDetailsReceiver`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *DocumentIssueRequestDetails) GetReceiverOk() (*DocumentIssueRequestDetailsReceiver, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *DocumentIssueRequestDetails) SetReceiver(v DocumentIssueRequestDetailsReceiver)`

SetReceiver sets Receiver field to given value.


### GetPaymentRequest

`func (o *DocumentIssueRequestDetails) GetPaymentRequest() PaymentRequest`

GetPaymentRequest returns the PaymentRequest field if non-nil, zero value otherwise.

### GetPaymentRequestOk

`func (o *DocumentIssueRequestDetails) GetPaymentRequestOk() (*PaymentRequest, bool)`

GetPaymentRequestOk returns a tuple with the PaymentRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequest

`func (o *DocumentIssueRequestDetails) SetPaymentRequest(v PaymentRequest)`

SetPaymentRequest sets PaymentRequest field to given value.

### HasPaymentRequest

`func (o *DocumentIssueRequestDetails) HasPaymentRequest() bool`

HasPaymentRequest returns a boolean if a field has been set.

### GetIssuedAtUtc

`func (o *DocumentIssueRequestDetails) GetIssuedAtUtc() time.Time`

GetIssuedAtUtc returns the IssuedAtUtc field if non-nil, zero value otherwise.

### GetIssuedAtUtcOk

`func (o *DocumentIssueRequestDetails) GetIssuedAtUtcOk() (*time.Time, bool)`

GetIssuedAtUtcOk returns a tuple with the IssuedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAtUtc

`func (o *DocumentIssueRequestDetails) SetIssuedAtUtc(v time.Time)`

SetIssuedAtUtc sets IssuedAtUtc field to given value.


### GetValidFromUtc

`func (o *DocumentIssueRequestDetails) GetValidFromUtc() time.Time`

GetValidFromUtc returns the ValidFromUtc field if non-nil, zero value otherwise.

### GetValidFromUtcOk

`func (o *DocumentIssueRequestDetails) GetValidFromUtcOk() (*time.Time, bool)`

GetValidFromUtcOk returns a tuple with the ValidFromUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFromUtc

`func (o *DocumentIssueRequestDetails) SetValidFromUtc(v time.Time)`

SetValidFromUtc sets ValidFromUtc field to given value.


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

### GetMetaData

`func (o *DocumentIssueRequestDetails) GetMetaData() interface{}`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *DocumentIssueRequestDetails) GetMetaDataOk() (*interface{}, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *DocumentIssueRequestDetails) SetMetaData(v interface{})`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *DocumentIssueRequestDetails) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### SetMetaDataNil

`func (o *DocumentIssueRequestDetails) SetMetaDataNil(b bool)`

 SetMetaDataNil sets the value for MetaData to be an explicit nil

### UnsetMetaData
`func (o *DocumentIssueRequestDetails) UnsetMetaData()`

UnsetMetaData ensures that no value is present for MetaData, not even an explicit nil
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



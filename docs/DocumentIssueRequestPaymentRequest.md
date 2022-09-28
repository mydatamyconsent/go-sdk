# DocumentIssueRequestPaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** |  | 
**Items** | [**[]PaymentOrderItem**](PaymentOrderItem.md) |  | 
**CurrencyCode** | **string** |  | 
**PaymentUrl** | Pointer to **string** |  | [optional] 
**Description** | **string** |  | 
**DueByUtc** | **time.Time** |  | 

## Methods

### NewDocumentIssueRequestPaymentRequest

`func NewDocumentIssueRequestPaymentRequest(identifier string, items []PaymentOrderItem, currencyCode string, description string, dueByUtc time.Time, ) *DocumentIssueRequestPaymentRequest`

NewDocumentIssueRequestPaymentRequest instantiates a new DocumentIssueRequestPaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentIssueRequestPaymentRequestWithDefaults

`func NewDocumentIssueRequestPaymentRequestWithDefaults() *DocumentIssueRequestPaymentRequest`

NewDocumentIssueRequestPaymentRequestWithDefaults instantiates a new DocumentIssueRequestPaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *DocumentIssueRequestPaymentRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *DocumentIssueRequestPaymentRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *DocumentIssueRequestPaymentRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetItems

`func (o *DocumentIssueRequestPaymentRequest) GetItems() []PaymentOrderItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DocumentIssueRequestPaymentRequest) GetItemsOk() (*[]PaymentOrderItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DocumentIssueRequestPaymentRequest) SetItems(v []PaymentOrderItem)`

SetItems sets Items field to given value.


### GetCurrencyCode

`func (o *DocumentIssueRequestPaymentRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *DocumentIssueRequestPaymentRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *DocumentIssueRequestPaymentRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetPaymentUrl

`func (o *DocumentIssueRequestPaymentRequest) GetPaymentUrl() string`

GetPaymentUrl returns the PaymentUrl field if non-nil, zero value otherwise.

### GetPaymentUrlOk

`func (o *DocumentIssueRequestPaymentRequest) GetPaymentUrlOk() (*string, bool)`

GetPaymentUrlOk returns a tuple with the PaymentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentUrl

`func (o *DocumentIssueRequestPaymentRequest) SetPaymentUrl(v string)`

SetPaymentUrl sets PaymentUrl field to given value.

### HasPaymentUrl

`func (o *DocumentIssueRequestPaymentRequest) HasPaymentUrl() bool`

HasPaymentUrl returns a boolean if a field has been set.

### GetDescription

`func (o *DocumentIssueRequestPaymentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DocumentIssueRequestPaymentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DocumentIssueRequestPaymentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDueByUtc

`func (o *DocumentIssueRequestPaymentRequest) GetDueByUtc() time.Time`

GetDueByUtc returns the DueByUtc field if non-nil, zero value otherwise.

### GetDueByUtcOk

`func (o *DocumentIssueRequestPaymentRequest) GetDueByUtcOk() (*time.Time, bool)`

GetDueByUtcOk returns a tuple with the DueByUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueByUtc

`func (o *DocumentIssueRequestPaymentRequest) SetDueByUtc(v time.Time)`

SetDueByUtc sets DueByUtc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



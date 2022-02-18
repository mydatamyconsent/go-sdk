# DataConsentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**TemplateId** | Pointer to **NullableString** |  | [optional] 
**RequestedAtUtc** | Pointer to **time.Time** |  | [optional] 
**RequestExpiredAtUtc** | Pointer to **time.Time** |  | [optional] 
**TransactionId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDataConsentRequest

`func NewDataConsentRequest() *DataConsentRequest`

NewDataConsentRequest instantiates a new DataConsentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataConsentRequestWithDefaults

`func NewDataConsentRequestWithDefaults() *DataConsentRequest`

NewDataConsentRequestWithDefaults instantiates a new DataConsentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataConsentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataConsentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataConsentRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataConsentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTemplateId

`func (o *DataConsentRequest) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *DataConsentRequest) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *DataConsentRequest) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *DataConsentRequest) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *DataConsentRequest) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *DataConsentRequest) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetRequestedAtUtc

`func (o *DataConsentRequest) GetRequestedAtUtc() time.Time`

GetRequestedAtUtc returns the RequestedAtUtc field if non-nil, zero value otherwise.

### GetRequestedAtUtcOk

`func (o *DataConsentRequest) GetRequestedAtUtcOk() (*time.Time, bool)`

GetRequestedAtUtcOk returns a tuple with the RequestedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAtUtc

`func (o *DataConsentRequest) SetRequestedAtUtc(v time.Time)`

SetRequestedAtUtc sets RequestedAtUtc field to given value.

### HasRequestedAtUtc

`func (o *DataConsentRequest) HasRequestedAtUtc() bool`

HasRequestedAtUtc returns a boolean if a field has been set.

### GetRequestExpiredAtUtc

`func (o *DataConsentRequest) GetRequestExpiredAtUtc() time.Time`

GetRequestExpiredAtUtc returns the RequestExpiredAtUtc field if non-nil, zero value otherwise.

### GetRequestExpiredAtUtcOk

`func (o *DataConsentRequest) GetRequestExpiredAtUtcOk() (*time.Time, bool)`

GetRequestExpiredAtUtcOk returns a tuple with the RequestExpiredAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestExpiredAtUtc

`func (o *DataConsentRequest) SetRequestExpiredAtUtc(v time.Time)`

SetRequestExpiredAtUtc sets RequestExpiredAtUtc field to given value.

### HasRequestExpiredAtUtc

`func (o *DataConsentRequest) HasRequestExpiredAtUtc() bool`

HasRequestExpiredAtUtc returns a boolean if a field has been set.

### GetTransactionId

`func (o *DataConsentRequest) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *DataConsentRequest) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *DataConsentRequest) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *DataConsentRequest) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### SetTransactionIdNil

`func (o *DataConsentRequest) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *DataConsentRequest) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



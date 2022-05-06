# PaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **NullableString** |  | [optional] 
**Amount** | Pointer to **NullableString** |  | [optional] 
**CurrencyCode** | Pointer to **NullableString** |  | [optional] 
**PaymentUrl** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**DueByUtc** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPaymentRequest

`func NewPaymentRequest() *PaymentRequest`

NewPaymentRequest instantiates a new PaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRequestWithDefaults

`func NewPaymentRequestWithDefaults() *PaymentRequest`

NewPaymentRequestWithDefaults instantiates a new PaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *PaymentRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PaymentRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PaymentRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *PaymentRequest) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *PaymentRequest) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *PaymentRequest) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetAmount

`func (o *PaymentRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PaymentRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *PaymentRequest) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *PaymentRequest) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetCurrencyCode

`func (o *PaymentRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PaymentRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PaymentRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *PaymentRequest) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *PaymentRequest) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *PaymentRequest) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetPaymentUrl

`func (o *PaymentRequest) GetPaymentUrl() string`

GetPaymentUrl returns the PaymentUrl field if non-nil, zero value otherwise.

### GetPaymentUrlOk

`func (o *PaymentRequest) GetPaymentUrlOk() (*string, bool)`

GetPaymentUrlOk returns a tuple with the PaymentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentUrl

`func (o *PaymentRequest) SetPaymentUrl(v string)`

SetPaymentUrl sets PaymentUrl field to given value.

### HasPaymentUrl

`func (o *PaymentRequest) HasPaymentUrl() bool`

HasPaymentUrl returns a boolean if a field has been set.

### SetPaymentUrlNil

`func (o *PaymentRequest) SetPaymentUrlNil(b bool)`

 SetPaymentUrlNil sets the value for PaymentUrl to be an explicit nil

### UnsetPaymentUrl
`func (o *PaymentRequest) UnsetPaymentUrl()`

UnsetPaymentUrl ensures that no value is present for PaymentUrl, not even an explicit nil
### GetDescription

`func (o *PaymentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PaymentRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PaymentRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDueByUtc

`func (o *PaymentRequest) GetDueByUtc() time.Time`

GetDueByUtc returns the DueByUtc field if non-nil, zero value otherwise.

### GetDueByUtcOk

`func (o *PaymentRequest) GetDueByUtcOk() (*time.Time, bool)`

GetDueByUtcOk returns a tuple with the DueByUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueByUtc

`func (o *PaymentRequest) SetDueByUtc(v time.Time)`

SetDueByUtc sets DueByUtc field to given value.

### HasDueByUtc

`func (o *PaymentRequest) HasDueByUtc() bool`

HasDueByUtc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



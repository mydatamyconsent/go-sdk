# CreditCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Identifier** | **string** |  | 
**Balance** | **float64** |  | 
**Profile** | [**CreditCardProfile**](CreditCardProfile.md) |  | 
**Summary** | [**CreditCardSummary**](CreditCardSummary.md) |  | 
**MaskedAccountNumber** | **string** |  | 
**LinkedAccountRef** | **string** |  | 
**Version** | **float32** |  | 

## Methods

### NewCreditCard

`func NewCreditCard(id string, name string, identifier string, balance float64, profile CreditCardProfile, summary CreditCardSummary, maskedAccountNumber string, linkedAccountRef string, version float32, ) *CreditCard`

NewCreditCard instantiates a new CreditCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditCardWithDefaults

`func NewCreditCardWithDefaults() *CreditCard`

NewCreditCardWithDefaults instantiates a new CreditCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreditCard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreditCard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreditCard) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CreditCard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreditCard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreditCard) SetName(v string)`

SetName sets Name field to given value.


### GetIdentifier

`func (o *CreditCard) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CreditCard) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CreditCard) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetBalance

`func (o *CreditCard) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *CreditCard) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *CreditCard) SetBalance(v float64)`

SetBalance sets Balance field to given value.


### GetProfile

`func (o *CreditCard) GetProfile() CreditCardProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *CreditCard) GetProfileOk() (*CreditCardProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *CreditCard) SetProfile(v CreditCardProfile)`

SetProfile sets Profile field to given value.


### GetSummary

`func (o *CreditCard) GetSummary() CreditCardSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *CreditCard) GetSummaryOk() (*CreditCardSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *CreditCard) SetSummary(v CreditCardSummary)`

SetSummary sets Summary field to given value.


### GetMaskedAccountNumber

`func (o *CreditCard) GetMaskedAccountNumber() string`

GetMaskedAccountNumber returns the MaskedAccountNumber field if non-nil, zero value otherwise.

### GetMaskedAccountNumberOk

`func (o *CreditCard) GetMaskedAccountNumberOk() (*string, bool)`

GetMaskedAccountNumberOk returns a tuple with the MaskedAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedAccountNumber

`func (o *CreditCard) SetMaskedAccountNumber(v string)`

SetMaskedAccountNumber sets MaskedAccountNumber field to given value.


### GetLinkedAccountRef

`func (o *CreditCard) GetLinkedAccountRef() string`

GetLinkedAccountRef returns the LinkedAccountRef field if non-nil, zero value otherwise.

### GetLinkedAccountRefOk

`func (o *CreditCard) GetLinkedAccountRefOk() (*string, bool)`

GetLinkedAccountRefOk returns a tuple with the LinkedAccountRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountRef

`func (o *CreditCard) SetLinkedAccountRef(v string)`

SetLinkedAccountRef sets LinkedAccountRef field to given value.


### GetVersion

`func (o *CreditCard) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreditCard) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreditCard) SetVersion(v float32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



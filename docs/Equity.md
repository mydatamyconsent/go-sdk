# Equity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Identifier** | **string** |  | 
**Balance** | **float64** |  | 
**Profile** | [**Profile**](Profile.md) |  | 
**Summary** | [**EquitySummary**](EquitySummary.md) |  | 
**MaskedAccountNumber** | **string** |  | 
**LinkedAccountRef** | **string** |  | 
**Version** | **float32** |  | 

## Methods

### NewEquity

`func NewEquity(id string, name string, identifier string, balance float64, profile Profile, summary EquitySummary, maskedAccountNumber string, linkedAccountRef string, version float32, ) *Equity`

NewEquity instantiates a new Equity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquityWithDefaults

`func NewEquityWithDefaults() *Equity`

NewEquityWithDefaults instantiates a new Equity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Equity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Equity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Equity) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Equity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Equity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Equity) SetName(v string)`

SetName sets Name field to given value.


### GetIdentifier

`func (o *Equity) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Equity) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Equity) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetBalance

`func (o *Equity) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *Equity) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *Equity) SetBalance(v float64)`

SetBalance sets Balance field to given value.


### GetProfile

`func (o *Equity) GetProfile() Profile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *Equity) GetProfileOk() (*Profile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *Equity) SetProfile(v Profile)`

SetProfile sets Profile field to given value.


### GetSummary

`func (o *Equity) GetSummary() EquitySummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Equity) GetSummaryOk() (*EquitySummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Equity) SetSummary(v EquitySummary)`

SetSummary sets Summary field to given value.


### GetMaskedAccountNumber

`func (o *Equity) GetMaskedAccountNumber() string`

GetMaskedAccountNumber returns the MaskedAccountNumber field if non-nil, zero value otherwise.

### GetMaskedAccountNumberOk

`func (o *Equity) GetMaskedAccountNumberOk() (*string, bool)`

GetMaskedAccountNumberOk returns a tuple with the MaskedAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedAccountNumber

`func (o *Equity) SetMaskedAccountNumber(v string)`

SetMaskedAccountNumber sets MaskedAccountNumber field to given value.


### GetLinkedAccountRef

`func (o *Equity) GetLinkedAccountRef() string`

GetLinkedAccountRef returns the LinkedAccountRef field if non-nil, zero value otherwise.

### GetLinkedAccountRefOk

`func (o *Equity) GetLinkedAccountRefOk() (*string, bool)`

GetLinkedAccountRefOk returns a tuple with the LinkedAccountRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountRef

`func (o *Equity) SetLinkedAccountRef(v string)`

SetLinkedAccountRef sets LinkedAccountRef field to given value.


### GetVersion

`func (o *Equity) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Equity) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Equity) SetVersion(v float32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



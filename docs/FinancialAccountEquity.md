# FinancialAccountEquity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
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

### NewFinancialAccountEquity

`func NewFinancialAccountEquity(type_ string, id string, name string, identifier string, balance float64, profile Profile, summary EquitySummary, maskedAccountNumber string, linkedAccountRef string, version float32, ) *FinancialAccountEquity`

NewFinancialAccountEquity instantiates a new FinancialAccountEquity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialAccountEquityWithDefaults

`func NewFinancialAccountEquityWithDefaults() *FinancialAccountEquity`

NewFinancialAccountEquityWithDefaults instantiates a new FinancialAccountEquity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FinancialAccountEquity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FinancialAccountEquity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FinancialAccountEquity) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *FinancialAccountEquity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FinancialAccountEquity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FinancialAccountEquity) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *FinancialAccountEquity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FinancialAccountEquity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FinancialAccountEquity) SetName(v string)`

SetName sets Name field to given value.


### GetIdentifier

`func (o *FinancialAccountEquity) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *FinancialAccountEquity) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *FinancialAccountEquity) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetBalance

`func (o *FinancialAccountEquity) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *FinancialAccountEquity) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *FinancialAccountEquity) SetBalance(v float64)`

SetBalance sets Balance field to given value.


### GetProfile

`func (o *FinancialAccountEquity) GetProfile() Profile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *FinancialAccountEquity) GetProfileOk() (*Profile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *FinancialAccountEquity) SetProfile(v Profile)`

SetProfile sets Profile field to given value.


### GetSummary

`func (o *FinancialAccountEquity) GetSummary() EquitySummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *FinancialAccountEquity) GetSummaryOk() (*EquitySummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *FinancialAccountEquity) SetSummary(v EquitySummary)`

SetSummary sets Summary field to given value.


### GetMaskedAccountNumber

`func (o *FinancialAccountEquity) GetMaskedAccountNumber() string`

GetMaskedAccountNumber returns the MaskedAccountNumber field if non-nil, zero value otherwise.

### GetMaskedAccountNumberOk

`func (o *FinancialAccountEquity) GetMaskedAccountNumberOk() (*string, bool)`

GetMaskedAccountNumberOk returns a tuple with the MaskedAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedAccountNumber

`func (o *FinancialAccountEquity) SetMaskedAccountNumber(v string)`

SetMaskedAccountNumber sets MaskedAccountNumber field to given value.


### GetLinkedAccountRef

`func (o *FinancialAccountEquity) GetLinkedAccountRef() string`

GetLinkedAccountRef returns the LinkedAccountRef field if non-nil, zero value otherwise.

### GetLinkedAccountRefOk

`func (o *FinancialAccountEquity) GetLinkedAccountRefOk() (*string, bool)`

GetLinkedAccountRefOk returns a tuple with the LinkedAccountRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountRef

`func (o *FinancialAccountEquity) SetLinkedAccountRef(v string)`

SetLinkedAccountRef sets LinkedAccountRef field to given value.


### GetVersion

`func (o *FinancialAccountEquity) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FinancialAccountEquity) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FinancialAccountEquity) SetVersion(v float32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



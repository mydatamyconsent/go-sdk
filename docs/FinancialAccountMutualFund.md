# FinancialAccountMutualFund

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Identifier** | **string** |  | 
**Balance** | **float64** |  | 
**Profile** | [**Profile**](Profile.md) |  | 
**Summary** | [**MutualFundSummary**](MutualFundSummary.md) |  | 
**MaskedAccountNumber** | **string** |  | 
**LinkedAccountRef** | **string** |  | 
**Version** | **float32** |  | 

## Methods

### NewFinancialAccountMutualFund

`func NewFinancialAccountMutualFund(type_ string, id string, name string, identifier string, balance float64, profile Profile, summary MutualFundSummary, maskedAccountNumber string, linkedAccountRef string, version float32, ) *FinancialAccountMutualFund`

NewFinancialAccountMutualFund instantiates a new FinancialAccountMutualFund object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialAccountMutualFundWithDefaults

`func NewFinancialAccountMutualFundWithDefaults() *FinancialAccountMutualFund`

NewFinancialAccountMutualFundWithDefaults instantiates a new FinancialAccountMutualFund object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FinancialAccountMutualFund) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FinancialAccountMutualFund) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FinancialAccountMutualFund) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *FinancialAccountMutualFund) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FinancialAccountMutualFund) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FinancialAccountMutualFund) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *FinancialAccountMutualFund) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FinancialAccountMutualFund) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FinancialAccountMutualFund) SetName(v string)`

SetName sets Name field to given value.


### GetIdentifier

`func (o *FinancialAccountMutualFund) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *FinancialAccountMutualFund) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *FinancialAccountMutualFund) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetBalance

`func (o *FinancialAccountMutualFund) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *FinancialAccountMutualFund) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *FinancialAccountMutualFund) SetBalance(v float64)`

SetBalance sets Balance field to given value.


### GetProfile

`func (o *FinancialAccountMutualFund) GetProfile() Profile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *FinancialAccountMutualFund) GetProfileOk() (*Profile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *FinancialAccountMutualFund) SetProfile(v Profile)`

SetProfile sets Profile field to given value.


### GetSummary

`func (o *FinancialAccountMutualFund) GetSummary() MutualFundSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *FinancialAccountMutualFund) GetSummaryOk() (*MutualFundSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *FinancialAccountMutualFund) SetSummary(v MutualFundSummary)`

SetSummary sets Summary field to given value.


### GetMaskedAccountNumber

`func (o *FinancialAccountMutualFund) GetMaskedAccountNumber() string`

GetMaskedAccountNumber returns the MaskedAccountNumber field if non-nil, zero value otherwise.

### GetMaskedAccountNumberOk

`func (o *FinancialAccountMutualFund) GetMaskedAccountNumberOk() (*string, bool)`

GetMaskedAccountNumberOk returns a tuple with the MaskedAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedAccountNumber

`func (o *FinancialAccountMutualFund) SetMaskedAccountNumber(v string)`

SetMaskedAccountNumber sets MaskedAccountNumber field to given value.


### GetLinkedAccountRef

`func (o *FinancialAccountMutualFund) GetLinkedAccountRef() string`

GetLinkedAccountRef returns the LinkedAccountRef field if non-nil, zero value otherwise.

### GetLinkedAccountRefOk

`func (o *FinancialAccountMutualFund) GetLinkedAccountRefOk() (*string, bool)`

GetLinkedAccountRefOk returns a tuple with the LinkedAccountRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountRef

`func (o *FinancialAccountMutualFund) SetLinkedAccountRef(v string)`

SetLinkedAccountRef sets LinkedAccountRef field to given value.


### GetVersion

`func (o *FinancialAccountMutualFund) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FinancialAccountMutualFund) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FinancialAccountMutualFund) SetVersion(v float32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



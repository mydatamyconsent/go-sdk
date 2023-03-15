# TermDepositAccountDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | [**TermDepositInformation**](TermDepositInformation.md) |  | 
**Holder** | [**Holder**](Holder.md) |  | 
**Bank** | [**TermDepositBankInformation**](TermDepositBankInformation.md) |  | 

## Methods

### NewTermDepositAccountDetails

`func NewTermDepositAccountDetails(info TermDepositInformation, holder Holder, bank TermDepositBankInformation, ) *TermDepositAccountDetails`

NewTermDepositAccountDetails instantiates a new TermDepositAccountDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTermDepositAccountDetailsWithDefaults

`func NewTermDepositAccountDetailsWithDefaults() *TermDepositAccountDetails`

NewTermDepositAccountDetailsWithDefaults instantiates a new TermDepositAccountDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *TermDepositAccountDetails) GetInfo() TermDepositInformation`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TermDepositAccountDetails) GetInfoOk() (*TermDepositInformation, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TermDepositAccountDetails) SetInfo(v TermDepositInformation)`

SetInfo sets Info field to given value.


### GetHolder

`func (o *TermDepositAccountDetails) GetHolder() Holder`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *TermDepositAccountDetails) GetHolderOk() (*Holder, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *TermDepositAccountDetails) SetHolder(v Holder)`

SetHolder sets Holder field to given value.


### GetBank

`func (o *TermDepositAccountDetails) GetBank() TermDepositBankInformation`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *TermDepositAccountDetails) GetBankOk() (*TermDepositBankInformation, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *TermDepositAccountDetails) SetBank(v TermDepositBankInformation)`

SetBank sets Bank field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



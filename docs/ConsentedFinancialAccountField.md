# ConsentedFinancialAccountField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldTitle** | **string** | Financial account field title. | 
**FieldSlug** | **string** | Financial account field slug. | 
**RequestedDetails** | [**[]FinancialAccountDetailsRequired**](FinancialAccountDetailsRequired.md) | Requested financial account details. | 
**TransactionPeriod** | Pointer to [**ConsentedFinancialAccountFieldTransactionPeriod**](ConsentedFinancialAccountFieldTransactionPeriod.md) |  | [optional] 
**ConsentedAccounts** | [**[]ConsentedFinancialAccount**](ConsentedFinancialAccount.md) | Consented financial accounts. | 

## Methods

### NewConsentedFinancialAccountField

`func NewConsentedFinancialAccountField(fieldTitle string, fieldSlug string, requestedDetails []FinancialAccountDetailsRequired, consentedAccounts []ConsentedFinancialAccount, ) *ConsentedFinancialAccountField`

NewConsentedFinancialAccountField instantiates a new ConsentedFinancialAccountField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsentedFinancialAccountFieldWithDefaults

`func NewConsentedFinancialAccountFieldWithDefaults() *ConsentedFinancialAccountField`

NewConsentedFinancialAccountFieldWithDefaults instantiates a new ConsentedFinancialAccountField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldTitle

`func (o *ConsentedFinancialAccountField) GetFieldTitle() string`

GetFieldTitle returns the FieldTitle field if non-nil, zero value otherwise.

### GetFieldTitleOk

`func (o *ConsentedFinancialAccountField) GetFieldTitleOk() (*string, bool)`

GetFieldTitleOk returns a tuple with the FieldTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldTitle

`func (o *ConsentedFinancialAccountField) SetFieldTitle(v string)`

SetFieldTitle sets FieldTitle field to given value.


### GetFieldSlug

`func (o *ConsentedFinancialAccountField) GetFieldSlug() string`

GetFieldSlug returns the FieldSlug field if non-nil, zero value otherwise.

### GetFieldSlugOk

`func (o *ConsentedFinancialAccountField) GetFieldSlugOk() (*string, bool)`

GetFieldSlugOk returns a tuple with the FieldSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldSlug

`func (o *ConsentedFinancialAccountField) SetFieldSlug(v string)`

SetFieldSlug sets FieldSlug field to given value.


### GetRequestedDetails

`func (o *ConsentedFinancialAccountField) GetRequestedDetails() []FinancialAccountDetailsRequired`

GetRequestedDetails returns the RequestedDetails field if non-nil, zero value otherwise.

### GetRequestedDetailsOk

`func (o *ConsentedFinancialAccountField) GetRequestedDetailsOk() (*[]FinancialAccountDetailsRequired, bool)`

GetRequestedDetailsOk returns a tuple with the RequestedDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedDetails

`func (o *ConsentedFinancialAccountField) SetRequestedDetails(v []FinancialAccountDetailsRequired)`

SetRequestedDetails sets RequestedDetails field to given value.


### GetTransactionPeriod

`func (o *ConsentedFinancialAccountField) GetTransactionPeriod() ConsentedFinancialAccountFieldTransactionPeriod`

GetTransactionPeriod returns the TransactionPeriod field if non-nil, zero value otherwise.

### GetTransactionPeriodOk

`func (o *ConsentedFinancialAccountField) GetTransactionPeriodOk() (*ConsentedFinancialAccountFieldTransactionPeriod, bool)`

GetTransactionPeriodOk returns a tuple with the TransactionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionPeriod

`func (o *ConsentedFinancialAccountField) SetTransactionPeriod(v ConsentedFinancialAccountFieldTransactionPeriod)`

SetTransactionPeriod sets TransactionPeriod field to given value.

### HasTransactionPeriod

`func (o *ConsentedFinancialAccountField) HasTransactionPeriod() bool`

HasTransactionPeriod returns a boolean if a field has been set.

### GetConsentedAccounts

`func (o *ConsentedFinancialAccountField) GetConsentedAccounts() []ConsentedFinancialAccount`

GetConsentedAccounts returns the ConsentedAccounts field if non-nil, zero value otherwise.

### GetConsentedAccountsOk

`func (o *ConsentedFinancialAccountField) GetConsentedAccountsOk() (*[]ConsentedFinancialAccount, bool)`

GetConsentedAccountsOk returns a tuple with the ConsentedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentedAccounts

`func (o *ConsentedFinancialAccountField) SetConsentedAccounts(v []ConsentedFinancialAccount)`

SetConsentedAccounts sets ConsentedAccounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Financial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | Pointer to **NullableString** |  | [optional] 
**CustomKey** | Pointer to **NullableString** |  | [optional] 
**Accounts** | Pointer to [**[]FinancialAccounts**](FinancialAccounts.md) |  | [optional] 
**Requirement** | Pointer to [**DocumentsRequired**](DocumentsRequired.md) |  | [optional] 

## Methods

### NewFinancial

`func NewFinancial() *Financial`

NewFinancial instantiates a new Financial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialWithDefaults

`func NewFinancialWithDefaults() *Financial`

NewFinancialWithDefaults instantiates a new Financial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *Financial) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *Financial) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *Financial) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *Financial) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### SetFieldNameNil

`func (o *Financial) SetFieldNameNil(b bool)`

 SetFieldNameNil sets the value for FieldName to be an explicit nil

### UnsetFieldName
`func (o *Financial) UnsetFieldName()`

UnsetFieldName ensures that no value is present for FieldName, not even an explicit nil
### GetCustomKey

`func (o *Financial) GetCustomKey() string`

GetCustomKey returns the CustomKey field if non-nil, zero value otherwise.

### GetCustomKeyOk

`func (o *Financial) GetCustomKeyOk() (*string, bool)`

GetCustomKeyOk returns a tuple with the CustomKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomKey

`func (o *Financial) SetCustomKey(v string)`

SetCustomKey sets CustomKey field to given value.

### HasCustomKey

`func (o *Financial) HasCustomKey() bool`

HasCustomKey returns a boolean if a field has been set.

### SetCustomKeyNil

`func (o *Financial) SetCustomKeyNil(b bool)`

 SetCustomKeyNil sets the value for CustomKey to be an explicit nil

### UnsetCustomKey
`func (o *Financial) UnsetCustomKey()`

UnsetCustomKey ensures that no value is present for CustomKey, not even an explicit nil
### GetAccounts

`func (o *Financial) GetAccounts() []FinancialAccounts`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *Financial) GetAccountsOk() (*[]FinancialAccounts, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *Financial) SetAccounts(v []FinancialAccounts)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *Financial) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### SetAccountsNil

`func (o *Financial) SetAccountsNil(b bool)`

 SetAccountsNil sets the value for Accounts to be an explicit nil

### UnsetAccounts
`func (o *Financial) UnsetAccounts()`

UnsetAccounts ensures that no value is present for Accounts, not even an explicit nil
### GetRequirement

`func (o *Financial) GetRequirement() DocumentsRequired`

GetRequirement returns the Requirement field if non-nil, zero value otherwise.

### GetRequirementOk

`func (o *Financial) GetRequirementOk() (*DocumentsRequired, bool)`

GetRequirementOk returns a tuple with the Requirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirement

`func (o *Financial) SetRequirement(v DocumentsRequired)`

SetRequirement sets Requirement field to given value.

### HasRequirement

`func (o *Financial) HasRequirement() bool`

HasRequirement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



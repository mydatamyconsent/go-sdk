# FinancialAccountField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldTitle** | **string** | Field title. | 
**FieldSlug** | **string** | Field slug. | 
**AccountTypes** | [**[]SelectedFinancialAccountType**](SelectedFinancialAccountType.md) | Selected financial account types. | 

## Methods

### NewFinancialAccountField

`func NewFinancialAccountField(fieldTitle string, fieldSlug string, accountTypes []SelectedFinancialAccountType, ) *FinancialAccountField`

NewFinancialAccountField instantiates a new FinancialAccountField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialAccountFieldWithDefaults

`func NewFinancialAccountFieldWithDefaults() *FinancialAccountField`

NewFinancialAccountFieldWithDefaults instantiates a new FinancialAccountField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldTitle

`func (o *FinancialAccountField) GetFieldTitle() string`

GetFieldTitle returns the FieldTitle field if non-nil, zero value otherwise.

### GetFieldTitleOk

`func (o *FinancialAccountField) GetFieldTitleOk() (*string, bool)`

GetFieldTitleOk returns a tuple with the FieldTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldTitle

`func (o *FinancialAccountField) SetFieldTitle(v string)`

SetFieldTitle sets FieldTitle field to given value.


### GetFieldSlug

`func (o *FinancialAccountField) GetFieldSlug() string`

GetFieldSlug returns the FieldSlug field if non-nil, zero value otherwise.

### GetFieldSlugOk

`func (o *FinancialAccountField) GetFieldSlugOk() (*string, bool)`

GetFieldSlugOk returns a tuple with the FieldSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldSlug

`func (o *FinancialAccountField) SetFieldSlug(v string)`

SetFieldSlug sets FieldSlug field to given value.


### GetAccountTypes

`func (o *FinancialAccountField) GetAccountTypes() []SelectedFinancialAccountType`

GetAccountTypes returns the AccountTypes field if non-nil, zero value otherwise.

### GetAccountTypesOk

`func (o *FinancialAccountField) GetAccountTypesOk() (*[]SelectedFinancialAccountType, bool)`

GetAccountTypesOk returns a tuple with the AccountTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTypes

`func (o *FinancialAccountField) SetAccountTypes(v []SelectedFinancialAccountType)`

SetAccountTypes sets AccountTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



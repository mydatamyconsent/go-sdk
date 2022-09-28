# SelectedFinancialAccountType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubCategory** | [**FinancialAccountSubCategoryType**](FinancialAccountSubCategoryType.md) |  | 
**Drns** | **[]string** | DRNs. | 

## Methods

### NewSelectedFinancialAccountType

`func NewSelectedFinancialAccountType(subCategory FinancialAccountSubCategoryType, drns []string, ) *SelectedFinancialAccountType`

NewSelectedFinancialAccountType instantiates a new SelectedFinancialAccountType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectedFinancialAccountTypeWithDefaults

`func NewSelectedFinancialAccountTypeWithDefaults() *SelectedFinancialAccountType`

NewSelectedFinancialAccountTypeWithDefaults instantiates a new SelectedFinancialAccountType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubCategory

`func (o *SelectedFinancialAccountType) GetSubCategory() FinancialAccountSubCategoryType`

GetSubCategory returns the SubCategory field if non-nil, zero value otherwise.

### GetSubCategoryOk

`func (o *SelectedFinancialAccountType) GetSubCategoryOk() (*FinancialAccountSubCategoryType, bool)`

GetSubCategoryOk returns a tuple with the SubCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCategory

`func (o *SelectedFinancialAccountType) SetSubCategory(v FinancialAccountSubCategoryType)`

SetSubCategory sets SubCategory field to given value.


### GetDrns

`func (o *SelectedFinancialAccountType) GetDrns() []string`

GetDrns returns the Drns field if non-nil, zero value otherwise.

### GetDrnsOk

`func (o *SelectedFinancialAccountType) GetDrnsOk() (*[]string, bool)`

GetDrnsOk returns a tuple with the Drns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrns

`func (o *SelectedFinancialAccountType) SetDrns(v []string)`

SetDrns sets Drns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



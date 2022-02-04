# OrganizationFinancialTransactionsDtoPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageIndex** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**TotalPages** | Pointer to **int32** |  | [optional] [readonly] 
**TotalItems** | Pointer to **int64** |  | [optional] 
**Items** | Pointer to [**[]OrganizationFinancialTransactionsDto**](OrganizationFinancialTransactionsDto.md) |  | [optional] 

## Methods

### NewOrganizationFinancialTransactionsDtoPaginatedList

`func NewOrganizationFinancialTransactionsDtoPaginatedList() *OrganizationFinancialTransactionsDtoPaginatedList`

NewOrganizationFinancialTransactionsDtoPaginatedList instantiates a new OrganizationFinancialTransactionsDtoPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationFinancialTransactionsDtoPaginatedListWithDefaults

`func NewOrganizationFinancialTransactionsDtoPaginatedListWithDefaults() *OrganizationFinancialTransactionsDtoPaginatedList`

NewOrganizationFinancialTransactionsDtoPaginatedListWithDefaults instantiates a new OrganizationFinancialTransactionsDtoPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageIndex

`func (o *OrganizationFinancialTransactionsDtoPaginatedList) GetPageIndex() int32`

GetPageIndex returns the PageIndex field if non-nil, zero value otherwise.

### GetPageIndexOk

`func (o *OrganizationFinancialTransactionsDtoPaginatedList) GetPageIndexOk() (*int32, bool)`

GetPageIndexOk returns a tuple with the PageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIndex

`func (o *OrganizationFinancialTransactionsDtoPaginatedList) SetPageIndex(v int32)`

SetPageIndex sets PageIndex field to given value.

### HasPageIndex

`func (o *OrganizationFinancialTransactionsDtoPaginatedList) HasPageIndex() bool`

HasPageIndex returns a boolean if a field has been set.

### GetPageSize

`func (o *OrganizationFinancialTransactionsDtoPaginatedList) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *OrganizationFinancialTransactionsDtoPaginatedList) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *OrganizationFinancialTransactionsDtoPaginatedList) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *OrganizationFinancialTransactionsDtoPaginatedList) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalPages

`func (o *OrganizationFinancialTransactionsDtoPaginatedList) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *OrganizationFinancialTransactionsDtoPaginatedList) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *OrganizationFinancialTransactionsDtoPaginatedList) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *OrganizationFinancialTransactionsDtoPaginatedList) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetTotalItems

`func (o *OrganizationFinancialTransactionsDtoPaginatedList) GetTotalItems() int64`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *OrganizationFinancialTransactionsDtoPaginatedList) GetTotalItemsOk() (*int64, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *OrganizationFinancialTransactionsDtoPaginatedList) SetTotalItems(v int64)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *OrganizationFinancialTransactionsDtoPaginatedList) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### GetItems

`func (o *OrganizationFinancialTransactionsDtoPaginatedList) GetItems() []OrganizationFinancialTransactionsDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *OrganizationFinancialTransactionsDtoPaginatedList) GetItemsOk() (*[]OrganizationFinancialTransactionsDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *OrganizationFinancialTransactionsDtoPaginatedList) SetItems(v []OrganizationFinancialTransactionsDto)`

SetItems sets Items field to given value.

### HasItems

`func (o *OrganizationFinancialTransactionsDtoPaginatedList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *OrganizationFinancialTransactionsDtoPaginatedList) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *OrganizationFinancialTransactionsDtoPaginatedList) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



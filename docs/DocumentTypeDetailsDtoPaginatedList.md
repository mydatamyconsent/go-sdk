# DocumentTypeDetailsDtoPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageIndex** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**TotalPages** | Pointer to **int32** |  | [optional] [readonly] 
**TotalItems** | Pointer to **int64** |  | [optional] 
**Items** | Pointer to [**[]DocumentTypeDetailsDto**](DocumentTypeDetailsDto.md) |  | [optional] 

## Methods

### NewDocumentTypeDetailsDtoPaginatedList

`func NewDocumentTypeDetailsDtoPaginatedList() *DocumentTypeDetailsDtoPaginatedList`

NewDocumentTypeDetailsDtoPaginatedList instantiates a new DocumentTypeDetailsDtoPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentTypeDetailsDtoPaginatedListWithDefaults

`func NewDocumentTypeDetailsDtoPaginatedListWithDefaults() *DocumentTypeDetailsDtoPaginatedList`

NewDocumentTypeDetailsDtoPaginatedListWithDefaults instantiates a new DocumentTypeDetailsDtoPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageIndex

`func (o *DocumentTypeDetailsDtoPaginatedList) GetPageIndex() int32`

GetPageIndex returns the PageIndex field if non-nil, zero value otherwise.

### GetPageIndexOk

`func (o *DocumentTypeDetailsDtoPaginatedList) GetPageIndexOk() (*int32, bool)`

GetPageIndexOk returns a tuple with the PageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIndex

`func (o *DocumentTypeDetailsDtoPaginatedList) SetPageIndex(v int32)`

SetPageIndex sets PageIndex field to given value.

### HasPageIndex

`func (o *DocumentTypeDetailsDtoPaginatedList) HasPageIndex() bool`

HasPageIndex returns a boolean if a field has been set.

### GetPageSize

`func (o *DocumentTypeDetailsDtoPaginatedList) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DocumentTypeDetailsDtoPaginatedList) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DocumentTypeDetailsDtoPaginatedList) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DocumentTypeDetailsDtoPaginatedList) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalPages

`func (o *DocumentTypeDetailsDtoPaginatedList) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *DocumentTypeDetailsDtoPaginatedList) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *DocumentTypeDetailsDtoPaginatedList) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *DocumentTypeDetailsDtoPaginatedList) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetTotalItems

`func (o *DocumentTypeDetailsDtoPaginatedList) GetTotalItems() int64`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *DocumentTypeDetailsDtoPaginatedList) GetTotalItemsOk() (*int64, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *DocumentTypeDetailsDtoPaginatedList) SetTotalItems(v int64)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *DocumentTypeDetailsDtoPaginatedList) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### GetItems

`func (o *DocumentTypeDetailsDtoPaginatedList) GetItems() []DocumentTypeDetailsDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DocumentTypeDetailsDtoPaginatedList) GetItemsOk() (*[]DocumentTypeDetailsDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DocumentTypeDetailsDtoPaginatedList) SetItems(v []DocumentTypeDetailsDto)`

SetItems sets Items field to given value.

### HasItems

`func (o *DocumentTypeDetailsDtoPaginatedList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *DocumentTypeDetailsDtoPaginatedList) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *DocumentTypeDetailsDtoPaginatedList) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



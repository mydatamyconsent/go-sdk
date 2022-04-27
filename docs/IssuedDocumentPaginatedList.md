# IssuedDocumentPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageIndex** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**TotalPages** | Pointer to **int32** |  | [optional] [readonly] 
**TotalItems** | Pointer to **int64** |  | [optional] 
**Items** | Pointer to [**[]IssuedDocument**](IssuedDocument.md) |  | [optional] 

## Methods

### NewIssuedDocumentPaginatedList

`func NewIssuedDocumentPaginatedList() *IssuedDocumentPaginatedList`

NewIssuedDocumentPaginatedList instantiates a new IssuedDocumentPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuedDocumentPaginatedListWithDefaults

`func NewIssuedDocumentPaginatedListWithDefaults() *IssuedDocumentPaginatedList`

NewIssuedDocumentPaginatedListWithDefaults instantiates a new IssuedDocumentPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageIndex

`func (o *IssuedDocumentPaginatedList) GetPageIndex() int32`

GetPageIndex returns the PageIndex field if non-nil, zero value otherwise.

### GetPageIndexOk

`func (o *IssuedDocumentPaginatedList) GetPageIndexOk() (*int32, bool)`

GetPageIndexOk returns a tuple with the PageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIndex

`func (o *IssuedDocumentPaginatedList) SetPageIndex(v int32)`

SetPageIndex sets PageIndex field to given value.

### HasPageIndex

`func (o *IssuedDocumentPaginatedList) HasPageIndex() bool`

HasPageIndex returns a boolean if a field has been set.

### GetPageSize

`func (o *IssuedDocumentPaginatedList) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *IssuedDocumentPaginatedList) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *IssuedDocumentPaginatedList) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *IssuedDocumentPaginatedList) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalPages

`func (o *IssuedDocumentPaginatedList) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *IssuedDocumentPaginatedList) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *IssuedDocumentPaginatedList) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *IssuedDocumentPaginatedList) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetTotalItems

`func (o *IssuedDocumentPaginatedList) GetTotalItems() int64`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *IssuedDocumentPaginatedList) GetTotalItemsOk() (*int64, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *IssuedDocumentPaginatedList) SetTotalItems(v int64)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *IssuedDocumentPaginatedList) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### GetItems

`func (o *IssuedDocumentPaginatedList) GetItems() []IssuedDocument`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *IssuedDocumentPaginatedList) GetItemsOk() (*[]IssuedDocument, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *IssuedDocumentPaginatedList) SetItems(v []IssuedDocument)`

SetItems sets Items field to given value.

### HasItems

`func (o *IssuedDocumentPaginatedList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *IssuedDocumentPaginatedList) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *IssuedDocumentPaginatedList) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



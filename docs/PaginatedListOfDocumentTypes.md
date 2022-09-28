# PaginatedListOfDocumentTypes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNo** | **int32** |  | 
**PageSize** | **int32** |  | 
**TotalPage** | **int32** |  | 
**Items** | [**[]DocumentType**](DocumentType.md) |  | 

## Methods

### NewPaginatedListOfDocumentTypes

`func NewPaginatedListOfDocumentTypes(pageNo int32, pageSize int32, totalPage int32, items []DocumentType, ) *PaginatedListOfDocumentTypes`

NewPaginatedListOfDocumentTypes instantiates a new PaginatedListOfDocumentTypes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedListOfDocumentTypesWithDefaults

`func NewPaginatedListOfDocumentTypesWithDefaults() *PaginatedListOfDocumentTypes`

NewPaginatedListOfDocumentTypesWithDefaults instantiates a new PaginatedListOfDocumentTypes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNo

`func (o *PaginatedListOfDocumentTypes) GetPageNo() int32`

GetPageNo returns the PageNo field if non-nil, zero value otherwise.

### GetPageNoOk

`func (o *PaginatedListOfDocumentTypes) GetPageNoOk() (*int32, bool)`

GetPageNoOk returns a tuple with the PageNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNo

`func (o *PaginatedListOfDocumentTypes) SetPageNo(v int32)`

SetPageNo sets PageNo field to given value.


### GetPageSize

`func (o *PaginatedListOfDocumentTypes) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PaginatedListOfDocumentTypes) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PaginatedListOfDocumentTypes) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetTotalPage

`func (o *PaginatedListOfDocumentTypes) GetTotalPage() int32`

GetTotalPage returns the TotalPage field if non-nil, zero value otherwise.

### GetTotalPageOk

`func (o *PaginatedListOfDocumentTypes) GetTotalPageOk() (*int32, bool)`

GetTotalPageOk returns a tuple with the TotalPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPage

`func (o *PaginatedListOfDocumentTypes) SetTotalPage(v int32)`

SetTotalPage sets TotalPage field to given value.


### GetItems

`func (o *PaginatedListOfDocumentTypes) GetItems() []DocumentType`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PaginatedListOfDocumentTypes) GetItemsOk() (*[]DocumentType, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PaginatedListOfDocumentTypes) SetItems(v []DocumentType)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PaginatedListOfIssuedDocuments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNo** | **int32** |  | 
**PageSize** | **int32** |  | 
**TotalPage** | **int32** |  | 
**Items** | [**[]IssuedDocument**](IssuedDocument.md) |  | 

## Methods

### NewPaginatedListOfIssuedDocuments

`func NewPaginatedListOfIssuedDocuments(pageNo int32, pageSize int32, totalPage int32, items []IssuedDocument, ) *PaginatedListOfIssuedDocuments`

NewPaginatedListOfIssuedDocuments instantiates a new PaginatedListOfIssuedDocuments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedListOfIssuedDocumentsWithDefaults

`func NewPaginatedListOfIssuedDocumentsWithDefaults() *PaginatedListOfIssuedDocuments`

NewPaginatedListOfIssuedDocumentsWithDefaults instantiates a new PaginatedListOfIssuedDocuments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNo

`func (o *PaginatedListOfIssuedDocuments) GetPageNo() int32`

GetPageNo returns the PageNo field if non-nil, zero value otherwise.

### GetPageNoOk

`func (o *PaginatedListOfIssuedDocuments) GetPageNoOk() (*int32, bool)`

GetPageNoOk returns a tuple with the PageNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNo

`func (o *PaginatedListOfIssuedDocuments) SetPageNo(v int32)`

SetPageNo sets PageNo field to given value.


### GetPageSize

`func (o *PaginatedListOfIssuedDocuments) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PaginatedListOfIssuedDocuments) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PaginatedListOfIssuedDocuments) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetTotalPage

`func (o *PaginatedListOfIssuedDocuments) GetTotalPage() int32`

GetTotalPage returns the TotalPage field if non-nil, zero value otherwise.

### GetTotalPageOk

`func (o *PaginatedListOfIssuedDocuments) GetTotalPageOk() (*int32, bool)`

GetTotalPageOk returns a tuple with the TotalPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPage

`func (o *PaginatedListOfIssuedDocuments) SetTotalPage(v int32)`

SetTotalPage sets TotalPage field to given value.


### GetItems

`func (o *PaginatedListOfIssuedDocuments) GetItems() []IssuedDocument`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PaginatedListOfIssuedDocuments) GetItemsOk() (*[]IssuedDocument, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PaginatedListOfIssuedDocuments) SetItems(v []IssuedDocument)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



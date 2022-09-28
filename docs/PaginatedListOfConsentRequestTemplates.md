# PaginatedListOfConsentRequestTemplates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNo** | **int32** |  | 
**PageSize** | **int32** |  | 
**TotalPage** | **int32** |  | 
**Items** | [**[]ConsentRequestTemplate**](ConsentRequestTemplate.md) |  | 

## Methods

### NewPaginatedListOfConsentRequestTemplates

`func NewPaginatedListOfConsentRequestTemplates(pageNo int32, pageSize int32, totalPage int32, items []ConsentRequestTemplate, ) *PaginatedListOfConsentRequestTemplates`

NewPaginatedListOfConsentRequestTemplates instantiates a new PaginatedListOfConsentRequestTemplates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedListOfConsentRequestTemplatesWithDefaults

`func NewPaginatedListOfConsentRequestTemplatesWithDefaults() *PaginatedListOfConsentRequestTemplates`

NewPaginatedListOfConsentRequestTemplatesWithDefaults instantiates a new PaginatedListOfConsentRequestTemplates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNo

`func (o *PaginatedListOfConsentRequestTemplates) GetPageNo() int32`

GetPageNo returns the PageNo field if non-nil, zero value otherwise.

### GetPageNoOk

`func (o *PaginatedListOfConsentRequestTemplates) GetPageNoOk() (*int32, bool)`

GetPageNoOk returns a tuple with the PageNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNo

`func (o *PaginatedListOfConsentRequestTemplates) SetPageNo(v int32)`

SetPageNo sets PageNo field to given value.


### GetPageSize

`func (o *PaginatedListOfConsentRequestTemplates) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PaginatedListOfConsentRequestTemplates) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PaginatedListOfConsentRequestTemplates) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetTotalPage

`func (o *PaginatedListOfConsentRequestTemplates) GetTotalPage() int32`

GetTotalPage returns the TotalPage field if non-nil, zero value otherwise.

### GetTotalPageOk

`func (o *PaginatedListOfConsentRequestTemplates) GetTotalPageOk() (*int32, bool)`

GetTotalPageOk returns a tuple with the TotalPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPage

`func (o *PaginatedListOfConsentRequestTemplates) SetTotalPage(v int32)`

SetTotalPage sets TotalPage field to given value.


### GetItems

`func (o *PaginatedListOfConsentRequestTemplates) GetItems() []ConsentRequestTemplate`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PaginatedListOfConsentRequestTemplates) GetItemsOk() (*[]ConsentRequestTemplate, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PaginatedListOfConsentRequestTemplates) SetItems(v []ConsentRequestTemplate)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



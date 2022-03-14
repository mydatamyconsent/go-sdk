# OrganizationDataConsentRequestDetailsPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageIndex** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**TotalPages** | Pointer to **int32** |  | [optional] [readonly] 
**TotalItems** | Pointer to **int64** |  | [optional] 
**Items** | Pointer to [**[]OrganizationDataConsentRequestDetails**](OrganizationDataConsentRequestDetails.md) |  | [optional] 

## Methods

### NewOrganizationDataConsentRequestDetailsPaginatedList

`func NewOrganizationDataConsentRequestDetailsPaginatedList() *OrganizationDataConsentRequestDetailsPaginatedList`

NewOrganizationDataConsentRequestDetailsPaginatedList instantiates a new OrganizationDataConsentRequestDetailsPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationDataConsentRequestDetailsPaginatedListWithDefaults

`func NewOrganizationDataConsentRequestDetailsPaginatedListWithDefaults() *OrganizationDataConsentRequestDetailsPaginatedList`

NewOrganizationDataConsentRequestDetailsPaginatedListWithDefaults instantiates a new OrganizationDataConsentRequestDetailsPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageIndex

`func (o *OrganizationDataConsentRequestDetailsPaginatedList) GetPageIndex() int32`

GetPageIndex returns the PageIndex field if non-nil, zero value otherwise.

### GetPageIndexOk

`func (o *OrganizationDataConsentRequestDetailsPaginatedList) GetPageIndexOk() (*int32, bool)`

GetPageIndexOk returns a tuple with the PageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIndex

`func (o *OrganizationDataConsentRequestDetailsPaginatedList) SetPageIndex(v int32)`

SetPageIndex sets PageIndex field to given value.

### HasPageIndex

`func (o *OrganizationDataConsentRequestDetailsPaginatedList) HasPageIndex() bool`

HasPageIndex returns a boolean if a field has been set.

### GetPageSize

`func (o *OrganizationDataConsentRequestDetailsPaginatedList) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *OrganizationDataConsentRequestDetailsPaginatedList) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *OrganizationDataConsentRequestDetailsPaginatedList) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *OrganizationDataConsentRequestDetailsPaginatedList) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalPages

`func (o *OrganizationDataConsentRequestDetailsPaginatedList) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *OrganizationDataConsentRequestDetailsPaginatedList) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *OrganizationDataConsentRequestDetailsPaginatedList) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *OrganizationDataConsentRequestDetailsPaginatedList) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetTotalItems

`func (o *OrganizationDataConsentRequestDetailsPaginatedList) GetTotalItems() int64`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *OrganizationDataConsentRequestDetailsPaginatedList) GetTotalItemsOk() (*int64, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *OrganizationDataConsentRequestDetailsPaginatedList) SetTotalItems(v int64)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *OrganizationDataConsentRequestDetailsPaginatedList) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### GetItems

`func (o *OrganizationDataConsentRequestDetailsPaginatedList) GetItems() []OrganizationDataConsentRequestDetails`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *OrganizationDataConsentRequestDetailsPaginatedList) GetItemsOk() (*[]OrganizationDataConsentRequestDetails, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *OrganizationDataConsentRequestDetailsPaginatedList) SetItems(v []OrganizationDataConsentRequestDetails)`

SetItems sets Items field to given value.

### HasItems

`func (o *OrganizationDataConsentRequestDetailsPaginatedList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *OrganizationDataConsentRequestDetailsPaginatedList) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *OrganizationDataConsentRequestDetailsPaginatedList) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



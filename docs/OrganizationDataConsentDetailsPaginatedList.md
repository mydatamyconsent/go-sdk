# OrganizationDataConsentDetailsPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageIndex** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**TotalPages** | Pointer to **int32** |  | [optional] [readonly] 
**TotalItems** | Pointer to **int64** |  | [optional] 
**Items** | Pointer to [**[]OrganizationDataConsentDetails**](OrganizationDataConsentDetails.md) |  | [optional] 

## Methods

### NewOrganizationDataConsentDetailsPaginatedList

`func NewOrganizationDataConsentDetailsPaginatedList() *OrganizationDataConsentDetailsPaginatedList`

NewOrganizationDataConsentDetailsPaginatedList instantiates a new OrganizationDataConsentDetailsPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationDataConsentDetailsPaginatedListWithDefaults

`func NewOrganizationDataConsentDetailsPaginatedListWithDefaults() *OrganizationDataConsentDetailsPaginatedList`

NewOrganizationDataConsentDetailsPaginatedListWithDefaults instantiates a new OrganizationDataConsentDetailsPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageIndex

`func (o *OrganizationDataConsentDetailsPaginatedList) GetPageIndex() int32`

GetPageIndex returns the PageIndex field if non-nil, zero value otherwise.

### GetPageIndexOk

`func (o *OrganizationDataConsentDetailsPaginatedList) GetPageIndexOk() (*int32, bool)`

GetPageIndexOk returns a tuple with the PageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIndex

`func (o *OrganizationDataConsentDetailsPaginatedList) SetPageIndex(v int32)`

SetPageIndex sets PageIndex field to given value.

### HasPageIndex

`func (o *OrganizationDataConsentDetailsPaginatedList) HasPageIndex() bool`

HasPageIndex returns a boolean if a field has been set.

### GetPageSize

`func (o *OrganizationDataConsentDetailsPaginatedList) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *OrganizationDataConsentDetailsPaginatedList) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *OrganizationDataConsentDetailsPaginatedList) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *OrganizationDataConsentDetailsPaginatedList) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalPages

`func (o *OrganizationDataConsentDetailsPaginatedList) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *OrganizationDataConsentDetailsPaginatedList) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *OrganizationDataConsentDetailsPaginatedList) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *OrganizationDataConsentDetailsPaginatedList) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetTotalItems

`func (o *OrganizationDataConsentDetailsPaginatedList) GetTotalItems() int64`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *OrganizationDataConsentDetailsPaginatedList) GetTotalItemsOk() (*int64, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *OrganizationDataConsentDetailsPaginatedList) SetTotalItems(v int64)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *OrganizationDataConsentDetailsPaginatedList) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### GetItems

`func (o *OrganizationDataConsentDetailsPaginatedList) GetItems() []OrganizationDataConsentDetails`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *OrganizationDataConsentDetailsPaginatedList) GetItemsOk() (*[]OrganizationDataConsentDetails, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *OrganizationDataConsentDetailsPaginatedList) SetItems(v []OrganizationDataConsentDetails)`

SetItems sets Items field to given value.

### HasItems

`func (o *OrganizationDataConsentDetailsPaginatedList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *OrganizationDataConsentDetailsPaginatedList) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *OrganizationDataConsentDetailsPaginatedList) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



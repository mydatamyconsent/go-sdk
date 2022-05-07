# BillPaymentOrderItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Amount** | Pointer to **float64** |  | [optional] 

## Methods

### NewBillPaymentOrderItem

`func NewBillPaymentOrderItem() *BillPaymentOrderItem`

NewBillPaymentOrderItem instantiates a new BillPaymentOrderItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillPaymentOrderItemWithDefaults

`func NewBillPaymentOrderItemWithDefaults() *BillPaymentOrderItem`

NewBillPaymentOrderItemWithDefaults instantiates a new BillPaymentOrderItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BillPaymentOrderItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillPaymentOrderItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillPaymentOrderItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillPaymentOrderItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BillPaymentOrderItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BillPaymentOrderItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAmount

`func (o *BillPaymentOrderItem) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BillPaymentOrderItem) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BillPaymentOrderItem) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BillPaymentOrderItem) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



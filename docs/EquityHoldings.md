# EquityHoldings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Holding** | [**EquityHolding**](EquityHolding.md) |  | 
**Type** | [**EquityHoldingMode**](EquityHoldingMode.md) |  | 

## Methods

### NewEquityHoldings

`func NewEquityHoldings(holding EquityHolding, type_ EquityHoldingMode, ) *EquityHoldings`

NewEquityHoldings instantiates a new EquityHoldings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquityHoldingsWithDefaults

`func NewEquityHoldingsWithDefaults() *EquityHoldings`

NewEquityHoldingsWithDefaults instantiates a new EquityHoldings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHolding

`func (o *EquityHoldings) GetHolding() EquityHolding`

GetHolding returns the Holding field if non-nil, zero value otherwise.

### GetHoldingOk

`func (o *EquityHoldings) GetHoldingOk() (*EquityHolding, bool)`

GetHoldingOk returns a tuple with the Holding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolding

`func (o *EquityHoldings) SetHolding(v EquityHolding)`

SetHolding sets Holding field to given value.


### GetType

`func (o *EquityHoldings) GetType() EquityHoldingMode`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EquityHoldings) GetTypeOk() (*EquityHoldingMode, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EquityHoldings) SetType(v EquityHoldingMode)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



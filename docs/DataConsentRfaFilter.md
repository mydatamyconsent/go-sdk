# DataConsentRfaFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | Pointer to [**FilterType**](FilterType.md) |  | [optional] 
**Operator** | Pointer to [**Operator**](Operator.md) |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDataConsentRfaFilter

`func NewDataConsentRfaFilter() *DataConsentRfaFilter`

NewDataConsentRfaFilter instantiates a new DataConsentRfaFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataConsentRfaFilterWithDefaults

`func NewDataConsentRfaFilterWithDefaults() *DataConsentRfaFilter`

NewDataConsentRfaFilterWithDefaults instantiates a new DataConsentRfaFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *DataConsentRfaFilter) GetFilterType() FilterType`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *DataConsentRfaFilter) GetFilterTypeOk() (*FilterType, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *DataConsentRfaFilter) SetFilterType(v FilterType)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *DataConsentRfaFilter) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### GetOperator

`func (o *DataConsentRfaFilter) GetOperator() Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *DataConsentRfaFilter) GetOperatorOk() (*Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *DataConsentRfaFilter) SetOperator(v Operator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *DataConsentRfaFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValue

`func (o *DataConsentRfaFilter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DataConsentRfaFilter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DataConsentRfaFilter) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DataConsentRfaFilter) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *DataConsentRfaFilter) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *DataConsentRfaFilter) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



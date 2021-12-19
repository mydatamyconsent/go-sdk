# DataConsentRfaFilterDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | Pointer to [**FilterType**](FilterType.md) |  | [optional] 
**Operator** | Pointer to [**Operator**](Operator.md) |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDataConsentRfaFilterDto

`func NewDataConsentRfaFilterDto() *DataConsentRfaFilterDto`

NewDataConsentRfaFilterDto instantiates a new DataConsentRfaFilterDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataConsentRfaFilterDtoWithDefaults

`func NewDataConsentRfaFilterDtoWithDefaults() *DataConsentRfaFilterDto`

NewDataConsentRfaFilterDtoWithDefaults instantiates a new DataConsentRfaFilterDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *DataConsentRfaFilterDto) GetFilterType() FilterType`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *DataConsentRfaFilterDto) GetFilterTypeOk() (*FilterType, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *DataConsentRfaFilterDto) SetFilterType(v FilterType)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *DataConsentRfaFilterDto) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### GetOperator

`func (o *DataConsentRfaFilterDto) GetOperator() Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *DataConsentRfaFilterDto) GetOperatorOk() (*Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *DataConsentRfaFilterDto) SetOperator(v Operator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *DataConsentRfaFilterDto) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValue

`func (o *DataConsentRfaFilterDto) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DataConsentRfaFilterDto) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DataConsentRfaFilterDto) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DataConsentRfaFilterDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *DataConsentRfaFilterDto) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *DataConsentRfaFilterDto) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



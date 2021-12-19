# IdentifierStringKeyValuePair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to [**Identifier**](Identifier.md) |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIdentifierStringKeyValuePair

`func NewIdentifierStringKeyValuePair() *IdentifierStringKeyValuePair`

NewIdentifierStringKeyValuePair instantiates a new IdentifierStringKeyValuePair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentifierStringKeyValuePairWithDefaults

`func NewIdentifierStringKeyValuePairWithDefaults() *IdentifierStringKeyValuePair`

NewIdentifierStringKeyValuePairWithDefaults instantiates a new IdentifierStringKeyValuePair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *IdentifierStringKeyValuePair) GetKey() Identifier`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IdentifierStringKeyValuePair) GetKeyOk() (*Identifier, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IdentifierStringKeyValuePair) SetKey(v Identifier)`

SetKey sets Key field to given value.

### HasKey

`func (o *IdentifierStringKeyValuePair) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *IdentifierStringKeyValuePair) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IdentifierStringKeyValuePair) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IdentifierStringKeyValuePair) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IdentifierStringKeyValuePair) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *IdentifierStringKeyValuePair) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *IdentifierStringKeyValuePair) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



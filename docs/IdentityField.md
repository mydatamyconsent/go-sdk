# IdentityField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Field title. | 
**Description** | Pointer to **string** | Field description. | [optional] 
**Key** | **string** | Field key. | 
**DataType** | **string** | Field data type. | 

## Methods

### NewIdentityField

`func NewIdentityField(title string, key string, dataType string, ) *IdentityField`

NewIdentityField instantiates a new IdentityField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityFieldWithDefaults

`func NewIdentityFieldWithDefaults() *IdentityField`

NewIdentityFieldWithDefaults instantiates a new IdentityField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *IdentityField) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IdentityField) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IdentityField) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *IdentityField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IdentityField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IdentityField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IdentityField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *IdentityField) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IdentityField) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IdentityField) SetKey(v string)`

SetKey sets Key field to given value.


### GetDataType

`func (o *IdentityField) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *IdentityField) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *IdentityField) SetDataType(v string)`

SetDataType sets DataType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



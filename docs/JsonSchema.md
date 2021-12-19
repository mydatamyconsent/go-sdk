# JsonSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keywords** | Pointer to **[]map[string]interface{}** |  | [optional] [readonly] 
**OtherData** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewJsonSchema

`func NewJsonSchema() *JsonSchema`

NewJsonSchema instantiates a new JsonSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonSchemaWithDefaults

`func NewJsonSchemaWithDefaults() *JsonSchema`

NewJsonSchemaWithDefaults instantiates a new JsonSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeywords

`func (o *JsonSchema) GetKeywords() []map[string]interface{}`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *JsonSchema) GetKeywordsOk() (*[]map[string]interface{}, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *JsonSchema) SetKeywords(v []map[string]interface{})`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *JsonSchema) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### SetKeywordsNil

`func (o *JsonSchema) SetKeywordsNil(b bool)`

 SetKeywordsNil sets the value for Keywords to be an explicit nil

### UnsetKeywords
`func (o *JsonSchema) UnsetKeywords()`

UnsetKeywords ensures that no value is present for Keywords, not even an explicit nil
### GetOtherData

`func (o *JsonSchema) GetOtherData() map[string]interface{}`

GetOtherData returns the OtherData field if non-nil, zero value otherwise.

### GetOtherDataOk

`func (o *JsonSchema) GetOtherDataOk() (*map[string]interface{}, bool)`

GetOtherDataOk returns a tuple with the OtherData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherData

`func (o *JsonSchema) SetOtherData(v map[string]interface{})`

SetOtherData sets OtherData field to given value.

### HasOtherData

`func (o *JsonSchema) HasOtherData() bool`

HasOtherData returns a boolean if a field has been set.

### SetOtherDataNil

`func (o *JsonSchema) SetOtherDataNil(b bool)`

 SetOtherDataNil sets the value for OtherData to be an explicit nil

### UnsetOtherData
`func (o *JsonSchema) UnsetOtherData()`

UnsetOtherData ensures that no value is present for OtherData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



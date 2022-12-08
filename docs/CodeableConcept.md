# CodeableConcept

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Coding** | Pointer to [**[]Coding**](Coding.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 

## Methods

### NewCodeableConcept

`func NewCodeableConcept() *CodeableConcept`

NewCodeableConcept instantiates a new CodeableConcept object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeableConceptWithDefaults

`func NewCodeableConceptWithDefaults() *CodeableConcept`

NewCodeableConceptWithDefaults instantiates a new CodeableConcept object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoding

`func (o *CodeableConcept) GetCoding() []Coding`

GetCoding returns the Coding field if non-nil, zero value otherwise.

### GetCodingOk

`func (o *CodeableConcept) GetCodingOk() (*[]Coding, bool)`

GetCodingOk returns a tuple with the Coding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoding

`func (o *CodeableConcept) SetCoding(v []Coding)`

SetCoding sets Coding field to given value.

### HasCoding

`func (o *CodeableConcept) HasCoding() bool`

HasCoding returns a boolean if a field has been set.

### GetId

`func (o *CodeableConcept) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CodeableConcept) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CodeableConcept) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CodeableConcept) HasId() bool`

HasId returns a boolean if a field has been set.

### GetText

`func (o *CodeableConcept) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CodeableConcept) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CodeableConcept) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *CodeableConcept) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



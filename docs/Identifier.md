# Identifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Identifier key. EMAIL, MOBILE_NUMBER, etc. | 
**Name** | **string** | Identifier name. Email, Mobile Number, etc. | 
**Description** | **string** | Identifier description. User&#39;s email, User&#39;s mobile number, etc. | 
**ExampleValue** | **string** | Example value. example@email.com, +919090909090, etc. | 

## Methods

### NewIdentifier

`func NewIdentifier(key string, name string, description string, exampleValue string, ) *Identifier`

NewIdentifier instantiates a new Identifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentifierWithDefaults

`func NewIdentifierWithDefaults() *Identifier`

NewIdentifierWithDefaults instantiates a new Identifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *Identifier) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Identifier) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Identifier) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *Identifier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Identifier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Identifier) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Identifier) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Identifier) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Identifier) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetExampleValue

`func (o *Identifier) GetExampleValue() string`

GetExampleValue returns the ExampleValue field if non-nil, zero value otherwise.

### GetExampleValueOk

`func (o *Identifier) GetExampleValueOk() (*string, bool)`

GetExampleValueOk returns a tuple with the ExampleValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExampleValue

`func (o *Identifier) SetExampleValue(v string)`

SetExampleValue sets ExampleValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



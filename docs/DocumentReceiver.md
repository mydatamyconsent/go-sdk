# DocumentReceiver

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifiers** | [**[]StringStringKeyValuePair**](StringStringKeyValuePair.md) |  | 
**IdentificationStrategy** | [**IdentificationStrategy**](IdentificationStrategy.md) |  | 

## Methods

### NewDocumentReceiver

`func NewDocumentReceiver(identifiers []StringStringKeyValuePair, identificationStrategy IdentificationStrategy, ) *DocumentReceiver`

NewDocumentReceiver instantiates a new DocumentReceiver object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentReceiverWithDefaults

`func NewDocumentReceiverWithDefaults() *DocumentReceiver`

NewDocumentReceiverWithDefaults instantiates a new DocumentReceiver object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifiers

`func (o *DocumentReceiver) GetIdentifiers() []StringStringKeyValuePair`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *DocumentReceiver) GetIdentifiersOk() (*[]StringStringKeyValuePair, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *DocumentReceiver) SetIdentifiers(v []StringStringKeyValuePair)`

SetIdentifiers sets Identifiers field to given value.


### GetIdentificationStrategy

`func (o *DocumentReceiver) GetIdentificationStrategy() IdentificationStrategy`

GetIdentificationStrategy returns the IdentificationStrategy field if non-nil, zero value otherwise.

### GetIdentificationStrategyOk

`func (o *DocumentReceiver) GetIdentificationStrategyOk() (*IdentificationStrategy, bool)`

GetIdentificationStrategyOk returns a tuple with the IdentificationStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationStrategy

`func (o *DocumentReceiver) SetIdentificationStrategy(v IdentificationStrategy)`

SetIdentificationStrategy sets IdentificationStrategy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



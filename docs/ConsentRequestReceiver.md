# ConsentRequestReceiver

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifiers** | Pointer to [**[]StringStringKeyValuePair**](StringStringKeyValuePair.md) | Consent request receiver identifiers | [optional] 
**IdentificationStrategy** | Pointer to [**IdentificationStrategy**](IdentificationStrategy.md) |  | [optional] 

## Methods

### NewConsentRequestReceiver

`func NewConsentRequestReceiver() *ConsentRequestReceiver`

NewConsentRequestReceiver instantiates a new ConsentRequestReceiver object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsentRequestReceiverWithDefaults

`func NewConsentRequestReceiverWithDefaults() *ConsentRequestReceiver`

NewConsentRequestReceiverWithDefaults instantiates a new ConsentRequestReceiver object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifiers

`func (o *ConsentRequestReceiver) GetIdentifiers() []StringStringKeyValuePair`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *ConsentRequestReceiver) GetIdentifiersOk() (*[]StringStringKeyValuePair, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *ConsentRequestReceiver) SetIdentifiers(v []StringStringKeyValuePair)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *ConsentRequestReceiver) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### SetIdentifiersNil

`func (o *ConsentRequestReceiver) SetIdentifiersNil(b bool)`

 SetIdentifiersNil sets the value for Identifiers to be an explicit nil

### UnsetIdentifiers
`func (o *ConsentRequestReceiver) UnsetIdentifiers()`

UnsetIdentifiers ensures that no value is present for Identifiers, not even an explicit nil
### GetIdentificationStrategy

`func (o *ConsentRequestReceiver) GetIdentificationStrategy() IdentificationStrategy`

GetIdentificationStrategy returns the IdentificationStrategy field if non-nil, zero value otherwise.

### GetIdentificationStrategyOk

`func (o *ConsentRequestReceiver) GetIdentificationStrategyOk() (*IdentificationStrategy, bool)`

GetIdentificationStrategyOk returns a tuple with the IdentificationStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationStrategy

`func (o *ConsentRequestReceiver) SetIdentificationStrategy(v IdentificationStrategy)`

SetIdentificationStrategy sets IdentificationStrategy field to given value.

### HasIdentificationStrategy

`func (o *ConsentRequestReceiver) HasIdentificationStrategy() bool`

HasIdentificationStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



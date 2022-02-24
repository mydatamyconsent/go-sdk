# Receiver

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ReceiverType**](ReceiverType.md) |  | [optional] 
**Identifiers** | Pointer to [**[]StringStringKeyValuePair**](StringStringKeyValuePair.md) | Consent request receiver identifiers | [optional] 
**IdentificationStrategy** | Pointer to [**IdentificationStrategy**](IdentificationStrategy.md) |  | [optional] 

## Methods

### NewReceiver

`func NewReceiver() *Receiver`

NewReceiver instantiates a new Receiver object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiverWithDefaults

`func NewReceiverWithDefaults() *Receiver`

NewReceiverWithDefaults instantiates a new Receiver object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Receiver) GetType() ReceiverType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Receiver) GetTypeOk() (*ReceiverType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Receiver) SetType(v ReceiverType)`

SetType sets Type field to given value.

### HasType

`func (o *Receiver) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIdentifiers

`func (o *Receiver) GetIdentifiers() []StringStringKeyValuePair`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *Receiver) GetIdentifiersOk() (*[]StringStringKeyValuePair, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *Receiver) SetIdentifiers(v []StringStringKeyValuePair)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *Receiver) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### SetIdentifiersNil

`func (o *Receiver) SetIdentifiersNil(b bool)`

 SetIdentifiersNil sets the value for Identifiers to be an explicit nil

### UnsetIdentifiers
`func (o *Receiver) UnsetIdentifiers()`

UnsetIdentifiers ensures that no value is present for Identifiers, not even an explicit nil
### GetIdentificationStrategy

`func (o *Receiver) GetIdentificationStrategy() IdentificationStrategy`

GetIdentificationStrategy returns the IdentificationStrategy field if non-nil, zero value otherwise.

### GetIdentificationStrategyOk

`func (o *Receiver) GetIdentificationStrategyOk() (*IdentificationStrategy, bool)`

GetIdentificationStrategyOk returns a tuple with the IdentificationStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationStrategy

`func (o *Receiver) SetIdentificationStrategy(v IdentificationStrategy)`

SetIdentificationStrategy sets IdentificationStrategy field to given value.

### HasIdentificationStrategy

`func (o *Receiver) HasIdentificationStrategy() bool`

HasIdentificationStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DocumentIssueRequestReceiver

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryIso2Code** | **string** |  | 
**Identifiers** | [**[]KeyValuePair**](KeyValuePair.md) |  | 
**IdentificationStrategy** | [**IdentificationStrategy**](IdentificationStrategy.md) |  | 

## Methods

### NewDocumentIssueRequestReceiver

`func NewDocumentIssueRequestReceiver(countryIso2Code string, identifiers []KeyValuePair, identificationStrategy IdentificationStrategy, ) *DocumentIssueRequestReceiver`

NewDocumentIssueRequestReceiver instantiates a new DocumentIssueRequestReceiver object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentIssueRequestReceiverWithDefaults

`func NewDocumentIssueRequestReceiverWithDefaults() *DocumentIssueRequestReceiver`

NewDocumentIssueRequestReceiverWithDefaults instantiates a new DocumentIssueRequestReceiver object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryIso2Code

`func (o *DocumentIssueRequestReceiver) GetCountryIso2Code() string`

GetCountryIso2Code returns the CountryIso2Code field if non-nil, zero value otherwise.

### GetCountryIso2CodeOk

`func (o *DocumentIssueRequestReceiver) GetCountryIso2CodeOk() (*string, bool)`

GetCountryIso2CodeOk returns a tuple with the CountryIso2Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIso2Code

`func (o *DocumentIssueRequestReceiver) SetCountryIso2Code(v string)`

SetCountryIso2Code sets CountryIso2Code field to given value.


### GetIdentifiers

`func (o *DocumentIssueRequestReceiver) GetIdentifiers() []KeyValuePair`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *DocumentIssueRequestReceiver) GetIdentifiersOk() (*[]KeyValuePair, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *DocumentIssueRequestReceiver) SetIdentifiers(v []KeyValuePair)`

SetIdentifiers sets Identifiers field to given value.


### GetIdentificationStrategy

`func (o *DocumentIssueRequestReceiver) GetIdentificationStrategy() IdentificationStrategy`

GetIdentificationStrategy returns the IdentificationStrategy field if non-nil, zero value otherwise.

### GetIdentificationStrategyOk

`func (o *DocumentIssueRequestReceiver) GetIdentificationStrategyOk() (*IdentificationStrategy, bool)`

GetIdentificationStrategyOk returns a tuple with the IdentificationStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationStrategy

`func (o *DocumentIssueRequestReceiver) SetIdentificationStrategy(v IdentificationStrategy)`

SetIdentificationStrategy sets IdentificationStrategy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



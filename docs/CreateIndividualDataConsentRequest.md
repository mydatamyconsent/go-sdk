# CreateIndividualDataConsentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsentTemplateId** | Pointer to **string** | Consent template id | [optional] 
**Receiver** | [**Receiver**](Receiver.md) |  | 

## Methods

### NewCreateIndividualDataConsentRequest

`func NewCreateIndividualDataConsentRequest(receiver Receiver, ) *CreateIndividualDataConsentRequest`

NewCreateIndividualDataConsentRequest instantiates a new CreateIndividualDataConsentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIndividualDataConsentRequestWithDefaults

`func NewCreateIndividualDataConsentRequestWithDefaults() *CreateIndividualDataConsentRequest`

NewCreateIndividualDataConsentRequestWithDefaults instantiates a new CreateIndividualDataConsentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsentTemplateId

`func (o *CreateIndividualDataConsentRequest) GetConsentTemplateId() string`

GetConsentTemplateId returns the ConsentTemplateId field if non-nil, zero value otherwise.

### GetConsentTemplateIdOk

`func (o *CreateIndividualDataConsentRequest) GetConsentTemplateIdOk() (*string, bool)`

GetConsentTemplateIdOk returns a tuple with the ConsentTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentTemplateId

`func (o *CreateIndividualDataConsentRequest) SetConsentTemplateId(v string)`

SetConsentTemplateId sets ConsentTemplateId field to given value.

### HasConsentTemplateId

`func (o *CreateIndividualDataConsentRequest) HasConsentTemplateId() bool`

HasConsentTemplateId returns a boolean if a field has been set.

### GetReceiver

`func (o *CreateIndividualDataConsentRequest) GetReceiver() Receiver`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *CreateIndividualDataConsentRequest) GetReceiverOk() (*Receiver, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *CreateIndividualDataConsentRequest) SetReceiver(v Receiver)`

SetReceiver sets Receiver field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



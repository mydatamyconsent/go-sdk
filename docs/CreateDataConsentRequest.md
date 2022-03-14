# CreateDataConsentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsentTemplateId** | **string** | Consent template id | 
**Receiver** | [**ConsentRequestReceiver**](ConsentRequestReceiver.md) |  | 

## Methods

### NewCreateDataConsentRequest

`func NewCreateDataConsentRequest(consentTemplateId string, receiver ConsentRequestReceiver, ) *CreateDataConsentRequest`

NewCreateDataConsentRequest instantiates a new CreateDataConsentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDataConsentRequestWithDefaults

`func NewCreateDataConsentRequestWithDefaults() *CreateDataConsentRequest`

NewCreateDataConsentRequestWithDefaults instantiates a new CreateDataConsentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsentTemplateId

`func (o *CreateDataConsentRequest) GetConsentTemplateId() string`

GetConsentTemplateId returns the ConsentTemplateId field if non-nil, zero value otherwise.

### GetConsentTemplateIdOk

`func (o *CreateDataConsentRequest) GetConsentTemplateIdOk() (*string, bool)`

GetConsentTemplateIdOk returns a tuple with the ConsentTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentTemplateId

`func (o *CreateDataConsentRequest) SetConsentTemplateId(v string)`

SetConsentTemplateId sets ConsentTemplateId field to given value.


### GetReceiver

`func (o *CreateDataConsentRequest) GetReceiver() ConsentRequestReceiver`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *CreateDataConsentRequest) GetReceiverOk() (*ConsentRequestReceiver, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *CreateDataConsentRequest) SetReceiver(v ConsentRequestReceiver)`

SetReceiver sets Receiver field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateOrganizationDataConsentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsentTemplateId** | Pointer to **string** | Consent template id | [optional] 
**Receiver** | [**Receiver**](Receiver.md) |  | 

## Methods

### NewCreateOrganizationDataConsentRequest

`func NewCreateOrganizationDataConsentRequest(receiver Receiver, ) *CreateOrganizationDataConsentRequest`

NewCreateOrganizationDataConsentRequest instantiates a new CreateOrganizationDataConsentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationDataConsentRequestWithDefaults

`func NewCreateOrganizationDataConsentRequestWithDefaults() *CreateOrganizationDataConsentRequest`

NewCreateOrganizationDataConsentRequestWithDefaults instantiates a new CreateOrganizationDataConsentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsentTemplateId

`func (o *CreateOrganizationDataConsentRequest) GetConsentTemplateId() string`

GetConsentTemplateId returns the ConsentTemplateId field if non-nil, zero value otherwise.

### GetConsentTemplateIdOk

`func (o *CreateOrganizationDataConsentRequest) GetConsentTemplateIdOk() (*string, bool)`

GetConsentTemplateIdOk returns a tuple with the ConsentTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentTemplateId

`func (o *CreateOrganizationDataConsentRequest) SetConsentTemplateId(v string)`

SetConsentTemplateId sets ConsentTemplateId field to given value.

### HasConsentTemplateId

`func (o *CreateOrganizationDataConsentRequest) HasConsentTemplateId() bool`

HasConsentTemplateId returns a boolean if a field has been set.

### GetReceiver

`func (o *CreateOrganizationDataConsentRequest) GetReceiver() Receiver`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *CreateOrganizationDataConsentRequest) GetReceiverOk() (*Receiver, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *CreateOrganizationDataConsentRequest) SetReceiver(v Receiver)`

SetReceiver sets Receiver field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



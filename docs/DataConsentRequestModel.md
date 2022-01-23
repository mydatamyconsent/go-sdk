# DataConsentRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsentTemplateId** | Pointer to **string** |  | [optional] 
**StartDateTime** | Pointer to **time.Time** |  | [optional] 
**ExpiryDateTime** | Pointer to **time.Time** |  | [optional] 
**Receiver** | [**Receiver**](Receiver.md) |  | 

## Methods

### NewDataConsentRequestModel

`func NewDataConsentRequestModel(receiver Receiver, ) *DataConsentRequestModel`

NewDataConsentRequestModel instantiates a new DataConsentRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataConsentRequestModelWithDefaults

`func NewDataConsentRequestModelWithDefaults() *DataConsentRequestModel`

NewDataConsentRequestModelWithDefaults instantiates a new DataConsentRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsentTemplateId

`func (o *DataConsentRequestModel) GetConsentTemplateId() string`

GetConsentTemplateId returns the ConsentTemplateId field if non-nil, zero value otherwise.

### GetConsentTemplateIdOk

`func (o *DataConsentRequestModel) GetConsentTemplateIdOk() (*string, bool)`

GetConsentTemplateIdOk returns a tuple with the ConsentTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentTemplateId

`func (o *DataConsentRequestModel) SetConsentTemplateId(v string)`

SetConsentTemplateId sets ConsentTemplateId field to given value.

### HasConsentTemplateId

`func (o *DataConsentRequestModel) HasConsentTemplateId() bool`

HasConsentTemplateId returns a boolean if a field has been set.

### GetStartDateTime

`func (o *DataConsentRequestModel) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *DataConsentRequestModel) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *DataConsentRequestModel) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *DataConsentRequestModel) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetExpiryDateTime

`func (o *DataConsentRequestModel) GetExpiryDateTime() time.Time`

GetExpiryDateTime returns the ExpiryDateTime field if non-nil, zero value otherwise.

### GetExpiryDateTimeOk

`func (o *DataConsentRequestModel) GetExpiryDateTimeOk() (*time.Time, bool)`

GetExpiryDateTimeOk returns a tuple with the ExpiryDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDateTime

`func (o *DataConsentRequestModel) SetExpiryDateTime(v time.Time)`

SetExpiryDateTime sets ExpiryDateTime field to given value.

### HasExpiryDateTime

`func (o *DataConsentRequestModel) HasExpiryDateTime() bool`

HasExpiryDateTime returns a boolean if a field has been set.

### GetReceiver

`func (o *DataConsentRequestModel) GetReceiver() Receiver`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *DataConsentRequestModel) GetReceiverOk() (*Receiver, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *DataConsentRequestModel) SetReceiver(v Receiver)`

SetReceiver sets Receiver field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



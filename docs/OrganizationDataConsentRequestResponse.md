# OrganizationDataConsentRequestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**TemplateId** | Pointer to **string** |  | [optional] 
**RequestedAtUtc** | Pointer to **time.Time** |  | [optional] 
**RequestExpiresAtUtc** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to [**DataConsentStatus**](DataConsentStatus.md) |  | [optional] 
**TransactionId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewOrganizationDataConsentRequestResponse

`func NewOrganizationDataConsentRequestResponse() *OrganizationDataConsentRequestResponse`

NewOrganizationDataConsentRequestResponse instantiates a new OrganizationDataConsentRequestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationDataConsentRequestResponseWithDefaults

`func NewOrganizationDataConsentRequestResponseWithDefaults() *OrganizationDataConsentRequestResponse`

NewOrganizationDataConsentRequestResponseWithDefaults instantiates a new OrganizationDataConsentRequestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationDataConsentRequestResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationDataConsentRequestResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationDataConsentRequestResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationDataConsentRequestResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTemplateId

`func (o *OrganizationDataConsentRequestResponse) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *OrganizationDataConsentRequestResponse) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *OrganizationDataConsentRequestResponse) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *OrganizationDataConsentRequestResponse) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetRequestedAtUtc

`func (o *OrganizationDataConsentRequestResponse) GetRequestedAtUtc() time.Time`

GetRequestedAtUtc returns the RequestedAtUtc field if non-nil, zero value otherwise.

### GetRequestedAtUtcOk

`func (o *OrganizationDataConsentRequestResponse) GetRequestedAtUtcOk() (*time.Time, bool)`

GetRequestedAtUtcOk returns a tuple with the RequestedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAtUtc

`func (o *OrganizationDataConsentRequestResponse) SetRequestedAtUtc(v time.Time)`

SetRequestedAtUtc sets RequestedAtUtc field to given value.

### HasRequestedAtUtc

`func (o *OrganizationDataConsentRequestResponse) HasRequestedAtUtc() bool`

HasRequestedAtUtc returns a boolean if a field has been set.

### GetRequestExpiresAtUtc

`func (o *OrganizationDataConsentRequestResponse) GetRequestExpiresAtUtc() time.Time`

GetRequestExpiresAtUtc returns the RequestExpiresAtUtc field if non-nil, zero value otherwise.

### GetRequestExpiresAtUtcOk

`func (o *OrganizationDataConsentRequestResponse) GetRequestExpiresAtUtcOk() (*time.Time, bool)`

GetRequestExpiresAtUtcOk returns a tuple with the RequestExpiresAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestExpiresAtUtc

`func (o *OrganizationDataConsentRequestResponse) SetRequestExpiresAtUtc(v time.Time)`

SetRequestExpiresAtUtc sets RequestExpiresAtUtc field to given value.

### HasRequestExpiresAtUtc

`func (o *OrganizationDataConsentRequestResponse) HasRequestExpiresAtUtc() bool`

HasRequestExpiresAtUtc returns a boolean if a field has been set.

### GetStatus

`func (o *OrganizationDataConsentRequestResponse) GetStatus() DataConsentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrganizationDataConsentRequestResponse) GetStatusOk() (*DataConsentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrganizationDataConsentRequestResponse) SetStatus(v DataConsentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrganizationDataConsentRequestResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTransactionId

`func (o *OrganizationDataConsentRequestResponse) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *OrganizationDataConsentRequestResponse) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *OrganizationDataConsentRequestResponse) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *OrganizationDataConsentRequestResponse) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### SetTransactionIdNil

`func (o *OrganizationDataConsentRequestResponse) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *OrganizationDataConsentRequestResponse) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



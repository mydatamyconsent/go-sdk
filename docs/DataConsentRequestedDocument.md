# DataConsentRequestedDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Drn** | Pointer to **NullableString** |  | [optional] 
**FromDatetime** | Pointer to **time.Time** |  | [optional] 
**ToDatetime** | Pointer to **time.Time** |  | [optional] 
**ProviderId** | Pointer to **string** |  | [optional] 
**DocumentTypeId** | Pointer to **NullableString** |  | [optional] 
**DocumentIdentifier** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDataConsentRequestedDocument

`func NewDataConsentRequestedDocument() *DataConsentRequestedDocument`

NewDataConsentRequestedDocument instantiates a new DataConsentRequestedDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataConsentRequestedDocumentWithDefaults

`func NewDataConsentRequestedDocumentWithDefaults() *DataConsentRequestedDocument`

NewDataConsentRequestedDocumentWithDefaults instantiates a new DataConsentRequestedDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrn

`func (o *DataConsentRequestedDocument) GetDrn() string`

GetDrn returns the Drn field if non-nil, zero value otherwise.

### GetDrnOk

`func (o *DataConsentRequestedDocument) GetDrnOk() (*string, bool)`

GetDrnOk returns a tuple with the Drn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrn

`func (o *DataConsentRequestedDocument) SetDrn(v string)`

SetDrn sets Drn field to given value.

### HasDrn

`func (o *DataConsentRequestedDocument) HasDrn() bool`

HasDrn returns a boolean if a field has been set.

### SetDrnNil

`func (o *DataConsentRequestedDocument) SetDrnNil(b bool)`

 SetDrnNil sets the value for Drn to be an explicit nil

### UnsetDrn
`func (o *DataConsentRequestedDocument) UnsetDrn()`

UnsetDrn ensures that no value is present for Drn, not even an explicit nil
### GetFromDatetime

`func (o *DataConsentRequestedDocument) GetFromDatetime() time.Time`

GetFromDatetime returns the FromDatetime field if non-nil, zero value otherwise.

### GetFromDatetimeOk

`func (o *DataConsentRequestedDocument) GetFromDatetimeOk() (*time.Time, bool)`

GetFromDatetimeOk returns a tuple with the FromDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDatetime

`func (o *DataConsentRequestedDocument) SetFromDatetime(v time.Time)`

SetFromDatetime sets FromDatetime field to given value.

### HasFromDatetime

`func (o *DataConsentRequestedDocument) HasFromDatetime() bool`

HasFromDatetime returns a boolean if a field has been set.

### GetToDatetime

`func (o *DataConsentRequestedDocument) GetToDatetime() time.Time`

GetToDatetime returns the ToDatetime field if non-nil, zero value otherwise.

### GetToDatetimeOk

`func (o *DataConsentRequestedDocument) GetToDatetimeOk() (*time.Time, bool)`

GetToDatetimeOk returns a tuple with the ToDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDatetime

`func (o *DataConsentRequestedDocument) SetToDatetime(v time.Time)`

SetToDatetime sets ToDatetime field to given value.

### HasToDatetime

`func (o *DataConsentRequestedDocument) HasToDatetime() bool`

HasToDatetime returns a boolean if a field has been set.

### GetProviderId

`func (o *DataConsentRequestedDocument) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *DataConsentRequestedDocument) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *DataConsentRequestedDocument) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *DataConsentRequestedDocument) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetDocumentTypeId

`func (o *DataConsentRequestedDocument) GetDocumentTypeId() string`

GetDocumentTypeId returns the DocumentTypeId field if non-nil, zero value otherwise.

### GetDocumentTypeIdOk

`func (o *DataConsentRequestedDocument) GetDocumentTypeIdOk() (*string, bool)`

GetDocumentTypeIdOk returns a tuple with the DocumentTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTypeId

`func (o *DataConsentRequestedDocument) SetDocumentTypeId(v string)`

SetDocumentTypeId sets DocumentTypeId field to given value.

### HasDocumentTypeId

`func (o *DataConsentRequestedDocument) HasDocumentTypeId() bool`

HasDocumentTypeId returns a boolean if a field has been set.

### SetDocumentTypeIdNil

`func (o *DataConsentRequestedDocument) SetDocumentTypeIdNil(b bool)`

 SetDocumentTypeIdNil sets the value for DocumentTypeId to be an explicit nil

### UnsetDocumentTypeId
`func (o *DataConsentRequestedDocument) UnsetDocumentTypeId()`

UnsetDocumentTypeId ensures that no value is present for DocumentTypeId, not even an explicit nil
### GetDocumentIdentifier

`func (o *DataConsentRequestedDocument) GetDocumentIdentifier() string`

GetDocumentIdentifier returns the DocumentIdentifier field if non-nil, zero value otherwise.

### GetDocumentIdentifierOk

`func (o *DataConsentRequestedDocument) GetDocumentIdentifierOk() (*string, bool)`

GetDocumentIdentifierOk returns a tuple with the DocumentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIdentifier

`func (o *DataConsentRequestedDocument) SetDocumentIdentifier(v string)`

SetDocumentIdentifier sets DocumentIdentifier field to given value.

### HasDocumentIdentifier

`func (o *DataConsentRequestedDocument) HasDocumentIdentifier() bool`

HasDocumentIdentifier returns a boolean if a field has been set.

### SetDocumentIdentifierNil

`func (o *DataConsentRequestedDocument) SetDocumentIdentifierNil(b bool)`

 SetDocumentIdentifierNil sets the value for DocumentIdentifier to be an explicit nil

### UnsetDocumentIdentifier
`func (o *DataConsentRequestedDocument) UnsetDocumentIdentifier()`

UnsetDocumentIdentifier ensures that no value is present for DocumentIdentifier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



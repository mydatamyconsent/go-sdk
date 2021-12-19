# DataConsentRequestedDocumentDto

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

### NewDataConsentRequestedDocumentDto

`func NewDataConsentRequestedDocumentDto() *DataConsentRequestedDocumentDto`

NewDataConsentRequestedDocumentDto instantiates a new DataConsentRequestedDocumentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataConsentRequestedDocumentDtoWithDefaults

`func NewDataConsentRequestedDocumentDtoWithDefaults() *DataConsentRequestedDocumentDto`

NewDataConsentRequestedDocumentDtoWithDefaults instantiates a new DataConsentRequestedDocumentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrn

`func (o *DataConsentRequestedDocumentDto) GetDrn() string`

GetDrn returns the Drn field if non-nil, zero value otherwise.

### GetDrnOk

`func (o *DataConsentRequestedDocumentDto) GetDrnOk() (*string, bool)`

GetDrnOk returns a tuple with the Drn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrn

`func (o *DataConsentRequestedDocumentDto) SetDrn(v string)`

SetDrn sets Drn field to given value.

### HasDrn

`func (o *DataConsentRequestedDocumentDto) HasDrn() bool`

HasDrn returns a boolean if a field has been set.

### SetDrnNil

`func (o *DataConsentRequestedDocumentDto) SetDrnNil(b bool)`

 SetDrnNil sets the value for Drn to be an explicit nil

### UnsetDrn
`func (o *DataConsentRequestedDocumentDto) UnsetDrn()`

UnsetDrn ensures that no value is present for Drn, not even an explicit nil
### GetFromDatetime

`func (o *DataConsentRequestedDocumentDto) GetFromDatetime() time.Time`

GetFromDatetime returns the FromDatetime field if non-nil, zero value otherwise.

### GetFromDatetimeOk

`func (o *DataConsentRequestedDocumentDto) GetFromDatetimeOk() (*time.Time, bool)`

GetFromDatetimeOk returns a tuple with the FromDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDatetime

`func (o *DataConsentRequestedDocumentDto) SetFromDatetime(v time.Time)`

SetFromDatetime sets FromDatetime field to given value.

### HasFromDatetime

`func (o *DataConsentRequestedDocumentDto) HasFromDatetime() bool`

HasFromDatetime returns a boolean if a field has been set.

### GetToDatetime

`func (o *DataConsentRequestedDocumentDto) GetToDatetime() time.Time`

GetToDatetime returns the ToDatetime field if non-nil, zero value otherwise.

### GetToDatetimeOk

`func (o *DataConsentRequestedDocumentDto) GetToDatetimeOk() (*time.Time, bool)`

GetToDatetimeOk returns a tuple with the ToDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDatetime

`func (o *DataConsentRequestedDocumentDto) SetToDatetime(v time.Time)`

SetToDatetime sets ToDatetime field to given value.

### HasToDatetime

`func (o *DataConsentRequestedDocumentDto) HasToDatetime() bool`

HasToDatetime returns a boolean if a field has been set.

### GetProviderId

`func (o *DataConsentRequestedDocumentDto) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *DataConsentRequestedDocumentDto) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *DataConsentRequestedDocumentDto) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *DataConsentRequestedDocumentDto) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetDocumentTypeId

`func (o *DataConsentRequestedDocumentDto) GetDocumentTypeId() string`

GetDocumentTypeId returns the DocumentTypeId field if non-nil, zero value otherwise.

### GetDocumentTypeIdOk

`func (o *DataConsentRequestedDocumentDto) GetDocumentTypeIdOk() (*string, bool)`

GetDocumentTypeIdOk returns a tuple with the DocumentTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTypeId

`func (o *DataConsentRequestedDocumentDto) SetDocumentTypeId(v string)`

SetDocumentTypeId sets DocumentTypeId field to given value.

### HasDocumentTypeId

`func (o *DataConsentRequestedDocumentDto) HasDocumentTypeId() bool`

HasDocumentTypeId returns a boolean if a field has been set.

### SetDocumentTypeIdNil

`func (o *DataConsentRequestedDocumentDto) SetDocumentTypeIdNil(b bool)`

 SetDocumentTypeIdNil sets the value for DocumentTypeId to be an explicit nil

### UnsetDocumentTypeId
`func (o *DataConsentRequestedDocumentDto) UnsetDocumentTypeId()`

UnsetDocumentTypeId ensures that no value is present for DocumentTypeId, not even an explicit nil
### GetDocumentIdentifier

`func (o *DataConsentRequestedDocumentDto) GetDocumentIdentifier() string`

GetDocumentIdentifier returns the DocumentIdentifier field if non-nil, zero value otherwise.

### GetDocumentIdentifierOk

`func (o *DataConsentRequestedDocumentDto) GetDocumentIdentifierOk() (*string, bool)`

GetDocumentIdentifierOk returns a tuple with the DocumentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIdentifier

`func (o *DataConsentRequestedDocumentDto) SetDocumentIdentifier(v string)`

SetDocumentIdentifier sets DocumentIdentifier field to given value.

### HasDocumentIdentifier

`func (o *DataConsentRequestedDocumentDto) HasDocumentIdentifier() bool`

HasDocumentIdentifier returns a boolean if a field has been set.

### SetDocumentIdentifierNil

`func (o *DataConsentRequestedDocumentDto) SetDocumentIdentifierNil(b bool)`

 SetDocumentIdentifierNil sets the value for DocumentIdentifier to be an explicit nil

### UnsetDocumentIdentifier
`func (o *DataConsentRequestedDocumentDto) UnsetDocumentIdentifier()`

UnsetDocumentIdentifier ensures that no value is present for DocumentIdentifier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



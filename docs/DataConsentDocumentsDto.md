# DataConsentDocumentsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Documents** | Pointer to [**[]Document**](Document.md) |  | [optional] 
**ApprovedDocuments** | Pointer to [**[]DataConsentRequestedDocument**](DataConsentRequestedDocument.md) |  | [optional] 

## Methods

### NewDataConsentDocumentsDto

`func NewDataConsentDocumentsDto() *DataConsentDocumentsDto`

NewDataConsentDocumentsDto instantiates a new DataConsentDocumentsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataConsentDocumentsDtoWithDefaults

`func NewDataConsentDocumentsDtoWithDefaults() *DataConsentDocumentsDto`

NewDataConsentDocumentsDtoWithDefaults instantiates a new DataConsentDocumentsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataConsentDocumentsDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataConsentDocumentsDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataConsentDocumentsDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataConsentDocumentsDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDocuments

`func (o *DataConsentDocumentsDto) GetDocuments() []Document`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *DataConsentDocumentsDto) GetDocumentsOk() (*[]Document, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *DataConsentDocumentsDto) SetDocuments(v []Document)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *DataConsentDocumentsDto) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### SetDocumentsNil

`func (o *DataConsentDocumentsDto) SetDocumentsNil(b bool)`

 SetDocumentsNil sets the value for Documents to be an explicit nil

### UnsetDocuments
`func (o *DataConsentDocumentsDto) UnsetDocuments()`

UnsetDocuments ensures that no value is present for Documents, not even an explicit nil
### GetApprovedDocuments

`func (o *DataConsentDocumentsDto) GetApprovedDocuments() []DataConsentRequestedDocument`

GetApprovedDocuments returns the ApprovedDocuments field if non-nil, zero value otherwise.

### GetApprovedDocumentsOk

`func (o *DataConsentDocumentsDto) GetApprovedDocumentsOk() (*[]DataConsentRequestedDocument, bool)`

GetApprovedDocumentsOk returns a tuple with the ApprovedDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedDocuments

`func (o *DataConsentDocumentsDto) SetApprovedDocuments(v []DataConsentRequestedDocument)`

SetApprovedDocuments sets ApprovedDocuments field to given value.

### HasApprovedDocuments

`func (o *DataConsentDocumentsDto) HasApprovedDocuments() bool`

HasApprovedDocuments returns a boolean if a field has been set.

### SetApprovedDocumentsNil

`func (o *DataConsentDocumentsDto) SetApprovedDocumentsNil(b bool)`

 SetApprovedDocumentsNil sets the value for ApprovedDocuments to be an explicit nil

### UnsetApprovedDocuments
`func (o *DataConsentDocumentsDto) UnsetApprovedDocuments()`

UnsetApprovedDocuments ensures that no value is present for ApprovedDocuments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



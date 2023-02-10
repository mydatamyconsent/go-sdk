# DataProcessingAgreement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Agreement id. | 
**Name** | **string** | Agreement name. | 
**IssuerType** | [**IssuerType**](IssuerType.md) |  | 
**AgreementUrl** | **string** | Agreement attachment file URL. | 

## Methods

### NewDataProcessingAgreement

`func NewDataProcessingAgreement(id string, name string, issuerType IssuerType, agreementUrl string, ) *DataProcessingAgreement`

NewDataProcessingAgreement instantiates a new DataProcessingAgreement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataProcessingAgreementWithDefaults

`func NewDataProcessingAgreementWithDefaults() *DataProcessingAgreement`

NewDataProcessingAgreementWithDefaults instantiates a new DataProcessingAgreement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataProcessingAgreement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataProcessingAgreement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataProcessingAgreement) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DataProcessingAgreement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataProcessingAgreement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataProcessingAgreement) SetName(v string)`

SetName sets Name field to given value.


### GetIssuerType

`func (o *DataProcessingAgreement) GetIssuerType() IssuerType`

GetIssuerType returns the IssuerType field if non-nil, zero value otherwise.

### GetIssuerTypeOk

`func (o *DataProcessingAgreement) GetIssuerTypeOk() (*IssuerType, bool)`

GetIssuerTypeOk returns a tuple with the IssuerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerType

`func (o *DataProcessingAgreement) SetIssuerType(v IssuerType)`

SetIssuerType sets IssuerType field to given value.


### GetAgreementUrl

`func (o *DataProcessingAgreement) GetAgreementUrl() string`

GetAgreementUrl returns the AgreementUrl field if non-nil, zero value otherwise.

### GetAgreementUrlOk

`func (o *DataProcessingAgreement) GetAgreementUrlOk() (*string, bool)`

GetAgreementUrlOk returns a tuple with the AgreementUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementUrl

`func (o *DataProcessingAgreement) SetAgreementUrl(v string)`

SetAgreementUrl sets AgreementUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



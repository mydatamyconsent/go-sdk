# CreateDataProcessingAgreement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Agreement version. Agreement body content. | 
**IssuerType** | [**IssuerType**](IssuerType.md) |  | 
**AgreementUrl** | **string** | Agreement attachment file URL. | 

## Methods

### NewCreateDataProcessingAgreement

`func NewCreateDataProcessingAgreement(name string, issuerType IssuerType, agreementUrl string, ) *CreateDataProcessingAgreement`

NewCreateDataProcessingAgreement instantiates a new CreateDataProcessingAgreement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDataProcessingAgreementWithDefaults

`func NewCreateDataProcessingAgreementWithDefaults() *CreateDataProcessingAgreement`

NewCreateDataProcessingAgreementWithDefaults instantiates a new CreateDataProcessingAgreement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateDataProcessingAgreement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDataProcessingAgreement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDataProcessingAgreement) SetName(v string)`

SetName sets Name field to given value.


### GetIssuerType

`func (o *CreateDataProcessingAgreement) GetIssuerType() IssuerType`

GetIssuerType returns the IssuerType field if non-nil, zero value otherwise.

### GetIssuerTypeOk

`func (o *CreateDataProcessingAgreement) GetIssuerTypeOk() (*IssuerType, bool)`

GetIssuerTypeOk returns a tuple with the IssuerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerType

`func (o *CreateDataProcessingAgreement) SetIssuerType(v IssuerType)`

SetIssuerType sets IssuerType field to given value.


### GetAgreementUrl

`func (o *CreateDataProcessingAgreement) GetAgreementUrl() string`

GetAgreementUrl returns the AgreementUrl field if non-nil, zero value otherwise.

### GetAgreementUrlOk

`func (o *CreateDataProcessingAgreement) GetAgreementUrlOk() (*string, bool)`

GetAgreementUrlOk returns a tuple with the AgreementUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementUrl

`func (o *CreateDataProcessingAgreement) SetAgreementUrl(v string)`

SetAgreementUrl sets AgreementUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



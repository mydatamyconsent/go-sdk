# ConsentedFinancialAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Financial account id. | 
**Name** | **string** | Financial account name. | 
**Category** | [**FinancialAccountCategoryType**](FinancialAccountCategoryType.md) |  | 
**SubCategory** | [**FinancialAccountSubCategoryType**](FinancialAccountSubCategoryType.md) |  | 
**Identifier** | **string** | Financial account identifier. | 
**IssuerId** | **string** | Financial account issuer id. | 
**IssuerName** | **string** | Financial account issuer name. | 

## Methods

### NewConsentedFinancialAccount

`func NewConsentedFinancialAccount(id string, name string, category FinancialAccountCategoryType, subCategory FinancialAccountSubCategoryType, identifier string, issuerId string, issuerName string, ) *ConsentedFinancialAccount`

NewConsentedFinancialAccount instantiates a new ConsentedFinancialAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsentedFinancialAccountWithDefaults

`func NewConsentedFinancialAccountWithDefaults() *ConsentedFinancialAccount`

NewConsentedFinancialAccountWithDefaults instantiates a new ConsentedFinancialAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConsentedFinancialAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsentedFinancialAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsentedFinancialAccount) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ConsentedFinancialAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsentedFinancialAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsentedFinancialAccount) SetName(v string)`

SetName sets Name field to given value.


### GetCategory

`func (o *ConsentedFinancialAccount) GetCategory() FinancialAccountCategoryType`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ConsentedFinancialAccount) GetCategoryOk() (*FinancialAccountCategoryType, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ConsentedFinancialAccount) SetCategory(v FinancialAccountCategoryType)`

SetCategory sets Category field to given value.


### GetSubCategory

`func (o *ConsentedFinancialAccount) GetSubCategory() FinancialAccountSubCategoryType`

GetSubCategory returns the SubCategory field if non-nil, zero value otherwise.

### GetSubCategoryOk

`func (o *ConsentedFinancialAccount) GetSubCategoryOk() (*FinancialAccountSubCategoryType, bool)`

GetSubCategoryOk returns a tuple with the SubCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCategory

`func (o *ConsentedFinancialAccount) SetSubCategory(v FinancialAccountSubCategoryType)`

SetSubCategory sets SubCategory field to given value.


### GetIdentifier

`func (o *ConsentedFinancialAccount) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ConsentedFinancialAccount) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ConsentedFinancialAccount) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetIssuerId

`func (o *ConsentedFinancialAccount) GetIssuerId() string`

GetIssuerId returns the IssuerId field if non-nil, zero value otherwise.

### GetIssuerIdOk

`func (o *ConsentedFinancialAccount) GetIssuerIdOk() (*string, bool)`

GetIssuerIdOk returns a tuple with the IssuerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerId

`func (o *ConsentedFinancialAccount) SetIssuerId(v string)`

SetIssuerId sets IssuerId field to given value.


### GetIssuerName

`func (o *ConsentedFinancialAccount) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *ConsentedFinancialAccount) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *ConsentedFinancialAccount) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



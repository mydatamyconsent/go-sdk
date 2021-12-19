# OrganizationFinancialAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**BeneficiaryName** | Pointer to **NullableString** |  | [optional] 
**AccountNumber** | Pointer to **NullableString** |  | [optional] 
**RoutingNumber** | Pointer to **NullableString** |  | [optional] 
**Organization** | Pointer to [**Organization**](Organization.md) |  | [optional] 
**IsPrimary** | Pointer to **bool** |  | [optional] 
**IsVerified** | Pointer to **bool** |  | [optional] 
**LogoUrl** | Pointer to **NullableString** |  | [optional] 
**BankName** | Pointer to **NullableString** |  | [optional] 
**BankAccountProofUrl** | Pointer to **NullableString** |  | [optional] 
**FileType** | Pointer to [**FileType**](FileType.md) |  | [optional] 
**BankAccountType** | Pointer to [**BankAccountType**](BankAccountType.md) |  | [optional] 

## Methods

### NewOrganizationFinancialAccount

`func NewOrganizationFinancialAccount() *OrganizationFinancialAccount`

NewOrganizationFinancialAccount instantiates a new OrganizationFinancialAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationFinancialAccountWithDefaults

`func NewOrganizationFinancialAccountWithDefaults() *OrganizationFinancialAccount`

NewOrganizationFinancialAccountWithDefaults instantiates a new OrganizationFinancialAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationFinancialAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationFinancialAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationFinancialAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationFinancialAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *OrganizationFinancialAccount) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OrganizationFinancialAccount) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OrganizationFinancialAccount) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *OrganizationFinancialAccount) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetBeneficiaryName

`func (o *OrganizationFinancialAccount) GetBeneficiaryName() string`

GetBeneficiaryName returns the BeneficiaryName field if non-nil, zero value otherwise.

### GetBeneficiaryNameOk

`func (o *OrganizationFinancialAccount) GetBeneficiaryNameOk() (*string, bool)`

GetBeneficiaryNameOk returns a tuple with the BeneficiaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryName

`func (o *OrganizationFinancialAccount) SetBeneficiaryName(v string)`

SetBeneficiaryName sets BeneficiaryName field to given value.

### HasBeneficiaryName

`func (o *OrganizationFinancialAccount) HasBeneficiaryName() bool`

HasBeneficiaryName returns a boolean if a field has been set.

### SetBeneficiaryNameNil

`func (o *OrganizationFinancialAccount) SetBeneficiaryNameNil(b bool)`

 SetBeneficiaryNameNil sets the value for BeneficiaryName to be an explicit nil

### UnsetBeneficiaryName
`func (o *OrganizationFinancialAccount) UnsetBeneficiaryName()`

UnsetBeneficiaryName ensures that no value is present for BeneficiaryName, not even an explicit nil
### GetAccountNumber

`func (o *OrganizationFinancialAccount) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *OrganizationFinancialAccount) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *OrganizationFinancialAccount) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *OrganizationFinancialAccount) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### SetAccountNumberNil

`func (o *OrganizationFinancialAccount) SetAccountNumberNil(b bool)`

 SetAccountNumberNil sets the value for AccountNumber to be an explicit nil

### UnsetAccountNumber
`func (o *OrganizationFinancialAccount) UnsetAccountNumber()`

UnsetAccountNumber ensures that no value is present for AccountNumber, not even an explicit nil
### GetRoutingNumber

`func (o *OrganizationFinancialAccount) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *OrganizationFinancialAccount) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *OrganizationFinancialAccount) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.

### HasRoutingNumber

`func (o *OrganizationFinancialAccount) HasRoutingNumber() bool`

HasRoutingNumber returns a boolean if a field has been set.

### SetRoutingNumberNil

`func (o *OrganizationFinancialAccount) SetRoutingNumberNil(b bool)`

 SetRoutingNumberNil sets the value for RoutingNumber to be an explicit nil

### UnsetRoutingNumber
`func (o *OrganizationFinancialAccount) UnsetRoutingNumber()`

UnsetRoutingNumber ensures that no value is present for RoutingNumber, not even an explicit nil
### GetOrganization

`func (o *OrganizationFinancialAccount) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *OrganizationFinancialAccount) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *OrganizationFinancialAccount) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *OrganizationFinancialAccount) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetIsPrimary

`func (o *OrganizationFinancialAccount) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *OrganizationFinancialAccount) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *OrganizationFinancialAccount) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *OrganizationFinancialAccount) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetIsVerified

`func (o *OrganizationFinancialAccount) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *OrganizationFinancialAccount) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *OrganizationFinancialAccount) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *OrganizationFinancialAccount) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetLogoUrl

`func (o *OrganizationFinancialAccount) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *OrganizationFinancialAccount) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *OrganizationFinancialAccount) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *OrganizationFinancialAccount) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### SetLogoUrlNil

`func (o *OrganizationFinancialAccount) SetLogoUrlNil(b bool)`

 SetLogoUrlNil sets the value for LogoUrl to be an explicit nil

### UnsetLogoUrl
`func (o *OrganizationFinancialAccount) UnsetLogoUrl()`

UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
### GetBankName

`func (o *OrganizationFinancialAccount) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *OrganizationFinancialAccount) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *OrganizationFinancialAccount) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *OrganizationFinancialAccount) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### SetBankNameNil

`func (o *OrganizationFinancialAccount) SetBankNameNil(b bool)`

 SetBankNameNil sets the value for BankName to be an explicit nil

### UnsetBankName
`func (o *OrganizationFinancialAccount) UnsetBankName()`

UnsetBankName ensures that no value is present for BankName, not even an explicit nil
### GetBankAccountProofUrl

`func (o *OrganizationFinancialAccount) GetBankAccountProofUrl() string`

GetBankAccountProofUrl returns the BankAccountProofUrl field if non-nil, zero value otherwise.

### GetBankAccountProofUrlOk

`func (o *OrganizationFinancialAccount) GetBankAccountProofUrlOk() (*string, bool)`

GetBankAccountProofUrlOk returns a tuple with the BankAccountProofUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountProofUrl

`func (o *OrganizationFinancialAccount) SetBankAccountProofUrl(v string)`

SetBankAccountProofUrl sets BankAccountProofUrl field to given value.

### HasBankAccountProofUrl

`func (o *OrganizationFinancialAccount) HasBankAccountProofUrl() bool`

HasBankAccountProofUrl returns a boolean if a field has been set.

### SetBankAccountProofUrlNil

`func (o *OrganizationFinancialAccount) SetBankAccountProofUrlNil(b bool)`

 SetBankAccountProofUrlNil sets the value for BankAccountProofUrl to be an explicit nil

### UnsetBankAccountProofUrl
`func (o *OrganizationFinancialAccount) UnsetBankAccountProofUrl()`

UnsetBankAccountProofUrl ensures that no value is present for BankAccountProofUrl, not even an explicit nil
### GetFileType

`func (o *OrganizationFinancialAccount) GetFileType() FileType`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *OrganizationFinancialAccount) GetFileTypeOk() (*FileType, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *OrganizationFinancialAccount) SetFileType(v FileType)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *OrganizationFinancialAccount) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### GetBankAccountType

`func (o *OrganizationFinancialAccount) GetBankAccountType() BankAccountType`

GetBankAccountType returns the BankAccountType field if non-nil, zero value otherwise.

### GetBankAccountTypeOk

`func (o *OrganizationFinancialAccount) GetBankAccountTypeOk() (*BankAccountType, bool)`

GetBankAccountTypeOk returns a tuple with the BankAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountType

`func (o *OrganizationFinancialAccount) SetBankAccountType(v BankAccountType)`

SetBankAccountType sets BankAccountType field to given value.

### HasBankAccountType

`func (o *OrganizationFinancialAccount) HasBankAccountType() bool`

HasBankAccountType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



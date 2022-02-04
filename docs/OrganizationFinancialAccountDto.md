# OrganizationFinancialAccountDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**BeneficiaryName** | Pointer to **NullableString** |  | [optional] 
**AccountNumber** | Pointer to **NullableString** |  | [optional] 
**RoutingNumber** | Pointer to **NullableString** |  | [optional] 
**IsPrimary** | Pointer to **bool** |  | [optional] 
**IsVerified** | Pointer to **bool** |  | [optional] 
**LogoUrl** | Pointer to **NullableString** |  | [optional] 
**BankName** | Pointer to **NullableString** |  | [optional] 
**BankAccountType** | Pointer to [**BankAccountType**](BankAccountType.md) |  | [optional] 
**BankAccountProofUrl** | Pointer to **NullableString** |  | [optional] 
**FileType** | Pointer to [**FileType**](FileType.md) |  | [optional] 

## Methods

### NewOrganizationFinancialAccountDto

`func NewOrganizationFinancialAccountDto() *OrganizationFinancialAccountDto`

NewOrganizationFinancialAccountDto instantiates a new OrganizationFinancialAccountDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationFinancialAccountDtoWithDefaults

`func NewOrganizationFinancialAccountDtoWithDefaults() *OrganizationFinancialAccountDto`

NewOrganizationFinancialAccountDtoWithDefaults instantiates a new OrganizationFinancialAccountDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationFinancialAccountDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationFinancialAccountDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationFinancialAccountDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationFinancialAccountDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *OrganizationFinancialAccountDto) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OrganizationFinancialAccountDto) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OrganizationFinancialAccountDto) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *OrganizationFinancialAccountDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOrganizationName

`func (o *OrganizationFinancialAccountDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *OrganizationFinancialAccountDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *OrganizationFinancialAccountDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *OrganizationFinancialAccountDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *OrganizationFinancialAccountDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *OrganizationFinancialAccountDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetBeneficiaryName

`func (o *OrganizationFinancialAccountDto) GetBeneficiaryName() string`

GetBeneficiaryName returns the BeneficiaryName field if non-nil, zero value otherwise.

### GetBeneficiaryNameOk

`func (o *OrganizationFinancialAccountDto) GetBeneficiaryNameOk() (*string, bool)`

GetBeneficiaryNameOk returns a tuple with the BeneficiaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryName

`func (o *OrganizationFinancialAccountDto) SetBeneficiaryName(v string)`

SetBeneficiaryName sets BeneficiaryName field to given value.

### HasBeneficiaryName

`func (o *OrganizationFinancialAccountDto) HasBeneficiaryName() bool`

HasBeneficiaryName returns a boolean if a field has been set.

### SetBeneficiaryNameNil

`func (o *OrganizationFinancialAccountDto) SetBeneficiaryNameNil(b bool)`

 SetBeneficiaryNameNil sets the value for BeneficiaryName to be an explicit nil

### UnsetBeneficiaryName
`func (o *OrganizationFinancialAccountDto) UnsetBeneficiaryName()`

UnsetBeneficiaryName ensures that no value is present for BeneficiaryName, not even an explicit nil
### GetAccountNumber

`func (o *OrganizationFinancialAccountDto) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *OrganizationFinancialAccountDto) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *OrganizationFinancialAccountDto) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *OrganizationFinancialAccountDto) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### SetAccountNumberNil

`func (o *OrganizationFinancialAccountDto) SetAccountNumberNil(b bool)`

 SetAccountNumberNil sets the value for AccountNumber to be an explicit nil

### UnsetAccountNumber
`func (o *OrganizationFinancialAccountDto) UnsetAccountNumber()`

UnsetAccountNumber ensures that no value is present for AccountNumber, not even an explicit nil
### GetRoutingNumber

`func (o *OrganizationFinancialAccountDto) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *OrganizationFinancialAccountDto) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *OrganizationFinancialAccountDto) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.

### HasRoutingNumber

`func (o *OrganizationFinancialAccountDto) HasRoutingNumber() bool`

HasRoutingNumber returns a boolean if a field has been set.

### SetRoutingNumberNil

`func (o *OrganizationFinancialAccountDto) SetRoutingNumberNil(b bool)`

 SetRoutingNumberNil sets the value for RoutingNumber to be an explicit nil

### UnsetRoutingNumber
`func (o *OrganizationFinancialAccountDto) UnsetRoutingNumber()`

UnsetRoutingNumber ensures that no value is present for RoutingNumber, not even an explicit nil
### GetIsPrimary

`func (o *OrganizationFinancialAccountDto) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *OrganizationFinancialAccountDto) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *OrganizationFinancialAccountDto) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *OrganizationFinancialAccountDto) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetIsVerified

`func (o *OrganizationFinancialAccountDto) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *OrganizationFinancialAccountDto) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *OrganizationFinancialAccountDto) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *OrganizationFinancialAccountDto) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetLogoUrl

`func (o *OrganizationFinancialAccountDto) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *OrganizationFinancialAccountDto) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *OrganizationFinancialAccountDto) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *OrganizationFinancialAccountDto) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### SetLogoUrlNil

`func (o *OrganizationFinancialAccountDto) SetLogoUrlNil(b bool)`

 SetLogoUrlNil sets the value for LogoUrl to be an explicit nil

### UnsetLogoUrl
`func (o *OrganizationFinancialAccountDto) UnsetLogoUrl()`

UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
### GetBankName

`func (o *OrganizationFinancialAccountDto) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *OrganizationFinancialAccountDto) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *OrganizationFinancialAccountDto) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *OrganizationFinancialAccountDto) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### SetBankNameNil

`func (o *OrganizationFinancialAccountDto) SetBankNameNil(b bool)`

 SetBankNameNil sets the value for BankName to be an explicit nil

### UnsetBankName
`func (o *OrganizationFinancialAccountDto) UnsetBankName()`

UnsetBankName ensures that no value is present for BankName, not even an explicit nil
### GetBankAccountType

`func (o *OrganizationFinancialAccountDto) GetBankAccountType() BankAccountType`

GetBankAccountType returns the BankAccountType field if non-nil, zero value otherwise.

### GetBankAccountTypeOk

`func (o *OrganizationFinancialAccountDto) GetBankAccountTypeOk() (*BankAccountType, bool)`

GetBankAccountTypeOk returns a tuple with the BankAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountType

`func (o *OrganizationFinancialAccountDto) SetBankAccountType(v BankAccountType)`

SetBankAccountType sets BankAccountType field to given value.

### HasBankAccountType

`func (o *OrganizationFinancialAccountDto) HasBankAccountType() bool`

HasBankAccountType returns a boolean if a field has been set.

### GetBankAccountProofUrl

`func (o *OrganizationFinancialAccountDto) GetBankAccountProofUrl() string`

GetBankAccountProofUrl returns the BankAccountProofUrl field if non-nil, zero value otherwise.

### GetBankAccountProofUrlOk

`func (o *OrganizationFinancialAccountDto) GetBankAccountProofUrlOk() (*string, bool)`

GetBankAccountProofUrlOk returns a tuple with the BankAccountProofUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountProofUrl

`func (o *OrganizationFinancialAccountDto) SetBankAccountProofUrl(v string)`

SetBankAccountProofUrl sets BankAccountProofUrl field to given value.

### HasBankAccountProofUrl

`func (o *OrganizationFinancialAccountDto) HasBankAccountProofUrl() bool`

HasBankAccountProofUrl returns a boolean if a field has been set.

### SetBankAccountProofUrlNil

`func (o *OrganizationFinancialAccountDto) SetBankAccountProofUrlNil(b bool)`

 SetBankAccountProofUrlNil sets the value for BankAccountProofUrl to be an explicit nil

### UnsetBankAccountProofUrl
`func (o *OrganizationFinancialAccountDto) UnsetBankAccountProofUrl()`

UnsetBankAccountProofUrl ensures that no value is present for BankAccountProofUrl, not even an explicit nil
### GetFileType

`func (o *OrganizationFinancialAccountDto) GetFileType() FileType`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *OrganizationFinancialAccountDto) GetFileTypeOk() (*FileType, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *OrganizationFinancialAccountDto) SetFileType(v FileType)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *OrganizationFinancialAccountDto) HasFileType() bool`

HasFileType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



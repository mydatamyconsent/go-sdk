# DataConsentRequestedAccountDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**AccountTypeId** | Pointer to **string** |  | [optional] 
**SuggestedAccounts** | Pointer to [**[]SuggestedAccountDto**](SuggestedAccountDto.md) |  | [optional] 
**Issuer** | Pointer to **[]string** |  | [optional] 
**IssuerLogoUrls** | Pointer to **[]string** |  | [optional] 
**RequestedDataType** | Pointer to **NullableString** |  | [optional] 
**Optional** | Pointer to **bool** |  | [optional] 

## Methods

### NewDataConsentRequestedAccountDto

`func NewDataConsentRequestedAccountDto() *DataConsentRequestedAccountDto`

NewDataConsentRequestedAccountDto instantiates a new DataConsentRequestedAccountDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataConsentRequestedAccountDtoWithDefaults

`func NewDataConsentRequestedAccountDtoWithDefaults() *DataConsentRequestedAccountDto`

NewDataConsentRequestedAccountDtoWithDefaults instantiates a new DataConsentRequestedAccountDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DataConsentRequestedAccountDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataConsentRequestedAccountDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataConsentRequestedAccountDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataConsentRequestedAccountDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DataConsentRequestedAccountDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DataConsentRequestedAccountDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAccountTypeId

`func (o *DataConsentRequestedAccountDto) GetAccountTypeId() string`

GetAccountTypeId returns the AccountTypeId field if non-nil, zero value otherwise.

### GetAccountTypeIdOk

`func (o *DataConsentRequestedAccountDto) GetAccountTypeIdOk() (*string, bool)`

GetAccountTypeIdOk returns a tuple with the AccountTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTypeId

`func (o *DataConsentRequestedAccountDto) SetAccountTypeId(v string)`

SetAccountTypeId sets AccountTypeId field to given value.

### HasAccountTypeId

`func (o *DataConsentRequestedAccountDto) HasAccountTypeId() bool`

HasAccountTypeId returns a boolean if a field has been set.

### GetSuggestedAccounts

`func (o *DataConsentRequestedAccountDto) GetSuggestedAccounts() []SuggestedAccountDto`

GetSuggestedAccounts returns the SuggestedAccounts field if non-nil, zero value otherwise.

### GetSuggestedAccountsOk

`func (o *DataConsentRequestedAccountDto) GetSuggestedAccountsOk() (*[]SuggestedAccountDto, bool)`

GetSuggestedAccountsOk returns a tuple with the SuggestedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedAccounts

`func (o *DataConsentRequestedAccountDto) SetSuggestedAccounts(v []SuggestedAccountDto)`

SetSuggestedAccounts sets SuggestedAccounts field to given value.

### HasSuggestedAccounts

`func (o *DataConsentRequestedAccountDto) HasSuggestedAccounts() bool`

HasSuggestedAccounts returns a boolean if a field has been set.

### SetSuggestedAccountsNil

`func (o *DataConsentRequestedAccountDto) SetSuggestedAccountsNil(b bool)`

 SetSuggestedAccountsNil sets the value for SuggestedAccounts to be an explicit nil

### UnsetSuggestedAccounts
`func (o *DataConsentRequestedAccountDto) UnsetSuggestedAccounts()`

UnsetSuggestedAccounts ensures that no value is present for SuggestedAccounts, not even an explicit nil
### GetIssuer

`func (o *DataConsentRequestedAccountDto) GetIssuer() []string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *DataConsentRequestedAccountDto) GetIssuerOk() (*[]string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *DataConsentRequestedAccountDto) SetIssuer(v []string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *DataConsentRequestedAccountDto) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### SetIssuerNil

`func (o *DataConsentRequestedAccountDto) SetIssuerNil(b bool)`

 SetIssuerNil sets the value for Issuer to be an explicit nil

### UnsetIssuer
`func (o *DataConsentRequestedAccountDto) UnsetIssuer()`

UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
### GetIssuerLogoUrls

`func (o *DataConsentRequestedAccountDto) GetIssuerLogoUrls() []string`

GetIssuerLogoUrls returns the IssuerLogoUrls field if non-nil, zero value otherwise.

### GetIssuerLogoUrlsOk

`func (o *DataConsentRequestedAccountDto) GetIssuerLogoUrlsOk() (*[]string, bool)`

GetIssuerLogoUrlsOk returns a tuple with the IssuerLogoUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerLogoUrls

`func (o *DataConsentRequestedAccountDto) SetIssuerLogoUrls(v []string)`

SetIssuerLogoUrls sets IssuerLogoUrls field to given value.

### HasIssuerLogoUrls

`func (o *DataConsentRequestedAccountDto) HasIssuerLogoUrls() bool`

HasIssuerLogoUrls returns a boolean if a field has been set.

### SetIssuerLogoUrlsNil

`func (o *DataConsentRequestedAccountDto) SetIssuerLogoUrlsNil(b bool)`

 SetIssuerLogoUrlsNil sets the value for IssuerLogoUrls to be an explicit nil

### UnsetIssuerLogoUrls
`func (o *DataConsentRequestedAccountDto) UnsetIssuerLogoUrls()`

UnsetIssuerLogoUrls ensures that no value is present for IssuerLogoUrls, not even an explicit nil
### GetRequestedDataType

`func (o *DataConsentRequestedAccountDto) GetRequestedDataType() string`

GetRequestedDataType returns the RequestedDataType field if non-nil, zero value otherwise.

### GetRequestedDataTypeOk

`func (o *DataConsentRequestedAccountDto) GetRequestedDataTypeOk() (*string, bool)`

GetRequestedDataTypeOk returns a tuple with the RequestedDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedDataType

`func (o *DataConsentRequestedAccountDto) SetRequestedDataType(v string)`

SetRequestedDataType sets RequestedDataType field to given value.

### HasRequestedDataType

`func (o *DataConsentRequestedAccountDto) HasRequestedDataType() bool`

HasRequestedDataType returns a boolean if a field has been set.

### SetRequestedDataTypeNil

`func (o *DataConsentRequestedAccountDto) SetRequestedDataTypeNil(b bool)`

 SetRequestedDataTypeNil sets the value for RequestedDataType to be an explicit nil

### UnsetRequestedDataType
`func (o *DataConsentRequestedAccountDto) UnsetRequestedDataType()`

UnsetRequestedDataType ensures that no value is present for RequestedDataType, not even an explicit nil
### GetOptional

`func (o *DataConsentRequestedAccountDto) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *DataConsentRequestedAccountDto) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *DataConsentRequestedAccountDto) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *DataConsentRequestedAccountDto) HasOptional() bool`

HasOptional returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



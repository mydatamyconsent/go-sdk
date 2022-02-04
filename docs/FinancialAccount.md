# FinancialAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Issuer** | Pointer to **NullableString** |  | [optional] 
**LogoUrl** | Pointer to **NullableString** |  | [optional] 
**CategoryId** | Pointer to **NullableString** |  | [optional] 
**CategoryName** | Pointer to **NullableString** |  | [optional] 
**AccountName** | Pointer to **NullableString** |  | [optional] 
**AccountLogoUrl** | Pointer to **NullableString** |  | [optional] 
**Balance** | Pointer to **NullableString** |  | [optional] 
**BalanceType** | Pointer to **NullableString** |  | [optional] 
**IsShared** | Pointer to **bool** |  | [optional] 
**SharedWith** | Pointer to [**[]SharedWith**](SharedWith.md) |  | [optional] 
**IsReceived** | Pointer to **bool** |  | [optional] 
**ExpiresAtUtc** | Pointer to **time.Time** |  | [optional] 
**Activities** | Pointer to [**[]Activity**](Activity.md) |  | [optional] 
**ApprovedConsentRequests** | Pointer to [**[]ApprovedConsentRequest**](ApprovedConsentRequest.md) |  | [optional] 

## Methods

### NewFinancialAccount

`func NewFinancialAccount() *FinancialAccount`

NewFinancialAccount instantiates a new FinancialAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialAccountWithDefaults

`func NewFinancialAccountWithDefaults() *FinancialAccount`

NewFinancialAccountWithDefaults instantiates a new FinancialAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FinancialAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FinancialAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FinancialAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FinancialAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssuer

`func (o *FinancialAccount) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *FinancialAccount) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *FinancialAccount) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *FinancialAccount) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### SetIssuerNil

`func (o *FinancialAccount) SetIssuerNil(b bool)`

 SetIssuerNil sets the value for Issuer to be an explicit nil

### UnsetIssuer
`func (o *FinancialAccount) UnsetIssuer()`

UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
### GetLogoUrl

`func (o *FinancialAccount) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *FinancialAccount) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *FinancialAccount) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *FinancialAccount) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### SetLogoUrlNil

`func (o *FinancialAccount) SetLogoUrlNil(b bool)`

 SetLogoUrlNil sets the value for LogoUrl to be an explicit nil

### UnsetLogoUrl
`func (o *FinancialAccount) UnsetLogoUrl()`

UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
### GetCategoryId

`func (o *FinancialAccount) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *FinancialAccount) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *FinancialAccount) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *FinancialAccount) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### SetCategoryIdNil

`func (o *FinancialAccount) SetCategoryIdNil(b bool)`

 SetCategoryIdNil sets the value for CategoryId to be an explicit nil

### UnsetCategoryId
`func (o *FinancialAccount) UnsetCategoryId()`

UnsetCategoryId ensures that no value is present for CategoryId, not even an explicit nil
### GetCategoryName

`func (o *FinancialAccount) GetCategoryName() string`

GetCategoryName returns the CategoryName field if non-nil, zero value otherwise.

### GetCategoryNameOk

`func (o *FinancialAccount) GetCategoryNameOk() (*string, bool)`

GetCategoryNameOk returns a tuple with the CategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryName

`func (o *FinancialAccount) SetCategoryName(v string)`

SetCategoryName sets CategoryName field to given value.

### HasCategoryName

`func (o *FinancialAccount) HasCategoryName() bool`

HasCategoryName returns a boolean if a field has been set.

### SetCategoryNameNil

`func (o *FinancialAccount) SetCategoryNameNil(b bool)`

 SetCategoryNameNil sets the value for CategoryName to be an explicit nil

### UnsetCategoryName
`func (o *FinancialAccount) UnsetCategoryName()`

UnsetCategoryName ensures that no value is present for CategoryName, not even an explicit nil
### GetAccountName

`func (o *FinancialAccount) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *FinancialAccount) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *FinancialAccount) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *FinancialAccount) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### SetAccountNameNil

`func (o *FinancialAccount) SetAccountNameNil(b bool)`

 SetAccountNameNil sets the value for AccountName to be an explicit nil

### UnsetAccountName
`func (o *FinancialAccount) UnsetAccountName()`

UnsetAccountName ensures that no value is present for AccountName, not even an explicit nil
### GetAccountLogoUrl

`func (o *FinancialAccount) GetAccountLogoUrl() string`

GetAccountLogoUrl returns the AccountLogoUrl field if non-nil, zero value otherwise.

### GetAccountLogoUrlOk

`func (o *FinancialAccount) GetAccountLogoUrlOk() (*string, bool)`

GetAccountLogoUrlOk returns a tuple with the AccountLogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLogoUrl

`func (o *FinancialAccount) SetAccountLogoUrl(v string)`

SetAccountLogoUrl sets AccountLogoUrl field to given value.

### HasAccountLogoUrl

`func (o *FinancialAccount) HasAccountLogoUrl() bool`

HasAccountLogoUrl returns a boolean if a field has been set.

### SetAccountLogoUrlNil

`func (o *FinancialAccount) SetAccountLogoUrlNil(b bool)`

 SetAccountLogoUrlNil sets the value for AccountLogoUrl to be an explicit nil

### UnsetAccountLogoUrl
`func (o *FinancialAccount) UnsetAccountLogoUrl()`

UnsetAccountLogoUrl ensures that no value is present for AccountLogoUrl, not even an explicit nil
### GetBalance

`func (o *FinancialAccount) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *FinancialAccount) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *FinancialAccount) SetBalance(v string)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *FinancialAccount) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### SetBalanceNil

`func (o *FinancialAccount) SetBalanceNil(b bool)`

 SetBalanceNil sets the value for Balance to be an explicit nil

### UnsetBalance
`func (o *FinancialAccount) UnsetBalance()`

UnsetBalance ensures that no value is present for Balance, not even an explicit nil
### GetBalanceType

`func (o *FinancialAccount) GetBalanceType() string`

GetBalanceType returns the BalanceType field if non-nil, zero value otherwise.

### GetBalanceTypeOk

`func (o *FinancialAccount) GetBalanceTypeOk() (*string, bool)`

GetBalanceTypeOk returns a tuple with the BalanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceType

`func (o *FinancialAccount) SetBalanceType(v string)`

SetBalanceType sets BalanceType field to given value.

### HasBalanceType

`func (o *FinancialAccount) HasBalanceType() bool`

HasBalanceType returns a boolean if a field has been set.

### SetBalanceTypeNil

`func (o *FinancialAccount) SetBalanceTypeNil(b bool)`

 SetBalanceTypeNil sets the value for BalanceType to be an explicit nil

### UnsetBalanceType
`func (o *FinancialAccount) UnsetBalanceType()`

UnsetBalanceType ensures that no value is present for BalanceType, not even an explicit nil
### GetIsShared

`func (o *FinancialAccount) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *FinancialAccount) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *FinancialAccount) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *FinancialAccount) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### GetSharedWith

`func (o *FinancialAccount) GetSharedWith() []SharedWith`

GetSharedWith returns the SharedWith field if non-nil, zero value otherwise.

### GetSharedWithOk

`func (o *FinancialAccount) GetSharedWithOk() (*[]SharedWith, bool)`

GetSharedWithOk returns a tuple with the SharedWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedWith

`func (o *FinancialAccount) SetSharedWith(v []SharedWith)`

SetSharedWith sets SharedWith field to given value.

### HasSharedWith

`func (o *FinancialAccount) HasSharedWith() bool`

HasSharedWith returns a boolean if a field has been set.

### SetSharedWithNil

`func (o *FinancialAccount) SetSharedWithNil(b bool)`

 SetSharedWithNil sets the value for SharedWith to be an explicit nil

### UnsetSharedWith
`func (o *FinancialAccount) UnsetSharedWith()`

UnsetSharedWith ensures that no value is present for SharedWith, not even an explicit nil
### GetIsReceived

`func (o *FinancialAccount) GetIsReceived() bool`

GetIsReceived returns the IsReceived field if non-nil, zero value otherwise.

### GetIsReceivedOk

`func (o *FinancialAccount) GetIsReceivedOk() (*bool, bool)`

GetIsReceivedOk returns a tuple with the IsReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReceived

`func (o *FinancialAccount) SetIsReceived(v bool)`

SetIsReceived sets IsReceived field to given value.

### HasIsReceived

`func (o *FinancialAccount) HasIsReceived() bool`

HasIsReceived returns a boolean if a field has been set.

### GetExpiresAtUtc

`func (o *FinancialAccount) GetExpiresAtUtc() time.Time`

GetExpiresAtUtc returns the ExpiresAtUtc field if non-nil, zero value otherwise.

### GetExpiresAtUtcOk

`func (o *FinancialAccount) GetExpiresAtUtcOk() (*time.Time, bool)`

GetExpiresAtUtcOk returns a tuple with the ExpiresAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAtUtc

`func (o *FinancialAccount) SetExpiresAtUtc(v time.Time)`

SetExpiresAtUtc sets ExpiresAtUtc field to given value.

### HasExpiresAtUtc

`func (o *FinancialAccount) HasExpiresAtUtc() bool`

HasExpiresAtUtc returns a boolean if a field has been set.

### GetActivities

`func (o *FinancialAccount) GetActivities() []Activity`

GetActivities returns the Activities field if non-nil, zero value otherwise.

### GetActivitiesOk

`func (o *FinancialAccount) GetActivitiesOk() (*[]Activity, bool)`

GetActivitiesOk returns a tuple with the Activities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivities

`func (o *FinancialAccount) SetActivities(v []Activity)`

SetActivities sets Activities field to given value.

### HasActivities

`func (o *FinancialAccount) HasActivities() bool`

HasActivities returns a boolean if a field has been set.

### SetActivitiesNil

`func (o *FinancialAccount) SetActivitiesNil(b bool)`

 SetActivitiesNil sets the value for Activities to be an explicit nil

### UnsetActivities
`func (o *FinancialAccount) UnsetActivities()`

UnsetActivities ensures that no value is present for Activities, not even an explicit nil
### GetApprovedConsentRequests

`func (o *FinancialAccount) GetApprovedConsentRequests() []ApprovedConsentRequest`

GetApprovedConsentRequests returns the ApprovedConsentRequests field if non-nil, zero value otherwise.

### GetApprovedConsentRequestsOk

`func (o *FinancialAccount) GetApprovedConsentRequestsOk() (*[]ApprovedConsentRequest, bool)`

GetApprovedConsentRequestsOk returns a tuple with the ApprovedConsentRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedConsentRequests

`func (o *FinancialAccount) SetApprovedConsentRequests(v []ApprovedConsentRequest)`

SetApprovedConsentRequests sets ApprovedConsentRequests field to given value.

### HasApprovedConsentRequests

`func (o *FinancialAccount) HasApprovedConsentRequests() bool`

HasApprovedConsentRequests returns a boolean if a field has been set.

### SetApprovedConsentRequestsNil

`func (o *FinancialAccount) SetApprovedConsentRequestsNil(b bool)`

 SetApprovedConsentRequestsNil sets the value for ApprovedConsentRequests to be an explicit nil

### UnsetApprovedConsentRequests
`func (o *FinancialAccount) UnsetApprovedConsentRequests()`

UnsetApprovedConsentRequests ensures that no value is present for ApprovedConsentRequests, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



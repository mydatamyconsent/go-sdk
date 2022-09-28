# CreditCardHolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreditCardId** | **string** |  | 
**Name** | **string** |  | 
**Dob** | **time.Time** |  | 
**Mobile** | **string** |  | 
**Nominee** | [**HoldingNominee**](HoldingNominee.md) |  | 
**Landline** | Pointer to **string** |  | [optional] 
**Address** | **string** |  | 
**Email** | **string** |  | 
**Pan** | **string** |  | 
**CkycCompliance** | **bool** |  | 

## Methods

### NewCreditCardHolder

`func NewCreditCardHolder(id string, creditCardId string, name string, dob time.Time, mobile string, nominee HoldingNominee, address string, email string, pan string, ckycCompliance bool, ) *CreditCardHolder`

NewCreditCardHolder instantiates a new CreditCardHolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditCardHolderWithDefaults

`func NewCreditCardHolderWithDefaults() *CreditCardHolder`

NewCreditCardHolderWithDefaults instantiates a new CreditCardHolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreditCardHolder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreditCardHolder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreditCardHolder) SetId(v string)`

SetId sets Id field to given value.


### GetCreditCardId

`func (o *CreditCardHolder) GetCreditCardId() string`

GetCreditCardId returns the CreditCardId field if non-nil, zero value otherwise.

### GetCreditCardIdOk

`func (o *CreditCardHolder) GetCreditCardIdOk() (*string, bool)`

GetCreditCardIdOk returns a tuple with the CreditCardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCardId

`func (o *CreditCardHolder) SetCreditCardId(v string)`

SetCreditCardId sets CreditCardId field to given value.


### GetName

`func (o *CreditCardHolder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreditCardHolder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreditCardHolder) SetName(v string)`

SetName sets Name field to given value.


### GetDob

`func (o *CreditCardHolder) GetDob() time.Time`

GetDob returns the Dob field if non-nil, zero value otherwise.

### GetDobOk

`func (o *CreditCardHolder) GetDobOk() (*time.Time, bool)`

GetDobOk returns a tuple with the Dob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDob

`func (o *CreditCardHolder) SetDob(v time.Time)`

SetDob sets Dob field to given value.


### GetMobile

`func (o *CreditCardHolder) GetMobile() string`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *CreditCardHolder) GetMobileOk() (*string, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *CreditCardHolder) SetMobile(v string)`

SetMobile sets Mobile field to given value.


### GetNominee

`func (o *CreditCardHolder) GetNominee() HoldingNominee`

GetNominee returns the Nominee field if non-nil, zero value otherwise.

### GetNomineeOk

`func (o *CreditCardHolder) GetNomineeOk() (*HoldingNominee, bool)`

GetNomineeOk returns a tuple with the Nominee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNominee

`func (o *CreditCardHolder) SetNominee(v HoldingNominee)`

SetNominee sets Nominee field to given value.


### GetLandline

`func (o *CreditCardHolder) GetLandline() string`

GetLandline returns the Landline field if non-nil, zero value otherwise.

### GetLandlineOk

`func (o *CreditCardHolder) GetLandlineOk() (*string, bool)`

GetLandlineOk returns a tuple with the Landline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandline

`func (o *CreditCardHolder) SetLandline(v string)`

SetLandline sets Landline field to given value.

### HasLandline

`func (o *CreditCardHolder) HasLandline() bool`

HasLandline returns a boolean if a field has been set.

### GetAddress

`func (o *CreditCardHolder) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreditCardHolder) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreditCardHolder) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetEmail

`func (o *CreditCardHolder) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreditCardHolder) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreditCardHolder) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPan

`func (o *CreditCardHolder) GetPan() string`

GetPan returns the Pan field if non-nil, zero value otherwise.

### GetPanOk

`func (o *CreditCardHolder) GetPanOk() (*string, bool)`

GetPanOk returns a tuple with the Pan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPan

`func (o *CreditCardHolder) SetPan(v string)`

SetPan sets Pan field to given value.


### GetCkycCompliance

`func (o *CreditCardHolder) GetCkycCompliance() bool`

GetCkycCompliance returns the CkycCompliance field if non-nil, zero value otherwise.

### GetCkycComplianceOk

`func (o *CreditCardHolder) GetCkycComplianceOk() (*bool, bool)`

GetCkycComplianceOk returns a tuple with the CkycCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCkycCompliance

`func (o *CreditCardHolder) SetCkycCompliance(v bool)`

SetCkycCompliance sets CkycCompliance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



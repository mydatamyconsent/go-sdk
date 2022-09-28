# Holder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Dob** | Pointer to **time.Time** |  | [optional] 
**Mobile** | Pointer to **string** |  | [optional] 
**Nominee** | Pointer to [**HoldingNominee**](HoldingNominee.md) |  | [optional] 
**DematId** | **string** |  | 
**Landline** | Pointer to **string** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**Email** | **string** |  | 
**Pan** | Pointer to **string** |  | [optional] 
**CkycCompliance** | **bool** |  | 

## Methods

### NewHolder

`func NewHolder(name string, dematId string, email string, ckycCompliance bool, ) *Holder`

NewHolder instantiates a new Holder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHolderWithDefaults

`func NewHolderWithDefaults() *Holder`

NewHolderWithDefaults instantiates a new Holder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Holder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Holder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Holder) SetName(v string)`

SetName sets Name field to given value.


### GetDob

`func (o *Holder) GetDob() time.Time`

GetDob returns the Dob field if non-nil, zero value otherwise.

### GetDobOk

`func (o *Holder) GetDobOk() (*time.Time, bool)`

GetDobOk returns a tuple with the Dob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDob

`func (o *Holder) SetDob(v time.Time)`

SetDob sets Dob field to given value.

### HasDob

`func (o *Holder) HasDob() bool`

HasDob returns a boolean if a field has been set.

### GetMobile

`func (o *Holder) GetMobile() string`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *Holder) GetMobileOk() (*string, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *Holder) SetMobile(v string)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *Holder) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetNominee

`func (o *Holder) GetNominee() HoldingNominee`

GetNominee returns the Nominee field if non-nil, zero value otherwise.

### GetNomineeOk

`func (o *Holder) GetNomineeOk() (*HoldingNominee, bool)`

GetNomineeOk returns a tuple with the Nominee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNominee

`func (o *Holder) SetNominee(v HoldingNominee)`

SetNominee sets Nominee field to given value.

### HasNominee

`func (o *Holder) HasNominee() bool`

HasNominee returns a boolean if a field has been set.

### GetDematId

`func (o *Holder) GetDematId() string`

GetDematId returns the DematId field if non-nil, zero value otherwise.

### GetDematIdOk

`func (o *Holder) GetDematIdOk() (*string, bool)`

GetDematIdOk returns a tuple with the DematId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDematId

`func (o *Holder) SetDematId(v string)`

SetDematId sets DematId field to given value.


### GetLandline

`func (o *Holder) GetLandline() string`

GetLandline returns the Landline field if non-nil, zero value otherwise.

### GetLandlineOk

`func (o *Holder) GetLandlineOk() (*string, bool)`

GetLandlineOk returns a tuple with the Landline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandline

`func (o *Holder) SetLandline(v string)`

SetLandline sets Landline field to given value.

### HasLandline

`func (o *Holder) HasLandline() bool`

HasLandline returns a boolean if a field has been set.

### GetAddress

`func (o *Holder) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Holder) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Holder) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Holder) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetEmail

`func (o *Holder) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Holder) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Holder) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPan

`func (o *Holder) GetPan() string`

GetPan returns the Pan field if non-nil, zero value otherwise.

### GetPanOk

`func (o *Holder) GetPanOk() (*string, bool)`

GetPanOk returns a tuple with the Pan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPan

`func (o *Holder) SetPan(v string)`

SetPan sets Pan field to given value.

### HasPan

`func (o *Holder) HasPan() bool`

HasPan returns a boolean if a field has been set.

### GetCkycCompliance

`func (o *Holder) GetCkycCompliance() bool`

GetCkycCompliance returns the CkycCompliance field if non-nil, zero value otherwise.

### GetCkycComplianceOk

`func (o *Holder) GetCkycComplianceOk() (*bool, bool)`

GetCkycComplianceOk returns a tuple with the CkycCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCkycCompliance

`func (o *Holder) SetCkycCompliance(v bool)`

SetCkycCompliance sets CkycCompliance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



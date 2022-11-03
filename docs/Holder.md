# Holder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**DateOfBirth** | Pointer to **time.Time** |  | [optional] 
**Mobile** | Pointer to **string** |  | [optional] 
**DematId** | **string** |  | 
**Email** | **string** |  | 
**Pan** | Pointer to **string** |  | [optional] 

## Methods

### NewHolder

`func NewHolder(name string, dematId string, email string, ) *Holder`

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


### GetDateOfBirth

`func (o *Holder) GetDateOfBirth() time.Time`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *Holder) GetDateOfBirthOk() (*time.Time, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *Holder) SetDateOfBirth(v time.Time)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *Holder) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



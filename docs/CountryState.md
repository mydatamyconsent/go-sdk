# CountryState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CountryId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**FipsCode** | Pointer to **int32** |  | [optional] 
**Country** | Pointer to [**Country**](Country.md) |  | [optional] 

## Methods

### NewCountryState

`func NewCountryState() *CountryState`

NewCountryState instantiates a new CountryState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryStateWithDefaults

`func NewCountryStateWithDefaults() *CountryState`

NewCountryStateWithDefaults instantiates a new CountryState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CountryState) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CountryState) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CountryState) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CountryState) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCountryId

`func (o *CountryState) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *CountryState) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *CountryState) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *CountryState) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### GetName

`func (o *CountryState) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CountryState) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CountryState) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CountryState) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CountryState) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CountryState) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFipsCode

`func (o *CountryState) GetFipsCode() int32`

GetFipsCode returns the FipsCode field if non-nil, zero value otherwise.

### GetFipsCodeOk

`func (o *CountryState) GetFipsCodeOk() (*int32, bool)`

GetFipsCodeOk returns a tuple with the FipsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFipsCode

`func (o *CountryState) SetFipsCode(v int32)`

SetFipsCode sets FipsCode field to given value.

### HasFipsCode

`func (o *CountryState) HasFipsCode() bool`

HasFipsCode returns a boolean if a field has been set.

### GetCountry

`func (o *CountryState) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CountryState) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CountryState) SetCountry(v Country)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CountryState) HasCountry() bool`

HasCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



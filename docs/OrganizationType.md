# OrganizationType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CountryId** | Pointer to **string** |  | [optional] 
**TypeName** | Pointer to **NullableString** |  | [optional] 
**ElfCode** | Pointer to **NullableString** |  | [optional] 
**Country** | Pointer to [**Country**](Country.md) |  | [optional] 

## Methods

### NewOrganizationType

`func NewOrganizationType() *OrganizationType`

NewOrganizationType instantiates a new OrganizationType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationTypeWithDefaults

`func NewOrganizationTypeWithDefaults() *OrganizationType`

NewOrganizationTypeWithDefaults instantiates a new OrganizationType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationType) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCountryId

`func (o *OrganizationType) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *OrganizationType) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *OrganizationType) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *OrganizationType) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### GetTypeName

`func (o *OrganizationType) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *OrganizationType) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *OrganizationType) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *OrganizationType) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### SetTypeNameNil

`func (o *OrganizationType) SetTypeNameNil(b bool)`

 SetTypeNameNil sets the value for TypeName to be an explicit nil

### UnsetTypeName
`func (o *OrganizationType) UnsetTypeName()`

UnsetTypeName ensures that no value is present for TypeName, not even an explicit nil
### GetElfCode

`func (o *OrganizationType) GetElfCode() string`

GetElfCode returns the ElfCode field if non-nil, zero value otherwise.

### GetElfCodeOk

`func (o *OrganizationType) GetElfCodeOk() (*string, bool)`

GetElfCodeOk returns a tuple with the ElfCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElfCode

`func (o *OrganizationType) SetElfCode(v string)`

SetElfCode sets ElfCode field to given value.

### HasElfCode

`func (o *OrganizationType) HasElfCode() bool`

HasElfCode returns a boolean if a field has been set.

### SetElfCodeNil

`func (o *OrganizationType) SetElfCodeNil(b bool)`

 SetElfCodeNil sets the value for ElfCode to be an explicit nil

### UnsetElfCode
`func (o *OrganizationType) UnsetElfCode()`

UnsetElfCode ensures that no value is present for ElfCode, not even an explicit nil
### GetCountry

`func (o *OrganizationType) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *OrganizationType) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *OrganizationType) SetCountry(v Country)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *OrganizationType) HasCountry() bool`

HasCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



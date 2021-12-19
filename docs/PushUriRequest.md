# PushUriRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UriDetails** | Pointer to [**UriDetails**](UriDetails.md) |  | [optional] 
**Ns2** | Pointer to **NullableString** |  | [optional] 
**Ver** | Pointer to **NullableString** |  | [optional] 
**Ts** | Pointer to **NullableString** |  | [optional] 
**Txn** | Pointer to **NullableString** |  | [optional] 
**OrgId** | Pointer to **NullableString** |  | [optional] 
**Keyhash** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPushUriRequest

`func NewPushUriRequest() *PushUriRequest`

NewPushUriRequest instantiates a new PushUriRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushUriRequestWithDefaults

`func NewPushUriRequestWithDefaults() *PushUriRequest`

NewPushUriRequestWithDefaults instantiates a new PushUriRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUriDetails

`func (o *PushUriRequest) GetUriDetails() UriDetails`

GetUriDetails returns the UriDetails field if non-nil, zero value otherwise.

### GetUriDetailsOk

`func (o *PushUriRequest) GetUriDetailsOk() (*UriDetails, bool)`

GetUriDetailsOk returns a tuple with the UriDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUriDetails

`func (o *PushUriRequest) SetUriDetails(v UriDetails)`

SetUriDetails sets UriDetails field to given value.

### HasUriDetails

`func (o *PushUriRequest) HasUriDetails() bool`

HasUriDetails returns a boolean if a field has been set.

### GetNs2

`func (o *PushUriRequest) GetNs2() string`

GetNs2 returns the Ns2 field if non-nil, zero value otherwise.

### GetNs2Ok

`func (o *PushUriRequest) GetNs2Ok() (*string, bool)`

GetNs2Ok returns a tuple with the Ns2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNs2

`func (o *PushUriRequest) SetNs2(v string)`

SetNs2 sets Ns2 field to given value.

### HasNs2

`func (o *PushUriRequest) HasNs2() bool`

HasNs2 returns a boolean if a field has been set.

### SetNs2Nil

`func (o *PushUriRequest) SetNs2Nil(b bool)`

 SetNs2Nil sets the value for Ns2 to be an explicit nil

### UnsetNs2
`func (o *PushUriRequest) UnsetNs2()`

UnsetNs2 ensures that no value is present for Ns2, not even an explicit nil
### GetVer

`func (o *PushUriRequest) GetVer() string`

GetVer returns the Ver field if non-nil, zero value otherwise.

### GetVerOk

`func (o *PushUriRequest) GetVerOk() (*string, bool)`

GetVerOk returns a tuple with the Ver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVer

`func (o *PushUriRequest) SetVer(v string)`

SetVer sets Ver field to given value.

### HasVer

`func (o *PushUriRequest) HasVer() bool`

HasVer returns a boolean if a field has been set.

### SetVerNil

`func (o *PushUriRequest) SetVerNil(b bool)`

 SetVerNil sets the value for Ver to be an explicit nil

### UnsetVer
`func (o *PushUriRequest) UnsetVer()`

UnsetVer ensures that no value is present for Ver, not even an explicit nil
### GetTs

`func (o *PushUriRequest) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *PushUriRequest) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *PushUriRequest) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *PushUriRequest) HasTs() bool`

HasTs returns a boolean if a field has been set.

### SetTsNil

`func (o *PushUriRequest) SetTsNil(b bool)`

 SetTsNil sets the value for Ts to be an explicit nil

### UnsetTs
`func (o *PushUriRequest) UnsetTs()`

UnsetTs ensures that no value is present for Ts, not even an explicit nil
### GetTxn

`func (o *PushUriRequest) GetTxn() string`

GetTxn returns the Txn field if non-nil, zero value otherwise.

### GetTxnOk

`func (o *PushUriRequest) GetTxnOk() (*string, bool)`

GetTxnOk returns a tuple with the Txn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxn

`func (o *PushUriRequest) SetTxn(v string)`

SetTxn sets Txn field to given value.

### HasTxn

`func (o *PushUriRequest) HasTxn() bool`

HasTxn returns a boolean if a field has been set.

### SetTxnNil

`func (o *PushUriRequest) SetTxnNil(b bool)`

 SetTxnNil sets the value for Txn to be an explicit nil

### UnsetTxn
`func (o *PushUriRequest) UnsetTxn()`

UnsetTxn ensures that no value is present for Txn, not even an explicit nil
### GetOrgId

`func (o *PushUriRequest) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *PushUriRequest) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *PushUriRequest) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *PushUriRequest) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *PushUriRequest) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *PushUriRequest) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetKeyhash

`func (o *PushUriRequest) GetKeyhash() string`

GetKeyhash returns the Keyhash field if non-nil, zero value otherwise.

### GetKeyhashOk

`func (o *PushUriRequest) GetKeyhashOk() (*string, bool)`

GetKeyhashOk returns a tuple with the Keyhash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyhash

`func (o *PushUriRequest) SetKeyhash(v string)`

SetKeyhash sets Keyhash field to given value.

### HasKeyhash

`func (o *PushUriRequest) HasKeyhash() bool`

HasKeyhash returns a boolean if a field has been set.

### SetKeyhashNil

`func (o *PushUriRequest) SetKeyhashNil(b bool)`

 SetKeyhashNil sets the value for Keyhash to be an explicit nil

### UnsetKeyhash
`func (o *PushUriRequest) UnsetKeyhash()`

UnsetKeyhash ensures that no value is present for Keyhash, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



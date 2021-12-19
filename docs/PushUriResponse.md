# PushUriResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseStatus** | Pointer to **NullableString** |  | [optional] 
**ResponseMessage** | Pointer to **NullableString** |  | [optional] 
**Ns2** | Pointer to **NullableString** |  | [optional] 
**Ver** | Pointer to **NullableString** |  | [optional] 
**Ts** | Pointer to **NullableString** |  | [optional] 
**Txn** | Pointer to **NullableString** |  | [optional] 
**OrgId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPushUriResponse

`func NewPushUriResponse() *PushUriResponse`

NewPushUriResponse instantiates a new PushUriResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushUriResponseWithDefaults

`func NewPushUriResponseWithDefaults() *PushUriResponse`

NewPushUriResponseWithDefaults instantiates a new PushUriResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseStatus

`func (o *PushUriResponse) GetResponseStatus() string`

GetResponseStatus returns the ResponseStatus field if non-nil, zero value otherwise.

### GetResponseStatusOk

`func (o *PushUriResponse) GetResponseStatusOk() (*string, bool)`

GetResponseStatusOk returns a tuple with the ResponseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatus

`func (o *PushUriResponse) SetResponseStatus(v string)`

SetResponseStatus sets ResponseStatus field to given value.

### HasResponseStatus

`func (o *PushUriResponse) HasResponseStatus() bool`

HasResponseStatus returns a boolean if a field has been set.

### SetResponseStatusNil

`func (o *PushUriResponse) SetResponseStatusNil(b bool)`

 SetResponseStatusNil sets the value for ResponseStatus to be an explicit nil

### UnsetResponseStatus
`func (o *PushUriResponse) UnsetResponseStatus()`

UnsetResponseStatus ensures that no value is present for ResponseStatus, not even an explicit nil
### GetResponseMessage

`func (o *PushUriResponse) GetResponseMessage() string`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *PushUriResponse) GetResponseMessageOk() (*string, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *PushUriResponse) SetResponseMessage(v string)`

SetResponseMessage sets ResponseMessage field to given value.

### HasResponseMessage

`func (o *PushUriResponse) HasResponseMessage() bool`

HasResponseMessage returns a boolean if a field has been set.

### SetResponseMessageNil

`func (o *PushUriResponse) SetResponseMessageNil(b bool)`

 SetResponseMessageNil sets the value for ResponseMessage to be an explicit nil

### UnsetResponseMessage
`func (o *PushUriResponse) UnsetResponseMessage()`

UnsetResponseMessage ensures that no value is present for ResponseMessage, not even an explicit nil
### GetNs2

`func (o *PushUriResponse) GetNs2() string`

GetNs2 returns the Ns2 field if non-nil, zero value otherwise.

### GetNs2Ok

`func (o *PushUriResponse) GetNs2Ok() (*string, bool)`

GetNs2Ok returns a tuple with the Ns2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNs2

`func (o *PushUriResponse) SetNs2(v string)`

SetNs2 sets Ns2 field to given value.

### HasNs2

`func (o *PushUriResponse) HasNs2() bool`

HasNs2 returns a boolean if a field has been set.

### SetNs2Nil

`func (o *PushUriResponse) SetNs2Nil(b bool)`

 SetNs2Nil sets the value for Ns2 to be an explicit nil

### UnsetNs2
`func (o *PushUriResponse) UnsetNs2()`

UnsetNs2 ensures that no value is present for Ns2, not even an explicit nil
### GetVer

`func (o *PushUriResponse) GetVer() string`

GetVer returns the Ver field if non-nil, zero value otherwise.

### GetVerOk

`func (o *PushUriResponse) GetVerOk() (*string, bool)`

GetVerOk returns a tuple with the Ver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVer

`func (o *PushUriResponse) SetVer(v string)`

SetVer sets Ver field to given value.

### HasVer

`func (o *PushUriResponse) HasVer() bool`

HasVer returns a boolean if a field has been set.

### SetVerNil

`func (o *PushUriResponse) SetVerNil(b bool)`

 SetVerNil sets the value for Ver to be an explicit nil

### UnsetVer
`func (o *PushUriResponse) UnsetVer()`

UnsetVer ensures that no value is present for Ver, not even an explicit nil
### GetTs

`func (o *PushUriResponse) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *PushUriResponse) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *PushUriResponse) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *PushUriResponse) HasTs() bool`

HasTs returns a boolean if a field has been set.

### SetTsNil

`func (o *PushUriResponse) SetTsNil(b bool)`

 SetTsNil sets the value for Ts to be an explicit nil

### UnsetTs
`func (o *PushUriResponse) UnsetTs()`

UnsetTs ensures that no value is present for Ts, not even an explicit nil
### GetTxn

`func (o *PushUriResponse) GetTxn() string`

GetTxn returns the Txn field if non-nil, zero value otherwise.

### GetTxnOk

`func (o *PushUriResponse) GetTxnOk() (*string, bool)`

GetTxnOk returns a tuple with the Txn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxn

`func (o *PushUriResponse) SetTxn(v string)`

SetTxn sets Txn field to given value.

### HasTxn

`func (o *PushUriResponse) HasTxn() bool`

HasTxn returns a boolean if a field has been set.

### SetTxnNil

`func (o *PushUriResponse) SetTxnNil(b bool)`

 SetTxnNil sets the value for Txn to be an explicit nil

### UnsetTxn
`func (o *PushUriResponse) UnsetTxn()`

UnsetTxn ensures that no value is present for Txn, not even an explicit nil
### GetOrgId

`func (o *PushUriResponse) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *PushUriResponse) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *PushUriResponse) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *PushUriResponse) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *PushUriResponse) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *PushUriResponse) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



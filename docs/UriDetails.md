# UriDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aadhaar** | **string** |  | 
**Uri** | **string** |  | 
**DocType** | **string** |  | 
**DocName** | **string** |  | 
**DocId** | **string** |  | 
**IssuedOn** | **string** |  | 
**ValidFrom** | **string** |  | 
**ValidTo** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableString** |  | [optional] 
**Action** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUriDetails

`func NewUriDetails(aadhaar string, uri string, docType string, docName string, docId string, issuedOn string, validFrom string, ) *UriDetails`

NewUriDetails instantiates a new UriDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUriDetailsWithDefaults

`func NewUriDetailsWithDefaults() *UriDetails`

NewUriDetailsWithDefaults instantiates a new UriDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAadhaar

`func (o *UriDetails) GetAadhaar() string`

GetAadhaar returns the Aadhaar field if non-nil, zero value otherwise.

### GetAadhaarOk

`func (o *UriDetails) GetAadhaarOk() (*string, bool)`

GetAadhaarOk returns a tuple with the Aadhaar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAadhaar

`func (o *UriDetails) SetAadhaar(v string)`

SetAadhaar sets Aadhaar field to given value.


### GetUri

`func (o *UriDetails) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *UriDetails) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *UriDetails) SetUri(v string)`

SetUri sets Uri field to given value.


### GetDocType

`func (o *UriDetails) GetDocType() string`

GetDocType returns the DocType field if non-nil, zero value otherwise.

### GetDocTypeOk

`func (o *UriDetails) GetDocTypeOk() (*string, bool)`

GetDocTypeOk returns a tuple with the DocType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocType

`func (o *UriDetails) SetDocType(v string)`

SetDocType sets DocType field to given value.


### GetDocName

`func (o *UriDetails) GetDocName() string`

GetDocName returns the DocName field if non-nil, zero value otherwise.

### GetDocNameOk

`func (o *UriDetails) GetDocNameOk() (*string, bool)`

GetDocNameOk returns a tuple with the DocName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocName

`func (o *UriDetails) SetDocName(v string)`

SetDocName sets DocName field to given value.


### GetDocId

`func (o *UriDetails) GetDocId() string`

GetDocId returns the DocId field if non-nil, zero value otherwise.

### GetDocIdOk

`func (o *UriDetails) GetDocIdOk() (*string, bool)`

GetDocIdOk returns a tuple with the DocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocId

`func (o *UriDetails) SetDocId(v string)`

SetDocId sets DocId field to given value.


### GetIssuedOn

`func (o *UriDetails) GetIssuedOn() string`

GetIssuedOn returns the IssuedOn field if non-nil, zero value otherwise.

### GetIssuedOnOk

`func (o *UriDetails) GetIssuedOnOk() (*string, bool)`

GetIssuedOnOk returns a tuple with the IssuedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedOn

`func (o *UriDetails) SetIssuedOn(v string)`

SetIssuedOn sets IssuedOn field to given value.


### GetValidFrom

`func (o *UriDetails) GetValidFrom() string`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *UriDetails) GetValidFromOk() (*string, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *UriDetails) SetValidFrom(v string)`

SetValidFrom sets ValidFrom field to given value.


### GetValidTo

`func (o *UriDetails) GetValidTo() string`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *UriDetails) GetValidToOk() (*string, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *UriDetails) SetValidTo(v string)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *UriDetails) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.

### SetValidToNil

`func (o *UriDetails) SetValidToNil(b bool)`

 SetValidToNil sets the value for ValidTo to be an explicit nil

### UnsetValidTo
`func (o *UriDetails) UnsetValidTo()`

UnsetValidTo ensures that no value is present for ValidTo, not even an explicit nil
### GetTimestamp

`func (o *UriDetails) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UriDetails) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UriDetails) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *UriDetails) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *UriDetails) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *UriDetails) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetAction

`func (o *UriDetails) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UriDetails) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UriDetails) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *UriDetails) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *UriDetails) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *UriDetails) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



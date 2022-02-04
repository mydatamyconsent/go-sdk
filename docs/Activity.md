# Activity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActorName** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ActorProfileUrl** | Pointer to **NullableString** |  | [optional] 
**DateTimeUtc** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewActivity

`func NewActivity() *Activity`

NewActivity instantiates a new Activity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityWithDefaults

`func NewActivityWithDefaults() *Activity`

NewActivityWithDefaults instantiates a new Activity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActorName

`func (o *Activity) GetActorName() string`

GetActorName returns the ActorName field if non-nil, zero value otherwise.

### GetActorNameOk

`func (o *Activity) GetActorNameOk() (*string, bool)`

GetActorNameOk returns a tuple with the ActorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorName

`func (o *Activity) SetActorName(v string)`

SetActorName sets ActorName field to given value.

### HasActorName

`func (o *Activity) HasActorName() bool`

HasActorName returns a boolean if a field has been set.

### SetActorNameNil

`func (o *Activity) SetActorNameNil(b bool)`

 SetActorNameNil sets the value for ActorName to be an explicit nil

### UnsetActorName
`func (o *Activity) UnsetActorName()`

UnsetActorName ensures that no value is present for ActorName, not even an explicit nil
### GetDescription

`func (o *Activity) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Activity) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Activity) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Activity) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Activity) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Activity) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetActorProfileUrl

`func (o *Activity) GetActorProfileUrl() string`

GetActorProfileUrl returns the ActorProfileUrl field if non-nil, zero value otherwise.

### GetActorProfileUrlOk

`func (o *Activity) GetActorProfileUrlOk() (*string, bool)`

GetActorProfileUrlOk returns a tuple with the ActorProfileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorProfileUrl

`func (o *Activity) SetActorProfileUrl(v string)`

SetActorProfileUrl sets ActorProfileUrl field to given value.

### HasActorProfileUrl

`func (o *Activity) HasActorProfileUrl() bool`

HasActorProfileUrl returns a boolean if a field has been set.

### SetActorProfileUrlNil

`func (o *Activity) SetActorProfileUrlNil(b bool)`

 SetActorProfileUrlNil sets the value for ActorProfileUrl to be an explicit nil

### UnsetActorProfileUrl
`func (o *Activity) UnsetActorProfileUrl()`

UnsetActorProfileUrl ensures that no value is present for ActorProfileUrl, not even an explicit nil
### GetDateTimeUtc

`func (o *Activity) GetDateTimeUtc() time.Time`

GetDateTimeUtc returns the DateTimeUtc field if non-nil, zero value otherwise.

### GetDateTimeUtcOk

`func (o *Activity) GetDateTimeUtcOk() (*time.Time, bool)`

GetDateTimeUtcOk returns a tuple with the DateTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeUtc

`func (o *Activity) SetDateTimeUtc(v time.Time)`

SetDateTimeUtc sets DateTimeUtc field to given value.

### HasDateTimeUtc

`func (o *Activity) HasDateTimeUtc() bool`

HasDateTimeUtc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



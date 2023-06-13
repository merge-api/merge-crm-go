# EngagementType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityType** | Pointer to [**NullableActivityTypeEnum**](ActivityTypeEnum.md) | The engagement type&#39;s activity type.  * &#x60;CALL&#x60; - CALL * &#x60;MEETING&#x60; - MEETING * &#x60;EMAIL&#x60; - EMAIL | [optional] 
**Name** | Pointer to **NullableString** | The engagement type&#39;s name. | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**ModifiedAt** | Pointer to **time.Time** | This is the datetime that this object was last updated by Merge | [optional] [readonly] 
**RemoteFields** | Pointer to [**[]RemoteField**](RemoteField.md) |  | [optional] [readonly] 

## Methods

### NewEngagementType

`func NewEngagementType() *EngagementType`

NewEngagementType instantiates a new EngagementType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngagementTypeWithDefaults

`func NewEngagementTypeWithDefaults() *EngagementType`

NewEngagementTypeWithDefaults instantiates a new EngagementType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityType

`func (o *EngagementType) GetActivityType() ActivityTypeEnum`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *EngagementType) GetActivityTypeOk() (*ActivityTypeEnum, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *EngagementType) SetActivityType(v ActivityTypeEnum)`

SetActivityType sets ActivityType field to given value.

### HasActivityType

`func (o *EngagementType) HasActivityType() bool`

HasActivityType returns a boolean if a field has been set.

### SetActivityTypeNil

`func (o *EngagementType) SetActivityTypeNil(b bool)`

 SetActivityTypeNil sets the value for ActivityType to be an explicit nil

### UnsetActivityType
`func (o *EngagementType) UnsetActivityType()`

UnsetActivityType ensures that no value is present for ActivityType, not even an explicit nil
### GetName

`func (o *EngagementType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EngagementType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EngagementType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EngagementType) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EngagementType) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EngagementType) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetId

`func (o *EngagementType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EngagementType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EngagementType) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EngagementType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteId

`func (o *EngagementType) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *EngagementType) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *EngagementType) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *EngagementType) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *EngagementType) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *EngagementType) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetModifiedAt

`func (o *EngagementType) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *EngagementType) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *EngagementType) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *EngagementType) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetRemoteFields

`func (o *EngagementType) GetRemoteFields() []RemoteField`

GetRemoteFields returns the RemoteFields field if non-nil, zero value otherwise.

### GetRemoteFieldsOk

`func (o *EngagementType) GetRemoteFieldsOk() (*[]RemoteField, bool)`

GetRemoteFieldsOk returns a tuple with the RemoteFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteFields

`func (o *EngagementType) SetRemoteFields(v []RemoteField)`

SetRemoteFields sets RemoteFields field to given value.

### HasRemoteFields

`func (o *EngagementType) HasRemoteFields() bool`

HasRemoteFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



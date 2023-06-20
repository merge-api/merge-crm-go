# PatchedEngagementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | Pointer to **NullableString** | The engagement&#39;s owner. | [optional] 
**Content** | Pointer to **NullableString** | The engagement&#39;s content. | [optional] 
**Subject** | Pointer to **NullableString** | The engagement&#39;s subject. | [optional] 
**Direction** | Pointer to [**NullableDirectionEnum**](DirectionEnum.md) | The engagement&#39;s direction.  * &#x60;INBOUND&#x60; - INBOUND * &#x60;OUTBOUND&#x60; - OUTBOUND | [optional] 
**EngagementType** | Pointer to **NullableString** | The engagement type of the engagement. | [optional] 
**StartTime** | Pointer to **NullableTime** | The time at which the engagement started. | [optional] 
**EndTime** | Pointer to **NullableTime** | The time at which the engagement ended. | [optional] 
**Account** | Pointer to **NullableString** | The account of the engagement. | [optional] 
**Contacts** | Pointer to **[]string** |  | [optional] 
**IntegrationParams** | Pointer to **map[string]interface{}** |  | [optional] 
**LinkedAccountParams** | Pointer to **map[string]interface{}** |  | [optional] 
**RemoteFields** | Pointer to [**[]RemoteFieldRequest**](RemoteFieldRequest.md) |  | [optional] 

## Methods

### NewPatchedEngagementRequest

`func NewPatchedEngagementRequest() *PatchedEngagementRequest`

NewPatchedEngagementRequest instantiates a new PatchedEngagementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedEngagementRequestWithDefaults

`func NewPatchedEngagementRequestWithDefaults() *PatchedEngagementRequest`

NewPatchedEngagementRequestWithDefaults instantiates a new PatchedEngagementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *PatchedEngagementRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedEngagementRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedEngagementRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedEngagementRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedEngagementRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedEngagementRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetContent

`func (o *PatchedEngagementRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PatchedEngagementRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PatchedEngagementRequest) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *PatchedEngagementRequest) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *PatchedEngagementRequest) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *PatchedEngagementRequest) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetSubject

`func (o *PatchedEngagementRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *PatchedEngagementRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *PatchedEngagementRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *PatchedEngagementRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *PatchedEngagementRequest) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *PatchedEngagementRequest) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetDirection

`func (o *PatchedEngagementRequest) GetDirection() DirectionEnum`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *PatchedEngagementRequest) GetDirectionOk() (*DirectionEnum, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *PatchedEngagementRequest) SetDirection(v DirectionEnum)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *PatchedEngagementRequest) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### SetDirectionNil

`func (o *PatchedEngagementRequest) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *PatchedEngagementRequest) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil
### GetEngagementType

`func (o *PatchedEngagementRequest) GetEngagementType() string`

GetEngagementType returns the EngagementType field if non-nil, zero value otherwise.

### GetEngagementTypeOk

`func (o *PatchedEngagementRequest) GetEngagementTypeOk() (*string, bool)`

GetEngagementTypeOk returns a tuple with the EngagementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementType

`func (o *PatchedEngagementRequest) SetEngagementType(v string)`

SetEngagementType sets EngagementType field to given value.

### HasEngagementType

`func (o *PatchedEngagementRequest) HasEngagementType() bool`

HasEngagementType returns a boolean if a field has been set.

### SetEngagementTypeNil

`func (o *PatchedEngagementRequest) SetEngagementTypeNil(b bool)`

 SetEngagementTypeNil sets the value for EngagementType to be an explicit nil

### UnsetEngagementType
`func (o *PatchedEngagementRequest) UnsetEngagementType()`

UnsetEngagementType ensures that no value is present for EngagementType, not even an explicit nil
### GetStartTime

`func (o *PatchedEngagementRequest) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *PatchedEngagementRequest) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *PatchedEngagementRequest) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *PatchedEngagementRequest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *PatchedEngagementRequest) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *PatchedEngagementRequest) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetEndTime

`func (o *PatchedEngagementRequest) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *PatchedEngagementRequest) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *PatchedEngagementRequest) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *PatchedEngagementRequest) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *PatchedEngagementRequest) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *PatchedEngagementRequest) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetAccount

`func (o *PatchedEngagementRequest) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PatchedEngagementRequest) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PatchedEngagementRequest) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PatchedEngagementRequest) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *PatchedEngagementRequest) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *PatchedEngagementRequest) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetContacts

`func (o *PatchedEngagementRequest) GetContacts() []string`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *PatchedEngagementRequest) GetContactsOk() (*[]string, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *PatchedEngagementRequest) SetContacts(v []string)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *PatchedEngagementRequest) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetIntegrationParams

`func (o *PatchedEngagementRequest) GetIntegrationParams() map[string]interface{}`

GetIntegrationParams returns the IntegrationParams field if non-nil, zero value otherwise.

### GetIntegrationParamsOk

`func (o *PatchedEngagementRequest) GetIntegrationParamsOk() (*map[string]interface{}, bool)`

GetIntegrationParamsOk returns a tuple with the IntegrationParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationParams

`func (o *PatchedEngagementRequest) SetIntegrationParams(v map[string]interface{})`

SetIntegrationParams sets IntegrationParams field to given value.

### HasIntegrationParams

`func (o *PatchedEngagementRequest) HasIntegrationParams() bool`

HasIntegrationParams returns a boolean if a field has been set.

### SetIntegrationParamsNil

`func (o *PatchedEngagementRequest) SetIntegrationParamsNil(b bool)`

 SetIntegrationParamsNil sets the value for IntegrationParams to be an explicit nil

### UnsetIntegrationParams
`func (o *PatchedEngagementRequest) UnsetIntegrationParams()`

UnsetIntegrationParams ensures that no value is present for IntegrationParams, not even an explicit nil
### GetLinkedAccountParams

`func (o *PatchedEngagementRequest) GetLinkedAccountParams() map[string]interface{}`

GetLinkedAccountParams returns the LinkedAccountParams field if non-nil, zero value otherwise.

### GetLinkedAccountParamsOk

`func (o *PatchedEngagementRequest) GetLinkedAccountParamsOk() (*map[string]interface{}, bool)`

GetLinkedAccountParamsOk returns a tuple with the LinkedAccountParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountParams

`func (o *PatchedEngagementRequest) SetLinkedAccountParams(v map[string]interface{})`

SetLinkedAccountParams sets LinkedAccountParams field to given value.

### HasLinkedAccountParams

`func (o *PatchedEngagementRequest) HasLinkedAccountParams() bool`

HasLinkedAccountParams returns a boolean if a field has been set.

### SetLinkedAccountParamsNil

`func (o *PatchedEngagementRequest) SetLinkedAccountParamsNil(b bool)`

 SetLinkedAccountParamsNil sets the value for LinkedAccountParams to be an explicit nil

### UnsetLinkedAccountParams
`func (o *PatchedEngagementRequest) UnsetLinkedAccountParams()`

UnsetLinkedAccountParams ensures that no value is present for LinkedAccountParams, not even an explicit nil
### GetRemoteFields

`func (o *PatchedEngagementRequest) GetRemoteFields() []RemoteFieldRequest`

GetRemoteFields returns the RemoteFields field if non-nil, zero value otherwise.

### GetRemoteFieldsOk

`func (o *PatchedEngagementRequest) GetRemoteFieldsOk() (*[]RemoteFieldRequest, bool)`

GetRemoteFieldsOk returns a tuple with the RemoteFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteFields

`func (o *PatchedEngagementRequest) SetRemoteFields(v []RemoteFieldRequest)`

SetRemoteFields sets RemoteFields field to given value.

### HasRemoteFields

`func (o *PatchedEngagementRequest) HasRemoteFields() bool`

HasRemoteFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



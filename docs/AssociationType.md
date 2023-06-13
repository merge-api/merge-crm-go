# AssociationType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceObjectClass** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**TargetObjectClasses** | Pointer to [**[]AssociationSubType**](AssociationSubType.md) |  | [optional] [readonly] 
**RemoteKeyName** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Cardinality** | Pointer to [**NullableCardinalityEnum**](CardinalityEnum.md) |  | [optional] 
**IsRequired** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**ModifiedAt** | Pointer to **time.Time** | This is the datetime that this object was last updated by Merge | [optional] [readonly] 

## Methods

### NewAssociationType

`func NewAssociationType() *AssociationType`

NewAssociationType instantiates a new AssociationType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssociationTypeWithDefaults

`func NewAssociationTypeWithDefaults() *AssociationType`

NewAssociationTypeWithDefaults instantiates a new AssociationType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceObjectClass

`func (o *AssociationType) GetSourceObjectClass() map[string]interface{}`

GetSourceObjectClass returns the SourceObjectClass field if non-nil, zero value otherwise.

### GetSourceObjectClassOk

`func (o *AssociationType) GetSourceObjectClassOk() (*map[string]interface{}, bool)`

GetSourceObjectClassOk returns a tuple with the SourceObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceObjectClass

`func (o *AssociationType) SetSourceObjectClass(v map[string]interface{})`

SetSourceObjectClass sets SourceObjectClass field to given value.

### HasSourceObjectClass

`func (o *AssociationType) HasSourceObjectClass() bool`

HasSourceObjectClass returns a boolean if a field has been set.

### GetTargetObjectClasses

`func (o *AssociationType) GetTargetObjectClasses() []AssociationSubType`

GetTargetObjectClasses returns the TargetObjectClasses field if non-nil, zero value otherwise.

### GetTargetObjectClassesOk

`func (o *AssociationType) GetTargetObjectClassesOk() (*[]AssociationSubType, bool)`

GetTargetObjectClassesOk returns a tuple with the TargetObjectClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetObjectClasses

`func (o *AssociationType) SetTargetObjectClasses(v []AssociationSubType)`

SetTargetObjectClasses sets TargetObjectClasses field to given value.

### HasTargetObjectClasses

`func (o *AssociationType) HasTargetObjectClasses() bool`

HasTargetObjectClasses returns a boolean if a field has been set.

### GetRemoteKeyName

`func (o *AssociationType) GetRemoteKeyName() string`

GetRemoteKeyName returns the RemoteKeyName field if non-nil, zero value otherwise.

### GetRemoteKeyNameOk

`func (o *AssociationType) GetRemoteKeyNameOk() (*string, bool)`

GetRemoteKeyNameOk returns a tuple with the RemoteKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteKeyName

`func (o *AssociationType) SetRemoteKeyName(v string)`

SetRemoteKeyName sets RemoteKeyName field to given value.

### HasRemoteKeyName

`func (o *AssociationType) HasRemoteKeyName() bool`

HasRemoteKeyName returns a boolean if a field has been set.

### SetRemoteKeyNameNil

`func (o *AssociationType) SetRemoteKeyNameNil(b bool)`

 SetRemoteKeyNameNil sets the value for RemoteKeyName to be an explicit nil

### UnsetRemoteKeyName
`func (o *AssociationType) UnsetRemoteKeyName()`

UnsetRemoteKeyName ensures that no value is present for RemoteKeyName, not even an explicit nil
### GetDisplayName

`func (o *AssociationType) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AssociationType) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AssociationType) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AssociationType) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *AssociationType) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *AssociationType) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetCardinality

`func (o *AssociationType) GetCardinality() CardinalityEnum`

GetCardinality returns the Cardinality field if non-nil, zero value otherwise.

### GetCardinalityOk

`func (o *AssociationType) GetCardinalityOk() (*CardinalityEnum, bool)`

GetCardinalityOk returns a tuple with the Cardinality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardinality

`func (o *AssociationType) SetCardinality(v CardinalityEnum)`

SetCardinality sets Cardinality field to given value.

### HasCardinality

`func (o *AssociationType) HasCardinality() bool`

HasCardinality returns a boolean if a field has been set.

### SetCardinalityNil

`func (o *AssociationType) SetCardinalityNil(b bool)`

 SetCardinalityNil sets the value for Cardinality to be an explicit nil

### UnsetCardinality
`func (o *AssociationType) UnsetCardinality()`

UnsetCardinality ensures that no value is present for Cardinality, not even an explicit nil
### GetIsRequired

`func (o *AssociationType) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *AssociationType) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *AssociationType) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *AssociationType) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetId

`func (o *AssociationType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssociationType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssociationType) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssociationType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteId

`func (o *AssociationType) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *AssociationType) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *AssociationType) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *AssociationType) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *AssociationType) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *AssociationType) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetModifiedAt

`func (o *AssociationType) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AssociationType) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AssociationType) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *AssociationType) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



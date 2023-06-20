# Association

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceObject** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**TargetObject** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**AssociationType** | Pointer to **NullableString** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** | This is the datetime that this object was last updated by Merge | [optional] [readonly] 

## Methods

### NewAssociation

`func NewAssociation() *Association`

NewAssociation instantiates a new Association object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssociationWithDefaults

`func NewAssociationWithDefaults() *Association`

NewAssociationWithDefaults instantiates a new Association object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceObject

`func (o *Association) GetSourceObject() map[string]interface{}`

GetSourceObject returns the SourceObject field if non-nil, zero value otherwise.

### GetSourceObjectOk

`func (o *Association) GetSourceObjectOk() (*map[string]interface{}, bool)`

GetSourceObjectOk returns a tuple with the SourceObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceObject

`func (o *Association) SetSourceObject(v map[string]interface{})`

SetSourceObject sets SourceObject field to given value.

### HasSourceObject

`func (o *Association) HasSourceObject() bool`

HasSourceObject returns a boolean if a field has been set.

### SetSourceObjectNil

`func (o *Association) SetSourceObjectNil(b bool)`

 SetSourceObjectNil sets the value for SourceObject to be an explicit nil

### UnsetSourceObject
`func (o *Association) UnsetSourceObject()`

UnsetSourceObject ensures that no value is present for SourceObject, not even an explicit nil
### GetTargetObject

`func (o *Association) GetTargetObject() map[string]interface{}`

GetTargetObject returns the TargetObject field if non-nil, zero value otherwise.

### GetTargetObjectOk

`func (o *Association) GetTargetObjectOk() (*map[string]interface{}, bool)`

GetTargetObjectOk returns a tuple with the TargetObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetObject

`func (o *Association) SetTargetObject(v map[string]interface{})`

SetTargetObject sets TargetObject field to given value.

### HasTargetObject

`func (o *Association) HasTargetObject() bool`

HasTargetObject returns a boolean if a field has been set.

### SetTargetObjectNil

`func (o *Association) SetTargetObjectNil(b bool)`

 SetTargetObjectNil sets the value for TargetObject to be an explicit nil

### UnsetTargetObject
`func (o *Association) UnsetTargetObject()`

UnsetTargetObject ensures that no value is present for TargetObject, not even an explicit nil
### GetAssociationType

`func (o *Association) GetAssociationType() string`

GetAssociationType returns the AssociationType field if non-nil, zero value otherwise.

### GetAssociationTypeOk

`func (o *Association) GetAssociationTypeOk() (*string, bool)`

GetAssociationTypeOk returns a tuple with the AssociationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationType

`func (o *Association) SetAssociationType(v string)`

SetAssociationType sets AssociationType field to given value.

### HasAssociationType

`func (o *Association) HasAssociationType() bool`

HasAssociationType returns a boolean if a field has been set.

### SetAssociationTypeNil

`func (o *Association) SetAssociationTypeNil(b bool)`

 SetAssociationTypeNil sets the value for AssociationType to be an explicit nil

### UnsetAssociationType
`func (o *Association) UnsetAssociationType()`

UnsetAssociationType ensures that no value is present for AssociationType, not even an explicit nil
### GetModifiedAt

`func (o *Association) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Association) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Association) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Association) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



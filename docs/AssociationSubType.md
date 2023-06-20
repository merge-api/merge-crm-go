# AssociationSubType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] [readonly] 
**OriginType** | Pointer to **NullableString** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** | This is the datetime that this object was last updated by Merge | [optional] [readonly] 

## Methods

### NewAssociationSubType

`func NewAssociationSubType() *AssociationSubType`

NewAssociationSubType instantiates a new AssociationSubType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssociationSubTypeWithDefaults

`func NewAssociationSubTypeWithDefaults() *AssociationSubType`

NewAssociationSubTypeWithDefaults instantiates a new AssociationSubType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AssociationSubType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssociationSubType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssociationSubType) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssociationSubType) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AssociationSubType) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AssociationSubType) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetOriginType

`func (o *AssociationSubType) GetOriginType() string`

GetOriginType returns the OriginType field if non-nil, zero value otherwise.

### GetOriginTypeOk

`func (o *AssociationSubType) GetOriginTypeOk() (*string, bool)`

GetOriginTypeOk returns a tuple with the OriginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginType

`func (o *AssociationSubType) SetOriginType(v string)`

SetOriginType sets OriginType field to given value.

### HasOriginType

`func (o *AssociationSubType) HasOriginType() bool`

HasOriginType returns a boolean if a field has been set.

### SetOriginTypeNil

`func (o *AssociationSubType) SetOriginTypeNil(b bool)`

 SetOriginTypeNil sets the value for OriginType to be an explicit nil

### UnsetOriginType
`func (o *AssociationSubType) UnsetOriginType()`

UnsetOriginType ensures that no value is present for OriginType, not even an explicit nil
### GetModifiedAt

`func (o *AssociationSubType) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AssociationSubType) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AssociationSubType) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *AssociationSubType) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



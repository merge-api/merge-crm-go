# CustomObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectClass** | Pointer to **NullableString** |  | [optional] 
**Fields** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** | This is the datetime that this object was last updated by Merge | [optional] [readonly] 
**RemoteFields** | Pointer to [**[]RemoteField**](RemoteField.md) |  | [optional] [readonly] 

## Methods

### NewCustomObject

`func NewCustomObject() *CustomObject`

NewCustomObject instantiates a new CustomObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomObjectWithDefaults

`func NewCustomObjectWithDefaults() *CustomObject`

NewCustomObjectWithDefaults instantiates a new CustomObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectClass

`func (o *CustomObject) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *CustomObject) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *CustomObject) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *CustomObject) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### SetObjectClassNil

`func (o *CustomObject) SetObjectClassNil(b bool)`

 SetObjectClassNil sets the value for ObjectClass to be an explicit nil

### UnsetObjectClass
`func (o *CustomObject) UnsetObjectClass()`

UnsetObjectClass ensures that no value is present for ObjectClass, not even an explicit nil
### GetFields

`func (o *CustomObject) GetFields() map[string]interface{}`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *CustomObject) GetFieldsOk() (*map[string]interface{}, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *CustomObject) SetFields(v map[string]interface{})`

SetFields sets Fields field to given value.

### HasFields

`func (o *CustomObject) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetRemoteId

`func (o *CustomObject) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *CustomObject) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *CustomObject) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *CustomObject) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *CustomObject) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *CustomObject) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetId

`func (o *CustomObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *CustomObject) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *CustomObject) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *CustomObject) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *CustomObject) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetRemoteFields

`func (o *CustomObject) GetRemoteFields() []RemoteField`

GetRemoteFields returns the RemoteFields field if non-nil, zero value otherwise.

### GetRemoteFieldsOk

`func (o *CustomObject) GetRemoteFieldsOk() (*[]RemoteField, bool)`

GetRemoteFieldsOk returns a tuple with the RemoteFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteFields

`func (o *CustomObject) SetRemoteFields(v []RemoteField)`

SetRemoteFields sets RemoteFields field to given value.

### HasRemoteFields

`func (o *CustomObject) HasRemoteFields() bool`

HasRemoteFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



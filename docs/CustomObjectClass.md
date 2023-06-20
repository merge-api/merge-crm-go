# CustomObjectClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] [readonly] 
**Fields** | Pointer to [**[]RemoteFieldClassForCustomObjectClass**](RemoteFieldClassForCustomObjectClass.md) |  | [optional] [readonly] 
**AssociationTypes** | Pointer to **[]map[string]interface{}** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewCustomObjectClass

`func NewCustomObjectClass() *CustomObjectClass`

NewCustomObjectClass instantiates a new CustomObjectClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomObjectClassWithDefaults

`func NewCustomObjectClassWithDefaults() *CustomObjectClass`

NewCustomObjectClassWithDefaults instantiates a new CustomObjectClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomObjectClass) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomObjectClass) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomObjectClass) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomObjectClass) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CustomObjectClass) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CustomObjectClass) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *CustomObjectClass) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomObjectClass) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomObjectClass) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomObjectClass) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CustomObjectClass) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CustomObjectClass) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLabels

`func (o *CustomObjectClass) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CustomObjectClass) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CustomObjectClass) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CustomObjectClass) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetFields

`func (o *CustomObjectClass) GetFields() []RemoteFieldClassForCustomObjectClass`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *CustomObjectClass) GetFieldsOk() (*[]RemoteFieldClassForCustomObjectClass, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *CustomObjectClass) SetFields(v []RemoteFieldClassForCustomObjectClass)`

SetFields sets Fields field to given value.

### HasFields

`func (o *CustomObjectClass) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetAssociationTypes

`func (o *CustomObjectClass) GetAssociationTypes() []map[string]interface{}`

GetAssociationTypes returns the AssociationTypes field if non-nil, zero value otherwise.

### GetAssociationTypesOk

`func (o *CustomObjectClass) GetAssociationTypesOk() (*[]map[string]interface{}, bool)`

GetAssociationTypesOk returns a tuple with the AssociationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationTypes

`func (o *CustomObjectClass) SetAssociationTypes(v []map[string]interface{})`

SetAssociationTypes sets AssociationTypes field to given value.

### HasAssociationTypes

`func (o *CustomObjectClass) HasAssociationTypes() bool`

HasAssociationTypes returns a boolean if a field has been set.

### SetAssociationTypesNil

`func (o *CustomObjectClass) SetAssociationTypesNil(b bool)`

 SetAssociationTypesNil sets the value for AssociationTypes to be an explicit nil

### UnsetAssociationTypes
`func (o *CustomObjectClass) UnsetAssociationTypes()`

UnsetAssociationTypes ensures that no value is present for AssociationTypes, not even an explicit nil
### GetId

`func (o *CustomObjectClass) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomObjectClass) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomObjectClass) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomObjectClass) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteId

`func (o *CustomObjectClass) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *CustomObjectClass) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *CustomObjectClass) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *CustomObjectClass) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *CustomObjectClass) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *CustomObjectClass) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetModifiedAt

`func (o *CustomObjectClass) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *CustomObjectClass) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *CustomObjectClass) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *CustomObjectClass) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



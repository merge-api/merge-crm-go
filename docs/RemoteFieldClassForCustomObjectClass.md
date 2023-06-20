# RemoteFieldClassForCustomObjectClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**RemoteKeyName** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**IsRequired** | Pointer to **bool** |  | [optional] 
**FieldType** | Pointer to [**FieldTypeEnum**](FieldTypeEnum.md) |  | [optional] [readonly] 
**FieldFormat** | Pointer to [**FieldFormatEnum**](FieldFormatEnum.md) |  | [optional] [readonly] 
**FieldChoices** | Pointer to **[]string** |  | [optional] [readonly] 
**ItemSchema** | Pointer to [**NullableRemoteFieldClassForCustomObjectClassItemSchema**](RemoteFieldClassForCustomObjectClassItemSchema.md) |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** | This is the datetime that this object was last updated by Merge | [optional] [readonly] 

## Methods

### NewRemoteFieldClassForCustomObjectClass

`func NewRemoteFieldClassForCustomObjectClass() *RemoteFieldClassForCustomObjectClass`

NewRemoteFieldClassForCustomObjectClass instantiates a new RemoteFieldClassForCustomObjectClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteFieldClassForCustomObjectClassWithDefaults

`func NewRemoteFieldClassForCustomObjectClassWithDefaults() *RemoteFieldClassForCustomObjectClass`

NewRemoteFieldClassForCustomObjectClassWithDefaults instantiates a new RemoteFieldClassForCustomObjectClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *RemoteFieldClassForCustomObjectClass) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RemoteFieldClassForCustomObjectClass) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *RemoteFieldClassForCustomObjectClass) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *RemoteFieldClassForCustomObjectClass) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *RemoteFieldClassForCustomObjectClass) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *RemoteFieldClassForCustomObjectClass) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetRemoteKeyName

`func (o *RemoteFieldClassForCustomObjectClass) GetRemoteKeyName() string`

GetRemoteKeyName returns the RemoteKeyName field if non-nil, zero value otherwise.

### GetRemoteKeyNameOk

`func (o *RemoteFieldClassForCustomObjectClass) GetRemoteKeyNameOk() (*string, bool)`

GetRemoteKeyNameOk returns a tuple with the RemoteKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteKeyName

`func (o *RemoteFieldClassForCustomObjectClass) SetRemoteKeyName(v string)`

SetRemoteKeyName sets RemoteKeyName field to given value.

### HasRemoteKeyName

`func (o *RemoteFieldClassForCustomObjectClass) HasRemoteKeyName() bool`

HasRemoteKeyName returns a boolean if a field has been set.

### SetRemoteKeyNameNil

`func (o *RemoteFieldClassForCustomObjectClass) SetRemoteKeyNameNil(b bool)`

 SetRemoteKeyNameNil sets the value for RemoteKeyName to be an explicit nil

### UnsetRemoteKeyName
`func (o *RemoteFieldClassForCustomObjectClass) UnsetRemoteKeyName()`

UnsetRemoteKeyName ensures that no value is present for RemoteKeyName, not even an explicit nil
### GetDescription

`func (o *RemoteFieldClassForCustomObjectClass) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RemoteFieldClassForCustomObjectClass) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RemoteFieldClassForCustomObjectClass) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RemoteFieldClassForCustomObjectClass) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *RemoteFieldClassForCustomObjectClass) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RemoteFieldClassForCustomObjectClass) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsRequired

`func (o *RemoteFieldClassForCustomObjectClass) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *RemoteFieldClassForCustomObjectClass) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *RemoteFieldClassForCustomObjectClass) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *RemoteFieldClassForCustomObjectClass) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetFieldType

`func (o *RemoteFieldClassForCustomObjectClass) GetFieldType() FieldTypeEnum`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *RemoteFieldClassForCustomObjectClass) GetFieldTypeOk() (*FieldTypeEnum, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *RemoteFieldClassForCustomObjectClass) SetFieldType(v FieldTypeEnum)`

SetFieldType sets FieldType field to given value.

### HasFieldType

`func (o *RemoteFieldClassForCustomObjectClass) HasFieldType() bool`

HasFieldType returns a boolean if a field has been set.

### GetFieldFormat

`func (o *RemoteFieldClassForCustomObjectClass) GetFieldFormat() FieldFormatEnum`

GetFieldFormat returns the FieldFormat field if non-nil, zero value otherwise.

### GetFieldFormatOk

`func (o *RemoteFieldClassForCustomObjectClass) GetFieldFormatOk() (*FieldFormatEnum, bool)`

GetFieldFormatOk returns a tuple with the FieldFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldFormat

`func (o *RemoteFieldClassForCustomObjectClass) SetFieldFormat(v FieldFormatEnum)`

SetFieldFormat sets FieldFormat field to given value.

### HasFieldFormat

`func (o *RemoteFieldClassForCustomObjectClass) HasFieldFormat() bool`

HasFieldFormat returns a boolean if a field has been set.

### GetFieldChoices

`func (o *RemoteFieldClassForCustomObjectClass) GetFieldChoices() []string`

GetFieldChoices returns the FieldChoices field if non-nil, zero value otherwise.

### GetFieldChoicesOk

`func (o *RemoteFieldClassForCustomObjectClass) GetFieldChoicesOk() (*[]string, bool)`

GetFieldChoicesOk returns a tuple with the FieldChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldChoices

`func (o *RemoteFieldClassForCustomObjectClass) SetFieldChoices(v []string)`

SetFieldChoices sets FieldChoices field to given value.

### HasFieldChoices

`func (o *RemoteFieldClassForCustomObjectClass) HasFieldChoices() bool`

HasFieldChoices returns a boolean if a field has been set.

### SetFieldChoicesNil

`func (o *RemoteFieldClassForCustomObjectClass) SetFieldChoicesNil(b bool)`

 SetFieldChoicesNil sets the value for FieldChoices to be an explicit nil

### UnsetFieldChoices
`func (o *RemoteFieldClassForCustomObjectClass) UnsetFieldChoices()`

UnsetFieldChoices ensures that no value is present for FieldChoices, not even an explicit nil
### GetItemSchema

`func (o *RemoteFieldClassForCustomObjectClass) GetItemSchema() RemoteFieldClassForCustomObjectClassItemSchema`

GetItemSchema returns the ItemSchema field if non-nil, zero value otherwise.

### GetItemSchemaOk

`func (o *RemoteFieldClassForCustomObjectClass) GetItemSchemaOk() (*RemoteFieldClassForCustomObjectClassItemSchema, bool)`

GetItemSchemaOk returns a tuple with the ItemSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSchema

`func (o *RemoteFieldClassForCustomObjectClass) SetItemSchema(v RemoteFieldClassForCustomObjectClassItemSchema)`

SetItemSchema sets ItemSchema field to given value.

### HasItemSchema

`func (o *RemoteFieldClassForCustomObjectClass) HasItemSchema() bool`

HasItemSchema returns a boolean if a field has been set.

### SetItemSchemaNil

`func (o *RemoteFieldClassForCustomObjectClass) SetItemSchemaNil(b bool)`

 SetItemSchemaNil sets the value for ItemSchema to be an explicit nil

### UnsetItemSchema
`func (o *RemoteFieldClassForCustomObjectClass) UnsetItemSchema()`

UnsetItemSchema ensures that no value is present for ItemSchema, not even an explicit nil
### GetModifiedAt

`func (o *RemoteFieldClassForCustomObjectClass) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *RemoteFieldClassForCustomObjectClass) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *RemoteFieldClassForCustomObjectClass) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *RemoteFieldClassForCustomObjectClass) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



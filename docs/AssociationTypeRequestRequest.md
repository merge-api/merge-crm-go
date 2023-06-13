# AssociationTypeRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceObjectClass** | [**ObjectClassDescriptionRequest**](ObjectClassDescriptionRequest.md) |  | 
**TargetObjectClasses** | [**[]ObjectClassDescriptionRequest**](ObjectClassDescriptionRequest.md) |  | 
**RemoteKeyName** | **string** |  | 
**DisplayName** | Pointer to **string** |  | [optional] 
**Cardinality** | Pointer to [**CardinalityEnum**](CardinalityEnum.md) |  | [optional] 
**IsRequired** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewAssociationTypeRequestRequest

`func NewAssociationTypeRequestRequest(sourceObjectClass ObjectClassDescriptionRequest, targetObjectClasses []ObjectClassDescriptionRequest, remoteKeyName string, ) *AssociationTypeRequestRequest`

NewAssociationTypeRequestRequest instantiates a new AssociationTypeRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssociationTypeRequestRequestWithDefaults

`func NewAssociationTypeRequestRequestWithDefaults() *AssociationTypeRequestRequest`

NewAssociationTypeRequestRequestWithDefaults instantiates a new AssociationTypeRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceObjectClass

`func (o *AssociationTypeRequestRequest) GetSourceObjectClass() ObjectClassDescriptionRequest`

GetSourceObjectClass returns the SourceObjectClass field if non-nil, zero value otherwise.

### GetSourceObjectClassOk

`func (o *AssociationTypeRequestRequest) GetSourceObjectClassOk() (*ObjectClassDescriptionRequest, bool)`

GetSourceObjectClassOk returns a tuple with the SourceObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceObjectClass

`func (o *AssociationTypeRequestRequest) SetSourceObjectClass(v ObjectClassDescriptionRequest)`

SetSourceObjectClass sets SourceObjectClass field to given value.


### GetTargetObjectClasses

`func (o *AssociationTypeRequestRequest) GetTargetObjectClasses() []ObjectClassDescriptionRequest`

GetTargetObjectClasses returns the TargetObjectClasses field if non-nil, zero value otherwise.

### GetTargetObjectClassesOk

`func (o *AssociationTypeRequestRequest) GetTargetObjectClassesOk() (*[]ObjectClassDescriptionRequest, bool)`

GetTargetObjectClassesOk returns a tuple with the TargetObjectClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetObjectClasses

`func (o *AssociationTypeRequestRequest) SetTargetObjectClasses(v []ObjectClassDescriptionRequest)`

SetTargetObjectClasses sets TargetObjectClasses field to given value.


### GetRemoteKeyName

`func (o *AssociationTypeRequestRequest) GetRemoteKeyName() string`

GetRemoteKeyName returns the RemoteKeyName field if non-nil, zero value otherwise.

### GetRemoteKeyNameOk

`func (o *AssociationTypeRequestRequest) GetRemoteKeyNameOk() (*string, bool)`

GetRemoteKeyNameOk returns a tuple with the RemoteKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteKeyName

`func (o *AssociationTypeRequestRequest) SetRemoteKeyName(v string)`

SetRemoteKeyName sets RemoteKeyName field to given value.


### GetDisplayName

`func (o *AssociationTypeRequestRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AssociationTypeRequestRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AssociationTypeRequestRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AssociationTypeRequestRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCardinality

`func (o *AssociationTypeRequestRequest) GetCardinality() CardinalityEnum`

GetCardinality returns the Cardinality field if non-nil, zero value otherwise.

### GetCardinalityOk

`func (o *AssociationTypeRequestRequest) GetCardinalityOk() (*CardinalityEnum, bool)`

GetCardinalityOk returns a tuple with the Cardinality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardinality

`func (o *AssociationTypeRequestRequest) SetCardinality(v CardinalityEnum)`

SetCardinality sets Cardinality field to given value.

### HasCardinality

`func (o *AssociationTypeRequestRequest) HasCardinality() bool`

HasCardinality returns a boolean if a field has been set.

### GetIsRequired

`func (o *AssociationTypeRequestRequest) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *AssociationTypeRequestRequest) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *AssociationTypeRequestRequest) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *AssociationTypeRequestRequest) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# RemoteField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteFieldClass** | [**RemoteFieldClass**](RemoteFieldClass.md) |  | 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRemoteField

`func NewRemoteField(remoteFieldClass RemoteFieldClass, ) *RemoteField`

NewRemoteField instantiates a new RemoteField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteFieldWithDefaults

`func NewRemoteFieldWithDefaults() *RemoteField`

NewRemoteFieldWithDefaults instantiates a new RemoteField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteFieldClass

`func (o *RemoteField) GetRemoteFieldClass() RemoteFieldClass`

GetRemoteFieldClass returns the RemoteFieldClass field if non-nil, zero value otherwise.

### GetRemoteFieldClassOk

`func (o *RemoteField) GetRemoteFieldClassOk() (*RemoteFieldClass, bool)`

GetRemoteFieldClassOk returns a tuple with the RemoteFieldClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteFieldClass

`func (o *RemoteField) SetRemoteFieldClass(v RemoteFieldClass)`

SetRemoteFieldClass sets RemoteFieldClass field to given value.


### GetValue

`func (o *RemoteField) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RemoteField) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RemoteField) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *RemoteField) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



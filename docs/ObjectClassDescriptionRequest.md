# ObjectClassDescriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**OriginType** | [**OriginTypeEnum**](OriginTypeEnum.md) |  | 

## Methods

### NewObjectClassDescriptionRequest

`func NewObjectClassDescriptionRequest(id string, originType OriginTypeEnum, ) *ObjectClassDescriptionRequest`

NewObjectClassDescriptionRequest instantiates a new ObjectClassDescriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectClassDescriptionRequestWithDefaults

`func NewObjectClassDescriptionRequestWithDefaults() *ObjectClassDescriptionRequest`

NewObjectClassDescriptionRequestWithDefaults instantiates a new ObjectClassDescriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ObjectClassDescriptionRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectClassDescriptionRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectClassDescriptionRequest) SetId(v string)`

SetId sets Id field to given value.


### GetOriginType

`func (o *ObjectClassDescriptionRequest) GetOriginType() OriginTypeEnum`

GetOriginType returns the OriginType field if non-nil, zero value otherwise.

### GetOriginTypeOk

`func (o *ObjectClassDescriptionRequest) GetOriginTypeOk() (*OriginTypeEnum, bool)`

GetOriginTypeOk returns a tuple with the OriginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginType

`func (o *ObjectClassDescriptionRequest) SetOriginType(v OriginTypeEnum)`

SetOriginType sets OriginType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



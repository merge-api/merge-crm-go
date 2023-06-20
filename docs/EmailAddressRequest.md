# EmailAddressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | Pointer to **NullableString** | The email address. | [optional] 
**EmailAddressType** | Pointer to **NullableString** | The email address&#39;s type. | [optional] 
**IntegrationParams** | Pointer to **map[string]interface{}** |  | [optional] 
**LinkedAccountParams** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewEmailAddressRequest

`func NewEmailAddressRequest() *EmailAddressRequest`

NewEmailAddressRequest instantiates a new EmailAddressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailAddressRequestWithDefaults

`func NewEmailAddressRequestWithDefaults() *EmailAddressRequest`

NewEmailAddressRequestWithDefaults instantiates a new EmailAddressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *EmailAddressRequest) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *EmailAddressRequest) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *EmailAddressRequest) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *EmailAddressRequest) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *EmailAddressRequest) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *EmailAddressRequest) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetEmailAddressType

`func (o *EmailAddressRequest) GetEmailAddressType() string`

GetEmailAddressType returns the EmailAddressType field if non-nil, zero value otherwise.

### GetEmailAddressTypeOk

`func (o *EmailAddressRequest) GetEmailAddressTypeOk() (*string, bool)`

GetEmailAddressTypeOk returns a tuple with the EmailAddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddressType

`func (o *EmailAddressRequest) SetEmailAddressType(v string)`

SetEmailAddressType sets EmailAddressType field to given value.

### HasEmailAddressType

`func (o *EmailAddressRequest) HasEmailAddressType() bool`

HasEmailAddressType returns a boolean if a field has been set.

### SetEmailAddressTypeNil

`func (o *EmailAddressRequest) SetEmailAddressTypeNil(b bool)`

 SetEmailAddressTypeNil sets the value for EmailAddressType to be an explicit nil

### UnsetEmailAddressType
`func (o *EmailAddressRequest) UnsetEmailAddressType()`

UnsetEmailAddressType ensures that no value is present for EmailAddressType, not even an explicit nil
### GetIntegrationParams

`func (o *EmailAddressRequest) GetIntegrationParams() map[string]interface{}`

GetIntegrationParams returns the IntegrationParams field if non-nil, zero value otherwise.

### GetIntegrationParamsOk

`func (o *EmailAddressRequest) GetIntegrationParamsOk() (*map[string]interface{}, bool)`

GetIntegrationParamsOk returns a tuple with the IntegrationParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationParams

`func (o *EmailAddressRequest) SetIntegrationParams(v map[string]interface{})`

SetIntegrationParams sets IntegrationParams field to given value.

### HasIntegrationParams

`func (o *EmailAddressRequest) HasIntegrationParams() bool`

HasIntegrationParams returns a boolean if a field has been set.

### SetIntegrationParamsNil

`func (o *EmailAddressRequest) SetIntegrationParamsNil(b bool)`

 SetIntegrationParamsNil sets the value for IntegrationParams to be an explicit nil

### UnsetIntegrationParams
`func (o *EmailAddressRequest) UnsetIntegrationParams()`

UnsetIntegrationParams ensures that no value is present for IntegrationParams, not even an explicit nil
### GetLinkedAccountParams

`func (o *EmailAddressRequest) GetLinkedAccountParams() map[string]interface{}`

GetLinkedAccountParams returns the LinkedAccountParams field if non-nil, zero value otherwise.

### GetLinkedAccountParamsOk

`func (o *EmailAddressRequest) GetLinkedAccountParamsOk() (*map[string]interface{}, bool)`

GetLinkedAccountParamsOk returns a tuple with the LinkedAccountParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountParams

`func (o *EmailAddressRequest) SetLinkedAccountParams(v map[string]interface{})`

SetLinkedAccountParams sets LinkedAccountParams field to given value.

### HasLinkedAccountParams

`func (o *EmailAddressRequest) HasLinkedAccountParams() bool`

HasLinkedAccountParams returns a boolean if a field has been set.

### SetLinkedAccountParamsNil

`func (o *EmailAddressRequest) SetLinkedAccountParamsNil(b bool)`

 SetLinkedAccountParamsNil sets the value for LinkedAccountParams to be an explicit nil

### UnsetLinkedAccountParams
`func (o *EmailAddressRequest) UnsetLinkedAccountParams()`

UnsetLinkedAccountParams ensures that no value is present for LinkedAccountParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



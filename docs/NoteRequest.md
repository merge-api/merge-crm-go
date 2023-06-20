# NoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | Pointer to **NullableString** | The note&#39;s owner. | [optional] 
**Content** | Pointer to **NullableString** | The note&#39;s content. | [optional] 
**Contact** | Pointer to **NullableString** | The note&#39;s contact. | [optional] 
**Account** | Pointer to **NullableString** | The note&#39;s account. | [optional] 
**Opportunity** | Pointer to **NullableString** | The note&#39;s opportunity. | [optional] 
**IntegrationParams** | Pointer to **map[string]interface{}** |  | [optional] 
**LinkedAccountParams** | Pointer to **map[string]interface{}** |  | [optional] 
**RemoteFields** | Pointer to [**[]RemoteFieldRequest**](RemoteFieldRequest.md) |  | [optional] 

## Methods

### NewNoteRequest

`func NewNoteRequest() *NoteRequest`

NewNoteRequest instantiates a new NoteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteRequestWithDefaults

`func NewNoteRequestWithDefaults() *NoteRequest`

NewNoteRequestWithDefaults instantiates a new NoteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *NoteRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NoteRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NoteRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *NoteRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *NoteRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *NoteRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetContent

`func (o *NoteRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NoteRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NoteRequest) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *NoteRequest) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *NoteRequest) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *NoteRequest) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetContact

`func (o *NoteRequest) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *NoteRequest) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *NoteRequest) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *NoteRequest) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *NoteRequest) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *NoteRequest) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetAccount

`func (o *NoteRequest) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NoteRequest) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NoteRequest) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *NoteRequest) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *NoteRequest) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *NoteRequest) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetOpportunity

`func (o *NoteRequest) GetOpportunity() string`

GetOpportunity returns the Opportunity field if non-nil, zero value otherwise.

### GetOpportunityOk

`func (o *NoteRequest) GetOpportunityOk() (*string, bool)`

GetOpportunityOk returns a tuple with the Opportunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunity

`func (o *NoteRequest) SetOpportunity(v string)`

SetOpportunity sets Opportunity field to given value.

### HasOpportunity

`func (o *NoteRequest) HasOpportunity() bool`

HasOpportunity returns a boolean if a field has been set.

### SetOpportunityNil

`func (o *NoteRequest) SetOpportunityNil(b bool)`

 SetOpportunityNil sets the value for Opportunity to be an explicit nil

### UnsetOpportunity
`func (o *NoteRequest) UnsetOpportunity()`

UnsetOpportunity ensures that no value is present for Opportunity, not even an explicit nil
### GetIntegrationParams

`func (o *NoteRequest) GetIntegrationParams() map[string]interface{}`

GetIntegrationParams returns the IntegrationParams field if non-nil, zero value otherwise.

### GetIntegrationParamsOk

`func (o *NoteRequest) GetIntegrationParamsOk() (*map[string]interface{}, bool)`

GetIntegrationParamsOk returns a tuple with the IntegrationParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationParams

`func (o *NoteRequest) SetIntegrationParams(v map[string]interface{})`

SetIntegrationParams sets IntegrationParams field to given value.

### HasIntegrationParams

`func (o *NoteRequest) HasIntegrationParams() bool`

HasIntegrationParams returns a boolean if a field has been set.

### SetIntegrationParamsNil

`func (o *NoteRequest) SetIntegrationParamsNil(b bool)`

 SetIntegrationParamsNil sets the value for IntegrationParams to be an explicit nil

### UnsetIntegrationParams
`func (o *NoteRequest) UnsetIntegrationParams()`

UnsetIntegrationParams ensures that no value is present for IntegrationParams, not even an explicit nil
### GetLinkedAccountParams

`func (o *NoteRequest) GetLinkedAccountParams() map[string]interface{}`

GetLinkedAccountParams returns the LinkedAccountParams field if non-nil, zero value otherwise.

### GetLinkedAccountParamsOk

`func (o *NoteRequest) GetLinkedAccountParamsOk() (*map[string]interface{}, bool)`

GetLinkedAccountParamsOk returns a tuple with the LinkedAccountParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountParams

`func (o *NoteRequest) SetLinkedAccountParams(v map[string]interface{})`

SetLinkedAccountParams sets LinkedAccountParams field to given value.

### HasLinkedAccountParams

`func (o *NoteRequest) HasLinkedAccountParams() bool`

HasLinkedAccountParams returns a boolean if a field has been set.

### SetLinkedAccountParamsNil

`func (o *NoteRequest) SetLinkedAccountParamsNil(b bool)`

 SetLinkedAccountParamsNil sets the value for LinkedAccountParams to be an explicit nil

### UnsetLinkedAccountParams
`func (o *NoteRequest) UnsetLinkedAccountParams()`

UnsetLinkedAccountParams ensures that no value is present for LinkedAccountParams, not even an explicit nil
### GetRemoteFields

`func (o *NoteRequest) GetRemoteFields() []RemoteFieldRequest`

GetRemoteFields returns the RemoteFields field if non-nil, zero value otherwise.

### GetRemoteFieldsOk

`func (o *NoteRequest) GetRemoteFieldsOk() (*[]RemoteFieldRequest, bool)`

GetRemoteFieldsOk returns a tuple with the RemoteFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteFields

`func (o *NoteRequest) SetRemoteFields(v []RemoteFieldRequest)`

SetRemoteFields sets RemoteFields field to given value.

### HasRemoteFields

`func (o *NoteRequest) HasRemoteFields() bool`

HasRemoteFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PhoneNumberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | Pointer to **NullableString** | The phone number. | [optional] 
**PhoneNumberType** | Pointer to **NullableString** | The phone number&#39;s type. | [optional] 

## Methods

### NewPhoneNumberRequest

`func NewPhoneNumberRequest() *PhoneNumberRequest`

NewPhoneNumberRequest instantiates a new PhoneNumberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneNumberRequestWithDefaults

`func NewPhoneNumberRequestWithDefaults() *PhoneNumberRequest`

NewPhoneNumberRequestWithDefaults instantiates a new PhoneNumberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *PhoneNumberRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *PhoneNumberRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *PhoneNumberRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *PhoneNumberRequest) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *PhoneNumberRequest) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *PhoneNumberRequest) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetPhoneNumberType

`func (o *PhoneNumberRequest) GetPhoneNumberType() string`

GetPhoneNumberType returns the PhoneNumberType field if non-nil, zero value otherwise.

### GetPhoneNumberTypeOk

`func (o *PhoneNumberRequest) GetPhoneNumberTypeOk() (*string, bool)`

GetPhoneNumberTypeOk returns a tuple with the PhoneNumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberType

`func (o *PhoneNumberRequest) SetPhoneNumberType(v string)`

SetPhoneNumberType sets PhoneNumberType field to given value.

### HasPhoneNumberType

`func (o *PhoneNumberRequest) HasPhoneNumberType() bool`

HasPhoneNumberType returns a boolean if a field has been set.

### SetPhoneNumberTypeNil

`func (o *PhoneNumberRequest) SetPhoneNumberTypeNil(b bool)`

 SetPhoneNumberTypeNil sets the value for PhoneNumberType to be an explicit nil

### UnsetPhoneNumberType
`func (o *PhoneNumberRequest) UnsetPhoneNumberType()`

UnsetPhoneNumberType ensures that no value is present for PhoneNumberType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



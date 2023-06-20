# EmailAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | Pointer to **NullableString** | The email address. | [optional] 
**EmailAddressType** | Pointer to **NullableString** | The email address&#39;s type. | [optional] 
**ModifiedAt** | Pointer to **time.Time** | This is the datetime that this object was last updated by Merge | [optional] [readonly] 

## Methods

### NewEmailAddress

`func NewEmailAddress() *EmailAddress`

NewEmailAddress instantiates a new EmailAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailAddressWithDefaults

`func NewEmailAddressWithDefaults() *EmailAddress`

NewEmailAddressWithDefaults instantiates a new EmailAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *EmailAddress) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *EmailAddress) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *EmailAddress) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *EmailAddress) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *EmailAddress) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *EmailAddress) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetEmailAddressType

`func (o *EmailAddress) GetEmailAddressType() string`

GetEmailAddressType returns the EmailAddressType field if non-nil, zero value otherwise.

### GetEmailAddressTypeOk

`func (o *EmailAddress) GetEmailAddressTypeOk() (*string, bool)`

GetEmailAddressTypeOk returns a tuple with the EmailAddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddressType

`func (o *EmailAddress) SetEmailAddressType(v string)`

SetEmailAddressType sets EmailAddressType field to given value.

### HasEmailAddressType

`func (o *EmailAddress) HasEmailAddressType() bool`

HasEmailAddressType returns a boolean if a field has been set.

### SetEmailAddressTypeNil

`func (o *EmailAddress) SetEmailAddressTypeNil(b bool)`

 SetEmailAddressTypeNil sets the value for EmailAddressType to be an explicit nil

### UnsetEmailAddressType
`func (o *EmailAddress) UnsetEmailAddressType()`

UnsetEmailAddressType ensures that no value is present for EmailAddressType, not even an explicit nil
### GetModifiedAt

`func (o *EmailAddress) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *EmailAddress) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *EmailAddress) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *EmailAddress) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



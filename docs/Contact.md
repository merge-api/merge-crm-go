# Contact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**FirstName** | Pointer to **NullableString** | The contact&#39;s first name. | [optional] 
**LastName** | Pointer to **NullableString** | The contact&#39;s last name. | [optional] 
**Account** | Pointer to **NullableString** |  | [optional] 
**Addresses** | Pointer to [**[]Address**](Address.md) |  | [optional] [readonly] 
**EmailAddresses** | Pointer to [**[]EmailAddress**](EmailAddress.md) |  | [optional] [readonly] 
**PhoneNumbers** | Pointer to [**[]PhoneNumber**](PhoneNumber.md) |  | [optional] [readonly] 
**LastActivityAt** | Pointer to **NullableTime** | When the contact&#39;s last activity occurred. | [optional] 
**RemoteCreatedAt** | Pointer to **NullableTime** | When the third party&#39;s contact was created. | [optional] 
**RemoteWasDeleted** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewContact

`func NewContact() *Contact`

NewContact instantiates a new Contact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactWithDefaults

`func NewContactWithDefaults() *Contact`

NewContactWithDefaults instantiates a new Contact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Contact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Contact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Contact) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Contact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteId

`func (o *Contact) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *Contact) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *Contact) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *Contact) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *Contact) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *Contact) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetFirstName

`func (o *Contact) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Contact) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Contact) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Contact) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *Contact) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *Contact) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *Contact) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Contact) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Contact) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Contact) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *Contact) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *Contact) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetAccount

`func (o *Contact) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Contact) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Contact) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Contact) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *Contact) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *Contact) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetAddresses

`func (o *Contact) GetAddresses() []Address`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *Contact) GetAddressesOk() (*[]Address, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *Contact) SetAddresses(v []Address)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *Contact) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetEmailAddresses

`func (o *Contact) GetEmailAddresses() []EmailAddress`

GetEmailAddresses returns the EmailAddresses field if non-nil, zero value otherwise.

### GetEmailAddressesOk

`func (o *Contact) GetEmailAddressesOk() (*[]EmailAddress, bool)`

GetEmailAddressesOk returns a tuple with the EmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddresses

`func (o *Contact) SetEmailAddresses(v []EmailAddress)`

SetEmailAddresses sets EmailAddresses field to given value.

### HasEmailAddresses

`func (o *Contact) HasEmailAddresses() bool`

HasEmailAddresses returns a boolean if a field has been set.

### GetPhoneNumbers

`func (o *Contact) GetPhoneNumbers() []PhoneNumber`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *Contact) GetPhoneNumbersOk() (*[]PhoneNumber, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *Contact) SetPhoneNumbers(v []PhoneNumber)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *Contact) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.

### GetLastActivityAt

`func (o *Contact) GetLastActivityAt() time.Time`

GetLastActivityAt returns the LastActivityAt field if non-nil, zero value otherwise.

### GetLastActivityAtOk

`func (o *Contact) GetLastActivityAtOk() (*time.Time, bool)`

GetLastActivityAtOk returns a tuple with the LastActivityAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityAt

`func (o *Contact) SetLastActivityAt(v time.Time)`

SetLastActivityAt sets LastActivityAt field to given value.

### HasLastActivityAt

`func (o *Contact) HasLastActivityAt() bool`

HasLastActivityAt returns a boolean if a field has been set.

### SetLastActivityAtNil

`func (o *Contact) SetLastActivityAtNil(b bool)`

 SetLastActivityAtNil sets the value for LastActivityAt to be an explicit nil

### UnsetLastActivityAt
`func (o *Contact) UnsetLastActivityAt()`

UnsetLastActivityAt ensures that no value is present for LastActivityAt, not even an explicit nil
### GetRemoteCreatedAt

`func (o *Contact) GetRemoteCreatedAt() time.Time`

GetRemoteCreatedAt returns the RemoteCreatedAt field if non-nil, zero value otherwise.

### GetRemoteCreatedAtOk

`func (o *Contact) GetRemoteCreatedAtOk() (*time.Time, bool)`

GetRemoteCreatedAtOk returns a tuple with the RemoteCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCreatedAt

`func (o *Contact) SetRemoteCreatedAt(v time.Time)`

SetRemoteCreatedAt sets RemoteCreatedAt field to given value.

### HasRemoteCreatedAt

`func (o *Contact) HasRemoteCreatedAt() bool`

HasRemoteCreatedAt returns a boolean if a field has been set.

### SetRemoteCreatedAtNil

`func (o *Contact) SetRemoteCreatedAtNil(b bool)`

 SetRemoteCreatedAtNil sets the value for RemoteCreatedAt to be an explicit nil

### UnsetRemoteCreatedAt
`func (o *Contact) UnsetRemoteCreatedAt()`

UnsetRemoteCreatedAt ensures that no value is present for RemoteCreatedAt, not even an explicit nil
### GetRemoteWasDeleted

`func (o *Contact) GetRemoteWasDeleted() bool`

GetRemoteWasDeleted returns the RemoteWasDeleted field if non-nil, zero value otherwise.

### GetRemoteWasDeletedOk

`func (o *Contact) GetRemoteWasDeletedOk() (*bool, bool)`

GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteWasDeleted

`func (o *Contact) SetRemoteWasDeleted(v bool)`

SetRemoteWasDeleted sets RemoteWasDeleted field to given value.

### HasRemoteWasDeleted

`func (o *Contact) HasRemoteWasDeleted() bool`

HasRemoteWasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



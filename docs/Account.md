# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**Owner** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** | The account&#39;s name. | [optional] 
**Description** | Pointer to **NullableString** | The account&#39;s description. | [optional] 
**Industry** | Pointer to **NullableString** | The account&#39;s industry. | [optional] 
**Website** | Pointer to **NullableString** | The account&#39;s website. | [optional] 
**NumberOfEmployees** | Pointer to **NullableInt32** | The account&#39;s number of employees. | [optional] 
**Addresses** | Pointer to [**[]Address**](Address.md) |  | [optional] [readonly] 
**PhoneNumbers** | Pointer to [**[]PhoneNumber**](PhoneNumber.md) |  | [optional] [readonly] 
**LastActivityAt** | Pointer to **NullableTime** | When the account&#39;s last activity  occurred. | [optional] 
**RemoteUpdatedAt** | Pointer to **NullableTime** | When the third party&#39;s account was updated. | [optional] 
**RemoteCreatedAt** | Pointer to **NullableTime** | When the third party&#39;s account was created. | [optional] 
**RemoteWasDeleted** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewAccount

`func NewAccount() *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Account) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Account) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Account) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Account) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteId

`func (o *Account) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *Account) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *Account) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *Account) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *Account) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *Account) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetOwner

`func (o *Account) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Account) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Account) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Account) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *Account) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *Account) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetName

`func (o *Account) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Account) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Account) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Account) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Account) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Account) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *Account) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Account) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Account) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Account) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Account) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Account) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIndustry

`func (o *Account) GetIndustry() string`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *Account) GetIndustryOk() (*string, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *Account) SetIndustry(v string)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *Account) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### SetIndustryNil

`func (o *Account) SetIndustryNil(b bool)`

 SetIndustryNil sets the value for Industry to be an explicit nil

### UnsetIndustry
`func (o *Account) UnsetIndustry()`

UnsetIndustry ensures that no value is present for Industry, not even an explicit nil
### GetWebsite

`func (o *Account) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *Account) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *Account) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *Account) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### SetWebsiteNil

`func (o *Account) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *Account) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil
### GetNumberOfEmployees

`func (o *Account) GetNumberOfEmployees() int32`

GetNumberOfEmployees returns the NumberOfEmployees field if non-nil, zero value otherwise.

### GetNumberOfEmployeesOk

`func (o *Account) GetNumberOfEmployeesOk() (*int32, bool)`

GetNumberOfEmployeesOk returns a tuple with the NumberOfEmployees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfEmployees

`func (o *Account) SetNumberOfEmployees(v int32)`

SetNumberOfEmployees sets NumberOfEmployees field to given value.

### HasNumberOfEmployees

`func (o *Account) HasNumberOfEmployees() bool`

HasNumberOfEmployees returns a boolean if a field has been set.

### SetNumberOfEmployeesNil

`func (o *Account) SetNumberOfEmployeesNil(b bool)`

 SetNumberOfEmployeesNil sets the value for NumberOfEmployees to be an explicit nil

### UnsetNumberOfEmployees
`func (o *Account) UnsetNumberOfEmployees()`

UnsetNumberOfEmployees ensures that no value is present for NumberOfEmployees, not even an explicit nil
### GetAddresses

`func (o *Account) GetAddresses() []Address`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *Account) GetAddressesOk() (*[]Address, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *Account) SetAddresses(v []Address)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *Account) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetPhoneNumbers

`func (o *Account) GetPhoneNumbers() []PhoneNumber`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *Account) GetPhoneNumbersOk() (*[]PhoneNumber, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *Account) SetPhoneNumbers(v []PhoneNumber)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *Account) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.

### GetLastActivityAt

`func (o *Account) GetLastActivityAt() time.Time`

GetLastActivityAt returns the LastActivityAt field if non-nil, zero value otherwise.

### GetLastActivityAtOk

`func (o *Account) GetLastActivityAtOk() (*time.Time, bool)`

GetLastActivityAtOk returns a tuple with the LastActivityAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityAt

`func (o *Account) SetLastActivityAt(v time.Time)`

SetLastActivityAt sets LastActivityAt field to given value.

### HasLastActivityAt

`func (o *Account) HasLastActivityAt() bool`

HasLastActivityAt returns a boolean if a field has been set.

### SetLastActivityAtNil

`func (o *Account) SetLastActivityAtNil(b bool)`

 SetLastActivityAtNil sets the value for LastActivityAt to be an explicit nil

### UnsetLastActivityAt
`func (o *Account) UnsetLastActivityAt()`

UnsetLastActivityAt ensures that no value is present for LastActivityAt, not even an explicit nil
### GetRemoteUpdatedAt

`func (o *Account) GetRemoteUpdatedAt() time.Time`

GetRemoteUpdatedAt returns the RemoteUpdatedAt field if non-nil, zero value otherwise.

### GetRemoteUpdatedAtOk

`func (o *Account) GetRemoteUpdatedAtOk() (*time.Time, bool)`

GetRemoteUpdatedAtOk returns a tuple with the RemoteUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUpdatedAt

`func (o *Account) SetRemoteUpdatedAt(v time.Time)`

SetRemoteUpdatedAt sets RemoteUpdatedAt field to given value.

### HasRemoteUpdatedAt

`func (o *Account) HasRemoteUpdatedAt() bool`

HasRemoteUpdatedAt returns a boolean if a field has been set.

### SetRemoteUpdatedAtNil

`func (o *Account) SetRemoteUpdatedAtNil(b bool)`

 SetRemoteUpdatedAtNil sets the value for RemoteUpdatedAt to be an explicit nil

### UnsetRemoteUpdatedAt
`func (o *Account) UnsetRemoteUpdatedAt()`

UnsetRemoteUpdatedAt ensures that no value is present for RemoteUpdatedAt, not even an explicit nil
### GetRemoteCreatedAt

`func (o *Account) GetRemoteCreatedAt() time.Time`

GetRemoteCreatedAt returns the RemoteCreatedAt field if non-nil, zero value otherwise.

### GetRemoteCreatedAtOk

`func (o *Account) GetRemoteCreatedAtOk() (*time.Time, bool)`

GetRemoteCreatedAtOk returns a tuple with the RemoteCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCreatedAt

`func (o *Account) SetRemoteCreatedAt(v time.Time)`

SetRemoteCreatedAt sets RemoteCreatedAt field to given value.

### HasRemoteCreatedAt

`func (o *Account) HasRemoteCreatedAt() bool`

HasRemoteCreatedAt returns a boolean if a field has been set.

### SetRemoteCreatedAtNil

`func (o *Account) SetRemoteCreatedAtNil(b bool)`

 SetRemoteCreatedAtNil sets the value for RemoteCreatedAt to be an explicit nil

### UnsetRemoteCreatedAt
`func (o *Account) UnsetRemoteCreatedAt()`

UnsetRemoteCreatedAt ensures that no value is present for RemoteCreatedAt, not even an explicit nil
### GetRemoteWasDeleted

`func (o *Account) GetRemoteWasDeleted() bool`

GetRemoteWasDeleted returns the RemoteWasDeleted field if non-nil, zero value otherwise.

### GetRemoteWasDeletedOk

`func (o *Account) GetRemoteWasDeletedOk() (*bool, bool)`

GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteWasDeleted

`func (o *Account) SetRemoteWasDeleted(v bool)`

SetRemoteWasDeleted sets RemoteWasDeleted field to given value.

### HasRemoteWasDeleted

`func (o *Account) HasRemoteWasDeleted() bool`

HasRemoteWasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



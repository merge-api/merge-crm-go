# Lead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | Pointer to **NullableString** | The lead&#39;s owner. | [optional] 
**LeadSource** | Pointer to **NullableString** | The lead&#39;s source. | [optional] 
**Title** | Pointer to **NullableString** | The lead&#39;s title. | [optional] 
**Company** | Pointer to **NullableString** | The lead&#39;s company. | [optional] 
**FirstName** | Pointer to **NullableString** | The lead&#39;s first name. | [optional] 
**LastName** | Pointer to **NullableString** | The lead&#39;s last name. | [optional] 
**Addresses** | Pointer to [**[]Address**](Address.md) |  | [optional] [readonly] 
**EmailAddresses** | Pointer to [**[]EmailAddress**](EmailAddress.md) |  | [optional] [readonly] 
**PhoneNumbers** | Pointer to [**[]PhoneNumber**](PhoneNumber.md) |  | [optional] [readonly] 
**RemoteUpdatedAt** | Pointer to **NullableTime** | When the third party&#39;s lead was updated. | [optional] 
**RemoteCreatedAt** | Pointer to **NullableTime** | When the third party&#39;s lead was created. | [optional] 
**ConvertedDate** | Pointer to **NullableTime** | When the lead was converted. | [optional] 
**ConvertedContact** | Pointer to **NullableString** | The contact of the converted lead. | [optional] 
**ConvertedAccount** | Pointer to **NullableString** | The account of the converted lead. | [optional] 
**RemoteWasDeleted** | Pointer to **bool** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**FieldMappings** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** | This is the datetime that this object was last updated by Merge | [optional] [readonly] 
**RemoteData** | Pointer to [**[]RemoteData**](RemoteData.md) |  | [optional] [readonly] 
**RemoteFields** | Pointer to [**[]RemoteField**](RemoteField.md) |  | [optional] [readonly] 

## Methods

### NewLead

`func NewLead() *Lead`

NewLead instantiates a new Lead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeadWithDefaults

`func NewLeadWithDefaults() *Lead`

NewLeadWithDefaults instantiates a new Lead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *Lead) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Lead) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Lead) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Lead) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *Lead) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *Lead) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetLeadSource

`func (o *Lead) GetLeadSource() string`

GetLeadSource returns the LeadSource field if non-nil, zero value otherwise.

### GetLeadSourceOk

`func (o *Lead) GetLeadSourceOk() (*string, bool)`

GetLeadSourceOk returns a tuple with the LeadSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSource

`func (o *Lead) SetLeadSource(v string)`

SetLeadSource sets LeadSource field to given value.

### HasLeadSource

`func (o *Lead) HasLeadSource() bool`

HasLeadSource returns a boolean if a field has been set.

### SetLeadSourceNil

`func (o *Lead) SetLeadSourceNil(b bool)`

 SetLeadSourceNil sets the value for LeadSource to be an explicit nil

### UnsetLeadSource
`func (o *Lead) UnsetLeadSource()`

UnsetLeadSource ensures that no value is present for LeadSource, not even an explicit nil
### GetTitle

`func (o *Lead) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Lead) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Lead) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Lead) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *Lead) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *Lead) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetCompany

`func (o *Lead) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Lead) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Lead) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *Lead) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *Lead) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *Lead) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetFirstName

`func (o *Lead) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Lead) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Lead) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Lead) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *Lead) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *Lead) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *Lead) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Lead) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Lead) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Lead) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *Lead) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *Lead) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetAddresses

`func (o *Lead) GetAddresses() []Address`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *Lead) GetAddressesOk() (*[]Address, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *Lead) SetAddresses(v []Address)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *Lead) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetEmailAddresses

`func (o *Lead) GetEmailAddresses() []EmailAddress`

GetEmailAddresses returns the EmailAddresses field if non-nil, zero value otherwise.

### GetEmailAddressesOk

`func (o *Lead) GetEmailAddressesOk() (*[]EmailAddress, bool)`

GetEmailAddressesOk returns a tuple with the EmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddresses

`func (o *Lead) SetEmailAddresses(v []EmailAddress)`

SetEmailAddresses sets EmailAddresses field to given value.

### HasEmailAddresses

`func (o *Lead) HasEmailAddresses() bool`

HasEmailAddresses returns a boolean if a field has been set.

### GetPhoneNumbers

`func (o *Lead) GetPhoneNumbers() []PhoneNumber`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *Lead) GetPhoneNumbersOk() (*[]PhoneNumber, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *Lead) SetPhoneNumbers(v []PhoneNumber)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *Lead) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.

### GetRemoteUpdatedAt

`func (o *Lead) GetRemoteUpdatedAt() time.Time`

GetRemoteUpdatedAt returns the RemoteUpdatedAt field if non-nil, zero value otherwise.

### GetRemoteUpdatedAtOk

`func (o *Lead) GetRemoteUpdatedAtOk() (*time.Time, bool)`

GetRemoteUpdatedAtOk returns a tuple with the RemoteUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUpdatedAt

`func (o *Lead) SetRemoteUpdatedAt(v time.Time)`

SetRemoteUpdatedAt sets RemoteUpdatedAt field to given value.

### HasRemoteUpdatedAt

`func (o *Lead) HasRemoteUpdatedAt() bool`

HasRemoteUpdatedAt returns a boolean if a field has been set.

### SetRemoteUpdatedAtNil

`func (o *Lead) SetRemoteUpdatedAtNil(b bool)`

 SetRemoteUpdatedAtNil sets the value for RemoteUpdatedAt to be an explicit nil

### UnsetRemoteUpdatedAt
`func (o *Lead) UnsetRemoteUpdatedAt()`

UnsetRemoteUpdatedAt ensures that no value is present for RemoteUpdatedAt, not even an explicit nil
### GetRemoteCreatedAt

`func (o *Lead) GetRemoteCreatedAt() time.Time`

GetRemoteCreatedAt returns the RemoteCreatedAt field if non-nil, zero value otherwise.

### GetRemoteCreatedAtOk

`func (o *Lead) GetRemoteCreatedAtOk() (*time.Time, bool)`

GetRemoteCreatedAtOk returns a tuple with the RemoteCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCreatedAt

`func (o *Lead) SetRemoteCreatedAt(v time.Time)`

SetRemoteCreatedAt sets RemoteCreatedAt field to given value.

### HasRemoteCreatedAt

`func (o *Lead) HasRemoteCreatedAt() bool`

HasRemoteCreatedAt returns a boolean if a field has been set.

### SetRemoteCreatedAtNil

`func (o *Lead) SetRemoteCreatedAtNil(b bool)`

 SetRemoteCreatedAtNil sets the value for RemoteCreatedAt to be an explicit nil

### UnsetRemoteCreatedAt
`func (o *Lead) UnsetRemoteCreatedAt()`

UnsetRemoteCreatedAt ensures that no value is present for RemoteCreatedAt, not even an explicit nil
### GetConvertedDate

`func (o *Lead) GetConvertedDate() time.Time`

GetConvertedDate returns the ConvertedDate field if non-nil, zero value otherwise.

### GetConvertedDateOk

`func (o *Lead) GetConvertedDateOk() (*time.Time, bool)`

GetConvertedDateOk returns a tuple with the ConvertedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertedDate

`func (o *Lead) SetConvertedDate(v time.Time)`

SetConvertedDate sets ConvertedDate field to given value.

### HasConvertedDate

`func (o *Lead) HasConvertedDate() bool`

HasConvertedDate returns a boolean if a field has been set.

### SetConvertedDateNil

`func (o *Lead) SetConvertedDateNil(b bool)`

 SetConvertedDateNil sets the value for ConvertedDate to be an explicit nil

### UnsetConvertedDate
`func (o *Lead) UnsetConvertedDate()`

UnsetConvertedDate ensures that no value is present for ConvertedDate, not even an explicit nil
### GetConvertedContact

`func (o *Lead) GetConvertedContact() string`

GetConvertedContact returns the ConvertedContact field if non-nil, zero value otherwise.

### GetConvertedContactOk

`func (o *Lead) GetConvertedContactOk() (*string, bool)`

GetConvertedContactOk returns a tuple with the ConvertedContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertedContact

`func (o *Lead) SetConvertedContact(v string)`

SetConvertedContact sets ConvertedContact field to given value.

### HasConvertedContact

`func (o *Lead) HasConvertedContact() bool`

HasConvertedContact returns a boolean if a field has been set.

### SetConvertedContactNil

`func (o *Lead) SetConvertedContactNil(b bool)`

 SetConvertedContactNil sets the value for ConvertedContact to be an explicit nil

### UnsetConvertedContact
`func (o *Lead) UnsetConvertedContact()`

UnsetConvertedContact ensures that no value is present for ConvertedContact, not even an explicit nil
### GetConvertedAccount

`func (o *Lead) GetConvertedAccount() string`

GetConvertedAccount returns the ConvertedAccount field if non-nil, zero value otherwise.

### GetConvertedAccountOk

`func (o *Lead) GetConvertedAccountOk() (*string, bool)`

GetConvertedAccountOk returns a tuple with the ConvertedAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertedAccount

`func (o *Lead) SetConvertedAccount(v string)`

SetConvertedAccount sets ConvertedAccount field to given value.

### HasConvertedAccount

`func (o *Lead) HasConvertedAccount() bool`

HasConvertedAccount returns a boolean if a field has been set.

### SetConvertedAccountNil

`func (o *Lead) SetConvertedAccountNil(b bool)`

 SetConvertedAccountNil sets the value for ConvertedAccount to be an explicit nil

### UnsetConvertedAccount
`func (o *Lead) UnsetConvertedAccount()`

UnsetConvertedAccount ensures that no value is present for ConvertedAccount, not even an explicit nil
### GetRemoteWasDeleted

`func (o *Lead) GetRemoteWasDeleted() bool`

GetRemoteWasDeleted returns the RemoteWasDeleted field if non-nil, zero value otherwise.

### GetRemoteWasDeletedOk

`func (o *Lead) GetRemoteWasDeletedOk() (*bool, bool)`

GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteWasDeleted

`func (o *Lead) SetRemoteWasDeleted(v bool)`

SetRemoteWasDeleted sets RemoteWasDeleted field to given value.

### HasRemoteWasDeleted

`func (o *Lead) HasRemoteWasDeleted() bool`

HasRemoteWasDeleted returns a boolean if a field has been set.

### GetId

`func (o *Lead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Lead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Lead) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Lead) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteId

`func (o *Lead) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *Lead) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *Lead) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *Lead) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *Lead) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *Lead) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetFieldMappings

`func (o *Lead) GetFieldMappings() map[string]interface{}`

GetFieldMappings returns the FieldMappings field if non-nil, zero value otherwise.

### GetFieldMappingsOk

`func (o *Lead) GetFieldMappingsOk() (*map[string]interface{}, bool)`

GetFieldMappingsOk returns a tuple with the FieldMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMappings

`func (o *Lead) SetFieldMappings(v map[string]interface{})`

SetFieldMappings sets FieldMappings field to given value.

### HasFieldMappings

`func (o *Lead) HasFieldMappings() bool`

HasFieldMappings returns a boolean if a field has been set.

### SetFieldMappingsNil

`func (o *Lead) SetFieldMappingsNil(b bool)`

 SetFieldMappingsNil sets the value for FieldMappings to be an explicit nil

### UnsetFieldMappings
`func (o *Lead) UnsetFieldMappings()`

UnsetFieldMappings ensures that no value is present for FieldMappings, not even an explicit nil
### GetModifiedAt

`func (o *Lead) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Lead) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Lead) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Lead) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetRemoteData

`func (o *Lead) GetRemoteData() []RemoteData`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *Lead) GetRemoteDataOk() (*[]RemoteData, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *Lead) SetRemoteData(v []RemoteData)`

SetRemoteData sets RemoteData field to given value.

### HasRemoteData

`func (o *Lead) HasRemoteData() bool`

HasRemoteData returns a boolean if a field has been set.

### SetRemoteDataNil

`func (o *Lead) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *Lead) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil
### GetRemoteFields

`func (o *Lead) GetRemoteFields() []RemoteField`

GetRemoteFields returns the RemoteFields field if non-nil, zero value otherwise.

### GetRemoteFieldsOk

`func (o *Lead) GetRemoteFieldsOk() (*[]RemoteField, bool)`

GetRemoteFieldsOk returns a tuple with the RemoteFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteFields

`func (o *Lead) SetRemoteFields(v []RemoteField)`

SetRemoteFields sets RemoteFields field to given value.

### HasRemoteFields

`func (o *Lead) HasRemoteFields() bool`

HasRemoteFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



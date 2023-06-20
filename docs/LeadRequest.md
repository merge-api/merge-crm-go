# LeadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | Pointer to **NullableString** | The lead&#39;s owner. | [optional] 
**LeadSource** | Pointer to **NullableString** | The lead&#39;s source. | [optional] 
**Title** | Pointer to **NullableString** | The lead&#39;s title. | [optional] 
**Company** | Pointer to **NullableString** | The lead&#39;s company. | [optional] 
**FirstName** | Pointer to **NullableString** | The lead&#39;s first name. | [optional] 
**LastName** | Pointer to **NullableString** | The lead&#39;s last name. | [optional] 
**Addresses** | Pointer to [**[]AddressRequest**](AddressRequest.md) |  | [optional] 
**EmailAddresses** | Pointer to [**[]EmailAddressRequest**](EmailAddressRequest.md) |  | [optional] 
**PhoneNumbers** | Pointer to [**[]PhoneNumberRequest**](PhoneNumberRequest.md) |  | [optional] 
**ConvertedDate** | Pointer to **NullableTime** | When the lead was converted. | [optional] 
**ConvertedContact** | Pointer to **NullableString** | The contact of the converted lead. | [optional] 
**ConvertedAccount** | Pointer to **NullableString** | The account of the converted lead. | [optional] 
**IntegrationParams** | Pointer to **map[string]interface{}** |  | [optional] 
**LinkedAccountParams** | Pointer to **map[string]interface{}** |  | [optional] 
**RemoteFields** | Pointer to [**[]RemoteFieldRequest**](RemoteFieldRequest.md) |  | [optional] 

## Methods

### NewLeadRequest

`func NewLeadRequest() *LeadRequest`

NewLeadRequest instantiates a new LeadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeadRequestWithDefaults

`func NewLeadRequestWithDefaults() *LeadRequest`

NewLeadRequestWithDefaults instantiates a new LeadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *LeadRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *LeadRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *LeadRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *LeadRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *LeadRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *LeadRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetLeadSource

`func (o *LeadRequest) GetLeadSource() string`

GetLeadSource returns the LeadSource field if non-nil, zero value otherwise.

### GetLeadSourceOk

`func (o *LeadRequest) GetLeadSourceOk() (*string, bool)`

GetLeadSourceOk returns a tuple with the LeadSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadSource

`func (o *LeadRequest) SetLeadSource(v string)`

SetLeadSource sets LeadSource field to given value.

### HasLeadSource

`func (o *LeadRequest) HasLeadSource() bool`

HasLeadSource returns a boolean if a field has been set.

### SetLeadSourceNil

`func (o *LeadRequest) SetLeadSourceNil(b bool)`

 SetLeadSourceNil sets the value for LeadSource to be an explicit nil

### UnsetLeadSource
`func (o *LeadRequest) UnsetLeadSource()`

UnsetLeadSource ensures that no value is present for LeadSource, not even an explicit nil
### GetTitle

`func (o *LeadRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LeadRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LeadRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *LeadRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *LeadRequest) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *LeadRequest) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetCompany

`func (o *LeadRequest) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *LeadRequest) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *LeadRequest) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *LeadRequest) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *LeadRequest) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *LeadRequest) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetFirstName

`func (o *LeadRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *LeadRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *LeadRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *LeadRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *LeadRequest) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *LeadRequest) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *LeadRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *LeadRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *LeadRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *LeadRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *LeadRequest) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *LeadRequest) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetAddresses

`func (o *LeadRequest) GetAddresses() []AddressRequest`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *LeadRequest) GetAddressesOk() (*[]AddressRequest, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *LeadRequest) SetAddresses(v []AddressRequest)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *LeadRequest) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetEmailAddresses

`func (o *LeadRequest) GetEmailAddresses() []EmailAddressRequest`

GetEmailAddresses returns the EmailAddresses field if non-nil, zero value otherwise.

### GetEmailAddressesOk

`func (o *LeadRequest) GetEmailAddressesOk() (*[]EmailAddressRequest, bool)`

GetEmailAddressesOk returns a tuple with the EmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddresses

`func (o *LeadRequest) SetEmailAddresses(v []EmailAddressRequest)`

SetEmailAddresses sets EmailAddresses field to given value.

### HasEmailAddresses

`func (o *LeadRequest) HasEmailAddresses() bool`

HasEmailAddresses returns a boolean if a field has been set.

### GetPhoneNumbers

`func (o *LeadRequest) GetPhoneNumbers() []PhoneNumberRequest`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *LeadRequest) GetPhoneNumbersOk() (*[]PhoneNumberRequest, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *LeadRequest) SetPhoneNumbers(v []PhoneNumberRequest)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *LeadRequest) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.

### GetConvertedDate

`func (o *LeadRequest) GetConvertedDate() time.Time`

GetConvertedDate returns the ConvertedDate field if non-nil, zero value otherwise.

### GetConvertedDateOk

`func (o *LeadRequest) GetConvertedDateOk() (*time.Time, bool)`

GetConvertedDateOk returns a tuple with the ConvertedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertedDate

`func (o *LeadRequest) SetConvertedDate(v time.Time)`

SetConvertedDate sets ConvertedDate field to given value.

### HasConvertedDate

`func (o *LeadRequest) HasConvertedDate() bool`

HasConvertedDate returns a boolean if a field has been set.

### SetConvertedDateNil

`func (o *LeadRequest) SetConvertedDateNil(b bool)`

 SetConvertedDateNil sets the value for ConvertedDate to be an explicit nil

### UnsetConvertedDate
`func (o *LeadRequest) UnsetConvertedDate()`

UnsetConvertedDate ensures that no value is present for ConvertedDate, not even an explicit nil
### GetConvertedContact

`func (o *LeadRequest) GetConvertedContact() string`

GetConvertedContact returns the ConvertedContact field if non-nil, zero value otherwise.

### GetConvertedContactOk

`func (o *LeadRequest) GetConvertedContactOk() (*string, bool)`

GetConvertedContactOk returns a tuple with the ConvertedContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertedContact

`func (o *LeadRequest) SetConvertedContact(v string)`

SetConvertedContact sets ConvertedContact field to given value.

### HasConvertedContact

`func (o *LeadRequest) HasConvertedContact() bool`

HasConvertedContact returns a boolean if a field has been set.

### SetConvertedContactNil

`func (o *LeadRequest) SetConvertedContactNil(b bool)`

 SetConvertedContactNil sets the value for ConvertedContact to be an explicit nil

### UnsetConvertedContact
`func (o *LeadRequest) UnsetConvertedContact()`

UnsetConvertedContact ensures that no value is present for ConvertedContact, not even an explicit nil
### GetConvertedAccount

`func (o *LeadRequest) GetConvertedAccount() string`

GetConvertedAccount returns the ConvertedAccount field if non-nil, zero value otherwise.

### GetConvertedAccountOk

`func (o *LeadRequest) GetConvertedAccountOk() (*string, bool)`

GetConvertedAccountOk returns a tuple with the ConvertedAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertedAccount

`func (o *LeadRequest) SetConvertedAccount(v string)`

SetConvertedAccount sets ConvertedAccount field to given value.

### HasConvertedAccount

`func (o *LeadRequest) HasConvertedAccount() bool`

HasConvertedAccount returns a boolean if a field has been set.

### SetConvertedAccountNil

`func (o *LeadRequest) SetConvertedAccountNil(b bool)`

 SetConvertedAccountNil sets the value for ConvertedAccount to be an explicit nil

### UnsetConvertedAccount
`func (o *LeadRequest) UnsetConvertedAccount()`

UnsetConvertedAccount ensures that no value is present for ConvertedAccount, not even an explicit nil
### GetIntegrationParams

`func (o *LeadRequest) GetIntegrationParams() map[string]interface{}`

GetIntegrationParams returns the IntegrationParams field if non-nil, zero value otherwise.

### GetIntegrationParamsOk

`func (o *LeadRequest) GetIntegrationParamsOk() (*map[string]interface{}, bool)`

GetIntegrationParamsOk returns a tuple with the IntegrationParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationParams

`func (o *LeadRequest) SetIntegrationParams(v map[string]interface{})`

SetIntegrationParams sets IntegrationParams field to given value.

### HasIntegrationParams

`func (o *LeadRequest) HasIntegrationParams() bool`

HasIntegrationParams returns a boolean if a field has been set.

### SetIntegrationParamsNil

`func (o *LeadRequest) SetIntegrationParamsNil(b bool)`

 SetIntegrationParamsNil sets the value for IntegrationParams to be an explicit nil

### UnsetIntegrationParams
`func (o *LeadRequest) UnsetIntegrationParams()`

UnsetIntegrationParams ensures that no value is present for IntegrationParams, not even an explicit nil
### GetLinkedAccountParams

`func (o *LeadRequest) GetLinkedAccountParams() map[string]interface{}`

GetLinkedAccountParams returns the LinkedAccountParams field if non-nil, zero value otherwise.

### GetLinkedAccountParamsOk

`func (o *LeadRequest) GetLinkedAccountParamsOk() (*map[string]interface{}, bool)`

GetLinkedAccountParamsOk returns a tuple with the LinkedAccountParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountParams

`func (o *LeadRequest) SetLinkedAccountParams(v map[string]interface{})`

SetLinkedAccountParams sets LinkedAccountParams field to given value.

### HasLinkedAccountParams

`func (o *LeadRequest) HasLinkedAccountParams() bool`

HasLinkedAccountParams returns a boolean if a field has been set.

### SetLinkedAccountParamsNil

`func (o *LeadRequest) SetLinkedAccountParamsNil(b bool)`

 SetLinkedAccountParamsNil sets the value for LinkedAccountParams to be an explicit nil

### UnsetLinkedAccountParams
`func (o *LeadRequest) UnsetLinkedAccountParams()`

UnsetLinkedAccountParams ensures that no value is present for LinkedAccountParams, not even an explicit nil
### GetRemoteFields

`func (o *LeadRequest) GetRemoteFields() []RemoteFieldRequest`

GetRemoteFields returns the RemoteFields field if non-nil, zero value otherwise.

### GetRemoteFieldsOk

`func (o *LeadRequest) GetRemoteFieldsOk() (*[]RemoteFieldRequest, bool)`

GetRemoteFieldsOk returns a tuple with the RemoteFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteFields

`func (o *LeadRequest) SetRemoteFields(v []RemoteFieldRequest)`

SetRemoteFields sets RemoteFields field to given value.

### HasRemoteFields

`func (o *LeadRequest) HasRemoteFields() bool`

HasRemoteFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



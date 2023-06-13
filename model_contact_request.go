/*
 * Merge CRM API
 *
 * The unified API for building rich integrations with multiple CRM platforms.
 *
 * API version: 1.0
 * Contact: hello@merge.dev
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merge_crm_client

import (
	"encoding/json"
	"time"
)

// ContactRequest # The Contact Object ### Description The `Contact` object is used to represent an existing point of contact at a company in a CRM system. ### Usage Example TODO
type ContactRequest struct {
	// The contact's first name.
	FirstName NullableString `json:"first_name,omitempty"`
	// The contact's last name.
	LastName NullableString `json:"last_name,omitempty"`
	// The contact's account.
	Account NullableString `json:"account,omitempty"`
	Addresses *[]AddressRequest `json:"addresses,omitempty"`
	EmailAddresses *[]EmailAddressRequest `json:"email_addresses,omitempty"`
	PhoneNumbers *[]PhoneNumberRequest `json:"phone_numbers,omitempty"`
	// When the contact's last activity occurred.
	LastActivityAt NullableTime `json:"last_activity_at,omitempty"`
	IntegrationParams map[string]interface{} `json:"integration_params,omitempty"`
	LinkedAccountParams map[string]interface{} `json:"linked_account_params,omitempty"`
	RemoteFields *[]RemoteFieldRequest `json:"remote_fields,omitempty"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewContactRequest instantiates a new ContactRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactRequest() *ContactRequest {
	this := ContactRequest{}
	return &this
}

// NewContactRequestWithDefaults instantiates a new ContactRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactRequestWithDefaults() *ContactRequest {
	this := ContactRequest{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactRequest) GetFirstName() string {
	if o == nil || o.FirstName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactRequest) GetFirstNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *ContactRequest) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *ContactRequest) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *ContactRequest) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *ContactRequest) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactRequest) GetLastName() string {
	if o == nil || o.LastName.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactRequest) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *ContactRequest) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *ContactRequest) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *ContactRequest) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *ContactRequest) UnsetLastName() {
	o.LastName.Unset()
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactRequest) GetAccount() string {
	if o == nil || o.Account.Get() == nil {
		var ret string
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactRequest) GetAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *ContactRequest) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableString and assigns it to the Account field.
func (o *ContactRequest) SetAccount(v string) {
	o.Account.Set(&v)
}
// SetAccountNil sets the value for Account to be an explicit nil
func (o *ContactRequest) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *ContactRequest) UnsetAccount() {
	o.Account.Unset()
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *ContactRequest) GetAddresses() []AddressRequest {
	if o == nil || o.Addresses == nil {
		var ret []AddressRequest
		return ret
	}
	return *o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactRequest) GetAddressesOk() (*[]AddressRequest, bool) {
	if o == nil || o.Addresses == nil {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *ContactRequest) HasAddresses() bool {
	if o != nil && o.Addresses != nil {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []AddressRequest and assigns it to the Addresses field.
func (o *ContactRequest) SetAddresses(v []AddressRequest) {
	o.Addresses = &v
}

// GetEmailAddresses returns the EmailAddresses field value if set, zero value otherwise.
func (o *ContactRequest) GetEmailAddresses() []EmailAddressRequest {
	if o == nil || o.EmailAddresses == nil {
		var ret []EmailAddressRequest
		return ret
	}
	return *o.EmailAddresses
}

// GetEmailAddressesOk returns a tuple with the EmailAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactRequest) GetEmailAddressesOk() (*[]EmailAddressRequest, bool) {
	if o == nil || o.EmailAddresses == nil {
		return nil, false
	}
	return o.EmailAddresses, true
}

// HasEmailAddresses returns a boolean if a field has been set.
func (o *ContactRequest) HasEmailAddresses() bool {
	if o != nil && o.EmailAddresses != nil {
		return true
	}

	return false
}

// SetEmailAddresses gets a reference to the given []EmailAddressRequest and assigns it to the EmailAddresses field.
func (o *ContactRequest) SetEmailAddresses(v []EmailAddressRequest) {
	o.EmailAddresses = &v
}

// GetPhoneNumbers returns the PhoneNumbers field value if set, zero value otherwise.
func (o *ContactRequest) GetPhoneNumbers() []PhoneNumberRequest {
	if o == nil || o.PhoneNumbers == nil {
		var ret []PhoneNumberRequest
		return ret
	}
	return *o.PhoneNumbers
}

// GetPhoneNumbersOk returns a tuple with the PhoneNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactRequest) GetPhoneNumbersOk() (*[]PhoneNumberRequest, bool) {
	if o == nil || o.PhoneNumbers == nil {
		return nil, false
	}
	return o.PhoneNumbers, true
}

// HasPhoneNumbers returns a boolean if a field has been set.
func (o *ContactRequest) HasPhoneNumbers() bool {
	if o != nil && o.PhoneNumbers != nil {
		return true
	}

	return false
}

// SetPhoneNumbers gets a reference to the given []PhoneNumberRequest and assigns it to the PhoneNumbers field.
func (o *ContactRequest) SetPhoneNumbers(v []PhoneNumberRequest) {
	o.PhoneNumbers = &v
}

// GetLastActivityAt returns the LastActivityAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactRequest) GetLastActivityAt() time.Time {
	if o == nil || o.LastActivityAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastActivityAt.Get()
}

// GetLastActivityAtOk returns a tuple with the LastActivityAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactRequest) GetLastActivityAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastActivityAt.Get(), o.LastActivityAt.IsSet()
}

// HasLastActivityAt returns a boolean if a field has been set.
func (o *ContactRequest) HasLastActivityAt() bool {
	if o != nil && o.LastActivityAt.IsSet() {
		return true
	}

	return false
}

// SetLastActivityAt gets a reference to the given NullableTime and assigns it to the LastActivityAt field.
func (o *ContactRequest) SetLastActivityAt(v time.Time) {
	o.LastActivityAt.Set(&v)
}
// SetLastActivityAtNil sets the value for LastActivityAt to be an explicit nil
func (o *ContactRequest) SetLastActivityAtNil() {
	o.LastActivityAt.Set(nil)
}

// UnsetLastActivityAt ensures that no value is present for LastActivityAt, not even an explicit nil
func (o *ContactRequest) UnsetLastActivityAt() {
	o.LastActivityAt.Unset()
}

// GetIntegrationParams returns the IntegrationParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactRequest) GetIntegrationParams() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.IntegrationParams
}

// GetIntegrationParamsOk returns a tuple with the IntegrationParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactRequest) GetIntegrationParamsOk() (*map[string]interface{}, bool) {
	if o == nil || o.IntegrationParams == nil {
		return nil, false
	}
	return &o.IntegrationParams, true
}

// HasIntegrationParams returns a boolean if a field has been set.
func (o *ContactRequest) HasIntegrationParams() bool {
	if o != nil && o.IntegrationParams != nil {
		return true
	}

	return false
}

// SetIntegrationParams gets a reference to the given map[string]interface{} and assigns it to the IntegrationParams field.
func (o *ContactRequest) SetIntegrationParams(v map[string]interface{}) {
	o.IntegrationParams = v
}

// GetLinkedAccountParams returns the LinkedAccountParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactRequest) GetLinkedAccountParams() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.LinkedAccountParams
}

// GetLinkedAccountParamsOk returns a tuple with the LinkedAccountParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactRequest) GetLinkedAccountParamsOk() (*map[string]interface{}, bool) {
	if o == nil || o.LinkedAccountParams == nil {
		return nil, false
	}
	return &o.LinkedAccountParams, true
}

// HasLinkedAccountParams returns a boolean if a field has been set.
func (o *ContactRequest) HasLinkedAccountParams() bool {
	if o != nil && o.LinkedAccountParams != nil {
		return true
	}

	return false
}

// SetLinkedAccountParams gets a reference to the given map[string]interface{} and assigns it to the LinkedAccountParams field.
func (o *ContactRequest) SetLinkedAccountParams(v map[string]interface{}) {
	o.LinkedAccountParams = v
}

// GetRemoteFields returns the RemoteFields field value if set, zero value otherwise.
func (o *ContactRequest) GetRemoteFields() []RemoteFieldRequest {
	if o == nil || o.RemoteFields == nil {
		var ret []RemoteFieldRequest
		return ret
	}
	return *o.RemoteFields
}

// GetRemoteFieldsOk returns a tuple with the RemoteFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactRequest) GetRemoteFieldsOk() (*[]RemoteFieldRequest, bool) {
	if o == nil || o.RemoteFields == nil {
		return nil, false
	}
	return o.RemoteFields, true
}

// HasRemoteFields returns a boolean if a field has been set.
func (o *ContactRequest) HasRemoteFields() bool {
	if o != nil && o.RemoteFields != nil {
		return true
	}

	return false
}

// SetRemoteFields gets a reference to the given []RemoteFieldRequest and assigns it to the RemoteFields field.
func (o *ContactRequest) SetRemoteFields(v []RemoteFieldRequest) {
	o.RemoteFields = &v
}

func (o ContactRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FirstName.IsSet() {
		toSerialize["first_name"] = o.FirstName.Get()
	}
	if o.LastName.IsSet() {
		toSerialize["last_name"] = o.LastName.Get()
	}
	if o.Account.IsSet() {
		toSerialize["account"] = o.Account.Get()
	}
	if o.Addresses != nil {
		toSerialize["addresses"] = o.Addresses
	}
	if o.EmailAddresses != nil {
		toSerialize["email_addresses"] = o.EmailAddresses
	}
	if o.PhoneNumbers != nil {
		toSerialize["phone_numbers"] = o.PhoneNumbers
	}
	if o.LastActivityAt.IsSet() {
		toSerialize["last_activity_at"] = o.LastActivityAt.Get()
	}
	if o.IntegrationParams != nil {
		toSerialize["integration_params"] = o.IntegrationParams
	}
	if o.LinkedAccountParams != nil {
		toSerialize["linked_account_params"] = o.LinkedAccountParams
	}
	if o.RemoteFields != nil {
		toSerialize["remote_fields"] = o.RemoteFields
	}
	return json.Marshal(toSerialize)
}

func (v *ContactRequest) UnmarshalJSON(src []byte) error {
    type ContactRequestUnmarshalTarget ContactRequest

	var intermediateResult ContactRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = ContactRequest(intermediateResult)
	return nil
}
type NullableContactRequest struct {
	value *ContactRequest
	isSet bool
}

func (v NullableContactRequest) Get() *ContactRequest {
	return v.value
}

func (v *NullableContactRequest) Set(val *ContactRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContactRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContactRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactRequest(val *ContactRequest) *NullableContactRequest {
	return &NullableContactRequest{value: val, isSet: true}
}

func (v NullableContactRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}



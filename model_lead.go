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

// Lead # The Lead Object ### Description The `Lead` object is used to represent an individual who is a potential customer. ### Usage Example TODO
type Lead struct {
	// The lead's owner.
	Owner NullableString `json:"owner,omitempty"`
	// The lead's source.
	LeadSource NullableString `json:"lead_source,omitempty"`
	// The lead's title.
	Title NullableString `json:"title,omitempty"`
	// The lead's company.
	Company NullableString `json:"company,omitempty"`
	// The lead's first name.
	FirstName NullableString `json:"first_name,omitempty"`
	// The lead's last name.
	LastName NullableString `json:"last_name,omitempty"`
	Addresses *[]Address `json:"addresses,omitempty"`
	EmailAddresses *[]EmailAddress `json:"email_addresses,omitempty"`
	PhoneNumbers *[]PhoneNumber `json:"phone_numbers,omitempty"`
	// When the third party's lead was updated.
	RemoteUpdatedAt NullableTime `json:"remote_updated_at,omitempty"`
	// When the third party's lead was created.
	RemoteCreatedAt NullableTime `json:"remote_created_at,omitempty"`
	// When the lead was converted.
	ConvertedDate NullableTime `json:"converted_date,omitempty"`
	// The contact of the converted lead.
	ConvertedContact NullableString `json:"converted_contact,omitempty"`
	// The account of the converted lead.
	ConvertedAccount NullableString `json:"converted_account,omitempty"`
	RemoteWasDeleted *bool `json:"remote_was_deleted,omitempty"`
	Id *string `json:"id,omitempty"`
	// The third-party API ID of the matching object.
	RemoteId NullableString `json:"remote_id,omitempty"`
	FieldMappings map[string]interface{} `json:"field_mappings,omitempty"`
	// This is the datetime that this object was last updated by Merge
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	RemoteData []RemoteData `json:"remote_data,omitempty"`
	RemoteFields *[]RemoteField `json:"remote_fields,omitempty"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewLead instantiates a new Lead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLead() *Lead {
	this := Lead{}
	return &this
}

// NewLeadWithDefaults instantiates a new Lead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLeadWithDefaults() *Lead {
	this := Lead{}
	return &this
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Lead) GetOwner() string {
	if o == nil || o.Owner.Get() == nil {
		var ret string
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Lead) GetOwnerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *Lead) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableString and assigns it to the Owner field.
func (o *Lead) SetOwner(v string) {
	o.Owner.Set(&v)
}
// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *Lead) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *Lead) UnsetOwner() {
	o.Owner.Unset()
}

// GetLeadSource returns the LeadSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Lead) GetLeadSource() string {
	if o == nil || o.LeadSource.Get() == nil {
		var ret string
		return ret
	}
	return *o.LeadSource.Get()
}

// GetLeadSourceOk returns a tuple with the LeadSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Lead) GetLeadSourceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LeadSource.Get(), o.LeadSource.IsSet()
}

// HasLeadSource returns a boolean if a field has been set.
func (o *Lead) HasLeadSource() bool {
	if o != nil && o.LeadSource.IsSet() {
		return true
	}

	return false
}

// SetLeadSource gets a reference to the given NullableString and assigns it to the LeadSource field.
func (o *Lead) SetLeadSource(v string) {
	o.LeadSource.Set(&v)
}
// SetLeadSourceNil sets the value for LeadSource to be an explicit nil
func (o *Lead) SetLeadSourceNil() {
	o.LeadSource.Set(nil)
}

// UnsetLeadSource ensures that no value is present for LeadSource, not even an explicit nil
func (o *Lead) UnsetLeadSource() {
	o.LeadSource.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Lead) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Lead) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *Lead) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *Lead) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *Lead) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *Lead) UnsetTitle() {
	o.Title.Unset()
}

// GetCompany returns the Company field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Lead) GetCompany() string {
	if o == nil || o.Company.Get() == nil {
		var ret string
		return ret
	}
	return *o.Company.Get()
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Lead) GetCompanyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Company.Get(), o.Company.IsSet()
}

// HasCompany returns a boolean if a field has been set.
func (o *Lead) HasCompany() bool {
	if o != nil && o.Company.IsSet() {
		return true
	}

	return false
}

// SetCompany gets a reference to the given NullableString and assigns it to the Company field.
func (o *Lead) SetCompany(v string) {
	o.Company.Set(&v)
}
// SetCompanyNil sets the value for Company to be an explicit nil
func (o *Lead) SetCompanyNil() {
	o.Company.Set(nil)
}

// UnsetCompany ensures that no value is present for Company, not even an explicit nil
func (o *Lead) UnsetCompany() {
	o.Company.Unset()
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Lead) GetFirstName() string {
	if o == nil || o.FirstName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Lead) GetFirstNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *Lead) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *Lead) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *Lead) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *Lead) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Lead) GetLastName() string {
	if o == nil || o.LastName.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Lead) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *Lead) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *Lead) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *Lead) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *Lead) UnsetLastName() {
	o.LastName.Unset()
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *Lead) GetAddresses() []Address {
	if o == nil || o.Addresses == nil {
		var ret []Address
		return ret
	}
	return *o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lead) GetAddressesOk() (*[]Address, bool) {
	if o == nil || o.Addresses == nil {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *Lead) HasAddresses() bool {
	if o != nil && o.Addresses != nil {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []Address and assigns it to the Addresses field.
func (o *Lead) SetAddresses(v []Address) {
	o.Addresses = &v
}

// GetEmailAddresses returns the EmailAddresses field value if set, zero value otherwise.
func (o *Lead) GetEmailAddresses() []EmailAddress {
	if o == nil || o.EmailAddresses == nil {
		var ret []EmailAddress
		return ret
	}
	return *o.EmailAddresses
}

// GetEmailAddressesOk returns a tuple with the EmailAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lead) GetEmailAddressesOk() (*[]EmailAddress, bool) {
	if o == nil || o.EmailAddresses == nil {
		return nil, false
	}
	return o.EmailAddresses, true
}

// HasEmailAddresses returns a boolean if a field has been set.
func (o *Lead) HasEmailAddresses() bool {
	if o != nil && o.EmailAddresses != nil {
		return true
	}

	return false
}

// SetEmailAddresses gets a reference to the given []EmailAddress and assigns it to the EmailAddresses field.
func (o *Lead) SetEmailAddresses(v []EmailAddress) {
	o.EmailAddresses = &v
}

// GetPhoneNumbers returns the PhoneNumbers field value if set, zero value otherwise.
func (o *Lead) GetPhoneNumbers() []PhoneNumber {
	if o == nil || o.PhoneNumbers == nil {
		var ret []PhoneNumber
		return ret
	}
	return *o.PhoneNumbers
}

// GetPhoneNumbersOk returns a tuple with the PhoneNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lead) GetPhoneNumbersOk() (*[]PhoneNumber, bool) {
	if o == nil || o.PhoneNumbers == nil {
		return nil, false
	}
	return o.PhoneNumbers, true
}

// HasPhoneNumbers returns a boolean if a field has been set.
func (o *Lead) HasPhoneNumbers() bool {
	if o != nil && o.PhoneNumbers != nil {
		return true
	}

	return false
}

// SetPhoneNumbers gets a reference to the given []PhoneNumber and assigns it to the PhoneNumbers field.
func (o *Lead) SetPhoneNumbers(v []PhoneNumber) {
	o.PhoneNumbers = &v
}

// GetRemoteUpdatedAt returns the RemoteUpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Lead) GetRemoteUpdatedAt() time.Time {
	if o == nil || o.RemoteUpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.RemoteUpdatedAt.Get()
}

// GetRemoteUpdatedAtOk returns a tuple with the RemoteUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Lead) GetRemoteUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteUpdatedAt.Get(), o.RemoteUpdatedAt.IsSet()
}

// HasRemoteUpdatedAt returns a boolean if a field has been set.
func (o *Lead) HasRemoteUpdatedAt() bool {
	if o != nil && o.RemoteUpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetRemoteUpdatedAt gets a reference to the given NullableTime and assigns it to the RemoteUpdatedAt field.
func (o *Lead) SetRemoteUpdatedAt(v time.Time) {
	o.RemoteUpdatedAt.Set(&v)
}
// SetRemoteUpdatedAtNil sets the value for RemoteUpdatedAt to be an explicit nil
func (o *Lead) SetRemoteUpdatedAtNil() {
	o.RemoteUpdatedAt.Set(nil)
}

// UnsetRemoteUpdatedAt ensures that no value is present for RemoteUpdatedAt, not even an explicit nil
func (o *Lead) UnsetRemoteUpdatedAt() {
	o.RemoteUpdatedAt.Unset()
}

// GetRemoteCreatedAt returns the RemoteCreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Lead) GetRemoteCreatedAt() time.Time {
	if o == nil || o.RemoteCreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.RemoteCreatedAt.Get()
}

// GetRemoteCreatedAtOk returns a tuple with the RemoteCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Lead) GetRemoteCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteCreatedAt.Get(), o.RemoteCreatedAt.IsSet()
}

// HasRemoteCreatedAt returns a boolean if a field has been set.
func (o *Lead) HasRemoteCreatedAt() bool {
	if o != nil && o.RemoteCreatedAt.IsSet() {
		return true
	}

	return false
}

// SetRemoteCreatedAt gets a reference to the given NullableTime and assigns it to the RemoteCreatedAt field.
func (o *Lead) SetRemoteCreatedAt(v time.Time) {
	o.RemoteCreatedAt.Set(&v)
}
// SetRemoteCreatedAtNil sets the value for RemoteCreatedAt to be an explicit nil
func (o *Lead) SetRemoteCreatedAtNil() {
	o.RemoteCreatedAt.Set(nil)
}

// UnsetRemoteCreatedAt ensures that no value is present for RemoteCreatedAt, not even an explicit nil
func (o *Lead) UnsetRemoteCreatedAt() {
	o.RemoteCreatedAt.Unset()
}

// GetConvertedDate returns the ConvertedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Lead) GetConvertedDate() time.Time {
	if o == nil || o.ConvertedDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ConvertedDate.Get()
}

// GetConvertedDateOk returns a tuple with the ConvertedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Lead) GetConvertedDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConvertedDate.Get(), o.ConvertedDate.IsSet()
}

// HasConvertedDate returns a boolean if a field has been set.
func (o *Lead) HasConvertedDate() bool {
	if o != nil && o.ConvertedDate.IsSet() {
		return true
	}

	return false
}

// SetConvertedDate gets a reference to the given NullableTime and assigns it to the ConvertedDate field.
func (o *Lead) SetConvertedDate(v time.Time) {
	o.ConvertedDate.Set(&v)
}
// SetConvertedDateNil sets the value for ConvertedDate to be an explicit nil
func (o *Lead) SetConvertedDateNil() {
	o.ConvertedDate.Set(nil)
}

// UnsetConvertedDate ensures that no value is present for ConvertedDate, not even an explicit nil
func (o *Lead) UnsetConvertedDate() {
	o.ConvertedDate.Unset()
}

// GetConvertedContact returns the ConvertedContact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Lead) GetConvertedContact() string {
	if o == nil || o.ConvertedContact.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConvertedContact.Get()
}

// GetConvertedContactOk returns a tuple with the ConvertedContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Lead) GetConvertedContactOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConvertedContact.Get(), o.ConvertedContact.IsSet()
}

// HasConvertedContact returns a boolean if a field has been set.
func (o *Lead) HasConvertedContact() bool {
	if o != nil && o.ConvertedContact.IsSet() {
		return true
	}

	return false
}

// SetConvertedContact gets a reference to the given NullableString and assigns it to the ConvertedContact field.
func (o *Lead) SetConvertedContact(v string) {
	o.ConvertedContact.Set(&v)
}
// SetConvertedContactNil sets the value for ConvertedContact to be an explicit nil
func (o *Lead) SetConvertedContactNil() {
	o.ConvertedContact.Set(nil)
}

// UnsetConvertedContact ensures that no value is present for ConvertedContact, not even an explicit nil
func (o *Lead) UnsetConvertedContact() {
	o.ConvertedContact.Unset()
}

// GetConvertedAccount returns the ConvertedAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Lead) GetConvertedAccount() string {
	if o == nil || o.ConvertedAccount.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConvertedAccount.Get()
}

// GetConvertedAccountOk returns a tuple with the ConvertedAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Lead) GetConvertedAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConvertedAccount.Get(), o.ConvertedAccount.IsSet()
}

// HasConvertedAccount returns a boolean if a field has been set.
func (o *Lead) HasConvertedAccount() bool {
	if o != nil && o.ConvertedAccount.IsSet() {
		return true
	}

	return false
}

// SetConvertedAccount gets a reference to the given NullableString and assigns it to the ConvertedAccount field.
func (o *Lead) SetConvertedAccount(v string) {
	o.ConvertedAccount.Set(&v)
}
// SetConvertedAccountNil sets the value for ConvertedAccount to be an explicit nil
func (o *Lead) SetConvertedAccountNil() {
	o.ConvertedAccount.Set(nil)
}

// UnsetConvertedAccount ensures that no value is present for ConvertedAccount, not even an explicit nil
func (o *Lead) UnsetConvertedAccount() {
	o.ConvertedAccount.Unset()
}

// GetRemoteWasDeleted returns the RemoteWasDeleted field value if set, zero value otherwise.
func (o *Lead) GetRemoteWasDeleted() bool {
	if o == nil || o.RemoteWasDeleted == nil {
		var ret bool
		return ret
	}
	return *o.RemoteWasDeleted
}

// GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lead) GetRemoteWasDeletedOk() (*bool, bool) {
	if o == nil || o.RemoteWasDeleted == nil {
		return nil, false
	}
	return o.RemoteWasDeleted, true
}

// HasRemoteWasDeleted returns a boolean if a field has been set.
func (o *Lead) HasRemoteWasDeleted() bool {
	if o != nil && o.RemoteWasDeleted != nil {
		return true
	}

	return false
}

// SetRemoteWasDeleted gets a reference to the given bool and assigns it to the RemoteWasDeleted field.
func (o *Lead) SetRemoteWasDeleted(v bool) {
	o.RemoteWasDeleted = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Lead) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lead) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Lead) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Lead) SetId(v string) {
	o.Id = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Lead) GetRemoteId() string {
	if o == nil || o.RemoteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemoteId.Get()
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Lead) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteId.Get(), o.RemoteId.IsSet()
}

// HasRemoteId returns a boolean if a field has been set.
func (o *Lead) HasRemoteId() bool {
	if o != nil && o.RemoteId.IsSet() {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given NullableString and assigns it to the RemoteId field.
func (o *Lead) SetRemoteId(v string) {
	o.RemoteId.Set(&v)
}
// SetRemoteIdNil sets the value for RemoteId to be an explicit nil
func (o *Lead) SetRemoteIdNil() {
	o.RemoteId.Set(nil)
}

// UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
func (o *Lead) UnsetRemoteId() {
	o.RemoteId.Unset()
}

// GetFieldMappings returns the FieldMappings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Lead) GetFieldMappings() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.FieldMappings
}

// GetFieldMappingsOk returns a tuple with the FieldMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Lead) GetFieldMappingsOk() (*map[string]interface{}, bool) {
	if o == nil || o.FieldMappings == nil {
		return nil, false
	}
	return &o.FieldMappings, true
}

// HasFieldMappings returns a boolean if a field has been set.
func (o *Lead) HasFieldMappings() bool {
	if o != nil && o.FieldMappings != nil {
		return true
	}

	return false
}

// SetFieldMappings gets a reference to the given map[string]interface{} and assigns it to the FieldMappings field.
func (o *Lead) SetFieldMappings(v map[string]interface{}) {
	o.FieldMappings = v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *Lead) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lead) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *Lead) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *Lead) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetRemoteData returns the RemoteData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Lead) GetRemoteData() []RemoteData {
	if o == nil  {
		var ret []RemoteData
		return ret
	}
	return o.RemoteData
}

// GetRemoteDataOk returns a tuple with the RemoteData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Lead) GetRemoteDataOk() (*[]RemoteData, bool) {
	if o == nil || o.RemoteData == nil {
		return nil, false
	}
	return &o.RemoteData, true
}

// HasRemoteData returns a boolean if a field has been set.
func (o *Lead) HasRemoteData() bool {
	if o != nil && o.RemoteData != nil {
		return true
	}

	return false
}

// SetRemoteData gets a reference to the given []RemoteData and assigns it to the RemoteData field.
func (o *Lead) SetRemoteData(v []RemoteData) {
	o.RemoteData = v
}

// GetRemoteFields returns the RemoteFields field value if set, zero value otherwise.
func (o *Lead) GetRemoteFields() []RemoteField {
	if o == nil || o.RemoteFields == nil {
		var ret []RemoteField
		return ret
	}
	return *o.RemoteFields
}

// GetRemoteFieldsOk returns a tuple with the RemoteFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lead) GetRemoteFieldsOk() (*[]RemoteField, bool) {
	if o == nil || o.RemoteFields == nil {
		return nil, false
	}
	return o.RemoteFields, true
}

// HasRemoteFields returns a boolean if a field has been set.
func (o *Lead) HasRemoteFields() bool {
	if o != nil && o.RemoteFields != nil {
		return true
	}

	return false
}

// SetRemoteFields gets a reference to the given []RemoteField and assigns it to the RemoteFields field.
func (o *Lead) SetRemoteFields(v []RemoteField) {
	o.RemoteFields = &v
}

func (o Lead) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Owner.IsSet() {
		toSerialize["owner"] = o.Owner.Get()
	}
	if o.LeadSource.IsSet() {
		toSerialize["lead_source"] = o.LeadSource.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Company.IsSet() {
		toSerialize["company"] = o.Company.Get()
	}
	if o.FirstName.IsSet() {
		toSerialize["first_name"] = o.FirstName.Get()
	}
	if o.LastName.IsSet() {
		toSerialize["last_name"] = o.LastName.Get()
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
	if o.RemoteUpdatedAt.IsSet() {
		toSerialize["remote_updated_at"] = o.RemoteUpdatedAt.Get()
	}
	if o.RemoteCreatedAt.IsSet() {
		toSerialize["remote_created_at"] = o.RemoteCreatedAt.Get()
	}
	if o.ConvertedDate.IsSet() {
		toSerialize["converted_date"] = o.ConvertedDate.Get()
	}
	if o.ConvertedContact.IsSet() {
		toSerialize["converted_contact"] = o.ConvertedContact.Get()
	}
	if o.ConvertedAccount.IsSet() {
		toSerialize["converted_account"] = o.ConvertedAccount.Get()
	}
	if o.RemoteWasDeleted != nil {
		toSerialize["remote_was_deleted"] = o.RemoteWasDeleted
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.RemoteId.IsSet() {
		toSerialize["remote_id"] = o.RemoteId.Get()
	}
	if o.FieldMappings != nil {
		toSerialize["field_mappings"] = o.FieldMappings
	}
	if o.ModifiedAt != nil {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if o.RemoteData != nil {
		toSerialize["remote_data"] = o.RemoteData
	}
	if o.RemoteFields != nil {
		toSerialize["remote_fields"] = o.RemoteFields
	}
	return json.Marshal(toSerialize)
}

func (v *Lead) UnmarshalJSON(src []byte) error {
    type LeadUnmarshalTarget Lead

	var intermediateResult LeadUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = Lead(intermediateResult)
	return nil
}
type NullableLead struct {
	value *Lead
	isSet bool
}

func (v NullableLead) Get() *Lead {
	return v.value
}

func (v *NullableLead) Set(val *Lead) {
	v.value = val
	v.isSet = true
}

func (v NullableLead) IsSet() bool {
	return v.isSet
}

func (v *NullableLead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLead(val *Lead) *NullableLead {
	return &NullableLead{value: val, isSet: true}
}

func (v NullableLead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}



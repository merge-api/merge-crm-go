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

// LeadRequest # The Lead Object ### Description The `Lead` object is used to represent a lead in the remote system. ### Usage Example TODO
type LeadRequest struct {
	// The third-party API ID of the matching object.
	RemoteId NullableString `json:"remote_id,omitempty"`
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
	// When the third party's lead was updated.
	RemoteUpdatedAt NullableTime `json:"remote_updated_at,omitempty"`
	// When the third party's lead was created.
	RemoteCreatedAt NullableTime `json:"remote_created_at,omitempty"`
	// When the lead was converted.
	ConvertedDate NullableTime `json:"converted_date,omitempty"`
	ConvertedContact NullableString `json:"converted_contact,omitempty"`
	ConvertedAccount NullableString `json:"converted_account,omitempty"`
	IntegrationParams map[string]interface{} `json:"integration_params,omitempty"`
	LinkedAccountParams map[string]interface{} `json:"linked_account_params,omitempty"`
    // raw json response by property name
    ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewLeadRequest instantiates a new LeadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLeadRequest() *LeadRequest {
	this := LeadRequest{}
	return &this
}

// NewLeadRequestWithDefaults instantiates a new LeadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLeadRequestWithDefaults() *LeadRequest {
	this := LeadRequest{}
	return &this
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LeadRequest) GetRemoteId() string {
	if o == nil || o.RemoteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemoteId.Get()
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LeadRequest) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteId.Get(), o.RemoteId.IsSet()
}

// HasRemoteId returns a boolean if a field has been set.
func (o *LeadRequest) HasRemoteId() bool {
	if o != nil && o.RemoteId.IsSet() {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given NullableString and assigns it to the RemoteId field.
func (o *LeadRequest) SetRemoteId(v string) {
	o.RemoteId.Set(&v)
}
// SetRemoteIdNil sets the value for RemoteId to be an explicit nil
func (o *LeadRequest) SetRemoteIdNil() {
	o.RemoteId.Set(nil)
}

// UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
func (o *LeadRequest) UnsetRemoteId() {
	o.RemoteId.Unset()
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LeadRequest) GetOwner() string {
	if o == nil || o.Owner.Get() == nil {
		var ret string
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LeadRequest) GetOwnerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *LeadRequest) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableString and assigns it to the Owner field.
func (o *LeadRequest) SetOwner(v string) {
	o.Owner.Set(&v)
}
// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *LeadRequest) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *LeadRequest) UnsetOwner() {
	o.Owner.Unset()
}

// GetLeadSource returns the LeadSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LeadRequest) GetLeadSource() string {
	if o == nil || o.LeadSource.Get() == nil {
		var ret string
		return ret
	}
	return *o.LeadSource.Get()
}

// GetLeadSourceOk returns a tuple with the LeadSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LeadRequest) GetLeadSourceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LeadSource.Get(), o.LeadSource.IsSet()
}

// HasLeadSource returns a boolean if a field has been set.
func (o *LeadRequest) HasLeadSource() bool {
	if o != nil && o.LeadSource.IsSet() {
		return true
	}

	return false
}

// SetLeadSource gets a reference to the given NullableString and assigns it to the LeadSource field.
func (o *LeadRequest) SetLeadSource(v string) {
	o.LeadSource.Set(&v)
}
// SetLeadSourceNil sets the value for LeadSource to be an explicit nil
func (o *LeadRequest) SetLeadSourceNil() {
	o.LeadSource.Set(nil)
}

// UnsetLeadSource ensures that no value is present for LeadSource, not even an explicit nil
func (o *LeadRequest) UnsetLeadSource() {
	o.LeadSource.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LeadRequest) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LeadRequest) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *LeadRequest) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *LeadRequest) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *LeadRequest) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *LeadRequest) UnsetTitle() {
	o.Title.Unset()
}

// GetCompany returns the Company field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LeadRequest) GetCompany() string {
	if o == nil || o.Company.Get() == nil {
		var ret string
		return ret
	}
	return *o.Company.Get()
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LeadRequest) GetCompanyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Company.Get(), o.Company.IsSet()
}

// HasCompany returns a boolean if a field has been set.
func (o *LeadRequest) HasCompany() bool {
	if o != nil && o.Company.IsSet() {
		return true
	}

	return false
}

// SetCompany gets a reference to the given NullableString and assigns it to the Company field.
func (o *LeadRequest) SetCompany(v string) {
	o.Company.Set(&v)
}
// SetCompanyNil sets the value for Company to be an explicit nil
func (o *LeadRequest) SetCompanyNil() {
	o.Company.Set(nil)
}

// UnsetCompany ensures that no value is present for Company, not even an explicit nil
func (o *LeadRequest) UnsetCompany() {
	o.Company.Unset()
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LeadRequest) GetFirstName() string {
	if o == nil || o.FirstName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LeadRequest) GetFirstNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *LeadRequest) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *LeadRequest) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *LeadRequest) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *LeadRequest) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LeadRequest) GetLastName() string {
	if o == nil || o.LastName.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LeadRequest) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *LeadRequest) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *LeadRequest) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *LeadRequest) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *LeadRequest) UnsetLastName() {
	o.LastName.Unset()
}

// GetRemoteUpdatedAt returns the RemoteUpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LeadRequest) GetRemoteUpdatedAt() time.Time {
	if o == nil || o.RemoteUpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.RemoteUpdatedAt.Get()
}

// GetRemoteUpdatedAtOk returns a tuple with the RemoteUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LeadRequest) GetRemoteUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteUpdatedAt.Get(), o.RemoteUpdatedAt.IsSet()
}

// HasRemoteUpdatedAt returns a boolean if a field has been set.
func (o *LeadRequest) HasRemoteUpdatedAt() bool {
	if o != nil && o.RemoteUpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetRemoteUpdatedAt gets a reference to the given NullableTime and assigns it to the RemoteUpdatedAt field.
func (o *LeadRequest) SetRemoteUpdatedAt(v time.Time) {
	o.RemoteUpdatedAt.Set(&v)
}
// SetRemoteUpdatedAtNil sets the value for RemoteUpdatedAt to be an explicit nil
func (o *LeadRequest) SetRemoteUpdatedAtNil() {
	o.RemoteUpdatedAt.Set(nil)
}

// UnsetRemoteUpdatedAt ensures that no value is present for RemoteUpdatedAt, not even an explicit nil
func (o *LeadRequest) UnsetRemoteUpdatedAt() {
	o.RemoteUpdatedAt.Unset()
}

// GetRemoteCreatedAt returns the RemoteCreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LeadRequest) GetRemoteCreatedAt() time.Time {
	if o == nil || o.RemoteCreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.RemoteCreatedAt.Get()
}

// GetRemoteCreatedAtOk returns a tuple with the RemoteCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LeadRequest) GetRemoteCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteCreatedAt.Get(), o.RemoteCreatedAt.IsSet()
}

// HasRemoteCreatedAt returns a boolean if a field has been set.
func (o *LeadRequest) HasRemoteCreatedAt() bool {
	if o != nil && o.RemoteCreatedAt.IsSet() {
		return true
	}

	return false
}

// SetRemoteCreatedAt gets a reference to the given NullableTime and assigns it to the RemoteCreatedAt field.
func (o *LeadRequest) SetRemoteCreatedAt(v time.Time) {
	o.RemoteCreatedAt.Set(&v)
}
// SetRemoteCreatedAtNil sets the value for RemoteCreatedAt to be an explicit nil
func (o *LeadRequest) SetRemoteCreatedAtNil() {
	o.RemoteCreatedAt.Set(nil)
}

// UnsetRemoteCreatedAt ensures that no value is present for RemoteCreatedAt, not even an explicit nil
func (o *LeadRequest) UnsetRemoteCreatedAt() {
	o.RemoteCreatedAt.Unset()
}

// GetConvertedDate returns the ConvertedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LeadRequest) GetConvertedDate() time.Time {
	if o == nil || o.ConvertedDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ConvertedDate.Get()
}

// GetConvertedDateOk returns a tuple with the ConvertedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LeadRequest) GetConvertedDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConvertedDate.Get(), o.ConvertedDate.IsSet()
}

// HasConvertedDate returns a boolean if a field has been set.
func (o *LeadRequest) HasConvertedDate() bool {
	if o != nil && o.ConvertedDate.IsSet() {
		return true
	}

	return false
}

// SetConvertedDate gets a reference to the given NullableTime and assigns it to the ConvertedDate field.
func (o *LeadRequest) SetConvertedDate(v time.Time) {
	o.ConvertedDate.Set(&v)
}
// SetConvertedDateNil sets the value for ConvertedDate to be an explicit nil
func (o *LeadRequest) SetConvertedDateNil() {
	o.ConvertedDate.Set(nil)
}

// UnsetConvertedDate ensures that no value is present for ConvertedDate, not even an explicit nil
func (o *LeadRequest) UnsetConvertedDate() {
	o.ConvertedDate.Unset()
}

// GetConvertedContact returns the ConvertedContact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LeadRequest) GetConvertedContact() string {
	if o == nil || o.ConvertedContact.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConvertedContact.Get()
}

// GetConvertedContactOk returns a tuple with the ConvertedContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LeadRequest) GetConvertedContactOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConvertedContact.Get(), o.ConvertedContact.IsSet()
}

// HasConvertedContact returns a boolean if a field has been set.
func (o *LeadRequest) HasConvertedContact() bool {
	if o != nil && o.ConvertedContact.IsSet() {
		return true
	}

	return false
}

// SetConvertedContact gets a reference to the given NullableString and assigns it to the ConvertedContact field.
func (o *LeadRequest) SetConvertedContact(v string) {
	o.ConvertedContact.Set(&v)
}
// SetConvertedContactNil sets the value for ConvertedContact to be an explicit nil
func (o *LeadRequest) SetConvertedContactNil() {
	o.ConvertedContact.Set(nil)
}

// UnsetConvertedContact ensures that no value is present for ConvertedContact, not even an explicit nil
func (o *LeadRequest) UnsetConvertedContact() {
	o.ConvertedContact.Unset()
}

// GetConvertedAccount returns the ConvertedAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LeadRequest) GetConvertedAccount() string {
	if o == nil || o.ConvertedAccount.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConvertedAccount.Get()
}

// GetConvertedAccountOk returns a tuple with the ConvertedAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LeadRequest) GetConvertedAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConvertedAccount.Get(), o.ConvertedAccount.IsSet()
}

// HasConvertedAccount returns a boolean if a field has been set.
func (o *LeadRequest) HasConvertedAccount() bool {
	if o != nil && o.ConvertedAccount.IsSet() {
		return true
	}

	return false
}

// SetConvertedAccount gets a reference to the given NullableString and assigns it to the ConvertedAccount field.
func (o *LeadRequest) SetConvertedAccount(v string) {
	o.ConvertedAccount.Set(&v)
}
// SetConvertedAccountNil sets the value for ConvertedAccount to be an explicit nil
func (o *LeadRequest) SetConvertedAccountNil() {
	o.ConvertedAccount.Set(nil)
}

// UnsetConvertedAccount ensures that no value is present for ConvertedAccount, not even an explicit nil
func (o *LeadRequest) UnsetConvertedAccount() {
	o.ConvertedAccount.Unset()
}

// GetIntegrationParams returns the IntegrationParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LeadRequest) GetIntegrationParams() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.IntegrationParams
}

// GetIntegrationParamsOk returns a tuple with the IntegrationParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LeadRequest) GetIntegrationParamsOk() (*map[string]interface{}, bool) {
	if o == nil || o.IntegrationParams == nil {
		return nil, false
	}
	return &o.IntegrationParams, true
}

// HasIntegrationParams returns a boolean if a field has been set.
func (o *LeadRequest) HasIntegrationParams() bool {
	if o != nil && o.IntegrationParams != nil {
		return true
	}

	return false
}

// SetIntegrationParams gets a reference to the given map[string]interface{} and assigns it to the IntegrationParams field.
func (o *LeadRequest) SetIntegrationParams(v map[string]interface{}) {
	o.IntegrationParams = v
}

// GetLinkedAccountParams returns the LinkedAccountParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LeadRequest) GetLinkedAccountParams() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.LinkedAccountParams
}

// GetLinkedAccountParamsOk returns a tuple with the LinkedAccountParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LeadRequest) GetLinkedAccountParamsOk() (*map[string]interface{}, bool) {
	if o == nil || o.LinkedAccountParams == nil {
		return nil, false
	}
	return &o.LinkedAccountParams, true
}

// HasLinkedAccountParams returns a boolean if a field has been set.
func (o *LeadRequest) HasLinkedAccountParams() bool {
	if o != nil && o.LinkedAccountParams != nil {
		return true
	}

	return false
}

// SetLinkedAccountParams gets a reference to the given map[string]interface{} and assigns it to the LinkedAccountParams field.
func (o *LeadRequest) SetLinkedAccountParams(v map[string]interface{}) {
	o.LinkedAccountParams = v
}

func (o LeadRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RemoteId.IsSet() {
		toSerialize["remote_id"] = o.RemoteId.Get()
	}
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
	if o.IntegrationParams != nil {
		toSerialize["integration_params"] = o.IntegrationParams
	}
	if o.LinkedAccountParams != nil {
		toSerialize["linked_account_params"] = o.LinkedAccountParams
	}
	return json.Marshal(toSerialize)
}

func (v *LeadRequest) UnmarshalJSON(src []byte) error {
    type LeadRequestUnmarshalTarget LeadRequest

	var intermediateResult LeadRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = LeadRequest(intermediateResult)
	return nil
}
type NullableLeadRequest struct {
	value *LeadRequest
	isSet bool
}

func (v NullableLeadRequest) Get() *LeadRequest {
	return v.value
}

func (v *NullableLeadRequest) Set(val *LeadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLeadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLeadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLeadRequest(val *LeadRequest) *NullableLeadRequest {
	return &NullableLeadRequest{value: val, isSet: true}
}

func (v NullableLeadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLeadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}



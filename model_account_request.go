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

// AccountRequest # The Account Object ### Description The `Account` object is used to represent a company in a CRM system. ### Usage Example TODO
type AccountRequest struct {
	// The account's owner.
	Owner NullableString `json:"owner,omitempty"`
	// The account's name.
	Name NullableString `json:"name,omitempty"`
	// The account's description.
	Description NullableString `json:"description,omitempty"`
	// The account's industry.
	Industry NullableString `json:"industry,omitempty"`
	// The account's website.
	Website NullableString `json:"website,omitempty"`
	// The account's number of employees.
	NumberOfEmployees NullableInt32 `json:"number_of_employees,omitempty"`
	// The last date (either most recent or furthest in the future) of when an activity occurs in an account.
	LastActivityAt NullableTime `json:"last_activity_at,omitempty"`
	IntegrationParams map[string]interface{} `json:"integration_params,omitempty"`
	LinkedAccountParams map[string]interface{} `json:"linked_account_params,omitempty"`
	RemoteFields *[]RemoteFieldRequest `json:"remote_fields,omitempty"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewAccountRequest instantiates a new AccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountRequest() *AccountRequest {
	this := AccountRequest{}
	return &this
}

// NewAccountRequestWithDefaults instantiates a new AccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountRequestWithDefaults() *AccountRequest {
	this := AccountRequest{}
	return &this
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountRequest) GetOwner() string {
	if o == nil || o.Owner.Get() == nil {
		var ret string
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountRequest) GetOwnerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *AccountRequest) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableString and assigns it to the Owner field.
func (o *AccountRequest) SetOwner(v string) {
	o.Owner.Set(&v)
}
// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *AccountRequest) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *AccountRequest) UnsetOwner() {
	o.Owner.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountRequest) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AccountRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AccountRequest) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AccountRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AccountRequest) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountRequest) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountRequest) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AccountRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AccountRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AccountRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AccountRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetIndustry returns the Industry field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountRequest) GetIndustry() string {
	if o == nil || o.Industry.Get() == nil {
		var ret string
		return ret
	}
	return *o.Industry.Get()
}

// GetIndustryOk returns a tuple with the Industry field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountRequest) GetIndustryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Industry.Get(), o.Industry.IsSet()
}

// HasIndustry returns a boolean if a field has been set.
func (o *AccountRequest) HasIndustry() bool {
	if o != nil && o.Industry.IsSet() {
		return true
	}

	return false
}

// SetIndustry gets a reference to the given NullableString and assigns it to the Industry field.
func (o *AccountRequest) SetIndustry(v string) {
	o.Industry.Set(&v)
}
// SetIndustryNil sets the value for Industry to be an explicit nil
func (o *AccountRequest) SetIndustryNil() {
	o.Industry.Set(nil)
}

// UnsetIndustry ensures that no value is present for Industry, not even an explicit nil
func (o *AccountRequest) UnsetIndustry() {
	o.Industry.Unset()
}

// GetWebsite returns the Website field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountRequest) GetWebsite() string {
	if o == nil || o.Website.Get() == nil {
		var ret string
		return ret
	}
	return *o.Website.Get()
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountRequest) GetWebsiteOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Website.Get(), o.Website.IsSet()
}

// HasWebsite returns a boolean if a field has been set.
func (o *AccountRequest) HasWebsite() bool {
	if o != nil && o.Website.IsSet() {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given NullableString and assigns it to the Website field.
func (o *AccountRequest) SetWebsite(v string) {
	o.Website.Set(&v)
}
// SetWebsiteNil sets the value for Website to be an explicit nil
func (o *AccountRequest) SetWebsiteNil() {
	o.Website.Set(nil)
}

// UnsetWebsite ensures that no value is present for Website, not even an explicit nil
func (o *AccountRequest) UnsetWebsite() {
	o.Website.Unset()
}

// GetNumberOfEmployees returns the NumberOfEmployees field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountRequest) GetNumberOfEmployees() int32 {
	if o == nil || o.NumberOfEmployees.Get() == nil {
		var ret int32
		return ret
	}
	return *o.NumberOfEmployees.Get()
}

// GetNumberOfEmployeesOk returns a tuple with the NumberOfEmployees field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountRequest) GetNumberOfEmployeesOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NumberOfEmployees.Get(), o.NumberOfEmployees.IsSet()
}

// HasNumberOfEmployees returns a boolean if a field has been set.
func (o *AccountRequest) HasNumberOfEmployees() bool {
	if o != nil && o.NumberOfEmployees.IsSet() {
		return true
	}

	return false
}

// SetNumberOfEmployees gets a reference to the given NullableInt32 and assigns it to the NumberOfEmployees field.
func (o *AccountRequest) SetNumberOfEmployees(v int32) {
	o.NumberOfEmployees.Set(&v)
}
// SetNumberOfEmployeesNil sets the value for NumberOfEmployees to be an explicit nil
func (o *AccountRequest) SetNumberOfEmployeesNil() {
	o.NumberOfEmployees.Set(nil)
}

// UnsetNumberOfEmployees ensures that no value is present for NumberOfEmployees, not even an explicit nil
func (o *AccountRequest) UnsetNumberOfEmployees() {
	o.NumberOfEmployees.Unset()
}

// GetLastActivityAt returns the LastActivityAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountRequest) GetLastActivityAt() time.Time {
	if o == nil || o.LastActivityAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastActivityAt.Get()
}

// GetLastActivityAtOk returns a tuple with the LastActivityAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountRequest) GetLastActivityAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastActivityAt.Get(), o.LastActivityAt.IsSet()
}

// HasLastActivityAt returns a boolean if a field has been set.
func (o *AccountRequest) HasLastActivityAt() bool {
	if o != nil && o.LastActivityAt.IsSet() {
		return true
	}

	return false
}

// SetLastActivityAt gets a reference to the given NullableTime and assigns it to the LastActivityAt field.
func (o *AccountRequest) SetLastActivityAt(v time.Time) {
	o.LastActivityAt.Set(&v)
}
// SetLastActivityAtNil sets the value for LastActivityAt to be an explicit nil
func (o *AccountRequest) SetLastActivityAtNil() {
	o.LastActivityAt.Set(nil)
}

// UnsetLastActivityAt ensures that no value is present for LastActivityAt, not even an explicit nil
func (o *AccountRequest) UnsetLastActivityAt() {
	o.LastActivityAt.Unset()
}

// GetIntegrationParams returns the IntegrationParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountRequest) GetIntegrationParams() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.IntegrationParams
}

// GetIntegrationParamsOk returns a tuple with the IntegrationParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountRequest) GetIntegrationParamsOk() (*map[string]interface{}, bool) {
	if o == nil || o.IntegrationParams == nil {
		return nil, false
	}
	return &o.IntegrationParams, true
}

// HasIntegrationParams returns a boolean if a field has been set.
func (o *AccountRequest) HasIntegrationParams() bool {
	if o != nil && o.IntegrationParams != nil {
		return true
	}

	return false
}

// SetIntegrationParams gets a reference to the given map[string]interface{} and assigns it to the IntegrationParams field.
func (o *AccountRequest) SetIntegrationParams(v map[string]interface{}) {
	o.IntegrationParams = v
}

// GetLinkedAccountParams returns the LinkedAccountParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountRequest) GetLinkedAccountParams() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.LinkedAccountParams
}

// GetLinkedAccountParamsOk returns a tuple with the LinkedAccountParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountRequest) GetLinkedAccountParamsOk() (*map[string]interface{}, bool) {
	if o == nil || o.LinkedAccountParams == nil {
		return nil, false
	}
	return &o.LinkedAccountParams, true
}

// HasLinkedAccountParams returns a boolean if a field has been set.
func (o *AccountRequest) HasLinkedAccountParams() bool {
	if o != nil && o.LinkedAccountParams != nil {
		return true
	}

	return false
}

// SetLinkedAccountParams gets a reference to the given map[string]interface{} and assigns it to the LinkedAccountParams field.
func (o *AccountRequest) SetLinkedAccountParams(v map[string]interface{}) {
	o.LinkedAccountParams = v
}

// GetRemoteFields returns the RemoteFields field value if set, zero value otherwise.
func (o *AccountRequest) GetRemoteFields() []RemoteFieldRequest {
	if o == nil || o.RemoteFields == nil {
		var ret []RemoteFieldRequest
		return ret
	}
	return *o.RemoteFields
}

// GetRemoteFieldsOk returns a tuple with the RemoteFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountRequest) GetRemoteFieldsOk() (*[]RemoteFieldRequest, bool) {
	if o == nil || o.RemoteFields == nil {
		return nil, false
	}
	return o.RemoteFields, true
}

// HasRemoteFields returns a boolean if a field has been set.
func (o *AccountRequest) HasRemoteFields() bool {
	if o != nil && o.RemoteFields != nil {
		return true
	}

	return false
}

// SetRemoteFields gets a reference to the given []RemoteFieldRequest and assigns it to the RemoteFields field.
func (o *AccountRequest) SetRemoteFields(v []RemoteFieldRequest) {
	o.RemoteFields = &v
}

func (o AccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Owner.IsSet() {
		toSerialize["owner"] = o.Owner.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Industry.IsSet() {
		toSerialize["industry"] = o.Industry.Get()
	}
	if o.Website.IsSet() {
		toSerialize["website"] = o.Website.Get()
	}
	if o.NumberOfEmployees.IsSet() {
		toSerialize["number_of_employees"] = o.NumberOfEmployees.Get()
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

func (v *AccountRequest) UnmarshalJSON(src []byte) error {
    type AccountRequestUnmarshalTarget AccountRequest

	var intermediateResult AccountRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = AccountRequest(intermediateResult)
	return nil
}
type NullableAccountRequest struct {
	value *AccountRequest
	isSet bool
}

func (v NullableAccountRequest) Get() *AccountRequest {
	return v.value
}

func (v *NullableAccountRequest) Set(val *AccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountRequest(val *AccountRequest) *NullableAccountRequest {
	return &NullableAccountRequest{value: val, isSet: true}
}

func (v NullableAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}



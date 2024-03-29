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

// PatchedEngagementRequest # The Engagement Object ### Description The `Engagement` object is used to represent an interaction noted in a CRM system. ### Usage Example TODO
type PatchedEngagementRequest struct {
	// The engagement's owner.
	Owner NullableString `json:"owner,omitempty"`
	// The engagement's content.
	Content NullableString `json:"content,omitempty"`
	// The engagement's subject.
	Subject NullableString `json:"subject,omitempty"`
	// The engagement's direction.  * `INBOUND` - INBOUND * `OUTBOUND` - OUTBOUND
	Direction NullableDirectionEnum `json:"direction,omitempty"`
	// The engagement type of the engagement.
	EngagementType NullableString `json:"engagement_type,omitempty"`
	// The time at which the engagement started.
	StartTime NullableTime `json:"start_time,omitempty"`
	// The time at which the engagement ended.
	EndTime NullableTime `json:"end_time,omitempty"`
	// The account of the engagement.
	Account NullableString `json:"account,omitempty"`
	Contacts *[]string `json:"contacts,omitempty"`
	IntegrationParams map[string]interface{} `json:"integration_params,omitempty"`
	LinkedAccountParams map[string]interface{} `json:"linked_account_params,omitempty"`
	RemoteFields *[]RemoteFieldRequest `json:"remote_fields,omitempty"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewPatchedEngagementRequest instantiates a new PatchedEngagementRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedEngagementRequest() *PatchedEngagementRequest {
	this := PatchedEngagementRequest{}
	return &this
}

// NewPatchedEngagementRequestWithDefaults instantiates a new PatchedEngagementRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedEngagementRequestWithDefaults() *PatchedEngagementRequest {
	this := PatchedEngagementRequest{}
	return &this
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedEngagementRequest) GetOwner() string {
	if o == nil || o.Owner.Get() == nil {
		var ret string
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedEngagementRequest) GetOwnerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *PatchedEngagementRequest) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableString and assigns it to the Owner field.
func (o *PatchedEngagementRequest) SetOwner(v string) {
	o.Owner.Set(&v)
}
// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *PatchedEngagementRequest) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *PatchedEngagementRequest) UnsetOwner() {
	o.Owner.Unset()
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedEngagementRequest) GetContent() string {
	if o == nil || o.Content.Get() == nil {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedEngagementRequest) GetContentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *PatchedEngagementRequest) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *PatchedEngagementRequest) SetContent(v string) {
	o.Content.Set(&v)
}
// SetContentNil sets the value for Content to be an explicit nil
func (o *PatchedEngagementRequest) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *PatchedEngagementRequest) UnsetContent() {
	o.Content.Unset()
}

// GetSubject returns the Subject field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedEngagementRequest) GetSubject() string {
	if o == nil || o.Subject.Get() == nil {
		var ret string
		return ret
	}
	return *o.Subject.Get()
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedEngagementRequest) GetSubjectOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Subject.Get(), o.Subject.IsSet()
}

// HasSubject returns a boolean if a field has been set.
func (o *PatchedEngagementRequest) HasSubject() bool {
	if o != nil && o.Subject.IsSet() {
		return true
	}

	return false
}

// SetSubject gets a reference to the given NullableString and assigns it to the Subject field.
func (o *PatchedEngagementRequest) SetSubject(v string) {
	o.Subject.Set(&v)
}
// SetSubjectNil sets the value for Subject to be an explicit nil
func (o *PatchedEngagementRequest) SetSubjectNil() {
	o.Subject.Set(nil)
}

// UnsetSubject ensures that no value is present for Subject, not even an explicit nil
func (o *PatchedEngagementRequest) UnsetSubject() {
	o.Subject.Unset()
}

// GetDirection returns the Direction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedEngagementRequest) GetDirection() DirectionEnum {
	if o == nil || o.Direction.Get() == nil {
		var ret DirectionEnum
		return ret
	}
	return *o.Direction.Get()
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedEngagementRequest) GetDirectionOk() (*DirectionEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Direction.Get(), o.Direction.IsSet()
}

// HasDirection returns a boolean if a field has been set.
func (o *PatchedEngagementRequest) HasDirection() bool {
	if o != nil && o.Direction.IsSet() {
		return true
	}

	return false
}

// SetDirection gets a reference to the given NullableDirectionEnum and assigns it to the Direction field.
func (o *PatchedEngagementRequest) SetDirection(v DirectionEnum) {
	o.Direction.Set(&v)
}
// SetDirectionNil sets the value for Direction to be an explicit nil
func (o *PatchedEngagementRequest) SetDirectionNil() {
	o.Direction.Set(nil)
}

// UnsetDirection ensures that no value is present for Direction, not even an explicit nil
func (o *PatchedEngagementRequest) UnsetDirection() {
	o.Direction.Unset()
}

// GetEngagementType returns the EngagementType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedEngagementRequest) GetEngagementType() string {
	if o == nil || o.EngagementType.Get() == nil {
		var ret string
		return ret
	}
	return *o.EngagementType.Get()
}

// GetEngagementTypeOk returns a tuple with the EngagementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedEngagementRequest) GetEngagementTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EngagementType.Get(), o.EngagementType.IsSet()
}

// HasEngagementType returns a boolean if a field has been set.
func (o *PatchedEngagementRequest) HasEngagementType() bool {
	if o != nil && o.EngagementType.IsSet() {
		return true
	}

	return false
}

// SetEngagementType gets a reference to the given NullableString and assigns it to the EngagementType field.
func (o *PatchedEngagementRequest) SetEngagementType(v string) {
	o.EngagementType.Set(&v)
}
// SetEngagementTypeNil sets the value for EngagementType to be an explicit nil
func (o *PatchedEngagementRequest) SetEngagementTypeNil() {
	o.EngagementType.Set(nil)
}

// UnsetEngagementType ensures that no value is present for EngagementType, not even an explicit nil
func (o *PatchedEngagementRequest) UnsetEngagementType() {
	o.EngagementType.Unset()
}

// GetStartTime returns the StartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedEngagementRequest) GetStartTime() time.Time {
	if o == nil || o.StartTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime.Get()
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedEngagementRequest) GetStartTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartTime.Get(), o.StartTime.IsSet()
}

// HasStartTime returns a boolean if a field has been set.
func (o *PatchedEngagementRequest) HasStartTime() bool {
	if o != nil && o.StartTime.IsSet() {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given NullableTime and assigns it to the StartTime field.
func (o *PatchedEngagementRequest) SetStartTime(v time.Time) {
	o.StartTime.Set(&v)
}
// SetStartTimeNil sets the value for StartTime to be an explicit nil
func (o *PatchedEngagementRequest) SetStartTimeNil() {
	o.StartTime.Set(nil)
}

// UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
func (o *PatchedEngagementRequest) UnsetStartTime() {
	o.StartTime.Unset()
}

// GetEndTime returns the EndTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedEngagementRequest) GetEndTime() time.Time {
	if o == nil || o.EndTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime.Get()
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedEngagementRequest) GetEndTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndTime.Get(), o.EndTime.IsSet()
}

// HasEndTime returns a boolean if a field has been set.
func (o *PatchedEngagementRequest) HasEndTime() bool {
	if o != nil && o.EndTime.IsSet() {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given NullableTime and assigns it to the EndTime field.
func (o *PatchedEngagementRequest) SetEndTime(v time.Time) {
	o.EndTime.Set(&v)
}
// SetEndTimeNil sets the value for EndTime to be an explicit nil
func (o *PatchedEngagementRequest) SetEndTimeNil() {
	o.EndTime.Set(nil)
}

// UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
func (o *PatchedEngagementRequest) UnsetEndTime() {
	o.EndTime.Unset()
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedEngagementRequest) GetAccount() string {
	if o == nil || o.Account.Get() == nil {
		var ret string
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedEngagementRequest) GetAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *PatchedEngagementRequest) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableString and assigns it to the Account field.
func (o *PatchedEngagementRequest) SetAccount(v string) {
	o.Account.Set(&v)
}
// SetAccountNil sets the value for Account to be an explicit nil
func (o *PatchedEngagementRequest) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *PatchedEngagementRequest) UnsetAccount() {
	o.Account.Unset()
}

// GetContacts returns the Contacts field value if set, zero value otherwise.
func (o *PatchedEngagementRequest) GetContacts() []string {
	if o == nil || o.Contacts == nil {
		var ret []string
		return ret
	}
	return *o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEngagementRequest) GetContactsOk() (*[]string, bool) {
	if o == nil || o.Contacts == nil {
		return nil, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *PatchedEngagementRequest) HasContacts() bool {
	if o != nil && o.Contacts != nil {
		return true
	}

	return false
}

// SetContacts gets a reference to the given []string and assigns it to the Contacts field.
func (o *PatchedEngagementRequest) SetContacts(v []string) {
	o.Contacts = &v
}

// GetIntegrationParams returns the IntegrationParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedEngagementRequest) GetIntegrationParams() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.IntegrationParams
}

// GetIntegrationParamsOk returns a tuple with the IntegrationParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedEngagementRequest) GetIntegrationParamsOk() (*map[string]interface{}, bool) {
	if o == nil || o.IntegrationParams == nil {
		return nil, false
	}
	return &o.IntegrationParams, true
}

// HasIntegrationParams returns a boolean if a field has been set.
func (o *PatchedEngagementRequest) HasIntegrationParams() bool {
	if o != nil && o.IntegrationParams != nil {
		return true
	}

	return false
}

// SetIntegrationParams gets a reference to the given map[string]interface{} and assigns it to the IntegrationParams field.
func (o *PatchedEngagementRequest) SetIntegrationParams(v map[string]interface{}) {
	o.IntegrationParams = v
}

// GetLinkedAccountParams returns the LinkedAccountParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedEngagementRequest) GetLinkedAccountParams() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.LinkedAccountParams
}

// GetLinkedAccountParamsOk returns a tuple with the LinkedAccountParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedEngagementRequest) GetLinkedAccountParamsOk() (*map[string]interface{}, bool) {
	if o == nil || o.LinkedAccountParams == nil {
		return nil, false
	}
	return &o.LinkedAccountParams, true
}

// HasLinkedAccountParams returns a boolean if a field has been set.
func (o *PatchedEngagementRequest) HasLinkedAccountParams() bool {
	if o != nil && o.LinkedAccountParams != nil {
		return true
	}

	return false
}

// SetLinkedAccountParams gets a reference to the given map[string]interface{} and assigns it to the LinkedAccountParams field.
func (o *PatchedEngagementRequest) SetLinkedAccountParams(v map[string]interface{}) {
	o.LinkedAccountParams = v
}

// GetRemoteFields returns the RemoteFields field value if set, zero value otherwise.
func (o *PatchedEngagementRequest) GetRemoteFields() []RemoteFieldRequest {
	if o == nil || o.RemoteFields == nil {
		var ret []RemoteFieldRequest
		return ret
	}
	return *o.RemoteFields
}

// GetRemoteFieldsOk returns a tuple with the RemoteFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEngagementRequest) GetRemoteFieldsOk() (*[]RemoteFieldRequest, bool) {
	if o == nil || o.RemoteFields == nil {
		return nil, false
	}
	return o.RemoteFields, true
}

// HasRemoteFields returns a boolean if a field has been set.
func (o *PatchedEngagementRequest) HasRemoteFields() bool {
	if o != nil && o.RemoteFields != nil {
		return true
	}

	return false
}

// SetRemoteFields gets a reference to the given []RemoteFieldRequest and assigns it to the RemoteFields field.
func (o *PatchedEngagementRequest) SetRemoteFields(v []RemoteFieldRequest) {
	o.RemoteFields = &v
}

func (o PatchedEngagementRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Owner.IsSet() {
		toSerialize["owner"] = o.Owner.Get()
	}
	if o.Content.IsSet() {
		toSerialize["content"] = o.Content.Get()
	}
	if o.Subject.IsSet() {
		toSerialize["subject"] = o.Subject.Get()
	}
	if o.Direction.IsSet() {
		toSerialize["direction"] = o.Direction.Get()
	}
	if o.EngagementType.IsSet() {
		toSerialize["engagement_type"] = o.EngagementType.Get()
	}
	if o.StartTime.IsSet() {
		toSerialize["start_time"] = o.StartTime.Get()
	}
	if o.EndTime.IsSet() {
		toSerialize["end_time"] = o.EndTime.Get()
	}
	if o.Account.IsSet() {
		toSerialize["account"] = o.Account.Get()
	}
	if o.Contacts != nil {
		toSerialize["contacts"] = o.Contacts
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

func (v *PatchedEngagementRequest) UnmarshalJSON(src []byte) error {
    type PatchedEngagementRequestUnmarshalTarget PatchedEngagementRequest

	var intermediateResult PatchedEngagementRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = PatchedEngagementRequest(intermediateResult)
	return nil
}
type NullablePatchedEngagementRequest struct {
	value *PatchedEngagementRequest
	isSet bool
}

func (v NullablePatchedEngagementRequest) Get() *PatchedEngagementRequest {
	return v.value
}

func (v *NullablePatchedEngagementRequest) Set(val *PatchedEngagementRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedEngagementRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedEngagementRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedEngagementRequest(val *PatchedEngagementRequest) *NullablePatchedEngagementRequest {
	return &NullablePatchedEngagementRequest{value: val, isSet: true}
}

func (v NullablePatchedEngagementRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedEngagementRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}



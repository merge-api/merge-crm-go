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

// Opportunity # The Opportunity Object ### Description The `Opportunity` object is used to represent a deal opportunity in a CRM system. ### Usage Example TODO
type Opportunity struct {
	// The opportunity's name.
	Name NullableString `json:"name,omitempty"`
	// The opportunity's description.
	Description NullableString `json:"description,omitempty"`
	// The opportunity's amount.
	Amount NullableInt32 `json:"amount,omitempty"`
	// The opportunity's owner.
	Owner NullableString `json:"owner,omitempty"`
	// The account of the opportunity.
	Account NullableString `json:"account,omitempty"`
	// The stage of the opportunity.
	Stage NullableString `json:"stage,omitempty"`
	// The opportunity's status.  * `OPEN` - OPEN * `WON` - WON * `LOST` - LOST
	Status NullableOpportunityStatusEnum `json:"status,omitempty"`
	// When the opportunity's last activity occurred.
	LastActivityAt NullableTime `json:"last_activity_at,omitempty"`
	// When the opportunity was closed.
	CloseDate NullableTime `json:"close_date,omitempty"`
	// When the third party's opportunity was created.
	RemoteCreatedAt NullableTime `json:"remote_created_at,omitempty"`
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

// NewOpportunity instantiates a new Opportunity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpportunity() *Opportunity {
	this := Opportunity{}
	return &this
}

// NewOpportunityWithDefaults instantiates a new Opportunity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpportunityWithDefaults() *Opportunity {
	this := Opportunity{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Opportunity) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Opportunity) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Opportunity) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Opportunity) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Opportunity) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Opportunity) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Opportunity) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Opportunity) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Opportunity) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Opportunity) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Opportunity) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Opportunity) UnsetDescription() {
	o.Description.Unset()
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Opportunity) GetAmount() int32 {
	if o == nil || o.Amount.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Opportunity) GetAmountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *Opportunity) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableInt32 and assigns it to the Amount field.
func (o *Opportunity) SetAmount(v int32) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *Opportunity) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *Opportunity) UnsetAmount() {
	o.Amount.Unset()
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Opportunity) GetOwner() string {
	if o == nil || o.Owner.Get() == nil {
		var ret string
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Opportunity) GetOwnerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *Opportunity) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableString and assigns it to the Owner field.
func (o *Opportunity) SetOwner(v string) {
	o.Owner.Set(&v)
}
// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *Opportunity) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *Opportunity) UnsetOwner() {
	o.Owner.Unset()
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Opportunity) GetAccount() string {
	if o == nil || o.Account.Get() == nil {
		var ret string
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Opportunity) GetAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *Opportunity) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableString and assigns it to the Account field.
func (o *Opportunity) SetAccount(v string) {
	o.Account.Set(&v)
}
// SetAccountNil sets the value for Account to be an explicit nil
func (o *Opportunity) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *Opportunity) UnsetAccount() {
	o.Account.Unset()
}

// GetStage returns the Stage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Opportunity) GetStage() string {
	if o == nil || o.Stage.Get() == nil {
		var ret string
		return ret
	}
	return *o.Stage.Get()
}

// GetStageOk returns a tuple with the Stage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Opportunity) GetStageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Stage.Get(), o.Stage.IsSet()
}

// HasStage returns a boolean if a field has been set.
func (o *Opportunity) HasStage() bool {
	if o != nil && o.Stage.IsSet() {
		return true
	}

	return false
}

// SetStage gets a reference to the given NullableString and assigns it to the Stage field.
func (o *Opportunity) SetStage(v string) {
	o.Stage.Set(&v)
}
// SetStageNil sets the value for Stage to be an explicit nil
func (o *Opportunity) SetStageNil() {
	o.Stage.Set(nil)
}

// UnsetStage ensures that no value is present for Stage, not even an explicit nil
func (o *Opportunity) UnsetStage() {
	o.Stage.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Opportunity) GetStatus() OpportunityStatusEnum {
	if o == nil || o.Status.Get() == nil {
		var ret OpportunityStatusEnum
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Opportunity) GetStatusOk() (*OpportunityStatusEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *Opportunity) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableOpportunityStatusEnum and assigns it to the Status field.
func (o *Opportunity) SetStatus(v OpportunityStatusEnum) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *Opportunity) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *Opportunity) UnsetStatus() {
	o.Status.Unset()
}

// GetLastActivityAt returns the LastActivityAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Opportunity) GetLastActivityAt() time.Time {
	if o == nil || o.LastActivityAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastActivityAt.Get()
}

// GetLastActivityAtOk returns a tuple with the LastActivityAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Opportunity) GetLastActivityAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastActivityAt.Get(), o.LastActivityAt.IsSet()
}

// HasLastActivityAt returns a boolean if a field has been set.
func (o *Opportunity) HasLastActivityAt() bool {
	if o != nil && o.LastActivityAt.IsSet() {
		return true
	}

	return false
}

// SetLastActivityAt gets a reference to the given NullableTime and assigns it to the LastActivityAt field.
func (o *Opportunity) SetLastActivityAt(v time.Time) {
	o.LastActivityAt.Set(&v)
}
// SetLastActivityAtNil sets the value for LastActivityAt to be an explicit nil
func (o *Opportunity) SetLastActivityAtNil() {
	o.LastActivityAt.Set(nil)
}

// UnsetLastActivityAt ensures that no value is present for LastActivityAt, not even an explicit nil
func (o *Opportunity) UnsetLastActivityAt() {
	o.LastActivityAt.Unset()
}

// GetCloseDate returns the CloseDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Opportunity) GetCloseDate() time.Time {
	if o == nil || o.CloseDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CloseDate.Get()
}

// GetCloseDateOk returns a tuple with the CloseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Opportunity) GetCloseDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CloseDate.Get(), o.CloseDate.IsSet()
}

// HasCloseDate returns a boolean if a field has been set.
func (o *Opportunity) HasCloseDate() bool {
	if o != nil && o.CloseDate.IsSet() {
		return true
	}

	return false
}

// SetCloseDate gets a reference to the given NullableTime and assigns it to the CloseDate field.
func (o *Opportunity) SetCloseDate(v time.Time) {
	o.CloseDate.Set(&v)
}
// SetCloseDateNil sets the value for CloseDate to be an explicit nil
func (o *Opportunity) SetCloseDateNil() {
	o.CloseDate.Set(nil)
}

// UnsetCloseDate ensures that no value is present for CloseDate, not even an explicit nil
func (o *Opportunity) UnsetCloseDate() {
	o.CloseDate.Unset()
}

// GetRemoteCreatedAt returns the RemoteCreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Opportunity) GetRemoteCreatedAt() time.Time {
	if o == nil || o.RemoteCreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.RemoteCreatedAt.Get()
}

// GetRemoteCreatedAtOk returns a tuple with the RemoteCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Opportunity) GetRemoteCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteCreatedAt.Get(), o.RemoteCreatedAt.IsSet()
}

// HasRemoteCreatedAt returns a boolean if a field has been set.
func (o *Opportunity) HasRemoteCreatedAt() bool {
	if o != nil && o.RemoteCreatedAt.IsSet() {
		return true
	}

	return false
}

// SetRemoteCreatedAt gets a reference to the given NullableTime and assigns it to the RemoteCreatedAt field.
func (o *Opportunity) SetRemoteCreatedAt(v time.Time) {
	o.RemoteCreatedAt.Set(&v)
}
// SetRemoteCreatedAtNil sets the value for RemoteCreatedAt to be an explicit nil
func (o *Opportunity) SetRemoteCreatedAtNil() {
	o.RemoteCreatedAt.Set(nil)
}

// UnsetRemoteCreatedAt ensures that no value is present for RemoteCreatedAt, not even an explicit nil
func (o *Opportunity) UnsetRemoteCreatedAt() {
	o.RemoteCreatedAt.Unset()
}

// GetRemoteWasDeleted returns the RemoteWasDeleted field value if set, zero value otherwise.
func (o *Opportunity) GetRemoteWasDeleted() bool {
	if o == nil || o.RemoteWasDeleted == nil {
		var ret bool
		return ret
	}
	return *o.RemoteWasDeleted
}

// GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Opportunity) GetRemoteWasDeletedOk() (*bool, bool) {
	if o == nil || o.RemoteWasDeleted == nil {
		return nil, false
	}
	return o.RemoteWasDeleted, true
}

// HasRemoteWasDeleted returns a boolean if a field has been set.
func (o *Opportunity) HasRemoteWasDeleted() bool {
	if o != nil && o.RemoteWasDeleted != nil {
		return true
	}

	return false
}

// SetRemoteWasDeleted gets a reference to the given bool and assigns it to the RemoteWasDeleted field.
func (o *Opportunity) SetRemoteWasDeleted(v bool) {
	o.RemoteWasDeleted = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Opportunity) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Opportunity) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Opportunity) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Opportunity) SetId(v string) {
	o.Id = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Opportunity) GetRemoteId() string {
	if o == nil || o.RemoteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemoteId.Get()
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Opportunity) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteId.Get(), o.RemoteId.IsSet()
}

// HasRemoteId returns a boolean if a field has been set.
func (o *Opportunity) HasRemoteId() bool {
	if o != nil && o.RemoteId.IsSet() {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given NullableString and assigns it to the RemoteId field.
func (o *Opportunity) SetRemoteId(v string) {
	o.RemoteId.Set(&v)
}
// SetRemoteIdNil sets the value for RemoteId to be an explicit nil
func (o *Opportunity) SetRemoteIdNil() {
	o.RemoteId.Set(nil)
}

// UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
func (o *Opportunity) UnsetRemoteId() {
	o.RemoteId.Unset()
}

// GetFieldMappings returns the FieldMappings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Opportunity) GetFieldMappings() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.FieldMappings
}

// GetFieldMappingsOk returns a tuple with the FieldMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Opportunity) GetFieldMappingsOk() (*map[string]interface{}, bool) {
	if o == nil || o.FieldMappings == nil {
		return nil, false
	}
	return &o.FieldMappings, true
}

// HasFieldMappings returns a boolean if a field has been set.
func (o *Opportunity) HasFieldMappings() bool {
	if o != nil && o.FieldMappings != nil {
		return true
	}

	return false
}

// SetFieldMappings gets a reference to the given map[string]interface{} and assigns it to the FieldMappings field.
func (o *Opportunity) SetFieldMappings(v map[string]interface{}) {
	o.FieldMappings = v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *Opportunity) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Opportunity) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *Opportunity) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *Opportunity) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetRemoteData returns the RemoteData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Opportunity) GetRemoteData() []RemoteData {
	if o == nil  {
		var ret []RemoteData
		return ret
	}
	return o.RemoteData
}

// GetRemoteDataOk returns a tuple with the RemoteData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Opportunity) GetRemoteDataOk() (*[]RemoteData, bool) {
	if o == nil || o.RemoteData == nil {
		return nil, false
	}
	return &o.RemoteData, true
}

// HasRemoteData returns a boolean if a field has been set.
func (o *Opportunity) HasRemoteData() bool {
	if o != nil && o.RemoteData != nil {
		return true
	}

	return false
}

// SetRemoteData gets a reference to the given []RemoteData and assigns it to the RemoteData field.
func (o *Opportunity) SetRemoteData(v []RemoteData) {
	o.RemoteData = v
}

// GetRemoteFields returns the RemoteFields field value if set, zero value otherwise.
func (o *Opportunity) GetRemoteFields() []RemoteField {
	if o == nil || o.RemoteFields == nil {
		var ret []RemoteField
		return ret
	}
	return *o.RemoteFields
}

// GetRemoteFieldsOk returns a tuple with the RemoteFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Opportunity) GetRemoteFieldsOk() (*[]RemoteField, bool) {
	if o == nil || o.RemoteFields == nil {
		return nil, false
	}
	return o.RemoteFields, true
}

// HasRemoteFields returns a boolean if a field has been set.
func (o *Opportunity) HasRemoteFields() bool {
	if o != nil && o.RemoteFields != nil {
		return true
	}

	return false
}

// SetRemoteFields gets a reference to the given []RemoteField and assigns it to the RemoteFields field.
func (o *Opportunity) SetRemoteFields(v []RemoteField) {
	o.RemoteFields = &v
}

func (o Opportunity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
	}
	if o.Owner.IsSet() {
		toSerialize["owner"] = o.Owner.Get()
	}
	if o.Account.IsSet() {
		toSerialize["account"] = o.Account.Get()
	}
	if o.Stage.IsSet() {
		toSerialize["stage"] = o.Stage.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.LastActivityAt.IsSet() {
		toSerialize["last_activity_at"] = o.LastActivityAt.Get()
	}
	if o.CloseDate.IsSet() {
		toSerialize["close_date"] = o.CloseDate.Get()
	}
	if o.RemoteCreatedAt.IsSet() {
		toSerialize["remote_created_at"] = o.RemoteCreatedAt.Get()
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

func (v *Opportunity) UnmarshalJSON(src []byte) error {
    type OpportunityUnmarshalTarget Opportunity

	var intermediateResult OpportunityUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = Opportunity(intermediateResult)
	return nil
}
type NullableOpportunity struct {
	value *Opportunity
	isSet bool
}

func (v NullableOpportunity) Get() *Opportunity {
	return v.value
}

func (v *NullableOpportunity) Set(val *Opportunity) {
	v.value = val
	v.isSet = true
}

func (v NullableOpportunity) IsSet() bool {
	return v.isSet
}

func (v *NullableOpportunity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpportunity(val *Opportunity) *NullableOpportunity {
	return &NullableOpportunity{value: val, isSet: true}
}

func (v NullableOpportunity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpportunity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}



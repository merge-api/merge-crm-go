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

// AssociationType # The AssociationType Object ### Description The `Association Type` object represents the relationship between two objects. ### Usage Example TODO
type AssociationType struct {
	SourceObjectClass *map[string]interface{} `json:"source_object_class,omitempty"`
	TargetObjectClasses *[]AssociationSubType `json:"target_object_classes,omitempty"`
	RemoteKeyName NullableString `json:"remote_key_name,omitempty"`
	DisplayName NullableString `json:"display_name,omitempty"`
	Cardinality NullableCardinalityEnum `json:"cardinality,omitempty"`
	IsRequired *bool `json:"is_required,omitempty"`
	Id *string `json:"id,omitempty"`
	// The third-party API ID of the matching object.
	RemoteId NullableString `json:"remote_id,omitempty"`
	// This is the datetime that this object was last updated by Merge
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewAssociationType instantiates a new AssociationType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssociationType() *AssociationType {
	this := AssociationType{}
	return &this
}

// NewAssociationTypeWithDefaults instantiates a new AssociationType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssociationTypeWithDefaults() *AssociationType {
	this := AssociationType{}
	return &this
}

// GetSourceObjectClass returns the SourceObjectClass field value if set, zero value otherwise.
func (o *AssociationType) GetSourceObjectClass() map[string]interface{} {
	if o == nil || o.SourceObjectClass == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.SourceObjectClass
}

// GetSourceObjectClassOk returns a tuple with the SourceObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociationType) GetSourceObjectClassOk() (*map[string]interface{}, bool) {
	if o == nil || o.SourceObjectClass == nil {
		return nil, false
	}
	return o.SourceObjectClass, true
}

// HasSourceObjectClass returns a boolean if a field has been set.
func (o *AssociationType) HasSourceObjectClass() bool {
	if o != nil && o.SourceObjectClass != nil {
		return true
	}

	return false
}

// SetSourceObjectClass gets a reference to the given map[string]interface{} and assigns it to the SourceObjectClass field.
func (o *AssociationType) SetSourceObjectClass(v map[string]interface{}) {
	o.SourceObjectClass = &v
}

// GetTargetObjectClasses returns the TargetObjectClasses field value if set, zero value otherwise.
func (o *AssociationType) GetTargetObjectClasses() []AssociationSubType {
	if o == nil || o.TargetObjectClasses == nil {
		var ret []AssociationSubType
		return ret
	}
	return *o.TargetObjectClasses
}

// GetTargetObjectClassesOk returns a tuple with the TargetObjectClasses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociationType) GetTargetObjectClassesOk() (*[]AssociationSubType, bool) {
	if o == nil || o.TargetObjectClasses == nil {
		return nil, false
	}
	return o.TargetObjectClasses, true
}

// HasTargetObjectClasses returns a boolean if a field has been set.
func (o *AssociationType) HasTargetObjectClasses() bool {
	if o != nil && o.TargetObjectClasses != nil {
		return true
	}

	return false
}

// SetTargetObjectClasses gets a reference to the given []AssociationSubType and assigns it to the TargetObjectClasses field.
func (o *AssociationType) SetTargetObjectClasses(v []AssociationSubType) {
	o.TargetObjectClasses = &v
}

// GetRemoteKeyName returns the RemoteKeyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssociationType) GetRemoteKeyName() string {
	if o == nil || o.RemoteKeyName.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemoteKeyName.Get()
}

// GetRemoteKeyNameOk returns a tuple with the RemoteKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssociationType) GetRemoteKeyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteKeyName.Get(), o.RemoteKeyName.IsSet()
}

// HasRemoteKeyName returns a boolean if a field has been set.
func (o *AssociationType) HasRemoteKeyName() bool {
	if o != nil && o.RemoteKeyName.IsSet() {
		return true
	}

	return false
}

// SetRemoteKeyName gets a reference to the given NullableString and assigns it to the RemoteKeyName field.
func (o *AssociationType) SetRemoteKeyName(v string) {
	o.RemoteKeyName.Set(&v)
}
// SetRemoteKeyNameNil sets the value for RemoteKeyName to be an explicit nil
func (o *AssociationType) SetRemoteKeyNameNil() {
	o.RemoteKeyName.Set(nil)
}

// UnsetRemoteKeyName ensures that no value is present for RemoteKeyName, not even an explicit nil
func (o *AssociationType) UnsetRemoteKeyName() {
	o.RemoteKeyName.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssociationType) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssociationType) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AssociationType) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *AssociationType) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *AssociationType) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *AssociationType) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetCardinality returns the Cardinality field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssociationType) GetCardinality() CardinalityEnum {
	if o == nil || o.Cardinality.Get() == nil {
		var ret CardinalityEnum
		return ret
	}
	return *o.Cardinality.Get()
}

// GetCardinalityOk returns a tuple with the Cardinality field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssociationType) GetCardinalityOk() (*CardinalityEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Cardinality.Get(), o.Cardinality.IsSet()
}

// HasCardinality returns a boolean if a field has been set.
func (o *AssociationType) HasCardinality() bool {
	if o != nil && o.Cardinality.IsSet() {
		return true
	}

	return false
}

// SetCardinality gets a reference to the given NullableCardinalityEnum and assigns it to the Cardinality field.
func (o *AssociationType) SetCardinality(v CardinalityEnum) {
	o.Cardinality.Set(&v)
}
// SetCardinalityNil sets the value for Cardinality to be an explicit nil
func (o *AssociationType) SetCardinalityNil() {
	o.Cardinality.Set(nil)
}

// UnsetCardinality ensures that no value is present for Cardinality, not even an explicit nil
func (o *AssociationType) UnsetCardinality() {
	o.Cardinality.Unset()
}

// GetIsRequired returns the IsRequired field value if set, zero value otherwise.
func (o *AssociationType) GetIsRequired() bool {
	if o == nil || o.IsRequired == nil {
		var ret bool
		return ret
	}
	return *o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociationType) GetIsRequiredOk() (*bool, bool) {
	if o == nil || o.IsRequired == nil {
		return nil, false
	}
	return o.IsRequired, true
}

// HasIsRequired returns a boolean if a field has been set.
func (o *AssociationType) HasIsRequired() bool {
	if o != nil && o.IsRequired != nil {
		return true
	}

	return false
}

// SetIsRequired gets a reference to the given bool and assigns it to the IsRequired field.
func (o *AssociationType) SetIsRequired(v bool) {
	o.IsRequired = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AssociationType) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociationType) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AssociationType) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AssociationType) SetId(v string) {
	o.Id = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssociationType) GetRemoteId() string {
	if o == nil || o.RemoteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemoteId.Get()
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssociationType) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteId.Get(), o.RemoteId.IsSet()
}

// HasRemoteId returns a boolean if a field has been set.
func (o *AssociationType) HasRemoteId() bool {
	if o != nil && o.RemoteId.IsSet() {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given NullableString and assigns it to the RemoteId field.
func (o *AssociationType) SetRemoteId(v string) {
	o.RemoteId.Set(&v)
}
// SetRemoteIdNil sets the value for RemoteId to be an explicit nil
func (o *AssociationType) SetRemoteIdNil() {
	o.RemoteId.Set(nil)
}

// UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
func (o *AssociationType) UnsetRemoteId() {
	o.RemoteId.Unset()
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *AssociationType) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociationType) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *AssociationType) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *AssociationType) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

func (o AssociationType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SourceObjectClass != nil {
		toSerialize["source_object_class"] = o.SourceObjectClass
	}
	if o.TargetObjectClasses != nil {
		toSerialize["target_object_classes"] = o.TargetObjectClasses
	}
	if o.RemoteKeyName.IsSet() {
		toSerialize["remote_key_name"] = o.RemoteKeyName.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["display_name"] = o.DisplayName.Get()
	}
	if o.Cardinality.IsSet() {
		toSerialize["cardinality"] = o.Cardinality.Get()
	}
	if o.IsRequired != nil {
		toSerialize["is_required"] = o.IsRequired
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.RemoteId.IsSet() {
		toSerialize["remote_id"] = o.RemoteId.Get()
	}
	if o.ModifiedAt != nil {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	return json.Marshal(toSerialize)
}

func (v *AssociationType) UnmarshalJSON(src []byte) error {
    type AssociationTypeUnmarshalTarget AssociationType

	var intermediateResult AssociationTypeUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = AssociationType(intermediateResult)
	return nil
}
type NullableAssociationType struct {
	value *AssociationType
	isSet bool
}

func (v NullableAssociationType) Get() *AssociationType {
	return v.value
}

func (v *NullableAssociationType) Set(val *AssociationType) {
	v.value = val
	v.isSet = true
}

func (v NullableAssociationType) IsSet() bool {
	return v.isSet
}

func (v *NullableAssociationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssociationType(val *AssociationType) *NullableAssociationType {
	return &NullableAssociationType{value: val, isSet: true}
}

func (v NullableAssociationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssociationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}



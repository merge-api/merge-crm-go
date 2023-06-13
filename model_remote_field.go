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
)

// RemoteField struct for RemoteField
type RemoteField struct {
	RemoteFieldClass RemoteFieldClass `json:"remote_field_class"`
	Value *map[string]interface{} `json:"value,omitempty"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewRemoteField instantiates a new RemoteField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteField(remoteFieldClass RemoteFieldClass) *RemoteField {
	this := RemoteField{}
	this.RemoteFieldClass = remoteFieldClass
	return &this
}

// NewRemoteFieldWithDefaults instantiates a new RemoteField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteFieldWithDefaults() *RemoteField {
	this := RemoteField{}
	return &this
}

// GetRemoteFieldClass returns the RemoteFieldClass field value
func (o *RemoteField) GetRemoteFieldClass() RemoteFieldClass {
	if o == nil {
		var ret RemoteFieldClass
		return ret
	}

	return o.RemoteFieldClass
}

// GetRemoteFieldClassOk returns a tuple with the RemoteFieldClass field value
// and a boolean to check if the value has been set.
func (o *RemoteField) GetRemoteFieldClassOk() (*RemoteFieldClass, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RemoteFieldClass, true
}

// SetRemoteFieldClass sets field value
func (o *RemoteField) SetRemoteFieldClass(v RemoteFieldClass) {
	o.RemoteFieldClass = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RemoteField) GetValue() map[string]interface{} {
	if o == nil || o.Value == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteField) GetValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RemoteField) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *RemoteField) SetValue(v map[string]interface{}) {
	o.Value = &v
}

func (o RemoteField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["remote_field_class"] = o.RemoteFieldClass
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

func (v *RemoteField) UnmarshalJSON(src []byte) error {
    type RemoteFieldUnmarshalTarget RemoteField

	var intermediateResult RemoteFieldUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = RemoteField(intermediateResult)
	return nil
}
type NullableRemoteField struct {
	value *RemoteField
	isSet bool
}

func (v NullableRemoteField) Get() *RemoteField {
	return v.value
}

func (v *NullableRemoteField) Set(val *RemoteField) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteField) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteField(val *RemoteField) *NullableRemoteField {
	return &NullableRemoteField{value: val, isSet: true}
}

func (v NullableRemoteField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


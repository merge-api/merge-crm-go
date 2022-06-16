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

// PhoneNumber # The PhoneNumber Object ### Description The `PhoneNumber` object is used to represent an entity's phone number. ### Usage Example Fetch from the `GET Contact` endpoint and view their phone numbers.
type PhoneNumber struct {
	// The phone number.
	PhoneNumber NullableString `json:"phone_number,omitempty"`
	// The phone number's type.
	PhoneNumberType NullableString `json:"phone_number_type,omitempty"`
    // raw json response by property name
    ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewPhoneNumber instantiates a new PhoneNumber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhoneNumber() *PhoneNumber {
	this := PhoneNumber{}
	return &this
}

// NewPhoneNumberWithDefaults instantiates a new PhoneNumber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhoneNumberWithDefaults() *PhoneNumber {
	this := PhoneNumber{}
	return &this
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhoneNumber) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber.Get()
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhoneNumber) GetPhoneNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PhoneNumber.Get(), o.PhoneNumber.IsSet()
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *PhoneNumber) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber.IsSet() {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given NullableString and assigns it to the PhoneNumber field.
func (o *PhoneNumber) SetPhoneNumber(v string) {
	o.PhoneNumber.Set(&v)
}
// SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil
func (o *PhoneNumber) SetPhoneNumberNil() {
	o.PhoneNumber.Set(nil)
}

// UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
func (o *PhoneNumber) UnsetPhoneNumber() {
	o.PhoneNumber.Unset()
}

// GetPhoneNumberType returns the PhoneNumberType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhoneNumber) GetPhoneNumberType() string {
	if o == nil || o.PhoneNumberType.Get() == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumberType.Get()
}

// GetPhoneNumberTypeOk returns a tuple with the PhoneNumberType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhoneNumber) GetPhoneNumberTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PhoneNumberType.Get(), o.PhoneNumberType.IsSet()
}

// HasPhoneNumberType returns a boolean if a field has been set.
func (o *PhoneNumber) HasPhoneNumberType() bool {
	if o != nil && o.PhoneNumberType.IsSet() {
		return true
	}

	return false
}

// SetPhoneNumberType gets a reference to the given NullableString and assigns it to the PhoneNumberType field.
func (o *PhoneNumber) SetPhoneNumberType(v string) {
	o.PhoneNumberType.Set(&v)
}
// SetPhoneNumberTypeNil sets the value for PhoneNumberType to be an explicit nil
func (o *PhoneNumber) SetPhoneNumberTypeNil() {
	o.PhoneNumberType.Set(nil)
}

// UnsetPhoneNumberType ensures that no value is present for PhoneNumberType, not even an explicit nil
func (o *PhoneNumber) UnsetPhoneNumberType() {
	o.PhoneNumberType.Unset()
}

func (o PhoneNumber) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PhoneNumber.IsSet() {
		toSerialize["phone_number"] = o.PhoneNumber.Get()
	}
	if o.PhoneNumberType.IsSet() {
		toSerialize["phone_number_type"] = o.PhoneNumberType.Get()
	}
	return json.Marshal(toSerialize)
}

func (v *PhoneNumber) UnmarshalJSON(src []byte) error {
    type PhoneNumberUnmarshalTarget PhoneNumber

	var intermediateResult PhoneNumberUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = PhoneNumber(intermediateResult)
	return nil
}
type NullablePhoneNumber struct {
	value *PhoneNumber
	isSet bool
}

func (v NullablePhoneNumber) Get() *PhoneNumber {
	return v.value
}

func (v *NullablePhoneNumber) Set(val *PhoneNumber) {
	v.value = val
	v.isSet = true
}

func (v NullablePhoneNumber) IsSet() bool {
	return v.isSet
}

func (v *NullablePhoneNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhoneNumber(val *PhoneNumber) *NullablePhoneNumber {
	return &NullablePhoneNumber{value: val, isSet: true}
}

func (v NullablePhoneNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhoneNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}



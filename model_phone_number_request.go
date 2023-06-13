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

// PhoneNumberRequest # The PhoneNumber Object ### Description The `PhoneNumber` object is used to represent an entity's phone number. ### Usage Example Fetch from the `GET Contact` endpoint and view their phone numbers.
type PhoneNumberRequest struct {
	// The phone number.
	PhoneNumber NullableString `json:"phone_number,omitempty"`
	// The phone number's type.
	PhoneNumberType NullableString `json:"phone_number_type,omitempty"`
	IntegrationParams map[string]interface{} `json:"integration_params,omitempty"`
	LinkedAccountParams map[string]interface{} `json:"linked_account_params,omitempty"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewPhoneNumberRequest instantiates a new PhoneNumberRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhoneNumberRequest() *PhoneNumberRequest {
	this := PhoneNumberRequest{}
	return &this
}

// NewPhoneNumberRequestWithDefaults instantiates a new PhoneNumberRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhoneNumberRequestWithDefaults() *PhoneNumberRequest {
	this := PhoneNumberRequest{}
	return &this
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhoneNumberRequest) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber.Get()
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhoneNumberRequest) GetPhoneNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PhoneNumber.Get(), o.PhoneNumber.IsSet()
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *PhoneNumberRequest) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber.IsSet() {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given NullableString and assigns it to the PhoneNumber field.
func (o *PhoneNumberRequest) SetPhoneNumber(v string) {
	o.PhoneNumber.Set(&v)
}
// SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil
func (o *PhoneNumberRequest) SetPhoneNumberNil() {
	o.PhoneNumber.Set(nil)
}

// UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
func (o *PhoneNumberRequest) UnsetPhoneNumber() {
	o.PhoneNumber.Unset()
}

// GetPhoneNumberType returns the PhoneNumberType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhoneNumberRequest) GetPhoneNumberType() string {
	if o == nil || o.PhoneNumberType.Get() == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumberType.Get()
}

// GetPhoneNumberTypeOk returns a tuple with the PhoneNumberType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhoneNumberRequest) GetPhoneNumberTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PhoneNumberType.Get(), o.PhoneNumberType.IsSet()
}

// HasPhoneNumberType returns a boolean if a field has been set.
func (o *PhoneNumberRequest) HasPhoneNumberType() bool {
	if o != nil && o.PhoneNumberType.IsSet() {
		return true
	}

	return false
}

// SetPhoneNumberType gets a reference to the given NullableString and assigns it to the PhoneNumberType field.
func (o *PhoneNumberRequest) SetPhoneNumberType(v string) {
	o.PhoneNumberType.Set(&v)
}
// SetPhoneNumberTypeNil sets the value for PhoneNumberType to be an explicit nil
func (o *PhoneNumberRequest) SetPhoneNumberTypeNil() {
	o.PhoneNumberType.Set(nil)
}

// UnsetPhoneNumberType ensures that no value is present for PhoneNumberType, not even an explicit nil
func (o *PhoneNumberRequest) UnsetPhoneNumberType() {
	o.PhoneNumberType.Unset()
}

// GetIntegrationParams returns the IntegrationParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhoneNumberRequest) GetIntegrationParams() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.IntegrationParams
}

// GetIntegrationParamsOk returns a tuple with the IntegrationParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhoneNumberRequest) GetIntegrationParamsOk() (*map[string]interface{}, bool) {
	if o == nil || o.IntegrationParams == nil {
		return nil, false
	}
	return &o.IntegrationParams, true
}

// HasIntegrationParams returns a boolean if a field has been set.
func (o *PhoneNumberRequest) HasIntegrationParams() bool {
	if o != nil && o.IntegrationParams != nil {
		return true
	}

	return false
}

// SetIntegrationParams gets a reference to the given map[string]interface{} and assigns it to the IntegrationParams field.
func (o *PhoneNumberRequest) SetIntegrationParams(v map[string]interface{}) {
	o.IntegrationParams = v
}

// GetLinkedAccountParams returns the LinkedAccountParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhoneNumberRequest) GetLinkedAccountParams() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.LinkedAccountParams
}

// GetLinkedAccountParamsOk returns a tuple with the LinkedAccountParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhoneNumberRequest) GetLinkedAccountParamsOk() (*map[string]interface{}, bool) {
	if o == nil || o.LinkedAccountParams == nil {
		return nil, false
	}
	return &o.LinkedAccountParams, true
}

// HasLinkedAccountParams returns a boolean if a field has been set.
func (o *PhoneNumberRequest) HasLinkedAccountParams() bool {
	if o != nil && o.LinkedAccountParams != nil {
		return true
	}

	return false
}

// SetLinkedAccountParams gets a reference to the given map[string]interface{} and assigns it to the LinkedAccountParams field.
func (o *PhoneNumberRequest) SetLinkedAccountParams(v map[string]interface{}) {
	o.LinkedAccountParams = v
}

func (o PhoneNumberRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PhoneNumber.IsSet() {
		toSerialize["phone_number"] = o.PhoneNumber.Get()
	}
	if o.PhoneNumberType.IsSet() {
		toSerialize["phone_number_type"] = o.PhoneNumberType.Get()
	}
	if o.IntegrationParams != nil {
		toSerialize["integration_params"] = o.IntegrationParams
	}
	if o.LinkedAccountParams != nil {
		toSerialize["linked_account_params"] = o.LinkedAccountParams
	}
	return json.Marshal(toSerialize)
}

func (v *PhoneNumberRequest) UnmarshalJSON(src []byte) error {
    type PhoneNumberRequestUnmarshalTarget PhoneNumberRequest

	var intermediateResult PhoneNumberRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = PhoneNumberRequest(intermediateResult)
	return nil
}
type NullablePhoneNumberRequest struct {
	value *PhoneNumberRequest
	isSet bool
}

func (v NullablePhoneNumberRequest) Get() *PhoneNumberRequest {
	return v.value
}

func (v *NullablePhoneNumberRequest) Set(val *PhoneNumberRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePhoneNumberRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePhoneNumberRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhoneNumberRequest(val *PhoneNumberRequest) *NullablePhoneNumberRequest {
	return &NullablePhoneNumberRequest{value: val, isSet: true}
}

func (v NullablePhoneNumberRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhoneNumberRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}



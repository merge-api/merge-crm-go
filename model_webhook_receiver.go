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

// WebhookReceiver struct for WebhookReceiver
type WebhookReceiver struct {
	Event string `json:"event"`
	IsActive bool `json:"is_active"`
	Key *string `json:"key,omitempty"`
    // raw json response by property name
    ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewWebhookReceiver instantiates a new WebhookReceiver object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookReceiver(event string, isActive bool) *WebhookReceiver {
	this := WebhookReceiver{}
	this.Event = event
	this.IsActive = isActive
	return &this
}

// NewWebhookReceiverWithDefaults instantiates a new WebhookReceiver object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookReceiverWithDefaults() *WebhookReceiver {
	this := WebhookReceiver{}
	return &this
}

// GetEvent returns the Event field value
func (o *WebhookReceiver) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *WebhookReceiver) GetEventOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *WebhookReceiver) SetEvent(v string) {
	o.Event = v
}

// GetIsActive returns the IsActive field value
func (o *WebhookReceiver) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *WebhookReceiver) GetIsActiveOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *WebhookReceiver) SetIsActive(v bool) {
	o.IsActive = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *WebhookReceiver) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookReceiver) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *WebhookReceiver) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *WebhookReceiver) SetKey(v string) {
	o.Key = &v
}

func (o WebhookReceiver) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["event"] = o.Event
	}
	if true {
		toSerialize["is_active"] = o.IsActive
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

func (v *WebhookReceiver) UnmarshalJSON(src []byte) error {
    type WebhookReceiverUnmarshalTarget WebhookReceiver

	var intermediateResult WebhookReceiverUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = WebhookReceiver(intermediateResult)
	return nil
}
type NullableWebhookReceiver struct {
	value *WebhookReceiver
	isSet bool
}

func (v NullableWebhookReceiver) Get() *WebhookReceiver {
	return v.value
}

func (v *NullableWebhookReceiver) Set(val *WebhookReceiver) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookReceiver) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookReceiver) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookReceiver(val *WebhookReceiver) *NullableWebhookReceiver {
	return &NullableWebhookReceiver{value: val, isSet: true}
}

func (v NullableWebhookReceiver) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookReceiver) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}



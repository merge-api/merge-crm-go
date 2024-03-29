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

// EngagementEndpointRequest struct for EngagementEndpointRequest
type EngagementEndpointRequest struct {
	Model EngagementRequest `json:"model"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewEngagementEndpointRequest instantiates a new EngagementEndpointRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEngagementEndpointRequest(model EngagementRequest) *EngagementEndpointRequest {
	this := EngagementEndpointRequest{}
	this.Model = model
	return &this
}

// NewEngagementEndpointRequestWithDefaults instantiates a new EngagementEndpointRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEngagementEndpointRequestWithDefaults() *EngagementEndpointRequest {
	this := EngagementEndpointRequest{}
	return &this
}

// GetModel returns the Model field value
func (o *EngagementEndpointRequest) GetModel() EngagementRequest {
	if o == nil {
		var ret EngagementRequest
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *EngagementEndpointRequest) GetModelOk() (*EngagementRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *EngagementEndpointRequest) SetModel(v EngagementRequest) {
	o.Model = v
}

func (o EngagementEndpointRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["model"] = o.Model
	}
	return json.Marshal(toSerialize)
}

func (v *EngagementEndpointRequest) UnmarshalJSON(src []byte) error {
    type EngagementEndpointRequestUnmarshalTarget EngagementEndpointRequest

	var intermediateResult EngagementEndpointRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = EngagementEndpointRequest(intermediateResult)
	return nil
}
type NullableEngagementEndpointRequest struct {
	value *EngagementEndpointRequest
	isSet bool
}

func (v NullableEngagementEndpointRequest) Get() *EngagementEndpointRequest {
	return v.value
}

func (v *NullableEngagementEndpointRequest) Set(val *EngagementEndpointRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEngagementEndpointRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEngagementEndpointRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEngagementEndpointRequest(val *EngagementEndpointRequest) *NullableEngagementEndpointRequest {
	return &NullableEngagementEndpointRequest{value: val, isSet: true}
}

func (v NullableEngagementEndpointRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEngagementEndpointRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}



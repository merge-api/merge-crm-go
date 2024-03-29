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

// PatchedTaskEndpointRequest struct for PatchedTaskEndpointRequest
type PatchedTaskEndpointRequest struct {
	Model PatchedTaskRequest `json:"model"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewPatchedTaskEndpointRequest instantiates a new PatchedTaskEndpointRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedTaskEndpointRequest(model PatchedTaskRequest) *PatchedTaskEndpointRequest {
	this := PatchedTaskEndpointRequest{}
	this.Model = model
	return &this
}

// NewPatchedTaskEndpointRequestWithDefaults instantiates a new PatchedTaskEndpointRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedTaskEndpointRequestWithDefaults() *PatchedTaskEndpointRequest {
	this := PatchedTaskEndpointRequest{}
	return &this
}

// GetModel returns the Model field value
func (o *PatchedTaskEndpointRequest) GetModel() PatchedTaskRequest {
	if o == nil {
		var ret PatchedTaskRequest
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *PatchedTaskEndpointRequest) GetModelOk() (*PatchedTaskRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *PatchedTaskEndpointRequest) SetModel(v PatchedTaskRequest) {
	o.Model = v
}

func (o PatchedTaskEndpointRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["model"] = o.Model
	}
	return json.Marshal(toSerialize)
}

func (v *PatchedTaskEndpointRequest) UnmarshalJSON(src []byte) error {
    type PatchedTaskEndpointRequestUnmarshalTarget PatchedTaskEndpointRequest

	var intermediateResult PatchedTaskEndpointRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = PatchedTaskEndpointRequest(intermediateResult)
	return nil
}
type NullablePatchedTaskEndpointRequest struct {
	value *PatchedTaskEndpointRequest
	isSet bool
}

func (v NullablePatchedTaskEndpointRequest) Get() *PatchedTaskEndpointRequest {
	return v.value
}

func (v *NullablePatchedTaskEndpointRequest) Set(val *PatchedTaskEndpointRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedTaskEndpointRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedTaskEndpointRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedTaskEndpointRequest(val *PatchedTaskEndpointRequest) *NullablePatchedTaskEndpointRequest {
	return &NullablePatchedTaskEndpointRequest{value: val, isSet: true}
}

func (v NullablePatchedTaskEndpointRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedTaskEndpointRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


